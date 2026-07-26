package Data_Array

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Control_Monad_ST_Internal "gopurs/output/Control.Monad.ST.Internal"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_FunctorWithIndex "gopurs/output/Data.FunctorWithIndex"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Data_Array_ST "gopurs/output/Data.Array.ST"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Array_ST_Iterator "gopurs/output/Data.Array.ST.Iterator"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	unsafe "unsafe"
)

var cache_traverse_ gopurs_runtime.Value
var once_traverse_ sync.Once
func Get_traverse_() gopurs_runtime.Value {
	once_traverse_.Do(func() {
		cache_traverse_ = gopurs_runtime.Apply(pkg_Data_Foldable.Get_traverse_(), pkg_Control_Monad_ST_Internal.Get_applicativeST())
	})
	return cache_traverse_
}

var cache_lessThan gopurs_runtime.Value
var once_lessThan sync.Once
func Get_lessThan() gopurs_runtime.Value {
	once_lessThan.Do(func() {
		cache_lessThan = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1527465420))
})
}()
	})
	return cache_lessThan
}

var cache_lessThanOrEq gopurs_runtime.Value
var once_lessThanOrEq sync.Once
func Get_lessThanOrEq() gopurs_runtime.Value {
	once_lessThanOrEq.Do(func() {
		cache_lessThanOrEq = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool(((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 380165415)) != (true))
})
}()
	})
	return cache_lessThanOrEq
}

var cache_discard gopurs_runtime.Value
var once_discard sync.Once
func Get_discard() gopurs_runtime.Value {
	once_discard.Do(func() {
		cache_discard = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), pkg_Control_Monad_ST_Internal.Get_bindST())
	})
	return cache_discard
}

var cache_void gopurs_runtime.Value
var once_void sync.Once
func Get_void() gopurs_runtime.Value {
	once_void.Do(func() {
		cache_void = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
	})
	return cache_void
}

var cache_intercalate1 gopurs_runtime.Value
var once_intercalate1 sync.Once
func Get_intercalate1() gopurs_runtime.Value {
	once_intercalate1.Do(func() {
		cache_intercalate1 = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intercalate1(dictMonoid_0_box)
})
	})
	return cache_intercalate1
}

var cache_greaterThanOrEq gopurs_runtime.Value
var once_greaterThanOrEq sync.Once
func Get_greaterThanOrEq() gopurs_runtime.Value {
	once_greaterThanOrEq.Do(func() {
		cache_greaterThanOrEq = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool(((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1527465420)) != (true))
})
}()
	})
	return cache_greaterThanOrEq
}

var cache_zipWith gopurs_runtime.Value
var once_zipWith sync.Once
func Get_zipWith() gopurs_runtime.Value {
	once_zipWith.Do(func() {
		cache_zipWith = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_zipWith(__local_var_0_box, func() []gopurs_runtime.Value {
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
	return cache_zipWith
}

var cache_zipWithA gopurs_runtime.Value
var once_zipWithA sync.Once
func Get_zipWithA() gopurs_runtime.Value {
	once_zipWithA.Do(func() {
		cache_zipWithA = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWithA(dictApplicative_0_box)
})
	})
	return cache_zipWithA
}

var cache_zip gopurs_runtime.Value
var once_zip sync.Once
func Get_zip() gopurs_runtime.Value {
	once_zip.Do(func() {
		cache_zip = gopurs_runtime.Apply(Get_zipWith(), pkg_Data_Tuple.Get_Tuple())
	})
	return cache_zip
}

var cache_updateAtIndices gopurs_runtime.Value
var once_updateAtIndices sync.Once
func Get_updateAtIndices() gopurs_runtime.Value {
	once_updateAtIndices.Do(func() {
		cache_updateAtIndices = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_updateAtIndices(dictFoldable_0_box)
})
	})
	return cache_updateAtIndices
}

var cache_updateAt gopurs_runtime.Value
var once_updateAt sync.Once
func Get_updateAt() gopurs_runtime.Value {
	once_updateAt.Do(func() {
		cache_updateAt = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_updateAt(__local_var_0_box.IntVal, __local_var_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_updateAt
}

var cache_unsafeIndex gopurs_runtime.Value
var once_unsafeIndex sync.Once
func Get_unsafeIndex() gopurs_runtime.Value {
	once_unsafeIndex.Do(func() {
		cache_unsafeIndex = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeIndex(_dollar__unused_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), __local_var_2_box.IntVal)
})
	})
	return cache_unsafeIndex
}

var cache_uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		cache_uncons = gopurs_runtime.Func(func(__local_var_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_uncons
}

var cache_toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		cache_toUnfoldable = gopurs_runtime.Func2(func(dictUnfoldable_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toUnfoldable(dictUnfoldable_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_toUnfoldable
}

var cache_tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		cache_tail = gopurs_runtime.Func(func(__local_var_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_tail(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_tail
}

var cache_sortBy gopurs_runtime.Value
var once_sortBy sync.Once
func Get_sortBy() gopurs_runtime.Value {
	once_sortBy.Do(func() {
		cache_sortBy = gopurs_runtime.Func2(func(comp_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_sortBy(comp_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_sortBy
}

var cache_sortWith gopurs_runtime.Value
var once_sortWith sync.Once
func Get_sortWith() gopurs_runtime.Value {
	once_sortWith.Do(func() {
		cache_sortWith = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sortWith(dictOrd_0_box, f_1_box)
})
	})
	return cache_sortWith
}

var cache_sort gopurs_runtime.Value
var once_sort sync.Once
func Get_sort() gopurs_runtime.Value {
	once_sort.Do(func() {
		cache_sort = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sort(dictOrd_0_box)
})
	})
	return cache_sort
}

var cache_snoc gopurs_runtime.Value
var once_snoc sync.Once
func Get_snoc() gopurs_runtime.Value {
	once_snoc.Do(func() {
		cache_snoc = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_snoc(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), x_1_box))
})
	})
	return cache_snoc
}

var cache_slice gopurs_runtime.Value
var once_slice sync.Once
func Get_slice() gopurs_runtime.Value {
	once_slice.Do(func() {
		cache_slice = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_slice(__local_var_0_box.IntVal, __local_var_1_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_slice
}

var cache_splitAt gopurs_runtime.Value
var once_splitAt sync.Once
func Get_splitAt() gopurs_runtime.Value {
	once_splitAt.Do(func() {
		cache_splitAt = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_splitAt(v_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v1_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_splitAt
}

var cache_take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		cache_take = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_take(n_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_take
}

var cache_singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		cache_singleton = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_singleton(a_0_box))
})
	})
	return cache_singleton
}

var cache_scanr gopurs_runtime.Value
var once_scanr sync.Once
func Get_scanr() gopurs_runtime.Value {
	once_scanr.Do(func() {
		cache_scanr = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_scanr(__local_var_0_box, __local_var_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_scanr
}

var cache_scanl gopurs_runtime.Value
var once_scanl sync.Once
func Get_scanl() gopurs_runtime.Value {
	once_scanl.Do(func() {
		cache_scanl = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_scanl(__local_var_0_box, __local_var_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_scanl
}

var cache_replicate gopurs_runtime.Value
var once_replicate sync.Once
func Get_replicate() gopurs_runtime.Value {
	once_replicate.Do(func() {
		cache_replicate = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_replicate(__local_var_0_box.IntVal, __local_var_1_box))
})
	})
	return cache_replicate
}

var cache_range_ gopurs_runtime.Value
var once_range_ sync.Once
func Get_range_() gopurs_runtime.Value {
	once_range_.Do(func() {
		cache_range_ = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_range_(__local_var_0_box.IntVal, __local_var_1_box.IntVal)
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_range_
}

var cache_partition gopurs_runtime.Value
var once_partition sync.Once
func Get_partition() gopurs_runtime.Value {
	once_partition.Do(func() {
		cache_partition = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_partition(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_partition
}

var cache_null gopurs_runtime.Value
var once_null sync.Once
func Get_null() gopurs_runtime.Value {
	once_null.Do(func() {
		cache_null = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_null(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_null
}

var cache_modifyAtIndices gopurs_runtime.Value
var once_modifyAtIndices sync.Once
func Get_modifyAtIndices() gopurs_runtime.Value {
	once_modifyAtIndices.Do(func() {
		cache_modifyAtIndices = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modifyAtIndices(dictFoldable_0_box)
})
	})
	return cache_modifyAtIndices
}

var cache_mapWithIndex gopurs_runtime.Value
var once_mapWithIndex sync.Once
func Get_mapWithIndex() gopurs_runtime.Value {
	once_mapWithIndex.Do(func() {
		cache_mapWithIndex = gopurs_runtime.RecordGet(pkg_Data_FunctorWithIndex.Get_functorWithIndexArray(), "mapWithIndex")
	})
	return cache_mapWithIndex
}

var cache_intersperse gopurs_runtime.Value
var once_intersperse sync.Once
func Get_intersperse() gopurs_runtime.Value {
	once_intersperse.Do(func() {
		cache_intersperse = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, arr_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_intersperse(a_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(arr_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_intersperse
}

var cache_intercalate gopurs_runtime.Value
var once_intercalate sync.Once
func Get_intercalate() gopurs_runtime.Value {
	once_intercalate.Do(func() {
		cache_intercalate = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intercalate(dictMonoid_0_box)
})
	})
	return cache_intercalate
}

var cache_insertAt gopurs_runtime.Value
var once_insertAt sync.Once
func Get_insertAt() gopurs_runtime.Value {
	once_insertAt.Do(func() {
		cache_insertAt = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_insertAt(__local_var_0_box.IntVal, __local_var_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_insertAt
}

var cache_init_ gopurs_runtime.Value
var once_init_ sync.Once
func Get_init_() gopurs_runtime.Value {
	once_init_.Do(func() {
		cache_init_ = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_init_(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_init_
}

var cache_index gopurs_runtime.Value
var once_index sync.Once
func Get_index() gopurs_runtime.Value {
	once_index.Do(func() {
		cache_index = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_index(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), __local_var_1_box.IntVal))}
})
	})
	return cache_index
}

var cache_last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		cache_last = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_last(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_last
}

var cache_unsnoc gopurs_runtime.Value
var once_unsnoc sync.Once
func Get_unsnoc() gopurs_runtime.Value {
	once_unsnoc.Do(func() {
		cache_unsnoc = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_unsnoc(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_unsnoc
}

var cache_modifyAt gopurs_runtime.Value
var once_modifyAt sync.Once
func Get_modifyAt() gopurs_runtime.Value {
	once_modifyAt.Do(func() {
		cache_modifyAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_modifyAt(i_0_box.IntVal, f_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_modifyAt
}

var cache_span gopurs_runtime.Value
var once_span sync.Once
func Get_span() gopurs_runtime.Value {
	once_span.Do(func() {
		cache_span = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, arr_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_span(p_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(arr_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_span
}

var cache_takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		cache_takeWhile = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_takeWhile(p_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_takeWhile
}

var cache_unzip gopurs_runtime.Value
var once_unzip sync.Once
func Get_unzip() gopurs_runtime.Value {
	once_unzip.Do(func() {
		cache_unzip = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unzip(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_unzip
}

var cache_head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		cache_head = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_head(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_head
}

var cache_nubBy gopurs_runtime.Value
var once_nubBy sync.Once
func Get_nubBy() gopurs_runtime.Value {
	once_nubBy.Do(func() {
		cache_nubBy = gopurs_runtime.Func2(func(comp_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_nubBy(comp_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_nubBy
}

var cache_nub gopurs_runtime.Value
var once_nub sync.Once
func Get_nub() gopurs_runtime.Value {
	once_nub.Do(func() {
		cache_nub = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_nub(dictOrd_0_box)
})
	})
	return cache_nub
}

var cache_groupBy gopurs_runtime.Value
var once_groupBy sync.Once
func Get_groupBy() gopurs_runtime.Value {
	once_groupBy.Do(func() {
		cache_groupBy = gopurs_runtime.Func2(func(op_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_groupBy(op_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_groupBy
}

var cache_groupAllBy gopurs_runtime.Value
var once_groupAllBy sync.Once
func Get_groupAllBy() gopurs_runtime.Value {
	once_groupAllBy.Do(func() {
		cache_groupAllBy = gopurs_runtime.Func(func(cmp_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_groupAllBy(cmp_0_box)
})
	})
	return cache_groupAllBy
}

var cache_groupAll gopurs_runtime.Value
var once_groupAll sync.Once
func Get_groupAll() gopurs_runtime.Value {
	once_groupAll.Do(func() {
		cache_groupAll = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_groupAll(dictOrd_0_box)
})
	})
	return cache_groupAll
}

var cache_group gopurs_runtime.Value
var once_group sync.Once
func Get_group() gopurs_runtime.Value {
	once_group.Do(func() {
		cache_group = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_group(dictEq_0_box)
})
	})
	return cache_group
}

var cache_fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		cache_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldable(dictFoldable_0_box)
})
	})
	return cache_fromFoldable
}

var cache_foldr gopurs_runtime.Value
var once_foldr sync.Once
func Get_foldr() gopurs_runtime.Value {
	once_foldr.Do(func() {
		cache_foldr = gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldr")
	})
	return cache_foldr
}

var cache_foldl gopurs_runtime.Value
var once_foldl sync.Once
func Get_foldl() gopurs_runtime.Value {
	once_foldl.Do(func() {
		cache_foldl = gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldl")
	})
	return cache_foldl
}

var cache_transpose gopurs_runtime.Value
var once_transpose sync.Once
func Get_transpose() gopurs_runtime.Value {
	once_transpose.Do(func() {
		cache_transpose = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_transpose(func() [][]gopurs_runtime.Value {
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
	return cache_transpose
}

var cache_foldRecM gopurs_runtime.Value
var once_foldRecM sync.Once
func Get_foldRecM() gopurs_runtime.Value {
	once_foldRecM.Do(func() {
		cache_foldRecM = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldRecM(dictMonadRec_0_box)
})
	})
	return cache_foldRecM
}

var cache_foldMap gopurs_runtime.Value
var once_foldMap sync.Once
func Get_foldMap() gopurs_runtime.Value {
	once_foldMap.Do(func() {
		cache_foldMap = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap(dictMonoid_0_box)
})
	})
	return cache_foldMap
}

var cache_foldM gopurs_runtime.Value
var once_foldM sync.Once
func Get_foldM() gopurs_runtime.Value {
	once_foldM.Do(func() {
		cache_foldM = gopurs_runtime.Func4(func(dictMonad_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value, __local_var_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldM(dictMonad_0_box, f_1_box, b_2_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_3_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_foldM
}

var cache_fold gopurs_runtime.Value
var once_fold sync.Once
func Get_fold() gopurs_runtime.Value {
	once_fold.Do(func() {
		cache_fold = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fold(dictMonoid_0_box)
})
	})
	return cache_fold
}

var cache_findMap gopurs_runtime.Value
var once_findMap sync.Once
func Get_findMap() gopurs_runtime.Value {
	once_findMap.Do(func() {
		cache_findMap = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findMap(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_findMap
}

var cache_findLastIndex gopurs_runtime.Value
var once_findLastIndex sync.Once
func Get_findLastIndex() gopurs_runtime.Value {
	once_findLastIndex.Do(func() {
		cache_findLastIndex = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findLastIndex(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_findLastIndex
}

var cache_insertBy gopurs_runtime.Value
var once_insertBy sync.Once
func Get_insertBy() gopurs_runtime.Value {
	once_insertBy.Do(func() {
		cache_insertBy = gopurs_runtime.Func3(func(cmp_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_insertBy(cmp_0_box, x_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(ys_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_insertBy
}

var cache_insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		cache_insert = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insert(dictOrd_0_box)
})
	})
	return cache_insert
}

var cache_findIndex gopurs_runtime.Value
var once_findIndex sync.Once
func Get_findIndex() gopurs_runtime.Value {
	once_findIndex.Do(func() {
		cache_findIndex = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findIndex(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_findIndex
}

var cache_find gopurs_runtime.Value
var once_find sync.Once
func Get_find() gopurs_runtime.Value {
	once_find.Do(func() {
		cache_find = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_find(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_find
}

var cache_filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		cache_filter = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_filter(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_filter
}

var cache_intersectBy gopurs_runtime.Value
var once_intersectBy sync.Once
func Get_intersectBy() gopurs_runtime.Value {
	once_intersectBy.Do(func() {
		cache_intersectBy = gopurs_runtime.Func3(func(eq2_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_intersectBy(eq2_0_box, func() []gopurs_runtime.Value {
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
	return cache_intersectBy
}

var cache_intersect gopurs_runtime.Value
var once_intersect sync.Once
func Get_intersect() gopurs_runtime.Value {
	once_intersect.Do(func() {
		cache_intersect = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intersect(dictEq_0_box)
})
	})
	return cache_intersect
}

var cache_elemLastIndex gopurs_runtime.Value
var once_elemLastIndex sync.Once
func Get_elemLastIndex() gopurs_runtime.Value {
	once_elemLastIndex.Do(func() {
		cache_elemLastIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_elemLastIndex(dictEq_0_box, x_1_box)
})
	})
	return cache_elemLastIndex
}

var cache_elemIndex gopurs_runtime.Value
var once_elemIndex sync.Once
func Get_elemIndex() gopurs_runtime.Value {
	once_elemIndex.Do(func() {
		cache_elemIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_elemIndex(dictEq_0_box, x_1_box)
})
	})
	return cache_elemIndex
}

var cache_notElem gopurs_runtime.Value
var once_notElem sync.Once
func Get_notElem() gopurs_runtime.Value {
	once_notElem.Do(func() {
		cache_notElem = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, arr_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_notElem(dictEq_0_box, a_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(arr_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_notElem
}

var cache_elem gopurs_runtime.Value
var once_elem sync.Once
func Get_elem() gopurs_runtime.Value {
	once_elem.Do(func() {
		cache_elem = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, arr_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_elem(dictEq_0_box, a_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(arr_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_elem
}

var cache_dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		cache_dropWhile = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_dropWhile(p_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_dropWhile
}

var cache_dropEnd gopurs_runtime.Value
var once_dropEnd sync.Once
func Get_dropEnd() gopurs_runtime.Value {
	once_dropEnd.Do(func() {
		cache_dropEnd = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_dropEnd(n_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_dropEnd
}

var cache_drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		cache_drop = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_drop(n_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_drop
}

var cache_takeEnd gopurs_runtime.Value
var once_takeEnd sync.Once
func Get_takeEnd() gopurs_runtime.Value {
	once_takeEnd.Do(func() {
		cache_takeEnd = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_takeEnd(n_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_takeEnd
}

var cache_deleteAt gopurs_runtime.Value
var once_deleteAt sync.Once
func Get_deleteAt() gopurs_runtime.Value {
	once_deleteAt.Do(func() {
		cache_deleteAt = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_deleteAt(__local_var_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_deleteAt
}

var cache_deleteBy gopurs_runtime.Value
var once_deleteBy sync.Once
func Get_deleteBy() gopurs_runtime.Value {
	once_deleteBy.Do(func() {
		cache_deleteBy = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_deleteBy(v_0_box, v1_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v2_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_deleteBy
}

var cache_delete_ gopurs_runtime.Value
var once_delete_ sync.Once
func Get_delete_() gopurs_runtime.Value {
	once_delete_.Do(func() {
		cache_delete_ = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_delete_(dictEq_0_box)
})
	})
	return cache_delete_
}

var cache_difference gopurs_runtime.Value
var once_difference sync.Once
func Get_difference() gopurs_runtime.Value {
	once_difference.Do(func() {
		cache_difference = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_difference(dictEq_0_box)
})
	})
	return cache_difference
}

var cache_cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cache_cons = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_cons(x_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_cons
}

var cache_some gopurs_runtime.Value
var once_some sync.Once
func Get_some() gopurs_runtime.Value {
	once_some.Do(func() {
		cache_some = gopurs_runtime.Func3(func(dictAlternative_0_box gopurs_runtime.Value, dictLazy_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_some(dictAlternative_0_box, dictLazy_1_box, v_2_box)
})
	})
	return cache_some
}

var cache_many gopurs_runtime.Value
var once_many sync.Once
func Get_many() gopurs_runtime.Value {
	once_many.Do(func() {
		cache_many = gopurs_runtime.Func3(func(dictAlternative_0_box gopurs_runtime.Value, dictLazy_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_many(dictAlternative_0_box, dictLazy_1_box, v_2_box)
})
	})
	return cache_many
}

var cache_concatMap gopurs_runtime.Value
var once_concatMap sync.Once
func Get_concatMap() gopurs_runtime.Value {
	once_concatMap.Do(func() {
		cache_concatMap = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_concatMap(b_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(a_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_concatMap
}

var cache_mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		cache_mapMaybe = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe(f_0_box)
})
	})
	return cache_mapMaybe
}

var cache_filterA gopurs_runtime.Value
var once_filterA sync.Once
func Get_filterA() gopurs_runtime.Value {
	once_filterA.Do(func() {
		cache_filterA = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filterA(dictApplicative_0_box)
})
	})
	return cache_filterA
}

var cache_catMaybes gopurs_runtime.Value
var once_catMaybes sync.Once
func Get_catMaybes() gopurs_runtime.Value {
	once_catMaybes.Do(func() {
		cache_catMaybes = Call_mapMaybe(gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_catMaybes
}

var cache_any gopurs_runtime.Value
var once_any sync.Once
func Get_any() gopurs_runtime.Value {
	once_any.Do(func() {
		cache_any = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_any(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_any
}

var cache_nubByEq gopurs_runtime.Value
var once_nubByEq sync.Once
func Get_nubByEq() gopurs_runtime.Value {
	once_nubByEq.Do(func() {
		cache_nubByEq = gopurs_runtime.Func2(func(eq2_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_nubByEq(eq2_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_nubByEq
}

var cache_nubEq gopurs_runtime.Value
var once_nubEq sync.Once
func Get_nubEq() gopurs_runtime.Value {
	once_nubEq.Do(func() {
		cache_nubEq = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_nubEq(dictEq_0_box)
})
	})
	return cache_nubEq
}

var cache_unionBy gopurs_runtime.Value
var once_unionBy sync.Once
func Get_unionBy() gopurs_runtime.Value {
	once_unionBy.Do(func() {
		cache_unionBy = gopurs_runtime.Func3(func(eq2_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_unionBy(eq2_0_box, func() []gopurs_runtime.Value {
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
	return cache_unionBy
}

var cache_union gopurs_runtime.Value
var once_union sync.Once
func Get_union() gopurs_runtime.Value {
	once_union.Do(func() {
		cache_union = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_union(dictEq_0_box)
})
	})
	return cache_union
}

var cache_alterAt gopurs_runtime.Value
var once_alterAt sync.Once
func Get_alterAt() gopurs_runtime.Value {
	once_alterAt.Do(func() {
		cache_alterAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_alterAt(i_0_box.IntVal, f_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_alterAt
}

var cache_all gopurs_runtime.Value
var once_all sync.Once
func Get_all() gopurs_runtime.Value {
	once_all.Do(func() {
		cache_all = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_all(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_all
}

func Call_intercalate1(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func2(func(sep_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldl"), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_5, "init").IntVal) != (0) {
__t2 = gopurs_runtime.RecordDict2("acc", "init", v1_6, gopurs_runtime.Bool(false))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), gopurs_runtime.RecordGet(v_5, "acc"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), sep_3, v1_6)), gopurs_runtime.Bool(false))
}
end_branch_2:
return __t2
}), gopurs_runtime.RecordDict2("acc", "init", mempty_2_1, gopurs_runtime.Bool(true)), xs_4), "acc")
})
}

func Call_zipWith(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_zipWithImpl(), __local_var_0, gopurs_runtime.Array(__local_var_1), gopurs_runtime.Array(__local_var_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_zipWithA(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
sequence1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Traversable.Get_traversableArray(), "sequence"), dictApplicative_0)
_ = sequence1_1_0
return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value, ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence1_1_0, gopurs_runtime.UncurriedApp3(Get_zipWithImpl(), f_2, xs_3, ys_4))
})
}

func Call_updateAtIndices(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
traverse_1_1_0 := gopurs_runtime.Apply(Get_traverse_(), dictFoldable_0)
_ = traverse_1_1_0
return gopurs_runtime.Func2(func(us_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_withArray(), gopurs_runtime.Func(func(res_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(traverse_1_1_0, gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_1 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1
_ = __local_var_6_1
__local_var_7_2 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0
_ = __local_var_7_2
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(pkg_Data_Array_ST.Get_pokeImpl(), __local_var_7_2, __local_var_6_1, res_4)
})
}), us_2)
}), xs_3))
})
}

func Call_updateAt(__local_var_0_loop int64, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value] {
var __local_var_0 int64 = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
return (*pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value])(gopurs_runtime.UncurriedApp5(Get__updateAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, gopurs_runtime.Int(__local_var_0), __local_var_1, gopurs_runtime.Array(__local_var_2)).UnsafePtr)
}

func Call_unsafeIndex(_dollar__unused_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value, __local_var_2_loop int64) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 int64 = __local_var_2_loop
_ = __local_var_2
return gopurs_runtime.ArrayAccess(gopurs_runtime.Array(__local_var_1), int(__local_var_2))
}

func Call_uncons(__local_var_0_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
return (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.UncurriedApp3(Get_unconsImpl(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
}), gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.RecordDict2("head", "tail", x_1, xs_2)})}
}), gopurs_runtime.Array(__local_var_0)).UnsafePtr)
}

func Call_toUnfoldable(dictUnfoldable_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
len_2_0 := gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Array(xs_1))))
_ = len_2_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThan(), i_3, len_2_0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.ArrayAccess(gopurs_runtime.Array(xs_1), int(i_3.IntVal)), gopurs_runtime.Int((i_3.IntVal) + (1))})}})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
}
end_branch_1:
return __t1
}), gopurs_runtime.Int(0))
}

func Call_tail(__local_var_0_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value] {
var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
return (*pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value])(gopurs_runtime.UncurriedApp3(Get_unconsImpl(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{xs_2})}
}), gopurs_runtime.Array(__local_var_0)).UnsafePtr)
}

func Call_sortBy(comp_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_sortByImpl(), comp_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 380165415) {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 902936544) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 1527465420) {
__t0 = gopurs_runtime.Int(-1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Array(__local_var_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_sortWith(dictOrd_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply(Get_sortBy(), gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), gopurs_runtime.Apply(f_1, x_2), gopurs_runtime.Apply(f_1, y_3))
}))
}

func Call_sort(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_sortBy(compare_1_0, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
}

func Call_snoc(xs_0_loop []gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) []gopurs_runtime.Value {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_withArray(), gopurs_runtime.Apply(pkg_Data_Array_ST.Get_push(), x_1), gopurs_runtime.Array(xs_0))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_slice(__local_var_0_loop int64, __local_var_1_loop int64, __local_var_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var __local_var_0 int64 = __local_var_0_loop
_ = __local_var_0
var __local_var_1 int64 = __local_var_1_loop
_ = __local_var_1
var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_sliceImpl(), gopurs_runtime.Int(__local_var_0), gopurs_runtime.Int(__local_var_1), gopurs_runtime.Array(__local_var_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_splitAt(v_0_loop int64, v1_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 []gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThanOrEq(), gopurs_runtime.Int(v_0), gopurs_runtime.Int(0)).IntVal) != (0) {
__t0 = gopurs_runtime.RecordDict2("after", "before", gopurs_runtime.Array(v1_1), gopurs_runtime.Array([]gopurs_runtime.Value{}))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict2("after", "before", gopurs_runtime.UncurriedApp3(Get_sliceImpl(), gopurs_runtime.Int(v_0), gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Array(v1_1)))), gopurs_runtime.Array(v1_1)), gopurs_runtime.UncurriedApp3(Get_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int(v_0), gopurs_runtime.Array(v1_1)))
}
end_branch_0:
return __t0
}

func Call_take(n_0_loop int64, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThan(), gopurs_runtime.Int(n_0), gopurs_runtime.Int(1)).IntVal) != (0) {
__t0 = gopurs_runtime.Array([]gopurs_runtime.Value{})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.UncurriedApp3(Get_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int(n_0), gopurs_runtime.Array(xs_1))
}
end_branch_0:
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t0.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_singleton(a_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{a_0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_scanr(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_scanrImpl(), __local_var_0, __local_var_1, gopurs_runtime.Array(__local_var_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_scanl(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_scanlImpl(), __local_var_0, __local_var_1, gopurs_runtime.Array(__local_var_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_replicate(__local_var_0_loop int64, __local_var_1_loop gopurs_runtime.Value) []gopurs_runtime.Value {
var __local_var_0 int64 = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_replicateImpl(), gopurs_runtime.Int(__local_var_0), __local_var_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_range_(__local_var_0_loop int64, __local_var_1_loop int64) []int64 {
var __local_var_0 int64 = __local_var_0_loop
_ = __local_var_0
var __local_var_1 int64 = __local_var_1_loop
_ = __local_var_1
return func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_rangeImpl(), gopurs_runtime.Int(__local_var_0), gopurs_runtime.Int(__local_var_1)).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
}

func Call_partition(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.UncurriedApp2(Get_partitionImpl(), __local_var_0, gopurs_runtime.Array(__local_var_1))
}

func Call_null(xs_0_loop []gopurs_runtime.Value) bool {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
return (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Array(xs_0)))).IntVal) == (0)
}

func Call_modifyAtIndices(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
traverse_1_1_0 := gopurs_runtime.Apply(Get_traverse_(), dictFoldable_0)
_ = traverse_1_1_0
return gopurs_runtime.Func3(func(is_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_withArray(), gopurs_runtime.Func(func(res_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(traverse_1_1_0, gopurs_runtime.Func(func(i_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Data_Array_ST.Get_modify(), i_6, f_3, res_5)
}), is_2)
}), xs_4))
})
}

func Call_intersperse(a_0_loop gopurs_runtime.Value, arr_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var arr_1 []gopurs_runtime.Value = arr_1_loop
_ = arr_1
v_2_0 := gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Array(arr_1))))
_ = v_2_0
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThan(), v_2_0, gopurs_runtime.Int(2)).IntVal) != (0) {
__t3 = gopurs_runtime.Array(arr_1)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), pkg_Data_Array_ST.Get_new_(), gopurs_runtime.Func(func(out_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.ArrayAccess(gopurs_runtime.Array(arr_1), 0)
_ = __local_var_4_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushImpl(), __local_var_4_1, out_3)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_discard(), gopurs_runtime.Apply3(pkg_Control_Monad_ST_Internal.Get_for_(), gopurs_runtime.Int(1), v_2_0, gopurs_runtime.Func(func(idx_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushImpl(), a_0, out_3)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_2 := gopurs_runtime.ArrayAccess(gopurs_runtime.Array(arr_1), int(idx_5.IntVal))
_ = __local_var_7_2
return gopurs_runtime.Apply(Get_void(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushImpl(), __local_var_7_2, out_3)
}))
}))
})), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "pure"), out_3)
}))
}))
})), pkg_Data_Array_ST.Get_unsafeFreeze()))
}
end_branch_3:
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t3.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_intercalate(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
return Call_intercalate1(dictMonoid_0)
}

func Call_insertAt(__local_var_0_loop int64, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value] {
var __local_var_0 int64 = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
return (*pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value])(gopurs_runtime.UncurriedApp5(Get__insertAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, gopurs_runtime.Int(__local_var_0), __local_var_1, gopurs_runtime.Array(__local_var_2)).UnsafePtr)
}

func Call_init_(xs_0_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value] {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Array(xs_0)))).IntVal) == (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.UncurriedApp3(Get_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Array(xs_0)))).IntVal) - (1)), gopurs_runtime.Array(xs_0))})}
}
end_branch_0:
return (*pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value])(__t0.UnsafePtr)
}

func Call_index(__local_var_0_loop []gopurs_runtime.Value, __local_var_1_loop int64) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 int64 = __local_var_1_loop
_ = __local_var_1
return (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, gopurs_runtime.Array(__local_var_0), gopurs_runtime.Int(__local_var_1)).UnsafePtr)
}

func Call_last(xs_0_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
return (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, gopurs_runtime.Array(xs_0), gopurs_runtime.Int((gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Array(xs_0)))).IntVal) - (1))).UnsafePtr)
}

func Call_unsnoc(xs_0_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Array(xs_0)))).IntVal) == (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.UncurriedApp3(Get_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Array(xs_0)))).IntVal) - (1)), gopurs_runtime.Array(xs_0))})}
}
end_branch_0:
return (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("init", "last", v_1, v1_2)
}), __t0), gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, gopurs_runtime.Array(xs_0), gopurs_runtime.Int((gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Array(xs_0)))).IntVal) - (1)))).UnsafePtr)
}

func Call_modifyAt(i_0_loop int64, f_1_loop gopurs_runtime.Value, xs_2_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value] {
var i_0 int64 = i_0_loop
_ = i_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 []gopurs_runtime.Value = xs_2_loop
_ = xs_2
__local_var_3_0 := gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, gopurs_runtime.Array(xs_2), gopurs_runtime.Int(i_0))
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 930809136 && __local_var_3_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 930809136 && __local_var_3_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.UncurriedApp5(Get__updateAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, gopurs_runtime.Int(i_0), gopurs_runtime.Apply(f_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0), gopurs_runtime.Array(xs_2))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return (*pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value])(__t1.UnsafePtr)
}

func Call_span(p_0_loop gopurs_runtime.Value, arr_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var arr_1 []gopurs_runtime.Value = arr_1_loop
_ = arr_1
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(i_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var i_3_loop gopurs_runtime.Value = i_3_loop_val
go__2_0:
for {
if false { continue go__2_0 }
var i_3 gopurs_runtime.Value = i_3_loop
_ = i_3
v_4_1 := gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, gopurs_runtime.Array(arr_1), i_3)
_ = v_4_1
var __t2 gopurs_runtime.Value
{
if (v_4_1.Type == 9 && v_4_1.IntVal == 930809136 && v_4_1.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(p_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_4_1.UnsafePtr).V0).IntVal) != (0) {
i_3_loop = gopurs_runtime.Int((i_3.IntVal) + (1))
continue go__2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{i_3})}
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
if (v_4_1.Type == 9 && v_4_1.IntVal == 930809136 && v_4_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
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
}()
})
breakIndex_3_4 := gopurs_runtime.Apply(go__2_0, gopurs_runtime.Int(0))
_ = breakIndex_3_4
var __t5 gopurs_runtime.Value
{
if (breakIndex_3_4.Type == 9 && breakIndex_3_4.IntVal == 930809136 && breakIndex_3_4.UnsafePtr != nil) {
var __t6 gopurs_runtime.Value
{
if ((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(breakIndex_3_4.UnsafePtr).V0.IntVal) == (0) {
__t6 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Array([]gopurs_runtime.Value{}), gopurs_runtime.Array(arr_1))
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.UncurriedApp3(Get_sliceImpl(), gopurs_runtime.Int(0), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(breakIndex_3_4.UnsafePtr).V0, gopurs_runtime.Array(arr_1)), gopurs_runtime.UncurriedApp3(Get_sliceImpl(), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(breakIndex_3_4.UnsafePtr).V0, gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Array(arr_1)))), gopurs_runtime.Array(arr_1)))
}
end_branch_6:
__t5 = __t6
goto end_branch_5
} else {

}
}
{
if (breakIndex_3_4.Type == 9 && breakIndex_3_4.IntVal == 930809136 && breakIndex_3_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Array(arr_1), gopurs_runtime.Array([]gopurs_runtime.Value{}))
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}

func Call_takeWhile(p_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(Call_span(p_0, xs_1), "init").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_unzip(xs_0_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), pkg_Data_Array_ST.Get_new_(), gopurs_runtime.Func(func(fsts_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), pkg_Data_Array_ST.Get_new_(), gopurs_runtime.Func(func(snds_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.Apply(pkg_Data_Array_ST_Iterator.Get_Iterator(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, gopurs_runtime.Array(xs_0), v_3)
})), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_ref_0 := gopurs_runtime.Int(0)
_ = __local_ref_0
return gopurs_runtime.Any(&__local_ref_0)
})), gopurs_runtime.Func(func(iter_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_discard(), gopurs_runtime.Apply2(pkg_Data_Array_ST_Iterator.Get_iterate(), iter_3, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_1 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0
_ = __local_var_5_1
__local_var_6_2 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1
_ = __local_var_6_2
return gopurs_runtime.Apply2(Get_discard(), gopurs_runtime.Apply(Get_void(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushImpl(), __local_var_5_1, fsts_1)
})), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_void(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushImpl(), __local_var_6_2, snds_2)
}))
}))
})), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_unsafeFreezeImpl(), fsts_1)
}), gopurs_runtime.Func(func(fsts_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_unsafeFreezeImpl(), snds_2)
}), gopurs_runtime.Func(func(snds_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{fsts_prime_5, snds_prime_6})})
}))
}))
}))
}))
}))
})))
}

func Call_head(xs_0_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
return (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, gopurs_runtime.Array(xs_0), gopurs_runtime.Int(0)).UnsafePtr)
}

func Call_nubBy(comp_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
indexedAndSorted_2_0 := gopurs_runtime.Array(Call_sortBy(gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(comp_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1)
}), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_FunctorWithIndex.Get_functorWithIndexArray(), "mapWithIndex"), pkg_Data_Tuple.Get_Tuple(), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
_ = indexedAndSorted_2_0
v_3_1 := gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, indexedAndSorted_2_0, gopurs_runtime.Int(0))
_ = v_3_1
var __t2 gopurs_runtime.Value
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 930809136 && v_3_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Array([]gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 930809136 && v_3_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorArray(), "map"), pkg_Data_Tuple.Get_snd(), gopurs_runtime.Apply(Call_sortWith(gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", pkg_Data_Eq.Get_eqIntImpl())
}), gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})), pkg_Data_Tuple.Get_fst()), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_unsafeThawImpl(), gopurs_runtime.Array([]gopurs_runtime.Value{(*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_3_1.UnsafePtr).V0}))
}), gopurs_runtime.Func(func(result_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_discard(), gopurs_runtime.Apply2(pkg_Control_Monad_ST_Internal.Get_foreach(), indexedAndSorted_2_0, gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_3 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1
_ = __local_var_6_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_4 := gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, x_7, gopurs_runtime.Int((gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(x_7))).IntVal) - (1)))
_ = __local_var_8_4
var __t5 gopurs_runtime.Value
{
if (__local_var_8_4.Type == 9 && __local_var_8_4.IntVal == 930809136 && __local_var_8_4.UnsafePtr != nil) {
__t5 = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_8_4.UnsafePtr).V0.UnsafePtr).V1
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_unsafeFreezeImpl(), result_4)
})), gopurs_runtime.Func(func(lst_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_6 := gopurs_runtime.Apply(Get_void(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushImpl(), v1_5, result_4)
}))
_ = __local_var_8_6
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Apply2(comp_0, lst_7, __local_var_6_3), gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}), gopurs_runtime.Bool(false)).IntVal) != (0) {
__t7 = __local_var_8_6
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "pure"), pkg_Data_Unit.Get_unit())
}
end_branch_7:
return __t7
}))
})), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_unsafeFreezeImpl(), result_4)
})
}))
})))))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_nub(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_nubBy(), gopurs_runtime.RecordGet(dictOrd_0, "compare"))
}

func Call_groupBy(op_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), pkg_Data_Array_ST.Get_new_(), gopurs_runtime.Func(func(result_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.Apply(pkg_Data_Array_ST_Iterator.Get_Iterator(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, gopurs_runtime.Array(xs_1), v_3)
})), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_ref_0 := gopurs_runtime.Int(0)
_ = __local_ref_0
return gopurs_runtime.Any(&__local_ref_0)
})), gopurs_runtime.Func(func(iter_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_discard(), gopurs_runtime.Apply2(pkg_Data_Array_ST_Iterator.Get_iterate(), iter_3, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_void(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), pkg_Data_Array_ST.Get_new_(), gopurs_runtime.Func(func(sub1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushImpl(), x_4, sub1_5)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_discard(), gopurs_runtime.Apply3(pkg_Data_Array_ST_Iterator.Get_pushWhile(), gopurs_runtime.Apply(op_0, x_4), iter_3, sub1_5), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_unsafeFreezeImpl(), sub1_5)
}), gopurs_runtime.Func(func(grp_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushImpl(), grp_8, result_2)
})
}))
}))
}))
})))
})), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_unsafeFreezeImpl(), result_2)
})
}))
}))
}))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_groupAllBy(cmp_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
__local_var_1_0 := gopurs_runtime.Apply(Get_groupBy(), gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Apply2(cmp_0, x_1, y_2), gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil})
}))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Array(Call_sortBy(cmp_0, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))
})
}

func Call_groupAll(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return Call_groupAllBy(gopurs_runtime.RecordGet(dictOrd_0, "compare"))
}

func Call_group(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
eq2_1_0 := gopurs_runtime.RecordGet(dictEq_0, "eq")
_ = eq2_1_0
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_groupBy(eq2_1_0, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
}

func Call_fromFoldable(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
__local_var_1_0 := gopurs_runtime.RecordGet(dictFoldable_0, "foldr")
_ = __local_var_1_0
return gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_fromFoldableImpl(), __local_var_1_0, __local_var_2)
})
}

func Call_transpose(xs_0_loop [][]gopurs_runtime.Value) [][]gopurs_runtime.Value {
var xs_0 [][]gopurs_runtime.Value = xs_0_loop
_ = xs_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(idx_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(allArrays_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var idx_2_loop gopurs_runtime.Value = idx_2_loop_val
var allArrays_3_loop gopurs_runtime.Value = allArrays_3_loop_val
go__1_0:
for {
if false { continue go__1_0 }
var idx_2 gopurs_runtime.Value = idx_2_loop
_ = idx_2
var allArrays_3 gopurs_runtime.Value = allArrays_3_loop
_ = allArrays_3
v_4_1 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldl"), gopurs_runtime.Func2(func(acc_4 gopurs_runtime.Value, nextArr_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_2 := gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, nextArr_5, idx_2)
_ = __local_var_6_2
var __t3 gopurs_runtime.Value
{
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 930809136 && __local_var_6_2.UnsafePtr == nil) {
__t3 = acc_4
goto end_branch_3
} else {

}
}
{
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 930809136 && __local_var_6_2.UnsafePtr != nil) {
var __t4 gopurs_runtime.Value
{
if (acc_4.Type == 9 && acc_4.IntVal == 930809136 && acc_4.UnsafePtr == nil) {
__t4 = gopurs_runtime.Array([]gopurs_runtime.Value{(*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V0})
goto end_branch_4
} else {

}
}
{
if (acc_4.Type == 9 && acc_4.IntVal == 930809136 && acc_4.UnsafePtr != nil) {
__t4 = gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_withArray(), gopurs_runtime.Apply(pkg_Data_Array_ST.Get_push(), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V0), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(acc_4.UnsafePtr).V0))
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{__t4})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}())
_ = v_4_1
var __t5 gopurs_runtime.Value
{
if (v_4_1.Type == 9 && v_4_1.IntVal == 930809136 && v_4_1.UnsafePtr == nil) {
__t5 = allArrays_3
goto end_branch_5
} else {

}
}
{
if (v_4_1.Type == 9 && v_4_1.IntVal == 930809136 && v_4_1.UnsafePtr != nil) {
idx_2_loop = gopurs_runtime.Int((idx_2.IntVal) + (1))
allArrays_3_loop = gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_withArray(), gopurs_runtime.Apply(pkg_Data_Array_ST.Get_push(), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_4_1.UnsafePtr).V0), allArrays_3))
continue go__1_0
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}()
})
})
return func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(go__1_0, gopurs_runtime.Int(0), gopurs_runtime.Array([]gopurs_runtime.Value{})).UnsafePtr)
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

func Call_foldRecM(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value, array_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(o_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_greaterThanOrEq(), gopurs_runtime.RecordGet(o_6, "b"), gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(array_5)))).IntVal) != (0) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.RecordGet(o_6, "a")})})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(f_3, gopurs_runtime.RecordGet(o_6, "a"), gopurs_runtime.ArrayAccess(array_5, int(gopurs_runtime.RecordGet(o_6, "b").IntVal))), gopurs_runtime.Func(func(res_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.RecordDict2("a", "b", res_prime_7, gopurs_runtime.Int((gopurs_runtime.RecordGet(o_6, "b").IntVal) + (1)))})})
}))
}
end_branch_2:
return __t2
}), gopurs_runtime.RecordDict2("a", "b", b_4, gopurs_runtime.Int(0)))
})
}

func Call_foldMap(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldMap"), dictMonoid_0)
}

func Call_foldM(dictMonad_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value, __local_var_3_loop []gopurs_runtime.Value) gopurs_runtime.Value {
foldM:
for {
if false { continue foldM }
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var __local_var_3 []gopurs_runtime.Value = __local_var_3_loop
_ = __local_var_3
return gopurs_runtime.UncurriedApp3(Get_unconsImpl(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), b_2)
}), gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, as_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(f_1, b_2, a_4), gopurs_runtime.Func(func(b_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldM(dictMonad_0, f_1, b_prime_6, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(as_5.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
}))
}), gopurs_runtime.Array(__local_var_3))
}
}

func Call_fold(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldMap"), dictMonoid_0, pkg_Data_Foldable.Get_identity())
}

func Call_findMap(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.UncurriedApp4(Get_findMapImpl(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, pkg_Data_Maybe.Get_isJust(), __local_var_0, gopurs_runtime.Array(__local_var_1)).UnsafePtr)
}

func Call_findLastIndex(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[int64] {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return (*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.UncurriedApp4(Get_findLastIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, __local_var_0, gopurs_runtime.Array(__local_var_1)).UnsafePtr)
}

func Call_insertBy(cmp_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, ys_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var ys_2 []gopurs_runtime.Value = ys_2_loop
_ = ys_2
__local_var_3_1 := gopurs_runtime.UncurriedApp4(Get_findLastIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Apply2(cmp_0, x_1, y_3), gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
}), gopurs_runtime.Array(ys_2))
_ = __local_var_3_1
var __t2 gopurs_runtime.Value
{
if (__local_var_3_1.Type == 9 && __local_var_3_1.IntVal == 930809136 && __local_var_3_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Int(0)
goto end_branch_2
} else {

}
}
{
if (__local_var_3_1.Type == 9 && __local_var_3_1.IntVal == 930809136 && __local_var_3_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.Int(((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_3_1.UnsafePtr).V0.IntVal) + (1))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__local_var_3_0 := gopurs_runtime.UncurriedApp5(Get__insertAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, __t2, x_1, gopurs_runtime.Array(ys_2))
_ = __local_var_3_0
var __t3 gopurs_runtime.Value
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 930809136 && __local_var_3_0.UnsafePtr != nil) {
__t3 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0
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

func Call_insert(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_insertBy(), gopurs_runtime.RecordGet(dictOrd_0, "compare"))
}

func Call_findIndex(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[int64] {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return (*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.UncurriedApp4(Get_findIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, __local_var_0, gopurs_runtime.Array(__local_var_1)).UnsafePtr)
}

func Call_find(f_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.ArrayAccess(gopurs_runtime.Array(xs_1), int(__local_var_2.IntVal))
}), gopurs_runtime.UncurriedApp4(Get_findIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, f_0, gopurs_runtime.Array(xs_1))).UnsafePtr)
}

func Call_filter(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_filterImpl(), __local_var_0, gopurs_runtime.Array(__local_var_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_intersectBy(eq2_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value, ys_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var eq2_0 gopurs_runtime.Value = eq2_0_loop
_ = eq2_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 []gopurs_runtime.Value = ys_2_loop
_ = ys_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_filterImpl(), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.UncurriedApp4(Get_findIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, gopurs_runtime.Apply(eq2_0, x_3), gopurs_runtime.Array(ys_2))
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
return __t1
}), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_intersect(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_intersectBy(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
}

func Call_elemLastIndex(dictEq_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_findLastIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), v_2, x_1)
}))
}

func Call_elemIndex(dictEq_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_findIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), v_2, x_1)
}))
}

func Call_notElem(dictEq_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, arr_2_loop []gopurs_runtime.Value) bool {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var arr_2 []gopurs_runtime.Value = arr_2_loop
_ = arr_2
__local_var_3_0 := gopurs_runtime.UncurriedApp4(Get_findIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), v_3, a_1)
}), gopurs_runtime.Array(arr_2))
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 930809136 && __local_var_3_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 930809136 && __local_var_3_0.UnsafePtr != nil) {
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

func Call_elem(dictEq_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, arr_2_loop []gopurs_runtime.Value) bool {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var arr_2 []gopurs_runtime.Value = arr_2_loop
_ = arr_2
__local_var_3_0 := gopurs_runtime.UncurriedApp4(Get_findIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), v_3, a_1)
}), gopurs_runtime.Array(arr_2))
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 930809136 && __local_var_3_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 930809136 && __local_var_3_0.UnsafePtr != nil) {
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

func Call_dropWhile(p_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(Call_span(p_0, xs_1), "rest").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_dropEnd(n_0_loop int64, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
__local_var_2_0 := (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Array(xs_1)))).IntVal) - (n_0)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThan(), gopurs_runtime.Int(__local_var_2_0), gopurs_runtime.Int(1)).IntVal) != (0) {
__t1 = gopurs_runtime.Array([]gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.UncurriedApp3(Get_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int(__local_var_2_0), gopurs_runtime.Array(xs_1))
}
end_branch_1:
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t1.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_drop(n_0_loop int64, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThan(), gopurs_runtime.Int(n_0), gopurs_runtime.Int(1)).IntVal) != (0) {
__t0 = gopurs_runtime.Array(xs_1)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.UncurriedApp3(Get_sliceImpl(), gopurs_runtime.Int(n_0), gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Array(xs_1)))), gopurs_runtime.Array(xs_1))
}
end_branch_0:
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t0.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_takeEnd(n_0_loop int64, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
__local_var_2_0 := (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Array(xs_1)))).IntVal) - (n_0)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThan(), gopurs_runtime.Int(__local_var_2_0), gopurs_runtime.Int(1)).IntVal) != (0) {
__t1 = gopurs_runtime.Array(xs_1)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.UncurriedApp3(Get_sliceImpl(), gopurs_runtime.Int(__local_var_2_0), gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Array(xs_1)))), gopurs_runtime.Array(xs_1))
}
end_branch_1:
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t1.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_deleteAt(__local_var_0_loop int64, __local_var_1_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value] {
var __local_var_0 int64 = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return (*pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value])(gopurs_runtime.UncurriedApp4(Get__deleteAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, gopurs_runtime.Int(__local_var_0), gopurs_runtime.Array(__local_var_1)).UnsafePtr)
}

func Call_deleteBy(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 []gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Array(v2_2)))).IntVal) == (0) {
__t4 = gopurs_runtime.Array([]gopurs_runtime.Value{})
goto end_branch_4
} else {

}
}
{
__local_var_3_0 := gopurs_runtime.UncurriedApp4(Get_findIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, gopurs_runtime.Apply(v_0, v1_1), gopurs_runtime.Array(v2_2))
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 930809136 && __local_var_3_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Array(v2_2)
goto end_branch_1
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 930809136 && __local_var_3_0.UnsafePtr != nil) {
__local_var_4_2 := gopurs_runtime.UncurriedApp4(Get__deleteAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, gopurs_runtime.Array(v2_2))
_ = __local_var_4_2
var __t3 gopurs_runtime.Value
{
if (__local_var_4_2.Type == 9 && __local_var_4_2.IntVal == 930809136 && __local_var_4_2.UnsafePtr != nil) {
__t3 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_4_2.UnsafePtr).V0
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t1 = __t3
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t4 = __t1
}
end_branch_4:
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t4.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_delete_(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_deleteBy(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
}

func Call_difference(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldr"), gopurs_runtime.Apply(Get_delete_(), dictEq_0))
}

func Call_cons(x_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"), gopurs_runtime.Array([]gopurs_runtime.Value{x_0}), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_some(dictAlternative_0_loop gopurs_runtime.Value, dictLazy_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
var dictLazy_1 gopurs_runtime.Value = dictLazy_1_loop
_ = dictLazy_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_cons(), v_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_1, "defer"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_many(dictAlternative_0, dictLazy_1, v_2)
})))
}

func Call_many(dictAlternative_0_loop gopurs_runtime.Value, dictLazy_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
var dictLazy_1 gopurs_runtime.Value = dictLazy_1_loop
_ = dictLazy_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "alt"), Call_some(dictAlternative_0, dictLazy_1, v_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Array([]gopurs_runtime.Value{})))
}

func Call_concatMap(b_0_loop gopurs_runtime.Value, a_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 []gopurs_runtime.Value = a_1_loop
_ = a_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_bindArray(), "bind"), gopurs_runtime.Array(a_1), b_0).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_mapMaybe(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_concatMap(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(f_0, x_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136 && __local_var_2_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Array([]gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136 && __local_var_2_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Array([]gopurs_runtime.Value{(*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_2_0.UnsafePtr).V0})
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

func Call_filterA(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
traverse1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Traversable.Get_traversableArray(), "traverse"), dictApplicative_0)
_ = traverse1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(p_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(traverse1_1_0, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), x_4), gopurs_runtime.Apply(p_3, x_4))
}))
_ = __local_var_4_2
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "map"), Call_mapMaybe(gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if ((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1.IntVal) != (0) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0})}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
}
end_branch_4:
return __t4
})))
_ = __local_var_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, gopurs_runtime.Apply(__local_var_4_2, x_6))
})
})
}

func Call_any(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) bool {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return (gopurs_runtime.UncurriedApp2(Get_anyImpl(), __local_var_0, gopurs_runtime.Array(__local_var_1)).IntVal) != (0)
}

func Call_nubByEq(eq2_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var eq2_0 gopurs_runtime.Value = eq2_0_loop
_ = eq2_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), pkg_Data_Array_ST.Get_new_(), gopurs_runtime.Func(func(arr_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_discard(), gopurs_runtime.Apply2(pkg_Control_Monad_ST_Internal.Get_foreach(), gopurs_runtime.Array(xs_1), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(Get_any(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(eq2_0, v_4, x_3)
}))
_ = __local_var_4_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Apply(__local_var_4_0, x_5))
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_unsafeFreezeImpl(), arr_2)
})), gopurs_runtime.Func(func(e_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_1 := gopurs_runtime.Apply(Get_void(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushImpl(), x_3, arr_2)
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
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "pure"), pkg_Data_Unit.Get_unit())
}
end_branch_2:
return __t2
}))
})), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_unsafeFreezeImpl(), arr_2)
})
}))
}))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_nubEq(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_nubByEq(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
}

func Call_unionBy(eq2_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value, ys_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var eq2_0 gopurs_runtime.Value = eq2_0_loop
_ = eq2_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 []gopurs_runtime.Value = ys_2_loop
_ = ys_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"), gopurs_runtime.Array(xs_1), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldl"), gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_deleteBy(eq2_0, a_4, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(b_3.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
}), gopurs_runtime.Array(Call_nubByEq(eq2_0, ys_2)), gopurs_runtime.Array(xs_1))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_union(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_unionBy(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
}

func Call_alterAt(i_0_loop int64, f_1_loop gopurs_runtime.Value, xs_2_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value] {
var i_0 int64 = i_0_loop
_ = i_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 []gopurs_runtime.Value = xs_2_loop
_ = xs_2
__local_var_3_0 := gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, gopurs_runtime.Array(xs_2), gopurs_runtime.Int(i_0))
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 930809136 && __local_var_3_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 930809136 && __local_var_3_0.UnsafePtr != nil) {
v_4_2 := gopurs_runtime.Apply(f_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0)
_ = v_4_2
var __t3 gopurs_runtime.Value
{
if (v_4_2.Type == 9 && v_4_2.IntVal == 930809136 && v_4_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.UncurriedApp4(Get__deleteAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, gopurs_runtime.Int(i_0), gopurs_runtime.Array(xs_2))
goto end_branch_3
} else {

}
}
{
if (v_4_2.Type == 9 && v_4_2.IntVal == 930809136 && v_4_2.UnsafePtr != nil) {
__t3 = gopurs_runtime.UncurriedApp5(Get__updateAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, gopurs_runtime.Int(i_0), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_4_2.UnsafePtr).V0, gopurs_runtime.Array(xs_2))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t1 = __t3
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return (*pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value])(__t1.UnsafePtr)
}

func Call_all(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) bool {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return (gopurs_runtime.UncurriedApp2(Get_allImpl(), __local_var_0, gopurs_runtime.Array(__local_var_1)).IntVal) != (0)
}

func Get__deleteAt() gopurs_runtime.Value {
	return _Gopurs__DeleteAt
}

func Get__insertAt() gopurs_runtime.Value {
	return _Gopurs__InsertAt
}

func Get__updateAt() gopurs_runtime.Value {
	return _Gopurs__UpdateAt
}

func Get_allImpl() gopurs_runtime.Value {
	return _Gopurs_AllImpl
}

func Get_anyImpl() gopurs_runtime.Value {
	return _Gopurs_AnyImpl
}

func Get_concat() gopurs_runtime.Value {
	return _Gopurs_Concat
}

func Get_filterImpl() gopurs_runtime.Value {
	return _Gopurs_FilterImpl
}

func Get_findIndexImpl() gopurs_runtime.Value {
	return _Gopurs_FindIndexImpl
}

func Get_findLastIndexImpl() gopurs_runtime.Value {
	return _Gopurs_FindLastIndexImpl
}

func Get_findMapImpl() gopurs_runtime.Value {
	return _Gopurs_FindMapImpl
}

func Get_fromFoldableImpl() gopurs_runtime.Value {
	return _Gopurs_FromFoldableImpl
}

func Get_indexImpl() gopurs_runtime.Value {
	return _Gopurs_IndexImpl
}

func Get_length() gopurs_runtime.Value {
	return _Gopurs_Length
}

func Get_partitionImpl() gopurs_runtime.Value {
	return _Gopurs_PartitionImpl
}

func Get_rangeImpl() gopurs_runtime.Value {
	return _Gopurs_RangeImpl
}

func Get_replicateImpl() gopurs_runtime.Value {
	return _Gopurs_ReplicateImpl
}

func Get_reverse() gopurs_runtime.Value {
	return _Gopurs_Reverse
}

func Get_scanlImpl() gopurs_runtime.Value {
	return _Gopurs_ScanlImpl
}

func Get_scanrImpl() gopurs_runtime.Value {
	return _Gopurs_ScanrImpl
}

func Get_sliceImpl() gopurs_runtime.Value {
	return _Gopurs_SliceImpl
}

func Get_sortByImpl() gopurs_runtime.Value {
	return _Gopurs_SortByImpl
}

func Get_unconsImpl() gopurs_runtime.Value {
	return _Gopurs_UnconsImpl
}

func Get_unsafeIndexImpl() gopurs_runtime.Value {
	return _Gopurs_UnsafeIndexImpl
}

func Get_zipWithImpl() gopurs_runtime.Value {
	return _Gopurs_ZipWithImpl
}

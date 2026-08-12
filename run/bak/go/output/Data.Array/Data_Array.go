package Data_Array

import (
	pkg_Control_Alt "gopurs/output/Control.Alt"
	pkg_Control_Alternative "gopurs/output/Control.Alternative"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Lazy "gopurs/output/Control.Lazy"
	pkg_Control_Monad "gopurs/output/Control.Monad"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
	pkg_Control_Monad_ST_Internal "gopurs/output/Control.Monad.ST.Internal"
	pkg_Control_Monad_ST_Uncurried "gopurs/output/Control.Monad.ST.Uncurried"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Array_ST "gopurs/output/Data.Array.ST"
	pkg_Data_Array_ST_Iterator "gopurs/output/Data.Array.ST.Iterator"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_FunctorWithIndex "gopurs/output/Data.FunctorWithIndex"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Monoid "gopurs/output/Data.Monoid"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_intercalate1 gopurs_runtime.Value
var once_intercalate1 sync.Once
func Get_intercalate1() gopurs_runtime.Value {
	once_intercalate1.Do(func() {
		cache_intercalate1 = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intercalate1(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_0_box))
})
	})
	return cache_intercalate1
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
		cache_zipWithA = gopurs_runtime.Func4(func(dictApplicative_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value, ys_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWithA(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), f_1_box, func() []gopurs_runtime.Value {
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
		cache_updateAtIndices = gopurs_runtime.Func3(func(dictFoldable_0_box gopurs_runtime.Value, us_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_updateAtIndices(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), us_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
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
return Call_toUnfoldable(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box), func() []gopurs_runtime.Value {
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
return Call_sortWith(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), f_1_box)
})
	})
	return cache_sortWith
}

var cache_sort gopurs_runtime.Value
var once_sort sync.Once
func Get_sort() gopurs_runtime.Value {
	once_sort.Do(func() {
		cache_sort = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sort(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
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

var cache_go__range gopurs_runtime.Value
var once_go__range sync.Once
func Get_go__range() gopurs_runtime.Value {
	once_go__range.Do(func() {
		cache_go__range = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_go__range(__local_var_0_box.IntVal, __local_var_1_box.IntVal)
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_go__range
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
		cache_modifyAtIndices = gopurs_runtime.Func4(func(dictFoldable_0_box gopurs_runtime.Value, is_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_modifyAtIndices(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), is_1_box, f_2_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_3_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
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
return Call_intercalate(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_0_box))
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

var cache_init gopurs_runtime.Value
var once_init sync.Once
func Get_init() gopurs_runtime.Value {
	once_init.Do(func() {
		cache_init = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_init(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_init
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

var cache_unzip gopurs_runtime.Value
var once_unzip sync.Once
func Get_unzip() gopurs_runtime.Value {
	once_unzip.Do(func() {
		cache_unzip = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_unzip(func() []*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v) }
					return unboxed
				}()))}
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
return Call_nub(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_nub
}

var cache_groupBy gopurs_runtime.Value
var once_groupBy sync.Once
func Get_groupBy() gopurs_runtime.Value {
	once_groupBy.Do(func() {
		cache_groupBy = gopurs_runtime.Func2(func(op_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_groupBy(op_0_box, func() []gopurs_runtime.Value {
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
return Call_groupAll(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_groupAll
}

var cache_group gopurs_runtime.Value
var once_group sync.Once
func Get_group() gopurs_runtime.Value {
	once_group.Do(func() {
		cache_group = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_group(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_group
}

var cache_fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		cache_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldable(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
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
return Call_foldRecM(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box))
})
	})
	return cache_foldRecM
}

var cache_foldMap gopurs_runtime.Value
var once_foldMap sync.Once
func Get_foldMap() gopurs_runtime.Value {
	once_foldMap.Do(func() {
		cache_foldMap = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_0_box))
})
	})
	return cache_foldMap
}

var cache_foldM gopurs_runtime.Value
var once_foldM sync.Once
func Get_foldM() gopurs_runtime.Value {
	once_foldM.Do(func() {
		cache_foldM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldM(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_foldM
}

var cache_fold gopurs_runtime.Value
var once_fold sync.Once
func Get_fold() gopurs_runtime.Value {
	once_fold.Do(func() {
		cache_fold = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fold(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_0_box))
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
return Call_insert(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
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
		cache_intersectBy = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_intersectBy(eq_0_box, func() []gopurs_runtime.Value {
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
return Call_intersect(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_intersect
}

var cache_elemLastIndex gopurs_runtime.Value
var once_elemLastIndex sync.Once
func Get_elemLastIndex() gopurs_runtime.Value {
	once_elemLastIndex.Do(func() {
		cache_elemLastIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_elemLastIndex(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
})
	})
	return cache_elemLastIndex
}

var cache_elemIndex gopurs_runtime.Value
var once_elemIndex sync.Once
func Get_elemIndex() gopurs_runtime.Value {
	once_elemIndex.Do(func() {
		cache_elemIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_elemIndex(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
})
	})
	return cache_elemIndex
}

var cache_notElem gopurs_runtime.Value
var once_notElem sync.Once
func Get_notElem() gopurs_runtime.Value {
	once_notElem.Do(func() {
		cache_notElem = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, arr_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_notElem(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box), a_1_box, func() []gopurs_runtime.Value {
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
return gopurs_runtime.Bool(Call_elem(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box), a_1_box, func() []gopurs_runtime.Value {
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

var cache_delete gopurs_runtime.Value
var once_delete sync.Once
func Get_delete() gopurs_runtime.Value {
	once_delete.Do(func() {
		cache_delete = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_delete(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_delete
}

var cache_difference gopurs_runtime.Value
var once_difference sync.Once
func Get_difference() gopurs_runtime.Value {
	once_difference.Do(func() {
		cache_difference = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_difference(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box))
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
		cache_some = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_some(gopurs_runtime.CoerceToStruct[pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]](dictAlternative_0_box))
})
	})
	return cache_some
}

var cache_many gopurs_runtime.Value
var once_many sync.Once
func Get_many() gopurs_runtime.Value {
	once_many.Do(func() {
		cache_many = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_many(gopurs_runtime.CoerceToStruct[pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]](dictAlternative_0_box))
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
return Call_filterA(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
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
		cache_nubByEq = gopurs_runtime.Func2(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_nubByEq(eq_0_box, func() []gopurs_runtime.Value {
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
return Call_nubEq(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_nubEq
}

var cache_unionBy gopurs_runtime.Value
var once_unionBy sync.Once
func Get_unionBy() gopurs_runtime.Value {
	once_unionBy.Do(func() {
		cache_unionBy = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_unionBy(eq_0_box, func() []gopurs_runtime.Value {
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
return Call_union(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box))
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

var cache_alt__267341625 gopurs_runtime.Value
var once_alt__267341625 sync.Once
func Get_alt__267341625() gopurs_runtime.Value {
	once_alt__267341625.Do(func() {
		cache_alt__267341625 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alt__267341625(gopurs_runtime.CoerceToStruct[pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_alt__267341625
}

var cache_alt__3999254335 gopurs_runtime.Value
var once_alt__3999254335 sync.Once
func Get_alt__3999254335() gopurs_runtime.Value {
	once_alt__3999254335.Do(func() {
		cache_alt__3999254335 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alt__3999254335(gopurs_runtime.CoerceToStruct[pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_alt__3999254335
}

var cache_pure__2935994064 gopurs_runtime.Value
var once_pure__2935994064 sync.Once
func Get_pure__2935994064() gopurs_runtime.Value {
	once_pure__2935994064.Do(func() {
		cache_pure__2935994064 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__2935994064(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__2935994064
}

var cache_pure__3215807376 gopurs_runtime.Value
var once_pure__3215807376 sync.Once
func Get_pure__3215807376() gopurs_runtime.Value {
	once_pure__3215807376.Do(func() {
		cache_pure__3215807376 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__3215807376(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__3215807376
}

var cache_pure__2960485136 gopurs_runtime.Value
var once_pure__2960485136 sync.Once
func Get_pure__2960485136() gopurs_runtime.Value {
	once_pure__2960485136.Do(func() {
		cache_pure__2960485136 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__2960485136(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__2960485136
}

var cache_pure__566620048 gopurs_runtime.Value
var once_pure__566620048 sync.Once
func Get_pure__566620048() gopurs_runtime.Value {
	once_pure__566620048.Do(func() {
		cache_pure__566620048 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__566620048(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__566620048
}

var cache_pure__3079134646 gopurs_runtime.Value
var once_pure__3079134646 sync.Once
func Get_pure__3079134646() gopurs_runtime.Value {
	once_pure__3079134646.Do(func() {
		cache_pure__3079134646 = gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "pure")
	})
	return cache_pure__3079134646
}

var cache_when__1954875249 gopurs_runtime.Value
var once_when__1954875249 sync.Once
func Get_when__1954875249() gopurs_runtime.Value {
	once_when__1954875249.Do(func() {
		cache_when__1954875249 = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_when__1954875249(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), (v_1_box.IntVal) != (0), v1_2_box)
})
	})
	return cache_when__1954875249
}

var cache_when__4245282944 gopurs_runtime.Value
var once_when__4245282944 sync.Once
func Get_when__4245282944() gopurs_runtime.Value {
	once_when__4245282944.Do(func() {
		cache_when__4245282944 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_when__4245282944((v_0_box.IntVal) != (0), v1_1_box)
})
	})
	return cache_when__4245282944
}

var cache_apply__353515660 gopurs_runtime.Value
var once_apply__353515660 sync.Once
func Get_apply__353515660() gopurs_runtime.Value {
	once_apply__353515660.Do(func() {
		cache_apply__353515660 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__353515660(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__353515660
}

var cache_apply__3582447820 gopurs_runtime.Value
var once_apply__3582447820 sync.Once
func Get_apply__3582447820() gopurs_runtime.Value {
	once_apply__3582447820.Do(func() {
		cache_apply__3582447820 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__3582447820(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__3582447820
}

var cache_apply__2485508138 gopurs_runtime.Value
var once_apply__2485508138 sync.Once
func Get_apply__2485508138() gopurs_runtime.Value {
	once_apply__2485508138.Do(func() {
		cache_apply__2485508138 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply")
	})
	return cache_apply__2485508138
}

var cache_bind__2601835655 gopurs_runtime.Value
var once_bind__2601835655 sync.Once
func Get_bind__2601835655() gopurs_runtime.Value {
	once_bind__2601835655.Do(func() {
		cache_bind__2601835655 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__2601835655(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__2601835655
}

var cache_bind__760543367 gopurs_runtime.Value
var once_bind__760543367 sync.Once
func Get_bind__760543367() gopurs_runtime.Value {
	once_bind__760543367.Do(func() {
		cache_bind__760543367 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__760543367(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__760543367
}

var cache_bind__3352508289 gopurs_runtime.Value
var once_bind__3352508289 sync.Once
func Get_bind__3352508289() gopurs_runtime.Value {
	once_bind__3352508289.Do(func() {
		cache_bind__3352508289 = gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind")
	})
	return cache_bind__3352508289
}

var cache_bind__1671968481 gopurs_runtime.Value
var once_bind__1671968481 sync.Once
func Get_bind__1671968481() gopurs_runtime.Value {
	once_bind__1671968481.Do(func() {
		cache_bind__1671968481 = gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind")
	})
	return cache_bind__1671968481
}

var cache_bind__3115293729 gopurs_runtime.Value
var once_bind__3115293729 sync.Once
func Get_bind__3115293729() gopurs_runtime.Value {
	once_bind__3115293729.Do(func() {
		cache_bind__3115293729 = gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind")
	})
	return cache_bind__3115293729
}

var cache_bind__568923361 gopurs_runtime.Value
var once_bind__568923361 sync.Once
func Get_bind__568923361() gopurs_runtime.Value {
	once_bind__568923361.Do(func() {
		cache_bind__568923361 = gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind")
	})
	return cache_bind__568923361
}

var cache_bind__3897039777 gopurs_runtime.Value
var once_bind__3897039777 sync.Once
func Get_bind__3897039777() gopurs_runtime.Value {
	once_bind__3897039777.Do(func() {
		cache_bind__3897039777 = gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind")
	})
	return cache_bind__3897039777
}

var cache_bind__1491025313 gopurs_runtime.Value
var once_bind__1491025313 sync.Once
func Get_bind__1491025313() gopurs_runtime.Value {
	once_bind__1491025313.Do(func() {
		cache_bind__1491025313 = gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind")
	})
	return cache_bind__1491025313
}

var cache_bind__555709729 gopurs_runtime.Value
var once_bind__555709729 sync.Once
func Get_bind__555709729() gopurs_runtime.Value {
	once_bind__555709729.Do(func() {
		cache_bind__555709729 = gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind")
	})
	return cache_bind__555709729
}

var cache_bind__2501247745 gopurs_runtime.Value
var once_bind__2501247745 sync.Once
func Get_bind__2501247745() gopurs_runtime.Value {
	once_bind__2501247745.Do(func() {
		cache_bind__2501247745 = gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind")
	})
	return cache_bind__2501247745
}

var cache_bind__1160188801 gopurs_runtime.Value
var once_bind__1160188801 sync.Once
func Get_bind__1160188801() gopurs_runtime.Value {
	once_bind__1160188801.Do(func() {
		cache_bind__1160188801 = gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind")
	})
	return cache_bind__1160188801
}

var cache_bind__3753603617 gopurs_runtime.Value
var once_bind__3753603617 sync.Once
func Get_bind__3753603617() gopurs_runtime.Value {
	once_bind__3753603617.Do(func() {
		cache_bind__3753603617 = gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind")
	})
	return cache_bind__3753603617
}

var cache_bind__3937147233 gopurs_runtime.Value
var once_bind__3937147233 sync.Once
func Get_bind__3937147233() gopurs_runtime.Value {
	once_bind__3937147233.Do(func() {
		cache_bind__3937147233 = gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind")
	})
	return cache_bind__3937147233
}

var cache_discard__1876171936 gopurs_runtime.Value
var once_discard__1876171936 sync.Once
func Get_discard__1876171936() gopurs_runtime.Value {
	once_discard__1876171936.Do(func() {
		cache_discard__1876171936 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](pkg_Control_Monad_ST_Internal.Get_bindST()))})
	})
	return cache_discard__1876171936
}

var cache_discard__317162198 gopurs_runtime.Value
var once_discard__317162198 sync.Once
func Get_discard__317162198() gopurs_runtime.Value {
	once_discard__317162198.Do(func() {
		cache_discard__317162198 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_discard__317162198(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Discard[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_discard__317162198
}

var cache_discardUnit__2687062302 gopurs_runtime.Value
var once_discardUnit__2687062302 sync.Once
func Get_discardUnit__2687062302() gopurs_runtime.Value {
	once_discardUnit__2687062302.Do(func() {
		cache_discardUnit__2687062302 = gopurs_runtime.RecordDict1("discard", gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictBind_0, "bind")
}))
	})
	return cache_discardUnit__2687062302
}

var cache_defer__3967925939 gopurs_runtime.Value
var once_defer__3967925939 sync.Once
func Get_defer__3967925939() gopurs_runtime.Value {
	once_defer__3967925939.Do(func() {
		cache_defer__3967925939 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_defer__3967925939(gopurs_runtime.CoerceToStruct[pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_defer__3967925939
}

var cache_defer__308916730 gopurs_runtime.Value
var once_defer__308916730 sync.Once
func Get_defer__308916730() gopurs_runtime.Value {
	once_defer__308916730.Do(func() {
		cache_defer__308916730 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_defer__308916730(gopurs_runtime.CoerceToStruct[pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_defer__308916730
}

var cache_tailRecM__3865988408 gopurs_runtime.Value
var once_tailRecM__3865988408 sync.Once
func Get_tailRecM__3865988408() gopurs_runtime.Value {
	once_tailRecM__3865988408.Do(func() {
		cache_tailRecM__3865988408 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM__3865988408(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_tailRecM__3865988408
}

var cache_tailRecM__1444729948 gopurs_runtime.Value
var once_tailRecM__1444729948 sync.Once
func Get_tailRecM__1444729948() gopurs_runtime.Value {
	once_tailRecM__1444729948.Do(func() {
		cache_tailRecM__1444729948 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM__1444729948(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_tailRecM__1444729948
}

var cache_tailRecM2__3241310373 gopurs_runtime.Value
var once_tailRecM2__3241310373 sync.Once
func Get_tailRecM2__3241310373() gopurs_runtime.Value {
	once_tailRecM2__3241310373.Do(func() {
		cache_tailRecM2__3241310373 = gopurs_runtime.Func4(func(dictMonadRec_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM2__3241310373(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), f_1_box, a_2_box, b_3_box)
})
	})
	return cache_tailRecM2__3241310373
}

var cache_tailRecM2__1943630176 gopurs_runtime.Value
var once_tailRecM2__1943630176 sync.Once
func Get_tailRecM2__1943630176() gopurs_runtime.Value {
	once_tailRecM2__1943630176.Do(func() {
		cache_tailRecM2__1943630176 = gopurs_runtime.Func4(func(dictMonadRec_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value, b_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM2__1943630176(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), f_1_box, a_2_box, b_3_box)
})
	})
	return cache_tailRecM2__1943630176
}

var cache_applicativeST__3091537981 gopurs_runtime.Value
var once_applicativeST__3091537981 sync.Once
func Get_applicativeST__3091537981() gopurs_runtime.Value {
	once_applicativeST__3091537981.Do(func() {
		cache_applicativeST__3091537981 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad_ST_Internal.Get_applyST()
}), pkg_Control_Monad_ST_Internal.Get_pure_())
	})
	return cache_applicativeST__3091537981
}

var cache_applyST__2741064779 gopurs_runtime.Value
var once_applyST__2741064779 sync.Once
func Get_applyST__2741064779() gopurs_runtime.Value {
	once_applyST__2741064779.Do(func() {
		cache_applyST__2741064779 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_monadST(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_monadST(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad_ST_Internal.Get_functorST()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_applyST__2741064779
}

var cache_bindST__2435660861 gopurs_runtime.Value
var once_bindST__2435660861 sync.Once
func Get_bindST__2435660861() gopurs_runtime.Value {
	once_bindST__2435660861.Do(func() {
		cache_bindST__2435660861 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad_ST_Internal.Get_applyST()
}), pkg_Control_Monad_ST_Internal.Get_bind_())
	})
	return cache_bindST__2435660861
}

var cache_for__1203933728 gopurs_runtime.Value
var once_for__1203933728 sync.Once
func Get_for__1203933728() gopurs_runtime.Value {
	once_for__1203933728.Do(func() {
		cache_for__1203933728 = pkg_Control_Monad_ST_Internal.Get_forImpl()
	})
	return cache_for__1203933728
}

var cache_for__956375728 gopurs_runtime.Value
var once_for__956375728 sync.Once
func Get_for__956375728() gopurs_runtime.Value {
	once_for__956375728.Do(func() {
		cache_for__956375728 = pkg_Control_Monad_ST_Internal.Get_forImpl()
	})
	return cache_for__956375728
}

var cache_functorST__4062753802 gopurs_runtime.Value
var once_functorST__4062753802 sync.Once
func Get_functorST__4062753802() gopurs_runtime.Value {
	once_functorST__4062753802.Do(func() {
		cache_functorST__4062753802 = gopurs_runtime.RecordDict1("map", pkg_Control_Monad_ST_Internal.Get_map_())
	})
	return cache_functorST__4062753802
}

var cache_functorST__2441840241 gopurs_runtime.Value
var once_functorST__2441840241 sync.Once
func Get_functorST__2441840241() gopurs_runtime.Value {
	once_functorST__2441840241.Do(func() {
		cache_functorST__2441840241 = gopurs_runtime.RecordDict1("map", pkg_Control_Monad_ST_Internal.Get_map_())
	})
	return cache_functorST__2441840241
}

var cache_modify__3866314397 gopurs_runtime.Value
var once_modify__3866314397 sync.Once
func Get_modify__3866314397() gopurs_runtime.Value {
	once_modify__3866314397.Do(func() {
		cache_modify__3866314397 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modify__3866314397(f_0_box)
})
	})
	return cache_modify__3866314397
}

var cache_modify__781734141 gopurs_runtime.Value
var once_modify__781734141 sync.Once
func Get_modify__781734141() gopurs_runtime.Value {
	once_modify__781734141.Do(func() {
		cache_modify__781734141 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modify__781734141(f_0_box)
})
	})
	return cache_modify__781734141
}

var cache_modify_prime__1497736571 gopurs_runtime.Value
var once_modify_prime__1497736571 sync.Once
func Get_modify_prime__1497736571() gopurs_runtime.Value {
	once_modify_prime__1497736571.Do(func() {
		cache_modify_prime__1497736571 = pkg_Control_Monad_ST_Internal.Get_modifyImpl()
	})
	return cache_modify_prime__1497736571
}

var cache_new__3579768924 gopurs_runtime.Value
var once_new__3579768924 sync.Once
func Get_new__3579768924() gopurs_runtime.Value {
	once_new__3579768924.Do(func() {
		cache_new__3579768924 = pkg_Control_Monad_ST_Internal.Get_newImpl()
	})
	return cache_new__3579768924
}

var cache_new__122671164 gopurs_runtime.Value
var once_new__122671164 sync.Once
func Get_new__122671164() gopurs_runtime.Value {
	once_new__122671164.Do(func() {
		cache_new__122671164 = pkg_Control_Monad_ST_Internal.Get_newImpl()
	})
	return cache_new__122671164
}

var cache_new__2010968700 gopurs_runtime.Value
var once_new__2010968700 sync.Once
func Get_new__2010968700() gopurs_runtime.Value {
	once_new__2010968700.Do(func() {
		cache_new__2010968700 = pkg_Control_Monad_ST_Internal.Get_newImpl()
	})
	return cache_new__2010968700
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

var cache_composeFlipped__2583068543 gopurs_runtime.Value
var once_composeFlipped__2583068543 sync.Once
func Get_composeFlipped__2583068543() gopurs_runtime.Value {
	once_composeFlipped__2583068543.Do(func() {
		cache_composeFlipped__2583068543 = gopurs_runtime.Func3(func(dictSemigroupoid_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_composeFlipped__2583068543(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dictSemigroupoid_0_box), f_1_box, g_2_box)
})
	})
	return cache_composeFlipped__2583068543
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

var cache_iterate__1936300090 gopurs_runtime.Value
var once_iterate__1936300090 sync.Once
func Get_iterate__1936300090() gopurs_runtime.Value {
	once_iterate__1936300090.Do(func() {
		cache_iterate__1936300090 = gopurs_runtime.Func2(func(iter_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_iterate__1936300090(gopurs_runtime.CoerceToStruct[pkg_Data_Array_ST_Iterator.Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](iter_0_box), f_1_box)
})
	})
	return cache_iterate__1936300090
}

var cache_iterate__1835948090 gopurs_runtime.Value
var once_iterate__1835948090 sync.Once
func Get_iterate__1835948090() gopurs_runtime.Value {
	once_iterate__1835948090.Do(func() {
		cache_iterate__1835948090 = gopurs_runtime.Func2(func(iter_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_iterate__1835948090(gopurs_runtime.CoerceToStruct[pkg_Data_Array_ST_Iterator.Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](iter_0_box), f_1_box)
})
	})
	return cache_iterate__1835948090
}

var cache_iterator__1149050118 gopurs_runtime.Value
var once_iterator__1149050118 sync.Once
func Get_iterator__1149050118() gopurs_runtime.Value {
	once_iterator__1149050118.Do(func() {
		cache_iterator__1149050118 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_iterator__1149050118(f_0_box)
})
	})
	return cache_iterator__1149050118
}

var cache_iterator__2947907462 gopurs_runtime.Value
var once_iterator__2947907462 sync.Once
func Get_iterator__2947907462() gopurs_runtime.Value {
	once_iterator__2947907462.Do(func() {
		cache_iterator__2947907462 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_iterator__2947907462(f_0_box)
})
	})
	return cache_iterator__2947907462
}

var cache_next__2731492779 gopurs_runtime.Value
var once_next__2731492779 sync.Once
func Get_next__2731492779() gopurs_runtime.Value {
	once_next__2731492779.Do(func() {
		cache_next__2731492779 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_next__2731492779(gopurs_runtime.CoerceToStruct[pkg_Data_Array_ST_Iterator.Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_next__2731492779
}

var cache_peek__2731492779 gopurs_runtime.Value
var once_peek__2731492779 sync.Once
func Get_peek__2731492779() gopurs_runtime.Value {
	once_peek__2731492779.Do(func() {
		cache_peek__2731492779 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_peek__2731492779(gopurs_runtime.CoerceToStruct[pkg_Data_Array_ST_Iterator.Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_peek__2731492779
}

var cache_pushWhile__2298419255 gopurs_runtime.Value
var once_pushWhile__2298419255 sync.Once
func Get_pushWhile__2298419255() gopurs_runtime.Value {
	once_pushWhile__2298419255.Do(func() {
		cache_pushWhile__2298419255 = gopurs_runtime.Func3(func(p_0_box gopurs_runtime.Value, iter_1_box gopurs_runtime.Value, array_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pushWhile__2298419255(p_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Array_ST_Iterator.Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](iter_1_box), array_2_box)
})
	})
	return cache_pushWhile__2298419255
}

var cache_modify__3999472078 gopurs_runtime.Value
var once_modify__3999472078 sync.Once
func Get_modify__3999472078() gopurs_runtime.Value {
	once_modify__3999472078.Do(func() {
		cache_modify__3999472078 = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modify__3999472078(i_0_box.IntVal, f_1_box, xs_2_box)
})
	})
	return cache_modify__3999472078
}

var cache_new__361268277 gopurs_runtime.Value
var once_new__361268277 sync.Once
func Get_new__361268277() gopurs_runtime.Value {
	once_new__361268277.Do(func() {
		cache_new__361268277 = pkg_Data_Array_ST.Get_newImpl()
	})
	return cache_new__361268277
}

var cache_new__108694771 gopurs_runtime.Value
var once_new__108694771 sync.Once
func Get_new__108694771() gopurs_runtime.Value {
	once_new__108694771.Do(func() {
		cache_new__108694771 = pkg_Data_Array_ST.Get_newImpl()
	})
	return cache_new__108694771
}

var cache_peek__1114289340 gopurs_runtime.Value
var once_peek__1114289340 sync.Once
func Get_peek__1114289340() gopurs_runtime.Value {
	once_peek__1114289340.Do(func() {
		cache_peek__1114289340 = gopurs_runtime.Apply3(pkg_Control_Monad_ST_Uncurried.Get_runSTFn4(), pkg_Data_Array_ST.Get_peekImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})
	})
	return cache_peek__1114289340
}

var cache_poke__1035212550 gopurs_runtime.Value
var once_poke__1035212550 sync.Once
func Get_poke__1035212550() gopurs_runtime.Value {
	once_poke__1035212550.Do(func() {
		cache_poke__1035212550 = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn3(), pkg_Data_Array_ST.Get_pokeImpl())
	})
	return cache_poke__1035212550
}

var cache_push__1557574173 gopurs_runtime.Value
var once_push__1557574173 sync.Once
func Get_push__1557574173() gopurs_runtime.Value {
	once_push__1557574173.Do(func() {
		cache_push__1557574173 = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn2(), pkg_Data_Array_ST.Get_pushImpl())
	})
	return cache_push__1557574173
}

var cache_push__2784843933 gopurs_runtime.Value
var once_push__2784843933 sync.Once
func Get_push__2784843933() gopurs_runtime.Value {
	once_push__2784843933.Do(func() {
		cache_push__2784843933 = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn2(), pkg_Data_Array_ST.Get_pushImpl())
	})
	return cache_push__2784843933
}

var cache_push__3387684349 gopurs_runtime.Value
var once_push__3387684349 sync.Once
func Get_push__3387684349() gopurs_runtime.Value {
	once_push__3387684349.Do(func() {
		cache_push__3387684349 = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn2(), pkg_Data_Array_ST.Get_pushImpl())
	})
	return cache_push__3387684349
}

var cache_run__1673253202 gopurs_runtime.Value
var once_run__1673253202 sync.Once
func Get_run__1673253202() gopurs_runtime.Value {
	once_run__1673253202.Do(func() {
		cache_run__1673253202 = gopurs_runtime.Func(func(st_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_run__1673253202(st_0_box))
})
	})
	return cache_run__1673253202
}

var cache_thaw__2850628855 gopurs_runtime.Value
var once_thaw__2850628855 sync.Once
func Get_thaw__2850628855() gopurs_runtime.Value {
	once_thaw__2850628855.Do(func() {
		cache_thaw__2850628855 = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), pkg_Data_Array_ST.Get_thawImpl())
	})
	return cache_thaw__2850628855
}

var cache_unsafeFreeze__2265612059 gopurs_runtime.Value
var once_unsafeFreeze__2265612059 sync.Once
func Get_unsafeFreeze__2265612059() gopurs_runtime.Value {
	once_unsafeFreeze__2265612059.Do(func() {
		cache_unsafeFreeze__2265612059 = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), pkg_Data_Array_ST.Get_unsafeFreezeImpl())
	})
	return cache_unsafeFreeze__2265612059
}

var cache_unsafeFreeze__3261307415 gopurs_runtime.Value
var once_unsafeFreeze__3261307415 sync.Once
func Get_unsafeFreeze__3261307415() gopurs_runtime.Value {
	once_unsafeFreeze__3261307415.Do(func() {
		cache_unsafeFreeze__3261307415 = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), pkg_Data_Array_ST.Get_unsafeFreezeImpl())
	})
	return cache_unsafeFreeze__3261307415
}

var cache_unsafeFreeze__3249621719 gopurs_runtime.Value
var once_unsafeFreeze__3249621719 sync.Once
func Get_unsafeFreeze__3249621719() gopurs_runtime.Value {
	once_unsafeFreeze__3249621719.Do(func() {
		cache_unsafeFreeze__3249621719 = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), pkg_Data_Array_ST.Get_unsafeFreezeImpl())
	})
	return cache_unsafeFreeze__3249621719
}

var cache_unsafeFreeze__4066092631 gopurs_runtime.Value
var once_unsafeFreeze__4066092631 sync.Once
func Get_unsafeFreeze__4066092631() gopurs_runtime.Value {
	once_unsafeFreeze__4066092631.Do(func() {
		cache_unsafeFreeze__4066092631 = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), pkg_Data_Array_ST.Get_unsafeFreezeImpl())
	})
	return cache_unsafeFreeze__4066092631
}

var cache_unsafeThaw__1529848091 gopurs_runtime.Value
var once_unsafeThaw__1529848091 sync.Once
func Get_unsafeThaw__1529848091() gopurs_runtime.Value {
	once_unsafeThaw__1529848091.Do(func() {
		cache_unsafeThaw__1529848091 = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn1(), pkg_Data_Array_ST.Get_unsafeThawImpl())
	})
	return cache_unsafeThaw__1529848091
}

var cache_withArray__126410905 gopurs_runtime.Value
var once_withArray__126410905 sync.Once
func Get_withArray__126410905() gopurs_runtime.Value {
	once_withArray__126410905.Do(func() {
		cache_withArray__126410905 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_withArray__126410905(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_withArray__126410905
}

var cache_withArray__3965688428 gopurs_runtime.Value
var once_withArray__3965688428 sync.Once
func Get_withArray__3965688428() gopurs_runtime.Value {
	once_withArray__3965688428.Do(func() {
		cache_withArray__3965688428 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_withArray__3965688428(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_withArray__3965688428
}

var cache_withArray__2779921436 gopurs_runtime.Value
var once_withArray__2779921436 sync.Once
func Get_withArray__2779921436() gopurs_runtime.Value {
	once_withArray__2779921436.Do(func() {
		cache_withArray__2779921436 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_withArray__2779921436(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_withArray__2779921436
}

var cache_any__3571149479 gopurs_runtime.Value
var once_any__3571149479 sync.Once
func Get_any__3571149479() gopurs_runtime.Value {
	once_any__3571149479.Do(func() {
		cache_any__3571149479 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_any__3571149479(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_any__3571149479
}

var cache_concatMap__435921434 gopurs_runtime.Value
var once_concatMap__435921434 sync.Once
func Get_concatMap__435921434() gopurs_runtime.Value {
	once_concatMap__435921434.Do(func() {
		cache_concatMap__435921434 = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_concatMap__435921434(b_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(a_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_concatMap__435921434
}

var cache_cons__3485573810 gopurs_runtime.Value
var once_cons__3485573810 sync.Once
func Get_cons__3485573810() gopurs_runtime.Value {
	once_cons__3485573810.Do(func() {
		cache_cons__3485573810 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_cons__3485573810(x_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_cons__3485573810
}

var cache_deleteAt__454851725 gopurs_runtime.Value
var once_deleteAt__454851725 sync.Once
func Get_deleteAt__454851725() gopurs_runtime.Value {
	once_deleteAt__454851725.Do(func() {
		cache_deleteAt__454851725 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_deleteAt__454851725(__local_var_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_deleteAt__454851725
}

var cache_deleteBy__519303411 gopurs_runtime.Value
var once_deleteBy__519303411 sync.Once
func Get_deleteBy__519303411() gopurs_runtime.Value {
	once_deleteBy__519303411.Do(func() {
		cache_deleteBy__519303411 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_deleteBy__519303411(v_0_box, v1_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v2_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_deleteBy__519303411
}

var cache_drop__1426757676 gopurs_runtime.Value
var once_drop__1426757676 sync.Once
func Get_drop__1426757676() gopurs_runtime.Value {
	once_drop__1426757676.Do(func() {
		cache_drop__1426757676 = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_drop__1426757676(n_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_drop__1426757676
}

var cache_elemIndex__33401498 gopurs_runtime.Value
var once_elemIndex__33401498 sync.Once
func Get_elemIndex__33401498() gopurs_runtime.Value {
	once_elemIndex__33401498.Do(func() {
		cache_elemIndex__33401498 = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_elemIndex__33401498(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
})
	})
	return cache_elemIndex__33401498
}

var cache_filter__377906483 gopurs_runtime.Value
var once_filter__377906483 sync.Once
func Get_filter__377906483() gopurs_runtime.Value {
	once_filter__377906483.Do(func() {
		cache_filter__377906483 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_filter__377906483(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_filter__377906483
}

var cache_findIndex__139581937 gopurs_runtime.Value
var once_findIndex__139581937 sync.Once
func Get_findIndex__139581937() gopurs_runtime.Value {
	once_findIndex__139581937.Do(func() {
		cache_findIndex__139581937 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findIndex__139581937(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_findIndex__139581937
}

var cache_findLastIndex__139581937 gopurs_runtime.Value
var once_findLastIndex__139581937 sync.Once
func Get_findLastIndex__139581937() gopurs_runtime.Value {
	once_findLastIndex__139581937.Do(func() {
		cache_findLastIndex__139581937 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findLastIndex__139581937(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_findLastIndex__139581937
}

var cache_foldM__2595407950 gopurs_runtime.Value
var once_foldM__2595407950 sync.Once
func Get_foldM__2595407950() gopurs_runtime.Value {
	once_foldM__2595407950.Do(func() {
		cache_foldM__2595407950 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldM__2595407950(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_foldM__2595407950
}

var cache_foldl__2208423996 gopurs_runtime.Value
var once_foldl__2208423996 sync.Once
func Get_foldl__2208423996() gopurs_runtime.Value {
	once_foldl__2208423996.Do(func() {
		cache_foldl__2208423996 = gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldl")
	})
	return cache_foldl__2208423996
}

var cache_foldl__1469296346 gopurs_runtime.Value
var once_foldl__1469296346 sync.Once
func Get_foldl__1469296346() gopurs_runtime.Value {
	once_foldl__1469296346.Do(func() {
		cache_foldl__1469296346 = gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldl")
	})
	return cache_foldl__1469296346
}

var cache_foldl__1522453594 gopurs_runtime.Value
var once_foldl__1522453594 sync.Once
func Get_foldl__1522453594() gopurs_runtime.Value {
	once_foldl__1522453594.Do(func() {
		cache_foldl__1522453594 = gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldl")
	})
	return cache_foldl__1522453594
}

var cache_foldr__1469296346 gopurs_runtime.Value
var once_foldr__1469296346 sync.Once
func Get_foldr__1469296346() gopurs_runtime.Value {
	once_foldr__1469296346.Do(func() {
		cache_foldr__1469296346 = gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldr")
	})
	return cache_foldr__1469296346
}

var cache_foldr__1916116122 gopurs_runtime.Value
var once_foldr__1916116122 sync.Once
func Get_foldr__1916116122() gopurs_runtime.Value {
	once_foldr__1916116122.Do(func() {
		cache_foldr__1916116122 = gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldr")
	})
	return cache_foldr__1916116122
}

var cache_groupAllBy__1923945894 gopurs_runtime.Value
var once_groupAllBy__1923945894 sync.Once
func Get_groupAllBy__1923945894() gopurs_runtime.Value {
	once_groupAllBy__1923945894.Do(func() {
		cache_groupAllBy__1923945894 = gopurs_runtime.Func(func(cmp_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_groupAllBy__1923945894(cmp_0_box)
})
	})
	return cache_groupAllBy__1923945894
}

var cache_groupBy__693635452 gopurs_runtime.Value
var once_groupBy__693635452 sync.Once
func Get_groupBy__693635452() gopurs_runtime.Value {
	once_groupBy__693635452.Do(func() {
		cache_groupBy__693635452 = gopurs_runtime.Func2(func(op_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_groupBy__693635452(op_0_box, func() []gopurs_runtime.Value {
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
	return cache_groupBy__693635452
}

var cache_head__2056156327 gopurs_runtime.Value
var once_head__2056156327 sync.Once
func Get_head__2056156327() gopurs_runtime.Value {
	once_head__2056156327.Do(func() {
		cache_head__2056156327 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_head__2056156327(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_head__2056156327
}

var cache_head__1956412839 gopurs_runtime.Value
var once_head__1956412839 sync.Once
func Get_head__1956412839() gopurs_runtime.Value {
	once_head__1956412839.Do(func() {
		cache_head__1956412839 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_head__1956412839(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_head__1956412839
}

var cache_index__2196477387 gopurs_runtime.Value
var once_index__2196477387 sync.Once
func Get_index__2196477387() gopurs_runtime.Value {
	once_index__2196477387.Do(func() {
		cache_index__2196477387 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_index__2196477387(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), __local_var_1_box.IntVal))}
})
	})
	return cache_index__2196477387
}

var cache_index__2659850123 gopurs_runtime.Value
var once_index__2659850123 sync.Once
func Get_index__2659850123() gopurs_runtime.Value {
	once_index__2659850123.Do(func() {
		cache_index__2659850123 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_index__2659850123(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), __local_var_1_box.IntVal))}
})
	})
	return cache_index__2659850123
}

var cache_index__3264285643 gopurs_runtime.Value
var once_index__3264285643 sync.Once
func Get_index__3264285643() gopurs_runtime.Value {
	once_index__3264285643.Do(func() {
		cache_index__3264285643 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_index__3264285643(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), __local_var_1_box.IntVal))}
})
	})
	return cache_index__3264285643
}

var cache_init__976795489 gopurs_runtime.Value
var once_init__976795489 sync.Once
func Get_init__976795489() gopurs_runtime.Value {
	once_init__976795489.Do(func() {
		cache_init__976795489 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_init__976795489(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_init__976795489
}

var cache_insertAt__388410084 gopurs_runtime.Value
var once_insertAt__388410084 sync.Once
func Get_insertAt__388410084() gopurs_runtime.Value {
	once_insertAt__388410084.Do(func() {
		cache_insertAt__388410084 = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_insertAt__388410084(__local_var_0_box.IntVal, __local_var_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_insertAt__388410084
}

var cache_insertBy__1563432905 gopurs_runtime.Value
var once_insertBy__1563432905 sync.Once
func Get_insertBy__1563432905() gopurs_runtime.Value {
	once_insertBy__1563432905.Do(func() {
		cache_insertBy__1563432905 = gopurs_runtime.Func3(func(cmp_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_insertBy__1563432905(cmp_0_box, x_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(ys_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_insertBy__1563432905
}

var cache_intersectBy__145374773 gopurs_runtime.Value
var once_intersectBy__145374773 sync.Once
func Get_intersectBy__145374773() gopurs_runtime.Value {
	once_intersectBy__145374773.Do(func() {
		cache_intersectBy__145374773 = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_intersectBy__145374773(eq_0_box, func() []gopurs_runtime.Value {
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
	return cache_intersectBy__145374773
}

var cache_last__2345912124 gopurs_runtime.Value
var once_last__2345912124 sync.Once
func Get_last__2345912124() gopurs_runtime.Value {
	once_last__2345912124.Do(func() {
		cache_last__2345912124 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_last__2345912124(func() []*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value], len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](v) }
					return unboxed
				}()))}
})
	})
	return cache_last__2345912124
}

var cache_last__2056156327 gopurs_runtime.Value
var once_last__2056156327 sync.Once
func Get_last__2056156327() gopurs_runtime.Value {
	once_last__2056156327.Do(func() {
		cache_last__2056156327 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_last__2056156327(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_last__2056156327
}

var cache_many__1839052385 gopurs_runtime.Value
var once_many__1839052385 sync.Once
func Get_many__1839052385() gopurs_runtime.Value {
	once_many__1839052385.Do(func() {
		cache_many__1839052385 = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_many__1839052385(gopurs_runtime.CoerceToStruct[pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]](dictAlternative_0_box))
})
	})
	return cache_many__1839052385
}

var cache_mapMaybe__1271145181 gopurs_runtime.Value
var once_mapMaybe__1271145181 sync.Once
func Get_mapMaybe__1271145181() gopurs_runtime.Value {
	once_mapMaybe__1271145181.Do(func() {
		cache_mapMaybe__1271145181 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe__1271145181(f_0_box)
})
	})
	return cache_mapMaybe__1271145181
}

var cache_mapMaybe__3137412285 gopurs_runtime.Value
var once_mapMaybe__3137412285 sync.Once
func Get_mapMaybe__3137412285() gopurs_runtime.Value {
	once_mapMaybe__3137412285.Do(func() {
		cache_mapMaybe__3137412285 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe__3137412285(f_0_box)
})
	})
	return cache_mapMaybe__3137412285
}

var cache_mapMaybe__2261006141 gopurs_runtime.Value
var once_mapMaybe__2261006141 sync.Once
func Get_mapMaybe__2261006141() gopurs_runtime.Value {
	once_mapMaybe__2261006141.Do(func() {
		cache_mapMaybe__2261006141 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe__2261006141(f_0_box)
})
	})
	return cache_mapMaybe__2261006141
}

var cache_mapWithIndex__3745622640 gopurs_runtime.Value
var once_mapWithIndex__3745622640 sync.Once
func Get_mapWithIndex__3745622640() gopurs_runtime.Value {
	once_mapWithIndex__3745622640.Do(func() {
		cache_mapWithIndex__3745622640 = gopurs_runtime.RecordGet(pkg_Data_FunctorWithIndex.Get_functorWithIndexArray(), "mapWithIndex")
	})
	return cache_mapWithIndex__3745622640
}

var cache_mapWithIndex__1705728720 gopurs_runtime.Value
var once_mapWithIndex__1705728720 sync.Once
func Get_mapWithIndex__1705728720() gopurs_runtime.Value {
	once_mapWithIndex__1705728720.Do(func() {
		cache_mapWithIndex__1705728720 = gopurs_runtime.RecordGet(pkg_Data_FunctorWithIndex.Get_functorWithIndexArray(), "mapWithIndex")
	})
	return cache_mapWithIndex__1705728720
}

var cache_nubBy__3347533344 gopurs_runtime.Value
var once_nubBy__3347533344 sync.Once
func Get_nubBy__3347533344() gopurs_runtime.Value {
	once_nubBy__3347533344.Do(func() {
		cache_nubBy__3347533344 = gopurs_runtime.Func2(func(comp_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_nubBy__3347533344(comp_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_nubBy__3347533344
}

var cache_nubByEq__3443670074 gopurs_runtime.Value
var once_nubByEq__3443670074 sync.Once
func Get_nubByEq__3443670074() gopurs_runtime.Value {
	once_nubByEq__3443670074.Do(func() {
		cache_nubByEq__3443670074 = gopurs_runtime.Func2(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_nubByEq__3443670074(eq_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_nubByEq__3443670074
}

var cache_null__1123412116 gopurs_runtime.Value
var once_null__1123412116 sync.Once
func Get_null__1123412116() gopurs_runtime.Value {
	once_null__1123412116.Do(func() {
		cache_null__1123412116 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_null__1123412116(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_null__1123412116
}

var cache_singleton__193199869 gopurs_runtime.Value
var once_singleton__193199869 sync.Once
func Get_singleton__193199869() gopurs_runtime.Value {
	once_singleton__193199869.Do(func() {
		cache_singleton__193199869 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_singleton__193199869(a_0_box))
})
	})
	return cache_singleton__193199869
}

var cache_singleton__2286220742 gopurs_runtime.Value
var once_singleton__2286220742 sync.Once
func Get_singleton__2286220742() gopurs_runtime.Value {
	once_singleton__2286220742.Do(func() {
		cache_singleton__2286220742 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_singleton__2286220742(a_0_box))
})
	})
	return cache_singleton__2286220742
}

var cache_singleton__2535196422 gopurs_runtime.Value
var once_singleton__2535196422 sync.Once
func Get_singleton__2535196422() gopurs_runtime.Value {
	once_singleton__2535196422.Do(func() {
		cache_singleton__2535196422 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_singleton__2535196422(a_0_box))
})
	})
	return cache_singleton__2535196422
}

var cache_slice__3011328576 gopurs_runtime.Value
var once_slice__3011328576 sync.Once
func Get_slice__3011328576() gopurs_runtime.Value {
	once_slice__3011328576.Do(func() {
		cache_slice__3011328576 = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_slice__3011328576(__local_var_0_box.IntVal, __local_var_1_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_slice__3011328576
}

var cache_snoc__3419689714 gopurs_runtime.Value
var once_snoc__3419689714 sync.Once
func Get_snoc__3419689714() gopurs_runtime.Value {
	once_snoc__3419689714.Do(func() {
		cache_snoc__3419689714 = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_snoc__3419689714(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), x_1_box))
})
	})
	return cache_snoc__3419689714
}

var cache_snoc__3911647657 gopurs_runtime.Value
var once_snoc__3911647657 sync.Once
func Get_snoc__3911647657() gopurs_runtime.Value {
	once_snoc__3911647657.Do(func() {
		cache_snoc__3911647657 = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_snoc__3911647657(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), x_1_box))
})
	})
	return cache_snoc__3911647657
}

var cache_snoc__1505998191 gopurs_runtime.Value
var once_snoc__1505998191 sync.Once
func Get_snoc__1505998191() gopurs_runtime.Value {
	once_snoc__1505998191.Do(func() {
		cache_snoc__1505998191 = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_snoc__1505998191(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), x_1_box))
})
	})
	return cache_snoc__1505998191
}

var cache_some__1839052385 gopurs_runtime.Value
var once_some__1839052385 sync.Once
func Get_some__1839052385() gopurs_runtime.Value {
	once_some__1839052385.Do(func() {
		cache_some__1839052385 = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_some__1839052385(gopurs_runtime.CoerceToStruct[pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]](dictAlternative_0_box))
})
	})
	return cache_some__1839052385
}

var cache_sortBy__3347533344 gopurs_runtime.Value
var once_sortBy__3347533344 sync.Once
func Get_sortBy__3347533344() gopurs_runtime.Value {
	once_sortBy__3347533344.Do(func() {
		cache_sortBy__3347533344 = gopurs_runtime.Func2(func(comp_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_sortBy__3347533344(comp_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_sortBy__3347533344
}

var cache_sortBy__1848798496 gopurs_runtime.Value
var once_sortBy__1848798496 sync.Once
func Get_sortBy__1848798496() gopurs_runtime.Value {
	once_sortBy__1848798496.Do(func() {
		cache_sortBy__1848798496 = gopurs_runtime.Func2(func(comp_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_sortBy__1848798496(comp_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_sortBy__1848798496
}

var cache_sortWith__1917042304 gopurs_runtime.Value
var once_sortWith__1917042304 sync.Once
func Get_sortWith__1917042304() gopurs_runtime.Value {
	once_sortWith__1917042304.Do(func() {
		cache_sortWith__1917042304 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sortWith__1917042304(f_0_box)
})
	})
	return cache_sortWith__1917042304
}

var cache_sortWith__3115909389 gopurs_runtime.Value
var once_sortWith__3115909389 sync.Once
func Get_sortWith__3115909389() gopurs_runtime.Value {
	once_sortWith__3115909389.Do(func() {
		cache_sortWith__3115909389 = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sortWith__3115909389(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), f_1_box)
})
	})
	return cache_sortWith__3115909389
}

var cache_span__174751768 gopurs_runtime.Value
var once_span__174751768 sync.Once
func Get_span__174751768() gopurs_runtime.Value {
	once_span__174751768.Do(func() {
		cache_span__174751768 = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, arr_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_span__174751768(p_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(arr_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_span__174751768
}

var cache_take__1426757676 gopurs_runtime.Value
var once_take__1426757676 sync.Once
func Get_take__1426757676() gopurs_runtime.Value {
	once_take__1426757676.Do(func() {
		cache_take__1426757676 = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_take__1426757676(n_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_take__1426757676
}

var cache_unionBy__145374773 gopurs_runtime.Value
var once_unionBy__145374773 sync.Once
func Get_unionBy__145374773() gopurs_runtime.Value {
	once_unionBy__145374773.Do(func() {
		cache_unionBy__145374773 = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_unionBy__145374773(eq_0_box, func() []gopurs_runtime.Value {
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
	return cache_unionBy__145374773
}

var cache_unsafeIndex__2808089623 gopurs_runtime.Value
var once_unsafeIndex__2808089623 sync.Once
func Get_unsafeIndex__2808089623() gopurs_runtime.Value {
	once_unsafeIndex__2808089623.Do(func() {
		cache_unsafeIndex__2808089623 = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeIndex__2808089623(_dollar__unused_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), __local_var_2_box.IntVal)
})
	})
	return cache_unsafeIndex__2808089623
}

var cache_updateAt__388410084 gopurs_runtime.Value
var once_updateAt__388410084 sync.Once
func Get_updateAt__388410084() gopurs_runtime.Value {
	once_updateAt__388410084.Do(func() {
		cache_updateAt__388410084 = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_updateAt__388410084(__local_var_0_box.IntVal, __local_var_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_updateAt__388410084
}

var cache_zipWith__1350747206 gopurs_runtime.Value
var once_zipWith__1350747206 sync.Once
func Get_zipWith__1350747206() gopurs_runtime.Value {
	once_zipWith__1350747206.Do(func() {
		cache_zipWith__1350747206 = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_zipWith__1350747206(__local_var_0_box, func() []gopurs_runtime.Value {
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
	return cache_zipWith__1350747206
}

var cache_zipWith__1220584870 gopurs_runtime.Value
var once_zipWith__1220584870 sync.Once
func Get_zipWith__1220584870() gopurs_runtime.Value {
	once_zipWith__1220584870.Do(func() {
		cache_zipWith__1220584870 = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_zipWith__1220584870(__local_var_0_box, func() []gopurs_runtime.Value {
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
	return cache_zipWith__1220584870
}

var cache_zipWith__2000246342 gopurs_runtime.Value
var once_zipWith__2000246342 sync.Once
func Get_zipWith__2000246342() gopurs_runtime.Value {
	once_zipWith__2000246342.Do(func() {
		cache_zipWith__2000246342 = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_zipWith__2000246342(__local_var_0_box, func() []gopurs_runtime.Value {
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
	return cache_zipWith__2000246342
}

var cache_eq__2843686287 gopurs_runtime.Value
var once_eq__2843686287 sync.Once
func Get_eq__2843686287() gopurs_runtime.Value {
	once_eq__2843686287.Do(func() {
		cache_eq__2843686287 = pkg_Data_Eq.Get_eqIntImpl()
	})
	return cache_eq__2843686287
}

var cache_eq__2276491096 gopurs_runtime.Value
var once_eq__2276491096 sync.Once
func Get_eq__2276491096() gopurs_runtime.Value {
	once_eq__2276491096.Do(func() {
		cache_eq__2276491096 = gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq")
	})
	return cache_eq__2276491096
}

var cache_eq__2384498378 gopurs_runtime.Value
var once_eq__2384498378 sync.Once
func Get_eq__2384498378() gopurs_runtime.Value {
	once_eq__2384498378.Do(func() {
		cache_eq__2384498378 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__2384498378
}

var cache_eq__1272715810 gopurs_runtime.Value
var once_eq__1272715810 sync.Once
func Get_eq__1272715810() gopurs_runtime.Value {
	once_eq__1272715810.Do(func() {
		cache_eq__1272715810 = gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq")
	})
	return cache_eq__1272715810
}

var cache_notEq__2384498378 gopurs_runtime.Value
var once_notEq__2384498378 sync.Once
func Get_notEq__2384498378() gopurs_runtime.Value {
	once_notEq__2384498378.Do(func() {
		cache_notEq__2384498378 = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_notEq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box, y_2_box))
})
	})
	return cache_notEq__2384498378
}

var cache_notEq__1272715810 gopurs_runtime.Value
var once_notEq__1272715810 sync.Once
func Get_notEq__1272715810() gopurs_runtime.Value {
	once_notEq__1272715810.Do(func() {
		cache_notEq__1272715810 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_notEq__1272715810(x_0_box, y_1_box))
})
	})
	return cache_notEq__1272715810
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

var cache_foldr__3591001499 gopurs_runtime.Value
var once_foldr__3591001499 sync.Once
func Get_foldr__3591001499() gopurs_runtime.Value {
	once_foldr__3591001499.Do(func() {
		cache_foldr__3591001499 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__3591001499(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__3591001499
}

var cache_traverse___996968168 gopurs_runtime.Value
var once_traverse___996968168 sync.Once
func Get_traverse___996968168() gopurs_runtime.Value {
	once_traverse___996968168.Do(func() {
		cache_traverse___996968168 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse___996968168(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_traverse___996968168
}

var cache_traverse___1229293625 gopurs_runtime.Value
var once_traverse___1229293625 sync.Once
func Get_traverse___1229293625() gopurs_runtime.Value {
	once_traverse___1229293625.Do(func() {
		cache_traverse___1229293625 = func() gopurs_runtime.Value {
__local_var_0_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "Apply0"), gopurs_runtime.Value{})
_ = __local_var_0_1
Functor0_1_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_2
applySecond_0_0 := gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_1, "apply"), gopurs_runtime.Apply2(Functor0_1_2.V0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return x_5
})
}), a_2), b_3)
})
})
_ = applySecond_0_0
return gopurs_runtime.Func(func(dictFoldable_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_1, "foldr"), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_0_0, gopurs_runtime.Apply(f_2, x_3))
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "pure"), pkg_Data_Unit.Get_unit()))
})
})
}()
	})
	return cache_traverse___1229293625
}

var cache_traverse___3202958137 gopurs_runtime.Value
var once_traverse___3202958137 sync.Once
func Get_traverse___3202958137() gopurs_runtime.Value {
	once_traverse___3202958137.Do(func() {
		cache_traverse___3202958137 = func() gopurs_runtime.Value {
__local_var_0_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "Apply0"), gopurs_runtime.Value{})
_ = __local_var_0_1
Functor0_1_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_0_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_2
applySecond_0_0 := gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_1, "apply"), gopurs_runtime.Apply2(Functor0_1_2.V0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return x_5
})
}), a_2), b_3)
})
})
_ = applySecond_0_0
return gopurs_runtime.Func(func(dictFoldable_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_1, "foldr"), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_0_0, gopurs_runtime.Apply(f_2, x_3))
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "pure"), pkg_Data_Unit.Get_unit()))
})
})
}()
	})
	return cache_traverse___3202958137
}

var cache_const__1243414737 gopurs_runtime.Value
var once_const__1243414737 sync.Once
func Get_const__1243414737() gopurs_runtime.Value {
	once_const__1243414737.Do(func() {
		cache_const__1243414737 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__1243414737(a_0_box, v_1_box)
})
	})
	return cache_const__1243414737
}

var cache_const__2082174484 gopurs_runtime.Value
var once_const__2082174484 sync.Once
func Get_const__2082174484() gopurs_runtime.Value {
	once_const__2082174484.Do(func() {
		cache_const__2082174484 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__2082174484(a_0_box, v_1_box)
})
	})
	return cache_const__2082174484
}

var cache_const__220790420 gopurs_runtime.Value
var once_const__220790420 sync.Once
func Get_const__220790420() gopurs_runtime.Value {
	once_const__220790420.Do(func() {
		cache_const__220790420 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__220790420(a_0_box, v_1_box)
})
	})
	return cache_const__220790420
}

var cache_const__4026847508 gopurs_runtime.Value
var once_const__4026847508 sync.Once
func Get_const__4026847508() gopurs_runtime.Value {
	once_const__4026847508.Do(func() {
		cache_const__4026847508 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__4026847508(a_0_box, v_1_box)
})
	})
	return cache_const__4026847508
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

var cache_const__153462564 gopurs_runtime.Value
var once_const__153462564 sync.Once
func Get_const__153462564() gopurs_runtime.Value {
	once_const__153462564.Do(func() {
		cache_const__153462564 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__153462564(a_0_box, v_1_box)
})
	})
	return cache_const__153462564
}

var cache_const__106415396 gopurs_runtime.Value
var once_const__106415396 sync.Once
func Get_const__106415396() gopurs_runtime.Value {
	once_const__106415396.Do(func() {
		cache_const__106415396 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__106415396(a_0_box, v_1_box)
})
	})
	return cache_const__106415396
}

var cache_flip__3525081280 gopurs_runtime.Value
var once_flip__3525081280 sync.Once
func Get_flip__3525081280() gopurs_runtime.Value {
	once_flip__3525081280.Do(func() {
		cache_flip__3525081280 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3525081280(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3525081280
}

var cache_flip__3021024128 gopurs_runtime.Value
var once_flip__3021024128 sync.Once
func Get_flip__3021024128() gopurs_runtime.Value {
	once_flip__3021024128.Do(func() {
		cache_flip__3021024128 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3021024128(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3021024128
}

var cache_flip__2857929472 gopurs_runtime.Value
var once_flip__2857929472 sync.Once
func Get_flip__2857929472() gopurs_runtime.Value {
	once_flip__2857929472.Do(func() {
		cache_flip__2857929472 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__2857929472(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__2857929472
}

var cache_flip__3136194560 gopurs_runtime.Value
var once_flip__3136194560 sync.Once
func Get_flip__3136194560() gopurs_runtime.Value {
	once_flip__3136194560.Do(func() {
		cache_flip__3136194560 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3136194560(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3136194560
}

var cache_flip__3772615392 gopurs_runtime.Value
var once_flip__3772615392 sync.Once
func Get_flip__3772615392() gopurs_runtime.Value {
	once_flip__3772615392.Do(func() {
		cache_flip__3772615392 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3772615392(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3772615392
}

var cache_functorArray__2747750794 gopurs_runtime.Value
var once_functorArray__2747750794 sync.Once
func Get_functorArray__2747750794() gopurs_runtime.Value {
	once_functorArray__2747750794.Do(func() {
		cache_functorArray__2747750794 = gopurs_runtime.RecordDict1("map", pkg_Data_Functor.Get_arrayMap())
	})
	return cache_functorArray__2747750794
}

var cache_functorArray__361387505 gopurs_runtime.Value
var once_functorArray__361387505 sync.Once
func Get_functorArray__361387505() gopurs_runtime.Value {
	once_functorArray__361387505.Do(func() {
		cache_functorArray__361387505 = gopurs_runtime.RecordDict1("map", pkg_Data_Functor.Get_arrayMap())
	})
	return cache_functorArray__361387505
}

var cache_map__2475811444 gopurs_runtime.Value
var once_map__2475811444 sync.Once
func Get_map__2475811444() gopurs_runtime.Value {
	once_map__2475811444.Do(func() {
		cache_map__2475811444 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2475811444(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2475811444
}

var cache_map__3240628980 gopurs_runtime.Value
var once_map__3240628980 sync.Once
func Get_map__3240628980() gopurs_runtime.Value {
	once_map__3240628980.Do(func() {
		cache_map__3240628980 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3240628980(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__3240628980
}

var cache_map__2199395572 gopurs_runtime.Value
var once_map__2199395572 sync.Once
func Get_map__2199395572() gopurs_runtime.Value {
	once_map__2199395572.Do(func() {
		cache_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2199395572(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2199395572
}

var cache_map__1319384564 gopurs_runtime.Value
var once_map__1319384564 sync.Once
func Get_map__1319384564() gopurs_runtime.Value {
	once_map__1319384564.Do(func() {
		cache_map__1319384564 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1319384564(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1319384564
}

var cache_map__1852651540 gopurs_runtime.Value
var once_map__1852651540 sync.Once
func Get_map__1852651540() gopurs_runtime.Value {
	once_map__1852651540.Do(func() {
		cache_map__1852651540 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1852651540(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1852651540
}

var cache_map__627657844 gopurs_runtime.Value
var once_map__627657844 sync.Once
func Get_map__627657844() gopurs_runtime.Value {
	once_map__627657844.Do(func() {
		cache_map__627657844 = gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorArray(), "map")
	})
	return cache_map__627657844
}

var cache_map__2174973445 gopurs_runtime.Value
var once_map__2174973445 sync.Once
func Get_map__2174973445() gopurs_runtime.Value {
	once_map__2174973445.Do(func() {
		cache_map__2174973445 = gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map")
	})
	return cache_map__2174973445
}

var cache_map__2107538812 gopurs_runtime.Value
var once_map__2107538812 sync.Once
func Get_map__2107538812() gopurs_runtime.Value {
	once_map__2107538812.Do(func() {
		cache_map__2107538812 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__2107538812
}

var cache_map__901270812 gopurs_runtime.Value
var once_map__901270812 sync.Once
func Get_map__901270812() gopurs_runtime.Value {
	once_map__901270812.Do(func() {
		cache_map__901270812 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__901270812
}

var cache_map__1171574364 gopurs_runtime.Value
var once_map__1171574364 sync.Once
func Get_map__1171574364() gopurs_runtime.Value {
	once_map__1171574364.Do(func() {
		cache_map__1171574364 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__1171574364
}

var cache_void__3020373336 gopurs_runtime.Value
var once_void__3020373336 sync.Once
func Get_void__3020373336() gopurs_runtime.Value {
	once_void__3020373336.Do(func() {
		cache_void__3020373336 = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_void__3020373336(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box))
})
	})
	return cache_void__3020373336
}

var cache_void__2104786761 gopurs_runtime.Value
var once_void__2104786761 sync.Once
func Get_void__2104786761() gopurs_runtime.Value {
	once_void__2104786761.Do(func() {
		cache_void__2104786761 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
	})
	return cache_void__2104786761
}

var cache_disj__3676519832 gopurs_runtime.Value
var once_disj__3676519832 sync.Once
func Get_disj__3676519832() gopurs_runtime.Value {
	once_disj__3676519832.Do(func() {
		cache_disj__3676519832 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj")
	})
	return cache_disj__3676519832
}

var cache_disj__3472268504 gopurs_runtime.Value
var once_disj__3472268504 sync.Once
func Get_disj__3472268504() gopurs_runtime.Value {
	once_disj__3472268504.Do(func() {
		cache_disj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_disj__3472268504
}

var cache_not__3201284355 gopurs_runtime.Value
var once_not__3201284355 sync.Once
func Get_not__3201284355() gopurs_runtime.Value {
	once_not__3201284355.Do(func() {
		cache_not__3201284355 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not")
	})
	return cache_not__3201284355
}

var cache_not__1505204753 gopurs_runtime.Value
var once_not__1505204753 sync.Once
func Get_not__1505204753() gopurs_runtime.Value {
	once_not__1505204753.Do(func() {
		cache_not__1505204753 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__1505204753(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_not__1505204753
}

var cache_applyMaybe__3561700045 gopurs_runtime.Value
var once_applyMaybe__3561700045 sync.Once
func Get_applyMaybe__3561700045() gopurs_runtime.Value {
	once_applyMaybe__3561700045.Do(func() {
		cache_applyMaybe__3561700045 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v1_1))})))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_applyMaybe__3561700045
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

var cache_fromJust__1478027324 gopurs_runtime.Value
var once_fromJust__1478027324 sync.Once
func Get_fromJust__1478027324() gopurs_runtime.Value {
	once_fromJust__1478027324.Do(func() {
		cache_fromJust__1478027324 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust__1478027324(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_fromJust__1478027324
}

var cache_functorMaybe__2569569018 gopurs_runtime.Value
var once_functorMaybe__2569569018 sync.Once
func Get_functorMaybe__2569569018() gopurs_runtime.Value {
	once_functorMaybe__2569569018.Do(func() {
		cache_functorMaybe__2569569018 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_functorMaybe__2569569018
}

var cache_functorMaybe__2097654001 gopurs_runtime.Value
var once_functorMaybe__2097654001 sync.Once
func Get_functorMaybe__2097654001() gopurs_runtime.Value {
	once_functorMaybe__2097654001.Do(func() {
		cache_functorMaybe__2097654001 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_functorMaybe__2097654001
}

var cache_isJust__2514352589 gopurs_runtime.Value
var once_isJust__2514352589 sync.Once
func Get_isJust__2514352589() gopurs_runtime.Value {
	once_isJust__2514352589.Do(func() {
		cache_isJust__2514352589 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isJust__2514352589(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](v2_0_box)))
})
	})
	return cache_isJust__2514352589
}

var cache_isJust__2591355336 gopurs_runtime.Value
var once_isJust__2591355336 sync.Once
func Get_isJust__2591355336() gopurs_runtime.Value {
	once_isJust__2591355336.Do(func() {
		cache_isJust__2591355336 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isJust__2591355336(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isJust__2591355336
}

var cache_isJust__1358705270 gopurs_runtime.Value
var once_isJust__1358705270 sync.Once
func Get_isJust__1358705270() gopurs_runtime.Value {
	once_isJust__1358705270.Do(func() {
		cache_isJust__1358705270 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isJust__1358705270(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isJust__1358705270
}

var cache_isJust__4206805139 gopurs_runtime.Value
var once_isJust__4206805139 sync.Once
func Get_isJust__4206805139() gopurs_runtime.Value {
	once_isJust__4206805139.Do(func() {
		cache_isJust__4206805139 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isJust__4206805139(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isJust__4206805139
}

var cache_isNothing__2514352589 gopurs_runtime.Value
var once_isNothing__2514352589 sync.Once
func Get_isNothing__2514352589() gopurs_runtime.Value {
	once_isNothing__2514352589.Do(func() {
		cache_isNothing__2514352589 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isNothing__2514352589(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](v2_0_box)))
})
	})
	return cache_isNothing__2514352589
}

var cache_maybe__919206801 gopurs_runtime.Value
var once_maybe__919206801 sync.Once
func Get_maybe__919206801() gopurs_runtime.Value {
	once_maybe__919206801.Do(func() {
		cache_maybe__919206801 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__919206801(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__919206801
}

var cache_maybe__3078346790 gopurs_runtime.Value
var once_maybe__3078346790 sync.Once
func Get_maybe__3078346790() gopurs_runtime.Value {
	once_maybe__3078346790.Do(func() {
		cache_maybe__3078346790 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__3078346790(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__3078346790
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

var cache_maybe__3718989812 gopurs_runtime.Value
var once_maybe__3718989812 sync.Once
func Get_maybe__3718989812() gopurs_runtime.Value {
	once_maybe__3718989812.Do(func() {
		cache_maybe__3718989812 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__3718989812(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__3718989812
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

var cache_maybe__2925953714 gopurs_runtime.Value
var once_maybe__2925953714 sync.Once
func Get_maybe__2925953714() gopurs_runtime.Value {
	once_maybe__2925953714.Do(func() {
		cache_maybe__2925953714 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__2925953714(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__2925953714
}

var cache_maybe__1408653394 gopurs_runtime.Value
var once_maybe__1408653394 sync.Once
func Get_maybe__1408653394() gopurs_runtime.Value {
	once_maybe__1408653394.Do(func() {
		cache_maybe__1408653394 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__1408653394(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__1408653394
}

var cache_maybe__727024722 gopurs_runtime.Value
var once_maybe__727024722 sync.Once
func Get_maybe__727024722() gopurs_runtime.Value {
	once_maybe__727024722.Do(func() {
		cache_maybe__727024722 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__727024722(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__727024722
}

var cache_maybe__1936890643 gopurs_runtime.Value
var once_maybe__1936890643 sync.Once
func Get_maybe__1936890643() gopurs_runtime.Value {
	once_maybe__1936890643.Do(func() {
		cache_maybe__1936890643 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__1936890643(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__1936890643
}

var cache_compare__821463600 gopurs_runtime.Value
var once_compare__821463600 sync.Once
func Get_compare__821463600() gopurs_runtime.Value {
	once_compare__821463600.Do(func() {
		cache_compare__821463600 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__821463600(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__821463600
}

var cache_comparing__3506074860 gopurs_runtime.Value
var once_comparing__3506074860 sync.Once
func Get_comparing__3506074860() gopurs_runtime.Value {
	once_comparing__3506074860.Do(func() {
		cache_comparing__3506074860 = gopurs_runtime.Func4(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_comparing__3506074860(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), f_1_box, x_2_box, y_3_box)), UnsafePtr: nil}
})
	})
	return cache_comparing__3506074860
}

var cache_greaterThanOrEq__4087042607 gopurs_runtime.Value
var once_greaterThanOrEq__4087042607 sync.Once
func Get_greaterThanOrEq__4087042607() gopurs_runtime.Value {
	once_greaterThanOrEq__4087042607.Do(func() {
		cache_greaterThanOrEq__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_greaterThanOrEq__4087042607
}

var cache_greaterThanOrEq__1409282474 gopurs_runtime.Value
var once_greaterThanOrEq__1409282474 sync.Once
func Get_greaterThanOrEq__1409282474() gopurs_runtime.Value {
	once_greaterThanOrEq__1409282474.Do(func() {
		cache_greaterThanOrEq__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThanOrEq__1409282474
}

var cache_lessThan__4087042607 gopurs_runtime.Value
var once_lessThan__4087042607 sync.Once
func Get_lessThan__4087042607() gopurs_runtime.Value {
	once_lessThan__4087042607.Do(func() {
		cache_lessThan__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThan__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_lessThan__4087042607
}

var cache_lessThan__1409282474 gopurs_runtime.Value
var once_lessThan__1409282474 sync.Once
func Get_lessThan__1409282474() gopurs_runtime.Value {
	once_lessThan__1409282474.Do(func() {
		cache_lessThan__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThan__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThan__1409282474
}

var cache_lessThanOrEq__4087042607 gopurs_runtime.Value
var once_lessThanOrEq__4087042607 sync.Once
func Get_lessThanOrEq__4087042607() gopurs_runtime.Value {
	once_lessThanOrEq__4087042607.Do(func() {
		cache_lessThanOrEq__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_lessThanOrEq__4087042607
}

var cache_lessThanOrEq__1409282474 gopurs_runtime.Value
var once_lessThanOrEq__1409282474 sync.Once
func Get_lessThanOrEq__1409282474() gopurs_runtime.Value {
	once_lessThanOrEq__1409282474.Do(func() {
		cache_lessThanOrEq__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThanOrEq__1409282474
}

var cache_negate__2635823316 gopurs_runtime.Value
var once_negate__2635823316 sync.Once
func Get_negate__2635823316() gopurs_runtime.Value {
	once_negate__2635823316.Do(func() {
		cache_negate__2635823316 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_negate__2635823316(a_0_box)
})
	})
	return cache_negate__2635823316
}

var cache_negate__1364373265 gopurs_runtime.Value
var once_negate__1364373265 sync.Once
func Get_negate__1364373265() gopurs_runtime.Value {
	once_negate__1364373265.Do(func() {
		cache_negate__1364373265 = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_negate__1364373265(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dictRing_0_box))
})
	})
	return cache_negate__1364373265
}

var cache_sub__1043827704 gopurs_runtime.Value
var once_sub__1043827704 sync.Once
func Get_sub__1043827704() gopurs_runtime.Value {
	once_sub__1043827704.Do(func() {
		cache_sub__1043827704 = pkg_Data_Ring.Get_intSub()
	})
	return cache_sub__1043827704
}

var cache_sub__3675938712 gopurs_runtime.Value
var once_sub__3675938712 sync.Once
func Get_sub__3675938712() gopurs_runtime.Value {
	once_sub__3675938712.Do(func() {
		cache_sub__3675938712 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__3675938712(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__3675938712
}

var cache_sub__2345699288 gopurs_runtime.Value
var once_sub__2345699288 sync.Once
func Get_sub__2345699288() gopurs_runtime.Value {
	once_sub__2345699288.Do(func() {
		cache_sub__2345699288 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__2345699288(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_sub__2345699288
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

var cache_append__365446200 gopurs_runtime.Value
var once_append__365446200 sync.Once
func Get_append__365446200() gopurs_runtime.Value {
	once_append__365446200.Do(func() {
		cache_append__365446200 = gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append")
	})
	return cache_append__365446200
}

var cache_semigroupArray__777842900 gopurs_runtime.Value
var once_semigroupArray__777842900 sync.Once
func Get_semigroupArray__777842900() gopurs_runtime.Value {
	once_semigroupArray__777842900.Do(func() {
		cache_semigroupArray__777842900 = gopurs_runtime.RecordDict1("append", pkg_Data_Semigroup.Get_concatArray())
	})
	return cache_semigroupArray__777842900
}

var cache_semigroupArray__3208067858 gopurs_runtime.Value
var once_semigroupArray__3208067858 sync.Once
func Get_semigroupArray__3208067858() gopurs_runtime.Value {
	once_semigroupArray__3208067858.Do(func() {
		cache_semigroupArray__3208067858 = gopurs_runtime.RecordDict1("append", pkg_Data_Semigroup.Get_concatArray())
	})
	return cache_semigroupArray__3208067858
}

var cache_add__560788792 gopurs_runtime.Value
var once_add__560788792 sync.Once
func Get_add__560788792() gopurs_runtime.Value {
	once_add__560788792.Do(func() {
		cache_add__560788792 = pkg_Data_Semiring.Get_intAdd()
	})
	return cache_add__560788792
}

var cache_add__1614463960 gopurs_runtime.Value
var once_add__1614463960 sync.Once
func Get_add__1614463960() gopurs_runtime.Value {
	once_add__1614463960.Do(func() {
		cache_add__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_add__1614463960
}

var cache_zero__1204848985 gopurs_runtime.Value
var once_zero__1204848985 sync.Once
func Get_zero__1204848985() gopurs_runtime.Value {
	once_zero__1204848985.Do(func() {
		cache_zero__1204848985 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zero__1204848985(dict_0_box)
})
	})
	return cache_zero__1204848985
}

var cache_sequence__1886310617 gopurs_runtime.Value
var once_sequence__1886310617 sync.Once
func Get_sequence__1886310617() gopurs_runtime.Value {
	once_sequence__1886310617.Do(func() {
		cache_sequence__1886310617 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequence__1886310617(gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sequence__1886310617
}

var cache_traversableArray__2090378122 gopurs_runtime.Value
var once_traversableArray__2090378122 sync.Once
func Get_traversableArray__2090378122() gopurs_runtime.Value {
	once_traversableArray__2090378122.Do(func() {
		cache_traversableArray__2090378122 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_traversableArray__2090378122
}

var cache_traverse__314957093 gopurs_runtime.Value
var once_traverse__314957093 sync.Once
func Get_traverse__314957093() gopurs_runtime.Value {
	once_traverse__314957093.Do(func() {
		cache_traverse__314957093 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse__314957093(gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_traverse__314957093
}

var cache_traverse__303957893 gopurs_runtime.Value
var once_traverse__303957893 sync.Once
func Get_traverse__303957893() gopurs_runtime.Value {
	once_traverse__303957893.Do(func() {
		cache_traverse__303957893 = gopurs_runtime.RecordGet(pkg_Data_Traversable.Get_traversableArray(), "traverse")
	})
	return cache_traverse__303957893
}

var cache_fst__3478736243 gopurs_runtime.Value
var once_fst__3478736243 sync.Once
func Get_fst__3478736243() gopurs_runtime.Value {
	once_fst__3478736243.Do(func() {
		cache_fst__3478736243 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_fst__3478736243(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](v_0_box)))
})
	})
	return cache_fst__3478736243
}

var cache_snd__1234761462 gopurs_runtime.Value
var once_snd__1234761462 sync.Once
func Get_snd__1234761462() gopurs_runtime.Value {
	once_snd__1234761462.Do(func() {
		cache_snd__1234761462 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_snd__1234761462(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_snd__1234761462
}

var cache_snd__1694562576 gopurs_runtime.Value
var once_snd__1694562576 sync.Once
func Get_snd__1694562576() gopurs_runtime.Value {
	once_snd__1694562576.Do(func() {
		cache_snd__1694562576 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_snd__1694562576(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_snd__1694562576
}

var cache_snd__395594805 gopurs_runtime.Value
var once_snd__395594805 sync.Once
func Get_snd__395594805() gopurs_runtime.Value {
	once_snd__395594805.Do(func() {
		cache_snd__395594805 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_snd__395594805(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_snd__395594805
}

var cache_unfoldr__1786322085 gopurs_runtime.Value
var once_unfoldr__1786322085 sync.Once
func Get_unfoldr__1786322085() gopurs_runtime.Value {
	once_unfoldr__1786322085.Do(func() {
		cache_unfoldr__1786322085 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr__1786322085(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr__1786322085
}

var cache_unfoldr__1128708256 gopurs_runtime.Value
var once_unfoldr__1128708256 sync.Once
func Get_unfoldr__1128708256() gopurs_runtime.Value {
	once_unfoldr__1128708256.Do(func() {
		cache_unfoldr__1128708256 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr__1128708256(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr__1128708256
}

var cache_unsafePartial__3178441094 gopurs_runtime.Value
var once_unsafePartial__3178441094 sync.Once
func Get_unsafePartial__3178441094() gopurs_runtime.Value {
	once_unsafePartial__3178441094.Do(func() {
		cache_unsafePartial__3178441094 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__3178441094
}

var cache_unsafePartial__1306634845 gopurs_runtime.Value
var once_unsafePartial__1306634845 sync.Once
func Get_unsafePartial__1306634845() gopurs_runtime.Value {
	once_unsafePartial__1306634845.Do(func() {
		cache_unsafePartial__1306634845 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1306634845
}

var cache_unsafePartial__1186189469 gopurs_runtime.Value
var once_unsafePartial__1186189469 sync.Once
func Get_unsafePartial__1186189469() gopurs_runtime.Value {
	once_unsafePartial__1186189469.Do(func() {
		cache_unsafePartial__1186189469 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1186189469
}

var cache_unsafePartial__823750045 gopurs_runtime.Value
var once_unsafePartial__823750045 sync.Once
func Get_unsafePartial__823750045() gopurs_runtime.Value {
	once_unsafePartial__823750045.Do(func() {
		cache_unsafePartial__823750045 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__823750045
}

var cache_unsafePartial__1709812605 gopurs_runtime.Value
var once_unsafePartial__1709812605 sync.Once
func Get_unsafePartial__1709812605() gopurs_runtime.Value {
	once_unsafePartial__1709812605.Do(func() {
		cache_unsafePartial__1709812605 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1709812605
}

func Call_intercalate1(dictMonoid_0_loop *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonoid_0 *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value] = dictMonoid_0_loop
_ = dictMonoid_0
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonoid_0.V0, gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := dictMonoid_0.V1
_ = mempty_2_1
return gopurs_runtime.Func(func(sep_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldl"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t2 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.RecordGet(v_5, "acc"), gopurs_runtime.Apply2(Semigroup0_1_0.V0, sep_3, v1_6)), gopurs_runtime.Bool(false))
}
end_branch_2:
return __t2
})
}), gopurs_runtime.RecordDict2("acc", "init", mempty_2_1, gopurs_runtime.Bool(true)), xs_4), "acc")
})
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

func Call_zipWithA(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, xs_2_loop []gopurs_runtime.Value, ys_3_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 []gopurs_runtime.Value = xs_2_loop
_ = xs_2
var ys_3 []gopurs_runtime.Value = ys_3_loop
_ = ys_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Traversable.Get_traversableArray(), "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_zipWithImpl(), f_1, gopurs_runtime.Array(xs_2), gopurs_runtime.Array(ys_3)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
}

func Call_updateAtIndices(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], us_1_loop gopurs_runtime.Value, xs_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var us_1 gopurs_runtime.Value = us_1_loop
_ = us_1
var xs_2 []gopurs_runtime.Value = xs_2_loop
_ = xs_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_withArray(), gopurs_runtime.Func(func(res_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(pkg_Data_Foldable.Get_traverse_(), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](pkg_Control_Monad_ST_Internal.Get_applicativeST()))}, gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_0)}, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(pkg_Data_Array_ST.Get_pokeImpl(), gopurs_runtime.Int((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0.IntVal), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, res_3)
})
}), us_1)
}), gopurs_runtime.Array(xs_2))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_updateAt(__local_var_0_loop int64, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value] {
var __local_var_0 int64 = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value]](gopurs_runtime.UncurriedApp5(Get__updateAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Int(__local_var_0), __local_var_1, gopurs_runtime.Array(__local_var_2)))
}

func Call_unsafeIndex(_dollar__unused_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value, __local_var_2_loop int64) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 int64 = __local_var_2_loop
_ = __local_var_2
return __local_var_1[__local_var_2]
}

func Call_uncons(__local_var_0_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_unconsImpl(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("head", "tail", x_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))})}
})
}), gopurs_runtime.Array(__local_var_0)))
}

func Call_toUnfoldable(dictUnfoldable_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value], xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
len_2_0 := gopurs_runtime.Int(int64(len(xs_1))).IntVal
_ = len_2_0
return gopurs_runtime.Apply2(dictUnfoldable_0.V1, gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(gopurs_runtime.Int(i_3.IntVal), gopurs_runtime.Int(len_2_0))).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return xs_1[i_3.IntVal]
})), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int(i_3.IntVal), gopurs_runtime.Int(1)).IntVal)})}})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, int64]]](__t1))}
}), gopurs_runtime.Int(0))
}

func Call_tail(__local_var_0_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value] {
var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_unconsImpl(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})}
})
}), gopurs_runtime.Array(__local_var_0)))
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
if (uint32(v_2.IntVal) == 380165415) {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
if (uint32(v_2.IntVal) == 902936544) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (uint32(v_2.IntVal) == 1527465420) {
__t0 = gopurs_runtime.Int(Call_negate__2635823316(gopurs_runtime.Int(1)).IntVal)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Int(__t0.IntVal)
}), gopurs_runtime.Array(__local_var_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_sortWith(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply(Get_sortBy(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(dictOrd_0.V1, gopurs_runtime.Apply(f_1, x_2), gopurs_runtime.Apply(f_1, y_3)).IntVal)), UnsafePtr: nil}
})
}))
}

func Call_sort(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.V1
_ = compare_1_0
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(Call_sortBy(compare_1_0, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
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
if (gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(gopurs_runtime.Int(v_0), gopurs_runtime.Int(0))).IntVal) != (0) {
__t0 = gopurs_runtime.RecordDict2("after", "before", gopurs_runtime.Array(v1_1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict2("after", "before", gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_sliceImpl(), gopurs_runtime.Int(v_0), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(v1_1))).IntVal), gopurs_runtime.Array(v1_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int(v_0), gopurs_runtime.Array(v1_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
}
end_branch_0:
return __t0
}

func Call_take(n_0_loop int64, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var __t0 []gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(gopurs_runtime.Int(n_0), gopurs_runtime.Int(1))).IntVal) != (0) {
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
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int(n_0), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_0:
return __t0
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

func Call_go__range(__local_var_0_loop int64, __local_var_1_loop int64) []int64 {
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
return (gopurs_runtime.Apply2(Get_eq__2843686287(), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(xs_0))).IntVal), gopurs_runtime.Int(0)).IntVal) != (0)
}

func Call_modifyAtIndices(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], is_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value, xs_3_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var is_1 gopurs_runtime.Value = is_1_loop
_ = is_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
var xs_3 []gopurs_runtime.Value = xs_3_loop
_ = xs_3
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_withArray(), gopurs_runtime.Func(func(res_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(pkg_Data_Foldable.Get_traverse_(), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](pkg_Control_Monad_ST_Internal.Get_applicativeST()))}, gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_0)}, gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Data_Array_ST.Get_modify(), gopurs_runtime.Int(i_5.IntVal), f_2, res_4)
}), is_1)
}), gopurs_runtime.Array(xs_3))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_intersperse(a_0_loop gopurs_runtime.Value, arr_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var arr_1 []gopurs_runtime.Value = arr_1_loop
_ = arr_1
v_2_0 := gopurs_runtime.Int(int64(len(arr_1))).IntVal
_ = v_2_0
var __t3 []gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(gopurs_runtime.Int(v_2_0), gopurs_runtime.Int(2))).IntVal) != (0) {
__t3 = arr_1
goto end_branch_3
} else {

}
}
{
__t3 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), pkg_Data_Array_ST.Get_newImpl(), gopurs_runtime.Func(func(out_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return arr_1[0]
}))
_ = __local_var_4_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushImpl(), __local_var_4_1, out_3)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](pkg_Control_Monad_ST_Internal.Get_bindST()))}, gopurs_runtime.Apply3(pkg_Control_Monad_ST_Internal.Get_forImpl(), gopurs_runtime.Int(1), gopurs_runtime.Int(v_2_0), gopurs_runtime.Func(func(idx_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushImpl(), a_0, out_3)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_2 := gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return arr_1[idx_5.IntVal]
}))
_ = __local_var_7_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushImpl(), __local_var_7_2, out_3)
}))
}))
})), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "pure"), out_3)
}))
}))
})), pkg_Data_Array_ST.Get_unsafeFreeze())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_3:
return __t3
}

func Call_intercalate(dictMonoid_0_loop *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonoid_0 *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value] = dictMonoid_0_loop
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
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value]](gopurs_runtime.UncurriedApp5(Get__insertAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Int(__local_var_0), __local_var_1, gopurs_runtime.Array(__local_var_2)))
}

func Call_init(xs_0_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value] {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_eq__2843686287(), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(xs_0))).IntVal), gopurs_runtime.Int(0)).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(xs_0))).IntVal), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Array(xs_0)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value]](__t0)
}

func Call_index(__local_var_0_loop []gopurs_runtime.Value, __local_var_1_loop int64) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 int64 = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(__local_var_0), gopurs_runtime.Int(__local_var_1)))
}

func Call_last(xs_0_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(xs_0), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(xs_0))).IntVal), gopurs_runtime.Int(1)).IntVal)))
}

func Call_unsnoc(xs_0_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v_1.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), v1_2)
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_init(xs_0))}))})))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(xs_0), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(xs_0))).IntVal), gopurs_runtime.Int(1)).IntVal))))}))
}

func Call_modifyAt(i_0_loop int64, f_1_loop gopurs_runtime.Value, xs_2_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value] {
var i_0 int64 = i_0_loop
_ = i_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 []gopurs_runtime.Value = xs_2_loop
_ = xs_2
__local_var_3_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(xs_2), gopurs_runtime.Int(i_0)))
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value]](gopurs_runtime.UncurriedApp5(Get__updateAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Int(i_0), gopurs_runtime.Apply(f_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.UnsafePtr).V0), gopurs_runtime.Array(xs_2))))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value]](__t1)
}

func Call_unzip(xs_0_loop []*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Tuple.Constructor_Tuple[[]gopurs_runtime.Value, []gopurs_runtime.Value] {
var xs_0 []*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = xs_0_loop
_ = xs_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[[]gopurs_runtime.Value, []gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), pkg_Data_Array_ST.Get_newImpl(), gopurs_runtime.Func(func(fsts_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), pkg_Data_Array_ST.Get_newImpl(), gopurs_runtime.Func(func(snds_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.Apply(pkg_Data_Array_ST_Iterator.Get_Iterator(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Int(v_3.IntVal))))}
})), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_newImpl(), gopurs_runtime.Int(0))), gopurs_runtime.Func(func(iter_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](pkg_Control_Monad_ST_Internal.Get_bindST()))}, gopurs_runtime.Apply2(pkg_Data_Array_ST_Iterator.Get_iterate(), gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Array_ST_Iterator.Constructor_Iterator[gopurs_runtime.Value, *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](iter_3))}, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_0 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1
_ = __local_var_5_0
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](pkg_Control_Monad_ST_Internal.Get_bindST()))}, gopurs_runtime.Apply(Get_void(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushImpl(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, fsts_1)
})), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_void(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushImpl(), __local_var_5_0, snds_2)
}))
}))
})), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_unsafeFreezeImpl(), fsts_1)
}), gopurs_runtime.Func(func(fsts_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_unsafeFreezeImpl(), snds_2)
}), gopurs_runtime.Func(func(snds_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
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

func Call_head(xs_0_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(xs_0), gopurs_runtime.Int(0)))
}

func Call_nubBy(comp_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
indexedAndSorted_2_0 := func() []*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(Call_sortBy(gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(comp_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1).IntVal)), UnsafePtr: nil}
})
}), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := func() []*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_FunctorWithIndex.Get_functorWithIndexArray(), "mapWithIndex"), pkg_Data_Tuple.Get_Tuple(), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value], len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](v) }
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
					unboxed := make([]*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value], len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](v) }
					return unboxed
				}()
_ = indexedAndSorted_2_0
v_3_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(func() []gopurs_runtime.Value {
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
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_3_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_3_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_3_1)}.UnsafePtr == nil) {
__t6 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_3_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_3_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_3_1)}.UnsafePtr != nil) {
__t6 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorArray(), "map"), pkg_Data_Tuple.Get_snd(), func() gopurs_runtime.Value {
					arr := func() []*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Call_sortWith(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&pkg_Data_Ord.Constructor_Ord[int64]{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", pkg_Data_Eq.Get_eqIntImpl())
}), gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})})}), pkg_Data_Tuple.Get_fst()), func() gopurs_runtime.Value {
					arr := func() []*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_unsafeThawImpl(), func() gopurs_runtime.Value {
					arr := func() []*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := []*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]{gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]]((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_3_1)}.UnsafePtr).V0)}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value], len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}())
}), gopurs_runtime.Func(func(result_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](pkg_Control_Monad_ST_Internal.Get_bindST()))}, gopurs_runtime.Apply2(pkg_Control_Monad_ST_Internal.Get_foreach(), func() gopurs_runtime.Value {
					arr := indexedAndSorted_2_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_2 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1
_ = __local_var_6_2
__local_var_7_3 := gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_4 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_8.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_8.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))).IntVal), gopurs_runtime.Int(1)).IntVal)))
_ = __local_var_9_4
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_9_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_9_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_9_4)}.UnsafePtr != nil) {
__t5 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_9_4)}.UnsafePtr).V0
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(__local_var_7_3, x_8).UnsafePtr).V1
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_unsafeFreezeImpl(), result_4)
})), gopurs_runtime.Func(func(lst_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_when__4245282944((gopurs_runtime.Bool(Call_notEq__1272715810(gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(comp_0, lst_7, __local_var_6_2).IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil})).IntVal) != (0), gopurs_runtime.Apply(Get_void(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushImpl(), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](v1_5))}, result_4)
})))
}))
})), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_unsafeFreezeImpl(), result_4)
})
}))
}))).UnsafePtr)
					unboxed := make([]*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value], len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value], len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t6.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_nub(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_nubBy(), dictOrd_0.V1)
}

func Call_groupBy(op_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) [][]gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), pkg_Data_Array_ST.Get_newImpl(), gopurs_runtime.Func(func(result_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.Apply(pkg_Data_Array_ST_Iterator.Get_Iterator(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(xs_1), gopurs_runtime.Int(v_3.IntVal))))}
})), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_newImpl(), gopurs_runtime.Int(0))), gopurs_runtime.Func(func(iter_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](pkg_Control_Monad_ST_Internal.Get_bindST()))}, gopurs_runtime.Apply2(pkg_Data_Array_ST_Iterator.Get_iterate(), gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Array_ST_Iterator.Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](iter_3))}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), pkg_Data_Array_ST.Get_newImpl(), gopurs_runtime.Func(func(sub_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushImpl(), x_4, sub_5)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](pkg_Control_Monad_ST_Internal.Get_bindST()))}, gopurs_runtime.Apply3(pkg_Data_Array_ST_Iterator.Get_pushWhile(), gopurs_runtime.Apply(op_0, x_4), gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Array_ST_Iterator.Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](iter_3))}, sub_5), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_unsafeFreezeImpl(), sub_5)
}), gopurs_runtime.Func(func(grp_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushImpl(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
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
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_unsafeFreezeImpl(), result_2)
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

func Call_groupAllBy(cmp_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
__local_var_1_0 := gopurs_runtime.Apply(Get_groupBy(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_eq__1272715810(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(cmp_0, x_1, y_2).IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}).IntVal) != (0))
})
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

func Call_groupAll(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return Call_groupAllBy(dictOrd_0.V1)
}

func Call_group(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
eq_1_0 := dictEq_0.V0
_ = eq_1_0
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := Call_groupBy(eq_1_0, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2.UnsafePtr)
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
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
}

func Call_fromFoldable(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
__local_var_1_0 := dictFoldable_0.V2
_ = __local_var_1_0
return gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_fromFoldableImpl(), __local_var_1_0, __local_var_2)
})
}

func Call_transpose(xs_0_loop [][]gopurs_runtime.Value) [][]gopurs_runtime.Value {
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
v_4_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(nextArr_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(nextArr_5.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Int(idx_2)))
_ = __local_var_6_2
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_6_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_6_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_6_2)}.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value]](acc_4))}
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_6_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_6_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_6_2)}.UnsafePtr != nil) {
var __t5 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value]](acc_4))}
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 930809136 && __t_tag_3.UnsafePtr == nil) {
__t5 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{(*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_6_2)}.UnsafePtr).V0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value]](acc_4))}
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 930809136 && __t_tag_4.UnsafePtr != nil) {
__t5 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_withArray(), gopurs_runtime.Apply(pkg_Data_Array_ST.Get_push(), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_6_2)}.UnsafePtr).V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(acc_4.UnsafePtr).V0.UnsafePtr)
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
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t5.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value]](__t6))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}()))
_ = v_4_1
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_4_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr == nil) {
__t7 = func() gopurs_runtime.Value {
					arr := allArrays_3
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}()
goto end_branch_7
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_4_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr != nil) {
idx_2_loop = gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int(idx_2), gopurs_runtime.Int(1)).IntVal
allArrays_3_loop = func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(pkg_Data_Array_ST.Get_withArray(), gopurs_runtime.Apply(pkg_Data_Array_ST.Get_push(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr).V0.UnsafePtr)
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
__t7 = func() gopurs_runtime.Value {
					arr := func() [][]gopurs_runtime.Value {
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
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}()
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return func() gopurs_runtime.Value {
					arr := func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t7.UnsafePtr)
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

func Call_foldRecM(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
Monad0_1_0 := gopurs_runtime.Apply(dictMonadRec_0.V0, gopurs_runtime.Value{})
_ = Monad0_1_0
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_1
Bind1_3_2 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_2
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(array_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(o_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_greaterThanOrEq__4087042607(gopurs_runtime.Int(gopurs_runtime.RecordGet(o_7, "b").IntVal), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(array_6.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))).IntVal))).IntVal) != (0) {
__t3 = gopurs_runtime.Apply(Applicative0_2_1.V1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(o_7, "a")})})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(Bind1_3_2.V1, gopurs_runtime.Apply2(f_4, gopurs_runtime.RecordGet(o_7, "a"), gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(array_6.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()[gopurs_runtime.RecordGet(o_7, "b").IntVal]
}))), gopurs_runtime.Func(func(res_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_2_1.V1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("a", "b", res_prime_8, gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int(gopurs_runtime.RecordGet(o_7, "b").IntVal), gopurs_runtime.Int(1)).IntVal))})})
}))
}
end_branch_3:
return __t3
}), gopurs_runtime.RecordDict2("a", "b", b_5, gopurs_runtime.Int(0)))
})
})
})
}

func Call_foldMap(dictMonoid_0_loop *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonoid_0 *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value] = dictMonoid_0_loop
_ = dictMonoid_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_0)})
}

func Call_foldM(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
foldM:
for {
if false { continue foldM }
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get_unconsImpl(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_0.V1, b_4)
}), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(as_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(f_3, b_4, a_6), gopurs_runtime.Func(func(b_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Call_foldM(dictMonad_0), f_3, b_prime_8, gopurs_runtime.Array(func() []gopurs_runtime.Value {
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

func Call_fold(dictMonoid_0_loop *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonoid_0 *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value] = dictMonoid_0_loop
_ = dictMonoid_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_0)}, pkg_Data_Foldable.Get_identity1())
}

func Call_findMap(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_findMapImpl(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, pkg_Data_Maybe.Get_isJust(), __local_var_0, gopurs_runtime.Array(__local_var_1)))
}

func Call_findLastIndex(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[int64] {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.UncurriedApp4(Get_findLastIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, __local_var_0, gopurs_runtime.Array(__local_var_1)))
}

func Call_insertBy(cmp_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, ys_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var ys_2 []gopurs_runtime.Value = ys_2_loop
_ = ys_2
i_3_0 := Call_maybe__919206801(gopurs_runtime.Int(0), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int(v_3.IntVal), gopurs_runtime.Int(1)).IntVal)
}), gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.UncurriedApp4(Get_findLastIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_eq__1272715810(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(cmp_0, x_1, y_3).IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}).IntVal) != (0))
}), gopurs_runtime.Array(ys_2))))}))
_ = i_3_0
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value]](gopurs_runtime.UncurriedApp5(Get__insertAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Int(i_3_0.IntVal), x_1, gopurs_runtime.Array(ys_2)))
_ = __local_var_5_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_1)}.UnsafePtr != nil) {
__t2 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_1)}.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t2.UnsafePtr)
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

func Call_insert(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_insertBy(), dictOrd_0.V1)
}

func Call_findIndex(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[int64] {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.UncurriedApp4(Get_findIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, __local_var_0, gopurs_runtime.Array(__local_var_1)))
}

func Call_span(p_0_loop gopurs_runtime.Value, arr_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var arr_1 []gopurs_runtime.Value = arr_1_loop
_ = arr_1
breakIndex_2_0 := gopurs_runtime.UncurriedApp4(Get_findIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply(Get_not__3201284355(), gopurs_runtime.Bool((gopurs_runtime.Apply(p_0, x_2).IntVal) != (0))).IntVal) != (0))
}), gopurs_runtime.Array(arr_1))
_ = breakIndex_2_0
var __t2 gopurs_runtime.Value
{
if (breakIndex_2_0.Type == 9 && breakIndex_2_0.IntVal == 930809136 && breakIndex_2_0.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if ((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(breakIndex_2_0.UnsafePtr).V0.IntVal) == (0) {
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
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(breakIndex_2_0.UnsafePtr).V0.IntVal), gopurs_runtime.Array(arr_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_sliceImpl(), gopurs_runtime.Int((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(breakIndex_2_0.UnsafePtr).V0.IntVal), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(arr_1))).IntVal), gopurs_runtime.Array(arr_1)).UnsafePtr)
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

func Call_find(f_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_3 gopurs_runtime.Value) gopurs_runtime.Value {
return xs_1[__local_var_3.IntVal]
})
})), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.UncurriedApp4(Get_findIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, f_0, gopurs_runtime.Array(xs_1))))}))
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

func Call_intersectBy(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value, ys_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 []gopurs_runtime.Value = ys_2_loop
_ = ys_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_filterImpl(), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Bool(Call_isJust__1358705270(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.UncurriedApp4(Get_findIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Apply(eq_0, x_3), gopurs_runtime.Array(ys_2))))}))).IntVal) != (0))
}), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_intersect(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_intersectBy(), dictEq_0.V0)
}

func Call_elemLastIndex(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_findLastIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(dictEq_0.V0, v_2, x_1).IntVal) != (0))
}))
}

func Call_elemIndex(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_findIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(dictEq_0.V0, v_2, x_1).IntVal) != (0))
}))
}

func Call_notElem(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value, arr_2_loop []gopurs_runtime.Value) bool {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var arr_2 []gopurs_runtime.Value = arr_2_loop
_ = arr_2
__local_var_3_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.UncurriedApp4(Get_findIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(dictEq_0.V0, v_3, a_1).IntVal) != (0))
}), gopurs_runtime.Array(arr_2)))
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.UnsafePtr != nil) {
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

func Call_elem(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value, arr_2_loop []gopurs_runtime.Value) bool {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var arr_2 []gopurs_runtime.Value = arr_2_loop
_ = arr_2
__local_var_3_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.UncurriedApp4(Get_findIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(dictEq_0.V0, v_3, a_1).IntVal) != (0))
}), gopurs_runtime.Array(arr_2)))
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.UnsafePtr != nil) {
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
__local_var_2_0 := gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(xs_1))).IntVal), gopurs_runtime.Int(n_0)).IntVal
_ = __local_var_2_0
var __t1 []gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(gopurs_runtime.Int(__local_var_2_0), gopurs_runtime.Int(1))).IntVal) != (0) {
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
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int(__local_var_2_0), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_1:
return __t1
}

func Call_drop(n_0_loop int64, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var __t0 []gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(gopurs_runtime.Int(n_0), gopurs_runtime.Int(1))).IntVal) != (0) {
__t0 = xs_1
goto end_branch_0
} else {

}
}
{
__t0 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_sliceImpl(), gopurs_runtime.Int(n_0), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(xs_1))).IntVal), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_0:
return __t0
}

func Call_takeEnd(n_0_loop int64, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
__local_var_2_0 := gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(xs_1))).IntVal), gopurs_runtime.Int(n_0)).IntVal
_ = __local_var_2_0
var __t1 []gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(gopurs_runtime.Int(__local_var_2_0), gopurs_runtime.Int(1))).IntVal) != (0) {
__t1 = xs_1
goto end_branch_1
} else {

}
}
{
__t1 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_sliceImpl(), gopurs_runtime.Int(__local_var_2_0), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(xs_1))).IntVal), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_1:
return __t1
}

func Call_deleteAt(__local_var_0_loop int64, __local_var_1_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value] {
var __local_var_0 int64 = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get__deleteAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Int(__local_var_0), gopurs_runtime.Array(__local_var_1)))
}

func Call_deleteBy(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
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
__local_var_3_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.UncurriedApp4(Get_findIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Apply(v_0, v1_1), gopurs_runtime.Array(v2_2)))
_ = __local_var_3_0
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.UnsafePtr == nil) {
__t4 = gopurs_runtime.Array(v2_2)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.UnsafePtr != nil) {
__local_var_4_1 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.UnsafePtr).V0
_ = __local_var_4_1
__t4 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get__deleteAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Int(__local_var_4_1.IntVal), gopurs_runtime.Array(v2_2)))
_ = __local_var_6_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_6_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_6_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_6_2)}.UnsafePtr != nil) {
__t3 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_6_2)}.UnsafePtr).V0
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

func Call_delete(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_deleteBy(), dictEq_0.V0)
}

func Call_difference(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldr"), Call_delete(dictEq_0))
}

func Call_cons(x_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
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

func Call_some(dictAlternative_0_loop *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V0, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V1, gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(dictLazy_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_cons(), v_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_3, "defer"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_many(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]](dictLazy_3))}, v_4)
})))
})
})
}

func Call_many(dictAlternative_0_loop *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
Alt0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V1, gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}))
_ = Alt0_1_0
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictAlternative_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func(func(dictLazy_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Alt0_1_0.V1, gopurs_runtime.Apply2(Call_some(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]](dictLazy_3))}, v_4), gopurs_runtime.Apply(Applicative0_2_1.V1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))
})
})
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
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{(*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_2_0.UnsafePtr).V0}).UnsafePtr)
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

func Call_filterA(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(p_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Traversable.Get_traversableArray(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), x_3), gopurs_runtime.Apply(p_2, x_3))
}))
_ = __local_var_3_1
__local_var_4_2 := gopurs_runtime.Apply(Functor0_1_0.V0, Call_mapMaybe(gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if ((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1.IntVal) != (0) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t3))}
})))
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_3_1, x_5))
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

func Call_nubByEq(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), pkg_Data_Array_ST.Get_newImpl(), gopurs_runtime.Func(func(arr_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](pkg_Control_Monad_ST_Internal.Get_bindST()))}, gopurs_runtime.Apply2(pkg_Control_Monad_ST_Internal.Get_foreach(), gopurs_runtime.Array(xs_1), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(Get_any(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(eq_0, v_4, x_3).IntVal) != (0))
}))
_ = __local_var_4_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Apply(__local_var_4_0, x_5))
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_unsafeFreezeImpl(), arr_2)
})), gopurs_runtime.Func(func(e_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_when__4245282944((e_4.IntVal) != (0), gopurs_runtime.Apply(Get_void(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushImpl(), x_3, arr_2)
})))
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

func Call_nubEq(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_nubByEq(), dictEq_0.V0)
}

func Call_unionBy(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value, ys_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 []gopurs_runtime.Value = ys_2_loop
_ = ys_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"), gopurs_runtime.Array(xs_1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_deleteBy(eq_0, a_4, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(b_3.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(Call_nubByEq(eq_0, ys_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_union(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_unionBy(), dictEq_0.V0)
}

func Call_alterAt(i_0_loop int64, f_1_loop gopurs_runtime.Value, xs_2_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value] {
var i_0 int64 = i_0_loop
_ = i_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 []gopurs_runtime.Value = xs_2_loop
_ = xs_2
__local_var_3_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(xs_2), gopurs_runtime.Int(i_0)))
_ = __local_var_3_0
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.UnsafePtr != nil) {
v_4_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.UnsafePtr).V0))
_ = v_4_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_4_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get__deleteAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Int(i_0), gopurs_runtime.Array(xs_2))))}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_4_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_4_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value]](gopurs_runtime.UncurriedApp5(Get__updateAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Int(i_0), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_4_1)}.UnsafePtr).V0, gopurs_runtime.Array(xs_2))))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value]](__t3)
}

func Call_all(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) bool {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return (gopurs_runtime.UncurriedApp2(Get_allImpl(), __local_var_0, gopurs_runtime.Array(__local_var_1)).IntVal) != (0)
}

func Call_alt__267341625(dict_0_loop *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_alt__3999254335(dict_0_loop *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__2935994064(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__2960485136(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__566620048(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_when__1954875249(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], v_1_loop bool, v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var v_1 bool = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t0 gopurs_runtime.Value
{
if v_1 {
__t0 = v1_2
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(dictApplicative_0.V1, pkg_Data_Unit.Get_unit())
}
end_branch_0:
return __t0
}

func Call_when__4245282944(v_0_loop bool, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 bool = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if v_0 {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "pure"), pkg_Data_Unit.Get_unit())
}
end_branch_0:
return __t0
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__3582447820(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__760543367(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_discard__317162198(dict_0_loop *pkg_Control_Bind.Constructor_Discard[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Discard[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_defer__3967925939(dict_0_loop *pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_defer__308916730(dict_0_loop *pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_tailRecM__3865988408(dict_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_tailRecM__1444729948(dict_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_tailRecM2__3241310373(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(o_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(o_4, "a"), gopurs_runtime.RecordGet(o_4, "b"))
}), gopurs_runtime.RecordDict2("a", "b", a_2, b_3))
}

func Call_tailRecM2__1943630176(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value, b_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
return gopurs_runtime.Apply2(dictMonadRec_0.V1, gopurs_runtime.Func(func(o_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.RecordGet(o_4, "a"), gopurs_runtime.RecordGet(o_4, "b"))
}), gopurs_runtime.RecordDict2("a", "b", a_2, b_3))
}

func Call_modify__3866314397(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_modifyImpl(), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
s_prime_2_0 := gopurs_runtime.Apply(f_0, s_1)
_ = s_prime_2_0
return gopurs_runtime.RecordDict2("state", "value", s_prime_2_0, s_prime_2_0)
}))
}

func Call_modify__781734141(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_modifyImpl(), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
s_prime_2_0 := gopurs_runtime.Apply(f_0, s_1)
_ = s_prime_2_0
return gopurs_runtime.RecordDict2("state", "value", s_prime_2_0, s_prime_2_0)
}))
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_composeFlipped__2583068543(dictSemigroupoid_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
return gopurs_runtime.Apply2(dictSemigroupoid_0.V0, g_2, f_1)
}

func Call_iterate__1936300090(iter_0_loop *pkg_Data_Array_ST_Iterator.Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var iter_0 *pkg_Data_Array_ST_Iterator.Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value] = iter_0_loop
_ = iter_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_newImpl(), gopurs_runtime.Bool(false)), gopurs_runtime.Func(func(go__break_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Monad_ST_Internal.Get_while(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(go__break_2.PtrVal().(*interface{}))).(gopurs_runtime.Value)
})), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply(pkg_Data_Array_ST_Iterator.Get_next(), gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(iter_0)}), gopurs_runtime.Func(func(mx_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (mx_3.Type == 9 && mx_3.IntVal == 930809136 && mx_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(f_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(mx_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (mx_3.Type == 9 && mx_3.IntVal == 930809136 && mx_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Apply(pkg_Data_Array_ST_Iterator.Get_void1(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(go__break_2.PtrVal().(*interface{})) = gopurs_runtime.Bool(true)
return gopurs_runtime.Bool(true)
}))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})))
}))
}

func Call_iterate__1835948090(iter_0_loop *pkg_Data_Array_ST_Iterator.Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var iter_0 *pkg_Data_Array_ST_Iterator.Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value] = iter_0_loop
_ = iter_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_newImpl(), gopurs_runtime.Bool(false)), gopurs_runtime.Func(func(go__break_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Monad_ST_Internal.Get_while(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(go__break_2.PtrVal().(*interface{}))).(gopurs_runtime.Value)
})), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply(pkg_Data_Array_ST_Iterator.Get_next(), gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(iter_0)}), gopurs_runtime.Func(func(mx_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (mx_3.Type == 9 && mx_3.IntVal == 930809136 && mx_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(f_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(mx_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (mx_3.Type == 9 && mx_3.IntVal == 930809136 && mx_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Apply(pkg_Data_Array_ST_Iterator.Get_void1(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(go__break_2.PtrVal().(*interface{})) = gopurs_runtime.Bool(true)
return gopurs_runtime.Bool(true)
}))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})))
}))
}

func Call_iterator__1149050118(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.Apply(pkg_Data_Array_ST_Iterator.Get_Iterator(), f_0), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_newImpl(), gopurs_runtime.Int(0)))
}

func Call_iterator__2947907462(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.Apply(pkg_Data_Array_ST_Iterator.Get_Iterator(), f_0), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_newImpl(), gopurs_runtime.Int(0)))
}

func Call_next__2731492779(v_0_loop *pkg_Data_Array_ST_Iterator.Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Array_ST_Iterator.Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
__local_var_1_0 := (*pkg_Data_Array_ST_Iterator.Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
_ = __local_var_1_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(__local_var_1_0.PtrVal().(*interface{}))).(gopurs_runtime.Value)
}), gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := (*(__local_var_1_0.PtrVal().(*interface{}))).(gopurs_runtime.Value)
_ = __local_var_3_1
*(__local_var_1_0.PtrVal().(*interface{})) = gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int(__local_var_3_1.IntVal), gopurs_runtime.Int(1)).IntVal)
return gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int(__local_var_3_1.IntVal), gopurs_runtime.Int(1)).IntVal)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply((*pkg_Data_Array_ST_Iterator.Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.Int(i_2.IntVal))))})
}))
}))
}

func Call_peek__2731492779(v_0_loop *pkg_Data_Array_ST_Iterator.Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Array_ST_Iterator.Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*((*pkg_Data_Array_ST_Iterator.Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1.PtrVal().(*interface{}))).(gopurs_runtime.Value)
}), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply((*pkg_Data_Array_ST_Iterator.Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.Int(i_1.IntVal))))})
}))
}

func Call_pushWhile__2298419255(p_0_loop gopurs_runtime.Value, iter_1_loop *pkg_Data_Array_ST_Iterator.Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value], array_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var iter_1 *pkg_Data_Array_ST_Iterator.Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value] = iter_1_loop
_ = iter_1
var array_2 gopurs_runtime.Value = array_2_loop
_ = array_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_newImpl(), gopurs_runtime.Bool(false)), gopurs_runtime.Func(func(go__break_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Monad_ST_Internal.Get_while(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(go__break_3.PtrVal().(*interface{}))).(gopurs_runtime.Value)
})), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply(pkg_Data_Array_ST_Iterator.Get_peek(), gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(iter_1)}), gopurs_runtime.Func(func(mx_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if ((mx_4.Type == 9 && mx_4.IntVal == 930809136 && mx_4.UnsafePtr != nil)) && ((gopurs_runtime.Apply(p_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(mx_4.UnsafePtr).V0).IntVal) != (0)) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushImpl(), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(mx_4.UnsafePtr).V0, array_2)
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Array_ST_Iterator.Get_void(), gopurs_runtime.Apply(pkg_Data_Array_ST_Iterator.Get_next(), gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(iter_1)}))
}))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(pkg_Data_Array_ST_Iterator.Get_void1(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(go__break_3.PtrVal().(*interface{})) = gopurs_runtime.Bool(true)
return gopurs_runtime.Bool(true)
}))
}
end_branch_0:
return __t0
})))
}))
}

func Call_modify__3999472078(i_0_loop int64, f_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var i_0 int64 = i_0_loop
_ = i_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(pkg_Data_Array_ST.Get_peekImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Int(i_0), xs_2)
}), gopurs_runtime.Func(func(entry_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (entry_3.Type == 9 && entry_3.IntVal == 930809136 && entry_3.UnsafePtr != nil) {
__local_var_4_0 := gopurs_runtime.Apply(f_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(entry_3.UnsafePtr).V0)
_ = __local_var_4_0
__t1 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(pkg_Data_Array_ST.Get_pokeImpl(), gopurs_runtime.Int(i_0), __local_var_4_0, xs_2)
})
goto end_branch_1
} else {

}
}
{
if (entry_3.Type == 9 && entry_3.IntVal == 930809136 && entry_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "pure"), gopurs_runtime.Bool(false))
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

func Call_run__1673253202(st_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
var st_0 gopurs_runtime.Value = st_0_loop
_ = st_0
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), st_0, pkg_Data_Array_ST.Get_unsafeFreeze())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_withArray__126410905(f_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_thawImpl(), gopurs_runtime.Array(xs_1))
}), gopurs_runtime.Func(func(result_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply(f_0, result_2), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_unsafeFreezeImpl(), result_2)
})
}))
}))
}

func Call_withArray__3965688428(f_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_thawImpl(), gopurs_runtime.Array(xs_1))
}), gopurs_runtime.Func(func(result_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply(f_0, result_2), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_unsafeFreezeImpl(), result_2)
})
}))
}))
}

func Call_withArray__2779921436(f_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_thawImpl(), gopurs_runtime.Array(xs_1))
}), gopurs_runtime.Func(func(result_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply(f_0, result_2), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_unsafeFreezeImpl(), result_2)
})
}))
}))
}

func Call_any__3571149479(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) bool {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return (gopurs_runtime.UncurriedApp2(Get_anyImpl(), __local_var_0, gopurs_runtime.Array(__local_var_1)).IntVal) != (0)
}

func Call_concatMap__435921434(b_0_loop gopurs_runtime.Value, a_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
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

func Call_cons__3485573810(x_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
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

func Call_deleteAt__454851725(__local_var_0_loop int64, __local_var_1_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value] {
var __local_var_0 int64 = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get__deleteAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Int(__local_var_0), gopurs_runtime.Array(__local_var_1)))
}

func Call_deleteBy__519303411(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
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
__local_var_3_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.UncurriedApp4(Get_findIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Apply(v_0, v1_1), gopurs_runtime.Array(v2_2)))
_ = __local_var_3_0
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.UnsafePtr == nil) {
__t4 = gopurs_runtime.Array(v2_2)
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.UnsafePtr != nil) {
__local_var_4_1 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.UnsafePtr).V0
_ = __local_var_4_1
__t4 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get__deleteAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Int(__local_var_4_1.IntVal), gopurs_runtime.Array(v2_2)))
_ = __local_var_6_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_6_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_6_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_6_2)}.UnsafePtr != nil) {
__t3 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_6_2)}.UnsafePtr).V0
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

func Call_drop__1426757676(n_0_loop int64, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var __t0 []gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(gopurs_runtime.Int(n_0), gopurs_runtime.Int(1))).IntVal) != (0) {
__t0 = xs_1
goto end_branch_0
} else {

}
}
{
__t0 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_sliceImpl(), gopurs_runtime.Int(n_0), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(xs_1))).IntVal), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_0:
return __t0
}

func Call_elemIndex__33401498(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_findIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(dictEq_0.V0, v_2, x_1).IntVal) != (0))
}))
}

func Call_filter__377906483(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
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

func Call_findIndex__139581937(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[int64] {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.UncurriedApp4(Get_findIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, __local_var_0, gopurs_runtime.Array(__local_var_1)))
}

func Call_findLastIndex__139581937(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[int64] {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.UncurriedApp4(Get_findLastIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, __local_var_0, gopurs_runtime.Array(__local_var_1)))
}

func Call_foldM__2595407950(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get_unconsImpl(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_0.V1, b_4)
}), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(as_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(f_3, b_4, a_6), gopurs_runtime.Func(func(b_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Call_foldM(dictMonad_0), f_3, b_prime_8, gopurs_runtime.Array(func() []gopurs_runtime.Value {
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

func Call_groupAllBy__1923945894(cmp_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
__local_var_1_0 := gopurs_runtime.Apply(Get_groupBy(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_eq__1272715810(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(cmp_0, x_1, y_2).IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}).IntVal) != (0))
})
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

func Call_groupBy__693635452(op_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) [][]gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), pkg_Data_Array_ST.Get_newImpl(), gopurs_runtime.Func(func(result_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.Apply(pkg_Data_Array_ST_Iterator.Get_Iterator(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(xs_1), gopurs_runtime.Int(v_3.IntVal))))}
})), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_newImpl(), gopurs_runtime.Int(0))), gopurs_runtime.Func(func(iter_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](pkg_Control_Monad_ST_Internal.Get_bindST()))}, gopurs_runtime.Apply2(pkg_Data_Array_ST_Iterator.Get_iterate(), gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Array_ST_Iterator.Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](iter_3))}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), pkg_Data_Array_ST.Get_newImpl(), gopurs_runtime.Func(func(sub_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushImpl(), x_4, sub_5)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](pkg_Control_Monad_ST_Internal.Get_bindST()))}, gopurs_runtime.Apply3(pkg_Data_Array_ST_Iterator.Get_pushWhile(), gopurs_runtime.Apply(op_0, x_4), gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Array_ST_Iterator.Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](iter_3))}, sub_5), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_unsafeFreezeImpl(), sub_5)
}), gopurs_runtime.Func(func(grp_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushImpl(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
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
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_unsafeFreezeImpl(), result_2)
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

func Call_head__2056156327(xs_0_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(xs_0), gopurs_runtime.Int(0)))
}

func Call_head__1956412839(xs_0_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(xs_0), gopurs_runtime.Int(0)))
}

func Call_index__2196477387(__local_var_0_loop []gopurs_runtime.Value, __local_var_1_loop int64) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 int64 = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(__local_var_0), gopurs_runtime.Int(__local_var_1)))
}

func Call_index__2659850123(__local_var_0_loop []gopurs_runtime.Value, __local_var_1_loop int64) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 int64 = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(__local_var_0), gopurs_runtime.Int(__local_var_1)))
}

func Call_index__3264285643(__local_var_0_loop []gopurs_runtime.Value, __local_var_1_loop int64) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 int64 = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(__local_var_0), gopurs_runtime.Int(__local_var_1)))
}

func Call_init__976795489(xs_0_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value] {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_eq__2843686287(), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(xs_0))).IntVal), gopurs_runtime.Int(0)).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(xs_0))).IntVal), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Array(xs_0)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value]](__t0)
}

func Call_insertAt__388410084(__local_var_0_loop int64, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value] {
var __local_var_0 int64 = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value]](gopurs_runtime.UncurriedApp5(Get__insertAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Int(__local_var_0), __local_var_1, gopurs_runtime.Array(__local_var_2)))
}

func Call_insertBy__1563432905(cmp_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, ys_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var ys_2 []gopurs_runtime.Value = ys_2_loop
_ = ys_2
i_3_0 := Call_maybe__919206801(gopurs_runtime.Int(0), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int(v_3.IntVal), gopurs_runtime.Int(1)).IntVal)
}), gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.UncurriedApp4(Get_findLastIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_eq__1272715810(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(cmp_0, x_1, y_3).IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}).IntVal) != (0))
}), gopurs_runtime.Array(ys_2))))}))
_ = i_3_0
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value]](gopurs_runtime.UncurriedApp5(Get__insertAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Int(i_3_0.IntVal), x_1, gopurs_runtime.Array(ys_2)))
_ = __local_var_5_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_1)}.UnsafePtr != nil) {
__t2 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_1)}.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t2.UnsafePtr)
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

func Call_intersectBy__145374773(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value, ys_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 []gopurs_runtime.Value = ys_2_loop
_ = ys_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_filterImpl(), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Bool(Call_isJust__1358705270(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.UncurriedApp4(Get_findIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Apply(eq_0, x_3), gopurs_runtime.Array(ys_2))))}))).IntVal) != (0))
}), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_last__2345912124(xs_0_loop []*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]] {
var xs_0 []*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value] = xs_0_loop
_ = xs_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]]](gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(xs_0))).IntVal), gopurs_runtime.Int(1)).IntVal)))
}

func Call_last__2056156327(xs_0_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(xs_0), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(xs_0))).IntVal), gopurs_runtime.Int(1)).IntVal)))
}

func Call_many__1839052385(dictAlternative_0_loop *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
Alt0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V1, gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}))
_ = Alt0_1_0
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictAlternative_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func(func(dictLazy_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Alt0_1_0.V1, gopurs_runtime.Apply2(Call_some(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]](dictLazy_3))}, v_4), gopurs_runtime.Apply(Applicative0_2_1.V1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))
})
})
}

func Call_mapMaybe__1271145181(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_concatMap(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
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
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{(*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_2_0.UnsafePtr).V0}).UnsafePtr)
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

func Call_mapMaybe__3137412285(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_concatMap(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
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
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{(*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_2_0.UnsafePtr).V0}).UnsafePtr)
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

func Call_mapMaybe__2261006141(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_concatMap(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
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
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{(*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_2_0.UnsafePtr).V0}).UnsafePtr)
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

func Call_nubBy__3347533344(comp_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
indexedAndSorted_2_0 := func() []*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(Call_sortBy(gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(comp_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1).IntVal)), UnsafePtr: nil}
})
}), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := func() []*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_FunctorWithIndex.Get_functorWithIndexArray(), "mapWithIndex"), pkg_Data_Tuple.Get_Tuple(), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value], len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](v) }
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
					unboxed := make([]*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value], len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](v) }
					return unboxed
				}()
_ = indexedAndSorted_2_0
v_3_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(func() []gopurs_runtime.Value {
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
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_3_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_3_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_3_1)}.UnsafePtr == nil) {
__t8 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
goto end_branch_8
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_3_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_3_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_3_1)}.UnsafePtr != nil) {
__t8 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorArray(), "map"), pkg_Data_Tuple.Get_snd(), func() gopurs_runtime.Value {
					arr := func() []*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Call_sortWith(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&pkg_Data_Ord.Constructor_Ord[int64]{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", pkg_Data_Eq.Get_eqIntImpl())
}), gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})})}), pkg_Data_Tuple.Get_fst()), func() gopurs_runtime.Value {
					arr := func() []*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_unsafeThawImpl(), func() gopurs_runtime.Value {
					arr := func() []*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := []*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]{gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]]((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_3_1)}.UnsafePtr).V0)}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value], len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}())
}), gopurs_runtime.Func(func(result_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](pkg_Control_Monad_ST_Internal.Get_bindST()))}, gopurs_runtime.Apply2(pkg_Control_Monad_ST_Internal.Get_foreach(), func() gopurs_runtime.Value {
					arr := indexedAndSorted_2_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_2 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1
_ = __local_var_6_2
__local_var_7_3 := gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_4 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_indexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_8.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_8.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))).IntVal), gopurs_runtime.Int(1)).IntVal)))
_ = __local_var_9_4
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_9_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_9_4)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_9_4)}.UnsafePtr != nil) {
__t5 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_9_4)}.UnsafePtr).V0
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(__local_var_7_3, x_8).UnsafePtr).V1
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(pkg_Data_Array_ST.Get_unsafeFreezeImpl(), result_4)
})), gopurs_runtime.Func(func(lst_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_6 := gopurs_runtime.Apply(Get_void(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushImpl(), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](v1_5))}, result_4)
}))
_ = __local_var_8_6
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_notEq__1272715810(gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(comp_0, lst_7, __local_var_6_2).IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil})).IntVal) != (0) {
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
}))).UnsafePtr)
					unboxed := make([]*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value], len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value], len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t8.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_nubByEq__3443670074(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), pkg_Data_Array_ST.Get_newImpl(), gopurs_runtime.Func(func(arr_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](pkg_Control_Monad_ST_Internal.Get_bindST()))}, gopurs_runtime.Apply2(pkg_Control_Monad_ST_Internal.Get_foreach(), gopurs_runtime.Array(xs_1), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(Get_any(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(eq_0, v_4, x_3).IntVal) != (0))
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

func Call_null__1123412116(xs_0_loop []gopurs_runtime.Value) bool {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
return (gopurs_runtime.Apply2(Get_eq__2843686287(), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(xs_0))).IntVal), gopurs_runtime.Int(0)).IntVal) != (0)
}

func Call_singleton__193199869(a_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{a_0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_singleton__2286220742(a_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{a_0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_singleton__2535196422(a_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{a_0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_slice__3011328576(__local_var_0_loop int64, __local_var_1_loop int64, __local_var_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
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

func Call_snoc__3419689714(xs_0_loop []gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) []gopurs_runtime.Value {
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

func Call_snoc__3911647657(xs_0_loop []gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) []gopurs_runtime.Value {
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

func Call_snoc__1505998191(xs_0_loop []gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) []gopurs_runtime.Value {
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

func Call_some__1839052385(dictAlternative_0_loop *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V0, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V1, gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(dictLazy_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_cons(), v_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_3, "defer"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_many(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Lazy.Constructor_Lazy[gopurs_runtime.Value]](dictLazy_3))}, v_4)
})))
})
})
}

func Call_sortBy__3347533344(comp_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_sortByImpl(), comp_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (uint32(v_2.IntVal) == 380165415) {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
if (uint32(v_2.IntVal) == 902936544) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (uint32(v_2.IntVal) == 1527465420) {
__t0 = gopurs_runtime.Int(Call_negate__2635823316(gopurs_runtime.Int(1)).IntVal)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Int(__t0.IntVal)
}), gopurs_runtime.Array(__local_var_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_sortBy__1848798496(comp_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_sortByImpl(), comp_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (uint32(v_2.IntVal) == 380165415) {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
if (uint32(v_2.IntVal) == 902936544) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (uint32(v_2.IntVal) == 1527465420) {
__t0 = gopurs_runtime.Int(Call_negate__2635823316(gopurs_runtime.Int(1)).IntVal)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Int(__t0.IntVal)
}), gopurs_runtime.Array(__local_var_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_sortWith__1917042304(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
__local_var_1_0 := &pkg_Data_Ord.Constructor_Ord[int64]{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", pkg_Data_Eq.Get_eqIntImpl())
}), gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})}
_ = __local_var_1_0
return gopurs_runtime.Apply(Get_sortBy(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(__local_var_1_0.V1, gopurs_runtime.Apply(f_0, x_2), gopurs_runtime.Apply(f_0, y_3)).IntVal)), UnsafePtr: nil}
})
}))
}

func Call_sortWith__3115909389(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply(Get_sortBy(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(dictOrd_0.V1, gopurs_runtime.Apply(f_1, x_2), gopurs_runtime.Apply(f_1, y_3)).IntVal)), UnsafePtr: nil}
})
}))
}

func Call_span__174751768(p_0_loop gopurs_runtime.Value, arr_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var arr_1 []gopurs_runtime.Value = arr_1_loop
_ = arr_1
breakIndex_2_0 := gopurs_runtime.UncurriedApp4(Get_findIndexImpl(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply(Get_not__3201284355(), gopurs_runtime.Bool((gopurs_runtime.Apply(p_0, x_2).IntVal) != (0))).IntVal) != (0))
}), gopurs_runtime.Array(arr_1))
_ = breakIndex_2_0
var __t2 gopurs_runtime.Value
{
if (breakIndex_2_0.Type == 9 && breakIndex_2_0.IntVal == 930809136 && breakIndex_2_0.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if ((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(breakIndex_2_0.UnsafePtr).V0.IntVal) == (0) {
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
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(breakIndex_2_0.UnsafePtr).V0.IntVal), gopurs_runtime.Array(arr_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_sliceImpl(), gopurs_runtime.Int((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(breakIndex_2_0.UnsafePtr).V0.IntVal), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(arr_1))).IntVal), gopurs_runtime.Array(arr_1)).UnsafePtr)
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

func Call_take__1426757676(n_0_loop int64, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var __t0 []gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(gopurs_runtime.Int(n_0), gopurs_runtime.Int(1))).IntVal) != (0) {
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
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int(n_0), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_0:
return __t0
}

func Call_unionBy__145374773(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value, ys_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 []gopurs_runtime.Value = ys_2_loop
_ = ys_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"), gopurs_runtime.Array(xs_1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_deleteBy(eq_0, a_4, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(b_3.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(Call_nubByEq(eq_0, ys_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_unsafeIndex__2808089623(_dollar__unused_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value, __local_var_2_loop int64) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 int64 = __local_var_2_loop
_ = __local_var_2
return __local_var_1[__local_var_2]
}

func Call_updateAt__388410084(__local_var_0_loop int64, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value] {
var __local_var_0 int64 = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value]](gopurs_runtime.UncurriedApp5(Get__updateAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Int(__local_var_0), __local_var_1, gopurs_runtime.Array(__local_var_2)))
}

func Call_zipWith__1350747206(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
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

func Call_zipWith__1220584870(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
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

func Call_zipWith__2000246342(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
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

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_notEq__2384498378(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) bool {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
return (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.Bool((gopurs_runtime.Apply2(dictEq_0.V0, x_1, y_2).IntVal) != (0)), gopurs_runtime.Bool(false)).IntVal) != (0)
}

func Call_notEq__1272715810(x_0_loop gopurs_runtime.Value, y_1_loop gopurs_runtime.Value) bool {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var y_1 gopurs_runtime.Value = y_1_loop
_ = y_1
return (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), x_0, y_1).IntVal) != (0)), gopurs_runtime.Bool(false)).IntVal) != (0)
}

func Call_foldr__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__3591001499(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_traverse___996968168(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_1 := gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{})
_ = __local_var_1_1
Functor0_2_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_2
applySecond_1_0 := gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply2(Functor0_2_2.V0, gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_6
})
}), a_3), b_4)
})
})
_ = applySecond_1_0
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_2, "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_1_0, gopurs_runtime.Apply(f_3, x_4))
}), gopurs_runtime.Apply(dictApplicative_0.V1, pkg_Data_Unit.Get_unit()))
})
})
}

func Call_const__1243414737(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__2082174484(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__220790420(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__4026847508(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__153462564(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__106415396(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_flip__3525081280(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__3021024128(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__2857929472(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__3136194560(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__3772615392(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_map__2475811444(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__3240628980(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1319384564(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1852651540(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_void__3020373336(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Apply(dictFunctor_0.V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
}

func Call_disj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_not__1505204753(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
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

func Call_fromJust__1478027324(v_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
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

func Call_isJust__2514352589(v2_0_loop *pkg_Data_Maybe.Constructor_Just[int64]) bool {
var v2_0 *pkg_Data_Maybe.Constructor_Just[int64] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(true)
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

func Call_isJust__2591355336(v2_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) bool {
var v2_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(true)
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

func Call_isJust__1358705270(v2_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) bool {
var v2_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(true)
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

func Call_isJust__4206805139(v2_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) bool {
var v2_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(true)
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

func Call_isNothing__2514352589(v2_0_loop *pkg_Data_Maybe.Constructor_Just[int64]) bool {
var v2_0 *pkg_Data_Maybe.Constructor_Just[int64] = v2_0_loop
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

func Call_maybe__919206801(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_maybe__3078346790(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_maybe__3718989812(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_maybe__2925953714(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_maybe__1408653394(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_maybe__727024722(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_maybe__1936890643(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_comparing__3506074860(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) uint32 {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return uint32(gopurs_runtime.Apply2(dictOrd_0.V1, gopurs_runtime.Apply(f_1, x_2), gopurs_runtime.Apply(f_1, y_3)).IntVal)
}

func Call_greaterThanOrEq__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) < (a2_1.IntVal) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
return __t0
}

func Call_greaterThanOrEq__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 1527465420) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_lessThan__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) < (a2_1.IntVal) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return __t0
}

func Call_lessThan__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 1527465420) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_lessThanOrEq__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) > (a2_1.IntVal) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
return __t0
}

func Call_lessThanOrEq__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_negate__2635823316(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Int(-(a_0.IntVal))
}

func Call_negate__1364373265(dictRing_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictRing_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dictRing_0_loop
_ = dictRing_0
Semiring0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](gopurs_runtime.Apply(dictRing_0.V0, gopurs_runtime.Value{}))
_ = Semiring0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictRing_0.V1, Semiring0_1_0.V3, a_2)
})
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_sub__2345699288(dict_0_loop *pkg_Data_Ring.Constructor_Ring[*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[*pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_zero__1204848985(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}

func Call_sequence__1886310617(dict_0_loop *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_traverse__314957093(dict_0_loop *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_fst__3478736243(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]) int64 {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.IntVal
}

func Call_snd__1234761462(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}

func Call_snd__1694562576(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}

func Call_snd__395594805(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}

func Call_unfoldr__1786322085(dict_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_unfoldr__1128708256(dict_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
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

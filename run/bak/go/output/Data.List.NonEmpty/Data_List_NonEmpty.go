package Data_List_NonEmpty

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad "gopurs/output/Control.Monad"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_List "gopurs/output/Data.List"
	pkg_Data_List_Internal "gopurs/output/Data.List.Internal"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Newtype "gopurs/output/Data.Newtype"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semigroup_Traversable "gopurs/output/Data.Semigroup.Traversable"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
	pkg_Partial "gopurs/output/Partial"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		cache_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_identity(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](x_0_box)))}
})
	})
	return cache_identity
}

var cache_zipWith gopurs_runtime.Value
var once_zipWith sync.Once
func Get_zipWith() gopurs_runtime.Value {
	once_zipWith.Do(func() {
		cache_zipWith = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_zipWith(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v1_2_box)))}
})
	})
	return cache_zipWith
}

var cache_zipWithA gopurs_runtime.Value
var once_zipWithA sync.Once
func Get_zipWithA() gopurs_runtime.Value {
	once_zipWithA.Do(func() {
		cache_zipWithA = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWithA(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
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

var cache_wrappedOperation2 gopurs_runtime.Value
var once_wrappedOperation2 sync.Once
func Get_wrappedOperation2() gopurs_runtime.Value {
	once_wrappedOperation2.Do(func() {
		cache_wrappedOperation2 = gopurs_runtime.Func4(func(name_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value, v1_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_wrappedOperation2(name_0_box.StrVal(), f_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2_box), gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v1_3_box)))}
})
	})
	return cache_wrappedOperation2
}

var cache_wrappedOperation gopurs_runtime.Value
var once_wrappedOperation sync.Once
func Get_wrappedOperation() gopurs_runtime.Value {
	once_wrappedOperation.Do(func() {
		cache_wrappedOperation = gopurs_runtime.Func3(func(name_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_wrappedOperation(name_0_box.StrVal(), f_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2_box)))}
})
	})
	return cache_wrappedOperation
}

var cache_updateAt gopurs_runtime.Value
var once_updateAt sync.Once
func Get_updateAt() gopurs_runtime.Value {
	once_updateAt.Do(func() {
		cache_updateAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_updateAt(i_0_box.IntVal, a_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2_box)))}
})
	})
	return cache_updateAt
}

var cache_unzip gopurs_runtime.Value
var once_unzip sync.Once
func Get_unzip() gopurs_runtime.Value {
	once_unzip.Do(func() {
		cache_unzip = gopurs_runtime.Func(func(ts_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_unzip(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](ts_0_box)))}
})
	})
	return cache_unzip
}

var cache_unsnoc gopurs_runtime.Value
var once_unsnoc sync.Once
func Get_unsnoc() gopurs_runtime.Value {
	once_unsnoc.Do(func() {
		cache_unsnoc = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsnoc(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_unsnoc
}

var cache_unionBy gopurs_runtime.Value
var once_unionBy sync.Once
func Get_unionBy() gopurs_runtime.Value {
	once_unionBy.Do(func() {
		cache_unionBy = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unionBy(x_0_box)
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

var cache_uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		cache_uncons = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_uncons(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_uncons
}

var cache_toList gopurs_runtime.Value
var once_toList sync.Once
func Get_toList() gopurs_runtime.Value {
	once_toList.Do(func() {
		cache_toList = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_toList(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_toList
}

var cache_toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		cache_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toUnfoldable(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
})
	})
	return cache_toUnfoldable
}

var cache_tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		cache_tail = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_tail(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_tail
}

var cache_sortBy gopurs_runtime.Value
var once_sortBy sync.Once
func Get_sortBy() gopurs_runtime.Value {
	once_sortBy.Do(func() {
		cache_sortBy = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sortBy(x_0_box)
})
	})
	return cache_sortBy
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
		cache_snoc = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_snoc(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box), y_1_box))}
})
	})
	return cache_snoc
}

var cache_singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		cache_singleton = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_singleton(x_0_box))}
})
	})
	return cache_singleton
}

var cache_snoc_prime gopurs_runtime.Value
var once_snoc_prime sync.Once
func Get_snoc_prime() gopurs_runtime.Value {
	once_snoc_prime.Do(func() {
		cache_snoc_prime = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_snoc_prime(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box), v1_1_box))}
})
	})
	return cache_snoc_prime
}

var cache_reverse gopurs_runtime.Value
var once_reverse sync.Once
func Get_reverse() gopurs_runtime.Value {
	once_reverse.Do(func() {
		cache_reverse = gopurs_runtime.Apply2(Get_wrappedOperation(), gopurs_runtime.Str("reverse"), pkg_Data_List.Get_reverse())
	})
	return cache_reverse
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

var cache_nubByEq gopurs_runtime.Value
var once_nubByEq sync.Once
func Get_nubByEq() gopurs_runtime.Value {
	once_nubByEq.Do(func() {
		cache_nubByEq = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_nubByEq(x_0_box)
})
	})
	return cache_nubByEq
}

var cache_nubBy gopurs_runtime.Value
var once_nubBy sync.Once
func Get_nubBy() gopurs_runtime.Value {
	once_nubBy.Do(func() {
		cache_nubBy = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_nubBy(x_0_box)
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

var cache_modifyAt gopurs_runtime.Value
var once_modifyAt sync.Once
func Get_modifyAt() gopurs_runtime.Value {
	once_modifyAt.Do(func() {
		cache_modifyAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_modifyAt(i_0_box.IntVal, f_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2_box)))}
})
	})
	return cache_modifyAt
}

var cache_mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		cache_mapMaybe = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe(x_0_box)
})
	})
	return cache_mapMaybe
}

var cache_partition gopurs_runtime.Value
var once_partition sync.Once
func Get_partition() gopurs_runtime.Value {
	once_partition.Do(func() {
		cache_partition = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_partition(x_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_partition
}

var cache_span gopurs_runtime.Value
var once_span sync.Once
func Get_span() gopurs_runtime.Value {
	once_span.Do(func() {
		cache_span = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_span(x_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_span
}

var cache_take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		cache_take = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_take(x_0_box.IntVal)
})
	})
	return cache_take
}

var cache_takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		cache_takeWhile = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_takeWhile(x_0_box)
})
	})
	return cache_takeWhile
}

var cache_length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		cache_length = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_length(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box)))
})
	})
	return cache_length
}

var cache_last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		cache_last = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_last(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_last
}

var cache_intersectBy gopurs_runtime.Value
var once_intersectBy sync.Once
func Get_intersectBy() gopurs_runtime.Value {
	once_intersectBy.Do(func() {
		cache_intersectBy = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intersectBy(x_0_box)
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

var cache_insertAt gopurs_runtime.Value
var once_insertAt sync.Once
func Get_insertAt() gopurs_runtime.Value {
	once_insertAt.Do(func() {
		cache_insertAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_insertAt(i_0_box.IntVal, a_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2_box)))}
})
	})
	return cache_insertAt
}

var cache_init gopurs_runtime.Value
var once_init sync.Once
func Get_init() gopurs_runtime.Value {
	once_init.Do(func() {
		cache_init = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_init(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_init
}

var cache_index gopurs_runtime.Value
var once_index sync.Once
func Get_index() gopurs_runtime.Value {
	once_index.Do(func() {
		cache_index = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, i_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_index(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box), i_1_box.IntVal))}
})
	})
	return cache_index
}

var cache_head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		cache_head = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_head(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_head
}

var cache_groupBy gopurs_runtime.Value
var once_groupBy sync.Once
func Get_groupBy() gopurs_runtime.Value {
	once_groupBy.Do(func() {
		cache_groupBy = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_groupBy(x_0_box)
})
	})
	return cache_groupBy
}

var cache_groupAllBy gopurs_runtime.Value
var once_groupAllBy sync.Once
func Get_groupAllBy() gopurs_runtime.Value {
	once_groupAllBy.Do(func() {
		cache_groupAllBy = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_groupAllBy(x_0_box)
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

var cache_fromList gopurs_runtime.Value
var once_fromList sync.Once
func Get_fromList() gopurs_runtime.Value {
	once_fromList.Do(func() {
		cache_fromList = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_fromList(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_fromList
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

var cache_findLastIndex gopurs_runtime.Value
var once_findLastIndex sync.Once
func Get_findLastIndex() gopurs_runtime.Value {
	once_findLastIndex.Do(func() {
		cache_findLastIndex = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findLastIndex(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box)))}
})
	})
	return cache_findLastIndex
}

var cache_findIndex gopurs_runtime.Value
var once_findIndex sync.Once
func Get_findIndex() gopurs_runtime.Value {
	once_findIndex.Do(func() {
		cache_findIndex = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findIndex(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box)))}
})
	})
	return cache_findIndex
}

var cache_filterM gopurs_runtime.Value
var once_filterM sync.Once
func Get_filterM() gopurs_runtime.Value {
	once_filterM.Do(func() {
		cache_filterM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filterM(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_filterM
}

var cache_filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		cache_filter = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filter(x_0_box)
})
	})
	return cache_filter
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

var cache_dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		cache_dropWhile = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_dropWhile(x_0_box)
})
	})
	return cache_dropWhile
}

var cache_drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		cache_drop = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_drop(x_0_box.IntVal, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box)))}
})
	})
	return cache_drop
}

var cache_cons_prime gopurs_runtime.Value
var once_cons_prime sync.Once
func Get_cons_prime() gopurs_runtime.Value {
	once_cons_prime.Do(func() {
		cache_cons_prime = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_cons_prime(x_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_1_box)))}
})
	})
	return cache_cons_prime
}

var cache_cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cache_cons = gopurs_runtime.Func2(func(y_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_cons(y_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box)))}
})
	})
	return cache_cons
}

var cache_concatMap gopurs_runtime.Value
var once_concatMap sync.Once
func Get_concatMap() gopurs_runtime.Value {
	once_concatMap.Do(func() {
		cache_concatMap = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_concatMap(b_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](a_1_box)))}
})
	})
	return cache_concatMap
}

var cache_concat gopurs_runtime.Value
var once_concat sync.Once
func Get_concat() gopurs_runtime.Value {
	once_concat.Do(func() {
		cache_concat = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_concat(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_concat
}

var cache_catMaybes gopurs_runtime.Value
var once_catMaybes sync.Once
func Get_catMaybes() gopurs_runtime.Value {
	once_catMaybes.Do(func() {
		cache_catMaybes = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_catMaybes(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]](v_0_box)))}
})
	})
	return cache_catMaybes
}

var cache_appendFoldable gopurs_runtime.Value
var once_appendFoldable sync.Once
func Get_appendFoldable() gopurs_runtime.Value {
	once_appendFoldable.Do(func() {
		cache_appendFoldable = gopurs_runtime.Func3(func(dictFoldable_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_appendFoldable(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box), ys_2_box))}
})
	})
	return cache_appendFoldable
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

var cache_pure__575667894 gopurs_runtime.Value
var once_pure__575667894 sync.Once
func Get_pure__575667894() gopurs_runtime.Value {
	once_pure__575667894.Do(func() {
		cache_pure__575667894 = gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_applicativeNonEmptyList(), "pure")
	})
	return cache_pure__575667894
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

var cache_apply__2169384906 gopurs_runtime.Value
var once_apply__2169384906 sync.Once
func Get_apply__2169384906() gopurs_runtime.Value {
	once_apply__2169384906.Do(func() {
		cache_apply__2169384906 = gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_applyList(), "apply")
	})
	return cache_apply__2169384906
}

var cache_lift2__2762258480 gopurs_runtime.Value
var once_lift2__2762258480 sync.Once
func Get_lift2__2762258480() gopurs_runtime.Value {
	once_lift2__2762258480.Do(func() {
		cache_lift2__2762258480 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__2762258480(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__2762258480
}

var cache_lift2__3213187376 gopurs_runtime.Value
var once_lift2__3213187376 sync.Once
func Get_lift2__3213187376() gopurs_runtime.Value {
	once_lift2__3213187376.Do(func() {
		cache_lift2__3213187376 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__3213187376(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__3213187376
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

var cache_bind__1872090113 gopurs_runtime.Value
var once_bind__1872090113 sync.Once
func Get_bind__1872090113() gopurs_runtime.Value {
	once_bind__1872090113.Do(func() {
		cache_bind__1872090113 = gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_bindList(), "bind")
	})
	return cache_bind__1872090113
}

var cache_bind__2389430209 gopurs_runtime.Value
var once_bind__2389430209 sync.Once
func Get_bind__2389430209() gopurs_runtime.Value {
	once_bind__2389430209.Do(func() {
		cache_bind__2389430209 = gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_bindNonEmptyList(), "bind")
	})
	return cache_bind__2389430209
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

var cache_any__4179648253 gopurs_runtime.Value
var once_any__4179648253 sync.Once
func Get_any__4179648253() gopurs_runtime.Value {
	once_any__4179648253.Do(func() {
		cache_any__4179648253 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_any__4179648253(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dictHeytingAlgebra_1_box))
})
	})
	return cache_any__4179648253
}

var cache_any__842931401 gopurs_runtime.Value
var once_any__842931401 sync.Once
func Get_any__842931401() gopurs_runtime.Value {
	once_any__842931401.Do(func() {
		cache_any__842931401 = func() gopurs_runtime.Value {
semigroupDisj1_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), v_0, v1_1)
})
}))
_ = semigroupDisj1_0_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldMap"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_0_0
}), gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "ff")))
}()
	})
	return cache_any__842931401
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

var cache_foldl__1754241693 gopurs_runtime.Value
var once_foldl__1754241693 sync.Once
func Get_foldl__1754241693() gopurs_runtime.Value {
	once_foldl__1754241693.Do(func() {
		cache_foldl__1754241693 = gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl")
	})
	return cache_foldl__1754241693
}

var cache_foldl__3943124669 gopurs_runtime.Value
var once_foldl__3943124669 sync.Once
func Get_foldl__3943124669() gopurs_runtime.Value {
	once_foldl__3943124669.Do(func() {
		cache_foldl__3943124669 = gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl")
	})
	return cache_foldl__3943124669
}

var cache_foldl__396932925 gopurs_runtime.Value
var once_foldl__396932925 sync.Once
func Get_foldl__396932925() gopurs_runtime.Value {
	once_foldl__396932925.Do(func() {
		cache_foldl__396932925 = gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl")
	})
	return cache_foldl__396932925
}

var cache_foldl__2928402749 gopurs_runtime.Value
var once_foldl__2928402749 sync.Once
func Get_foldl__2928402749() gopurs_runtime.Value {
	once_foldl__2928402749.Do(func() {
		cache_foldl__2928402749 = gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl")
	})
	return cache_foldl__2928402749
}

var cache_foldl__3459294429 gopurs_runtime.Value
var once_foldl__3459294429 sync.Once
func Get_foldl__3459294429() gopurs_runtime.Value {
	once_foldl__3459294429.Do(func() {
		cache_foldl__3459294429 = gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl")
	})
	return cache_foldl__3459294429
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

var cache_foldr__2829803163 gopurs_runtime.Value
var once_foldr__2829803163 sync.Once
func Get_foldr__2829803163() gopurs_runtime.Value {
	once_foldr__2829803163.Do(func() {
		cache_foldr__2829803163 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2829803163(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__2829803163
}

var cache_foldr__2979608669 gopurs_runtime.Value
var once_foldr__2979608669 sync.Once
func Get_foldr__2979608669() gopurs_runtime.Value {
	once_foldr__2979608669.Do(func() {
		cache_foldr__2979608669 = gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr")
	})
	return cache_foldr__2979608669
}

var cache_foldr__3489910557 gopurs_runtime.Value
var once_foldr__3489910557 sync.Once
func Get_foldr__3489910557() gopurs_runtime.Value {
	once_foldr__3489910557.Do(func() {
		cache_foldr__3489910557 = gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr")
	})
	return cache_foldr__3489910557
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

var cache_flip__3563101792 gopurs_runtime.Value
var once_flip__3563101792 sync.Once
func Get_flip__3563101792() gopurs_runtime.Value {
	once_flip__3563101792.Do(func() {
		cache_flip__3563101792 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3563101792(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3563101792
}

var cache_flip__1833071808 gopurs_runtime.Value
var once_flip__1833071808 sync.Once
func Get_flip__1833071808() gopurs_runtime.Value {
	once_flip__1833071808.Do(func() {
		cache_flip__1833071808 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__1833071808(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__1833071808
}

var cache_flip__1570464832 gopurs_runtime.Value
var once_flip__1570464832 sync.Once
func Get_flip__1570464832() gopurs_runtime.Value {
	once_flip__1570464832.Do(func() {
		cache_flip__1570464832 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__1570464832(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__1570464832
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

var cache_map__1256368628 gopurs_runtime.Value
var once_map__1256368628 sync.Once
func Get_map__1256368628() gopurs_runtime.Value {
	once_map__1256368628.Do(func() {
		cache_map__1256368628 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1256368628(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1256368628
}

var cache_map__1483545076 gopurs_runtime.Value
var once_map__1483545076 sync.Once
func Get_map__1483545076() gopurs_runtime.Value {
	once_map__1483545076.Do(func() {
		cache_map__1483545076 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1483545076(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1483545076
}

var cache_map__438443400 gopurs_runtime.Value
var once_map__438443400 sync.Once
func Get_map__438443400() gopurs_runtime.Value {
	once_map__438443400.Do(func() {
		cache_map__438443400 = gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_functorList(), "map")
	})
	return cache_map__438443400
}

var cache_map__291265340 gopurs_runtime.Value
var once_map__291265340 sync.Once
func Get_map__291265340() gopurs_runtime.Value {
	once_map__291265340.Do(func() {
		cache_map__291265340 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__291265340
}

var cache_map__140514012 gopurs_runtime.Value
var once_map__140514012 sync.Once
func Get_map__140514012() gopurs_runtime.Value {
	once_map__140514012.Do(func() {
		cache_map__140514012 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__140514012
}

var cache_map__3210082748 gopurs_runtime.Value
var once_map__3210082748 sync.Once
func Get_map__3210082748() gopurs_runtime.Value {
	once_map__3210082748.Do(func() {
		cache_map__3210082748 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__3210082748
}

var cache_map__2202537180 gopurs_runtime.Value
var once_map__2202537180 sync.Once
func Get_map__2202537180() gopurs_runtime.Value {
	once_map__2202537180.Do(func() {
		cache_map__2202537180 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__2202537180
}

var cache_map__1681779388 gopurs_runtime.Value
var once_map__1681779388 sync.Once
func Get_map__1681779388() gopurs_runtime.Value {
	once_map__1681779388.Do(func() {
		cache_map__1681779388 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__1681779388
}

var cache_map__3486165692 gopurs_runtime.Value
var once_map__3486165692 sync.Once
func Get_map__3486165692() gopurs_runtime.Value {
	once_map__3486165692.Do(func() {
		cache_map__3486165692 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__3486165692
}

var cache_map__91618268 gopurs_runtime.Value
var once_map__91618268 sync.Once
func Get_map__91618268() gopurs_runtime.Value {
	once_map__91618268.Do(func() {
		cache_map__91618268 = gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_functorNonEmptyList(), "map")
	})
	return cache_map__91618268
}

var cache_mapFlipped__4215217780 gopurs_runtime.Value
var once_mapFlipped__4215217780 sync.Once
func Get_mapFlipped__4215217780() gopurs_runtime.Value {
	once_mapFlipped__4215217780.Do(func() {
		cache_mapFlipped__4215217780 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapFlipped__4215217780(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_mapFlipped__4215217780
}

var cache_mapFlipped__2919806324 gopurs_runtime.Value
var once_mapFlipped__2919806324 sync.Once
func Get_mapFlipped__2919806324() gopurs_runtime.Value {
	once_mapFlipped__2919806324.Do(func() {
		cache_mapFlipped__2919806324 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapFlipped__2919806324(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_mapFlipped__2919806324
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

var cache_emptySet__2398681994 gopurs_runtime.Value
var once_emptySet__2398681994 sync.Once
func Get_emptySet__2398681994() gopurs_runtime.Value {
	once_emptySet__2398681994.Do(func() {
		cache_emptySet__2398681994 = gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)}
	})
	return cache_emptySet__2398681994
}

var cache_fromZipper__1019554324 gopurs_runtime.Value
var once_fromZipper__1019554324 sync.Once
func Get_fromZipper__1019554324() gopurs_runtime.Value {
	once_fromZipper__1019554324.Do(func() {
		cache_fromZipper__1019554324 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromZipper__1019554324(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box), v1_1_box)
})
	})
	return cache_fromZipper__1019554324
}

var cache_insertAndLookupBy__3244745033 gopurs_runtime.Value
var once_insertAndLookupBy__3244745033 sync.Once
func Get_insertAndLookupBy__3244745033() gopurs_runtime.Value {
	once_insertAndLookupBy__3244745033.Do(func() {
		cache_insertAndLookupBy__3244745033 = gopurs_runtime.Func3(func(comp_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value, orig_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insertAndLookupBy__3244745033(comp_0_box, k_1_box, orig_2_box)
})
	})
	return cache_insertAndLookupBy__3244745033
}

var cache_findIndex__1077906787 gopurs_runtime.Value
var once_findIndex__1077906787 sync.Once
func Get_findIndex__1077906787() gopurs_runtime.Value {
	once_findIndex__1077906787.Do(func() {
		cache_findIndex__1077906787 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findIndex__1077906787(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box)))}
})
	})
	return cache_findIndex__1077906787
}

var cache_findLastIndex__1077906787 gopurs_runtime.Value
var once_findLastIndex__1077906787 sync.Once
func Get_findLastIndex__1077906787() gopurs_runtime.Value {
	once_findLastIndex__1077906787.Do(func() {
		cache_findLastIndex__1077906787 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findLastIndex__1077906787(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box)))}
})
	})
	return cache_findLastIndex__1077906787
}

var cache_fromList__2793484475 gopurs_runtime.Value
var once_fromList__2793484475 sync.Once
func Get_fromList__2793484475() gopurs_runtime.Value {
	once_fromList__2793484475.Do(func() {
		cache_fromList__2793484475 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_fromList__2793484475(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_fromList__2793484475
}

var cache_lift__1094295643 gopurs_runtime.Value
var once_lift__1094295643 sync.Once
func Get_lift__1094295643() gopurs_runtime.Value {
	once_lift__1094295643.Do(func() {
		cache_lift__1094295643 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_lift__1094295643(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box)))}
})
	})
	return cache_lift__1094295643
}

var cache_lift__184746683 gopurs_runtime.Value
var once_lift__184746683 sync.Once
func Get_lift__184746683() gopurs_runtime.Value {
	once_lift__184746683.Do(func() {
		cache_lift__184746683 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift__184746683(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_lift__184746683
}

var cache_lift__243257371 gopurs_runtime.Value
var once_lift__243257371 sync.Once
func Get_lift__243257371() gopurs_runtime.Value {
	once_lift__243257371.Do(func() {
		cache_lift__243257371 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift__243257371(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_lift__243257371
}

var cache_lift__2580243643 gopurs_runtime.Value
var once_lift__2580243643 sync.Once
func Get_lift__2580243643() gopurs_runtime.Value {
	once_lift__2580243643.Do(func() {
		cache_lift__2580243643 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift__2580243643(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_lift__2580243643
}

var cache_lift__3667893565 gopurs_runtime.Value
var once_lift__3667893565 sync.Once
func Get_lift__3667893565() gopurs_runtime.Value {
	once_lift__3667893565.Do(func() {
		cache_lift__3667893565 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift__3667893565(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_lift__3667893565
}

var cache_lift__1609676349 gopurs_runtime.Value
var once_lift__1609676349 sync.Once
func Get_lift__1609676349() gopurs_runtime.Value {
	once_lift__1609676349.Do(func() {
		cache_lift__1609676349 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift__1609676349(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_lift__1609676349
}

var cache_singleton__3219659348 gopurs_runtime.Value
var once_singleton__3219659348 sync.Once
func Get_singleton__3219659348() gopurs_runtime.Value {
	once_singleton__3219659348.Do(func() {
		cache_singleton__3219659348 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_singleton__3219659348(x_0_box))}
})
	})
	return cache_singleton__3219659348
}

var cache_sortBy__2726669792 gopurs_runtime.Value
var once_sortBy__2726669792 sync.Once
func Get_sortBy__2726669792() gopurs_runtime.Value {
	once_sortBy__2726669792.Do(func() {
		cache_sortBy__2726669792 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sortBy__2726669792(x_0_box)
})
	})
	return cache_sortBy__2726669792
}

var cache_toList__2859885498 gopurs_runtime.Value
var once_toList__2859885498 sync.Once
func Get_toList__2859885498() gopurs_runtime.Value {
	once_toList__2859885498.Do(func() {
		cache_toList__2859885498 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_toList__2859885498(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_toList__2859885498
}

var cache_wrappedOperation__2422223318 gopurs_runtime.Value
var once_wrappedOperation__2422223318 sync.Once
func Get_wrappedOperation__2422223318() gopurs_runtime.Value {
	once_wrappedOperation__2422223318.Do(func() {
		cache_wrappedOperation__2422223318 = gopurs_runtime.Func3(func(name_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_wrappedOperation__2422223318(name_0_box.StrVal(), f_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2_box)))}
})
	})
	return cache_wrappedOperation__2422223318
}

var cache_wrappedOperation__1897406582 gopurs_runtime.Value
var once_wrappedOperation__1897406582 sync.Once
func Get_wrappedOperation__1897406582() gopurs_runtime.Value {
	once_wrappedOperation__1897406582.Do(func() {
		cache_wrappedOperation__1897406582 = gopurs_runtime.Func3(func(name_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_wrappedOperation__1897406582(name_0_box.StrVal(), f_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2_box)))}
})
	})
	return cache_wrappedOperation__1897406582
}

var cache_wrappedOperation2__4016046219 gopurs_runtime.Value
var once_wrappedOperation2__4016046219 sync.Once
func Get_wrappedOperation2__4016046219() gopurs_runtime.Value {
	once_wrappedOperation2__4016046219.Do(func() {
		cache_wrappedOperation2__4016046219 = gopurs_runtime.Func4(func(name_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value, v1_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_wrappedOperation2__4016046219(name_0_box.StrVal(), f_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2_box), gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v1_3_box)))}
})
	})
	return cache_wrappedOperation2__4016046219
}

var cache_zipWith__1818047060 gopurs_runtime.Value
var once_zipWith__1818047060 sync.Once
func Get_zipWith__1818047060() gopurs_runtime.Value {
	once_zipWith__1818047060.Do(func() {
		cache_zipWith__1818047060 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_zipWith__1818047060(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v1_2_box)))}
})
	})
	return cache_zipWith__1818047060
}

var cache_zipWith__716394740 gopurs_runtime.Value
var once_zipWith__716394740 sync.Once
func Get_zipWith__716394740() gopurs_runtime.Value {
	once_zipWith__716394740.Do(func() {
		cache_zipWith__716394740 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_zipWith__716394740(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v1_2_box)))}
})
	})
	return cache_zipWith__716394740
}

var cache_zipWith__3875859348 gopurs_runtime.Value
var once_zipWith__3875859348 sync.Once
func Get_zipWith__3875859348() gopurs_runtime.Value {
	once_zipWith__3875859348.Do(func() {
		cache_zipWith__3875859348 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_zipWith__3875859348(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v1_2_box)))}
})
	})
	return cache_zipWith__3875859348
}

var cache_applicativeNonEmptyList__1156428081 gopurs_runtime.Value
var once_applicativeNonEmptyList__1156428081 sync.Once
func Get_applicativeNonEmptyList__1156428081() gopurs_runtime.Value {
	once_applicativeNonEmptyList__1156428081.Do(func() {
		cache_applicativeNonEmptyList__1156428081 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Types.Get_applyNonEmptyList()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_0, gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_plusList(), "empty")})}
}))
	})
	return cache_applicativeNonEmptyList__1156428081
}

var cache_applicativeNonEmptyList__3820246605 gopurs_runtime.Value
var once_applicativeNonEmptyList__3820246605 sync.Once
func Get_applicativeNonEmptyList__3820246605() gopurs_runtime.Value {
	once_applicativeNonEmptyList__3820246605.Do(func() {
		cache_applicativeNonEmptyList__3820246605 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Types.Get_applyNonEmptyList()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_0, gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_plusList(), "empty")})}
}))
	})
	return cache_applicativeNonEmptyList__3820246605
}

var cache_applyList__3072763993 gopurs_runtime.Value
var once_applyList__3072763993 sync.Once
func Get_applyList__3072763993() gopurs_runtime.Value {
	once_applyList__3072763993.Do(func() {
		cache_applyList__3072763993 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Types.Get_functorList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_functorList(), "map"), (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1))})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1))})))})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_applyList__3072763993
}

var cache_applyList__1109325167 gopurs_runtime.Value
var once_applyList__1109325167 sync.Once
func Get_applyList__1109325167() gopurs_runtime.Value {
	once_applyList__1109325167.Do(func() {
		cache_applyList__1109325167 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Types.Get_functorList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_functorList(), "map"), (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1))})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1))})))})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_applyList__1109325167
}

var cache_applyNonEmptyList__602103086 gopurs_runtime.Value
var once_applyNonEmptyList__602103086 sync.Once
func Get_applyNonEmptyList__602103086() gopurs_runtime.Value {
	once_applyNonEmptyList__602103086.Do(func() {
		cache_applyNonEmptyList__602103086 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Types.Get_functorNonEmptyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1))})))})))}})}
})
}))
	})
	return cache_applyNonEmptyList__602103086
}

var cache_bindList__241263065 gopurs_runtime.Value
var once_bindList__241263065 sync.Once
func Get_bindList__241263065() gopurs_runtime.Value {
	once_bindList__241263065.Do(func() {
		cache_bindList__241263065 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Types.Get_applyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V0)))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_bindList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V1)}, v1_1)))})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_bindList__241263065
}

var cache_bindNonEmptyList__1408886065 gopurs_runtime.Value
var once_bindNonEmptyList__1408886065 sync.Once
func Get_bindNonEmptyList__1408886065() gopurs_runtime.Value {
	once_bindNonEmptyList__1408886065.Do(func() {
		cache_bindNonEmptyList__1408886065 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Types.Get_applyNonEmptyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
v1_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0))
_ = v1_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v1_2_0)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v1_2_0)}.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_bindList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1))}, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(f_1, x_3)
_ = __local_var_4_1
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1)})}
}))))})))}})}
})
}))
	})
	return cache_bindNonEmptyList__1408886065
}

var cache_bindNonEmptyList__132736205 gopurs_runtime.Value
var once_bindNonEmptyList__132736205 sync.Once
func Get_bindNonEmptyList__132736205() gopurs_runtime.Value {
	once_bindNonEmptyList__132736205.Do(func() {
		cache_bindNonEmptyList__132736205 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Types.Get_applyNonEmptyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
v1_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0))
_ = v1_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v1_2_0)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v1_2_0)}.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_bindList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1))}, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(f_1, x_3)
_ = __local_var_4_1
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1)})}
}))))})))}})}
})
}))
	})
	return cache_bindNonEmptyList__132736205
}

var cache_foldable1NonEmptyList__1746670655 gopurs_runtime.Value
var once_foldable1NonEmptyList__1746670655 sync.Once
func Get_foldable1NonEmptyList__1746670655() gopurs_runtime.Value {
	once_foldable1NonEmptyList__1746670655.Do(func() {
		cache_foldable1NonEmptyList__1746670655 = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_foldable1NonEmpty(), pkg_Data_List_Types.Get_foldableList())
	})
	return cache_foldable1NonEmptyList__1746670655
}

var cache_foldableList__1753400174 gopurs_runtime.Value
var once_foldableList__1753400174 sync.Once
func Get_foldableList__1753400174() gopurs_runtime.Value {
	once_foldableList__1753400174.Do(func() {
		cache_foldableList__1753400174 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(Semigroup0_1_0.V0, acc_4)
_ = __local_var_5_2
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, gopurs_runtime.Apply(f_3, x_6))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_3_5 gopurs_runtime.Value
go__go_1_3_5 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_3_5:
for {
if false { continue go__go_1_3_5 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t4 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t4 = b_2
goto end_branch_4
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(f_0, b_2, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
continue go__go_1_3_5
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}()
})
})
return go__go_1_3_5
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_5 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
})
}), b_1)
_ = __local_var_2_5
var go__go_3_7_6 gopurs_runtime.Value
go__go_3_7_6 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_7_6:
for {
if false { continue go__go_3_7_6 }
var v_4 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t8 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}
goto end_branch_8
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
continue go__go_3_7_6
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t8))}
}
}()
})
})
__local_var_3_6 := gopurs_runtime.Apply(go__go_3_7_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_3_6
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_5, gopurs_runtime.Apply(__local_var_3_6, x_4))
})
})
}))
	})
	return cache_foldableList__1753400174
}

var cache_functorList__4121998062 gopurs_runtime.Value
var once_functorList__4121998062 sync.Once
func Get_functorList__4121998062() gopurs_runtime.Value {
	once_functorList__4121998062.Do(func() {
		cache_functorList__4121998062 = gopurs_runtime.RecordDict1("map", pkg_Data_List_Types.Get_listMap())
	})
	return cache_functorList__4121998062
}

var cache_functorList__1783129585 gopurs_runtime.Value
var once_functorList__1783129585 sync.Once
func Get_functorList__1783129585() gopurs_runtime.Value {
	once_functorList__1783129585.Do(func() {
		cache_functorList__1783129585 = gopurs_runtime.RecordDict1("map", pkg_Data_List_Types.Get_listMap())
	})
	return cache_functorList__1783129585
}

var cache_functorNonEmptyList__2834508934 gopurs_runtime.Value
var once_functorNonEmptyList__2834508934 sync.Once
func Get_functorNonEmptyList__2834508934() gopurs_runtime.Value {
	once_functorNonEmptyList__2834508934.Do(func() {
		cache_functorNonEmptyList__2834508934 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_functorList(), "map"), f_0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_functorNonEmptyList__2834508934
}

var cache_functorNonEmptyList__1593940346 gopurs_runtime.Value
var once_functorNonEmptyList__1593940346 sync.Once
func Get_functorNonEmptyList__1593940346() gopurs_runtime.Value {
	once_functorNonEmptyList__1593940346.Do(func() {
		cache_functorNonEmptyList__1593940346 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_functorList(), "map"), f_0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_functorNonEmptyList__1593940346
}

var cache_functorNonEmptyList__257963697 gopurs_runtime.Value
var once_functorNonEmptyList__257963697 sync.Once
func Get_functorNonEmptyList__257963697() gopurs_runtime.Value {
	once_functorNonEmptyList__257963697.Do(func() {
		cache_functorNonEmptyList__257963697 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_functorList(), "map"), f_0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_functorNonEmptyList__257963697
}

var cache_listMap__4135416762 gopurs_runtime.Value
var once_listMap__4135416762 sync.Once
func Get_listMap__4135416762() gopurs_runtime.Value {
	once_listMap__4135416762.Do(func() {
		cache_listMap__4135416762 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_listMap__4135416762(f_0_box)
})
	})
	return cache_listMap__4135416762
}

var cache_nelCons__195558898 gopurs_runtime.Value
var once_nelCons__195558898 sync.Once
func Get_nelCons__195558898() gopurs_runtime.Value {
	once_nelCons__195558898.Do(func() {
		cache_nelCons__195558898 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_nelCons__195558898(a_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box)))}
})
	})
	return cache_nelCons__195558898
}

var cache_semigroupList__2766094215 gopurs_runtime.Value
var once_semigroupList__2766094215 sync.Once
func Get_semigroupList__2766094215() gopurs_runtime.Value {
	once_semigroupList__2766094215.Do(func() {
		cache_semigroupList__2766094215 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](ys_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_0))})))}
})
}))
	})
	return cache_semigroupList__2766094215
}

var cache_semigroupList__2410751346 gopurs_runtime.Value
var once_semigroupList__2410751346 sync.Once
func Get_semigroupList__2410751346() gopurs_runtime.Value {
	once_semigroupList__2410751346.Do(func() {
		cache_semigroupList__2410751346 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](ys_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_0))})))}
})
}))
	})
	return cache_semigroupList__2410751346
}

var cache_toList__2859885498 gopurs_runtime.Value
var once_toList__2859885498 sync.Once
func Get_toList__2859885498() gopurs_runtime.Value {
	once_toList__2859885498.Do(func() {
		cache_toList__2859885498 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_toList__2859885498(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_toList__2859885498
}

var cache_traversable1NonEmptyList__1171985061 gopurs_runtime.Value
var once_traversable1NonEmptyList__1171985061 sync.Once
func Get_traversable1NonEmptyList__1171985061() gopurs_runtime.Value {
	once_traversable1NonEmptyList__1171985061.Do(func() {
		cache_traversable1NonEmptyList__1171985061 = gopurs_runtime.RecordDict4("Foldable10", "Traversable1", "sequence1", "traverse1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Types.Get_foldable1NonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Types.Get_traversableNonEmptyList()
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_traversable1NonEmptyList(), "traverse1"), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0))}, pkg_Data_List_Types.Get_identity1())
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V1)})}})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_applicativeNonEmptyList(), "pure"), (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0)))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1))})))}
}), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_5_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_2
__local_var_5_1 := gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(Functor0_5_2.V0, gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_8, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_7.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_7.UnsafePtr).V1)})}})}
})
}), acc_4), b_6)
})
_ = __local_var_5_1
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_1, gopurs_runtime.Apply(f_2, x_6))
})
}), gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_applicativeNonEmptyList(), "pure"), gopurs_runtime.Apply(f_2, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))}))
})
})
}))
	})
	return cache_traversable1NonEmptyList__1171985061
}

var cache_traversable1NonEmptyList__1516899673 gopurs_runtime.Value
var once_traversable1NonEmptyList__1516899673 sync.Once
func Get_traversable1NonEmptyList__1516899673() gopurs_runtime.Value {
	once_traversable1NonEmptyList__1516899673.Do(func() {
		cache_traversable1NonEmptyList__1516899673 = gopurs_runtime.RecordDict4("Foldable10", "Traversable1", "sequence1", "traverse1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Types.Get_foldable1NonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Types.Get_traversableNonEmptyList()
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_traversable1NonEmptyList(), "traverse1"), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0))}, pkg_Data_List_Types.Get_identity1())
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V1)})}})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_applicativeNonEmptyList(), "pure"), (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0)))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1))})))}
}), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_5_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_2
__local_var_5_1 := gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(Functor0_5_2.V0, gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_8, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_7.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_7.UnsafePtr).V1)})}})}
})
}), acc_4), b_6)
})
_ = __local_var_5_1
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_1, gopurs_runtime.Apply(f_2, x_6))
})
}), gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_applicativeNonEmptyList(), "pure"), gopurs_runtime.Apply(f_2, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))}))
})
})
}))
	})
	return cache_traversable1NonEmptyList__1516899673
}

var cache_traversableNonEmptyList__1085933743 gopurs_runtime.Value
var once_traversableNonEmptyList__1085933743 sync.Once
func Get_traversableNonEmptyList__1085933743() gopurs_runtime.Value {
	once_traversableNonEmptyList__1085933743.Do(func() {
		cache_traversableNonEmptyList__1085933743 = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_traversableNonEmpty(), pkg_Data_List_Types.Get_traversableList())
	})
	return cache_traversableNonEmptyList__1085933743
}

var cache_alterAt__3453373293 gopurs_runtime.Value
var once_alterAt__3453373293 sync.Once
func Get_alterAt__3453373293() gopurs_runtime.Value {
	once_alterAt__3453373293.Do(func() {
		cache_alterAt__3453373293 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_alterAt__3453373293(v_0_box.IntVal, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_2_box)))}
})
	})
	return cache_alterAt__3453373293
}

var cache_catMaybes__3687414234 gopurs_runtime.Value
var once_catMaybes__3687414234 sync.Once
func Get_catMaybes__3687414234() gopurs_runtime.Value {
	once_catMaybes__3687414234.Do(func() {
		cache_catMaybes__3687414234 = gopurs_runtime.Apply(pkg_Data_List.Get_mapMaybe(), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_catMaybes__3687414234
}

var cache_deleteBy__697302515 gopurs_runtime.Value
var once_deleteBy__697302515 sync.Once
func Get_deleteBy__697302515() gopurs_runtime.Value {
	once_deleteBy__697302515.Do(func() {
		cache_deleteBy__697302515 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_deleteBy__697302515(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_2_box)))}
})
	})
	return cache_deleteBy__697302515
}

var cache_drop__551729751 gopurs_runtime.Value
var once_drop__551729751 sync.Once
func Get_drop__551729751() gopurs_runtime.Value {
	once_drop__551729751.Do(func() {
		cache_drop__551729751 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_drop__551729751(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_drop__551729751
}

var cache_drop__1836090668 gopurs_runtime.Value
var once_drop__1836090668 sync.Once
func Get_drop__1836090668() gopurs_runtime.Value {
	once_drop__1836090668.Do(func() {
		cache_drop__1836090668 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_drop__1836090668(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_drop__1836090668
}

var cache_dropWhile__2352021032 gopurs_runtime.Value
var once_dropWhile__2352021032 sync.Once
func Get_dropWhile__2352021032() gopurs_runtime.Value {
	once_dropWhile__2352021032.Do(func() {
		cache_dropWhile__2352021032 = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_dropWhile__2352021032(p_0_box)
})
	})
	return cache_dropWhile__2352021032
}

var cache_filter__2352021032 gopurs_runtime.Value
var once_filter__2352021032 sync.Once
func Get_filter__2352021032() gopurs_runtime.Value {
	once_filter__2352021032.Do(func() {
		cache_filter__2352021032 = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filter__2352021032(p_0_box)
})
	})
	return cache_filter__2352021032
}

var cache_filter__1617261107 gopurs_runtime.Value
var once_filter__1617261107 sync.Once
func Get_filter__1617261107() gopurs_runtime.Value {
	once_filter__1617261107.Do(func() {
		cache_filter__1617261107 = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filter__1617261107(p_0_box)
})
	})
	return cache_filter__1617261107
}

var cache_findIndex__2366045378 gopurs_runtime.Value
var once_findIndex__2366045378 sync.Once
func Get_findIndex__2366045378() gopurs_runtime.Value {
	once_findIndex__2366045378.Do(func() {
		cache_findIndex__2366045378 = gopurs_runtime.Func(func(fn_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_findIndex__2366045378(fn_0_box)
})
	})
	return cache_findIndex__2366045378
}

var cache_findLastIndex__2366045378 gopurs_runtime.Value
var once_findLastIndex__2366045378 sync.Once
func Get_findLastIndex__2366045378() gopurs_runtime.Value {
	once_findLastIndex__2366045378.Do(func() {
		cache_findLastIndex__2366045378 = gopurs_runtime.Func2(func(fn_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findLastIndex__2366045378(fn_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_1_box)))}
})
	})
	return cache_findLastIndex__2366045378
}

var cache_foldM__3577257629 gopurs_runtime.Value
var once_foldM__3577257629 sync.Once
func Get_foldM__3577257629() gopurs_runtime.Value {
	once_foldM__3577257629.Do(func() {
		cache_foldM__3577257629 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldM__3577257629(gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_foldM__3577257629
}

var cache_fromFoldable__614070391 gopurs_runtime.Value
var once_fromFoldable__614070391 sync.Once
func Get_fromFoldable__614070391() gopurs_runtime.Value {
	once_fromFoldable__614070391.Do(func() {
		cache_fromFoldable__614070391 = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldable__614070391(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
})
	})
	return cache_fromFoldable__614070391
}

var cache_groupAllBy__3934374991 gopurs_runtime.Value
var once_groupAllBy__3934374991 sync.Once
func Get_groupAllBy__3934374991() gopurs_runtime.Value {
	once_groupAllBy__3934374991.Do(func() {
		cache_groupAllBy__3934374991 = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_groupAllBy__3934374991(p_0_box)
})
	})
	return cache_groupAllBy__3934374991
}

var cache_groupBy__2162447253 gopurs_runtime.Value
var once_groupBy__2162447253 sync.Once
func Get_groupBy__2162447253() gopurs_runtime.Value {
	once_groupBy__2162447253.Do(func() {
		cache_groupBy__2162447253 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_groupBy__2162447253(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_groupBy__2162447253
}

var cache_groupBy__1039549870 gopurs_runtime.Value
var once_groupBy__1039549870 sync.Once
func Get_groupBy__1039549870() gopurs_runtime.Value {
	once_groupBy__1039549870.Do(func() {
		cache_groupBy__1039549870 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_groupBy__1039549870(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_groupBy__1039549870
}

var cache_index__304299960 gopurs_runtime.Value
var once_index__304299960 sync.Once
func Get_index__304299960() gopurs_runtime.Value {
	once_index__304299960.Do(func() {
		cache_index__304299960 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_index__304299960(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box), v1_1_box.IntVal))}
})
	})
	return cache_index__304299960
}

var cache_init__2496605985 gopurs_runtime.Value
var once_init__2496605985 sync.Once
func Get_init__2496605985() gopurs_runtime.Value {
	once_init__2496605985.Do(func() {
		cache_init__2496605985 = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_init__2496605985(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](lst_0_box)))}
})
	})
	return cache_init__2496605985
}

var cache_insertAt__2634211748 gopurs_runtime.Value
var once_insertAt__2634211748 sync.Once
func Get_insertAt__2634211748() gopurs_runtime.Value {
	once_insertAt__2634211748.Do(func() {
		cache_insertAt__2634211748 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_insertAt__2634211748(v_0_box.IntVal, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_2_box)))}
})
	})
	return cache_insertAt__2634211748
}

var cache_intersectBy__588351261 gopurs_runtime.Value
var once_intersectBy__588351261 sync.Once
func Get_intersectBy__588351261() gopurs_runtime.Value {
	once_intersectBy__588351261.Do(func() {
		cache_intersectBy__588351261 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_intersectBy__588351261(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_2_box)))}
})
	})
	return cache_intersectBy__588351261
}

var cache_last__4043133652 gopurs_runtime.Value
var once_last__4043133652 sync.Once
func Get_last__4043133652() gopurs_runtime.Value {
	once_last__4043133652.Do(func() {
		cache_last__4043133652 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_last__4043133652(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_last__4043133652
}

var cache_length__3003998832 gopurs_runtime.Value
var once_length__3003998832 sync.Once
func Get_length__3003998832() gopurs_runtime.Value {
	once_length__3003998832.Do(func() {
		cache_length__3003998832 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int(acc_0.IntVal), gopurs_runtime.Int(1)).IntVal)
})
}), gopurs_runtime.Int(0))
	})
	return cache_length__3003998832
}

var cache_mapMaybe__3262563995 gopurs_runtime.Value
var once_mapMaybe__3262563995 sync.Once
func Get_mapMaybe__3262563995() gopurs_runtime.Value {
	once_mapMaybe__3262563995.Do(func() {
		cache_mapMaybe__3262563995 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe__3262563995(f_0_box)
})
	})
	return cache_mapMaybe__3262563995
}

var cache_mapMaybe__1486753757 gopurs_runtime.Value
var once_mapMaybe__1486753757 sync.Once
func Get_mapMaybe__1486753757() gopurs_runtime.Value {
	once_mapMaybe__1486753757.Do(func() {
		cache_mapMaybe__1486753757 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe__1486753757(f_0_box)
})
	})
	return cache_mapMaybe__1486753757
}

var cache_modifyAt__1886983628 gopurs_runtime.Value
var once_modifyAt__1886983628 sync.Once
func Get_modifyAt__1886983628() gopurs_runtime.Value {
	once_modifyAt__1886983628.Do(func() {
		cache_modifyAt__1886983628 = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modifyAt__1886983628(n_0_box.IntVal, f_1_box)
})
	})
	return cache_modifyAt__1886983628
}

var cache_nubBy__2103943131 gopurs_runtime.Value
var once_nubBy__2103943131 sync.Once
func Get_nubBy__2103943131() gopurs_runtime.Value {
	once_nubBy__2103943131.Do(func() {
		cache_nubBy__2103943131 = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_nubBy__2103943131(p_0_box)
})
	})
	return cache_nubBy__2103943131
}

var cache_nubByEq__3956095361 gopurs_runtime.Value
var once_nubByEq__3956095361 sync.Once
func Get_nubByEq__3956095361() gopurs_runtime.Value {
	once_nubByEq__3956095361.Do(func() {
		cache_nubByEq__3956095361 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_nubByEq__3956095361(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_nubByEq__3956095361
}

var cache_nubByEq__3655321914 gopurs_runtime.Value
var once_nubByEq__3655321914 sync.Once
func Get_nubByEq__3655321914() gopurs_runtime.Value {
	once_nubByEq__3655321914.Do(func() {
		cache_nubByEq__3655321914 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_nubByEq__3655321914(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_nubByEq__3655321914
}

var cache_partition__1623965204 gopurs_runtime.Value
var once_partition__1623965204 sync.Once
func Get_partition__1623965204() gopurs_runtime.Value {
	once_partition__1623965204.Do(func() {
		cache_partition__1623965204 = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_partition__1623965204(p_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_1_box))
})
	})
	return cache_partition__1623965204
}

var cache_reverse__1174136571 gopurs_runtime.Value
var once_reverse__1174136571 sync.Once
func Get_reverse__1174136571() gopurs_runtime.Value {
	once_reverse__1174136571.Do(func() {
		cache_reverse__1174136571 = func() gopurs_runtime.Value {
var go__go_0_0_22 gopurs_runtime.Value
go__go_0_0_22 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_1_loop_val)
var v1_2_loop gopurs_runtime.Value = v1_2_loop_val
go__go_0_0_22:
for {
if false { continue go__go_0_0_22 }
var v_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t1 gopurs_runtime.Value
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}
goto end_branch_1
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr != nil) {
v_1_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V0, v_1})})
v1_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V1)}
continue go__go_0_0_22
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}
}()
})
})
return gopurs_runtime.Apply(go__go_0_0_22, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}()
	})
	return cache_reverse__1174136571
}

var cache_reverse__4230102656 gopurs_runtime.Value
var once_reverse__4230102656 sync.Once
func Get_reverse__4230102656() gopurs_runtime.Value {
	once_reverse__4230102656.Do(func() {
		cache_reverse__4230102656 = func() gopurs_runtime.Value {
var go__go_0_0_23 gopurs_runtime.Value
go__go_0_0_23 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_1_loop_val)
var v1_2_loop gopurs_runtime.Value = v1_2_loop_val
go__go_0_0_23:
for {
if false { continue go__go_0_0_23 }
var v_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t1 gopurs_runtime.Value
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}
goto end_branch_1
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr != nil) {
v_1_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V0, v_1})})
v1_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V1)}
continue go__go_0_0_23
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}
}()
})
})
return gopurs_runtime.Apply(go__go_0_0_23, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}()
	})
	return cache_reverse__4230102656
}

var cache_singleton__707062261 gopurs_runtime.Value
var once_singleton__707062261 sync.Once
func Get_singleton__707062261() gopurs_runtime.Value {
	once_singleton__707062261.Do(func() {
		cache_singleton__707062261 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_singleton__707062261(a_0_box))}
})
	})
	return cache_singleton__707062261
}

var cache_singleton__3932757557 gopurs_runtime.Value
var once_singleton__3932757557 sync.Once
func Get_singleton__3932757557() gopurs_runtime.Value {
	once_singleton__3932757557.Do(func() {
		cache_singleton__3932757557 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_singleton__3932757557(a_0_box))}
})
	})
	return cache_singleton__3932757557
}

var cache_snoc__4290067657 gopurs_runtime.Value
var once_snoc__4290067657 sync.Once
func Get_snoc__4290067657() gopurs_runtime.Value {
	once_snoc__4290067657.Do(func() {
		cache_snoc__4290067657 = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_snoc__4290067657(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_0_box), x_1_box))}
})
	})
	return cache_snoc__4290067657
}

var cache_sortBy__2103943131 gopurs_runtime.Value
var once_sortBy__2103943131 sync.Once
func Get_sortBy__2103943131() gopurs_runtime.Value {
	once_sortBy__2103943131.Do(func() {
		cache_sortBy__2103943131 = gopurs_runtime.Func(func(cmp_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sortBy__2103943131(cmp_0_box)
})
	})
	return cache_sortBy__2103943131
}

var cache_sortBy__1502591776 gopurs_runtime.Value
var once_sortBy__1502591776 sync.Once
func Get_sortBy__1502591776() gopurs_runtime.Value {
	once_sortBy__1502591776.Do(func() {
		cache_sortBy__1502591776 = gopurs_runtime.Func(func(cmp_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sortBy__1502591776(cmp_0_box)
})
	})
	return cache_sortBy__1502591776
}

var cache_span__1918198736 gopurs_runtime.Value
var once_span__1918198736 sync.Once
func Get_span__1918198736() gopurs_runtime.Value {
	once_span__1918198736.Do(func() {
		cache_span__1918198736 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_span__1918198736(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1_box))
})
	})
	return cache_span__1918198736
}

var cache_span__799093643 gopurs_runtime.Value
var once_span__799093643 sync.Once
func Get_span__799093643() gopurs_runtime.Value {
	once_span__799093643.Do(func() {
		cache_span__799093643 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_span__799093643(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1_box))
})
	})
	return cache_span__799093643
}

var cache_take__551729751 gopurs_runtime.Value
var once_take__551729751 sync.Once
func Get_take__551729751() gopurs_runtime.Value {
	once_take__551729751.Do(func() {
		cache_take__551729751 = func() gopurs_runtime.Value {
var go__go_0_0_36 gopurs_runtime.Value
go__go_0_0_36 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_1_loop_val)
var v1_2_loop int64 = v1_2_loop_val.IntVal
var v2_3_loop gopurs_runtime.Value = v2_3_loop_val
go__go_0_0_36:
for {
if false { continue go__go_0_0_36 }
var v_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_1_loop
_ = v_1
var v1_2 int64 = v1_2_loop
_ = v1_2
var v2_3 gopurs_runtime.Value = v2_3_loop
_ = v2_3
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(gopurs_runtime.Int(v1_2), gopurs_runtime.Int(1))).IntVal) != (0) {
var go__go_4_1_37 gopurs_runtime.Value
go__go_4_1_37 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_37:
for {
if false { continue go__go_4_1_37 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_1_37
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}
}()
})
})
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_4_1_37, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)})))}
goto end_branch_5
} else {

}
}
{
if (v2_3.Type == 9 && v2_3.IntVal == 1358893437 && v2_3.UnsafePtr == nil) {
var go__go_4_3_38 gopurs_runtime.Value
go__go_4_3_38 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_3_38:
for {
if false { continue go__go_4_3_38 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t4 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_4
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_3_38
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t4))}
}
}()
})
})
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_4_3_38, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)})))}
goto end_branch_5
} else {

}
}
{
if (v2_3.Type == 9 && v2_3.IntVal == 1358893437 && v2_3.UnsafePtr != nil) {
v_1_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_3.UnsafePtr).V0, v_1})})
v1_2_loop = gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(v1_2), gopurs_runtime.Int(1)).IntVal
v2_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_3.UnsafePtr).V1)}
continue go__go_0_0_36
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t5))}
}
}()
})
})
})
return gopurs_runtime.Apply(go__go_0_0_36, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}()
	})
	return cache_take__551729751
}

var cache_takeWhile__2352021032 gopurs_runtime.Value
var once_takeWhile__2352021032 sync.Once
func Get_takeWhile__2352021032() gopurs_runtime.Value {
	once_takeWhile__2352021032.Do(func() {
		cache_takeWhile__2352021032 = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_takeWhile__2352021032(p_0_box)
})
	})
	return cache_takeWhile__2352021032
}

var cache_uncons__3009258782 gopurs_runtime.Value
var once_uncons__3009258782 sync.Once
func Get_uncons__3009258782() gopurs_runtime.Value {
	once_uncons__3009258782.Do(func() {
		cache_uncons__3009258782 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons__3009258782(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_uncons__3009258782
}

var cache_unionBy__588351261 gopurs_runtime.Value
var once_unionBy__588351261 sync.Once
func Get_unionBy__588351261() gopurs_runtime.Value {
	once_unionBy__588351261.Do(func() {
		cache_unionBy__588351261 = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_unionBy__588351261(eq_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](ys_2_box)))}
})
	})
	return cache_unionBy__588351261
}

var cache_unsnoc__2942606998 gopurs_runtime.Value
var once_unsnoc__2942606998 sync.Once
func Get_unsnoc__2942606998() gopurs_runtime.Value {
	once_unsnoc__2942606998.Do(func() {
		cache_unsnoc__2942606998 = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_unsnoc__2942606998(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](lst_0_box)))}
})
	})
	return cache_unsnoc__2942606998
}

var cache_updateAt__2634211748 gopurs_runtime.Value
var once_updateAt__2634211748 sync.Once
func Get_updateAt__2634211748() gopurs_runtime.Value {
	once_updateAt__2634211748.Do(func() {
		cache_updateAt__2634211748 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_updateAt__2634211748(v_0_box.IntVal, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_2_box)))}
})
	})
	return cache_updateAt__2634211748
}

var cache_zipWith__884793877 gopurs_runtime.Value
var once_zipWith__884793877 sync.Once
func Get_zipWith__884793877() gopurs_runtime.Value {
	once_zipWith__884793877.Do(func() {
		cache_zipWith__884793877 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_zipWith__884793877(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](ys_2_box)))}
})
	})
	return cache_zipWith__884793877
}

var cache_fromMaybe__430429096 gopurs_runtime.Value
var once_fromMaybe__430429096 sync.Once
func Get_fromMaybe__430429096() gopurs_runtime.Value {
	once_fromMaybe__430429096.Do(func() {
		cache_fromMaybe__430429096 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromMaybe__430429096(a_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_1_box))
})
	})
	return cache_fromMaybe__430429096
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

var cache_maybe__1246883617 gopurs_runtime.Value
var once_maybe__1246883617 sync.Once
func Get_maybe__1246883617() gopurs_runtime.Value {
	once_maybe__1246883617.Do(func() {
		cache_maybe__1246883617 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__1246883617(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__1246883617
}

var cache_alaF__4085337484 gopurs_runtime.Value
var once_alaF__4085337484 sync.Once
func Get_alaF__4085337484() gopurs_runtime.Value {
	once_alaF__4085337484.Do(func() {
		cache_alaF__4085337484 = gopurs_runtime.Func5(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, _dollar__unused_2_box gopurs_runtime.Value, _dollar__unused_3_box gopurs_runtime.Value, v_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alaF__4085337484(_dollar__unused_0_box, _dollar__unused_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_2_box), gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_3_box), v_4_box)
})
	})
	return cache_alaF__4085337484
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

var cache_sequence1__4078930642 gopurs_runtime.Value
var once_sequence1__4078930642 sync.Once
func Get_sequence1__4078930642() gopurs_runtime.Value {
	once_sequence1__4078930642.Do(func() {
		cache_sequence1__4078930642 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequence1__4078930642(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Traversable.Constructor_Traversable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sequence1__4078930642
}

var cache_sequence1__2726498490 gopurs_runtime.Value
var once_sequence1__2726498490 sync.Once
func Get_sequence1__2726498490() gopurs_runtime.Value {
	once_sequence1__2726498490.Do(func() {
		cache_sequence1__2726498490 = gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_traversable1NonEmptyList(), "sequence1")
	})
	return cache_sequence1__2726498490
}

var cache_traverse1__157785005 gopurs_runtime.Value
var once_traverse1__157785005 sync.Once
func Get_traverse1__157785005() gopurs_runtime.Value {
	once_traverse1__157785005.Do(func() {
		cache_traverse1__157785005 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse1__157785005(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Traversable.Constructor_Traversable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_traverse1__157785005
}

var cache_traverse1__42886725 gopurs_runtime.Value
var once_traverse1__42886725 sync.Once
func Get_traverse1__42886725() gopurs_runtime.Value {
	once_traverse1__42886725.Do(func() {
		cache_traverse1__42886725 = gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_traversable1NonEmptyList(), "traverse1")
	})
	return cache_traverse1__42886725
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

var cache_append__2013893496 gopurs_runtime.Value
var once_append__2013893496 sync.Once
func Get_append__2013893496() gopurs_runtime.Value {
	once_append__2013893496.Do(func() {
		cache_append__2013893496 = gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_semigroupList(), "append")
	})
	return cache_append__2013893496
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

var cache_unfoldr__3827943605 gopurs_runtime.Value
var once_unfoldr__3827943605 sync.Once
func Get_unfoldr__3827943605() gopurs_runtime.Value {
	once_unfoldr__3827943605.Do(func() {
		cache_unfoldr__3827943605 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr__3827943605(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr__3827943605
}

var cache_unsafeCrashWith__69763299 gopurs_runtime.Value
var once_unsafeCrashWith__69763299 sync.Once
func Get_unsafeCrashWith__69763299() gopurs_runtime.Value {
	once_unsafeCrashWith__69763299.Do(func() {
		cache_unsafeCrashWith__69763299 = gopurs_runtime.Func(func(msg_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeCrashWith__69763299(msg_0_box.StrVal())
})
	})
	return cache_unsafeCrashWith__69763299
}

var cache_unsafeCrashWith__2941848695 gopurs_runtime.Value
var once_unsafeCrashWith__2941848695 sync.Once
func Get_unsafeCrashWith__2941848695() gopurs_runtime.Value {
	once_unsafeCrashWith__2941848695.Do(func() {
		cache_unsafeCrashWith__2941848695 = gopurs_runtime.Func(func(msg_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeCrashWith__2941848695(msg_0_box.StrVal())
})
	})
	return cache_unsafeCrashWith__2941848695
}

var cache_unsafePartial__1306634845 gopurs_runtime.Value
var once_unsafePartial__1306634845 sync.Once
func Get_unsafePartial__1306634845() gopurs_runtime.Value {
	once_unsafePartial__1306634845.Do(func() {
		cache_unsafePartial__1306634845 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1306634845
}

var cache_crashWith__1894115486 gopurs_runtime.Value
var once_crashWith__1894115486 sync.Once
func Get_crashWith__1894115486() gopurs_runtime.Value {
	once_crashWith__1894115486.Do(func() {
		cache_crashWith__1894115486 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_crashWith__1894115486(_dollar__unused_0_box)
})
	})
	return cache_crashWith__1894115486
}

func Call_identity(x_0_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var x_0 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = x_0_loop
_ = x_0
return x_0
}

func Call_zipWith(f_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value], v1_2_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
var v1_2 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v1_2_loop
_ = v1_2
return gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v1_2)}.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(pkg_Data_List.Get_zipWith(), f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v1_2)}.UnsafePtr).V1))})))}})})
}

func Call_zipWithA(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{}))
_ = Apply0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_traversable1NonEmptyList(), "sequence1"), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Apply0_1_0)}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_zipWith(f_2, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](xs_3), gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](ys_4)))}))})
})
})
})
}

func Call_wrappedOperation2(name_0_loop string, f_1_loop gopurs_runtime.Value, v_2_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value], v1_3_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var name_0 string = name_0_loop
_ = name_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v1_3_loop
_ = v1_3
v2_4_0 := gopurs_runtime.Apply2(f_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V1)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)})})
_ = v2_4_0
var __t2 gopurs_runtime.Value
{
if (v2_4_0.Type == 9 && v2_4_0.IntVal == 1358893437 && v2_4_0.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_4_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_4_0.UnsafePtr).V1)}})}
goto end_branch_2
} else {

}
}
{
if (v2_4_0.Type == 9 && v2_4_0.IntVal == 1358893437 && v2_4_0.UnsafePtr == nil) {
__local_var_5_1 := gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str("Impossible: empty list in NonEmptyList "), gopurs_runtime.Str(name_0)).StrVal()
_ = __local_var_5_1
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str(__local_var_5_1))
}))))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](__t2)
}

func Call_wrappedOperation(name_0_loop string, f_1_loop gopurs_runtime.Value, v_2_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var name_0 string = name_0_loop
_ = name_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_2_loop
_ = v_2
v1_3_0 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V1)})}))
_ = v1_3_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3_0)}.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3_0)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3_0)}.UnsafePtr).V1)}})}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3_0)}.UnsafePtr == nil) {
__local_var_4_1 := gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str("Impossible: empty list in NonEmptyList "), gopurs_runtime.Str(name_0)).StrVal()
_ = __local_var_4_1
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str(__local_var_4_1))
}))))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](__t2)
}

func Call_updateAt(i_0_loop int64, a_1_loop gopurs_runtime.Value, v_2_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
var i_0 int64 = i_0_loop
_ = i_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var v_2 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_eq__2843686287(), gopurs_runtime.Int(i_0), gopurs_runtime.Int(0)).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V1))}})}})}
goto end_branch_1
} else {

}
}
{
__local_var_3_0 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V0
_ = __local_var_3_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_3_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](x_4))}})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(pkg_Data_List.Get_updateAt(), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(i_0), gopurs_runtime.Int(1)).IntVal), a_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V1))})))})))}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](__t1)
}

func Call_unzip(ts_0_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *pkg_Data_Tuple.Constructor_Tuple[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value], *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
var ts_0 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = ts_0_loop
_ = ts_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value], *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_map__91618268(), pkg_Data_Tuple.Get_fst(), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(ts_0)})))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_map__91618268(), pkg_Data_Tuple.Get_snd(), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(ts_0)})))}})})
}

func Call_unsnoc(v_0_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
v1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_List.Get_unsnoc(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1))}))
_ = v1_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1_0)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1_0)}.UnsafePtr != nil) {
__t1 = gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1_0)}.UnsafePtr).V0, "init"))})}, gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1_0)}.UnsafePtr).V0, "last"))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}

func Call_unionBy(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply2(Get_wrappedOperation2(), gopurs_runtime.Str("unionBy"), gopurs_runtime.Apply(pkg_Data_List.Get_unionBy(), x_0))
}

func Call_union(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply2(Get_wrappedOperation2(), gopurs_runtime.Str("union"), gopurs_runtime.Apply(pkg_Data_List.Get_union(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_0)}))
}

func Call_uncons(v_0_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.RecordDict2("head", "tail", (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1))})
}

func Call_toList(v_0_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)})})
}

func Call_toUnfoldable(dictUnfoldable_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
__local_var_1_0 := gopurs_runtime.Apply(dictUnfoldable_0.V1, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_1))}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1358893437 && __t_tag_1.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_1))}
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1358893437 && __t_tag_2.UnsafePtr != nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("head", "tail", (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(xs_1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(xs_1.UnsafePtr).V1)})})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(rec_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(rec_2, "head"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(rec_2, "tail")))}})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t3))})))}
}))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1)})})
})
}

func Call_tail(v_0_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)
}

func Call_sortBy(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply2(Get_wrappedOperation(), gopurs_runtime.Str("sortBy"), gopurs_runtime.Apply(pkg_Data_List.Get_sortBy(), x_0))
}

func Call_sort(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.V1
_ = compare_1_0
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_wrappedOperation("sortBy", gopurs_runtime.Apply(pkg_Data_List.Get_sortBy(), compare_1_0), gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](xs_2)))}))}
})
}

func Call_snoc(v_0_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value], y_1_loop gopurs_runtime.Value) *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var v_0 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
var y_1 gopurs_runtime.Value = y_1_loop
_ = y_1
return gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, y_1, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1))})))}})})
}

func Call_singleton(x_0_loop gopurs_runtime.Value) *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_0, gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_plusList(), "empty")})})
}

func Call_snoc_prime(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], v1_1_loop gopurs_runtime.Value) *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v1_1, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)})))}})}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, v1_1, gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_plusList(), "empty")})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](__t0)
}

func Call_nubEq(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply2(Get_wrappedOperation(), gopurs_runtime.Str("nubEq"), gopurs_runtime.Apply(pkg_Data_List.Get_nubEq(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_0)}))
}

func Call_nubByEq(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply2(Get_wrappedOperation(), gopurs_runtime.Str("nubByEq"), gopurs_runtime.Apply(pkg_Data_List.Get_nubByEq(), x_0))
}

func Call_nubBy(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply2(Get_wrappedOperation(), gopurs_runtime.Str("nubBy"), gopurs_runtime.Apply(pkg_Data_List.Get_nubBy(), x_0))
}

func Call_nub(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply2(Get_wrappedOperation(), gopurs_runtime.Str("nub"), gopurs_runtime.Apply(pkg_Data_List.Get_nubBy(), dictOrd_0.V1))
}

func Call_modifyAt(i_0_loop int64, f_1_loop gopurs_runtime.Value, v_2_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
var i_0 int64 = i_0_loop
_ = i_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_eq__2843686287(), gopurs_runtime.Int(i_0), gopurs_runtime.Int(0)).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V1))}})}})}
goto end_branch_1
} else {

}
}
{
__local_var_3_0 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V0
_ = __local_var_3_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_3_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](x_4))}})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(pkg_Data_List.Get_alterAt(), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(i_0), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, x_4)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V1))})))})))}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](__t1)
}

func Call_mapMaybe(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_List.Get_mapMaybe(), x_0)
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)})})
})
}

func Call_partition(x_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var v_1 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(pkg_Data_List.Get_partition(), x_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)})})
}

func Call_span(x_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var v_1 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(pkg_Data_List.Get_span(), x_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)})})
}

func Call_take(x_0_loop int64) gopurs_runtime.Value {
var x_0 int64 = x_0_loop
_ = x_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_List.Get_take(), gopurs_runtime.Int(x_0))
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)})})
})
}

func Call_takeWhile(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var go__go_1_1_0 gopurs_runtime.Value
go__go_1_1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_1_0:
for {
if false { continue go__go_1_1_0 }
var v_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t4 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
if ((v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil)) && ((gopurs_runtime.Apply(x_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0).IntVal) != (0)) {
v_2_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, v_2})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_1_0
__t4 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
goto end_branch_4
} else {

}
}
{
var go__go_4_2_1 gopurs_runtime.Value
go__go_4_2_1 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_2_1:
for {
if false { continue go__go_4_2_1 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t3 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_3
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_2_1
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t3))}
}
}()
})
})
__t4 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_4_2_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}))
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t4)}
}
}()
})
})
__local_var_1_0 := gopurs_runtime.Apply(go__go_1_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)})})
})
}

func Call_length(v_0_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) int64 {
var v_0 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int(1), gopurs_runtime.Int(gopurs_runtime.Apply(pkg_Data_List.Get_length(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1))}).IntVal)).IntVal
}

func Call_last(v_0_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t4 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1))}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1358893437 && __t_tag_1.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1.UnsafePtr).V1)}
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1358893437 && __t_tag_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1.UnsafePtr).V0})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_List.Get_last(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1.UnsafePtr).V1)})))}
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_4:
__local_var_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t4)
_ = __local_var_1_0
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_1_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_1_0)}.UnsafePtr == nil) {
__t5 = (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_1_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_1_0)}.UnsafePtr != nil) {
__t5 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_1_0)}.UnsafePtr).V0
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

func Call_intersectBy(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply2(Get_wrappedOperation2(), gopurs_runtime.Str("intersectBy"), gopurs_runtime.Apply(pkg_Data_List.Get_intersectBy(), x_0))
}

func Call_intersect(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply2(Get_wrappedOperation2(), gopurs_runtime.Str("intersect"), gopurs_runtime.Apply(pkg_Data_List.Get_intersect(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_0)}))
}

func Call_insertAt(i_0_loop int64, a_1_loop gopurs_runtime.Value, v_2_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
var i_0 int64 = i_0_loop
_ = i_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var v_2 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_eq__2843686287(), gopurs_runtime.Int(i_0), gopurs_runtime.Int(0)).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V1)})}})}})}
goto end_branch_1
} else {

}
}
{
__local_var_3_0 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V0
_ = __local_var_3_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_3_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](x_4))}})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(pkg_Data_List.Get_insertAt(), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(i_0), gopurs_runtime.Int(1)).IntVal), a_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V1))})))})))}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](__t1)
}

func Call_init(v_0_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
__local_var_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v_1, "init")))}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_List.Get_unsnoc(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1))})))}))
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_1_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_1_0)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_1_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_1_0)}.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_1_0)}.UnsafePtr).V0)})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1)
}

func Call_index(v_0_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value], i_1_loop int64) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
var i_1 int64 = i_1_loop
_ = i_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_eq__2843686287(), gopurs_runtime.Int(i_1), gopurs_runtime.Int(0)).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(pkg_Data_List.Get_index(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1))}, gopurs_runtime.Int(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(i_1), gopurs_runtime.Int(1)).IntVal))))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0)
}

func Call_head(v_0_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
}

func Call_groupBy(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply2(Get_wrappedOperation(), gopurs_runtime.Str("groupBy"), gopurs_runtime.Apply(pkg_Data_List.Get_groupBy(), x_0))
}

func Call_groupAllBy(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply2(Get_wrappedOperation(), gopurs_runtime.Str("groupAllBy"), gopurs_runtime.Apply(pkg_Data_List.Get_groupAllBy(), x_0))
}

func Call_groupAll(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply2(Get_wrappedOperation(), gopurs_runtime.Str("groupAll"), gopurs_runtime.Apply(pkg_Data_List.Get_groupAll(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}))
}

func Call_group(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply2(Get_wrappedOperation(), gopurs_runtime.Str("group"), gopurs_runtime.Apply(pkg_Data_List.Get_group(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_0)}))
}

func Call_fromList(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}})}})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](__t0)
}

func Call_fromFoldable(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
__local_var_1_0 := gopurs_runtime.Apply2(dictFoldable_0.V2, pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
_ = __local_var_3_1
var __t2 gopurs_runtime.Value
{
if (__local_var_3_1.Type == 9 && __local_var_3_1.IntVal == 1358893437 && __local_var_3_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
if (__local_var_3_1.Type == 9 && __local_var_3_1.IntVal == 1358893437 && __local_var_3_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(__local_var_3_1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(__local_var_3_1.UnsafePtr).V1)}})}})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](__t2))}
})
}

func Call_foldM(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Bind1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_1 := (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1
_ = __local_var_5_1
return gopurs_runtime.Apply2(Bind1_1_0.V1, gopurs_runtime.Apply2(f_2, b_3, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Func(func(b_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(pkg_Data_List.Get_foldM(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)}, f_2, b_prime_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__local_var_5_1))})
}))
})
})
})
}

func Call_findLastIndex(f_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[int64] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
v1_2_0 := gopurs_runtime.Apply2(pkg_Data_List.Get_findLastIndex(), f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1))})
_ = v1_2_0
var __t2 gopurs_runtime.Value
{
if (v1_2_0.Type == 9 && v1_2_0.IntVal == 930809136 && v1_2_0.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_2_0.UnsafePtr).V0.IntVal), gopurs_runtime.Int(1)).IntVal)})}
goto end_branch_2
} else {

}
}
{
if (v1_2_0.Type == 9 && v1_2_0.IntVal == 930809136 && v1_2_0.UnsafePtr == nil) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(0)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](__t2)
}

func Call_findIndex(f_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[int64] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply2(Get_map__291265340(), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int(v1_2.IntVal), gopurs_runtime.Int(1)).IntVal)
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply2(pkg_Data_List.Get_findIndex(), f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1))})))})))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](__t0)
}

func Call_filterM(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_List.Get_filterM(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)})})
})
})
}

func Call_filter(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_List.Get_filter(), x_0)
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)})})
})
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

func Call_dropWhile(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var go__go_1_1_2 gopurs_runtime.Value
go__go_1_1_2 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2_loop_val)
go__go_1_1_2:
for {
if false { continue go__go_1_1_2 }
var v_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var __t2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
if ((gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr != nil)) && ((gopurs_runtime.Apply(x_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V0).IntVal) != (0)) {
v_2_loop = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V1
continue go__go_1_1_2
__t2 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
__t2 = v_2
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t2)}
}
}()
})
__local_var_1_0 := go__go_1_1_2
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)})})
})
}

func Call_drop(x_0_loop int64, v_1_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var x_0 int64 = x_0_loop
_ = x_0
var v_1 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(pkg_Data_List.Get_drop(), gopurs_runtime.Int(x_0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)})}))
}

func Call_cons_prime(x_0_loop gopurs_runtime.Value, xs_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = xs_1_loop
_ = xs_1
return gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}})})
}

func Call_cons(y_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var y_0 gopurs_runtime.Value = y_0_loop
_ = y_0
var v_1 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, y_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)})}})})
}

func Call_concatMap(b_0_loop gopurs_runtime.Value, a_1_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = a_1_loop
_ = a_1
return gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_bindNonEmptyList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(a_1)}, b_0))
}

func Call_concat(v_0_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var v_0 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_bind__2389430209(), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}, Get_identity()))
}

func Call_catMaybes(v_0_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]] = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_List.Get_catMaybes(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)})}))
}

func Call_appendFoldable(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], v_1_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value], ys_2_loop gopurs_runtime.Value) *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var v_1 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
return gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(dictFoldable_0.V2, pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, ys_2)))})))}})})
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_lift2__2762258480(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
})
})
}

func Call_lift2__3213187376(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
})
})
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
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

func Call_any__4179648253(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], dictHeytingAlgebra_1_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictHeytingAlgebra_1 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
semigroupDisj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictHeytingAlgebra_1.V1, v_2, v1_3)
})
}))
_ = semigroupDisj1_2_0
return gopurs_runtime.Apply(dictFoldable_0.V0, gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_2_0
}), dictHeytingAlgebra_1.V2))
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

func Call_foldr__2829803163(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
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

func Call_flip__3563101792(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__1833071808(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__1570464832(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1256368628(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1483545076(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mapFlipped__4215217780(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, f_2, fa_1)
}

func Call_mapFlipped__2919806324(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, f_2, fa_1)
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

func Call_fromZipper__1019554324(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t6 = v1_1
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t5 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 1304506903) {
__t5 = gopurs_runtime.Apply2(pkg_Data_List_Internal.Get_fromZipper(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value]{1, v1_1, (*pkg_Data_List_Internal.Constructor_TwoLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_TwoLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1})})
goto end_branch_5
} else {

}
}
{
var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 2884341868) {
__t5 = gopurs_runtime.Apply2(pkg_Data_List_Internal.Get_fromZipper(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_TwoRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_TwoRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1, v1_1})})
goto end_branch_5
} else {

}
}
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2195694037) {
__t5 = gopurs_runtime.Apply2(pkg_Data_List_Internal.Get_fromZipper(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value]{1, v1_1, (*pkg_Data_List_Internal.Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V2, (*pkg_Data_List_Internal.Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V3})})
goto end_branch_5
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 1584522659) {
__t5 = gopurs_runtime.Apply2(pkg_Data_List_Internal.Get_fromZipper(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1, v1_1, (*pkg_Data_List_Internal.Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V2, (*pkg_Data_List_Internal.Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V3})})
goto end_branch_5
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 3952671150) {
__t5 = gopurs_runtime.Apply2(pkg_Data_List_Internal.Get_fromZipper(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V2, (*pkg_Data_List_Internal.Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V3, v1_1})})
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}

func Call_insertAndLookupBy__3244745033(comp_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value, orig_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var orig_2 gopurs_runtime.Value = orig_2_loop
_ = orig_2
var up_3_0_3 gopurs_runtime.Value
up_3_0_3 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
up_3_0_3:
for {
if false { continue up_3_0_3 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t7 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V2})}
goto end_branch_7
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
var __t6 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1304506903) {
__t6 = gopurs_runtime.Apply2(pkg_Data_List_Internal.Get_fromZipper(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V2, (*pkg_Data_List_Internal.Constructor_TwoLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_TwoLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V1})})
goto end_branch_6
} else {

}
}
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2884341868) {
__t6 = gopurs_runtime.Apply2(pkg_Data_List_Internal.Get_fromZipper(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_TwoRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_TwoRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V2})})
goto end_branch_6
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2195694037) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V2})}, (*pkg_Data_List_Internal.Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V2, (*pkg_Data_List_Internal.Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V3})}})}
continue up_3_0_3
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 1584522659) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V0})}, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V2, (*pkg_Data_List_Internal.Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V2, (*pkg_Data_List_Internal.Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V3})}})}
continue up_3_0_3
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 3952671150) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V2})}, (*pkg_Data_List_Internal.Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V2})}})}
continue up_3_0_3
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
__t7 = __t6
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}
}()
})
})
var down_4_8_4 gopurs_runtime.Value
down_4_8_4 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
down_4_8_4:
for {
if false { continue down_4_8_4 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t15 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 2764020654) {
__t15 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(false), gopurs_runtime.Apply2(up_3_0_3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}, gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_KickUp[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)}, k_1, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)}})}))
goto end_branch_15
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1177901036) {
v2_7_9 := gopurs_runtime.Apply2(comp_0, k_1, (*pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)
_ = v2_7_9
var __t10 gopurs_runtime.Value
{
if (uint32(v2_7_9.IntVal) == 902936544) {
__t10 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(true), orig_2)
goto end_branch_10
} else {

}
}
{
if (uint32(v2_7_9.IntVal) == 1527465420) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1304506903, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_TwoLeft[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V2})}, v_5})})
v1_6_loop = (*pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
continue down_4_8_4
__t10 = gopurs_runtime.Value{}
goto end_branch_10
} else {

}
}
{
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2884341868, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_TwoRight[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V1})}, v_5})})
v1_6_loop = (*pkg_Data_List_Internal.Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V2
continue down_4_8_4
__t10 = gopurs_runtime.Value{}
}
end_branch_10:
__t15 = __t10
goto end_branch_15
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1064476974) {
v2_7_11 := gopurs_runtime.Apply2(comp_0, k_1, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)
_ = v2_7_11
var __t14 gopurs_runtime.Value
{
if (uint32(v2_7_11.IntVal) == 902936544) {
__t14 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(true), orig_2)
goto end_branch_14
} else {

}
}
{
v3_8_12 := gopurs_runtime.Apply2(comp_0, k_1, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V3)
_ = v3_8_12
var __t13 gopurs_runtime.Value
{
if (uint32(v3_8_12.IntVal) == 902936544) {
__t13 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(true), orig_2)
goto end_branch_13
} else {

}
}
{
if (uint32(v2_7_11.IntVal) == 1527465420) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2195694037, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_ThreeLeft[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V2, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V3, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V4})}, v_5})})
v1_6_loop = (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
continue down_4_8_4
__t13 = gopurs_runtime.Value{}
goto end_branch_13
} else {

}
}
{
if ((uint32(v2_7_11.IntVal) == 380165415)) && ((uint32(v3_8_12.IntVal) == 1527465420)) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1584522659, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_ThreeMiddle[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V3, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V4})}, v_5})})
v1_6_loop = (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V2
continue down_4_8_4
__t13 = gopurs_runtime.Value{}
goto end_branch_13
} else {

}
}
{
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3952671150, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Constructor_ThreeRight[gopurs_runtime.Value]{1, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V2, (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V3})}, v_5})})
v1_6_loop = (*pkg_Data_List_Internal.Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V4
continue down_4_8_4
__t13 = gopurs_runtime.Value{}
}
end_branch_13:
__t14 = __t13
}
end_branch_14:
__t15 = __t14
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
}
}()
})
})
return gopurs_runtime.Apply2(down_4_8_4, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, orig_2)
}

func Call_findIndex__1077906787(f_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[int64] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int(v1_2.IntVal), gopurs_runtime.Int(1)).IntVal)
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply2(pkg_Data_List.Get_findIndex(), f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1))})))})))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](__t0)
}

func Call_findLastIndex__1077906787(f_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[int64] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
v1_2_0 := gopurs_runtime.Apply2(pkg_Data_List.Get_findLastIndex(), f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1))})
_ = v1_2_0
var __t2 gopurs_runtime.Value
{
if (v1_2_0.Type == 9 && v1_2_0.IntVal == 930809136 && v1_2_0.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_2_0.UnsafePtr).V0.IntVal), gopurs_runtime.Int(1)).IntVal)})}
goto end_branch_2
} else {

}
}
{
if (v1_2_0.Type == 9 && v1_2_0.IntVal == 930809136 && v1_2_0.UnsafePtr == nil) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(0)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](__t2)
}

func Call_fromList__2793484475(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}})}})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](__t0)
}

func Call_lift__1094295643(f_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)})}))
}

func Call_lift__184746683(f_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)})})
}

func Call_lift__243257371(f_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)})})
}

func Call_lift__2580243643(f_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)})})
}

func Call_lift__3667893565(f_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)})})
}

func Call_lift__1609676349(f_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)})})
}

func Call_singleton__3219659348(x_0_loop gopurs_runtime.Value) *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_0, gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_plusList(), "empty")})})
}

func Call_sortBy__2726669792(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply2(Get_wrappedOperation(), gopurs_runtime.Str("sortBy"), gopurs_runtime.Apply(pkg_Data_List.Get_sortBy(), x_0))
}

func Call_toList__2859885498(v_0_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)})})
}

func Call_wrappedOperation__2422223318(name_0_loop string, f_1_loop gopurs_runtime.Value, v_2_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var name_0 string = name_0_loop
_ = name_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_2_loop
_ = v_2
v1_3_0 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V1)})}))
_ = v1_3_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3_0)}.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3_0)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3_0)}.UnsafePtr).V1)}})}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3_0)}.UnsafePtr == nil) {
__local_var_4_1 := gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str("Impossible: empty list in NonEmptyList "), gopurs_runtime.Str(name_0)).StrVal()
_ = __local_var_4_1
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str(__local_var_4_1))
}))))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](__t2)
}

func Call_wrappedOperation__1897406582(name_0_loop string, f_1_loop gopurs_runtime.Value, v_2_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var name_0 string = name_0_loop
_ = name_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_2_loop
_ = v_2
v1_3_0 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V1)})}))
_ = v1_3_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3_0)}.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3_0)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3_0)}.UnsafePtr).V1)}})}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3_0)}.UnsafePtr == nil) {
__local_var_4_1 := gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str("Impossible: empty list in NonEmptyList "), gopurs_runtime.Str(name_0)).StrVal()
_ = __local_var_4_1
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str(__local_var_4_1))
}))))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](__t2)
}

func Call_wrappedOperation2__4016046219(name_0_loop string, f_1_loop gopurs_runtime.Value, v_2_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value], v1_3_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var name_0 string = name_0_loop
_ = name_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v1_3_loop
_ = v1_3
v2_4_0 := gopurs_runtime.Apply2(f_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V1)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)})})
_ = v2_4_0
var __t2 gopurs_runtime.Value
{
if (v2_4_0.Type == 9 && v2_4_0.IntVal == 1358893437 && v2_4_0.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_4_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_4_0.UnsafePtr).V1)}})}
goto end_branch_2
} else {

}
}
{
if (v2_4_0.Type == 9 && v2_4_0.IntVal == 1358893437 && v2_4_0.UnsafePtr == nil) {
__local_var_5_1 := gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str("Impossible: empty list in NonEmptyList "), gopurs_runtime.Str(name_0)).StrVal()
_ = __local_var_5_1
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str(__local_var_5_1))
}))))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](__t2)
}

func Call_zipWith__1818047060(f_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value], v1_2_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
var v1_2 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v1_2_loop
_ = v1_2
return gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v1_2)}.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(pkg_Data_List.Get_zipWith(), f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v1_2)}.UnsafePtr).V1))})))}})})
}

func Call_zipWith__716394740(f_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value], v1_2_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
var v1_2 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v1_2_loop
_ = v1_2
return gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v1_2)}.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(pkg_Data_List.Get_zipWith(), f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v1_2)}.UnsafePtr).V1))})))}})})
}

func Call_zipWith__3875859348(f_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value], v1_2_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
var v1_2 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v1_2_loop
_ = v1_2
return gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v1_2)}.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(pkg_Data_List.Get_zipWith(), f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v1_2)}.UnsafePtr).V1))})))}})})
}

func Call_listMap__4135416762(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var chunkedRevMap_1_0_7 gopurs_runtime.Value
chunkedRevMap_1_0_7 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](v_2_loop_val)
var v1_3_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_3_loop_val)
chunkedRevMap_1_0_7:
for {
if false { continue chunkedRevMap_1_0_7 }
var v_2 *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = v_2_loop
_ = v_2
var v1_3 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t19 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
var __t_and_18 bool = false
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr != nil) {

var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
var __t_and_17 bool = false
if (__t_tag_15.Type == 9 && __t_tag_15.IntVal == 1358893437 && __t_tag_15.UnsafePtr != nil) {

var __t_tag_16 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}
__t_and_17 = (__t_tag_16.Type == 9 && __t_tag_16.IntVal == 1358893437 && __t_tag_16.UnsafePtr != nil)
}
__t_and_18 = __t_and_17
}
if __t_and_18 {
v_2_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)})})})
v1_3_loop = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr).V1
continue chunkedRevMap_1_0_7
__t19 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
goto end_branch_19
} else {

}
}
{
var reverseUnrolledMap_4_1_8 gopurs_runtime.Value
reverseUnrolledMap_4_1_8 = gopurs_runtime.Func(func(v2_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v2_5_loop gopurs_runtime.Value = v2_5_loop_val
var v3_6_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v3_6_loop_val)
reverseUnrolledMap_4_1_8:
for {
if false { continue reverseUnrolledMap_4_1_8 }
var v2_5 gopurs_runtime.Value = v2_5_loop
_ = v2_5
var v3_6 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v3_6_loop
_ = v3_6
var __t8 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
var __t_and_7 bool = false
if (v2_5.Type == 9 && v2_5.IntVal == 1358893437 && v2_5.UnsafePtr != nil) {

var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0
var __t_and_6 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1358893437 && __t_tag_2.UnsafePtr != nil) {

var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}
var __t_and_5 bool = false
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 1358893437 && __t_tag_3.UnsafePtr != nil) {

var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V1)}
__t_and_5 = (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 1358893437 && __t_tag_4.UnsafePtr != nil)
}
__t_and_6 = __t_and_5
}
__t_and_7 = __t_and_6
}
if __t_and_7 {
v2_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V1)}
v3_6_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr).V0), v3_6})})})})})})
continue reverseUnrolledMap_4_1_8
__t8 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
goto end_branch_8
} else {

}
}
{
__t8 = v3_6
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t8)}
}
}()
})
})
var __t14 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr != nil) {
var __t13 gopurs_runtime.Value
{
var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 1358893437 && __t_tag_9.UnsafePtr != nil) {
var __t11 gopurs_runtime.Value
{
var __t_tag_10 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}
if (__t_tag_10.Type == 9 && __t_tag_10.IntVal == 1358893437 && __t_tag_10.UnsafePtr == nil) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})})})}
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_11:
__t13 = __t11
goto end_branch_13
} else {

}
}
{
var __t_tag_12 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 1358893437 && __t_tag_12.UnsafePtr == nil) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})}
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_13:
__t14 = __t13
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_14:
__t19 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(reverseUnrolledMap_4_1_8, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t14))}))
}
end_branch_19:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t19)}
}
}()
})
})
return gopurs_runtime.Apply(chunkedRevMap_1_0_7, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}

func Call_nelCons__195558898(a_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)})}})})
}

func Call_toList__2859885498(v_0_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 *pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)})})
}

func Call_alterAt__3453373293(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
if (v_0) == (0) {
v3_3_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0))
_ = v3_3_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v3_3_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v3_3_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v3_3_1)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1)}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v3_3_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v3_3_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v3_3_1)}.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v3_3_1)}.UnsafePtr).V0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}})}
goto end_branch_3
} else {

}
}
{
__local_var_3_0 := (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0
_ = __local_var_3_0
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v3_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_3_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v3_4)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(pkg_Data_List.Get_alterAt(), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal), v1_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1)})))})))}
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_4:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t4)
}

func Call_deleteBy__697302515(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(v_0, v1_1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1)}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(pkg_Data_List.Get_deleteBy(), v_0, v1_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1)}))})}
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1)
}

func Call_drop__551729751(v_0_loop int64, v1_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(gopurs_runtime.Int(v_0), gopurs_runtime.Int(1))).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(pkg_Data_List.Get_drop(), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t0)
}

func Call_drop__1836090668(v_0_loop int64, v1_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(gopurs_runtime.Int(v_0), gopurs_runtime.Int(1))).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(pkg_Data_List.Get_drop(), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t0)
}

func Call_dropWhile__2352021032(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_9 gopurs_runtime.Value
go__go_1_0_9 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2_loop_val)
go__go_1_0_9:
for {
if false { continue go__go_1_0_9 }
var v_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var __t1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
if ((gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr != nil)) && ((gopurs_runtime.Apply(p_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V0).IntVal) != (0)) {
v_2_loop = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}.UnsafePtr).V1
continue go__go_1_0_9
__t1 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = v_2
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
return go__go_1_0_9
}

func Call_filter__2352021032(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_10 gopurs_runtime.Value
go__go_1_0_10 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_10:
for {
if false { continue go__go_1_0_10 }
var v_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t4 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
var go__go_4_1_11 gopurs_runtime.Value
go__go_4_1_11 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_11:
for {
if false { continue go__go_4_1_11 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_1_11
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}
}()
})
})
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_4_1_11, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)})))}
goto end_branch_4
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
var __t3 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply(p_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0).IntVal) != (0) {
v_2_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, v_2})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_10
__t3 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
v_2_loop = v_2
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_10
__t3 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t4))}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_10, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}

func Call_filter__1617261107(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_12 gopurs_runtime.Value
go__go_1_0_12 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_12:
for {
if false { continue go__go_1_0_12 }
var v_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t4 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
var go__go_4_1_13 gopurs_runtime.Value
go__go_4_1_13 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_13:
for {
if false { continue go__go_4_1_13 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_1_13
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}
}()
})
})
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_4_1_13, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)})))}
goto end_branch_4
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
var __t3 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply(p_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0).IntVal) != (0) {
v_2_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, v_2})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_12
__t3 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
v_2_loop = v_2
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_12
__t3 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t4))}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_12, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}

func Call_findIndex__2366045378(fn_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var go__go_1_0_14 gopurs_runtime.Value
go__go_1_0_14 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop int64 = v_2_loop_val.IntVal
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_14:
for {
if false { continue go__go_1_0_14 }
var v_2 int64 = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t2 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(fn_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(v_2)})}
goto end_branch_1
} else {

}
}
{
v_2_loop = gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Int(v_2), gopurs_runtime.Int(1)).IntVal
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_14
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Value{}))}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](__t2))}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_14, gopurs_runtime.Int(0))
}

func Call_findLastIndex__2366045378(fn_0_loop gopurs_runtime.Value, xs_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[int64] {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var xs_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = xs_1_loop
_ = xs_1
var go__go_2_0_15 gopurs_runtime.Value
go__go_2_0_15 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_3_loop_val)
var v1_4_loop gopurs_runtime.Value = v1_4_loop_val
go__go_2_0_15:
for {
if false { continue go__go_2_0_15 }
var v_3 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_3_loop
_ = v_3
var v1_4 gopurs_runtime.Value = v1_4_loop
_ = v1_4
var __t1 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_3)}
goto end_branch_1
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr != nil) {
v_3_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V0, v_3})})
v1_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V1)}
continue go__go_2_0_15
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(gopurs_runtime.Apply(pkg_Data_List.Get_length(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}).IntVal), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Int(v_2.IntVal)).IntVal)
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply2(pkg_Data_List.Get_findIndex(), fn_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_2_0_15, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)})))})))}))
}

func Call_foldM__3577257629(dictMonad_0_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 1358893437 && v2_5.UnsafePtr == nil) {
__t3 = gopurs_runtime.Apply(Applicative0_1_0.V1, v1_4)
goto end_branch_3
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 1358893437 && v2_5.UnsafePtr != nil) {
__local_var_6_2 := (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V1
_ = __local_var_6_2
__t3 = gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(v_3, v1_4, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0), gopurs_runtime.Func(func(b_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(pkg_Data_List.Get_foldM(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)}, v_3, b_prime_7, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__local_var_6_2)})
}))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
})
})
})
}

func Call_fromFoldable__614070391(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply2(dictFoldable_0.V2, pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}

func Call_groupAllBy__3934374991(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_List.Get_groupBy(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(p_0, x_1, y_2).IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}).IntVal) != (0))
})
}))
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(pkg_Data_List.Get_sortBy(), p_0)
_ = __local_var_2_1
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply(__local_var_2_1, x_3))
})
}

func Call_groupBy__2162447253(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
v2_2_0 := gopurs_runtime.Apply2(pkg_Data_List.Get_span(), gopurs_runtime.Apply(v_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)})
_ = v2_2_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v2_2_0, "init")))}})}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](gopurs_runtime.Apply2(pkg_Data_List.Get_groupBy(), v_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v2_2_0, "rest")))})))})})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](__t1)
}

func Call_groupBy__1039549870(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
v2_2_0 := gopurs_runtime.Apply2(pkg_Data_List.Get_span(), gopurs_runtime.Apply(v_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)})
_ = v2_2_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v2_2_0, "init")))}})}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](gopurs_runtime.Apply2(pkg_Data_List.Get_groupBy(), v_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v2_2_0, "rest")))})))})})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](__t1)
}

func Call_index__304299960(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], v1_1_loop int64) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var v1_1 int64 = v1_1_loop
_ = v1_1
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t0 gopurs_runtime.Value
{
if (v1_1) == (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(pkg_Data_List.Get_index(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}, gopurs_runtime.Int(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(v1_1), gopurs_runtime.Int(1)).IntVal))))}
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t1)
}

func Call_init__2496605985(lst_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] {
var lst_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = lst_0_loop
_ = lst_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v_1, "init")))}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_List.Get_unsnoc(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(lst_0)})))}))
}

func Call_insertAt__2634211748(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t1 gopurs_runtime.Value
{
if (v_0) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v1_1, v2_2})}})}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__local_var_3_0 := (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0
_ = __local_var_3_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v3_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_3_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v3_4)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(pkg_Data_List.Get_insertAt(), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal), v1_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1)})))})))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t1)
}

func Call_intersectBy__588351261(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], v2_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(pkg_Data_List.Get_filter(), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupDisj1_4_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), v_4, v1_5)
})
}))
_ = semigroupDisj1_4_0
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldMap"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_4_0
}), gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "ff")), gopurs_runtime.Apply(v_0, x_3), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}).IntVal) != (0))
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)})))}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1)
}

func Call_last__4043133652(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t5 gopurs_runtime.Value
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 1358893437 && __t_tag_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0})}
goto end_branch_5
} else {

}
}
{
var __t3 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 1358893437 && __t_tag_0.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}.UnsafePtr).V1)}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1358893437 && __t_tag_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}.UnsafePtr).V0})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_List.Get_last(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}.UnsafePtr).V1)})))}
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_3:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t3))}
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_6:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t6)
}

func Call_mapMaybe__3262563995(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_16 gopurs_runtime.Value
go__go_1_0_16 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_16:
for {
if false { continue go__go_1_0_16 }
var v_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t5 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
var go__go_4_1_17 gopurs_runtime.Value
go__go_4_1_17 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_17:
for {
if false { continue go__go_4_1_17 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_1_17
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}
}()
})
})
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_4_1_17, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)})))}
goto end_branch_5
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
v2_4_3 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0))
_ = v2_4_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.UnsafePtr == nil) {
v_2_loop = v_2
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_16
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.UnsafePtr != nil) {
v_2_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.UnsafePtr).V0, v_2})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_16
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t4))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t5))}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_16, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}

func Call_mapMaybe__1486753757(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_18 gopurs_runtime.Value
go__go_1_0_18 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_18:
for {
if false { continue go__go_1_0_18 }
var v_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t5 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
var go__go_4_1_19 gopurs_runtime.Value
go__go_4_1_19 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_19:
for {
if false { continue go__go_4_1_19 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_1_19
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}
}()
})
})
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_4_1_19, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)})))}
goto end_branch_5
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
v2_4_3 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0))
_ = v2_4_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.UnsafePtr == nil) {
v_2_loop = v_2
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_18
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.UnsafePtr != nil) {
v_2_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_4_3)}.UnsafePtr).V0, v_2})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_18
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t4))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t5))}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_18, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}

func Call_modifyAt__1886983628(n_0_loop int64, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(pkg_Data_List.Get_alterAt(), gopurs_runtime.Int(n_0), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, x_2)})}
}))
}

func Call_nubBy__2103943131(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_20 gopurs_runtime.Value
go__go_1_0_20 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_3_loop_val)
var v2_4_loop gopurs_runtime.Value = v2_4_loop_val
go__go_1_0_20:
for {
if false { continue go__go_1_0_20 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var v2_4 gopurs_runtime.Value = v2_4_loop
_ = v2_4
var __t3 gopurs_runtime.Value
{
if (v2_4.Type == 9 && v2_4.IntVal == 1358893437 && v2_4.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}
goto end_branch_3
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 1358893437 && v2_4.UnsafePtr != nil) {
v3_5_1 := gopurs_runtime.Apply3(pkg_Data_List_Internal.Get_insertAndLookupBy(), p_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_4.UnsafePtr).V0, v_2)
_ = v3_5_1
var __t2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
if (gopurs_runtime.RecordGet(v3_5_1, "found").IntVal) != (0) {
v_2_loop = gopurs_runtime.RecordGet(v3_5_1, "result")
v1_3_loop = v1_3
v2_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_4.UnsafePtr).V1)}
continue go__go_1_0_20
__t2 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
v_2_loop = gopurs_runtime.RecordGet(v3_5_1, "result")
v1_3_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_4.UnsafePtr).V0, v1_3})})
v2_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_4.UnsafePtr).V1)}
continue go__go_1_0_20
__t2 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t2)}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t3))}
}
}()
})
})
})
__local_var_2_4 := gopurs_runtime.Apply2(go__go_1_0_20, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_5_21 gopurs_runtime.Value
go__go_4_5_21 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_5_21:
for {
if false { continue go__go_4_5_21 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t6 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_6
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_5_21
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t6))}
}
}()
})
})
return gopurs_runtime.Apply2(go__go_4_5_21, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Apply(__local_var_2_4, x_3))
})
}

func Call_nubByEq__3956095361(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
__local_var_2_0 := (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0
_ = __local_var_2_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_2_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(pkg_Data_List.Get_nubByEq(), v_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(pkg_Data_List.Get_filter(), gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Bool((gopurs_runtime.Apply2(v_0, __local_var_2_0, y_3).IntVal) != (0))).IntVal) != (0))
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)})))}))})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1)
}

func Call_nubByEq__3655321914(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
__local_var_2_0 := (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0
_ = __local_var_2_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_2_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(pkg_Data_List.Get_nubByEq(), v_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(pkg_Data_List.Get_filter(), gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Bool((gopurs_runtime.Apply2(v_0, __local_var_2_0, y_3).IntVal) != (0))).IntVal) != (0))
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)})))}))})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1)
}

func Call_partition__1623965204(p_0_loop gopurs_runtime.Value, xs_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var xs_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(p_0, x_2).IntVal) != (0) {
__t0 = gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v_3, "no")))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_2, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v_3, "yes"))})})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_2, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v_3, "no"))})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v_3, "yes")))})
}
end_branch_0:
return __t0
})
}), gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)})
}

func Call_singleton__707062261(a_0_loop gopurs_runtime.Value) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})})
}

func Call_singleton__3932757557(a_0_loop gopurs_runtime.Value) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})})
}

func Call_snoc__4290067657(xs_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var xs_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = xs_0_loop
_ = xs_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, x_1, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_0)}))
}

func Call_sortBy__2103943131(cmp_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
var merge_1_0_24 gopurs_runtime.Value
_ = merge_1_0_24
merge_1_0_24 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
var __t1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(cmp_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0).IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}).IntVal) != (0) {
__t1 = &pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(merge_1_0_24, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}))}
goto end_branch_1
} else {

}
}
{
__t1 = &pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(merge_1_0_24, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_3))}))}
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)}
goto end_branch_2
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_3))}
goto end_branch_3
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t3))}
})
})
var mergePairs_2_4_25 gopurs_runtime.Value
_ = mergePairs_2_4_25
mergePairs_2_4_25 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
var __t_and_6 bool = false
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {

var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
__t_and_6 = (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 1358893437 && __t_tag_5.UnsafePtr != nil)
}
if __t_and_6 {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(merge_1_0_24, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}.UnsafePtr).V0))})))}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(mergePairs_2_4_25, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}.UnsafePtr).V1)})))})})}
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](v_3))}
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t7))}
})
var mergeAll_3_8_26 gopurs_runtime.Value
mergeAll_3_8_26 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](v_4_loop_val)
mergeAll_3_8_26:
for {
if false { continue mergeAll_3_8_26 }
var v_4 *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = v_4_loop
_ = v_4
var __t11 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
var __t_and_10 bool = false
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}.UnsafePtr != nil) {

var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}.UnsafePtr).V1)}
__t_and_10 = (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 1358893437 && __t_tag_9.UnsafePtr == nil)
}
if __t_and_10 {
__t11 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}.UnsafePtr).V0)
goto end_branch_11
} else {

}
}
{
v_4_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(mergePairs_2_4_25, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}))
continue mergeAll_3_8_26
__t11 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t11)}
}
}()
})
var sequences_4_12_27 gopurs_runtime.Value
_ = sequences_4_12_27
var descending_4_13_28 gopurs_runtime.Value
_ = descending_4_13_28
var ascending_4_14_29 gopurs_runtime.Value
_ = ascending_4_14_29
sequences_4_12_27 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t19 gopurs_runtime.Value
{
var __t_and_16 bool = false
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {

var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}
__t_and_16 = (__t_tag_15.Type == 9 && __t_tag_15.IntVal == 1358893437 && __t_tag_15.UnsafePtr != nil)
}
if __t_and_16 {
var __t18 *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(cmp_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}.UnsafePtr).V0).IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}).IntVal) != (0) {
__t18 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(descending_4_13_28, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}.UnsafePtr).V1)}))
goto end_branch_18
} else {

}
}
{
__local_var_6_17 := (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V0
_ = __local_var_6_17
__t18 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(ascending_4_14_29, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}.UnsafePtr).V0, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_6_17, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_7)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}.UnsafePtr).V1)}))
}
end_branch_18:
__t19 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t18)}
goto end_branch_19
} else {

}
}
{
__t19 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5))}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})}
}
end_branch_19:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t19))}
})
descending_4_13_28 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t20 gopurs_runtime.Value
{
if ((v2_7.Type == 9 && v2_7.IntVal == 1358893437 && v2_7.UnsafePtr != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(cmp_0, v_5, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_7.UnsafePtr).V0).IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}).IntVal) != (0)) {
__t20 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(descending_4_13_28, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_7.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v_5, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_6)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_7.UnsafePtr).V1)})))}
goto end_branch_20
} else {

}
}
{
__t20 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v_5, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_6)})}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(sequences_4_12_27, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_7))})))})})}
}
end_branch_20:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t20))}
})
})
})
ascending_4_14_29 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t21 gopurs_runtime.Value
{
if ((v2_7.Type == 9 && v2_7.IntVal == 1358893437 && v2_7.UnsafePtr != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(cmp_0, v_5, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_7.UnsafePtr).V0).IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}).IntVal) != (0)), gopurs_runtime.Bool(false)).IntVal) != (0)) {
__t21 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(ascending_4_14_29, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_7.UnsafePtr).V0, gopurs_runtime.Func(func(ys_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v_5, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](ys_8)})})))}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_7.UnsafePtr).V1)})))}
goto end_branch_21
} else {

}
}
{
__t21 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v_5, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})})))}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(sequences_4_12_27, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_7))})))})})}
}
end_branch_21:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t21))}
})
})
})
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(mergeAll_3_8_26, gopurs_runtime.Apply(sequences_4_12_27, x_5))
})
}

func Call_sortBy__1502591776(cmp_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
var merge_1_0_30 gopurs_runtime.Value
_ = merge_1_0_30
merge_1_0_30 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
var __t1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(cmp_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0).IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}).IntVal) != (0) {
__t1 = &pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(merge_1_0_30, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}))}
goto end_branch_1
} else {

}
}
{
__t1 = &pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(merge_1_0_30, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_3))}))}
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)}
goto end_branch_2
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_3))}
goto end_branch_3
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t3))}
})
})
var mergePairs_2_4_31 gopurs_runtime.Value
_ = mergePairs_2_4_31
mergePairs_2_4_31 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
var __t_and_6 bool = false
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {

var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
__t_and_6 = (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 1358893437 && __t_tag_5.UnsafePtr != nil)
}
if __t_and_6 {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(merge_1_0_30, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}.UnsafePtr).V0))})))}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(mergePairs_2_4_31, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}.UnsafePtr).V1)})))})})}
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](v_3))}
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t7))}
})
var mergeAll_3_8_32 gopurs_runtime.Value
mergeAll_3_8_32 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](v_4_loop_val)
mergeAll_3_8_32:
for {
if false { continue mergeAll_3_8_32 }
var v_4 *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = v_4_loop
_ = v_4
var __t11 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
var __t_and_10 bool = false
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}.UnsafePtr != nil) {

var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}.UnsafePtr).V1)}
__t_and_10 = (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 1358893437 && __t_tag_9.UnsafePtr == nil)
}
if __t_and_10 {
__t11 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}.UnsafePtr).V0)
goto end_branch_11
} else {

}
}
{
v_4_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(mergePairs_2_4_31, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}))
continue mergeAll_3_8_32
__t11 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t11)}
}
}()
})
var sequences_4_12_33 gopurs_runtime.Value
_ = sequences_4_12_33
var descending_4_13_34 gopurs_runtime.Value
_ = descending_4_13_34
var ascending_4_14_35 gopurs_runtime.Value
_ = ascending_4_14_35
sequences_4_12_33 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t19 gopurs_runtime.Value
{
var __t_and_16 bool = false
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {

var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}
__t_and_16 = (__t_tag_15.Type == 9 && __t_tag_15.IntVal == 1358893437 && __t_tag_15.UnsafePtr != nil)
}
if __t_and_16 {
var __t18 *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(cmp_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}.UnsafePtr).V0).IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}).IntVal) != (0) {
__t18 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(descending_4_13_34, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}.UnsafePtr).V1)}))
goto end_branch_18
} else {

}
}
{
__local_var_6_17 := (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V0
_ = __local_var_6_17
__t18 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(ascending_4_14_35, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}.UnsafePtr).V0, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_6_17, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_7)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}.UnsafePtr).V1)}))
}
end_branch_18:
__t19 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t18)}
goto end_branch_19
} else {

}
}
{
__t19 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5))}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})}
}
end_branch_19:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t19))}
})
descending_4_13_34 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t20 gopurs_runtime.Value
{
if ((v2_7.Type == 9 && v2_7.IntVal == 1358893437 && v2_7.UnsafePtr != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(cmp_0, v_5, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_7.UnsafePtr).V0).IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}).IntVal) != (0)) {
__t20 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(descending_4_13_34, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_7.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v_5, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_6)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_7.UnsafePtr).V1)})))}
goto end_branch_20
} else {

}
}
{
__t20 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v_5, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_6)})}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(sequences_4_12_33, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_7))})))})})}
}
end_branch_20:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t20))}
})
})
})
ascending_4_14_35 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t21 gopurs_runtime.Value
{
if ((v2_7.Type == 9 && v2_7.IntVal == 1358893437 && v2_7.UnsafePtr != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(cmp_0, v_5, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_7.UnsafePtr).V0).IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}).IntVal) != (0)), gopurs_runtime.Bool(false)).IntVal) != (0)) {
__t21 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(ascending_4_14_35, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_7.UnsafePtr).V0, gopurs_runtime.Func(func(ys_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v_5, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](ys_8)})})))}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_7.UnsafePtr).V1)})))}
goto end_branch_21
} else {

}
}
{
__t21 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v_5, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})})))}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(sequences_4_12_33, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_7))})))})})}
}
end_branch_21:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t21))}
})
})
})
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(mergeAll_3_8_32, gopurs_runtime.Apply(sequences_4_12_33, x_5))
})
}

func Call_span__1918198736(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t1 gopurs_runtime.Value
{
if ((gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil)) && ((gopurs_runtime.Apply(v_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0).IntVal) != (0)) {
v2_2_0 := gopurs_runtime.Apply2(pkg_Data_List.Get_span(), v_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)})
_ = v2_2_0
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v2_2_0, "init"))})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v2_2_0, "rest")))})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)})
}
end_branch_1:
return __t1
}

func Call_span__799093643(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t1 gopurs_runtime.Value
{
if ((gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil)) && ((gopurs_runtime.Apply(v_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0).IntVal) != (0)) {
v2_2_0 := gopurs_runtime.Apply2(pkg_Data_List.Get_span(), v_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)})
_ = v2_2_0
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v2_2_0, "init"))})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v2_2_0, "rest")))})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)})
}
end_branch_1:
return __t1
}

func Call_takeWhile__2352021032(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_39 gopurs_runtime.Value
go__go_1_0_39 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_39:
for {
if false { continue go__go_1_0_39 }
var v_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t3 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
if ((v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil)) && ((gopurs_runtime.Apply(p_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0).IntVal) != (0)) {
v_2_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, v_2})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__go_1_0_39
__t3 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
var go__go_4_1_40 gopurs_runtime.Value
go__go_4_1_40 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_40:
for {
if false { continue go__go_4_1_40 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_1_40
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}
}()
})
})
__t3 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_4_1_40, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}))
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_39, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}

func Call_uncons__3009258782(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("head", "tail", (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)})})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0)
}

func Call_unionBy__588351261(eq_0_loop gopurs_runtime.Value, xs_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], ys_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = xs_1_loop
_ = xs_1
var ys_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = ys_2_loop
_ = ys_2
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Data_List.Get_deleteBy(), eq_0, a_4, b_3)
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(pkg_Data_List.Get_nubByEq(), eq_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(ys_2)})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)})))}))
}

func Call_unsnoc__2942606998(lst_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var lst_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = lst_0_loop
_ = lst_0
var go__go_1_0_41 gopurs_runtime.Value
go__go_1_0_41 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_3_loop_val)
go__go_1_0_41:
for {
if false { continue go__go_1_0_41 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t3 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1358893437 && __t_tag_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("last", "revInit", (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)})})}
goto end_branch_2
} else {

}
}
{
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V1)}
v1_3_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr).V0, v1_3})})
continue go__go_1_0_41
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t3))}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(h_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_4_42 gopurs_runtime.Value
go__go_3_4_42 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_4_42:
for {
if false { continue go__go_3_4_42 }
var v_4 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t5 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}
goto end_branch_5
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
continue go__go_3_4_42
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t5))}
}
}()
})
})
return gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_3_4_42, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(h_2, "revInit")))})))}, gopurs_runtime.RecordGet(h_2, "last"))
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_1_0_41, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(lst_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})))}))
}

func Call_updateAt__2634211748(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (v_0) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, v1_1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1})}})}
goto end_branch_1
} else {

}
}
{
__local_var_3_0 := (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0
_ = __local_var_3_0
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v3_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, __local_var_3_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v3_4)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(pkg_Data_List.Get_updateAt(), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal), v1_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V1)})))})))}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](__t2)
}

func Call_zipWith__884793877(f_0_loop gopurs_runtime.Value, xs_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], ys_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = xs_1_loop
_ = xs_1
var ys_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = ys_2_loop
_ = ys_2
var go__go_3_0_43 gopurs_runtime.Value
go__go_3_0_43 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
var v2_6_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v2_6_loop_val)
go__go_3_0_43:
for {
if false { continue go__go_3_0_43 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var v2_6 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v2_6_loop
_ = v2_6
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_6)}
goto end_branch_1
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_6)}
goto end_branch_1
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil)) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
v2_6_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0), v2_6})})
continue go__go_3_0_43
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}
}()
})
})
})
var go__go_4_2_44 gopurs_runtime.Value
go__go_4_2_44 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_2_44:
for {
if false { continue go__go_4_2_44 }
var v_5 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t3 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_5)}
goto end_branch_3
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_2_44
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t3))}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_4_2_44, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(go__go_3_0_43, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(ys_2)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})))}))
}

func Call_fromMaybe__430429096(a_0_loop gopurs_runtime.Value, v2_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v2_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_1_loop
_ = v2_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr == nil) {
__t0 = a_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_1)}.UnsafePtr).V0
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

func Call_maybe__1246883617(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_alaF__4085337484(_dollar__unused_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, _dollar__unused_2_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value], _dollar__unused_3_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value], v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var _dollar__unused_2 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_2_loop
_ = _dollar__unused_2
var _dollar__unused_3 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_3_loop
_ = _dollar__unused_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
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

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_sequence1__4078930642(dict_0_loop *pkg_Data_Semigroup_Traversable.Constructor_Traversable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup_Traversable.Constructor_Traversable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_traverse1__157785005(dict_0_loop *pkg_Data_Semigroup_Traversable.Constructor_Traversable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup_Traversable.Constructor_Traversable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
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

func Call_unfoldr__1128708256(dict_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_unfoldr__3827943605(dict_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_unsafeCrashWith__69763299(msg_0_loop string) gopurs_runtime.Value {
var msg_0 string = msg_0_loop
_ = msg_0
return gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str(msg_0))
}))
}

func Call_unsafeCrashWith__2941848695(msg_0_loop string) gopurs_runtime.Value {
var msg_0 string = msg_0_loop
_ = msg_0
return gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str(msg_0))
}))
}

func Call_crashWith__1894115486(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Partial.Get__crashWith()
}



package Data_List_NonEmpty

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_List "gopurs/output/Data.List"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	pkg_Partial "gopurs/output/Partial"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	unsafe "unsafe"
)

var cache_identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		cache_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity(x_0_box)
})
	})
	return cache_identity
}

var cache_zipWith gopurs_runtime.Value
var once_zipWith sync.Once
func Get_zipWith() gopurs_runtime.Value {
	once_zipWith.Do(func() {
		cache_zipWith = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWith(f_0_box, v_1_box, v1_2_box)
})
	})
	return cache_zipWith
}

var cache_zipWithA gopurs_runtime.Value
var once_zipWithA sync.Once
func Get_zipWithA() gopurs_runtime.Value {
	once_zipWithA.Do(func() {
		cache_zipWithA = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWithA((*Record_pure_gopurs_runtime_Value)(dictApplicative_0_box.UnsafePtr))
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
return Call_wrappedOperation2(name_0_box.StrVal(), f_1_box, v_2_box, v1_3_box)
})
	})
	return cache_wrappedOperation2
}

var cache_wrappedOperation gopurs_runtime.Value
var once_wrappedOperation sync.Once
func Get_wrappedOperation() gopurs_runtime.Value {
	once_wrappedOperation.Do(func() {
		cache_wrappedOperation = gopurs_runtime.Func3(func(name_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_wrappedOperation(name_0_box.StrVal(), f_1_box, v_2_box)
})
	})
	return cache_wrappedOperation
}

var cache_updateAt gopurs_runtime.Value
var once_updateAt sync.Once
func Get_updateAt() gopurs_runtime.Value {
	once_updateAt.Do(func() {
		cache_updateAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_updateAt(i_0_box.IntVal, a_1_box, v_2_box)
})
	})
	return cache_updateAt
}

var cache_unzip gopurs_runtime.Value
var once_unzip sync.Once
func Get_unzip() gopurs_runtime.Value {
	once_unzip.Do(func() {
		cache_unzip = gopurs_runtime.Func(func(ts_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unzip(ts_0_box)
})
	})
	return cache_unzip
}

var cache_unsnoc gopurs_runtime.Value
var once_unsnoc sync.Once
func Get_unsnoc() gopurs_runtime.Value {
	once_unsnoc.Do(func() {
		cache_unsnoc = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsnoc(v_0_box)
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
return Call_union((*Record_eq_gopurs_runtime_Value)(dictEq_0_box.UnsafePtr))
})
	})
	return cache_union
}

var cache_uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		cache_uncons = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_uncons(v_0_box)
})
	})
	return cache_uncons
}

var cache_toList gopurs_runtime.Value
var once_toList sync.Once
func Get_toList() gopurs_runtime.Value {
	once_toList.Do(func() {
		cache_toList = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toList(v_0_box)
})
	})
	return cache_toList
}

var cache_toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		cache_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toUnfoldable((*Record_unfoldr_gopurs_runtime_Value)(dictUnfoldable_0_box.UnsafePtr))
})
	})
	return cache_toUnfoldable
}

var cache_tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		cache_tail = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tail(v_0_box)
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
return Call_sort((*Record_compare_gopurs_runtime_Value)(dictOrd_0_box.UnsafePtr))
})
	})
	return cache_sort
}

var cache_snoc gopurs_runtime.Value
var once_snoc sync.Once
func Get_snoc() gopurs_runtime.Value {
	once_snoc.Do(func() {
		cache_snoc = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_snoc(v_0_box, y_1_box)
})
	})
	return cache_snoc
}

var cache_singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		cache_singleton = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_singleton(x_0_box)
})
	})
	return cache_singleton
}

var cache_snoc_prime gopurs_runtime.Value
var once_snoc_prime sync.Once
func Get_snoc_prime() gopurs_runtime.Value {
	once_snoc_prime.Do(func() {
		cache_snoc_prime = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_snoc_prime(v_0_box, v1_1_box)
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
return Call_nubEq((*Record_eq_gopurs_runtime_Value)(dictEq_0_box.UnsafePtr))
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
return Call_nub((*Record_compare_gopurs_runtime_Value)(dictOrd_0_box.UnsafePtr))
})
	})
	return cache_nub
}

var cache_modifyAt gopurs_runtime.Value
var once_modifyAt sync.Once
func Get_modifyAt() gopurs_runtime.Value {
	once_modifyAt.Do(func() {
		cache_modifyAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modifyAt(i_0_box.IntVal, f_1_box, v_2_box)
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
return Call_partition(x_0_box, v_1_box)
})
	})
	return cache_partition
}

var cache_span gopurs_runtime.Value
var once_span sync.Once
func Get_span() gopurs_runtime.Value {
	once_span.Do(func() {
		cache_span = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_span(x_0_box, v_1_box)
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
return gopurs_runtime.Int(Call_length(v_0_box))
})
	})
	return cache_length
}

var cache_last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		cache_last = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_last(v_0_box)
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
return Call_intersect((*Record_eq_gopurs_runtime_Value)(dictEq_0_box.UnsafePtr))
})
	})
	return cache_intersect
}

var cache_insertAt gopurs_runtime.Value
var once_insertAt sync.Once
func Get_insertAt() gopurs_runtime.Value {
	once_insertAt.Do(func() {
		cache_insertAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insertAt(i_0_box.IntVal, a_1_box, v_2_box)
})
	})
	return cache_insertAt
}

var cache_init_ gopurs_runtime.Value
var once_init_ sync.Once
func Get_init_() gopurs_runtime.Value {
	once_init_.Do(func() {
		cache_init_ = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_init_(v_0_box)
})
	})
	return cache_init_
}

var cache_index gopurs_runtime.Value
var once_index sync.Once
func Get_index() gopurs_runtime.Value {
	once_index.Do(func() {
		cache_index = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, i_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_index(v_0_box, i_1_box.IntVal)
})
	})
	return cache_index
}

var cache_head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		cache_head = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_head(v_0_box)
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
return Call_groupAll((*Record_compare_gopurs_runtime_Value)(dictOrd_0_box.UnsafePtr))
})
	})
	return cache_groupAll
}

var cache_group gopurs_runtime.Value
var once_group sync.Once
func Get_group() gopurs_runtime.Value {
	once_group.Do(func() {
		cache_group = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_group((*Record_eq_gopurs_runtime_Value)(dictEq_0_box.UnsafePtr))
})
	})
	return cache_group
}

var cache_fromList gopurs_runtime.Value
var once_fromList sync.Once
func Get_fromList() gopurs_runtime.Value {
	once_fromList.Do(func() {
		cache_fromList = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromList(v_0_box)
})
	})
	return cache_fromList
}

var cache_fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		cache_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldable((*Record_foldMap_gopurs_runtime_Value_foldl_gopurs_runtime_Value_foldr_gopurs_runtime_Value)(dictFoldable_0_box.UnsafePtr))
})
	})
	return cache_fromFoldable
}

var cache_foldM gopurs_runtime.Value
var once_foldM sync.Once
func Get_foldM() gopurs_runtime.Value {
	once_foldM.Do(func() {
		cache_foldM = gopurs_runtime.Func4(func(dictMonad_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldM((*Record_)(dictMonad_0_box.UnsafePtr), f_1_box, b_2_box, v_3_box)
})
	})
	return cache_foldM
}

var cache_findLastIndex gopurs_runtime.Value
var once_findLastIndex sync.Once
func Get_findLastIndex() gopurs_runtime.Value {
	once_findLastIndex.Do(func() {
		cache_findLastIndex = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_findLastIndex(f_0_box, v_1_box)
})
	})
	return cache_findLastIndex
}

var cache_findIndex gopurs_runtime.Value
var once_findIndex sync.Once
func Get_findIndex() gopurs_runtime.Value {
	once_findIndex.Do(func() {
		cache_findIndex = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_findIndex(f_0_box, v_1_box)
})
	})
	return cache_findIndex
}

var cache_filterM gopurs_runtime.Value
var once_filterM sync.Once
func Get_filterM() gopurs_runtime.Value {
	once_filterM.Do(func() {
		cache_filterM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filterM((*Record_)(dictMonad_0_box.UnsafePtr))
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
return Call_elemLastIndex((*Record_eq_gopurs_runtime_Value)(dictEq_0_box.UnsafePtr), x_1_box)
})
	})
	return cache_elemLastIndex
}

var cache_elemIndex gopurs_runtime.Value
var once_elemIndex sync.Once
func Get_elemIndex() gopurs_runtime.Value {
	once_elemIndex.Do(func() {
		cache_elemIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_elemIndex((*Record_eq_gopurs_runtime_Value)(dictEq_0_box.UnsafePtr), x_1_box)
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
return Call_drop(x_0_box.IntVal, v_1_box)
})
	})
	return cache_drop
}

var cache_cons_prime gopurs_runtime.Value
var once_cons_prime sync.Once
func Get_cons_prime() gopurs_runtime.Value {
	once_cons_prime.Do(func() {
		cache_cons_prime = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cons_prime(x_0_box, xs_1_box)
})
	})
	return cache_cons_prime
}

var cache_cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cache_cons = gopurs_runtime.Func2(func(y_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cons(y_0_box, v_1_box)
})
	})
	return cache_cons
}

var cache_concatMap gopurs_runtime.Value
var once_concatMap sync.Once
func Get_concatMap() gopurs_runtime.Value {
	once_concatMap.Do(func() {
		cache_concatMap = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_concatMap(b_0_box, a_1_box)
})
	})
	return cache_concatMap
}

var cache_concat gopurs_runtime.Value
var once_concat sync.Once
func Get_concat() gopurs_runtime.Value {
	once_concat.Do(func() {
		cache_concat = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_concat(v_0_box)
})
	})
	return cache_concat
}

var cache_catMaybes gopurs_runtime.Value
var once_catMaybes sync.Once
func Get_catMaybes() gopurs_runtime.Value {
	once_catMaybes.Do(func() {
		cache_catMaybes = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_catMaybes(v_0_box)
})
	})
	return cache_catMaybes
}

var cache_appendFoldable gopurs_runtime.Value
var once_appendFoldable sync.Once
func Get_appendFoldable() gopurs_runtime.Value {
	once_appendFoldable.Do(func() {
		cache_appendFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_appendFoldable((*Record_foldMap_gopurs_runtime_Value_foldl_gopurs_runtime_Value_foldr_gopurs_runtime_Value)(dictFoldable_0_box.UnsafePtr))
})
	})
	return cache_appendFoldable
}

type Record_alt_gopurs_runtime_Value struct {
	alt gopurs_runtime.Value
}

type Record_ struct {
	
}

type Record_pure_gopurs_runtime_Value struct {
	pure gopurs_runtime.Value
}

type Record_apply_gopurs_runtime_Value struct {
	apply gopurs_runtime.Value
}

type Record_bipure_gopurs_runtime_Value struct {
	bipure gopurs_runtime.Value
}

type Record_biapply_gopurs_runtime_Value struct {
	biapply gopurs_runtime.Value
}

type Record_bind_gopurs_runtime_Value struct {
	bind gopurs_runtime.Value
}

type Record_discard_gopurs_runtime_Value struct {
	discard gopurs_runtime.Value
}

type Record_identity_gopurs_runtime_Value struct {
	identity gopurs_runtime.Value
}

type Record_ask_gopurs_runtime_Value struct {
	ask gopurs_runtime.Value
}

type Record_local_gopurs_runtime_Value struct {
	local gopurs_runtime.Value
}

type Record_peek_gopurs_runtime_Value_pos_gopurs_runtime_Value struct {
	peek gopurs_runtime.Value
	pos gopurs_runtime.Value
}

type Record_track_gopurs_runtime_Value struct {
	track gopurs_runtime.Value
}

type Record_extract_gopurs_runtime_Value struct {
	extract gopurs_runtime.Value
}

type Record_extend_gopurs_runtime_Value struct {
	extend gopurs_runtime.Value
}

type Record_defer__gopurs_runtime_Value struct {
	defer_ gopurs_runtime.Value
}

type Record_callCC_gopurs_runtime_Value struct {
	callCC gopurs_runtime.Value
}

type Record_catchError_gopurs_runtime_Value struct {
	catchError gopurs_runtime.Value
}

type Record_throwError_gopurs_runtime_Value struct {
	throwError gopurs_runtime.Value
}

type Record_chooseBool_gopurs_runtime_Value_chooseFloat_gopurs_runtime_Value_chooseInt_gopurs_runtime_Value_resize_gopurs_runtime_Value_sized_gopurs_runtime_Value struct {
	chooseBool gopurs_runtime.Value
	chooseFloat gopurs_runtime.Value
	chooseInt gopurs_runtime.Value
	resize gopurs_runtime.Value
	sized gopurs_runtime.Value
}

type Record_foldMap1_gopurs_runtime_Value_foldl1_gopurs_runtime_Value_foldr1_gopurs_runtime_Value struct {
	foldMap1 gopurs_runtime.Value
	foldl1 gopurs_runtime.Value
	foldr1 gopurs_runtime.Value
}

type Record_append__gopurs_runtime_Value struct {
	append_ gopurs_runtime.Value
}

type Record_tailRecM_gopurs_runtime_Value struct {
	tailRecM gopurs_runtime.Value
}

type Record_unfoldr_gopurs_runtime_Value struct {
	unfoldr gopurs_runtime.Value
}

type Record_map__gopurs_runtime_Value struct {
	map_ gopurs_runtime.Value
}

type Record_state_gopurs_runtime_Value struct {
	state gopurs_runtime.Value
}

type Record_lift_gopurs_runtime_Value struct {
	lift gopurs_runtime.Value
}

type Record_listen_gopurs_runtime_Value_pass_gopurs_runtime_Value struct {
	listen gopurs_runtime.Value
	pass gopurs_runtime.Value
}

type Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value struct {
	parallel gopurs_runtime.Value
	sequential gopurs_runtime.Value
}

type Record_foldMap_gopurs_runtime_Value_foldl_gopurs_runtime_Value_foldr_gopurs_runtime_Value struct {
	foldMap gopurs_runtime.Value
	foldl gopurs_runtime.Value
	foldr gopurs_runtime.Value
}

type Record_mempty_gopurs_runtime_Value struct {
	mempty gopurs_runtime.Value
}

type Record_sequence_gopurs_runtime_Value_traverse_gopurs_runtime_Value struct {
	sequence gopurs_runtime.Value
	traverse gopurs_runtime.Value
}

type Record_empty_gopurs_runtime_Value struct {
	empty gopurs_runtime.Value
}

type Record_compose_gopurs_runtime_Value struct {
	compose gopurs_runtime.Value
}

type Record_eq_gopurs_runtime_Value struct {
	eq gopurs_runtime.Value
}

type Record_compare_gopurs_runtime_Value struct {
	compare gopurs_runtime.Value
}

type Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value struct {
	bifoldMap gopurs_runtime.Value
	bifoldl gopurs_runtime.Value
	bifoldr gopurs_runtime.Value
}

type Record_bimap_gopurs_runtime_Value struct {
	bimap gopurs_runtime.Value
}

type Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value struct {
	bisequence gopurs_runtime.Value
	bitraverse gopurs_runtime.Value
}

type Record_genericBottom_prime_gopurs_runtime_Value struct {
	genericBottom_prime gopurs_runtime.Value
}

type Record_genericTop_prime_gopurs_runtime_Value struct {
	genericTop_prime gopurs_runtime.Value
}

type Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value struct {
	bottom gopurs_runtime.Value
	top gopurs_runtime.Value
}

type Record_lose_gopurs_runtime_Value struct {
	lose gopurs_runtime.Value
}

type Record_choose_gopurs_runtime_Value struct {
	choose gopurs_runtime.Value
}

type Record_collect_gopurs_runtime_Value_distribute_gopurs_runtime_Value struct {
	collect gopurs_runtime.Value
	distribute gopurs_runtime.Value
}

type Record_divide_gopurs_runtime_Value struct {
	divide gopurs_runtime.Value
}

type Record_recip_gopurs_runtime_Value struct {
	recip gopurs_runtime.Value
}

type Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value struct {
	genericCardinality_prime gopurs_runtime.Value
	genericFromEnum_prime gopurs_runtime.Value
	genericToEnum_prime gopurs_runtime.Value
}

type Record_genericPred_prime_gopurs_runtime_Value_genericSucc_prime_gopurs_runtime_Value struct {
	genericPred_prime gopurs_runtime.Value
	genericSucc_prime gopurs_runtime.Value
}

type Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value struct {
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
}

type Record_unfoldr1_gopurs_runtime_Value struct {
	unfoldr1 gopurs_runtime.Value
}

type Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value struct {
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
}

type Record_genericEq_prime_gopurs_runtime_Value struct {
	genericEq_prime gopurs_runtime.Value
}

type Record_eq1_gopurs_runtime_Value struct {
	eq1 gopurs_runtime.Value
}

type Record_degree_gopurs_runtime_Value_div_gopurs_runtime_Value_mod_gopurs_runtime_Value struct {
	degree gopurs_runtime.Value
	div gopurs_runtime.Value
	mod gopurs_runtime.Value
}

type Record_conj_gopurs_runtime_Value_disj_gopurs_runtime_Value_ff_gopurs_runtime_Value_implies_gopurs_runtime_Value_not_gopurs_runtime_Value_tt_gopurs_runtime_Value struct {
	conj gopurs_runtime.Value
	disj gopurs_runtime.Value
	ff gopurs_runtime.Value
	implies gopurs_runtime.Value
	not gopurs_runtime.Value
	tt gopurs_runtime.Value
}

type Record_add_gopurs_runtime_Value_mul_gopurs_runtime_Value_one_gopurs_runtime_Value_zero_gopurs_runtime_Value struct {
	add gopurs_runtime.Value
	mul gopurs_runtime.Value
	one gopurs_runtime.Value
	zero gopurs_runtime.Value
}

type Record_foldMapWithIndex_gopurs_runtime_Value_foldlWithIndex_gopurs_runtime_Value_foldrWithIndex_gopurs_runtime_Value struct {
	foldMapWithIndex gopurs_runtime.Value
	foldlWithIndex gopurs_runtime.Value
	foldrWithIndex gopurs_runtime.Value
}

type Record_cmap_gopurs_runtime_Value struct {
	cmap gopurs_runtime.Value
}

type Record_imap_gopurs_runtime_Value struct {
	imap gopurs_runtime.Value
}

type Record_mapWithIndex_gopurs_runtime_Value struct {
	mapWithIndex gopurs_runtime.Value
}

type Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value struct {
	from gopurs_runtime.Value
	to gopurs_runtime.Value
}

type Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value struct {
	genericConj_prime gopurs_runtime.Value
	genericDisj_prime gopurs_runtime.Value
	genericFF_prime gopurs_runtime.Value
	genericImplies_prime gopurs_runtime.Value
	genericNot_prime gopurs_runtime.Value
	genericTT_prime gopurs_runtime.Value
}

type Record_conj_gopurs_runtime_Value_disj_gopurs_runtime_Value_ff_bool_implies_gopurs_runtime_Value_not_gopurs_runtime_Value_tt_bool struct {
	conj gopurs_runtime.Value
	disj gopurs_runtime.Value
	ff bool
	implies gopurs_runtime.Value
	not gopurs_runtime.Value
	tt bool
}

type Record_genericMempty_prime_gopurs_runtime_Value struct {
	genericMempty_prime gopurs_runtime.Value
}

type Record_genericCompare_prime_gopurs_runtime_Value struct {
	genericCompare_prime gopurs_runtime.Value
}

type Record_sub_gopurs_runtime_Value struct {
	sub gopurs_runtime.Value
}

type Record_compare1_gopurs_runtime_Value struct {
	compare1 gopurs_runtime.Value
}

type Record_left_gopurs_runtime_Value_right_gopurs_runtime_Value struct {
	left gopurs_runtime.Value
	right gopurs_runtime.Value
}

type Record_first_gopurs_runtime_Value_second_gopurs_runtime_Value struct {
	first gopurs_runtime.Value
	second gopurs_runtime.Value
}

type Record_dimap_gopurs_runtime_Value struct {
	dimap gopurs_runtime.Value
}

type Record_genericSub_prime_gopurs_runtime_Value struct {
	genericSub_prime gopurs_runtime.Value
}

type Record_genericAppend_prime_gopurs_runtime_Value struct {
	genericAppend_prime gopurs_runtime.Value
}

type Record_sequence1_gopurs_runtime_Value_traverse1_gopurs_runtime_Value struct {
	sequence1 gopurs_runtime.Value
	traverse1 gopurs_runtime.Value
}

type Record_genericAdd_prime_gopurs_runtime_Value_genericMul_prime_gopurs_runtime_Value_genericOne_prime_gopurs_runtime_Value_genericZero_prime_gopurs_runtime_Value struct {
	genericAdd_prime gopurs_runtime.Value
	genericMul_prime gopurs_runtime.Value
	genericOne_prime gopurs_runtime.Value
	genericZero_prime gopurs_runtime.Value
}

type Record_genericShow_prime_gopurs_runtime_Value struct {
	genericShow_prime gopurs_runtime.Value
}

type Record_genericShowArgs_gopurs_runtime_Value struct {
	genericShowArgs gopurs_runtime.Value
}

type Record_show_gopurs_runtime_Value struct {
	show gopurs_runtime.Value
}

type Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value struct {
	fromDuration gopurs_runtime.Value
	toDuration gopurs_runtime.Value
}

type Record_traverseWithIndex_gopurs_runtime_Value struct {
	traverseWithIndex gopurs_runtime.Value
}

type Record_liftEffect_gopurs_runtime_Value struct {
	liftEffect gopurs_runtime.Value
}

type Record_mappend__gopurs_runtime_Value_mempty__gopurs_runtime_Value struct {
	mappend_ gopurs_runtime.Value
	mempty_ gopurs_runtime.Value
}

type Record_proof_gopurs_runtime_Value struct {
	proof gopurs_runtime.Value
}

type Record_lower_gopurs_runtime_Value struct {
	lower gopurs_runtime.Value
}

type Record_liftST_gopurs_runtime_Value struct {
	liftST gopurs_runtime.Value
}

type Record_tell_gopurs_runtime_Value struct {
	tell gopurs_runtime.Value
}

type Record_reflectSymbol_gopurs_runtime_Value struct {
	reflectSymbol gopurs_runtime.Value
}

type Record_bottomRecord_gopurs_runtime_Value_topRecord_gopurs_runtime_Value struct {
	bottomRecord gopurs_runtime.Value
	topRecord gopurs_runtime.Value
}

type Record_conquer_gopurs_runtime_Value struct {
	conquer gopurs_runtime.Value
}

type Record_inj_gopurs_runtime_Value_prj_gopurs_runtime_Value struct {
	inj gopurs_runtime.Value
	prj gopurs_runtime.Value
}

type Record_eqRecord_gopurs_runtime_Value struct {
	eqRecord gopurs_runtime.Value
}

type Record_conjRecord_gopurs_runtime_Value_disjRecord_gopurs_runtime_Value_ffRecord_gopurs_runtime_Value_impliesRecord_gopurs_runtime_Value_notRecord_gopurs_runtime_Value_ttRecord_gopurs_runtime_Value struct {
	conjRecord gopurs_runtime.Value
	disjRecord gopurs_runtime.Value
	ffRecord gopurs_runtime.Value
	impliesRecord gopurs_runtime.Value
	notRecord gopurs_runtime.Value
	ttRecord gopurs_runtime.Value
}

type Record_memptyRecord_gopurs_runtime_Value struct {
	memptyRecord gopurs_runtime.Value
}

type Record_compareRecord_gopurs_runtime_Value struct {
	compareRecord gopurs_runtime.Value
}

type Record_closed_gopurs_runtime_Value struct {
	closed gopurs_runtime.Value
}

type Record_unleft_gopurs_runtime_Value_unright_gopurs_runtime_Value struct {
	unleft gopurs_runtime.Value
	unright gopurs_runtime.Value
}

type Record_unfirst_gopurs_runtime_Value_unsecond_gopurs_runtime_Value struct {
	unfirst gopurs_runtime.Value
	unsecond gopurs_runtime.Value
}

type Record_reflectType_gopurs_runtime_Value struct {
	reflectType gopurs_runtime.Value
}

type Record_subRecord_gopurs_runtime_Value struct {
	subRecord gopurs_runtime.Value
}

type Record_appendRecord_gopurs_runtime_Value struct {
	appendRecord gopurs_runtime.Value
}

type Record_addRecord_gopurs_runtime_Value_mulRecord_gopurs_runtime_Value_oneRecord_gopurs_runtime_Value_zeroRecord_gopurs_runtime_Value struct {
	addRecord gopurs_runtime.Value
	mulRecord gopurs_runtime.Value
	oneRecord gopurs_runtime.Value
	zeroRecord gopurs_runtime.Value
}

type Record_showRecordFields_gopurs_runtime_Value struct {
	showRecordFields gopurs_runtime.Value
}

type Record_nes_gopurs_runtime_Value struct {
	nes gopurs_runtime.Value
}

type Record_liftAff_gopurs_runtime_Value struct {
	liftAff gopurs_runtime.Value
}

func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_zipWith(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var go__3_0 gopurs_runtime.Value
go__3_0 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
var v2_6_loop gopurs_runtime.Value = v2_6_loop_val
go__3_0:
for {
if false { continue go__3_0 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var v2_6 gopurs_runtime.Value = v2_6_loop
_ = v2_6
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 786377863) {
__t1 = v2_6
goto end_branch_1
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 786377863) {
__t1 = v2_6
goto end_branch_1
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 1358893437)) && ((v1_5.Type == 9 && v1_5.IntVal == 1358893437)) {
v_4_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V1
v1_5_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_5.UnsafePtr).V1
v2_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Apply2(f_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_5.UnsafePtr).V0), v2_6})}
continue go__3_0
__t1 = gopurs_runtime.Value{}
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
}()
})
})
})
var go__4_2 gopurs_runtime.Value
go__4_2 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__4_2:
for {
if false { continue go__4_2 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t3 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 786377863) {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})}
v1_6_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V1
continue go__4_2
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty{gopurs_runtime.Apply2(f_0, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_1.UnsafePtr).V0, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v1_2.UnsafePtr).V0), gopurs_runtime.Apply2(go__4_2, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, gopurs_runtime.Apply3(go__3_0, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_1.UnsafePtr).V1, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v1_2.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}))})}
}

func Call_zipWithA(dictApplicative_0_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
var dictApplicative_0 *Record_pure_gopurs_runtime_Value = dictApplicative_0_loop
_ = dictApplicative_0
sequence11_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_traversable1NonEmptyList(), "sequence1"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, "Apply0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = sequence11_1_0
return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value, ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence11_1_0, Call_zipWith(f_2, xs_3, ys_4))
})
}

func Call_wrappedOperation2(name_0_loop string, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value, v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var name_0 string = name_0_loop
_ = name_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
v2_4_0 := gopurs_runtime.Apply2(f_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v1_3.UnsafePtr).V0, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v1_3.UnsafePtr).V1})})
_ = v2_4_0
var __t1 gopurs_runtime.Value
{
if (v2_4_0.Type == 9 && v2_4_0.IntVal == 1358893437) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_4_0.UnsafePtr).V0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_4_0.UnsafePtr).V1})}
goto end_branch_1
} else {

}
}
{
if (v2_4_0.Type == 9 && v2_4_0.IntVal == 786377863) {
__t1 = gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("Impossible: empty list in NonEmptyList "), gopurs_runtime.Str(name_0)))
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

func Call_wrappedOperation(name_0_loop string, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var name_0 string = name_0_loop
_ = name_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
v1_3_0 := gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1})})
_ = v1_3_0
var __t1 gopurs_runtime.Value
{
if (v1_3_0.Type == 9 && v1_3_0.IntVal == 1358893437) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3_0.UnsafePtr).V0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3_0.UnsafePtr).V1})}
goto end_branch_1
} else {

}
}
{
if (v1_3_0.Type == 9 && v1_3_0.IntVal == 786377863) {
__t1 = gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("Impossible: empty list in NonEmptyList "), gopurs_runtime.Str(name_0)))
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

func Call_updateAt(i_0_loop int64, a_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var i_0 int64 = i_0_loop
_ = i_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (i_0) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty{a_1, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1})}})}
goto end_branch_1
} else {

}
}
{
__local_var_3_0 := (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0
_ = __local_var_3_0
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty{__local_var_3_0, x_4})}
}), gopurs_runtime.Apply3(pkg_Data_List.Get_updateAt(), gopurs_runtime.Int((i_0) - (1)), a_1, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1))
}
end_branch_1:
return __t1
}

func Call_unzip(ts_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var ts_0 gopurs_runtime.Value = ts_0_loop
_ = ts_0
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_functorNonEmptyList(), "map"), pkg_Data_Tuple.Get_fst(), ts_0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_functorNonEmptyList(), "map"), pkg_Data_Tuple.Get_snd(), ts_0)})}
}

func Call_unsnoc(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
v1_1_0 := gopurs_runtime.Apply(pkg_Data_List.Get_unsnoc(), (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1)
_ = v1_1_0
var __t1 gopurs_runtime.Value
{
if (v1_1_0.Type == 9 && v1_1_0.IntVal == 3589588149) {
__t1 = gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
if (v1_1_0.Type == 9 && v1_1_0.IntVal == 930809136) {
__t1 = gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V0, gopurs_runtime.RecordGet((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v1_1_0.UnsafePtr).V0, "init")})}, gopurs_runtime.RecordGet((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v1_1_0.UnsafePtr).V0, "last"))
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

func Call_union(dictEq_0_loop *Record_eq_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEq_0 *Record_eq_gopurs_runtime_Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply2(Get_wrappedOperation2(), gopurs_runtime.Str("union"), gopurs_runtime.Apply(pkg_Data_List.Get_union(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictEq_0)}))
}

func Call_uncons(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.RecordDict2("head", "tail", (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V0, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1)
}

func Call_toList(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V0, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1})}
}

func Call_toUnfoldable(dictUnfoldable_0_loop *Record_unfoldr_gopurs_runtime_Value) gopurs_runtime.Value {
var dictUnfoldable_0 *Record_unfoldr_gopurs_runtime_Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
__local_var_1_0 := gopurs_runtime.Apply(dictUnfoldable_0.unfoldr, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (xs_1.Type == 9 && xs_1.IntVal == 786377863) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if (xs_1.Type == 9 && xs_1.IntVal == 1358893437) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.RecordDict2("head", "tail", (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(xs_1.UnsafePtr).V0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(xs_1.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(rec_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{gopurs_runtime.RecordGet(rec_2, "head"), gopurs_runtime.RecordGet(rec_2, "tail")})}
}), __t1)
}))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V0, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V1})})
})
}

func Call_tail(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1
}

func Call_sortBy(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply2(Get_wrappedOperation(), gopurs_runtime.Str("sortBy"), gopurs_runtime.Apply(pkg_Data_List.Get_sortBy(), x_0))
}

func Call_sort(dictOrd_0_loop *Record_compare_gopurs_runtime_Value) gopurs_runtime.Value {
var dictOrd_0 *Record_compare_gopurs_runtime_Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.compare
_ = compare_1_0
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_wrappedOperation("sortBy", gopurs_runtime.Apply(pkg_Data_List.Get_sortBy(), compare_1_0), xs_2)
})
}

func Call_snoc(v_0_loop gopurs_runtime.Value, y_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var y_1 gopurs_runtime.Value = y_1_loop
_ = y_1
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty{(*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{y_1, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}})}, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1)})}
}

func Call_singleton(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty{x_0, gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_plusList(), "empty")})}
}

func Call_snoc_prime(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{v1_1, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}})}, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 786377863) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty{v1_1, gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_plusList(), "empty")})}
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

func Call_nubEq(dictEq_0_loop *Record_eq_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEq_0 *Record_eq_gopurs_runtime_Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply2(Get_wrappedOperation(), gopurs_runtime.Str("nubEq"), gopurs_runtime.Apply(pkg_Data_List.Get_nubEq(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictEq_0)}))
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

func Call_nub(dictOrd_0_loop *Record_compare_gopurs_runtime_Value) gopurs_runtime.Value {
var dictOrd_0 *Record_compare_gopurs_runtime_Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply2(Get_wrappedOperation(), gopurs_runtime.Str("nub"), gopurs_runtime.Apply(pkg_Data_List.Get_nubBy(), dictOrd_0.compare))
}

func Call_modifyAt(i_0_loop int64, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var i_0 int64 = i_0_loop
_ = i_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (i_0) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty{gopurs_runtime.Apply(f_1, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0), (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1})}})}
goto end_branch_1
} else {

}
}
{
__local_var_3_0 := (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0
_ = __local_var_3_0
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty{__local_var_3_0, x_4})}
}), gopurs_runtime.Apply3(pkg_Data_List.Get_alterAt(), gopurs_runtime.Int((i_0) - (1)), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Apply(f_1, x_4)})}
}), (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1))
}
end_branch_1:
return __t1
}

func Call_mapMaybe(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__1_0:
for {
if false { continue go__1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 786377863) {
var go__4_2 gopurs_runtime.Value
go__4_2 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__4_2:
for {
if false { continue go__4_2 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t3 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 786377863) {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})}
v1_6_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V1
continue go__4_2
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}()
})
})
__t1 = gopurs_runtime.Apply2(go__4_2, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, v_2)
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437) {
v2_4_4 := gopurs_runtime.Apply(x_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0)
_ = v2_4_4
var __t5 gopurs_runtime.Value
{
if (v2_4_4.Type == 9 && v2_4_4.IntVal == 3589588149) {
v_2_loop = v_2
v1_3_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V1
continue go__1_0
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
if (v2_4_4.Type == 9 && v2_4_4.IntVal == 930809136) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_Maybe.Data_Data_Maybe_Just)(v2_4_4.UnsafePtr).V0, v_2})}
v1_3_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V1
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
__t1 = __t5
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
}()
})
})
__local_var_2_6 := gopurs_runtime.Apply(go__1_0, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil})
_ = __local_var_2_6
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1})})
})
}

func Call_partition(x_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(pkg_Data_List.Get_partition(), x_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_1.UnsafePtr).V0, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_1.UnsafePtr).V1})})
}

func Call_span(x_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(pkg_Data_List.Get_span(), x_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_1.UnsafePtr).V0, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_1.UnsafePtr).V1})})
}

func Call_take(x_0_loop int64) gopurs_runtime.Value {
var x_0 int64 = x_0_loop
_ = x_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_List.Get_take(), gopurs_runtime.Int(x_0))
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1})})
})
}

func Call_takeWhile(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__1_0:
for {
if false { continue go__1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t3 gopurs_runtime.Value
{
if ((v1_3.Type == 9 && v1_3.IntVal == 1358893437)) && ((gopurs_runtime.Apply(x_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0).IntVal) != (0)) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0, v_2})}
v1_3_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V1
continue go__1_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
var go__4_1 gopurs_runtime.Value
go__4_1 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__4_1:
for {
if false { continue go__4_1 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 786377863) {
__t2 = v_5
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})}
v1_6_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V1
continue go__4_1
__t2 = gopurs_runtime.Value{}
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
})
__t3 = gopurs_runtime.Apply2(go__4_1, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, v_2)
}
end_branch_3:
return __t3
}
}()
})
})
__local_var_2_4 := gopurs_runtime.Apply(go__1_0, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil})
_ = __local_var_2_4
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_4, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1})})
})
}

func Call_length(v_0_loop gopurs_runtime.Value) int64 {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (1) + (gopurs_runtime.Apply(pkg_Data_List.Get_length(), (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1).IntVal)
}

func Call_last(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1358893437) {
var __t2 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)((*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1.UnsafePtr).V1
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 786377863) {
__t2 = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)((*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Apply(pkg_Data_List.Get_last(), (*pkg_Data_List_Types.Data_Data_List_Types_Cons)((*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1.UnsafePtr).V1)
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 3589588149) {
__t2 = (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Apply(pkg_Data_List.Get_last(), (*pkg_Data_List_Types.Data_Data_List_Types_Cons)((*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1.UnsafePtr).V1)
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 930809136) {
__t2 = (*pkg_Data_Maybe.Data_Data_Maybe_Just)(gopurs_runtime.Apply(pkg_Data_List.Get_last(), (*pkg_Data_List_Types.Data_Data_List_Types_Cons)((*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1.UnsafePtr).V1).UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t0 = __t2
goto end_branch_0
} else {

}
}
{
__t0 = (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V0
}
end_branch_0:
return __t0
}

func Call_intersectBy(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply2(Get_wrappedOperation2(), gopurs_runtime.Str("intersectBy"), gopurs_runtime.Apply(pkg_Data_List.Get_intersectBy(), x_0))
}

func Call_intersect(dictEq_0_loop *Record_eq_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEq_0 *Record_eq_gopurs_runtime_Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply2(Get_wrappedOperation2(), gopurs_runtime.Str("intersect"), gopurs_runtime.Apply(pkg_Data_List.Get_intersect(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictEq_0)}))
}

func Call_insertAt(i_0_loop int64, a_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var i_0 int64 = i_0_loop
_ = i_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (i_0) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty{a_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1})}})}})}
goto end_branch_1
} else {

}
}
{
__local_var_3_0 := (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0
_ = __local_var_3_0
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty{__local_var_3_0, x_4})}
}), gopurs_runtime.Apply3(pkg_Data_List.Get_insertAt(), gopurs_runtime.Int((i_0) - (1)), a_1, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1))
}
end_branch_1:
return __t1
}

func Call_init_(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
__local_var_1_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v_1, "init")
}), gopurs_runtime.Apply(pkg_Data_List.Get_unsnoc(), (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1))
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 3589588149) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V0, (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_1_0.UnsafePtr).V0})}
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

func Call_index(v_0_loop gopurs_runtime.Value, i_1_loop int64) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var i_1 int64 = i_1_loop
_ = i_1
var __t0 gopurs_runtime.Value
{
if (i_1) == (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply2(pkg_Data_List.Get_index(), (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1, gopurs_runtime.Int((i_1) - (1)))
}
end_branch_0:
return __t0
}

func Call_head(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V0
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

func Call_groupAll(dictOrd_0_loop *Record_compare_gopurs_runtime_Value) gopurs_runtime.Value {
var dictOrd_0 *Record_compare_gopurs_runtime_Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply2(Get_wrappedOperation(), gopurs_runtime.Str("groupAll"), gopurs_runtime.Apply(pkg_Data_List.Get_groupAll(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictOrd_0)}))
}

func Call_group(dictEq_0_loop *Record_eq_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEq_0 *Record_eq_gopurs_runtime_Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply2(Get_wrappedOperation(), gopurs_runtime.Str("group"), gopurs_runtime.Apply(pkg_Data_List.Get_group(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictEq_0)}))
}

func Call_fromList(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 786377863) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V1})}})}
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

func Call_fromFoldable(dictFoldable_0_loop *Record_foldMap_gopurs_runtime_Value_foldl_gopurs_runtime_Value_foldr_gopurs_runtime_Value) gopurs_runtime.Value {
var dictFoldable_0 *Record_foldMap_gopurs_runtime_Value_foldl_gopurs_runtime_Value_foldr_gopurs_runtime_Value = dictFoldable_0_loop
_ = dictFoldable_0
__local_var_1_0 := gopurs_runtime.Apply2(dictFoldable_0.foldr, pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
_ = __local_var_3_1
var __t2 gopurs_runtime.Value
{
if (__local_var_3_1.Type == 9 && __local_var_3_1.IntVal == 786377863) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_2
} else {

}
}
{
if (__local_var_3_1.Type == 9 && __local_var_3_1.IntVal == 1358893437) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(__local_var_3_1.UnsafePtr).V0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(__local_var_3_1.UnsafePtr).V1})}})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
})
}

func Call_foldM(dictMonad_0_loop *Record_, f_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value, v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
__local_var_4_0 := (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1
_ = __local_var_4_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Bind1_NOT_FOUND"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(f_1, b_2, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0), gopurs_runtime.Func(func(b_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(pkg_Data_List.Get_foldM(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, f_1, b_prime_5, __local_var_4_0)
}))
}

func Call_findLastIndex(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
v1_2_0 := gopurs_runtime.Apply2(pkg_Data_List.Get_findLastIndex(), f_0, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_1.UnsafePtr).V1)
_ = v1_2_0
var __t1 gopurs_runtime.Value
{
if (v1_2_0.Type == 9 && v1_2_0.IntVal == 930809136) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Int(((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v1_2_0.UnsafePtr).V0.IntVal) + (1))})}
goto end_branch_1
} else {

}
}
{
if (v1_2_0.Type == 9 && v1_2_0.IntVal == 3589588149) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_0, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_1.UnsafePtr).V0).IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Int(0)})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
}
end_branch_2:
__t1 = __t2
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

func Call_findIndex(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_0, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_1.UnsafePtr).V0).IntVal) != (0) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Int(0)})}
goto end_branch_3
} else {

}
}
{
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
var v1_4_loop gopurs_runtime.Value = v1_4_loop_val
go__2_0:
for {
if false { continue go__2_0 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var v1_4 gopurs_runtime.Value = v1_4_loop
_ = v1_4
var __t1 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_4.UnsafePtr).V0).IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{v_3})}
goto end_branch_2
} else {

}
}
{
v_3_loop = gopurs_runtime.Int((v_3.IntVal) + (1))
v1_4_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_4.UnsafePtr).V1
continue go__2_0
__t2 = gopurs_runtime.Value{}
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 786377863) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
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
}()
})
})
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((v1_2.IntVal) + (1))
}), gopurs_runtime.Apply2(go__2_0, gopurs_runtime.Int(0), (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_1.UnsafePtr).V1))
}
end_branch_3:
return __t3
}

func Call_filterM(dictMonad_0_loop *Record_) gopurs_runtime.Value {
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_List.Get_filterM(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1})})
})
})
}

func Call_filter(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__1_0:
for {
if false { continue go__1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 786377863) {
var go__4_2 gopurs_runtime.Value
go__4_2 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__4_2:
for {
if false { continue go__4_2 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t3 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 786377863) {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})}
v1_6_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V1
continue go__4_2
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}()
})
})
__t1 = gopurs_runtime.Apply2(go__4_2, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, v_2)
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437) {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(x_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0).IntVal) != (0) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0, v_2})}
v1_3_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V1
continue go__1_0
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
v_2_loop = v_2
v1_3_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V1
continue go__1_0
__t4 = gopurs_runtime.Value{}
}
end_branch_4:
__t1 = __t4
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
}()
})
})
__local_var_2_5 := gopurs_runtime.Apply(go__1_0, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil})
_ = __local_var_2_5
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_5, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1})})
})
}

func Call_elemLastIndex(dictEq_0_loop *Record_eq_gopurs_runtime_Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 *Record_eq_gopurs_runtime_Value = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_findLastIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictEq_0.eq, v_2, x_1)
}))
}

func Call_elemIndex(dictEq_0_loop *Record_eq_gopurs_runtime_Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 *Record_eq_gopurs_runtime_Value = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_findIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictEq_0.eq, v_2, x_1)
}))
}

func Call_dropWhile(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__1_0:
for {
if false { continue go__1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if ((v_2.Type == 9 && v_2.IntVal == 1358893437)) && ((gopurs_runtime.Apply(x_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_2.UnsafePtr).V0).IntVal) != (0)) {
v_2_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_2.UnsafePtr).V1
continue go__1_0
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
__t1 = v_2
}
end_branch_1:
return __t1
}
}()
})
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1})})
})
}

func Call_drop(x_0_loop int64, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 int64 = x_0_loop
_ = x_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(pkg_Data_List.Get_drop(), gopurs_runtime.Int(x_0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_1.UnsafePtr).V0, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_1.UnsafePtr).V1})})
}

func Call_cons_prime(x_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty{x_0, xs_1})}
}

func Call_cons(y_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var y_0 gopurs_runtime.Value = y_0_loop
_ = y_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty{y_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_1.UnsafePtr).V0, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_1.UnsafePtr).V1})}})}
}

func Call_concatMap(b_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_bindNonEmptyList(), "bind"), a_1, b_0)
}

func Call_concat(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_bindNonEmptyList(), "bind"), v_0, Get_identity())
}

func Call_catMaybes(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__1_0:
for {
if false { continue go__1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 786377863) {
var go__4_2 gopurs_runtime.Value
go__4_2 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__4_2:
for {
if false { continue go__4_2 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t3 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 786377863) {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})}
v1_6_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V1
continue go__4_2
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}()
})
})
__t1 = gopurs_runtime.Apply2(go__4_2, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, v_2)
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437) {
var __t4 gopurs_runtime.Value
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 3589588149) {
v_2_loop = v_2
v1_3_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V1
continue go__1_0
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 930809136) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_Maybe.Data_Data_Maybe_Just)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0.UnsafePtr).V0, v_2})}
v1_3_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V1
continue go__1_0
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t1 = __t4
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
}()
})
})
return gopurs_runtime.Apply2(go__1_0, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V0, (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1})})
}

func Call_appendFoldable(dictFoldable_0_loop *Record_foldMap_gopurs_runtime_Value_foldl_gopurs_runtime_Value_foldr_gopurs_runtime_Value) gopurs_runtime.Value {
var dictFoldable_0 *Record_foldMap_gopurs_runtime_Value_foldl_gopurs_runtime_Value_foldr_gopurs_runtime_Value = dictFoldable_0_loop
_ = dictFoldable_0
fromFoldable1_1_0 := gopurs_runtime.Apply2(dictFoldable_0.foldr, pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil})
_ = fromFoldable1_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty{(*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_semigroupList(), "append"), (*pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1, gopurs_runtime.Apply(fromFoldable1_1_0, ys_3))})}
})
}



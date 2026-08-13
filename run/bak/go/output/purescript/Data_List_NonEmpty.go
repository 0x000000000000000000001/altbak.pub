package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_List_NonEmpty_identity gopurs_runtime.Value
var once_Data_List_NonEmpty_identity sync.Once
func Get_Data_List_NonEmpty_identity() gopurs_runtime.Value {
	once_Data_List_NonEmpty_identity.Do(func() {
		cache_Data_List_NonEmpty_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_identity(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](x_0_box)))}
})
	})
	return cache_Data_List_NonEmpty_identity
}

var cache_Data_List_NonEmpty_zipWith gopurs_runtime.Value
var once_Data_List_NonEmpty_zipWith sync.Once
func Get_Data_List_NonEmpty_zipWith() gopurs_runtime.Value {
	once_Data_List_NonEmpty_zipWith.Do(func() {
		cache_Data_List_NonEmpty_zipWith = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_zipWith(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v1_2_box)))}
})
	})
	return cache_Data_List_NonEmpty_zipWith
}

var cache_Data_List_NonEmpty_zipWithA gopurs_runtime.Value
var once_Data_List_NonEmpty_zipWithA sync.Once
func Get_Data_List_NonEmpty_zipWithA() gopurs_runtime.Value {
	once_Data_List_NonEmpty_zipWithA.Do(func() {
		cache_Data_List_NonEmpty_zipWithA = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_zipWithA(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box))
})
	})
	return cache_Data_List_NonEmpty_zipWithA
}

var cache_Data_List_NonEmpty_zip gopurs_runtime.Value
var once_Data_List_NonEmpty_zip sync.Once
func Get_Data_List_NonEmpty_zip() gopurs_runtime.Value {
	once_Data_List_NonEmpty_zip.Do(func() {
		cache_Data_List_NonEmpty_zip = gopurs_runtime.Apply(Get_Data_List_NonEmpty_zipWith(), Get_Data_Tuple_Tuple())
	})
	return cache_Data_List_NonEmpty_zip
}

var cache_Data_List_NonEmpty_wrappedOperation2 gopurs_runtime.Value
var once_Data_List_NonEmpty_wrappedOperation2 sync.Once
func Get_Data_List_NonEmpty_wrappedOperation2() gopurs_runtime.Value {
	once_Data_List_NonEmpty_wrappedOperation2.Do(func() {
		cache_Data_List_NonEmpty_wrappedOperation2 = gopurs_runtime.Func4(func(name_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value, v1_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_wrappedOperation2(name_0_box.StrVal(), f_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_2_box), gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v1_3_box)))}
})
	})
	return cache_Data_List_NonEmpty_wrappedOperation2
}

var cache_Data_List_NonEmpty_wrappedOperation gopurs_runtime.Value
var once_Data_List_NonEmpty_wrappedOperation sync.Once
func Get_Data_List_NonEmpty_wrappedOperation() gopurs_runtime.Value {
	once_Data_List_NonEmpty_wrappedOperation.Do(func() {
		cache_Data_List_NonEmpty_wrappedOperation = gopurs_runtime.Func3(func(name_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_wrappedOperation(name_0_box.StrVal(), f_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_2_box)))}
})
	})
	return cache_Data_List_NonEmpty_wrappedOperation
}

var cache_Data_List_NonEmpty_updateAt gopurs_runtime.Value
var once_Data_List_NonEmpty_updateAt sync.Once
func Get_Data_List_NonEmpty_updateAt() gopurs_runtime.Value {
	once_Data_List_NonEmpty_updateAt.Do(func() {
		cache_Data_List_NonEmpty_updateAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_updateAt(i_0_box.IntVal, a_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_2_box)))}
})
	})
	return cache_Data_List_NonEmpty_updateAt
}

var cache_Data_List_NonEmpty_unzip gopurs_runtime.Value
var once_Data_List_NonEmpty_unzip sync.Once
func Get_Data_List_NonEmpty_unzip() gopurs_runtime.Value {
	once_Data_List_NonEmpty_unzip.Do(func() {
		cache_Data_List_NonEmpty_unzip = gopurs_runtime.Func(func(ts_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_unzip(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](ts_0_box)))}
})
	})
	return cache_Data_List_NonEmpty_unzip
}

var cache_Data_List_NonEmpty_unsnoc gopurs_runtime.Value
var once_Data_List_NonEmpty_unsnoc sync.Once
func Get_Data_List_NonEmpty_unsnoc() gopurs_runtime.Value {
	once_Data_List_NonEmpty_unsnoc.Do(func() {
		cache_Data_List_NonEmpty_unsnoc = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_unsnoc(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_0_box))
})
	})
	return cache_Data_List_NonEmpty_unsnoc
}

var cache_Data_List_NonEmpty_unionBy gopurs_runtime.Value
var once_Data_List_NonEmpty_unionBy sync.Once
func Get_Data_List_NonEmpty_unionBy() gopurs_runtime.Value {
	once_Data_List_NonEmpty_unionBy.Do(func() {
		cache_Data_List_NonEmpty_unionBy = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_unionBy(x_0_box)
})
	})
	return cache_Data_List_NonEmpty_unionBy
}

var cache_Data_List_NonEmpty_union gopurs_runtime.Value
var once_Data_List_NonEmpty_union sync.Once
func Get_Data_List_NonEmpty_union() gopurs_runtime.Value {
	once_Data_List_NonEmpty_union.Do(func() {
		cache_Data_List_NonEmpty_union = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_union(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_List_NonEmpty_union
}

var cache_Data_List_NonEmpty_uncons gopurs_runtime.Value
var once_Data_List_NonEmpty_uncons sync.Once
func Get_Data_List_NonEmpty_uncons() gopurs_runtime.Value {
	once_Data_List_NonEmpty_uncons.Do(func() {
		cache_Data_List_NonEmpty_uncons = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_uncons(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_0_box))
})
	})
	return cache_Data_List_NonEmpty_uncons
}

var cache_Data_List_NonEmpty_toList gopurs_runtime.Value
var once_Data_List_NonEmpty_toList sync.Once
func Get_Data_List_NonEmpty_toList() gopurs_runtime.Value {
	once_Data_List_NonEmpty_toList.Do(func() {
		cache_Data_List_NonEmpty_toList = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_toList(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_0_box)))}
})
	})
	return cache_Data_List_NonEmpty_toList
}

var cache_Data_List_NonEmpty_toUnfoldable gopurs_runtime.Value
var once_Data_List_NonEmpty_toUnfoldable sync.Once
func Get_Data_List_NonEmpty_toUnfoldable() gopurs_runtime.Value {
	once_Data_List_NonEmpty_toUnfoldable.Do(func() {
		cache_Data_List_NonEmpty_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dictUnfoldable_0_box))
})
	})
	return cache_Data_List_NonEmpty_toUnfoldable
}

var cache_Data_List_NonEmpty_tail gopurs_runtime.Value
var once_Data_List_NonEmpty_tail sync.Once
func Get_Data_List_NonEmpty_tail() gopurs_runtime.Value {
	once_Data_List_NonEmpty_tail.Do(func() {
		cache_Data_List_NonEmpty_tail = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_tail(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_0_box)))}
})
	})
	return cache_Data_List_NonEmpty_tail
}

var cache_Data_List_NonEmpty_sortBy gopurs_runtime.Value
var once_Data_List_NonEmpty_sortBy sync.Once
func Get_Data_List_NonEmpty_sortBy() gopurs_runtime.Value {
	once_Data_List_NonEmpty_sortBy.Do(func() {
		cache_Data_List_NonEmpty_sortBy = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_sortBy(x_0_box)
})
	})
	return cache_Data_List_NonEmpty_sortBy
}

var cache_Data_List_NonEmpty_sort gopurs_runtime.Value
var once_Data_List_NonEmpty_sort sync.Once
func Get_Data_List_NonEmpty_sort() gopurs_runtime.Value {
	once_Data_List_NonEmpty_sort.Do(func() {
		cache_Data_List_NonEmpty_sort = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_sort(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_List_NonEmpty_sort
}

var cache_Data_List_NonEmpty_snoc gopurs_runtime.Value
var once_Data_List_NonEmpty_snoc sync.Once
func Get_Data_List_NonEmpty_snoc() gopurs_runtime.Value {
	once_Data_List_NonEmpty_snoc.Do(func() {
		cache_Data_List_NonEmpty_snoc = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_snoc(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_0_box), y_1_box))}
})
	})
	return cache_Data_List_NonEmpty_snoc
}

var cache_Data_List_NonEmpty_singleton gopurs_runtime.Value
var once_Data_List_NonEmpty_singleton sync.Once
func Get_Data_List_NonEmpty_singleton() gopurs_runtime.Value {
	once_Data_List_NonEmpty_singleton.Do(func() {
		cache_Data_List_NonEmpty_singleton = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_singleton(x_0_box))}
})
	})
	return cache_Data_List_NonEmpty_singleton
}

var cache_Data_List_NonEmpty_snoc_prime gopurs_runtime.Value
var once_Data_List_NonEmpty_snoc_prime sync.Once
func Get_Data_List_NonEmpty_snoc_prime() gopurs_runtime.Value {
	once_Data_List_NonEmpty_snoc_prime.Do(func() {
		cache_Data_List_NonEmpty_snoc_prime = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_snoc_prime(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box), v1_1_box))}
})
	})
	return cache_Data_List_NonEmpty_snoc_prime
}

var cache_Data_List_NonEmpty_reverse gopurs_runtime.Value
var once_Data_List_NonEmpty_reverse sync.Once
func Get_Data_List_NonEmpty_reverse() gopurs_runtime.Value {
	once_Data_List_NonEmpty_reverse.Do(func() {
		cache_Data_List_NonEmpty_reverse = gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation(), gopurs_runtime.Str("reverse"), Get_Data_List_reverse())
	})
	return cache_Data_List_NonEmpty_reverse
}

var cache_Data_List_NonEmpty_nubEq gopurs_runtime.Value
var once_Data_List_NonEmpty_nubEq sync.Once
func Get_Data_List_NonEmpty_nubEq() gopurs_runtime.Value {
	once_Data_List_NonEmpty_nubEq.Do(func() {
		cache_Data_List_NonEmpty_nubEq = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_nubEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_List_NonEmpty_nubEq
}

var cache_Data_List_NonEmpty_nubByEq gopurs_runtime.Value
var once_Data_List_NonEmpty_nubByEq sync.Once
func Get_Data_List_NonEmpty_nubByEq() gopurs_runtime.Value {
	once_Data_List_NonEmpty_nubByEq.Do(func() {
		cache_Data_List_NonEmpty_nubByEq = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_nubByEq(x_0_box)
})
	})
	return cache_Data_List_NonEmpty_nubByEq
}

var cache_Data_List_NonEmpty_nubBy gopurs_runtime.Value
var once_Data_List_NonEmpty_nubBy sync.Once
func Get_Data_List_NonEmpty_nubBy() gopurs_runtime.Value {
	once_Data_List_NonEmpty_nubBy.Do(func() {
		cache_Data_List_NonEmpty_nubBy = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_nubBy(x_0_box)
})
	})
	return cache_Data_List_NonEmpty_nubBy
}

var cache_Data_List_NonEmpty_nub gopurs_runtime.Value
var once_Data_List_NonEmpty_nub sync.Once
func Get_Data_List_NonEmpty_nub() gopurs_runtime.Value {
	once_Data_List_NonEmpty_nub.Do(func() {
		cache_Data_List_NonEmpty_nub = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_nub(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_List_NonEmpty_nub
}

var cache_Data_List_NonEmpty_modifyAt gopurs_runtime.Value
var once_Data_List_NonEmpty_modifyAt sync.Once
func Get_Data_List_NonEmpty_modifyAt() gopurs_runtime.Value {
	once_Data_List_NonEmpty_modifyAt.Do(func() {
		cache_Data_List_NonEmpty_modifyAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_modifyAt(i_0_box.IntVal, f_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_2_box)))}
})
	})
	return cache_Data_List_NonEmpty_modifyAt
}

var cache_Data_List_NonEmpty_lift gopurs_runtime.Value
var once_Data_List_NonEmpty_lift sync.Once
func Get_Data_List_NonEmpty_lift() gopurs_runtime.Value {
	once_Data_List_NonEmpty_lift.Do(func() {
		cache_Data_List_NonEmpty_lift = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_lift(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_1_box))
})
	})
	return cache_Data_List_NonEmpty_lift
}

var cache_Data_List_NonEmpty_mapMaybe gopurs_runtime.Value
var once_Data_List_NonEmpty_mapMaybe sync.Once
func Get_Data_List_NonEmpty_mapMaybe() gopurs_runtime.Value {
	once_Data_List_NonEmpty_mapMaybe.Do(func() {
		cache_Data_List_NonEmpty_mapMaybe = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_mapMaybe(x_0_box)
})
	})
	return cache_Data_List_NonEmpty_mapMaybe
}

var cache_Data_List_NonEmpty_partition gopurs_runtime.Value
var once_Data_List_NonEmpty_partition sync.Once
func Get_Data_List_NonEmpty_partition() gopurs_runtime.Value {
	once_Data_List_NonEmpty_partition.Do(func() {
		cache_Data_List_NonEmpty_partition = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_partition(x_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_1_box))
})
	})
	return cache_Data_List_NonEmpty_partition
}

var cache_Data_List_NonEmpty_span gopurs_runtime.Value
var once_Data_List_NonEmpty_span sync.Once
func Get_Data_List_NonEmpty_span() gopurs_runtime.Value {
	once_Data_List_NonEmpty_span.Do(func() {
		cache_Data_List_NonEmpty_span = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_span(x_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_1_box))
})
	})
	return cache_Data_List_NonEmpty_span
}

var cache_Data_List_NonEmpty_take gopurs_runtime.Value
var once_Data_List_NonEmpty_take sync.Once
func Get_Data_List_NonEmpty_take() gopurs_runtime.Value {
	once_Data_List_NonEmpty_take.Do(func() {
		cache_Data_List_NonEmpty_take = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_take(x_0_box.IntVal)
})
	})
	return cache_Data_List_NonEmpty_take
}

var cache_Data_List_NonEmpty_takeWhile gopurs_runtime.Value
var once_Data_List_NonEmpty_takeWhile sync.Once
func Get_Data_List_NonEmpty_takeWhile() gopurs_runtime.Value {
	once_Data_List_NonEmpty_takeWhile.Do(func() {
		cache_Data_List_NonEmpty_takeWhile = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_takeWhile(x_0_box)
})
	})
	return cache_Data_List_NonEmpty_takeWhile
}

var cache_Data_List_NonEmpty_length gopurs_runtime.Value
var once_Data_List_NonEmpty_length sync.Once
func Get_Data_List_NonEmpty_length() gopurs_runtime.Value {
	once_Data_List_NonEmpty_length.Do(func() {
		cache_Data_List_NonEmpty_length = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_List_NonEmpty_length(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_0_box)))
})
	})
	return cache_Data_List_NonEmpty_length
}

var cache_Data_List_NonEmpty_last gopurs_runtime.Value
var once_Data_List_NonEmpty_last sync.Once
func Get_Data_List_NonEmpty_last() gopurs_runtime.Value {
	once_Data_List_NonEmpty_last.Do(func() {
		cache_Data_List_NonEmpty_last = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_last(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_0_box))
})
	})
	return cache_Data_List_NonEmpty_last
}

var cache_Data_List_NonEmpty_intersectBy gopurs_runtime.Value
var once_Data_List_NonEmpty_intersectBy sync.Once
func Get_Data_List_NonEmpty_intersectBy() gopurs_runtime.Value {
	once_Data_List_NonEmpty_intersectBy.Do(func() {
		cache_Data_List_NonEmpty_intersectBy = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_intersectBy(x_0_box)
})
	})
	return cache_Data_List_NonEmpty_intersectBy
}

var cache_Data_List_NonEmpty_intersect gopurs_runtime.Value
var once_Data_List_NonEmpty_intersect sync.Once
func Get_Data_List_NonEmpty_intersect() gopurs_runtime.Value {
	once_Data_List_NonEmpty_intersect.Do(func() {
		cache_Data_List_NonEmpty_intersect = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_intersect(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_List_NonEmpty_intersect
}

var cache_Data_List_NonEmpty_insertAt gopurs_runtime.Value
var once_Data_List_NonEmpty_insertAt sync.Once
func Get_Data_List_NonEmpty_insertAt() gopurs_runtime.Value {
	once_Data_List_NonEmpty_insertAt.Do(func() {
		cache_Data_List_NonEmpty_insertAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_insertAt(i_0_box.IntVal, a_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_2_box)))}
})
	})
	return cache_Data_List_NonEmpty_insertAt
}

var cache_Data_List_NonEmpty_init gopurs_runtime.Value
var once_Data_List_NonEmpty_init sync.Once
func Get_Data_List_NonEmpty_init() gopurs_runtime.Value {
	once_Data_List_NonEmpty_init.Do(func() {
		cache_Data_List_NonEmpty_init = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_init(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_0_box)))}
})
	})
	return cache_Data_List_NonEmpty_init
}

var cache_Data_List_NonEmpty_index gopurs_runtime.Value
var once_Data_List_NonEmpty_index sync.Once
func Get_Data_List_NonEmpty_index() gopurs_runtime.Value {
	once_Data_List_NonEmpty_index.Do(func() {
		cache_Data_List_NonEmpty_index = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, i_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_index(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_0_box), i_1_box.IntVal))}
})
	})
	return cache_Data_List_NonEmpty_index
}

var cache_Data_List_NonEmpty_head gopurs_runtime.Value
var once_Data_List_NonEmpty_head sync.Once
func Get_Data_List_NonEmpty_head() gopurs_runtime.Value {
	once_Data_List_NonEmpty_head.Do(func() {
		cache_Data_List_NonEmpty_head = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_head(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_0_box))
})
	})
	return cache_Data_List_NonEmpty_head
}

var cache_Data_List_NonEmpty_groupBy gopurs_runtime.Value
var once_Data_List_NonEmpty_groupBy sync.Once
func Get_Data_List_NonEmpty_groupBy() gopurs_runtime.Value {
	once_Data_List_NonEmpty_groupBy.Do(func() {
		cache_Data_List_NonEmpty_groupBy = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_groupBy(x_0_box)
})
	})
	return cache_Data_List_NonEmpty_groupBy
}

var cache_Data_List_NonEmpty_groupAllBy gopurs_runtime.Value
var once_Data_List_NonEmpty_groupAllBy sync.Once
func Get_Data_List_NonEmpty_groupAllBy() gopurs_runtime.Value {
	once_Data_List_NonEmpty_groupAllBy.Do(func() {
		cache_Data_List_NonEmpty_groupAllBy = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_groupAllBy(x_0_box)
})
	})
	return cache_Data_List_NonEmpty_groupAllBy
}

var cache_Data_List_NonEmpty_groupAll gopurs_runtime.Value
var once_Data_List_NonEmpty_groupAll sync.Once
func Get_Data_List_NonEmpty_groupAll() gopurs_runtime.Value {
	once_Data_List_NonEmpty_groupAll.Do(func() {
		cache_Data_List_NonEmpty_groupAll = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_groupAll(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_List_NonEmpty_groupAll
}

var cache_Data_List_NonEmpty_group gopurs_runtime.Value
var once_Data_List_NonEmpty_group sync.Once
func Get_Data_List_NonEmpty_group() gopurs_runtime.Value {
	once_Data_List_NonEmpty_group.Do(func() {
		cache_Data_List_NonEmpty_group = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_group(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_List_NonEmpty_group
}

var cache_Data_List_NonEmpty_fromList gopurs_runtime.Value
var once_Data_List_NonEmpty_fromList sync.Once
func Get_Data_List_NonEmpty_fromList() gopurs_runtime.Value {
	once_Data_List_NonEmpty_fromList.Do(func() {
		cache_Data_List_NonEmpty_fromList = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_fromList(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box)))}
})
	})
	return cache_Data_List_NonEmpty_fromList
}

var cache_Data_List_NonEmpty_fromFoldable gopurs_runtime.Value
var once_Data_List_NonEmpty_fromFoldable sync.Once
func Get_Data_List_NonEmpty_fromFoldable() gopurs_runtime.Value {
	once_Data_List_NonEmpty_fromFoldable.Do(func() {
		cache_Data_List_NonEmpty_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box))
})
	})
	return cache_Data_List_NonEmpty_fromFoldable
}

var cache_Data_List_NonEmpty_foldM gopurs_runtime.Value
var once_Data_List_NonEmpty_foldM sync.Once
func Get_Data_List_NonEmpty_foldM() gopurs_runtime.Value {
	once_Data_List_NonEmpty_foldM.Do(func() {
		cache_Data_List_NonEmpty_foldM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_foldM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Data_List_NonEmpty_foldM
}

var cache_Data_List_NonEmpty_findLastIndex gopurs_runtime.Value
var once_Data_List_NonEmpty_findLastIndex sync.Once
func Get_Data_List_NonEmpty_findLastIndex() gopurs_runtime.Value {
	once_Data_List_NonEmpty_findLastIndex.Do(func() {
		cache_Data_List_NonEmpty_findLastIndex = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_findLastIndex(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_1_box)))}
})
	})
	return cache_Data_List_NonEmpty_findLastIndex
}

var cache_Data_List_NonEmpty_findIndex gopurs_runtime.Value
var once_Data_List_NonEmpty_findIndex sync.Once
func Get_Data_List_NonEmpty_findIndex() gopurs_runtime.Value {
	once_Data_List_NonEmpty_findIndex.Do(func() {
		cache_Data_List_NonEmpty_findIndex = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_findIndex(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_1_box)))}
})
	})
	return cache_Data_List_NonEmpty_findIndex
}

var cache_Data_List_NonEmpty_filterM gopurs_runtime.Value
var once_Data_List_NonEmpty_filterM sync.Once
func Get_Data_List_NonEmpty_filterM() gopurs_runtime.Value {
	once_Data_List_NonEmpty_filterM.Do(func() {
		cache_Data_List_NonEmpty_filterM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_filterM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Data_List_NonEmpty_filterM
}

var cache_Data_List_NonEmpty_filter gopurs_runtime.Value
var once_Data_List_NonEmpty_filter sync.Once
func Get_Data_List_NonEmpty_filter() gopurs_runtime.Value {
	once_Data_List_NonEmpty_filter.Do(func() {
		cache_Data_List_NonEmpty_filter = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_filter(x_0_box)
})
	})
	return cache_Data_List_NonEmpty_filter
}

var cache_Data_List_NonEmpty_elemLastIndex gopurs_runtime.Value
var once_Data_List_NonEmpty_elemLastIndex sync.Once
func Get_Data_List_NonEmpty_elemLastIndex() gopurs_runtime.Value {
	once_Data_List_NonEmpty_elemLastIndex.Do(func() {
		cache_Data_List_NonEmpty_elemLastIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_elemLastIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), x_1_box)
})
	})
	return cache_Data_List_NonEmpty_elemLastIndex
}

var cache_Data_List_NonEmpty_elemIndex gopurs_runtime.Value
var once_Data_List_NonEmpty_elemIndex sync.Once
func Get_Data_List_NonEmpty_elemIndex() gopurs_runtime.Value {
	once_Data_List_NonEmpty_elemIndex.Do(func() {
		cache_Data_List_NonEmpty_elemIndex = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_elemIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), x_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_2_box)))}
})
	})
	return cache_Data_List_NonEmpty_elemIndex
}

var cache_Data_List_NonEmpty_dropWhile gopurs_runtime.Value
var once_Data_List_NonEmpty_dropWhile sync.Once
func Get_Data_List_NonEmpty_dropWhile() gopurs_runtime.Value {
	once_Data_List_NonEmpty_dropWhile.Do(func() {
		cache_Data_List_NonEmpty_dropWhile = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_dropWhile(x_0_box)
})
	})
	return cache_Data_List_NonEmpty_dropWhile
}

var cache_Data_List_NonEmpty_drop gopurs_runtime.Value
var once_Data_List_NonEmpty_drop sync.Once
func Get_Data_List_NonEmpty_drop() gopurs_runtime.Value {
	once_Data_List_NonEmpty_drop.Do(func() {
		cache_Data_List_NonEmpty_drop = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_drop(x_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_1_box)))}
})
	})
	return cache_Data_List_NonEmpty_drop
}

var cache_Data_List_NonEmpty_cons_prime gopurs_runtime.Value
var once_Data_List_NonEmpty_cons_prime sync.Once
func Get_Data_List_NonEmpty_cons_prime() gopurs_runtime.Value {
	once_Data_List_NonEmpty_cons_prime.Do(func() {
		cache_Data_List_NonEmpty_cons_prime = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_cons_prime(x_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_1_box)))}
})
	})
	return cache_Data_List_NonEmpty_cons_prime
}

var cache_Data_List_NonEmpty_cons gopurs_runtime.Value
var once_Data_List_NonEmpty_cons sync.Once
func Get_Data_List_NonEmpty_cons() gopurs_runtime.Value {
	once_Data_List_NonEmpty_cons.Do(func() {
		cache_Data_List_NonEmpty_cons = gopurs_runtime.Func2(func(y_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_cons(y_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_1_box)))}
})
	})
	return cache_Data_List_NonEmpty_cons
}

var cache_Data_List_NonEmpty_concatMap gopurs_runtime.Value
var once_Data_List_NonEmpty_concatMap sync.Once
func Get_Data_List_NonEmpty_concatMap() gopurs_runtime.Value {
	once_Data_List_NonEmpty_concatMap.Do(func() {
		cache_Data_List_NonEmpty_concatMap = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_concatMap(b_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](a_1_box)))}
})
	})
	return cache_Data_List_NonEmpty_concatMap
}

var cache_Data_List_NonEmpty_concat gopurs_runtime.Value
var once_Data_List_NonEmpty_concat sync.Once
func Get_Data_List_NonEmpty_concat() gopurs_runtime.Value {
	once_Data_List_NonEmpty_concat.Do(func() {
		cache_Data_List_NonEmpty_concat = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_concat(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_0_box)))}
})
	})
	return cache_Data_List_NonEmpty_concat
}

var cache_Data_List_NonEmpty_catMaybes gopurs_runtime.Value
var once_Data_List_NonEmpty_catMaybes sync.Once
func Get_Data_List_NonEmpty_catMaybes() gopurs_runtime.Value {
	once_Data_List_NonEmpty_catMaybes.Do(func() {
		cache_Data_List_NonEmpty_catMaybes = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_catMaybes(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_0_box)))}
})
	})
	return cache_Data_List_NonEmpty_catMaybes
}

var cache_Data_List_NonEmpty_appendFoldable gopurs_runtime.Value
var once_Data_List_NonEmpty_appendFoldable sync.Once
func Get_Data_List_NonEmpty_appendFoldable() gopurs_runtime.Value {
	once_Data_List_NonEmpty_appendFoldable.Do(func() {
		cache_Data_List_NonEmpty_appendFoldable = gopurs_runtime.Func3(func(dictFoldable_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_appendFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_1_box), ys_2_box))}
})
	})
	return cache_Data_List_NonEmpty_appendFoldable
}

var cache_Data_List_NonEmpty_findIndex__1077906787 gopurs_runtime.Value
var once_Data_List_NonEmpty_findIndex__1077906787 sync.Once
func Get_Data_List_NonEmpty_findIndex__1077906787() gopurs_runtime.Value {
	once_Data_List_NonEmpty_findIndex__1077906787.Do(func() {
		cache_Data_List_NonEmpty_findIndex__1077906787 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_findIndex__1077906787(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_1_box)))}
})
	})
	return cache_Data_List_NonEmpty_findIndex__1077906787
}

var cache_Data_List_NonEmpty_findLastIndex__1077906787 gopurs_runtime.Value
var once_Data_List_NonEmpty_findLastIndex__1077906787 sync.Once
func Get_Data_List_NonEmpty_findLastIndex__1077906787() gopurs_runtime.Value {
	once_Data_List_NonEmpty_findLastIndex__1077906787.Do(func() {
		cache_Data_List_NonEmpty_findLastIndex__1077906787 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_findLastIndex__1077906787(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_1_box)))}
})
	})
	return cache_Data_List_NonEmpty_findLastIndex__1077906787
}

var cache_Data_List_NonEmpty_fromList__2793484475 gopurs_runtime.Value
var once_Data_List_NonEmpty_fromList__2793484475 sync.Once
func Get_Data_List_NonEmpty_fromList__2793484475() gopurs_runtime.Value {
	once_Data_List_NonEmpty_fromList__2793484475.Do(func() {
		cache_Data_List_NonEmpty_fromList__2793484475 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_fromList__2793484475(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box)))}
})
	})
	return cache_Data_List_NonEmpty_fromList__2793484475
}

var cache_Data_List_NonEmpty_fromList__970809024 gopurs_runtime.Value
var once_Data_List_NonEmpty_fromList__970809024 sync.Once
func Get_Data_List_NonEmpty_fromList__970809024() gopurs_runtime.Value {
	once_Data_List_NonEmpty_fromList__970809024.Do(func() {
		cache_Data_List_NonEmpty_fromList__970809024 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_fromList__970809024(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box)))}
})
	})
	return cache_Data_List_NonEmpty_fromList__970809024
}

var cache_Data_List_NonEmpty_fromList__1312353984 gopurs_runtime.Value
var once_Data_List_NonEmpty_fromList__1312353984 sync.Once
func Get_Data_List_NonEmpty_fromList__1312353984() gopurs_runtime.Value {
	once_Data_List_NonEmpty_fromList__1312353984.Do(func() {
		cache_Data_List_NonEmpty_fromList__1312353984 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_fromList__1312353984(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box)))}
})
	})
	return cache_Data_List_NonEmpty_fromList__1312353984
}

var cache_Data_List_NonEmpty_lift__1094295643 gopurs_runtime.Value
var once_Data_List_NonEmpty_lift__1094295643 sync.Once
func Get_Data_List_NonEmpty_lift__1094295643() gopurs_runtime.Value {
	once_Data_List_NonEmpty_lift__1094295643.Do(func() {
		cache_Data_List_NonEmpty_lift__1094295643 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_lift__1094295643(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_1_box)))}
})
	})
	return cache_Data_List_NonEmpty_lift__1094295643
}

var cache_Data_List_NonEmpty_lift__184746683 gopurs_runtime.Value
var once_Data_List_NonEmpty_lift__184746683 sync.Once
func Get_Data_List_NonEmpty_lift__184746683() gopurs_runtime.Value {
	once_Data_List_NonEmpty_lift__184746683.Do(func() {
		cache_Data_List_NonEmpty_lift__184746683 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_lift__184746683(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_1_box))
})
	})
	return cache_Data_List_NonEmpty_lift__184746683
}

var cache_Data_List_NonEmpty_lift__243257371 gopurs_runtime.Value
var once_Data_List_NonEmpty_lift__243257371 sync.Once
func Get_Data_List_NonEmpty_lift__243257371() gopurs_runtime.Value {
	once_Data_List_NonEmpty_lift__243257371.Do(func() {
		cache_Data_List_NonEmpty_lift__243257371 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_lift__243257371(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_1_box))
})
	})
	return cache_Data_List_NonEmpty_lift__243257371
}

var cache_Data_List_NonEmpty_lift__2580243643 gopurs_runtime.Value
var once_Data_List_NonEmpty_lift__2580243643 sync.Once
func Get_Data_List_NonEmpty_lift__2580243643() gopurs_runtime.Value {
	once_Data_List_NonEmpty_lift__2580243643.Do(func() {
		cache_Data_List_NonEmpty_lift__2580243643 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_lift__2580243643(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_1_box))
})
	})
	return cache_Data_List_NonEmpty_lift__2580243643
}

var cache_Data_List_NonEmpty_lift__3667893565 gopurs_runtime.Value
var once_Data_List_NonEmpty_lift__3667893565 sync.Once
func Get_Data_List_NonEmpty_lift__3667893565() gopurs_runtime.Value {
	once_Data_List_NonEmpty_lift__3667893565.Do(func() {
		cache_Data_List_NonEmpty_lift__3667893565 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_lift__3667893565(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_1_box))
})
	})
	return cache_Data_List_NonEmpty_lift__3667893565
}

var cache_Data_List_NonEmpty_lift__1609676349 gopurs_runtime.Value
var once_Data_List_NonEmpty_lift__1609676349 sync.Once
func Get_Data_List_NonEmpty_lift__1609676349() gopurs_runtime.Value {
	once_Data_List_NonEmpty_lift__1609676349.Do(func() {
		cache_Data_List_NonEmpty_lift__1609676349 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_lift__1609676349(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_1_box)))}
})
	})
	return cache_Data_List_NonEmpty_lift__1609676349
}

var cache_Data_List_NonEmpty_singleton__3219659348 gopurs_runtime.Value
var once_Data_List_NonEmpty_singleton__3219659348 sync.Once
func Get_Data_List_NonEmpty_singleton__3219659348() gopurs_runtime.Value {
	once_Data_List_NonEmpty_singleton__3219659348.Do(func() {
		cache_Data_List_NonEmpty_singleton__3219659348 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_singleton__3219659348(x_0_box))}
})
	})
	return cache_Data_List_NonEmpty_singleton__3219659348
}

var cache_Data_List_NonEmpty_sortBy__2726669792 gopurs_runtime.Value
var once_Data_List_NonEmpty_sortBy__2726669792 sync.Once
func Get_Data_List_NonEmpty_sortBy__2726669792() gopurs_runtime.Value {
	once_Data_List_NonEmpty_sortBy__2726669792.Do(func() {
		cache_Data_List_NonEmpty_sortBy__2726669792 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_sortBy__2726669792(x_0_box)
})
	})
	return cache_Data_List_NonEmpty_sortBy__2726669792
}

var cache_Data_List_NonEmpty_toList__2859885498 gopurs_runtime.Value
var once_Data_List_NonEmpty_toList__2859885498 sync.Once
func Get_Data_List_NonEmpty_toList__2859885498() gopurs_runtime.Value {
	once_Data_List_NonEmpty_toList__2859885498.Do(func() {
		cache_Data_List_NonEmpty_toList__2859885498 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_toList__2859885498(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_0_box)))}
})
	})
	return cache_Data_List_NonEmpty_toList__2859885498
}

var cache_Data_List_NonEmpty_wrappedOperation__2422223318 gopurs_runtime.Value
var once_Data_List_NonEmpty_wrappedOperation__2422223318 sync.Once
func Get_Data_List_NonEmpty_wrappedOperation__2422223318() gopurs_runtime.Value {
	once_Data_List_NonEmpty_wrappedOperation__2422223318.Do(func() {
		cache_Data_List_NonEmpty_wrappedOperation__2422223318 = gopurs_runtime.Func3(func(name_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_wrappedOperation__2422223318(name_0_box.StrVal(), f_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_2_box)))}
})
	})
	return cache_Data_List_NonEmpty_wrappedOperation__2422223318
}

var cache_Data_List_NonEmpty_wrappedOperation__1897406582 gopurs_runtime.Value
var once_Data_List_NonEmpty_wrappedOperation__1897406582 sync.Once
func Get_Data_List_NonEmpty_wrappedOperation__1897406582() gopurs_runtime.Value {
	once_Data_List_NonEmpty_wrappedOperation__1897406582.Do(func() {
		cache_Data_List_NonEmpty_wrappedOperation__1897406582 = gopurs_runtime.Func3(func(name_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_wrappedOperation__1897406582(name_0_box.StrVal(), f_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_2_box)))}
})
	})
	return cache_Data_List_NonEmpty_wrappedOperation__1897406582
}

var cache_Data_List_NonEmpty_wrappedOperation2__4016046219 gopurs_runtime.Value
var once_Data_List_NonEmpty_wrappedOperation2__4016046219 sync.Once
func Get_Data_List_NonEmpty_wrappedOperation2__4016046219() gopurs_runtime.Value {
	once_Data_List_NonEmpty_wrappedOperation2__4016046219.Do(func() {
		cache_Data_List_NonEmpty_wrappedOperation2__4016046219 = gopurs_runtime.Func4(func(name_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value, v1_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_wrappedOperation2__4016046219(name_0_box.StrVal(), f_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_2_box), gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v1_3_box)))}
})
	})
	return cache_Data_List_NonEmpty_wrappedOperation2__4016046219
}

var cache_Data_List_NonEmpty_zipWith__1818047060 gopurs_runtime.Value
var once_Data_List_NonEmpty_zipWith__1818047060 sync.Once
func Get_Data_List_NonEmpty_zipWith__1818047060() gopurs_runtime.Value {
	once_Data_List_NonEmpty_zipWith__1818047060.Do(func() {
		cache_Data_List_NonEmpty_zipWith__1818047060 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_zipWith__1818047060(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v1_2_box)))}
})
	})
	return cache_Data_List_NonEmpty_zipWith__1818047060
}

var cache_Data_List_NonEmpty_zipWith__716394740 gopurs_runtime.Value
var once_Data_List_NonEmpty_zipWith__716394740 sync.Once
func Get_Data_List_NonEmpty_zipWith__716394740() gopurs_runtime.Value {
	once_Data_List_NonEmpty_zipWith__716394740.Do(func() {
		cache_Data_List_NonEmpty_zipWith__716394740 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_zipWith__716394740(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v1_2_box)))}
})
	})
	return cache_Data_List_NonEmpty_zipWith__716394740
}

var cache_Data_List_NonEmpty_zipWith__3875859348 gopurs_runtime.Value
var once_Data_List_NonEmpty_zipWith__3875859348 sync.Once
func Get_Data_List_NonEmpty_zipWith__3875859348() gopurs_runtime.Value {
	once_Data_List_NonEmpty_zipWith__3875859348.Do(func() {
		cache_Data_List_NonEmpty_zipWith__3875859348 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_zipWith__3875859348(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v1_2_box)))}
})
	})
	return cache_Data_List_NonEmpty_zipWith__3875859348
}

func Call_Data_List_NonEmpty_identity(x_0_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_NonEmpty_NonEmpty {
var x_0 *Constructor_Data_NonEmpty_NonEmpty = x_0_loop
_ = x_0
return x_0
}

func Call_Data_List_NonEmpty_zipWith(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty, v1_2_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_NonEmpty_NonEmpty {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty = v_1_loop
_ = v_1
var v1_2 *Constructor_Data_NonEmpty_NonEmpty = v1_2_loop
_ = v1_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply2(f_0, (v_1).V0, (v1_2).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(Get_Data_List_zipWith(), f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_1).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v1_2).V1))})))}})})
}

func Call_Data_List_NonEmpty_zipWithA(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}))
_ = Apply0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_traversable1NonEmptyList(), "sequence1"), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Apply0_1_0)}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_zipWith(f_2, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](xs_3), gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](ys_4)))})
})
})
})
}

func Call_Data_List_NonEmpty_wrappedOperation2(name_0_loop string, f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_NonEmpty_NonEmpty, v1_3_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_NonEmpty_NonEmpty {
var name_0 string = name_0_loop
_ = name_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *Constructor_Data_NonEmpty_NonEmpty = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_NonEmpty_NonEmpty = v1_3_loop
_ = v1_3
// TAST (Let): v2_4_0 -> gopurs_runtime.Value
v2_4_0 := gopurs_runtime.Apply2(f_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v_2).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_2).V1)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v1_3).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v1_3).V1)})})
_ = v2_4_0
var __t2 *Constructor_Data_NonEmpty_NonEmpty
{
if (v2_4_0.Type == 9 && v2_4_0.IntVal == 1358893437 && v2_4_0.UnsafePtr != nil) {
__t2 = &Constructor_Data_NonEmpty_NonEmpty{1, (*Constructor_Data_List_Types_Cons)(v2_4_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v2_4_0.UnsafePtr).V1)}}
goto end_branch_2
} else {

}
}
{
if (v2_4_0.Type == 9 && v2_4_0.IntVal == 1358893437 && v2_4_0.UnsafePtr == nil) {
// TAST (Let): __local_var_5_1 -> string
__local_var_5_1 := ("Impossible: empty list in NonEmptyList ") + (name_0)
_ = __local_var_5_1
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.Str(__local_var_5_1))
}))))})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__t2)})
}

func Call_Data_List_NonEmpty_wrappedOperation(name_0_loop string, f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_NonEmpty_NonEmpty {
var name_0 string = name_0_loop
_ = name_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *Constructor_Data_NonEmpty_NonEmpty = v_2_loop
_ = v_2
// TAST (Let): v1_3_0 -> *Constructor_Data_List_Types_Cons
v1_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v_2).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_2).V1)})}))
_ = v1_3_0
var __t2 *Constructor_Data_NonEmpty_NonEmpty
{
if (v1_3_0 != nil) {
__t2 = &Constructor_Data_NonEmpty_NonEmpty{1, (v1_3_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v1_3_0).V1)}}
goto end_branch_2
} else {

}
}
{
if (v1_3_0 == nil) {
// TAST (Let): __local_var_4_1 -> string
__local_var_4_1 := ("Impossible: empty list in NonEmptyList ") + (name_0)
_ = __local_var_4_1
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.Str(__local_var_4_1))
}))))})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__t2)})
}

func Call_Data_List_NonEmpty_updateAt(i_0_loop int64, a_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_Maybe_Just {
var i_0 int64 = i_0_loop
_ = i_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var v_2 *Constructor_Data_NonEmpty_NonEmpty = v_2_loop
_ = v_2
var __t1 *Constructor_Data_Maybe_Just
{
if (i_0) == (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, a_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_2).V1))}})}})})
goto end_branch_1
} else {

}
}
{
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := (v_2).V0
_ = __local_var_3_0
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, __local_var_3_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](x_4))}})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_List_updateAt(), gopurs_runtime.Int((i_0) - (1)), a_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_2).V1))})))}))
}
end_branch_1:
return __t1
}

func Call_Data_List_NonEmpty_unzip(ts_0_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_Tuple_Tuple {
var ts_0 *Constructor_Data_NonEmpty_NonEmpty = ts_0_loop
_ = ts_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_functorNonEmptyList(), "map"), Get_Data_Tuple_fst(), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(ts_0)})))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_functorNonEmptyList(), "map"), Get_Data_Tuple_snd(), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(ts_0)})))}})})
}

func Call_Data_List_NonEmpty_unsnoc(v_0_loop *Constructor_Data_NonEmpty_NonEmpty) gopurs_runtime.Value {
var v_0 *Constructor_Data_NonEmpty_NonEmpty = v_0_loop
_ = v_0
// TAST (Let): v1_1_0 -> *Constructor_Data_Maybe_Just
v1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_List_unsnoc(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_0).V1))}))
_ = v1_1_0
var __t1 gopurs_runtime.Value
{
if (v1_1_0 == nil) {
__t1 = gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, (v_0).V0)
goto end_branch_1
} else {

}
}
{
if (v1_1_0 != nil) {
__t1 = gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v_0).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet((v1_1_0).V0, "init"))})}, gopurs_runtime.RecordGet((v1_1_0).V0, "last"))
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

func Call_Data_List_NonEmpty_unionBy(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation2(), gopurs_runtime.Str("unionBy"), gopurs_runtime.Apply(Get_Data_List_unionBy(), x_0))
}

func Call_Data_List_NonEmpty_union(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation2(), gopurs_runtime.Str("union"), gopurs_runtime.Apply(Get_Data_List_union(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_0)}))
}

func Call_Data_List_NonEmpty_uncons(v_0_loop *Constructor_Data_NonEmpty_NonEmpty) gopurs_runtime.Value {
var v_0 *Constructor_Data_NonEmpty_NonEmpty = v_0_loop
_ = v_0
return gopurs_runtime.RecordDict2("head", "tail", (v_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_0).V1))})
}

func Call_Data_List_NonEmpty_toList(v_0_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_List_Types_Cons {
var v_0 *Constructor_Data_NonEmpty_NonEmpty = v_0_loop
_ = v_0
return &Constructor_Data_List_Types_Cons{1, (v_0).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_0).V1)}
}

func Call_Data_List_NonEmpty_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable = dictUnfoldable_0_loop
_ = dictUnfoldable_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable_0.V1), gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just
{
var __t_tag_1 *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_1)
if (__t_tag_1 == nil) {
__t3 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_1)
if (__t_tag_2 != nil) {
__t3 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("head", "tail", (*Constructor_Data_List_Types_Cons)(xs_1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(xs_1.UnsafePtr).V1)})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(rec_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.RecordGet(rec_2, "head"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(rec_2, "tail")))}})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)})))}
}))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V1)})})
})
}

func Call_Data_List_NonEmpty_tail(v_0_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_List_Types_Cons {
var v_0 *Constructor_Data_NonEmpty_NonEmpty = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_0).V1)
}

func Call_Data_List_NonEmpty_sortBy(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation(), gopurs_runtime.Str("sortBy"), gopurs_runtime.Apply(Get_Data_List_sortBy(), x_0))
}

func Call_Data_List_NonEmpty_sort(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_wrappedOperation("sortBy", gopurs_runtime.Apply(Get_Data_List_sortBy(), compare_1_0), gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](xs_2)))}
})
}

func Call_Data_List_NonEmpty_snoc(v_0_loop *Constructor_Data_NonEmpty_NonEmpty, y_1_loop gopurs_runtime.Value) *Constructor_Data_NonEmpty_NonEmpty {
var v_0 *Constructor_Data_NonEmpty_NonEmpty = v_0_loop
_ = v_0
var y_1 gopurs_runtime.Value = y_1_loop
_ = y_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (v_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldr"), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, y_1, (*Constructor_Data_List_Types_Cons)(nil)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_0).V1))})))}})})
}

func Call_Data_List_NonEmpty_singleton(x_0_loop gopurs_runtime.Value) *Constructor_Data_NonEmpty_NonEmpty {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, x_0, gopurs_runtime.RecordGet(Get_Data_List_Types_plusList(), "empty")})})
}

func Call_Data_List_NonEmpty_snoc_prime(v_0_loop *Constructor_Data_List_Types_Cons, v1_1_loop gopurs_runtime.Value) *Constructor_Data_NonEmpty_NonEmpty {
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_NonEmpty_NonEmpty
{
if (v_0 != nil) {
__t0 = &Constructor_Data_NonEmpty_NonEmpty{1, (v_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldr"), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, v1_1, (*Constructor_Data_List_Types_Cons)(nil)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)})))}}
goto end_branch_0
} else {

}
}
{
if (v_0 == nil) {
__t0 = &Constructor_Data_NonEmpty_NonEmpty{1, v1_1, gopurs_runtime.RecordGet(Get_Data_List_Types_plusList(), "empty")}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__t0)})
}

func Call_Data_List_NonEmpty_nubEq(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation(), gopurs_runtime.Str("nubEq"), gopurs_runtime.Apply(Get_Data_List_nubEq(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_0)}))
}

func Call_Data_List_NonEmpty_nubByEq(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation(), gopurs_runtime.Str("nubByEq"), gopurs_runtime.Apply(Get_Data_List_nubByEq(), x_0))
}

func Call_Data_List_NonEmpty_nubBy(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation(), gopurs_runtime.Str("nubBy"), gopurs_runtime.Apply(Get_Data_List_nubBy(), x_0))
}

func Call_Data_List_NonEmpty_nub(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation(), gopurs_runtime.Str("nub"), gopurs_runtime.Apply(Get_Data_List_nubBy(), gopurs_runtime.Box(dictOrd_0.V1)))
}

func Call_Data_List_NonEmpty_modifyAt(i_0_loop int64, f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_Maybe_Just {
var i_0 int64 = i_0_loop
_ = i_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *Constructor_Data_NonEmpty_NonEmpty = v_2_loop
_ = v_2
var __t1 *Constructor_Data_Maybe_Just
{
if (i_0) == (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(f_1, (v_2).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_2).V1))}})}})})
goto end_branch_1
} else {

}
}
{
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := (v_2).V0
_ = __local_var_3_0
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, __local_var_3_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](x_4))}})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_List_alterAt(), gopurs_runtime.Int((i_0) - (1)), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_1, x_4)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_2).V1))})))}))
}
end_branch_1:
return __t1
}

func Call_Data_List_NonEmpty_lift(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_1).V1)})})
}

func Call_Data_List_NonEmpty_mapMaybe(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_List_mapMaybe(), x_0)
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)})})
})
}

func Call_Data_List_NonEmpty_partition(x_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(Get_Data_List_partition(), x_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_1).V1)})})
}

func Call_Data_List_NonEmpty_span(x_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(Get_Data_List_span(), x_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_1).V1)})})
}

func Call_Data_List_NonEmpty_take(x_0_loop int64) gopurs_runtime.Value {
var x_0 int64 = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_List_take(), gopurs_runtime.Int(x_0))
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)})})
})
}

func Call_Data_List_NonEmpty_takeWhile(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var go__go_1_1_0 gopurs_runtime.Value
go__go_1_1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_1_0:
for {
if false { continue go__go_1_1_0 }
var v_2 *Constructor_Data_List_Types_Cons = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t4 *Constructor_Data_List_Types_Cons
{
if ((v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil)) && ((gopurs_runtime.Apply(x_0, (*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0).IntVal) != (0)) {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0, v_2})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}
continue go__go_1_1_0
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_4
} else {

}
}
{
var go__go_4_2_1 gopurs_runtime.Value
go__go_4_2_1 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_2_1:
for {
if false { continue go__go_4_2_1 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t3 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_2_1
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_2_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}))
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t4)}
}
}()
})
})
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(go__go_1_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)})})
})
}

func Call_Data_List_NonEmpty_length(v_0_loop *Constructor_Data_NonEmpty_NonEmpty) int64 {
var v_0 *Constructor_Data_NonEmpty_NonEmpty = v_0_loop
_ = v_0
return (1) + (gopurs_runtime.Apply(Get_Data_List_length(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_0).V1))}).IntVal)
}

func Call_Data_List_NonEmpty_last(v_0_loop *Constructor_Data_NonEmpty_NonEmpty) gopurs_runtime.Value {
var v_0 *Constructor_Data_NonEmpty_NonEmpty = v_0_loop
_ = v_0
var __t4 *Constructor_Data_Maybe_Just
{
var __t_tag_1 *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_0).V1)
if (__t_tag_1 != nil) {
var __t3 *Constructor_Data_Maybe_Just
{
var __t_tag_2 *Constructor_Data_List_Types_Cons = (*Constructor_Data_List_Types_Cons)((v_0).V1.UnsafePtr).V1
if (__t_tag_2 == nil) {
__t3 = &Constructor_Data_Maybe_Just{1, (*Constructor_Data_List_Types_Cons)((v_0).V1.UnsafePtr).V0}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_List_last(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)((v_0).V1.UnsafePtr).V1)}))
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_4:
// TAST (Let): __local_var_1_0 -> *Constructor_Data_Maybe_Just
var __local_var_1_0 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)})
var __t5 gopurs_runtime.Value
{
if (__local_var_1_0 == nil) {
__t5 = (v_0).V0
goto end_branch_5
} else {

}
}
{
if (__local_var_1_0 != nil) {
__t5 = (__local_var_1_0).V0
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

func Call_Data_List_NonEmpty_intersectBy(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation2(), gopurs_runtime.Str("intersectBy"), gopurs_runtime.Apply(Get_Data_List_intersectBy(), x_0))
}

func Call_Data_List_NonEmpty_intersect(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation2(), gopurs_runtime.Str("intersect"), gopurs_runtime.Apply(Get_Data_List_intersect(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_0)}))
}

func Call_Data_List_NonEmpty_insertAt(i_0_loop int64, a_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_Maybe_Just {
var i_0 int64 = i_0_loop
_ = i_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var v_2 *Constructor_Data_NonEmpty_NonEmpty = v_2_loop
_ = v_2
var __t1 *Constructor_Data_Maybe_Just
{
if (i_0) == (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, a_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v_2).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_2).V1)})}})}})})
goto end_branch_1
} else {

}
}
{
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := (v_2).V0
_ = __local_var_3_0
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, __local_var_3_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](x_4))}})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_List_insertAt(), gopurs_runtime.Int((i_0) - (1)), a_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_2).V1))})))}))
}
end_branch_1:
return __t1
}

func Call_Data_List_NonEmpty_init(v_0_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_List_Types_Cons {
var v_0 *Constructor_Data_NonEmpty_NonEmpty = v_0_loop
_ = v_0
// TAST (Let): __local_var_1_0 -> *Constructor_Data_Maybe_Just
__local_var_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v_1, "init")))}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_List_unsnoc(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_0).V1))})))}))
_ = __local_var_1_0
var __t1 *Constructor_Data_List_Types_Cons
{
if (__local_var_1_0 == nil) {
__t1 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_1
} else {

}
}
{
if (__local_var_1_0 != nil) {
__t1 = &Constructor_Data_List_Types_Cons{1, (v_0).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((__local_var_1_0).V0)}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return __t1
}

func Call_Data_List_NonEmpty_index(v_0_loop *Constructor_Data_NonEmpty_NonEmpty, i_1_loop int64) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_NonEmpty_NonEmpty = v_0_loop
_ = v_0
var i_1 int64 = i_1_loop
_ = i_1
var __t0 *Constructor_Data_Maybe_Just
{
if (i_1) == (0) {
__t0 = &Constructor_Data_Maybe_Just{1, (v_0).V0}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(Get_Data_List_index(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_0).V1))}, gopurs_runtime.Int((i_1) - (1))))
}
end_branch_0:
return __t0
}

func Call_Data_List_NonEmpty_head(v_0_loop *Constructor_Data_NonEmpty_NonEmpty) gopurs_runtime.Value {
var v_0 *Constructor_Data_NonEmpty_NonEmpty = v_0_loop
_ = v_0
return (v_0).V0
}

func Call_Data_List_NonEmpty_groupBy(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation(), gopurs_runtime.Str("groupBy"), gopurs_runtime.Apply(Get_Data_List_groupBy(), x_0))
}

func Call_Data_List_NonEmpty_groupAllBy(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation(), gopurs_runtime.Str("groupAllBy"), gopurs_runtime.Apply(Get_Data_List_groupAllBy(), x_0))
}

func Call_Data_List_NonEmpty_groupAll(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation(), gopurs_runtime.Str("groupAll"), gopurs_runtime.Apply(Get_Data_List_groupAll(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}))
}

func Call_Data_List_NonEmpty_group(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation(), gopurs_runtime.Str("group"), gopurs_runtime.Apply(Get_Data_List_group(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_0)}))
}

func Call_Data_List_NonEmpty_fromList(v_0_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 == nil) {
__t0 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_0
} else {

}
}
{
if (v_0 != nil) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (v_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)}})}}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)})
}

func Call_Data_List_NonEmpty_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V2), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
_ = __local_var_3_1
var __t2 *Constructor_Data_Maybe_Just
{
if (__local_var_3_1.Type == 9 && __local_var_3_1.IntVal == 1358893437 && __local_var_3_1.UnsafePtr == nil) {
__t2 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_2
} else {

}
}
{
if (__local_var_3_1.Type == 9 && __local_var_3_1.IntVal == 1358893437 && __local_var_3_1.UnsafePtr != nil) {
__t2 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (*Constructor_Data_List_Types_Cons)(__local_var_3_1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(__local_var_3_1.UnsafePtr).V1)}})}}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
})
}

func Call_Data_List_NonEmpty_foldM(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_1 -> gopurs_runtime.Value
__local_var_5_1 := (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1
_ = __local_var_5_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply2(f_2, b_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0), gopurs_runtime.Func(func(b_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(Get_Data_List_foldM(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)}, f_2, b_prime_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](__local_var_5_1))})
}))
})
})
})
}

func Call_Data_List_NonEmpty_findLastIndex(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_Maybe_Just {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty = v_1_loop
_ = v_1
// TAST (Let): v1_2_0 -> gopurs_runtime.Value
v1_2_0 := gopurs_runtime.Apply2(Get_Data_List_findLastIndex(), f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_1).V1))})
_ = v1_2_0
var __t2 *Constructor_Data_Maybe_Just
{
if (v1_2_0.Type == 9 && v1_2_0.IntVal == 930809136 && v1_2_0.UnsafePtr != nil) {
__t2 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(((*Constructor_Data_Maybe_Just)(v1_2_0.UnsafePtr).V0.IntVal) + (1))}
goto end_branch_2
} else {

}
}
{
if (v1_2_0.Type == 9 && v1_2_0.IntVal == 930809136 && v1_2_0.UnsafePtr == nil) {
var __t1 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.Apply(f_0, (v_1).V0).IntVal) != (0) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(0)}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)})
}

func Call_Data_List_NonEmpty_findIndex(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_Maybe_Just {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty = v_1_loop
_ = v_1
var __t5 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.Apply(f_0, (v_1).V0).IntVal) != (0) {
__t5 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(0)}
goto end_branch_5
} else {

}
}
{
var go__go_2_1_2 gopurs_runtime.Value
go__go_2_1_2 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop int64 = v_3_loop_val.IntVal
var v1_4_loop gopurs_runtime.Value = v1_4_loop_val
go__go_2_1_2:
for {
if false { continue go__go_2_1_2 }
var v_3 int64 = v_3_loop
_ = v_3
var v1_4 gopurs_runtime.Value = v1_4_loop
_ = v1_4
var __t3 *Constructor_Data_Maybe_Just
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr != nil) {
var __t2 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons)(v1_4.UnsafePtr).V0).IntVal) != (0) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(v_3)})})
goto end_branch_2
} else {

}
}
{
v_3_loop = (v_3) + (1)
v1_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_4.UnsafePtr).V1)}
continue go__go_2_1_2
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
// TAST (Let): __local_var_2_0 -> *Constructor_Data_Maybe_Just
__local_var_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(go__go_2_1_2, gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_1).V1))}))
_ = __local_var_2_0
var __t4 *Constructor_Data_Maybe_Just
{
if (__local_var_2_0 != nil) {
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(((__local_var_2_0).V0.IntVal) + (1))}
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_4:
__t5 = __t4
}
end_branch_5:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)})
}

func Call_Data_List_NonEmpty_filterM(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_List_filterM(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1)})})
})
})
}

func Call_Data_List_NonEmpty_filter(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_List_filter(), x_0)
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)})})
})
}

func Call_Data_List_NonEmpty_elemLastIndex(dictEq_0_loop *Constructor_Data_Eq_Eq, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_Data_List_NonEmpty_findLastIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), v_2, x_1).IntVal) != (0))
}))
}

func Call_Data_List_NonEmpty_elemIndex(dictEq_0_loop *Constructor_Data_Eq_Eq, x_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_Maybe_Just {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var v_2 *Constructor_Data_NonEmpty_NonEmpty = v_2_loop
_ = v_2
var __t5 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), (v_2).V0, x_1).IntVal) != (0) {
__t5 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(0)}
goto end_branch_5
} else {

}
}
{
var go__go_3_1_3 gopurs_runtime.Value
go__go_3_1_3 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop int64 = v_4_loop_val.IntVal
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_1_3:
for {
if false { continue go__go_3_1_3 }
var v_4 int64 = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 *Constructor_Data_Maybe_Just
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
var __t2 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, x_1).IntVal) != (0) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(v_4)})})
goto end_branch_2
} else {

}
}
{
v_4_loop = (v_4) + (1)
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_1_3
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
// TAST (Let): __local_var_3_0 -> *Constructor_Data_Maybe_Just
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(go__go_3_1_3, gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_2).V1))}))
_ = __local_var_3_0
var __t4 *Constructor_Data_Maybe_Just
{
if (__local_var_3_0 != nil) {
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(((__local_var_3_0).V0.IntVal) + (1))}
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_4:
__t5 = __t4
}
end_branch_5:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)})
}

func Call_Data_List_NonEmpty_dropWhile(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var go__go_1_1_4 gopurs_runtime.Value
go__go_1_1_4 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2_loop_val)
go__go_1_1_4:
for {
if false { continue go__go_1_1_4 }
var v_2 *Constructor_Data_List_Types_Cons = v_2_loop
_ = v_2
var __t2 *Constructor_Data_List_Types_Cons
{
if ((v_2 != nil)) && ((gopurs_runtime.Apply(x_0, (v_2).V0).IntVal) != (0)) {
v_2_loop = (v_2).V1
continue go__go_1_1_4
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
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
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := go__go_1_1_4
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)})})
})
}

func Call_Data_List_NonEmpty_drop(x_0_loop int64, v_1_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_List_Types_Cons {
var x_0 int64 = x_0_loop
_ = x_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(Get_Data_List_drop(), gopurs_runtime.Int(x_0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_1).V1)})}))
}

func Call_Data_List_NonEmpty_cons_prime(x_0_loop gopurs_runtime.Value, xs_1_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_NonEmpty_NonEmpty {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 *Constructor_Data_List_Types_Cons = xs_1_loop
_ = xs_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, x_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}})})
}

func Call_Data_List_NonEmpty_cons(y_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_NonEmpty_NonEmpty {
var y_0 gopurs_runtime.Value = y_0_loop
_ = y_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, y_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_1).V1)})}})})
}

func Call_Data_List_NonEmpty_concatMap(b_0_loop gopurs_runtime.Value, a_1_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_NonEmpty_NonEmpty {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 *Constructor_Data_NonEmpty_NonEmpty = a_1_loop
_ = a_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_bindNonEmptyList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(a_1)}, b_0))
}

func Call_Data_List_NonEmpty_concat(v_0_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_NonEmpty_NonEmpty {
var v_0 *Constructor_Data_NonEmpty_NonEmpty = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply2(Get_Control_Bind_bind__2389430209(), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
})))
}

func Call_Data_List_NonEmpty_catMaybes(v_0_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_List_Types_Cons {
var v_0 *Constructor_Data_NonEmpty_NonEmpty = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(Get_Data_List_catMaybes(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v_0).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_0).V1)})}))
}

func Call_Data_List_NonEmpty_appendFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, v_1_loop *Constructor_Data_NonEmpty_NonEmpty, ys_2_loop gopurs_runtime.Value) *Constructor_Data_NonEmpty_NonEmpty {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty = v_1_loop
_ = v_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (v_1).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_1).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable_0.V2), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, ys_2)))})))}})})
}

func Call_Data_List_NonEmpty_findIndex__1077906787(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_Maybe_Just {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty = v_1_loop
_ = v_1
var __t5 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.Apply(f_0, (v_1).V0).IntVal) != (0) {
__t5 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(0)}
goto end_branch_5
} else {

}
}
{
var go__go_2_1_5 gopurs_runtime.Value
go__go_2_1_5 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop int64 = v_3_loop_val.IntVal
var v1_4_loop gopurs_runtime.Value = v1_4_loop_val
go__go_2_1_5:
for {
if false { continue go__go_2_1_5 }
var v_3 int64 = v_3_loop
_ = v_3
var v1_4 gopurs_runtime.Value = v1_4_loop
_ = v1_4
var __t3 *Constructor_Data_Maybe_Just
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr != nil) {
var __t2 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons)(v1_4.UnsafePtr).V0).IntVal) != (0) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(v_3)})})
goto end_branch_2
} else {

}
}
{
v_3_loop = (v_3) + (1)
v1_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_4.UnsafePtr).V1)}
continue go__go_2_1_5
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
// TAST (Let): __local_var_2_0 -> *Constructor_Data_Maybe_Just
__local_var_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(go__go_2_1_5, gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_1).V1))}))
_ = __local_var_2_0
var __t4 *Constructor_Data_Maybe_Just
{
if (__local_var_2_0 != nil) {
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(((__local_var_2_0).V0.IntVal) + (1))}
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_4:
__t5 = __t4
}
end_branch_5:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)})
}

func Call_Data_List_NonEmpty_findLastIndex__1077906787(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_Maybe_Just {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty = v_1_loop
_ = v_1
// TAST (Let): v1_2_0 -> gopurs_runtime.Value
v1_2_0 := gopurs_runtime.Apply2(Get_Data_List_findLastIndex(), f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_1).V1))})
_ = v1_2_0
var __t2 *Constructor_Data_Maybe_Just
{
if (v1_2_0.Type == 9 && v1_2_0.IntVal == 930809136 && v1_2_0.UnsafePtr != nil) {
__t2 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(((*Constructor_Data_Maybe_Just)(v1_2_0.UnsafePtr).V0.IntVal) + (1))}
goto end_branch_2
} else {

}
}
{
if (v1_2_0.Type == 9 && v1_2_0.IntVal == 930809136 && v1_2_0.UnsafePtr == nil) {
var __t1 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.Apply(f_0, (v_1).V0).IntVal) != (0) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(0)}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)})
}

func Call_Data_List_NonEmpty_fromList__2793484475(v_0_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 == nil) {
__t0 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_0
} else {

}
}
{
if (v_0 != nil) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (v_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)}})}}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)})
}

func Call_Data_List_NonEmpty_fromList__970809024(v_0_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 == nil) {
__t0 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_0
} else {

}
}
{
if (v_0 != nil) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (v_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)}})}}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)})
}

func Call_Data_List_NonEmpty_fromList__1312353984(v_0_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 == nil) {
__t0 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_0
} else {

}
}
{
if (v_0 != nil) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (v_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)}})}}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)})
}

func Call_Data_List_NonEmpty_lift__1094295643(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_List_Types_Cons {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_1).V1)})}))
}

func Call_Data_List_NonEmpty_lift__184746683(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_1).V1)})})
}

func Call_Data_List_NonEmpty_lift__243257371(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_1).V1)})})
}

func Call_Data_List_NonEmpty_lift__2580243643(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_1).V1)})})
}

func Call_Data_List_NonEmpty_lift__3667893565(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_1).V1)})})
}

func Call_Data_List_NonEmpty_lift__1609676349(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_List_Types_Cons {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((v_1).V0))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_1).V1))})})}))
}

func Call_Data_List_NonEmpty_singleton__3219659348(x_0_loop gopurs_runtime.Value) *Constructor_Data_NonEmpty_NonEmpty {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, x_0, gopurs_runtime.RecordGet(Get_Data_List_Types_plusList(), "empty")})})
}

func Call_Data_List_NonEmpty_sortBy__2726669792(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation(), gopurs_runtime.Str("sortBy"), gopurs_runtime.Apply(Get_Data_List_sortBy(), x_0))
}

func Call_Data_List_NonEmpty_toList__2859885498(v_0_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_List_Types_Cons {
var v_0 *Constructor_Data_NonEmpty_NonEmpty = v_0_loop
_ = v_0
return &Constructor_Data_List_Types_Cons{1, (v_0).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_0).V1)}
}

func Call_Data_List_NonEmpty_wrappedOperation__2422223318(name_0_loop string, f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_NonEmpty_NonEmpty {
var name_0 string = name_0_loop
_ = name_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *Constructor_Data_NonEmpty_NonEmpty = v_2_loop
_ = v_2
// TAST (Let): v1_3_0 -> *Constructor_Data_List_Types_Cons
v1_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v_2).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_2).V1)})}))
_ = v1_3_0
var __t2 *Constructor_Data_NonEmpty_NonEmpty
{
if (v1_3_0 != nil) {
__t2 = &Constructor_Data_NonEmpty_NonEmpty{1, (v1_3_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v1_3_0).V1)}}
goto end_branch_2
} else {

}
}
{
if (v1_3_0 == nil) {
// TAST (Let): __local_var_4_1 -> string
__local_var_4_1 := ("Impossible: empty list in NonEmptyList ") + (name_0)
_ = __local_var_4_1
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.Str(__local_var_4_1))
}))))})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__t2)})
}

func Call_Data_List_NonEmpty_wrappedOperation__1897406582(name_0_loop string, f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_NonEmpty_NonEmpty {
var name_0 string = name_0_loop
_ = name_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *Constructor_Data_NonEmpty_NonEmpty = v_2_loop
_ = v_2
// TAST (Let): v1_3_0 -> *Constructor_Data_List_Types_Cons
v1_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v_2).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_2).V1)})}))
_ = v1_3_0
var __t2 *Constructor_Data_NonEmpty_NonEmpty
{
if (v1_3_0 != nil) {
__t2 = &Constructor_Data_NonEmpty_NonEmpty{1, (v1_3_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v1_3_0).V1)}}
goto end_branch_2
} else {

}
}
{
if (v1_3_0 == nil) {
// TAST (Let): __local_var_4_1 -> string
__local_var_4_1 := ("Impossible: empty list in NonEmptyList ") + (name_0)
_ = __local_var_4_1
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.Str(__local_var_4_1))
}))))})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__t2)})
}

func Call_Data_List_NonEmpty_wrappedOperation2__4016046219(name_0_loop string, f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_NonEmpty_NonEmpty, v1_3_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_NonEmpty_NonEmpty {
var name_0 string = name_0_loop
_ = name_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *Constructor_Data_NonEmpty_NonEmpty = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_NonEmpty_NonEmpty = v1_3_loop
_ = v1_3
// TAST (Let): v2_4_0 -> gopurs_runtime.Value
v2_4_0 := gopurs_runtime.Apply2(f_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v_2).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_2).V1)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v1_3).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v1_3).V1)})})
_ = v2_4_0
var __t2 *Constructor_Data_NonEmpty_NonEmpty
{
if (v2_4_0.Type == 9 && v2_4_0.IntVal == 1358893437 && v2_4_0.UnsafePtr != nil) {
__t2 = &Constructor_Data_NonEmpty_NonEmpty{1, (*Constructor_Data_List_Types_Cons)(v2_4_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v2_4_0.UnsafePtr).V1)}}
goto end_branch_2
} else {

}
}
{
if (v2_4_0.Type == 9 && v2_4_0.IntVal == 1358893437 && v2_4_0.UnsafePtr == nil) {
// TAST (Let): __local_var_5_1 -> string
__local_var_5_1 := ("Impossible: empty list in NonEmptyList ") + (name_0)
_ = __local_var_5_1
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.Str(__local_var_5_1))
}))))})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__t2)})
}

func Call_Data_List_NonEmpty_zipWith__1818047060(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty, v1_2_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_NonEmpty_NonEmpty {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty = v_1_loop
_ = v_1
var v1_2 *Constructor_Data_NonEmpty_NonEmpty = v1_2_loop
_ = v1_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply2(f_0, (v_1).V0, (v1_2).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(Get_Data_List_zipWith(), f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_1).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v1_2).V1))})))}})})
}

func Call_Data_List_NonEmpty_zipWith__716394740(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty, v1_2_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_NonEmpty_NonEmpty {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty = v_1_loop
_ = v_1
var v1_2 *Constructor_Data_NonEmpty_NonEmpty = v1_2_loop
_ = v1_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(f_0, (v_1).V0, (v1_2).V0)))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(Get_Data_List_zipWith(), f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_1).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v1_2).V1))})))}})})
}

func Call_Data_List_NonEmpty_zipWith__3875859348(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty, v1_2_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_NonEmpty_NonEmpty {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty = v_1_loop
_ = v_1
var v1_2 *Constructor_Data_NonEmpty_NonEmpty = v1_2_loop
_ = v1_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply2(f_0, (v_1).V0, (v1_2).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(Get_Data_List_zipWith(), f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_1).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v1_2).V1))})))}})})
}



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
		cache_Data_List_NonEmpty_identity = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
	})
	return cache_Data_List_NonEmpty_identity
}

var cache_Data_List_NonEmpty_zipWith gopurs_runtime.Value
var once_Data_List_NonEmpty_zipWith sync.Once
func Get_Data_List_NonEmpty_zipWith() gopurs_runtime.Value {
	once_Data_List_NonEmpty_zipWith.Do(func() {
		cache_Data_List_NonEmpty_zipWith = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Call_Data_List_NonEmpty_zipWith(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v1_2_box))))}
})
	})
	return cache_Data_List_NonEmpty_zipWith
}

var cache_Data_List_NonEmpty_zipWithA gopurs_runtime.Value
var once_Data_List_NonEmpty_zipWithA sync.Once
func Get_Data_List_NonEmpty_zipWithA() gopurs_runtime.Value {
	once_Data_List_NonEmpty_zipWithA.Do(func() {
		cache_Data_List_NonEmpty_zipWithA = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_zipWithA(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
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
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Call_Data_List_NonEmpty_wrappedOperation2(name_0_box.StrVal(), f_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2_box), gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v1_3_box))))}
})
	})
	return cache_Data_List_NonEmpty_wrappedOperation2
}

var cache_Data_List_NonEmpty_wrappedOperation gopurs_runtime.Value
var once_Data_List_NonEmpty_wrappedOperation sync.Once
func Get_Data_List_NonEmpty_wrappedOperation() gopurs_runtime.Value {
	once_Data_List_NonEmpty_wrappedOperation.Do(func() {
		cache_Data_List_NonEmpty_wrappedOperation = gopurs_runtime.Func3(func(name_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Call_Data_List_NonEmpty_wrappedOperation(name_0_box.StrVal(), f_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2_box))))}
})
	})
	return cache_Data_List_NonEmpty_wrappedOperation
}

var cache_Data_List_NonEmpty_updateAt gopurs_runtime.Value
var once_Data_List_NonEmpty_updateAt sync.Once
func Get_Data_List_NonEmpty_updateAt() gopurs_runtime.Value {
	once_Data_List_NonEmpty_updateAt.Do(func() {
		cache_Data_List_NonEmpty_updateAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_NonEmpty_updateAt(i_0_box.IntVal, a_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_NonEmpty_updateAt
}

var cache_Data_List_NonEmpty_unzip gopurs_runtime.Value
var once_Data_List_NonEmpty_unzip sync.Once
func Get_Data_List_NonEmpty_unzip() gopurs_runtime.Value {
	once_Data_List_NonEmpty_unzip.Do(func() {
		cache_Data_List_NonEmpty_unzip = gopurs_runtime.Func(func(ts_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_NonEmpty_unzip(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](ts_0_box))
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()
})
	})
	return cache_Data_List_NonEmpty_unzip
}

var cache_Data_List_NonEmpty_unsnoc gopurs_runtime.Value
var once_Data_List_NonEmpty_unsnoc sync.Once
func Get_Data_List_NonEmpty_unsnoc() gopurs_runtime.Value {
	once_Data_List_NonEmpty_unsnoc.Do(func() {
		cache_Data_List_NonEmpty_unsnoc = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := Call_Data_List_NonEmpty_unsnoc(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box))
				_ = orig
				return gopurs_runtime.RecordDict([]string{"init", "last"}, []gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.go__init)}, orig.last})
				}()
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
return Call_Data_List_NonEmpty_union(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_Data_List_NonEmpty_union
}

var cache_Data_List_NonEmpty_uncons gopurs_runtime.Value
var once_Data_List_NonEmpty_uncons sync.Once
func Get_Data_List_NonEmpty_uncons() gopurs_runtime.Value {
	once_Data_List_NonEmpty_uncons.Do(func() {
		cache_Data_List_NonEmpty_uncons = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := Call_Data_List_NonEmpty_uncons(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box))
				_ = orig
				return gopurs_runtime.RecordDict([]string{"head", "tail"}, []gopurs_runtime.Value{orig.head, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.tail)}})
				}()
})
	})
	return cache_Data_List_NonEmpty_uncons
}

var cache_Data_List_NonEmpty_toList gopurs_runtime.Value
var once_Data_List_NonEmpty_toList sync.Once
func Get_Data_List_NonEmpty_toList() gopurs_runtime.Value {
	once_Data_List_NonEmpty_toList.Do(func() {
		cache_Data_List_NonEmpty_toList = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_toList(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_Data_List_NonEmpty_toList
}

var cache_Data_List_NonEmpty_toUnfoldable gopurs_runtime.Value
var once_Data_List_NonEmpty_toUnfoldable sync.Once
func Get_Data_List_NonEmpty_toUnfoldable() gopurs_runtime.Value {
	once_Data_List_NonEmpty_toUnfoldable.Do(func() {
		cache_Data_List_NonEmpty_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
})
	})
	return cache_Data_List_NonEmpty_toUnfoldable
}

var cache_Data_List_NonEmpty_tail gopurs_runtime.Value
var once_Data_List_NonEmpty_tail sync.Once
func Get_Data_List_NonEmpty_tail() gopurs_runtime.Value {
	once_Data_List_NonEmpty_tail.Do(func() {
		cache_Data_List_NonEmpty_tail = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_tail(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box)))}
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
return Call_Data_List_NonEmpty_sort(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_List_NonEmpty_sort
}

var cache_Data_List_NonEmpty_snoc gopurs_runtime.Value
var once_Data_List_NonEmpty_snoc sync.Once
func Get_Data_List_NonEmpty_snoc() gopurs_runtime.Value {
	once_Data_List_NonEmpty_snoc.Do(func() {
		cache_Data_List_NonEmpty_snoc = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Call_Data_List_NonEmpty_snoc(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box), y_1_box)))}
})
	})
	return cache_Data_List_NonEmpty_snoc
}

var cache_Data_List_NonEmpty_singleton gopurs_runtime.Value
var once_Data_List_NonEmpty_singleton sync.Once
func Get_Data_List_NonEmpty_singleton() gopurs_runtime.Value {
	once_Data_List_NonEmpty_singleton.Do(func() {
		cache_Data_List_NonEmpty_singleton = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Call_Data_List_NonEmpty_singleton(x_0_box)))}
})
	})
	return cache_Data_List_NonEmpty_singleton
}

var cache_Data_List_NonEmpty_snoc_prime_ gopurs_runtime.Value
var once_Data_List_NonEmpty_snoc_prime_ sync.Once
func Get_Data_List_NonEmpty_snoc_prime_() gopurs_runtime.Value {
	once_Data_List_NonEmpty_snoc_prime_.Do(func() {
		cache_Data_List_NonEmpty_snoc_prime_ = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Call_Data_List_NonEmpty_snoc_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0_box), v1_1_box)))}
})
	})
	return cache_Data_List_NonEmpty_snoc_prime_
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
return Call_Data_List_NonEmpty_nubEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
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
return Call_Data_List_NonEmpty_nub(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_List_NonEmpty_nub
}

var cache_Data_List_NonEmpty_modifyAt gopurs_runtime.Value
var once_Data_List_NonEmpty_modifyAt sync.Once
func Get_Data_List_NonEmpty_modifyAt() gopurs_runtime.Value {
	once_Data_List_NonEmpty_modifyAt.Do(func() {
		cache_Data_List_NonEmpty_modifyAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_NonEmpty_modifyAt(i_0_box.IntVal, f_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_NonEmpty_modifyAt
}

var cache_Data_List_NonEmpty_lift gopurs_runtime.Value
var once_Data_List_NonEmpty_lift sync.Once
func Get_Data_List_NonEmpty_lift() gopurs_runtime.Value {
	once_Data_List_NonEmpty_lift.Do(func() {
		cache_Data_List_NonEmpty_lift = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_lift(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box))
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
return func() gopurs_runtime.Value {
				orig := Call_Data_List_NonEmpty_partition(x_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box))
				_ = orig
				return gopurs_runtime.RecordDict([]string{"no", "yes"}, []gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.no)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.yes)}})
				}()
})
	})
	return cache_Data_List_NonEmpty_partition
}

var cache_Data_List_NonEmpty_span gopurs_runtime.Value
var once_Data_List_NonEmpty_span sync.Once
func Get_Data_List_NonEmpty_span() gopurs_runtime.Value {
	once_Data_List_NonEmpty_span.Do(func() {
		cache_Data_List_NonEmpty_span = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := Call_Data_List_NonEmpty_span(x_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box))
				_ = orig
				return gopurs_runtime.RecordDict([]string{"init", "rest"}, []gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.go__init)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.rest)}})
				}()
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
return gopurs_runtime.Int(Call_Data_List_NonEmpty_length(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box)))
})
	})
	return cache_Data_List_NonEmpty_length
}

var cache_Data_List_NonEmpty_last gopurs_runtime.Value
var once_Data_List_NonEmpty_last sync.Once
func Get_Data_List_NonEmpty_last() gopurs_runtime.Value {
	once_Data_List_NonEmpty_last.Do(func() {
		cache_Data_List_NonEmpty_last = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_last(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box))
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
return Call_Data_List_NonEmpty_intersect(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_Data_List_NonEmpty_intersect
}

var cache_Data_List_NonEmpty_insertAt gopurs_runtime.Value
var once_Data_List_NonEmpty_insertAt sync.Once
func Get_Data_List_NonEmpty_insertAt() gopurs_runtime.Value {
	once_Data_List_NonEmpty_insertAt.Do(func() {
		cache_Data_List_NonEmpty_insertAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_NonEmpty_insertAt(i_0_box.IntVal, a_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_NonEmpty_insertAt
}

var cache_Data_List_NonEmpty_go__init gopurs_runtime.Value
var once_Data_List_NonEmpty_go__init sync.Once
func Get_Data_List_NonEmpty_go__init() gopurs_runtime.Value {
	once_Data_List_NonEmpty_go__init.Do(func() {
		cache_Data_List_NonEmpty_go__init = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_go__init(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_Data_List_NonEmpty_go__init
}

var cache_Data_List_NonEmpty_index gopurs_runtime.Value
var once_Data_List_NonEmpty_index sync.Once
func Get_Data_List_NonEmpty_index() gopurs_runtime.Value {
	once_Data_List_NonEmpty_index.Do(func() {
		cache_Data_List_NonEmpty_index = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, i_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_NonEmpty_index(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box), i_1_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_NonEmpty_index
}

var cache_Data_List_NonEmpty_head gopurs_runtime.Value
var once_Data_List_NonEmpty_head sync.Once
func Get_Data_List_NonEmpty_head() gopurs_runtime.Value {
	once_Data_List_NonEmpty_head.Do(func() {
		cache_Data_List_NonEmpty_head = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_head(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box))
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
return Call_Data_List_NonEmpty_groupAll(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_List_NonEmpty_groupAll
}

var cache_Data_List_NonEmpty_group gopurs_runtime.Value
var once_Data_List_NonEmpty_group sync.Once
func Get_Data_List_NonEmpty_group() gopurs_runtime.Value {
	once_Data_List_NonEmpty_group.Do(func() {
		cache_Data_List_NonEmpty_group = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_group(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_Data_List_NonEmpty_group
}

var cache_Data_List_NonEmpty_fromList gopurs_runtime.Value
var once_Data_List_NonEmpty_fromList sync.Once
func Get_Data_List_NonEmpty_fromList() gopurs_runtime.Value {
	once_Data_List_NonEmpty_fromList.Do(func() {
		cache_Data_List_NonEmpty_fromList = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_NonEmpty_fromList(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_NonEmpty_fromList
}

var cache_Data_List_NonEmpty_fromList__2835274500 gopurs_runtime.Value
var once_Data_List_NonEmpty_fromList__2835274500 sync.Once
func Get_Data_List_NonEmpty_fromList__2835274500() gopurs_runtime.Value {
	once_Data_List_NonEmpty_fromList__2835274500.Do(func() {
		cache_Data_List_NonEmpty_fromList__2835274500 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_NonEmpty_fromList__2835274500(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_NonEmpty_fromList__2835274500
}

var cache_Data_List_NonEmpty_fromFoldable gopurs_runtime.Value
var once_Data_List_NonEmpty_fromFoldable sync.Once
func Get_Data_List_NonEmpty_fromFoldable() gopurs_runtime.Value {
	once_Data_List_NonEmpty_fromFoldable.Do(func() {
		cache_Data_List_NonEmpty_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_NonEmpty_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_NonEmpty_fromFoldable
}

var cache_Data_List_NonEmpty_foldM gopurs_runtime.Value
var once_Data_List_NonEmpty_foldM sync.Once
func Get_Data_List_NonEmpty_foldM() gopurs_runtime.Value {
	once_Data_List_NonEmpty_foldM.Do(func() {
		cache_Data_List_NonEmpty_foldM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_foldM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_Data_List_NonEmpty_foldM
}

var cache_Data_List_NonEmpty_findLastIndex gopurs_runtime.Value
var once_Data_List_NonEmpty_findLastIndex sync.Once
func Get_Data_List_NonEmpty_findLastIndex() gopurs_runtime.Value {
	once_Data_List_NonEmpty_findLastIndex.Do(func() {
		cache_Data_List_NonEmpty_findLastIndex = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_NonEmpty_findLastIndex(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_NonEmpty_findLastIndex
}

var cache_Data_List_NonEmpty_findIndex gopurs_runtime.Value
var once_Data_List_NonEmpty_findIndex sync.Once
func Get_Data_List_NonEmpty_findIndex() gopurs_runtime.Value {
	once_Data_List_NonEmpty_findIndex.Do(func() {
		cache_Data_List_NonEmpty_findIndex = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_NonEmpty_findIndex(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_NonEmpty_findIndex
}

var cache_Data_List_NonEmpty_filterM gopurs_runtime.Value
var once_Data_List_NonEmpty_filterM sync.Once
func Get_Data_List_NonEmpty_filterM() gopurs_runtime.Value {
	once_Data_List_NonEmpty_filterM.Do(func() {
		cache_Data_List_NonEmpty_filterM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_NonEmpty_filterM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
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
return func() gopurs_runtime.Value {
				_v := Call_Data_List_NonEmpty_elemLastIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_NonEmpty_elemLastIndex
}

var cache_Data_List_NonEmpty_elemIndex gopurs_runtime.Value
var once_Data_List_NonEmpty_elemIndex sync.Once
func Get_Data_List_NonEmpty_elemIndex() gopurs_runtime.Value {
	once_Data_List_NonEmpty_elemIndex.Do(func() {
		cache_Data_List_NonEmpty_elemIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_NonEmpty_elemIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
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
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_drop(x_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box)))}
})
	})
	return cache_Data_List_NonEmpty_drop
}

var cache_Data_List_NonEmpty_cons_prime_ gopurs_runtime.Value
var once_Data_List_NonEmpty_cons_prime_ sync.Once
func Get_Data_List_NonEmpty_cons_prime_() gopurs_runtime.Value {
	once_Data_List_NonEmpty_cons_prime_.Do(func() {
		cache_Data_List_NonEmpty_cons_prime_ = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Call_Data_List_NonEmpty_cons_prime_(x_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1_box))))}
})
	})
	return cache_Data_List_NonEmpty_cons_prime_
}

var cache_Data_List_NonEmpty_cons gopurs_runtime.Value
var once_Data_List_NonEmpty_cons sync.Once
func Get_Data_List_NonEmpty_cons() gopurs_runtime.Value {
	once_Data_List_NonEmpty_cons.Do(func() {
		cache_Data_List_NonEmpty_cons = gopurs_runtime.Func2(func(y_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Call_Data_List_NonEmpty_cons(y_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box))))}
})
	})
	return cache_Data_List_NonEmpty_cons
}

var cache_Data_List_NonEmpty_concatMap gopurs_runtime.Value
var once_Data_List_NonEmpty_concatMap sync.Once
func Get_Data_List_NonEmpty_concatMap() gopurs_runtime.Value {
	once_Data_List_NonEmpty_concatMap.Do(func() {
		cache_Data_List_NonEmpty_concatMap = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Call_Data_List_NonEmpty_concatMap(b_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](a_1_box))))}
})
	})
	return cache_Data_List_NonEmpty_concatMap
}

var cache_Data_List_NonEmpty_concat gopurs_runtime.Value
var once_Data_List_NonEmpty_concat sync.Once
func Get_Data_List_NonEmpty_concat() gopurs_runtime.Value {
	once_Data_List_NonEmpty_concat.Do(func() {
		cache_Data_List_NonEmpty_concat = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Call_Data_List_NonEmpty_concat(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box))))}
})
	})
	return cache_Data_List_NonEmpty_concat
}

var cache_Data_List_NonEmpty_catMaybes gopurs_runtime.Value
var once_Data_List_NonEmpty_catMaybes sync.Once
func Get_Data_List_NonEmpty_catMaybes() gopurs_runtime.Value {
	once_Data_List_NonEmpty_catMaybes.Do(func() {
		cache_Data_List_NonEmpty_catMaybes = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_NonEmpty_catMaybes(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](v_0_box)))}
})
	})
	return cache_Data_List_NonEmpty_catMaybes
}

var cache_Data_List_NonEmpty_appendFoldable gopurs_runtime.Value
var once_Data_List_NonEmpty_appendFoldable sync.Once
func Get_Data_List_NonEmpty_appendFoldable() gopurs_runtime.Value {
	once_Data_List_NonEmpty_appendFoldable.Do(func() {
		cache_Data_List_NonEmpty_appendFoldable = gopurs_runtime.Func3(func(dictFoldable_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Call_Data_List_NonEmpty_appendFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box), ys_2_box)))}
})
	})
	return cache_Data_List_NonEmpty_appendFoldable
}

func Call_Data_List_NonEmpty_zipWith(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value], v1_2_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
var v1_2 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v1_2_loop
_ = v1_2
var Call_local_Data_List_NonEmpty_go__go_3_0_0 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_NonEmpty_go__go_3_0_0
var go__go_3_0_0 gopurs_runtime.Value
_ = go__go_3_0_0
Call_local_Data_List_NonEmpty_go__go_3_0_0 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v2_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_3_0_0:
for {
if false { continue go__go_3_0_0 }
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_5_loop
_ = v1_5
var v2_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v2_6_loop
_ = v2_6
var __t1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v_4 == nil) {
__t1 = v2_6
goto end_branch_1
} else {

}
}
{
if (v1_5 == nil) {
__t1 = v2_6
goto end_branch_1
} else {

}
}
{
if ((v_4 != nil)) && ((v1_5 != nil)) {
v_4_loop = (v_4).V1
v1_5_loop = (v1_5).V1
v2_6_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (v_4).V0, (v1_5).V0), v2_6})
continue go__go_3_0_0
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](func() gopurs_runtime.Value { panic("unreachable") }())
goto end_branch_1
} else {

}
}
{
__t1 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}
go__go_3_0_0 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_NonEmpty_go__go_3_0_0(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_6_loop_val)))}
})
})
})
var Call_local_Data_List_NonEmpty_go__go_4_2_1 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_NonEmpty_go__go_4_2_1
var go__go_4_2_1 gopurs_runtime.Value
_ = go__go_4_2_1
Call_local_Data_List_NonEmpty_go__go_4_2_1 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_4_2_1:
for {
if false { continue go__go_4_2_1 }
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_6 == nil) {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if (v1_6 != nil) {
v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
v1_6_loop = (v1_6).V1
continue go__go_4_2_1
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}
go__go_4_2_1 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_NonEmpty_go__go_4_2_1(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
})
})
return (&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (v_1).V0, (v1_2).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_NonEmpty_go__go_4_2_1((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), Call_local_Data_List_NonEmpty_go__go_3_0_0(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_1).V1), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v1_2).V1), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))))}})
}

func Call_Data_List_NonEmpty_zipWithA(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Apply0_1_0 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m)])
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}))
_ = Apply0_1_0
return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value, ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_traversable1NonEmptyList()).V3), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Apply0_1_0)}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Call_Data_List_NonEmpty_zipWith(f_2, Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xs_3)), Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](ys_4)))))})
})
}

func Call_Data_List_NonEmpty_wrappedOperation2(name_0_loop string, f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value], v1_3_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var name_0 string = name_0_loop
_ = name_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v1_3_loop
_ = v1_3
// TAST (Let): v2_4_0 shape=App(Other) bindingType=(ADT ["Data","List","Types","List"] [(TypeVar c)])
v2_4_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(f_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_2).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_2).V1)}))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_3).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v1_3).V1)}))}))
_ = v2_4_0
var __t1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]
{
if (v2_4_0 != nil) {
__t1 = (&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, (v2_4_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v2_4_0).V1)}})
goto end_branch_1
} else {

}
}
{
if (v2_4_0 == nil) {
__t1 = Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.Str(("Impossible: empty list in NonEmptyList ") + (name_0)))))
goto end_branch_1
} else {

}
}
{
__t1 = func() *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}

func Call_Data_List_NonEmpty_wrappedOperation(name_0_loop string, f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var name_0 string = name_0_loop
_ = name_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_2_loop
_ = v_2
// TAST (Let): v1_3_0 shape=App(Other) bindingType=(ADT ["Data","List","Types","List"] [(TypeVar b)])
v1_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_2).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_2).V1)}))}))
_ = v1_3_0
var __t1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]
{
if (v1_3_0 != nil) {
__t1 = (&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, (v1_3_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v1_3_0).V1)}})
goto end_branch_1
} else {

}
}
{
if (v1_3_0 == nil) {
__t1 = Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.Str(("Impossible: empty list in NonEmptyList ") + (name_0)))))
goto end_branch_1
} else {

}
}
{
__t1 = func() *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}

func Call_Data_List_NonEmpty_updateAt(i_0_loop int64, a_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var i_0 int64 = i_0_loop
_ = i_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var v_2 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_2_loop
_ = v_2
var __t2 gopurs_runtime.Value
{
if (i_0) == (int64(0)) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_470062885_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_1, (v_2).V1}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
goto end_branch_2
} else {

}
}
{
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","List","Types","List"] [(TypeVar a)])])
__local_var_3_0 := Rebox_Data_List_NonEmpty_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply3(Get_Data_List_updateAt(), gopurs_runtime.Int((i_0) - (int64(1))), a_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_2).V1))})))
_ = __local_var_3_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_3_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952((&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, (v_2).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((__local_var_3_0).V0)}})))}, true}
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
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}
end_branch_2:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_470062885_3094389156(Rebox_Data_List_NonEmpty_3094389156_470062885(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t2))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_List_NonEmpty_unzip(ts_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
var ts_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = ts_0_loop
_ = ts_0
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_1673906728_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, ((ts_0).V0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_List_Types_listMap(), Get_Data_Tuple_fst(), (ts_0).V1)))}}))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, ((ts_0).V0).V1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_List_Types_listMap(), Get_Data_Tuple_snd(), (ts_0).V1)))}}))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
				_p := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{V0: _p.V0, V1: _p.V1}
			}()
}

func Call_Data_List_NonEmpty_unsnoc(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) struct{
	go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	last gopurs_runtime.Value
} {
var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
// TAST (Let): v1_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [init: (ADT ["Data","List","Types","List"] [(TypeVar a)]), last: (TypeVar a)] Any))])
v1_1_0 := Rebox_Data_List_NonEmpty_3094389156_3371309921(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_unsnoc(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_0).V1))})))
_ = v1_1_0
var __t1 struct{
	go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	last gopurs_runtime.Value
}
{
if (v1_1_0 == nil) {
__t1 = struct{
	go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	last gopurs_runtime.Value
}{(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), (v_0).V0}
goto end_branch_1
} else {

}
}
{
if (v1_1_0 != nil) {
__t1 = struct{
	go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	last gopurs_runtime.Value
}{(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_0).V0, (v1_1_0).V0.go__init}), (v1_1_0).V0.last}
goto end_branch_1
} else {

}
}
{
__t1 = func() struct{
	go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	last gopurs_runtime.Value
} { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}

func Call_Data_List_NonEmpty_unionBy(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation2(), gopurs_runtime.Str("unionBy"), gopurs_runtime.Apply(Get_Data_List_unionBy(), x_0))
}

func Call_Data_List_NonEmpty_union(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation2(), gopurs_runtime.Str("union"), gopurs_runtime.Apply(Get_Data_List_union(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_0)}))
}

func Call_Data_List_NonEmpty_uncons(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) struct{
	head gopurs_runtime.Value
	tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
} {
var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
return struct{
	head gopurs_runtime.Value
	tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}{(v_0).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_0).V1)}
}

func Call_Data_List_NonEmpty_toList(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
return (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_0).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_0).V1)})
}

func Call_Data_List_NonEmpty_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=Any
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable_0.V1), gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1)
if (__t_tag_1 == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
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
var __t_tag_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1)
if (__t_tag_2 != nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_2351501316_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(xs_1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(xs_1.UnsafePtr).V1)}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}, true}
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
__t3 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
}))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1)}))})
})
}

func Call_Data_List_NonEmpty_tail(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_0).V1)
}

func Call_Data_List_NonEmpty_sortBy(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation(), gopurs_runtime.Str("sortBy"), gopurs_runtime.Apply(Get_Data_List_sortBy(), x_0))
}

func Call_Data_List_NonEmpty_sort(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 shape=Other bindingType=(Func [(TypeVar a), (TypeVar a)] (ADT ["Data","Ordering","Ordering"] []))
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(Call_Data_List_NonEmpty_wrappedOperation("sortBy", gopurs_runtime.Apply(Get_Data_List_sortBy(), compare_1_0), Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](xs_2)))))}
})
}

func Call_Data_List_NonEmpty_snoc(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value], y_1_loop gopurs_runtime.Value) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
var y_1 gopurs_runtime.Value = y_1_loop
_ = y_1
return (&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, (v_0).V0, gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_foldableList()).V2), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, y_1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)}))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_0).V1))})})
}

func Call_Data_List_NonEmpty_singleton(x_0_loop gopurs_runtime.Value) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return Rebox_Data_List_NonEmpty_1293498952_3123684004((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}}))
}

func Call_Data_List_NonEmpty_snoc_prime_(v_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_1_loop gopurs_runtime.Value) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var v_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v_0 != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952((&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, (v_0).V0, gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_foldableList()).V2), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, v1_1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)}))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)})})))}
goto end_branch_0
} else {

}
}
{
if (v_0 == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, v1_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(func() *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] { panic("Failed pattern match") }()))}
}
end_branch_0:
return Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](__t0))
}

func Call_Data_List_NonEmpty_nubEq(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
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

func Call_Data_List_NonEmpty_nub(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation(), gopurs_runtime.Str("nub"), gopurs_runtime.Apply(Get_Data_List_nubBy(), gopurs_runtime.Box(dictOrd_0.V1)))
}

func Call_Data_List_NonEmpty_modifyAt(i_0_loop int64, f_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var i_0 int64 = i_0_loop
_ = i_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_2_loop
_ = v_2
var __t2 gopurs_runtime.Value
{
if (i_0) == (int64(0)) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_470062885_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, (v_2).V0), (v_2).V1}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
goto end_branch_2
} else {

}
}
{
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","List","Types","List"] [(TypeVar a)])])
__local_var_3_0 := Rebox_Data_List_NonEmpty_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply3(Get_Data_List_alterAt(), gopurs_runtime.Int((i_0) - (int64(1))), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, x_3)}))}
}), (v_2).V1)))
_ = __local_var_3_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_3_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952((&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, (v_2).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((__local_var_3_0).V0)}})))}, true}
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
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}
end_branch_2:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_470062885_3094389156(Rebox_Data_List_NonEmpty_3094389156_470062885(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t2))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_List_NonEmpty_lift(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_1).V1)}))})
}

func Call_Data_List_NonEmpty_mapMaybe(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var Call_local_Data_List_NonEmpty_go__go_1_0_2 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_NonEmpty_go__go_1_0_2
var go__go_1_0_2 gopurs_runtime.Value
_ = go__go_1_0_2
Call_local_Data_List_NonEmpty_go__go_1_0_2 = func(v_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_1_0_2:
for {
if false { continue go__go_1_0_2 }
var v_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_3 == nil) {
var Call_local_Data_List_NonEmpty_go__go_4_1_3 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_NonEmpty_go__go_4_1_3
var go__go_4_1_3 gopurs_runtime.Value
_ = go__go_4_1_3
Call_local_Data_List_NonEmpty_go__go_4_1_3 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_4_1_3:
for {
if false { continue go__go_4_1_3 }
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var __t2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_6 == nil) {
__t2 = v_5
goto end_branch_2
} else {

}
}
{
if (v1_6 != nil) {
v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
v1_6_loop = (v1_6).V1
continue go__go_4_1_3
__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}
go__go_4_1_3 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_NonEmpty_go__go_4_1_3(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
})
})
__t5 = Call_local_Data_List_NonEmpty_go__go_4_1_3((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v_2)
goto end_branch_5
} else {

}
}
{
if (v1_3 != nil) {
// TAST (Let): v2_4_3 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
v2_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(x_0, (v1_3).V0))
_ = v2_4_3
var __t4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v2_4_3 == nil) {
v_2_loop = v_2
v1_3_loop = (v1_3).V1
continue go__go_1_0_2
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_4
} else {

}
}
{
if (v2_4_3 != nil) {
v_2_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v2_4_3).V0, v_2})
v1_3_loop = (v1_3).V1
continue go__go_1_0_2
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}
go__go_1_0_2 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_NonEmpty_go__go_1_0_2(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val)))}
})
})
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_NonEmpty_go__go_1_0_2((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)})))}
})
}

func Call_Data_List_NonEmpty_partition(x_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) struct{
	no *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	yes *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
} {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
return func() struct{
	no *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	yes *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
} {
					orig := func() gopurs_runtime.Value {
				orig := func() struct{
	no *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	yes *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
} {
					orig := gopurs_runtime.Apply2(Get_Data_List_partition(), x_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_1).V1)}))})
					_ = orig
					clone := struct{
	no *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	yes *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}{}
					clone.no = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "no"))
					clone.yes = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "yes"))
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict([]string{"no", "yes"}, []gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.no)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.yes)}})
				}()
					_ = orig
					clone := struct{
	no *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	yes *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}{}
					clone.no = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "no"))
					clone.yes = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "yes"))
					return clone
				}()
}

func Call_Data_List_NonEmpty_span(x_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) struct{
	go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	rest *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
} {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
return func() struct{
	go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	rest *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
} {
					orig := func() gopurs_runtime.Value {
				orig := func() struct{
	go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	rest *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
} {
					orig := gopurs_runtime.Apply2(Get_Data_List_span(), x_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_1).V1)}))})
					_ = orig
					clone := struct{
	go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	rest *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}{}
					clone.go__init = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "init"))
					clone.rest = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "rest"))
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict([]string{"init", "rest"}, []gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.go__init)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.rest)}})
				}()
					_ = orig
					clone := struct{
	go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	rest *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}{}
					clone.go__init = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "init"))
					clone.rest = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "rest"))
					return clone
				}()
}

func Call_Data_List_NonEmpty_take(x_0_loop int64) gopurs_runtime.Value {
var x_0 int64 = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [(ADT ["Data","List","Types","List"] [(TypeVar a)])] (ADT ["Data","List","Types","List"] [(TypeVar a)]))
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_List_take(), gopurs_runtime.Int(x_0))
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)}))})
})
}

func Call_Data_List_NonEmpty_takeWhile(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var Call_local_Data_List_NonEmpty_go__go_1_0_4 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_NonEmpty_go__go_1_0_4
var go__go_1_0_4 gopurs_runtime.Value
_ = go__go_1_0_4
Call_local_Data_List_NonEmpty_go__go_1_0_4 = func(v_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_1_0_4:
for {
if false { continue go__go_1_0_4 }
var v_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if ((v1_3 != nil)) && ((gopurs_runtime.Apply(x_0, (v1_3).V0).IntVal) != (0)) {
v_2_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_3).V0, v_2})
v1_3_loop = (v1_3).V1
continue go__go_1_0_4
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_3
} else {

}
}
{
var Call_local_Data_List_NonEmpty_go__go_4_1_5 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_NonEmpty_go__go_4_1_5
var go__go_4_1_5 gopurs_runtime.Value
_ = go__go_4_1_5
Call_local_Data_List_NonEmpty_go__go_4_1_5 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_4_1_5:
for {
if false { continue go__go_4_1_5 }
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var __t2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_6 == nil) {
__t2 = v_5
goto end_branch_2
} else {

}
}
{
if (v1_6 != nil) {
v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
v1_6_loop = (v1_6).V1
continue go__go_4_1_5
__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}
go__go_4_1_5 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_NonEmpty_go__go_4_1_5(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
})
})
__t3 = Call_local_Data_List_NonEmpty_go__go_4_1_5((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v_2)
}
end_branch_3:
return __t3
}
}
go__go_1_0_4 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_NonEmpty_go__go_1_0_4(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val)))}
})
})
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_NonEmpty_go__go_1_0_4((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)})))}
})
}

func Call_Data_List_NonEmpty_length(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) int64 {
var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
var Call_local_Data_List_NonEmpty_go__go_1_0_6 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_NonEmpty_go__go_1_0_6
var go__go_1_0_6 gopurs_runtime.Value
_ = go__go_1_0_6
Call_local_Data_List_NonEmpty_go__go_1_0_6 = func(b_2_loop gopurs_runtime.Value, v_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_1_0_6:
for {
if false { continue go__go_1_0_6 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var v_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (v_3 == nil) {
__t1 = b_2
goto end_branch_1
} else {

}
}
{
if (v_3 != nil) {
b_2_loop = gopurs_runtime.Int((b_2.IntVal) + (int64(1)))
v_3_loop = (v_3).V1
continue go__go_1_0_6
__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_1_0_6 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_NonEmpty_go__go_1_0_6(b_2_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_3_loop_val))
})
})
return (int64(1)) + (Call_local_Data_List_NonEmpty_go__go_1_0_6(gopurs_runtime.Int(int64(0)), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_0).V1)).IntVal)
}

func Call_Data_List_NonEmpty_last(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t6 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = (v_0).V1
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 1358893437 && __t_tag_0.UnsafePtr != nil) {
var __t5 gopurs_runtime.Value
{
var __t_tag_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((v_0).V1.UnsafePtr).V1
if (__t_tag_4 == nil) {
__t5 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((v_0).V1.UnsafePtr).V0
goto end_branch_5
} else {

}
}
{
var __t3 gopurs_runtime.Value
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_last(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((v_0).V1.UnsafePtr).V1)}))
if (__t_tag_1 == nil) {
__t3 = (v_0).V0
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_last(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((v_0).V1.UnsafePtr).V1)}))
if (__t_tag_2 != nil) {
__t3 = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Apply(Get_Data_List_last(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((v_0).V1.UnsafePtr).V1)}).UnsafePtr).V0
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t5 = __t3
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
__t6 = (v_0).V0
}
end_branch_6:
return __t6
}

func Call_Data_List_NonEmpty_intersectBy(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation2(), gopurs_runtime.Str("intersectBy"), gopurs_runtime.Apply(Get_Data_List_intersectBy(), x_0))
}

func Call_Data_List_NonEmpty_intersect(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation2(), gopurs_runtime.Str("intersect"), gopurs_runtime.Apply(Get_Data_List_intersect(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_0)}))
}

func Call_Data_List_NonEmpty_insertAt(i_0_loop int64, a_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var i_0 int64 = i_0_loop
_ = i_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var v_2 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_2_loop
_ = v_2
var __t2 gopurs_runtime.Value
{
if (i_0) == (int64(0)) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_470062885_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_2).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_2).V1)}))}}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
goto end_branch_2
} else {

}
}
{
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","List","Types","List"] [(TypeVar a)])])
__local_var_3_0 := Rebox_Data_List_NonEmpty_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply3(Get_Data_List_insertAt(), gopurs_runtime.Int((i_0) - (int64(1))), a_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_2).V1))})))
_ = __local_var_3_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_3_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952((&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, (v_2).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((__local_var_3_0).V0)}})))}, true}
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
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}
end_branch_2:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_470062885_3094389156(Rebox_Data_List_NonEmpty_3094389156_470062885(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t2))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_List_NonEmpty_go__init(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
// TAST (Let): __local_var_1_1 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [init: (ADT ["Data","List","Types","List"] [(TypeVar a)]), last: (TypeVar a)] Any))])
__local_var_1_1 := Rebox_Data_List_NonEmpty_3094389156_3371309921(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_unsnoc(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_0).V1))})))
_ = __local_var_1_1
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_1_1 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((__local_var_1_1).V0.go__init)}, true}
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
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_2:
// TAST (Let): __local_var_1_0 shape=Let(Branch(Other, def=Other)) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","List","Types","List"] [(TypeVar a)])])
var __local_var_1_0 *Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)})
var __t3 gopurs_runtime.Value
{
if (__local_var_1_0 == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (__local_var_1_0 != nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_0).V0, (__local_var_1_0).V0}))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__t3)
}

func Call_Data_List_NonEmpty_index(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value], i_1_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
var i_1 int64 = i_1_loop
_ = i_1
var __t0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (i_1) == (int64(0)) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(v_0).V0, true}
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
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_List_index(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_0).V1))}, gopurs_runtime.Int((i_1) - (int64(1)))))
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_List_NonEmpty_head(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
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

func Call_Data_List_NonEmpty_groupAll(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation(), gopurs_runtime.Str("groupAll"), gopurs_runtime.Apply(Get_Data_List_groupAll(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}))
}

func Call_Data_List_NonEmpty_group(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply2(Get_Data_List_NonEmpty_wrappedOperation(), gopurs_runtime.Str("group"), gopurs_runtime.Apply(Get_Data_List_group(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_0)}))
}

func Call_Data_List_NonEmpty_fromList(v_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var v_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]
{
if (v_0 == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](func() gopurs_runtime.Value {
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
if (v_0 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)}}))}, true}
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
__t0 = func() *Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] { panic("Failed pattern match") }()
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_470062885_3094389156(__t0))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_List_NonEmpty_fromList__2835274500(v_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
fromList__2835274500:
for {
if false { continue fromList__2835274500 }
var v_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]
{
if (v_0 == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](func() gopurs_runtime.Value {
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
if (v_0 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)}}))}, true}
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
__t0 = func() *Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] { panic("Failed pattern match") }()
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_470062885_3094389156(__t0))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_List_NonEmpty_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(TypeVar a)])] (ADT ["Data","List","Types","List"] [(TypeVar a)]))
__local_var_1_0 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V2), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_1_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 shape=App(Other) bindingType=Any
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
_ = __local_var_3_1
var __t2 *Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]
{
if (__local_var_3_1.Type == 9 && __local_var_3_1.IntVal == 1358893437 && __local_var_3_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](func() gopurs_runtime.Value {
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
if (__local_var_3_1.Type == 9 && __local_var_3_1.IntVal == 1358893437 && __local_var_3_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(__local_var_3_1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(__local_var_3_1.UnsafePtr).V1)}}))}, true}
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
__t2 = func() *Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_470062885_3094389156(__t2))}
})
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_List_NonEmpty_foldM(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_1 shape=Other bindingType=Any
__local_var_5_1 := (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1
_ = __local_var_5_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply2(f_2, b_3, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Func(func(b_prime__6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(Get_Data_List_foldM(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)}, f_2, b_prime__6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__local_var_5_1))})
}))
})
}

func Call_Data_List_NonEmpty_findLastIndex(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
// TAST (Let): v1_2_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
v1_2_0 := Rebox_Data_List_NonEmpty_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_List_findLastIndex(), f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_1).V1))})))
_ = v1_2_0
var __t2 gopurs_runtime.Value
{
if (v1_2_0 != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(((v1_2_0).V0) + (int64(1)))}))}
goto end_branch_2
} else {

}
}
{
if (v1_2_0 == nil) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_0, (v_1).V0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(int64(0))}))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_1170268447_3094389156(func() *Constructor_Data_Maybe_Just[int64] { panic("Failed pattern match") }()))}
}
end_branch_2:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_1170268447_3094389156(Rebox_Data_List_NonEmpty_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t2))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_List_NonEmpty_findIndex(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
var __t10 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply(f_0, (v_1).V0).IntVal) != (0) {
__t10 = (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(int64(0))})
goto end_branch_10
} else {

}
}
{
var Call_local_Data_List_NonEmpty_go__816245670_2_1_7 func(int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64]
_ = Call_local_Data_List_NonEmpty_go__816245670_2_1_7
var go__816245670_2_1_7 gopurs_runtime.Value
_ = go__816245670_2_1_7
Call_local_Data_List_NonEmpty_go__816245670_2_1_7 = func(v_3_loop int64, v1_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
go__816245670_2_1_7:
for {
if false { continue go__816245670_2_1_7 }
var v_3 int64 = v_3_loop
_ = v_3
var v1_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_4_loop
_ = v1_4
var __t3 *Constructor_Data_Maybe_Just[int64]
{
if (v1_4 != nil) {
var __t2 *Constructor_Data_Maybe_Just[int64]
{
if (gopurs_runtime.Apply(f_0, (v1_4).V0).IntVal) != (0) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int(v_3), true}
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
v_3_loop = (v_3) + (int64(1))
v1_4_loop = (v1_4).V1
continue go__816245670_2_1_7
__t2 = func() *Constructor_Data_Maybe_Just[int64] { panic("unreachable") }()
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
if (v1_4 == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
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
__t3 = func() *Constructor_Data_Maybe_Just[int64] { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}
go__816245670_2_1_7 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_1170268447_3094389156(Call_local_Data_List_NonEmpty_go__816245670_2_1_7(v_3_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_4_loop_val))))}
})
})
var go__go_3_4_8 gopurs_runtime.Value
_ = go__go_3_4_8
// FALLBACK TCO: isLoop=false len=1
go__go_3_4_8 = gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
var __t_tag_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5)
if (__t_tag_5 != nil) {
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0).IntVal) != (0) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(v_4.IntVal)}))}
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_1170268447_3094389156(Call_local_Data_List_NonEmpty_go__816245670_2_1_7((v_4.IntVal) + (int64(1)), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)))}
}
end_branch_6:
__t8 = __t6
goto end_branch_8
} else {

}
}
{
var __t_tag_7 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5)
if (__t_tag_7 == nil) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_1170268447_3094389156(func() *Constructor_Data_Maybe_Just[int64] { panic("Failed pattern match") }()))}
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_1170268447_3094389156(Rebox_Data_List_NonEmpty_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t8))))}
})
// TAST (Let): __local_var_2_0 shape=LetRec(LetRec(App(Other))) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
__local_var_2_0 := Call_local_Data_List_NonEmpty_go__816245670_2_1_7(int64(0), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_1).V1))
_ = __local_var_2_0
var __t9 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_2_0 != nil) {
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int(((__local_var_2_0).V0) + (int64(1))), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_9:
__t10 = __t9
}
end_branch_10:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t10)}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_List_NonEmpty_filterM(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [Boolean])), (ADT ["Data","List","Types","List"] [(TypeVar a)])] (TypeApp (TypeVar m) [(ADT ["Data","List","Types","List"] [(TypeVar a)])]))
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_List_filterM(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 shape=App(Other) bindingType=Any
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)}))})
})
})
}

func Call_Data_List_NonEmpty_filter(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var Call_local_Data_List_NonEmpty_go__go_1_0_9 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_NonEmpty_go__go_1_0_9
var go__go_1_0_9 gopurs_runtime.Value
_ = go__go_1_0_9
Call_local_Data_List_NonEmpty_go__go_1_0_9 = func(v_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_1_0_9:
for {
if false { continue go__go_1_0_9 }
var v_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_3 == nil) {
var Call_local_Data_List_NonEmpty_go__go_4_1_10 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_NonEmpty_go__go_4_1_10
var go__go_4_1_10 gopurs_runtime.Value
_ = go__go_4_1_10
Call_local_Data_List_NonEmpty_go__go_4_1_10 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_4_1_10:
for {
if false { continue go__go_4_1_10 }
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var __t2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_6 == nil) {
__t2 = v_5
goto end_branch_2
} else {

}
}
{
if (v1_6 != nil) {
v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
v1_6_loop = (v1_6).V1
continue go__go_4_1_10
__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}
go__go_4_1_10 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_NonEmpty_go__go_4_1_10(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
})
})
__t4 = Call_local_Data_List_NonEmpty_go__go_4_1_10((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v_2)
goto end_branch_4
} else {

}
}
{
if (v1_3 != nil) {
var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply(x_0, (v1_3).V0).IntVal) != (0) {
v_2_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_3).V0, v_2})
v1_3_loop = (v1_3).V1
continue go__go_1_0_9
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_3
} else {

}
}
{
v_2_loop = v_2
v1_3_loop = (v1_3).V1
continue go__go_1_0_9
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}
go__go_1_0_9 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_NonEmpty_go__go_1_0_9(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val)))}
})
})
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_NonEmpty_go__go_1_0_9((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)})))}
})
}

func Call_Data_List_NonEmpty_elemLastIndex(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply(Get_Data_List_NonEmpty_findLastIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), v_2, x_1).IntVal) != (0))
}))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_List_NonEmpty_elemIndex(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply(Get_Data_List_NonEmpty_findIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), v_2, x_1).IntVal) != (0))
}))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_List_NonEmpty_dropWhile(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var Call_local_Data_List_NonEmpty_go__go_1_0_11 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_NonEmpty_go__go_1_0_11
var go__go_1_0_11 gopurs_runtime.Value
_ = go__go_1_0_11
Call_local_Data_List_NonEmpty_go__go_1_0_11 = func(v_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_1_0_11:
for {
if false { continue go__go_1_0_11 }
var v_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var __t1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if ((v_2 != nil)) && ((gopurs_runtime.Apply(x_0, (v_2).V0).IntVal) != (0)) {
v_2_loop = (v_2).V1
continue go__go_1_0_11
__t1 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
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
}
go__go_1_0_11 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_NonEmpty_go__go_1_0_11(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val)))}
})
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_NonEmpty_go__go_1_0_11((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)})))}
})
}

func Call_Data_List_NonEmpty_drop(x_0_loop int64, v_1_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
var x_0 int64 = x_0_loop
_ = x_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_List_drop(), gopurs_runtime.Int(x_0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_1).V1)}))}))
}

func Call_Data_List_NonEmpty_cons_prime_(x_0_loop gopurs_runtime.Value, xs_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = xs_1_loop
_ = xs_1
return (&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, x_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}})
}

func Call_Data_List_NonEmpty_cons(y_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var y_0 gopurs_runtime.Value = y_0_loop
_ = y_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
return (&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, y_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_1).V1)}))}})
}

func Call_Data_List_NonEmpty_concatMap(b_0_loop gopurs_runtime.Value, a_1_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = a_1_loop
_ = a_1
return Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_bindNonEmptyList()).V1), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(a_1))}, b_0)))
}

func Call_Data_List_NonEmpty_concat(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
return Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_bindNonEmptyList()).V1), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(v_0))}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))))
}

func Call_Data_List_NonEmpty_catMaybes(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_Maybe_Just[gopurs_runtime.Value]] = v_0_loop
_ = v_0
var Call_local_Data_List_NonEmpty_go__go_1_0_12 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_NonEmpty_go__go_1_0_12
var go__go_1_0_12 gopurs_runtime.Value
_ = go__go_1_0_12
Call_local_Data_List_NonEmpty_go__go_1_0_12 = func(v_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_1_0_12:
for {
if false { continue go__go_1_0_12 }
var v_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_3 == nil) {
var Call_local_Data_List_NonEmpty_go__go_4_1_13 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_NonEmpty_go__go_4_1_13
var go__go_4_1_13 gopurs_runtime.Value
_ = go__go_4_1_13
Call_local_Data_List_NonEmpty_go__go_4_1_13 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_4_1_13:
for {
if false { continue go__go_4_1_13 }
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var __t2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_6 == nil) {
__t2 = v_5
goto end_branch_2
} else {

}
}
{
if (v1_6 != nil) {
v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
v1_6_loop = (v1_6).V1
continue go__go_4_1_13
__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}
go__go_4_1_13 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_NonEmpty_go__go_4_1_13(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
})
})
__t6 = Call_local_Data_List_NonEmpty_go__go_4_1_13((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v_2)
goto end_branch_6
} else {

}
}
{
if (v1_3 != nil) {
var __t5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]]((v1_3).V0)
if (__t_tag_3 == nil) {
v_2_loop = v_2
v1_3_loop = (v1_3).V1
continue go__go_1_0_12
__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]]((v1_3).V0)
if (__t_tag_4 != nil) {
v_2_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])((v1_3).V0.UnsafePtr).V0, v_2})
v1_3_loop = (v1_3).V1
continue go__go_1_0_12
__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}
}
go__go_1_0_12 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_NonEmpty_go__go_1_0_12(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val)))}
})
})
return Call_local_Data_List_NonEmpty_go__go_1_0_12((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((v_0).V0)}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_0).V1)}))
}

func Call_Data_List_NonEmpty_appendFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], v_1_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value], ys_2_loop gopurs_runtime.Value) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
var Call_local_Data_List_NonEmpty_go__go_3_0_14 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_NonEmpty_go__go_3_0_14
var go__go_3_0_14 gopurs_runtime.Value
_ = go__go_3_0_14
Call_local_Data_List_NonEmpty_go__go_3_0_14 = func(b_4_loop gopurs_runtime.Value, v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_3_0_14:
for {
if false { continue go__go_3_0_14 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var __t1 gopurs_runtime.Value
{
if (v_5 == nil) {
__t1 = b_4
goto end_branch_1
} else {

}
}
{
if (v_5 != nil) {
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_5).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_4)}))}
v_5_loop = (v_5).V1
continue go__go_3_0_14
__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_3_0_14 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_NonEmpty_go__go_3_0_14(b_4_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val))
})
})
var Call_local_Data_List_NonEmpty_go__go_4_2_15 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_NonEmpty_go__go_4_2_15
var go__go_4_2_15 gopurs_runtime.Value
_ = go__go_4_2_15
Call_local_Data_List_NonEmpty_go__go_4_2_15 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_4_2_15:
for {
if false { continue go__go_4_2_15 }
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_6 == nil) {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if (v1_6 != nil) {
v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
v1_6_loop = (v1_6).V1
continue go__go_4_2_15
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}
go__go_4_2_15 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_NonEmpty_go__go_4_2_15(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
})
})
return (&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, (v_1).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_local_Data_List_NonEmpty_go__go_3_0_14(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable_0.V2), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, ys_2)))}, Call_local_Data_List_NonEmpty_go__go_4_2_15((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_1).V1)))))}})
}

func Rebox_Data_List_NonEmpty_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_List_NonEmpty_1293498952_3123684004(in *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_NonEmpty_1673906728_138441832(in *Constructor_Data_Tuple_Tuple[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(in.V0))}
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(in.V1))}
	return out
}

func Rebox_Data_List_NonEmpty_2351501316_138441832(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}

func Rebox_Data_List_NonEmpty_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[int64]{}
		out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_List_NonEmpty_3094389156_1978018568(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](in.V0)
	return out
}

func Rebox_Data_List_NonEmpty_3094389156_3371309921(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct{
	go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	last gopurs_runtime.Value
}] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[struct{
	go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	last gopurs_runtime.Value
}]{}
		out.V0 = func() struct{
	go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	last gopurs_runtime.Value
} {
					orig := in.V0
					_ = orig
					clone := struct{
	go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	last gopurs_runtime.Value
}{}
					clone.go__init = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "init"))
					clone.last = gopurs_runtime.RecordGet(orig, "last")
					return clone
				}()
	return out
}

func Rebox_Data_List_NonEmpty_3094389156_470062885(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{}
		out.V0 = Rebox_Data_List_NonEmpty_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
	return out
}

func Rebox_Data_List_NonEmpty_3123684004_1293498952(in *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_NonEmpty_470062885_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_NonEmpty_3123684004_1293498952(in.V0))}
	return out
}



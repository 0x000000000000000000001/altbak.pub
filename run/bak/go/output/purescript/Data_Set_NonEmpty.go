package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Set_NonEmpty_coerce gopurs_runtime.Value
var once_Data_Set_NonEmpty_coerce sync.Once
func Get_Data_Set_NonEmpty_coerce() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_coerce.Do(func() {
		cache_Data_Set_NonEmpty_coerce = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Set_NonEmpty_coerce
}

var cache_Data_Set_NonEmpty_NonEmptySet gopurs_runtime.Value
var once_Data_Set_NonEmpty_NonEmptySet sync.Once
func Get_Data_Set_NonEmpty_NonEmptySet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_NonEmptySet.Do(func() {
		cache_Data_Set_NonEmpty_NonEmptySet = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_NonEmptySet(x_0_box)
})
	})
	return cache_Data_Set_NonEmpty_NonEmptySet
}

var cache_Data_Set_NonEmpty_unionSet gopurs_runtime.Value
var once_Data_Set_NonEmpty_unionSet sync.Once
func Get_Data_Set_NonEmpty_unionSet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_unionSet.Do(func() {
		cache_Data_Set_NonEmpty_unionSet = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_unionSet(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Set_NonEmpty_unionSet
}

var cache_Data_Set_NonEmpty_toUnfoldable1 gopurs_runtime.Value
var once_Data_Set_NonEmpty_toUnfoldable1 sync.Once
func Get_Data_Set_NonEmpty_toUnfoldable1() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_toUnfoldable1.Do(func() {
		cache_Data_Set_NonEmpty_toUnfoldable1 = gopurs_runtime.Func(func(dictUnfoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_toUnfoldable1(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](dictUnfoldable1_0_box))
})
	})
	return cache_Data_Set_NonEmpty_toUnfoldable1
}

var cache_Data_Set_NonEmpty_toUnfoldable11 gopurs_runtime.Value
var once_Data_Set_NonEmpty_toUnfoldable11 sync.Once
func Get_Data_Set_NonEmpty_toUnfoldable11() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_toUnfoldable11.Do(func() {
		cache_Data_Set_NonEmpty_toUnfoldable11 = Call_Data_Set_NonEmpty_toUnfoldable1(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](Get_Data_List_Types_unfoldable1NonEmptyList()))
	})
	return cache_Data_Set_NonEmpty_toUnfoldable11
}

var cache_Data_Set_NonEmpty_toUnfoldable gopurs_runtime.Value
var once_Data_Set_NonEmpty_toUnfoldable sync.Once
func Get_Data_Set_NonEmpty_toUnfoldable() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_toUnfoldable.Do(func() {
		cache_Data_Set_NonEmpty_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dictUnfoldable_0_box))
})
	})
	return cache_Data_Set_NonEmpty_toUnfoldable
}

var cache_Data_Set_NonEmpty_toSet gopurs_runtime.Value
var once_Data_Set_NonEmpty_toSet sync.Once
func Get_Data_Set_NonEmpty_toSet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_toSet.Do(func() {
		cache_Data_Set_NonEmpty_toSet = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Set_NonEmpty_toSet(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box)))}
})
	})
	return cache_Data_Set_NonEmpty_toSet
}

var cache_Data_Set_NonEmpty_subset gopurs_runtime.Value
var once_Data_Set_NonEmpty_subset sync.Once
func Get_Data_Set_NonEmpty_subset() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_subset.Do(func() {
		cache_Data_Set_NonEmpty_subset = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_subset(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Set_NonEmpty_subset
}

var cache_Data_Set_NonEmpty_size gopurs_runtime.Value
var once_Data_Set_NonEmpty_size sync.Once
func Get_Data_Set_NonEmpty_size() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_size.Do(func() {
		cache_Data_Set_NonEmpty_size = Get_Data_Map_Internal_size()
	})
	return cache_Data_Set_NonEmpty_size
}

var cache_Data_Set_NonEmpty_singleton gopurs_runtime.Value
var once_Data_Set_NonEmpty_singleton sync.Once
func Get_Data_Set_NonEmpty_singleton() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_singleton.Do(func() {
		cache_Data_Set_NonEmpty_singleton = Get_Data_Set_singleton()
	})
	return cache_Data_Set_NonEmpty_singleton
}

var cache_Data_Set_NonEmpty_showNonEmptySet gopurs_runtime.Value
var once_Data_Set_NonEmpty_showNonEmptySet sync.Once
func Get_Data_Set_NonEmpty_showNonEmptySet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_showNonEmptySet.Do(func() {
		cache_Data_Set_NonEmpty_showNonEmptySet = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_showNonEmptySet(dictShow_0_box)
})
	})
	return cache_Data_Set_NonEmpty_showNonEmptySet
}

var cache_Data_Set_NonEmpty_semigroupNonEmptySet gopurs_runtime.Value
var once_Data_Set_NonEmpty_semigroupNonEmptySet sync.Once
func Get_Data_Set_NonEmpty_semigroupNonEmptySet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_semigroupNonEmptySet.Do(func() {
		cache_Data_Set_NonEmpty_semigroupNonEmptySet = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_semigroupNonEmptySet(dictOrd_0_box)
})
	})
	return cache_Data_Set_NonEmpty_semigroupNonEmptySet
}

var cache_Data_Set_NonEmpty_properSubset gopurs_runtime.Value
var once_Data_Set_NonEmpty_properSubset sync.Once
func Get_Data_Set_NonEmpty_properSubset() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_properSubset.Do(func() {
		cache_Data_Set_NonEmpty_properSubset = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_properSubset(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Set_NonEmpty_properSubset
}

var cache_Data_Set_NonEmpty_ordNonEmptySet gopurs_runtime.Value
var once_Data_Set_NonEmpty_ordNonEmptySet sync.Once
func Get_Data_Set_NonEmpty_ordNonEmptySet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_ordNonEmptySet.Do(func() {
		cache_Data_Set_NonEmpty_ordNonEmptySet = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_ordNonEmptySet(dictOrd_0_box)
})
	})
	return cache_Data_Set_NonEmpty_ordNonEmptySet
}

var cache_Data_Set_NonEmpty_ord1NonEmptySet gopurs_runtime.Value
var once_Data_Set_NonEmpty_ord1NonEmptySet sync.Once
func Get_Data_Set_NonEmpty_ord1NonEmptySet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_ord1NonEmptySet.Do(func() {
		cache_Data_Set_NonEmpty_ord1NonEmptySet = Get_Data_Set_ord1Set()
	})
	return cache_Data_Set_NonEmpty_ord1NonEmptySet
}

var cache_Data_Set_NonEmpty_min gopurs_runtime.Value
var once_Data_Set_NonEmpty_min sync.Once
func Get_Data_Set_NonEmpty_min() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_min.Do(func() {
		cache_Data_Set_NonEmpty_min = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_min(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box))
})
	})
	return cache_Data_Set_NonEmpty_min
}

var cache_Data_Set_NonEmpty_member gopurs_runtime.Value
var once_Data_Set_NonEmpty_member sync.Once
func Get_Data_Set_NonEmpty_member() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_member.Do(func() {
		cache_Data_Set_NonEmpty_member = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_member(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), k_1_box)
})
	})
	return cache_Data_Set_NonEmpty_member
}

var cache_Data_Set_NonEmpty_max gopurs_runtime.Value
var once_Data_Set_NonEmpty_max sync.Once
func Get_Data_Set_NonEmpty_max() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_max.Do(func() {
		cache_Data_Set_NonEmpty_max = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_max(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box))
})
	})
	return cache_Data_Set_NonEmpty_max
}

var cache_Data_Set_NonEmpty_mapMaybe gopurs_runtime.Value
var once_Data_Set_NonEmpty_mapMaybe sync.Once
func Get_Data_Set_NonEmpty_mapMaybe() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_mapMaybe.Do(func() {
		cache_Data_Set_NonEmpty_mapMaybe = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_mapMaybe(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Set_NonEmpty_mapMaybe
}

var cache_Data_Set_NonEmpty_go__map gopurs_runtime.Value
var once_Data_Set_NonEmpty_go__map sync.Once
func Get_Data_Set_NonEmpty_go__map() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_go__map.Do(func() {
		cache_Data_Set_NonEmpty_go__map = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Set_NonEmpty_go__map
}

var cache_Data_Set_NonEmpty_insert gopurs_runtime.Value
var once_Data_Set_NonEmpty_insert sync.Once
func Get_Data_Set_NonEmpty_insert() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_insert.Do(func() {
		cache_Data_Set_NonEmpty_insert = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_insert(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Set_NonEmpty_insert
}

var cache_Data_Set_NonEmpty_fromSet gopurs_runtime.Value
var once_Data_Set_NonEmpty_fromSet sync.Once
func Get_Data_Set_NonEmpty_fromSet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_fromSet.Do(func() {
		cache_Data_Set_NonEmpty_fromSet = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Set_NonEmpty_fromSet(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s_0_box)))}
})
	})
	return cache_Data_Set_NonEmpty_fromSet
}

var cache_Data_Set_NonEmpty_intersection gopurs_runtime.Value
var once_Data_Set_NonEmpty_intersection sync.Once
func Get_Data_Set_NonEmpty_intersection() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_intersection.Do(func() {
		cache_Data_Set_NonEmpty_intersection = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Set_NonEmpty_intersection(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v1_2_box)))}
})
	})
	return cache_Data_Set_NonEmpty_intersection
}

var cache_Data_Set_NonEmpty_fromFoldable1 gopurs_runtime.Value
var once_Data_Set_NonEmpty_fromFoldable1 sync.Once
func Get_Data_Set_NonEmpty_fromFoldable1() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_fromFoldable1.Do(func() {
		cache_Data_Set_NonEmpty_fromFoldable1 = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_fromFoldable1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_1_box))
})
	})
	return cache_Data_Set_NonEmpty_fromFoldable1
}

var cache_Data_Set_NonEmpty_fromFoldable gopurs_runtime.Value
var once_Data_Set_NonEmpty_fromFoldable sync.Once
func Get_Data_Set_NonEmpty_fromFoldable() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_fromFoldable.Do(func() {
		cache_Data_Set_NonEmpty_fromFoldable = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_1_box))
})
	})
	return cache_Data_Set_NonEmpty_fromFoldable
}

var cache_Data_Set_NonEmpty_foldableNonEmptySet gopurs_runtime.Value
var once_Data_Set_NonEmpty_foldableNonEmptySet sync.Once
func Get_Data_Set_NonEmpty_foldableNonEmptySet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_foldableNonEmptySet.Do(func() {
		cache_Data_Set_NonEmpty_foldableNonEmptySet = Get_Data_Set_foldableSet()
	})
	return cache_Data_Set_NonEmpty_foldableNonEmptySet
}

var cache_Data_Set_NonEmpty_foldable1NonEmptySet gopurs_runtime.Value
var once_Data_Set_NonEmpty_foldable1NonEmptySet sync.Once
func Get_Data_Set_NonEmpty_foldable1NonEmptySet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_foldable1NonEmptySet.Do(func() {
		cache_Data_Set_NonEmpty_foldable1NonEmptySet = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Set_foldableSet()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldable1NonEmptyList(), "foldMap1"), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dictSemigroup_0))}, f_1)
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(Get_Data_Set_NonEmpty_toUnfoldable11(), x_3))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Types_foldable1NonEmptyList(), "foldl1"), f_0)
_ = __local_var_1_1
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_1, gopurs_runtime.Apply(Get_Data_Set_NonEmpty_toUnfoldable11(), x_2))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Types_foldable1NonEmptyList(), "foldr1"), f_0)
_ = __local_var_1_2
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_2, gopurs_runtime.Apply(Get_Data_Set_NonEmpty_toUnfoldable11(), x_2))
})
}))
	})
	return cache_Data_Set_NonEmpty_foldable1NonEmptySet
}

var cache_Data_Set_NonEmpty_filter gopurs_runtime.Value
var once_Data_Set_NonEmpty_filter sync.Once
func Get_Data_Set_NonEmpty_filter() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_filter.Do(func() {
		cache_Data_Set_NonEmpty_filter = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_filter(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Set_NonEmpty_filter
}

var cache_Data_Set_NonEmpty_eqNonEmptySet gopurs_runtime.Value
var once_Data_Set_NonEmpty_eqNonEmptySet sync.Once
func Get_Data_Set_NonEmpty_eqNonEmptySet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_eqNonEmptySet.Do(func() {
		cache_Data_Set_NonEmpty_eqNonEmptySet = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_eqNonEmptySet(dictEq_0_box)
})
	})
	return cache_Data_Set_NonEmpty_eqNonEmptySet
}

var cache_Data_Set_NonEmpty_eq1NonEmptySet gopurs_runtime.Value
var once_Data_Set_NonEmpty_eq1NonEmptySet sync.Once
func Get_Data_Set_NonEmpty_eq1NonEmptySet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_eq1NonEmptySet.Do(func() {
		cache_Data_Set_NonEmpty_eq1NonEmptySet = Get_Data_Set_eq1Set()
	})
	return cache_Data_Set_NonEmpty_eq1NonEmptySet
}

var cache_Data_Set_NonEmpty_difference gopurs_runtime.Value
var once_Data_Set_NonEmpty_difference sync.Once
func Get_Data_Set_NonEmpty_difference() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_difference.Do(func() {
		cache_Data_Set_NonEmpty_difference = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Set_NonEmpty_difference(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v1_2_box)))}
})
	})
	return cache_Data_Set_NonEmpty_difference
}

var cache_Data_Set_NonEmpty_delete gopurs_runtime.Value
var once_Data_Set_NonEmpty_delete sync.Once
func Get_Data_Set_NonEmpty_delete() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_delete.Do(func() {
		cache_Data_Set_NonEmpty_delete = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Set_NonEmpty_delete(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), a_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_2_box)))}
})
	})
	return cache_Data_Set_NonEmpty_delete
}

var cache_Data_Set_NonEmpty_cons gopurs_runtime.Value
var once_Data_Set_NonEmpty_cons sync.Once
func Get_Data_Set_NonEmpty_cons() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_cons.Do(func() {
		cache_Data_Set_NonEmpty_cons = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_cons(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Set_NonEmpty_cons
}

var cache_Data_Set_NonEmpty_foldableNonEmptySet__1081688304 gopurs_runtime.Value
var once_Data_Set_NonEmpty_foldableNonEmptySet__1081688304 sync.Once
func Get_Data_Set_NonEmpty_foldableNonEmptySet__1081688304() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_foldableNonEmptySet__1081688304.Do(func() {
		cache_Data_Set_NonEmpty_foldableNonEmptySet__1081688304 = Get_Data_Set_foldableSet()
	})
	return cache_Data_Set_NonEmpty_foldableNonEmptySet__1081688304
}

var cache_Data_Set_NonEmpty_fromSet__3199996154 gopurs_runtime.Value
var once_Data_Set_NonEmpty_fromSet__3199996154 sync.Once
func Get_Data_Set_NonEmpty_fromSet__3199996154() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_fromSet__3199996154.Do(func() {
		cache_Data_Set_NonEmpty_fromSet__3199996154 = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Set_NonEmpty_fromSet__3199996154(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s_0_box)))}
})
	})
	return cache_Data_Set_NonEmpty_fromSet__3199996154
}

var cache_Data_Set_NonEmpty_fromSet__1805959329 gopurs_runtime.Value
var once_Data_Set_NonEmpty_fromSet__1805959329 sync.Once
func Get_Data_Set_NonEmpty_fromSet__1805959329() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_fromSet__1805959329.Do(func() {
		cache_Data_Set_NonEmpty_fromSet__1805959329 = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Set_NonEmpty_fromSet__1805959329(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s_0_box)))}
})
	})
	return cache_Data_Set_NonEmpty_fromSet__1805959329
}

var cache_Data_Set_NonEmpty_singleton__3724491835 gopurs_runtime.Value
var once_Data_Set_NonEmpty_singleton__3724491835 sync.Once
func Get_Data_Set_NonEmpty_singleton__3724491835() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_singleton__3724491835.Do(func() {
		cache_Data_Set_NonEmpty_singleton__3724491835 = Get_Data_Set_singleton()
	})
	return cache_Data_Set_NonEmpty_singleton__3724491835
}

var cache_Data_Set_NonEmpty_toUnfoldable1__800752263 gopurs_runtime.Value
var once_Data_Set_NonEmpty_toUnfoldable1__800752263 sync.Once
func Get_Data_Set_NonEmpty_toUnfoldable1__800752263() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_toUnfoldable1__800752263.Do(func() {
		cache_Data_Set_NonEmpty_toUnfoldable1__800752263 = gopurs_runtime.Func(func(dictUnfoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_toUnfoldable1__800752263(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](dictUnfoldable1_0_box))
})
	})
	return cache_Data_Set_NonEmpty_toUnfoldable1__800752263
}

func Call_Data_Set_NonEmpty_NonEmptySet(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Set_NonEmpty_unionSet(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_0, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_3))})))}
})
})
}

func Call_Data_Set_NonEmpty_toUnfoldable1(dictUnfoldable1_0_loop *Constructor_Data_Unfoldable1_Unfoldable1) gopurs_runtime.Value {
var dictUnfoldable1_0 *Constructor_Data_Unfoldable1_Unfoldable1 = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
// TAST (Let): stepNext_1_0 -> gopurs_runtime.Value
stepNext_1_0 := gopurs_runtime.Apply3(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapL(), gopurs_runtime.Func3(func(__local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value, __local_var_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, __local_var_1, __local_var_3})}})}
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}))
_ = stepNext_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable1_0.V0), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Tuple_functorTuple(), "map"), stepNext_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_2))})))}
}))
_ = __local_var_2_1
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply3(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapL(), gopurs_runtime.Func3(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, __local_var_3, __local_var_5})}
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.Str("toUnfoldable1: impossible"))
}))))}
}))
_ = __local_var_3_3
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_4), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})})
})
_ = __local_var_3_2
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_1, gopurs_runtime.Apply(__local_var_3_2, x_4))
})
}

func Call_Data_Set_NonEmpty_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable = dictUnfoldable_0_loop
_ = dictUnfoldable_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_List_toUnfoldable(), gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(dictUnfoldable_0)})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(Get_Data_Map_Internal_keys(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_2))})))})
})
}

func Call_Data_Set_NonEmpty_toSet(v_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Map_Internal_Node {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Set_NonEmpty_subset(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Set_subset(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_Data_Set_NonEmpty_showNonEmptySet(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): showNonEmptyArray_1_0 -> *Constructor_Data_Show_Show
showNonEmptyArray_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](gopurs_runtime.Apply(Get_Data_Array_NonEmpty_Internal_showNonEmptyArray(), dictShow_0))
_ = showNonEmptyArray_1_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(fromFoldable1 ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showNonEmptyArray_1_0.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Call_Data_Set_NonEmpty_toUnfoldable1(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](Get_Data_Unfoldable1_unfoldable1Array()))})), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s_2))}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).StrVal())) + (")"))
}))
}

func Call_Data_Set_NonEmpty_semigroupNonEmptySet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_0, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_3))})))}
})
}))
}

func Call_Data_Set_NonEmpty_properSubset(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Set_properSubset(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_Data_Set_NonEmpty_ordNonEmptySet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Set_ordSet(), dictOrd_0)
}

func Call_Data_Set_NonEmpty_min(v_0_loop *Constructor_Data_Map_Internal_Node) gopurs_runtime.Value {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
return gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> *Constructor_Data_Maybe_Just
__local_var_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v1_2, "key")
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Map_Internal_findMin(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)})))}))
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0 != nil) {
__t1 = (__local_var_2_0).V0
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

func Call_Data_Set_NonEmpty_member(dictOrd_0_loop *Constructor_Data_Ord_Ord, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_0 gopurs_runtime.Value
go__go_2_0_0 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_2_0_0:
for {
if false { continue go__go_2_0_0 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t3 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Bool(false)
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t2 gopurs_runtime.Value
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)}
continue go__go_2_0_0
__t2 = gopurs_runtime.Bool((gopurs_runtime.Value{}.IntVal) != (0))
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)}
continue go__go_2_0_0
__t2 = gopurs_runtime.Bool((gopurs_runtime.Value{}.IntVal) != (0))
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.Bool(true)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Bool((__t2.IntVal) != (0))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Bool((__t3.IntVal) != (0))
}
}()
})
return go__go_2_0_0
}

func Call_Data_Set_NonEmpty_max(v_0_loop *Constructor_Data_Map_Internal_Node) gopurs_runtime.Value {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
return gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> *Constructor_Data_Maybe_Just
__local_var_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v1_2, "key")
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Map_Internal_findMax(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)})))}))
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0 != nil) {
__t1 = (__local_var_2_0).V0
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

func Call_Data_Set_NonEmpty_mapMaybe(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Set_mapMaybe(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_Data_Set_NonEmpty_go__map(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Set_go__map(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_Data_Set_NonEmpty_insert(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Set_insert(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_Data_Set_NonEmpty_fromSet(s_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Maybe_Just {
var s_0 *Constructor_Data_Map_Internal_Node = s_0_loop
_ = s_0
var __t1 gopurs_runtime.Value
{
var __t0 bool
{
if (s_0 == nil) {
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s_0)}})}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t1)
}

func Call_Data_Set_NonEmpty_intersection(dictOrd_0_loop *Constructor_Data_Ord_Ord, v_1_loop *Constructor_Data_Map_Internal_Node, v1_2_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Maybe_Just {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var v_1 *Constructor_Data_Map_Internal_Node = v_1_loop
_ = v_1
var v1_2 *Constructor_Data_Map_Internal_Node = v1_2_loop
_ = v1_2
// TAST (Let): __local_var_3_0 -> *Constructor_Data_Map_Internal_Node
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), gopurs_runtime.Box(dictOrd_0.V1), Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v1_2)}))
_ = __local_var_3_0
var __t2 gopurs_runtime.Value
{
var __t1 bool
{
if (__local_var_3_0 == nil) {
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
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__local_var_3_0)}})}
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t2)
}

func Call_Data_Set_NonEmpty_fromFoldable1(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1, dictOrd_1_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictOrd_1 *Constructor_Data_Ord_Ord = dictOrd_1_loop
_ = dictOrd_1
// TAST (Let): compare_2_0 -> gopurs_runtime.Value
compare_2_0 := gopurs_runtime.Box(dictOrd_1.V1)
_ = compare_2_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_2_0, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_4))})))}
})
})})}, Get_Data_Set_singleton())
}

func Call_Data_Set_NonEmpty_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictOrd_1_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictOrd_1 *Constructor_Data_Ord_Ord = dictOrd_1_loop
_ = dictOrd_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(Get_Data_Set_fromFoldable(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_1)})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(__local_var_2_0, x_3)
_ = __local_var_4_1
var __t4 gopurs_runtime.Value
{
var __t3 bool
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_4_1)
if (__t_tag_2 == nil) {
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
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_4_1))}})}
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t4))}
})
}

func Call_Data_Set_NonEmpty_filter(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Set_filter(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_Data_Set_NonEmpty_eqNonEmptySet(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_Set_eqSet(), dictEq_0)
}

func Call_Data_Set_NonEmpty_difference(dictOrd_0_loop *Constructor_Data_Ord_Ord, v_1_loop *Constructor_Data_Map_Internal_Node, v1_2_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Maybe_Just {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var v_1 *Constructor_Data_Map_Internal_Node = v_1_loop
_ = v_1
var v1_2 *Constructor_Data_Map_Internal_Node = v1_2_loop
_ = v1_2
// TAST (Let): __local_var_3_0 -> *Constructor_Data_Map_Internal_Node
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeDifference(), gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v1_2)}))
_ = __local_var_3_0
var __t2 gopurs_runtime.Value
{
var __t1 bool
{
if (__local_var_3_0 == nil) {
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
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__local_var_3_0)}})}
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t2)
}

func Call_Data_Set_NonEmpty_delete(dictOrd_0_loop *Constructor_Data_Ord_Ord, a_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Maybe_Just {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var v_2 *Constructor_Data_Map_Internal_Node = v_2_loop
_ = v_2
// TAST (Let): __local_var_3_0 -> *Constructor_Data_Map_Internal_Node
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply3(Get_Data_Map_Internal_delete(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, a_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_2)}))
_ = __local_var_3_0
var __t2 gopurs_runtime.Value
{
var __t1 bool
{
if (__local_var_3_0 == nil) {
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
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__local_var_3_0)}})}
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t2)
}

func Call_Data_Set_NonEmpty_cons(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Set_insert(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_Data_Set_NonEmpty_fromSet__3199996154(s_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Maybe_Just {
var s_0 *Constructor_Data_Map_Internal_Node = s_0_loop
_ = s_0
var __t1 gopurs_runtime.Value
{
var __t0 bool
{
if (s_0 == nil) {
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s_0)}})}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t1)
}

func Call_Data_Set_NonEmpty_fromSet__1805959329(s_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Maybe_Just {
var s_0 *Constructor_Data_Map_Internal_Node = s_0_loop
_ = s_0
var __t1 gopurs_runtime.Value
{
var __t0 bool
{
if (s_0 == nil) {
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s_0)}})}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t1)
}

func Call_Data_Set_NonEmpty_toUnfoldable1__800752263(dictUnfoldable1_0_loop *Constructor_Data_Unfoldable1_Unfoldable1) gopurs_runtime.Value {
var dictUnfoldable1_0 *Constructor_Data_Unfoldable1_Unfoldable1 = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
// TAST (Let): stepNext_1_0 -> gopurs_runtime.Value
stepNext_1_0 := gopurs_runtime.Apply3(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapL(), gopurs_runtime.Func3(func(__local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value, __local_var_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, __local_var_1, __local_var_3})}})}
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}))
_ = stepNext_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable1_0.V0), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Tuple_functorTuple(), "map"), stepNext_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_2))})))}
}))
_ = __local_var_2_1
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply3(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapL(), gopurs_runtime.Func3(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, __local_var_3, __local_var_5})}
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.Str("toUnfoldable1: impossible"))
}))))}
}))
_ = __local_var_3_3
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_4), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})})
})
_ = __local_var_3_2
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_1, gopurs_runtime.Apply(__local_var_3_2, x_4))
})
}



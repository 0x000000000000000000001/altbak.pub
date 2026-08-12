package Data_Set_NonEmpty

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Array_NonEmpty_Internal "gopurs/output/Data.Array.NonEmpty.Internal"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_FoldableWithIndex "gopurs/output/Data.FoldableWithIndex"
	pkg_Data_Function "gopurs/output/Data.Function"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_List "gopurs/output/Data.List"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_Map_Internal "gopurs/output/Data.Map.Internal"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semigroup_Foldable "gopurs/output/Data.Semigroup.Foldable"
	pkg_Data_Set "gopurs/output/Data.Set"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
	pkg_Data_Unfoldable1 "gopurs/output/Data.Unfoldable1"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Partial "gopurs/output/Partial"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_unionSet gopurs_runtime.Value
var once_unionSet sync.Once
func Get_unionSet() gopurs_runtime.Value {
	once_unionSet.Do(func() {
		cache_unionSet = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unionSet(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_unionSet
}

var cache_toUnfoldable1 gopurs_runtime.Value
var once_toUnfoldable1 sync.Once
func Get_toUnfoldable1() gopurs_runtime.Value {
	once_toUnfoldable1.Do(func() {
		cache_toUnfoldable1 = gopurs_runtime.Func(func(dictUnfoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toUnfoldable1(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_0_box))
})
	})
	return cache_toUnfoldable1
}

var cache_toUnfoldable11 gopurs_runtime.Value
var once_toUnfoldable11 sync.Once
func Get_toUnfoldable11() gopurs_runtime.Value {
	once_toUnfoldable11.Do(func() {
		cache_toUnfoldable11 = Call_toUnfoldable1(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value]](pkg_Data_List_Types.Get_unfoldable1NonEmptyList()))
	})
	return cache_toUnfoldable11
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

var cache_toSet gopurs_runtime.Value
var once_toSet sync.Once
func Get_toSet() gopurs_runtime.Value {
	once_toSet.Do(func() {
		cache_toSet = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_toSet(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_toSet
}

var cache_subset gopurs_runtime.Value
var once_subset sync.Once
func Get_subset() gopurs_runtime.Value {
	once_subset.Do(func() {
		cache_subset = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_subset(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_subset
}

var cache_size gopurs_runtime.Value
var once_size sync.Once
func Get_size() gopurs_runtime.Value {
	once_size.Do(func() {
		cache_size = pkg_Data_Map_Internal.Get_size()
	})
	return cache_size
}

var cache_singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		cache_singleton = pkg_Data_Set.Get_singleton()
	})
	return cache_singleton
}

var cache_showNonEmptySet gopurs_runtime.Value
var once_showNonEmptySet sync.Once
func Get_showNonEmptySet() gopurs_runtime.Value {
	once_showNonEmptySet.Do(func() {
		cache_showNonEmptySet = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showNonEmptySet(dictShow_0_box)
})
	})
	return cache_showNonEmptySet
}

var cache_semigroupNonEmptySet gopurs_runtime.Value
var once_semigroupNonEmptySet sync.Once
func Get_semigroupNonEmptySet() gopurs_runtime.Value {
	once_semigroupNonEmptySet.Do(func() {
		cache_semigroupNonEmptySet = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupNonEmptySet(dictOrd_0_box)
})
	})
	return cache_semigroupNonEmptySet
}

var cache_properSubset gopurs_runtime.Value
var once_properSubset sync.Once
func Get_properSubset() gopurs_runtime.Value {
	once_properSubset.Do(func() {
		cache_properSubset = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_properSubset(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_properSubset
}

var cache_ordNonEmptySet gopurs_runtime.Value
var once_ordNonEmptySet sync.Once
func Get_ordNonEmptySet() gopurs_runtime.Value {
	once_ordNonEmptySet.Do(func() {
		cache_ordNonEmptySet = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordNonEmptySet(dictOrd_0_box)
})
	})
	return cache_ordNonEmptySet
}

var cache_ord1NonEmptySet gopurs_runtime.Value
var once_ord1NonEmptySet sync.Once
func Get_ord1NonEmptySet() gopurs_runtime.Value {
	once_ord1NonEmptySet.Do(func() {
		cache_ord1NonEmptySet = pkg_Data_Set.Get_ord1Set()
	})
	return cache_ord1NonEmptySet
}

var cache_min gopurs_runtime.Value
var once_min sync.Once
func Get_min() gopurs_runtime.Value {
	once_min.Do(func() {
		cache_min = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_min(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_min
}

var cache_member gopurs_runtime.Value
var once_member sync.Once
func Get_member() gopurs_runtime.Value {
	once_member.Do(func() {
		cache_member = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_member(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
})
	})
	return cache_member
}

var cache_max gopurs_runtime.Value
var once_max sync.Once
func Get_max() gopurs_runtime.Value {
	once_max.Do(func() {
		cache_max = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_max(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_max
}

var cache_mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		cache_mapMaybe = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_mapMaybe
}

var cache_go__map gopurs_runtime.Value
var once_go__map sync.Once
func Get_go__map() gopurs_runtime.Value {
	once_go__map.Do(func() {
		cache_go__map = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_go__map(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_go__map
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

var cache_fromSet gopurs_runtime.Value
var once_fromSet sync.Once
func Get_fromSet() gopurs_runtime.Value {
	once_fromSet.Do(func() {
		cache_fromSet = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_fromSet(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](s_0_box)))}
})
	})
	return cache_fromSet
}

var cache_intersection gopurs_runtime.Value
var once_intersection sync.Once
func Get_intersection() gopurs_runtime.Value {
	once_intersection.Do(func() {
		cache_intersection = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_intersection(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v1_2_box)))}
})
	})
	return cache_intersection
}

var cache_fromFoldable1 gopurs_runtime.Value
var once_fromFoldable1 sync.Once
func Get_fromFoldable1() gopurs_runtime.Value {
	once_fromFoldable1.Do(func() {
		cache_fromFoldable1 = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldable1(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_1_box))
})
	})
	return cache_fromFoldable1
}

var cache_fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		cache_fromFoldable = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldable(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_1_box))
})
	})
	return cache_fromFoldable
}

var cache_foldableNonEmptySet gopurs_runtime.Value
var once_foldableNonEmptySet sync.Once
func Get_foldableNonEmptySet() gopurs_runtime.Value {
	once_foldableNonEmptySet.Do(func() {
		cache_foldableNonEmptySet = pkg_Data_Set.Get_foldableSet()
	})
	return cache_foldableNonEmptySet
}

var cache_foldable1NonEmptySet gopurs_runtime.Value
var once_foldable1NonEmptySet sync.Once
func Get_foldable1NonEmptySet() gopurs_runtime.Value {
	once_foldable1NonEmptySet.Do(func() {
		cache_foldable1NonEmptySet = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Set.Get_foldableSet()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldable1NonEmptyList(), "foldMap1"), dictSemigroup_0, f_1)
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(Get_toUnfoldable11(), x_3))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldable1NonEmptyList(), "foldl1"), f_0)
_ = __local_var_1_1
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_1, gopurs_runtime.Apply(Get_toUnfoldable11(), x_2))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldable1NonEmptyList(), "foldr1"), f_0)
_ = __local_var_1_2
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_2, gopurs_runtime.Apply(Get_toUnfoldable11(), x_2))
})
}))
	})
	return cache_foldable1NonEmptySet
}

var cache_filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		cache_filter = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filter(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_filter
}

var cache_eqNonEmptySet gopurs_runtime.Value
var once_eqNonEmptySet sync.Once
func Get_eqNonEmptySet() gopurs_runtime.Value {
	once_eqNonEmptySet.Do(func() {
		cache_eqNonEmptySet = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqNonEmptySet(dictEq_0_box)
})
	})
	return cache_eqNonEmptySet
}

var cache_eq1NonEmptySet gopurs_runtime.Value
var once_eq1NonEmptySet sync.Once
func Get_eq1NonEmptySet() gopurs_runtime.Value {
	once_eq1NonEmptySet.Do(func() {
		cache_eq1NonEmptySet = pkg_Data_Set.Get_eq1Set()
	})
	return cache_eq1NonEmptySet
}

var cache_difference gopurs_runtime.Value
var once_difference sync.Once
func Get_difference() gopurs_runtime.Value {
	once_difference.Do(func() {
		cache_difference = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_difference(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v1_2_box)))}
})
	})
	return cache_difference
}

var cache_delete gopurs_runtime.Value
var once_delete sync.Once
func Get_delete() gopurs_runtime.Value {
	once_delete.Do(func() {
		cache_delete = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_delete(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_2_box)))}
})
	})
	return cache_delete
}

var cache_cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cache_cons = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cons(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_cons
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

var cache_unfoldable1NonEmptyArray__3769668500 gopurs_runtime.Value
var once_unfoldable1NonEmptyArray__3769668500 sync.Once
func Get_unfoldable1NonEmptyArray__3769668500() gopurs_runtime.Value {
	once_unfoldable1NonEmptyArray__3769668500.Do(func() {
		cache_unfoldable1NonEmptyArray__3769668500 = pkg_Data_Unfoldable1.Get_unfoldable1Array()
	})
	return cache_unfoldable1NonEmptyArray__3769668500
}

var cache_foldMap__4098395794 gopurs_runtime.Value
var once_foldMap__4098395794 sync.Once
func Get_foldMap__4098395794() gopurs_runtime.Value {
	once_foldMap__4098395794.Do(func() {
		cache_foldMap__4098395794 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap__4098395794(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap__4098395794
}

var cache_foldMap__3661646260 gopurs_runtime.Value
var once_foldMap__3661646260 sync.Once
func Get_foldMap__3661646260() gopurs_runtime.Value {
	once_foldMap__3661646260.Do(func() {
		cache_foldMap__3661646260 = gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldMap")
	})
	return cache_foldMap__3661646260
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

var cache_foldl__3943124669 gopurs_runtime.Value
var once_foldl__3943124669 sync.Once
func Get_foldl__3943124669() gopurs_runtime.Value {
	once_foldl__3943124669.Do(func() {
		cache_foldl__3943124669 = gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl")
	})
	return cache_foldl__3943124669
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

var cache_foldr__3943124669 gopurs_runtime.Value
var once_foldr__3943124669 sync.Once
func Get_foldr__3943124669() gopurs_runtime.Value {
	once_foldr__3943124669.Do(func() {
		cache_foldr__3943124669 = gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr")
	})
	return cache_foldr__3943124669
}

var cache_foldrWithIndex__2986161357 gopurs_runtime.Value
var once_foldrWithIndex__2986161357 sync.Once
func Get_foldrWithIndex__2986161357() gopurs_runtime.Value {
	once_foldrWithIndex__2986161357.Do(func() {
		cache_foldrWithIndex__2986161357 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrWithIndex__2986161357(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldrWithIndex__2986161357
}

var cache_foldrWithIndex__3511467915 gopurs_runtime.Value
var once_foldrWithIndex__3511467915 sync.Once
func Get_foldrWithIndex__3511467915() gopurs_runtime.Value {
	once_foldrWithIndex__3511467915.Do(func() {
		cache_foldrWithIndex__3511467915 = gopurs_runtime.RecordGet(pkg_Data_Map_Internal.Get_foldableWithIndexMap(), "foldrWithIndex")
	})
	return cache_foldrWithIndex__3511467915
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

var cache_flip__437353440 gopurs_runtime.Value
var once_flip__437353440 sync.Once
func Get_flip__437353440() gopurs_runtime.Value {
	once_flip__437353440.Do(func() {
		cache_flip__437353440 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__437353440(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__437353440
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

var cache_map__1206962620 gopurs_runtime.Value
var once_map__1206962620 sync.Once
func Get_map__1206962620() gopurs_runtime.Value {
	once_map__1206962620.Do(func() {
		cache_map__1206962620 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__1206962620
}

var cache_map__2311960860 gopurs_runtime.Value
var once_map__2311960860 sync.Once
func Get_map__2311960860() gopurs_runtime.Value {
	once_map__2311960860.Do(func() {
		cache_map__2311960860 = gopurs_runtime.RecordGet(pkg_Data_Tuple.Get_functorTuple(), "map")
	})
	return cache_map__2311960860
}

var cache_foldable1NonEmptyList__2239557029 gopurs_runtime.Value
var once_foldable1NonEmptyList__2239557029 sync.Once
func Get_foldable1NonEmptyList__2239557029() gopurs_runtime.Value {
	once_foldable1NonEmptyList__2239557029.Do(func() {
		cache_foldable1NonEmptyList__2239557029 = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_foldable1NonEmpty(), pkg_Data_List_Types.Get_foldableList())
	})
	return cache_foldable1NonEmptyList__2239557029
}

var cache_foldable1NonEmptyList__2630190169 gopurs_runtime.Value
var once_foldable1NonEmptyList__2630190169 sync.Once
func Get_foldable1NonEmptyList__2630190169() gopurs_runtime.Value {
	once_foldable1NonEmptyList__2630190169.Do(func() {
		cache_foldable1NonEmptyList__2630190169 = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_foldable1NonEmpty(), pkg_Data_List_Types.Get_foldableList())
	})
	return cache_foldable1NonEmptyList__2630190169
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
var go__go_1_3_1 gopurs_runtime.Value
go__go_1_3_1 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_3_1:
for {
if false { continue go__go_1_3_1 }
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
continue go__go_1_3_1
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
return go__go_1_3_1
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_5 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
})
}), b_1)
_ = __local_var_2_5
var go__go_3_6_2 gopurs_runtime.Value
go__go_3_6_2 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_6_2:
for {
if false { continue go__go_3_6_2 }
var v_4 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t7 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}
goto end_branch_7
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
continue go__go_3_6_2
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t7))}
}
}()
})
})
__local_var_4_8 := gopurs_runtime.Apply(go__go_3_6_2, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_4_8
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_5, gopurs_runtime.Apply(__local_var_4_8, x_5))
})
})
}))
	})
	return cache_foldableList__1753400174
}

var cache_findMax__2266220649 gopurs_runtime.Value
var once_findMax__2266220649 sync.Once
func Get_findMax__2266220649() gopurs_runtime.Value {
	once_findMax__2266220649.Do(func() {
		cache_findMax__2266220649 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findMax__2266220649(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_findMax__2266220649
}

var cache_findMax__528468393 gopurs_runtime.Value
var once_findMax__528468393 sync.Once
func Get_findMax__528468393() gopurs_runtime.Value {
	once_findMax__528468393.Do(func() {
		cache_findMax__528468393 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findMax__528468393(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_findMax__528468393
}

var cache_findMin__2266220649 gopurs_runtime.Value
var once_findMin__2266220649 sync.Once
func Get_findMin__2266220649() gopurs_runtime.Value {
	once_findMin__2266220649.Do(func() {
		cache_findMin__2266220649 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findMin__2266220649(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_findMin__2266220649
}

var cache_findMin__528468393 gopurs_runtime.Value
var once_findMin__528468393 sync.Once
func Get_findMin__528468393() gopurs_runtime.Value {
	once_findMin__528468393.Do(func() {
		cache_findMin__528468393 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findMin__528468393(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_findMin__528468393
}

var cache_foldableMap__373570208 gopurs_runtime.Value
var once_foldableMap__373570208 sync.Once
func Get_foldableMap__373570208() gopurs_runtime.Value {
	once_foldableMap__373570208.Do(func() {
		cache_foldableMap__373570208 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_3 gopurs_runtime.Value
_ = go__go_3_1_3
go__go_3_1_3 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr == nil) {
__t2 = gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr != nil) {
__t2 = gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(go__go_3_1_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}), gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(f_2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_1_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)})))
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
return go__go_3_1_3
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_3_4 gopurs_runtime.Value
_ = go__go_2_3_4
go__go_2_3_4 = gopurs_runtime.Func2(func(z_prime_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr == nil) {
__t4 = z_prime_3
goto end_branch_4
} else {

}
}
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr != nil) {
__t4 = gopurs_runtime.UncurriedApp2(go__go_2_3_4, gopurs_runtime.Apply2(f_0, gopurs_runtime.UncurriedApp2(go__go_2_3_4, z_prime_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V4)}), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V5)})
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
})
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_2_3_4, z_1, m_3)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_5_5 gopurs_runtime.Value
_ = go__go_2_5_5
go__go_2_5_5 = gopurs_runtime.Func2(func(m_prime_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (m_prime_3.Type == 9 && m_prime_3.IntVal == 324739070 && m_prime_3.UnsafePtr == nil) {
__t6 = __local_var_4
goto end_branch_6
} else {

}
}
{
if (m_prime_3.Type == 9 && m_prime_3.IntVal == 324739070 && m_prime_3.UnsafePtr != nil) {
__t6 = gopurs_runtime.UncurriedApp2(go__go_2_5_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime_3.UnsafePtr).V4)}, gopurs_runtime.Apply2(f_0, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_2_5_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime_3.UnsafePtr).V5)}, __local_var_4)))
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
})
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_2_5_5, m_3, z_1)
})
})
}))
	})
	return cache_foldableMap__373570208
}

var cache_foldableWithIndexMap__1634502082 gopurs_runtime.Value
var once_foldableWithIndexMap__1634502082 sync.Once
func Get_foldableWithIndexMap__1634502082() gopurs_runtime.Value {
	once_foldableWithIndexMap__1634502082.Do(func() {
		cache_foldableWithIndexMap__1634502082 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Map_Internal.Get_foldableMap()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_6 gopurs_runtime.Value
_ = go__go_3_1_6
go__go_3_1_6 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr == nil) {
__t2 = gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr != nil) {
__t2 = gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(go__go_3_1_6, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V4)}), gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply2(f_2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V3), gopurs_runtime.Apply(go__go_3_1_6, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V5)})))
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
return go__go_3_1_6
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_3_7 gopurs_runtime.Value
_ = go__go_2_3_7
go__go_2_3_7 = gopurs_runtime.Func2(func(z_prime_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr == nil) {
__t4 = z_prime_3
goto end_branch_4
} else {

}
}
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr != nil) {
__t4 = gopurs_runtime.UncurriedApp2(go__go_2_3_7, gopurs_runtime.Apply3(f_0, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V2, gopurs_runtime.UncurriedApp2(go__go_2_3_7, z_prime_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V4)}), (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V3), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4.UnsafePtr).V5)})
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
})
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_2_3_7, z_1, m_3)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_5_8 gopurs_runtime.Value
_ = go__go_2_5_8
go__go_2_5_8 = gopurs_runtime.Func2(func(m_prime_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (m_prime_3.Type == 9 && m_prime_3.IntVal == 324739070 && m_prime_3.UnsafePtr == nil) {
__t6 = __local_var_4
goto end_branch_6
} else {

}
}
{
if (m_prime_3.Type == 9 && m_prime_3.IntVal == 324739070 && m_prime_3.UnsafePtr != nil) {
__t6 = gopurs_runtime.UncurriedApp2(go__go_2_5_8, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime_3.UnsafePtr).V4)}, gopurs_runtime.Apply3(f_0, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime_3.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime_3.UnsafePtr).V3, gopurs_runtime.UncurriedApp2(go__go_2_5_8, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime_3.UnsafePtr).V5)}, __local_var_4)))
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
})
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(go__go_2_5_8, m_3, z_1)
})
})
}))
	})
	return cache_foldableWithIndexMap__1634502082
}

var cache_isEmpty__1620059593 gopurs_runtime.Value
var once_isEmpty__1620059593 sync.Once
func Get_isEmpty__1620059593() gopurs_runtime.Value {
	once_isEmpty__1620059593.Do(func() {
		cache_isEmpty__1620059593 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isEmpty__1620059593(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))
})
	})
	return cache_isEmpty__1620059593
}

var cache_iterMapL__878452066 gopurs_runtime.Value
var once_iterMapL__878452066 sync.Once
func Get_iterMapL__878452066() gopurs_runtime.Value {
	once_iterMapL__878452066.Do(func() {
		cache_iterMapL__878452066 = func() gopurs_runtime.Value {
var go__go_0_0_9 gopurs_runtime.Value
go__go_0_0_9 = gopurs_runtime.Func(func(iter_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var iter_1_loop gopurs_runtime.Value = iter_1_loop_val
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_0_0_9:
for {
if false { continue go__go_0_0_9 }
var iter_1 gopurs_runtime.Value = iter_1_loop
_ = iter_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t3 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t3 = iter_1
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 324739070 && __t_tag_1.UnsafePtr == nil) {
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3, iter_1})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
continue go__go_0_0_9
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5, iter_1})}})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
continue go__go_0_0_9
__t2 = gopurs_runtime.Value{}
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
return __t3
}
}()
})
})
return go__go_0_0_9
}()
	})
	return cache_iterMapL__878452066
}

var cache_keys__3504999702 gopurs_runtime.Value
var once_keys__3504999702 sync.Once
func Get_keys__3504999702() gopurs_runtime.Value {
	once_keys__3504999702.Do(func() {
		cache_keys__3504999702 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Map_Internal.Get_foldableWithIndexMap(), "foldrWithIndex"), gopurs_runtime.Func(func(k_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, k_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](acc_2)})}
})
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
	})
	return cache_keys__3504999702
}

var cache_keys__2406038214 gopurs_runtime.Value
var once_keys__2406038214 sync.Once
func Get_keys__2406038214() gopurs_runtime.Value {
	once_keys__2406038214.Do(func() {
		cache_keys__2406038214 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Map_Internal.Get_foldableWithIndexMap(), "foldrWithIndex"), gopurs_runtime.Func(func(k_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, k_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](acc_2)})}
})
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
	})
	return cache_keys__2406038214
}

var cache_singleton__943571066 gopurs_runtime.Value
var once_singleton__943571066 sync.Once
func Get_singleton__943571066() gopurs_runtime.Value {
	once_singleton__943571066.Do(func() {
		cache_singleton__943571066 = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_singleton__943571066(k_0_box, v_1_box))}
})
	})
	return cache_singleton__943571066
}

var cache_singleton__2450056090 gopurs_runtime.Value
var once_singleton__2450056090 sync.Once
func Get_singleton__2450056090() gopurs_runtime.Value {
	once_singleton__2450056090.Do(func() {
		cache_singleton__2450056090 = gopurs_runtime.Func2(func(k_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_singleton__2450056090(k_0_box, v_1_box))}
})
	})
	return cache_singleton__2450056090
}

var cache_size__909390430 gopurs_runtime.Value
var once_size__909390430 sync.Once
func Get_size__909390430() gopurs_runtime.Value {
	once_size__909390430.Do(func() {
		cache_size__909390430 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_size__909390430(v_0_box)
})
	})
	return cache_size__909390430
}

var cache_stepAscCps__4029696238 gopurs_runtime.Value
var once_stepAscCps__4029696238 sync.Once
func Get_stepAscCps__4029696238() gopurs_runtime.Value {
	once_stepAscCps__4029696238.Do(func() {
		cache_stepAscCps__4029696238 = gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_stepWith(), pkg_Data_Map_Internal.Get_iterMapL())
	})
	return cache_stepAscCps__4029696238
}

var cache_stepAscCps__3405418095 gopurs_runtime.Value
var once_stepAscCps__3405418095 sync.Once
func Get_stepAscCps__3405418095() gopurs_runtime.Value {
	once_stepAscCps__3405418095.Do(func() {
		cache_stepAscCps__3405418095 = gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_stepWith(), pkg_Data_Map_Internal.Get_iterMapL())
	})
	return cache_stepAscCps__3405418095
}

var cache_stepAscCps__3090303421 gopurs_runtime.Value
var once_stepAscCps__3090303421 sync.Once
func Get_stepAscCps__3090303421() gopurs_runtime.Value {
	once_stepAscCps__3090303421.Do(func() {
		cache_stepAscCps__3090303421 = gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_stepWith(), pkg_Data_Map_Internal.Get_iterMapL())
	})
	return cache_stepAscCps__3090303421
}

var cache_stepWith__3186376421 gopurs_runtime.Value
var once_stepWith__3186376421 sync.Once
func Get_stepWith__3186376421() gopurs_runtime.Value {
	once_stepWith__3186376421.Do(func() {
		cache_stepWith__3186376421 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, next_1_box gopurs_runtime.Value, done_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stepWith__3186376421(f_0_box, next_1_box, done_2_box)
})
	})
	return cache_stepWith__3186376421
}

var cache_toMapIter__1799172593 gopurs_runtime.Value
var once_toMapIter__1799172593 sync.Once
func Get_toMapIter__1799172593() gopurs_runtime.Value {
	once_toMapIter__1799172593.Do(func() {
		cache_toMapIter__1799172593 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toMapIter__1799172593(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](a_0_box))
})
	})
	return cache_toMapIter__1799172593
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

var cache_isNothing__2591355336 gopurs_runtime.Value
var once_isNothing__2591355336 sync.Once
func Get_isNothing__2591355336() gopurs_runtime.Value {
	once_isNothing__2591355336.Do(func() {
		cache_isNothing__2591355336 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isNothing__2591355336(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isNothing__2591355336
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

var cache_mempty__2312420373 gopurs_runtime.Value
var once_mempty__2312420373 sync.Once
func Get_mempty__2312420373() gopurs_runtime.Value {
	once_mempty__2312420373.Do(func() {
		cache_mempty__2312420373 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mempty__2312420373(dict_0_box)
})
	})
	return cache_mempty__2312420373
}

var cache_foldMap1__3342855683 gopurs_runtime.Value
var once_foldMap1__3342855683 sync.Once
func Get_foldMap1__3342855683() gopurs_runtime.Value {
	once_foldMap1__3342855683.Do(func() {
		cache_foldMap1__3342855683 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap1__3342855683(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap1__3342855683
}

var cache_foldMap1__3962279491 gopurs_runtime.Value
var once_foldMap1__3962279491 sync.Once
func Get_foldMap1__3962279491() gopurs_runtime.Value {
	once_foldMap1__3962279491.Do(func() {
		cache_foldMap1__3962279491 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap1__3962279491(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap1__3962279491
}

var cache_foldMap1__3810372805 gopurs_runtime.Value
var once_foldMap1__3810372805 sync.Once
func Get_foldMap1__3810372805() gopurs_runtime.Value {
	once_foldMap1__3810372805.Do(func() {
		cache_foldMap1__3810372805 = gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldable1NonEmptyList(), "foldMap1")
	})
	return cache_foldMap1__3810372805
}

var cache_foldl1__3059734942 gopurs_runtime.Value
var once_foldl1__3059734942 sync.Once
func Get_foldl1__3059734942() gopurs_runtime.Value {
	once_foldl1__3059734942.Do(func() {
		cache_foldl1__3059734942 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl1__3059734942(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl1__3059734942
}

var cache_foldl1__2849185176 gopurs_runtime.Value
var once_foldl1__2849185176 sync.Once
func Get_foldl1__2849185176() gopurs_runtime.Value {
	once_foldl1__2849185176.Do(func() {
		cache_foldl1__2849185176 = gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldable1NonEmptyList(), "foldl1")
	})
	return cache_foldl1__2849185176
}

var cache_foldr1__3059734942 gopurs_runtime.Value
var once_foldr1__3059734942 sync.Once
func Get_foldr1__3059734942() gopurs_runtime.Value {
	once_foldr1__3059734942.Do(func() {
		cache_foldr1__3059734942 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr1__3059734942(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr1__3059734942
}

var cache_foldr1__2849185176 gopurs_runtime.Value
var once_foldr1__2849185176 sync.Once
func Get_foldr1__2849185176() gopurs_runtime.Value {
	once_foldr1__2849185176.Do(func() {
		cache_foldr1__2849185176 = gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldable1NonEmptyList(), "foldr1")
	})
	return cache_foldr1__2849185176
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

var cache_foldableNonEmptySet__1081688304 gopurs_runtime.Value
var once_foldableNonEmptySet__1081688304 sync.Once
func Get_foldableNonEmptySet__1081688304() gopurs_runtime.Value {
	once_foldableNonEmptySet__1081688304.Do(func() {
		cache_foldableNonEmptySet__1081688304 = pkg_Data_Set.Get_foldableSet()
	})
	return cache_foldableNonEmptySet__1081688304
}

var cache_fromSet__3199996154 gopurs_runtime.Value
var once_fromSet__3199996154 sync.Once
func Get_fromSet__3199996154() gopurs_runtime.Value {
	once_fromSet__3199996154.Do(func() {
		cache_fromSet__3199996154 = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_fromSet__3199996154(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](s_0_box)))}
})
	})
	return cache_fromSet__3199996154
}

var cache_fromSet__1805959329 gopurs_runtime.Value
var once_fromSet__1805959329 sync.Once
func Get_fromSet__1805959329() gopurs_runtime.Value {
	once_fromSet__1805959329.Do(func() {
		cache_fromSet__1805959329 = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_fromSet__1805959329(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](s_0_box)))}
})
	})
	return cache_fromSet__1805959329
}

var cache_singleton__3724491835 gopurs_runtime.Value
var once_singleton__3724491835 sync.Once
func Get_singleton__3724491835() gopurs_runtime.Value {
	once_singleton__3724491835.Do(func() {
		cache_singleton__3724491835 = pkg_Data_Set.Get_singleton()
	})
	return cache_singleton__3724491835
}

var cache_toUnfoldable1__800752263 gopurs_runtime.Value
var once_toUnfoldable1__800752263 sync.Once
func Get_toUnfoldable1__800752263() gopurs_runtime.Value {
	once_toUnfoldable1__800752263.Do(func() {
		cache_toUnfoldable1__800752263 = gopurs_runtime.Func(func(dictUnfoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toUnfoldable1__800752263(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_0_box))
})
	})
	return cache_toUnfoldable1__800752263
}

var cache_delete__4217907800 gopurs_runtime.Value
var once_delete__4217907800 sync.Once
func Get_delete__4217907800() gopurs_runtime.Value {
	once_delete__4217907800.Do(func() {
		cache_delete__4217907800 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_delete__4217907800(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_delete__4217907800
}

var cache_difference__833514008 gopurs_runtime.Value
var once_difference__833514008 sync.Once
func Get_difference__833514008() gopurs_runtime.Value {
	once_difference__833514008.Do(func() {
		cache_difference__833514008 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_difference__833514008(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_difference__833514008
}

var cache_eq1Set__656709182 gopurs_runtime.Value
var once_eq1Set__656709182 sync.Once
func Get_eq1Set__656709182() gopurs_runtime.Value {
	once_eq1Set__656709182.Do(func() {
		cache_eq1Set__656709182 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(pkg_Data_Set.Get_eqSet(), dictEq_0), "eq")
}))
	})
	return cache_eq1Set__656709182
}

var cache_findMax__611847841 gopurs_runtime.Value
var once_findMax__611847841 sync.Once
func Get_findMax__611847841() gopurs_runtime.Value {
	once_findMax__611847841.Do(func() {
		cache_findMax__611847841 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findMax__611847841(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_findMax__611847841
}

var cache_findMin__611847841 gopurs_runtime.Value
var once_findMin__611847841 sync.Once
func Get_findMin__611847841() gopurs_runtime.Value {
	once_findMin__611847841.Do(func() {
		cache_findMin__611847841 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_findMin__611847841(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_findMin__611847841
}

var cache_foldableSet__1081688304 gopurs_runtime.Value
var once_foldableSet__1081688304 sync.Once
func Get_foldableSet__1081688304() gopurs_runtime.Value {
	once_foldableSet__1081688304.Do(func() {
		cache_foldableSet__1081688304 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldMap"), dictMonoid_0, f_1)
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_keys(), x_3))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), f_0, x_1)
_ = __local_var_2_1
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_1, gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_keys(), x_3))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), f_0, x_1)
_ = __local_var_2_2
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_2, gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_keys(), x_3))
})
})
}))
	})
	return cache_foldableSet__1081688304
}

var cache_intersection__833514008 gopurs_runtime.Value
var once_intersection__833514008 sync.Once
func Get_intersection__833514008() gopurs_runtime.Value {
	once_intersection__833514008.Do(func() {
		cache_intersection__833514008 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intersection__833514008(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_intersection__833514008
}

var cache_isEmpty__1390394866 gopurs_runtime.Value
var once_isEmpty__1390394866 sync.Once
func Get_isEmpty__1390394866() gopurs_runtime.Value {
	once_isEmpty__1390394866.Do(func() {
		cache_isEmpty__1390394866 = pkg_Data_Map_Internal.Get_isEmpty()
	})
	return cache_isEmpty__1390394866
}

var cache_ord1Set__1315498729 gopurs_runtime.Value
var once_ord1Set__1315498729 sync.Once
func Get_ord1Set__1315498729() gopurs_runtime.Value {
	once_ord1Set__1315498729.Do(func() {
		cache_ord1Set__1315498729 = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Set.Get_eq1Set()
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(pkg_Data_Set.Get_ordSet(), dictOrd_0), "compare")
}))
	})
	return cache_ord1Set__1315498729
}

var cache_singleton__3724491835 gopurs_runtime.Value
var once_singleton__3724491835 sync.Once
func Get_singleton__3724491835() gopurs_runtime.Value {
	once_singleton__3724491835.Do(func() {
		cache_singleton__3724491835 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_singleton__3724491835(a_0_box))}
})
	})
	return cache_singleton__3724491835
}

var cache_size__909390430 gopurs_runtime.Value
var once_size__909390430 sync.Once
func Get_size__909390430() gopurs_runtime.Value {
	once_size__909390430.Do(func() {
		cache_size__909390430 = pkg_Data_Map_Internal.Get_size()
	})
	return cache_size__909390430
}

var cache_toList__319294638 gopurs_runtime.Value
var once_toList__319294638 sync.Once
func Get_toList__319294638() gopurs_runtime.Value {
	once_toList__319294638.Do(func() {
		cache_toList__319294638 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_toList__319294638(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_toList__319294638
}

var cache_toMap__556171355 gopurs_runtime.Value
var once_toMap__556171355 sync.Once
func Get_toMap__556171355() gopurs_runtime.Value {
	once_toMap__556171355.Do(func() {
		cache_toMap__556171355 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_toMap__556171355(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_toMap__556171355
}

var cache_show__2742601362 gopurs_runtime.Value
var once_show__2742601362 sync.Once
func Get_show__2742601362() gopurs_runtime.Value {
	once_show__2742601362.Do(func() {
		cache_show__2742601362 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__2742601362(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__2742601362
}

var cache_show__4031350738 gopurs_runtime.Value
var once_show__4031350738 sync.Once
func Get_show__4031350738() gopurs_runtime.Value {
	once_show__4031350738.Do(func() {
		cache_show__4031350738 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__4031350738(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[[]gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__4031350738
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

var cache_functorTuple__2544689875 gopurs_runtime.Value
var once_functorTuple__2544689875 sync.Once
func Get_functorTuple__2544689875() gopurs_runtime.Value {
	once_functorTuple__2544689875.Do(func() {
		cache_functorTuple__2544689875 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0, gopurs_runtime.Apply(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_functorTuple__2544689875
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

var cache_unfoldable1Array__4196906331 gopurs_runtime.Value
var once_unfoldable1Array__4196906331 sync.Once
func Get_unfoldable1Array__4196906331() gopurs_runtime.Value {
	once_unfoldable1Array__4196906331.Do(func() {
		cache_unfoldable1Array__4196906331 = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Apply4(pkg_Data_Unfoldable1.Get_unfoldr1ArrayImpl(), pkg_Data_Maybe.Get_isNothing(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), pkg_Data_Tuple.Get_fst(), pkg_Data_Tuple.Get_snd()))
	})
	return cache_unfoldable1Array__4196906331
}

var cache_unfoldr1__2402610528 gopurs_runtime.Value
var once_unfoldr1__2402610528 sync.Once
func Get_unfoldr1__2402610528() gopurs_runtime.Value {
	once_unfoldr1__2402610528.Do(func() {
		cache_unfoldr1__2402610528 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr1__2402610528(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr1__2402610528
}

var cache_unfoldr1__78405858 gopurs_runtime.Value
var once_unfoldr1__78405858 sync.Once
func Get_unfoldr1__78405858() gopurs_runtime.Value {
	once_unfoldr1__78405858.Do(func() {
		cache_unfoldr1__78405858 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr1__78405858(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr1__78405858
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

var cache_unsafeCrashWith__2326648737 gopurs_runtime.Value
var once_unsafeCrashWith__2326648737 sync.Once
func Get_unsafeCrashWith__2326648737() gopurs_runtime.Value {
	once_unsafeCrashWith__2326648737.Do(func() {
		cache_unsafeCrashWith__2326648737 = gopurs_runtime.Func(func(msg_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeCrashWith__2326648737(msg_0_box.StrVal())
})
	})
	return cache_unsafeCrashWith__2326648737
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

var cache_coerce__529298068 gopurs_runtime.Value
var once_coerce__529298068 sync.Once
func Get_coerce__529298068() gopurs_runtime.Value {
	once_coerce__529298068.Do(func() {
		cache_coerce__529298068 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_coerce__529298068(_dollar__unused_0_box)
})
	})
	return cache_coerce__529298068
}

var cache_coerce__1801851476 gopurs_runtime.Value
var once_coerce__1801851476 sync.Once
func Get_coerce__1801851476() gopurs_runtime.Value {
	once_coerce__1801851476.Do(func() {
		cache_coerce__1801851476 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_coerce__1801851476
}

var cache_coerce__275695412 gopurs_runtime.Value
var once_coerce__275695412 sync.Once
func Get_coerce__275695412() gopurs_runtime.Value {
	once_coerce__275695412.Do(func() {
		cache_coerce__275695412 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_coerce__275695412
}

var cache_coerce__1245591956 gopurs_runtime.Value
var once_coerce__1245591956 sync.Once
func Get_coerce__1245591956() gopurs_runtime.Value {
	once_coerce__1245591956.Do(func() {
		cache_coerce__1245591956 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_coerce__1245591956
}

var cache_coerce__413730068 gopurs_runtime.Value
var once_coerce__413730068 sync.Once
func Get_coerce__413730068() gopurs_runtime.Value {
	once_coerce__413730068.Do(func() {
		cache_coerce__413730068 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_coerce__413730068
}

var cache_coerce__2766566164 gopurs_runtime.Value
var once_coerce__2766566164 sync.Once
func Get_coerce__2766566164() gopurs_runtime.Value {
	once_coerce__2766566164.Do(func() {
		cache_coerce__2766566164 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_coerce__2766566164
}

var cache_coerce__4188666772 gopurs_runtime.Value
var once_coerce__4188666772 sync.Once
func Get_coerce__4188666772() gopurs_runtime.Value {
	once_coerce__4188666772.Do(func() {
		cache_coerce__4188666772 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_coerce__4188666772
}

var cache_coerce__3944714388 gopurs_runtime.Value
var once_coerce__3944714388 sync.Once
func Get_coerce__3944714388() gopurs_runtime.Value {
	once_coerce__3944714388.Do(func() {
		cache_coerce__3944714388 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_coerce__3944714388
}

var cache_coerce__3115373716 gopurs_runtime.Value
var once_coerce__3115373716 sync.Once
func Get_coerce__3115373716() gopurs_runtime.Value {
	once_coerce__3115373716.Do(func() {
		cache_coerce__3115373716 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_coerce__3115373716
}

var cache_coerce__3108103700 gopurs_runtime.Value
var once_coerce__3108103700 sync.Once
func Get_coerce__3108103700() gopurs_runtime.Value {
	once_coerce__3108103700.Do(func() {
		cache_coerce__3108103700 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_coerce__3108103700
}

var cache_coerce__3836515092 gopurs_runtime.Value
var once_coerce__3836515092 sync.Once
func Get_coerce__3836515092() gopurs_runtime.Value {
	once_coerce__3836515092.Do(func() {
		cache_coerce__3836515092 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_coerce__3836515092
}

var cache_coerce__3671183764 gopurs_runtime.Value
var once_coerce__3671183764 sync.Once
func Get_coerce__3671183764() gopurs_runtime.Value {
	once_coerce__3671183764.Do(func() {
		cache_coerce__3671183764 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_coerce__3671183764
}

var cache_coerce__2052038804 gopurs_runtime.Value
var once_coerce__2052038804 sync.Once
func Get_coerce__2052038804() gopurs_runtime.Value {
	once_coerce__2052038804.Do(func() {
		cache_coerce__2052038804 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_coerce__2052038804
}

var cache_coerce__1870437204 gopurs_runtime.Value
var once_coerce__1870437204 sync.Once
func Get_coerce__1870437204() gopurs_runtime.Value {
	once_coerce__1870437204.Do(func() {
		cache_coerce__1870437204 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_coerce__1870437204
}

func Call_unionSet(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.V1
_ = compare_1_0
return gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), compare_1_0, pkg_Data_Function.Get_go__const(), m1_2, m2_3)))}
})
})
}

func Call_toUnfoldable1(dictUnfoldable1_0_loop *pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable1_0 *pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
stepNext_1_0 := gopurs_runtime.Apply3(pkg_Data_Map_Internal.Get_stepWith(), pkg_Data_Map_Internal.Get_iterMapL(), gopurs_runtime.Func3(func(k_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value, next_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, k_1, next_3})}})}
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}))
_ = stepNext_1_0
__local_var_2_1 := gopurs_runtime.Apply(dictUnfoldable1_0.V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Tuple.Get_functorTuple(), "map"), stepNext_1_0, v_2)))}
}))
_ = __local_var_2_1
__local_var_3_2 := gopurs_runtime.Apply3(pkg_Data_Map_Internal.Get_stepWith(), pkg_Data_Map_Internal.Get_iterMapL(), gopurs_runtime.Func3(func(k_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value, next_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, k_3, next_5})}
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str("toUnfoldable1: impossible"))))}
}))
_ = __local_var_3_2
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_1, gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_4), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}))
})
}

func Call_toUnfoldable(dictUnfoldable_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_List.Get_toUnfoldable(), gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(dictUnfoldable_0)})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_keys(), x_2))
})
}

func Call_toSet(v_0_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var v_0 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return v_0
}

func Call_subset(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Set.Get_subset(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_showNonEmptySet(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
showNonEmptyArray_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[[]gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Array_NonEmpty_Internal.Get_showNonEmptyArray(), dictShow_0))
_ = showNonEmptyArray_1_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str("(fromFoldable1 "), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Apply(showNonEmptyArray_1_0.V0, gopurs_runtime.Apply(Call_toUnfoldable1(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value]](pkg_Data_Unfoldable1.Get_unfoldable1Array())), s_2)), gopurs_runtime.Str(")"))).StrVal())
}))
}

func Call_semigroupNonEmptySet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), compare_1_0, pkg_Data_Function.Get_go__const(), m1_2, m2_3)))}
})
}))
}

func Call_properSubset(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Set.Get_properSubset(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_ordNonEmptySet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Set.Get_ordSet(), dictOrd_0)
}

func Call_min(v_0_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
__local_var_1_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v1_1, "key")
}), gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_findMin(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}))
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136 && __local_var_1_0.UnsafePtr != nil) {
__t1 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0
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

func Call_member(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
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
v1_4_1 := gopurs_runtime.Apply2(dictOrd_0.V1, k_1, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t2 gopurs_runtime.Value
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)}
continue go__go_2_0_0
__t2 = gopurs_runtime.Bool((gopurs_runtime.Value{}.IntVal) != (0))
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)}
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

func Call_max(v_0_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
__local_var_1_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v1_1, "key")
}), gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_findMax(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}))
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136 && __local_var_1_0.UnsafePtr != nil) {
__t1 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0
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

func Call_mapMaybe(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Set.Get_mapMaybe(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_go__map(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Set.Get_go__map(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_insert(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Set.Get_insert(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_fromSet(s_0_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
var s_0 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = s_0_loop
_ = s_0
var __t1 gopurs_runtime.Value
{
var __t0 bool
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s_0)}.UnsafePtr == nil) {
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s_0)}})}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]]](__t1)
}

func Call_intersection(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], v_1_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value], v1_2_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var v_1 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
_ = v_1
var v1_2 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v1_2_loop
_ = v1_2
__local_var_3_0 := gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeIntersectionWith(), dictOrd_0.V1, pkg_Data_Function.Get_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v1_2)})
_ = __local_var_3_0
var __t2 gopurs_runtime.Value
{
var __t1 bool
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 324739070 && __local_var_3_0.UnsafePtr == nil) {
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
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, __local_var_3_0})}
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]]](__t2)
}

func Call_fromFoldable1(dictFoldable1_0_loop *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value], dictOrd_1_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictOrd_1 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_1_loop
_ = dictOrd_1
compare_2_0 := dictOrd_1.V1
_ = compare_2_0
return gopurs_runtime.Apply2(dictFoldable1_0.V1, gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), compare_2_0, pkg_Data_Function.Get_go__const(), m1_3, m2_4)))}
})
})), pkg_Data_Set.Get_singleton())
}

func Call_fromFoldable(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], dictOrd_1_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictOrd_1 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_1_loop
_ = dictOrd_1
__local_var_2_0 := gopurs_runtime.Apply2(pkg_Data_Set.Get_fromFoldable(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_1)})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(__local_var_2_0, x_3)
_ = __local_var_4_1
var __t3 gopurs_runtime.Value
{
var __t2 bool
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 324739070 && __local_var_4_1.UnsafePtr == nil) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
if __t2 {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, __local_var_4_1})}
}
end_branch_3:
return __t3
})
}

func Call_filter(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Set.Get_filter(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_eqNonEmptySet(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(pkg_Data_Set.Get_eqSet(), dictEq_0)
}

func Call_difference(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], v_1_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value], v1_2_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var v_1 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
_ = v_1
var v1_2 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v1_2_loop
_ = v1_2
__local_var_3_0 := gopurs_runtime.UncurriedApp3(pkg_Data_Map_Internal.Get_unsafeDifference(), dictOrd_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v1_2)})
_ = __local_var_3_0
var __t2 gopurs_runtime.Value
{
var __t1 bool
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 324739070 && __local_var_3_0.UnsafePtr == nil) {
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
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, __local_var_3_0})}
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]]](__t2)
}

func Call_delete(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value, v_2_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var v_2 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_2_loop
_ = v_2
__local_var_3_0 := gopurs_runtime.Apply3(pkg_Data_Map_Internal.Get_delete(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, a_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_2)})
_ = __local_var_3_0
var __t2 gopurs_runtime.Value
{
var __t1 bool
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 324739070 && __local_var_3_0.UnsafePtr == nil) {
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
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, __local_var_3_0})}
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]]](__t2)
}

func Call_cons(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Set.Get_insert(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_foldMap__4098395794(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
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

func Call_foldrWithIndex__2986161357(dict_0_loop *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_const__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_flip__437353440(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_findMax__2266220649(v_0_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V5)}
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 324739070 && __t_tag_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("key", "value", (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V3)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_findMax(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V5)})))}
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t1))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t2)
}

func Call_findMax__528468393(v_0_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V5)}
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 324739070 && __t_tag_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("key", "value", (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V3)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_findMax(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V5)})))}
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t1))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t2)
}

func Call_findMin__2266220649(v_0_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V4)}
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 324739070 && __t_tag_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("key", "value", (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V3)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_findMin(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V4)})))}
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t1))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t2)
}

func Call_findMin__528468393(v_0_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V4)}
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 324739070 && __t_tag_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("key", "value", (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V3)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_findMin(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V4)})))}
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t1))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t2)
}

func Call_isEmpty__1620059593(v_0_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) bool {
var v_0 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 bool
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
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

func Call_singleton__943571066(k_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, k_0, v_1, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})})
}

func Call_singleton__2450056090(k_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, k_0, v_1, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})})
}

func Call_size__909390430(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 324739070 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 324739070 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Int(__t0.IntVal)
}

func Call_stepWith__3186376421(f_0_loop gopurs_runtime.Value, next_1_loop gopurs_runtime.Value, done_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var next_1 gopurs_runtime.Value = next_1_loop
_ = next_1
var done_2 gopurs_runtime.Value = done_2_loop
_ = done_2
var go__go_3_0_10 gopurs_runtime.Value
go__go_3_0_10 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_0_10:
for {
if false { continue go__go_3_0_10 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 2509360378) {
__t1 = gopurs_runtime.Apply(done_2, pkg_Data_Unit.Get_unit())
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1343415489) {
__t1 = gopurs_runtime.UncurriedApp3(next_1, (*pkg_Data_Map_Internal.Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*pkg_Data_Map_Internal.Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*pkg_Data_Map_Internal.Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2)
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 2861335956) {
v_4_loop = gopurs_runtime.Apply2(f_0, (*pkg_Data_Map_Internal.Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)})
continue go__go_3_0_10
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
return go__go_3_0_10
}

func Call_toMapIter__1799172593(a_0_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = a_0_loop
_ = a_0
return gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}
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

func Call_isNothing__2591355336(v2_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) bool {
var v2_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_0_loop
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

func Call_mempty__2312420373(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty")
}

func Call_foldMap1__3342855683(dict_0_loop *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldMap1__3962279491(dict_0_loop *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl1__3059734942(dict_0_loop *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr1__3059734942(dict_0_loop *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_fromSet__3199996154(s_0_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
var s_0 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = s_0_loop
_ = s_0
var __t1 gopurs_runtime.Value
{
var __t0 bool
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s_0)}.UnsafePtr == nil) {
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s_0)}})}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]]](__t1)
}

func Call_fromSet__1805959329(s_0_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
var s_0 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = s_0_loop
_ = s_0
var __t1 gopurs_runtime.Value
{
var __t0 bool
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s_0)}.UnsafePtr == nil) {
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s_0)}})}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]]](__t1)
}

func Call_toUnfoldable1__800752263(dictUnfoldable1_0_loop *pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable1_0 *pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
stepNext_1_0 := gopurs_runtime.Apply3(pkg_Data_Map_Internal.Get_stepWith(), pkg_Data_Map_Internal.Get_iterMapL(), gopurs_runtime.Func3(func(k_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value, next_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, k_1, next_3})}})}
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}))
_ = stepNext_1_0
__local_var_2_1 := gopurs_runtime.Apply(dictUnfoldable1_0.V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Tuple.Get_functorTuple(), "map"), stepNext_1_0, v_2)))}
}))
_ = __local_var_2_1
__local_var_3_2 := gopurs_runtime.Apply3(pkg_Data_Map_Internal.Get_stepWith(), pkg_Data_Map_Internal.Get_iterMapL(), gopurs_runtime.Func3(func(k_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value, next_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, k_3, next_5})}
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str("toUnfoldable1: impossible"))))}
}))
_ = __local_var_3_2
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_1, gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_4), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}))
})
}

func Call_delete__4217907800(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_delete(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_difference__833514008(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.V1
_ = compare_1_0
return gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(pkg_Data_Map_Internal.Get_unsafeDifference(), compare_1_0, m1_2, m2_3)))}
})
})
}

func Call_findMax__611847841(v_0_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v1_1, "key")
}), gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_findMax(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)})))
}

func Call_findMin__611847841(v_0_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v1_1, "key")
}), gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_findMin(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)})))
}

func Call_intersection__833514008(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.V1
_ = compare_1_0
return gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeIntersectionWith(), compare_1_0, pkg_Data_Function.Get_go__const(), m1_2, m2_3)))}
})
})
}

func Call_singleton__3724491835(a_0_loop gopurs_runtime.Value) *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, 1, 1, a_0, pkg_Data_Unit.Get_unit(), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}), gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})})
}

func Call_toList__319294638(v_0_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_keys(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}))
}

func Call_toMap__556171355(v_0_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var v_0 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return v_0
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__4031350738(dict_0_loop *pkg_Data_Show.Constructor_Show[[]gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[[]gopurs_runtime.Value] = dict_0_loop
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

func Call_unfoldr1__2402610528(dict_0_loop *pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_unfoldr1__78405858(dict_0_loop *pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_unsafeCrashWith__69763299(msg_0_loop string) gopurs_runtime.Value {
var msg_0 string = msg_0_loop
_ = msg_0
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str(msg_0))
}

func Call_unsafeCrashWith__2326648737(msg_0_loop string) gopurs_runtime.Value {
var msg_0 string = msg_0_loop
_ = msg_0
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str(msg_0))
}

func Call_crashWith__1894115486(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Partial.Get__crashWith()
}

func Call_coerce__529298068(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}



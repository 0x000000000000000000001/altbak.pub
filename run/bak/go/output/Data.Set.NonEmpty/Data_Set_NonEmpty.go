package Data_Set_NonEmpty

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unfoldable1 "gopurs/output/Data.Unfoldable1"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_Map_Internal "gopurs/output/Data.Map.Internal"
	pkg_Data_Set "gopurs/output/Data.Set"
	pkg_Data_Function "gopurs/output/Data.Function"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Partial "gopurs/output/Partial"
	pkg_Data_List "gopurs/output/Data.List"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Array_NonEmpty_Internal "gopurs/output/Data.Array.NonEmpty.Internal"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	unsafe "unsafe"
)

var cache_unionSet gopurs_runtime.Value
var once_unionSet sync.Once
func Get_unionSet() gopurs_runtime.Value {
	once_unionSet.Do(func() {
		cache_unionSet = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unionSet(dictOrd_0_box)
})
	})
	return cache_unionSet
}

var cache_toUnfoldable1 gopurs_runtime.Value
var once_toUnfoldable1 sync.Once
func Get_toUnfoldable1() gopurs_runtime.Value {
	once_toUnfoldable1.Do(func() {
		cache_toUnfoldable1 = gopurs_runtime.Func(func(dictUnfoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toUnfoldable1(dictUnfoldable1_0_box)
})
	})
	return cache_toUnfoldable1
}

var cache_toUnfoldable11 gopurs_runtime.Value
var once_toUnfoldable11 sync.Once
func Get_toUnfoldable11() gopurs_runtime.Value {
	once_toUnfoldable11.Do(func() {
		cache_toUnfoldable11 = Call_toUnfoldable1(pkg_Data_Unfoldable1.Get_unfoldable1Array())
	})
	return cache_toUnfoldable11
}

var cache_toUnfoldable12 gopurs_runtime.Value
var once_toUnfoldable12 sync.Once
func Get_toUnfoldable12() gopurs_runtime.Value {
	once_toUnfoldable12.Do(func() {
		cache_toUnfoldable12 = Call_toUnfoldable1(pkg_Data_List_Types.Get_unfoldable1NonEmptyList())
	})
	return cache_toUnfoldable12
}

var cache_toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		cache_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toUnfoldable(dictUnfoldable_0_box)
})
	})
	return cache_toUnfoldable
}

var cache_toSet gopurs_runtime.Value
var once_toSet sync.Once
func Get_toSet() gopurs_runtime.Value {
	once_toSet.Do(func() {
		cache_toSet = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_toSet((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_0_box.UnsafePtr)))}
})
	})
	return cache_toSet
}

var cache_subset gopurs_runtime.Value
var once_subset sync.Once
func Get_subset() gopurs_runtime.Value {
	once_subset.Do(func() {
		cache_subset = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_subset(dictOrd_0_box)
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
return Call_properSubset(dictOrd_0_box)
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
return Call_min((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_0_box.UnsafePtr))
})
	})
	return cache_min
}

var cache_member gopurs_runtime.Value
var once_member sync.Once
func Get_member() gopurs_runtime.Value {
	once_member.Do(func() {
		cache_member = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_member(dictOrd_0_box, k_1_box)
})
	})
	return cache_member
}

var cache_max gopurs_runtime.Value
var once_max sync.Once
func Get_max() gopurs_runtime.Value {
	once_max.Do(func() {
		cache_max = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_max((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_0_box.UnsafePtr))
})
	})
	return cache_max
}

var cache_mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		cache_mapMaybe = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe(dictOrd_0_box)
})
	})
	return cache_mapMaybe
}

var cache_map_ gopurs_runtime.Value
var once_map_ sync.Once
func Get_map_() gopurs_runtime.Value {
	once_map_.Do(func() {
		cache_map_ = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map_(dictOrd_0_box)
})
	})
	return cache_map_
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

var cache_fromSet gopurs_runtime.Value
var once_fromSet sync.Once
func Get_fromSet() gopurs_runtime.Value {
	once_fromSet.Do(func() {
		cache_fromSet = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_fromSet((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(s_0_box.UnsafePtr)))}
})
	})
	return cache_fromSet
}

var cache_intersection gopurs_runtime.Value
var once_intersection sync.Once
func Get_intersection() gopurs_runtime.Value {
	once_intersection.Do(func() {
		cache_intersection = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intersection(dictOrd_0_box)
})
	})
	return cache_intersection
}

var cache_fromFoldable1 gopurs_runtime.Value
var once_fromFoldable1 sync.Once
func Get_fromFoldable1() gopurs_runtime.Value {
	once_fromFoldable1.Do(func() {
		cache_fromFoldable1 = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldable1(dictFoldable1_0_box, dictOrd_1_box)
})
	})
	return cache_fromFoldable1
}

var cache_fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		cache_fromFoldable = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldable(dictFoldable_0_box, dictOrd_1_box)
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
foldMap11_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldable1NonEmptyList(), "foldMap1"), dictSemigroup_0)
_ = foldMap11_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(foldMap11_1_0, f_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(Get_toUnfoldable12(), x_4))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldable1NonEmptyList(), "foldl1"), f_0)
_ = __local_var_1_2
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_2, gopurs_runtime.Apply(Get_toUnfoldable12(), x_2))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldable1NonEmptyList(), "foldr1"), f_0)
_ = __local_var_1_3
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_3, gopurs_runtime.Apply(Get_toUnfoldable12(), x_2))
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
return Call_filter(dictOrd_0_box)
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
		cache_difference = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_difference(dictOrd_0_box)
})
	})
	return cache_difference
}

var cache_delete_ gopurs_runtime.Value
var once_delete_ sync.Once
func Get_delete_() gopurs_runtime.Value {
	once_delete_.Do(func() {
		cache_delete_ = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_delete_(dictOrd_0_box, a_1_box, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2_box.UnsafePtr)))}
})
	})
	return cache_delete_
}

var cache_cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cache_cons = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_cons(dictOrd_0_box)
})
	})
	return cache_cons
}

func Call_unionSet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), compare_1_0, pkg_Data_Function.Get_const_(), m1_2, m2_3)
})
}

func Call_toUnfoldable1(dictUnfoldable1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable1_0 gopurs_runtime.Value = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
stepNext_1_0 := gopurs_runtime.Apply3(pkg_Data_Map_Internal.Get_stepWith(), pkg_Data_Map_Internal.Get_iterMapL(), gopurs_runtime.Func3(func(k_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, next_3 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
if ((v_2.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(v_2.UnsafePtr).Rc) == (1)) {
(*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0 = func() gopurs_runtime.Value {
if ((v_2.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(v_2.UnsafePtr).Rc) == (1)) {
(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0 = k_1
(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1 = next_3
return v_2
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, k_1, next_3})}
}
}()
return v_2
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, func() gopurs_runtime.Value {
if ((v_2.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(v_2.UnsafePtr).Rc) == (1)) {
(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0 = k_1
(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1 = next_3
return v_2
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, k_1, next_3})}
}
}()})}
}
}()
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
}))
_ = stepNext_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable1_0, "unfoldr1"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Tuple.Get_functorTuple(), "map"), stepNext_1_0, v_2)
}))
_ = __local_var_2_1
__local_var_3_2 := gopurs_runtime.Apply3(pkg_Data_Map_Internal.Get_stepWith(), pkg_Data_Map_Internal.Get_iterMapL(), gopurs_runtime.Func3(func(k_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value, next_5 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
if ((__local_var_2_1.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(__local_var_2_1.UnsafePtr).Rc) == (1)) {
(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2_1.UnsafePtr).V0 = k_3
(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2_1.UnsafePtr).V1 = next_5
return __local_var_2_1
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, k_3, next_5})}
}
}()
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str("toUnfoldable1: impossible"))
}))
_ = __local_var_3_2
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_1, gopurs_runtime.Apply(__local_var_3_2, func() gopurs_runtime.Value {
if ((__local_var_2_1.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(__local_var_2_1.UnsafePtr).Rc) == (1)) {
(*pkg_Data_Map_Internal.Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2_1.UnsafePtr).V0 = (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr)
(*pkg_Data_Map_Internal.Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2_1.UnsafePtr).V1 = gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: nil}
return __local_var_2_1
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: nil}})}
}
}()))
})
}

func Call_toUnfoldable(dictUnfoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_List.Get_toUnfoldable(), dictUnfoldable_0)
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

func Call_subset(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func2(func(s1_2 gopurs_runtime.Value, s2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.UncurriedApp3(pkg_Data_Map_Internal.Get_unsafeDifference(), compare_1_0, s1_2, s2_3)
return gopurs_runtime.Bool((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 324739070 && __t_tag_1.UnsafePtr == nil))
})
}

func Call_showNonEmptySet(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(fromFoldable1 "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(pkg_Data_Array_NonEmpty_Internal.Get_showNonEmptyArray(), dictShow_0), "show"), gopurs_runtime.Apply(Get_toUnfoldable11(), s_1)), gopurs_runtime.Str(")")))
}))
}

func Call_semigroupNonEmptySet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), compare_1_0, pkg_Data_Function.Get_const_(), m1_2, m2_3)
}))
}

func Call_properSubset(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Set.Get_properSubset(), dictOrd_0)
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

func Call_member(dictOrd_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__2_0:
for {
if false { continue go__2_0 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
v1_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_1, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_2
var __t3 gopurs_runtime.Value
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 1527465420) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)}
continue go__2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 380165415) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)}
continue go__2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (v1_4_2.Type == 9 && v1_4_2.IntVal == 902936544) {
__t3 = gopurs_runtime.Bool(true)
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
return __t1
}
}()
})
return go__2_0
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

func Call_mapMaybe(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Set.Get_mapMaybe(), dictOrd_0)
}

func Call_map_(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Set.Get_map_(), dictOrd_0)
}

func Call_insert(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Set.Get_insert(), dictOrd_0)
}

func Call_fromSet(s_0_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
var s_0 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = s_0_loop
_ = s_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s_0)}})}
}
end_branch_0:
return (*pkg_Data_Maybe.Constructor_Just[*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]])(__t0.UnsafePtr)
}

func Call_intersection(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeIntersectionWith(), compare_1_0, pkg_Data_Function.Get_const_(), v_2, v1_3)
_ = __local_var_4_1
var __t2 gopurs_runtime.Value
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 324739070 && __local_var_4_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value {
if ((__local_var_4_1.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(__local_var_4_1.UnsafePtr).Rc) == (1)) {
(*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0 = __local_var_4_1
return __local_var_4_1
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, __local_var_4_1})}
}
}()
}
end_branch_2:
return __t2
})
}

func Call_fromFoldable1(dictFoldable1_0_loop gopurs_runtime.Value, dictOrd_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 gopurs_runtime.Value = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictOrd_1 gopurs_runtime.Value = dictOrd_1_loop
_ = dictOrd_1
compare_2_0 := gopurs_runtime.RecordGet(dictOrd_1, "compare")
_ = compare_2_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_0, "foldMap1"), gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(m1_3 gopurs_runtime.Value, m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), compare_2_0, pkg_Data_Function.Get_const_(), m1_3, m2_4)
})), pkg_Data_Set.Get_singleton())
}

func Call_fromFoldable(dictFoldable_0_loop gopurs_runtime.Value, dictOrd_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictOrd_1 gopurs_runtime.Value = dictOrd_1_loop
_ = dictOrd_1
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(m_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(pkg_Data_Map_Internal.Get_insert(), dictOrd_1, a_3, pkg_Data_Unit.Get_unit(), m_2)
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: nil})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(__local_var_2_0, x_3)
_ = __local_var_4_1
var __t2 gopurs_runtime.Value
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 324739070 && __local_var_4_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value {
if ((__local_var_2_0.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(__local_var_2_0.UnsafePtr).Rc) == (1)) {
(*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_2_0.UnsafePtr).V0 = __local_var_4_1
return __local_var_2_0
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, __local_var_4_1})}
}
}()
}
end_branch_2:
return __t2
})
}

func Call_filter(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Set.Get_filter(), dictOrd_0)
}

func Call_eqNonEmptySet(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(pkg_Data_Map_Internal.Get_eqMap(), dictEq_0, pkg_Data_Eq.Get_eqUnit()), "eq"), v_1, v1_2)
}))
}

func Call_difference(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.UncurriedApp3(pkg_Data_Map_Internal.Get_unsafeDifference(), compare_1_0, v_2, v1_3)
_ = __local_var_4_1
var __t2 gopurs_runtime.Value
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 324739070 && __local_var_4_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value {
if ((__local_var_4_1.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(__local_var_4_1.UnsafePtr).Rc) == (1)) {
(*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0 = __local_var_4_1
return __local_var_4_1
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, __local_var_4_1})}
}
}()
}
end_branch_2:
return __t2
})
}

func Call_delete_(dictOrd_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, v_2_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var v_2 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_2_loop
_ = v_2
__local_var_3_0 := gopurs_runtime.Apply3(pkg_Data_Map_Internal.Get_delete_(), dictOrd_0, a_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_2)})
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 324739070 && __local_var_3_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value {
if ((__local_var_3_0.UnsafePtr) != (nil)) && (((*struct{Rc uint32})(__local_var_3_0.UnsafePtr).Rc) == (1)) {
(*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0 = __local_var_3_0
return __local_var_3_0
} else {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, __local_var_3_0})}
}
}()
}
end_branch_1:
return (*pkg_Data_Maybe.Constructor_Just[*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]])(__t1.UnsafePtr)
}

func Call_cons(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Set.Get_insert(), dictOrd_0)
}



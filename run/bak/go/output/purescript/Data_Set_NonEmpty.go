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
		cache_Data_Set_NonEmpty_toUnfoldable11 = func() gopurs_runtime.Value {
// TAST (Let): stepNext_0_0 -> gopurs_runtime.Value
stepNext_0_0 := gopurs_runtime.Apply3(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapL(), gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, __local_var_0, __local_var_2})}})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}))
_ = stepNext_0_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](Get_Data_List_Types_unfoldable1NonEmptyList()).V0), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V0, gopurs_runtime.Apply(stepNext_0_0, (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V1)})}
}))
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply3(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapL(), gopurs_runtime.Func3(func(__local_var_2 gopurs_runtime.Value, __local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, __local_var_2, __local_var_4})}
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.Str("toUnfoldable1: impossible"))
}))))}
}))
_ = __local_var_2_3
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})})
})
_ = __local_var_2_2
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_1, gopurs_runtime.Apply(__local_var_2_2, x_3))
})
}()
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
		cache_Data_Set_NonEmpty_subset = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, s1_1_box gopurs_runtime.Value, s2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Set_NonEmpty_subset(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s1_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s2_2_box)))
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
		cache_Data_Set_NonEmpty_properSubset = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, s1_1_box gopurs_runtime.Value, s2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Set_NonEmpty_properSubset(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s1_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s2_2_box)))
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
		cache_Data_Set_NonEmpty_ord1NonEmptySet = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1](Get_Data_Set_ord1Set()))}
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
		cache_Data_Set_NonEmpty_mapMaybe = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_mapMaybe(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), f_1_box)
})
	})
	return cache_Data_Set_NonEmpty_mapMaybe
}

var cache_Data_Set_NonEmpty_go__map gopurs_runtime.Value
var once_Data_Set_NonEmpty_go__map sync.Once
func Get_Data_Set_NonEmpty_go__map() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_go__map.Do(func() {
		cache_Data_Set_NonEmpty_go__map = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), f_1_box)
})
	})
	return cache_Data_Set_NonEmpty_go__map
}

var cache_Data_Set_NonEmpty_insert gopurs_runtime.Value
var once_Data_Set_NonEmpty_insert sync.Once
func Get_Data_Set_NonEmpty_insert() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_insert.Do(func() {
		cache_Data_Set_NonEmpty_insert = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Set_NonEmpty_insert(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), a_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_2_box)))}
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
		cache_Data_Set_NonEmpty_foldableNonEmptySet = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_Set_foldableSet()))}
	})
	return cache_Data_Set_NonEmpty_foldableNonEmptySet
}

var cache_Data_Set_NonEmpty_foldable1NonEmptySet gopurs_runtime.Value
var once_Data_Set_NonEmpty_foldable1NonEmptySet sync.Once
func Get_Data_Set_NonEmpty_foldable1NonEmptySet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_foldable1NonEmptySet.Do(func() {
		cache_Data_Set_NonEmpty_foldable1NonEmptySet = gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Foldable_Foldable1{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_Set_foldableSet()))}
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): stepNext_2_1 -> gopurs_runtime.Value
stepNext_2_1 := gopurs_runtime.Apply3(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapL(), gopurs_runtime.Func3(func(__local_var_2 gopurs_runtime.Value, __local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, __local_var_2, __local_var_4})}})}
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}))
_ = stepNext_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](Get_Data_List_Types_unfoldable1NonEmptyList()).V0), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V0, gopurs_runtime.Apply(stepNext_2_1, (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V1)})}
}))
_ = __local_var_3_2
// TAST (Let): __local_var_4_4 -> gopurs_runtime.Value
__local_var_4_4 := gopurs_runtime.Apply3(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapL(), gopurs_runtime.Func3(func(__local_var_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, __local_var_4, __local_var_6})}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.Str("toUnfoldable1: impossible"))
}))))}
}))
_ = __local_var_4_4
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_5), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})})
})
_ = __local_var_4_3
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(__local_var_4_3, x_5))
})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_5 -> gopurs_runtime.Value
__local_var_4_5 := gopurs_runtime.Apply(__local_var_2_0, x_3)
_ = __local_var_4_5
var go__go_5_6_16 gopurs_runtime.Value
go__go_5_6_16 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_6_16:
for {
if false { continue go__go_5_6_16 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t7 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr == nil) {
__t7 = b_6
goto end_branch_7
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr != nil) {
b_6_loop = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), b_6, gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0))
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_6_16
__t7 = gopurs_runtime.Value{}
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
return gopurs_runtime.Apply2(go__go_5_6_16, gopurs_runtime.Apply(f_1, (*Constructor_Data_NonEmpty_NonEmpty)(__local_var_4_5.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(__local_var_4_5.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): stepNext_1_9 -> gopurs_runtime.Value
stepNext_1_9 := gopurs_runtime.Apply3(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapL(), gopurs_runtime.Func3(func(__local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value, __local_var_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, __local_var_1, __local_var_3})}})}
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}))
_ = stepNext_1_9
// TAST (Let): __local_var_2_10 -> gopurs_runtime.Value
__local_var_2_10 := gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](Get_Data_List_Types_unfoldable1NonEmptyList()).V0), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0, gopurs_runtime.Apply(stepNext_1_9, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1)})}
}))
_ = __local_var_2_10
// TAST (Let): __local_var_3_12 -> gopurs_runtime.Value
__local_var_3_12 := gopurs_runtime.Apply3(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapL(), gopurs_runtime.Func3(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, __local_var_3, __local_var_5})}
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.Str("toUnfoldable1: impossible"))
}))))}
}))
_ = __local_var_3_12
// TAST (Let): __local_var_3_11 -> gopurs_runtime.Value
__local_var_3_11 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_12, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_4), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})})
})
_ = __local_var_3_11
// TAST (Let): __local_var_1_8 -> gopurs_runtime.Value
__local_var_1_8 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_10, gopurs_runtime.Apply(__local_var_3_11, x_4))
})
_ = __local_var_1_8
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_13 -> gopurs_runtime.Value
__local_var_3_13 := gopurs_runtime.Apply(__local_var_1_8, x_2)
_ = __local_var_3_13
var go__go_4_14_17 gopurs_runtime.Value
go__go_4_14_17 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_14_17:
for {
if false { continue go__go_4_14_17 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t15 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t15 = b_5
goto end_branch_15
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Apply2(f_0, b_5, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0)
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_14_17
__t15 = gopurs_runtime.Value{}
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
return gopurs_runtime.Apply2(go__go_4_14_17, (*Constructor_Data_NonEmpty_NonEmpty)(__local_var_3_13.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(__local_var_3_13.UnsafePtr).V1)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): stepNext_1_17 -> gopurs_runtime.Value
stepNext_1_17 := gopurs_runtime.Apply3(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapL(), gopurs_runtime.Func3(func(__local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value, __local_var_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, __local_var_1, __local_var_3})}})}
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}))
_ = stepNext_1_17
// TAST (Let): __local_var_2_18 -> gopurs_runtime.Value
__local_var_2_18 := gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](Get_Data_List_Types_unfoldable1NonEmptyList()).V0), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0, gopurs_runtime.Apply(stepNext_1_17, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1)})}
}))
_ = __local_var_2_18
// TAST (Let): __local_var_3_20 -> gopurs_runtime.Value
__local_var_3_20 := gopurs_runtime.Apply3(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapL(), gopurs_runtime.Func3(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, __local_var_3, __local_var_5})}
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.Str("toUnfoldable1: impossible"))
}))))}
}))
_ = __local_var_3_20
// TAST (Let): __local_var_3_19 -> gopurs_runtime.Value
__local_var_3_19 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_20, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_4), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})})
})
_ = __local_var_3_19
// TAST (Let): __local_var_1_16 -> gopurs_runtime.Value
__local_var_1_16 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_18, gopurs_runtime.Apply(__local_var_3_19, x_4))
})
_ = __local_var_1_16
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_21 -> gopurs_runtime.Value
__local_var_3_21 := gopurs_runtime.Apply(__local_var_1_16, x_2)
_ = __local_var_3_21
// TAST (Let): __local_var_4_22 -> gopurs_runtime.Value
__local_var_4_22 := gopurs_runtime.Apply(f_0, (*Constructor_Data_NonEmpty_NonEmpty)(__local_var_3_21.UnsafePtr).V0)
_ = __local_var_4_22
var go__go_5_24_18 gopurs_runtime.Value
go__go_5_24_18 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_24_18:
for {
if false { continue go__go_5_24_18 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t27 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr == nil) {
__t27 = b_6
goto end_branch_27
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr != nil) {
// TAST (Let): __local_var_8_25 -> gopurs_runtime.Value
__local_var_8_25 := gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0)
_ = __local_var_8_25
var __t26 gopurs_runtime.Value
{
if (b_6.Type == 9 && b_6.IntVal == 930809136 && b_6.UnsafePtr == nil) {
__t26 = (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0
goto end_branch_26
} else {

}
}
{
if (b_6.Type == 9 && b_6.IntVal == 930809136 && b_6.UnsafePtr != nil) {
__t26 = gopurs_runtime.Apply(__local_var_8_25, (*Constructor_Data_Maybe_Just)(b_6.UnsafePtr).V0)
goto end_branch_26
} else {

}
}
{
__t26 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_26:
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, __t26})}
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_24_18
__t27 = gopurs_runtime.Value{}
goto end_branch_27
} else {

}
}
{
__t27 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_27:
return __t27
}
}()
})
})
var go__go_6_28_19 gopurs_runtime.Value
go__go_6_28_19 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_7_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_7_loop_val)
var v1_8_loop gopurs_runtime.Value = v1_8_loop_val
go__go_6_28_19:
for {
if false { continue go__go_6_28_19 }
var v_7 *Constructor_Data_List_Types_Cons = v_7_loop
_ = v_7
var v1_8 gopurs_runtime.Value = v1_8_loop
_ = v1_8
var __t29 *Constructor_Data_List_Types_Cons
{
if (v1_8.Type == 9 && v1_8.IntVal == 1358893437 && v1_8.UnsafePtr == nil) {
__t29 = v_7
goto end_branch_29
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 1358893437 && v1_8.UnsafePtr != nil) {
v_7_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_8.UnsafePtr).V0, v_7})})
v1_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_8.UnsafePtr).V1)}
continue go__go_6_28_19
__t29 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_29
} else {

}
}
{
__t29 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_29:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t29)}
}
}()
})
})
// TAST (Let): __local_var_5_23 -> *Constructor_Data_Maybe_Just
__local_var_5_23 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(go__go_5_24_18, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Apply2(go__go_6_28_19, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(__local_var_3_21.UnsafePtr).V1)))
_ = __local_var_5_23
var __t30 gopurs_runtime.Value
{
if (__local_var_5_23 == nil) {
__t30 = (*Constructor_Data_NonEmpty_NonEmpty)(__local_var_3_21.UnsafePtr).V0
goto end_branch_30
} else {

}
}
{
if (__local_var_5_23 != nil) {
__t30 = gopurs_runtime.Apply(__local_var_4_22, (__local_var_5_23).V0)
goto end_branch_30
} else {

}
}
{
__t30 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_30:
return __t30
})
})})}
	})
	return cache_Data_Set_NonEmpty_foldable1NonEmptySet
}

var cache_Data_Set_NonEmpty_filter gopurs_runtime.Value
var once_Data_Set_NonEmpty_filter sync.Once
func Get_Data_Set_NonEmpty_filter() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_filter.Do(func() {
		cache_Data_Set_NonEmpty_filter = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_filter(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), f_1_box)
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
		cache_Data_Set_NonEmpty_eq1NonEmptySet = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1](Get_Data_Set_eq1Set()))}
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
		cache_Data_Set_NonEmpty_cons = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Set_NonEmpty_cons(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), a_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_2_box)))}
})
	})
	return cache_Data_Set_NonEmpty_cons
}

var cache_Data_Set_NonEmpty_foldableNonEmptySet__2661034898 gopurs_runtime.Value
var once_Data_Set_NonEmpty_foldableNonEmptySet__2661034898 sync.Once
func Get_Data_Set_NonEmpty_foldableNonEmptySet__2661034898() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_foldableNonEmptySet__2661034898.Do(func() {
		cache_Data_Set_NonEmpty_foldableNonEmptySet__2661034898 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_Set_foldableSet()))}
	})
	return cache_Data_Set_NonEmpty_foldableNonEmptySet__2661034898
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
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0, gopurs_runtime.Apply(stepNext_1_0, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1)})}
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
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable_0.V1), gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_Maybe_Just
{
var __t_tag_2 *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_1)
if (__t_tag_2 == nil) {
__t4 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_4
} else {

}
}
{
var __t_tag_3 *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_1)
if (__t_tag_3 != nil) {
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("head", "tail", (*Constructor_Data_List_Types_Cons)(xs_1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(xs_1.UnsafePtr).V1)})}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
// TAST (Let): __local_var_2_1 -> *Constructor_Data_Maybe_Just
var __local_var_2_1 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)})
var __t5 *Constructor_Data_Maybe_Just
{
if (__local_var_2_1 != nil) {
__t5 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.RecordGet((__local_var_2_1).V0, "head"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet((__local_var_2_1).V0, "tail")))}})}}
goto end_branch_5
} else {

}
}
{
__t5 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)}
}))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_6_0 gopurs_runtime.Value
_ = go__go_3_6_0
go__go_3_6_0 = gopurs_runtime.Func2(func(__local_var_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr == nil) {
__t7 = __local_var_5
goto end_branch_7
} else {

}
}
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr != nil) {
__t7 = gopurs_runtime.UncurriedApp2(go__go_3_6_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_3_6_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V5)}, __local_var_5))})})
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
})
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_3_6_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
}

func Call_Data_Set_NonEmpty_toSet(v_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Map_Internal_Node {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Set_NonEmpty_subset(dictOrd_0_loop *Constructor_Data_Ord_Ord, s1_1_loop *Constructor_Data_Map_Internal_Node, s2_2_loop *Constructor_Data_Map_Internal_Node) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var s1_1 *Constructor_Data_Map_Internal_Node = s1_1_loop
_ = s1_1
var s2_2 *Constructor_Data_Map_Internal_Node = s2_2_loop
_ = s2_2
var __t1 bool
{
var __t_tag_0 *Constructor_Data_Map_Internal_Node = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeDifference(), gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s1_1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s2_2)}))
if (__t_tag_0 == nil) {
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

func Call_Data_Set_NonEmpty_showNonEmptySet(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): showArray_1_1 -> *Constructor_Data_Show_Show
showArray_1_1 := &Constructor_Data_Show_Show{1, gopurs_runtime.Apply(Get_Data_Show_showArrayImpl(), gopurs_runtime.RecordGet(dictShow_0, "show"))}
_ = showArray_1_1
// TAST (Let): showNonEmptyArray_1_0 -> *Constructor_Data_Show_Show
showNonEmptyArray_1_0 := &Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(NonEmptyArray ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showArray_1_1.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).StrVal())) + (")"))
})}
_ = showNonEmptyArray_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): stepNext_3_2 -> gopurs_runtime.Value
stepNext_3_2 := gopurs_runtime.Apply3(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapL(), gopurs_runtime.Func3(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, __local_var_3, __local_var_5})}})}
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}))
_ = stepNext_3_2
return gopurs_runtime.Str((("(fromFoldable1 ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showNonEmptyArray_1_0.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply6(Get_Data_Unfoldable1_unfoldr1ArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 930809136 && v_5.UnsafePtr != nil) {
__t3 = (*Constructor_Data_Maybe_Just)(v_5.UnsafePtr).V0
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
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0, gopurs_runtime.Apply(stepNext_3_2, (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1)})}
}), gopurs_runtime.Apply4(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapL(), gopurs_runtime.Func3(func(__local_var_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, __local_var_4, __local_var_6})}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.Str("toUnfoldable1: impossible"))
}))))}
}), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s_2), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})})).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).StrVal())) + (")"))
})})}
}

func Call_Data_Set_NonEmpty_semigroupNonEmptySet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_0, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_3))})))}
})
})})}
}

func Call_Data_Set_NonEmpty_properSubset(dictOrd_0_loop *Constructor_Data_Ord_Ord, s1_1_loop *Constructor_Data_Map_Internal_Node, s2_2_loop *Constructor_Data_Map_Internal_Node) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var s1_1 *Constructor_Data_Map_Internal_Node = s1_1_loop
_ = s1_1
var s2_2 *Constructor_Data_Map_Internal_Node = s2_2_loop
_ = s2_2
var __t0 int64
{
if (s1_1 == nil) {
__t0 = 0
goto end_branch_0
} else {

}
}
{
if (s1_1 != nil) {
__t0 = (s1_1).V1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_0:
var __t1 int64
{
if (s2_2 == nil) {
__t1 = 0
goto end_branch_1
} else {

}
}
{
if (s2_2 != nil) {
__t1 = (s2_2).V1
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_1:
var __t_and_4 bool = false
if ((__t0) == (__t1)) != (true) {

var __t3 bool
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeDifference(), gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s1_1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s2_2)}))
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
__t_and_4 = __t3
}
return __t_and_4
}

func Call_Data_Set_NonEmpty_ordNonEmptySet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_2
// TAST (Let): eqList1_1_1 -> *Constructor_Data_Eq_Eq
eqList1_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_3_1 gopurs_runtime.Value
go__go_4_3_1 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
var v2_7_loop bool = (v2_7_loop_val.IntVal) != (0)
go__go_4_3_1:
for {
if false { continue go__go_4_3_1 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var v2_7 bool = v2_7_loop
_ = v2_7
var __t5 bool
{
if (v2_7) != (true) {
__t5 = false
goto end_branch_5
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
var __t4 bool
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t4 = v2_7
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
if ((v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil)) && ((v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil)) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
v2_7_loop = (v2_7) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "eq"), (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0).IntVal) != (0))
continue go__go_4_3_1
__t5 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_5
} else {

}
}
{
__t5 = false
}
end_branch_5:
return gopurs_runtime.Bool(__t5)
}
}()
})
})
})
return gopurs_runtime.Bool((gopurs_runtime.Apply3(go__go_4_3_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_3))}, gopurs_runtime.Bool(true)).IntVal) != (0))
})
})))
_ = eqList1_1_1
// TAST (Let): ordList_1_0 -> *Constructor_Data_Ord_Ord
ordList_1_0 := &Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqList1_1_1)}
}), gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_6_2 gopurs_runtime.Value
go__go_4_6_2 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_6_2:
for {
if false { continue go__go_4_6_2 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t10 uint32
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
var __t7 uint32
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t7 = 902936544
goto end_branch_7
} else {

}
}
{
__t7 = 1527465420
}
end_branch_7:
__t10 = __t7
goto end_branch_10
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t10 = 380165415
goto end_branch_10
} else {

}
}
{
if ((v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil)) && ((v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil)) {
// TAST (Let): v2_7_8 -> gopurs_runtime.Value
v2_7_8 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0)
_ = v2_7_8
var __t9 uint32
{
if (uint32(v2_7_8.IntVal) == 902936544) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_6_2
__t9 = uint32(gopurs_runtime.Value{}.IntVal)
goto end_branch_9
} else {

}
}
{
__t9 = uint32(v2_7_8.IntVal)
}
end_branch_9:
__t10 = __t9
goto end_branch_10
} else {

}
}
{
__t10 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t10), UnsafePtr: nil}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(go__go_4_6_2, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_3))}).IntVal)), UnsafePtr: nil}
})
})}
_ = ordList_1_0
// TAST (Let): __local_var_2_12 -> gopurs_runtime.Value
__local_var_2_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_12
var go__go_3_15_3 gopurs_runtime.Value
go__go_3_15_3 = gopurs_runtime.Func(func(a_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_4_loop gopurs_runtime.Value = a_4_loop_val
var b_5_loop gopurs_runtime.Value = b_5_loop_val
go__go_3_15_3:
for {
if false { continue go__go_3_15_3 }
var a_4 gopurs_runtime.Value = a_4_loop
_ = a_4
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
// TAST (Let): v_6_16 -> *Constructor_Data_Map_Internal_IterNext
v_6_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_4))
_ = v_6_16
var __t19 bool
{
if (v_6_16 != nil) {
// TAST (Let): v2_7_17 -> *Constructor_Data_Map_Internal_IterNext
v2_7_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_5))
_ = v2_7_17
var __t18 bool
{
if ((v2_7_17 != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_12, "eq"), (v_6_16).V0, (v2_7_17).V0).IntVal) != (0)) {
a_4_loop = (v_6_16).V2
b_5_loop = (v2_7_17).V2
continue go__go_3_15_3
__t18 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_18
} else {

}
}
{
__t18 = false
}
end_branch_18:
__t19 = __t18
goto end_branch_19
} else {

}
}
{
if (v_6_16 == nil) {
__t19 = true
goto end_branch_19
} else {

}
}
{
__t19 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_19:
return gopurs_runtime.Bool(__t19)
}
}()
})
})
// TAST (Let): eqMapIter2_3_14 -> *Constructor_Data_Eq_Eq
eqMapIter2_3_14 := &Constructor_Data_Eq_Eq{1, go__go_3_15_3}
_ = eqMapIter2_3_14
// TAST (Let): eqMap_3_13 -> *Constructor_Data_Eq_Eq
eqMap_3_13 := &Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t22 bool
{
if (xs_4.Type == 9 && xs_4.IntVal == 324739070 && xs_4.UnsafePtr == nil) {
var __t20 bool
{
if (ys_5.Type == 9 && ys_5.IntVal == 324739070 && ys_5.UnsafePtr == nil) {
__t20 = true
goto end_branch_20
} else {

}
}
{
__t20 = false
}
end_branch_20:
__t22 = __t20
goto end_branch_22
} else {

}
}
{
if (xs_4.Type == 9 && xs_4.IntVal == 324739070 && xs_4.UnsafePtr != nil) {
var __t21 bool
{
if ((ys_5.Type == 9 && ys_5.IntVal == 324739070 && ys_5.UnsafePtr != nil)) && (((*Constructor_Data_Map_Internal_Node)(xs_4.UnsafePtr).V1) == ((*Constructor_Data_Map_Internal_Node)(ys_5.UnsafePtr).V1)) {
__t21 = (gopurs_runtime.Apply2(gopurs_runtime.Box(eqMapIter2_3_14.V0), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](xs_4), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](ys_5), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}).IntVal) != (0)
goto end_branch_21
} else {

}
}
{
__t21 = false
}
end_branch_21:
__t22 = __t21
goto end_branch_22
} else {

}
}
{
__t22 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_22:
return gopurs_runtime.Bool(__t22)
})
})}
_ = eqMap_3_13
// TAST (Let): eqSet1_2_11 -> *Constructor_Data_Eq_Eq
eqSet1_2_11 := &Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(eqMap_3_13.V0), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_4))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v1_5))}).IntVal) != (0))
})
})}
_ = eqSet1_2_11
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqSet1_2_11)}
}), gopurs_runtime.Func(func(s1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_23_4 gopurs_runtime.Value
_ = go__go_5_23_4
go__go_5_23_4 = gopurs_runtime.Func2(func(__local_var_6 gopurs_runtime.Value, __local_var_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t24 gopurs_runtime.Value
{
if (__local_var_6.Type == 9 && __local_var_6.IntVal == 324739070 && __local_var_6.UnsafePtr == nil) {
__t24 = __local_var_7
goto end_branch_24
} else {

}
}
{
if (__local_var_6.Type == 9 && __local_var_6.IntVal == 324739070 && __local_var_6.UnsafePtr != nil) {
__t24 = gopurs_runtime.UncurriedApp2(go__go_5_23_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_5_23_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V5)}, __local_var_7))})})
goto end_branch_24
} else {

}
}
{
__t24 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_24:
return __t24
})
var go__go_5_25_5 gopurs_runtime.Value
_ = go__go_5_25_5
go__go_5_25_5 = gopurs_runtime.Func2(func(__local_var_6 gopurs_runtime.Value, __local_var_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t26 gopurs_runtime.Value
{
if (__local_var_6.Type == 9 && __local_var_6.IntVal == 324739070 && __local_var_6.UnsafePtr == nil) {
__t26 = __local_var_7
goto end_branch_26
} else {

}
}
{
if (__local_var_6.Type == 9 && __local_var_6.IntVal == 324739070 && __local_var_6.UnsafePtr != nil) {
__t26 = gopurs_runtime.UncurriedApp2(go__go_5_25_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_5_25_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V5)}, __local_var_7))})})
goto end_branch_26
} else {

}
}
{
__t26 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_26:
return __t26
})
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(ordList_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_5_23_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_5_25_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s2_4))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}).IntVal)), UnsafePtr: nil}
})
})})}
}

func Call_Data_Set_NonEmpty_min(v_0_loop *Constructor_Data_Map_Internal_Node) gopurs_runtime.Value {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
return gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 -> *Constructor_Data_Maybe_Just
__local_var_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Map_Internal_findMin(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}))
_ = __local_var_2_1
var __t2 *Constructor_Data_Maybe_Just
{
if (__local_var_2_1 != nil) {
__t2 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordGet((__local_var_2_1).V0, "key")}
goto end_branch_2
} else {

}
}
{
__t2 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_2:
// TAST (Let): __local_var_2_0 -> *Constructor_Data_Maybe_Just
var __local_var_2_0 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)})
var __t3 gopurs_runtime.Value
{
if (__local_var_2_0 != nil) {
__t3 = (__local_var_2_0).V0
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}))
}

func Call_Data_Set_NonEmpty_member(dictOrd_0_loop *Constructor_Data_Ord_Ord, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_6 gopurs_runtime.Value
go__go_2_0_6 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_2_0_6:
for {
if false { continue go__go_2_0_6 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t3 bool
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t3 = false
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t2 bool
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)}
continue go__go_2_0_6
__t2 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)}
continue go__go_2_0_6
__t2 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_3:
return gopurs_runtime.Bool(__t3)
}
}()
})
return go__go_2_0_6
}

func Call_Data_Set_NonEmpty_max(v_0_loop *Constructor_Data_Map_Internal_Node) gopurs_runtime.Value {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
return gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 -> *Constructor_Data_Maybe_Just
__local_var_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Map_Internal_findMax(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}))
_ = __local_var_2_1
var __t2 *Constructor_Data_Maybe_Just
{
if (__local_var_2_1 != nil) {
__t2 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordGet((__local_var_2_1).V0, "key")}
goto end_branch_2
} else {

}
}
{
__t2 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_2:
// TAST (Let): __local_var_2_0 -> *Constructor_Data_Maybe_Just
var __local_var_2_0 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)})
var __t3 gopurs_runtime.Value
{
if (__local_var_2_0 != nil) {
__t3 = (__local_var_2_0).V0
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}))
}

func Call_Data_Set_NonEmpty_mapMaybe(dictOrd_0_loop *Constructor_Data_Ord_Ord, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_2_7 gopurs_runtime.Value
go__go_2_2_7 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_2_7:
for {
if false { continue go__go_2_2_7 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t10 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t10 = b_3
goto end_branch_10
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
// TAST (Let): __local_var_5_3 -> *Constructor_Data_Maybe_Just
__local_var_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0))
_ = __local_var_5_3
var __t9 *Constructor_Data_Map_Internal_Node
{
if (__local_var_5_3 == nil) {
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](b_3)
goto end_branch_9
} else {

}
}
{
if (__local_var_5_3 != nil) {
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := (__local_var_5_3).V0
_ = __local_var_6_4
var go__go_7_5_8 gopurs_runtime.Value
_ = go__go_7_5_8
go__go_7_5_8 = gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 *Constructor_Data_Map_Internal_Node
{
if (v1_8.Type == 9 && v1_8.IntVal == 324739070 && v1_8.UnsafePtr == nil) {
__t8 = &Constructor_Data_Map_Internal_Node{1, 1, 1, __local_var_6_4, Get_Data_Unit_unit(), (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_8
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 324739070 && v1_8.UnsafePtr != nil) {
// TAST (Let): v2_9_6 -> gopurs_runtime.Value
v2_9_6 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), __local_var_6_4, (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V2)
_ = v2_9_6
var __t7 *Constructor_Data_Map_Internal_Node
{
if (uint32(v2_9_6.IntVal) == 1527465420) {
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_7_5_8, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V5)}))
goto end_branch_7
} else {

}
}
{
if (uint32(v2_9_6.IntVal) == 380165415) {
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_7_5_8, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V5)})))}))
goto end_branch_7
} else {

}
}
{
if (uint32(v2_9_6.IntVal) == 902936544) {
__t7 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V1, __local_var_6_4, Get_Data_Unit_unit(), (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V5}
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_7:
__t8 = __t7
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t8)}
})
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_7_5_8, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](b_3))}))
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_9:
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t9)}
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_2_7
__t10 = gopurs_runtime.Value{}
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
}
}()
})
})
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(go__go_2_2_7, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
_ = __local_var_2_1
var go__go_3_12_9 gopurs_runtime.Value
go__go_3_12_9 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_12_9:
for {
if false { continue go__go_3_12_9 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t13 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t13 = v_4
goto end_branch_13
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_12_9
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_13:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t13)}
}
}()
})
})
// TAST (Let): __local_var_3_11 -> gopurs_runtime.Value
__local_var_3_11 := gopurs_runtime.Apply(go__go_3_12_9, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
_ = __local_var_3_11
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_1, gopurs_runtime.Apply(__local_var_3_11, x_4))
})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_14_10 gopurs_runtime.Value
_ = go__go_4_14_10
go__go_4_14_10 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t15 = __local_var_6
goto end_branch_15
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t15 = gopurs_runtime.UncurriedApp2(go__go_4_14_10, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_14_10, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
})
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_14_10, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
}

func Call_Data_Set_NonEmpty_go__map(dictOrd_0_loop *Constructor_Data_Ord_Ord, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_1_11 gopurs_runtime.Value
go__go_2_1_11 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_1_11:
for {
if false { continue go__go_2_1_11 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t7 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t7 = b_3
goto end_branch_7
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0)
_ = __local_var_5_2
var go__go_6_3_12 gopurs_runtime.Value
_ = go__go_6_3_12
go__go_6_3_12 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 *Constructor_Data_Map_Internal_Node
{
if (v1_7.Type == 9 && v1_7.IntVal == 324739070 && v1_7.UnsafePtr == nil) {
__t6 = &Constructor_Data_Map_Internal_Node{1, 1, 1, __local_var_5_2, Get_Data_Unit_unit(), (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_6
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 324739070 && v1_7.UnsafePtr != nil) {
// TAST (Let): v2_8_4 -> gopurs_runtime.Value
v2_8_4 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), __local_var_5_2, (*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V2)
_ = v2_8_4
var __t5 *Constructor_Data_Map_Internal_Node
{
if (uint32(v2_8_4.IntVal) == 1527465420) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_6_3_12, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V5)}))
goto end_branch_5
} else {

}
}
{
if (uint32(v2_8_4.IntVal) == 380165415) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_6_3_12, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V5)})))}))
goto end_branch_5
} else {

}
}
{
if (uint32(v2_8_4.IntVal) == 902936544) {
__t5 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V1, __local_var_5_2, Get_Data_Unit_unit(), (*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V5}
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t6)}
})
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_6_3_12, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](b_3))})))}
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_1_11
__t7 = gopurs_runtime.Value{}
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
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply(go__go_2_1_11, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_8_13 gopurs_runtime.Value
_ = go__go_4_8_13
go__go_4_8_13 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t9 = __local_var_6
goto end_branch_9
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t9 = gopurs_runtime.UncurriedApp2(go__go_4_8_13, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_8_13, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
})
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_8_13, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
}

func Call_Data_Set_NonEmpty_insert(dictOrd_0_loop *Constructor_Data_Ord_Ord, a_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Map_Internal_Node {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var v_2 *Constructor_Data_Map_Internal_Node = v_2_loop
_ = v_2
var go__go_3_0_14 gopurs_runtime.Value
_ = go__go_3_0_14
go__go_3_0_14 = gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr == nil) {
__t3 = &Constructor_Data_Map_Internal_Node{1, 1, 1, a_1, Get_Data_Unit_unit(), (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr != nil) {
// TAST (Let): v2_5_1 -> gopurs_runtime.Value
v2_5_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), a_1, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2)
_ = v2_5_1
var __t2 *Constructor_Data_Map_Internal_Node
{
if (uint32(v2_5_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_14, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5)}))
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_14, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5)})))}))
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 902936544) {
__t2 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V1, a_1, Get_Data_Unit_unit(), (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}
})
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_14, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_2)}))
}

func Call_Data_Set_NonEmpty_fromSet(s_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Maybe_Just {
var s_0 *Constructor_Data_Map_Internal_Node = s_0_loop
_ = s_0
var __t1 *Constructor_Data_Maybe_Just
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
__t1 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s_0)}}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)})
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
var __t2 *Constructor_Data_Maybe_Just
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
__t2 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_2
} else {

}
}
{
__t2 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__local_var_3_0)}}
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)})
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
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_15 gopurs_runtime.Value
_ = go__go_4_1_15
go__go_4_1_15 = gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_Map_Internal_Node
{
if (v1_5.Type == 9 && v1_5.IntVal == 324739070 && v1_5.UnsafePtr == nil) {
__t4 = &Constructor_Data_Map_Internal_Node{1, 1, 1, a_3, Get_Data_Unit_unit(), (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_4
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 324739070 && v1_5.UnsafePtr != nil) {
// TAST (Let): v2_6_2 -> gopurs_runtime.Value
v2_6_2 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_1.V1), a_3, (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V2)
_ = v2_6_2
var __t3 *Constructor_Data_Map_Internal_Node
{
if (uint32(v2_6_2.IntVal) == 1527465420) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_4_1_15, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V5)}))
goto end_branch_3
} else {

}
}
{
if (uint32(v2_6_2.IntVal) == 380165415) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_4_1_15, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V5)})))}))
goto end_branch_3
} else {

}
}
{
if (uint32(v2_6_2.IntVal) == 902936544) {
__t3 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V1, a_3, Get_Data_Unit_unit(), (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V5}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t4)}
})
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_4_1_15, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_2))})))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_5 -> gopurs_runtime.Value
__local_var_4_5 := gopurs_runtime.Apply(__local_var_2_0, x_3)
_ = __local_var_4_5
var __t8 *Constructor_Data_Maybe_Just
{
var __t7 bool
{
var __t_tag_6 *Constructor_Data_Map_Internal_Node = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_4_5)
if (__t_tag_6 == nil) {
__t7 = true
goto end_branch_7
} else {

}
}
{
__t7 = false
}
end_branch_7:
if __t7 {
__t8 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_8
} else {

}
}
{
__t8 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](__local_var_4_5))}}
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t8)}
})
}

func Call_Data_Set_NonEmpty_filter(dictOrd_0_loop *Constructor_Data_Ord_Ord, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_0_20 gopurs_runtime.Value
_ = go__go_2_0_20
go__go_2_0_20 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Map_Internal_Node
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t2 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_2
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
var __t1 *Constructor_Data_Map_Internal_Node
{
if (gopurs_runtime.Apply(f_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2).IntVal) != (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_20, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_20, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)})))}))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_20, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_20, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)})))}))
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t2)}
})
return go__go_2_0_20
}

func Call_Data_Set_NonEmpty_eqNonEmptySet(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var go__go_1_2_21 gopurs_runtime.Value
go__go_1_2_21 = gopurs_runtime.Func(func(a_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_2_loop gopurs_runtime.Value = a_2_loop_val
var b_3_loop gopurs_runtime.Value = b_3_loop_val
go__go_1_2_21:
for {
if false { continue go__go_1_2_21 }
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
// TAST (Let): v_4_3 -> *Constructor_Data_Map_Internal_IterNext
v_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_2))
_ = v_4_3
var __t6 bool
{
if (v_4_3 != nil) {
// TAST (Let): v2_5_4 -> *Constructor_Data_Map_Internal_IterNext
v2_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_3))
_ = v2_5_4
var __t5 bool
{
if ((v2_5_4 != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (v_4_3).V0, (v2_5_4).V0).IntVal) != (0)) {
a_2_loop = (v_4_3).V2
b_3_loop = (v2_5_4).V2
continue go__go_1_2_21
__t5 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_5
} else {

}
}
{
__t5 = false
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
if (v_4_3 == nil) {
__t6 = true
goto end_branch_6
} else {

}
}
{
__t6 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_6:
return gopurs_runtime.Bool(__t6)
}
}()
})
})
// TAST (Let): eqMapIter2_1_1 -> *Constructor_Data_Eq_Eq
eqMapIter2_1_1 := &Constructor_Data_Eq_Eq{1, go__go_1_2_21}
_ = eqMapIter2_1_1
// TAST (Let): eqMap_1_0 -> *Constructor_Data_Eq_Eq
eqMap_1_0 := &Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 bool
{
if (xs_2.Type == 9 && xs_2.IntVal == 324739070 && xs_2.UnsafePtr == nil) {
var __t7 bool
{
if (ys_3.Type == 9 && ys_3.IntVal == 324739070 && ys_3.UnsafePtr == nil) {
__t7 = true
goto end_branch_7
} else {

}
}
{
__t7 = false
}
end_branch_7:
__t9 = __t7
goto end_branch_9
} else {

}
}
{
if (xs_2.Type == 9 && xs_2.IntVal == 324739070 && xs_2.UnsafePtr != nil) {
var __t8 bool
{
if ((ys_3.Type == 9 && ys_3.IntVal == 324739070 && ys_3.UnsafePtr != nil)) && (((*Constructor_Data_Map_Internal_Node)(xs_2.UnsafePtr).V1) == ((*Constructor_Data_Map_Internal_Node)(ys_3.UnsafePtr).V1)) {
__t8 = (gopurs_runtime.Apply2(gopurs_runtime.Box(eqMapIter2_1_1.V0), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](xs_2), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](ys_3), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}).IntVal) != (0)
goto end_branch_8
} else {

}
}
{
__t8 = false
}
end_branch_8:
__t9 = __t8
goto end_branch_9
} else {

}
}
{
__t9 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_9:
return gopurs_runtime.Bool(__t9)
})
})}
_ = eqMap_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(eqMap_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v1_3))}).IntVal) != (0))
})
})})}
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
var __t2 *Constructor_Data_Maybe_Just
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
__t2 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_2
} else {

}
}
{
__t2 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__local_var_3_0)}}
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)})
}

func Call_Data_Set_NonEmpty_delete(dictOrd_0_loop *Constructor_Data_Ord_Ord, a_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Maybe_Just {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var v_2 *Constructor_Data_Map_Internal_Node = v_2_loop
_ = v_2
var go__go_3_1_22 gopurs_runtime.Value
_ = go__go_3_1_22
go__go_3_1_22 = gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_Map_Internal_Node
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr == nil) {
__t4 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_4
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 324739070 && v_4.UnsafePtr != nil) {
// TAST (Let): v1_5_2 -> gopurs_runtime.Value
v1_5_2 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), a_1, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V2)
_ = v1_5_2
var __t3 *Constructor_Data_Map_Internal_Node
{
if (uint32(v1_5_2.IntVal) == 1527465420) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_1_22, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V5)}))
goto end_branch_3
} else {

}
}
{
if (uint32(v1_5_2.IntVal) == 380165415) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_1_22, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V5)})))}))
goto end_branch_3
} else {

}
}
{
if (uint32(v1_5_2.IntVal) == 902936544) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_4.UnsafePtr).V5)}))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t4)}
})
// TAST (Let): __local_var_3_0 -> *Constructor_Data_Map_Internal_Node
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_1_22, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_2)}))
_ = __local_var_3_0
var __t6 *Constructor_Data_Maybe_Just
{
var __t5 bool
{
if (__local_var_3_0 == nil) {
__t5 = true
goto end_branch_5
} else {

}
}
{
__t5 = false
}
end_branch_5:
if __t5 {
__t6 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_6
} else {

}
}
{
__t6 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__local_var_3_0)}}
}
end_branch_6:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t6)})
}

func Call_Data_Set_NonEmpty_cons(dictOrd_0_loop *Constructor_Data_Ord_Ord, a_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Map_Internal_Node {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var v_2 *Constructor_Data_Map_Internal_Node = v_2_loop
_ = v_2
var go__go_3_0_23 gopurs_runtime.Value
_ = go__go_3_0_23
go__go_3_0_23 = gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr == nil) {
__t3 = &Constructor_Data_Map_Internal_Node{1, 1, 1, a_1, Get_Data_Unit_unit(), (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr != nil) {
// TAST (Let): v2_5_1 -> gopurs_runtime.Value
v2_5_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), a_1, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2)
_ = v2_5_1
var __t2 *Constructor_Data_Map_Internal_Node
{
if (uint32(v2_5_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_23, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5)}))
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_23, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5)})))}))
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 902936544) {
__t2 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V1, a_1, Get_Data_Unit_unit(), (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t3)}
})
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_23, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_2)}))
}

func Call_Data_Set_NonEmpty_fromSet__3199996154(s_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Maybe_Just {
var s_0 *Constructor_Data_Map_Internal_Node = s_0_loop
_ = s_0
var __t1 *Constructor_Data_Maybe_Just
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
__t1 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s_0)}}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)})
}

func Call_Data_Set_NonEmpty_fromSet__1805959329(s_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Maybe_Just {
var s_0 *Constructor_Data_Map_Internal_Node = s_0_loop
_ = s_0
var __t1 *Constructor_Data_Maybe_Just
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
__t1 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s_0)}}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)})
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
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0, gopurs_runtime.Apply(stepNext_1_0, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1)})}
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



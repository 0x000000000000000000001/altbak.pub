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
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Set_NonEmpty_NonEmptySet(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_0_box)))}
})
	})
	return cache_Data_Set_NonEmpty_NonEmptySet
}

var cache_Data_Set_NonEmpty_unionSet gopurs_runtime.Value
var once_Data_Set_NonEmpty_unionSet sync.Once
func Get_Data_Set_NonEmpty_unionSet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_unionSet.Do(func() {
		cache_Data_Set_NonEmpty_unionSet = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_unionSet(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Set_NonEmpty_unionSet
}

var cache_Data_Set_NonEmpty_toUnfoldable1 gopurs_runtime.Value
var once_Data_Set_NonEmpty_toUnfoldable1 sync.Once
func Get_Data_Set_NonEmpty_toUnfoldable1() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_toUnfoldable1.Do(func() {
		cache_Data_Set_NonEmpty_toUnfoldable1 = gopurs_runtime.Func(func(dictUnfoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_toUnfoldable1(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_0_box))
})
	})
	return cache_Data_Set_NonEmpty_toUnfoldable1
}

var cache_Data_Set_NonEmpty_toUnfoldable11 gopurs_runtime.Value
var once_Data_Set_NonEmpty_toUnfoldable11 sync.Once
func Get_Data_Set_NonEmpty_toUnfoldable11() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_toUnfoldable11.Do(func() {
		cache_Data_Set_NonEmpty_toUnfoldable11 = Call_Data_Set_NonEmpty_toUnfoldable1(Rebox_Data_Set_NonEmpty_2435023311_2187088110(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_unfoldable1NonEmptyList())))
	})
	return cache_Data_Set_NonEmpty_toUnfoldable11
}

var cache_Data_Set_NonEmpty_toUnfoldable gopurs_runtime.Value
var once_Data_Set_NonEmpty_toUnfoldable sync.Once
func Get_Data_Set_NonEmpty_toUnfoldable() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_toUnfoldable.Do(func() {
		cache_Data_Set_NonEmpty_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
})
	})
	return cache_Data_Set_NonEmpty_toUnfoldable
}

var cache_Data_Set_NonEmpty_toSet gopurs_runtime.Value
var once_Data_Set_NonEmpty_toSet sync.Once
func Get_Data_Set_NonEmpty_toSet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_toSet.Do(func() {
		cache_Data_Set_NonEmpty_toSet = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Set_NonEmpty_toSet(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_Data_Set_NonEmpty_toSet
}

var cache_Data_Set_NonEmpty_subset gopurs_runtime.Value
var once_Data_Set_NonEmpty_subset sync.Once
func Get_Data_Set_NonEmpty_subset() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_subset.Do(func() {
		cache_Data_Set_NonEmpty_subset = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_subset(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Set_NonEmpty_subset
}

var cache_Data_Set_NonEmpty_size gopurs_runtime.Value
var once_Data_Set_NonEmpty_size sync.Once
func Get_Data_Set_NonEmpty_size() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_size.Do(func() {
		cache_Data_Set_NonEmpty_size = gopurs_runtime.Int(Get_Data_Map_Internal_size().IntVal)
	})
	return cache_Data_Set_NonEmpty_size
}

var cache_Data_Set_NonEmpty_singleton gopurs_runtime.Value
var once_Data_Set_NonEmpty_singleton sync.Once
func Get_Data_Set_NonEmpty_singleton() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_singleton.Do(func() {
		cache_Data_Set_NonEmpty_singleton = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Set_singleton()))}
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
return Call_Data_Set_NonEmpty_properSubset(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
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
		cache_Data_Set_NonEmpty_ord1NonEmptySet = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_849406934_3985601471(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Set_ord1Set())))}
	})
	return cache_Data_Set_NonEmpty_ord1NonEmptySet
}

var cache_Data_Set_NonEmpty_min gopurs_runtime.Value
var once_Data_Set_NonEmpty_min sync.Once
func Get_Data_Set_NonEmpty_min() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_min.Do(func() {
		cache_Data_Set_NonEmpty_min = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_min(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_Data_Set_NonEmpty_min
}

var cache_Data_Set_NonEmpty_member gopurs_runtime.Value
var once_Data_Set_NonEmpty_member sync.Once
func Get_Data_Set_NonEmpty_member() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_member.Do(func() {
		cache_Data_Set_NonEmpty_member = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_member(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
})
	})
	return cache_Data_Set_NonEmpty_member
}

var cache_Data_Set_NonEmpty_max gopurs_runtime.Value
var once_Data_Set_NonEmpty_max sync.Once
func Get_Data_Set_NonEmpty_max() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_max.Do(func() {
		cache_Data_Set_NonEmpty_max = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_max(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_Data_Set_NonEmpty_max
}

var cache_Data_Set_NonEmpty_mapMaybe gopurs_runtime.Value
var once_Data_Set_NonEmpty_mapMaybe sync.Once
func Get_Data_Set_NonEmpty_mapMaybe() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_mapMaybe.Do(func() {
		cache_Data_Set_NonEmpty_mapMaybe = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_mapMaybe(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Set_NonEmpty_mapMaybe
}

var cache_Data_Set_NonEmpty_go__map gopurs_runtime.Value
var once_Data_Set_NonEmpty_go__map sync.Once
func Get_Data_Set_NonEmpty_go__map() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_go__map.Do(func() {
		cache_Data_Set_NonEmpty_go__map = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Set_NonEmpty_go__map
}

var cache_Data_Set_NonEmpty_insert gopurs_runtime.Value
var once_Data_Set_NonEmpty_insert sync.Once
func Get_Data_Set_NonEmpty_insert() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_insert.Do(func() {
		cache_Data_Set_NonEmpty_insert = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_insert(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Set_NonEmpty_insert
}

var cache_Data_Set_NonEmpty_fromSet gopurs_runtime.Value
var once_Data_Set_NonEmpty_fromSet sync.Once
func Get_Data_Set_NonEmpty_fromSet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_fromSet.Do(func() {
		cache_Data_Set_NonEmpty_fromSet = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Set_NonEmpty_fromSet(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](s_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Set_NonEmpty_fromSet
}

var cache_Data_Set_NonEmpty_intersection gopurs_runtime.Value
var once_Data_Set_NonEmpty_intersection sync.Once
func Get_Data_Set_NonEmpty_intersection() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_intersection.Do(func() {
		cache_Data_Set_NonEmpty_intersection = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Set_NonEmpty_intersection(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v1_2_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Set_NonEmpty_intersection
}

var cache_Data_Set_NonEmpty_fromFoldable1 gopurs_runtime.Value
var once_Data_Set_NonEmpty_fromFoldable1 sync.Once
func Get_Data_Set_NonEmpty_fromFoldable1() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_fromFoldable1.Do(func() {
		cache_Data_Set_NonEmpty_fromFoldable1 = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_fromFoldable1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_1_box))
})
	})
	return cache_Data_Set_NonEmpty_fromFoldable1
}

var cache_Data_Set_NonEmpty_fromFoldable gopurs_runtime.Value
var once_Data_Set_NonEmpty_fromFoldable sync.Once
func Get_Data_Set_NonEmpty_fromFoldable() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_fromFoldable.Do(func() {
		cache_Data_Set_NonEmpty_fromFoldable = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Set_NonEmpty_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_1_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Set_NonEmpty_fromFoldable
}

var cache_Data_Set_NonEmpty_foldableNonEmptySet gopurs_runtime.Value
var once_Data_Set_NonEmpty_foldableNonEmptySet sync.Once
func Get_Data_Set_NonEmpty_foldableNonEmptySet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_foldableNonEmptySet.Do(func() {
		cache_Data_Set_NonEmpty_foldableNonEmptySet = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_3596835815_1680800814(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Set_foldableSet())))}
	})
	return cache_Data_Set_NonEmpty_foldableNonEmptySet
}

var cache_Data_Set_NonEmpty_foldable1NonEmptySet gopurs_runtime.Value
var once_Data_Set_NonEmpty_foldable1NonEmptySet sync.Once
func Get_Data_Set_NonEmpty_foldable1NonEmptySet() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_foldable1NonEmptySet.Do(func() {
		cache_Data_Set_NonEmpty_foldable1NonEmptySet = gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_68075396_4151366573((&Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_3596835815_1680800814(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Set_foldableSet())))}
}), gopurs_runtime.Func2(func(dictSemigroup_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_foldable1NonEmptyList()).V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_0))}, f_1)
_ = __local_var_2_0
// TAST (Let): __local_var_3_1 shape=App(Var) bindingType=(Func [(ADT ["Data","Map","Internal","Map"] [(TypeVar a), Unit])] (ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (TypeVar a)]))
__local_var_3_1 := Call_Data_Set_NonEmpty_toUnfoldable1(Rebox_Data_Set_NonEmpty_2435023311_2187088110(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_unfoldable1NonEmptyList())))
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_2 shape=App(Other) bindingType=Any
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_foldable1NonEmptyList()).V2), f_0)
_ = __local_var_1_2
// TAST (Let): __local_var_2_3 shape=App(Var) bindingType=(Func [(ADT ["Data","Map","Internal","Map"] [(TypeVar a), Unit])] (ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (TypeVar a)]))
__local_var_2_3 := Call_Data_Set_NonEmpty_toUnfoldable1(Rebox_Data_Set_NonEmpty_2435023311_2187088110(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_unfoldable1NonEmptyList())))
_ = __local_var_2_3
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_2, gopurs_runtime.Apply(__local_var_2_3, x_3))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_4 shape=App(Other) bindingType=Any
__local_var_1_4 := gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_foldable1NonEmptyList()).V3), f_0)
_ = __local_var_1_4
// TAST (Let): __local_var_2_5 shape=App(Var) bindingType=(Func [(ADT ["Data","Map","Internal","Map"] [(TypeVar a), Unit])] (ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (TypeVar a)]))
__local_var_2_5 := Call_Data_Set_NonEmpty_toUnfoldable1(Rebox_Data_Set_NonEmpty_2435023311_2187088110(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_unfoldable1NonEmptyList())))
_ = __local_var_2_5
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_4, gopurs_runtime.Apply(__local_var_2_5, x_3))
})
})})))}
	})
	return cache_Data_Set_NonEmpty_foldable1NonEmptySet
}

var cache_Data_Set_NonEmpty_filter gopurs_runtime.Value
var once_Data_Set_NonEmpty_filter sync.Once
func Get_Data_Set_NonEmpty_filter() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_filter.Do(func() {
		cache_Data_Set_NonEmpty_filter = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_filter(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
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
		cache_Data_Set_NonEmpty_eq1NonEmptySet = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_3691144502_1766074591(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Set_eq1Set())))}
	})
	return cache_Data_Set_NonEmpty_eq1NonEmptySet
}

var cache_Data_Set_NonEmpty_difference gopurs_runtime.Value
var once_Data_Set_NonEmpty_difference sync.Once
func Get_Data_Set_NonEmpty_difference() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_difference.Do(func() {
		cache_Data_Set_NonEmpty_difference = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Set_NonEmpty_difference(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v1_2_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Set_NonEmpty_difference
}

var cache_Data_Set_NonEmpty_go__delete gopurs_runtime.Value
var once_Data_Set_NonEmpty_go__delete sync.Once
func Get_Data_Set_NonEmpty_go__delete() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_go__delete.Do(func() {
		cache_Data_Set_NonEmpty_go__delete = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Set_NonEmpty_go__delete(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), a_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_2_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Set_NonEmpty_go__delete
}

var cache_Data_Set_NonEmpty_cons gopurs_runtime.Value
var once_Data_Set_NonEmpty_cons sync.Once
func Get_Data_Set_NonEmpty_cons() gopurs_runtime.Value {
	once_Data_Set_NonEmpty_cons.Do(func() {
		cache_Data_Set_NonEmpty_cons = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_NonEmpty_cons(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Set_NonEmpty_cons
}

func Call_Data_Set_NonEmpty_NonEmptySet(x_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Set_NonEmpty_unionSet(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 shape=Other bindingType=(Func [(TypeVar k), (TypeVar k)] (ADT ["Data","Ordering","Ordering"] []))
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_0, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
}

func Call_Data_Set_NonEmpty_toUnfoldable1(dictUnfoldable1_0_loop *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable1_0 *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
// TAST (Let): stepNext_1_0 shape=App(Var) bindingType=(TypeVar r)
stepNext_1_0 := gopurs_runtime.Apply3(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapL(), gopurs_runtime.Func3(func(k_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value, __local_var_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_4010058633_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, k_1, __local_var_3}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_4010058633_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}))
_ = stepNext_1_0
// TAST (Let): __local_var_2_1 shape=App(Other) bindingType=Any
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable1_0.V0), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply(stepNext_1_0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}))
_ = __local_var_2_1
// TAST (Let): __local_var_3_3 shape=App(Var) bindingType=Any
__local_var_3_3 := gopurs_runtime.Apply3(Get_Data_Map_Internal_stepWith(), Get_Data_Map_Internal_iterMapL(), gopurs_runtime.Func3(func(k_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{k_3, __local_var_5}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.Str("toUnfoldable1: impossible"))))}
}))
_ = __local_var_3_3
// TAST (Let): __local_var_3_2 shape=Let(Abs(App(Other))) bindingType=(Func [(ADT ["Data","Map","Internal","Map"] [(TypeVar a), Unit])] (ADT ["Data","Tuple","Tuple"] [(TypeVar a), (ADT ["Data","Map","Internal","MapIter"] [(TypeVar a), Unit])]))
__local_var_3_2 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_4), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}}))})
})
_ = __local_var_3_2
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_1, gopurs_runtime.Apply(__local_var_3_2, x_4))
})
}

func Call_Data_Set_NonEmpty_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
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
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_2351501316_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
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
var go__go_3_4_0 gopurs_runtime.Value
_ = go__go_3_4_0
// FALLBACK TCO: isLoop=false len=1
go__go_3_4_0 = gopurs_runtime.Func2(func(m_prime__4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
var __t_tag_5 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__4)
if (__t_tag_5 == nil) {
__t7 = __local_var_5
goto end_branch_7
} else {

}
}
{
var __t_tag_6 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__4)
if (__t_tag_6 != nil) {
__t7 = gopurs_runtime.UncurriedApp2(go__go_3_4_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__4.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_3_4_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__4.UnsafePtr).V5)}, __local_var_5))}))})
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
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_3_4_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})))})
})
}

func Call_Data_Set_NonEmpty_toSet(v_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var v_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Set_NonEmpty_subset(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Set_subset(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_Data_Set_NonEmpty_showNonEmptySet(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): showNonEmptyArray_1_0 shape=App(Var) bindingType=(ADT ["Data","Show","Show"] [(Array (TypeVar a))])
showNonEmptyArray_1_0 := Rebox_Data_Set_NonEmpty_1386611502_1356436936(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Array_NonEmpty_Internal_showNonEmptyArray(), dictShow_0)))
_ = showNonEmptyArray_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_2638796135_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(fromFoldable1 ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showNonEmptyArray_1_0.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Call_Data_Set_NonEmpty_toUnfoldable1(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Get_Data_Unfoldable1_unfoldable1Array())), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](s_2))}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).StrVal())) + (")"))
})})))}
}

func Call_Data_Set_NonEmpty_semigroupNonEmptySet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 shape=Other bindingType=(Func [(TypeVar k), (TypeVar k)] (ADT ["Data","Ordering","Ordering"] []))
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_3854229351_4179793454((&Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_0, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})})))}
}

func Call_Data_Set_NonEmpty_properSubset(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Set_properSubset(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_Data_Set_NonEmpty_ordNonEmptySet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_1910448679_4177771502(Rebox_Data_Set_NonEmpty_4177771502_1910448679(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Set_ordSet(), dictOrd_0)))))}
}

func Call_Data_Set_NonEmpty_min(v_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
// TAST (Let): __local_var_1_1 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [key: (TypeVar a), value: Unit] Any))])
__local_var_1_1 := Rebox_Data_Set_NonEmpty_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_findMin(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)})))
_ = __local_var_1_1
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_1_1 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(__local_var_1_1).V0.key, true}
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
// TAST (Let): __local_var_1_0 shape=Let(Branch(Other, def=Other)) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
__local_var_1_0 := __t2
_ = __local_var_1_0
var __t3 gopurs_runtime.Value
{
if (__local_var_1_0 != nil) {
__t3 = (__local_var_1_0).V0
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

func Call_Data_Set_NonEmpty_member(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var Call_local_Data_Set_NonEmpty_go__go_2_0_1 func(*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) bool
_ = Call_local_Data_Set_NonEmpty_go__go_2_0_1
var go__go_2_0_1 gopurs_runtime.Value
_ = go__go_2_0_1
Call_local_Data_Set_NonEmpty_go__go_2_0_1 = func(v_3_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) bool {
go__go_2_0_1:
for {
if false { continue go__go_2_0_1 }
var v_3 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_3_loop
_ = v_3
var __t3 bool
{
if (v_3 == nil) {
__t3 = false
goto end_branch_3
} else {

}
}
{
if (v_3 != nil) {
// TAST (Let): v1_4_1 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v1_4_1 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (v_3).V2).IntVal)
_ = v1_4_1
var __t2 bool
{
if (v1_4_1 == 1527465420) {
v_3_loop = (v_3).V4
continue go__go_2_0_1
__t2 = func() bool { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
if (v1_4_1 == 380165415) {
v_3_loop = (v_3).V5
continue go__go_2_0_1
__t2 = func() bool { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
if (v1_4_1 == 902936544) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = func() bool { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = func() bool { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}
go__go_2_0_1 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_local_Data_Set_NonEmpty_go__go_2_0_1(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_3_loop_val)))
})
return go__go_2_0_1
}

func Call_Data_Set_NonEmpty_max(v_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
// TAST (Let): __local_var_1_1 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [key: (TypeVar a), value: Unit] Any))])
__local_var_1_1 := Rebox_Data_Set_NonEmpty_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_findMax(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)})))
_ = __local_var_1_1
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_1_1 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(__local_var_1_1).V0.key, true}
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
// TAST (Let): __local_var_1_0 shape=Let(Branch(Other, def=Other)) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
__local_var_1_0 := __t2
_ = __local_var_1_0
var __t3 gopurs_runtime.Value
{
if (__local_var_1_0 != nil) {
__t3 = (__local_var_1_0).V0
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

func Call_Data_Set_NonEmpty_mapMaybe(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Set_mapMaybe(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_Data_Set_NonEmpty_go__map(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Set_go__map(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_Data_Set_NonEmpty_insert(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Set_insert(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_Data_Set_NonEmpty_fromSet(s_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var s_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = s_0_loop
_ = s_0
var __t0 *Constructor_Data_Maybe_Just[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]
{
if (s_0 == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]](func() gopurs_runtime.Value {
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
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s_0)}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_2503784013_3094389156(__t0))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Set_NonEmpty_intersection(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], v_1_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value], v1_2_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var v_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
_ = v_1
var v1_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v1_2_loop
_ = v1_2
// TAST (Let): __local_var_3_0 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar c)])
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), gopurs_runtime.Box(dictOrd_0.V1), Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v1_2)}))
_ = __local_var_3_0
var __t1 *Constructor_Data_Maybe_Just[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]
{
if (__local_var_3_0 == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
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
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__local_var_3_0)}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_1:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_2503784013_3094389156(__t1))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Set_NonEmpty_fromFoldable1(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value], dictOrd_1_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictOrd_1 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_1_loop
_ = dictOrd_1
// TAST (Let): compare_2_0 shape=Other bindingType=(Func [(TypeVar k), (TypeVar k)] (ADT ["Data","Ordering","Ordering"] []))
compare_2_0 := gopurs_runtime.Box(dictOrd_1.V1)
_ = compare_2_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_3854229351_4179793454((&Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(m1_3 gopurs_runtime.Value, m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_2_0, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_4))})))}
})})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Set_singleton()))})
}

func Call_Data_Set_NonEmpty_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], dictOrd_1_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictOrd_1 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_1_loop
_ = dictOrd_1
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar f) [(TypeVar a)])] (ADT ["Data","Map","Internal","Map"] [(TypeVar a), Unit]))
__local_var_2_0 := gopurs_runtime.Apply2(Get_Data_Set_fromFoldable(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_1)})
_ = __local_var_2_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=Any
__local_var_4_1 := gopurs_runtime.Apply(__local_var_2_0, x_3)
_ = __local_var_4_1
var __t2 *Constructor_Data_Maybe_Just[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 324739070 && __local_var_4_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]](func() gopurs_runtime.Value {
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
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_4_1))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_2503784013_3094389156(__t2))}
})
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Set_NonEmpty_filter(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Set_filter(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_Data_Set_NonEmpty_eqNonEmptySet(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_1444241223_3790796878(Rebox_Data_Set_NonEmpty_3790796878_1444241223(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Set_eqSet(), dictEq_0)))))}
}

func Call_Data_Set_NonEmpty_difference(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], v_1_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value], v1_2_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var v_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
_ = v_1
var v1_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v1_2_loop
_ = v1_2
// TAST (Let): __local_var_3_0 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Map","Internal","Map"] [(TypeVar k), (TypeVar v)])
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeDifference(), gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v1_2)}))
_ = __local_var_3_0
var __t1 *Constructor_Data_Maybe_Just[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]
{
if (__local_var_3_0 == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
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
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__local_var_3_0)}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_1:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_2503784013_3094389156(__t1))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Set_NonEmpty_go__delete(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var v_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_2_loop
_ = v_2
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=Any
__local_var_3_0 := gopurs_runtime.Apply3(Get_Data_Map_Internal_go__delete(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, a_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_2)})
_ = __local_var_3_0
var __t1 *Constructor_Data_Maybe_Just[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 324739070 && __local_var_3_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
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
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__local_var_3_0))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_1:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_NonEmpty_2503784013_3094389156(__t1))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Set_NonEmpty_cons(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Set_insert(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Rebox_Data_Set_NonEmpty_1386611502_1356436936(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[[]gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[[]gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Set_NonEmpty_1444241223_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Set_NonEmpty_1910448679_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Set_NonEmpty_2351501316_138441832(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}

func Rebox_Data_Set_NonEmpty_2435023311_2187088110(in *Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Set_NonEmpty_2503784013_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_Set_NonEmpty_2638796135_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Set_NonEmpty_3094389156_2758605161(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}]{}
		out.V0 = func() struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
} {
					orig := in.V0
					_ = orig
					clone := struct{
	key gopurs_runtime.Value
	value gopurs_runtime.Value
}{}
					clone.key = gopurs_runtime.RecordGet(orig, "key")
					clone.value = gopurs_runtime.RecordGet(orig, "value")
					return clone
				}()
	return out
}

func Rebox_Data_Set_NonEmpty_3596835815_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Set_NonEmpty_3691144502_1766074591(in *Constructor_Data_Eq_Eq1[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Set_NonEmpty_3790796878_1444241223(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Set_NonEmpty_3854229351_4179793454(in *Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Set_NonEmpty_4010058633_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_Set_NonEmpty_4177771502_1910448679(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Set_NonEmpty_68075396_4151366573(in *Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Set_NonEmpty_849406934_3985601471(in *Constructor_Data_Ord_Ord1[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}



package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Set_identity gopurs_runtime.Value
var once_Data_Set_identity sync.Once
func Get_Data_Set_identity() gopurs_runtime.Value {
	once_Data_Set_identity.Do(func() {
		cache_Data_Set_identity = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Data_Set_identity
}

var cache_Data_Set_Set gopurs_runtime.Value
var once_Data_Set_Set sync.Once
func Get_Data_Set_Set() gopurs_runtime.Value {
	once_Data_Set_Set.Do(func() {
		cache_Data_Set_Set = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Set_Set(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_0_box)))}
})
	})
	return cache_Data_Set_Set
}

var cache_Data_Set_union gopurs_runtime.Value
var once_Data_Set_union sync.Once
func Get_Data_Set_union() gopurs_runtime.Value {
	once_Data_Set_union.Do(func() {
		cache_Data_Set_union = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_union(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Set_union
}

var cache_Data_Set_toggle gopurs_runtime.Value
var once_Data_Set_toggle sync.Once
func Get_Data_Set_toggle() gopurs_runtime.Value {
	once_Data_Set_toggle.Do(func() {
		cache_Data_Set_toggle = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Set_toggle(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), a_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_2_box)))}
})
	})
	return cache_Data_Set_toggle
}

var cache_Data_Set_toMap gopurs_runtime.Value
var once_Data_Set_toMap sync.Once
func Get_Data_Set_toMap() gopurs_runtime.Value {
	once_Data_Set_toMap.Do(func() {
		cache_Data_Set_toMap = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Set_toMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_Data_Set_toMap
}

var cache_Data_Set_toList gopurs_runtime.Value
var once_Data_Set_toList sync.Once
func Get_Data_Set_toList() gopurs_runtime.Value {
	once_Data_Set_toList.Do(func() {
		cache_Data_Set_toList = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Set_toList(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_Data_Set_toList
}

var cache_Data_Set_toUnfoldable gopurs_runtime.Value
var once_Data_Set_toUnfoldable sync.Once
func Get_Data_Set_toUnfoldable() gopurs_runtime.Value {
	once_Data_Set_toUnfoldable.Do(func() {
		cache_Data_Set_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
})
	})
	return cache_Data_Set_toUnfoldable
}

var cache_Data_Set_size gopurs_runtime.Value
var once_Data_Set_size sync.Once
func Get_Data_Set_size() gopurs_runtime.Value {
	once_Data_Set_size.Do(func() {
		cache_Data_Set_size = Get_Data_Map_Internal_size()
	})
	return cache_Data_Set_size
}

var cache_Data_Set_singleton gopurs_runtime.Value
var once_Data_Set_singleton sync.Once
func Get_Data_Set_singleton() gopurs_runtime.Value {
	once_Data_Set_singleton.Do(func() {
		cache_Data_Set_singleton = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Set_singleton(a_0_box))}
})
	})
	return cache_Data_Set_singleton
}

var cache_Data_Set_showSet gopurs_runtime.Value
var once_Data_Set_showSet sync.Once
func Get_Data_Set_showSet() gopurs_runtime.Value {
	once_Data_Set_showSet.Do(func() {
		cache_Data_Set_showSet = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_showSet(dictShow_0_box)
})
	})
	return cache_Data_Set_showSet
}

var cache_Data_Set_semigroupSet gopurs_runtime.Value
var once_Data_Set_semigroupSet sync.Once
func Get_Data_Set_semigroupSet() gopurs_runtime.Value {
	once_Data_Set_semigroupSet.Do(func() {
		cache_Data_Set_semigroupSet = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_semigroupSet(dictOrd_0_box)
})
	})
	return cache_Data_Set_semigroupSet
}

var cache_Data_Set_member gopurs_runtime.Value
var once_Data_Set_member sync.Once
func Get_Data_Set_member() gopurs_runtime.Value {
	once_Data_Set_member.Do(func() {
		cache_Data_Set_member = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_member(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Set_member
}

var cache_Data_Set_isEmpty gopurs_runtime.Value
var once_Data_Set_isEmpty sync.Once
func Get_Data_Set_isEmpty() gopurs_runtime.Value {
	once_Data_Set_isEmpty.Do(func() {
		cache_Data_Set_isEmpty = Get_Data_Map_Internal_isEmpty()
	})
	return cache_Data_Set_isEmpty
}

var cache_Data_Set_intersection gopurs_runtime.Value
var once_Data_Set_intersection sync.Once
func Get_Data_Set_intersection() gopurs_runtime.Value {
	once_Data_Set_intersection.Do(func() {
		cache_Data_Set_intersection = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_intersection(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Set_intersection
}

var cache_Data_Set_insert gopurs_runtime.Value
var once_Data_Set_insert sync.Once
func Get_Data_Set_insert() gopurs_runtime.Value {
	once_Data_Set_insert.Do(func() {
		cache_Data_Set_insert = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Set_insert(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), a_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_2_box)))}
})
	})
	return cache_Data_Set_insert
}

var cache_Data_Set_fromMap gopurs_runtime.Value
var once_Data_Set_fromMap sync.Once
func Get_Data_Set_fromMap() gopurs_runtime.Value {
	once_Data_Set_fromMap.Do(func() {
		cache_Data_Set_fromMap = Get_Data_Set_Set()
	})
	return cache_Data_Set_fromMap
}

var cache_Data_Set_foldableSet gopurs_runtime.Value
var once_Data_Set_foldableSet sync.Once
func Get_Data_Set_foldableSet() gopurs_runtime.Value {
	once_Data_Set_foldableSet.Do(func() {
		cache_Data_Set_foldableSet = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_3596835815_1680800814((&Constructor_Data_Foldable_Foldable[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply2(Rebox_Data_Set_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())).V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))}, f_1), Get_Data_Set_toList())
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_Set_go__go_2_0_1 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_Set_go__go_2_0_1
var go__go_2_0_1 gopurs_runtime.Value
_ = go__go_2_0_1
Call_local_Data_Set_go__go_2_0_1 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_0_1:
for {
if false { continue go__go_2_0_1 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4 == nil) {
__t1 = b_3
goto end_branch_1
} else {

}
}
{
if (v_4 != nil) {
b_3_loop = gopurs_runtime.Apply2(f_0, b_3, (v_4).V0)
v_4_loop = (v_4).V1
continue go__go_2_0_1
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
go__go_2_0_1 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_Set_go__go_2_0_1(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(go__go_2_0_1, x_1), Get_Data_Set_toList())
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply2(Rebox_Data_Set_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())).V2, f_0, x_1), Get_Data_Set_toList())
})})))}
	})
	return cache_Data_Set_foldableSet
}

var cache_Data_Set_findMin gopurs_runtime.Value
var once_Data_Set_findMin sync.Once
func Get_Data_Set_findMin() gopurs_runtime.Value {
	once_Data_Set_findMin.Do(func() {
		cache_Data_Set_findMin = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Set_findMin(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Set_findMin
}

var cache_Data_Set_findMax gopurs_runtime.Value
var once_Data_Set_findMax sync.Once
func Get_Data_Set_findMax() gopurs_runtime.Value {
	once_Data_Set_findMax.Do(func() {
		cache_Data_Set_findMax = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Set_findMax(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_Set_findMax
}

var cache_Data_Set_filter gopurs_runtime.Value
var once_Data_Set_filter sync.Once
func Get_Data_Set_filter() gopurs_runtime.Value {
	once_Data_Set_filter.Do(func() {
		cache_Data_Set_filter = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_filter(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Set_filter
}

var cache_Data_Set_eqSet gopurs_runtime.Value
var once_Data_Set_eqSet sync.Once
func Get_Data_Set_eqSet() gopurs_runtime.Value {
	once_Data_Set_eqSet.Do(func() {
		cache_Data_Set_eqSet = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_eqSet(dictEq_0_box)
})
	})
	return cache_Data_Set_eqSet
}

var cache_Data_Set_ordSet gopurs_runtime.Value
var once_Data_Set_ordSet sync.Once
func Get_Data_Set_ordSet() gopurs_runtime.Value {
	once_Data_Set_ordSet.Do(func() {
		cache_Data_Set_ordSet = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_ordSet(dictOrd_0_box)
})
	})
	return cache_Data_Set_ordSet
}

var cache_Data_Set_eq1Set gopurs_runtime.Value
var once_Data_Set_eq1Set sync.Once
func Get_Data_Set_eq1Set() gopurs_runtime.Value {
	once_Data_Set_eq1Set.Do(func() {
		cache_Data_Set_eq1Set = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_3691144502_1766074591((&Constructor_Data_Eq_Eq1[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Set_eqSet(dictEq_0)))
})})))}
	})
	return cache_Data_Set_eq1Set
}

var cache_Data_Set_ord1Set gopurs_runtime.Value
var once_Data_Set_ord1Set sync.Once
func Get_Data_Set_ord1Set() gopurs_runtime.Value {
	once_Data_Set_ord1Set.Do(func() {
		cache_Data_Set_ord1Set = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_849406934_3985601471((&Constructor_Data_Ord_Ord1[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_3691144502_1766074591(Rebox_Data_Set_1766074591_3691144502(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Set_eq1Set()))))}
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Call_Data_Set_ordSet(dictOrd_0)))
})})))}
	})
	return cache_Data_Set_ord1Set
}

var cache_Data_Set_empty gopurs_runtime.Value
var once_Data_Set_empty sync.Once
func Get_Data_Set_empty() gopurs_runtime.Value {
	once_Data_Set_empty.Do(func() {
		cache_Data_Set_empty = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Map_Internal_empty()))}
	})
	return cache_Data_Set_empty
}

var cache_Data_Set_fromFoldable gopurs_runtime.Value
var once_Data_Set_fromFoldable sync.Once
func Get_Data_Set_fromFoldable() gopurs_runtime.Value {
	once_Data_Set_fromFoldable.Do(func() {
		cache_Data_Set_fromFoldable = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_1_box))
})
	})
	return cache_Data_Set_fromFoldable
}

var cache_Data_Set_go__map gopurs_runtime.Value
var once_Data_Set_go__map sync.Once
func Get_Data_Set_go__map() gopurs_runtime.Value {
	once_Data_Set_go__map.Do(func() {
		cache_Data_Set_go__map = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), f_1_box)
})
	})
	return cache_Data_Set_go__map
}

var cache_Data_Set_mapMaybe gopurs_runtime.Value
var once_Data_Set_mapMaybe sync.Once
func Get_Data_Set_mapMaybe() gopurs_runtime.Value {
	once_Data_Set_mapMaybe.Do(func() {
		cache_Data_Set_mapMaybe = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_mapMaybe(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), f_1_box)
})
	})
	return cache_Data_Set_mapMaybe
}

var cache_Data_Set_monoidSet gopurs_runtime.Value
var once_Data_Set_monoidSet sync.Once
func Get_Data_Set_monoidSet() gopurs_runtime.Value {
	once_Data_Set_monoidSet.Do(func() {
		cache_Data_Set_monoidSet = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_monoidSet(dictOrd_0_box)
})
	})
	return cache_Data_Set_monoidSet
}

var cache_Data_Set_unions gopurs_runtime.Value
var once_Data_Set_unions sync.Once
func Get_Data_Set_unions() gopurs_runtime.Value {
	once_Data_Set_unions.Do(func() {
		cache_Data_Set_unions = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_unions(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_1_box))
})
	})
	return cache_Data_Set_unions
}

var cache_Data_Set_difference gopurs_runtime.Value
var once_Data_Set_difference sync.Once
func Get_Data_Set_difference() gopurs_runtime.Value {
	once_Data_Set_difference.Do(func() {
		cache_Data_Set_difference = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_difference(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Set_difference
}

var cache_Data_Set_subset gopurs_runtime.Value
var once_Data_Set_subset sync.Once
func Get_Data_Set_subset() gopurs_runtime.Value {
	once_Data_Set_subset.Do(func() {
		cache_Data_Set_subset = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, s1_1_box gopurs_runtime.Value, s2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Set_subset(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](s1_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](s2_2_box)))
})
	})
	return cache_Data_Set_subset
}

var cache_Data_Set_properSubset gopurs_runtime.Value
var once_Data_Set_properSubset sync.Once
func Get_Data_Set_properSubset() gopurs_runtime.Value {
	once_Data_Set_properSubset.Do(func() {
		cache_Data_Set_properSubset = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, s1_1_box gopurs_runtime.Value, s2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Set_properSubset(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](s1_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](s2_2_box)))
})
	})
	return cache_Data_Set_properSubset
}

var cache_Data_Set_go__delete gopurs_runtime.Value
var once_Data_Set_go__delete sync.Once
func Get_Data_Set_go__delete() gopurs_runtime.Value {
	once_Data_Set_go__delete.Do(func() {
		cache_Data_Set_go__delete = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_go__delete(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Set_go__delete
}

var cache_Data_Set_checkValid gopurs_runtime.Value
var once_Data_Set_checkValid sync.Once
func Get_Data_Set_checkValid() gopurs_runtime.Value {
	once_Data_Set_checkValid.Do(func() {
		cache_Data_Set_checkValid = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_checkValid(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Set_checkValid
}

var cache_Data_Set_catMaybes gopurs_runtime.Value
var once_Data_Set_catMaybes sync.Once
func Get_Data_Set_catMaybes() gopurs_runtime.Value {
	once_Data_Set_catMaybes.Do(func() {
		cache_Data_Set_catMaybes = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_catMaybes(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Set_catMaybes
}

func Call_Data_Set_Set(x_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Set_union(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return Call_Data_Map_Internal_union(dictOrd_0)
}

func Call_Data_Set_toggle(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var v_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_2_loop
_ = v_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply3(Call_Data_Map_Internal_alter(dictOrd_0), gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_3)
_ = __t_tag_0
if (__t_tag_0 == nil) {
__t2 = (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, Get_Data_Unit_unit()})
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_3)
_ = __t_tag_1
if (__t_tag_1 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
}), a_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_2)}))
}

func Call_Data_Set_toMap(v_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var v_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Set_toList(v_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
var v_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
var go__go_1_0_0 gopurs_runtime.Value
_ = go__go_1_0_0
var go__go_1_0_0_cell *gopurs_runtime.Value
_ = go__go_1_0_0_cell
// FALLBACK TCO: isLoop=false len=1
go__go_1_0_0 = gopurs_runtime.Func2(func(m_prime__2 gopurs_runtime.Value, __local_var_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__2)
_ = __t_tag_1
if (__t_tag_1 == nil) {
__t3 = __local_var_3
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__2)
_ = __t_tag_2
if (__t_tag_2 != nil) {
__t3 = gopurs_runtime.UncurriedApp2((*go__go_1_0_0_cell), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__2.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2((*go__go_1_0_0_cell), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__2.UnsafePtr).V5)}, __local_var_3))}))})
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
go__go_1_0_0_cell = &go__go_1_0_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_1_0_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}))
}

func Call_Data_Set_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_List_toUnfoldable(dictUnfoldable_0), Get_Data_Set_toList())
}

func Call_Data_Set_singleton(a_0_loop gopurs_runtime.Value) *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return (&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, int64(1), int64(1), a_0, Get_Data_Unit_unit(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil)})
}

func Call_Data_Set_showSet(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): showArray_1_0 shape=App(Var) bindingType=(ADT ["Data","Show","Show"] [(Array (TypeVar a$scope22))])
showArray_1_0 := Rebox_Data_Set_1386611502_1356436936(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_Show_showArray(dictShow_0)))
_ = showArray_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_2638796135_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(fromFoldable ") + (gopurs_runtime.Apply(showArray_1_0.V0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_List_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_Unfoldable_unfoldableArray())), Get_Data_Set_toList(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](s_2))}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).StrVal())) + (")"))
})})))}
}

func Call_Data_Set_semigroupSet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_3854229351_4179793454((&Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, Call_Data_Set_union(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0))})))}
}

func Call_Data_Set_member(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Map_Internal_member(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_Data_Set_intersection(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return Call_Data_Map_Internal_intersection(dictOrd_0)
}

func Call_Data_Set_insert(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var v_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_2_loop
_ = v_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Map_Internal_insert(dictOrd_0, a_1, Get_Data_Unit_unit()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_2)}))
}

func Call_Data_Set_findMin(v_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var v_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [key: (TypeVar a$scope62), value: Unit] Empty))])
__local_var_1_0 := Rebox_Data_Set_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Data_Map_Internal_findMin(v_0)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
_ = __local_var_1_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_1_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(__local_var_1_0).V0.key, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_1:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Set_findMax(v_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var v_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [key: (TypeVar a$scope64), value: Unit] Empty))])
__local_var_1_0 := Rebox_Data_Set_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Data_Map_Internal_findMax(v_0)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
_ = __local_var_1_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_1_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(__local_var_1_0).V0.key, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_1:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Set_filter(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Map_Internal_filterKeys(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_Data_Set_eqSet(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
// TAST (Let): eqMap_1_0 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope68), Unit])])
eqMap_1_0 := Rebox_Data_Set_3790796878_1444241223(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Map_Internal_eqMap(dictEq_0, gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqUnit()))})))
_ = eqMap_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_1444241223_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(eqMap_1_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v1_3))}).IntVal) != (0))
})})))}
}

func Call_Data_Set_ordSet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): ordList_1_0 shape=App(Var) bindingType=(ADT ["Data","Ord","Ord"] [(ADT ["Data","List","Types","List"] [(TypeVar a$scope28)])])
ordList_1_0 := Rebox_Data_Set_4177771502_4210054658(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Call_Data_List_Types_ordList(dictOrd_0)))
_ = ordList_1_0
// TAST (Let): eqSet1_2_1 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope28), Unit])])
eqSet1_2_1 := Rebox_Data_Set_3790796878_1444241223(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Set_eqSet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))))
_ = eqSet1_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_1910448679_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_1444241223_3790796878(eqSet1_2_1))}
}), gopurs_runtime.Func2(func(s1_3 gopurs_runtime.Value, s2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_2_2 gopurs_runtime.Value
_ = go__go_5_2_2
var go__go_5_2_2_cell *gopurs_runtime.Value
_ = go__go_5_2_2_cell
// FALLBACK TCO: isLoop=false len=1
go__go_5_2_2 = gopurs_runtime.Func2(func(m_prime__6 gopurs_runtime.Value, __local_var_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
var __t_tag_3 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__6)
_ = __t_tag_3
if (__t_tag_3 == nil) {
__t5 = __local_var_7
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__6)
_ = __t_tag_4
if (__t_tag_4 != nil) {
__t5 = gopurs_runtime.UncurriedApp2((*go__go_5_2_2_cell), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__6.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__6.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2((*go__go_5_2_2_cell), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__6.UnsafePtr).V5)}, __local_var_7))}))})
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
go__go_5_2_2_cell = &go__go_5_2_2
var go__go_5_6_3 gopurs_runtime.Value
_ = go__go_5_6_3
var go__go_5_6_3_cell *gopurs_runtime.Value
_ = go__go_5_6_3_cell
// FALLBACK TCO: isLoop=false len=1
go__go_5_6_3 = gopurs_runtime.Func2(func(m_prime__6 gopurs_runtime.Value, __local_var_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
var __t_tag_7 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__6)
_ = __t_tag_7
if (__t_tag_7 == nil) {
__t9 = __local_var_7
goto end_branch_9
} else {

}
}
{
var __t_tag_8 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__6)
_ = __t_tag_8
if (__t_tag_8 != nil) {
__t9 = gopurs_runtime.UncurriedApp2((*go__go_5_6_3_cell), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__6.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__6.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2((*go__go_5_6_3_cell), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__6.UnsafePtr).V5)}, __local_var_7))}))})
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
go__go_5_6_3_cell = &go__go_5_6_3
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(ordList_1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_5_2_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](s1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_5_6_3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](s2_4))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})))}).IntVal)), UnsafePtr: nil}
})})))}
}

func Call_Data_Set_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], dictOrd_1_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictOrd_1 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_1_loop
_ = dictOrd_1
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func2(func(m_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Map_Internal_insert(dictOrd_1, a_3, Get_Data_Unit_unit()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_2))})))}
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Set_empty()))})
}

func Call_Data_Set_go__map(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var Call_local_Data_Set_go__go_2_0_4 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_Set_go__go_2_0_4
var go__go_2_0_4 gopurs_runtime.Value
_ = go__go_2_0_4
Call_local_Data_Set_go__go_2_0_4 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_0_4:
for {
if false { continue go__go_2_0_4 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4 == nil) {
__t1 = b_3
goto end_branch_1
} else {

}
}
{
if (v_4 != nil) {
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Map_Internal_insert(dictOrd_0, gopurs_runtime.Apply(f_1, (v_4).V0), Get_Data_Unit_unit()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](b_3))})))}
v_4_loop = (v_4).V1
continue go__go_2_0_4
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
go__go_2_0_4 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_Set_go__go_2_0_4(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(go__go_2_0_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Set_empty()))}), Get_Data_Set_toList())
}

func Call_Data_Set_mapMaybe(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply2(Rebox_Data_Set_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())).V2, gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope87)])
__local_var_4_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, a_2))
_ = __local_var_4_0
var __t1 gopurs_runtime.Value
{
if (__local_var_4_0 == nil) {
__t1 = acc_3
goto end_branch_1
} else {

}
}
{
if (__local_var_4_0 != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_Map_Internal_insert(dictOrd_0, (__local_var_4_0).V0, Get_Data_Unit_unit()), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](acc_3))})))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t1))}
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Set_empty()))}), Get_Data_Set_toList())
}

func Call_Data_Set_monoidSet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): semigroupSet1_1_0 shape=App(Var) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(ADT ["Data","Map","Internal","Map"] [(TypeVar a$scope89), Unit])])
semigroupSet1_1_0 := Rebox_Data_Set_4179793454_3854229351(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Call_Data_Set_semigroupSet(dictOrd_0)))
_ = semigroupSet1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_1291127239_1201789390((&Constructor_Data_Monoid_Monoid[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_3854229351_4179793454(semigroupSet1_1_0))}
}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Map_Internal_empty())})))}
}

func Call_Data_Set_unions(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], dictOrd_1_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictOrd_1 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_1_loop
_ = dictOrd_1
return gopurs_runtime.Apply2(dictFoldable_0.V1, Call_Data_Set_union(dictOrd_1), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Set_empty()))})
}

func Call_Data_Set_difference(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return Call_Data_Map_Internal_difference(dictOrd_0)
}

func Call_Data_Set_subset(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], s1_1_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value], s2_2_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var s1_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = s1_1_loop
_ = s1_1
var s2_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = s2_2_loop
_ = s2_2
return (gopurs_runtime.Apply(Get_Data_Set_isEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(Call_Data_Map_Internal_difference(dictOrd_0), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s1_1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s2_2)})))}).IntVal) != (0)
}

func Call_Data_Set_properSubset(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], s1_1_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value], s2_2_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var s1_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = s1_1_loop
_ = s1_1
var s2_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = s2_2_loop
_ = s2_2
return (((Call_Data_Map_Internal_size(s1_1)) == (Call_Data_Map_Internal_size(s2_2))) != (true)) && ((gopurs_runtime.Apply(Get_Data_Set_isEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(Call_Data_Map_Internal_difference(dictOrd_0), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s1_1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s2_2)})))}).IntVal) != (0))
}

func Call_Data_Set_go__delete(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Map_Internal_go__delete(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_Data_Set_checkValid(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return Call_Data_Map_Internal_checkValid(dictOrd_0)
}

func Call_Data_Set_catMaybes(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return Call_Data_Set_mapMaybe(dictOrd_0, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))
}

func Rebox_Data_Set_1291127239_1201789390(in *Constructor_Data_Monoid_Monoid[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}

func Rebox_Data_Set_1386611502_1356436936(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[[]gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[[]gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Set_1444241223_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Set_1680800814_1022383170(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Set_1766074591_3691144502(in *Constructor_Data_Eq_Eq1[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq1[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq1[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Set_1910448679_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Set_2638796135_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Set_3094389156_2758605161(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct{
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

func Rebox_Data_Set_3596835815_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Set_3691144502_1766074591(in *Constructor_Data_Eq_Eq1[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Set_3790796878_1444241223(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Set_3854229351_4179793454(in *Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Set_4177771502_4210054658(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Set_4179793454_3854229351(in *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Set_849406934_3985601471(in *Constructor_Data_Ord_Ord1[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}



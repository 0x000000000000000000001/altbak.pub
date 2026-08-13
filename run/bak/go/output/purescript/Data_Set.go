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
		cache_Data_Set_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Set_identity(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](x_0_box)))}
})
	})
	return cache_Data_Set_identity
}

var cache_Data_Set_Set gopurs_runtime.Value
var once_Data_Set_Set sync.Once
func Get_Data_Set_Set() gopurs_runtime.Value {
	once_Data_Set_Set.Do(func() {
		cache_Data_Set_Set = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_Set(x_0_box)
})
	})
	return cache_Data_Set_Set
}

var cache_Data_Set_union gopurs_runtime.Value
var once_Data_Set_union sync.Once
func Get_Data_Set_union() gopurs_runtime.Value {
	once_Data_Set_union.Do(func() {
		cache_Data_Set_union = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_union(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Set_union
}

var cache_Data_Set_toggle gopurs_runtime.Value
var once_Data_Set_toggle sync.Once
func Get_Data_Set_toggle() gopurs_runtime.Value {
	once_Data_Set_toggle.Do(func() {
		cache_Data_Set_toggle = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Set_toggle(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), a_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_2_box)))}
})
	})
	return cache_Data_Set_toggle
}

var cache_Data_Set_toMap gopurs_runtime.Value
var once_Data_Set_toMap sync.Once
func Get_Data_Set_toMap() gopurs_runtime.Value {
	once_Data_Set_toMap.Do(func() {
		cache_Data_Set_toMap = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Set_toMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box)))}
})
	})
	return cache_Data_Set_toMap
}

var cache_Data_Set_toList gopurs_runtime.Value
var once_Data_Set_toList sync.Once
func Get_Data_Set_toList() gopurs_runtime.Value {
	once_Data_Set_toList.Do(func() {
		cache_Data_Set_toList = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Set_toList(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box)))}
})
	})
	return cache_Data_Set_toList
}

var cache_Data_Set_toUnfoldable gopurs_runtime.Value
var once_Data_Set_toUnfoldable sync.Once
func Get_Data_Set_toUnfoldable() gopurs_runtime.Value {
	once_Data_Set_toUnfoldable.Do(func() {
		cache_Data_Set_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dictUnfoldable_0_box))
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
		cache_Data_Set_member = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_member(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), k_1_box)
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
return Call_Data_Set_intersection(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Set_intersection
}

var cache_Data_Set_insert gopurs_runtime.Value
var once_Data_Set_insert sync.Once
func Get_Data_Set_insert() gopurs_runtime.Value {
	once_Data_Set_insert.Do(func() {
		cache_Data_Set_insert = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Set_insert(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), a_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_2_box)))}
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
		cache_Data_Set_foldableSet = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_0))}, f_1)
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_4 gopurs_runtime.Value
_ = go__go_4_1_4
go__go_4_1_4 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t2 = __local_var_6
goto end_branch_2
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t2 = gopurs_runtime.UncurriedApp2(go__go_4_1_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_1_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
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
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_1_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), f_0, x_1)
_ = __local_var_2_3
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_4_5 gopurs_runtime.Value
_ = go__go_4_4_5
go__go_4_4_5 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t5 = __local_var_6
goto end_branch_5
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t5 = gopurs_runtime.UncurriedApp2(go__go_4_4_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_4_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
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
return gopurs_runtime.Apply(__local_var_2_3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_4_5, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_6 -> gopurs_runtime.Value
__local_var_2_6 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldr"), f_0, x_1)
_ = __local_var_2_6
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_7_6 gopurs_runtime.Value
_ = go__go_4_7_6
go__go_4_7_6 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t8 = __local_var_6
goto end_branch_8
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t8 = gopurs_runtime.UncurriedApp2(go__go_4_7_6, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_7_6, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
})
return gopurs_runtime.Apply(__local_var_2_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_7_6, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
})
}))
	})
	return cache_Data_Set_foldableSet
}

var cache_Data_Set_findMin gopurs_runtime.Value
var once_Data_Set_findMin sync.Once
func Get_Data_Set_findMin() gopurs_runtime.Value {
	once_Data_Set_findMin.Do(func() {
		cache_Data_Set_findMin = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Set_findMin(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box)))}
})
	})
	return cache_Data_Set_findMin
}

var cache_Data_Set_findMax gopurs_runtime.Value
var once_Data_Set_findMax sync.Once
func Get_Data_Set_findMax() gopurs_runtime.Value {
	once_Data_Set_findMax.Do(func() {
		cache_Data_Set_findMax = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Set_findMax(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box)))}
})
	})
	return cache_Data_Set_findMax
}

var cache_Data_Set_filter gopurs_runtime.Value
var once_Data_Set_filter sync.Once
func Get_Data_Set_filter() gopurs_runtime.Value {
	once_Data_Set_filter.Do(func() {
		cache_Data_Set_filter = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_filter(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
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
		cache_Data_Set_eq1Set = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqMap_1_0 -> *Constructor_Data_Eq_Eq
eqMap_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.Apply2(Get_Data_Map_Internal_eqMap(), dictEq_0, Get_Data_Eq_eqUnit()))
_ = eqMap_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(eqMap_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v1_3))}).IntVal) != (0))
})
})
}))
	})
	return cache_Data_Set_eq1Set
}

var cache_Data_Set_ord1Set gopurs_runtime.Value
var once_Data_Set_ord1Set sync.Once
func Get_Data_Set_ord1Set() gopurs_runtime.Value {
	once_Data_Set_ord1Set.Do(func() {
		cache_Data_Set_ord1Set = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqMap_2_0 -> *Constructor_Data_Eq_Eq
eqMap_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.Apply2(Get_Data_Map_Internal_eqMap(), dictEq_1, Get_Data_Eq_eqUnit()))
_ = eqMap_2_0
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(eqMap_2_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v1_4))}).IntVal) != (0))
})
})
}))
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): ordList_1_1 -> *Constructor_Data_Ord_Ord
ordList_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](gopurs_runtime.Apply(Get_Data_List_Types_ordList(), dictOrd_0))
_ = ordList_1_1
return gopurs_runtime.Func(func(s1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_9 gopurs_runtime.Value
_ = go__go_4_2_9
go__go_4_2_9 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t3 = __local_var_6
goto end_branch_3
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t3 = gopurs_runtime.UncurriedApp2(go__go_4_2_9, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_2_9, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
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
var go__go_4_4_10 gopurs_runtime.Value
_ = go__go_4_4_10
go__go_4_4_10 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t5 = __local_var_6
goto end_branch_5
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t5 = gopurs_runtime.UncurriedApp2(go__go_4_4_10, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_4_10, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
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
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(ordList_1_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_2_9, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_4_10, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s2_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}).IntVal)), UnsafePtr: nil}
})
})
}))
	})
	return cache_Data_Set_ord1Set
}

var cache_Data_Set_empty gopurs_runtime.Value
var once_Data_Set_empty sync.Once
func Get_Data_Set_empty() gopurs_runtime.Value {
	once_Data_Set_empty.Do(func() {
		cache_Data_Set_empty = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))}
	})
	return cache_Data_Set_empty
}

var cache_Data_Set_fromFoldable gopurs_runtime.Value
var once_Data_Set_fromFoldable sync.Once
func Get_Data_Set_fromFoldable() gopurs_runtime.Value {
	once_Data_Set_fromFoldable.Do(func() {
		cache_Data_Set_fromFoldable = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictOrd_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_1_box))
})
	})
	return cache_Data_Set_fromFoldable
}

var cache_Data_Set_go__map gopurs_runtime.Value
var once_Data_Set_go__map sync.Once
func Get_Data_Set_go__map() gopurs_runtime.Value {
	once_Data_Set_go__map.Do(func() {
		cache_Data_Set_go__map = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), f_1_box)
})
	})
	return cache_Data_Set_go__map
}

var cache_Data_Set_mapMaybe gopurs_runtime.Value
var once_Data_Set_mapMaybe sync.Once
func Get_Data_Set_mapMaybe() gopurs_runtime.Value {
	once_Data_Set_mapMaybe.Do(func() {
		cache_Data_Set_mapMaybe = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_mapMaybe(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), f_1_box)
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
return Call_Data_Set_unions(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_1_box))
})
	})
	return cache_Data_Set_unions
}

var cache_Data_Set_difference gopurs_runtime.Value
var once_Data_Set_difference sync.Once
func Get_Data_Set_difference() gopurs_runtime.Value {
	once_Data_Set_difference.Do(func() {
		cache_Data_Set_difference = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_difference(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Set_difference
}

var cache_Data_Set_subset gopurs_runtime.Value
var once_Data_Set_subset sync.Once
func Get_Data_Set_subset() gopurs_runtime.Value {
	once_Data_Set_subset.Do(func() {
		cache_Data_Set_subset = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, s1_1_box gopurs_runtime.Value, s2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Set_subset(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s1_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s2_2_box)))
})
	})
	return cache_Data_Set_subset
}

var cache_Data_Set_properSubset gopurs_runtime.Value
var once_Data_Set_properSubset sync.Once
func Get_Data_Set_properSubset() gopurs_runtime.Value {
	once_Data_Set_properSubset.Do(func() {
		cache_Data_Set_properSubset = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, s1_1_box gopurs_runtime.Value, s2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Set_properSubset(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s1_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s2_2_box)))
})
	})
	return cache_Data_Set_properSubset
}

var cache_Data_Set_delete gopurs_runtime.Value
var once_Data_Set_delete sync.Once
func Get_Data_Set_delete() gopurs_runtime.Value {
	once_Data_Set_delete.Do(func() {
		cache_Data_Set_delete = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_delete(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Set_delete
}

var cache_Data_Set_checkValid gopurs_runtime.Value
var once_Data_Set_checkValid sync.Once
func Get_Data_Set_checkValid() gopurs_runtime.Value {
	once_Data_Set_checkValid.Do(func() {
		cache_Data_Set_checkValid = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_checkValid(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Set_checkValid
}

var cache_Data_Set_catMaybes gopurs_runtime.Value
var once_Data_Set_catMaybes sync.Once
func Get_Data_Set_catMaybes() gopurs_runtime.Value {
	once_Data_Set_catMaybes.Do(func() {
		cache_Data_Set_catMaybes = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_catMaybes(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Set_catMaybes
}

var cache_Data_Set_delete__4217907800 gopurs_runtime.Value
var once_Data_Set_delete__4217907800 sync.Once
func Get_Data_Set_delete__4217907800() gopurs_runtime.Value {
	once_Data_Set_delete__4217907800.Do(func() {
		cache_Data_Set_delete__4217907800 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_delete__4217907800(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Set_delete__4217907800
}

var cache_Data_Set_difference__833514008 gopurs_runtime.Value
var once_Data_Set_difference__833514008 sync.Once
func Get_Data_Set_difference__833514008() gopurs_runtime.Value {
	once_Data_Set_difference__833514008.Do(func() {
		cache_Data_Set_difference__833514008 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_difference__833514008(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Set_difference__833514008
}

var cache_Data_Set_empty__2198260019 gopurs_runtime.Value
var once_Data_Set_empty__2198260019 sync.Once
func Get_Data_Set_empty__2198260019() gopurs_runtime.Value {
	once_Data_Set_empty__2198260019.Do(func() {
		cache_Data_Set_empty__2198260019 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))}
	})
	return cache_Data_Set_empty__2198260019
}

var cache_Data_Set_empty__3279308360 gopurs_runtime.Value
var once_Data_Set_empty__3279308360 sync.Once
func Get_Data_Set_empty__3279308360() gopurs_runtime.Value {
	once_Data_Set_empty__3279308360.Do(func() {
		cache_Data_Set_empty__3279308360 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))}
	})
	return cache_Data_Set_empty__3279308360
}

var cache_Data_Set_eq1Set__656709182 gopurs_runtime.Value
var once_Data_Set_eq1Set__656709182 sync.Once
func Get_Data_Set_eq1Set__656709182() gopurs_runtime.Value {
	once_Data_Set_eq1Set__656709182.Do(func() {
		cache_Data_Set_eq1Set__656709182 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqMap_1_0 -> *Constructor_Data_Eq_Eq
eqMap_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.Apply2(Get_Data_Map_Internal_eqMap(), dictEq_0, Get_Data_Eq_eqUnit()))
_ = eqMap_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(eqMap_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v1_3))}).IntVal) != (0))
})
})
}))
	})
	return cache_Data_Set_eq1Set__656709182
}

var cache_Data_Set_findMax__611847841 gopurs_runtime.Value
var once_Data_Set_findMax__611847841 sync.Once
func Get_Data_Set_findMax__611847841() gopurs_runtime.Value {
	once_Data_Set_findMax__611847841.Do(func() {
		cache_Data_Set_findMax__611847841 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Set_findMax__611847841(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box)))}
})
	})
	return cache_Data_Set_findMax__611847841
}

var cache_Data_Set_findMin__611847841 gopurs_runtime.Value
var once_Data_Set_findMin__611847841 sync.Once
func Get_Data_Set_findMin__611847841() gopurs_runtime.Value {
	once_Data_Set_findMin__611847841.Do(func() {
		cache_Data_Set_findMin__611847841 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Set_findMin__611847841(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box)))}
})
	})
	return cache_Data_Set_findMin__611847841
}

var cache_Data_Set_foldableSet__2661034898 gopurs_runtime.Value
var once_Data_Set_foldableSet__2661034898 sync.Once
func Get_Data_Set_foldableSet__2661034898() gopurs_runtime.Value {
	once_Data_Set_foldableSet__2661034898.Do(func() {
		cache_Data_Set_foldableSet__2661034898 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_0))}, f_1)
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_14 gopurs_runtime.Value
_ = go__go_4_1_14
go__go_4_1_14 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t2 = __local_var_6
goto end_branch_2
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t2 = gopurs_runtime.UncurriedApp2(go__go_4_1_14, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_1_14, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
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
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_1_14, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), f_0, x_1)
_ = __local_var_2_3
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_4_15 gopurs_runtime.Value
_ = go__go_4_4_15
go__go_4_4_15 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t5 = __local_var_6
goto end_branch_5
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t5 = gopurs_runtime.UncurriedApp2(go__go_4_4_15, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_4_15, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
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
return gopurs_runtime.Apply(__local_var_2_3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_4_15, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_6 -> gopurs_runtime.Value
__local_var_2_6 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldr"), f_0, x_1)
_ = __local_var_2_6
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_7_16 gopurs_runtime.Value
_ = go__go_4_7_16
go__go_4_7_16 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t8 = __local_var_6
goto end_branch_8
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t8 = gopurs_runtime.UncurriedApp2(go__go_4_7_16, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_7_16, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
})
return gopurs_runtime.Apply(__local_var_2_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_7_16, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
})
}))
	})
	return cache_Data_Set_foldableSet__2661034898
}

var cache_Data_Set_foldableSet__3201980275 gopurs_runtime.Value
var once_Data_Set_foldableSet__3201980275 sync.Once
func Get_Data_Set_foldableSet__3201980275() gopurs_runtime.Value {
	once_Data_Set_foldableSet__3201980275.Do(func() {
		cache_Data_Set_foldableSet__3201980275 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_0))}, f_1)
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_17 gopurs_runtime.Value
_ = go__go_4_1_17
go__go_4_1_17 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t2 = __local_var_6
goto end_branch_2
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t2 = gopurs_runtime.UncurriedApp2(go__go_4_1_17, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_1_17, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
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
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_1_17, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), f_0, x_1)
_ = __local_var_2_3
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_4_18 gopurs_runtime.Value
_ = go__go_4_4_18
go__go_4_4_18 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t5 = __local_var_6
goto end_branch_5
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t5 = gopurs_runtime.UncurriedApp2(go__go_4_4_18, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_4_18, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
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
return gopurs_runtime.Apply(__local_var_2_3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_4_18, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_6 -> gopurs_runtime.Value
__local_var_2_6 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldr"), f_0, x_1)
_ = __local_var_2_6
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_7_19 gopurs_runtime.Value
_ = go__go_4_7_19
go__go_4_7_19 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t8 = __local_var_6
goto end_branch_8
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t8 = gopurs_runtime.UncurriedApp2(go__go_4_7_19, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_7_19, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
})
return gopurs_runtime.Apply(__local_var_2_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_7_19, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
})
}))
	})
	return cache_Data_Set_foldableSet__3201980275
}

var cache_Data_Set_foldableSet__1081688304 gopurs_runtime.Value
var once_Data_Set_foldableSet__1081688304 sync.Once
func Get_Data_Set_foldableSet__1081688304() gopurs_runtime.Value {
	once_Data_Set_foldableSet__1081688304.Do(func() {
		cache_Data_Set_foldableSet__1081688304 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_0))}, f_1)
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_20 gopurs_runtime.Value
_ = go__go_4_1_20
go__go_4_1_20 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t2 = __local_var_6
goto end_branch_2
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t2 = gopurs_runtime.UncurriedApp2(go__go_4_1_20, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_1_20, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
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
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_1_20, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), f_0, x_1)
_ = __local_var_2_3
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_4_21 gopurs_runtime.Value
_ = go__go_4_4_21
go__go_4_4_21 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t5 = __local_var_6
goto end_branch_5
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t5 = gopurs_runtime.UncurriedApp2(go__go_4_4_21, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_4_21, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
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
return gopurs_runtime.Apply(__local_var_2_3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_4_21, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_6 -> gopurs_runtime.Value
__local_var_2_6 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldr"), f_0, x_1)
_ = __local_var_2_6
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_7_22 gopurs_runtime.Value
_ = go__go_4_7_22
go__go_4_7_22 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t8 = __local_var_6
goto end_branch_8
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t8 = gopurs_runtime.UncurriedApp2(go__go_4_7_22, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_7_22, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
})
return gopurs_runtime.Apply(__local_var_2_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_7_22, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
})
}))
	})
	return cache_Data_Set_foldableSet__1081688304
}

var cache_Data_Set_fromMap__556171355 gopurs_runtime.Value
var once_Data_Set_fromMap__556171355 sync.Once
func Get_Data_Set_fromMap__556171355() gopurs_runtime.Value {
	once_Data_Set_fromMap__556171355.Do(func() {
		cache_Data_Set_fromMap__556171355 = Get_Data_Set_Set()
	})
	return cache_Data_Set_fromMap__556171355
}

var cache_Data_Set_insert__4217907800 gopurs_runtime.Value
var once_Data_Set_insert__4217907800 sync.Once
func Get_Data_Set_insert__4217907800() gopurs_runtime.Value {
	once_Data_Set_insert__4217907800.Do(func() {
		cache_Data_Set_insert__4217907800 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Set_insert__4217907800(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), a_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_2_box)))}
})
	})
	return cache_Data_Set_insert__4217907800
}

var cache_Data_Set_insert__3935803192 gopurs_runtime.Value
var once_Data_Set_insert__3935803192 sync.Once
func Get_Data_Set_insert__3935803192() gopurs_runtime.Value {
	once_Data_Set_insert__3935803192.Do(func() {
		cache_Data_Set_insert__3935803192 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Set_insert__3935803192(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](a_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_2_box)))}
})
	})
	return cache_Data_Set_insert__3935803192
}

var cache_Data_Set_intersection__833514008 gopurs_runtime.Value
var once_Data_Set_intersection__833514008 sync.Once
func Get_Data_Set_intersection__833514008() gopurs_runtime.Value {
	once_Data_Set_intersection__833514008.Do(func() {
		cache_Data_Set_intersection__833514008 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_intersection__833514008(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Set_intersection__833514008
}

var cache_Data_Set_isEmpty__1620059593 gopurs_runtime.Value
var once_Data_Set_isEmpty__1620059593 sync.Once
func Get_Data_Set_isEmpty__1620059593() gopurs_runtime.Value {
	once_Data_Set_isEmpty__1620059593.Do(func() {
		cache_Data_Set_isEmpty__1620059593 = Get_Data_Map_Internal_isEmpty()
	})
	return cache_Data_Set_isEmpty__1620059593
}

var cache_Data_Set_isEmpty__1390394866 gopurs_runtime.Value
var once_Data_Set_isEmpty__1390394866 sync.Once
func Get_Data_Set_isEmpty__1390394866() gopurs_runtime.Value {
	once_Data_Set_isEmpty__1390394866.Do(func() {
		cache_Data_Set_isEmpty__1390394866 = Get_Data_Map_Internal_isEmpty()
	})
	return cache_Data_Set_isEmpty__1390394866
}

var cache_Data_Set_mapMaybe__2748927564 gopurs_runtime.Value
var once_Data_Set_mapMaybe__2748927564 sync.Once
func Get_Data_Set_mapMaybe__2748927564() gopurs_runtime.Value {
	once_Data_Set_mapMaybe__2748927564.Do(func() {
		cache_Data_Set_mapMaybe__2748927564 = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_mapMaybe__2748927564(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), f_1_box)
})
	})
	return cache_Data_Set_mapMaybe__2748927564
}

var cache_Data_Set_mapMaybe__4234899916 gopurs_runtime.Value
var once_Data_Set_mapMaybe__4234899916 sync.Once
func Get_Data_Set_mapMaybe__4234899916() gopurs_runtime.Value {
	once_Data_Set_mapMaybe__4234899916.Do(func() {
		cache_Data_Set_mapMaybe__4234899916 = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_mapMaybe__4234899916(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), f_1_box)
})
	})
	return cache_Data_Set_mapMaybe__4234899916
}

var cache_Data_Set_ord1Set__1315498729 gopurs_runtime.Value
var once_Data_Set_ord1Set__1315498729 sync.Once
func Get_Data_Set_ord1Set__1315498729() gopurs_runtime.Value {
	once_Data_Set_ord1Set__1315498729.Do(func() {
		cache_Data_Set_ord1Set__1315498729 = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): eqMap_2_0 -> *Constructor_Data_Eq_Eq
eqMap_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.Apply2(Get_Data_Map_Internal_eqMap(), dictEq_1, Get_Data_Eq_eqUnit()))
_ = eqMap_2_0
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(eqMap_2_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v1_4))}).IntVal) != (0))
})
})
}))
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): ordList_1_1 -> *Constructor_Data_Ord_Ord
ordList_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](gopurs_runtime.Apply(Get_Data_List_Types_ordList(), dictOrd_0))
_ = ordList_1_1
return gopurs_runtime.Func(func(s1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_25 gopurs_runtime.Value
_ = go__go_4_2_25
go__go_4_2_25 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t3 = __local_var_6
goto end_branch_3
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t3 = gopurs_runtime.UncurriedApp2(go__go_4_2_25, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_2_25, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
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
var go__go_4_4_26 gopurs_runtime.Value
_ = go__go_4_4_26
go__go_4_4_26 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t5 = __local_var_6
goto end_branch_5
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t5 = gopurs_runtime.UncurriedApp2(go__go_4_4_26, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_4_26, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
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
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(ordList_1_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_2_25, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_4_26, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s2_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}).IntVal)), UnsafePtr: nil}
})
})
}))
	})
	return cache_Data_Set_ord1Set__1315498729
}

var cache_Data_Set_singleton__3724491835 gopurs_runtime.Value
var once_Data_Set_singleton__3724491835 sync.Once
func Get_Data_Set_singleton__3724491835() gopurs_runtime.Value {
	once_Data_Set_singleton__3724491835.Do(func() {
		cache_Data_Set_singleton__3724491835 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Set_singleton__3724491835(a_0_box))}
})
	})
	return cache_Data_Set_singleton__3724491835
}

var cache_Data_Set_size__909390430 gopurs_runtime.Value
var once_Data_Set_size__909390430 sync.Once
func Get_Data_Set_size__909390430() gopurs_runtime.Value {
	once_Data_Set_size__909390430.Do(func() {
		cache_Data_Set_size__909390430 = Get_Data_Map_Internal_size()
	})
	return cache_Data_Set_size__909390430
}

var cache_Data_Set_size__667039429 gopurs_runtime.Value
var once_Data_Set_size__667039429 sync.Once
func Get_Data_Set_size__667039429() gopurs_runtime.Value {
	once_Data_Set_size__667039429.Do(func() {
		cache_Data_Set_size__667039429 = Get_Data_Map_Internal_size()
	})
	return cache_Data_Set_size__667039429
}

var cache_Data_Set_subset__3164788426 gopurs_runtime.Value
var once_Data_Set_subset__3164788426 sync.Once
func Get_Data_Set_subset__3164788426() gopurs_runtime.Value {
	once_Data_Set_subset__3164788426.Do(func() {
		cache_Data_Set_subset__3164788426 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, s1_1_box gopurs_runtime.Value, s2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Set_subset__3164788426(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s1_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s2_2_box)))
})
	})
	return cache_Data_Set_subset__3164788426
}

var cache_Data_Set_toList__319294638 gopurs_runtime.Value
var once_Data_Set_toList__319294638 sync.Once
func Get_Data_Set_toList__319294638() gopurs_runtime.Value {
	once_Data_Set_toList__319294638.Do(func() {
		cache_Data_Set_toList__319294638 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Set_toList__319294638(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box)))}
})
	})
	return cache_Data_Set_toList__319294638
}

var cache_Data_Set_toList__2510070446 gopurs_runtime.Value
var once_Data_Set_toList__2510070446 sync.Once
func Get_Data_Set_toList__2510070446() gopurs_runtime.Value {
	once_Data_Set_toList__2510070446.Do(func() {
		cache_Data_Set_toList__2510070446 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Set_toList__2510070446(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box)))}
})
	})
	return cache_Data_Set_toList__2510070446
}

var cache_Data_Set_toList__2521903733 gopurs_runtime.Value
var once_Data_Set_toList__2521903733 sync.Once
func Get_Data_Set_toList__2521903733() gopurs_runtime.Value {
	once_Data_Set_toList__2521903733.Do(func() {
		cache_Data_Set_toList__2521903733 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Set_toList__2521903733(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box)))}
})
	})
	return cache_Data_Set_toList__2521903733
}

var cache_Data_Set_toMap__556171355 gopurs_runtime.Value
var once_Data_Set_toMap__556171355 sync.Once
func Get_Data_Set_toMap__556171355() gopurs_runtime.Value {
	once_Data_Set_toMap__556171355.Do(func() {
		cache_Data_Set_toMap__556171355 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Set_toMap__556171355(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box)))}
})
	})
	return cache_Data_Set_toMap__556171355
}

var cache_Data_Set_toUnfoldable__871784487 gopurs_runtime.Value
var once_Data_Set_toUnfoldable__871784487 sync.Once
func Get_Data_Set_toUnfoldable__871784487() gopurs_runtime.Value {
	once_Data_Set_toUnfoldable__871784487.Do(func() {
		cache_Data_Set_toUnfoldable__871784487 = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_toUnfoldable__871784487(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dictUnfoldable_0_box))
})
	})
	return cache_Data_Set_toUnfoldable__871784487
}

func Call_Data_Set_identity(x_0_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var x_0 *Constructor_Data_Maybe_Just = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Set_Set(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Set_union(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
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

func Call_Data_Set_toggle(dictOrd_0_loop *Constructor_Data_Ord_Ord, a_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Map_Internal_Node {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var v_2 *Constructor_Data_Map_Internal_Node = v_2_loop
_ = v_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply4(Get_Data_Map_Internal_alter(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just
{
if (v2_3.Type == 9 && v2_3.IntVal == 930809136 && v2_3.UnsafePtr == nil) {
__t0 = &Constructor_Data_Maybe_Just{1, Get_Data_Unit_unit()}
goto end_branch_0
} else {

}
}
{
if (v2_3.Type == 9 && v2_3.IntVal == 930809136 && v2_3.UnsafePtr != nil) {
__t0 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
}), a_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_2)}))
}

func Call_Data_Set_toMap(v_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Map_Internal_Node {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Set_toList(v_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_List_Types_Cons {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
var go__go_1_0_0 gopurs_runtime.Value
_ = go__go_1_0_0
go__go_1_0_0 = gopurs_runtime.Func2(func(__local_var_2 gopurs_runtime.Value, __local_var_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t1 = __local_var_3
goto end_branch_1
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.UncurriedApp2(go__go_1_0_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_1_0_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)}, __local_var_3))})})
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
})
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_1_0_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}))
}

func Call_Data_Set_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable = dictUnfoldable_0_loop
_ = dictUnfoldable_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_List_toUnfoldable(), gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(dictUnfoldable_0)})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_1 gopurs_runtime.Value
_ = go__go_3_1_1
go__go_3_1_1 = gopurs_runtime.Func2(func(__local_var_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr == nil) {
__t2 = __local_var_5
goto end_branch_2
} else {

}
}
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr != nil) {
__t2 = gopurs_runtime.UncurriedApp2(go__go_3_1_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_3_1_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V5)}, __local_var_5))})})
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
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_3_1_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
}

func Call_Data_Set_singleton(a_0_loop gopurs_runtime.Value) *Constructor_Data_Map_Internal_Node {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return &Constructor_Data_Map_Internal_Node{1, 1, 1, a_0, Get_Data_Unit_unit(), (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
}

func Call_Data_Set_showSet(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): showArray_1_0 -> *Constructor_Data_Show_Show
showArray_1_0 := &Constructor_Data_Show_Show{1, gopurs_runtime.Apply(Get_Data_Show_showArrayImpl(), gopurs_runtime.RecordGet(dictShow_0, "show"))}
_ = showArray_1_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_2 gopurs_runtime.Value
_ = go__go_3_1_2
go__go_3_1_2 = gopurs_runtime.Func2(func(__local_var_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr == nil) {
__t2 = __local_var_5
goto end_branch_2
} else {

}
}
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr != nil) {
__t2 = gopurs_runtime.UncurriedApp2(go__go_3_1_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_3_1_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V5)}, __local_var_5))})})
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
return gopurs_runtime.Str((("(fromFoldable ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showArray_1_0.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_List_toUnfoldable(), gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](Get_Data_Unfoldable_unfoldableArray()))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_3_1_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).StrVal())) + (")"))
}))
}

func Call_Data_Set_semigroupSet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_Data_Set_member(dictOrd_0_loop *Constructor_Data_Ord_Ord, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_3 gopurs_runtime.Value
go__go_2_0_3 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_2_0_3:
for {
if false { continue go__go_2_0_3 }
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
continue go__go_2_0_3
__t2 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)}
continue go__go_2_0_3
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
return go__go_2_0_3
}

func Call_Data_Set_intersection(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), compare_1_0, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_3))})))}
})
})
}

func Call_Data_Set_insert(dictOrd_0_loop *Constructor_Data_Ord_Ord, a_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Map_Internal_Node {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var v_2 *Constructor_Data_Map_Internal_Node = v_2_loop
_ = v_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply4(Get_Data_Map_Internal_insert(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, a_1, Get_Data_Unit_unit(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_2)}))
}

func Call_Data_Set_findMin(v_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v1_1, "key")
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Map_Internal_findMin(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)})))}))
}

func Call_Data_Set_findMax(v_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v1_1, "key")
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Map_Internal_findMax(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)})))}))
}

func Call_Data_Set_filter(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Map_Internal_filterKeys(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_Data_Set_eqSet(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
// TAST (Let): eqMap_1_0 -> *Constructor_Data_Eq_Eq
eqMap_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.Apply2(Get_Data_Map_Internal_eqMap(), dictEq_0, Get_Data_Eq_eqUnit()))
_ = eqMap_1_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(eqMap_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v1_3))}).IntVal) != (0))
})
}))
}

func Call_Data_Set_ordSet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): ordList_1_0 -> *Constructor_Data_Ord_Ord
ordList_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](gopurs_runtime.Apply(Get_Data_List_Types_ordList(), dictOrd_0))
_ = ordList_1_0
// TAST (Let): eqMap_2_2 -> *Constructor_Data_Eq_Eq
eqMap_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.Apply2(Get_Data_Map_Internal_eqMap(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}), Get_Data_Eq_eqUnit()))
_ = eqMap_2_2
// TAST (Let): eqSet1_2_1 -> gopurs_runtime.Value
eqSet1_2_1 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(eqMap_2_2.V0), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v1_4))}).IntVal) != (0))
})
}))
_ = eqSet1_2_1
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eqSet1_2_1
}), gopurs_runtime.Func(func(s1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_3_7 gopurs_runtime.Value
_ = go__go_5_3_7
go__go_5_3_7 = gopurs_runtime.Func2(func(__local_var_6 gopurs_runtime.Value, __local_var_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (__local_var_6.Type == 9 && __local_var_6.IntVal == 324739070 && __local_var_6.UnsafePtr == nil) {
__t4 = __local_var_7
goto end_branch_4
} else {

}
}
{
if (__local_var_6.Type == 9 && __local_var_6.IntVal == 324739070 && __local_var_6.UnsafePtr != nil) {
__t4 = gopurs_runtime.UncurriedApp2(go__go_5_3_7, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_5_3_7, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V5)}, __local_var_7))})})
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
var go__go_5_5_8 gopurs_runtime.Value
_ = go__go_5_5_8
go__go_5_5_8 = gopurs_runtime.Func2(func(__local_var_6 gopurs_runtime.Value, __local_var_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (__local_var_6.Type == 9 && __local_var_6.IntVal == 324739070 && __local_var_6.UnsafePtr == nil) {
__t6 = __local_var_7
goto end_branch_6
} else {

}
}
{
if (__local_var_6.Type == 9 && __local_var_6.IntVal == 324739070 && __local_var_6.UnsafePtr != nil) {
__t6 = gopurs_runtime.UncurriedApp2(go__go_5_5_8, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_5_5_8, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V5)}, __local_var_7))})})
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
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(ordList_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_5_3_7, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_5_5_8, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s2_4))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}).IntVal)), UnsafePtr: nil}
})
}))
}

func Call_Data_Set_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictOrd_1_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictOrd_1 *Constructor_Data_Ord_Ord = dictOrd_1_loop
_ = dictOrd_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply4(Get_Data_Map_Internal_insert(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_1)}, a_3, Get_Data_Unit_unit(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_2))})))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
}

func Call_Data_Set_go__map(dictOrd_0_loop *Constructor_Data_Ord_Ord, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply4(Get_Data_Map_Internal_insert(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, gopurs_runtime.Apply(f_1, a_3), Get_Data_Unit_unit(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_2))})))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_11 gopurs_runtime.Value
_ = go__go_4_1_11
go__go_4_1_11 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t2 = __local_var_6
goto end_branch_2
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t2 = gopurs_runtime.UncurriedApp2(go__go_4_1_11, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_1_11, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
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
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_1_11, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
}

func Call_Data_Set_mapMaybe(dictOrd_0_loop *Constructor_Data_Ord_Ord, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldr"), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> *Constructor_Data_Maybe_Just
__local_var_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_1, a_2))
_ = __local_var_4_1
var __t2 *Constructor_Data_Map_Internal_Node
{
if (__local_var_4_1 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](acc_3)
goto end_branch_2
} else {

}
}
{
if (__local_var_4_1 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply4(Get_Data_Map_Internal_insert(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, (__local_var_4_1).V0, Get_Data_Unit_unit(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](acc_3))}))
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
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_3_12 gopurs_runtime.Value
_ = go__go_4_3_12
go__go_4_3_12 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t4 = __local_var_6
goto end_branch_4
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t4 = gopurs_runtime.UncurriedApp2(go__go_4_3_12, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_3_12, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
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
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_3_12, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
}

func Call_Data_Set_monoidSet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_1 -> gopurs_runtime.Value
compare_1_1 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_1
// TAST (Let): semigroupSet1_1_0 -> gopurs_runtime.Value
semigroupSet1_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_1, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_3))})))}
})
}))
_ = semigroupSet1_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupSet1_1_0
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
}

func Call_Data_Set_unions(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictOrd_1_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictOrd_1 *Constructor_Data_Ord_Ord = dictOrd_1_loop
_ = dictOrd_1
// TAST (Let): compare_2_0 -> gopurs_runtime.Value
compare_2_0 := gopurs_runtime.Box(dictOrd_1.V1)
_ = compare_2_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(m1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_2_0, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_4))})))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
}

func Call_Data_Set_difference(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeDifference(), compare_1_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_3))})))}
})
})
}

func Call_Data_Set_subset(dictOrd_0_loop *Constructor_Data_Ord_Ord, s1_1_loop *Constructor_Data_Map_Internal_Node, s2_2_loop *Constructor_Data_Map_Internal_Node) bool {
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

func Call_Data_Set_properSubset(dictOrd_0_loop *Constructor_Data_Ord_Ord, s1_1_loop *Constructor_Data_Map_Internal_Node, s2_2_loop *Constructor_Data_Map_Internal_Node) bool {
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

func Call_Data_Set_delete(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Map_Internal_delete(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_Data_Set_checkValid(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Map_Internal_checkValid(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_Data_Set_catMaybes(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldr"), gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Map_Internal_Node
{
var __t_tag_1 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](a_1)
if (__t_tag_1 == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](acc_2)
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](a_1)
if (__t_tag_2 != nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply4(Get_Data_Map_Internal_insert(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, (*Constructor_Data_Maybe_Just)(a_1.UnsafePtr).V0, Get_Data_Unit_unit(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](acc_2))}))
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
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_4_13 gopurs_runtime.Value
_ = go__go_3_4_13
go__go_3_4_13 = gopurs_runtime.Func2(func(__local_var_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr == nil) {
__t5 = __local_var_5
goto end_branch_5
} else {

}
}
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr != nil) {
__t5 = gopurs_runtime.UncurriedApp2(go__go_3_4_13, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_3_4_13, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V5)}, __local_var_5))})})
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
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_3_4_13, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
}

func Call_Data_Set_delete__4217907800(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Map_Internal_delete(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_Data_Set_difference__833514008(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeDifference(), compare_1_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_3))})))}
})
})
}

func Call_Data_Set_findMax__611847841(v_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v1_1, "key")
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Map_Internal_findMax(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)})))}))
}

func Call_Data_Set_findMin__611847841(v_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v1_1, "key")
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Map_Internal_findMin(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)})))}))
}

func Call_Data_Set_insert__4217907800(dictOrd_0_loop *Constructor_Data_Ord_Ord, a_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Map_Internal_Node {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var v_2 *Constructor_Data_Map_Internal_Node = v_2_loop
_ = v_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply4(Get_Data_Map_Internal_insert(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, a_1, Get_Data_Unit_unit(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_2)}))
}

func Call_Data_Set_insert__3935803192(dictOrd_0_loop *Constructor_Data_Ord_Ord, a_1_loop *Constructor_Data_Maybe_Just, v_2_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Map_Internal_Node {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a_1 *Constructor_Data_Maybe_Just = a_1_loop
_ = a_1
var v_2 *Constructor_Data_Map_Internal_Node = v_2_loop
_ = v_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply4(Get_Data_Map_Internal_insert(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(a_1)}, Get_Data_Unit_unit(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_2)}))
}

func Call_Data_Set_intersection__833514008(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), compare_1_0, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_3))})))}
})
})
}

func Call_Data_Set_mapMaybe__2748927564(dictOrd_0_loop *Constructor_Data_Ord_Ord, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldr"), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> *Constructor_Data_Maybe_Just
__local_var_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_1, a_2))
_ = __local_var_4_1
var __t2 *Constructor_Data_Map_Internal_Node
{
if (__local_var_4_1 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](acc_3)
goto end_branch_2
} else {

}
}
{
if (__local_var_4_1 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply4(Get_Data_Map_Internal_insert(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, (__local_var_4_1).V0, Get_Data_Unit_unit(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](acc_3))}))
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
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_3_23 gopurs_runtime.Value
_ = go__go_4_3_23
go__go_4_3_23 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t4 = __local_var_6
goto end_branch_4
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t4 = gopurs_runtime.UncurriedApp2(go__go_4_3_23, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_3_23, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
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
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_3_23, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
}

func Call_Data_Set_mapMaybe__4234899916(dictOrd_0_loop *Constructor_Data_Ord_Ord, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldr"), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> *Constructor_Data_Maybe_Just
__local_var_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](a_2))}))
_ = __local_var_4_1
var __t2 *Constructor_Data_Map_Internal_Node
{
if (__local_var_4_1 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](acc_3)
goto end_branch_2
} else {

}
}
{
if (__local_var_4_1 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply4(Get_Data_Map_Internal_insert(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, (__local_var_4_1).V0, Get_Data_Unit_unit(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](acc_3))}))
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
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_3_24 gopurs_runtime.Value
_ = go__go_4_3_24
go__go_4_3_24 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t4 = __local_var_6
goto end_branch_4
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t4 = gopurs_runtime.UncurriedApp2(go__go_4_3_24, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_3_24, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
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
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_3_24, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
}

func Call_Data_Set_singleton__3724491835(a_0_loop gopurs_runtime.Value) *Constructor_Data_Map_Internal_Node {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return &Constructor_Data_Map_Internal_Node{1, 1, 1, a_0, Get_Data_Unit_unit(), (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
}

func Call_Data_Set_subset__3164788426(dictOrd_0_loop *Constructor_Data_Ord_Ord, s1_1_loop *Constructor_Data_Map_Internal_Node, s2_2_loop *Constructor_Data_Map_Internal_Node) bool {
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

func Call_Data_Set_toList__319294638(v_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_List_Types_Cons {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
var go__go_1_0_27 gopurs_runtime.Value
_ = go__go_1_0_27
go__go_1_0_27 = gopurs_runtime.Func2(func(__local_var_2 gopurs_runtime.Value, __local_var_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t1 = __local_var_3
goto end_branch_1
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.UncurriedApp2(go__go_1_0_27, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_1_0_27, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)}, __local_var_3))})})
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
})
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_1_0_27, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}))
}

func Call_Data_Set_toList__2510070446(v_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_List_Types_Cons {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
var go__go_1_0_28 gopurs_runtime.Value
_ = go__go_1_0_28
go__go_1_0_28 = gopurs_runtime.Func2(func(__local_var_2 gopurs_runtime.Value, __local_var_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t1 = __local_var_3
goto end_branch_1
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.UncurriedApp2(go__go_1_0_28, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_1_0_28, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)}, __local_var_3))})})
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
})
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_1_0_28, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}))
}

func Call_Data_Set_toList__2521903733(v_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_List_Types_Cons {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
var go__go_1_0_29 gopurs_runtime.Value
_ = go__go_1_0_29
go__go_1_0_29 = gopurs_runtime.Func2(func(__local_var_2 gopurs_runtime.Value, __local_var_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr == nil) {
__t1 = __local_var_3
goto end_branch_1
} else {

}
}
{
if (__local_var_2.Type == 9 && __local_var_2.IntVal == 324739070 && __local_var_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.UncurriedApp2(go__go_1_0_29, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_1_0_29, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)}, __local_var_3))})})
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
})
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_1_0_29, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}))
}

func Call_Data_Set_toMap__556171355(v_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Map_Internal_Node {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Set_toUnfoldable__871784487(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable = dictUnfoldable_0_loop
_ = dictUnfoldable_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_List_toUnfoldable(), gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(dictUnfoldable_0)})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_30 gopurs_runtime.Value
_ = go__go_3_1_30
go__go_3_1_30 = gopurs_runtime.Func2(func(__local_var_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr == nil) {
__t2 = __local_var_5
goto end_branch_2
} else {

}
}
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr != nil) {
__t2 = gopurs_runtime.UncurriedApp2(go__go_3_1_30, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_3_1_30, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V5)}, __local_var_5))})})
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
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_3_1_30, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
}



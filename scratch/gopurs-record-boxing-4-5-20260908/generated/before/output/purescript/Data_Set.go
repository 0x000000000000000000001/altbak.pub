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
		cache_Data_Set_identity = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
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
		cache_Data_Set_size = gopurs_runtime.Int(Get_Data_Map_Internal_size().IntVal)
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
return Call_Data_Set_member(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
})
	})
	return cache_Data_Set_member
}

var cache_Data_Set_isEmpty gopurs_runtime.Value
var once_Data_Set_isEmpty sync.Once
func Get_Data_Set_isEmpty() gopurs_runtime.Value {
	once_Data_Set_isEmpty.Do(func() {
		cache_Data_Set_isEmpty = gopurs_runtime.Bool((Get_Data_Map_Internal_isEmpty().IntVal) != (0))
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
		cache_Data_Set_fromMap = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Set_fromMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_0_box)))}
})
	})
	return cache_Data_Set_fromMap
}

var cache_Data_Set_foldableSet gopurs_runtime.Value
var once_Data_Set_foldableSet sync.Once
func Get_Data_Set_foldableSet() gopurs_runtime.Value {
	once_Data_Set_foldableSet.Do(func() {
		cache_Data_Set_foldableSet = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_3596835815_1680800814((&Constructor_Data_Foldable_Foldable[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_foldableList()).V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))}, f_1)
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_4 gopurs_runtime.Value
_ = go__go_4_1_4
// FALLBACK TCO: isLoop=false len=1
go__go_4_1_4 = gopurs_runtime.Func2(func(m_prime__5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__5)
if (__t_tag_2 == nil) {
__t4 = __local_var_6
goto end_branch_4
} else {

}
}
{
var __t_tag_3 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__5)
if (__t_tag_3 != nil) {
__t4 = gopurs_runtime.UncurriedApp2(go__go_4_1_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_4_1_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__5.UnsafePtr).V5)}, __local_var_6))}))})
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
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_4_1_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})))})
})
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_Set_go__go_2_5_5 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_Set_go__go_2_5_5
var go__go_2_5_5 gopurs_runtime.Value
_ = go__go_2_5_5
Call_local_Data_Set_go__go_2_5_5 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_5_5:
for {
if false { continue go__go_2_5_5 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var __t6 gopurs_runtime.Value
{
if (v_4 == nil) {
__t6 = b_3
goto end_branch_6
} else {

}
}
{
if (v_4 != nil) {
b_3_loop = gopurs_runtime.Apply2(f_0, b_3, (v_4).V0)
v_4_loop = (v_4).V1
continue go__go_2_5_5
__t6 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}
}
go__go_2_5_5 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_Set_go__go_2_5_5(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
// TAST (Let): __local_var_3_7 shape=App(Other) bindingType=(TypeVar b)
__local_var_3_7 := gopurs_runtime.Apply(go__go_2_5_5, x_1)
_ = __local_var_3_7
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_8_6 gopurs_runtime.Value
_ = go__go_5_8_6
// FALLBACK TCO: isLoop=false len=1
go__go_5_8_6 = gopurs_runtime.Func2(func(m_prime__6 gopurs_runtime.Value, __local_var_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
{
var __t_tag_9 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__6)
if (__t_tag_9 == nil) {
__t11 = __local_var_7
goto end_branch_11
} else {

}
}
{
var __t_tag_10 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__6)
if (__t_tag_10 != nil) {
__t11 = gopurs_runtime.UncurriedApp2(go__go_5_8_6, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__6.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__6.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_5_8_6, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__6.UnsafePtr).V5)}, __local_var_7))}))})
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
})
return gopurs_runtime.Apply(__local_var_3_7, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_5_8_6, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_4))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})))})
})
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_12 shape=App(Other) bindingType=Any
__local_var_2_12 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_foldableList()).V2), f_0, x_1)
_ = __local_var_2_12
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_13_7 gopurs_runtime.Value
_ = go__go_4_13_7
// FALLBACK TCO: isLoop=false len=1
go__go_4_13_7 = gopurs_runtime.Func2(func(m_prime__5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 gopurs_runtime.Value
{
var __t_tag_14 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__5)
if (__t_tag_14 == nil) {
__t16 = __local_var_6
goto end_branch_16
} else {

}
}
{
var __t_tag_15 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__5)
if (__t_tag_15 != nil) {
__t16 = gopurs_runtime.UncurriedApp2(go__go_4_13_7, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_4_13_7, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__5.UnsafePtr).V5)}, __local_var_6))}))})
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
return __t16
})
return gopurs_runtime.Apply(__local_var_2_12, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_4_13_7, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})))})
})
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
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
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
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
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
var go__go_1_2_12 gopurs_runtime.Value
_ = go__go_1_2_12
// FALLBACK TCO: isLoop=false len=1
go__go_1_2_12 = gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_4_3 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_2))
_ = v_4_3
var __t5 bool
{
if (v_4_3 != nil) {
// TAST (Let): v2_5_4 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v2_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_3))
_ = v2_5_4
__t5 = ((v2_5_4 != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (v_4_3).V0, (v2_5_4).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(go__go_1_2_12, (v_4_3).V2, (v2_5_4).V2).IntVal) != (0)))
goto end_branch_5
} else {

}
}
{
if (v_4_3 == nil) {
__t5 = true
goto end_branch_5
} else {

}
}
{
__t5 = func() bool { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Bool(__t5)
})
// TAST (Let): eqMapIter2_1_1 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Map","Internal","MapIter"] [(TypeVar k), (TypeVar v)])])
eqMapIter2_1_1 := (&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, go__go_1_2_12})
_ = eqMapIter2_1_1
// TAST (Let): eqMap_1_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Map","Internal","Map"] [(TypeVar a), Unit])])
eqMap_1_0 := (&Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(xs_2 gopurs_runtime.Value, ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 bool
{
var __t_tag_6 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_2)
if (__t_tag_6 == nil) {
var __t_tag_7 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_3)
__t10 = (__t_tag_7 == nil)
goto end_branch_10
} else {

}
}
{
var __t_tag_8 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_2)
if (__t_tag_8 != nil) {
var __t_tag_9 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_3)
__t10 = ((__t_tag_9 != nil)) && ((((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(xs_2.UnsafePtr).V1) == ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(ys_3.UnsafePtr).V1)) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqMapIter2_1_1.V0), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_2), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}}))}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_3), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}}))}).IntVal) != (0)))
goto end_branch_10
} else {

}
}
{
__t10 = func() bool { panic("Failed pattern match") }()
}
end_branch_10:
return gopurs_runtime.Bool(__t10)
})})
_ = eqMap_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(eqMap_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v1_3))}).IntVal) != (0))
})
})})))}
	})
	return cache_Data_Set_eq1Set
}

var cache_Data_Set_ord1Set gopurs_runtime.Value
var once_Data_Set_ord1Set sync.Once
func Get_Data_Set_ord1Set() gopurs_runtime.Value {
	once_Data_Set_ord1Set.Do(func() {
		cache_Data_Set_ord1Set = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_849406934_3985601471((&Constructor_Data_Ord_Ord1[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_3691144502_1766074591(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Set_eq1Set())))}
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): ordList_1_0 shape=App(Var) bindingType=(ADT ["Data","Ord","Ord"] [(ADT ["Data","List","Types","List"] [(TypeVar a)])])
ordList_1_0 := Rebox_Data_Set_4177771502_4210054658(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Types_ordList(), dictOrd_0)))
_ = ordList_1_0
return gopurs_runtime.Func2(func(s1_2 gopurs_runtime.Value, s2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_13 gopurs_runtime.Value
_ = go__go_4_1_13
// FALLBACK TCO: isLoop=false len=1
go__go_4_1_13 = gopurs_runtime.Func2(func(m_prime__5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__5)
if (__t_tag_2 == nil) {
__t4 = __local_var_6
goto end_branch_4
} else {

}
}
{
var __t_tag_3 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__5)
if (__t_tag_3 != nil) {
__t4 = gopurs_runtime.UncurriedApp2(go__go_4_1_13, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_4_1_13, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__5.UnsafePtr).V5)}, __local_var_6))}))})
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
var go__go_4_5_14 gopurs_runtime.Value
_ = go__go_4_5_14
// FALLBACK TCO: isLoop=false len=1
go__go_4_5_14 = gopurs_runtime.Func2(func(m_prime__5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
var __t_tag_6 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__5)
if (__t_tag_6 == nil) {
__t8 = __local_var_6
goto end_branch_8
} else {

}
}
{
var __t_tag_7 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__5)
if (__t_tag_7 != nil) {
__t8 = gopurs_runtime.UncurriedApp2(go__go_4_5_14, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_4_5_14, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__5.UnsafePtr).V5)}, __local_var_6))}))})
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
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(ordList_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_4_1_13, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](s1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_4_5_14, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](s2_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})))}).IntVal)), UnsafePtr: nil}
})
})})))}
	})
	return cache_Data_Set_ord1Set
}

var cache_Data_Set_empty gopurs_runtime.Value
var once_Data_Set_empty sync.Once
func Get_Data_Set_empty() gopurs_runtime.Value {
	once_Data_Set_empty.Do(func() {
		cache_Data_Set_empty = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))}
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
// TAST (Let): compare_1_0 shape=Other bindingType=(Func [(TypeVar k), (TypeVar k)] (ADT ["Data","Ordering","Ordering"] []))
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_0, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
}

func Call_Data_Set_toggle(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var v_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_2_loop
_ = v_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply4(Get_Data_Map_Internal_alter(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_3)
if (__t_tag_0 == nil) {
__t2 = (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, Get_Data_Unit_unit()})
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_3)
if (__t_tag_1 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
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
// FALLBACK TCO: isLoop=false len=1
go__go_1_0_0 = gopurs_runtime.Func2(func(m_prime__2 gopurs_runtime.Value, __local_var_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__2)
if (__t_tag_1 == nil) {
__t3 = __local_var_3
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__2)
if (__t_tag_2 != nil) {
__t3 = gopurs_runtime.UncurriedApp2(go__go_1_0_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__2.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_1_0_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__2.UnsafePtr).V5)}, __local_var_3))}))})
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
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_1_0_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}))
}

func Call_Data_Set_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
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
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_2351501316_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
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
var go__go_3_4_1 gopurs_runtime.Value
_ = go__go_3_4_1
// FALLBACK TCO: isLoop=false len=1
go__go_3_4_1 = gopurs_runtime.Func2(func(m_prime__4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t7 = gopurs_runtime.UncurriedApp2(go__go_3_4_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__4.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_3_4_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__4.UnsafePtr).V5)}, __local_var_5))}))})
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
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_3_4_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})))})
})
}

func Call_Data_Set_singleton(a_0_loop gopurs_runtime.Value) *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return (&Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{1, int64(1), int64(1), a_0, Get_Data_Unit_unit(), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil), (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil)})
}

func Call_Data_Set_showSet(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): showArray_1_0 shape=LitRecord bindingType=(ADT ["Data","Show","Show"] [(Array (TypeVar a))])
showArray_1_0 := (&Constructor_Data_Show_Show[[]gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Show_showArrayImpl(), gopurs_runtime.RecordGet(dictShow_0, "show"))})
_ = showArray_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_2638796135_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_4_2 gopurs_runtime.Value
_ = go__go_3_4_2
// FALLBACK TCO: isLoop=false len=1
go__go_3_4_2 = gopurs_runtime.Func2(func(m_prime__4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t7 = gopurs_runtime.UncurriedApp2(go__go_3_4_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__4.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_3_4_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__4.UnsafePtr).V5)}, __local_var_5))}))})
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
return gopurs_runtime.Str((("(fromFoldable ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showArray_1_0.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_Unfoldable_unfoldableArray()).V1), gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_3)
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
var __t_tag_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_3)
if (__t_tag_2 != nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_2351501316_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(xs_3.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(xs_3.UnsafePtr).V1)}}
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
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_3_4_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](s_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})))}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).StrVal())) + (")"))
})})))}
}

func Call_Data_Set_semigroupSet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 shape=Other bindingType=(Func [(TypeVar k), (TypeVar k)] (ADT ["Data","Ordering","Ordering"] []))
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_3854229351_4179793454((&Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_0, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})})))}
}

func Call_Data_Set_member(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var Call_local_Data_Set_go__go_2_0_3 func(*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) bool
_ = Call_local_Data_Set_go__go_2_0_3
var go__go_2_0_3 gopurs_runtime.Value
_ = go__go_2_0_3
Call_local_Data_Set_go__go_2_0_3 = func(v_3_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) bool {
go__go_2_0_3:
for {
if false { continue go__go_2_0_3 }
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
continue go__go_2_0_3
__t2 = func() bool { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
if (v1_4_1 == 380165415) {
v_3_loop = (v_3).V5
continue go__go_2_0_3
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
go__go_2_0_3 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_local_Data_Set_go__go_2_0_3(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_3_loop_val)))
})
return go__go_2_0_3
}

func Call_Data_Set_intersection(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 shape=Other bindingType=(Func [(TypeVar k), (TypeVar k)] (ADT ["Data","Ordering","Ordering"] []))
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeIntersectionWith(), compare_1_0, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
}

func Call_Data_Set_insert(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var v_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_2_loop
_ = v_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply4(Get_Data_Map_Internal_insert(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, a_1, Get_Data_Unit_unit(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_2)}))
}

func Call_Data_Set_fromMap(x_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Set_findMin(v_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var v_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [key: (TypeVar a), value: Unit] Any))])
__local_var_1_0 := Rebox_Data_Set_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_findMin(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)})))
_ = __local_var_1_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_1_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(__local_var_1_0).V0.key, true}
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
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Set_findMax(v_0_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var v_0 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [key: (TypeVar a), value: Unit] Any))])
__local_var_1_0 := Rebox_Data_Set_3094389156_2758605161(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_findMax(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)})))
_ = __local_var_1_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_1_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(__local_var_1_0).V0.key, true}
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
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
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
var go__go_1_2_8 gopurs_runtime.Value
_ = go__go_1_2_8
// FALLBACK TCO: isLoop=false len=1
go__go_1_2_8 = gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_4_3 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_2))
_ = v_4_3
var __t5 bool
{
if (v_4_3 != nil) {
// TAST (Let): v2_5_4 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v2_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_3))
_ = v2_5_4
__t5 = ((v2_5_4 != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (v_4_3).V0, (v2_5_4).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(go__go_1_2_8, (v_4_3).V2, (v2_5_4).V2).IntVal) != (0)))
goto end_branch_5
} else {

}
}
{
if (v_4_3 == nil) {
__t5 = true
goto end_branch_5
} else {

}
}
{
__t5 = func() bool { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Bool(__t5)
})
// TAST (Let): eqMapIter2_1_1 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Map","Internal","MapIter"] [(TypeVar k), (TypeVar v)])])
eqMapIter2_1_1 := (&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, go__go_1_2_8})
_ = eqMapIter2_1_1
// TAST (Let): eqMap_1_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Map","Internal","Map"] [(TypeVar a), Unit])])
eqMap_1_0 := (&Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(xs_2 gopurs_runtime.Value, ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 bool
{
var __t_tag_6 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_2)
if (__t_tag_6 == nil) {
var __t_tag_7 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_3)
__t10 = (__t_tag_7 == nil)
goto end_branch_10
} else {

}
}
{
var __t_tag_8 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_2)
if (__t_tag_8 != nil) {
var __t_tag_9 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_3)
__t10 = ((__t_tag_9 != nil)) && ((((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(xs_2.UnsafePtr).V1) == ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(ys_3.UnsafePtr).V1)) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqMapIter2_1_1.V0), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_2), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}}))}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_3), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}}))}).IntVal) != (0)))
goto end_branch_10
} else {

}
}
{
__t10 = func() bool { panic("Failed pattern match") }()
}
end_branch_10:
return gopurs_runtime.Bool(__t10)
})})
_ = eqMap_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_1444241223_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(eqMap_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v1_3))}).IntVal) != (0))
})})))}
}

func Call_Data_Set_ordSet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): ordList_1_0 shape=App(Var) bindingType=(ADT ["Data","Ord","Ord"] [(ADT ["Data","List","Types","List"] [(TypeVar a)])])
ordList_1_0 := Rebox_Data_Set_4177771502_4210054658(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Types_ordList(), dictOrd_0)))
_ = ordList_1_0
// TAST (Let): __local_var_2_2 shape=App(Other) bindingType=Any
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_2
var go__go_3_5_9 gopurs_runtime.Value
_ = go__go_3_5_9
// FALLBACK TCO: isLoop=false len=1
go__go_3_5_9 = gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_6_6 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_4))
_ = v_6_6
var __t8 bool
{
if (v_6_6 != nil) {
// TAST (Let): v2_7_7 shape=App(Var) bindingType=(ADT ["Data","Map","Internal","MapIterStep"] [(TypeVar k), (TypeVar v)])
v2_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_5))
_ = v2_7_7
__t8 = ((v2_7_7 != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_2, "eq"), (v_6_6).V0, (v2_7_7).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(go__go_3_5_9, (v_6_6).V2, (v2_7_7).V2).IntVal) != (0)))
goto end_branch_8
} else {

}
}
{
if (v_6_6 == nil) {
__t8 = true
goto end_branch_8
} else {

}
}
{
__t8 = func() bool { panic("Failed pattern match") }()
}
end_branch_8:
return gopurs_runtime.Bool(__t8)
})
// TAST (Let): eqMapIter2_3_4 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Map","Internal","MapIter"] [(TypeVar k), (TypeVar v)])])
eqMapIter2_3_4 := (&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, go__go_3_5_9})
_ = eqMapIter2_3_4
// TAST (Let): eqMap_3_3 shape=Let(LitRecord) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Map","Internal","Map"] [(TypeVar a), Unit])])
eqMap_3_3 := (&Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(xs_4 gopurs_runtime.Value, ys_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 bool
{
var __t_tag_9 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_4)
if (__t_tag_9 == nil) {
var __t_tag_10 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_5)
__t13 = (__t_tag_10 == nil)
goto end_branch_13
} else {

}
}
{
var __t_tag_11 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_4)
if (__t_tag_11 != nil) {
var __t_tag_12 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_5)
__t13 = ((__t_tag_12 != nil)) && ((((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(xs_4.UnsafePtr).V1) == ((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(ys_5.UnsafePtr).V1)) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(eqMapIter2_3_4.V0), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](xs_4), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}}))}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer((&Constructor_Data_Map_Internal_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](ys_5), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}}))}).IntVal) != (0)))
goto end_branch_13
} else {

}
}
{
__t13 = func() bool { panic("Failed pattern match") }()
}
end_branch_13:
return gopurs_runtime.Bool(__t13)
})})
_ = eqMap_3_3
// TAST (Let): eqSet1_2_1 shape=Let(Let(LitRecord)) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Map","Internal","Map"] [(TypeVar a), Unit])])
eqSet1_2_1 := (&Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(eqMap_3_3.V0), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_4))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v1_5))}).IntVal) != (0))
})})
_ = eqSet1_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_1910448679_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_1444241223_3790796878(eqSet1_2_1))}
}), gopurs_runtime.Func2(func(s1_3 gopurs_runtime.Value, s2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_14_10 gopurs_runtime.Value
_ = go__go_5_14_10
// FALLBACK TCO: isLoop=false len=1
go__go_5_14_10 = gopurs_runtime.Func2(func(m_prime__6 gopurs_runtime.Value, __local_var_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t17 gopurs_runtime.Value
{
var __t_tag_15 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__6)
if (__t_tag_15 == nil) {
__t17 = __local_var_7
goto end_branch_17
} else {

}
}
{
var __t_tag_16 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__6)
if (__t_tag_16 != nil) {
__t17 = gopurs_runtime.UncurriedApp2(go__go_5_14_10, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__6.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__6.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_5_14_10, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__6.UnsafePtr).V5)}, __local_var_7))}))})
goto end_branch_17
} else {

}
}
{
__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_17:
return __t17
})
var go__go_5_18_11 gopurs_runtime.Value
_ = go__go_5_18_11
// FALLBACK TCO: isLoop=false len=1
go__go_5_18_11 = gopurs_runtime.Func2(func(m_prime__6 gopurs_runtime.Value, __local_var_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t21 gopurs_runtime.Value
{
var __t_tag_19 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__6)
if (__t_tag_19 == nil) {
__t21 = __local_var_7
goto end_branch_21
} else {

}
}
{
var __t_tag_20 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__6)
if (__t_tag_20 != nil) {
__t21 = gopurs_runtime.UncurriedApp2(go__go_5_18_11, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__6.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__6.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_5_18_11, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__6.UnsafePtr).V5)}, __local_var_7))}))})
goto end_branch_21
} else {

}
}
{
__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_21:
return __t21
})
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(ordList_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_5_14_10, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](s1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_5_18_11, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](s2_4))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})))}).IntVal)), UnsafePtr: nil}
})})))}
}

func Call_Data_Set_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], dictOrd_1_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictOrd_1 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_1_loop
_ = dictOrd_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func2(func(m_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply4(Get_Data_Map_Internal_insert(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_1)}, a_3, Get_Data_Unit_unit(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_2))})))}
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
}

func Call_Data_Set_go__map(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var Call_local_Data_Set_go__go_2_0_15 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_Set_go__go_2_0_15
var go__go_2_0_15 gopurs_runtime.Value
_ = go__go_2_0_15
Call_local_Data_Set_go__go_2_0_15 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_0_15:
for {
if false { continue go__go_2_0_15 }
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
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply4(Get_Data_Map_Internal_insert(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, gopurs_runtime.Apply(f_1, (v_4).V0), Get_Data_Unit_unit(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](b_3))})))}
v_4_loop = (v_4).V1
continue go__go_2_0_15
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
go__go_2_0_15 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_Set_go__go_2_0_15(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
// TAST (Let): __local_var_3_2 shape=App(Other) bindingType=(TypeVar b)
__local_var_3_2 := gopurs_runtime.Apply(go__go_2_0_15, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
_ = __local_var_3_2
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_3_16 gopurs_runtime.Value
_ = go__go_5_3_16
// FALLBACK TCO: isLoop=false len=1
go__go_5_3_16 = gopurs_runtime.Func2(func(m_prime__6 gopurs_runtime.Value, __local_var_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
var __t_tag_4 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__6)
if (__t_tag_4 == nil) {
__t6 = __local_var_7
goto end_branch_6
} else {

}
}
{
var __t_tag_5 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__6)
if (__t_tag_5 != nil) {
__t6 = gopurs_runtime.UncurriedApp2(go__go_5_3_16, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__6.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__6.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_5_3_16, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__6.UnsafePtr).V5)}, __local_var_7))}))})
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
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_5_3_16, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_4))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})))})
})
}

func Call_Data_Set_mapMaybe(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_foldableList()).V2), gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
__local_var_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, a_2))
_ = __local_var_4_1
var __t2 gopurs_runtime.Value
{
if (__local_var_4_1 == nil) {
__t2 = acc_3
goto end_branch_2
} else {

}
}
{
if (__local_var_4_1 != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply4(Get_Data_Map_Internal_insert(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, (__local_var_4_1).V0, Get_Data_Unit_unit(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](acc_3))})))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__t2))}
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_3_17 gopurs_runtime.Value
_ = go__go_4_3_17
// FALLBACK TCO: isLoop=false len=1
go__go_4_3_17 = gopurs_runtime.Func2(func(m_prime__5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
var __t_tag_4 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__5)
if (__t_tag_4 == nil) {
__t6 = __local_var_6
goto end_branch_6
} else {

}
}
{
var __t_tag_5 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_prime__5)
if (__t_tag_5 != nil) {
__t6 = gopurs_runtime.UncurriedApp2(go__go_4_3_17, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_4_3_17, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(m_prime__5.UnsafePtr).V5)}, __local_var_6))}))})
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
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_4_3_17, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})))})
})
}

func Call_Data_Set_monoidSet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_1 shape=Other bindingType=(Func [(TypeVar k), (TypeVar k)] (ADT ["Data","Ordering","Ordering"] []))
compare_1_1 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_1
// TAST (Let): semigroupSet1_1_0 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(ADT ["Data","Map","Internal","Map"] [(TypeVar a), Unit])])
semigroupSet1_1_0 := (&Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_1, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})})
_ = semigroupSet1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_1291127239_1201789390((&Constructor_Data_Monoid_Monoid[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Set_3854229351_4179793454(semigroupSet1_1_0))}
}), gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})})))}
}

func Call_Data_Set_unions(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], dictOrd_1_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictOrd_1 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_1_loop
_ = dictOrd_1
// TAST (Let): compare_2_0 shape=Other bindingType=(Func [(TypeVar k), (TypeVar k)] (ADT ["Data","Ordering","Ordering"] []))
compare_2_0 := gopurs_runtime.Box(dictOrd_1.V1)
_ = compare_2_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func2(func(m1_3 gopurs_runtime.Value, m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_2_0, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_4))})))}
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(nil))})
}

func Call_Data_Set_difference(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 shape=Other bindingType=(Func [(TypeVar k), (TypeVar k)] (ADT ["Data","Ordering","Ordering"] []))
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeDifference(), compare_1_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m2_3))})))}
})
}

func Call_Data_Set_subset(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], s1_1_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value], s2_2_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var s1_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = s1_1_loop
_ = s1_1
var s2_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = s2_2_loop
_ = s2_2
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeDifference(), gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s1_1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s2_2)})
return (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 324739070 && __t_tag_0.UnsafePtr == nil)
}

func Call_Data_Set_properSubset(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], s1_1_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value], s2_2_loop *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) bool {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var s1_1 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = s1_1_loop
_ = s1_1
var s2_2 *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] = s2_2_loop
_ = s2_2
var __t0 int64
{
if (s1_1 == nil) {
__t0 = int64(0)
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
__t1 = int64(0)
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
__t1 = func() int64 { panic("Failed pattern match") }()
}
end_branch_1:
var __t_and_3 bool = false
if (gopurs_runtime.Bool((__t0) == (__t1)).IntVal) == (gopurs_runtime.Bool(false).IntVal) {

var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeDifference(), gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s1_1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(s2_2)})
__t_and_3 = (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 324739070 && __t_tag_2.UnsafePtr == nil)
}
return __t_and_3
}

func Call_Data_Set_go__delete(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Map_Internal_go__delete(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)})
}

func Call_Data_Set_checkValid(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Bool((gopurs_runtime.Apply(Get_Data_Map_Internal_checkValid(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}).IntVal) != (0))
}

func Call_Data_Set_catMaybes(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return Call_Data_Set_mapMaybe(dictOrd_0, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}

func Rebox_Data_Set_1291127239_1201789390(in *Constructor_Data_Monoid_Monoid[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}

func Rebox_Data_Set_1444241223_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
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

func Rebox_Data_Set_2351501316_138441832(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(in.V1)}
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

func Rebox_Data_Set_849406934_3985601471(in *Constructor_Data_Ord_Ord1[*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}



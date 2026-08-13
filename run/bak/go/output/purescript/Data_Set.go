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
		cache_Data_Set_fromMap = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Set_fromMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_0_box)))}
})
	})
	return cache_Data_Set_fromMap
}

var cache_Data_Set_foldableSet gopurs_runtime.Value
var once_Data_Set_foldableSet sync.Once
func Get_Data_Set_foldableSet() gopurs_runtime.Value {
	once_Data_Set_foldableSet.Do(func() {
		cache_Data_Set_foldableSet = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(&Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_2_1 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_1
var go__go_3_2_5 gopurs_runtime.Value
go__go_3_2_5 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_2_5:
for {
if false { continue go__go_3_2_5 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t3 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t3 = b_4
goto end_branch_3
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), b_4, gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0))
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_2_5
__t3 = gopurs_runtime.Value{}
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
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply(go__go_3_2_5, gopurs_runtime.RecordGet(dictMonoid_0, "mempty"))
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_4_6 gopurs_runtime.Value
_ = go__go_4_4_6
go__go_4_4_6 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t5 = gopurs_runtime.UncurriedApp2(go__go_4_4_6, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_4_6, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
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
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_4_6, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_7_7 gopurs_runtime.Value
go__go_2_7_7 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_7_7:
for {
if false { continue go__go_2_7_7 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t8 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t8 = b_3
goto end_branch_8
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
b_3_loop = gopurs_runtime.Apply2(f_0, b_3, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0)
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_7_7
__t8 = gopurs_runtime.Value{}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
}
}()
})
})
// TAST (Let): __local_var_2_6 -> gopurs_runtime.Value
__local_var_2_6 := gopurs_runtime.Apply(go__go_2_7_7, x_1)
_ = __local_var_2_6
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_9_8 gopurs_runtime.Value
_ = go__go_4_9_8
go__go_4_9_8 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t10 = __local_var_6
goto end_branch_10
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t10 = gopurs_runtime.UncurriedApp2(go__go_4_9_8, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_9_8, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
})
return gopurs_runtime.Apply(__local_var_2_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_9_8, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_13_9 gopurs_runtime.Value
go__go_2_13_9 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_13_9:
for {
if false { continue go__go_2_13_9 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t14 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t14 = b_3
goto end_branch_14
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
b_3_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, b_3)
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_13_9
__t14 = gopurs_runtime.Value{}
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
return __t14
}
}()
})
})
// TAST (Let): __local_var_2_12 -> gopurs_runtime.Value
__local_var_2_12 := gopurs_runtime.Apply(go__go_2_13_9, x_1)
_ = __local_var_2_12
var go__go_3_16_10 gopurs_runtime.Value
go__go_3_16_10 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_16_10:
for {
if false { continue go__go_3_16_10 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t17 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t17 = v_4
goto end_branch_17
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_16_10
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_17:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t17)}
}
}()
})
})
// TAST (Let): __local_var_3_15 -> gopurs_runtime.Value
__local_var_3_15 := gopurs_runtime.Apply(go__go_3_16_10, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
_ = __local_var_3_15
// TAST (Let): __local_var_2_11 -> gopurs_runtime.Value
__local_var_2_11 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_12, gopurs_runtime.Apply(__local_var_3_15, x_4))
})
_ = __local_var_2_11
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_18_11 gopurs_runtime.Value
_ = go__go_4_18_11
go__go_4_18_11 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t19 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t19 = __local_var_6
goto end_branch_19
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t19 = gopurs_runtime.UncurriedApp2(go__go_4_18_11, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_18_11, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
return __t19
})
return gopurs_runtime.Apply(__local_var_2_11, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_18_11, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
})
})})}
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
		cache_Data_Set_filter = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_filter(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), f_1_box)
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
		cache_Data_Set_eq1Set = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq1{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_19 gopurs_runtime.Value
go__go_1_2_19 = gopurs_runtime.Func(func(a_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_2_loop gopurs_runtime.Value = a_2_loop_val
var b_3_loop gopurs_runtime.Value = b_3_loop_val
go__go_1_2_19:
for {
if false { continue go__go_1_2_19 }
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
continue go__go_1_2_19
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
eqMapIter2_1_1 := &Constructor_Data_Eq_Eq{1, go__go_1_2_19}
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
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(eqMap_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v1_3))}).IntVal) != (0))
})
})
})})}
	})
	return cache_Data_Set_eq1Set
}

var cache_Data_Set_ord1Set gopurs_runtime.Value
var once_Data_Set_ord1Set sync.Once
func Get_Data_Set_ord1Set() gopurs_runtime.Value {
	once_Data_Set_ord1Set.Do(func() {
		cache_Data_Set_ord1Set = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord1{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1](Get_Data_Set_eq1Set()))}
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_2
// TAST (Let): eqList1_1_1 -> *Constructor_Data_Eq_Eq
eqList1_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_3_20 gopurs_runtime.Value
go__go_4_3_20 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
var v2_7_loop bool = (v2_7_loop_val.IntVal) != (0)
go__go_4_3_20:
for {
if false { continue go__go_4_3_20 }
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
continue go__go_4_3_20
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
return gopurs_runtime.Bool((gopurs_runtime.Apply3(go__go_4_3_20, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_3))}, gopurs_runtime.Bool(true)).IntVal) != (0))
})
})))
_ = eqList1_1_1
// TAST (Let): ordList_1_0 -> *Constructor_Data_Ord_Ord
ordList_1_0 := &Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqList1_1_1)}
}), gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_6_21 gopurs_runtime.Value
go__go_4_6_21 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_6_21:
for {
if false { continue go__go_4_6_21 }
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
continue go__go_4_6_21
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
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(go__go_4_6_21, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_3))}).IntVal)), UnsafePtr: nil}
})
})}
_ = ordList_1_0
return gopurs_runtime.Func(func(s1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_11_22 gopurs_runtime.Value
_ = go__go_4_11_22
go__go_4_11_22 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t12 = __local_var_6
goto end_branch_12
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t12 = gopurs_runtime.UncurriedApp2(go__go_4_11_22, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_11_22, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
})
var go__go_4_13_23 gopurs_runtime.Value
_ = go__go_4_13_23
go__go_4_13_23 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t14 = __local_var_6
goto end_branch_14
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t14 = gopurs_runtime.UncurriedApp2(go__go_4_13_23, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_13_23, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
return __t14
})
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(ordList_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_11_22, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_13_23, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s2_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}).IntVal)), UnsafePtr: nil}
})
})
})})}
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
		cache_Data_Set_delete = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_delete(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), k_1_box)
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
		cache_Data_Set_delete__4217907800 = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Set_delete__4217907800(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), k_1_box)
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

var cache_Data_Set_eq1Set__892523203 gopurs_runtime.Value
var once_Data_Set_eq1Set__892523203 sync.Once
func Get_Data_Set_eq1Set__892523203() gopurs_runtime.Value {
	once_Data_Set_eq1Set__892523203.Do(func() {
		cache_Data_Set_eq1Set__892523203 = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq1{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_39 gopurs_runtime.Value
go__go_1_2_39 = gopurs_runtime.Func(func(a_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_2_loop gopurs_runtime.Value = a_2_loop_val
var b_3_loop gopurs_runtime.Value = b_3_loop_val
go__go_1_2_39:
for {
if false { continue go__go_1_2_39 }
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
continue go__go_1_2_39
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
eqMapIter2_1_1 := &Constructor_Data_Eq_Eq{1, go__go_1_2_39}
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
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(eqMap_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v1_3))}).IntVal) != (0))
})
})
})})}
	})
	return cache_Data_Set_eq1Set__892523203
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
		cache_Data_Set_foldableSet__2661034898 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(&Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_2_1 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_1
var go__go_3_2_40 gopurs_runtime.Value
go__go_3_2_40 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_2_40:
for {
if false { continue go__go_3_2_40 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t3 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t3 = b_4
goto end_branch_3
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), b_4, gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0))
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_2_40
__t3 = gopurs_runtime.Value{}
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
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply(go__go_3_2_40, gopurs_runtime.RecordGet(dictMonoid_0, "mempty"))
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_4_41 gopurs_runtime.Value
_ = go__go_4_4_41
go__go_4_4_41 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t5 = gopurs_runtime.UncurriedApp2(go__go_4_4_41, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_4_41, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
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
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_4_41, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_7_42 gopurs_runtime.Value
go__go_2_7_42 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_7_42:
for {
if false { continue go__go_2_7_42 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t8 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t8 = b_3
goto end_branch_8
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
b_3_loop = gopurs_runtime.Apply2(f_0, b_3, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0)
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_7_42
__t8 = gopurs_runtime.Value{}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
}
}()
})
})
// TAST (Let): __local_var_2_6 -> gopurs_runtime.Value
__local_var_2_6 := gopurs_runtime.Apply(go__go_2_7_42, x_1)
_ = __local_var_2_6
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_9_43 gopurs_runtime.Value
_ = go__go_4_9_43
go__go_4_9_43 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t10 = __local_var_6
goto end_branch_10
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t10 = gopurs_runtime.UncurriedApp2(go__go_4_9_43, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_9_43, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
})
return gopurs_runtime.Apply(__local_var_2_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_9_43, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_13_44 gopurs_runtime.Value
go__go_2_13_44 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_13_44:
for {
if false { continue go__go_2_13_44 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t14 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t14 = b_3
goto end_branch_14
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
b_3_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, b_3)
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_13_44
__t14 = gopurs_runtime.Value{}
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
return __t14
}
}()
})
})
// TAST (Let): __local_var_2_12 -> gopurs_runtime.Value
__local_var_2_12 := gopurs_runtime.Apply(go__go_2_13_44, x_1)
_ = __local_var_2_12
var go__go_3_16_45 gopurs_runtime.Value
go__go_3_16_45 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_16_45:
for {
if false { continue go__go_3_16_45 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t17 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t17 = v_4
goto end_branch_17
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_16_45
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_17:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t17)}
}
}()
})
})
// TAST (Let): __local_var_3_15 -> gopurs_runtime.Value
__local_var_3_15 := gopurs_runtime.Apply(go__go_3_16_45, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
_ = __local_var_3_15
// TAST (Let): __local_var_2_11 -> gopurs_runtime.Value
__local_var_2_11 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_12, gopurs_runtime.Apply(__local_var_3_15, x_4))
})
_ = __local_var_2_11
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_18_46 gopurs_runtime.Value
_ = go__go_4_18_46
go__go_4_18_46 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t19 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t19 = __local_var_6
goto end_branch_19
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t19 = gopurs_runtime.UncurriedApp2(go__go_4_18_46, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_18_46, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
return __t19
})
return gopurs_runtime.Apply(__local_var_2_11, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_18_46, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
})
})})}
	})
	return cache_Data_Set_foldableSet__2661034898
}

var cache_Data_Set_foldableSet__3201980275 gopurs_runtime.Value
var once_Data_Set_foldableSet__3201980275 sync.Once
func Get_Data_Set_foldableSet__3201980275() gopurs_runtime.Value {
	once_Data_Set_foldableSet__3201980275.Do(func() {
		cache_Data_Set_foldableSet__3201980275 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(&Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_2_1 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_1
var go__go_3_2_47 gopurs_runtime.Value
go__go_3_2_47 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_2_47:
for {
if false { continue go__go_3_2_47 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t3 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t3 = b_4
goto end_branch_3
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), b_4, gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0))
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_2_47
__t3 = gopurs_runtime.Value{}
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
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply(go__go_3_2_47, gopurs_runtime.RecordGet(dictMonoid_0, "mempty"))
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_4_48 gopurs_runtime.Value
_ = go__go_4_4_48
go__go_4_4_48 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t5 = gopurs_runtime.UncurriedApp2(go__go_4_4_48, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_4_48, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
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
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_4_48, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_7_49 gopurs_runtime.Value
go__go_2_7_49 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_7_49:
for {
if false { continue go__go_2_7_49 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t8 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t8 = b_3
goto end_branch_8
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
b_3_loop = gopurs_runtime.Apply2(f_0, b_3, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0)
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_7_49
__t8 = gopurs_runtime.Value{}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
}
}()
})
})
// TAST (Let): __local_var_2_6 -> gopurs_runtime.Value
__local_var_2_6 := gopurs_runtime.Apply(go__go_2_7_49, x_1)
_ = __local_var_2_6
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_9_50 gopurs_runtime.Value
_ = go__go_4_9_50
go__go_4_9_50 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t10 = __local_var_6
goto end_branch_10
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t10 = gopurs_runtime.UncurriedApp2(go__go_4_9_50, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_9_50, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
})
return gopurs_runtime.Apply(__local_var_2_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_9_50, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_13_51 gopurs_runtime.Value
go__go_2_13_51 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_13_51:
for {
if false { continue go__go_2_13_51 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t14 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t14 = b_3
goto end_branch_14
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
b_3_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, b_3)
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_13_51
__t14 = gopurs_runtime.Value{}
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
return __t14
}
}()
})
})
// TAST (Let): __local_var_2_12 -> gopurs_runtime.Value
__local_var_2_12 := gopurs_runtime.Apply(go__go_2_13_51, x_1)
_ = __local_var_2_12
var go__go_3_16_52 gopurs_runtime.Value
go__go_3_16_52 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_16_52:
for {
if false { continue go__go_3_16_52 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t17 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t17 = v_4
goto end_branch_17
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_16_52
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_17
} else {

}
}
{
__t17 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_17:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t17)}
}
}()
})
})
// TAST (Let): __local_var_3_15 -> gopurs_runtime.Value
__local_var_3_15 := gopurs_runtime.Apply(go__go_3_16_52, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
_ = __local_var_3_15
// TAST (Let): __local_var_2_11 -> gopurs_runtime.Value
__local_var_2_11 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_12, gopurs_runtime.Apply(__local_var_3_15, x_4))
})
_ = __local_var_2_11
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_18_53 gopurs_runtime.Value
_ = go__go_4_18_53
go__go_4_18_53 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t19 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t19 = __local_var_6
goto end_branch_19
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t19 = gopurs_runtime.UncurriedApp2(go__go_4_18_53, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_18_53, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
return __t19
})
return gopurs_runtime.Apply(__local_var_2_11, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_18_53, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
})
})})}
	})
	return cache_Data_Set_foldableSet__3201980275
}

var cache_Data_Set_fromMap__556171355 gopurs_runtime.Value
var once_Data_Set_fromMap__556171355 sync.Once
func Get_Data_Set_fromMap__556171355() gopurs_runtime.Value {
	once_Data_Set_fromMap__556171355.Do(func() {
		cache_Data_Set_fromMap__556171355 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Set_fromMap__556171355(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_0_box)))}
})
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

var cache_Data_Set_ord1Set__3128423011 gopurs_runtime.Value
var once_Data_Set_ord1Set__3128423011 sync.Once
func Get_Data_Set_ord1Set__3128423011() gopurs_runtime.Value {
	once_Data_Set_ord1Set__3128423011.Do(func() {
		cache_Data_Set_ord1Set__3128423011 = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord1{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1](Get_Data_Set_eq1Set()))}
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_2
// TAST (Let): eqList1_1_1 -> *Constructor_Data_Eq_Eq
eqList1_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_3_64 gopurs_runtime.Value
go__go_4_3_64 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
var v2_7_loop bool = (v2_7_loop_val.IntVal) != (0)
go__go_4_3_64:
for {
if false { continue go__go_4_3_64 }
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
continue go__go_4_3_64
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
return gopurs_runtime.Bool((gopurs_runtime.Apply3(go__go_4_3_64, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_3))}, gopurs_runtime.Bool(true)).IntVal) != (0))
})
})))
_ = eqList1_1_1
// TAST (Let): ordList_1_0 -> *Constructor_Data_Ord_Ord
ordList_1_0 := &Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqList1_1_1)}
}), gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_6_65 gopurs_runtime.Value
go__go_4_6_65 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_6_65:
for {
if false { continue go__go_4_6_65 }
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
continue go__go_4_6_65
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
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(go__go_4_6_65, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_3))}).IntVal)), UnsafePtr: nil}
})
})}
_ = ordList_1_0
return gopurs_runtime.Func(func(s1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_11_66 gopurs_runtime.Value
_ = go__go_4_11_66
go__go_4_11_66 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t12 = __local_var_6
goto end_branch_12
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t12 = gopurs_runtime.UncurriedApp2(go__go_4_11_66, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_11_66, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
})
var go__go_4_13_67 gopurs_runtime.Value
_ = go__go_4_13_67
go__go_4_13_67 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 gopurs_runtime.Value
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr == nil) {
__t14 = __local_var_6
goto end_branch_14
} else {

}
}
{
if (__local_var_5.Type == 9 && __local_var_5.IntVal == 324739070 && __local_var_5.UnsafePtr != nil) {
__t14 = gopurs_runtime.UncurriedApp2(go__go_4_13_67, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_13_67, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
return __t14
})
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(ordList_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_11_66, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_13_67, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s2_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}).IntVal)), UnsafePtr: nil}
})
})
})})}
	})
	return cache_Data_Set_ord1Set__3128423011
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
// TAST (Let): v_3_0 -> gopurs_runtime.Value
v_3_0 := gopurs_runtime.UncurriedApp3(Get_Data_Map_Internal_unsafeSplit(), gopurs_runtime.Box(dictOrd_0.V1), a_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_2)})
_ = v_3_0
var __t4 *Constructor_Data_Maybe_Just
{
var __t_tag_2 *Constructor_Data_Maybe_Just = (*Constructor_Data_Map_Internal_Split)(v_3_0.UnsafePtr).V0
if (__t_tag_2 == nil) {
__t4 = &Constructor_Data_Maybe_Just{1, Get_Data_Unit_unit()}
goto end_branch_4
} else {

}
}
{
var __t_tag_3 *Constructor_Data_Maybe_Just = (*Constructor_Data_Map_Internal_Split)(v_3_0.UnsafePtr).V0
if (__t_tag_3 != nil) {
__t4 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
// TAST (Let): v2_4_1 -> *Constructor_Data_Maybe_Just
var v2_4_1 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)})
var __t5 *Constructor_Data_Map_Internal_Node
{
if (v2_4_1 == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_3_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_3_0.UnsafePtr).V2)}))
goto end_branch_5
} else {

}
}
{
if (v2_4_1 != nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), a_1, (v2_4_1).V0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_3_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Split)(v_3_0.UnsafePtr).V2)}))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return __t5
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
var go__go_3_6_1 gopurs_runtime.Value
_ = go__go_3_6_1
go__go_3_6_1 = gopurs_runtime.Func2(func(__local_var_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t7 = gopurs_runtime.UncurriedApp2(go__go_3_6_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_3_6_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V5)}, __local_var_5))})})
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
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_3_6_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
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
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_7_2 gopurs_runtime.Value
_ = go__go_3_7_2
go__go_3_7_2 = gopurs_runtime.Func2(func(__local_var_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr == nil) {
__t8 = __local_var_5
goto end_branch_8
} else {

}
}
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr != nil) {
__t8 = gopurs_runtime.UncurriedApp2(go__go_3_7_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_3_7_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V5)}, __local_var_5))})})
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
return gopurs_runtime.Str((("(fromFoldable ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showArray_1_0.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply6(Get_Data_Unfoldable_unfoldrArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr != nil) {
__t1 = (*Constructor_Data_Maybe_Just)(v_4.UnsafePtr).V0
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
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Maybe_Just
{
var __t_tag_3 *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_3)
if (__t_tag_3 == nil) {
__t5 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_3)
if (__t_tag_4 != nil) {
__t5 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("head", "tail", (*Constructor_Data_List_Types_Cons)(xs_3.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(xs_3.UnsafePtr).V1)})}
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
// TAST (Let): __local_var_4_2 -> *Constructor_Data_Maybe_Just
var __local_var_4_2 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)})
var __t6 *Constructor_Data_Maybe_Just
{
if (__local_var_4_2 != nil) {
__t6 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.RecordGet((__local_var_4_2).V0, "head"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet((__local_var_4_2).V0, "tail")))}})}}
goto end_branch_6
} else {

}
}
{
__t6 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t6)}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_3_7_2, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).StrVal())) + (")"))
})})}
}

func Call_Data_Set_semigroupSet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
var go__go_3_0_4 gopurs_runtime.Value
_ = go__go_3_0_4
go__go_3_0_4 = gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5)}))
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5)})))}))
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
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_4, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_2)}))
}

func Call_Data_Set_fromMap(x_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Map_Internal_Node {
var x_0 *Constructor_Data_Map_Internal_Node = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Set_findMin(v_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
// TAST (Let): __local_var_1_0 -> *Constructor_Data_Maybe_Just
__local_var_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Map_Internal_findMin(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}))
_ = __local_var_1_0
var __t1 *Constructor_Data_Maybe_Just
{
if (__local_var_1_0 != nil) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordGet((__local_var_1_0).V0, "key")}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
return __t1
}

func Call_Data_Set_findMax(v_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
// TAST (Let): __local_var_1_0 -> *Constructor_Data_Maybe_Just
__local_var_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Map_Internal_findMax(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}))
_ = __local_var_1_0
var __t1 *Constructor_Data_Maybe_Just
{
if (__local_var_1_0 != nil) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordGet((__local_var_1_0).V0, "key")}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
return __t1
}

func Call_Data_Set_filter(dictOrd_0_loop *Constructor_Data_Ord_Ord, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_0_12 gopurs_runtime.Value
_ = go__go_2_0_12
go__go_2_0_12 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_12, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_12, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)})))}))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_12, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_12, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)})))}))
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
return go__go_2_0_12
}

func Call_Data_Set_eqSet(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var go__go_1_2_13 gopurs_runtime.Value
go__go_1_2_13 = gopurs_runtime.Func(func(a_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_2_loop gopurs_runtime.Value = a_2_loop_val
var b_3_loop gopurs_runtime.Value = b_3_loop_val
go__go_1_2_13:
for {
if false { continue go__go_1_2_13 }
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
continue go__go_1_2_13
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
eqMapIter2_1_1 := &Constructor_Data_Eq_Eq{1, go__go_1_2_13}
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

func Call_Data_Set_ordSet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_2
// TAST (Let): eqList1_1_1 -> *Constructor_Data_Eq_Eq
eqList1_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_3_14 gopurs_runtime.Value
go__go_4_3_14 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
var v2_7_loop bool = (v2_7_loop_val.IntVal) != (0)
go__go_4_3_14:
for {
if false { continue go__go_4_3_14 }
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
continue go__go_4_3_14
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
return gopurs_runtime.Bool((gopurs_runtime.Apply3(go__go_4_3_14, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_3))}, gopurs_runtime.Bool(true)).IntVal) != (0))
})
})))
_ = eqList1_1_1
// TAST (Let): ordList_1_0 -> *Constructor_Data_Ord_Ord
ordList_1_0 := &Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqList1_1_1)}
}), gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_6_15 gopurs_runtime.Value
go__go_4_6_15 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_6_15:
for {
if false { continue go__go_4_6_15 }
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
continue go__go_4_6_15
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
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(go__go_4_6_15, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_3))}).IntVal)), UnsafePtr: nil}
})
})}
_ = ordList_1_0
// TAST (Let): __local_var_2_12 -> gopurs_runtime.Value
__local_var_2_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_12
var go__go_3_15_16 gopurs_runtime.Value
go__go_3_15_16 = gopurs_runtime.Func(func(a_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_4_loop gopurs_runtime.Value = a_4_loop_val
var b_5_loop gopurs_runtime.Value = b_5_loop_val
go__go_3_15_16:
for {
if false { continue go__go_3_15_16 }
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
continue go__go_3_15_16
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
eqMapIter2_3_14 := &Constructor_Data_Eq_Eq{1, go__go_3_15_16}
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
var go__go_5_23_17 gopurs_runtime.Value
_ = go__go_5_23_17
go__go_5_23_17 = gopurs_runtime.Func2(func(__local_var_6 gopurs_runtime.Value, __local_var_7 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t24 = gopurs_runtime.UncurriedApp2(go__go_5_23_17, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_5_23_17, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V5)}, __local_var_7))})})
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
var go__go_5_25_18 gopurs_runtime.Value
_ = go__go_5_25_18
go__go_5_25_18 = gopurs_runtime.Func2(func(__local_var_6 gopurs_runtime.Value, __local_var_7 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t26 = gopurs_runtime.UncurriedApp2(go__go_5_25_18, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_5_25_18, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_6.UnsafePtr).V5)}, __local_var_7))})})
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
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(ordList_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_5_23_17, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s1_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_5_25_18, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](s2_4))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}).IntVal)), UnsafePtr: nil}
})
})})}
}

func Call_Data_Set_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictOrd_1_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictOrd_1 *Constructor_Data_Ord_Ord = dictOrd_1_loop
_ = dictOrd_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_0_24 gopurs_runtime.Value
_ = go__go_4_0_24
go__go_4_0_24 = gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (v1_5.Type == 9 && v1_5.IntVal == 324739070 && v1_5.UnsafePtr == nil) {
__t3 = &Constructor_Data_Map_Internal_Node{1, 1, 1, a_3, Get_Data_Unit_unit(), (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 324739070 && v1_5.UnsafePtr != nil) {
// TAST (Let): v2_6_1 -> gopurs_runtime.Value
v2_6_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_1.V1), a_3, (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V2)
_ = v2_6_1
var __t2 *Constructor_Data_Map_Internal_Node
{
if (uint32(v2_6_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_4_0_24, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V5)}))
goto end_branch_2
} else {

}
}
{
if (uint32(v2_6_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_4_0_24, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V5)})))}))
goto end_branch_2
} else {

}
}
{
if (uint32(v2_6_1.IntVal) == 902936544) {
__t2 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V1, a_3, Get_Data_Unit_unit(), (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node)(v1_5.UnsafePtr).V5}
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
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_4_0_24, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_2))})))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
}

func Call_Data_Set_go__map(dictOrd_0_loop *Constructor_Data_Ord_Ord, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_1_25 gopurs_runtime.Value
go__go_2_1_25 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_1_25:
for {
if false { continue go__go_2_1_25 }
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
var go__go_6_3_26 gopurs_runtime.Value
_ = go__go_6_3_26
go__go_6_3_26 = gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_6_3_26, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V5)}))
goto end_branch_5
} else {

}
}
{
if (uint32(v2_8_4.IntVal) == 380165415) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_6_3_26, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_7.UnsafePtr).V5)})))}))
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
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_6_3_26, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](b_3))})))}
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_1_25
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
__local_var_2_0 := gopurs_runtime.Apply(go__go_2_1_25, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_8_27 gopurs_runtime.Value
_ = go__go_4_8_27
go__go_4_8_27 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t9 = gopurs_runtime.UncurriedApp2(go__go_4_8_27, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_8_27, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
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
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_8_27, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
}

func Call_Data_Set_mapMaybe(dictOrd_0_loop *Constructor_Data_Ord_Ord, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_2_28 gopurs_runtime.Value
go__go_2_2_28 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_2_28:
for {
if false { continue go__go_2_2_28 }
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
var go__go_7_5_29 gopurs_runtime.Value
_ = go__go_7_5_29
go__go_7_5_29 = gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_7_5_29, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V5)}))
goto end_branch_7
} else {

}
}
{
if (uint32(v2_9_6.IntVal) == 380165415) {
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_7_5_29, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V5)})))}))
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
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_7_5_29, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](b_3))}))
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
continue go__go_2_2_28
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
__local_var_2_1 := gopurs_runtime.Apply(go__go_2_2_28, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
_ = __local_var_2_1
var go__go_3_12_30 gopurs_runtime.Value
go__go_3_12_30 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_12_30:
for {
if false { continue go__go_3_12_30 }
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
continue go__go_3_12_30
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
__local_var_3_11 := gopurs_runtime.Apply(go__go_3_12_30, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
_ = __local_var_3_11
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_1, gopurs_runtime.Apply(__local_var_3_11, x_4))
})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_14_31 gopurs_runtime.Value
_ = go__go_4_14_31
go__go_4_14_31 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t15 = gopurs_runtime.UncurriedApp2(go__go_4_14_31, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_14_31, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
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
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_14_31, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
}

func Call_Data_Set_monoidSet(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_1 -> gopurs_runtime.Value
compare_1_1 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_1
// TAST (Let): semigroupSet1_1_0 -> *Constructor_Data_Semigroup_Semigroup
semigroupSet1_1_0 := &Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeUnionWith(), compare_1_1, Get_Data_Function_go__const(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m1_2))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m2_3))})))}
})
})}
_ = semigroupSet1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupSet1_1_0)}
}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))}})}
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

func Call_Data_Set_delete(dictOrd_0_loop *Constructor_Data_Ord_Ord, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_32 gopurs_runtime.Value
_ = go__go_2_0_32
go__go_2_0_32 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t3 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t2 *Constructor_Data_Map_Internal_Node
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_32, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)}))
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_32, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)})))}))
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)}))
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
return go__go_2_0_32
}

func Call_Data_Set_checkValid(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var go__go_1_0_33 gopurs_runtime.Value
_ = go__go_1_0_33
go__go_1_0_33 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t36 bool
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t36 = true
goto end_branch_36
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
var __t35 bool
{
var __t_tag_1 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4
if (__t_tag_1 == nil) {
var __t10 bool
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5
if (__t_tag_2 == nil) {
__t10 = true
goto end_branch_10
} else {

}
}
{
var __t_tag_3 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5
if (__t_tag_3 != nil) {
var __t_and_9 bool = false
if ((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V0) == (2) {

var __t_and_8 bool = false
if (((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5).V0) == (1) {

var __t4 bool
{
if ((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V1) > (((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5).V1) {
__t4 = true
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
var __t_and_7 bool = false
if __t4 {

var __t6 bool
{
var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), ((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5).V2, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2)
if (uint32(__t_tag_5.IntVal) == 380165415) {
__t6 = true
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
__t_and_7 = (__t6) && ((gopurs_runtime.Apply(go__go_1_0_33, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5)}).IntVal) != (0))
}
__t_and_8 = __t_and_7
}
__t_and_9 = __t_and_8
}
__t10 = __t_and_9
goto end_branch_10
} else {

}
}
{
__t10 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_10:
__t35 = __t10
goto end_branch_35
} else {

}
}
{
var __t_tag_11 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4
if (__t_tag_11 != nil) {
var __t34 bool
{
var __t_tag_12 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5
if (__t_tag_12 == nil) {
var __t_and_18 bool = false
if ((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V0) == (2) {

var __t_and_17 bool = false
if (((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4).V0) == (1) {

var __t13 bool
{
if ((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V1) > (((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4).V1) {
__t13 = true
goto end_branch_13
} else {

}
}
{
__t13 = false
}
end_branch_13:
var __t_and_16 bool = false
if __t13 {

var __t15 bool
{
var __t_tag_14 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), ((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4).V2, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2)
if (uint32(__t_tag_14.IntVal) == 1527465420) {
__t15 = true
goto end_branch_15
} else {

}
}
{
__t15 = false
}
end_branch_15:
__t_and_16 = (__t15) && ((gopurs_runtime.Apply(go__go_1_0_33, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4)}).IntVal) != (0))
}
__t_and_17 = __t_and_16
}
__t_and_18 = __t_and_17
}
__t34 = __t_and_18
goto end_branch_34
} else {

}
}
{
var __t_tag_19 *Constructor_Data_Map_Internal_Node = (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5
if (__t_tag_19 != nil) {
var __t20 bool
{
if ((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5).V0) {
__t20 = true
goto end_branch_20
} else {

}
}
{
__t20 = false
}
end_branch_20:
var __t_and_33 bool = false
if __t20 {

var __t22 bool
{
var __t_tag_21 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), ((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5).V2, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2)
if (uint32(__t_tag_21.IntVal) == 380165415) {
__t22 = true
goto end_branch_22
} else {

}
}
{
__t22 = false
}
end_branch_22:
var __t_and_32 bool = false
if __t22 {

var __t23 bool
{
if ((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V0) > (((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4).V0) {
__t23 = true
goto end_branch_23
} else {

}
}
{
__t23 = false
}
end_branch_23:
var __t_and_31 bool = false
if __t23 {

var __t25 bool
{
var __t_tag_24 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), ((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4).V2, (*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V2)
if (uint32(__t_tag_24.IntVal) == 1527465420) {
__t25 = true
goto end_branch_25
} else {

}
}
{
__t25 = false
}
end_branch_25:
var __t_and_30 bool = false
if __t25 {

var __t29 bool
{
// TAST (Let): __local_var_3_26 -> int64
__local_var_3_26 := (((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5).V0) - (((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4).V0)
_ = __local_var_3_26
var __t28 int64
{
var __t27 bool
{
if (gopurs_runtime.Int(__local_var_3_26).IntVal) < (gopurs_runtime.Int(0).IntVal) {
__t27 = false
goto end_branch_27
} else {

}
}
{
__t27 = true
}
end_branch_27:
if __t27 {
__t28 = __local_var_3_26
goto end_branch_28
} else {

}
}
{
__t28 = gopurs_runtime.Int(-(gopurs_runtime.Int(__local_var_3_26).IntVal)).IntVal
}
end_branch_28:
if (__t28) < (2) {
__t29 = true
goto end_branch_29
} else {

}
}
{
__t29 = false
}
end_branch_29:
__t_and_30 = (__t29) && (((((((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5).V1) + (((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4).V1)) + (1)) == ((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V1)) && (((gopurs_runtime.Apply(go__go_1_0_33, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V4)}).IntVal) != (0)) && ((gopurs_runtime.Apply(go__go_1_0_33, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_2.UnsafePtr).V5)}).IntVal) != (0))))
}
__t_and_31 = __t_and_30
}
__t_and_32 = __t_and_31
}
__t_and_33 = __t_and_32
}
__t34 = __t_and_33
goto end_branch_34
} else {

}
}
{
__t34 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_34:
__t35 = __t34
goto end_branch_35
} else {

}
}
{
__t35 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_35:
__t36 = __t35
goto end_branch_36
} else {

}
}
{
__t36 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_36:
return gopurs_runtime.Bool(__t36)
})
return go__go_1_0_33
}

func Call_Data_Set_catMaybes(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var go__go_1_2_34 gopurs_runtime.Value
go__go_1_2_34 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_2_34:
for {
if false { continue go__go_1_2_34 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t11 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t11 = b_2
goto end_branch_11
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
var __t10 *Constructor_Data_Map_Internal_Node
{
var __t_tag_3 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V0)
if (__t_tag_3 == nil) {
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](b_2)
goto end_branch_10
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V0)
if (__t_tag_4 != nil) {
// TAST (Let): __local_var_4_5 -> gopurs_runtime.Value
__local_var_4_5 := (*Constructor_Data_Maybe_Just)((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V0.UnsafePtr).V0
_ = __local_var_4_5
var go__go_5_6_35 gopurs_runtime.Value
_ = go__go_5_6_35
go__go_5_6_35 = gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 *Constructor_Data_Map_Internal_Node
{
if (v1_6.Type == 9 && v1_6.IntVal == 324739070 && v1_6.UnsafePtr == nil) {
__t9 = &Constructor_Data_Map_Internal_Node{1, 1, 1, __local_var_4_5, Get_Data_Unit_unit(), (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_9
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 324739070 && v1_6.UnsafePtr != nil) {
// TAST (Let): v2_7_7 -> gopurs_runtime.Value
v2_7_7 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), __local_var_4_5, (*Constructor_Data_Map_Internal_Node)(v1_6.UnsafePtr).V2)
_ = v2_7_7
var __t8 *Constructor_Data_Map_Internal_Node
{
if (uint32(v2_7_7.IntVal) == 1527465420) {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_6.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_6.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_5_6_35, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_6.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_6.UnsafePtr).V5)}))
goto end_branch_8
} else {

}
}
{
if (uint32(v2_7_7.IntVal) == 380165415) {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_6.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_6.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_6.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_5_6_35, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_6.UnsafePtr).V5)})))}))
goto end_branch_8
} else {

}
}
{
if (uint32(v2_7_7.IntVal) == 902936544) {
__t8 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v1_6.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v1_6.UnsafePtr).V1, __local_var_4_5, Get_Data_Unit_unit(), (*Constructor_Data_Map_Internal_Node)(v1_6.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node)(v1_6.UnsafePtr).V5}
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_8:
__t9 = __t8
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t9)}
})
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_5_6_35, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](b_2))}))
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_10:
b_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t10)}
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V1)}
continue go__go_1_2_34
__t11 = gopurs_runtime.Value{}
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
}
}()
})
})
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(go__go_1_2_34, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
_ = __local_var_1_1
var go__go_2_13_36 gopurs_runtime.Value
go__go_2_13_36 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_3_loop_val)
var v1_4_loop gopurs_runtime.Value = v1_4_loop_val
go__go_2_13_36:
for {
if false { continue go__go_2_13_36 }
var v_3 *Constructor_Data_List_Types_Cons = v_3_loop
_ = v_3
var v1_4 gopurs_runtime.Value = v1_4_loop
_ = v1_4
var __t14 *Constructor_Data_List_Types_Cons
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr == nil) {
__t14 = v_3
goto end_branch_14
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr != nil) {
v_3_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_4.UnsafePtr).V0, v_3})})
v1_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_4.UnsafePtr).V1)}
continue go__go_2_13_36
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_14:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t14)}
}
}()
})
})
// TAST (Let): __local_var_2_12 -> gopurs_runtime.Value
__local_var_2_12 := gopurs_runtime.Apply(go__go_2_13_36, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
_ = __local_var_2_12
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_1, gopurs_runtime.Apply(__local_var_2_12, x_3))
})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_15_37 gopurs_runtime.Value
_ = go__go_3_15_37
go__go_3_15_37 = gopurs_runtime.Func2(func(__local_var_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 gopurs_runtime.Value
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr == nil) {
__t16 = __local_var_5
goto end_branch_16
} else {

}
}
{
if (__local_var_4.Type == 9 && __local_var_4.IntVal == 324739070 && __local_var_4.UnsafePtr != nil) {
__t16 = gopurs_runtime.UncurriedApp2(go__go_3_15_37, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_3_15_37, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V5)}, __local_var_5))})})
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
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_3_15_37, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
}

func Call_Data_Set_delete__4217907800(dictOrd_0_loop *Constructor_Data_Ord_Ord, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_38 gopurs_runtime.Value
_ = go__go_2_0_38
go__go_2_0_38 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t3 = (*Constructor_Data_Map_Internal_Node)(nil)
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
// TAST (Let): v1_4_1 -> gopurs_runtime.Value
v1_4_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), k_1, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t2 *Constructor_Data_Map_Internal_Node
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_38, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)}))
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_2_0_38, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)})))}))
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp2(Get_Data_Map_Internal_unsafeJoinNodes(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_3.UnsafePtr).V5)}))
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
return go__go_2_0_38
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
// TAST (Let): __local_var_1_0 -> *Constructor_Data_Maybe_Just
__local_var_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Map_Internal_findMax(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}))
_ = __local_var_1_0
var __t1 *Constructor_Data_Maybe_Just
{
if (__local_var_1_0 != nil) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordGet((__local_var_1_0).V0, "key")}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
return __t1
}

func Call_Data_Set_findMin__611847841(v_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
// TAST (Let): __local_var_1_0 -> *Constructor_Data_Maybe_Just
__local_var_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_Map_Internal_findMin(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}))
_ = __local_var_1_0
var __t1 *Constructor_Data_Maybe_Just
{
if (__local_var_1_0 != nil) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordGet((__local_var_1_0).V0, "key")}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
return __t1
}

func Call_Data_Set_fromMap__556171355(x_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Map_Internal_Node {
var x_0 *Constructor_Data_Map_Internal_Node = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Set_insert__4217907800(dictOrd_0_loop *Constructor_Data_Ord_Ord, a_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Map_Internal_Node {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var v_2 *Constructor_Data_Map_Internal_Node = v_2_loop
_ = v_2
var go__go_3_0_54 gopurs_runtime.Value
_ = go__go_3_0_54
go__go_3_0_54 = gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_54, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5)}))
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_54, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5)})))}))
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
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_54, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_2)}))
}

func Call_Data_Set_insert__3935803192(dictOrd_0_loop *Constructor_Data_Ord_Ord, a_1_loop *Constructor_Data_Maybe_Just, v_2_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Map_Internal_Node {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var a_1 *Constructor_Data_Maybe_Just = a_1_loop
_ = a_1
var v_2 *Constructor_Data_Map_Internal_Node = v_2_loop
_ = v_2
var go__go_3_0_55 gopurs_runtime.Value
_ = go__go_3_0_55
go__go_3_0_55 = gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Map_Internal_Node
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr == nil) {
__t3 = &Constructor_Data_Map_Internal_Node{1, 1, 1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(a_1)}, Get_Data_Unit_unit(), (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 324739070 && v1_4.UnsafePtr != nil) {
// TAST (Let): v2_5_1 -> gopurs_runtime.Value
v2_5_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(a_1)}, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2)
_ = v2_5_1
var __t2 *Constructor_Data_Map_Internal_Node
{
if (uint32(v2_5_1.IntVal) == 1527465420) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_55, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5)}))
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_55, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5)})))}))
goto end_branch_2
} else {

}
}
{
if (uint32(v2_5_1.IntVal) == 902936544) {
__t2 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(a_1)}, Get_Data_Unit_unit(), (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node)(v1_4.UnsafePtr).V5}
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
return gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_3_0_55, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_2)}))
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
var go__go_2_2_56 gopurs_runtime.Value
go__go_2_2_56 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_2_56:
for {
if false { continue go__go_2_2_56 }
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
var go__go_7_5_57 gopurs_runtime.Value
_ = go__go_7_5_57
go__go_7_5_57 = gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_7_5_57, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V5)}))
goto end_branch_7
} else {

}
}
{
if (uint32(v2_9_6.IntVal) == 380165415) {
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_7_5_57, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V5)})))}))
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
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_7_5_57, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](b_3))}))
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
continue go__go_2_2_56
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
__local_var_2_1 := gopurs_runtime.Apply(go__go_2_2_56, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
_ = __local_var_2_1
var go__go_3_12_58 gopurs_runtime.Value
go__go_3_12_58 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_12_58:
for {
if false { continue go__go_3_12_58 }
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
continue go__go_3_12_58
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
__local_var_3_11 := gopurs_runtime.Apply(go__go_3_12_58, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
_ = __local_var_3_11
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_1, gopurs_runtime.Apply(__local_var_3_11, x_4))
})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_14_59 gopurs_runtime.Value
_ = go__go_4_14_59
go__go_4_14_59 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t15 = gopurs_runtime.UncurriedApp2(go__go_4_14_59, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_14_59, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
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
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_14_59, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
}

func Call_Data_Set_mapMaybe__4234899916(dictOrd_0_loop *Constructor_Data_Ord_Ord, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_2_60 gopurs_runtime.Value
go__go_2_2_60 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_2_60:
for {
if false { continue go__go_2_2_60 }
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
__local_var_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0))}))
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
var go__go_7_5_61 gopurs_runtime.Value
_ = go__go_7_5_61
go__go_7_5_61 = gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_7_5_61, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V5)}))
goto end_branch_7
} else {

}
}
{
if (uint32(v2_9_6.IntVal) == 380165415) {
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_7_5_61, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_8.UnsafePtr).V5)})))}))
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
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_7_5_61, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](b_3))}))
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
continue go__go_2_2_60
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
__local_var_2_1 := gopurs_runtime.Apply(go__go_2_2_60, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
_ = __local_var_2_1
var go__go_3_12_62 gopurs_runtime.Value
go__go_3_12_62 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_12_62:
for {
if false { continue go__go_3_12_62 }
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
continue go__go_3_12_62
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
__local_var_3_11 := gopurs_runtime.Apply(go__go_3_12_62, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
_ = __local_var_3_11
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_1, gopurs_runtime.Apply(__local_var_3_11, x_4))
})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_14_63 gopurs_runtime.Value
_ = go__go_4_14_63
go__go_4_14_63 = gopurs_runtime.Func2(func(__local_var_5 gopurs_runtime.Value, __local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t15 = gopurs_runtime.UncurriedApp2(go__go_4_14_63, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_14_63, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_5.UnsafePtr).V5)}, __local_var_6))})})
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
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_4_14_63, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
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
var go__go_1_0_68 gopurs_runtime.Value
_ = go__go_1_0_68
go__go_1_0_68 = gopurs_runtime.Func2(func(__local_var_2 gopurs_runtime.Value, __local_var_3 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = gopurs_runtime.UncurriedApp2(go__go_1_0_68, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_1_0_68, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)}, __local_var_3))})})
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
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_1_0_68, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}))
}

func Call_Data_Set_toList__2510070446(v_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_List_Types_Cons {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
var go__go_1_0_69 gopurs_runtime.Value
_ = go__go_1_0_69
go__go_1_0_69 = gopurs_runtime.Func2(func(__local_var_2 gopurs_runtime.Value, __local_var_3 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = gopurs_runtime.UncurriedApp2(go__go_1_0_69, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_1_0_69, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)}, __local_var_3))})})
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
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_1_0_69, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}))
}

func Call_Data_Set_toList__2521903733(v_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_List_Types_Cons {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
var go__go_1_0_70 gopurs_runtime.Value
_ = go__go_1_0_70
go__go_1_0_70 = gopurs_runtime.Func2(func(__local_var_2 gopurs_runtime.Value, __local_var_3 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = gopurs_runtime.UncurriedApp2(go__go_1_0_70, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_1_0_70, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_2.UnsafePtr).V5)}, __local_var_3))})})
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
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_1_0_70, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}))
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
var go__go_3_6_71 gopurs_runtime.Value
_ = go__go_3_6_71
go__go_3_6_71 = gopurs_runtime.Func2(func(__local_var_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t7 = gopurs_runtime.UncurriedApp2(go__go_3_6_71, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_3_6_71, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_4.UnsafePtr).V5)}, __local_var_5))})})
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
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_3_6_71, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
})
}



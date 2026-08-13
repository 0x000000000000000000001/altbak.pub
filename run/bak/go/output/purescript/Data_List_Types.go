package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_List_Types_identity gopurs_runtime.Value
var once_Data_List_Types_identity sync.Once
func Get_Data_List_Types_identity() gopurs_runtime.Value {
	once_Data_List_Types_identity.Do(func() {
		cache_Data_List_Types_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Types_identity(x_0_box)
})
	})
	return cache_Data_List_Types_identity
}

var cache_Data_List_Types_identity1 gopurs_runtime.Value
var once_Data_List_Types_identity1 sync.Once
func Get_Data_List_Types_identity1() gopurs_runtime.Value {
	once_Data_List_Types_identity1.Do(func() {
		cache_Data_List_Types_identity1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Types_identity1(x_0_box)
})
	})
	return cache_Data_List_Types_identity1
}

var cache_Data_List_Types_Nil gopurs_runtime.Value
var once_Data_List_Types_Nil sync.Once
func Get_Data_List_Types_Nil() gopurs_runtime.Value {
	once_Data_List_Types_Nil.Do(func() {
		cache_Data_List_Types_Nil = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}
	})
	return cache_Data_List_Types_Nil
}

var cache_Data_List_Types_Cons gopurs_runtime.Value
var once_Data_List_Types_Cons sync.Once
func Get_Data_List_Types_Cons() gopurs_runtime.Value {
	once_Data_List_Types_Cons.Do(func() {
		cache_Data_List_Types_Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](value1)})}
})
})
	})
	return cache_Data_List_Types_Cons
}

var cache_Data_List_Types_NonEmptyList gopurs_runtime.Value
var once_Data_List_Types_NonEmptyList sync.Once
func Get_Data_List_Types_NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_NonEmptyList.Do(func() {
		cache_Data_List_Types_NonEmptyList = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Types_NonEmptyList(x_0_box)
})
	})
	return cache_Data_List_Types_NonEmptyList
}

var cache_Data_List_Types_toList gopurs_runtime.Value
var once_Data_List_Types_toList sync.Once
func Get_Data_List_Types_toList() gopurs_runtime.Value {
	once_Data_List_Types_toList.Do(func() {
		cache_Data_List_Types_toList = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_Types_toList(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_0_box)))}
})
	})
	return cache_Data_List_Types_toList
}

var cache_Data_List_Types_newtypeNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_newtypeNonEmptyList sync.Once
func Get_Data_List_Types_newtypeNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_newtypeNonEmptyList.Do(func() {
		cache_Data_List_Types_newtypeNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(&Constructor_Data_Newtype_Newtype{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})}
	})
	return cache_Data_List_Types_newtypeNonEmptyList
}

var cache_Data_List_Types_nelCons gopurs_runtime.Value
var once_Data_List_Types_nelCons sync.Once
func Get_Data_List_Types_nelCons() gopurs_runtime.Value {
	once_Data_List_Types_nelCons.Do(func() {
		cache_Data_List_Types_nelCons = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_Types_nelCons(a_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_1_box)))}
})
	})
	return cache_Data_List_Types_nelCons
}

var cache_Data_List_Types_listMap gopurs_runtime.Value
var once_Data_List_Types_listMap sync.Once
func Get_Data_List_Types_listMap() gopurs_runtime.Value {
	once_Data_List_Types_listMap.Do(func() {
		cache_Data_List_Types_listMap = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Types_listMap(f_0_box)
})
	})
	return cache_Data_List_Types_listMap
}

var cache_Data_List_Types_functorList gopurs_runtime.Value
var once_Data_List_Types_functorList sync.Once
func Get_Data_List_Types_functorList() gopurs_runtime.Value {
	once_Data_List_Types_functorList.Do(func() {
		cache_Data_List_Types_functorList = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, Get_Data_List_Types_listMap()})}
	})
	return cache_Data_List_Types_functorList
}

var cache_Data_List_Types_functorNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_functorNonEmptyList sync.Once
func Get_Data_List_Types_functorNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_functorNonEmptyList.Do(func() {
		cache_Data_List_Types_functorNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_NonEmpty_NonEmpty)(m_1.UnsafePtr).V0), gopurs_runtime.Apply(Call_Data_List_Types_listMap(f_0), (*Constructor_Data_NonEmpty_NonEmpty)(m_1.UnsafePtr).V1)})}
})
})})}
	})
	return cache_Data_List_Types_functorNonEmptyList
}

var cache_Data_List_Types_foldableList gopurs_runtime.Value
var once_Data_List_Types_foldableList sync.Once
func Get_Data_List_Types_foldableList() gopurs_runtime.Value {
	once_Data_List_Types_foldableList.Do(func() {
		cache_Data_List_Types_foldableList = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(&Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Types_foldableList()).V1), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Box(Semigroup0_1_0.V0), acc_4)
_ = __local_var_5_2
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, gopurs_runtime.Apply(f_3, x_6))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_3_2 gopurs_runtime.Value
go__go_1_3_2 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_3_2:
for {
if false { continue go__go_1_3_2 }
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
b_2_loop = gopurs_runtime.Apply2(f_0, b_2, (*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V0)
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V1)}
continue go__go_1_3_2
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
return go__go_1_3_2
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_5 -> gopurs_runtime.Value
__local_var_2_5 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Types_foldableList()).V1), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
})
}), b_1)
_ = __local_var_2_5
var go__go_3_7_3 gopurs_runtime.Value
go__go_3_7_3 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_7_3:
for {
if false { continue go__go_3_7_3 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t8 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t8 = v_4
goto end_branch_8
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_7_3
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t8)}
}
}()
})
})
// TAST (Let): __local_var_3_6 -> gopurs_runtime.Value
__local_var_3_6 := gopurs_runtime.Apply(go__go_3_7_3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
_ = __local_var_3_6
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_5, gopurs_runtime.Apply(__local_var_3_6, x_4))
})
})
})})}
	})
	return cache_Data_List_Types_foldableList
}

var cache_Data_List_Types_foldableNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_foldableNonEmptyList sync.Once
func Get_Data_List_Types_foldableNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_foldableNonEmptyList.Do(func() {
		cache_Data_List_Types_foldableNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(&Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_4_1 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_4_1
var go__go_5_2_4 gopurs_runtime.Value
go__go_5_2_4 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_2_4:
for {
if false { continue go__go_5_2_4 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t3 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr == nil) {
__t3 = b_6
goto end_branch_3
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr != nil) {
b_6_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_4_1.V0), b_6, gopurs_runtime.Apply(f_2, (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0))
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_2_4
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
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0), gopurs_runtime.Apply2(go__go_5_2_4, gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_4_5 gopurs_runtime.Value
go__go_3_4_5 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_4_5:
for {
if false { continue go__go_3_4_5 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t5 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t5 = b_4
goto end_branch_5
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Apply2(f_0, b_4, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0)
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_4_5
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}()
})
})
return gopurs_runtime.Apply2(go__go_3_4_5, gopurs_runtime.Apply2(f_0, b_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_6_6 gopurs_runtime.Value
go__go_3_6_6 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_6_6:
for {
if false { continue go__go_3_6_6 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t7 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t7 = b_4
goto end_branch_7
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0, b_4)
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_6_6
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
var go__go_4_8_7 gopurs_runtime.Value
go__go_4_8_7 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_8_7:
for {
if false { continue go__go_4_8_7 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t9 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t9 = v_5
goto end_branch_9
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_8_7
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t9)}
}
}()
})
})
return gopurs_runtime.Apply2(f_0, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_3_6_6, b_1, gopurs_runtime.Apply2(go__go_4_8_7, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)))
})
})
})})}
	})
	return cache_Data_List_Types_foldableNonEmptyList
}

var cache_Data_List_Types_foldableWithIndexList gopurs_runtime.Value
var once_Data_List_Types_foldableWithIndexList sync.Once
func Get_Data_List_Types_foldableWithIndexList() gopurs_runtime.Value {
	once_Data_List_Types_foldableWithIndexList.Do(func() {
		cache_Data_List_Types_foldableWithIndexList = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(&Constructor_Data_FoldableWithIndex_FoldableWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Types_foldableList()))}
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex](Get_Data_List_Types_foldableWithIndexList()).V2), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 -> gopurs_runtime.Value
__local_var_6_2 := gopurs_runtime.Apply(gopurs_runtime.Box(Semigroup0_1_0.V0), acc_5)
_ = __local_var_6_2
// TAST (Let): __local_var_7_3 -> gopurs_runtime.Value
__local_var_7_3 := gopurs_runtime.Apply(f_3, gopurs_runtime.Int(i_4.IntVal))
_ = __local_var_7_3
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_2, gopurs_runtime.Apply(__local_var_7_3, x_8))
})
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_5_8 gopurs_runtime.Value
go__go_2_5_8 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_5_8:
for {
if false { continue go__go_2_5_8 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t6 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t6 = b_3
goto end_branch_6
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_3.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple)(b_3.UnsafePtr).V0.IntVal), (*Constructor_Data_Tuple_Tuple)(b_3.UnsafePtr).V1, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0)})}
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_5_8
__t6 = gopurs_runtime.Value{}
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
}()
})
})
// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
__local_var_2_4 := gopurs_runtime.Apply(go__go_2_5_8, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), acc_1})})
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply(__local_var_2_4, x_3).UnsafePtr).V1
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_8_9 gopurs_runtime.Value
go__go_3_8_9 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_8_9:
for {
if false { continue go__go_3_8_9 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t9 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t9 = b_4
goto end_branch_9
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V1)})}})}
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_8_9
__t9 = gopurs_runtime.Value{}
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
}
}()
})
})
// TAST (Let): v_3_7 -> *Constructor_Data_Tuple_Tuple
v_3_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(go__go_3_8_9, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_2))}))
_ = v_3_7
var go__go_4_10_10 gopurs_runtime.Value
go__go_4_10_10 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_10_10:
for {
if false { continue go__go_4_10_10 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t11 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t11 = b_5
goto end_branch_11
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) - (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) - (1)), (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V1)})}
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_10_10
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
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_4_10_10, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int((v_3_7).V0.IntVal), b_1})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_3_7).V1))}).UnsafePtr).V1
})
})
})})}
	})
	return cache_Data_List_Types_foldableWithIndexList
}

var cache_Data_List_Types_foldableWithIndexNonEmpty gopurs_runtime.Value
var once_Data_List_Types_foldableWithIndexNonEmpty sync.Once
func Get_Data_List_Types_foldableWithIndexNonEmpty() gopurs_runtime.Value {
	once_Data_List_Types_foldableWithIndexNonEmpty.Do(func() {
		cache_Data_List_Types_foldableWithIndexNonEmpty = func() gopurs_runtime.Value {
// TAST (Let): foldableNonEmpty1_0_0 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_0_0 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_1 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_4_2 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_4_2
var go__go_5_3_11 gopurs_runtime.Value
go__go_5_3_11 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_3_11:
for {
if false { continue go__go_5_3_11 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t4 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr == nil) {
__t4 = b_6
goto end_branch_4
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr != nil) {
b_6_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_4_2.V0), b_6, gopurs_runtime.Apply(f_2, (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0))
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_3_11
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
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_1.V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0), gopurs_runtime.Apply2(go__go_5_3_11, gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_5_12 gopurs_runtime.Value
go__go_3_5_12 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_5_12:
for {
if false { continue go__go_3_5_12 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t6 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t6 = b_4
goto end_branch_6
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Apply2(f_0, b_4, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0)
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_5_12
__t6 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Apply2(go__go_3_5_12, gopurs_runtime.Apply2(f_0, b_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_7_13 gopurs_runtime.Value
go__go_3_7_13 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_7_13:
for {
if false { continue go__go_3_7_13 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t8 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t8 = b_4
goto end_branch_8
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0, b_4)
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_7_13
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
var go__go_4_9_14 gopurs_runtime.Value
go__go_4_9_14 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_9_14:
for {
if false { continue go__go_4_9_14 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t10 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t10 = v_5
goto end_branch_10
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_9_14
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t10)}
}
}()
})
})
return gopurs_runtime.Apply2(f_0, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_3_7_13, b_1, gopurs_runtime.Apply2(go__go_4_9_14, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)))
})
})
})}
_ = foldableNonEmpty1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(&Constructor_Data_FoldableWithIndex_FoldableWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_0_0)}
}), gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_2_11 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_11
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_5_12 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_5_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_5_12
var go__go_6_13_15 gopurs_runtime.Value
go__go_6_13_15 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_7_loop gopurs_runtime.Value = b_7_loop_val
var v_8_loop gopurs_runtime.Value = v_8_loop_val
go__go_6_13_15:
for {
if false { continue go__go_6_13_15 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var v_8 gopurs_runtime.Value = v_8_loop
_ = v_8
var __t14 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 1358893437 && v_8.UnsafePtr == nil) {
__t14 = b_7
goto end_branch_14
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 1358893437 && v_8.UnsafePtr != nil) {
b_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_7.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_12.V0), (*Constructor_Data_Tuple_Tuple)(b_7.UnsafePtr).V1, gopurs_runtime.Apply2(f_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple)(b_7.UnsafePtr).V0.IntVal)})}, (*Constructor_Data_List_Types_Cons)(v_8.UnsafePtr).V0))})}
v_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_8.UnsafePtr).V1)}
continue go__go_6_13_15
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
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_11.V0), gopurs_runtime.Apply2(f_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_6_13_15, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.RecordGet(dictMonoid_1, "mempty")})}, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1).UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_15_16 gopurs_runtime.Value
go__go_4_15_16 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_15_16:
for {
if false { continue go__go_4_15_16 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t16 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t16 = b_5
goto end_branch_16
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply3(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal)})}, (*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V1, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0)})}
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_15_16
__t16 = gopurs_runtime.Value{}
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
return __t16
}
}()
})
})
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_4_15_16, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.Apply3(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, b_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0)})}, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1).UnsafePtr).V1
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_18_17 gopurs_runtime.Value
go__go_4_18_17 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_18_17:
for {
if false { continue go__go_4_18_17 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t19 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t19 = b_5
goto end_branch_19
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V1)})}})}
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_18_17
__t19 = gopurs_runtime.Value{}
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
return __t19
}
}()
})
})
// TAST (Let): v_4_17 -> *Constructor_Data_Tuple_Tuple
v_4_17 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(go__go_4_18_17, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1))}))
_ = v_4_17
var go__go_5_20_18 gopurs_runtime.Value
go__go_5_20_18 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_20_18:
for {
if false { continue go__go_5_20_18 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t21 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr == nil) {
__t21 = b_6
goto end_branch_21
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr != nil) {
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V0.IntVal) - (1)), gopurs_runtime.Apply3(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V0.IntVal) - (1))})}, (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V1)})}
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_20_18
__t21 = gopurs_runtime.Value{}
goto end_branch_21
} else {

}
}
{
__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_21:
return __t21
}
}()
})
})
return gopurs_runtime.Apply3(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_5_20_18, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int((v_4_17).V0.IntVal), b_2})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_4_17).V1))}).UnsafePtr).V1)
})
})
})})}
}()
	})
	return cache_Data_List_Types_foldableWithIndexNonEmpty
}

var cache_Data_List_Types_foldableWithIndexNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_foldableWithIndexNonEmptyList sync.Once
func Get_Data_List_Types_foldableWithIndexNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_foldableWithIndexNonEmptyList.Do(func() {
		cache_Data_List_Types_foldableWithIndexNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(&Constructor_Data_FoldableWithIndex_FoldableWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Types_foldableNonEmptyList()))}
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_3_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_3_0
var go__go_4_1_19 gopurs_runtime.Value
go__go_4_1_19 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_1_19:
for {
if false { continue go__go_4_1_19 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t2 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t2 = b_5
goto end_branch_2
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_0.V0), (*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V1, gopurs_runtime.Apply2(f_1, gopurs_runtime.Int((1) + ((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal)), (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0))})}
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_1_19
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply2(f_1, gopurs_runtime.Int(0), (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_4_1_19, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})}, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1).UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_3_20 gopurs_runtime.Value
go__go_3_3_20 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_3_20:
for {
if false { continue go__go_3_3_20 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t4 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t4 = b_4
goto end_branch_4
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((1) + ((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V0.IntVal)), (*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V1, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0)})}
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_3_20
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
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_3_3_20, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(0), b_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0)})}, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1).UnsafePtr).V1
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_6_21 gopurs_runtime.Value
go__go_3_6_21 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_6_21:
for {
if false { continue go__go_3_6_21 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t7 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t7 = b_4
goto end_branch_7
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V1)})}})}
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_6_21
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
// TAST (Let): v_3_5 -> *Constructor_Data_Tuple_Tuple
v_3_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(go__go_3_6_21, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1))}))
_ = v_3_5
var go__go_4_8_22 gopurs_runtime.Value
go__go_4_8_22 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_8_22:
for {
if false { continue go__go_4_8_22 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t9 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t9 = b_5
goto end_branch_9
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) - (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((1) + (((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) - (1))), (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V1)})}
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_8_22
__t9 = gopurs_runtime.Value{}
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
}
}()
})
})
return gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(0), (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_4_8_22, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int((v_3_5).V0.IntVal), b_1})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_3_5).V1))}).UnsafePtr).V1)
})
})
})})}
	})
	return cache_Data_List_Types_foldableWithIndexNonEmptyList
}

var cache_Data_List_Types_functorWithIndexList gopurs_runtime.Value
var once_Data_List_Types_functorWithIndexList sync.Once
func Get_Data_List_Types_functorWithIndexList() gopurs_runtime.Value {
	once_Data_List_Types_functorWithIndexList.Do(func() {
		cache_Data_List_Types_functorWithIndexList = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(&Constructor_Data_FunctorWithIndex_FunctorWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Types_functorList()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_1_23 gopurs_runtime.Value
go__go_2_1_23 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_1_23:
for {
if false { continue go__go_2_1_23 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t2 = b_3
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_3.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_Tuple_Tuple)(b_3.UnsafePtr).V1)})}})}
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_1_23
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}()
})
})
// TAST (Let): v_2_0 -> *Constructor_Data_Tuple_Tuple
v_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(go__go_2_1_23, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_1))}))
_ = v_2_0
var go__go_3_3_24 gopurs_runtime.Value
go__go_3_3_24 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_3_24:
for {
if false { continue go__go_3_3_24 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t4 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t4 = b_4
goto end_branch_4
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V0.IntVal) - (1)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Apply2(f_0, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V0.IntVal) - (1)), (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V1)})}})}
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_3_24
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
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_3_3_24, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int((v_2_0).V0.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_2_0).V1))}).UnsafePtr).V1
})
})})}
	})
	return cache_Data_List_Types_functorWithIndexList
}

var cache_Data_List_Types_functorWithIndex gopurs_runtime.Value
var once_Data_List_Types_functorWithIndex sync.Once
func Get_Data_List_Types_functorWithIndex() gopurs_runtime.Value {
	once_Data_List_Types_functorWithIndex.Do(func() {
		cache_Data_List_Types_functorWithIndex = func() gopurs_runtime.Value {
// TAST (Let): functorNonEmpty1_0_0 -> *Constructor_Data_Functor_Functor
functorNonEmpty1_0_0 := &Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_NonEmpty_NonEmpty)(m_1.UnsafePtr).V0), gopurs_runtime.Apply(Call_Data_List_Types_listMap(f_0), (*Constructor_Data_NonEmpty_NonEmpty)(m_1.UnsafePtr).V1)})}
})
})}
_ = functorNonEmpty1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(&Constructor_Data_FunctorWithIndex_FunctorWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorNonEmpty1_0_0)}
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_2_25 gopurs_runtime.Value
go__go_3_2_25 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_2_25:
for {
if false { continue go__go_3_2_25 }
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
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V1)})}})}
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_2_25
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
// TAST (Let): v_3_1 -> *Constructor_Data_Tuple_Tuple
v_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(go__go_3_2_25, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1))}))
_ = v_3_1
var go__go_4_4_26 gopurs_runtime.Value
go__go_4_4_26 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_4_26:
for {
if false { continue go__go_4_4_26 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t5 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t5 = b_5
goto end_branch_5
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) - (1)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Apply2(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) - (1))})}, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V1)})}})}
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_4_26
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply2(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_4_4_26, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int((v_3_1).V0.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_3_1).V1))}).UnsafePtr).V1})}
})
})})}
}()
	})
	return cache_Data_List_Types_functorWithIndex
}

var cache_Data_List_Types_functorWithIndexNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_functorWithIndexNonEmptyList sync.Once
func Get_Data_List_Types_functorWithIndexNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_functorWithIndexNonEmptyList.Do(func() {
		cache_Data_List_Types_functorWithIndexNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(&Constructor_Data_FunctorWithIndex_FunctorWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Types_functorNonEmptyList()))}
}), gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_1_27 gopurs_runtime.Value
go__go_2_1_27 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_1_27:
for {
if false { continue go__go_2_1_27 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t2 = b_3
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_3.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_Tuple_Tuple)(b_3.UnsafePtr).V1)})}})}
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_1_27
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}()
})
})
// TAST (Let): v_2_0 -> *Constructor_Data_Tuple_Tuple
v_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(go__go_2_1_27, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_1.UnsafePtr).V1))}))
_ = v_2_0
var go__go_3_3_28 gopurs_runtime.Value
go__go_3_3_28 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_3_28:
for {
if false { continue go__go_3_3_28 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t4 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t4 = b_4
goto end_branch_4
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V0.IntVal) - (1)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Apply2(fn_0, gopurs_runtime.Int((1) + (((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V0.IntVal) - (1))), (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V1)})}})}
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_3_28
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
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply2(fn_0, gopurs_runtime.Int(0), (*Constructor_Data_NonEmpty_NonEmpty)(v_1.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_3_3_28, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int((v_2_0).V0.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_2_0).V1))}).UnsafePtr).V1})}
})
})})}
	})
	return cache_Data_List_Types_functorWithIndexNonEmptyList
}

var cache_Data_List_Types_semigroupList gopurs_runtime.Value
var once_Data_List_Types_semigroupList sync.Once
func Get_Data_List_Types_semigroupList() gopurs_runtime.Value {
	once_Data_List_Types_semigroupList.Do(func() {
		cache_Data_List_Types_semigroupList = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_29 gopurs_runtime.Value
go__go_2_0_29 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_0_29:
for {
if false { continue go__go_2_0_29 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t1 = b_3
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_3)})}
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_0_29
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
})
var go__go_3_2_30 gopurs_runtime.Value
go__go_3_2_30 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_2_30:
for {
if false { continue go__go_3_2_30 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t3 = v_4
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_2_30
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_2_0_29, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_1))}, gopurs_runtime.Apply2(go__go_3_2_30, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_0))}))))}
})
})})}
	})
	return cache_Data_List_Types_semigroupList
}

var cache_Data_List_Types_monoidList gopurs_runtime.Value
var once_Data_List_Types_monoidList sync.Once
func Get_Data_List_Types_monoidList() gopurs_runtime.Value {
	once_Data_List_Types_monoidList.Do(func() {
		cache_Data_List_Types_monoidList = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Types_semigroupList()))}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})}
	})
	return cache_Data_List_Types_monoidList
}

var cache_Data_List_Types_semigroupNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_semigroupNonEmptyList sync.Once
func Get_Data_List_Types_semigroupNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_semigroupNonEmptyList.Do(func() {
		cache_Data_List_Types_semigroupNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(as_prime_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_31 gopurs_runtime.Value
go__go_2_0_31 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_0_31:
for {
if false { continue go__go_2_0_31 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t1 = b_3
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_3)})}
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_0_31
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
})
var go__go_3_2_32 gopurs_runtime.Value
go__go_3_2_32 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_2_32:
for {
if false { continue go__go_3_2_32 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t3 = v_4
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_2_32
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_2_0_31, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(as_prime_1.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(as_prime_1.UnsafePtr).V1)})}, gopurs_runtime.Apply2(go__go_3_2_32, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1))}))))}})}
})
})})}
	})
	return cache_Data_List_Types_semigroupNonEmptyList
}

var cache_Data_List_Types_showList gopurs_runtime.Value
var once_Data_List_Types_showList sync.Once
func Get_Data_List_Types_showList() gopurs_runtime.Value {
	once_Data_List_Types_showList.Do(func() {
		cache_Data_List_Types_showList = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Types_showList(dictShow_0_box)
})
	})
	return cache_Data_List_Types_showList
}

var cache_Data_List_Types_showNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_showNonEmptyList sync.Once
func Get_Data_List_Types_showNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_showNonEmptyList.Do(func() {
		cache_Data_List_Types_showNonEmptyList = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Types_showNonEmptyList(dictShow_0_box)
})
	})
	return cache_Data_List_Types_showNonEmptyList
}

var cache_Data_List_Types_traversableList gopurs_runtime.Value
var once_Data_List_Types_traversableList sync.Once
func Get_Data_List_Types_traversableList() gopurs_runtime.Value {
	once_Data_List_Types_traversableList.Do(func() {
		cache_Data_List_Types_traversableList = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(&Constructor_Data_Traversable_Traversable{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Types_foldableList()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Types_functorList()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](Get_Data_List_Types_traversableList()).V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0))}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): Apply0_2_1 -> *Constructor_Control_Apply_Apply
Apply0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_3_33 gopurs_runtime.Value
go__go_4_3_33 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_3_33:
for {
if false { continue go__go_4_3_33 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t4 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t4 = b_5
goto end_branch_4
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_5)})}
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_3_33
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
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Apply(go__go_4_3_33, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}))
_ = __local_var_4_2
var go__go_5_6_34 gopurs_runtime.Value
go__go_5_6_34 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_6_34:
for {
if false { continue go__go_5_6_34 }
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
b_6_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_2_1.V0), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, a_9, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_8)})}
})
}), b_6), gopurs_runtime.Apply(f_3, (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0))
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_6_34
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
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(go__go_5_6_34, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}))
_ = __local_var_5_5
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_5_5, x_6))
})
})
})})}
	})
	return cache_Data_List_Types_traversableList
}

var cache_Data_List_Types_traversableNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_traversableNonEmptyList sync.Once
func Get_Data_List_Types_traversableNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_traversableNonEmptyList.Do(func() {
		cache_Data_List_Types_traversableNonEmptyList = func() gopurs_runtime.Value {
// TAST (Let): functorNonEmpty1_0_0 -> *Constructor_Data_Functor_Functor
functorNonEmpty1_0_0 := &Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_NonEmpty_NonEmpty)(m_1.UnsafePtr).V0), gopurs_runtime.Apply(Call_Data_List_Types_listMap(f_0), (*Constructor_Data_NonEmpty_NonEmpty)(m_1.UnsafePtr).V1)})}
})
})}
_ = functorNonEmpty1_0_0
// TAST (Let): foldableNonEmpty1_1_1 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_1_1 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_2_2 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_2
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_5_3 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_5_3
var go__go_6_4_35 gopurs_runtime.Value
go__go_6_4_35 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_7_loop gopurs_runtime.Value = b_7_loop_val
var v_8_loop gopurs_runtime.Value = v_8_loop_val
go__go_6_4_35:
for {
if false { continue go__go_6_4_35 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var v_8 gopurs_runtime.Value = v_8_loop
_ = v_8
var __t5 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 1358893437 && v_8.UnsafePtr == nil) {
__t5 = b_7
goto end_branch_5
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 1358893437 && v_8.UnsafePtr != nil) {
b_7_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_3.V0), b_7, gopurs_runtime.Apply(f_3, (*Constructor_Data_List_Types_Cons)(v_8.UnsafePtr).V0))
v_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_8.UnsafePtr).V1)}
continue go__go_6_4_35
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_2.V0), gopurs_runtime.Apply(f_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0), gopurs_runtime.Apply2(go__go_6_4_35, gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_6_36 gopurs_runtime.Value
go__go_4_6_36 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_6_36:
for {
if false { continue go__go_4_6_36 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t7 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t7 = b_5
goto end_branch_7
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Apply2(f_1, b_5, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0)
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_6_36
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
return gopurs_runtime.Apply2(go__go_4_6_36, gopurs_runtime.Apply2(f_1, b_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_8_37 gopurs_runtime.Value
go__go_4_8_37 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_8_37:
for {
if false { continue go__go_4_8_37 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t9 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t9 = b_5
goto end_branch_9
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Apply2(f_1, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0, b_5)
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_8_37
__t9 = gopurs_runtime.Value{}
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
}
}()
})
})
var go__go_5_10_38 gopurs_runtime.Value
go__go_5_10_38 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_6_loop_val)
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_10_38:
for {
if false { continue go__go_5_10_38 }
var v_6 *Constructor_Data_List_Types_Cons = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t11 *Constructor_Data_List_Types_Cons
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr == nil) {
__t11 = v_6
goto end_branch_11
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V0, v_6})})
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V1)}
continue go__go_5_10_38
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t11)}
}
}()
})
})
return gopurs_runtime.Apply2(f_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_4_8_37, b_2, gopurs_runtime.Apply2(go__go_5_10_38, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1)))
})
})
})}
_ = foldableNonEmpty1_1_1
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(&Constructor_Data_Traversable_Traversable{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_1_1)}
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorNonEmpty1_0_0)}
}), gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_3_12 -> *Constructor_Control_Apply_Apply
Apply0_3_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_3_12
// TAST (Let): Functor0_4_13 -> *Constructor_Data_Functor_Functor
Functor0_4_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_13
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_6_14 -> *Constructor_Control_Apply_Apply
Apply0_6_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_6_14
var go__go_7_15_39 gopurs_runtime.Value
go__go_7_15_39 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_8_loop gopurs_runtime.Value = b_8_loop_val
var v_9_loop gopurs_runtime.Value = v_9_loop_val
go__go_7_15_39:
for {
if false { continue go__go_7_15_39 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var v_9 gopurs_runtime.Value = v_9_loop
_ = v_9
var __t16 gopurs_runtime.Value
{
if (v_9.Type == 9 && v_9.IntVal == 1358893437 && v_9.UnsafePtr == nil) {
__t16 = b_8
goto end_branch_16
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 1358893437 && v_9.UnsafePtr != nil) {
b_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_9.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_8)})}
v_9_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_9.UnsafePtr).V1)}
continue go__go_7_15_39
__t16 = gopurs_runtime.Value{}
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
return __t16
}
}()
})
})
var go__go_7_17_40 gopurs_runtime.Value
go__go_7_17_40 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_8_loop gopurs_runtime.Value = b_8_loop_val
var v_9_loop gopurs_runtime.Value = v_9_loop_val
go__go_7_17_40:
for {
if false { continue go__go_7_17_40 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var v_9 gopurs_runtime.Value = v_9_loop
_ = v_9
var __t18 gopurs_runtime.Value
{
if (v_9.Type == 9 && v_9.IntVal == 1358893437 && v_9.UnsafePtr == nil) {
__t18 = b_8
goto end_branch_18
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 1358893437 && v_9.UnsafePtr != nil) {
b_8_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_6_14.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_6_14.V0), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(b_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, a_11, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_10)})}
})
}), b_8), (*Constructor_Data_List_Types_Cons)(v_9.UnsafePtr).V0)
v_9_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_9.UnsafePtr).V1)}
continue go__go_7_17_40
__t18 = gopurs_runtime.Value{}
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
return __t18
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_12.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_13.V0), Get_Data_NonEmpty_NonEmpty(), (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(go__go_7_15_39, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), gopurs_runtime.Apply2(go__go_7_17_40, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1)))
})
}), gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_3_19 -> *Constructor_Control_Apply_Apply
Apply0_3_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_3_19
// TAST (Let): Functor0_4_20 -> *Constructor_Data_Functor_Functor
Functor0_4_20 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_20
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_7_21 -> *Constructor_Control_Apply_Apply
Apply0_7_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_7_21
var go__go_8_22_41 gopurs_runtime.Value
go__go_8_22_41 = gopurs_runtime.Func(func(b_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_9_loop gopurs_runtime.Value = b_9_loop_val
var v_10_loop gopurs_runtime.Value = v_10_loop_val
go__go_8_22_41:
for {
if false { continue go__go_8_22_41 }
var b_9 gopurs_runtime.Value = b_9_loop
_ = b_9
var v_10 gopurs_runtime.Value = v_10_loop
_ = v_10
var __t23 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 1358893437 && v_10.UnsafePtr == nil) {
__t23 = b_9
goto end_branch_23
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 1358893437 && v_10.UnsafePtr != nil) {
b_9_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_10.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_9)})}
v_10_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_10.UnsafePtr).V1)}
continue go__go_8_22_41
__t23 = gopurs_runtime.Value{}
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
return __t23
}
}()
})
})
var go__go_8_24_42 gopurs_runtime.Value
go__go_8_24_42 = gopurs_runtime.Func(func(b_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_9_loop gopurs_runtime.Value = b_9_loop_val
var v_10_loop gopurs_runtime.Value = v_10_loop_val
go__go_8_24_42:
for {
if false { continue go__go_8_24_42 }
var b_9 gopurs_runtime.Value = b_9_loop
_ = b_9
var v_10 gopurs_runtime.Value = v_10_loop
_ = v_10
var __t25 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 1358893437 && v_10.UnsafePtr == nil) {
__t25 = b_9
goto end_branch_25
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 1358893437 && v_10.UnsafePtr != nil) {
b_9_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_7_21.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_7_21.V0), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(b_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, a_12, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_11)})}
})
}), b_9), gopurs_runtime.Apply(f_5, (*Constructor_Data_List_Types_Cons)(v_10.UnsafePtr).V0))
v_10_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_10.UnsafePtr).V1)}
continue go__go_8_24_42
__t25 = gopurs_runtime.Value{}
goto end_branch_25
} else {

}
}
{
__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_25:
return __t25
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_19.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_20.V0), Get_Data_NonEmpty_NonEmpty(), gopurs_runtime.Apply(f_5, (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(go__go_8_22_41, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), gopurs_runtime.Apply2(go__go_8_24_42, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V1)))
})
})
})})}
}()
	})
	return cache_Data_List_Types_traversableNonEmptyList
}

var cache_Data_List_Types_traversableWithIndexList gopurs_runtime.Value
var once_Data_List_Types_traversableWithIndexList sync.Once
func Get_Data_List_Types_traversableWithIndexList() gopurs_runtime.Value {
	once_Data_List_Types_traversableWithIndexList.Do(func() {
		cache_Data_List_Types_traversableWithIndexList = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(&Constructor_Data_TraversableWithIndex_TraversableWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex](Get_Data_List_Types_foldableWithIndexList()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex](Get_Data_List_Types_functorWithIndexList()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](Get_Data_List_Types_traversableList()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): Apply0_2_1 -> *Constructor_Control_Apply_Apply
Apply0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_3_43 gopurs_runtime.Value
go__go_4_3_43 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_3_43:
for {
if false { continue go__go_4_3_43 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t4 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t4 = b_5
goto end_branch_4
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_5)})}
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_3_43
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
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(go__go_4_3_43, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))})
_ = __local_var_4_2
var go__go_5_7_44 gopurs_runtime.Value
go__go_5_7_44 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_7_44:
for {
if false { continue go__go_5_7_44 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t8 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr == nil) {
__t8 = b_6
goto end_branch_8
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr != nil) {
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_2_1.V0), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, a_9, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_8)})}
})
}), (*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V1), gopurs_runtime.Apply2(f_3, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V0.IntVal), (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0))})}
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_7_44
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
// TAST (Let): __local_var_5_6 -> gopurs_runtime.Value
__local_var_5_6 := gopurs_runtime.Apply(go__go_5_7_44, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})})})
_ = __local_var_5_6
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply(__local_var_5_6, x_6).UnsafePtr).V1
})
_ = __local_var_5_5
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_5_5, x_6))
})
})
})})}
	})
	return cache_Data_List_Types_traversableWithIndexList
}

var cache_Data_List_Types_traversableWithIndexNonEmpty gopurs_runtime.Value
var once_Data_List_Types_traversableWithIndexNonEmpty sync.Once
func Get_Data_List_Types_traversableWithIndexNonEmpty() gopurs_runtime.Value {
	once_Data_List_Types_traversableWithIndexNonEmpty.Do(func() {
		cache_Data_List_Types_traversableWithIndexNonEmpty = func() gopurs_runtime.Value {
// TAST (Let): functorNonEmpty1_0_1 -> *Constructor_Data_Functor_Functor
functorNonEmpty1_0_1 := &Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_NonEmpty_NonEmpty)(m_1.UnsafePtr).V0), gopurs_runtime.Apply(Call_Data_List_Types_listMap(f_0), (*Constructor_Data_NonEmpty_NonEmpty)(m_1.UnsafePtr).V1)})}
})
})}
_ = functorNonEmpty1_0_1
// TAST (Let): functorWithIndex1_0_0 -> *Constructor_Data_FunctorWithIndex_FunctorWithIndex
functorWithIndex1_0_0 := &Constructor_Data_FunctorWithIndex_FunctorWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorNonEmpty1_0_1)}
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_3_45 gopurs_runtime.Value
go__go_3_3_45 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_3_45:
for {
if false { continue go__go_3_3_45 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t4 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t4 = b_4
goto end_branch_4
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V1)})}})}
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_3_45
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
// TAST (Let): v_3_2 -> *Constructor_Data_Tuple_Tuple
v_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(go__go_3_3_45, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1))}))
_ = v_3_2
var go__go_4_5_46 gopurs_runtime.Value
go__go_4_5_46 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_5_46:
for {
if false { continue go__go_4_5_46 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t6 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t6 = b_5
goto end_branch_6
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) - (1)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Apply2(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) - (1))})}, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V1)})}})}
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_5_46
__t6 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply2(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_4_5_46, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int((v_3_2).V0.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_3_2).V1))}).UnsafePtr).V1})}
})
})}
_ = functorWithIndex1_0_0
// TAST (Let): foldableNonEmpty1_1_8 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_1_8 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_2_9 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_9
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_5_10 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_5_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_5_10
var go__go_6_11_47 gopurs_runtime.Value
go__go_6_11_47 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_7_loop gopurs_runtime.Value = b_7_loop_val
var v_8_loop gopurs_runtime.Value = v_8_loop_val
go__go_6_11_47:
for {
if false { continue go__go_6_11_47 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var v_8 gopurs_runtime.Value = v_8_loop
_ = v_8
var __t12 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 1358893437 && v_8.UnsafePtr == nil) {
__t12 = b_7
goto end_branch_12
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 1358893437 && v_8.UnsafePtr != nil) {
b_7_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_10.V0), b_7, gopurs_runtime.Apply(f_3, (*Constructor_Data_List_Types_Cons)(v_8.UnsafePtr).V0))
v_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_8.UnsafePtr).V1)}
continue go__go_6_11_47
__t12 = gopurs_runtime.Value{}
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_9.V0), gopurs_runtime.Apply(f_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0), gopurs_runtime.Apply2(go__go_6_11_47, gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_13_48 gopurs_runtime.Value
go__go_4_13_48 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_13_48:
for {
if false { continue go__go_4_13_48 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t14 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t14 = b_5
goto end_branch_14
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Apply2(f_1, b_5, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0)
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_13_48
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
return gopurs_runtime.Apply2(go__go_4_13_48, gopurs_runtime.Apply2(f_1, b_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_15_49 gopurs_runtime.Value
go__go_4_15_49 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_15_49:
for {
if false { continue go__go_4_15_49 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t16 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t16 = b_5
goto end_branch_16
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Apply2(f_1, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0, b_5)
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_15_49
__t16 = gopurs_runtime.Value{}
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
return __t16
}
}()
})
})
var go__go_5_17_50 gopurs_runtime.Value
go__go_5_17_50 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_6_loop_val)
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_17_50:
for {
if false { continue go__go_5_17_50 }
var v_6 *Constructor_Data_List_Types_Cons = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t18 *Constructor_Data_List_Types_Cons
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr == nil) {
__t18 = v_6
goto end_branch_18
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V0, v_6})})
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V1)}
continue go__go_5_17_50
__t18 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_18
} else {

}
}
{
__t18 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_18:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t18)}
}
}()
})
})
return gopurs_runtime.Apply2(f_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_4_15_49, b_2, gopurs_runtime.Apply2(go__go_5_17_50, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1)))
})
})
})}
_ = foldableNonEmpty1_1_8
// TAST (Let): foldableWithIndexNonEmpty1_1_7 -> *Constructor_Data_FoldableWithIndex_FoldableWithIndex
foldableWithIndexNonEmpty1_1_7 := &Constructor_Data_FoldableWithIndex_FoldableWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_1_8)}
}), gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_3_19 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_3_19 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_3_19
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_6_20 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_6_20 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_6_20
var go__go_7_21_51 gopurs_runtime.Value
go__go_7_21_51 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_8_loop gopurs_runtime.Value = b_8_loop_val
var v_9_loop gopurs_runtime.Value = v_9_loop_val
go__go_7_21_51:
for {
if false { continue go__go_7_21_51 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var v_9 gopurs_runtime.Value = v_9_loop
_ = v_9
var __t22 gopurs_runtime.Value
{
if (v_9.Type == 9 && v_9.IntVal == 1358893437 && v_9.UnsafePtr == nil) {
__t22 = b_8
goto end_branch_22
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 1358893437 && v_9.UnsafePtr != nil) {
b_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_8.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_20.V0), (*Constructor_Data_Tuple_Tuple)(b_8.UnsafePtr).V1, gopurs_runtime.Apply2(f_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple)(b_8.UnsafePtr).V0.IntVal)})}, (*Constructor_Data_List_Types_Cons)(v_9.UnsafePtr).V0))})}
v_9_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_9.UnsafePtr).V1)}
continue go__go_7_21_51
__t22 = gopurs_runtime.Value{}
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_22:
return __t22
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_19.V0), gopurs_runtime.Apply2(f_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_7_21_51, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.RecordGet(dictMonoid_2, "mempty")})}, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1).UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_23_52 gopurs_runtime.Value
go__go_5_23_52 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_23_52:
for {
if false { continue go__go_5_23_52 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t24 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr == nil) {
__t24 = b_6
goto end_branch_24
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr != nil) {
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply3(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V0.IntVal)})}, (*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V1, (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0)})}
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_23_52
__t24 = gopurs_runtime.Value{}
goto end_branch_24
} else {

}
}
{
__t24 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_24:
return __t24
}
}()
})
})
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_5_23_52, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.Apply3(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, b_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0)})}, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1).UnsafePtr).V1
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_26_53 gopurs_runtime.Value
go__go_5_26_53 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_26_53:
for {
if false { continue go__go_5_26_53 }
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
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V1)})}})}
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_26_53
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
// TAST (Let): v_5_25 -> *Constructor_Data_Tuple_Tuple
v_5_25 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(go__go_5_26_53, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1))}))
_ = v_5_25
var go__go_6_28_54 gopurs_runtime.Value
go__go_6_28_54 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_7_loop gopurs_runtime.Value = b_7_loop_val
var v_8_loop gopurs_runtime.Value = v_8_loop_val
go__go_6_28_54:
for {
if false { continue go__go_6_28_54 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var v_8 gopurs_runtime.Value = v_8_loop
_ = v_8
var __t29 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 1358893437 && v_8.UnsafePtr == nil) {
__t29 = b_7
goto end_branch_29
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 1358893437 && v_8.UnsafePtr != nil) {
b_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_7.UnsafePtr).V0.IntVal) - (1)), gopurs_runtime.Apply3(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_7.UnsafePtr).V0.IntVal) - (1))})}, (*Constructor_Data_List_Types_Cons)(v_8.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(b_7.UnsafePtr).V1)})}
v_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_8.UnsafePtr).V1)}
continue go__go_6_28_54
__t29 = gopurs_runtime.Value{}
goto end_branch_29
} else {

}
}
{
__t29 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_29:
return __t29
}
}()
})
})
return gopurs_runtime.Apply3(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_6_28_54, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int((v_5_25).V0.IntVal), b_3})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_5_25).V1))}).UnsafePtr).V1)
})
})
})}
_ = foldableWithIndexNonEmpty1_1_7
// TAST (Let): functorNonEmpty1_2_31 -> *Constructor_Data_Functor_Functor
functorNonEmpty1_2_31 := &Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(m_3.UnsafePtr).V0), gopurs_runtime.Apply(Call_Data_List_Types_listMap(f_2), (*Constructor_Data_NonEmpty_NonEmpty)(m_3.UnsafePtr).V1)})}
})
})}
_ = functorNonEmpty1_2_31
// TAST (Let): foldableNonEmpty1_3_32 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_3_32 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_4_33 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_4_33 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_4_33
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_7_34 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_7_34 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_7_34
var go__go_8_35_55 gopurs_runtime.Value
go__go_8_35_55 = gopurs_runtime.Func(func(b_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_9_loop gopurs_runtime.Value = b_9_loop_val
var v_10_loop gopurs_runtime.Value = v_10_loop_val
go__go_8_35_55:
for {
if false { continue go__go_8_35_55 }
var b_9 gopurs_runtime.Value = b_9_loop
_ = b_9
var v_10 gopurs_runtime.Value = v_10_loop
_ = v_10
var __t36 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 1358893437 && v_10.UnsafePtr == nil) {
__t36 = b_9
goto end_branch_36
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 1358893437 && v_10.UnsafePtr != nil) {
b_9_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_7_34.V0), b_9, gopurs_runtime.Apply(f_5, (*Constructor_Data_List_Types_Cons)(v_10.UnsafePtr).V0))
v_10_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_10.UnsafePtr).V1)}
continue go__go_8_35_55
__t36 = gopurs_runtime.Value{}
goto end_branch_36
} else {

}
}
{
__t36 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_36:
return __t36
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_4_33.V0), gopurs_runtime.Apply(f_5, (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V0), gopurs_runtime.Apply2(go__go_8_35_55, gopurs_runtime.RecordGet(dictMonoid_3, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_6_37_56 gopurs_runtime.Value
go__go_6_37_56 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_7_loop gopurs_runtime.Value = b_7_loop_val
var v_8_loop gopurs_runtime.Value = v_8_loop_val
go__go_6_37_56:
for {
if false { continue go__go_6_37_56 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var v_8 gopurs_runtime.Value = v_8_loop
_ = v_8
var __t38 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 1358893437 && v_8.UnsafePtr == nil) {
__t38 = b_7
goto end_branch_38
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 1358893437 && v_8.UnsafePtr != nil) {
b_7_loop = gopurs_runtime.Apply2(f_3, b_7, (*Constructor_Data_List_Types_Cons)(v_8.UnsafePtr).V0)
v_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_8.UnsafePtr).V1)}
continue go__go_6_37_56
__t38 = gopurs_runtime.Value{}
goto end_branch_38
} else {

}
}
{
__t38 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_38:
return __t38
}
}()
})
})
return gopurs_runtime.Apply2(go__go_6_37_56, gopurs_runtime.Apply2(f_3, b_4, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_6_39_57 gopurs_runtime.Value
go__go_6_39_57 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_7_loop gopurs_runtime.Value = b_7_loop_val
var v_8_loop gopurs_runtime.Value = v_8_loop_val
go__go_6_39_57:
for {
if false { continue go__go_6_39_57 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var v_8 gopurs_runtime.Value = v_8_loop
_ = v_8
var __t40 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 1358893437 && v_8.UnsafePtr == nil) {
__t40 = b_7
goto end_branch_40
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 1358893437 && v_8.UnsafePtr != nil) {
b_7_loop = gopurs_runtime.Apply2(f_3, (*Constructor_Data_List_Types_Cons)(v_8.UnsafePtr).V0, b_7)
v_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_8.UnsafePtr).V1)}
continue go__go_6_39_57
__t40 = gopurs_runtime.Value{}
goto end_branch_40
} else {

}
}
{
__t40 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_40:
return __t40
}
}()
})
})
var go__go_7_41_58 gopurs_runtime.Value
go__go_7_41_58 = gopurs_runtime.Func(func(v_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_8_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_8_loop_val)
var v1_9_loop gopurs_runtime.Value = v1_9_loop_val
go__go_7_41_58:
for {
if false { continue go__go_7_41_58 }
var v_8 *Constructor_Data_List_Types_Cons = v_8_loop
_ = v_8
var v1_9 gopurs_runtime.Value = v1_9_loop
_ = v1_9
var __t42 *Constructor_Data_List_Types_Cons
{
if (v1_9.Type == 9 && v1_9.IntVal == 1358893437 && v1_9.UnsafePtr == nil) {
__t42 = v_8
goto end_branch_42
} else {

}
}
{
if (v1_9.Type == 9 && v1_9.IntVal == 1358893437 && v1_9.UnsafePtr != nil) {
v_8_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_9.UnsafePtr).V0, v_8})})
v1_9_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_9.UnsafePtr).V1)}
continue go__go_7_41_58
__t42 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_42
} else {

}
}
{
__t42 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_42:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t42)}
}
}()
})
})
return gopurs_runtime.Apply2(f_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_6_39_57, b_4, gopurs_runtime.Apply2(go__go_7_41_58, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1)))
})
})
})}
_ = foldableNonEmpty1_3_32
// TAST (Let): traversableNonEmpty1_2_30 -> *Constructor_Data_Traversable_Traversable
traversableNonEmpty1_2_30 := &Constructor_Data_Traversable_Traversable{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_3_32)}
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorNonEmpty1_2_31)}
}), gopurs_runtime.Func(func(dictApplicative_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_5_43 -> *Constructor_Control_Apply_Apply
Apply0_5_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_5_43
// TAST (Let): Functor0_6_44 -> *Constructor_Data_Functor_Functor
Functor0_6_44 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_44
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_8_45 -> *Constructor_Control_Apply_Apply
Apply0_8_45 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_8_45
var go__go_9_46_59 gopurs_runtime.Value
go__go_9_46_59 = gopurs_runtime.Func(func(b_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_10_loop gopurs_runtime.Value = b_10_loop_val
var v_11_loop gopurs_runtime.Value = v_11_loop_val
go__go_9_46_59:
for {
if false { continue go__go_9_46_59 }
var b_10 gopurs_runtime.Value = b_10_loop
_ = b_10
var v_11 gopurs_runtime.Value = v_11_loop
_ = v_11
var __t47 gopurs_runtime.Value
{
if (v_11.Type == 9 && v_11.IntVal == 1358893437 && v_11.UnsafePtr == nil) {
__t47 = b_10
goto end_branch_47
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 1358893437 && v_11.UnsafePtr != nil) {
b_10_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_11.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_10)})}
v_11_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_11.UnsafePtr).V1)}
continue go__go_9_46_59
__t47 = gopurs_runtime.Value{}
goto end_branch_47
} else {

}
}
{
__t47 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_47:
return __t47
}
}()
})
})
var go__go_9_48_60 gopurs_runtime.Value
go__go_9_48_60 = gopurs_runtime.Func(func(b_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_10_loop gopurs_runtime.Value = b_10_loop_val
var v_11_loop gopurs_runtime.Value = v_11_loop_val
go__go_9_48_60:
for {
if false { continue go__go_9_48_60 }
var b_10 gopurs_runtime.Value = b_10_loop
_ = b_10
var v_11 gopurs_runtime.Value = v_11_loop
_ = v_11
var __t49 gopurs_runtime.Value
{
if (v_11.Type == 9 && v_11.IntVal == 1358893437 && v_11.UnsafePtr == nil) {
__t49 = b_10
goto end_branch_49
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 1358893437 && v_11.UnsafePtr != nil) {
b_10_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_8_45.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_8_45.V0), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(b_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, a_13, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_12)})}
})
}), b_10), (*Constructor_Data_List_Types_Cons)(v_11.UnsafePtr).V0)
v_11_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_11.UnsafePtr).V1)}
continue go__go_9_48_60
__t49 = gopurs_runtime.Value{}
goto end_branch_49
} else {

}
}
{
__t49 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_49:
return __t49
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_5_43.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_44.V0), Get_Data_NonEmpty_NonEmpty(), (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(go__go_9_46_59, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), gopurs_runtime.Apply2(go__go_9_48_60, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1)))
})
}), gopurs_runtime.Func(func(dictApplicative_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_5_50 -> *Constructor_Control_Apply_Apply
Apply0_5_50 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_5_50
// TAST (Let): Functor0_6_51 -> *Constructor_Data_Functor_Functor
Functor0_6_51 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_51
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_9_52 -> *Constructor_Control_Apply_Apply
Apply0_9_52 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_9_52
var go__go_10_53_61 gopurs_runtime.Value
go__go_10_53_61 = gopurs_runtime.Func(func(b_11_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_11_loop gopurs_runtime.Value = b_11_loop_val
var v_12_loop gopurs_runtime.Value = v_12_loop_val
go__go_10_53_61:
for {
if false { continue go__go_10_53_61 }
var b_11 gopurs_runtime.Value = b_11_loop
_ = b_11
var v_12 gopurs_runtime.Value = v_12_loop
_ = v_12
var __t54 gopurs_runtime.Value
{
if (v_12.Type == 9 && v_12.IntVal == 1358893437 && v_12.UnsafePtr == nil) {
__t54 = b_11
goto end_branch_54
} else {

}
}
{
if (v_12.Type == 9 && v_12.IntVal == 1358893437 && v_12.UnsafePtr != nil) {
b_11_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_12.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_11)})}
v_12_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_12.UnsafePtr).V1)}
continue go__go_10_53_61
__t54 = gopurs_runtime.Value{}
goto end_branch_54
} else {

}
}
{
__t54 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_54:
return __t54
}
}()
})
})
var go__go_10_55_62 gopurs_runtime.Value
go__go_10_55_62 = gopurs_runtime.Func(func(b_11_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_11_loop gopurs_runtime.Value = b_11_loop_val
var v_12_loop gopurs_runtime.Value = v_12_loop_val
go__go_10_55_62:
for {
if false { continue go__go_10_55_62 }
var b_11 gopurs_runtime.Value = b_11_loop
_ = b_11
var v_12 gopurs_runtime.Value = v_12_loop
_ = v_12
var __t56 gopurs_runtime.Value
{
if (v_12.Type == 9 && v_12.IntVal == 1358893437 && v_12.UnsafePtr == nil) {
__t56 = b_11
goto end_branch_56
} else {

}
}
{
if (v_12.Type == 9 && v_12.IntVal == 1358893437 && v_12.UnsafePtr != nil) {
b_11_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_9_52.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_9_52.V0), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(b_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, a_14, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_13)})}
})
}), b_11), gopurs_runtime.Apply(f_7, (*Constructor_Data_List_Types_Cons)(v_12.UnsafePtr).V0))
v_12_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_12.UnsafePtr).V1)}
continue go__go_10_55_62
__t56 = gopurs_runtime.Value{}
goto end_branch_56
} else {

}
}
{
__t56 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_56:
return __t56
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_5_50.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_51.V0), Get_Data_NonEmpty_NonEmpty(), gopurs_runtime.Apply(f_7, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(go__go_10_53_61, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), gopurs_runtime.Apply2(go__go_10_55_62, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1)))
})
})
})}
_ = traversableNonEmpty1_2_30
return gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(&Constructor_Data_TraversableWithIndex_TraversableWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(foldableWithIndexNonEmpty1_1_7)}
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(functorWithIndex1_0_0)}
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(traversableNonEmpty1_2_30)}
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_4_57 -> *Constructor_Control_Apply_Apply
Apply0_4_57 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_4_57
// TAST (Let): Functor0_5_58 -> *Constructor_Data_Functor_Functor
Functor0_5_58 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_58
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_8_59 -> *Constructor_Control_Apply_Apply
Apply0_8_59 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_8_59
var go__go_9_60_63 gopurs_runtime.Value
go__go_9_60_63 = gopurs_runtime.Func(func(b_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_10_loop gopurs_runtime.Value = b_10_loop_val
var v_11_loop gopurs_runtime.Value = v_11_loop_val
go__go_9_60_63:
for {
if false { continue go__go_9_60_63 }
var b_10 gopurs_runtime.Value = b_10_loop
_ = b_10
var v_11 gopurs_runtime.Value = v_11_loop
_ = v_11
var __t61 gopurs_runtime.Value
{
if (v_11.Type == 9 && v_11.IntVal == 1358893437 && v_11.UnsafePtr == nil) {
__t61 = b_10
goto end_branch_61
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 1358893437 && v_11.UnsafePtr != nil) {
b_10_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_11.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_10)})}
v_11_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_11.UnsafePtr).V1)}
continue go__go_9_60_63
__t61 = gopurs_runtime.Value{}
goto end_branch_61
} else {

}
}
{
__t61 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_61:
return __t61
}
}()
})
})
var go__go_9_62_64 gopurs_runtime.Value
go__go_9_62_64 = gopurs_runtime.Func(func(b_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_10_loop gopurs_runtime.Value = b_10_loop_val
var v_11_loop gopurs_runtime.Value = v_11_loop_val
go__go_9_62_64:
for {
if false { continue go__go_9_62_64 }
var b_10 gopurs_runtime.Value = b_10_loop
_ = b_10
var v_11 gopurs_runtime.Value = v_11_loop
_ = v_11
var __t63 gopurs_runtime.Value
{
if (v_11.Type == 9 && v_11.IntVal == 1358893437 && v_11.UnsafePtr == nil) {
__t63 = b_10
goto end_branch_63
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 1358893437 && v_11.UnsafePtr != nil) {
b_10_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_10.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_8_59.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_8_59.V0), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(b_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, a_13, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_12)})}
})
}), (*Constructor_Data_Tuple_Tuple)(b_10.UnsafePtr).V1), gopurs_runtime.Apply2(f_6, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple)(b_10.UnsafePtr).V0.IntVal)})}, (*Constructor_Data_List_Types_Cons)(v_11.UnsafePtr).V0))})}
v_11_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_11.UnsafePtr).V1)}
continue go__go_9_62_64
__t63 = gopurs_runtime.Value{}
goto end_branch_63
} else {

}
}
{
__t63 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_63:
return __t63
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_57.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_58.V0), Get_Data_NonEmpty_NonEmpty(), gopurs_runtime.Apply2(f_6, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(go__go_9_60_63, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}, (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_9_62_64, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})})}, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1).UnsafePtr).V1))
})
})
})})}
}()
	})
	return cache_Data_List_Types_traversableWithIndexNonEmpty
}

var cache_Data_List_Types_traversableWithIndexNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_traversableWithIndexNonEmptyList sync.Once
func Get_Data_List_Types_traversableWithIndexNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_traversableWithIndexNonEmptyList.Do(func() {
		cache_Data_List_Types_traversableWithIndexNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(&Constructor_Data_TraversableWithIndex_TraversableWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex](Get_Data_List_Types_foldableWithIndexNonEmptyList()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex](Get_Data_List_Types_functorWithIndexNonEmptyList()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorNonEmpty1_1_0 -> *Constructor_Data_Functor_Functor
functorNonEmpty1_1_0 := &Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(f_1, (*Constructor_Data_NonEmpty_NonEmpty)(m_2.UnsafePtr).V0), gopurs_runtime.Apply(Call_Data_List_Types_listMap(f_1), (*Constructor_Data_NonEmpty_NonEmpty)(m_2.UnsafePtr).V1)})}
})
})}
_ = functorNonEmpty1_1_0
// TAST (Let): foldableNonEmpty1_2_1 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_2_1 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_3_2 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_3_2
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_6_3 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_6_3
var go__go_7_4_65 gopurs_runtime.Value
go__go_7_4_65 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_8_loop gopurs_runtime.Value = b_8_loop_val
var v_9_loop gopurs_runtime.Value = v_9_loop_val
go__go_7_4_65:
for {
if false { continue go__go_7_4_65 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var v_9 gopurs_runtime.Value = v_9_loop
_ = v_9
var __t5 gopurs_runtime.Value
{
if (v_9.Type == 9 && v_9.IntVal == 1358893437 && v_9.UnsafePtr == nil) {
__t5 = b_8
goto end_branch_5
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 1358893437 && v_9.UnsafePtr != nil) {
b_8_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_3.V0), b_8, gopurs_runtime.Apply(f_4, (*Constructor_Data_List_Types_Cons)(v_9.UnsafePtr).V0))
v_9_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_9.UnsafePtr).V1)}
continue go__go_7_4_65
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_2.V0), gopurs_runtime.Apply(f_4, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0), gopurs_runtime.Apply2(go__go_7_4_65, gopurs_runtime.RecordGet(dictMonoid_2, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_6_66 gopurs_runtime.Value
go__go_5_6_66 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_6_66:
for {
if false { continue go__go_5_6_66 }
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
b_6_loop = gopurs_runtime.Apply2(f_2, b_6, (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0)
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_6_66
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
return gopurs_runtime.Apply2(go__go_5_6_66, gopurs_runtime.Apply2(f_2, b_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_8_67 gopurs_runtime.Value
go__go_5_8_67 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_8_67:
for {
if false { continue go__go_5_8_67 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t9 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr == nil) {
__t9 = b_6
goto end_branch_9
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr != nil) {
b_6_loop = gopurs_runtime.Apply2(f_2, (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0, b_6)
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_8_67
__t9 = gopurs_runtime.Value{}
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
}
}()
})
})
var go__go_6_10_68 gopurs_runtime.Value
go__go_6_10_68 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_7_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_7_loop_val)
var v1_8_loop gopurs_runtime.Value = v1_8_loop_val
go__go_6_10_68:
for {
if false { continue go__go_6_10_68 }
var v_7 *Constructor_Data_List_Types_Cons = v_7_loop
_ = v_7
var v1_8 gopurs_runtime.Value = v1_8_loop
_ = v1_8
var __t11 *Constructor_Data_List_Types_Cons
{
if (v1_8.Type == 9 && v1_8.IntVal == 1358893437 && v1_8.UnsafePtr == nil) {
__t11 = v_7
goto end_branch_11
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 1358893437 && v1_8.UnsafePtr != nil) {
v_7_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_8.UnsafePtr).V0, v_7})})
v1_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_8.UnsafePtr).V1)}
continue go__go_6_10_68
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t11)}
}
}()
})
})
return gopurs_runtime.Apply2(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_5_8_67, b_3, gopurs_runtime.Apply2(go__go_6_10_68, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1)))
})
})
})}
_ = foldableNonEmpty1_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(&Constructor_Data_Traversable_Traversable{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_2_1)}
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorNonEmpty1_1_0)}
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_4_12 -> *Constructor_Control_Apply_Apply
Apply0_4_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_4_12
// TAST (Let): Functor0_5_13 -> *Constructor_Data_Functor_Functor
Functor0_5_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_13
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_7_14 -> *Constructor_Control_Apply_Apply
Apply0_7_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_7_14
var go__go_8_15_69 gopurs_runtime.Value
go__go_8_15_69 = gopurs_runtime.Func(func(b_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_9_loop gopurs_runtime.Value = b_9_loop_val
var v_10_loop gopurs_runtime.Value = v_10_loop_val
go__go_8_15_69:
for {
if false { continue go__go_8_15_69 }
var b_9 gopurs_runtime.Value = b_9_loop
_ = b_9
var v_10 gopurs_runtime.Value = v_10_loop
_ = v_10
var __t16 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 1358893437 && v_10.UnsafePtr == nil) {
__t16 = b_9
goto end_branch_16
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 1358893437 && v_10.UnsafePtr != nil) {
b_9_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_10.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_9)})}
v_10_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_10.UnsafePtr).V1)}
continue go__go_8_15_69
__t16 = gopurs_runtime.Value{}
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
return __t16
}
}()
})
})
var go__go_8_17_70 gopurs_runtime.Value
go__go_8_17_70 = gopurs_runtime.Func(func(b_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_9_loop gopurs_runtime.Value = b_9_loop_val
var v_10_loop gopurs_runtime.Value = v_10_loop_val
go__go_8_17_70:
for {
if false { continue go__go_8_17_70 }
var b_9 gopurs_runtime.Value = b_9_loop
_ = b_9
var v_10 gopurs_runtime.Value = v_10_loop
_ = v_10
var __t18 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 1358893437 && v_10.UnsafePtr == nil) {
__t18 = b_9
goto end_branch_18
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 1358893437 && v_10.UnsafePtr != nil) {
b_9_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_7_14.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_7_14.V0), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(b_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, a_12, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_11)})}
})
}), b_9), (*Constructor_Data_List_Types_Cons)(v_10.UnsafePtr).V0)
v_10_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_10.UnsafePtr).V1)}
continue go__go_8_17_70
__t18 = gopurs_runtime.Value{}
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
return __t18
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_12.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_13.V0), Get_Data_NonEmpty_NonEmpty(), (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(go__go_8_15_69, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), gopurs_runtime.Apply2(go__go_8_17_70, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V1)))
})
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_4_19 -> *Constructor_Control_Apply_Apply
Apply0_4_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_4_19
// TAST (Let): Functor0_5_20 -> *Constructor_Data_Functor_Functor
Functor0_5_20 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_20
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_8_21 -> *Constructor_Control_Apply_Apply
Apply0_8_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_8_21
var go__go_9_22_71 gopurs_runtime.Value
go__go_9_22_71 = gopurs_runtime.Func(func(b_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_10_loop gopurs_runtime.Value = b_10_loop_val
var v_11_loop gopurs_runtime.Value = v_11_loop_val
go__go_9_22_71:
for {
if false { continue go__go_9_22_71 }
var b_10 gopurs_runtime.Value = b_10_loop
_ = b_10
var v_11 gopurs_runtime.Value = v_11_loop
_ = v_11
var __t23 gopurs_runtime.Value
{
if (v_11.Type == 9 && v_11.IntVal == 1358893437 && v_11.UnsafePtr == nil) {
__t23 = b_10
goto end_branch_23
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 1358893437 && v_11.UnsafePtr != nil) {
b_10_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_11.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_10)})}
v_11_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_11.UnsafePtr).V1)}
continue go__go_9_22_71
__t23 = gopurs_runtime.Value{}
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
return __t23
}
}()
})
})
var go__go_9_24_72 gopurs_runtime.Value
go__go_9_24_72 = gopurs_runtime.Func(func(b_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_10_loop gopurs_runtime.Value = b_10_loop_val
var v_11_loop gopurs_runtime.Value = v_11_loop_val
go__go_9_24_72:
for {
if false { continue go__go_9_24_72 }
var b_10 gopurs_runtime.Value = b_10_loop
_ = b_10
var v_11 gopurs_runtime.Value = v_11_loop
_ = v_11
var __t25 gopurs_runtime.Value
{
if (v_11.Type == 9 && v_11.IntVal == 1358893437 && v_11.UnsafePtr == nil) {
__t25 = b_10
goto end_branch_25
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 1358893437 && v_11.UnsafePtr != nil) {
b_10_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_8_21.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_8_21.V0), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(b_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, a_13, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_12)})}
})
}), b_10), gopurs_runtime.Apply(f_6, (*Constructor_Data_List_Types_Cons)(v_11.UnsafePtr).V0))
v_11_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_11.UnsafePtr).V1)}
continue go__go_9_24_72
__t25 = gopurs_runtime.Value{}
goto end_branch_25
} else {

}
}
{
__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_25:
return __t25
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_19.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_20.V0), Get_Data_NonEmpty_NonEmpty(), gopurs_runtime.Apply(f_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(go__go_9_22_71, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), gopurs_runtime.Apply2(go__go_9_24_72, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1)))
})
})
})})}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_26 -> *Constructor_Data_Functor_Functor
Functor0_1_26 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_26
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_4_27 -> *Constructor_Control_Apply_Apply
Apply0_4_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_4_27
var go__go_5_28_73 gopurs_runtime.Value
go__go_5_28_73 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_28_73:
for {
if false { continue go__go_5_28_73 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t29 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr == nil) {
__t29 = b_6
goto end_branch_29
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr != nil) {
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_6)})}
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_28_73
__t29 = gopurs_runtime.Value{}
goto end_branch_29
} else {

}
}
{
__t29 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_29:
return __t29
}
}()
})
})
var go__go_5_30_74 gopurs_runtime.Value
go__go_5_30_74 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_30_74:
for {
if false { continue go__go_5_30_74 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t31 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr == nil) {
__t31 = b_6
goto end_branch_31
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr != nil) {
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_27.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_4_27.V0), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, a_9, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_8)})}
})
}), (*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V1), gopurs_runtime.Apply2(f_2, gopurs_runtime.Int((1) + ((*Constructor_Data_Tuple_Tuple)(b_6.UnsafePtr).V0.IntVal)), (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0))})}
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_30_74
__t31 = gopurs_runtime.Value{}
goto end_branch_31
} else {

}
}
{
__t31 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_31:
return __t31
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_26.V0), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return x_4
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_NonEmpty_NonEmpty(), gopurs_runtime.Apply2(f_2, gopurs_runtime.Int(0), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(go__go_5_28_73, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}, (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_5_30_74, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})})}, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1).UnsafePtr).V1)))
})
})
})})}
	})
	return cache_Data_List_Types_traversableWithIndexNonEmptyList
}

var cache_Data_List_Types_unfoldable1List gopurs_runtime.Value
var once_Data_List_Types_unfoldable1List sync.Once
func Get_Data_List_Types_unfoldable1List() gopurs_runtime.Value {
	once_Data_List_Types_unfoldable1List.Do(func() {
		cache_Data_List_Types_unfoldable1List = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(&Constructor_Data_Unfoldable1_Unfoldable1{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_75 gopurs_runtime.Value
go__go_2_0_75 = gopurs_runtime.Func(func(source_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var source_3_loop gopurs_runtime.Value = source_3_loop_val
var memo_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](memo_4_loop_val)
go__go_2_0_75:
for {
if false { continue go__go_2_0_75 }
var source_3 gopurs_runtime.Value = source_3_loop
_ = source_3
var memo_4 *Constructor_Data_List_Types_Cons = memo_4_loop
_ = memo_4
// TAST (Let): v_5_1 -> *Constructor_Data_Tuple_Tuple
v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_0, source_3))
_ = v_5_1
var __t6 *Constructor_Data_List_Types_Cons
{
var __t_tag_2 gopurs_runtime.Value = (v_5_1).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 930809136 && __t_tag_2.UnsafePtr != nil) {
source_3_loop = (*Constructor_Data_Maybe_Just)((v_5_1).V1.UnsafePtr).V0
memo_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v_5_1).V0, memo_4})})
continue go__go_2_0_75
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_6
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (v_5_1).V1
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 930809136 && __t_tag_3.UnsafePtr == nil) {
var go__go_6_4_76 gopurs_runtime.Value
go__go_6_4_76 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_7_loop gopurs_runtime.Value = b_7_loop_val
var v_8_loop gopurs_runtime.Value = v_8_loop_val
go__go_6_4_76:
for {
if false { continue go__go_6_4_76 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var v_8 gopurs_runtime.Value = v_8_loop
_ = v_8
var __t5 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 1358893437 && v_8.UnsafePtr == nil) {
__t5 = b_7
goto end_branch_5
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 1358893437 && v_8.UnsafePtr != nil) {
b_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_8.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_7)})}
v_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_8.UnsafePtr).V1)}
continue go__go_6_4_76
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}()
})
})
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_6_4_76, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v_5_1).V0, memo_4})}))
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t6)}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_2_0_75, b_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}
})
})})}
	})
	return cache_Data_List_Types_unfoldable1List
}

var cache_Data_List_Types_unfoldableList gopurs_runtime.Value
var once_Data_List_Types_unfoldableList sync.Once
func Get_Data_List_Types_unfoldableList() gopurs_runtime.Value {
	once_Data_List_Types_unfoldableList.Do(func() {
		cache_Data_List_Types_unfoldableList = gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(&Constructor_Data_Unfoldable_Unfoldable{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](Get_Data_List_Types_unfoldable1List()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_77 gopurs_runtime.Value
go__go_2_0_77 = gopurs_runtime.Func(func(source_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var source_3_loop gopurs_runtime.Value = source_3_loop_val
var memo_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](memo_4_loop_val)
go__go_2_0_77:
for {
if false { continue go__go_2_0_77 }
var source_3 gopurs_runtime.Value = source_3_loop
_ = source_3
var memo_4 *Constructor_Data_List_Types_Cons = memo_4_loop
_ = memo_4
// TAST (Let): v_5_1 -> *Constructor_Data_Maybe_Just
v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_0, source_3))
_ = v_5_1
var __t4 *Constructor_Data_List_Types_Cons
{
if (v_5_1 == nil) {
var go__go_6_2_78 gopurs_runtime.Value
go__go_6_2_78 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_7_loop gopurs_runtime.Value = b_7_loop_val
var v_8_loop gopurs_runtime.Value = v_8_loop_val
go__go_6_2_78:
for {
if false { continue go__go_6_2_78 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var v_8 gopurs_runtime.Value = v_8_loop
_ = v_8
var __t3 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 1358893437 && v_8.UnsafePtr == nil) {
__t3 = b_7
goto end_branch_3
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 1358893437 && v_8.UnsafePtr != nil) {
b_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_8.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_7)})}
v_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_8.UnsafePtr).V1)}
continue go__go_6_2_78
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
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_6_2_78, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(memo_4)}))
goto end_branch_4
} else {

}
}
{
if (v_5_1 != nil) {
source_3_loop = (*Constructor_Data_Tuple_Tuple)((v_5_1).V0.UnsafePtr).V1
memo_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Tuple_Tuple)((v_5_1).V0.UnsafePtr).V0, memo_4})})
continue go__go_2_0_77
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t4)}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_2_0_77, b_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}
})
})})}
	})
	return cache_Data_List_Types_unfoldableList
}

var cache_Data_List_Types_unfoldable1NonEmptyList gopurs_runtime.Value
var once_Data_List_Types_unfoldable1NonEmptyList sync.Once
func Get_Data_List_Types_unfoldable1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_unfoldable1NonEmptyList.Do(func() {
		cache_Data_List_Types_unfoldable1NonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(&Constructor_Data_Unfoldable1_Unfoldable1{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 -> *Constructor_Data_Tuple_Tuple
__local_var_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(f_0, b_1))
_ = __local_var_2_1
var go__go_3_2_79 gopurs_runtime.Value
go__go_3_2_79 = gopurs_runtime.Func(func(source_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var source_4_loop gopurs_runtime.Value = source_4_loop_val
var memo_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](memo_5_loop_val)
go__go_3_2_79:
for {
if false { continue go__go_3_2_79 }
var source_4 gopurs_runtime.Value = source_4_loop
_ = source_4
var memo_5 *Constructor_Data_List_Types_Cons = memo_5_loop
_ = memo_5
var __t4 *Constructor_Data_Maybe_Just
{
if (source_4.Type == 9 && source_4.IntVal == 930809136 && source_4.UnsafePtr != nil) {
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Maybe_Just)(source_4.UnsafePtr).V0)}
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_4:
// TAST (Let): v_6_3 -> *Constructor_Data_Maybe_Just
var v_6_3 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)})
var __t7 *Constructor_Data_List_Types_Cons
{
if (v_6_3 == nil) {
var go__go_7_5_80 gopurs_runtime.Value
go__go_7_5_80 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_8_loop gopurs_runtime.Value = b_8_loop_val
var v_9_loop gopurs_runtime.Value = v_9_loop_val
go__go_7_5_80:
for {
if false { continue go__go_7_5_80 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var v_9 gopurs_runtime.Value = v_9_loop
_ = v_9
var __t6 gopurs_runtime.Value
{
if (v_9.Type == 9 && v_9.IntVal == 1358893437 && v_9.UnsafePtr == nil) {
__t6 = b_8
goto end_branch_6
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 1358893437 && v_9.UnsafePtr != nil) {
b_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_9.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_8)})}
v_9_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_9.UnsafePtr).V1)}
continue go__go_7_5_80
__t6 = gopurs_runtime.Value{}
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
}()
})
})
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_7_5_80, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(memo_5)}))
goto end_branch_7
} else {

}
}
{
if (v_6_3 != nil) {
source_4_loop = (*Constructor_Data_Tuple_Tuple)((v_6_3).V0.UnsafePtr).V1
memo_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Tuple_Tuple)((v_6_3).V0.UnsafePtr).V0, memo_5})})
continue go__go_3_2_79
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t7)}
}
}()
})
})
// TAST (Let): __local_var_2_0 -> *Constructor_Data_Tuple_Tuple
var __local_var_2_0 *Constructor_Data_Tuple_Tuple = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (__local_var_2_1).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_3_2_79, (__local_var_2_1).V1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}})})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (__local_var_2_0).V0, (__local_var_2_0).V1})}
})
})})}
	})
	return cache_Data_List_Types_unfoldable1NonEmptyList
}

var cache_Data_List_Types_foldable1NonEmptyList gopurs_runtime.Value
var once_Data_List_Types_foldable1NonEmptyList sync.Once
func Get_Data_List_Types_foldable1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_foldable1NonEmptyList.Do(func() {
		cache_Data_List_Types_foldable1NonEmptyList = func() gopurs_runtime.Value {
// TAST (Let): foldableNonEmpty1_0_0 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_0_0 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_1 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_4_2 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_4_2
var go__go_5_3_81 gopurs_runtime.Value
go__go_5_3_81 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_3_81:
for {
if false { continue go__go_5_3_81 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t4 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr == nil) {
__t4 = b_6
goto end_branch_4
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr != nil) {
b_6_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_4_2.V0), b_6, gopurs_runtime.Apply(f_2, (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0))
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_3_81
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
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_1.V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0), gopurs_runtime.Apply2(go__go_5_3_81, gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_5_82 gopurs_runtime.Value
go__go_3_5_82 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_5_82:
for {
if false { continue go__go_3_5_82 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t6 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t6 = b_4
goto end_branch_6
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Apply2(f_0, b_4, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0)
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_5_82
__t6 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Apply2(go__go_3_5_82, gopurs_runtime.Apply2(f_0, b_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_7_83 gopurs_runtime.Value
go__go_3_7_83 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_7_83:
for {
if false { continue go__go_3_7_83 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t8 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t8 = b_4
goto end_branch_8
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0, b_4)
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_7_83
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
var go__go_4_9_84 gopurs_runtime.Value
go__go_4_9_84 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_9_84:
for {
if false { continue go__go_4_9_84 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t10 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t10 = v_5
goto end_branch_10
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_9_84
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t10)}
}
}()
})
})
return gopurs_runtime.Apply2(f_0, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_3_7_83, b_1, gopurs_runtime.Apply2(go__go_4_9_84, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)))
})
})
})}
_ = foldableNonEmpty1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Foldable_Foldable1{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_0_0)}
}), gopurs_runtime.Func(func(dictSemigroup_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_11_85 gopurs_runtime.Value
go__go_4_11_85 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_11_85:
for {
if false { continue go__go_4_11_85 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t12 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t12 = b_5
goto end_branch_12
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_1, "append"), b_5, gopurs_runtime.Apply(f_2, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0))
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_11_85
__t12 = gopurs_runtime.Value{}
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
}
}()
})
})
return gopurs_runtime.Apply2(go__go_4_11_85, gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_13_86 gopurs_runtime.Value
go__go_3_13_86 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_13_86:
for {
if false { continue go__go_3_13_86 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t14 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t14 = b_4
goto end_branch_14
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Apply2(f_1, b_4, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0)
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_13_86
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
return gopurs_runtime.Apply2(go__go_3_13_86, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_15 -> gopurs_runtime.Value
__local_var_3_15 := gopurs_runtime.Apply(f_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0)
_ = __local_var_3_15
var go__go_4_17_87 gopurs_runtime.Value
go__go_4_17_87 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_17_87:
for {
if false { continue go__go_4_17_87 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t20 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t20 = b_5
goto end_branch_20
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
// TAST (Let): __local_var_7_18 -> gopurs_runtime.Value
__local_var_7_18 := gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0)
_ = __local_var_7_18
var __t19 gopurs_runtime.Value
{
if (b_5.Type == 9 && b_5.IntVal == 930809136 && b_5.UnsafePtr == nil) {
__t19 = (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0
goto end_branch_19
} else {

}
}
{
if (b_5.Type == 9 && b_5.IntVal == 930809136 && b_5.UnsafePtr != nil) {
__t19 = gopurs_runtime.Apply(__local_var_7_18, (*Constructor_Data_Maybe_Just)(b_5.UnsafePtr).V0)
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, __t19})}
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_17_87
__t20 = gopurs_runtime.Value{}
goto end_branch_20
} else {

}
}
{
__t20 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_20:
return __t20
}
}()
})
})
var go__go_5_21_88 gopurs_runtime.Value
go__go_5_21_88 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_6_loop_val)
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_21_88:
for {
if false { continue go__go_5_21_88 }
var v_6 *Constructor_Data_List_Types_Cons = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t22 *Constructor_Data_List_Types_Cons
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr == nil) {
__t22 = v_6
goto end_branch_22
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V0, v_6})})
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V1)}
continue go__go_5_21_88
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_22
} else {

}
}
{
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_22:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t22)}
}
}()
})
})
// TAST (Let): __local_var_4_16 -> *Constructor_Data_Maybe_Just
__local_var_4_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(go__go_4_17_87, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Apply2(go__go_5_21_88, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)))
_ = __local_var_4_16
var __t23 gopurs_runtime.Value
{
if (__local_var_4_16 == nil) {
__t23 = (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0
goto end_branch_23
} else {

}
}
{
if (__local_var_4_16 != nil) {
__t23 = gopurs_runtime.Apply(__local_var_3_15, (__local_var_4_16).V0)
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
return __t23
})
})})}
}()
	})
	return cache_Data_List_Types_foldable1NonEmptyList
}

var cache_Data_List_Types_extendNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_extendNonEmptyList sync.Once
func Get_Data_List_Types_extendNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_extendNonEmptyList.Do(func() {
		cache_Data_List_Types_extendNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(&Constructor_Control_Extend_Extend{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Types_functorNonEmptyList()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_89 gopurs_runtime.Value
go__go_2_0_89 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_0_89:
for {
if false { continue go__go_2_0_89 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t1 = b_3
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
b_3_loop = gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(b_3, "acc"))})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(b_3, "acc")))}})}), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(b_3, "val"))})})
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_0_89
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
})
var go__go_3_2_90 gopurs_runtime.Value
go__go_3_2_90 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_2_90:
for {
if false { continue go__go_3_2_90 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t3 = v_4
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_2_90
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_1))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(gopurs_runtime.Apply2(go__go_2_0_89, gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), gopurs_runtime.Apply2(go__go_3_2_90, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_1.UnsafePtr).V1))})), "val")))}})}
})
})})}
	})
	return cache_Data_List_Types_extendNonEmptyList
}

var cache_Data_List_Types_extendList gopurs_runtime.Value
var once_Data_List_Types_extendList sync.Once
func Get_Data_List_Types_extendList() gopurs_runtime.Value {
	once_Data_List_Types_extendList.Do(func() {
		cache_Data_List_Types_extendList = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(&Constructor_Control_Extend_Extend{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Types_functorList()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_List_Types_Cons
{
if (v1_1.Type == 9 && v1_1.IntVal == 1358893437 && v1_1.UnsafePtr == nil) {
__t4 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_4
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 1358893437 && v1_1.UnsafePtr != nil) {
var go__go_2_0_91 gopurs_runtime.Value
go__go_2_0_91 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_0_91:
for {
if false { continue go__go_2_0_91 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t1 = b_3
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
b_3_loop = gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(b_3, "acc"))})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(b_3, "acc"))})}), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(b_3, "val"))})})
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_0_91
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
})
var go__go_3_2_92 gopurs_runtime.Value
go__go_3_2_92 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_2_92:
for {
if false { continue go__go_3_2_92 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t3 = v_4
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_2_92
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
__t4 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_1))}), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(gopurs_runtime.Apply2(go__go_2_0_91, gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), gopurs_runtime.Apply2(go__go_3_2_92, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_1.UnsafePtr).V1)})), "val"))}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t4)}
})
})})}
	})
	return cache_Data_List_Types_extendList
}

var cache_Data_List_Types_eq1List gopurs_runtime.Value
var once_Data_List_Types_eq1List sync.Once
func Get_Data_List_Types_eq1List() gopurs_runtime.Value {
	once_Data_List_Types_eq1List.Do(func() {
		cache_Data_List_Types_eq1List = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq1{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_0_93 gopurs_runtime.Value
go__go_3_0_93 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
var v2_6_loop bool = (v2_6_loop_val.IntVal) != (0)
go__go_3_0_93:
for {
if false { continue go__go_3_0_93 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var v2_6 bool = v2_6_loop
_ = v2_6
var __t2 bool
{
if (v2_6) != (true) {
__t2 = false
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
var __t1 bool
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t1 = v2_6
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil)) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
v2_6_loop = (v2_6) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0).IntVal) != (0))
continue go__go_3_0_93
__t2 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
return gopurs_runtime.Bool(__t2)
}
}()
})
})
})
return gopurs_runtime.Bool((gopurs_runtime.Apply3(go__go_3_0_93, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_2))}, gopurs_runtime.Bool(true)).IntVal) != (0))
})
})
})})}
	})
	return cache_Data_List_Types_eq1List
}

var cache_Data_List_Types_eq1 gopurs_runtime.Value
var once_Data_List_Types_eq1 sync.Once
func Get_Data_List_Types_eq1() gopurs_runtime.Value {
	once_Data_List_Types_eq1.Do(func() {
		cache_Data_List_Types_eq1 = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_List_Types_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_2_box)))
})
	})
	return cache_Data_List_Types_eq1
}

var cache_Data_List_Types_eqNonEmpty gopurs_runtime.Value
var once_Data_List_Types_eqNonEmpty sync.Once
func Get_Data_List_Types_eqNonEmpty() gopurs_runtime.Value {
	once_Data_List_Types_eqNonEmpty.Do(func() {
		cache_Data_List_Types_eqNonEmpty = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Types_eqNonEmpty(dictEq_0_box)
})
	})
	return cache_Data_List_Types_eqNonEmpty
}

var cache_Data_List_Types_eq1NonEmptyList gopurs_runtime.Value
var once_Data_List_Types_eq1NonEmptyList sync.Once
func Get_Data_List_Types_eq1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_eq1NonEmptyList.Do(func() {
		cache_Data_List_Types_eq1NonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq1{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_and_3 bool = false
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_NonEmpty_NonEmpty)(x_1.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(y_2.UnsafePtr).V0).IntVal) != (0) {

var go__go_3_0_96 gopurs_runtime.Value
go__go_3_0_96 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
var v2_6_loop bool = (v2_6_loop_val.IntVal) != (0)
go__go_3_0_96:
for {
if false { continue go__go_3_0_96 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var v2_6 bool = v2_6_loop
_ = v2_6
var __t2 bool
{
if (v2_6) != (true) {
__t2 = false
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
var __t1 bool
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t1 = v2_6
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil)) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
v2_6_loop = (v2_6) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0).IntVal) != (0))
continue go__go_3_0_96
__t2 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
return gopurs_runtime.Bool(__t2)
}
}()
})
})
})
__t_and_3 = (gopurs_runtime.Apply3(go__go_3_0_96, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(x_1.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(y_2.UnsafePtr).V1))}, gopurs_runtime.Bool(true)).IntVal) != (0)
}
return gopurs_runtime.Bool(__t_and_3)
})
})
})})}
	})
	return cache_Data_List_Types_eq1NonEmptyList
}

var cache_Data_List_Types_eqList gopurs_runtime.Value
var once_Data_List_Types_eqList sync.Once
func Get_Data_List_Types_eqList() gopurs_runtime.Value {
	once_Data_List_Types_eqList.Do(func() {
		cache_Data_List_Types_eqList = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Types_eqList(dictEq_0_box)
})
	})
	return cache_Data_List_Types_eqList
}

var cache_Data_List_Types_eqNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_eqNonEmptyList sync.Once
func Get_Data_List_Types_eqNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_eqNonEmptyList.Do(func() {
		cache_Data_List_Types_eqNonEmptyList = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Types_eqNonEmptyList(dictEq_0_box)
})
	})
	return cache_Data_List_Types_eqNonEmptyList
}

var cache_Data_List_Types_ord1List gopurs_runtime.Value
var once_Data_List_Types_ord1List sync.Once
func Get_Data_List_Types_ord1List() gopurs_runtime.Value {
	once_Data_List_Types_ord1List.Do(func() {
		cache_Data_List_Types_ord1List = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord1{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1](Get_Data_List_Types_eq1List()))}
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_0_99 gopurs_runtime.Value
go__go_3_0_99 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_0_99:
for {
if false { continue go__go_3_0_99 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t4 uint32
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
var __t1 uint32
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t1 = 902936544
goto end_branch_1
} else {

}
}
{
__t1 = 1527465420
}
end_branch_1:
__t4 = __t1
goto end_branch_4
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t4 = 380165415
goto end_branch_4
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil)) {
// TAST (Let): v2_6_2 -> gopurs_runtime.Value
v2_6_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0)
_ = v2_6_2
var __t3 uint32
{
if (uint32(v2_6_2.IntVal) == 902936544) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_0_99
__t3 = uint32(gopurs_runtime.Value{}.IntVal)
goto end_branch_3
} else {

}
}
{
__t3 = uint32(v2_6_2.IntVal)
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t4), UnsafePtr: nil}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(go__go_3_0_99, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_2))}).IntVal)), UnsafePtr: nil}
})
})
})})}
	})
	return cache_Data_List_Types_ord1List
}

var cache_Data_List_Types_compare1 gopurs_runtime.Value
var once_Data_List_Types_compare1 sync.Once
func Get_Data_List_Types_compare1() gopurs_runtime.Value {
	once_Data_List_Types_compare1.Do(func() {
		cache_Data_List_Types_compare1 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_List_Types_compare1(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_2_box))), UnsafePtr: nil}
})
	})
	return cache_Data_List_Types_compare1
}

var cache_Data_List_Types_ordNonEmpty gopurs_runtime.Value
var once_Data_List_Types_ordNonEmpty sync.Once
func Get_Data_List_Types_ordNonEmpty() gopurs_runtime.Value {
	once_Data_List_Types_ordNonEmpty.Do(func() {
		cache_Data_List_Types_ordNonEmpty = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Types_ordNonEmpty(dictOrd_0_box)
})
	})
	return cache_Data_List_Types_ordNonEmpty
}

var cache_Data_List_Types_ord1NonEmptyList gopurs_runtime.Value
var once_Data_List_Types_ord1NonEmptyList sync.Once
func Get_Data_List_Types_ord1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_ord1NonEmptyList.Do(func() {
		cache_Data_List_Types_ord1NonEmptyList = func() gopurs_runtime.Value {
// TAST (Let): eq1NonEmpty1_0_0 -> *Constructor_Data_Eq_Eq1
eq1NonEmpty1_0_0 := &Constructor_Data_Eq_Eq1{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_and_4 bool = false
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_NonEmpty_NonEmpty)(x_1.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(y_2.UnsafePtr).V0).IntVal) != (0) {

var go__go_3_1_103 gopurs_runtime.Value
go__go_3_1_103 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
var v2_6_loop bool = (v2_6_loop_val.IntVal) != (0)
go__go_3_1_103:
for {
if false { continue go__go_3_1_103 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var v2_6 bool = v2_6_loop
_ = v2_6
var __t3 bool
{
if (v2_6) != (true) {
__t3 = false
goto end_branch_3
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
var __t2 bool
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t2 = v2_6
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil)) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
v2_6_loop = (v2_6) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0).IntVal) != (0))
continue go__go_3_1_103
__t3 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
return gopurs_runtime.Bool(__t3)
}
}()
})
})
})
__t_and_4 = (gopurs_runtime.Apply3(go__go_3_1_103, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(x_1.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(y_2.UnsafePtr).V1))}, gopurs_runtime.Bool(true)).IntVal) != (0)
}
return gopurs_runtime.Bool(__t_and_4)
})
})
})}
_ = eq1NonEmpty1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord1{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(eq1NonEmpty1_0_0)}
}), gopurs_runtime.Func(func(dictOrd_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_4_5 -> gopurs_runtime.Value
v_4_5 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_1, "compare"), (*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(y_3.UnsafePtr).V0)
_ = v_4_5
var __t11 uint32
{
if (uint32(v_4_5.IntVal) == 1527465420) {
__t11 = 1527465420
goto end_branch_11
} else {

}
}
{
if (uint32(v_4_5.IntVal) == 380165415) {
__t11 = 380165415
goto end_branch_11
} else {

}
}
{
var go__go_5_6_104 gopurs_runtime.Value
go__go_5_6_104 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_6_104:
for {
if false { continue go__go_5_6_104 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t10 uint32
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
var __t7 uint32
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr == nil) {
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
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr == nil) {
__t10 = 380165415
goto end_branch_10
} else {

}
}
{
if ((v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil)) && ((v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr != nil)) {
// TAST (Let): v2_8_8 -> gopurs_runtime.Value
v2_8_8 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_1, "compare"), (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V0)
_ = v2_8_8
var __t9 uint32
{
if (uint32(v2_8_8.IntVal) == 902936544) {
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V1)}
continue go__go_5_6_104
__t9 = uint32(gopurs_runtime.Value{}.IntVal)
goto end_branch_9
} else {

}
}
{
__t9 = uint32(v2_8_8.IntVal)
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
__t11 = uint32(gopurs_runtime.Apply2(go__go_5_6_104, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(y_3.UnsafePtr).V1))}).IntVal)
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t11), UnsafePtr: nil}
})
})
})})}
}()
	})
	return cache_Data_List_Types_ord1NonEmptyList
}

var cache_Data_List_Types_ordList gopurs_runtime.Value
var once_Data_List_Types_ordList sync.Once
func Get_Data_List_Types_ordList() gopurs_runtime.Value {
	once_Data_List_Types_ordList.Do(func() {
		cache_Data_List_Types_ordList = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Types_ordList(dictOrd_0_box)
})
	})
	return cache_Data_List_Types_ordList
}

var cache_Data_List_Types_ordNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_ordNonEmptyList sync.Once
func Get_Data_List_Types_ordNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_ordNonEmptyList.Do(func() {
		cache_Data_List_Types_ordNonEmptyList = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Types_ordNonEmptyList(dictOrd_0_box)
})
	})
	return cache_Data_List_Types_ordNonEmptyList
}

var cache_Data_List_Types_comonadNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_comonadNonEmptyList sync.Once
func Get_Data_List_Types_comonadNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_comonadNonEmptyList.Do(func() {
		cache_Data_List_Types_comonadNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(&Constructor_Control_Comonad_Comonad{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend](Get_Data_List_Types_extendNonEmptyList()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V0
})})}
	})
	return cache_Data_List_Types_comonadNonEmptyList
}

var cache_Data_List_Types_applyList gopurs_runtime.Value
var once_Data_List_Types_applyList sync.Once
func Get_Data_List_Types_applyList() gopurs_runtime.Value {
	once_Data_List_Types_applyList.Do(func() {
		cache_Data_List_Types_applyList = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Types_functorList()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_List_Types_Cons
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr == nil) {
__t4 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_4
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr != nil) {
var go__go_2_0_109 gopurs_runtime.Value
go__go_2_0_109 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_0_109:
for {
if false { continue go__go_2_0_109 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t1 = b_3
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_3)})}
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_0_109
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
})
var go__go_3_2_110 gopurs_runtime.Value
go__go_3_2_110 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_2_110:
for {
if false { continue go__go_3_2_110 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t3 = v_4
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_2_110
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_2_0_109, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_List_Types_applyList()).V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_1))})))}, gopurs_runtime.Apply2(go__go_3_2_110, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(Call_Data_List_Types_listMap((*Constructor_Data_List_Types_Cons)(v_0.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_1))})))})))
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t4)}
})
})})}
	})
	return cache_Data_List_Types_applyList
}

var cache_Data_List_Types_applyNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_applyNonEmptyList sync.Once
func Get_Data_List_Types_applyNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_applyNonEmptyList.Do(func() {
		cache_Data_List_Types_applyNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Types_functorNonEmptyList()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_111 gopurs_runtime.Value
go__go_2_0_111 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_0_111:
for {
if false { continue go__go_2_0_111 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t1 = b_3
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_3)})}
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_0_111
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
})
// TAST (Let): __local_var_3_2 -> *Constructor_Data_List_Types_Cons
var __local_var_3_2 *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1)})})
var go__go_4_3_112 gopurs_runtime.Value
go__go_4_3_112 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_3_112:
for {
if false { continue go__go_4_3_112 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t4 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t4 = b_5
goto end_branch_4
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_5)})}
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_3_112
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
var go__go_5_5_113 gopurs_runtime.Value
go__go_5_5_113 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_6_loop_val)
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_5_113:
for {
if false { continue go__go_5_5_113 }
var v_6 *Constructor_Data_List_Types_Cons = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t6 *Constructor_Data_List_Types_Cons
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr == nil) {
__t6 = v_6
goto end_branch_6
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V0, v_6})})
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V1)}
continue go__go_5_5_113
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t6)}
}
}()
})
})
var go__go_3_7_114 gopurs_runtime.Value
go__go_3_7_114 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_7_114:
for {
if false { continue go__go_3_7_114 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t8 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t8 = v_4
goto end_branch_8
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_7_114
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t8)}
}
}()
})
})
var __t15 *Constructor_Data_List_Types_Cons
{
var __t_tag_9 *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1)
if (__t_tag_9 == nil) {
__t15 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_15
} else {

}
}
{
var __t_tag_10 *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1)
if (__t_tag_10 != nil) {
var go__go_4_11_115 gopurs_runtime.Value
go__go_4_11_115 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_11_115:
for {
if false { continue go__go_4_11_115 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t12 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t12 = b_5
goto end_branch_12
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_5)})}
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_11_115
__t12 = gopurs_runtime.Value{}
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
}
}()
})
})
var go__go_5_13_116 gopurs_runtime.Value
go__go_5_13_116 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_6_loop_val)
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_13_116:
for {
if false { continue go__go_5_13_116 }
var v_6 *Constructor_Data_List_Types_Cons = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t14 *Constructor_Data_List_Types_Cons
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr == nil) {
__t14 = v_6
goto end_branch_14
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V0, v_6})})
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V1)}
continue go__go_5_13_116
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
__t15 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_11_115, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_List_Types_applyList()).V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(v1_1.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(nil)})})))}, gopurs_runtime.Apply2(go__go_5_13_116, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(Call_Data_List_Types_listMap((*Constructor_Data_List_Types_Cons)((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(v1_1.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(nil)})})))})))
goto end_branch_15
} else {

}
}
{
__t15 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_15:
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(v1_1.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_2_0_111, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_3_112, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_List_Types_applyList()).V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((__local_var_3_2).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v1_1.UnsafePtr).V1))})))}, gopurs_runtime.Apply2(go__go_5_5_113, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(Call_Data_List_Types_listMap((__local_var_3_2).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v1_1.UnsafePtr).V1))})))}))))}, gopurs_runtime.Apply2(go__go_3_7_114, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t15)}))))}})}
})
})})}
	})
	return cache_Data_List_Types_applyNonEmptyList
}

var cache_Data_List_Types_bindList gopurs_runtime.Value
var once_Data_List_Types_bindList sync.Once
func Get_Data_List_Types_bindList() gopurs_runtime.Value {
	once_Data_List_Types_bindList.Do(func() {
		cache_Data_List_Types_bindList = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_List_Types_applyList()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_List_Types_Cons
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr == nil) {
__t4 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_4
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr != nil) {
var go__go_2_0_117 gopurs_runtime.Value
go__go_2_0_117 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_0_117:
for {
if false { continue go__go_2_0_117 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t1 = b_3
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_3)})}
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_0_117
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
})
var go__go_3_2_118 gopurs_runtime.Value
go__go_3_2_118 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_2_118:
for {
if false { continue go__go_3_2_118 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t3 = v_4
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_2_118
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_2_0_117, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Types_bindList()).V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_0.UnsafePtr).V1)}, v1_1)))}, gopurs_runtime.Apply2(go__go_3_2_118, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(v1_1, (*Constructor_Data_List_Types_Cons)(v_0.UnsafePtr).V0)))})))
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t4)}
})
})})}
	})
	return cache_Data_List_Types_bindList
}

var cache_Data_List_Types_bindNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_bindNonEmptyList sync.Once
func Get_Data_List_Types_bindNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_bindNonEmptyList.Do(func() {
		cache_Data_List_Types_bindNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_List_Types_applyNonEmptyList()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_2_0 -> *Constructor_Data_NonEmpty_NonEmpty
v1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(f_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V0))
_ = v1_2_0
var go__go_3_1_119 gopurs_runtime.Value
go__go_3_1_119 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_1_119:
for {
if false { continue go__go_3_1_119 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t2 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t2 = b_4
goto end_branch_2
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_4)})}
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_1_119
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}()
})
})
var __t11 *Constructor_Data_List_Types_Cons
{
var __t_tag_3 *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1)
if (__t_tag_3 == nil) {
__t11 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_11
} else {

}
}
{
var __t_tag_4 *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1)
if (__t_tag_4 != nil) {
var go__go_4_5_120 gopurs_runtime.Value
go__go_4_5_120 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_5_120:
for {
if false { continue go__go_4_5_120 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t6 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t6 = b_5
goto end_branch_6
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_5)})}
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_5_120
__t6 = gopurs_runtime.Value{}
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
}()
})
})
var go__go_5_8_121 gopurs_runtime.Value
go__go_5_8_121 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_6_loop_val)
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_8_121:
for {
if false { continue go__go_5_8_121 }
var v_6 *Constructor_Data_List_Types_Cons = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t9 *Constructor_Data_List_Types_Cons
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr == nil) {
__t9 = v_6
goto end_branch_9
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V0, v_6})})
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V1)}
continue go__go_5_8_121
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t9)}
}
}()
})
})
// TAST (Let): __local_var_6_10 -> gopurs_runtime.Value
__local_var_6_10 := gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Types_Cons)((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1.UnsafePtr).V0)
_ = __local_var_6_10
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_5_120, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Types_bindList()).V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1.UnsafePtr).V1)}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.Apply(f_1, x_5)
_ = __local_var_6_7
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(__local_var_6_7.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(__local_var_6_7.UnsafePtr).V1)})}
}))))}, gopurs_runtime.Apply2(go__go_5_8_121, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(__local_var_6_10.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(__local_var_6_10.UnsafePtr).V1)})})))
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_11:
var go__go_4_12_122 gopurs_runtime.Value
go__go_4_12_122 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_12_122:
for {
if false { continue go__go_4_12_122 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t13 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t13 = v_5
goto end_branch_13
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_12_122
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
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (v1_2_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_3_1_119, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t11)}, gopurs_runtime.Apply2(go__go_4_12_122, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v1_2_0).V1))}))))}})}
})
})})}
	})
	return cache_Data_List_Types_bindNonEmptyList
}

var cache_Data_List_Types_applicativeList gopurs_runtime.Value
var once_Data_List_Types_applicativeList sync.Once
func Get_Data_List_Types_applicativeList() gopurs_runtime.Value {
	once_Data_List_Types_applicativeList.Do(func() {
		cache_Data_List_Types_applicativeList = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_List_Types_applyList()))}
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, a_0, (*Constructor_Data_List_Types_Cons)(nil)})}
})})}
	})
	return cache_Data_List_Types_applicativeList
}

var cache_Data_List_Types_monadList gopurs_runtime.Value
var once_Data_List_Types_monadList sync.Once
func Get_Data_List_Types_monadList() gopurs_runtime.Value {
	once_Data_List_Types_monadList.Do(func() {
		cache_Data_List_Types_monadList = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_List_Types_applicativeList()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Types_bindList()))}
})})}
	})
	return cache_Data_List_Types_monadList
}

var cache_Data_List_Types_altNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_altNonEmptyList sync.Once
func Get_Data_List_Types_altNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_altNonEmptyList.Do(func() {
		cache_Data_List_Types_altNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(&Constructor_Control_Alt_Alt{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Types_functorNonEmptyList()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(as_prime_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_123 gopurs_runtime.Value
go__go_2_0_123 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_0_123:
for {
if false { continue go__go_2_0_123 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t1 = b_3
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_3)})}
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_0_123
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
})
var go__go_3_2_124 gopurs_runtime.Value
go__go_3_2_124 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_2_124:
for {
if false { continue go__go_3_2_124 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t3 = v_4
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_2_124
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_2_0_123, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(as_prime_1.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(as_prime_1.UnsafePtr).V1)})}, gopurs_runtime.Apply2(go__go_3_2_124, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1))}))))}})}
})
})})}
	})
	return cache_Data_List_Types_altNonEmptyList
}

var cache_Data_List_Types_altList gopurs_runtime.Value
var once_Data_List_Types_altList sync.Once
func Get_Data_List_Types_altList() gopurs_runtime.Value {
	once_Data_List_Types_altList.Do(func() {
		cache_Data_List_Types_altList = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(&Constructor_Control_Alt_Alt{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Types_functorList()))}
}), gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_125 gopurs_runtime.Value
go__go_2_0_125 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_0_125:
for {
if false { continue go__go_2_0_125 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t1 = b_3
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_3)})}
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_0_125
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
})
var go__go_3_2_126 gopurs_runtime.Value
go__go_3_2_126 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_2_126:
for {
if false { continue go__go_3_2_126 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t3 = v_4
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_2_126
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_2_0_125, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_1))}, gopurs_runtime.Apply2(go__go_3_2_126, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_0))}))))}
})
})})}
	})
	return cache_Data_List_Types_altList
}

var cache_Data_List_Types_plusList gopurs_runtime.Value
var once_Data_List_Types_plusList sync.Once
func Get_Data_List_Types_plusList() gopurs_runtime.Value {
	once_Data_List_Types_plusList.Do(func() {
		cache_Data_List_Types_plusList = gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(&Constructor_Control_Plus_Plus{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt](Get_Data_List_Types_altList()))}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})}
	})
	return cache_Data_List_Types_plusList
}

var cache_Data_List_Types_alternativeList gopurs_runtime.Value
var once_Data_List_Types_alternativeList sync.Once
func Get_Data_List_Types_alternativeList() gopurs_runtime.Value {
	once_Data_List_Types_alternativeList.Do(func() {
		cache_Data_List_Types_alternativeList = gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(&Constructor_Control_Alternative_Alternative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_List_Types_applicativeList()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus](Get_Data_List_Types_plusList()))}
})})}
	})
	return cache_Data_List_Types_alternativeList
}

var cache_Data_List_Types_monadPlusList gopurs_runtime.Value
var once_Data_List_Types_monadPlusList sync.Once
func Get_Data_List_Types_monadPlusList() gopurs_runtime.Value {
	once_Data_List_Types_monadPlusList.Do(func() {
		cache_Data_List_Types_monadPlusList = gopurs_runtime.Value{Type: 9, IntVal: 3236234573, UnsafePtr: unsafe.Pointer(&Constructor_Control_MonadPlus_MonadPlus{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative](Get_Data_List_Types_alternativeList()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](Get_Data_List_Types_monadList()))}
})})}
	})
	return cache_Data_List_Types_monadPlusList
}

var cache_Data_List_Types_applicativeNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_applicativeNonEmptyList sync.Once
func Get_Data_List_Types_applicativeNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_applicativeNonEmptyList.Do(func() {
		cache_Data_List_Types_applicativeNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_List_Types_applyNonEmptyList()))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, x_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})}
})})}
	})
	return cache_Data_List_Types_applicativeNonEmptyList
}

var cache_Data_List_Types_pure gopurs_runtime.Value
var once_Data_List_Types_pure sync.Once
func Get_Data_List_Types_pure() gopurs_runtime.Value {
	once_Data_List_Types_pure.Do(func() {
		cache_Data_List_Types_pure = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_Types_pure(x_0_box))}
})
	})
	return cache_Data_List_Types_pure
}

var cache_Data_List_Types_monadNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_monadNonEmptyList sync.Once
func Get_Data_List_Types_monadNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_monadNonEmptyList.Do(func() {
		cache_Data_List_Types_monadNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_List_Types_applicativeNonEmptyList()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Types_bindNonEmptyList()))}
})})}
	})
	return cache_Data_List_Types_monadNonEmptyList
}

var cache_Data_List_Types_traversable1NonEmptyList gopurs_runtime.Value
var once_Data_List_Types_traversable1NonEmptyList sync.Once
func Get_Data_List_Types_traversable1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_traversable1NonEmptyList.Do(func() {
		cache_Data_List_Types_traversable1NonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 1596088409, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Traversable_Traversable1{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldableNonEmpty1_1_0 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_1_0 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_2_1 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_5_2 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_5_2
var go__go_6_3_127 gopurs_runtime.Value
go__go_6_3_127 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_7_loop gopurs_runtime.Value = b_7_loop_val
var v_8_loop gopurs_runtime.Value = v_8_loop_val
go__go_6_3_127:
for {
if false { continue go__go_6_3_127 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var v_8 gopurs_runtime.Value = v_8_loop
_ = v_8
var __t4 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 1358893437 && v_8.UnsafePtr == nil) {
__t4 = b_7
goto end_branch_4
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 1358893437 && v_8.UnsafePtr != nil) {
b_7_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_2.V0), b_7, gopurs_runtime.Apply(f_3, (*Constructor_Data_List_Types_Cons)(v_8.UnsafePtr).V0))
v_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_8.UnsafePtr).V1)}
continue go__go_6_3_127
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
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), gopurs_runtime.Apply(f_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0), gopurs_runtime.Apply2(go__go_6_3_127, gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_5_128 gopurs_runtime.Value
go__go_4_5_128 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_5_128:
for {
if false { continue go__go_4_5_128 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t6 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t6 = b_5
goto end_branch_6
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Apply2(f_1, b_5, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0)
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_5_128
__t6 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Apply2(go__go_4_5_128, gopurs_runtime.Apply2(f_1, b_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_7_129 gopurs_runtime.Value
go__go_4_7_129 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_7_129:
for {
if false { continue go__go_4_7_129 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t8 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t8 = b_5
goto end_branch_8
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Apply2(f_1, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0, b_5)
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_7_129
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
var go__go_5_9_130 gopurs_runtime.Value
go__go_5_9_130 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_6_loop_val)
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_9_130:
for {
if false { continue go__go_5_9_130 }
var v_6 *Constructor_Data_List_Types_Cons = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t10 *Constructor_Data_List_Types_Cons
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr == nil) {
__t10 = v_6
goto end_branch_10
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V0, v_6})})
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V1)}
continue go__go_5_9_130
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t10)}
}
}()
})
})
return gopurs_runtime.Apply2(f_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_4_7_129, b_2, gopurs_runtime.Apply2(go__go_5_9_130, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1)))
})
})
})}
_ = foldableNonEmpty1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Foldable_Foldable1{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_1_0)}
}), gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_11_131 gopurs_runtime.Value
go__go_5_11_131 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_11_131:
for {
if false { continue go__go_5_11_131 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t12 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr == nil) {
__t12 = b_6
goto end_branch_12
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr != nil) {
b_6_loop = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_2, "append"), b_6, gopurs_runtime.Apply(f_3, (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0))
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_11_131
__t12 = gopurs_runtime.Value{}
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
}
}()
})
})
return gopurs_runtime.Apply2(go__go_5_11_131, gopurs_runtime.Apply(f_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_13_132 gopurs_runtime.Value
go__go_4_13_132 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_13_132:
for {
if false { continue go__go_4_13_132 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t14 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t14 = b_5
goto end_branch_14
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Apply2(f_2, b_5, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0)
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_13_132
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
return gopurs_runtime.Apply2(go__go_4_13_132, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1)
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_15 -> gopurs_runtime.Value
__local_var_4_15 := gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0)
_ = __local_var_4_15
var go__go_5_17_133 gopurs_runtime.Value
go__go_5_17_133 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_17_133:
for {
if false { continue go__go_5_17_133 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t20 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr == nil) {
__t20 = b_6
goto end_branch_20
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr != nil) {
// TAST (Let): __local_var_8_18 -> gopurs_runtime.Value
__local_var_8_18 := gopurs_runtime.Apply(f_2, (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0)
_ = __local_var_8_18
var __t19 gopurs_runtime.Value
{
if (b_6.Type == 9 && b_6.IntVal == 930809136 && b_6.UnsafePtr == nil) {
__t19 = (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0
goto end_branch_19
} else {

}
}
{
if (b_6.Type == 9 && b_6.IntVal == 930809136 && b_6.UnsafePtr != nil) {
__t19 = gopurs_runtime.Apply(__local_var_8_18, (*Constructor_Data_Maybe_Just)(b_6.UnsafePtr).V0)
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, __t19})}
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_17_133
__t20 = gopurs_runtime.Value{}
goto end_branch_20
} else {

}
}
{
__t20 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_20:
return __t20
}
}()
})
})
var go__go_6_21_134 gopurs_runtime.Value
go__go_6_21_134 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_7_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_7_loop_val)
var v1_8_loop gopurs_runtime.Value = v1_8_loop_val
go__go_6_21_134:
for {
if false { continue go__go_6_21_134 }
var v_7 *Constructor_Data_List_Types_Cons = v_7_loop
_ = v_7
var v1_8 gopurs_runtime.Value = v1_8_loop
_ = v1_8
var __t22 *Constructor_Data_List_Types_Cons
{
if (v1_8.Type == 9 && v1_8.IntVal == 1358893437 && v1_8.UnsafePtr == nil) {
__t22 = v_7
goto end_branch_22
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 1358893437 && v1_8.UnsafePtr != nil) {
v_7_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_8.UnsafePtr).V0, v_7})})
v1_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_8.UnsafePtr).V1)}
continue go__go_6_21_134
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_22
} else {

}
}
{
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_22:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t22)}
}
}()
})
})
// TAST (Let): __local_var_5_16 -> *Constructor_Data_Maybe_Just
__local_var_5_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(go__go_5_17_133, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Apply2(go__go_6_21_134, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1)))
_ = __local_var_5_16
var __t23 gopurs_runtime.Value
{
if (__local_var_5_16 == nil) {
__t23 = (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0
goto end_branch_23
} else {

}
}
{
if (__local_var_5_16 != nil) {
__t23 = gopurs_runtime.Apply(__local_var_4_15, (__local_var_5_16).V0)
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
return __t23
})
})})}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorNonEmpty1_1_24 -> *Constructor_Data_Functor_Functor
functorNonEmpty1_1_24 := &Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(f_1, (*Constructor_Data_NonEmpty_NonEmpty)(m_2.UnsafePtr).V0), gopurs_runtime.Apply(Call_Data_List_Types_listMap(f_1), (*Constructor_Data_NonEmpty_NonEmpty)(m_2.UnsafePtr).V1)})}
})
})}
_ = functorNonEmpty1_1_24
// TAST (Let): foldableNonEmpty1_2_25 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_2_25 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_3_26 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_3_26 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_3_26
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_6_27 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_6_27 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_6_27
var go__go_7_28_135 gopurs_runtime.Value
go__go_7_28_135 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_8_loop gopurs_runtime.Value = b_8_loop_val
var v_9_loop gopurs_runtime.Value = v_9_loop_val
go__go_7_28_135:
for {
if false { continue go__go_7_28_135 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var v_9 gopurs_runtime.Value = v_9_loop
_ = v_9
var __t29 gopurs_runtime.Value
{
if (v_9.Type == 9 && v_9.IntVal == 1358893437 && v_9.UnsafePtr == nil) {
__t29 = b_8
goto end_branch_29
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 1358893437 && v_9.UnsafePtr != nil) {
b_8_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_27.V0), b_8, gopurs_runtime.Apply(f_4, (*Constructor_Data_List_Types_Cons)(v_9.UnsafePtr).V0))
v_9_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_9.UnsafePtr).V1)}
continue go__go_7_28_135
__t29 = gopurs_runtime.Value{}
goto end_branch_29
} else {

}
}
{
__t29 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_29:
return __t29
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_26.V0), gopurs_runtime.Apply(f_4, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0), gopurs_runtime.Apply2(go__go_7_28_135, gopurs_runtime.RecordGet(dictMonoid_2, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_30_136 gopurs_runtime.Value
go__go_5_30_136 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_30_136:
for {
if false { continue go__go_5_30_136 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t31 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr == nil) {
__t31 = b_6
goto end_branch_31
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr != nil) {
b_6_loop = gopurs_runtime.Apply2(f_2, b_6, (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0)
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_30_136
__t31 = gopurs_runtime.Value{}
goto end_branch_31
} else {

}
}
{
__t31 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_31:
return __t31
}
}()
})
})
return gopurs_runtime.Apply2(go__go_5_30_136, gopurs_runtime.Apply2(f_2, b_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_32_137 gopurs_runtime.Value
go__go_5_32_137 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_32_137:
for {
if false { continue go__go_5_32_137 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t33 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr == nil) {
__t33 = b_6
goto end_branch_33
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr != nil) {
b_6_loop = gopurs_runtime.Apply2(f_2, (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0, b_6)
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_32_137
__t33 = gopurs_runtime.Value{}
goto end_branch_33
} else {

}
}
{
__t33 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_33:
return __t33
}
}()
})
})
var go__go_6_34_138 gopurs_runtime.Value
go__go_6_34_138 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_7_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_7_loop_val)
var v1_8_loop gopurs_runtime.Value = v1_8_loop_val
go__go_6_34_138:
for {
if false { continue go__go_6_34_138 }
var v_7 *Constructor_Data_List_Types_Cons = v_7_loop
_ = v_7
var v1_8 gopurs_runtime.Value = v1_8_loop
_ = v1_8
var __t35 *Constructor_Data_List_Types_Cons
{
if (v1_8.Type == 9 && v1_8.IntVal == 1358893437 && v1_8.UnsafePtr == nil) {
__t35 = v_7
goto end_branch_35
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 1358893437 && v1_8.UnsafePtr != nil) {
v_7_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_8.UnsafePtr).V0, v_7})})
v1_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_8.UnsafePtr).V1)}
continue go__go_6_34_138
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_35
} else {

}
}
{
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_35:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t35)}
}
}()
})
})
return gopurs_runtime.Apply2(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_5_32_137, b_3, gopurs_runtime.Apply2(go__go_6_34_138, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1)))
})
})
})}
_ = foldableNonEmpty1_2_25
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(&Constructor_Data_Traversable_Traversable{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_2_25)}
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorNonEmpty1_1_24)}
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_4_36 -> *Constructor_Control_Apply_Apply
Apply0_4_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_4_36
// TAST (Let): Functor0_5_37 -> *Constructor_Data_Functor_Functor
Functor0_5_37 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_37
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_7_38 -> *Constructor_Control_Apply_Apply
Apply0_7_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_7_38
var go__go_8_39_139 gopurs_runtime.Value
go__go_8_39_139 = gopurs_runtime.Func(func(b_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_9_loop gopurs_runtime.Value = b_9_loop_val
var v_10_loop gopurs_runtime.Value = v_10_loop_val
go__go_8_39_139:
for {
if false { continue go__go_8_39_139 }
var b_9 gopurs_runtime.Value = b_9_loop
_ = b_9
var v_10 gopurs_runtime.Value = v_10_loop
_ = v_10
var __t40 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 1358893437 && v_10.UnsafePtr == nil) {
__t40 = b_9
goto end_branch_40
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 1358893437 && v_10.UnsafePtr != nil) {
b_9_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_10.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_9)})}
v_10_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_10.UnsafePtr).V1)}
continue go__go_8_39_139
__t40 = gopurs_runtime.Value{}
goto end_branch_40
} else {

}
}
{
__t40 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_40:
return __t40
}
}()
})
})
var go__go_8_41_140 gopurs_runtime.Value
go__go_8_41_140 = gopurs_runtime.Func(func(b_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_9_loop gopurs_runtime.Value = b_9_loop_val
var v_10_loop gopurs_runtime.Value = v_10_loop_val
go__go_8_41_140:
for {
if false { continue go__go_8_41_140 }
var b_9 gopurs_runtime.Value = b_9_loop
_ = b_9
var v_10 gopurs_runtime.Value = v_10_loop
_ = v_10
var __t42 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 1358893437 && v_10.UnsafePtr == nil) {
__t42 = b_9
goto end_branch_42
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 1358893437 && v_10.UnsafePtr != nil) {
b_9_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_7_38.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_7_38.V0), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(b_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, a_12, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_11)})}
})
}), b_9), (*Constructor_Data_List_Types_Cons)(v_10.UnsafePtr).V0)
v_10_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_10.UnsafePtr).V1)}
continue go__go_8_41_140
__t42 = gopurs_runtime.Value{}
goto end_branch_42
} else {

}
}
{
__t42 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_42:
return __t42
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_36.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_37.V0), Get_Data_NonEmpty_NonEmpty(), (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(go__go_8_39_139, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), gopurs_runtime.Apply2(go__go_8_41_140, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V1)))
})
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_4_43 -> *Constructor_Control_Apply_Apply
Apply0_4_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_4_43
// TAST (Let): Functor0_5_44 -> *Constructor_Data_Functor_Functor
Functor0_5_44 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_44
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_8_45 -> *Constructor_Control_Apply_Apply
Apply0_8_45 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_8_45
var go__go_9_46_141 gopurs_runtime.Value
go__go_9_46_141 = gopurs_runtime.Func(func(b_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_10_loop gopurs_runtime.Value = b_10_loop_val
var v_11_loop gopurs_runtime.Value = v_11_loop_val
go__go_9_46_141:
for {
if false { continue go__go_9_46_141 }
var b_10 gopurs_runtime.Value = b_10_loop
_ = b_10
var v_11 gopurs_runtime.Value = v_11_loop
_ = v_11
var __t47 gopurs_runtime.Value
{
if (v_11.Type == 9 && v_11.IntVal == 1358893437 && v_11.UnsafePtr == nil) {
__t47 = b_10
goto end_branch_47
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 1358893437 && v_11.UnsafePtr != nil) {
b_10_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_11.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_10)})}
v_11_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_11.UnsafePtr).V1)}
continue go__go_9_46_141
__t47 = gopurs_runtime.Value{}
goto end_branch_47
} else {

}
}
{
__t47 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_47:
return __t47
}
}()
})
})
var go__go_9_48_142 gopurs_runtime.Value
go__go_9_48_142 = gopurs_runtime.Func(func(b_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_10_loop gopurs_runtime.Value = b_10_loop_val
var v_11_loop gopurs_runtime.Value = v_11_loop_val
go__go_9_48_142:
for {
if false { continue go__go_9_48_142 }
var b_10 gopurs_runtime.Value = b_10_loop
_ = b_10
var v_11 gopurs_runtime.Value = v_11_loop
_ = v_11
var __t49 gopurs_runtime.Value
{
if (v_11.Type == 9 && v_11.IntVal == 1358893437 && v_11.UnsafePtr == nil) {
__t49 = b_10
goto end_branch_49
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 1358893437 && v_11.UnsafePtr != nil) {
b_10_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_8_45.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_8_45.V0), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(b_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, a_13, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_12)})}
})
}), b_10), gopurs_runtime.Apply(f_6, (*Constructor_Data_List_Types_Cons)(v_11.UnsafePtr).V0))
v_11_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_11.UnsafePtr).V1)}
continue go__go_9_48_142
__t49 = gopurs_runtime.Value{}
goto end_branch_49
} else {

}
}
{
__t49 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_49:
return __t49
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_43.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_44.V0), Get_Data_NonEmpty_NonEmpty(), gopurs_runtime.Apply(f_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(go__go_9_46_141, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), gopurs_runtime.Apply2(go__go_9_48_142, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1)))
})
})
})})}
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1](Get_Data_List_Types_traversable1NonEmptyList()).V3), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0))}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_50 -> *Constructor_Data_Functor_Functor
Functor0_1_50 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_50
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_53_144 gopurs_runtime.Value
go__go_4_53_144 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_53_144:
for {
if false { continue go__go_4_53_144 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t54 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t54 = b_5
goto end_branch_54
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, a_8, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(b_7.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(b_7.UnsafePtr).V1)})}})}
})
}), b_5), gopurs_runtime.Apply(f_2, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0))
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_53_144
__t54 = gopurs_runtime.Value{}
goto end_branch_54
} else {

}
}
{
__t54 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_54:
return __t54
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_50.V0), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_51_143 gopurs_runtime.Value
go__go_5_51_143 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_51_143:
for {
if false { continue go__go_5_51_143 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t52 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr == nil) {
__t52 = b_6
goto end_branch_52
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr != nil) {
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(b_6.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(b_6.UnsafePtr).V1)})}})}
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_51_143
__t52 = gopurs_runtime.Value{}
goto end_branch_52
} else {

}
}
{
__t52 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_52:
return __t52
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply2(go__go_5_51_143, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (*Constructor_Data_NonEmpty_NonEmpty)(v1_4.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v1_4.UnsafePtr).V1))})))}
}), gopurs_runtime.Apply2(go__go_4_53_144, gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_50.V0), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, x_5, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})}
}), gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1))}))
})
})
})})}
	})
	return cache_Data_List_Types_traversable1NonEmptyList
}

var cache_Data_List_Types_applicativeNonEmptyList__1156428081 gopurs_runtime.Value
var once_Data_List_Types_applicativeNonEmptyList__1156428081 sync.Once
func Get_Data_List_Types_applicativeNonEmptyList__1156428081() gopurs_runtime.Value {
	once_Data_List_Types_applicativeNonEmptyList__1156428081.Do(func() {
		cache_Data_List_Types_applicativeNonEmptyList__1156428081 = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_List_Types_applyNonEmptyList()))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, x_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})}
})})}
	})
	return cache_Data_List_Types_applicativeNonEmptyList__1156428081
}

var cache_Data_List_Types_applicativeNonEmptyList__3820246605 gopurs_runtime.Value
var once_Data_List_Types_applicativeNonEmptyList__3820246605 sync.Once
func Get_Data_List_Types_applicativeNonEmptyList__3820246605() gopurs_runtime.Value {
	once_Data_List_Types_applicativeNonEmptyList__3820246605.Do(func() {
		cache_Data_List_Types_applicativeNonEmptyList__3820246605 = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_List_Types_applyNonEmptyList()))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, x_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})}
})})}
	})
	return cache_Data_List_Types_applicativeNonEmptyList__3820246605
}

var cache_Data_List_Types_applyNonEmptyList__1066888753 gopurs_runtime.Value
var once_Data_List_Types_applyNonEmptyList__1066888753 sync.Once
func Get_Data_List_Types_applyNonEmptyList__1066888753() gopurs_runtime.Value {
	once_Data_List_Types_applyNonEmptyList__1066888753.Do(func() {
		cache_Data_List_Types_applyNonEmptyList__1066888753 = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Types_functorNonEmptyList()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_145 gopurs_runtime.Value
go__go_2_0_145 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_0_145:
for {
if false { continue go__go_2_0_145 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t1 = b_3
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_3)})}
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_0_145
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
})
// TAST (Let): __local_var_3_2 -> *Constructor_Data_List_Types_Cons
var __local_var_3_2 *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1)})})
var go__go_4_3_146 gopurs_runtime.Value
go__go_4_3_146 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_3_146:
for {
if false { continue go__go_4_3_146 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t4 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t4 = b_5
goto end_branch_4
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_5)})}
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_3_146
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
var go__go_5_5_147 gopurs_runtime.Value
go__go_5_5_147 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_6_loop_val)
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_5_147:
for {
if false { continue go__go_5_5_147 }
var v_6 *Constructor_Data_List_Types_Cons = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t6 *Constructor_Data_List_Types_Cons
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr == nil) {
__t6 = v_6
goto end_branch_6
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V0, v_6})})
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V1)}
continue go__go_5_5_147
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t6)}
}
}()
})
})
var go__go_3_7_148 gopurs_runtime.Value
go__go_3_7_148 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_7_148:
for {
if false { continue go__go_3_7_148 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t8 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t8 = v_4
goto end_branch_8
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_7_148
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t8)}
}
}()
})
})
var __t15 *Constructor_Data_List_Types_Cons
{
var __t_tag_9 *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1)
if (__t_tag_9 == nil) {
__t15 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_15
} else {

}
}
{
var __t_tag_10 *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1)
if (__t_tag_10 != nil) {
var go__go_4_11_149 gopurs_runtime.Value
go__go_4_11_149 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_11_149:
for {
if false { continue go__go_4_11_149 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t12 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t12 = b_5
goto end_branch_12
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_5)})}
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_11_149
__t12 = gopurs_runtime.Value{}
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
}
}()
})
})
var go__go_5_13_150 gopurs_runtime.Value
go__go_5_13_150 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_6_loop_val)
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_13_150:
for {
if false { continue go__go_5_13_150 }
var v_6 *Constructor_Data_List_Types_Cons = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t14 *Constructor_Data_List_Types_Cons
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr == nil) {
__t14 = v_6
goto end_branch_14
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V0, v_6})})
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V1)}
continue go__go_5_13_150
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
__t15 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_11_149, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_List_Types_applyList()).V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(v1_1.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(nil)})})))}, gopurs_runtime.Apply2(go__go_5_13_150, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(Call_Data_List_Types_listMap((*Constructor_Data_List_Types_Cons)((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(v1_1.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(nil)})})))})))
goto end_branch_15
} else {

}
}
{
__t15 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_15:
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(v1_1.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_2_0_145, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_3_146, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_List_Types_applyList()).V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((__local_var_3_2).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v1_1.UnsafePtr).V1))})))}, gopurs_runtime.Apply2(go__go_5_5_147, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(Call_Data_List_Types_listMap((__local_var_3_2).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v1_1.UnsafePtr).V1))})))}))))}, gopurs_runtime.Apply2(go__go_3_7_148, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t15)}))))}})}
})
})})}
	})
	return cache_Data_List_Types_applyNonEmptyList__1066888753
}

var cache_Data_List_Types_applyNonEmptyList__538673485 gopurs_runtime.Value
var once_Data_List_Types_applyNonEmptyList__538673485 sync.Once
func Get_Data_List_Types_applyNonEmptyList__538673485() gopurs_runtime.Value {
	once_Data_List_Types_applyNonEmptyList__538673485.Do(func() {
		cache_Data_List_Types_applyNonEmptyList__538673485 = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Types_functorNonEmptyList()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_151 gopurs_runtime.Value
go__go_2_0_151 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_0_151:
for {
if false { continue go__go_2_0_151 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t1 = b_3
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_3)})}
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_0_151
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
})
// TAST (Let): __local_var_3_2 -> *Constructor_Data_List_Types_Cons
var __local_var_3_2 *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1)})})
var go__go_4_3_152 gopurs_runtime.Value
go__go_4_3_152 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_3_152:
for {
if false { continue go__go_4_3_152 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t4 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t4 = b_5
goto end_branch_4
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_5)})}
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_3_152
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
var go__go_5_5_153 gopurs_runtime.Value
go__go_5_5_153 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_6_loop_val)
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_5_153:
for {
if false { continue go__go_5_5_153 }
var v_6 *Constructor_Data_List_Types_Cons = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t6 *Constructor_Data_List_Types_Cons
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr == nil) {
__t6 = v_6
goto end_branch_6
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V0, v_6})})
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V1)}
continue go__go_5_5_153
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t6)}
}
}()
})
})
var go__go_3_7_154 gopurs_runtime.Value
go__go_3_7_154 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_7_154:
for {
if false { continue go__go_3_7_154 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t8 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t8 = v_4
goto end_branch_8
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_7_154
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t8)}
}
}()
})
})
var __t15 *Constructor_Data_List_Types_Cons
{
var __t_tag_9 *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1)
if (__t_tag_9 == nil) {
__t15 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_15
} else {

}
}
{
var __t_tag_10 *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1)
if (__t_tag_10 != nil) {
var go__go_4_11_155 gopurs_runtime.Value
go__go_4_11_155 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_11_155:
for {
if false { continue go__go_4_11_155 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t12 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t12 = b_5
goto end_branch_12
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_5)})}
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_11_155
__t12 = gopurs_runtime.Value{}
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
}
}()
})
})
var go__go_5_13_156 gopurs_runtime.Value
go__go_5_13_156 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_6_loop_val)
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_13_156:
for {
if false { continue go__go_5_13_156 }
var v_6 *Constructor_Data_List_Types_Cons = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t14 *Constructor_Data_List_Types_Cons
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr == nil) {
__t14 = v_6
goto end_branch_14
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V0, v_6})})
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V1)}
continue go__go_5_13_156
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
__t15 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_11_155, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_List_Types_applyList()).V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v1_1.UnsafePtr).V0))}, (*Constructor_Data_List_Types_Cons)(nil)})})))}, gopurs_runtime.Apply2(go__go_5_13_156, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(Call_Data_List_Types_listMap((*Constructor_Data_List_Types_Cons)((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v1_1.UnsafePtr).V0))}, (*Constructor_Data_List_Types_Cons)(nil)})})))})))
goto end_branch_15
} else {

}
}
{
__t15 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_15:
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v1_1.UnsafePtr).V0))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_2_0_151, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_3_152, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_List_Types_applyList()).V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((__local_var_3_2).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v1_1.UnsafePtr).V1))})))}, gopurs_runtime.Apply2(go__go_5_5_153, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(Call_Data_List_Types_listMap((__local_var_3_2).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v1_1.UnsafePtr).V1))})))}))))}, gopurs_runtime.Apply2(go__go_3_7_154, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t15)}))))}})}
})
})})}
	})
	return cache_Data_List_Types_applyNonEmptyList__538673485
}

var cache_Data_List_Types_bindNonEmptyList__1408886065 gopurs_runtime.Value
var once_Data_List_Types_bindNonEmptyList__1408886065 sync.Once
func Get_Data_List_Types_bindNonEmptyList__1408886065() gopurs_runtime.Value {
	once_Data_List_Types_bindNonEmptyList__1408886065.Do(func() {
		cache_Data_List_Types_bindNonEmptyList__1408886065 = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Data_List_Types_applyNonEmptyList()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_2_0 -> *Constructor_Data_NonEmpty_NonEmpty
v1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply(f_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V0))
_ = v1_2_0
var go__go_3_1_157 gopurs_runtime.Value
go__go_3_1_157 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_1_157:
for {
if false { continue go__go_3_1_157 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t2 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t2 = b_4
goto end_branch_2
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_4)})}
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_1_157
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}()
})
})
var __t11 *Constructor_Data_List_Types_Cons
{
var __t_tag_3 *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1)
if (__t_tag_3 == nil) {
__t11 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_11
} else {

}
}
{
var __t_tag_4 *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1)
if (__t_tag_4 != nil) {
var go__go_4_5_158 gopurs_runtime.Value
go__go_4_5_158 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_5_158:
for {
if false { continue go__go_4_5_158 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t6 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t6 = b_5
goto end_branch_6
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_5)})}
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_5_158
__t6 = gopurs_runtime.Value{}
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
}()
})
})
var go__go_5_8_159 gopurs_runtime.Value
go__go_5_8_159 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_6_loop_val)
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_8_159:
for {
if false { continue go__go_5_8_159 }
var v_6 *Constructor_Data_List_Types_Cons = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t9 *Constructor_Data_List_Types_Cons
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr == nil) {
__t9 = v_6
goto end_branch_9
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V0, v_6})})
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V1)}
continue go__go_5_8_159
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t9)}
}
}()
})
})
// TAST (Let): __local_var_6_10 -> gopurs_runtime.Value
__local_var_6_10 := gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Types_Cons)((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1.UnsafePtr).V0)
_ = __local_var_6_10
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_5_158, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Types_bindList()).V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)((*Constructor_Data_NonEmpty_NonEmpty)(v_0.UnsafePtr).V1.UnsafePtr).V1)}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.Apply(f_1, x_5)
_ = __local_var_6_7
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(__local_var_6_7.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(__local_var_6_7.UnsafePtr).V1)})}
}))))}, gopurs_runtime.Apply2(go__go_5_8_159, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(__local_var_6_10.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(__local_var_6_10.UnsafePtr).V1)})})))
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_11:
var go__go_4_12_160 gopurs_runtime.Value
go__go_4_12_160 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_12_160:
for {
if false { continue go__go_4_12_160 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t13 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t13 = v_5
goto end_branch_13
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_12_160
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
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (v1_2_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_3_1_157, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t11)}, gopurs_runtime.Apply2(go__go_4_12_160, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v1_2_0).V1))}))))}})}
})
})})}
	})
	return cache_Data_List_Types_bindNonEmptyList__1408886065
}

var cache_Data_List_Types_extendNonEmptyList__688763217 gopurs_runtime.Value
var once_Data_List_Types_extendNonEmptyList__688763217 sync.Once
func Get_Data_List_Types_extendNonEmptyList__688763217() gopurs_runtime.Value {
	once_Data_List_Types_extendNonEmptyList__688763217.Do(func() {
		cache_Data_List_Types_extendNonEmptyList__688763217 = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(&Constructor_Control_Extend_Extend{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Types_functorNonEmptyList()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_161 gopurs_runtime.Value
go__go_2_0_161 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_0_161:
for {
if false { continue go__go_2_0_161 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t1 = b_3
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
b_3_loop = gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(b_3, "acc"))})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(b_3, "acc")))}})}), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(b_3, "val"))})})
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_0_161
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
})
var go__go_3_2_162 gopurs_runtime.Value
go__go_3_2_162 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_2_162:
for {
if false { continue go__go_3_2_162 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t3 = v_4
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_2_162
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_1))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(gopurs_runtime.Apply2(go__go_2_0_161, gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), gopurs_runtime.Apply2(go__go_3_2_162, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_1.UnsafePtr).V1))})), "val")))}})}
})
})})}
	})
	return cache_Data_List_Types_extendNonEmptyList__688763217
}

var cache_Data_List_Types_foldable1NonEmptyList__2239557029 gopurs_runtime.Value
var once_Data_List_Types_foldable1NonEmptyList__2239557029 sync.Once
func Get_Data_List_Types_foldable1NonEmptyList__2239557029() gopurs_runtime.Value {
	once_Data_List_Types_foldable1NonEmptyList__2239557029.Do(func() {
		cache_Data_List_Types_foldable1NonEmptyList__2239557029 = func() gopurs_runtime.Value {
// TAST (Let): foldableNonEmpty1_0_0 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_0_0 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_1 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_4_2 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_4_2
var go__go_5_3_163 gopurs_runtime.Value
go__go_5_3_163 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_3_163:
for {
if false { continue go__go_5_3_163 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t4 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr == nil) {
__t4 = b_6
goto end_branch_4
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr != nil) {
b_6_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_4_2.V0), b_6, gopurs_runtime.Apply(f_2, (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0))
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_3_163
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
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_1.V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0), gopurs_runtime.Apply2(go__go_5_3_163, gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_5_164 gopurs_runtime.Value
go__go_3_5_164 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_5_164:
for {
if false { continue go__go_3_5_164 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t6 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t6 = b_4
goto end_branch_6
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Apply2(f_0, b_4, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0)
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_5_164
__t6 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Apply2(go__go_3_5_164, gopurs_runtime.Apply2(f_0, b_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_7_165 gopurs_runtime.Value
go__go_3_7_165 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_7_165:
for {
if false { continue go__go_3_7_165 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t8 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t8 = b_4
goto end_branch_8
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0, b_4)
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_7_165
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
var go__go_4_9_166 gopurs_runtime.Value
go__go_4_9_166 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_9_166:
for {
if false { continue go__go_4_9_166 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t10 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t10 = v_5
goto end_branch_10
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_9_166
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t10)}
}
}()
})
})
return gopurs_runtime.Apply2(f_0, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_3_7_165, b_1, gopurs_runtime.Apply2(go__go_4_9_166, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)))
})
})
})}
_ = foldableNonEmpty1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Foldable_Foldable1{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_0_0)}
}), gopurs_runtime.Func(func(dictSemigroup_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_11_167 gopurs_runtime.Value
go__go_4_11_167 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_11_167:
for {
if false { continue go__go_4_11_167 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t12 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t12 = b_5
goto end_branch_12
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_1, "append"), b_5, gopurs_runtime.Apply(f_2, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0))
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_11_167
__t12 = gopurs_runtime.Value{}
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
}
}()
})
})
return gopurs_runtime.Apply2(go__go_4_11_167, gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_13_168 gopurs_runtime.Value
go__go_3_13_168 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_13_168:
for {
if false { continue go__go_3_13_168 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t14 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t14 = b_4
goto end_branch_14
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Apply2(f_1, b_4, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0)
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_13_168
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
return gopurs_runtime.Apply2(go__go_3_13_168, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_15 -> gopurs_runtime.Value
__local_var_3_15 := gopurs_runtime.Apply(f_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0)
_ = __local_var_3_15
var go__go_4_17_169 gopurs_runtime.Value
go__go_4_17_169 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_17_169:
for {
if false { continue go__go_4_17_169 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t20 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t20 = b_5
goto end_branch_20
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
// TAST (Let): __local_var_7_18 -> gopurs_runtime.Value
__local_var_7_18 := gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0)
_ = __local_var_7_18
var __t19 gopurs_runtime.Value
{
if (b_5.Type == 9 && b_5.IntVal == 930809136 && b_5.UnsafePtr == nil) {
__t19 = (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0
goto end_branch_19
} else {

}
}
{
if (b_5.Type == 9 && b_5.IntVal == 930809136 && b_5.UnsafePtr != nil) {
__t19 = gopurs_runtime.Apply(__local_var_7_18, (*Constructor_Data_Maybe_Just)(b_5.UnsafePtr).V0)
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, __t19})}
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_17_169
__t20 = gopurs_runtime.Value{}
goto end_branch_20
} else {

}
}
{
__t20 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_20:
return __t20
}
}()
})
})
var go__go_5_21_170 gopurs_runtime.Value
go__go_5_21_170 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_6_loop_val)
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_21_170:
for {
if false { continue go__go_5_21_170 }
var v_6 *Constructor_Data_List_Types_Cons = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t22 *Constructor_Data_List_Types_Cons
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr == nil) {
__t22 = v_6
goto end_branch_22
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V0, v_6})})
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V1)}
continue go__go_5_21_170
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_22
} else {

}
}
{
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_22:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t22)}
}
}()
})
})
// TAST (Let): __local_var_4_16 -> *Constructor_Data_Maybe_Just
__local_var_4_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(go__go_4_17_169, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Apply2(go__go_5_21_170, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)))
_ = __local_var_4_16
var __t23 gopurs_runtime.Value
{
if (__local_var_4_16 == nil) {
__t23 = (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0
goto end_branch_23
} else {

}
}
{
if (__local_var_4_16 != nil) {
__t23 = gopurs_runtime.Apply(__local_var_3_15, (__local_var_4_16).V0)
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
return __t23
})
})})}
}()
	})
	return cache_Data_List_Types_foldable1NonEmptyList__2239557029
}

var cache_Data_List_Types_foldable1NonEmptyList__2630190169 gopurs_runtime.Value
var once_Data_List_Types_foldable1NonEmptyList__2630190169 sync.Once
func Get_Data_List_Types_foldable1NonEmptyList__2630190169() gopurs_runtime.Value {
	once_Data_List_Types_foldable1NonEmptyList__2630190169.Do(func() {
		cache_Data_List_Types_foldable1NonEmptyList__2630190169 = func() gopurs_runtime.Value {
// TAST (Let): foldableNonEmpty1_0_0 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_0_0 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_1 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_4_2 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_4_2
var go__go_5_3_171 gopurs_runtime.Value
go__go_5_3_171 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_3_171:
for {
if false { continue go__go_5_3_171 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t4 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr == nil) {
__t4 = b_6
goto end_branch_4
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr != nil) {
b_6_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_4_2.V0), b_6, gopurs_runtime.Apply(f_2, (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0))
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_3_171
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
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_1.V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0), gopurs_runtime.Apply2(go__go_5_3_171, gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_5_172 gopurs_runtime.Value
go__go_3_5_172 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_5_172:
for {
if false { continue go__go_3_5_172 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t6 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t6 = b_4
goto end_branch_6
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Apply2(f_0, b_4, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0)
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_5_172
__t6 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Apply2(go__go_3_5_172, gopurs_runtime.Apply2(f_0, b_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_7_173 gopurs_runtime.Value
go__go_3_7_173 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_7_173:
for {
if false { continue go__go_3_7_173 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t8 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t8 = b_4
goto end_branch_8
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0, b_4)
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_7_173
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
var go__go_4_9_174 gopurs_runtime.Value
go__go_4_9_174 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_9_174:
for {
if false { continue go__go_4_9_174 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t10 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t10 = v_5
goto end_branch_10
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_9_174
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t10)}
}
}()
})
})
return gopurs_runtime.Apply2(f_0, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_3_7_173, b_1, gopurs_runtime.Apply2(go__go_4_9_174, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)))
})
})
})}
_ = foldableNonEmpty1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Foldable_Foldable1{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_0_0)}
}), gopurs_runtime.Func(func(dictSemigroup_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_11_175 gopurs_runtime.Value
go__go_4_11_175 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_11_175:
for {
if false { continue go__go_4_11_175 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t12 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t12 = b_5
goto end_branch_12
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_1, "append"), b_5, gopurs_runtime.Apply(f_2, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0))
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_11_175
__t12 = gopurs_runtime.Value{}
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
}
}()
})
})
return gopurs_runtime.Apply2(go__go_4_11_175, gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_13_176 gopurs_runtime.Value
go__go_3_13_176 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_13_176:
for {
if false { continue go__go_3_13_176 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t14 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t14 = b_4
goto end_branch_14
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Apply2(f_1, b_4, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0)
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_13_176
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
return gopurs_runtime.Apply2(go__go_3_13_176, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_15 -> gopurs_runtime.Value
__local_var_3_15 := gopurs_runtime.Apply(f_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0)
_ = __local_var_3_15
var go__go_4_17_177 gopurs_runtime.Value
go__go_4_17_177 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_17_177:
for {
if false { continue go__go_4_17_177 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t20 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t20 = b_5
goto end_branch_20
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
// TAST (Let): __local_var_7_18 -> gopurs_runtime.Value
__local_var_7_18 := gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0)
_ = __local_var_7_18
var __t19 gopurs_runtime.Value
{
if (b_5.Type == 9 && b_5.IntVal == 930809136 && b_5.UnsafePtr == nil) {
__t19 = (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0
goto end_branch_19
} else {

}
}
{
if (b_5.Type == 9 && b_5.IntVal == 930809136 && b_5.UnsafePtr != nil) {
__t19 = gopurs_runtime.Apply(__local_var_7_18, (*Constructor_Data_Maybe_Just)(b_5.UnsafePtr).V0)
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, __t19})}
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_17_177
__t20 = gopurs_runtime.Value{}
goto end_branch_20
} else {

}
}
{
__t20 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_20:
return __t20
}
}()
})
})
var go__go_5_21_178 gopurs_runtime.Value
go__go_5_21_178 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_6_loop_val)
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_21_178:
for {
if false { continue go__go_5_21_178 }
var v_6 *Constructor_Data_List_Types_Cons = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t22 *Constructor_Data_List_Types_Cons
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr == nil) {
__t22 = v_6
goto end_branch_22
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V0, v_6})})
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V1)}
continue go__go_5_21_178
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_22
} else {

}
}
{
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_22:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t22)}
}
}()
})
})
// TAST (Let): __local_var_4_16 -> *Constructor_Data_Maybe_Just
__local_var_4_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(go__go_4_17_177, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Apply2(go__go_5_21_178, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)))
_ = __local_var_4_16
var __t23 gopurs_runtime.Value
{
if (__local_var_4_16 == nil) {
__t23 = (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0
goto end_branch_23
} else {

}
}
{
if (__local_var_4_16 != nil) {
__t23 = gopurs_runtime.Apply(__local_var_3_15, (__local_var_4_16).V0)
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
return __t23
})
})})}
}()
	})
	return cache_Data_List_Types_foldable1NonEmptyList__2630190169
}

var cache_Data_List_Types_foldableNonEmptyList__1181575686 gopurs_runtime.Value
var once_Data_List_Types_foldableNonEmptyList__1181575686 sync.Once
func Get_Data_List_Types_foldableNonEmptyList__1181575686() gopurs_runtime.Value {
	once_Data_List_Types_foldableNonEmptyList__1181575686.Do(func() {
		cache_Data_List_Types_foldableNonEmptyList__1181575686 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(&Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_4_1 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_4_1
var go__go_5_2_179 gopurs_runtime.Value
go__go_5_2_179 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_2_179:
for {
if false { continue go__go_5_2_179 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t3 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr == nil) {
__t3 = b_6
goto end_branch_3
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr != nil) {
b_6_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_4_1.V0), b_6, gopurs_runtime.Apply(f_2, (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0))
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_2_179
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
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0), gopurs_runtime.Apply2(go__go_5_2_179, gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_4_180 gopurs_runtime.Value
go__go_3_4_180 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_4_180:
for {
if false { continue go__go_3_4_180 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t5 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t5 = b_4
goto end_branch_5
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Apply2(f_0, b_4, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0)
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_4_180
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}()
})
})
return gopurs_runtime.Apply2(go__go_3_4_180, gopurs_runtime.Apply2(f_0, b_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_6_181 gopurs_runtime.Value
go__go_3_6_181 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_6_181:
for {
if false { continue go__go_3_6_181 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t7 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t7 = b_4
goto end_branch_7
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Apply2(f_0, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0, b_4)
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_6_181
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
var go__go_4_8_182 gopurs_runtime.Value
go__go_4_8_182 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_8_182:
for {
if false { continue go__go_4_8_182 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t9 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t9 = v_5
goto end_branch_9
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_8_182
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t9)}
}
}()
})
})
return gopurs_runtime.Apply2(f_0, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_3_6_181, b_1, gopurs_runtime.Apply2(go__go_4_8_182, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1)))
})
})
})})}
	})
	return cache_Data_List_Types_foldableNonEmptyList__1181575686
}

var cache_Data_List_Types_foldableWithIndexNonEmptyList__909823402 gopurs_runtime.Value
var once_Data_List_Types_foldableWithIndexNonEmptyList__909823402 sync.Once
func Get_Data_List_Types_foldableWithIndexNonEmptyList__909823402() gopurs_runtime.Value {
	once_Data_List_Types_foldableWithIndexNonEmptyList__909823402.Do(func() {
		cache_Data_List_Types_foldableWithIndexNonEmptyList__909823402 = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(&Constructor_Data_FoldableWithIndex_FoldableWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Types_foldableNonEmptyList()))}
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_3_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_3_0
var go__go_4_1_183 gopurs_runtime.Value
go__go_4_1_183 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_1_183:
for {
if false { continue go__go_4_1_183 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t2 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t2 = b_5
goto end_branch_2
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_0.V0), (*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V1, gopurs_runtime.Apply2(f_1, gopurs_runtime.Int((1) + ((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal)), (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0))})}
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_1_183
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply2(f_1, gopurs_runtime.Int(0), (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_4_1_183, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})}, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1).UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_3_184 gopurs_runtime.Value
go__go_3_3_184 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_3_184:
for {
if false { continue go__go_3_3_184 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t4 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t4 = b_4
goto end_branch_4
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((1) + ((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V0.IntVal)), (*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V1, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0)})}
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_3_184
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
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_3_3_184, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(0), b_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0)})}, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1).UnsafePtr).V1
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_6_185 gopurs_runtime.Value
go__go_3_6_185 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_6_185:
for {
if false { continue go__go_3_6_185 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t7 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t7 = b_4
goto end_branch_7
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V1)})}})}
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_6_185
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
// TAST (Let): v_3_5 -> *Constructor_Data_Tuple_Tuple
v_3_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(go__go_3_6_185, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1))}))
_ = v_3_5
var go__go_4_8_186 gopurs_runtime.Value
go__go_4_8_186 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_8_186:
for {
if false { continue go__go_4_8_186 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t9 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t9 = b_5
goto end_branch_9
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) - (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((1) + (((*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V0.IntVal) - (1))), (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(b_5.UnsafePtr).V1)})}
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_8_186
__t9 = gopurs_runtime.Value{}
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
}
}()
})
})
return gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(0), (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_4_8_186, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int((v_3_5).V0.IntVal), b_1})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_3_5).V1))}).UnsafePtr).V1)
})
})
})})}
	})
	return cache_Data_List_Types_foldableWithIndexNonEmptyList__909823402
}

var cache_Data_List_Types_functorNonEmptyList__2834508934 gopurs_runtime.Value
var once_Data_List_Types_functorNonEmptyList__2834508934 sync.Once
func Get_Data_List_Types_functorNonEmptyList__2834508934() gopurs_runtime.Value {
	once_Data_List_Types_functorNonEmptyList__2834508934.Do(func() {
		cache_Data_List_Types_functorNonEmptyList__2834508934 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_NonEmpty_NonEmpty)(m_1.UnsafePtr).V0), gopurs_runtime.Apply(Call_Data_List_Types_listMap(f_0), (*Constructor_Data_NonEmpty_NonEmpty)(m_1.UnsafePtr).V1)})}
})
})})}
	})
	return cache_Data_List_Types_functorNonEmptyList__2834508934
}

var cache_Data_List_Types_functorNonEmptyList__1593940346 gopurs_runtime.Value
var once_Data_List_Types_functorNonEmptyList__1593940346 sync.Once
func Get_Data_List_Types_functorNonEmptyList__1593940346() gopurs_runtime.Value {
	once_Data_List_Types_functorNonEmptyList__1593940346.Do(func() {
		cache_Data_List_Types_functorNonEmptyList__1593940346 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_NonEmpty_NonEmpty)(m_1.UnsafePtr).V0), gopurs_runtime.Apply(Call_Data_List_Types_listMap(f_0), (*Constructor_Data_NonEmpty_NonEmpty)(m_1.UnsafePtr).V1)})}
})
})})}
	})
	return cache_Data_List_Types_functorNonEmptyList__1593940346
}

var cache_Data_List_Types_functorWithIndexNonEmptyList__1542383978 gopurs_runtime.Value
var once_Data_List_Types_functorWithIndexNonEmptyList__1542383978 sync.Once
func Get_Data_List_Types_functorWithIndexNonEmptyList__1542383978() gopurs_runtime.Value {
	once_Data_List_Types_functorWithIndexNonEmptyList__1542383978.Do(func() {
		cache_Data_List_Types_functorWithIndexNonEmptyList__1542383978 = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(&Constructor_Data_FunctorWithIndex_FunctorWithIndex{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Types_functorNonEmptyList()))}
}), gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_1_187 gopurs_runtime.Value
go__go_2_1_187 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_1_187:
for {
if false { continue go__go_2_1_187 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t2 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t2 = b_3
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_3.UnsafePtr).V0.IntVal) + (1)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_Tuple_Tuple)(b_3.UnsafePtr).V1)})}})}
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_1_187
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}()
})
})
// TAST (Let): v_2_0 -> *Constructor_Data_Tuple_Tuple
v_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(go__go_2_1_187, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_1.UnsafePtr).V1))}))
_ = v_2_0
var go__go_3_3_188 gopurs_runtime.Value
go__go_3_3_188 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_3_188:
for {
if false { continue go__go_3_3_188 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t4 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t4 = b_4
goto end_branch_4
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V0.IntVal) - (1)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Apply2(fn_0, gopurs_runtime.Int((1) + (((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V0.IntVal) - (1))), (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_Tuple_Tuple)(b_4.UnsafePtr).V1)})}})}
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_3_188
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
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply2(fn_0, gopurs_runtime.Int(0), (*Constructor_Data_NonEmpty_NonEmpty)(v_1.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply2(go__go_3_3_188, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int((v_2_0).V0.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_2_0).V1))}).UnsafePtr).V1})}
})
})})}
	})
	return cache_Data_List_Types_functorWithIndexNonEmptyList__1542383978
}

var cache_Data_List_Types_listMap__858544730 gopurs_runtime.Value
var once_Data_List_Types_listMap__858544730 sync.Once
func Get_Data_List_Types_listMap__858544730() gopurs_runtime.Value {
	once_Data_List_Types_listMap__858544730.Do(func() {
		cache_Data_List_Types_listMap__858544730 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Types_listMap__858544730(f_0_box)
})
	})
	return cache_Data_List_Types_listMap__858544730
}

var cache_Data_List_Types_listMap__4135416762 gopurs_runtime.Value
var once_Data_List_Types_listMap__4135416762 sync.Once
func Get_Data_List_Types_listMap__4135416762() gopurs_runtime.Value {
	once_Data_List_Types_listMap__4135416762.Do(func() {
		cache_Data_List_Types_listMap__4135416762 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Types_listMap__4135416762(f_0_box)
})
	})
	return cache_Data_List_Types_listMap__4135416762
}

var cache_Data_List_Types_nelCons__195558898 gopurs_runtime.Value
var once_Data_List_Types_nelCons__195558898 sync.Once
func Get_Data_List_Types_nelCons__195558898() gopurs_runtime.Value {
	once_Data_List_Types_nelCons__195558898.Do(func() {
		cache_Data_List_Types_nelCons__195558898 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_Types_nelCons__195558898(a_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_1_box)))}
})
	})
	return cache_Data_List_Types_nelCons__195558898
}

var cache_Data_List_Types_nelCons__2148523118 gopurs_runtime.Value
var once_Data_List_Types_nelCons__2148523118 sync.Once
func Get_Data_List_Types_nelCons__2148523118() gopurs_runtime.Value {
	once_Data_List_Types_nelCons__2148523118.Do(func() {
		cache_Data_List_Types_nelCons__2148523118 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_Types_nelCons__2148523118(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](a_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_1_box)))}
})
	})
	return cache_Data_List_Types_nelCons__2148523118
}

var cache_Data_List_Types_semigroupList__2766094215 gopurs_runtime.Value
var once_Data_List_Types_semigroupList__2766094215 sync.Once
func Get_Data_List_Types_semigroupList__2766094215() gopurs_runtime.Value {
	once_Data_List_Types_semigroupList__2766094215.Do(func() {
		cache_Data_List_Types_semigroupList__2766094215 = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_193 gopurs_runtime.Value
go__go_2_0_193 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_0_193:
for {
if false { continue go__go_2_0_193 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t1 = b_3
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_3)})}
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_0_193
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
})
var go__go_3_2_194 gopurs_runtime.Value
go__go_3_2_194 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_2_194:
for {
if false { continue go__go_3_2_194 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t3 = v_4
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_2_194
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_2_0_193, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_1))}, gopurs_runtime.Apply2(go__go_3_2_194, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_0))}))))}
})
})})}
	})
	return cache_Data_List_Types_semigroupList__2766094215
}

var cache_Data_List_Types_semigroupList__3527039931 gopurs_runtime.Value
var once_Data_List_Types_semigroupList__3527039931 sync.Once
func Get_Data_List_Types_semigroupList__3527039931() gopurs_runtime.Value {
	once_Data_List_Types_semigroupList__3527039931.Do(func() {
		cache_Data_List_Types_semigroupList__3527039931 = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_195 gopurs_runtime.Value
go__go_2_0_195 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_2_0_195:
for {
if false { continue go__go_2_0_195 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t1 = b_3
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_3)})}
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
continue go__go_2_0_195
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
})
var go__go_3_2_196 gopurs_runtime.Value
go__go_3_2_196 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_2_196:
for {
if false { continue go__go_3_2_196 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t3 = v_4
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_2_196
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_2_0_195, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_1))}, gopurs_runtime.Apply2(go__go_3_2_196, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_0))}))))}
})
})})}
	})
	return cache_Data_List_Types_semigroupList__3527039931
}

var cache_Data_List_Types_toList__2859885498 gopurs_runtime.Value
var once_Data_List_Types_toList__2859885498 sync.Once
func Get_Data_List_Types_toList__2859885498() gopurs_runtime.Value {
	once_Data_List_Types_toList__2859885498.Do(func() {
		cache_Data_List_Types_toList__2859885498 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_Types_toList__2859885498(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_0_box)))}
})
	})
	return cache_Data_List_Types_toList__2859885498
}

var cache_Data_List_Types_toList__1324737658 gopurs_runtime.Value
var once_Data_List_Types_toList__1324737658 sync.Once
func Get_Data_List_Types_toList__1324737658() gopurs_runtime.Value {
	once_Data_List_Types_toList__1324737658.Do(func() {
		cache_Data_List_Types_toList__1324737658 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_Types_toList__1324737658(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_0_box)))}
})
	})
	return cache_Data_List_Types_toList__1324737658
}

var cache_Data_List_Types_toList__2402503393 gopurs_runtime.Value
var once_Data_List_Types_toList__2402503393 sync.Once
func Get_Data_List_Types_toList__2402503393() gopurs_runtime.Value {
	once_Data_List_Types_toList__2402503393.Do(func() {
		cache_Data_List_Types_toList__2402503393 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_Types_toList__2402503393(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_0_box)))}
})
	})
	return cache_Data_List_Types_toList__2402503393
}

var cache_Data_List_Types_traversable1NonEmptyList__1171985061 gopurs_runtime.Value
var once_Data_List_Types_traversable1NonEmptyList__1171985061 sync.Once
func Get_Data_List_Types_traversable1NonEmptyList__1171985061() gopurs_runtime.Value {
	once_Data_List_Types_traversable1NonEmptyList__1171985061.Do(func() {
		cache_Data_List_Types_traversable1NonEmptyList__1171985061 = gopurs_runtime.Value{Type: 9, IntVal: 1596088409, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Traversable_Traversable1{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldableNonEmpty1_1_0 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_1_0 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_2_1 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_5_2 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_5_2
var go__go_6_3_197 gopurs_runtime.Value
go__go_6_3_197 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_7_loop gopurs_runtime.Value = b_7_loop_val
var v_8_loop gopurs_runtime.Value = v_8_loop_val
go__go_6_3_197:
for {
if false { continue go__go_6_3_197 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var v_8 gopurs_runtime.Value = v_8_loop
_ = v_8
var __t4 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 1358893437 && v_8.UnsafePtr == nil) {
__t4 = b_7
goto end_branch_4
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 1358893437 && v_8.UnsafePtr != nil) {
b_7_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_2.V0), b_7, gopurs_runtime.Apply(f_3, (*Constructor_Data_List_Types_Cons)(v_8.UnsafePtr).V0))
v_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_8.UnsafePtr).V1)}
continue go__go_6_3_197
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
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), gopurs_runtime.Apply(f_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0), gopurs_runtime.Apply2(go__go_6_3_197, gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_5_198 gopurs_runtime.Value
go__go_4_5_198 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_5_198:
for {
if false { continue go__go_4_5_198 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t6 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t6 = b_5
goto end_branch_6
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Apply2(f_1, b_5, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0)
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_5_198
__t6 = gopurs_runtime.Value{}
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
}()
})
})
return gopurs_runtime.Apply2(go__go_4_5_198, gopurs_runtime.Apply2(f_1, b_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_7_199 gopurs_runtime.Value
go__go_4_7_199 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_7_199:
for {
if false { continue go__go_4_7_199 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t8 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t8 = b_5
goto end_branch_8
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Apply2(f_1, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0, b_5)
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_7_199
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
var go__go_5_9_200 gopurs_runtime.Value
go__go_5_9_200 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_6_loop_val)
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_9_200:
for {
if false { continue go__go_5_9_200 }
var v_6 *Constructor_Data_List_Types_Cons = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t10 *Constructor_Data_List_Types_Cons
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr == nil) {
__t10 = v_6
goto end_branch_10
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V0, v_6})})
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V1)}
continue go__go_5_9_200
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t10)}
}
}()
})
})
return gopurs_runtime.Apply2(f_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_4_7_199, b_2, gopurs_runtime.Apply2(go__go_5_9_200, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1)))
})
})
})}
_ = foldableNonEmpty1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Foldable_Foldable1{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_1_0)}
}), gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_11_201 gopurs_runtime.Value
go__go_5_11_201 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_11_201:
for {
if false { continue go__go_5_11_201 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t12 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr == nil) {
__t12 = b_6
goto end_branch_12
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr != nil) {
b_6_loop = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_2, "append"), b_6, gopurs_runtime.Apply(f_3, (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0))
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_11_201
__t12 = gopurs_runtime.Value{}
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
}
}()
})
})
return gopurs_runtime.Apply2(go__go_5_11_201, gopurs_runtime.Apply(f_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_13_202 gopurs_runtime.Value
go__go_4_13_202 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_13_202:
for {
if false { continue go__go_4_13_202 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t14 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t14 = b_5
goto end_branch_14
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Apply2(f_2, b_5, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0)
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_13_202
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
return gopurs_runtime.Apply2(go__go_4_13_202, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1)
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_15 -> gopurs_runtime.Value
__local_var_4_15 := gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0)
_ = __local_var_4_15
var go__go_5_17_203 gopurs_runtime.Value
go__go_5_17_203 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_17_203:
for {
if false { continue go__go_5_17_203 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t20 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr == nil) {
__t20 = b_6
goto end_branch_20
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr != nil) {
// TAST (Let): __local_var_8_18 -> gopurs_runtime.Value
__local_var_8_18 := gopurs_runtime.Apply(f_2, (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0)
_ = __local_var_8_18
var __t19 gopurs_runtime.Value
{
if (b_6.Type == 9 && b_6.IntVal == 930809136 && b_6.UnsafePtr == nil) {
__t19 = (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0
goto end_branch_19
} else {

}
}
{
if (b_6.Type == 9 && b_6.IntVal == 930809136 && b_6.UnsafePtr != nil) {
__t19 = gopurs_runtime.Apply(__local_var_8_18, (*Constructor_Data_Maybe_Just)(b_6.UnsafePtr).V0)
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, __t19})}
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_17_203
__t20 = gopurs_runtime.Value{}
goto end_branch_20
} else {

}
}
{
__t20 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_20:
return __t20
}
}()
})
})
var go__go_6_21_204 gopurs_runtime.Value
go__go_6_21_204 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_7_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_7_loop_val)
var v1_8_loop gopurs_runtime.Value = v1_8_loop_val
go__go_6_21_204:
for {
if false { continue go__go_6_21_204 }
var v_7 *Constructor_Data_List_Types_Cons = v_7_loop
_ = v_7
var v1_8 gopurs_runtime.Value = v1_8_loop
_ = v1_8
var __t22 *Constructor_Data_List_Types_Cons
{
if (v1_8.Type == 9 && v1_8.IntVal == 1358893437 && v1_8.UnsafePtr == nil) {
__t22 = v_7
goto end_branch_22
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 1358893437 && v1_8.UnsafePtr != nil) {
v_7_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_8.UnsafePtr).V0, v_7})})
v1_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_8.UnsafePtr).V1)}
continue go__go_6_21_204
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_22
} else {

}
}
{
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_22:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t22)}
}
}()
})
})
// TAST (Let): __local_var_5_16 -> *Constructor_Data_Maybe_Just
__local_var_5_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(go__go_5_17_203, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Apply2(go__go_6_21_204, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1)))
_ = __local_var_5_16
var __t23 gopurs_runtime.Value
{
if (__local_var_5_16 == nil) {
__t23 = (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0
goto end_branch_23
} else {

}
}
{
if (__local_var_5_16 != nil) {
__t23 = gopurs_runtime.Apply(__local_var_4_15, (__local_var_5_16).V0)
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
return __t23
})
})})}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorNonEmpty1_1_24 -> *Constructor_Data_Functor_Functor
functorNonEmpty1_1_24 := &Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(f_1, (*Constructor_Data_NonEmpty_NonEmpty)(m_2.UnsafePtr).V0), gopurs_runtime.Apply(Call_Data_List_Types_listMap(f_1), (*Constructor_Data_NonEmpty_NonEmpty)(m_2.UnsafePtr).V1)})}
})
})}
_ = functorNonEmpty1_1_24
// TAST (Let): foldableNonEmpty1_2_25 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_2_25 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_3_26 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_3_26 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_3_26
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_6_27 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_6_27 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_6_27
var go__go_7_28_205 gopurs_runtime.Value
go__go_7_28_205 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_8_loop gopurs_runtime.Value = b_8_loop_val
var v_9_loop gopurs_runtime.Value = v_9_loop_val
go__go_7_28_205:
for {
if false { continue go__go_7_28_205 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var v_9 gopurs_runtime.Value = v_9_loop
_ = v_9
var __t29 gopurs_runtime.Value
{
if (v_9.Type == 9 && v_9.IntVal == 1358893437 && v_9.UnsafePtr == nil) {
__t29 = b_8
goto end_branch_29
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 1358893437 && v_9.UnsafePtr != nil) {
b_8_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_27.V0), b_8, gopurs_runtime.Apply(f_4, (*Constructor_Data_List_Types_Cons)(v_9.UnsafePtr).V0))
v_9_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_9.UnsafePtr).V1)}
continue go__go_7_28_205
__t29 = gopurs_runtime.Value{}
goto end_branch_29
} else {

}
}
{
__t29 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_29:
return __t29
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_26.V0), gopurs_runtime.Apply(f_4, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0), gopurs_runtime.Apply2(go__go_7_28_205, gopurs_runtime.RecordGet(dictMonoid_2, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_30_206 gopurs_runtime.Value
go__go_5_30_206 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_30_206:
for {
if false { continue go__go_5_30_206 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t31 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr == nil) {
__t31 = b_6
goto end_branch_31
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr != nil) {
b_6_loop = gopurs_runtime.Apply2(f_2, b_6, (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0)
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_30_206
__t31 = gopurs_runtime.Value{}
goto end_branch_31
} else {

}
}
{
__t31 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_31:
return __t31
}
}()
})
})
return gopurs_runtime.Apply2(go__go_5_30_206, gopurs_runtime.Apply2(f_2, b_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_32_207 gopurs_runtime.Value
go__go_5_32_207 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_32_207:
for {
if false { continue go__go_5_32_207 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t33 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr == nil) {
__t33 = b_6
goto end_branch_33
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr != nil) {
b_6_loop = gopurs_runtime.Apply2(f_2, (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0, b_6)
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_32_207
__t33 = gopurs_runtime.Value{}
goto end_branch_33
} else {

}
}
{
__t33 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_33:
return __t33
}
}()
})
})
var go__go_6_34_208 gopurs_runtime.Value
go__go_6_34_208 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_7_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_7_loop_val)
var v1_8_loop gopurs_runtime.Value = v1_8_loop_val
go__go_6_34_208:
for {
if false { continue go__go_6_34_208 }
var v_7 *Constructor_Data_List_Types_Cons = v_7_loop
_ = v_7
var v1_8 gopurs_runtime.Value = v1_8_loop
_ = v1_8
var __t35 *Constructor_Data_List_Types_Cons
{
if (v1_8.Type == 9 && v1_8.IntVal == 1358893437 && v1_8.UnsafePtr == nil) {
__t35 = v_7
goto end_branch_35
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 1358893437 && v1_8.UnsafePtr != nil) {
v_7_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_8.UnsafePtr).V0, v_7})})
v1_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_8.UnsafePtr).V1)}
continue go__go_6_34_208
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_35
} else {

}
}
{
__t35 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_35:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t35)}
}
}()
})
})
return gopurs_runtime.Apply2(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_5_32_207, b_3, gopurs_runtime.Apply2(go__go_6_34_208, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1)))
})
})
})}
_ = foldableNonEmpty1_2_25
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(&Constructor_Data_Traversable_Traversable{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_2_25)}
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorNonEmpty1_1_24)}
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_4_36 -> *Constructor_Control_Apply_Apply
Apply0_4_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_4_36
// TAST (Let): Functor0_5_37 -> *Constructor_Data_Functor_Functor
Functor0_5_37 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_37
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_7_38 -> *Constructor_Control_Apply_Apply
Apply0_7_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_7_38
var go__go_8_39_209 gopurs_runtime.Value
go__go_8_39_209 = gopurs_runtime.Func(func(b_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_9_loop gopurs_runtime.Value = b_9_loop_val
var v_10_loop gopurs_runtime.Value = v_10_loop_val
go__go_8_39_209:
for {
if false { continue go__go_8_39_209 }
var b_9 gopurs_runtime.Value = b_9_loop
_ = b_9
var v_10 gopurs_runtime.Value = v_10_loop
_ = v_10
var __t40 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 1358893437 && v_10.UnsafePtr == nil) {
__t40 = b_9
goto end_branch_40
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 1358893437 && v_10.UnsafePtr != nil) {
b_9_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_10.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_9)})}
v_10_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_10.UnsafePtr).V1)}
continue go__go_8_39_209
__t40 = gopurs_runtime.Value{}
goto end_branch_40
} else {

}
}
{
__t40 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_40:
return __t40
}
}()
})
})
var go__go_8_41_210 gopurs_runtime.Value
go__go_8_41_210 = gopurs_runtime.Func(func(b_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_9_loop gopurs_runtime.Value = b_9_loop_val
var v_10_loop gopurs_runtime.Value = v_10_loop_val
go__go_8_41_210:
for {
if false { continue go__go_8_41_210 }
var b_9 gopurs_runtime.Value = b_9_loop
_ = b_9
var v_10 gopurs_runtime.Value = v_10_loop
_ = v_10
var __t42 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 1358893437 && v_10.UnsafePtr == nil) {
__t42 = b_9
goto end_branch_42
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 1358893437 && v_10.UnsafePtr != nil) {
b_9_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_7_38.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_7_38.V0), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(b_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, a_12, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_11)})}
})
}), b_9), (*Constructor_Data_List_Types_Cons)(v_10.UnsafePtr).V0)
v_10_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_10.UnsafePtr).V1)}
continue go__go_8_41_210
__t42 = gopurs_runtime.Value{}
goto end_branch_42
} else {

}
}
{
__t42 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_42:
return __t42
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_36.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_37.V0), Get_Data_NonEmpty_NonEmpty(), (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(go__go_8_39_209, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), gopurs_runtime.Apply2(go__go_8_41_210, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V1)))
})
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_4_43 -> *Constructor_Control_Apply_Apply
Apply0_4_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_4_43
// TAST (Let): Functor0_5_44 -> *Constructor_Data_Functor_Functor
Functor0_5_44 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_44
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_8_45 -> *Constructor_Control_Apply_Apply
Apply0_8_45 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_8_45
var go__go_9_46_211 gopurs_runtime.Value
go__go_9_46_211 = gopurs_runtime.Func(func(b_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_10_loop gopurs_runtime.Value = b_10_loop_val
var v_11_loop gopurs_runtime.Value = v_11_loop_val
go__go_9_46_211:
for {
if false { continue go__go_9_46_211 }
var b_10 gopurs_runtime.Value = b_10_loop
_ = b_10
var v_11 gopurs_runtime.Value = v_11_loop
_ = v_11
var __t47 gopurs_runtime.Value
{
if (v_11.Type == 9 && v_11.IntVal == 1358893437 && v_11.UnsafePtr == nil) {
__t47 = b_10
goto end_branch_47
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 1358893437 && v_11.UnsafePtr != nil) {
b_10_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_11.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_10)})}
v_11_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_11.UnsafePtr).V1)}
continue go__go_9_46_211
__t47 = gopurs_runtime.Value{}
goto end_branch_47
} else {

}
}
{
__t47 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_47:
return __t47
}
}()
})
})
var go__go_9_48_212 gopurs_runtime.Value
go__go_9_48_212 = gopurs_runtime.Func(func(b_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_10_loop gopurs_runtime.Value = b_10_loop_val
var v_11_loop gopurs_runtime.Value = v_11_loop_val
go__go_9_48_212:
for {
if false { continue go__go_9_48_212 }
var b_10 gopurs_runtime.Value = b_10_loop
_ = b_10
var v_11 gopurs_runtime.Value = v_11_loop
_ = v_11
var __t49 gopurs_runtime.Value
{
if (v_11.Type == 9 && v_11.IntVal == 1358893437 && v_11.UnsafePtr == nil) {
__t49 = b_10
goto end_branch_49
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 1358893437 && v_11.UnsafePtr != nil) {
b_10_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_8_45.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_8_45.V0), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(b_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, a_13, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_12)})}
})
}), b_10), gopurs_runtime.Apply(f_6, (*Constructor_Data_List_Types_Cons)(v_11.UnsafePtr).V0))
v_11_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_11.UnsafePtr).V1)}
continue go__go_9_48_212
__t49 = gopurs_runtime.Value{}
goto end_branch_49
} else {

}
}
{
__t49 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_49:
return __t49
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_43.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_44.V0), Get_Data_NonEmpty_NonEmpty(), gopurs_runtime.Apply(f_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(go__go_9_46_211, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), gopurs_runtime.Apply2(go__go_9_48_212, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1)))
})
})
})})}
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_50 -> *Constructor_Data_Functor_Functor
Functor0_1_50 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_50
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_53_214 gopurs_runtime.Value
go__go_3_53_214 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_53_214:
for {
if false { continue go__go_3_53_214 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t54 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t54 = b_4
goto end_branch_54
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, a_7, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(b_6.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(b_6.UnsafePtr).V1)})}})}
})
}), b_4), (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0)
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_53_214
__t54 = gopurs_runtime.Value{}
goto end_branch_54
} else {

}
}
{
__t54 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_54:
return __t54
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_50.V0), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_51_213 gopurs_runtime.Value
go__go_4_51_213 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_51_213:
for {
if false { continue go__go_4_51_213 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t52 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t52 = b_5
goto end_branch_52
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(b_5.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(b_5.UnsafePtr).V1)})}})}
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_51_213
__t52 = gopurs_runtime.Value{}
goto end_branch_52
} else {

}
}
{
__t52 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_52:
return __t52
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply2(go__go_4_51_213, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (*Constructor_Data_NonEmpty_NonEmpty)(v1_3.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v1_3.UnsafePtr).V1))})))}
}), gopurs_runtime.Apply2(go__go_3_53_214, gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_50.V0), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, x_4, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})}
}), (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1))}))
})
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_55 -> *Constructor_Data_Functor_Functor
Functor0_1_55 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_55
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_58_216 gopurs_runtime.Value
go__go_4_58_216 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_58_216:
for {
if false { continue go__go_4_58_216 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t59 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t59 = b_5
goto end_branch_59
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, a_8, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(b_7.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(b_7.UnsafePtr).V1)})}})}
})
}), b_5), gopurs_runtime.Apply(f_2, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0))
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_58_216
__t59 = gopurs_runtime.Value{}
goto end_branch_59
} else {

}
}
{
__t59 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_59:
return __t59
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_55.V0), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_56_215 gopurs_runtime.Value
go__go_5_56_215 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_6_loop gopurs_runtime.Value = b_6_loop_val
var v_7_loop gopurs_runtime.Value = v_7_loop_val
go__go_5_56_215:
for {
if false { continue go__go_5_56_215 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t57 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr == nil) {
__t57 = b_6
goto end_branch_57
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 1358893437 && v_7.UnsafePtr != nil) {
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_NonEmpty_NonEmpty)(b_6.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(b_6.UnsafePtr).V1)})}})}
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_7.UnsafePtr).V1)}
continue go__go_5_56_215
__t57 = gopurs_runtime.Value{}
goto end_branch_57
} else {

}
}
{
__t57 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_57:
return __t57
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply2(go__go_5_56_215, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (*Constructor_Data_NonEmpty_NonEmpty)(v1_4.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v1_4.UnsafePtr).V1))})))}
}), gopurs_runtime.Apply2(go__go_4_58_216, gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_55.V0), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, x_5, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})}
}), gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1))}))
})
})
})})}
	})
	return cache_Data_List_Types_traversable1NonEmptyList__1171985061
}

var cache_Data_List_Types_traversableNonEmptyList__2666586758 gopurs_runtime.Value
var once_Data_List_Types_traversableNonEmptyList__2666586758 sync.Once
func Get_Data_List_Types_traversableNonEmptyList__2666586758() gopurs_runtime.Value {
	once_Data_List_Types_traversableNonEmptyList__2666586758.Do(func() {
		cache_Data_List_Types_traversableNonEmptyList__2666586758 = func() gopurs_runtime.Value {
// TAST (Let): functorNonEmpty1_0_0 -> *Constructor_Data_Functor_Functor
functorNonEmpty1_0_0 := &Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_NonEmpty_NonEmpty)(m_1.UnsafePtr).V0), gopurs_runtime.Apply(Call_Data_List_Types_listMap(f_0), (*Constructor_Data_NonEmpty_NonEmpty)(m_1.UnsafePtr).V1)})}
})
})}
_ = functorNonEmpty1_0_0
// TAST (Let): foldableNonEmpty1_1_1 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_1_1 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_2_2 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_2
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_5_3 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_5_3
var go__go_6_4_217 gopurs_runtime.Value
go__go_6_4_217 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_7_loop gopurs_runtime.Value = b_7_loop_val
var v_8_loop gopurs_runtime.Value = v_8_loop_val
go__go_6_4_217:
for {
if false { continue go__go_6_4_217 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var v_8 gopurs_runtime.Value = v_8_loop
_ = v_8
var __t5 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 1358893437 && v_8.UnsafePtr == nil) {
__t5 = b_7
goto end_branch_5
} else {

}
}
{
if (v_8.Type == 9 && v_8.IntVal == 1358893437 && v_8.UnsafePtr != nil) {
b_7_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_3.V0), b_7, gopurs_runtime.Apply(f_3, (*Constructor_Data_List_Types_Cons)(v_8.UnsafePtr).V0))
v_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_8.UnsafePtr).V1)}
continue go__go_6_4_217
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_2.V0), gopurs_runtime.Apply(f_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0), gopurs_runtime.Apply2(go__go_6_4_217, gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_6_218 gopurs_runtime.Value
go__go_4_6_218 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_6_218:
for {
if false { continue go__go_4_6_218 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t7 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t7 = b_5
goto end_branch_7
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Apply2(f_1, b_5, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0)
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_6_218
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
return gopurs_runtime.Apply2(go__go_4_6_218, gopurs_runtime.Apply2(f_1, b_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_8_219 gopurs_runtime.Value
go__go_4_8_219 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_8_219:
for {
if false { continue go__go_4_8_219 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t9 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t9 = b_5
goto end_branch_9
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Apply2(f_1, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0, b_5)
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_8_219
__t9 = gopurs_runtime.Value{}
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
}
}()
})
})
var go__go_5_10_220 gopurs_runtime.Value
go__go_5_10_220 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_6_loop_val)
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_10_220:
for {
if false { continue go__go_5_10_220 }
var v_6 *Constructor_Data_List_Types_Cons = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t11 *Constructor_Data_List_Types_Cons
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr == nil) {
__t11 = v_6
goto end_branch_11
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V0, v_6})})
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V1)}
continue go__go_5_10_220
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t11)}
}
}()
})
})
return gopurs_runtime.Apply2(f_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_4_8_219, b_2, gopurs_runtime.Apply2(go__go_5_10_220, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1)))
})
})
})}
_ = foldableNonEmpty1_1_1
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(&Constructor_Data_Traversable_Traversable{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_1_1)}
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorNonEmpty1_0_0)}
}), gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_3_12 -> *Constructor_Control_Apply_Apply
Apply0_3_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_3_12
// TAST (Let): Functor0_4_13 -> *Constructor_Data_Functor_Functor
Functor0_4_13 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_13
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_6_14 -> *Constructor_Control_Apply_Apply
Apply0_6_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_6_14
var go__go_7_15_221 gopurs_runtime.Value
go__go_7_15_221 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_8_loop gopurs_runtime.Value = b_8_loop_val
var v_9_loop gopurs_runtime.Value = v_9_loop_val
go__go_7_15_221:
for {
if false { continue go__go_7_15_221 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var v_9 gopurs_runtime.Value = v_9_loop
_ = v_9
var __t16 gopurs_runtime.Value
{
if (v_9.Type == 9 && v_9.IntVal == 1358893437 && v_9.UnsafePtr == nil) {
__t16 = b_8
goto end_branch_16
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 1358893437 && v_9.UnsafePtr != nil) {
b_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_9.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_8)})}
v_9_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_9.UnsafePtr).V1)}
continue go__go_7_15_221
__t16 = gopurs_runtime.Value{}
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
return __t16
}
}()
})
})
var go__go_7_17_222 gopurs_runtime.Value
go__go_7_17_222 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_8_loop gopurs_runtime.Value = b_8_loop_val
var v_9_loop gopurs_runtime.Value = v_9_loop_val
go__go_7_17_222:
for {
if false { continue go__go_7_17_222 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var v_9 gopurs_runtime.Value = v_9_loop
_ = v_9
var __t18 gopurs_runtime.Value
{
if (v_9.Type == 9 && v_9.IntVal == 1358893437 && v_9.UnsafePtr == nil) {
__t18 = b_8
goto end_branch_18
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 1358893437 && v_9.UnsafePtr != nil) {
b_8_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_6_14.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_6_14.V0), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(b_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, a_11, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_10)})}
})
}), b_8), (*Constructor_Data_List_Types_Cons)(v_9.UnsafePtr).V0)
v_9_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_9.UnsafePtr).V1)}
continue go__go_7_17_222
__t18 = gopurs_runtime.Value{}
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
return __t18
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_12.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_13.V0), Get_Data_NonEmpty_NonEmpty(), (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(go__go_7_15_221, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), gopurs_runtime.Apply2(go__go_7_17_222, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1)))
})
}), gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_3_19 -> *Constructor_Control_Apply_Apply
Apply0_3_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_3_19
// TAST (Let): Functor0_4_20 -> *Constructor_Data_Functor_Functor
Functor0_4_20 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_20
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_7_21 -> *Constructor_Control_Apply_Apply
Apply0_7_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_7_21
var go__go_8_22_223 gopurs_runtime.Value
go__go_8_22_223 = gopurs_runtime.Func(func(b_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_9_loop gopurs_runtime.Value = b_9_loop_val
var v_10_loop gopurs_runtime.Value = v_10_loop_val
go__go_8_22_223:
for {
if false { continue go__go_8_22_223 }
var b_9 gopurs_runtime.Value = b_9_loop
_ = b_9
var v_10 gopurs_runtime.Value = v_10_loop
_ = v_10
var __t23 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 1358893437 && v_10.UnsafePtr == nil) {
__t23 = b_9
goto end_branch_23
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 1358893437 && v_10.UnsafePtr != nil) {
b_9_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_10.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_9)})}
v_10_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_10.UnsafePtr).V1)}
continue go__go_8_22_223
__t23 = gopurs_runtime.Value{}
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
return __t23
}
}()
})
})
var go__go_8_24_224 gopurs_runtime.Value
go__go_8_24_224 = gopurs_runtime.Func(func(b_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_9_loop gopurs_runtime.Value = b_9_loop_val
var v_10_loop gopurs_runtime.Value = v_10_loop_val
go__go_8_24_224:
for {
if false { continue go__go_8_24_224 }
var b_9 gopurs_runtime.Value = b_9_loop
_ = b_9
var v_10 gopurs_runtime.Value = v_10_loop
_ = v_10
var __t25 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 1358893437 && v_10.UnsafePtr == nil) {
__t25 = b_9
goto end_branch_25
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 1358893437 && v_10.UnsafePtr != nil) {
b_9_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_7_21.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_7_21.V0), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(b_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, a_12, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_11)})}
})
}), b_9), gopurs_runtime.Apply(f_5, (*Constructor_Data_List_Types_Cons)(v_10.UnsafePtr).V0))
v_10_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_10.UnsafePtr).V1)}
continue go__go_8_24_224
__t25 = gopurs_runtime.Value{}
goto end_branch_25
} else {

}
}
{
__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_25:
return __t25
}
}()
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_19.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_20.V0), Get_Data_NonEmpty_NonEmpty(), gopurs_runtime.Apply(f_5, (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(go__go_8_22_223, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), gopurs_runtime.Apply2(go__go_8_24_224, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_2, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V1)))
})
})
})})}
}()
	})
	return cache_Data_List_Types_traversableNonEmptyList__2666586758
}

type Constructor_Data_List_Types_Nil struct {
	Rc uint32
}


type Constructor_Data_List_Types_Cons struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 *Constructor_Data_List_Types_Cons
}


func Call_Data_List_Types_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_List_Types_identity1(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_List_Types_NonEmptyList(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_List_Types_toList(v_0_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_List_Types_Cons {
var v_0 *Constructor_Data_NonEmpty_NonEmpty = v_0_loop
_ = v_0
return &Constructor_Data_List_Types_Cons{1, (v_0).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_0).V1)}
}

func Call_Data_List_Types_nelCons(a_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_NonEmpty_NonEmpty {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_1).V1)})}})})
}

func Call_Data_List_Types_listMap(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var chunkedRevMap_1_0_0 gopurs_runtime.Value
chunkedRevMap_1_0_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2_loop_val)
var v1_3_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_3_loop_val)
chunkedRevMap_1_0_0:
for {
if false { continue chunkedRevMap_1_0_0 }
var v_2 *Constructor_Data_List_Types_Cons = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Types_Cons = v1_3_loop
_ = v1_3
var __t19 *Constructor_Data_List_Types_Cons
{
var __t_and_18 bool = false
if (v1_3 != nil) {

var __t_tag_15 *Constructor_Data_List_Types_Cons = (v1_3).V1
var __t_and_17 bool = false
if (__t_tag_15 != nil) {

var __t_tag_16 *Constructor_Data_List_Types_Cons = ((v1_3).V1).V1
__t_and_17 = (__t_tag_16 != nil)
}
__t_and_18 = __t_and_17
}
if __t_and_18 {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)})})})
v1_3_loop = (((v1_3).V1).V1).V1
continue chunkedRevMap_1_0_0
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_19
} else {

}
}
{
var reverseUnrolledMap_4_1_1 gopurs_runtime.Value
reverseUnrolledMap_4_1_1 = gopurs_runtime.Func(func(v2_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v2_5_loop gopurs_runtime.Value = v2_5_loop_val
var v3_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v3_6_loop_val)
reverseUnrolledMap_4_1_1:
for {
if false { continue reverseUnrolledMap_4_1_1 }
var v2_5 gopurs_runtime.Value = v2_5_loop
_ = v2_5
var v3_6 *Constructor_Data_List_Types_Cons = v3_6_loop
_ = v3_6
var __t8 *Constructor_Data_List_Types_Cons
{
var __t_and_7 bool = false
if (v2_5.Type == 9 && v2_5.IntVal == 1358893437 && v2_5.UnsafePtr != nil) {

var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_List_Types_Cons)(v2_5.UnsafePtr).V0
var __t_and_6 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1358893437 && __t_tag_2.UnsafePtr != nil) {

var __t_tag_3 *Constructor_Data_List_Types_Cons = (*Constructor_Data_List_Types_Cons)((*Constructor_Data_List_Types_Cons)(v2_5.UnsafePtr).V0.UnsafePtr).V1
var __t_and_5 bool = false
if (__t_tag_3 != nil) {

var __t_tag_4 *Constructor_Data_List_Types_Cons = ((*Constructor_Data_List_Types_Cons)((*Constructor_Data_List_Types_Cons)(v2_5.UnsafePtr).V0.UnsafePtr).V1).V1
__t_and_5 = (__t_tag_4 != nil)
}
__t_and_6 = __t_and_5
}
__t_and_7 = __t_and_6
}
if __t_and_7 {
v2_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v2_5.UnsafePtr).V1)}
v3_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons)((*Constructor_Data_List_Types_Cons)(v2_5.UnsafePtr).V0.UnsafePtr).V0), &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Apply(f_0, ((*Constructor_Data_List_Types_Cons)((*Constructor_Data_List_Types_Cons)(v2_5.UnsafePtr).V0.UnsafePtr).V1).V0), &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Apply(f_0, (((*Constructor_Data_List_Types_Cons)((*Constructor_Data_List_Types_Cons)(v2_5.UnsafePtr).V0.UnsafePtr).V1).V1).V0), v3_6}}})})
continue reverseUnrolledMap_4_1_1
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_8
} else {

}
}
{
__t8 = v3_6
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t8)}
}
}()
})
})
var __t14 *Constructor_Data_List_Types_Cons
{
if (v1_3 != nil) {
var __t13 *Constructor_Data_List_Types_Cons
{
var __t_tag_9 *Constructor_Data_List_Types_Cons = (v1_3).V1
if (__t_tag_9 != nil) {
var __t11 *Constructor_Data_List_Types_Cons
{
var __t_tag_10 *Constructor_Data_List_Types_Cons = ((v1_3).V1).V1
if (__t_tag_10 == nil) {
__t11 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Apply(f_0, (v1_3).V0), &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Apply(f_0, ((v1_3).V1).V0), (*Constructor_Data_List_Types_Cons)(nil)}}
goto end_branch_11
} else {

}
}
{
__t11 = (*Constructor_Data_List_Types_Cons)(nil)
}
end_branch_11:
__t13 = __t11
goto end_branch_13
} else {

}
}
{
var __t_tag_12 *Constructor_Data_List_Types_Cons = (v1_3).V1
if (__t_tag_12 == nil) {
__t13 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Apply(f_0, (v1_3).V0), (*Constructor_Data_List_Types_Cons)(nil)}
goto end_branch_13
} else {

}
}
{
__t13 = (*Constructor_Data_List_Types_Cons)(nil)
}
end_branch_13:
__t14 = __t13
goto end_branch_14
} else {

}
}
{
__t14 = (*Constructor_Data_List_Types_Cons)(nil)
}
end_branch_14:
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(reverseUnrolledMap_4_1_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t14)}))
}
end_branch_19:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t19)}
}
}()
})
})
return gopurs_runtime.Apply(chunkedRevMap_1_0_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}

func Call_Data_List_Types_showList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): show_1_0 -> gopurs_runtime.Value
show_1_0 := gopurs_runtime.RecordGet(dictShow_0, "show")
_ = show_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 string
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr == nil) {
__t1 = "Nil"
goto end_branch_1
} else {

}
}
{
__t1 = (("(") + (gopurs_runtime.Apply2(Get_Data_Foldable_intercalate__2937349250(), gopurs_runtime.Str(" : "), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(Call_Data_List_Types_listMap(show_1_0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2))})))}).StrVal())) + (" : Nil)")
}
end_branch_1:
return gopurs_runtime.Str(__t1)
})})}
}

func Call_Data_List_Types_showNonEmptyList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): show_1_1 -> gopurs_runtime.Value
show_1_1 := gopurs_runtime.RecordGet(dictShow_0, "show")
_ = show_1_1
// TAST (Let): __local_var_2_2 -> *Constructor_Data_Show_Show
__local_var_2_2 := &Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 string
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr == nil) {
__t3 = "Nil"
goto end_branch_3
} else {

}
}
{
__t3 = (("(") + (gopurs_runtime.Apply2(Get_Data_Foldable_intercalate__2937349250(), gopurs_runtime.Str(" : "), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(Call_Data_List_Types_listMap(show_1_1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2))})))}).StrVal())) + (" : Nil)")
}
end_branch_3:
return gopurs_runtime.Str(__t3)
})}
_ = __local_var_2_2
// TAST (Let): showNonEmpty_1_0 -> *Constructor_Data_Show_Show
showNonEmpty_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((((("(NonEmpty ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_2_2.V0), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1).StrVal())) + (")"))
})))
_ = showNonEmpty_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(NonEmptyList ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showNonEmpty_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_2))}).StrVal())) + (")"))
})})}
}

func Call_Data_List_Types_eq1(dictEq_0_loop *Constructor_Data_Eq_Eq, xs_1_loop *Constructor_Data_List_Types_Cons, ys_2_loop *Constructor_Data_List_Types_Cons) bool {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var xs_1 *Constructor_Data_List_Types_Cons = xs_1_loop
_ = xs_1
var ys_2 *Constructor_Data_List_Types_Cons = ys_2_loop
_ = ys_2
var go__go_3_0_94 gopurs_runtime.Value
go__go_3_0_94 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
var v2_6_loop bool = (v2_6_loop_val.IntVal) != (0)
go__go_3_0_94:
for {
if false { continue go__go_3_0_94 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var v2_6 bool = v2_6_loop
_ = v2_6
var __t2 bool
{
if (v2_6) != (true) {
__t2 = false
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
var __t1 bool
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t1 = v2_6
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil)) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
v2_6_loop = (v2_6) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0).IntVal) != (0))
continue go__go_3_0_94
__t2 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
return gopurs_runtime.Bool(__t2)
}
}()
})
})
})
return (gopurs_runtime.Apply3(go__go_3_0_94, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(ys_2)}, gopurs_runtime.Bool(true)).IntVal) != (0)
}

func Call_Data_List_Types_eqNonEmpty(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_and_3 bool = false
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_NonEmpty_NonEmpty)(x_1.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(y_2.UnsafePtr).V0).IntVal) != (0) {

var go__go_3_0_95 gopurs_runtime.Value
go__go_3_0_95 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
var v2_6_loop bool = (v2_6_loop_val.IntVal) != (0)
go__go_3_0_95:
for {
if false { continue go__go_3_0_95 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var v2_6 bool = v2_6_loop
_ = v2_6
var __t2 bool
{
if (v2_6) != (true) {
__t2 = false
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
var __t1 bool
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t1 = v2_6
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil)) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
v2_6_loop = (v2_6) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0).IntVal) != (0))
continue go__go_3_0_95
__t2 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
return gopurs_runtime.Bool(__t2)
}
}()
})
})
})
__t_and_3 = (gopurs_runtime.Apply3(go__go_3_0_95, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(x_1.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(y_2.UnsafePtr).V1))}, gopurs_runtime.Bool(true)).IntVal) != (0)
}
return gopurs_runtime.Bool(__t_and_3)
})
})})}
}

func Call_Data_List_Types_eqList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_0_97 gopurs_runtime.Value
go__go_3_0_97 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
var v2_6_loop bool = (v2_6_loop_val.IntVal) != (0)
go__go_3_0_97:
for {
if false { continue go__go_3_0_97 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var v2_6 bool = v2_6_loop
_ = v2_6
var __t2 bool
{
if (v2_6) != (true) {
__t2 = false
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
var __t1 bool
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t1 = v2_6
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil)) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
v2_6_loop = (v2_6) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0).IntVal) != (0))
continue go__go_3_0_97
__t2 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
return gopurs_runtime.Bool(__t2)
}
}()
})
})
})
return gopurs_runtime.Bool((gopurs_runtime.Apply3(go__go_3_0_97, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_2))}, gopurs_runtime.Bool(true)).IntVal) != (0))
})
})})}
}

func Call_Data_List_Types_eqNonEmptyList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_and_3 bool = false
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_NonEmpty_NonEmpty)(x_1.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(y_2.UnsafePtr).V0).IntVal) != (0) {

var go__go_3_0_98 gopurs_runtime.Value
go__go_3_0_98 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
var v2_6_loop bool = (v2_6_loop_val.IntVal) != (0)
go__go_3_0_98:
for {
if false { continue go__go_3_0_98 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var v2_6 bool = v2_6_loop
_ = v2_6
var __t2 bool
{
if (v2_6) != (true) {
__t2 = false
goto end_branch_2
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
var __t1 bool
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t1 = v2_6
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil)) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
v2_6_loop = (v2_6) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0).IntVal) != (0))
continue go__go_3_0_98
__t2 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
return gopurs_runtime.Bool(__t2)
}
}()
})
})
})
__t_and_3 = (gopurs_runtime.Apply3(go__go_3_0_98, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(x_1.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(y_2.UnsafePtr).V1))}, gopurs_runtime.Bool(true)).IntVal) != (0)
}
return gopurs_runtime.Bool(__t_and_3)
})
})})}
}

func Call_Data_List_Types_compare1(dictOrd_0_loop *Constructor_Data_Ord_Ord, xs_1_loop *Constructor_Data_List_Types_Cons, ys_2_loop *Constructor_Data_List_Types_Cons) uint32 {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var xs_1 *Constructor_Data_List_Types_Cons = xs_1_loop
_ = xs_1
var ys_2 *Constructor_Data_List_Types_Cons = ys_2_loop
_ = ys_2
var go__go_3_0_100 gopurs_runtime.Value
go__go_3_0_100 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_0_100:
for {
if false { continue go__go_3_0_100 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t4 uint32
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
var __t1 uint32
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t1 = 902936544
goto end_branch_1
} else {

}
}
{
__t1 = 1527465420
}
end_branch_1:
__t4 = __t1
goto end_branch_4
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t4 = 380165415
goto end_branch_4
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil)) {
// TAST (Let): v2_6_2 -> gopurs_runtime.Value
v2_6_2 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0)
_ = v2_6_2
var __t3 uint32
{
if (uint32(v2_6_2.IntVal) == 902936544) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_0_100
__t3 = uint32(gopurs_runtime.Value{}.IntVal)
goto end_branch_3
} else {

}
}
{
__t3 = uint32(v2_6_2.IntVal)
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t4), UnsafePtr: nil}
}
}()
})
})
return uint32(gopurs_runtime.Apply2(go__go_3_0_100, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(ys_2)}).IntVal)
}

func Call_Data_List_Types_ordNonEmpty(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): eqNonEmpty2_1_0 -> *Constructor_Data_Eq_Eq
eqNonEmpty2_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_and_5 bool = false
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "eq"), (*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(y_3.UnsafePtr).V0).IntVal) != (0) {

var go__go_4_2_101 gopurs_runtime.Value
go__go_4_2_101 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
var v2_7_loop bool = (v2_7_loop_val.IntVal) != (0)
go__go_4_2_101:
for {
if false { continue go__go_4_2_101 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var v2_7 bool = v2_7_loop
_ = v2_7
var __t4 bool
{
if (v2_7) != (true) {
__t4 = false
goto end_branch_4
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
var __t3 bool
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t3 = v2_7
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
if ((v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil)) && ((v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil)) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
v2_7_loop = (v2_7) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "eq"), (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0).IntVal) != (0))
continue go__go_4_2_101
__t4 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
return gopurs_runtime.Bool(__t4)
}
}()
})
})
})
__t_and_5 = (gopurs_runtime.Apply3(go__go_4_2_101, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(y_3.UnsafePtr).V1))}, gopurs_runtime.Bool(true)).IntVal) != (0)
}
return gopurs_runtime.Bool(__t_and_5)
})
})))
_ = eqNonEmpty2_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqNonEmpty2_1_0)}
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_4_6 -> gopurs_runtime.Value
v_4_6 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(y_3.UnsafePtr).V0)
_ = v_4_6
var __t12 uint32
{
if (uint32(v_4_6.IntVal) == 1527465420) {
__t12 = 1527465420
goto end_branch_12
} else {

}
}
{
if (uint32(v_4_6.IntVal) == 380165415) {
__t12 = 380165415
goto end_branch_12
} else {

}
}
{
var go__go_5_7_102 gopurs_runtime.Value
go__go_5_7_102 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_7_102:
for {
if false { continue go__go_5_7_102 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t11 uint32
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
var __t8 uint32
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr == nil) {
__t8 = 902936544
goto end_branch_8
} else {

}
}
{
__t8 = 1527465420
}
end_branch_8:
__t11 = __t8
goto end_branch_11
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr == nil) {
__t11 = 380165415
goto end_branch_11
} else {

}
}
{
if ((v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil)) && ((v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr != nil)) {
// TAST (Let): v2_8_9 -> gopurs_runtime.Value
v2_8_9 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V0)
_ = v2_8_9
var __t10 uint32
{
if (uint32(v2_8_9.IntVal) == 902936544) {
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V1)}
continue go__go_5_7_102
__t10 = uint32(gopurs_runtime.Value{}.IntVal)
goto end_branch_10
} else {

}
}
{
__t10 = uint32(v2_8_9.IntVal)
}
end_branch_10:
__t11 = __t10
goto end_branch_11
} else {

}
}
{
__t11 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t11), UnsafePtr: nil}
}
}()
})
})
__t12 = uint32(gopurs_runtime.Apply2(go__go_5_7_102, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(y_3.UnsafePtr).V1))}).IntVal)
}
end_branch_12:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t12), UnsafePtr: nil}
})
})})}
}

func Call_Data_List_Types_ordList(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): eqList1_1_0 -> *Constructor_Data_Eq_Eq
eqList1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_105 gopurs_runtime.Value
go__go_4_2_105 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
var v2_7_loop bool = (v2_7_loop_val.IntVal) != (0)
go__go_4_2_105:
for {
if false { continue go__go_4_2_105 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var v2_7 bool = v2_7_loop
_ = v2_7
var __t4 bool
{
if (v2_7) != (true) {
__t4 = false
goto end_branch_4
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
var __t3 bool
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t3 = v2_7
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
if ((v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil)) && ((v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil)) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
v2_7_loop = (v2_7) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "eq"), (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0).IntVal) != (0))
continue go__go_4_2_105
__t4 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
return gopurs_runtime.Bool(__t4)
}
}()
})
})
})
return gopurs_runtime.Bool((gopurs_runtime.Apply3(go__go_4_2_105, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_3))}, gopurs_runtime.Bool(true)).IntVal) != (0))
})
})))
_ = eqList1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqList1_1_0)}
}), gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_5_106 gopurs_runtime.Value
go__go_4_5_106 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_5_106:
for {
if false { continue go__go_4_5_106 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t9 uint32
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
var __t6 uint32
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t6 = 902936544
goto end_branch_6
} else {

}
}
{
__t6 = 1527465420
}
end_branch_6:
__t9 = __t6
goto end_branch_9
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t9 = 380165415
goto end_branch_9
} else {

}
}
{
if ((v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil)) && ((v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil)) {
// TAST (Let): v2_7_7 -> gopurs_runtime.Value
v2_7_7 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0)
_ = v2_7_7
var __t8 uint32
{
if (uint32(v2_7_7.IntVal) == 902936544) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_5_106
__t8 = uint32(gopurs_runtime.Value{}.IntVal)
goto end_branch_8
} else {

}
}
{
__t8 = uint32(v2_7_7.IntVal)
}
end_branch_8:
__t9 = __t8
goto end_branch_9
} else {

}
}
{
__t9 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t9), UnsafePtr: nil}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(go__go_4_5_106, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_3))}).IntVal)), UnsafePtr: nil}
})
})})}
}

func Call_Data_List_Types_ordNonEmptyList(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): eqNonEmpty2_1_0 -> *Constructor_Data_Eq_Eq
eqNonEmpty2_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_and_5 bool = false
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "eq"), (*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(y_3.UnsafePtr).V0).IntVal) != (0) {

var go__go_4_2_107 gopurs_runtime.Value
go__go_4_2_107 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
var v2_7_loop bool = (v2_7_loop_val.IntVal) != (0)
go__go_4_2_107:
for {
if false { continue go__go_4_2_107 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var v2_7 bool = v2_7_loop
_ = v2_7
var __t4 bool
{
if (v2_7) != (true) {
__t4 = false
goto end_branch_4
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
var __t3 bool
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t3 = v2_7
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
if ((v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil)) && ((v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil)) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
v2_7_loop = (v2_7) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "eq"), (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0).IntVal) != (0))
continue go__go_4_2_107
__t4 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
return gopurs_runtime.Bool(__t4)
}
}()
})
})
})
__t_and_5 = (gopurs_runtime.Apply3(go__go_4_2_107, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(y_3.UnsafePtr).V1))}, gopurs_runtime.Bool(true)).IntVal) != (0)
}
return gopurs_runtime.Bool(__t_and_5)
})
})))
_ = eqNonEmpty2_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqNonEmpty2_1_0)}
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_4_6 -> gopurs_runtime.Value
v_4_6 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty)(y_3.UnsafePtr).V0)
_ = v_4_6
var __t12 uint32
{
if (uint32(v_4_6.IntVal) == 1527465420) {
__t12 = 1527465420
goto end_branch_12
} else {

}
}
{
if (uint32(v_4_6.IntVal) == 380165415) {
__t12 = 380165415
goto end_branch_12
} else {

}
}
{
var go__go_5_7_108 gopurs_runtime.Value
go__go_5_7_108 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_7_108:
for {
if false { continue go__go_5_7_108 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t11 uint32
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
var __t8 uint32
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr == nil) {
__t8 = 902936544
goto end_branch_8
} else {

}
}
{
__t8 = 1527465420
}
end_branch_8:
__t11 = __t8
goto end_branch_11
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr == nil) {
__t11 = 380165415
goto end_branch_11
} else {

}
}
{
if ((v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil)) && ((v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr != nil)) {
// TAST (Let): v2_8_9 -> gopurs_runtime.Value
v2_8_9 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V0)
_ = v2_8_9
var __t10 uint32
{
if (uint32(v2_8_9.IntVal) == 902936544) {
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V1)}
continue go__go_5_7_108
__t10 = uint32(gopurs_runtime.Value{}.IntVal)
goto end_branch_10
} else {

}
}
{
__t10 = uint32(v2_8_9.IntVal)
}
end_branch_10:
__t11 = __t10
goto end_branch_11
} else {

}
}
{
__t11 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t11), UnsafePtr: nil}
}
}()
})
})
__t12 = uint32(gopurs_runtime.Apply2(go__go_5_7_108, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(x_2.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_NonEmpty_NonEmpty)(y_3.UnsafePtr).V1))}).IntVal)
}
end_branch_12:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t12), UnsafePtr: nil}
})
})})}
}

func Call_Data_List_Types_pure(x_0_loop gopurs_runtime.Value) *Constructor_Data_NonEmpty_NonEmpty {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, x_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})})
}

func Call_Data_List_Types_listMap__858544730(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var chunkedRevMap_1_0_189 gopurs_runtime.Value
chunkedRevMap_1_0_189 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2_loop_val)
var v1_3_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_3_loop_val)
chunkedRevMap_1_0_189:
for {
if false { continue chunkedRevMap_1_0_189 }
var v_2 *Constructor_Data_List_Types_Cons = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Types_Cons = v1_3_loop
_ = v1_3
var __t19 *Constructor_Data_List_Types_Cons
{
var __t_and_18 bool = false
if (v1_3 != nil) {

var __t_tag_15 *Constructor_Data_List_Types_Cons = (v1_3).V1
var __t_and_17 bool = false
if (__t_tag_15 != nil) {

var __t_tag_16 *Constructor_Data_List_Types_Cons = ((v1_3).V1).V1
__t_and_17 = (__t_tag_16 != nil)
}
__t_and_18 = __t_and_17
}
if __t_and_18 {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)})})})
v1_3_loop = (((v1_3).V1).V1).V1
continue chunkedRevMap_1_0_189
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_19
} else {

}
}
{
var reverseUnrolledMap_4_1_190 gopurs_runtime.Value
reverseUnrolledMap_4_1_190 = gopurs_runtime.Func(func(v2_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v2_5_loop gopurs_runtime.Value = v2_5_loop_val
var v3_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v3_6_loop_val)
reverseUnrolledMap_4_1_190:
for {
if false { continue reverseUnrolledMap_4_1_190 }
var v2_5 gopurs_runtime.Value = v2_5_loop
_ = v2_5
var v3_6 *Constructor_Data_List_Types_Cons = v3_6_loop
_ = v3_6
var __t8 *Constructor_Data_List_Types_Cons
{
var __t_and_7 bool = false
if (v2_5.Type == 9 && v2_5.IntVal == 1358893437 && v2_5.UnsafePtr != nil) {

var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_List_Types_Cons)(v2_5.UnsafePtr).V0
var __t_and_6 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1358893437 && __t_tag_2.UnsafePtr != nil) {

var __t_tag_3 *Constructor_Data_List_Types_Cons = (*Constructor_Data_List_Types_Cons)((*Constructor_Data_List_Types_Cons)(v2_5.UnsafePtr).V0.UnsafePtr).V1
var __t_and_5 bool = false
if (__t_tag_3 != nil) {

var __t_tag_4 *Constructor_Data_List_Types_Cons = ((*Constructor_Data_List_Types_Cons)((*Constructor_Data_List_Types_Cons)(v2_5.UnsafePtr).V0.UnsafePtr).V1).V1
__t_and_5 = (__t_tag_4 != nil)
}
__t_and_6 = __t_and_5
}
__t_and_7 = __t_and_6
}
if __t_and_7 {
v2_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v2_5.UnsafePtr).V1)}
v3_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Str(gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons)((*Constructor_Data_List_Types_Cons)(v2_5.UnsafePtr).V0.UnsafePtr).V0).StrVal()), &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Str(gopurs_runtime.Apply(f_0, ((*Constructor_Data_List_Types_Cons)((*Constructor_Data_List_Types_Cons)(v2_5.UnsafePtr).V0.UnsafePtr).V1).V0).StrVal()), &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Str(gopurs_runtime.Apply(f_0, (((*Constructor_Data_List_Types_Cons)((*Constructor_Data_List_Types_Cons)(v2_5.UnsafePtr).V0.UnsafePtr).V1).V1).V0).StrVal()), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v3_6)})}}})})
continue reverseUnrolledMap_4_1_190
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_8
} else {

}
}
{
__t8 = v3_6
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t8)}
}
}()
})
})
var __t14 *Constructor_Data_List_Types_Cons
{
if (v1_3 != nil) {
var __t13 *Constructor_Data_List_Types_Cons
{
var __t_tag_9 *Constructor_Data_List_Types_Cons = (v1_3).V1
if (__t_tag_9 != nil) {
var __t11 *Constructor_Data_List_Types_Cons
{
var __t_tag_10 *Constructor_Data_List_Types_Cons = ((v1_3).V1).V1
if (__t_tag_10 == nil) {
__t11 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Str(gopurs_runtime.Apply(f_0, (v1_3).V0).StrVal()), &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Str(gopurs_runtime.Apply(f_0, ((v1_3).V1).V0).StrVal()), (*Constructor_Data_List_Types_Cons)(nil)}}
goto end_branch_11
} else {

}
}
{
__t11 = (*Constructor_Data_List_Types_Cons)(nil)
}
end_branch_11:
__t13 = __t11
goto end_branch_13
} else {

}
}
{
var __t_tag_12 *Constructor_Data_List_Types_Cons = (v1_3).V1
if (__t_tag_12 == nil) {
__t13 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Str(gopurs_runtime.Apply(f_0, (v1_3).V0).StrVal()), (*Constructor_Data_List_Types_Cons)(nil)}
goto end_branch_13
} else {

}
}
{
__t13 = (*Constructor_Data_List_Types_Cons)(nil)
}
end_branch_13:
__t14 = __t13
goto end_branch_14
} else {

}
}
{
__t14 = (*Constructor_Data_List_Types_Cons)(nil)
}
end_branch_14:
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(reverseUnrolledMap_4_1_190, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t14)}))
}
end_branch_19:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t19)}
}
}()
})
})
return gopurs_runtime.Apply(chunkedRevMap_1_0_189, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}

func Call_Data_List_Types_listMap__4135416762(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var chunkedRevMap_1_0_191 gopurs_runtime.Value
chunkedRevMap_1_0_191 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2_loop_val)
var v1_3_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_3_loop_val)
chunkedRevMap_1_0_191:
for {
if false { continue chunkedRevMap_1_0_191 }
var v_2 *Constructor_Data_List_Types_Cons = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Types_Cons = v1_3_loop
_ = v1_3
var __t19 *Constructor_Data_List_Types_Cons
{
var __t_and_18 bool = false
if (v1_3 != nil) {

var __t_tag_15 *Constructor_Data_List_Types_Cons = (v1_3).V1
var __t_and_17 bool = false
if (__t_tag_15 != nil) {

var __t_tag_16 *Constructor_Data_List_Types_Cons = ((v1_3).V1).V1
__t_and_17 = (__t_tag_16 != nil)
}
__t_and_18 = __t_and_17
}
if __t_and_18 {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)})})})
v1_3_loop = (((v1_3).V1).V1).V1
continue chunkedRevMap_1_0_191
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_19
} else {

}
}
{
var reverseUnrolledMap_4_1_192 gopurs_runtime.Value
reverseUnrolledMap_4_1_192 = gopurs_runtime.Func(func(v2_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v2_5_loop gopurs_runtime.Value = v2_5_loop_val
var v3_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v3_6_loop_val)
reverseUnrolledMap_4_1_192:
for {
if false { continue reverseUnrolledMap_4_1_192 }
var v2_5 gopurs_runtime.Value = v2_5_loop
_ = v2_5
var v3_6 *Constructor_Data_List_Types_Cons = v3_6_loop
_ = v3_6
var __t8 *Constructor_Data_List_Types_Cons
{
var __t_and_7 bool = false
if (v2_5.Type == 9 && v2_5.IntVal == 1358893437 && v2_5.UnsafePtr != nil) {

var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_List_Types_Cons)(v2_5.UnsafePtr).V0
var __t_and_6 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1358893437 && __t_tag_2.UnsafePtr != nil) {

var __t_tag_3 *Constructor_Data_List_Types_Cons = (*Constructor_Data_List_Types_Cons)((*Constructor_Data_List_Types_Cons)(v2_5.UnsafePtr).V0.UnsafePtr).V1
var __t_and_5 bool = false
if (__t_tag_3 != nil) {

var __t_tag_4 *Constructor_Data_List_Types_Cons = ((*Constructor_Data_List_Types_Cons)((*Constructor_Data_List_Types_Cons)(v2_5.UnsafePtr).V0.UnsafePtr).V1).V1
__t_and_5 = (__t_tag_4 != nil)
}
__t_and_6 = __t_and_5
}
__t_and_7 = __t_and_6
}
if __t_and_7 {
v2_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v2_5.UnsafePtr).V1)}
v3_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons)((*Constructor_Data_List_Types_Cons)(v2_5.UnsafePtr).V0.UnsafePtr).V0), &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Apply(f_0, ((*Constructor_Data_List_Types_Cons)((*Constructor_Data_List_Types_Cons)(v2_5.UnsafePtr).V0.UnsafePtr).V1).V0), &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Apply(f_0, (((*Constructor_Data_List_Types_Cons)((*Constructor_Data_List_Types_Cons)(v2_5.UnsafePtr).V0.UnsafePtr).V1).V1).V0), v3_6}}})})
continue reverseUnrolledMap_4_1_192
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_8
} else {

}
}
{
__t8 = v3_6
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t8)}
}
}()
})
})
var __t14 *Constructor_Data_List_Types_Cons
{
if (v1_3 != nil) {
var __t13 *Constructor_Data_List_Types_Cons
{
var __t_tag_9 *Constructor_Data_List_Types_Cons = (v1_3).V1
if (__t_tag_9 != nil) {
var __t11 *Constructor_Data_List_Types_Cons
{
var __t_tag_10 *Constructor_Data_List_Types_Cons = ((v1_3).V1).V1
if (__t_tag_10 == nil) {
__t11 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Apply(f_0, (v1_3).V0), &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Apply(f_0, ((v1_3).V1).V0), (*Constructor_Data_List_Types_Cons)(nil)}}
goto end_branch_11
} else {

}
}
{
__t11 = (*Constructor_Data_List_Types_Cons)(nil)
}
end_branch_11:
__t13 = __t11
goto end_branch_13
} else {

}
}
{
var __t_tag_12 *Constructor_Data_List_Types_Cons = (v1_3).V1
if (__t_tag_12 == nil) {
__t13 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Apply(f_0, (v1_3).V0), (*Constructor_Data_List_Types_Cons)(nil)}
goto end_branch_13
} else {

}
}
{
__t13 = (*Constructor_Data_List_Types_Cons)(nil)
}
end_branch_13:
__t14 = __t13
goto end_branch_14
} else {

}
}
{
__t14 = (*Constructor_Data_List_Types_Cons)(nil)
}
end_branch_14:
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(reverseUnrolledMap_4_1_192, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t14)}))
}
end_branch_19:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t19)}
}
}()
})
})
return gopurs_runtime.Apply(chunkedRevMap_1_0_191, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}

func Call_Data_List_Types_nelCons__195558898(a_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_NonEmpty_NonEmpty {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_1).V1)})}})})
}

func Call_Data_List_Types_nelCons__2148523118(a_0_loop *Constructor_Data_List_Types_Cons, v_1_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_NonEmpty_NonEmpty {
var a_0 *Constructor_Data_List_Types_Cons = a_0_loop
_ = a_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(a_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_1).V0))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_1).V1))})})}})})
}

func Call_Data_List_Types_toList__2859885498(v_0_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_List_Types_Cons {
var v_0 *Constructor_Data_NonEmpty_NonEmpty = v_0_loop
_ = v_0
return &Constructor_Data_List_Types_Cons{1, (v_0).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_0).V1)}
}

func Call_Data_List_Types_toList__1324737658(v_0_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_List_Types_Cons {
var v_0 *Constructor_Data_NonEmpty_NonEmpty = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_0).V0))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_0).V1))})})})
}

func Call_Data_List_Types_toList__2402503393(v_0_loop *Constructor_Data_NonEmpty_NonEmpty) *Constructor_Data_List_Types_Cons {
var v_0 *Constructor_Data_NonEmpty_NonEmpty = v_0_loop
_ = v_0
return &Constructor_Data_List_Types_Cons{1, (v_0).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_0).V1)}
}



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
		cache_Data_List_Types_Nil = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}
	})
	return cache_Data_List_Types_Nil
}

var cache_Data_List_Types_Cons gopurs_runtime.Value
var once_Data_List_Types_Cons sync.Once
func Get_Data_List_Types_Cons() gopurs_runtime.Value {
	once_Data_List_Types_Cons.Do(func() {
		cache_Data_List_Types_Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](value1)})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_Types_toList(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_Data_List_Types_toList
}

var cache_Data_List_Types_newtypeNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_newtypeNonEmptyList sync.Once
func Get_Data_List_Types_newtypeNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_newtypeNonEmptyList.Do(func() {
		cache_Data_List_Types_newtypeNonEmptyList = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_List_Types_newtypeNonEmptyList
}

var cache_Data_List_Types_nelCons gopurs_runtime.Value
var once_Data_List_Types_nelCons sync.Once
func Get_Data_List_Types_nelCons() gopurs_runtime.Value {
	once_Data_List_Types_nelCons.Do(func() {
		cache_Data_List_Types_nelCons = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_Types_nelCons(a_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box)))}
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
		cache_Data_List_Types_functorList = gopurs_runtime.RecordDict1("map", Get_Data_List_Types_listMap())
	})
	return cache_Data_List_Types_functorList
}

var cache_Data_List_Types_functorNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_functorNonEmptyList sync.Once
func Get_Data_List_Types_functorNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_functorNonEmptyList.Do(func() {
		cache_Data_List_Types_functorNonEmptyList = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_functorList(), "map"), f_0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_Data_List_Types_functorNonEmptyList
}

var cache_Data_List_Types_foldableList gopurs_runtime.Value
var once_Data_List_Types_foldableList sync.Once
func Get_Data_List_Types_foldableList() gopurs_runtime.Value {
	once_Data_List_Types_foldableList.Do(func() {
		cache_Data_List_Types_foldableList = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
b_2_loop = gopurs_runtime.Apply2(f_0, b_2, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
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
__local_var_2_5 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
})
}), b_1)
_ = __local_var_2_5
var go__go_3_7_3 gopurs_runtime.Value
go__go_3_7_3 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_7_3:
for {
if false { continue go__go_3_7_3 }
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t8 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}
goto end_branch_8
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
continue go__go_3_7_3
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__t8))}
}
}()
})
})
// TAST (Let): __local_var_3_6 -> gopurs_runtime.Value
__local_var_3_6 := gopurs_runtime.Apply(go__go_3_7_3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_3_6
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_5, gopurs_runtime.Apply(__local_var_3_6, x_4))
})
})
}))
	})
	return cache_Data_List_Types_foldableList
}

var cache_Data_List_Types_foldableNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_foldableNonEmptyList sync.Once
func Get_Data_List_Types_foldableNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_foldableNonEmptyList.Do(func() {
		cache_Data_List_Types_foldableNonEmptyList = gopurs_runtime.Apply(Get_Data_NonEmpty_foldableNonEmpty(), Get_Data_List_Types_foldableList())
	})
	return cache_Data_List_Types_foldableNonEmptyList
}

var cache_Data_List_Types_foldableWithIndexList gopurs_runtime.Value
var once_Data_List_Types_foldableWithIndexList sync.Once
func Get_Data_List_Types_foldableWithIndexList() gopurs_runtime.Value {
	once_Data_List_Types_foldableWithIndexList.Do(func() {
		cache_Data_List_Types_foldableWithIndexList = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_foldableList()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableWithIndexList(), "foldlWithIndex"), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
__local_var_2_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_5 -> gopurs_runtime.Value
__local_var_3_5 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_3_5
// TAST (Let): __local_var_4_6 -> gopurs_runtime.Value
__local_var_4_6 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_4_6
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_4_6.IntVal) + (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_4_6.IntVal), __local_var_3_5, a_5)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), acc_1})})
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(__local_var_2_4, x_3).UnsafePtr).V1
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_3_7 -> *Constructor_Data_Tuple_Tuple[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]
v_3_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_8 -> gopurs_runtime.Value
__local_var_4_8 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_4_8
// TAST (Let): __local_var_5_9 -> gopurs_runtime.Value
__local_var_5_9 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0
_ = __local_var_5_9
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_5_9.IntVal) + (1)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_6, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__local_var_4_8)})}})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_2))}))
_ = v_3_7
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_10 -> gopurs_runtime.Value
__local_var_5_10 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_10
// TAST (Let): __local_var_6_11 -> gopurs_runtime.Value
__local_var_6_11 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0
_ = __local_var_6_11
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_6_11.IntVal) - (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((__local_var_6_11.IntVal) - (1)), a_7, __local_var_5_10)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_7)}.UnsafePtr).V0.IntVal), b_1})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_7)}.UnsafePtr).V1))}).UnsafePtr).V1
})
})
}))
	})
	return cache_Data_List_Types_foldableWithIndexList
}

var cache_Data_List_Types_foldableWithIndexNonEmpty gopurs_runtime.Value
var once_Data_List_Types_foldableWithIndexNonEmpty sync.Once
func Get_Data_List_Types_foldableWithIndexNonEmpty() gopurs_runtime.Value {
	once_Data_List_Types_foldableWithIndexNonEmpty.Do(func() {
		cache_Data_List_Types_foldableWithIndexNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](gopurs_runtime.Apply(Get_Data_NonEmpty_foldableWithIndexNonEmpty(), Get_Data_List_Types_foldableWithIndexList())))}
	})
	return cache_Data_List_Types_foldableWithIndexNonEmpty
}

var cache_Data_List_Types_foldableWithIndexNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_foldableWithIndexNonEmptyList sync.Once
func Get_Data_List_Types_foldableWithIndexNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_foldableWithIndexNonEmptyList.Do(func() {
		cache_Data_List_Types_foldableWithIndexNonEmptyList = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_foldableNonEmptyList()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_foldableWithIndexNonEmpty()).V1), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))}, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Int((1) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0.IntVal))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Apply(f_1, gopurs_runtime.Int(__t0.IntVal))
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_foldableWithIndexNonEmpty()).V2), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Int(0)
goto end_branch_1
} else {

}
}
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr != nil) {
__t1 = gopurs_runtime.Int((1) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0.IntVal))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Apply(f_0, gopurs_runtime.Int(__t1.IntVal))
}), b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_foldableWithIndexNonEmpty()).V3), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr == nil) {
__t2 = gopurs_runtime.Int(0)
goto end_branch_2
} else {

}
}
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr != nil) {
__t2 = gopurs_runtime.Int((1) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0.IntVal))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Apply(f_0, gopurs_runtime.Int(__t2.IntVal))
}), b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2))})
})
})
}))
	})
	return cache_Data_List_Types_foldableWithIndexNonEmptyList
}

var cache_Data_List_Types_functorWithIndexList gopurs_runtime.Value
var once_Data_List_Types_functorWithIndexList sync.Once
func Get_Data_List_Types_functorWithIndexList() gopurs_runtime.Value {
	once_Data_List_Types_functorWithIndexList.Do(func() {
		cache_Data_List_Types_functorWithIndexList = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_functorList()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableWithIndexList(), "foldrWithIndex"), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, gopurs_runtime.Int(i_1.IntVal), x_2), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](acc_3)})}
})
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
}))
	})
	return cache_Data_List_Types_functorWithIndexList
}

var cache_Data_List_Types_functorWithIndex gopurs_runtime.Value
var once_Data_List_Types_functorWithIndex sync.Once
func Get_Data_List_Types_functorWithIndex() gopurs_runtime.Value {
	once_Data_List_Types_functorWithIndex.Do(func() {
		cache_Data_List_Types_functorWithIndex = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_1 -> gopurs_runtime.Value
__local_var_0_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Types_functorWithIndexList(), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_0_1
// TAST (Let): functorNonEmpty1_0_0 -> gopurs_runtime.Value
functorNonEmpty1_0_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_1, "map"), f_1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1)})}
})
}))
_ = functorNonEmpty1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return functorNonEmpty1_0_0
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_functorWithIndexList(), "mapWithIndex"), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, x_3})})
}), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
})
}))))}
}()
	})
	return cache_Data_List_Types_functorWithIndex
}

var cache_Data_List_Types_functorWithIndexNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_functorWithIndexNonEmptyList sync.Once
func Get_Data_List_Types_functorWithIndexNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_functorWithIndexNonEmptyList.Do(func() {
		cache_Data_List_Types_functorWithIndexNonEmptyList = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_functorNonEmptyList()
}), gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_functorWithIndex()).V1), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_2.Type == 9 && x_2.IntVal == 930809136 && x_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (x_2.Type == 9 && x_2.IntVal == 930809136 && x_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Int((1) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_2.UnsafePtr).V0.IntVal))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Apply(fn_0, gopurs_runtime.Int(__t0.IntVal))
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1))})))}
})
}))
	})
	return cache_Data_List_Types_functorWithIndexNonEmptyList
}

var cache_Data_List_Types_semigroupList gopurs_runtime.Value
var once_Data_List_Types_semigroupList sync.Once
func Get_Data_List_Types_semigroupList() gopurs_runtime.Value {
	once_Data_List_Types_semigroupList.Do(func() {
		cache_Data_List_Types_semigroupList = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldr"), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_0))})))}
})
}))
	})
	return cache_Data_List_Types_semigroupList
}

var cache_Data_List_Types_monoidList gopurs_runtime.Value
var once_Data_List_Types_monoidList sync.Once
func Get_Data_List_Types_monoidList() gopurs_runtime.Value {
	once_Data_List_Types_monoidList.Do(func() {
		cache_Data_List_Types_monoidList = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_semigroupList()
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
	})
	return cache_Data_List_Types_monoidList
}

var cache_Data_List_Types_semigroupNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_semigroupNonEmptyList sync.Once
func Get_Data_List_Types_semigroupNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_semigroupNonEmptyList.Do(func() {
		cache_Data_List_Types_semigroupNonEmptyList = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(as_prime_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(as_prime_1.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(as_prime_1.UnsafePtr).V1)})})))}})}
})
}))
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
		cache_Data_List_Types_traversableList = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_foldableList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_functorList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_traversableList(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, Get_Data_List_Types_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): Apply0_2_1 -> *Constructor_Control_Apply_Apply[gopurs_runtime.Value]
Apply0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_5, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_4)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}))
_ = __local_var_4_2
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_6_5 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_2_1.V0), gopurs_runtime.Value{}))
_ = Functor0_6_5
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_5.V0), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_9, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_8)})}
})
}), acc_5), b_7)
})
_ = __local_var_6_4
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, gopurs_runtime.Apply(f_3, x_7))
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}))
_ = __local_var_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_5_3, x_6))
})
})
}))
	})
	return cache_Data_List_Types_traversableList
}

var cache_Data_List_Types_traversableNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_traversableNonEmptyList sync.Once
func Get_Data_List_Types_traversableNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_traversableNonEmptyList.Do(func() {
		cache_Data_List_Types_traversableNonEmptyList = gopurs_runtime.Apply(Get_Data_NonEmpty_traversableNonEmpty(), Get_Data_List_Types_traversableList())
	})
	return cache_Data_List_Types_traversableNonEmptyList
}

var cache_Data_List_Types_traversableWithIndexList gopurs_runtime.Value
var once_Data_List_Types_traversableWithIndexList sync.Once
func Get_Data_List_Types_traversableWithIndexList() gopurs_runtime.Value {
	once_Data_List_Types_traversableWithIndexList.Do(func() {
		cache_Data_List_Types_traversableWithIndexList = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_foldableWithIndexList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_functorWithIndexList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_traversableList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): Apply0_2_1 -> *Constructor_Control_Apply_Apply[gopurs_runtime.Value]
Apply0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_5, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_4)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})))})
_ = __local_var_4_2
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableWithIndexList(), "foldlWithIndex"), gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_7_5 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_7_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_2_1.V0), gopurs_runtime.Value{}))
_ = Functor0_7_5
// TAST (Let): __local_var_7_4 -> gopurs_runtime.Value
__local_var_7_4 := gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_5.V0), gopurs_runtime.Func(func(b_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_10, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_9)})}
})
}), acc_6), b_8)
})
_ = __local_var_7_4
// TAST (Let): __local_var_8_6 -> gopurs_runtime.Value
__local_var_8_6 := gopurs_runtime.Apply(f_3, gopurs_runtime.Int(i_5.IntVal))
_ = __local_var_8_6
return gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_4, gopurs_runtime.Apply(__local_var_8_6, x_9))
})
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}))
_ = __local_var_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_5_3, x_6))
})
})
}))
	})
	return cache_Data_List_Types_traversableWithIndexList
}

var cache_Data_List_Types_traversableWithIndexNonEmpty gopurs_runtime.Value
var once_Data_List_Types_traversableWithIndexNonEmpty sync.Once
func Get_Data_List_Types_traversableWithIndexNonEmpty() gopurs_runtime.Value {
	once_Data_List_Types_traversableWithIndexNonEmpty.Do(func() {
		cache_Data_List_Types_traversableWithIndexNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](gopurs_runtime.Apply(Get_Data_NonEmpty_traversableWithIndexNonEmpty(), Get_Data_List_Types_traversableWithIndexList())))}
	})
	return cache_Data_List_Types_traversableWithIndexNonEmpty
}

var cache_Data_List_Types_traversableWithIndexNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_traversableWithIndexNonEmptyList sync.Once
func Get_Data_List_Types_traversableWithIndexNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_traversableWithIndexNonEmptyList.Do(func() {
		cache_Data_List_Types_traversableWithIndexNonEmptyList = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_foldableWithIndexNonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_functorWithIndexNonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_traversableNonEmptyList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_List_Types_NonEmptyList(), gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_traversableWithIndexNonEmpty()).V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (x_4.Type == 9 && x_4.IntVal == 930809136 && x_4.UnsafePtr == nil) {
__t1 = gopurs_runtime.Int(0)
goto end_branch_1
} else {

}
}
{
if (x_4.Type == 9 && x_4.IntVal == 930809136 && x_4.UnsafePtr != nil) {
__t1 = gopurs_runtime.Int((1) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_4.UnsafePtr).V0.IntVal))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Apply(f_2, gopurs_runtime.Int(__t1.IntVal))
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_3))}))
})
})
}))
	})
	return cache_Data_List_Types_traversableWithIndexNonEmptyList
}

var cache_Data_List_Types_unfoldable1List gopurs_runtime.Value
var once_Data_List_Types_unfoldable1List sync.Once
func Get_Data_List_Types_unfoldable1List() gopurs_runtime.Value {
	once_Data_List_Types_unfoldable1List.Do(func() {
		cache_Data_List_Types_unfoldable1List = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_4 gopurs_runtime.Value
go__go_2_0_4 = gopurs_runtime.Func(func(source_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var source_3_loop gopurs_runtime.Value = source_3_loop_val
var memo_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](memo_4_loop_val)
go__go_2_0_4:
for {
if false { continue go__go_2_0_4 }
var source_3 gopurs_runtime.Value = source_3_loop
_ = source_3
var memo_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = memo_4_loop
_ = memo_4
// TAST (Let): v_5_1 -> *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]
v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](gopurs_runtime.Apply(f_0, source_3))
_ = v_5_1
var __t4 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 930809136 && __t_tag_2.UnsafePtr != nil) {
source_3_loop = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V1.UnsafePtr).V0
memo_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V0, memo_4})})
continue go__go_2_0_4
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V1
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 930809136 && __t_tag_3.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_7, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_6)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V0, memo_4})})))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__t4))}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_2_0_4, b_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})))}
})
}))
	})
	return cache_Data_List_Types_unfoldable1List
}

var cache_Data_List_Types_unfoldableList gopurs_runtime.Value
var once_Data_List_Types_unfoldableList sync.Once
func Get_Data_List_Types_unfoldableList() gopurs_runtime.Value {
	once_Data_List_Types_unfoldableList.Do(func() {
		cache_Data_List_Types_unfoldableList = gopurs_runtime.RecordDict2("Unfoldable10", "unfoldr", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_unfoldable1List()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_5 gopurs_runtime.Value
go__go_2_0_5 = gopurs_runtime.Func(func(source_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var source_3_loop gopurs_runtime.Value = source_3_loop_val
var memo_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](memo_4_loop_val)
go__go_2_0_5:
for {
if false { continue go__go_2_0_5 }
var source_3 gopurs_runtime.Value = source_3_loop
_ = source_3
var memo_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = memo_4_loop
_ = memo_4
// TAST (Let): v_5_1 -> *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]
v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(f_0, source_3))
_ = v_5_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_7, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_6)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(memo_4)})))}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr != nil) {
source_3_loop = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V0.UnsafePtr).V1
memo_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V0.UnsafePtr).V0, memo_4})})
continue go__go_2_0_5
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__t2))}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_2_0_5, b_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})))}
})
}))
	})
	return cache_Data_List_Types_unfoldableList
}

var cache_Data_List_Types_unfoldable1NonEmptyList gopurs_runtime.Value
var once_Data_List_Types_unfoldable1NonEmptyList sync.Once
func Get_Data_List_Types_unfoldable1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_unfoldable1NonEmptyList.Do(func() {
		cache_Data_List_Types_unfoldable1NonEmptyList = gopurs_runtime.Apply(Get_Data_NonEmpty_unfoldable1NonEmpty(), Get_Data_List_Types_unfoldableList())
	})
	return cache_Data_List_Types_unfoldable1NonEmptyList
}

var cache_Data_List_Types_foldable1NonEmptyList gopurs_runtime.Value
var once_Data_List_Types_foldable1NonEmptyList sync.Once
func Get_Data_List_Types_foldable1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_foldable1NonEmptyList.Do(func() {
		cache_Data_List_Types_foldable1NonEmptyList = gopurs_runtime.Apply(Get_Data_NonEmpty_foldable1NonEmpty(), Get_Data_List_Types_foldableList())
	})
	return cache_Data_List_Types_foldable1NonEmptyList
}

var cache_Data_List_Types_extendNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_extendNonEmptyList sync.Once
func Get_Data_List_Types_extendNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_extendNonEmptyList.Do(func() {
		cache_Data_List_Types_extendNonEmptyList = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_functorNonEmptyList()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldr"), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v1_3, "acc"))})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_2, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v1_3, "acc")))}})}), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v1_3, "val"))})})
})
}), gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1))}), "val")))}})}
})
}))
	})
	return cache_Data_List_Types_extendNonEmptyList
}

var cache_Data_List_Types_extendList gopurs_runtime.Value
var once_Data_List_Types_extendList sync.Once
func Get_Data_List_Types_extendList() gopurs_runtime.Value {
	once_Data_List_Types_extendList.Do(func() {
		cache_Data_List_Types_extendList = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_functorList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 1358893437 && v1_1.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 1358893437 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1))}), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldr"), gopurs_runtime.Func(func(a_prime_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_prime_2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v2_3, "acc"))})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_prime_2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v2_3, "acc"))})}), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v2_3, "val"))})})
})
}), gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_1.UnsafePtr).V1)}), "val"))})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_Data_List_Types_extendList
}

var cache_Data_List_Types_eq1List gopurs_runtime.Value
var once_Data_List_Types_eq1List sync.Once
func Get_Data_List_Types_eq1List() gopurs_runtime.Value {
	once_Data_List_Types_eq1List.Do(func() {
		cache_Data_List_Types_eq1List = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_0_6 gopurs_runtime.Value
go__go_3_0_6 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
var v2_6_loop bool = (v2_6_loop_val.IntVal) != (0)
go__go_3_0_6:
for {
if false { continue go__go_3_0_6 }
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
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
v2_6_loop = (v2_6) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0).IntVal) != (0))
continue go__go_3_0_6
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
return gopurs_runtime.Bool((gopurs_runtime.Apply3(go__go_3_0_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_2))}, gopurs_runtime.Bool(true)).IntVal) != (0))
})
})
}))
	})
	return cache_Data_List_Types_eq1List
}

var cache_Data_List_Types_eq1 gopurs_runtime.Value
var once_Data_List_Types_eq1 sync.Once
func Get_Data_List_Types_eq1() gopurs_runtime.Value {
	once_Data_List_Types_eq1.Do(func() {
		cache_Data_List_Types_eq1 = gopurs_runtime.RecordGet(Get_Data_List_Types_eq1List(), "eq1")
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
		cache_Data_List_Types_eq1NonEmptyList = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_eq1List(), "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0))}, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V1).IntVal) != (0)))
})
})
}))
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
		cache_Data_List_Types_ord1List = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_eq1List()
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_0_7 gopurs_runtime.Value
go__go_3_0_7 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_0_7:
for {
if false { continue go__go_3_0_7 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t4 gopurs_runtime.Value
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
__t4 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t1), UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil)) {
// TAST (Let): v2_6_2 -> gopurs_runtime.Value
v2_6_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0)
_ = v2_6_2
var __t3 uint32
{
if (uint32(v2_6_2.IntVal) == 902936544) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
continue go__go_3_0_7
__t3 = uint32(gopurs_runtime.Value{}.IntVal)
goto end_branch_3
} else {

}
}
{
__t3 = uint32(v2_6_2.IntVal)
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t3), UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t4.IntVal)), UnsafePtr: nil}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(go__go_3_0_7, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_2))}).IntVal)), UnsafePtr: nil}
})
})
}))
	})
	return cache_Data_List_Types_ord1List
}

var cache_Data_List_Types_compare1 gopurs_runtime.Value
var once_Data_List_Types_compare1 sync.Once
func Get_Data_List_Types_compare1() gopurs_runtime.Value {
	once_Data_List_Types_compare1.Do(func() {
		cache_Data_List_Types_compare1 = gopurs_runtime.RecordGet(Get_Data_List_Types_ord1List(), "compare1")
	})
	return cache_Data_List_Types_compare1
}

var cache_Data_List_Types_ordNonEmpty gopurs_runtime.Value
var once_Data_List_Types_ordNonEmpty sync.Once
func Get_Data_List_Types_ordNonEmpty() gopurs_runtime.Value {
	once_Data_List_Types_ordNonEmpty.Do(func() {
		cache_Data_List_Types_ordNonEmpty = gopurs_runtime.Apply(Get_Data_NonEmpty_ordNonEmpty(), Get_Data_List_Types_ord1List())
	})
	return cache_Data_List_Types_ordNonEmpty
}

var cache_Data_List_Types_ord1NonEmptyList gopurs_runtime.Value
var once_Data_List_Types_ord1NonEmptyList sync.Once
func Get_Data_List_Types_ord1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_ord1NonEmptyList.Do(func() {
		cache_Data_List_Types_ord1NonEmptyList = gopurs_runtime.Apply(Get_Data_NonEmpty_ord1NonEmpty(), Get_Data_List_Types_ord1List())
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
		cache_Data_List_Types_comonadNonEmptyList = gopurs_runtime.RecordDict2("Extend0", "extract", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_extendNonEmptyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0
}))
	})
	return cache_Data_List_Types_comonadNonEmptyList
}

var cache_Data_List_Types_applyList gopurs_runtime.Value
var once_Data_List_Types_applyList sync.Once
func Get_Data_List_Types_applyList() gopurs_runtime.Value {
	once_Data_List_Types_applyList.Do(func() {
		cache_Data_List_Types_applyList = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_functorList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_functorList(), "map"), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1))})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1))})))})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_Data_List_Types_applyList
}

var cache_Data_List_Types_applyNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_applyNonEmptyList sync.Once
func Get_Data_List_Types_applyNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_applyNonEmptyList.Do(func() {
		cache_Data_List_Types_applyNonEmptyList = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_functorNonEmptyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})})})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1))})))})))}})}
})
}))
	})
	return cache_Data_List_Types_applyNonEmptyList
}

var cache_Data_List_Types_bindList gopurs_runtime.Value
var once_Data_List_Types_bindList sync.Once
func Get_Data_List_Types_bindList() gopurs_runtime.Value {
	once_Data_List_Types_bindList.Do(func() {
		cache_Data_List_Types_bindList = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_applyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V0)))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_bindList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V1)}, v1_1)))})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_Data_List_Types_bindList
}

var cache_Data_List_Types_bindNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_bindNonEmptyList sync.Once
func Get_Data_List_Types_bindNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_bindNonEmptyList.Do(func() {
		cache_Data_List_Types_bindNonEmptyList = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_applyNonEmptyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_2_0 -> *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]
v1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0))
_ = v1_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v1_2_0)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v1_2_0)}.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_bindList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1))}, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(f_1, x_3)
_ = __local_var_4_1
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1)})}
}))))})))}})}
})
}))
	})
	return cache_Data_List_Types_bindNonEmptyList
}

var cache_Data_List_Types_applicativeList gopurs_runtime.Value
var once_Data_List_Types_applicativeList sync.Once
func Get_Data_List_Types_applicativeList() gopurs_runtime.Value {
	once_Data_List_Types_applicativeList.Do(func() {
		cache_Data_List_Types_applicativeList = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_applyList()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})})}
}))
	})
	return cache_Data_List_Types_applicativeList
}

var cache_Data_List_Types_monadList gopurs_runtime.Value
var once_Data_List_Types_monadList sync.Once
func Get_Data_List_Types_monadList() gopurs_runtime.Value {
	once_Data_List_Types_monadList.Do(func() {
		cache_Data_List_Types_monadList = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_applicativeList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_bindList()
}))
	})
	return cache_Data_List_Types_monadList
}

var cache_Data_List_Types_altNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_altNonEmptyList sync.Once
func Get_Data_List_Types_altNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_altNonEmptyList.Do(func() {
		cache_Data_List_Types_altNonEmptyList = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_functorNonEmptyList()
}), gopurs_runtime.RecordGet(Get_Data_List_Types_semigroupNonEmptyList(), "append"))
	})
	return cache_Data_List_Types_altNonEmptyList
}

var cache_Data_List_Types_altList gopurs_runtime.Value
var once_Data_List_Types_altList sync.Once
func Get_Data_List_Types_altList() gopurs_runtime.Value {
	once_Data_List_Types_altList.Do(func() {
		cache_Data_List_Types_altList = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_functorList()
}), gopurs_runtime.RecordGet(Get_Data_List_Types_semigroupList(), "append"))
	})
	return cache_Data_List_Types_altList
}

var cache_Data_List_Types_plusList gopurs_runtime.Value
var once_Data_List_Types_plusList sync.Once
func Get_Data_List_Types_plusList() gopurs_runtime.Value {
	once_Data_List_Types_plusList.Do(func() {
		cache_Data_List_Types_plusList = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_altList()
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
	})
	return cache_Data_List_Types_plusList
}

var cache_Data_List_Types_alternativeList gopurs_runtime.Value
var once_Data_List_Types_alternativeList sync.Once
func Get_Data_List_Types_alternativeList() gopurs_runtime.Value {
	once_Data_List_Types_alternativeList.Do(func() {
		cache_Data_List_Types_alternativeList = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_applicativeList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_plusList()
}))
	})
	return cache_Data_List_Types_alternativeList
}

var cache_Data_List_Types_monadPlusList gopurs_runtime.Value
var once_Data_List_Types_monadPlusList sync.Once
func Get_Data_List_Types_monadPlusList() gopurs_runtime.Value {
	once_Data_List_Types_monadPlusList.Do(func() {
		cache_Data_List_Types_monadPlusList = gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_alternativeList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_monadList()
}))
	})
	return cache_Data_List_Types_monadPlusList
}

var cache_Data_List_Types_applicativeNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_applicativeNonEmptyList sync.Once
func Get_Data_List_Types_applicativeNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_applicativeNonEmptyList.Do(func() {
		cache_Data_List_Types_applicativeNonEmptyList = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_applyNonEmptyList()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_0, gopurs_runtime.RecordGet(Get_Data_List_Types_plusList(), "empty")})}
}))
	})
	return cache_Data_List_Types_applicativeNonEmptyList
}

var cache_Data_List_Types_pure gopurs_runtime.Value
var once_Data_List_Types_pure sync.Once
func Get_Data_List_Types_pure() gopurs_runtime.Value {
	once_Data_List_Types_pure.Do(func() {
		cache_Data_List_Types_pure = gopurs_runtime.RecordGet(Get_Data_List_Types_applicativeNonEmptyList(), "pure")
	})
	return cache_Data_List_Types_pure
}

var cache_Data_List_Types_monadNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_monadNonEmptyList sync.Once
func Get_Data_List_Types_monadNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_monadNonEmptyList.Do(func() {
		cache_Data_List_Types_monadNonEmptyList = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_applicativeNonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_bindNonEmptyList()
}))
	})
	return cache_Data_List_Types_monadNonEmptyList
}

var cache_Data_List_Types_traversable1NonEmptyList gopurs_runtime.Value
var once_Data_List_Types_traversable1NonEmptyList sync.Once
func Get_Data_List_Types_traversable1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_traversable1NonEmptyList.Do(func() {
		cache_Data_List_Types_traversable1NonEmptyList = gopurs_runtime.RecordDict4("Foldable10", "Traversable1", "sequence1", "traverse1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_foldable1NonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_traversableNonEmptyList()
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_traversable1NonEmptyList(), "traverse1"), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0))}, Get_Data_List_Types_identity1())
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V1)})}})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0, gopurs_runtime.RecordGet(Get_Data_List_Types_plusList(), "empty")})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1))})))}
}), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_5_2 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_2
// TAST (Let): __local_var_5_1 -> gopurs_runtime.Value
__local_var_5_1 := gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_2.V0), gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_8, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_7.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_7.UnsafePtr).V1)})}})}
})
}), acc_4), b_6)
})
_ = __local_var_5_1
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_1, gopurs_runtime.Apply(f_2, x_6))
})
}), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.RecordGet(Get_Data_List_Types_applicativeNonEmptyList(), "pure"), gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))}))
})
})
}))
	})
	return cache_Data_List_Types_traversable1NonEmptyList
}

var cache_Data_List_Types_altList__614667287 gopurs_runtime.Value
var once_Data_List_Types_altList__614667287 sync.Once
func Get_Data_List_Types_altList__614667287() gopurs_runtime.Value {
	once_Data_List_Types_altList__614667287.Do(func() {
		cache_Data_List_Types_altList__614667287 = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_functorList()
}), gopurs_runtime.RecordGet(Get_Data_List_Types_semigroupList(), "append"))
	})
	return cache_Data_List_Types_altList__614667287
}

var cache_Data_List_Types_alternativeList__4245871901 gopurs_runtime.Value
var once_Data_List_Types_alternativeList__4245871901 sync.Once
func Get_Data_List_Types_alternativeList__4245871901() gopurs_runtime.Value {
	once_Data_List_Types_alternativeList__4245871901.Do(func() {
		cache_Data_List_Types_alternativeList__4245871901 = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_applicativeList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_plusList()
}))
	})
	return cache_Data_List_Types_alternativeList__4245871901
}

var cache_Data_List_Types_applicativeList__615687001 gopurs_runtime.Value
var once_Data_List_Types_applicativeList__615687001 sync.Once
func Get_Data_List_Types_applicativeList__615687001() gopurs_runtime.Value {
	once_Data_List_Types_applicativeList__615687001.Do(func() {
		cache_Data_List_Types_applicativeList__615687001 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_applyList()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})})}
}))
	})
	return cache_Data_List_Types_applicativeList__615687001
}

var cache_Data_List_Types_applicativeList__2027879016 gopurs_runtime.Value
var once_Data_List_Types_applicativeList__2027879016 sync.Once
func Get_Data_List_Types_applicativeList__2027879016() gopurs_runtime.Value {
	once_Data_List_Types_applicativeList__2027879016.Do(func() {
		cache_Data_List_Types_applicativeList__2027879016 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_applyList()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})})}
}))
	})
	return cache_Data_List_Types_applicativeList__2027879016
}

var cache_Data_List_Types_applicativeNonEmptyList__1156428081 gopurs_runtime.Value
var once_Data_List_Types_applicativeNonEmptyList__1156428081 sync.Once
func Get_Data_List_Types_applicativeNonEmptyList__1156428081() gopurs_runtime.Value {
	once_Data_List_Types_applicativeNonEmptyList__1156428081.Do(func() {
		cache_Data_List_Types_applicativeNonEmptyList__1156428081 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_applyNonEmptyList()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_0, gopurs_runtime.RecordGet(Get_Data_List_Types_plusList(), "empty")})}
}))
	})
	return cache_Data_List_Types_applicativeNonEmptyList__1156428081
}

var cache_Data_List_Types_applicativeNonEmptyList__3820246605 gopurs_runtime.Value
var once_Data_List_Types_applicativeNonEmptyList__3820246605 sync.Once
func Get_Data_List_Types_applicativeNonEmptyList__3820246605() gopurs_runtime.Value {
	once_Data_List_Types_applicativeNonEmptyList__3820246605.Do(func() {
		cache_Data_List_Types_applicativeNonEmptyList__3820246605 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_applyNonEmptyList()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_0, gopurs_runtime.RecordGet(Get_Data_List_Types_plusList(), "empty")})}
}))
	})
	return cache_Data_List_Types_applicativeNonEmptyList__3820246605
}

var cache_Data_List_Types_applicativeNonEmptyList__233027336 gopurs_runtime.Value
var once_Data_List_Types_applicativeNonEmptyList__233027336 sync.Once
func Get_Data_List_Types_applicativeNonEmptyList__233027336() gopurs_runtime.Value {
	once_Data_List_Types_applicativeNonEmptyList__233027336.Do(func() {
		cache_Data_List_Types_applicativeNonEmptyList__233027336 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_applyNonEmptyList()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_0, gopurs_runtime.RecordGet(Get_Data_List_Types_plusList(), "empty")})}
}))
	})
	return cache_Data_List_Types_applicativeNonEmptyList__233027336
}

var cache_Data_List_Types_applyList__3072763993 gopurs_runtime.Value
var once_Data_List_Types_applyList__3072763993 sync.Once
func Get_Data_List_Types_applyList__3072763993() gopurs_runtime.Value {
	once_Data_List_Types_applyList__3072763993.Do(func() {
		cache_Data_List_Types_applyList__3072763993 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_functorList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_functorList(), "map"), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1))})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1))})))})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_Data_List_Types_applyList__3072763993
}

var cache_Data_List_Types_applyList__1109325167 gopurs_runtime.Value
var once_Data_List_Types_applyList__1109325167 sync.Once
func Get_Data_List_Types_applyList__1109325167() gopurs_runtime.Value {
	once_Data_List_Types_applyList__1109325167.Do(func() {
		cache_Data_List_Types_applyList__1109325167 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_functorList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_functorList(), "map"), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1))})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1))})))})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_Data_List_Types_applyList__1109325167
}

var cache_Data_List_Types_applyNonEmptyList__602103086 gopurs_runtime.Value
var once_Data_List_Types_applyNonEmptyList__602103086 sync.Once
func Get_Data_List_Types_applyNonEmptyList__602103086() gopurs_runtime.Value {
	once_Data_List_Types_applyNonEmptyList__602103086.Do(func() {
		cache_Data_List_Types_applyNonEmptyList__602103086 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_functorNonEmptyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})})})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1))})))})))}})}
})
}))
	})
	return cache_Data_List_Types_applyNonEmptyList__602103086
}

var cache_Data_List_Types_bindList__241263065 gopurs_runtime.Value
var once_Data_List_Types_bindList__241263065 sync.Once
func Get_Data_List_Types_bindList__241263065() gopurs_runtime.Value {
	once_Data_List_Types_bindList__241263065.Do(func() {
		cache_Data_List_Types_bindList__241263065 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_applyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V0)))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_bindList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V1)}, v1_1)))})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_Data_List_Types_bindList__241263065
}

var cache_Data_List_Types_bindList__3903132455 gopurs_runtime.Value
var once_Data_List_Types_bindList__3903132455 sync.Once
func Get_Data_List_Types_bindList__3903132455() gopurs_runtime.Value {
	once_Data_List_Types_bindList__3903132455.Do(func() {
		cache_Data_List_Types_bindList__3903132455 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_applyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V0)))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_bindList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V1)}, v1_1)))})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_Data_List_Types_bindList__3903132455
}

var cache_Data_List_Types_bindNonEmptyList__1408886065 gopurs_runtime.Value
var once_Data_List_Types_bindNonEmptyList__1408886065 sync.Once
func Get_Data_List_Types_bindNonEmptyList__1408886065() gopurs_runtime.Value {
	once_Data_List_Types_bindNonEmptyList__1408886065.Do(func() {
		cache_Data_List_Types_bindNonEmptyList__1408886065 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_applyNonEmptyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_2_0 -> *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]
v1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0))
_ = v1_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v1_2_0)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v1_2_0)}.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_bindList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1))}, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(f_1, x_3)
_ = __local_var_4_1
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1)})}
}))))})))}})}
})
}))
	})
	return cache_Data_List_Types_bindNonEmptyList__1408886065
}

var cache_Data_List_Types_bindNonEmptyList__1309464679 gopurs_runtime.Value
var once_Data_List_Types_bindNonEmptyList__1309464679 sync.Once
func Get_Data_List_Types_bindNonEmptyList__1309464679() gopurs_runtime.Value {
	once_Data_List_Types_bindNonEmptyList__1309464679.Do(func() {
		cache_Data_List_Types_bindNonEmptyList__1309464679 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_applyNonEmptyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_2_0 -> *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]
v1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0))
_ = v1_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v1_2_0)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v1_2_0)}.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_bindList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1))}, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(f_1, x_3)
_ = __local_var_4_1
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1)})}
}))))})))}})}
})
}))
	})
	return cache_Data_List_Types_bindNonEmptyList__1309464679
}

var cache_Data_List_Types_eq1List__1109645470 gopurs_runtime.Value
var once_Data_List_Types_eq1List__1109645470 sync.Once
func Get_Data_List_Types_eq1List__1109645470() gopurs_runtime.Value {
	once_Data_List_Types_eq1List__1109645470.Do(func() {
		cache_Data_List_Types_eq1List__1109645470 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_0_8 gopurs_runtime.Value
go__go_3_0_8 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
var v2_6_loop bool = (v2_6_loop_val.IntVal) != (0)
go__go_3_0_8:
for {
if false { continue go__go_3_0_8 }
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
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
v2_6_loop = (v2_6) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0).IntVal) != (0))
continue go__go_3_0_8
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
return gopurs_runtime.Bool((gopurs_runtime.Apply3(go__go_3_0_8, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_2))}, gopurs_runtime.Bool(true)).IntVal) != (0))
})
})
}))
	})
	return cache_Data_List_Types_eq1List__1109645470
}

var cache_Data_List_Types_extendNonEmptyList__2163561628 gopurs_runtime.Value
var once_Data_List_Types_extendNonEmptyList__2163561628 sync.Once
func Get_Data_List_Types_extendNonEmptyList__2163561628() gopurs_runtime.Value {
	once_Data_List_Types_extendNonEmptyList__2163561628.Do(func() {
		cache_Data_List_Types_extendNonEmptyList__2163561628 = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_functorNonEmptyList()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldr"), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v1_3, "acc"))})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_2, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v1_3, "acc")))}})}), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v1_3, "val"))})})
})
}), gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1))}), "val")))}})}
})
}))
	})
	return cache_Data_List_Types_extendNonEmptyList__2163561628
}

var cache_Data_List_Types_foldable1NonEmptyList__2239557029 gopurs_runtime.Value
var once_Data_List_Types_foldable1NonEmptyList__2239557029 sync.Once
func Get_Data_List_Types_foldable1NonEmptyList__2239557029() gopurs_runtime.Value {
	once_Data_List_Types_foldable1NonEmptyList__2239557029.Do(func() {
		cache_Data_List_Types_foldable1NonEmptyList__2239557029 = gopurs_runtime.Apply(Get_Data_NonEmpty_foldable1NonEmpty(), Get_Data_List_Types_foldableList())
	})
	return cache_Data_List_Types_foldable1NonEmptyList__2239557029
}

var cache_Data_List_Types_foldable1NonEmptyList__2630190169 gopurs_runtime.Value
var once_Data_List_Types_foldable1NonEmptyList__2630190169 sync.Once
func Get_Data_List_Types_foldable1NonEmptyList__2630190169() gopurs_runtime.Value {
	once_Data_List_Types_foldable1NonEmptyList__2630190169.Do(func() {
		cache_Data_List_Types_foldable1NonEmptyList__2630190169 = gopurs_runtime.Apply(Get_Data_NonEmpty_foldable1NonEmpty(), Get_Data_List_Types_foldableList())
	})
	return cache_Data_List_Types_foldable1NonEmptyList__2630190169
}

var cache_Data_List_Types_foldable1NonEmptyList__1746670655 gopurs_runtime.Value
var once_Data_List_Types_foldable1NonEmptyList__1746670655 sync.Once
func Get_Data_List_Types_foldable1NonEmptyList__1746670655() gopurs_runtime.Value {
	once_Data_List_Types_foldable1NonEmptyList__1746670655.Do(func() {
		cache_Data_List_Types_foldable1NonEmptyList__1746670655 = gopurs_runtime.Apply(Get_Data_NonEmpty_foldable1NonEmpty(), Get_Data_List_Types_foldableList())
	})
	return cache_Data_List_Types_foldable1NonEmptyList__1746670655
}

var cache_Data_List_Types_foldableList__1753400174 gopurs_runtime.Value
var once_Data_List_Types_foldableList__1753400174 sync.Once
func Get_Data_List_Types_foldableList__1753400174() gopurs_runtime.Value {
	once_Data_List_Types_foldableList__1753400174.Do(func() {
		cache_Data_List_Types_foldableList__1753400174 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Box(Semigroup0_1_0.V0), acc_4)
_ = __local_var_5_2
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, gopurs_runtime.Apply(f_3, x_6))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_3_9 gopurs_runtime.Value
go__go_1_3_9 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_3_9:
for {
if false { continue go__go_1_3_9 }
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
b_2_loop = gopurs_runtime.Apply2(f_0, b_2, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
continue go__go_1_3_9
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
return go__go_1_3_9
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_5 -> gopurs_runtime.Value
__local_var_2_5 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
})
}), b_1)
_ = __local_var_2_5
var go__go_3_7_10 gopurs_runtime.Value
go__go_3_7_10 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_7_10:
for {
if false { continue go__go_3_7_10 }
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t8 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}
goto end_branch_8
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
continue go__go_3_7_10
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__t8))}
}
}()
})
})
// TAST (Let): __local_var_3_6 -> gopurs_runtime.Value
__local_var_3_6 := gopurs_runtime.Apply(go__go_3_7_10, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_3_6
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_5, gopurs_runtime.Apply(__local_var_3_6, x_4))
})
})
}))
	})
	return cache_Data_List_Types_foldableList__1753400174
}

var cache_Data_List_Types_foldableList__46095397 gopurs_runtime.Value
var once_Data_List_Types_foldableList__46095397 sync.Once
func Get_Data_List_Types_foldableList__46095397() gopurs_runtime.Value {
	once_Data_List_Types_foldableList__46095397.Do(func() {
		cache_Data_List_Types_foldableList__46095397 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Box(Semigroup0_1_0.V0), acc_4)
_ = __local_var_5_2
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, gopurs_runtime.Apply(f_3, x_6))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_3_11 gopurs_runtime.Value
go__go_1_3_11 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_3_11:
for {
if false { continue go__go_1_3_11 }
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
b_2_loop = gopurs_runtime.Apply2(f_0, b_2, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
continue go__go_1_3_11
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
return go__go_1_3_11
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_5 -> gopurs_runtime.Value
__local_var_2_5 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
})
}), b_1)
_ = __local_var_2_5
var go__go_3_7_12 gopurs_runtime.Value
go__go_3_7_12 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_7_12:
for {
if false { continue go__go_3_7_12 }
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t8 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}
goto end_branch_8
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
continue go__go_3_7_12
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__t8))}
}
}()
})
})
// TAST (Let): __local_var_3_6 -> gopurs_runtime.Value
__local_var_3_6 := gopurs_runtime.Apply(go__go_3_7_12, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_3_6
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_5, gopurs_runtime.Apply(__local_var_3_6, x_4))
})
})
}))
	})
	return cache_Data_List_Types_foldableList__46095397
}

var cache_Data_List_Types_foldableNonEmptyList__3933232868 gopurs_runtime.Value
var once_Data_List_Types_foldableNonEmptyList__3933232868 sync.Once
func Get_Data_List_Types_foldableNonEmptyList__3933232868() gopurs_runtime.Value {
	once_Data_List_Types_foldableNonEmptyList__3933232868.Do(func() {
		cache_Data_List_Types_foldableNonEmptyList__3933232868 = gopurs_runtime.Apply(Get_Data_NonEmpty_foldableNonEmpty(), Get_Data_List_Types_foldableList())
	})
	return cache_Data_List_Types_foldableNonEmptyList__3933232868
}

var cache_Data_List_Types_foldableWithIndexList__3168570882 gopurs_runtime.Value
var once_Data_List_Types_foldableWithIndexList__3168570882 sync.Once
func Get_Data_List_Types_foldableWithIndexList__3168570882() gopurs_runtime.Value {
	once_Data_List_Types_foldableWithIndexList__3168570882.Do(func() {
		cache_Data_List_Types_foldableWithIndexList__3168570882 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_foldableList()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableWithIndexList(), "foldlWithIndex"), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
__local_var_2_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_5 -> gopurs_runtime.Value
__local_var_3_5 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_3_5
// TAST (Let): __local_var_4_6 -> gopurs_runtime.Value
__local_var_4_6 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_4_6
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_4_6.IntVal) + (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_4_6.IntVal), __local_var_3_5, a_5)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), acc_1})})
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(__local_var_2_4, x_3).UnsafePtr).V1
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_3_7 -> *Constructor_Data_Tuple_Tuple[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]
v_3_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_8 -> gopurs_runtime.Value
__local_var_4_8 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_4_8
// TAST (Let): __local_var_5_9 -> gopurs_runtime.Value
__local_var_5_9 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0
_ = __local_var_5_9
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_5_9.IntVal) + (1)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_6, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__local_var_4_8)})}})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_2))}))
_ = v_3_7
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_10 -> gopurs_runtime.Value
__local_var_5_10 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_10
// TAST (Let): __local_var_6_11 -> gopurs_runtime.Value
__local_var_6_11 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0
_ = __local_var_6_11
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_6_11.IntVal) - (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((__local_var_6_11.IntVal) - (1)), a_7, __local_var_5_10)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_7)}.UnsafePtr).V0.IntVal), b_1})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_7)}.UnsafePtr).V1))}).UnsafePtr).V1
})
})
}))
	})
	return cache_Data_List_Types_foldableWithIndexList__3168570882
}

var cache_Data_List_Types_foldableWithIndexList__184979230 gopurs_runtime.Value
var once_Data_List_Types_foldableWithIndexList__184979230 sync.Once
func Get_Data_List_Types_foldableWithIndexList__184979230() gopurs_runtime.Value {
	once_Data_List_Types_foldableWithIndexList__184979230.Do(func() {
		cache_Data_List_Types_foldableWithIndexList__184979230 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_foldableList()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableWithIndexList(), "foldlWithIndex"), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
__local_var_2_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_5 -> gopurs_runtime.Value
__local_var_3_5 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_3_5
// TAST (Let): __local_var_4_6 -> gopurs_runtime.Value
__local_var_4_6 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_4_6
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_4_6.IntVal) + (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_4_6.IntVal), __local_var_3_5, a_5)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), acc_1})})
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(__local_var_2_4, x_3).UnsafePtr).V1
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_3_7 -> *Constructor_Data_Tuple_Tuple[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]
v_3_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_8 -> gopurs_runtime.Value
__local_var_4_8 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_4_8
// TAST (Let): __local_var_5_9 -> gopurs_runtime.Value
__local_var_5_9 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0
_ = __local_var_5_9
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_5_9.IntVal) + (1)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_6, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__local_var_4_8)})}})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_2))}))
_ = v_3_7
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_10 -> gopurs_runtime.Value
__local_var_5_10 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_10
// TAST (Let): __local_var_6_11 -> gopurs_runtime.Value
__local_var_6_11 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0
_ = __local_var_6_11
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_6_11.IntVal) - (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((__local_var_6_11.IntVal) - (1)), a_7, __local_var_5_10)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_7)}.UnsafePtr).V0.IntVal), b_1})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_7)}.UnsafePtr).V1))}).UnsafePtr).V1
})
})
}))
	})
	return cache_Data_List_Types_foldableWithIndexList__184979230
}

var cache_Data_List_Types_foldableWithIndexNonEmptyList__3425888830 gopurs_runtime.Value
var once_Data_List_Types_foldableWithIndexNonEmptyList__3425888830 sync.Once
func Get_Data_List_Types_foldableWithIndexNonEmptyList__3425888830() gopurs_runtime.Value {
	once_Data_List_Types_foldableWithIndexNonEmptyList__3425888830.Do(func() {
		cache_Data_List_Types_foldableWithIndexNonEmptyList__3425888830 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_foldableNonEmptyList()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_foldableWithIndexNonEmpty()).V1), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))}, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Int((1) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0.IntVal))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Apply(f_1, gopurs_runtime.Int(__t0.IntVal))
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_foldableWithIndexNonEmpty()).V2), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Int(0)
goto end_branch_1
} else {

}
}
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr != nil) {
__t1 = gopurs_runtime.Int((1) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0.IntVal))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Apply(f_0, gopurs_runtime.Int(__t1.IntVal))
}), b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_foldableWithIndexNonEmpty()).V3), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr == nil) {
__t2 = gopurs_runtime.Int(0)
goto end_branch_2
} else {

}
}
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr != nil) {
__t2 = gopurs_runtime.Int((1) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0.IntVal))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Apply(f_0, gopurs_runtime.Int(__t2.IntVal))
}), b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2))})
})
})
}))
	})
	return cache_Data_List_Types_foldableWithIndexNonEmptyList__3425888830
}

var cache_Data_List_Types_functorList__4121998062 gopurs_runtime.Value
var once_Data_List_Types_functorList__4121998062 sync.Once
func Get_Data_List_Types_functorList__4121998062() gopurs_runtime.Value {
	once_Data_List_Types_functorList__4121998062.Do(func() {
		cache_Data_List_Types_functorList__4121998062 = gopurs_runtime.RecordDict1("map", Get_Data_List_Types_listMap())
	})
	return cache_Data_List_Types_functorList__4121998062
}

var cache_Data_List_Types_functorList__1783129585 gopurs_runtime.Value
var once_Data_List_Types_functorList__1783129585 sync.Once
func Get_Data_List_Types_functorList__1783129585() gopurs_runtime.Value {
	once_Data_List_Types_functorList__1783129585.Do(func() {
		cache_Data_List_Types_functorList__1783129585 = gopurs_runtime.RecordDict1("map", Get_Data_List_Types_listMap())
	})
	return cache_Data_List_Types_functorList__1783129585
}

var cache_Data_List_Types_functorNonEmptyList__2834508934 gopurs_runtime.Value
var once_Data_List_Types_functorNonEmptyList__2834508934 sync.Once
func Get_Data_List_Types_functorNonEmptyList__2834508934() gopurs_runtime.Value {
	once_Data_List_Types_functorNonEmptyList__2834508934.Do(func() {
		cache_Data_List_Types_functorNonEmptyList__2834508934 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_functorList(), "map"), f_0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_Data_List_Types_functorNonEmptyList__2834508934
}

var cache_Data_List_Types_functorNonEmptyList__1593940346 gopurs_runtime.Value
var once_Data_List_Types_functorNonEmptyList__1593940346 sync.Once
func Get_Data_List_Types_functorNonEmptyList__1593940346() gopurs_runtime.Value {
	once_Data_List_Types_functorNonEmptyList__1593940346.Do(func() {
		cache_Data_List_Types_functorNonEmptyList__1593940346 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_functorList(), "map"), f_0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_Data_List_Types_functorNonEmptyList__1593940346
}

var cache_Data_List_Types_functorNonEmptyList__257963697 gopurs_runtime.Value
var once_Data_List_Types_functorNonEmptyList__257963697 sync.Once
func Get_Data_List_Types_functorNonEmptyList__257963697() gopurs_runtime.Value {
	once_Data_List_Types_functorNonEmptyList__257963697.Do(func() {
		cache_Data_List_Types_functorNonEmptyList__257963697 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_functorList(), "map"), f_0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_Data_List_Types_functorNonEmptyList__257963697
}

var cache_Data_List_Types_functorWithIndexList__1995002722 gopurs_runtime.Value
var once_Data_List_Types_functorWithIndexList__1995002722 sync.Once
func Get_Data_List_Types_functorWithIndexList__1995002722() gopurs_runtime.Value {
	once_Data_List_Types_functorWithIndexList__1995002722.Do(func() {
		cache_Data_List_Types_functorWithIndexList__1995002722 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_functorList()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableWithIndexList(), "foldrWithIndex"), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, gopurs_runtime.Int(i_1.IntVal), x_2), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](acc_3)})}
})
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
}))
	})
	return cache_Data_List_Types_functorWithIndexList__1995002722
}

var cache_Data_List_Types_functorWithIndexNonEmptyList__3683208290 gopurs_runtime.Value
var once_Data_List_Types_functorWithIndexNonEmptyList__3683208290 sync.Once
func Get_Data_List_Types_functorWithIndexNonEmptyList__3683208290() gopurs_runtime.Value {
	once_Data_List_Types_functorWithIndexNonEmptyList__3683208290.Do(func() {
		cache_Data_List_Types_functorWithIndexNonEmptyList__3683208290 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_functorNonEmptyList()
}), gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_functorWithIndex()).V1), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_2.Type == 9 && x_2.IntVal == 930809136 && x_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (x_2.Type == 9 && x_2.IntVal == 930809136 && x_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Int((1) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_2.UnsafePtr).V0.IntVal))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Apply(fn_0, gopurs_runtime.Int(__t0.IntVal))
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1))})))}
})
}))
	})
	return cache_Data_List_Types_functorWithIndexNonEmptyList__3683208290
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

var cache_Data_List_Types_monadList__4139186259 gopurs_runtime.Value
var once_Data_List_Types_monadList__4139186259 sync.Once
func Get_Data_List_Types_monadList__4139186259() gopurs_runtime.Value {
	once_Data_List_Types_monadList__4139186259.Do(func() {
		cache_Data_List_Types_monadList__4139186259 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_applicativeList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_bindList()
}))
	})
	return cache_Data_List_Types_monadList__4139186259
}

var cache_Data_List_Types_nelCons__195558898 gopurs_runtime.Value
var once_Data_List_Types_nelCons__195558898 sync.Once
func Get_Data_List_Types_nelCons__195558898() gopurs_runtime.Value {
	once_Data_List_Types_nelCons__195558898.Do(func() {
		cache_Data_List_Types_nelCons__195558898 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_Types_nelCons__195558898(a_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box)))}
})
	})
	return cache_Data_List_Types_nelCons__195558898
}

var cache_Data_List_Types_nelCons__2148523118 gopurs_runtime.Value
var once_Data_List_Types_nelCons__2148523118 sync.Once
func Get_Data_List_Types_nelCons__2148523118() gopurs_runtime.Value {
	once_Data_List_Types_nelCons__2148523118.Do(func() {
		cache_Data_List_Types_nelCons__2148523118 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_List_Types_nelCons__2148523118(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](a_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](v_1_box)))}
})
	})
	return cache_Data_List_Types_nelCons__2148523118
}

var cache_Data_List_Types_plusList__598824825 gopurs_runtime.Value
var once_Data_List_Types_plusList__598824825 sync.Once
func Get_Data_List_Types_plusList__598824825() gopurs_runtime.Value {
	once_Data_List_Types_plusList__598824825.Do(func() {
		cache_Data_List_Types_plusList__598824825 = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_altList()
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
	})
	return cache_Data_List_Types_plusList__598824825
}

var cache_Data_List_Types_plusList__2201439314 gopurs_runtime.Value
var once_Data_List_Types_plusList__2201439314 sync.Once
func Get_Data_List_Types_plusList__2201439314() gopurs_runtime.Value {
	once_Data_List_Types_plusList__2201439314.Do(func() {
		cache_Data_List_Types_plusList__2201439314 = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_altList()
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
	})
	return cache_Data_List_Types_plusList__2201439314
}

var cache_Data_List_Types_semigroupList__2766094215 gopurs_runtime.Value
var once_Data_List_Types_semigroupList__2766094215 sync.Once
func Get_Data_List_Types_semigroupList__2766094215() gopurs_runtime.Value {
	once_Data_List_Types_semigroupList__2766094215.Do(func() {
		cache_Data_List_Types_semigroupList__2766094215 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldr"), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_0))})))}
})
}))
	})
	return cache_Data_List_Types_semigroupList__2766094215
}

var cache_Data_List_Types_semigroupList__3527039931 gopurs_runtime.Value
var once_Data_List_Types_semigroupList__3527039931 sync.Once
func Get_Data_List_Types_semigroupList__3527039931() gopurs_runtime.Value {
	once_Data_List_Types_semigroupList__3527039931.Do(func() {
		cache_Data_List_Types_semigroupList__3527039931 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldr"), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_0))})))}
})
}))
	})
	return cache_Data_List_Types_semigroupList__3527039931
}

var cache_Data_List_Types_semigroupList__3410686552 gopurs_runtime.Value
var once_Data_List_Types_semigroupList__3410686552 sync.Once
func Get_Data_List_Types_semigroupList__3410686552() gopurs_runtime.Value {
	once_Data_List_Types_semigroupList__3410686552.Do(func() {
		cache_Data_List_Types_semigroupList__3410686552 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldr"), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_0))})))}
})
}))
	})
	return cache_Data_List_Types_semigroupList__3410686552
}

var cache_Data_List_Types_toList__2859885498 gopurs_runtime.Value
var once_Data_List_Types_toList__2859885498 sync.Once
func Get_Data_List_Types_toList__2859885498() gopurs_runtime.Value {
	once_Data_List_Types_toList__2859885498.Do(func() {
		cache_Data_List_Types_toList__2859885498 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_Types_toList__2859885498(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_Data_List_Types_toList__2859885498
}

var cache_Data_List_Types_toList__1324737658 gopurs_runtime.Value
var once_Data_List_Types_toList__1324737658 sync.Once
func Get_Data_List_Types_toList__1324737658() gopurs_runtime.Value {
	once_Data_List_Types_toList__1324737658.Do(func() {
		cache_Data_List_Types_toList__1324737658 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_Types_toList__1324737658(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](v_0_box)))}
})
	})
	return cache_Data_List_Types_toList__1324737658
}

var cache_Data_List_Types_toList__2402503393 gopurs_runtime.Value
var once_Data_List_Types_toList__2402503393 sync.Once
func Get_Data_List_Types_toList__2402503393() gopurs_runtime.Value {
	once_Data_List_Types_toList__2402503393.Do(func() {
		cache_Data_List_Types_toList__2402503393 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_Types_toList__2402503393(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_Data_List_Types_toList__2402503393
}

var cache_Data_List_Types_traversable1NonEmptyList__1171985061 gopurs_runtime.Value
var once_Data_List_Types_traversable1NonEmptyList__1171985061 sync.Once
func Get_Data_List_Types_traversable1NonEmptyList__1171985061() gopurs_runtime.Value {
	once_Data_List_Types_traversable1NonEmptyList__1171985061.Do(func() {
		cache_Data_List_Types_traversable1NonEmptyList__1171985061 = gopurs_runtime.RecordDict4("Foldable10", "Traversable1", "sequence1", "traverse1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_foldable1NonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_traversableNonEmptyList()
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_traversable1NonEmptyList(), "traverse1"), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0))}, Get_Data_List_Types_identity1())
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V1)})}})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0, gopurs_runtime.RecordGet(Get_Data_List_Types_plusList(), "empty")})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1))})))}
}), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_5_2 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_2
// TAST (Let): __local_var_5_1 -> gopurs_runtime.Value
__local_var_5_1 := gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_2.V0), gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_8, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_7.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_7.UnsafePtr).V1)})}})}
})
}), acc_4), b_6)
})
_ = __local_var_5_1
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_1, gopurs_runtime.Apply(f_2, x_6))
})
}), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.RecordGet(Get_Data_List_Types_applicativeNonEmptyList(), "pure"), gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))}))
})
})
}))
	})
	return cache_Data_List_Types_traversable1NonEmptyList__1171985061
}

var cache_Data_List_Types_traversableList__3361437934 gopurs_runtime.Value
var once_Data_List_Types_traversableList__3361437934 sync.Once
func Get_Data_List_Types_traversableList__3361437934() gopurs_runtime.Value {
	once_Data_List_Types_traversableList__3361437934.Do(func() {
		cache_Data_List_Types_traversableList__3361437934 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_foldableList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_functorList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_traversableList(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, Get_Data_List_Types_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): Apply0_2_1 -> *Constructor_Control_Apply_Apply[gopurs_runtime.Value]
Apply0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_5, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_4)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}))
_ = __local_var_4_2
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_6_5 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_2_1.V0), gopurs_runtime.Value{}))
_ = Functor0_6_5
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_5.V0), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_9, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_8)})}
})
}), acc_5), b_7)
})
_ = __local_var_6_4
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, gopurs_runtime.Apply(f_3, x_7))
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}))
_ = __local_var_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_5_3, x_6))
})
})
}))
	})
	return cache_Data_List_Types_traversableList__3361437934
}

var cache_Data_List_Types_traversableList__365792206 gopurs_runtime.Value
var once_Data_List_Types_traversableList__365792206 sync.Once
func Get_Data_List_Types_traversableList__365792206() gopurs_runtime.Value {
	once_Data_List_Types_traversableList__365792206.Do(func() {
		cache_Data_List_Types_traversableList__365792206 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_foldableList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_functorList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_traversableList(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, Get_Data_List_Types_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): Apply0_2_1 -> *Constructor_Control_Apply_Apply[gopurs_runtime.Value]
Apply0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_5, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_4)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}))
_ = __local_var_4_2
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_6_5 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_2_1.V0), gopurs_runtime.Value{}))
_ = Functor0_6_5
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_5.V0), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_9, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_8)})}
})
}), acc_5), b_7)
})
_ = __local_var_6_4
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, gopurs_runtime.Apply(f_3, x_7))
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}))
_ = __local_var_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_5_3, x_6))
})
})
}))
	})
	return cache_Data_List_Types_traversableList__365792206
}

var cache_Data_List_Types_traversableNonEmptyList__1085933743 gopurs_runtime.Value
var once_Data_List_Types_traversableNonEmptyList__1085933743 sync.Once
func Get_Data_List_Types_traversableNonEmptyList__1085933743() gopurs_runtime.Value {
	once_Data_List_Types_traversableNonEmptyList__1085933743.Do(func() {
		cache_Data_List_Types_traversableNonEmptyList__1085933743 = gopurs_runtime.Apply(Get_Data_NonEmpty_traversableNonEmpty(), Get_Data_List_Types_traversableList())
	})
	return cache_Data_List_Types_traversableNonEmptyList__1085933743
}

var cache_Data_List_Types_unfoldable1List__3672302568 gopurs_runtime.Value
var once_Data_List_Types_unfoldable1List__3672302568 sync.Once
func Get_Data_List_Types_unfoldable1List__3672302568() gopurs_runtime.Value {
	once_Data_List_Types_unfoldable1List__3672302568.Do(func() {
		cache_Data_List_Types_unfoldable1List__3672302568 = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_17 gopurs_runtime.Value
go__go_2_0_17 = gopurs_runtime.Func(func(source_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var source_3_loop gopurs_runtime.Value = source_3_loop_val
var memo_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](memo_4_loop_val)
go__go_2_0_17:
for {
if false { continue go__go_2_0_17 }
var source_3 gopurs_runtime.Value = source_3_loop
_ = source_3
var memo_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = memo_4_loop
_ = memo_4
// TAST (Let): v_5_1 -> *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]
v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](gopurs_runtime.Apply(f_0, source_3))
_ = v_5_1
var __t4 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 930809136 && __t_tag_2.UnsafePtr != nil) {
source_3_loop = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V1.UnsafePtr).V0
memo_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V0, memo_4})})
continue go__go_2_0_17
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V1
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 930809136 && __t_tag_3.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_7, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_6)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V0, memo_4})})))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__t4))}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_2_0_17, b_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})))}
})
}))
	})
	return cache_Data_List_Types_unfoldable1List__3672302568
}

var cache_Data_List_Types_unfoldableList__2633941518 gopurs_runtime.Value
var once_Data_List_Types_unfoldableList__2633941518 sync.Once
func Get_Data_List_Types_unfoldableList__2633941518() gopurs_runtime.Value {
	once_Data_List_Types_unfoldableList__2633941518.Do(func() {
		cache_Data_List_Types_unfoldableList__2633941518 = gopurs_runtime.RecordDict2("Unfoldable10", "unfoldr", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Types_unfoldable1List()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_18 gopurs_runtime.Value
go__go_2_0_18 = gopurs_runtime.Func(func(source_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var source_3_loop gopurs_runtime.Value = source_3_loop_val
var memo_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](memo_4_loop_val)
go__go_2_0_18:
for {
if false { continue go__go_2_0_18 }
var source_3 gopurs_runtime.Value = source_3_loop
_ = source_3
var memo_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = memo_4_loop
_ = memo_4
// TAST (Let): v_5_1 -> *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]
v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(f_0, source_3))
_ = v_5_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_7, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_6)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(memo_4)})))}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr != nil) {
source_3_loop = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V0.UnsafePtr).V1
memo_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V0.UnsafePtr).V0, memo_4})})
continue go__go_2_0_18
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__t2))}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_2_0_18, b_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})))}
})
}))
	})
	return cache_Data_List_Types_unfoldableList__2633941518
}

type Constructor_Data_List_Types_Nil[T_a any] struct {
	Rc uint32
}


type Constructor_Data_List_Types_Cons[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
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

func Call_Data_List_Types_toList(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)})})
}

func Call_Data_List_Types_nelCons(a_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)})}})})
}

func Call_Data_List_Types_listMap(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var chunkedRevMap_1_0_0 gopurs_runtime.Value
chunkedRevMap_1_0_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](v_2_loop_val)
var v1_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val)
chunkedRevMap_1_0_0:
for {
if false { continue chunkedRevMap_1_0_0 }
var v_2 *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t19 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_and_18 bool = false
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr != nil) {

var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
var __t_and_17 bool = false
if (__t_tag_15.Type == 9 && __t_tag_15.IntVal == 1358893437 && __t_tag_15.UnsafePtr != nil) {

var __t_tag_16 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}
__t_and_17 = (__t_tag_16.Type == 9 && __t_tag_16.IntVal == 1358893437 && __t_tag_16.UnsafePtr != nil)
}
__t_and_18 = __t_and_17
}
if __t_and_18 {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)})})})
v1_3_loop = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr).V1
continue chunkedRevMap_1_0_0
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
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
var v3_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v3_6_loop_val)
reverseUnrolledMap_4_1_1:
for {
if false { continue reverseUnrolledMap_4_1_1 }
var v2_5 gopurs_runtime.Value = v2_5_loop
_ = v2_5
var v3_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v3_6_loop
_ = v3_6
var __t8 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_and_7 bool = false
if (v2_5.Type == 9 && v2_5.IntVal == 1358893437 && v2_5.UnsafePtr != nil) {

var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0
var __t_and_6 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1358893437 && __t_tag_2.UnsafePtr != nil) {

var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}
var __t_and_5 bool = false
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 1358893437 && __t_tag_3.UnsafePtr != nil) {

var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V1)}
__t_and_5 = (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 1358893437 && __t_tag_4.UnsafePtr != nil)
}
__t_and_6 = __t_and_5
}
__t_and_7 = __t_and_6
}
if __t_and_7 {
v2_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V1)}
v3_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr).V0), v3_6})})})})})})
continue reverseUnrolledMap_4_1_1
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
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
var __t14 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr != nil) {
var __t13 gopurs_runtime.Value
{
var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 1358893437 && __t_tag_9.UnsafePtr != nil) {
var __t11 gopurs_runtime.Value
{
var __t_tag_10 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}
if (__t_tag_10.Type == 9 && __t_tag_10.IntVal == 1358893437 && __t_tag_10.UnsafePtr == nil) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})})})})}
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_11:
__t13 = __t11
goto end_branch_13
} else {

}
}
{
var __t_tag_12 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 1358893437 && __t_tag_12.UnsafePtr == nil) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})})}
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_13:
__t14 = __t13
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_14:
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(reverseUnrolledMap_4_1_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__t14))}))
}
end_branch_19:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t19)}
}
}()
})
})
return gopurs_runtime.Apply(chunkedRevMap_1_0_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
}

func Call_Data_List_Types_showList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): show_1_0 -> gopurs_runtime.Value
show_1_0 := gopurs_runtime.RecordGet(dictShow_0, "show")
_ = show_1_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 string
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr == nil) {
__t1 = "Nil"
goto end_branch_1
} else {

}
}
{
__t1 = (("(") + (gopurs_runtime.Apply2(Get_Data_Foldable_intercalate__2937349250(), gopurs_runtime.Str(" : "), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[string]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_functorList(), "map"), show_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2))})))}).StrVal())) + (" : Nil)")
}
end_branch_1:
return gopurs_runtime.Str(__t1)
}))
}

func Call_Data_List_Types_showNonEmptyList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := Call_Data_List_Types_showList(dictShow_0)
_ = __local_var_1_1
// TAST (Let): showNonEmpty_1_0 -> *Constructor_Data_Show_Show[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]
showNonEmpty_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((((("(NonEmpty ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "show"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal())) + (")"))
})))
_ = showNonEmpty_1_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(NonEmptyList ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showNonEmpty_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2))}).StrVal())) + (")"))
}))
}

func Call_Data_List_Types_eqNonEmpty(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_eq1List(), "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0))}, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V1).IntVal) != (0)))
})
}))
}

func Call_Data_List_Types_eqList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Types_eq1List(), "eq1"), dictEq_0))
}

func Call_Data_List_Types_eqNonEmptyList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return Call_Data_List_Types_eqNonEmpty(dictEq_0)
}

func Call_Data_List_Types_ordList(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): eqList1_1_0 -> gopurs_runtime.Value
eqList1_1_0 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Types_eq1List(), "eq1"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})))
_ = eqList1_1_0
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return eqList1_1_0
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Types_ord1List(), "compare1"), dictOrd_0))
}

func Call_Data_List_Types_ordNonEmptyList(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_List_Types_ordNonEmpty(), dictOrd_0)
}

func Call_Data_List_Types_listMap__858544730(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var chunkedRevMap_1_0_13 gopurs_runtime.Value
chunkedRevMap_1_0_13 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](v_2_loop_val)
var v1_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val)
chunkedRevMap_1_0_13:
for {
if false { continue chunkedRevMap_1_0_13 }
var v_2 *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t19 *Constructor_Data_List_Types_Cons[string]
{
var __t_and_18 bool = false
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr != nil) {

var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
var __t_and_17 bool = false
if (__t_tag_15.Type == 9 && __t_tag_15.IntVal == 1358893437 && __t_tag_15.UnsafePtr != nil) {

var __t_tag_16 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}
__t_and_17 = (__t_tag_16.Type == 9 && __t_tag_16.IntVal == 1358893437 && __t_tag_16.UnsafePtr != nil)
}
__t_and_18 = __t_and_17
}
if __t_and_18 {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)})})})
v1_3_loop = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr).V1
continue chunkedRevMap_1_0_13
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[string]](gopurs_runtime.Value{})
goto end_branch_19
} else {

}
}
{
var reverseUnrolledMap_4_1_14 gopurs_runtime.Value
reverseUnrolledMap_4_1_14 = gopurs_runtime.Func(func(v2_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v2_5_loop gopurs_runtime.Value = v2_5_loop_val
var v3_6_loop *Constructor_Data_List_Types_Cons[string] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[string]](v3_6_loop_val)
reverseUnrolledMap_4_1_14:
for {
if false { continue reverseUnrolledMap_4_1_14 }
var v2_5 gopurs_runtime.Value = v2_5_loop
_ = v2_5
var v3_6 *Constructor_Data_List_Types_Cons[string] = v3_6_loop
_ = v3_6
var __t8 *Constructor_Data_List_Types_Cons[string]
{
var __t_and_7 bool = false
if (v2_5.Type == 9 && v2_5.IntVal == 1358893437 && v2_5.UnsafePtr != nil) {

var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0
var __t_and_6 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1358893437 && __t_tag_2.UnsafePtr != nil) {

var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}
var __t_and_5 bool = false
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 1358893437 && __t_tag_3.UnsafePtr != nil) {

var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V1)}
__t_and_5 = (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 1358893437 && __t_tag_4.UnsafePtr != nil)
}
__t_and_6 = __t_and_5
}
__t_and_7 = __t_and_6
}
if __t_and_7 {
v2_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V1)}
v3_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[string]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V0).StrVal()), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V0).StrVal()), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr).V0).StrVal()), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v3_6)})})})})})})})
continue reverseUnrolledMap_4_1_14
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[string]](gopurs_runtime.Value{})
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
var __t14 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr != nil) {
var __t13 gopurs_runtime.Value
{
var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 1358893437 && __t_tag_9.UnsafePtr != nil) {
var __t11 gopurs_runtime.Value
{
var __t_tag_10 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}
if (__t_tag_10.Type == 9 && __t_tag_10.IntVal == 1358893437 && __t_tag_10.UnsafePtr == nil) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V0).StrVal()), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V0).StrVal()), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})})})})}
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_11:
__t13 = __t11
goto end_branch_13
} else {

}
}
{
var __t_tag_12 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 1358893437 && __t_tag_12.UnsafePtr == nil) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V0).StrVal()), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})})}
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_13:
__t14 = __t13
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_14:
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[string]](gopurs_runtime.Apply2(reverseUnrolledMap_4_1_14, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[string]](__t14))}))
}
end_branch_19:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t19)}
}
}()
})
})
return gopurs_runtime.Apply(chunkedRevMap_1_0_13, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
}

func Call_Data_List_Types_listMap__4135416762(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var chunkedRevMap_1_0_15 gopurs_runtime.Value
chunkedRevMap_1_0_15 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](v_2_loop_val)
var v1_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val)
chunkedRevMap_1_0_15:
for {
if false { continue chunkedRevMap_1_0_15 }
var v_2 *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t19 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_and_18 bool = false
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr != nil) {

var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
var __t_and_17 bool = false
if (__t_tag_15.Type == 9 && __t_tag_15.IntVal == 1358893437 && __t_tag_15.UnsafePtr != nil) {

var __t_tag_16 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}
__t_and_17 = (__t_tag_16.Type == 9 && __t_tag_16.IntVal == 1358893437 && __t_tag_16.UnsafePtr != nil)
}
__t_and_18 = __t_and_17
}
if __t_and_18 {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)})})})
v1_3_loop = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr).V1
continue chunkedRevMap_1_0_15
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
goto end_branch_19
} else {

}
}
{
var reverseUnrolledMap_4_1_16 gopurs_runtime.Value
reverseUnrolledMap_4_1_16 = gopurs_runtime.Func(func(v2_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v2_5_loop gopurs_runtime.Value = v2_5_loop_val
var v3_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v3_6_loop_val)
reverseUnrolledMap_4_1_16:
for {
if false { continue reverseUnrolledMap_4_1_16 }
var v2_5 gopurs_runtime.Value = v2_5_loop
_ = v2_5
var v3_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v3_6_loop
_ = v3_6
var __t8 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_and_7 bool = false
if (v2_5.Type == 9 && v2_5.IntVal == 1358893437 && v2_5.UnsafePtr != nil) {

var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0
var __t_and_6 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1358893437 && __t_tag_2.UnsafePtr != nil) {

var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}
var __t_and_5 bool = false
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 1358893437 && __t_tag_3.UnsafePtr != nil) {

var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V1)}
__t_and_5 = (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 1358893437 && __t_tag_4.UnsafePtr != nil)
}
__t_and_6 = __t_and_5
}
__t_and_7 = __t_and_6
}
if __t_and_7 {
v2_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V1)}
v3_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr).V0), v3_6})})})})})})
continue reverseUnrolledMap_4_1_16
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
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
var __t14 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr != nil) {
var __t13 gopurs_runtime.Value
{
var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 1358893437 && __t_tag_9.UnsafePtr != nil) {
var __t11 gopurs_runtime.Value
{
var __t_tag_10 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}
if (__t_tag_10.Type == 9 && __t_tag_10.IntVal == 1358893437 && __t_tag_10.UnsafePtr == nil) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})})})})}
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_11:
__t13 = __t11
goto end_branch_13
} else {

}
}
{
var __t_tag_12 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 1358893437 && __t_tag_12.UnsafePtr == nil) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})})}
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_13:
__t14 = __t13
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_14:
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(reverseUnrolledMap_4_1_16, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__t14))}))
}
end_branch_19:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t19)}
}
}()
})
})
return gopurs_runtime.Apply(chunkedRevMap_1_0_15, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
}

func Call_Data_List_Types_nelCons__195558898(a_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)})}})})
}

func Call_Data_List_Types_nelCons__2148523118(a_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v_1_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
var a_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = a_0_loop
_ = a_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(a_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1))})})}})})
}

func Call_Data_List_Types_toList__2859885498(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)})})
}

func Call_Data_List_Types_toList__1324737658(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1))})})})
}

func Call_Data_List_Types_toList__2402503393(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)})})
}



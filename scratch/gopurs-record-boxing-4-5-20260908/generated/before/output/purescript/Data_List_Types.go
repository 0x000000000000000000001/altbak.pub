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
		cache_Data_List_Types_identity = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
	})
	return cache_Data_List_Types_identity
}

var cache_Data_List_Types_identity1 gopurs_runtime.Value
var once_Data_List_Types_identity1 sync.Once
func Get_Data_List_Types_identity1() gopurs_runtime.Value {
	once_Data_List_Types_identity1.Do(func() {
		cache_Data_List_Types_identity1 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
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
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](value1)}))}
})
})
	})
	return cache_Data_List_Types_Cons
}

var cache_Data_List_Types_Cons__1062349399 gopurs_runtime.Value
var once_Data_List_Types_Cons__1062349399 sync.Once
func Get_Data_List_Types_Cons__1062349399() gopurs_runtime.Value {
	once_Data_List_Types_Cons__1062349399.Do(func() {
		cache_Data_List_Types_Cons__1062349399 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2442833393_849153993(Call_Data_List_Types_Cons__1062349399(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[uint32, float64]](__eta_norm_1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]](__eta_norm_0_1_box))))}
})
	})
	return cache_Data_List_Types_Cons__1062349399
}

var cache_Data_List_Types_Cons__2126472648 gopurs_runtime.Value
var once_Data_List_Types_Cons__2126472648 sync.Once
func Get_Data_List_Types_Cons__2126472648() gopurs_runtime.Value {
	once_Data_List_Types_Cons__2126472648.Do(func() {
		cache_Data_List_Types_Cons__2126472648 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3704040722_849153993(Call_Data_List_Types_Cons__2126472648(__eta_norm_1_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[int64]](__eta_norm_0_1_box))))}
})
	})
	return cache_Data_List_Types_Cons__2126472648
}

var cache_Data_List_Types_Cons__2681173055 gopurs_runtime.Value
var once_Data_List_Types_Cons__2681173055 sync.Once
func Get_Data_List_Types_Cons__2681173055() gopurs_runtime.Value {
	once_Data_List_Types_Cons__2681173055.Do(func() {
		cache_Data_List_Types_Cons__2681173055 = gopurs_runtime.Func2(func(__eta_norm_0_unused_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3704040722_849153993(Call_Data_List_Types_Cons__2681173055(__eta_norm_0_unused_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[int64]](__eta_norm_1_1_box))))}
})
	})
	return cache_Data_List_Types_Cons__2681173055
}

var cache_Data_List_Types_NonEmptyList gopurs_runtime.Value
var once_Data_List_Types_NonEmptyList sync.Once
func Get_Data_List_Types_NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_NonEmptyList.Do(func() {
		cache_Data_List_Types_NonEmptyList = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952(Call_Data_List_Types_NonEmptyList(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](x_0_box))))}
})
	})
	return cache_Data_List_Types_NonEmptyList
}

var cache_Data_List_Types_NonEmptyList__3166204772 gopurs_runtime.Value
var once_Data_List_Types_NonEmptyList__3166204772 sync.Once
func Get_Data_List_Types_NonEmptyList__3166204772() gopurs_runtime.Value {
	once_Data_List_Types_NonEmptyList__3166204772.Do(func() {
		cache_Data_List_Types_NonEmptyList__3166204772 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952(Call_Data_List_Types_NonEmptyList__3166204772(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](x_0_box))))}
})
	})
	return cache_Data_List_Types_NonEmptyList__3166204772
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
		cache_Data_List_Types_newtypeNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3742495784_385277032((&Constructor_Data_Newtype_Newtype[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})))}
	})
	return cache_Data_List_Types_newtypeNonEmptyList
}

var cache_Data_List_Types_nelCons gopurs_runtime.Value
var once_Data_List_Types_nelCons sync.Once
func Get_Data_List_Types_nelCons() gopurs_runtime.Value {
	once_Data_List_Types_nelCons.Do(func() {
		cache_Data_List_Types_nelCons = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952(Call_Data_List_Types_nelCons(a_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box))))}
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
		cache_Data_List_Types_functorList = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3262795586_2812149806((&Constructor_Data_Functor_Functor[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, Get_Data_List_Types_listMap()})))}
	})
	return cache_Data_List_Types_functorList
}

var cache_Data_List_Types_functorNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_functorNonEmptyList sync.Once
func Get_Data_List_Types_functorNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_functorNonEmptyList.Do(func() {
		cache_Data_List_Types_functorNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2801299215_2812149806((&Constructor_Data_Functor_Functor[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_List_Types_listMap(f_0), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V1)))}}))}
})})))}
	})
	return cache_Data_List_Types_functorNonEmptyList
}

var cache_Data_List_Types_foldableList gopurs_runtime.Value
var once_Data_List_Types_foldableList sync.Once
func Get_Data_List_Types_foldableList() gopurs_runtime.Value {
	once_Data_List_Types_foldableList.Do(func() {
		cache_Data_List_Types_foldableList = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1022383170_1680800814((&Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 shape=Other bindingType=(TypeVar m)
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_foldableList()).V1), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_2 shape=App(Other) bindingType=Any
__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Box(Semigroup0_1_0.V0), acc_4)
_ = __local_var_5_2
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, gopurs_runtime.Apply(f_3, x_6))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_1_3_2 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_1_3_2
var go__go_1_3_2 gopurs_runtime.Value
_ = go__go_1_3_2
Call_local_Data_List_Types_go__go_1_3_2 = func(b_2_loop gopurs_runtime.Value, v_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_1_3_2:
for {
if false { continue go__go_1_3_2 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var v_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_3_loop
_ = v_3
var __t4 gopurs_runtime.Value
{
if (v_3 == nil) {
__t4 = b_2
goto end_branch_4
} else {

}
}
{
if (v_3 != nil) {
b_2_loop = gopurs_runtime.Apply2(f_0, b_2, (v_3).V0)
v_3_loop = (v_3).V1
continue go__go_1_3_2
__t4 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_1_3_2 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_1_3_2(b_2_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_3_loop_val))
})
})
return go__go_1_3_2
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_5 shape=App(Other) bindingType=Any
__local_var_2_5 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_foldableList()).V1), gopurs_runtime.Func2(func(b_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
}), b_1)
_ = __local_var_2_5
var Call_local_Data_List_Types_go__go_3_7_3 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_3_7_3
var go__go_3_7_3 gopurs_runtime.Value
_ = go__go_3_7_3
Call_local_Data_List_Types_go__go_3_7_3 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_3_7_3:
for {
if false { continue go__go_3_7_3 }
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_5_loop
_ = v1_5
var __t8 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_5 == nil) {
__t8 = v_4
goto end_branch_8
} else {

}
}
{
if (v1_5 != nil) {
v_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_5).V0, v_4})
v1_5_loop = (v1_5).V1
continue go__go_3_7_3
__t8 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_8
} else {

}
}
{
__t8 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
}
}
go__go_3_7_3 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_3_7_3(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val)))}
})
})
// TAST (Let): __local_var_3_6 shape=LetRec(App(Other)) bindingType=(Func [(ADT ["Data","List","Types","List"] [(TypeVar a)])] (ADT ["Data","List","Types","List"] [(TypeVar a)]))
__local_var_3_6 := gopurs_runtime.Apply(go__go_3_7_3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_3_6
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_5, gopurs_runtime.Apply(__local_var_3_6, x_4))
})
})})))}
	})
	return cache_Data_List_Types_foldableList
}

var cache_Data_List_Types_foldableNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_foldableNonEmptyList sync.Once
func Get_Data_List_Types_foldableNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_foldableNonEmptyList.Do(func() {
		cache_Data_List_Types_foldableNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_4064382095_1680800814((&Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_4_1 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
Semigroup0_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_4_1
var Call_local_Data_List_Types_go__go_5_2_4 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_5_2_4
var go__go_5_2_4 gopurs_runtime.Value
_ = go__go_5_2_4
Call_local_Data_List_Types_go__go_5_2_4 = func(b_6_loop gopurs_runtime.Value, v_7_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_5_2_4:
for {
if false { continue go__go_5_2_4 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_7_loop
_ = v_7
var __t3 gopurs_runtime.Value
{
if (v_7 == nil) {
__t3 = b_6
goto end_branch_3
} else {

}
}
{
if (v_7 != nil) {
b_6_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_4_1.V0), b_6, gopurs_runtime.Apply(f_2, (v_7).V0))
v_7_loop = (v_7).V1
continue go__go_5_2_4
__t3 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_5_2_4 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_5_2_4(b_6_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_7_loop_val))
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), Call_local_Data_List_Types_go__go_5_2_4(gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)))
})
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_3_4_5 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_3_4_5
var go__go_3_4_5 gopurs_runtime.Value
_ = go__go_3_4_5
Call_local_Data_List_Types_go__go_3_4_5 = func(b_4_loop gopurs_runtime.Value, v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_3_4_5:
for {
if false { continue go__go_3_4_5 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var __t5 gopurs_runtime.Value
{
if (v_5 == nil) {
__t5 = b_4
goto end_branch_5
} else {

}
}
{
if (v_5 != nil) {
b_4_loop = gopurs_runtime.Apply2(f_0, b_4, (v_5).V0)
v_5_loop = (v_5).V1
continue go__go_3_4_5
__t5 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_3_4_5 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_3_4_5(b_4_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val))
})
})
return Call_local_Data_List_Types_go__go_3_4_5(gopurs_runtime.Apply2(f_0, b_1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1))
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_3_6_6 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_3_6_6
var go__go_3_6_6 gopurs_runtime.Value
_ = go__go_3_6_6
Call_local_Data_List_Types_go__go_3_6_6 = func(b_4_loop gopurs_runtime.Value, v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_3_6_6:
for {
if false { continue go__go_3_6_6 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var __t7 gopurs_runtime.Value
{
if (v_5 == nil) {
__t7 = b_4
goto end_branch_7
} else {

}
}
{
if (v_5 != nil) {
b_4_loop = gopurs_runtime.Apply2(f_0, (v_5).V0, b_4)
v_5_loop = (v_5).V1
continue go__go_3_6_6
__t7 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_3_6_6 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_3_6_6(b_4_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val))
})
})
var Call_local_Data_List_Types_go__go_4_8_7 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_4_8_7
var go__go_4_8_7 gopurs_runtime.Value
_ = go__go_4_8_7
Call_local_Data_List_Types_go__go_4_8_7 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_4_8_7:
for {
if false { continue go__go_4_8_7 }
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var __t9 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_6 == nil) {
__t9 = v_5
goto end_branch_9
} else {

}
}
{
if (v1_6 != nil) {
v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
v1_6_loop = (v1_6).V1
continue go__go_4_8_7
__t9 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_9
} else {

}
}
{
__t9 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
}
}
go__go_4_8_7 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_4_8_7(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
})
})
return gopurs_runtime.Apply2(f_0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, Call_local_Data_List_Types_go__go_3_6_6(b_1, Call_local_Data_List_Types_go__go_4_8_7((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1))))
})})))}
	})
	return cache_Data_List_Types_foldableNonEmptyList
}

var cache_Data_List_Types_foldableWithIndexList gopurs_runtime.Value
var once_Data_List_Types_foldableWithIndexList sync.Once
func Get_Data_List_Types_foldableWithIndexList() gopurs_runtime.Value {
	once_Data_List_Types_foldableWithIndexList.Do(func() {
		cache_Data_List_Types_foldableWithIndexList = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_538456415_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1022383170_1680800814(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_foldableList())))}
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 shape=Other bindingType=(TypeVar m)
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_foldableWithIndexList()).V2), gopurs_runtime.Func2(func(i_4 gopurs_runtime.Value, acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 shape=App(Other) bindingType=Any
__local_var_6_2 := gopurs_runtime.Apply(gopurs_runtime.Box(Semigroup0_1_0.V0), acc_5)
_ = __local_var_6_2
// TAST (Let): __local_var_7_3 shape=App(Other) bindingType=(Func [(TypeVar a)] (TypeVar m))
__local_var_7_3 := gopurs_runtime.Apply(f_3, gopurs_runtime.Int(i_4.IntVal))
_ = __local_var_7_3
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_2, gopurs_runtime.Apply(__local_var_7_3, x_8))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, acc_1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_2_5_8 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_2_5_8
var go__go_2_5_8 gopurs_runtime.Value
_ = go__go_2_5_8
Call_local_Data_List_Types_go__go_2_5_8 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_5_8:
for {
if false { continue go__go_2_5_8 }
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
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_3.UnsafePtr).V0.IntVal) + (int64(1))), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_3.UnsafePtr).V0.IntVal), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_3.UnsafePtr).V1, (v_4).V0)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
v_4_loop = (v_4).V1
continue go__go_2_5_8
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
go__go_2_5_8 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_2_5_8(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
// TAST (Let): __local_var_2_4 shape=LetRec(App(Other)) bindingType=(Func [(ADT ["Data","List","Types","List"] [(TypeVar a)])] (ADT ["Data","Tuple","Tuple"] [Int, (TypeVar b)]))
__local_var_2_4 := Rebox_Data_List_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_2_5_8, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(int64(0)), acc_1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))})))
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1728839155_138441832(__local_var_2_4))}, x_3).UnsafePtr).V1
})
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_3_8_9 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_3_8_9
var go__go_3_8_9 gopurs_runtime.Value
_ = go__go_3_8_9
Call_local_Data_List_Types_go__go_3_8_9 = func(b_4_loop gopurs_runtime.Value, v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_3_8_9:
for {
if false { continue go__go_3_8_9 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var __t9 gopurs_runtime.Value
{
if (v_5 == nil) {
__t9 = b_4
goto end_branch_9
} else {

}
}
{
if (v_5 != nil) {
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3800170591_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V0.IntVal) + (int64(1))), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_5).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V1)}))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
v_5_loop = (v_5).V1
continue go__go_3_8_9
__t9 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_3_8_9 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_3_8_9(b_4_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val))
})
})
// TAST (Let): v_3_7 shape=LetRec(App(Other)) bindingType=(ADT ["Data","Tuple","Tuple"] [Int, (ADT ["Data","List","Types","List"] [(TypeVar a)])])
v_3_7 := Rebox_Data_List_Types_138441832_3800170591(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](Call_local_Data_List_Types_go__go_3_8_9(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3800170591_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(int64(0)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_2))))
_ = v_3_7
var Call_local_Data_List_Types_go__go_4_10_10 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_4_10_10
var go__go_4_10_10 gopurs_runtime.Value
_ = go__go_4_10_10
Call_local_Data_List_Types_go__go_4_10_10 = func(b_5_loop gopurs_runtime.Value, v_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_4_10_10:
for {
if false { continue go__go_4_10_10 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_6_loop
_ = v_6
var __t11 gopurs_runtime.Value
{
if (v_6 == nil) {
__t11 = b_5
goto end_branch_11
} else {

}
}
{
if (v_6 != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V0.IntVal) - (int64(1))), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V0.IntVal) - (int64(1))), (v_6).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
v_6_loop = (v_6).V1
continue go__go_4_10_10
__t11 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_4_10_10 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_4_10_10(b_5_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_6_loop_val))
})
})
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(Call_local_Data_List_Types_go__go_4_10_10(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int((v_3_7).V0), b_1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}, (v_3_7).V1).UnsafePtr).V1
})})))}
	})
	return cache_Data_List_Types_foldableWithIndexList
}

var cache_Data_List_Types_foldableWithIndexNonEmpty gopurs_runtime.Value
var once_Data_List_Types_foldableWithIndexNonEmpty sync.Once
func Get_Data_List_Types_foldableWithIndexNonEmpty() gopurs_runtime.Value {
	once_Data_List_Types_foldableWithIndexNonEmpty.Do(func() {
		cache_Data_List_Types_foldableWithIndexNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_4016503379_3725484264(Rebox_Data_List_Types_3725484264_4016503379(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_NonEmpty_foldableWithIndexNonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_538456415_3725484264(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_foldableWithIndexList())))})))))}
	})
	return cache_Data_List_Types_foldableWithIndexNonEmpty
}

var cache_Data_List_Types_foldableWithIndexNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_foldableWithIndexNonEmptyList sync.Once
func Get_Data_List_Types_foldableWithIndexNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_foldableWithIndexNonEmptyList.Do(func() {
		cache_Data_List_Types_foldableWithIndexNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1951493170_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_4064382095_1680800814(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_foldableNonEmptyList())))}
}), gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_Data_NonEmpty_foldableWithIndexNonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_538456415_3725484264(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_foldableWithIndexList())))}), "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))}, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 int64
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr == nil) {
__t0 = int64(0)
goto end_branch_0
} else {

}
}
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr != nil) {
__t0 = (int64(1)) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0.IntVal)
goto end_branch_0
} else {

}
}
{
__t0 = func() int64 { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Apply(f_1, gopurs_runtime.Int(__t0))
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952(Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_2))))})
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_Data_NonEmpty_foldableWithIndexNonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_538456415_3725484264(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_foldableWithIndexList())))}), "foldlWithIndex"), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 int64
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr == nil) {
__t1 = int64(0)
goto end_branch_1
} else {

}
}
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr != nil) {
__t1 = (int64(1)) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0.IntVal)
goto end_branch_1
} else {

}
}
{
__t1 = func() int64 { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Apply(f_0, gopurs_runtime.Int(__t1))
}), b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952(Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_2))))})
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_Data_NonEmpty_foldableWithIndexNonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_538456415_3725484264(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_foldableWithIndexList())))}), "foldrWithIndex"), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 int64
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr == nil) {
__t2 = int64(0)
goto end_branch_2
} else {

}
}
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr != nil) {
__t2 = (int64(1)) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0.IntVal)
goto end_branch_2
} else {

}
}
{
__t2 = func() int64 { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Apply(f_0, gopurs_runtime.Int(__t2))
}), b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952(Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_2))))})
})})))}
	})
	return cache_Data_List_Types_foldableWithIndexNonEmptyList
}

var cache_Data_List_Types_functorWithIndexList gopurs_runtime.Value
var once_Data_List_Types_functorWithIndexList sync.Once
func Get_Data_List_Types_functorWithIndexList() gopurs_runtime.Value {
	once_Data_List_Types_functorWithIndexList.Do(func() {
		cache_Data_List_Types_functorWithIndexList = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_721753375_2412140840((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3262795586_2812149806(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_functorList())))}
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_2_1_11 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_2_1_11
var go__go_2_1_11 gopurs_runtime.Value
_ = go__go_2_1_11
Call_local_Data_List_Types_go__go_2_1_11 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_1_11:
for {
if false { continue go__go_2_1_11 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var __t2 gopurs_runtime.Value
{
if (v_4 == nil) {
__t2 = b_3
goto end_branch_2
} else {

}
}
{
if (v_4 != nil) {
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3800170591_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_3.UnsafePtr).V0.IntVal) + (int64(1))), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_4).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_3.UnsafePtr).V1)}))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
v_4_loop = (v_4).V1
continue go__go_2_1_11
__t2 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_2_1_11 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_2_1_11(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
// TAST (Let): v_2_0 shape=LetRec(App(Other)) bindingType=(ADT ["Data","Tuple","Tuple"] [Int, (ADT ["Data","List","Types","List"] [(TypeVar a)])])
v_2_0 := Rebox_Data_List_Types_138441832_3800170591(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](Call_local_Data_List_Types_go__go_2_1_11(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3800170591_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(int64(0)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1))))
_ = v_2_0
var Call_local_Data_List_Types_go__go_3_3_12 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_3_3_12
var go__go_3_3_12 gopurs_runtime.Value
_ = go__go_3_3_12
Call_local_Data_List_Types_go__go_3_3_12 = func(b_4_loop gopurs_runtime.Value, v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_3_3_12:
for {
if false { continue go__go_3_3_12 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var __t4 gopurs_runtime.Value
{
if (v_5 == nil) {
__t4 = b_4
goto end_branch_4
} else {

}
}
{
if (v_5 != nil) {
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V0.IntVal) - (int64(1))), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V0.IntVal) - (int64(1))), (v_5).V0), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V1)}))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
v_5_loop = (v_5).V1
continue go__go_3_3_12
__t4 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_3_3_12 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_3_3_12(b_4_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val))
})
})
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(Call_local_Data_List_Types_go__go_3_3_12(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int((v_2_0).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}, (v_2_0).V1).UnsafePtr).V1
})})))}
	})
	return cache_Data_List_Types_functorWithIndexList
}

var cache_Data_List_Types_functorWithIndex gopurs_runtime.Value
var once_Data_List_Types_functorWithIndex sync.Once
func Get_Data_List_Types_functorWithIndex() gopurs_runtime.Value {
	once_Data_List_Types_functorWithIndex.Do(func() {
		cache_Data_List_Types_functorWithIndex = func() gopurs_runtime.Value {
// TAST (Let): functorNonEmpty1_0_0 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [(ADT ["Data","NonEmpty","NonEmpty"] [(TypeVar f)])])
functorNonEmpty1_0_0 := (&Constructor_Data_Functor_Functor[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_List_Types_listMap(f_0), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V1)))}}))}
})})
_ = functorNonEmpty1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_111597075_2412140840((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1468200963_2812149806(functorNonEmpty1_0_0))}
}), gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_3_2_13 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_3_2_13
var go__go_3_2_13 gopurs_runtime.Value
_ = go__go_3_2_13
Call_local_Data_List_Types_go__go_3_2_13 = func(b_4_loop gopurs_runtime.Value, v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_3_2_13:
for {
if false { continue go__go_3_2_13 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var __t3 gopurs_runtime.Value
{
if (v_5 == nil) {
__t3 = b_4
goto end_branch_3
} else {

}
}
{
if (v_5 != nil) {
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3800170591_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V0.IntVal) + (int64(1))), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_5).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V1)}))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
v_5_loop = (v_5).V1
continue go__go_3_2_13
__t3 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_3_2_13 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_3_2_13(b_4_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val))
})
})
// TAST (Let): v_3_1 shape=LetRec(App(Other)) bindingType=(ADT ["Data","Tuple","Tuple"] [Int, (ADT ["Data","List","Types","List"] [(TypeVar a)])])
v_3_1 := Rebox_Data_List_Types_138441832_3800170591(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](Call_local_Data_List_Types_go__go_3_2_13(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3800170591_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(int64(0)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1))))
_ = v_3_1
var Call_local_Data_List_Types_go__go_4_4_14 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_4_4_14
var go__go_4_4_14 gopurs_runtime.Value
_ = go__go_4_4_14
Call_local_Data_List_Types_go__go_4_4_14 = func(b_5_loop gopurs_runtime.Value, v_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_4_4_14:
for {
if false { continue go__go_4_4_14 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_6_loop
_ = v_6
var __t5 gopurs_runtime.Value
{
if (v_6 == nil) {
__t5 = b_5
goto end_branch_5
} else {

}
}
{
if (v_6 != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V0.IntVal) - (int64(1))), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V0.IntVal) - (int64(1))), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, (v_6).V0), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V1)}))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
v_6_loop = (v_6).V1
continue go__go_4_4_14
__t5 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_4_4_14 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_4_4_14(b_5_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_6_loop_val))
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))}, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(Call_local_Data_List_Types_go__go_4_4_14(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int((v_3_1).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}, (v_3_1).V1).UnsafePtr).V1}))}
})})))}
}()
	})
	return cache_Data_List_Types_functorWithIndex
}

var cache_Data_List_Types_functorWithIndexNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_functorWithIndexNonEmptyList sync.Once
func Get_Data_List_Types_functorWithIndexNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_functorWithIndexNonEmptyList.Do(func() {
		cache_Data_List_Types_functorWithIndexNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3339399026_2412140840((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2801299215_2812149806(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_functorNonEmptyList())))}
}), gopurs_runtime.Func2(func(fn_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 int64
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
if (__t_tag_0 == nil) {
__t2 = int64(0)
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
if (__t_tag_1 != nil) {
__t2 = (int64(1)) + (((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil)).V0.IntVal)
goto end_branch_2
} else {

}
}
{
__t2 = func() int64 { panic("Failed pattern match") }()
}
end_branch_2:
var Call_local_Data_List_Types_go__go_2_4_15 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_2_4_15
var go__go_2_4_15 gopurs_runtime.Value
_ = go__go_2_4_15
Call_local_Data_List_Types_go__go_2_4_15 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_4_15:
for {
if false { continue go__go_2_4_15 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var __t5 gopurs_runtime.Value
{
if (v_4 == nil) {
__t5 = b_3
goto end_branch_5
} else {

}
}
{
if (v_4 != nil) {
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3800170591_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_3.UnsafePtr).V0.IntVal) + (int64(1))), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_4).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_3.UnsafePtr).V1)}))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
v_4_loop = (v_4).V1
continue go__go_2_4_15
__t5 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_2_4_15 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_2_4_15(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
// TAST (Let): v_2_3 shape=LetRec(App(Other)) bindingType=(ADT ["Data","Tuple","Tuple"] [Int, (ADT ["Data","List","Types","List"] [(TypeVar a)])])
v_2_3 := Rebox_Data_List_Types_138441832_3800170591(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](Call_local_Data_List_Types_go__go_2_4_15(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3800170591_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(int64(0)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1))))
_ = v_2_3
var Call_local_Data_List_Types_go__go_3_6_16 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_3_6_16
var go__go_3_6_16 gopurs_runtime.Value
_ = go__go_3_6_16
Call_local_Data_List_Types_go__go_3_6_16 = func(b_4_loop gopurs_runtime.Value, v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_3_6_16:
for {
if false { continue go__go_3_6_16 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var __t7 gopurs_runtime.Value
{
if (v_5 == nil) {
__t7 = b_4
goto end_branch_7
} else {

}
}
{
if (v_5 != nil) {
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V0.IntVal) - (int64(1))), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(fn_0, gopurs_runtime.Int((int64(1)) + (((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V0.IntVal) - (int64(1)))), (v_5).V0), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V1)}))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
v_5_loop = (v_5).V1
continue go__go_3_6_16
__t7 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_3_6_16 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_3_6_16(b_4_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val))
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(fn_0, gopurs_runtime.Int(__t2), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(Call_local_Data_List_Types_go__go_3_6_16(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int((v_2_3).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}, (v_2_3).V1).UnsafePtr).V1}))}
})})))}
	})
	return cache_Data_List_Types_functorWithIndexNonEmptyList
}

var cache_Data_List_Types_semigroupList gopurs_runtime.Value
var once_Data_List_Types_semigroupList sync.Once
func Get_Data_List_Types_semigroupList() gopurs_runtime.Value {
	once_Data_List_Types_semigroupList.Do(func() {
		cache_Data_List_Types_semigroupList = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_4022093634_4179793454((&Constructor_Data_Semigroup_Semigroup[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(xs_0 gopurs_runtime.Value, ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_2_0_17 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_2_0_17
var go__go_2_0_17 gopurs_runtime.Value
_ = go__go_2_0_17
Call_local_Data_List_Types_go__go_2_0_17 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_0_17:
for {
if false { continue go__go_2_0_17 }
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
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_4).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_3)}))}
v_4_loop = (v_4).V1
continue go__go_2_0_17
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
go__go_2_0_17 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_2_0_17(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
var Call_local_Data_List_Types_go__go_3_2_18 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_3_2_18
var go__go_3_2_18 gopurs_runtime.Value
_ = go__go_3_2_18
Call_local_Data_List_Types_go__go_3_2_18 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_3_2_18:
for {
if false { continue go__go_3_2_18 }
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_5_loop
_ = v1_5
var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_5 == nil) {
__t3 = v_4
goto end_branch_3
} else {

}
}
{
if (v1_5 != nil) {
v_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_5).V0, v_4})
v1_5_loop = (v1_5).V1
continue go__go_3_2_18
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}
go__go_3_2_18 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_3_2_18(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val)))}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_local_Data_List_Types_go__go_2_0_17(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_1))}, Call_local_Data_List_Types_go__go_3_2_18((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_0)))))}
})})))}
	})
	return cache_Data_List_Types_semigroupList
}

var cache_Data_List_Types_monoidList gopurs_runtime.Value
var once_Data_List_Types_monoidList sync.Once
func Get_Data_List_Types_monoidList() gopurs_runtime.Value {
	once_Data_List_Types_monoidList.Do(func() {
		cache_Data_List_Types_monoidList = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_640718306_1201789390((&Constructor_Data_Monoid_Monoid[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_4022093634_4179793454(Rebox_Data_List_Types_4179793454_4022093634(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_List_Types_semigroupList()))))}
}), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})})))}
	})
	return cache_Data_List_Types_monoidList
}

var cache_Data_List_Types_semigroupNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_semigroupNonEmptyList sync.Once
func Get_Data_List_Types_semigroupNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_semigroupNonEmptyList.Do(func() {
		cache_Data_List_Types_semigroupNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_655843087_4179793454((&Constructor_Data_Semigroup_Semigroup[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, as_prime__1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_2_0_19 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_2_0_19
var go__go_2_0_19 gopurs_runtime.Value
_ = go__go_2_0_19
Call_local_Data_List_Types_go__go_2_0_19 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_0_19:
for {
if false { continue go__go_2_0_19 }
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
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_4).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_3)}))}
v_4_loop = (v_4).V1
continue go__go_2_0_19
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
go__go_2_0_19 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_2_0_19(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
var Call_local_Data_List_Types_go__go_3_2_20 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_3_2_20
var go__go_3_2_20 gopurs_runtime.Value
_ = go__go_3_2_20
Call_local_Data_List_Types_go__go_3_2_20 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_3_2_20:
for {
if false { continue go__go_3_2_20 }
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_5_loop
_ = v1_5
var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_5 == nil) {
__t3 = v_4
goto end_branch_3
} else {

}
}
{
if (v1_5 != nil) {
v_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_5).V0, v_4})
v1_5_loop = (v1_5).V1
continue go__go_3_2_20
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}
go__go_3_2_20 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_3_2_20(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val)))}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952((&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_local_Data_List_Types_go__go_2_0_19(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(as_prime__1.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(as_prime__1.UnsafePtr).V1)}))}, Call_local_Data_List_Types_go__go_3_2_20((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1)))))}})))}
})})))}
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
		cache_Data_List_Types_traversableList = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3037784642_3043886126((&Constructor_Data_Traversable_Traversable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1022383170_1680800814(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_foldableList())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3262795586_2812149806(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_functorList())))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_traversableList()).V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): Apply0_2_1 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m)])
Apply0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_4_3_23 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_4_3_23
var go__go_4_3_23 gopurs_runtime.Value
_ = go__go_4_3_23
Call_local_Data_List_Types_go__go_4_3_23 = func(b_5_loop gopurs_runtime.Value, v_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_4_3_23:
for {
if false { continue go__go_4_3_23 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_6_loop
_ = v_6
var __t4 gopurs_runtime.Value
{
if (v_6 == nil) {
__t4 = b_5
goto end_branch_4
} else {

}
}
{
if (v_6 != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_6).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_5)}))}
v_6_loop = (v_6).V1
continue go__go_4_3_23
__t4 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_4_3_23 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_4_3_23(b_5_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_6_loop_val))
})
})
// TAST (Let): __local_var_4_2 shape=App(Other) bindingType=Any
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_4_3_23, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})))})
_ = __local_var_4_2
var Call_local_Data_List_Types_go__go_5_6_24 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_5_6_24
var go__go_5_6_24 gopurs_runtime.Value
_ = go__go_5_6_24
Call_local_Data_List_Types_go__go_5_6_24 = func(b_6_loop gopurs_runtime.Value, v_7_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_5_6_24:
for {
if false { continue go__go_5_6_24 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_7_loop
_ = v_7
var __t7 gopurs_runtime.Value
{
if (v_7 == nil) {
__t7 = b_6
goto end_branch_7
} else {

}
}
{
if (v_7 != nil) {
b_6_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_2_1.V0), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func2(func(b_8 gopurs_runtime.Value, a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_9, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_8)}))}
}), b_6), gopurs_runtime.Apply(f_3, (v_7).V0))
v_7_loop = (v_7).V1
continue go__go_5_6_24
__t7 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_5_6_24 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_5_6_24(b_6_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_7_loop_val))
})
})
// TAST (Let): __local_var_5_5 shape=LetRec(App(Other)) bindingType=(Func [(ADT ["Data","List","Types","List"] [(TypeVar a)])] (TypeApp (TypeVar m) [(ADT ["Data","List","Types","List"] [(TypeVar b)])]))
__local_var_5_5 := gopurs_runtime.Apply(go__go_5_6_24, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}))}))
_ = __local_var_5_5
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_5_5, x_6))
})
})
})})))}
	})
	return cache_Data_List_Types_traversableList
}

var cache_Data_List_Types_traversableNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_traversableNonEmptyList sync.Once
func Get_Data_List_Types_traversableNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_traversableNonEmptyList.Do(func() {
		cache_Data_List_Types_traversableNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3425342735_3043886126(Rebox_Data_List_Types_3043886126_3425342735(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_NonEmpty_traversableNonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3037784642_3043886126(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_traversableList())))})))))}
	})
	return cache_Data_List_Types_traversableNonEmptyList
}

var cache_Data_List_Types_traversableWithIndexList gopurs_runtime.Value
var once_Data_List_Types_traversableWithIndexList sync.Once
func Get_Data_List_Types_traversableWithIndexList() gopurs_runtime.Value {
	once_Data_List_Types_traversableWithIndexList.Do(func() {
		cache_Data_List_Types_traversableWithIndexList = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1544744991_1812164904((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_538456415_3725484264(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_foldableWithIndexList())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_721753375_2412140840(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_functorWithIndexList())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3037784642_3043886126(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_traversableList())))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): Apply0_2_1 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m)])
Apply0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_4_3_25 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_4_3_25
var go__go_4_3_25 gopurs_runtime.Value
_ = go__go_4_3_25
Call_local_Data_List_Types_go__go_4_3_25 = func(b_5_loop gopurs_runtime.Value, v_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_4_3_25:
for {
if false { continue go__go_4_3_25 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_6_loop
_ = v_6
var __t4 gopurs_runtime.Value
{
if (v_6 == nil) {
__t4 = b_5
goto end_branch_4
} else {

}
}
{
if (v_6 != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_6).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_5)}))}
v_6_loop = (v_6).V1
continue go__go_4_3_25
__t4 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_4_3_25 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_4_3_25(b_5_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_6_loop_val))
})
})
// TAST (Let): __local_var_4_2 shape=App(Other) bindingType=Any
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Apply(go__go_4_3_25, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}))
_ = __local_var_4_2
var Call_local_Data_List_Types_go__go_5_7_26 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_5_7_26
var go__go_5_7_26 gopurs_runtime.Value
_ = go__go_5_7_26
Call_local_Data_List_Types_go__go_5_7_26 = func(b_6_loop gopurs_runtime.Value, v_7_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_5_7_26:
for {
if false { continue go__go_5_7_26 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_7_loop
_ = v_7
var __t8 gopurs_runtime.Value
{
if (v_7 == nil) {
__t8 = b_6
goto end_branch_8
} else {

}
}
{
if (v_7 != nil) {
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_6.UnsafePtr).V0.IntVal) + (int64(1))), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_2_1.V0), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func2(func(b_8 gopurs_runtime.Value, a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_9, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_8)}))}
}), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_6.UnsafePtr).V1), gopurs_runtime.Apply2(f_3, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_6.UnsafePtr).V0.IntVal), (v_7).V0))}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
v_7_loop = (v_7).V1
continue go__go_5_7_26
__t8 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_5_7_26 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_5_7_26(b_6_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_7_loop_val))
})
})
// TAST (Let): __local_var_5_6 shape=LetRec(App(Other)) bindingType=(Func [(ADT ["Data","List","Types","List"] [(TypeVar a)])] (ADT ["Data","Tuple","Tuple"] [Int, (TypeVar b)]))
__local_var_5_6 := Rebox_Data_List_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_5_7_26, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1728839155_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(int64(0)), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}))})}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))})))
_ = __local_var_5_6
// TAST (Let): __local_var_5_5 shape=Let(Abs(Other)) bindingType=(Func [(ADT ["Data","List","Types","List"] [(TypeVar a)])] (TypeApp (TypeVar m) [(ADT ["Data","List","Types","List"] [(TypeVar b)])]))
__local_var_5_5 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1728839155_138441832(__local_var_5_6))}, x_6).UnsafePtr).V1
})
_ = __local_var_5_5
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_5_5, x_6))
})
})
})})))}
	})
	return cache_Data_List_Types_traversableWithIndexList
}

var cache_Data_List_Types_traversableWithIndexNonEmpty gopurs_runtime.Value
var once_Data_List_Types_traversableWithIndexNonEmpty sync.Once
func Get_Data_List_Types_traversableWithIndexNonEmpty() gopurs_runtime.Value {
	once_Data_List_Types_traversableWithIndexNonEmpty.Do(func() {
		cache_Data_List_Types_traversableWithIndexNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_4122471187_1812164904(Rebox_Data_List_Types_1812164904_4122471187(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_NonEmpty_traversableWithIndexNonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1544744991_1812164904(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_traversableWithIndexList())))})))))}
	})
	return cache_Data_List_Types_traversableWithIndexNonEmpty
}

var cache_Data_List_Types_traversableWithIndexNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_traversableWithIndexNonEmptyList sync.Once
func Get_Data_List_Types_traversableWithIndexNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_traversableWithIndexNonEmptyList.Do(func() {
		cache_Data_List_Types_traversableWithIndexNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_677368690_1812164904((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1951493170_3725484264(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_foldableWithIndexNonEmptyList())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3339399026_2412140840(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_functorWithIndexNonEmptyList())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3425342735_3043886126(Rebox_Data_List_Types_3043886126_3425342735(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_NonEmpty_traversableNonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3037784642_3043886126(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_traversableList())))})))))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return x_4
}), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_Data_NonEmpty_traversableWithIndexNonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1544744991_1812164904(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_traversableWithIndexList())))}), "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 int64
{
if (x_4.Type == 9 && x_4.IntVal == 930809136 && x_4.UnsafePtr == nil) {
__t1 = int64(0)
goto end_branch_1
} else {

}
}
{
if (x_4.Type == 9 && x_4.IntVal == 930809136 && x_4.UnsafePtr != nil) {
__t1 = (int64(1)) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(x_4.UnsafePtr).V0.IntVal)
goto end_branch_1
} else {

}
}
{
__t1 = func() int64 { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Apply(f_2, gopurs_runtime.Int(__t1))
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952(Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_3))))}))
})
})})))}
	})
	return cache_Data_List_Types_traversableWithIndexNonEmptyList
}

var cache_Data_List_Types_unfoldable1List gopurs_runtime.Value
var once_Data_List_Types_unfoldable1List sync.Once
func Get_Data_List_Types_unfoldable1List() gopurs_runtime.Value {
	once_Data_List_Types_unfoldable1List.Do(func() {
		cache_Data_List_Types_unfoldable1List = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_4073635714_2187088110((&Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_2_0_27 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_2_0_27
var go__go_2_0_27 gopurs_runtime.Value
_ = go__go_2_0_27
Call_local_Data_List_Types_go__go_2_0_27 = func(source_3_loop gopurs_runtime.Value, memo_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_2_0_27:
for {
if false { continue go__go_2_0_27 }
var source_3 gopurs_runtime.Value = source_3_loop
_ = source_3
var memo_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = memo_4_loop
_ = memo_4
// TAST (Let): v_5_1 shape=App(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (ADT ["Data","Maybe","Maybe"] [(TypeVar b)])])
v_5_1 := Rebox_Data_List_Types_138441832_3804580809(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, source_3)))
_ = v_5_1
var __t6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = (v_5_1).V1
if (__t_tag_2 != nil) {
source_3_loop = ((v_5_1).V1).V0
memo_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_5_1).V0, memo_4})
continue go__go_2_0_27
__t6 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_6
} else {

}
}
{
var __t_tag_3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = (v_5_1).V1
if (__t_tag_3 == nil) {
var Call_local_Data_List_Types_go__go_6_4_28 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_6_4_28
var go__go_6_4_28 gopurs_runtime.Value
_ = go__go_6_4_28
Call_local_Data_List_Types_go__go_6_4_28 = func(b_7_loop gopurs_runtime.Value, v_8_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_6_4_28:
for {
if false { continue go__go_6_4_28 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var v_8 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_8_loop
_ = v_8
var __t5 gopurs_runtime.Value
{
if (v_8 == nil) {
__t5 = b_7
goto end_branch_5
} else {

}
}
{
if (v_8 != nil) {
b_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_8).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_7)}))}
v_8_loop = (v_8).V1
continue go__go_6_4_28
__t5 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_6_4_28 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_6_4_28(b_7_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_8_loop_val))
})
})
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_local_Data_List_Types_go__go_6_4_28(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_5_1).V0, memo_4})))
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}
}
go__go_2_0_27 = gopurs_runtime.Func(func(source_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_2_0_27(source_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](memo_4_loop_val)))}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_2_0_27(b_1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)))}
})})))}
	})
	return cache_Data_List_Types_unfoldable1List
}

var cache_Data_List_Types_unfoldableList gopurs_runtime.Value
var once_Data_List_Types_unfoldableList sync.Once
func Get_Data_List_Types_unfoldableList() gopurs_runtime.Value {
	once_Data_List_Types_unfoldableList.Do(func() {
		cache_Data_List_Types_unfoldableList = gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_132451362_2738507278((&Constructor_Data_Unfoldable_Unfoldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_4073635714_2187088110(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_unfoldable1List())))}
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_2_0_29 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_2_0_29
var go__go_2_0_29 gopurs_runtime.Value
_ = go__go_2_0_29
Call_local_Data_List_Types_go__go_2_0_29 = func(source_3_loop gopurs_runtime.Value, memo_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_2_0_29:
for {
if false { continue go__go_2_0_29 }
var source_3 gopurs_runtime.Value = source_3_loop
_ = source_3
var memo_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = memo_4_loop
_ = memo_4
// TAST (Let): v_5_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar b)])])
v_5_1 := Rebox_Data_List_Types_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, source_3)))
_ = v_5_1
var __t4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v_5_1 == nil) {
var Call_local_Data_List_Types_go__go_6_2_30 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_6_2_30
var go__go_6_2_30 gopurs_runtime.Value
_ = go__go_6_2_30
Call_local_Data_List_Types_go__go_6_2_30 = func(b_7_loop gopurs_runtime.Value, v_8_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_6_2_30:
for {
if false { continue go__go_6_2_30 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var v_8 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_8_loop
_ = v_8
var __t3 gopurs_runtime.Value
{
if (v_8 == nil) {
__t3 = b_7
goto end_branch_3
} else {

}
}
{
if (v_8 != nil) {
b_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_8).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_7)}))}
v_8_loop = (v_8).V1
continue go__go_6_2_30
__t3 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_6_2_30 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_6_2_30(b_7_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_8_loop_val))
})
})
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_local_Data_List_Types_go__go_6_2_30(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, memo_4))
goto end_branch_4
} else {

}
}
{
if (v_5_1 != nil) {
source_3_loop = ((v_5_1).V0).V1
memo_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, ((v_5_1).V0).V0, memo_4})
continue go__go_2_0_29
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}
go__go_2_0_29 = gopurs_runtime.Func(func(source_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_2_0_29(source_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](memo_4_loop_val)))}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_2_0_29(b_1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)))}
})})))}
	})
	return cache_Data_List_Types_unfoldableList
}

var cache_Data_List_Types_unfoldable1NonEmptyList gopurs_runtime.Value
var once_Data_List_Types_unfoldable1NonEmptyList sync.Once
func Get_Data_List_Types_unfoldable1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_unfoldable1NonEmptyList.Do(func() {
		cache_Data_List_Types_unfoldable1NonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2435023311_2187088110((&Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 shape=App(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (ADT ["Data","Maybe","Maybe"] [(TypeVar b)])])
__local_var_2_1 := Rebox_Data_List_Types_138441832_3804580809(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, b_1)))
_ = __local_var_2_1
var Call_local_Data_List_Types_go__go_3_2_31 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_3_2_31
var go__go_3_2_31 gopurs_runtime.Value
_ = go__go_3_2_31
Call_local_Data_List_Types_go__go_3_2_31 = func(source_4_loop gopurs_runtime.Value, memo_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_3_2_31:
for {
if false { continue go__go_3_2_31 }
var source_4 gopurs_runtime.Value = source_4_loop
_ = source_4
var memo_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = memo_5_loop
_ = memo_5
var __t5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (source_4.Type == 9 && source_4.IntVal == 930809136 && source_4.UnsafePtr != nil) {
source_4_loop = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(f_0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(source_4.UnsafePtr).V0).UnsafePtr).V1
memo_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(f_0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(source_4.UnsafePtr).V0).UnsafePtr).V0, memo_5})
continue go__go_3_2_31
__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_5
} else {

}
}
{
var Call_local_Data_List_Types_go__go_6_3_32 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_6_3_32
var go__go_6_3_32 gopurs_runtime.Value
_ = go__go_6_3_32
Call_local_Data_List_Types_go__go_6_3_32 = func(b_7_loop gopurs_runtime.Value, v_8_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_6_3_32:
for {
if false { continue go__go_6_3_32 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var v_8 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_8_loop
_ = v_8
var __t4 gopurs_runtime.Value
{
if (v_8 == nil) {
__t4 = b_7
goto end_branch_4
} else {

}
}
{
if (v_8 != nil) {
b_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_8).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_7)}))}
v_8_loop = (v_8).V1
continue go__go_6_3_32
__t4 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_6_3_32 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_6_3_32(b_7_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_8_loop_val))
})
})
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_local_Data_List_Types_go__go_6_3_32(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, memo_5))
}
end_branch_5:
return __t5
}
}
go__go_3_2_31 = gopurs_runtime.Func(func(source_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_3_2_31(source_4_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](memo_5_loop_val)))}
})
})
// TAST (Let): __local_var_2_0 shape=Let(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeApp (TypeVar f) [(TypeVar a)])])
__local_var_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(__local_var_2_1).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_3_2_31(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((__local_var_2_1).V1)}, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())
_ = __local_var_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (__local_var_2_0).V0, (__local_var_2_0).V1}))}
})})))}
	})
	return cache_Data_List_Types_unfoldable1NonEmptyList
}

var cache_Data_List_Types_foldable1NonEmptyList gopurs_runtime.Value
var once_Data_List_Types_foldable1NonEmptyList sync.Once
func Get_Data_List_Types_foldable1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_foldable1NonEmptyList.Do(func() {
		cache_Data_List_Types_foldable1NonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3368202604_4151366573(Rebox_Data_List_Types_4151366573_3368202604(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_NonEmpty_foldable1NonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1022383170_1680800814(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_foldableList())))})))))}
	})
	return cache_Data_List_Types_foldable1NonEmptyList
}

var cache_Data_List_Types_extendNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_extendNonEmptyList sync.Once
func Get_Data_List_Types_extendNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_extendNonEmptyList.Do(func() {
		cache_Data_List_Types_extendNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2476964120_3290176857((&Constructor_Control_Extend_Extend[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2801299215_2812149806(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_functorNonEmptyList())))}
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_2_0_33 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_2_0_33
var go__go_2_0_33 gopurs_runtime.Value
_ = go__go_2_0_33
Call_local_Data_List_Types_go__go_2_0_33 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_0_33:
for {
if false { continue go__go_2_0_33 }
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
b_3_loop = func() gopurs_runtime.Value {
				orig := struct{
	acc *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	val *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}{(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_4).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(b_3, "acc"))}), (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952((&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, (v_4).V0, gopurs_runtime.RecordGet(b_3, "acc")})))}), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(b_3, "val"))})}
				_ = orig
				return gopurs_runtime.RecordDict([]string{"acc", "val"}, []gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.acc)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.val)}})
				}()
v_4_loop = (v_4).V1
continue go__go_2_0_33
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
go__go_2_0_33 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_2_0_33(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
var Call_local_Data_List_Types_go__go_3_2_34 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_3_2_34
var go__go_3_2_34 gopurs_runtime.Value
_ = go__go_3_2_34
Call_local_Data_List_Types_go__go_3_2_34 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_3_2_34:
for {
if false { continue go__go_3_2_34 }
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_5_loop
_ = v1_5
var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_5 == nil) {
__t3 = v_4
goto end_branch_3
} else {

}
}
{
if (v1_5 != nil) {
v_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_5).V0, v_4})
v1_5_loop = (v1_5).V1
continue go__go_3_2_34
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}
go__go_3_2_34 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_3_2_34(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val)))}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952((&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952(Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_1))))}), gopurs_runtime.RecordGet(Call_local_Data_List_Types_go__go_2_0_33(func() gopurs_runtime.Value {
				orig := struct{
	acc *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	val *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}{(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)}
				_ = orig
				return gopurs_runtime.RecordDict([]string{"acc", "val"}, []gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.acc)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.val)}})
				}(), Call_local_Data_List_Types_go__go_3_2_34((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1))), "val")})))}
})})))}
	})
	return cache_Data_List_Types_extendNonEmptyList
}

var cache_Data_List_Types_extendList gopurs_runtime.Value
var once_Data_List_Types_extendList sync.Once
func Get_Data_List_Types_extendList() gopurs_runtime.Value {
	once_Data_List_Types_extendList.Do(func() {
		cache_Data_List_Types_extendList = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2171623093_3290176857((&Constructor_Control_Extend_Extend[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3262795586_2812149806(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_functorList())))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1)
if (__t_tag_0 == nil) {
__t6 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_6
} else {

}
}
{
var __t_tag_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1)
if (__t_tag_1 != nil) {
var Call_local_Data_List_Types_go__go_2_2_35 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_2_2_35
var go__go_2_2_35 gopurs_runtime.Value
_ = go__go_2_2_35
Call_local_Data_List_Types_go__go_2_2_35 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_2_35:
for {
if false { continue go__go_2_2_35 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var __t3 gopurs_runtime.Value
{
if (v_4 == nil) {
__t3 = b_3
goto end_branch_3
} else {

}
}
{
if (v_4 != nil) {
b_3_loop = func() gopurs_runtime.Value {
				orig := struct{
	acc *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	val *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}{(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_4).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(b_3, "acc"))}), (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_4).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(b_3, "acc"))}))}), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(b_3, "val"))})}
				_ = orig
				return gopurs_runtime.RecordDict([]string{"acc", "val"}, []gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.acc)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.val)}})
				}()
v_4_loop = (v_4).V1
continue go__go_2_2_35
__t3 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_2_2_35 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_2_2_35(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
var Call_local_Data_List_Types_go__go_3_4_36 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_3_4_36
var go__go_3_4_36 gopurs_runtime.Value
_ = go__go_3_4_36
Call_local_Data_List_Types_go__go_3_4_36 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_3_4_36:
for {
if false { continue go__go_3_4_36 }
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_5_loop
_ = v1_5
var __t5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_5 == nil) {
__t5 = v_4
goto end_branch_5
} else {

}
}
{
if (v1_5 != nil) {
v_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_5).V0, v_4})
v1_5_loop = (v1_5).V1
continue go__go_3_4_36
__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}
go__go_3_4_36 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_3_4_36(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val)))}
})
})
__t6 = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1))}), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(Call_local_Data_List_Types_go__go_2_2_35(func() gopurs_runtime.Value {
				orig := struct{
	acc *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	val *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}{(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)}
				_ = orig
				return gopurs_runtime.RecordDict([]string{"acc", "val"}, []gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.acc)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.val)}})
				}(), Call_local_Data_List_Types_go__go_3_4_36((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_1.UnsafePtr).V1)), "val"))})
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t6)}
})})))}
	})
	return cache_Data_List_Types_extendList
}

var cache_Data_List_Types_eq1List gopurs_runtime.Value
var once_Data_List_Types_eq1List sync.Once
func Get_Data_List_Types_eq1List() gopurs_runtime.Value {
	once_Data_List_Types_eq1List.Do(func() {
		cache_Data_List_Types_eq1List = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3554500787_1766074591((&Constructor_Data_Eq_Eq1[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func3(func(dictEq_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value, ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_0_37 gopurs_runtime.Value
_ = go__go_3_0_37
// FALLBACK TCO: isLoop=false len=1
go__go_3_0_37 = gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 bool
{
if ((v2_6.IntVal) != (0)) != (true) {
__t6 = false
goto end_branch_6
} else {

}
}
{
var __t_tag_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4)
if (__t_tag_4 == nil) {
var __t_tag_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5)
__t6 = ((__t_tag_5 == nil)) && ((v2_6.IntVal) != (0))
goto end_branch_6
} else {

}
}
{
var __t_tag_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4)
var __t_and_3 bool = false
if (__t_tag_1 != nil) {

var __t_tag_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5)
__t_and_3 = ((__t_tag_2 != nil)) && ((gopurs_runtime.Apply3(go__go_3_0_37, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}, gopurs_runtime.Bool(((v2_6.IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0).IntVal) != (0)))).IntVal) != (0))
}
__t6 = __t_and_3
}
end_branch_6:
return gopurs_runtime.Bool(__t6)
})
return gopurs_runtime.Bool((gopurs_runtime.Apply3(go__go_3_0_37, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_2))}, gopurs_runtime.Bool(true)).IntVal) != (0))
})})))}
	})
	return cache_Data_List_Types_eq1List
}

var cache_Data_List_Types_eq1 gopurs_runtime.Value
var once_Data_List_Types_eq1 sync.Once
func Get_Data_List_Types_eq1() gopurs_runtime.Value {
	once_Data_List_Types_eq1.Do(func() {
		cache_Data_List_Types_eq1 = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_List_Types_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_2_box)))
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
		cache_Data_List_Types_eq1NonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2625657118_1766074591((&Constructor_Data_Eq_Eq1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func3(func(dictEq_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_and_7 bool = false
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal) != (0) {

var go__go_3_0_40 gopurs_runtime.Value
_ = go__go_3_0_40
// FALLBACK TCO: isLoop=false len=1
go__go_3_0_40 = gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 bool
{
if ((v2_6.IntVal) != (0)) != (true) {
__t6 = false
goto end_branch_6
} else {

}
}
{
var __t_tag_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4)
if (__t_tag_4 == nil) {
var __t_tag_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5)
__t6 = ((__t_tag_5 == nil)) && ((v2_6.IntVal) != (0))
goto end_branch_6
} else {

}
}
{
var __t_tag_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4)
var __t_and_3 bool = false
if (__t_tag_1 != nil) {

var __t_tag_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5)
__t_and_3 = ((__t_tag_2 != nil)) && ((gopurs_runtime.Apply3(go__go_3_0_40, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}, gopurs_runtime.Bool(((v2_6.IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0).IntVal) != (0)))).IntVal) != (0))
}
__t6 = __t_and_3
}
end_branch_6:
return gopurs_runtime.Bool(__t6)
})
__t_and_7 = (gopurs_runtime.Apply3(go__go_3_0_40, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V1))}, gopurs_runtime.Bool(true)).IntVal) != (0)
}
return gopurs_runtime.Bool(__t_and_7)
})})))}
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
		cache_Data_List_Types_ord1List = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3739089555_3985601471((&Constructor_Data_Ord_Ord1[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3554500787_1766074591(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_eq1List())))}
}), gopurs_runtime.Func3(func(dictOrd_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value, ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_3_0_43 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) uint32
_ = Call_local_Data_List_Types_go__go_3_0_43
var go__go_3_0_43 gopurs_runtime.Value
_ = go__go_3_0_43
Call_local_Data_List_Types_go__go_3_0_43 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) uint32 {
go__go_3_0_43:
for {
if false { continue go__go_3_0_43 }
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_5_loop
_ = v1_5
var __t4 uint32
{
if (v_4 == nil) {
var __t1 uint32
{
if (v1_5 == nil) {
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
if (v1_5 == nil) {
__t4 = 380165415
goto end_branch_4
} else {

}
}
{
if ((v_4 != nil)) && ((v1_5 != nil)) {
// TAST (Let): v2_6_2 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v2_6_2 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (v_4).V0, (v1_5).V0).IntVal)
_ = v2_6_2
var __t3 uint32
{
if (v2_6_2 == 902936544) {
v_4_loop = (v_4).V1
v1_5_loop = (v1_5).V1
continue go__go_3_0_43
__t3 = func() uint32 { panic("unreachable") }()
goto end_branch_3
} else {

}
}
{
__t3 = v2_6_2
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}
go__go_3_0_43 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_List_Types_go__go_3_0_43(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val))), UnsafePtr: nil}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_List_Types_go__go_3_0_43(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_2))), UnsafePtr: nil}
})})))}
	})
	return cache_Data_List_Types_ord1List
}

var cache_Data_List_Types_compare1 gopurs_runtime.Value
var once_Data_List_Types_compare1 sync.Once
func Get_Data_List_Types_compare1() gopurs_runtime.Value {
	once_Data_List_Types_compare1.Do(func() {
		cache_Data_List_Types_compare1 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_List_Types_compare1(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_2_box))), UnsafePtr: nil}
})
	})
	return cache_Data_List_Types_compare1
}

var cache_Data_List_Types_ordNonEmpty gopurs_runtime.Value
var once_Data_List_Types_ordNonEmpty sync.Once
func Get_Data_List_Types_ordNonEmpty() gopurs_runtime.Value {
	once_Data_List_Types_ordNonEmpty.Do(func() {
		cache_Data_List_Types_ordNonEmpty = gopurs_runtime.Apply(Get_Data_NonEmpty_ordNonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3739089555_3985601471(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_ord1List())))})
	})
	return cache_Data_List_Types_ordNonEmpty
}

var cache_Data_List_Types_ord1NonEmptyList gopurs_runtime.Value
var once_Data_List_Types_ord1NonEmptyList sync.Once
func Get_Data_List_Types_ord1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_ord1NonEmptyList.Do(func() {
		cache_Data_List_Types_ord1NonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3290489278_3985601471(Rebox_Data_List_Types_3985601471_3290489278(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_NonEmpty_ord1NonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3739089555_3985601471(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_ord1List())))})))))}
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
		cache_Data_List_Types_comonadNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1262583736_2550391993((&Constructor_Control_Comonad_Comonad[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2476964120_3290176857(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_extendNonEmptyList())))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0
})})))}
	})
	return cache_Data_List_Types_comonadNonEmptyList
}

var cache_Data_List_Types_applyList gopurs_runtime.Value
var once_Data_List_Types_applyList sync.Once
func Get_Data_List_Types_applyList() gopurs_runtime.Value {
	once_Data_List_Types_applyList.Do(func() {
		cache_Data_List_Types_applyList = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1305434581_3741347833((&Constructor_Control_Apply_Apply[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3262795586_2812149806(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_functorList())))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0)
if (__t_tag_0 == nil) {
__t6 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_6
} else {

}
}
{
var __t_tag_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0)
if (__t_tag_1 != nil) {
var Call_local_Data_List_Types_go__go_2_2_47 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_2_2_47
var go__go_2_2_47 gopurs_runtime.Value
_ = go__go_2_2_47
Call_local_Data_List_Types_go__go_2_2_47 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_2_47:
for {
if false { continue go__go_2_2_47 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var __t3 gopurs_runtime.Value
{
if (v_4 == nil) {
__t3 = b_3
goto end_branch_3
} else {

}
}
{
if (v_4 != nil) {
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_4).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_3)}))}
v_4_loop = (v_4).V1
continue go__go_2_2_47
__t3 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_2_2_47 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_2_2_47(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
var Call_local_Data_List_Types_go__go_3_4_48 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_3_4_48
var go__go_3_4_48 gopurs_runtime.Value
_ = go__go_3_4_48
Call_local_Data_List_Types_go__go_3_4_48 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_3_4_48:
for {
if false { continue go__go_3_4_48 }
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_5_loop
_ = v1_5
var __t5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_5 == nil) {
__t5 = v_4
goto end_branch_5
} else {

}
}
{
if (v1_5 != nil) {
v_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_5).V0, v_4})
v1_5_loop = (v1_5).V1
continue go__go_3_4_48
__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}
go__go_3_4_48 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_3_4_48(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val)))}
})
})
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_local_Data_List_Types_go__go_2_2_47(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_applyList()).V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1))})))}, Call_local_Data_List_Types_go__go_3_4_48((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_List_Types_listMap((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1))})))))
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t6)}
})})))}
	})
	return cache_Data_List_Types_applyList
}

var cache_Data_List_Types_applyNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_applyNonEmptyList sync.Once
func Get_Data_List_Types_applyNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_applyNonEmptyList.Do(func() {
		cache_Data_List_Types_applyNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3970790648_3741347833((&Constructor_Control_Apply_Apply[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2801299215_2812149806(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_functorNonEmptyList())))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_2_0_49 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_2_0_49
var go__go_2_0_49 gopurs_runtime.Value
_ = go__go_2_0_49
Call_local_Data_List_Types_go__go_2_0_49 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_0_49:
for {
if false { continue go__go_2_0_49 }
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
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_4).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_3)}))}
v_4_loop = (v_4).V1
continue go__go_2_0_49
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
go__go_2_0_49 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_2_0_49(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
var Call_local_Data_List_Types_go__go_3_2_50 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_3_2_50
var go__go_3_2_50 gopurs_runtime.Value
_ = go__go_3_2_50
Call_local_Data_List_Types_go__go_3_2_50 = func(b_4_loop gopurs_runtime.Value, v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_3_2_50:
for {
if false { continue go__go_3_2_50 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var __t3 gopurs_runtime.Value
{
if (v_5 == nil) {
__t3 = b_4
goto end_branch_3
} else {

}
}
{
if (v_5 != nil) {
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_5).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_4)}))}
v_5_loop = (v_5).V1
continue go__go_3_2_50
__t3 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_3_2_50 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_3_2_50(b_4_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val))
})
})
var Call_local_Data_List_Types_go__go_4_4_51 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_4_4_51
var go__go_4_4_51 gopurs_runtime.Value
_ = go__go_4_4_51
Call_local_Data_List_Types_go__go_4_4_51 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_4_4_51:
for {
if false { continue go__go_4_4_51 }
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var __t5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_6 == nil) {
__t5 = v_5
goto end_branch_5
} else {

}
}
{
if (v1_6 != nil) {
v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
v1_6_loop = (v1_6).V1
continue go__go_4_4_51
__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}
go__go_4_4_51 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_4_4_51(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
})
})
var Call_local_Data_List_Types_go__go_3_6_52 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_3_6_52
var go__go_3_6_52 gopurs_runtime.Value
_ = go__go_3_6_52
Call_local_Data_List_Types_go__go_3_6_52 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_3_6_52:
for {
if false { continue go__go_3_6_52 }
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_5_loop
_ = v1_5
var __t7 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_5 == nil) {
__t7 = v_4
goto end_branch_7
} else {

}
}
{
if (v1_5 != nil) {
v_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_5).V0, v_4})
v1_5_loop = (v1_5).V1
continue go__go_3_6_52
__t7 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_7
} else {

}
}
{
__t7 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}
}
go__go_3_6_52 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_3_6_52(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val)))}
})
})
var __t14 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_8 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1)
if (__t_tag_8 == nil) {
__t14 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_14
} else {

}
}
{
var __t_tag_9 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1)
if (__t_tag_9 != nil) {
var Call_local_Data_List_Types_go__go_4_10_53 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_4_10_53
var go__go_4_10_53 gopurs_runtime.Value
_ = go__go_4_10_53
Call_local_Data_List_Types_go__go_4_10_53 = func(b_5_loop gopurs_runtime.Value, v_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_4_10_53:
for {
if false { continue go__go_4_10_53 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_6_loop
_ = v_6
var __t11 gopurs_runtime.Value
{
if (v_6 == nil) {
__t11 = b_5
goto end_branch_11
} else {

}
}
{
if (v_6 != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_6).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_5)}))}
v_6_loop = (v_6).V1
continue go__go_4_10_53
__t11 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_4_10_53 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_4_10_53(b_5_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_6_loop_val))
})
})
var Call_local_Data_List_Types_go__go_5_12_54 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_5_12_54
var go__go_5_12_54 gopurs_runtime.Value
_ = go__go_5_12_54
Call_local_Data_List_Types_go__go_5_12_54 = func(v_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_7_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_5_12_54:
for {
if false { continue go__go_5_12_54 }
var v_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_6_loop
_ = v_6
var v1_7 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_7_loop
_ = v1_7
var __t13 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_7 == nil) {
__t13 = v_6
goto end_branch_13
} else {

}
}
{
if (v1_7 != nil) {
v_6_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_7).V0, v_6})
v1_7_loop = (v1_7).V1
continue go__go_5_12_54
__t13 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_13
} else {

}
}
{
__t13 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
}
}
go__go_5_12_54 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_5_12_54(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_6_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_7_loop_val)))}
})
})
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_local_Data_List_Types_go__go_4_10_53(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_applyList()).V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)}))})))}, Call_local_Data_List_Types_go__go_5_12_54((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_List_Types_listMap((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)}))})))))
goto end_branch_14
} else {

}
}
{
__t14 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_14:
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952((&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, gopurs_runtime.Apply((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_local_Data_List_Types_go__go_2_0_49(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_local_Data_List_Types_go__go_3_2_50(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_applyList()).V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1))})))}, Call_local_Data_List_Types_go__go_4_4_51((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_List_Types_listMap((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1))}))))))}, Call_local_Data_List_Types_go__go_3_6_52((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), __t14))))}})))}
})})))}
	})
	return cache_Data_List_Types_applyNonEmptyList
}

var cache_Data_List_Types_bindList gopurs_runtime.Value
var once_Data_List_Types_bindList sync.Once
func Get_Data_List_Types_bindList() gopurs_runtime.Value {
	once_Data_List_Types_bindList.Do(func() {
		cache_Data_List_Types_bindList = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2183599445_2748095225((&Constructor_Control_Bind_Bind[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1305434581_3741347833(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_applyList())))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0)
if (__t_tag_0 == nil) {
__t6 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_6
} else {

}
}
{
var __t_tag_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0)
if (__t_tag_1 != nil) {
var Call_local_Data_List_Types_go__go_2_2_55 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_2_2_55
var go__go_2_2_55 gopurs_runtime.Value
_ = go__go_2_2_55
Call_local_Data_List_Types_go__go_2_2_55 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_2_55:
for {
if false { continue go__go_2_2_55 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var __t3 gopurs_runtime.Value
{
if (v_4 == nil) {
__t3 = b_3
goto end_branch_3
} else {

}
}
{
if (v_4 != nil) {
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_4).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_3)}))}
v_4_loop = (v_4).V1
continue go__go_2_2_55
__t3 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_2_2_55 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_2_2_55(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
var Call_local_Data_List_Types_go__go_3_4_56 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_3_4_56
var go__go_3_4_56 gopurs_runtime.Value
_ = go__go_3_4_56
Call_local_Data_List_Types_go__go_3_4_56 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_3_4_56:
for {
if false { continue go__go_3_4_56 }
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_5_loop
_ = v1_5
var __t5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_5 == nil) {
__t5 = v_4
goto end_branch_5
} else {

}
}
{
if (v1_5 != nil) {
v_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_5).V0, v_4})
v1_5_loop = (v1_5).V1
continue go__go_3_4_56
__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}
go__go_3_4_56 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_3_4_56(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val)))}
})
})
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_local_Data_List_Types_go__go_2_2_55(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_bindList()).V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V1)}, v1_1)))}, Call_local_Data_List_Types_go__go_3_4_56((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V0)))))
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t6)}
})})))}
	})
	return cache_Data_List_Types_bindList
}

var cache_Data_List_Types_bindNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_bindNonEmptyList sync.Once
func Get_Data_List_Types_bindNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_bindNonEmptyList.Do(func() {
		cache_Data_List_Types_bindNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3054666744_2748095225((&Constructor_Control_Bind_Bind[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3970790648_3741347833(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_applyNonEmptyList())))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_2_0 shape=App(Other) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (TypeVar b)])
v1_2_0 := Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0)))
_ = v1_2_0
var Call_local_Data_List_Types_go__go_3_1_57 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_3_1_57
var go__go_3_1_57 gopurs_runtime.Value
_ = go__go_3_1_57
Call_local_Data_List_Types_go__go_3_1_57 = func(b_4_loop gopurs_runtime.Value, v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_3_1_57:
for {
if false { continue go__go_3_1_57 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var __t2 gopurs_runtime.Value
{
if (v_5 == nil) {
__t2 = b_4
goto end_branch_2
} else {

}
}
{
if (v_5 != nil) {
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_5).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_4)}))}
v_5_loop = (v_5).V1
continue go__go_3_1_57
__t2 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_3_1_57 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_3_1_57(b_4_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val))
})
})
var __t11 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1)
if (__t_tag_3 == nil) {
__t11 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_11
} else {

}
}
{
var __t_tag_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1)
if (__t_tag_4 != nil) {
var Call_local_Data_List_Types_go__go_4_5_58 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_4_5_58
var go__go_4_5_58 gopurs_runtime.Value
_ = go__go_4_5_58
Call_local_Data_List_Types_go__go_4_5_58 = func(b_5_loop gopurs_runtime.Value, v_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_4_5_58:
for {
if false { continue go__go_4_5_58 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_6_loop
_ = v_6
var __t6 gopurs_runtime.Value
{
if (v_6 == nil) {
__t6 = b_5
goto end_branch_6
} else {

}
}
{
if (v_6 != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_6).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_5)}))}
v_6_loop = (v_6).V1
continue go__go_4_5_58
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
go__go_4_5_58 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_4_5_58(b_5_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_6_loop_val))
})
})
var Call_local_Data_List_Types_go__go_5_8_59 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_5_8_59
var go__go_5_8_59 gopurs_runtime.Value
_ = go__go_5_8_59
Call_local_Data_List_Types_go__go_5_8_59 = func(v_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_7_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_5_8_59:
for {
if false { continue go__go_5_8_59 }
var v_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_6_loop
_ = v_6
var v1_7 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_7_loop
_ = v1_7
var __t9 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_7 == nil) {
__t9 = v_6
goto end_branch_9
} else {

}
}
{
if (v1_7 != nil) {
v_6_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_7).V0, v_6})
v1_7_loop = (v1_7).V1
continue go__go_5_8_59
__t9 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_9
} else {

}
}
{
__t9 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
}
}
go__go_5_8_59 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_5_8_59(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_6_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_7_loop_val)))}
})
})
// TAST (Let): __local_var_6_10 shape=App(Other) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (TypeVar b)])
__local_var_6_10 := Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1.UnsafePtr).V0)))
_ = __local_var_6_10
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_local_Data_List_Types_go__go_4_5_58(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_bindList()).V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1.UnsafePtr).V1)}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_7 shape=App(Other) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (TypeVar b)])
__local_var_6_7 := Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, x_5)))
_ = __local_var_6_7
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (__local_var_6_7).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((__local_var_6_7).V1)}))}
}))))}, Call_local_Data_List_Types_go__go_5_8_59((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (__local_var_6_10).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((__local_var_6_10).V1)}))))
goto end_branch_11
} else {

}
}
{
__t11 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_11:
var Call_local_Data_List_Types_go__go_4_12_60 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_4_12_60
var go__go_4_12_60 gopurs_runtime.Value
_ = go__go_4_12_60
Call_local_Data_List_Types_go__go_4_12_60 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_4_12_60:
for {
if false { continue go__go_4_12_60 }
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var __t13 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_6 == nil) {
__t13 = v_5
goto end_branch_13
} else {

}
}
{
if (v1_6 != nil) {
v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
v1_6_loop = (v1_6).V1
continue go__go_4_12_60
__t13 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_13
} else {

}
}
{
__t13 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
}
}
go__go_4_12_60 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_4_12_60(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952((&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, (v1_2_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_local_Data_List_Types_go__go_3_1_57(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t11)}, Call_local_Data_List_Types_go__go_4_12_60((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v1_2_0).V1)))))}})))}
})})))}
	})
	return cache_Data_List_Types_bindNonEmptyList
}

var cache_Data_List_Types_applicativeList gopurs_runtime.Value
var once_Data_List_Types_applicativeList sync.Once
func Get_Data_List_Types_applicativeList() gopurs_runtime.Value {
	once_Data_List_Types_applicativeList.Do(func() {
		cache_Data_List_Types_applicativeList = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2870828117_1439734649((&Constructor_Control_Applicative_Applicative[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1305434581_3741347833(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_applyList())))}
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)}))}
})})))}
	})
	return cache_Data_List_Types_applicativeList
}

var cache_Data_List_Types_monadList gopurs_runtime.Value
var once_Data_List_Types_monadList sync.Once
func Get_Data_List_Types_monadList() gopurs_runtime.Value {
	once_Data_List_Types_monadList.Do(func() {
		cache_Data_List_Types_monadList = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3994390613_2568689657((&Constructor_Control_Monad_Monad[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2870828117_1439734649(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_applicativeList())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2183599445_2748095225(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_bindList())))}
})})))}
	})
	return cache_Data_List_Types_monadList
}

var cache_Data_List_Types_altNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_altNonEmptyList sync.Once
func Get_Data_List_Types_altNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_altNonEmptyList.Do(func() {
		cache_Data_List_Types_altNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_203237368_3421983481((&Constructor_Control_Alt_Alt[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2801299215_2812149806(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_functorNonEmptyList())))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, as_prime__1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_2_0_61 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_2_0_61
var go__go_2_0_61 gopurs_runtime.Value
_ = go__go_2_0_61
Call_local_Data_List_Types_go__go_2_0_61 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_0_61:
for {
if false { continue go__go_2_0_61 }
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
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_4).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_3)}))}
v_4_loop = (v_4).V1
continue go__go_2_0_61
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
go__go_2_0_61 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_2_0_61(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
var Call_local_Data_List_Types_go__go_3_2_62 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_3_2_62
var go__go_3_2_62 gopurs_runtime.Value
_ = go__go_3_2_62
Call_local_Data_List_Types_go__go_3_2_62 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_3_2_62:
for {
if false { continue go__go_3_2_62 }
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_5_loop
_ = v1_5
var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_5 == nil) {
__t3 = v_4
goto end_branch_3
} else {

}
}
{
if (v1_5 != nil) {
v_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_5).V0, v_4})
v1_5_loop = (v1_5).V1
continue go__go_3_2_62
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}
go__go_3_2_62 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_3_2_62(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val)))}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952((&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_local_Data_List_Types_go__go_2_0_61(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(as_prime__1.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(as_prime__1.UnsafePtr).V1)}))}, Call_local_Data_List_Types_go__go_3_2_62((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1)))))}})))}
})})))}
	})
	return cache_Data_List_Types_altNonEmptyList
}

var cache_Data_List_Types_altList gopurs_runtime.Value
var once_Data_List_Types_altList sync.Once
func Get_Data_List_Types_altList() gopurs_runtime.Value {
	once_Data_List_Types_altList.Do(func() {
		cache_Data_List_Types_altList = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_257347157_3421983481((&Constructor_Control_Alt_Alt[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3262795586_2812149806(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_functorList())))}
}), gopurs_runtime.Func2(func(xs_0 gopurs_runtime.Value, ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_2_0_63 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_2_0_63
var go__go_2_0_63 gopurs_runtime.Value
_ = go__go_2_0_63
Call_local_Data_List_Types_go__go_2_0_63 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_0_63:
for {
if false { continue go__go_2_0_63 }
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
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_4).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_3)}))}
v_4_loop = (v_4).V1
continue go__go_2_0_63
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
go__go_2_0_63 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_2_0_63(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
var Call_local_Data_List_Types_go__go_3_2_64 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_3_2_64
var go__go_3_2_64 gopurs_runtime.Value
_ = go__go_3_2_64
Call_local_Data_List_Types_go__go_3_2_64 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_3_2_64:
for {
if false { continue go__go_3_2_64 }
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_5_loop
_ = v1_5
var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_5 == nil) {
__t3 = v_4
goto end_branch_3
} else {

}
}
{
if (v1_5 != nil) {
v_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_5).V0, v_4})
v1_5_loop = (v1_5).V1
continue go__go_3_2_64
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}
go__go_3_2_64 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_3_2_64(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val)))}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_local_Data_List_Types_go__go_2_0_63(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_1))}, Call_local_Data_List_Types_go__go_3_2_64((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_0)))))}
})})))}
	})
	return cache_Data_List_Types_altList
}

var cache_Data_List_Types_plusList gopurs_runtime.Value
var once_Data_List_Types_plusList sync.Once
func Get_Data_List_Types_plusList() gopurs_runtime.Value {
	once_Data_List_Types_plusList.Do(func() {
		cache_Data_List_Types_plusList = gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2173899317_3706288089((&Constructor_Control_Plus_Plus[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_257347157_3421983481(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_altList())))}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}})))}
	})
	return cache_Data_List_Types_plusList
}

var cache_Data_List_Types_alternativeList gopurs_runtime.Value
var once_Data_List_Types_alternativeList sync.Once
func Get_Data_List_Types_alternativeList() gopurs_runtime.Value {
	once_Data_List_Types_alternativeList.Do(func() {
		cache_Data_List_Types_alternativeList = gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_739440277_2307501113((&Constructor_Control_Alternative_Alternative[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2870828117_1439734649(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_applicativeList())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2173899317_3706288089(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_plusList())))}
})})))}
	})
	return cache_Data_List_Types_alternativeList
}

var cache_Data_List_Types_monadPlusList gopurs_runtime.Value
var once_Data_List_Types_monadPlusList sync.Once
func Get_Data_List_Types_monadPlusList() gopurs_runtime.Value {
	once_Data_List_Types_monadPlusList.Do(func() {
		cache_Data_List_Types_monadPlusList = gopurs_runtime.Value{Type: 9, IntVal: 3236234573, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2656832021_2394491833((&Constructor_Control_MonadPlus_MonadPlus[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_739440277_2307501113(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_alternativeList())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3994390613_2568689657(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_monadList())))}
})})))}
	})
	return cache_Data_List_Types_monadPlusList
}

var cache_Data_List_Types_applicativeNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_applicativeNonEmptyList sync.Once
func Get_Data_List_Types_applicativeNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_applicativeNonEmptyList.Do(func() {
		cache_Data_List_Types_applicativeNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2334388216_1439734649((&Constructor_Control_Applicative_Applicative[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3970790648_3741347833(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_applyNonEmptyList())))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}}))}
})})))}
	})
	return cache_Data_List_Types_applicativeNonEmptyList
}

var cache_Data_List_Types_pure gopurs_runtime.Value
var once_Data_List_Types_pure sync.Once
func Get_Data_List_Types_pure() gopurs_runtime.Value {
	once_Data_List_Types_pure.Do(func() {
		cache_Data_List_Types_pure = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952(Call_Data_List_Types_pure(x_0_box)))}
})
	})
	return cache_Data_List_Types_pure
}

var cache_Data_List_Types_monadNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_monadNonEmptyList sync.Once
func Get_Data_List_Types_monadNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_monadNonEmptyList.Do(func() {
		cache_Data_List_Types_monadNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1757775096_2568689657((&Constructor_Control_Monad_Monad[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2334388216_1439734649(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_applicativeNonEmptyList())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3054666744_2748095225(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_bindNonEmptyList())))}
})})))}
	})
	return cache_Data_List_Types_monadNonEmptyList
}

var cache_Data_List_Types_traversable1NonEmptyList gopurs_runtime.Value
var once_Data_List_Types_traversable1NonEmptyList sync.Once
func Get_Data_List_Types_traversable1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_traversable1NonEmptyList.Do(func() {
		cache_Data_List_Types_traversable1NonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 1596088409, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3842311788_306175789((&Constructor_Data_Semigroup_Traversable_Traversable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3368202604_4151366573(Rebox_Data_List_Types_4151366573_3368202604(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_NonEmpty_foldable1NonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1022383170_1680800814(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_foldableList())))})))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3425342735_3043886126(Rebox_Data_List_Types_3043886126_3425342735(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_NonEmpty_traversableNonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3037784642_3043886126(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_traversableList())))})))))}
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_traversable1NonEmptyList()).V3), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0))}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_4_3_66 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_4_3_66
var go__go_4_3_66 gopurs_runtime.Value
_ = go__go_4_3_66
Call_local_Data_List_Types_go__go_4_3_66 = func(b_5_loop gopurs_runtime.Value, v_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_4_3_66:
for {
if false { continue go__go_4_3_66 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_6_loop
_ = v_6
var __t4 gopurs_runtime.Value
{
if (v_6 == nil) {
__t4 = b_5
goto end_branch_4
} else {

}
}
{
if (v_6 != nil) {
b_5_loop = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func2(func(b_7 gopurs_runtime.Value, a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952((&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, a_8, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_7.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_7.UnsafePtr).V1)}))}})))}
}), b_5), gopurs_runtime.Apply(f_2, (v_6).V0))
v_6_loop = (v_6).V1
continue go__go_4_3_66
__t4 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_4_3_66 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_4_3_66(b_5_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_6_loop_val))
})
})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_5_1_65 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_5_1_65
var go__go_5_1_65 gopurs_runtime.Value
_ = go__go_5_1_65
Call_local_Data_List_Types_go__go_5_1_65 = func(b_6_loop gopurs_runtime.Value, v_7_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_5_1_65:
for {
if false { continue go__go_5_1_65 }
var b_6 gopurs_runtime.Value = b_6_loop
_ = b_6
var v_7 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_7_loop
_ = v_7
var __t2 gopurs_runtime.Value
{
if (v_7 == nil) {
__t2 = b_6
goto end_branch_2
} else {

}
}
{
if (v_7 != nil) {
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952((&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, (v_7).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_6.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_6.UnsafePtr).V1)}))}})))}
v_7_loop = (v_7).V1
continue go__go_5_1_65
__t2 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_5_1_65 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_5_1_65(b_6_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_7_loop_val))
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952(Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](Call_local_Data_List_Types_go__go_5_1_65(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}}))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1))))))}
}), Call_local_Data_List_Types_go__go_4_3_66(gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_5, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}}))}
}), gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)))
})
})})))}
	})
	return cache_Data_List_Types_traversable1NonEmptyList
}

type Constructor_Data_List_Types_Nil[T_a any] struct {
	Rc uint32
}


type Constructor_Data_List_Types_Cons[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 *Constructor_Data_List_Types_Cons[T_a]
}


func Call_Data_List_Types_Cons__1062349399(__eta_norm_1_0_loop *Constructor_Data_Tuple_Tuple[uint32, float64], __eta_norm_0_1_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
Cons__1062349399:
for {
if false { continue Cons__1062349399 }
var __eta_norm_1_0 *Constructor_Data_Tuple_Tuple[uint32, float64] = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return Rebox_Data_List_Types_849153993_2442833393(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](value1)}))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3132786365_138441832(__eta_norm_1_0))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2442833393_849153993(__eta_norm_0_1))})))
}
}

func Call_Data_List_Types_Cons__2126472648(__eta_norm_1_0_loop int64, __eta_norm_0_1_loop *Constructor_Data_List_Types_Cons[int64]) *Constructor_Data_List_Types_Cons[int64] {
Cons__2126472648:
for {
if false { continue Cons__2126472648 }
var __eta_norm_1_0 int64 = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 *Constructor_Data_List_Types_Cons[int64] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return Rebox_Data_List_Types_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](value1)}))}
})
}), gopurs_runtime.Int(__eta_norm_1_0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3704040722_849153993(__eta_norm_0_1))})))
}
}

func Call_Data_List_Types_Cons__2681173055(__eta_norm_0_unused_0_loop int64, __eta_norm_1_1_loop *Constructor_Data_List_Types_Cons[int64]) *Constructor_Data_List_Types_Cons[int64] {
Cons__2681173055:
for {
if false { continue Cons__2681173055 }
var __eta_norm_0_unused_0 int64 = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
var __eta_norm_1_1 *Constructor_Data_List_Types_Cons[int64] = __eta_norm_1_1_loop
_ = __eta_norm_1_1
return Rebox_Data_List_Types_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](value1)}))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3704040722_849153993(__eta_norm_1_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3704040722_849153993((*Constructor_Data_List_Types_Cons[int64])(nil)))})))
}
}

func Call_Data_List_Types_NonEmptyList(x_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var x_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = x_0_loop
_ = x_0
return x_0
}

func Call_Data_List_Types_NonEmptyList__3166204772(x_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
NonEmptyList__3166204772:
for {
if false { continue NonEmptyList__3166204772 }
var x_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = x_0_loop
_ = x_0
return x_0
}
}

func Call_Data_List_Types_toList(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
var v_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
return (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_0).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_0).V1)})
}

func Call_Data_List_Types_nelCons(a_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
return (&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v_1).V1)}))}})
}

func Call_Data_List_Types_listMap(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var Call_local_Data_List_Types_chunkedRevMap_1_0_0 func(*Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_chunkedRevMap_1_0_0
var chunkedRevMap_1_0_0 gopurs_runtime.Value
_ = chunkedRevMap_1_0_0
Call_local_Data_List_Types_chunkedRevMap_1_0_0 = func(v_2_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]], v1_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
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
if (v1_3 != nil) {

var __t_tag_15 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = (v1_3).V1
var __t_and_17 bool = false
if (__t_tag_15 != nil) {

var __t_tag_16 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = ((v1_3).V1).V1
__t_and_17 = (__t_tag_16 != nil)
}
__t_and_18 = __t_and_17
}
if __t_and_18 {
v_2_loop = (&Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, v1_3, v_2})
v1_3_loop = (((v1_3).V1).V1).V1
continue chunkedRevMap_1_0_0
__t19 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_19
} else {

}
}
{
var Call_local_Data_List_Types_reverseUnrolledMap_4_1_1 func(*Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_reverseUnrolledMap_4_1_1
var reverseUnrolledMap_4_1_1 gopurs_runtime.Value
_ = reverseUnrolledMap_4_1_1
Call_local_Data_List_Types_reverseUnrolledMap_4_1_1 = func(v2_5_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]], v3_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
reverseUnrolledMap_4_1_1:
for {
if false { continue reverseUnrolledMap_4_1_1 }
var v2_5 *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] = v2_5_loop
_ = v2_5
var v3_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v3_6_loop
_ = v3_6
var __t8 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_and_7 bool = false
if (v2_5 != nil) {

var __t_tag_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = (v2_5).V0
var __t_and_6 bool = false
if (__t_tag_2 != nil) {

var __t_tag_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = ((v2_5).V0).V1
var __t_and_5 bool = false
if (__t_tag_3 != nil) {

var __t_tag_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = (((v2_5).V0).V1).V1
__t_and_5 = (__t_tag_4 != nil)
}
__t_and_6 = __t_and_5
}
__t_and_7 = __t_and_6
}
if __t_and_7 {
v2_5_loop = (v2_5).V1
v3_6_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, ((v2_5).V0).V0), (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (((v2_5).V0).V1).V0), (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, ((((v2_5).V0).V1).V1).V0), v3_6})})})
continue reverseUnrolledMap_4_1_1
__t8 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_8
} else {

}
}
{
__t8 = v3_6
}
end_branch_8:
return __t8
}
}
reverseUnrolledMap_4_1_1 = gopurs_runtime.Func(func(v2_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_reverseUnrolledMap_4_1_1(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](v2_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v3_6_loop_val)))}
})
})
var __t14 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_3 != nil) {
var __t13 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_9 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = (v1_3).V1
if (__t_tag_9 != nil) {
var __t11 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_10 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = ((v1_3).V1).V1
if (__t_tag_10 == nil) {
__t11 = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (v1_3).V0), (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, ((v1_3).V1).V0), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)})})
goto end_branch_11
} else {

}
}
{
__t11 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
}
end_branch_11:
__t13 = __t11
goto end_branch_13
} else {

}
}
{
var __t_tag_12 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = (v1_3).V1
if (__t_tag_12 == nil) {
__t13 = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (v1_3).V0), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)})
goto end_branch_13
} else {

}
}
{
__t13 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
}
end_branch_13:
__t14 = __t13
goto end_branch_14
} else {

}
}
{
__t14 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
}
end_branch_14:
__t19 = Call_local_Data_List_Types_reverseUnrolledMap_4_1_1(v_2, __t14)
}
end_branch_19:
return __t19
}
}
chunkedRevMap_1_0_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_chunkedRevMap_1_0_0(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](v_2_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val)))}
})
})
return gopurs_runtime.Apply(chunkedRevMap_1_0_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1116310629_849153993((*Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(nil)))})
}

func Call_Data_List_Types_showList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): show_1_0 shape=Other bindingType=(Func [(TypeVar a)] String)
show_1_0 := gopurs_runtime.RecordGet(dictShow_0, "show")
_ = show_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3351995458_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 string
{
var __t_tag_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2)
if (__t_tag_4 == nil) {
__t5 = "Nil"
goto end_branch_5
} else {

}
}
{
var Call_local_Data_List_Types_go__go_3_1_21 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_3_1_21
var go__go_3_1_21 gopurs_runtime.Value
_ = go__go_3_1_21
Call_local_Data_List_Types_go__go_3_1_21 = func(b_4_loop gopurs_runtime.Value, v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_3_1_21:
for {
if false { continue go__go_3_1_21 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var __t3 gopurs_runtime.Value
{
if (v_5 == nil) {
__t3 = b_4
goto end_branch_3
} else {

}
}
{
if (v_5 != nil) {
var __t2 struct{
	acc string
	go__init bool
}
{
if (gopurs_runtime.RecordGet(b_4, "init").IntVal) != (0) {
__t2 = struct{
	acc string
	go__init bool
}{(v_5).V0.StrVal(), false}
goto end_branch_2
} else {

}
}
{
__t2 = struct{
	acc string
	go__init bool
}{((gopurs_runtime.RecordGet(b_4, "acc").StrVal()) + (" : ")) + ((v_5).V0.StrVal()), false}
}
end_branch_2:
b_4_loop = func() gopurs_runtime.Value {
				orig := __t2
				_ = orig
				return gopurs_runtime.RecordDict([]string{"acc", "init"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.acc), gopurs_runtime.Bool(orig.go__init)})
				}()
v_5_loop = (v_5).V1
continue go__go_3_1_21
__t3 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_3_1_21 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_3_1_21(b_4_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val))
})
})
__t5 = (("(") + (gopurs_runtime.RecordGet(Call_local_Data_List_Types_go__go_3_1_21(func() gopurs_runtime.Value {
				orig := struct{
	acc string
	go__init bool
}{"", true}
				_ = orig
				return gopurs_runtime.RecordDict([]string{"acc", "init"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.acc), gopurs_runtime.Bool(orig.go__init)})
				}(), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_List_Types_listMap(show_1_0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2))}))), "acc").StrVal())) + (" : Nil)")
}
end_branch_5:
return gopurs_runtime.Str(__t5)
})})))}
}

func Call_Data_List_Types_showNonEmptyList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): show_1_1 shape=Other bindingType=(Func [(TypeVar a)] String)
show_1_1 := gopurs_runtime.RecordGet(dictShow_0, "show")
_ = show_1_1
// TAST (Let): __local_var_2_2 shape=LitRecord bindingType=(ADT ["Data","Show","Show"] [(ADT ["Data","List","Types","List"] [(TypeVar a)])])
__local_var_2_2 := (&Constructor_Data_Show_Show[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 string
{
var __t_tag_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2)
if (__t_tag_6 == nil) {
__t7 = "Nil"
goto end_branch_7
} else {

}
}
{
var Call_local_Data_List_Types_go__go_3_3_22 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_3_3_22
var go__go_3_3_22 gopurs_runtime.Value
_ = go__go_3_3_22
Call_local_Data_List_Types_go__go_3_3_22 = func(b_4_loop gopurs_runtime.Value, v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_3_3_22:
for {
if false { continue go__go_3_3_22 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var __t5 gopurs_runtime.Value
{
if (v_5 == nil) {
__t5 = b_4
goto end_branch_5
} else {

}
}
{
if (v_5 != nil) {
var __t4 struct{
	acc string
	go__init bool
}
{
if (gopurs_runtime.RecordGet(b_4, "init").IntVal) != (0) {
__t4 = struct{
	acc string
	go__init bool
}{(v_5).V0.StrVal(), false}
goto end_branch_4
} else {

}
}
{
__t4 = struct{
	acc string
	go__init bool
}{((gopurs_runtime.RecordGet(b_4, "acc").StrVal()) + (" : ")) + ((v_5).V0.StrVal()), false}
}
end_branch_4:
b_4_loop = func() gopurs_runtime.Value {
				orig := __t4
				_ = orig
				return gopurs_runtime.RecordDict([]string{"acc", "init"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.acc), gopurs_runtime.Bool(orig.go__init)})
				}()
v_5_loop = (v_5).V1
continue go__go_3_3_22
__t5 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_3_3_22 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_3_3_22(b_4_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val))
})
})
__t7 = (("(") + (gopurs_runtime.RecordGet(Call_local_Data_List_Types_go__go_3_3_22(func() gopurs_runtime.Value {
				orig := struct{
	acc string
	go__init bool
}{"", true}
				_ = orig
				return gopurs_runtime.RecordDict([]string{"acc", "init"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.acc), gopurs_runtime.Bool(orig.go__init)})
				}(), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_List_Types_listMap(show_1_1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2))}))), "acc").StrVal())) + (" : Nil)")
}
end_branch_7:
return gopurs_runtime.Str(__t7)
})})
_ = __local_var_2_2
// TAST (Let): showNonEmpty_1_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Data","Show","Show"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (TypeVar a)])])
showNonEmpty_1_0 := (&Constructor_Data_Show_Show[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((((("(NonEmpty ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_2_2.V0), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1).StrVal())) + (")"))
})})
_ = showNonEmpty_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_4061813263_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(NonEmptyList ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showNonEmpty_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952(Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_2))))}).StrVal())) + (")"))
})})))}
}

func Call_Data_List_Types_eq1(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], xs_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], ys_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) bool {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var xs_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = xs_1_loop
_ = xs_1
var ys_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = ys_2_loop
_ = ys_2
var go__go_3_0_38 gopurs_runtime.Value
_ = go__go_3_0_38
// FALLBACK TCO: isLoop=false len=1
go__go_3_0_38 = gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 bool
{
if ((v2_6.IntVal) != (0)) != (true) {
__t6 = false
goto end_branch_6
} else {

}
}
{
var __t_tag_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4)
if (__t_tag_4 == nil) {
var __t_tag_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5)
__t6 = ((__t_tag_5 == nil)) && ((v2_6.IntVal) != (0))
goto end_branch_6
} else {

}
}
{
var __t_tag_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4)
var __t_and_3 bool = false
if (__t_tag_1 != nil) {

var __t_tag_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5)
__t_and_3 = ((__t_tag_2 != nil)) && ((gopurs_runtime.Apply3(go__go_3_0_38, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}, gopurs_runtime.Bool(((v2_6.IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0).IntVal) != (0)))).IntVal) != (0))
}
__t6 = __t_and_3
}
end_branch_6:
return gopurs_runtime.Bool(__t6)
})
return (gopurs_runtime.Apply3(go__go_3_0_38, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(ys_2)}, gopurs_runtime.Bool(true)).IntVal) != (0)
}

func Call_Data_List_Types_eqNonEmpty(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2176830691_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_and_7 bool = false
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal) != (0) {

var go__go_3_0_39 gopurs_runtime.Value
_ = go__go_3_0_39
// FALLBACK TCO: isLoop=false len=1
go__go_3_0_39 = gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 bool
{
if ((v2_6.IntVal) != (0)) != (true) {
__t6 = false
goto end_branch_6
} else {

}
}
{
var __t_tag_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4)
if (__t_tag_4 == nil) {
var __t_tag_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5)
__t6 = ((__t_tag_5 == nil)) && ((v2_6.IntVal) != (0))
goto end_branch_6
} else {

}
}
{
var __t_tag_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4)
var __t_and_3 bool = false
if (__t_tag_1 != nil) {

var __t_tag_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5)
__t_and_3 = ((__t_tag_2 != nil)) && ((gopurs_runtime.Apply3(go__go_3_0_39, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}, gopurs_runtime.Bool(((v2_6.IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0).IntVal) != (0)))).IntVal) != (0))
}
__t6 = __t_and_3
}
end_branch_6:
return gopurs_runtime.Bool(__t6)
})
__t_and_7 = (gopurs_runtime.Apply3(go__go_3_0_39, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V1))}, gopurs_runtime.Bool(true)).IntVal) != (0)
}
return gopurs_runtime.Bool(__t_and_7)
})})))}
}

func Call_Data_List_Types_eqList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1636902754_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(xs_1 gopurs_runtime.Value, ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_0_41 gopurs_runtime.Value
_ = go__go_3_0_41
// FALLBACK TCO: isLoop=false len=1
go__go_3_0_41 = gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 bool
{
if ((v2_6.IntVal) != (0)) != (true) {
__t6 = false
goto end_branch_6
} else {

}
}
{
var __t_tag_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4)
if (__t_tag_4 == nil) {
var __t_tag_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5)
__t6 = ((__t_tag_5 == nil)) && ((v2_6.IntVal) != (0))
goto end_branch_6
} else {

}
}
{
var __t_tag_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4)
var __t_and_3 bool = false
if (__t_tag_1 != nil) {

var __t_tag_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5)
__t_and_3 = ((__t_tag_2 != nil)) && ((gopurs_runtime.Apply3(go__go_3_0_41, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}, gopurs_runtime.Bool(((v2_6.IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0).IntVal) != (0)))).IntVal) != (0))
}
__t6 = __t_and_3
}
end_branch_6:
return gopurs_runtime.Bool(__t6)
})
return gopurs_runtime.Bool((gopurs_runtime.Apply3(go__go_3_0_41, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_2))}, gopurs_runtime.Bool(true)).IntVal) != (0))
})})))}
}

func Call_Data_List_Types_eqNonEmptyList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3912963311_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_and_7 bool = false
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V0).IntVal) != (0) {

var go__go_3_0_42 gopurs_runtime.Value
_ = go__go_3_0_42
// FALLBACK TCO: isLoop=false len=1
go__go_3_0_42 = gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 bool
{
if ((v2_6.IntVal) != (0)) != (true) {
__t6 = false
goto end_branch_6
} else {

}
}
{
var __t_tag_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4)
if (__t_tag_4 == nil) {
var __t_tag_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5)
__t6 = ((__t_tag_5 == nil)) && ((v2_6.IntVal) != (0))
goto end_branch_6
} else {

}
}
{
var __t_tag_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4)
var __t_and_3 bool = false
if (__t_tag_1 != nil) {

var __t_tag_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5)
__t_and_3 = ((__t_tag_2 != nil)) && ((gopurs_runtime.Apply3(go__go_3_0_42, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}, gopurs_runtime.Bool(((v2_6.IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0).IntVal) != (0)))).IntVal) != (0))
}
__t6 = __t_and_3
}
end_branch_6:
return gopurs_runtime.Bool(__t6)
})
__t_and_7 = (gopurs_runtime.Apply3(go__go_3_0_42, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_2.UnsafePtr).V1))}, gopurs_runtime.Bool(true)).IntVal) != (0)
}
return gopurs_runtime.Bool(__t_and_7)
})})))}
}

func Call_Data_List_Types_compare1(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], xs_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], ys_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) uint32 {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var xs_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = xs_1_loop
_ = xs_1
var ys_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = ys_2_loop
_ = ys_2
var Call_local_Data_List_Types_go__go_3_0_44 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) uint32
_ = Call_local_Data_List_Types_go__go_3_0_44
var go__go_3_0_44 gopurs_runtime.Value
_ = go__go_3_0_44
Call_local_Data_List_Types_go__go_3_0_44 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) uint32 {
go__go_3_0_44:
for {
if false { continue go__go_3_0_44 }
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_5_loop
_ = v1_5
var __t4 uint32
{
if (v_4 == nil) {
var __t1 uint32
{
if (v1_5 == nil) {
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
if (v1_5 == nil) {
__t4 = 380165415
goto end_branch_4
} else {

}
}
{
if ((v_4 != nil)) && ((v1_5 != nil)) {
// TAST (Let): v2_6_2 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v2_6_2 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), (v_4).V0, (v1_5).V0).IntVal)
_ = v2_6_2
var __t3 uint32
{
if (v2_6_2 == 902936544) {
v_4_loop = (v_4).V1
v1_5_loop = (v1_5).V1
continue go__go_3_0_44
__t3 = func() uint32 { panic("unreachable") }()
goto end_branch_3
} else {

}
}
{
__t3 = v2_6_2
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}
go__go_3_0_44 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_List_Types_go__go_3_0_44(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val))), UnsafePtr: nil}
})
})
return Call_local_Data_List_Types_go__go_3_0_44(xs_1, ys_2)
}

func Call_Data_List_Types_ordList(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): eqList1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","List","Types","List"] [(TypeVar a)])])
eqList1_1_0 := (&Constructor_Data_Eq_Eq[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(xs_2 gopurs_runtime.Value, ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_45 gopurs_runtime.Value
_ = go__go_4_2_45
// FALLBACK TCO: isLoop=false len=1
go__go_4_2_45 = gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 bool
{
if ((v2_7.IntVal) != (0)) != (true) {
__t8 = false
goto end_branch_8
} else {

}
}
{
var __t_tag_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5)
if (__t_tag_6 == nil) {
var __t_tag_7 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6)
__t8 = ((__t_tag_7 == nil)) && ((v2_7.IntVal) != (0))
goto end_branch_8
} else {

}
}
{
var __t_tag_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5)
var __t_and_5 bool = false
if (__t_tag_3 != nil) {

var __t_tag_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6)
__t_and_5 = ((__t_tag_4 != nil)) && ((gopurs_runtime.Apply3(go__go_4_2_45, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}, gopurs_runtime.Bool(((v2_7.IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "eq"), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V0).IntVal) != (0)))).IntVal) != (0))
}
__t8 = __t_and_5
}
end_branch_8:
return gopurs_runtime.Bool(__t8)
})
return gopurs_runtime.Bool((gopurs_runtime.Apply3(go__go_4_2_45, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_3))}, gopurs_runtime.Bool(true)).IntVal) != (0))
})})
_ = eqList1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_4210054658_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1636902754_3790796878(eqList1_1_0))}
}), gopurs_runtime.Func2(func(xs_2 gopurs_runtime.Value, ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_4_9_46 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) uint32
_ = Call_local_Data_List_Types_go__go_4_9_46
var go__go_4_9_46 gopurs_runtime.Value
_ = go__go_4_9_46
Call_local_Data_List_Types_go__go_4_9_46 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) uint32 {
go__go_4_9_46:
for {
if false { continue go__go_4_9_46 }
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var __t13 uint32
{
if (v_5 == nil) {
var __t10 uint32
{
if (v1_6 == nil) {
__t10 = 902936544
goto end_branch_10
} else {

}
}
{
__t10 = 1527465420
}
end_branch_10:
__t13 = __t10
goto end_branch_13
} else {

}
}
{
if (v1_6 == nil) {
__t13 = 380165415
goto end_branch_13
} else {

}
}
{
if ((v_5 != nil)) && ((v1_6 != nil)) {
// TAST (Let): v2_7_11 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v2_7_11 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (v_5).V0, (v1_6).V0).IntVal)
_ = v2_7_11
var __t12 uint32
{
if (v2_7_11 == 902936544) {
v_5_loop = (v_5).V1
v1_6_loop = (v1_6).V1
continue go__go_4_9_46
__t12 = func() uint32 { panic("unreachable") }()
goto end_branch_12
} else {

}
}
{
__t12 = v2_7_11
}
end_branch_12:
__t13 = __t12
goto end_branch_13
} else {

}
}
{
__t13 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
}
}
go__go_4_9_46 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_List_Types_go__go_4_9_46(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val))), UnsafePtr: nil}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_List_Types_go__go_4_9_46(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_2), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_3))), UnsafePtr: nil}
})})))}
}

func Call_Data_List_Types_ordNonEmptyList(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1180711887_4177771502(Rebox_Data_List_Types_4177771502_1180711887(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_NonEmpty_ordNonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3739089555_3985601471(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_ord1List())))}, dictOrd_0)))))}
}

func Call_Data_List_Types_pure(x_0_loop gopurs_runtime.Value) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return Rebox_Data_List_Types_1293498952_3123684004((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}}))
}

func Rebox_Data_List_Types_1022383170_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_List_Types_111597075_2412140840(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_1116310629_849153993(in *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(in.V0)}
		out.V1 = Rebox_Data_List_Types_1116310629_849153993(in.V1)
	return out
}

func Rebox_Data_List_Types_1180711887_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_1262583736_2550391993(in *Constructor_Control_Comonad_Comonad[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Control_Comonad_Comonad[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_1293498952_3123684004(in *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_1305434581_3741347833(in *Constructor_Control_Apply_Apply[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Control_Apply_Apply[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Apply_Apply[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_132451362_2738507278(in *Constructor_Data_Unfoldable_Unfoldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_138441832_1728839155(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]{}
		out.V0 = in.V0.IntVal
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_138441832_3132786365(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[uint32, float64] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[uint32, float64]{}
		out.V0 = uint32(in.V0.IntVal)
		out.V1 = in.V1.FloatVal()
	return out
}

func Rebox_Data_List_Types_138441832_3800170591(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
		out.V0 = in.V0.IntVal
		out.V1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](in.V1)
	return out
}

func Rebox_Data_List_Types_138441832_3804580809(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](in.V1)
	return out
}

func Rebox_Data_List_Types_1468200963_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Types_1544744991_1812164904(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_List_Types_1636902754_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Types_1728839155_138441832(in *Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_1757775096_2568689657(in *Constructor_Control_Monad_Monad[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Control_Monad_Monad[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Monad[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_1812164904_4122471187(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_TraversableWithIndex_TraversableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_List_Types_1951493170_3725484264(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_List_Types_203237368_3421983481(in *Constructor_Control_Alt_Alt[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Control_Alt_Alt[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Alt_Alt[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_2171623093_3290176857(in *Constructor_Control_Extend_Extend[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Control_Extend_Extend[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Extend_Extend[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_2173899317_3706288089(in *Constructor_Control_Plus_Plus[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Control_Plus_Plus[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Plus_Plus[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_2176830691_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Types_2183599445_2748095225(in *Constructor_Control_Bind_Bind[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Control_Bind_Bind[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Bind_Bind[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_2334388216_1439734649(in *Constructor_Control_Applicative_Applicative[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_2435023311_2187088110(in *Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Types_2442833393_849153993(in *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3132786365_138441832(in.V0))}
		out.V1 = Rebox_Data_List_Types_2442833393_849153993(in.V1)
	return out
}

func Rebox_Data_List_Types_2476964120_3290176857(in *Constructor_Control_Extend_Extend[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Control_Extend_Extend[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Extend_Extend[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_257347157_3421983481(in *Constructor_Control_Alt_Alt[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Control_Alt_Alt[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Alt_Alt[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_2625657118_1766074591(in *Constructor_Data_Eq_Eq1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Types_2656832021_2394491833(in *Constructor_Control_MonadPlus_MonadPlus[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Control_MonadPlus_MonadPlus[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_MonadPlus_MonadPlus[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_2801299215_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Types_2870828117_1439734649(in *Constructor_Control_Applicative_Applicative[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_3037784642_3043886126(in *Constructor_Data_Traversable_Traversable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_List_Types_3043886126_3425342735(in *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) *Constructor_Data_Traversable_Traversable[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Traversable_Traversable[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_List_Types_3054666744_2748095225(in *Constructor_Control_Bind_Bind[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Control_Bind_Bind[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Bind_Bind[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_3094389156_4010058633(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0)
	return out
}

func Rebox_Data_List_Types_3123684004_1293498952(in *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_3132786365_138441832(in *Constructor_Data_Tuple_Tuple[uint32, float64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
		out.V1 = gopurs_runtime.Float(in.V1)
	return out
}

func Rebox_Data_List_Types_3262795586_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Types_3290489278_3985601471(in *Constructor_Data_Ord_Ord1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_3339399026_2412140840(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_3351995458_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Types_3368202604_4151366573(in *Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_List_Types_3425342735_3043886126(in *Constructor_Data_Traversable_Traversable[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_List_Types_3554500787_1766074591(in *Constructor_Data_Eq_Eq1[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Types_3704040722_849153993(in *Constructor_Data_List_Types_Cons[int64]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
		out.V1 = Rebox_Data_List_Types_3704040722_849153993(in.V1)
	return out
}

func Rebox_Data_List_Types_3725484264_4016503379(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_List_Types_3739089555_3985601471(in *Constructor_Data_Ord_Ord1[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_3742495784_385277032(in *Constructor_Data_Newtype_Newtype[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Types_3800170591_138441832(in *Constructor_Data_Tuple_Tuple[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}

func Rebox_Data_List_Types_3842311788_306175789(in *Constructor_Data_Semigroup_Traversable_Traversable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_List_Types_3912963311_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Types_3970790648_3741347833(in *Constructor_Control_Apply_Apply[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Control_Apply_Apply[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Apply_Apply[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_3985601471_3290489278(in *Constructor_Data_Ord_Ord1[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_3994390613_2568689657(in *Constructor_Control_Monad_Monad[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Control_Monad_Monad[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Monad[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_4016503379_3725484264(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_List_Types_4022093634_4179793454(in *Constructor_Data_Semigroup_Semigroup[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Types_4061813263_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Types_4064382095_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_List_Types_4073635714_2187088110(in *Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Types_4122471187_1812164904(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_List_Types_4151366573_3368202604(in *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_List_Types_4177771502_1180711887(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_4179793454_4022093634(in *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Semigroup[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Types_4210054658_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_538456415_3725484264(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_List_Types_640718306_1201789390(in *Constructor_Data_Monoid_Monoid[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}

func Rebox_Data_List_Types_655843087_4179793454(in *Constructor_Data_Semigroup_Semigroup[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Types_677368690_1812164904(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_List_Types_721753375_2412140840(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_739440277_2307501113(in *Constructor_Control_Alternative_Alternative[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_849153993_2442833393(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]{}
		out.V0 = Rebox_Data_List_Types_138441832_3132786365(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
		out.V1 = Rebox_Data_List_Types_849153993_2442833393(in.V1)
	return out
}

func Rebox_Data_List_Types_849153993_3704040722(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[int64]{}
		out.V0 = in.V0.IntVal
		out.V1 = Rebox_Data_List_Types_849153993_3704040722(in.V1)
	return out
}



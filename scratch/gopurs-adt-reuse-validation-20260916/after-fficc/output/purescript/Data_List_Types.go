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
		cache_Data_List_Types_identity = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Data_List_Types_identity
}

var cache_Data_List_Types_identity1 gopurs_runtime.Value
var once_Data_List_Types_identity1 sync.Once
func Get_Data_List_Types_identity1() gopurs_runtime.Value {
	once_Data_List_Types_identity1.Do(func() {
		cache_Data_List_Types_identity1 = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
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

var cache_Data_List_Types_Cons__1027980219 gopurs_runtime.Value
var once_Data_List_Types_Cons__1027980219 sync.Once
func Get_Data_List_Types_Cons__1027980219() gopurs_runtime.Value {
	once_Data_List_Types_Cons__1027980219.Do(func() {
		cache_Data_List_Types_Cons__1027980219 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2442833393_849153993(Call_Data_List_Types_Cons__1027980219(Rebox_Data_List_Types_138441832_3132786365(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](__eta_norm_1_0_box)), Rebox_Data_List_Types_849153993_2442833393(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__eta_norm_0_1_box)))))}
})
	})
	return cache_Data_List_Types_Cons__1027980219
}

var cache_Data_List_Types_Cons__351658551 gopurs_runtime.Value
var once_Data_List_Types_Cons__351658551 sync.Once
func Get_Data_List_Types_Cons__351658551() gopurs_runtime.Value {
	once_Data_List_Types_Cons__351658551.Do(func() {
		cache_Data_List_Types_Cons__351658551 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3704040722_849153993(Call_Data_List_Types_Cons__351658551(__eta_norm_1_0_box.IntVal, Rebox_Data_List_Types_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__eta_norm_0_unused_1_box)))))}
})
	})
	return cache_Data_List_Types_Cons__351658551
}

var cache_Data_List_Types_Cons__317927332 gopurs_runtime.Value
var once_Data_List_Types_Cons__317927332 sync.Once
func Get_Data_List_Types_Cons__317927332() gopurs_runtime.Value {
	once_Data_List_Types_Cons__317927332.Do(func() {
		cache_Data_List_Types_Cons__317927332 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3704040722_849153993(Call_Data_List_Types_Cons__317927332(__eta_norm_1_0_box.IntVal, Rebox_Data_List_Types_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__eta_norm_0_1_box)))))}
})
	})
	return cache_Data_List_Types_Cons__317927332
}

var cache_Data_List_Types_NonEmptyList gopurs_runtime.Value
var once_Data_List_Types_NonEmptyList sync.Once
func Get_Data_List_Types_NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_NonEmptyList.Do(func() {
		cache_Data_List_Types_NonEmptyList = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952(Call_Data_List_Types_NonEmptyList(Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](x_0_box)))))}
})
	})
	return cache_Data_List_Types_NonEmptyList
}

var cache_Data_List_Types_NonEmptyList__2988773242 gopurs_runtime.Value
var once_Data_List_Types_NonEmptyList__2988773242 sync.Once
func Get_Data_List_Types_NonEmptyList__2988773242() gopurs_runtime.Value {
	once_Data_List_Types_NonEmptyList__2988773242.Do(func() {
		cache_Data_List_Types_NonEmptyList__2988773242 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952(Call_Data_List_Types_NonEmptyList__2988773242(Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](x_0_box)))))}
})
	})
	return cache_Data_List_Types_NonEmptyList__2988773242
}

var cache_Data_List_Types_toList gopurs_runtime.Value
var once_Data_List_Types_toList sync.Once
func Get_Data_List_Types_toList() gopurs_runtime.Value {
	once_Data_List_Types_toList.Do(func() {
		cache_Data_List_Types_toList = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_Types_toList(Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952(Call_Data_List_Types_nelCons(a_0_box, Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box)))))}
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
		cache_Data_List_Types_functorNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2801299215_2812149806(Rebox_Data_List_Types_2812149806_2801299215(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Data_NonEmpty_functorNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3262795586_2812149806(Rebox_Data_List_Types_2812149806_3262795586(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Types_functorList()))))})))))}
	})
	return cache_Data_List_Types_functorNonEmptyList
}

var cache_Data_List_Types_foldableList gopurs_runtime.Value
var once_Data_List_Types_foldableList sync.Once
func Get_Data_List_Types_foldableList() gopurs_runtime.Value {
	once_Data_List_Types_foldableList.Do(func() {
		cache_Data_List_Types_foldableList = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1022383170_1680800814((&Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m$scope221)])
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 shape=App(Var) bindingType=(TypeVar m$scope221)
mempty_2_1 := Call_Data_Monoid_mempty(dictMonoid_0)
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Rebox_Data_List_Types_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())).V1, gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Semigroup0_1_0.V0, acc_4), f_3)
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_1_2_2 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_1_2_2
var go__go_1_2_2 gopurs_runtime.Value
_ = go__go_1_2_2
Call_local_Data_List_Types_go__go_1_2_2 = func(b_2_loop gopurs_runtime.Value, v_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_1_2_2:
for {
if false { continue go__go_1_2_2 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var v_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_3_loop
_ = v_3
var __t3 gopurs_runtime.Value
{
if (v_3 == nil) {
__t3 = b_2
goto end_branch_3
} else {

}
}
{
if (v_3 != nil) {
b_2_loop = gopurs_runtime.Apply2(f_0, b_2, (v_3).V0)
v_3_loop = (v_3).V1
continue go__go_1_2_2
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
go__go_1_2_2 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_1_2_2(b_2_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_3_loop_val))
})
})
return go__go_1_2_2
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_2_4_3 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_2_4_3
var go__go_2_4_3 gopurs_runtime.Value
_ = go__go_2_4_3
Call_local_Data_List_Types_go__go_2_4_3 = func(v_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_2_4_3:
for {
if false { continue go__go_2_4_3 }
var v_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_3_loop
_ = v_3
var v1_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_4_loop
_ = v1_4
var __t5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_4 == nil) {
__t5 = v_3
goto end_branch_5
} else {

}
}
{
if (v1_4 != nil) {
v_3_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_4).V0, v_3})
v1_4_loop = (v1_4).V1
continue go__go_2_4_3
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
go__go_2_4_3 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_2_4_3(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_3_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_4_loop_val)))}
})
})
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply2(Rebox_Data_List_Types_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())).V1, gopurs_runtime.Func2(func(b_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
}), b_1), gopurs_runtime.Apply(go__go_2_4_3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}))
})})))}
	})
	return cache_Data_List_Types_foldableList
}

var cache_Data_List_Types_foldableNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_foldableNonEmptyList sync.Once
func Get_Data_List_Types_foldableNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_foldableNonEmptyList.Do(func() {
		cache_Data_List_Types_foldableNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_4064382095_1680800814(Rebox_Data_List_Types_1680800814_4064382095(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Data_NonEmpty_foldableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1022383170_1680800814(Rebox_Data_List_Types_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList()))))})))))}
	})
	return cache_Data_List_Types_foldableNonEmptyList
}

var cache_Data_List_Types_foldableWithIndexList gopurs_runtime.Value
var once_Data_List_Types_foldableWithIndexList sync.Once
func Get_Data_List_Types_foldableWithIndexList() gopurs_runtime.Value {
	once_Data_List_Types_foldableWithIndexList.Do(func() {
		cache_Data_List_Types_foldableWithIndexList = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_538456415_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1022383170_1680800814(Rebox_Data_List_Types_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList()))))}
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m$scope204)])
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 shape=App(Var) bindingType=(TypeVar m$scope204)
mempty_2_1 := Call_Data_Monoid_mempty(dictMonoid_0)
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Rebox_Data_List_Types_3725484264_538456415(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Types_foldableWithIndexList())).V2, gopurs_runtime.Func2(func(i_4 gopurs_runtime.Value, acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Semigroup0_1_0.V0, acc_5), gopurs_runtime.Apply(f_3, gopurs_runtime.Int(i_4.IntVal)))
}), mempty_2_1)
})
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, acc_1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_2_2_4 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_2_2_4
var go__go_2_2_4 gopurs_runtime.Value
_ = go__go_2_2_4
Call_local_Data_List_Types_go__go_2_2_4 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_2_4:
for {
if false { continue go__go_2_2_4 }
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
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1728839155_138441832(Rebox_Data_List_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_3.UnsafePtr).V0.IntVal) + (int64(1))), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_3.UnsafePtr).V0.IntVal), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_3.UnsafePtr).V1, (v_4).V0)}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}
v_4_loop = (v_4).V1
continue go__go_2_2_4
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
go__go_2_2_4 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_2_2_4(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Tuple_snd(), gopurs_runtime.Apply(go__go_2_2_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1728839155_138441832(Rebox_Data_List_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(int64(0)), acc_1}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}))
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_3_5_5 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_3_5_5
var go__go_3_5_5 gopurs_runtime.Value
_ = go__go_3_5_5
Call_local_Data_List_Types_go__go_3_5_5 = func(b_4_loop gopurs_runtime.Value, v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_3_5_5:
for {
if false { continue go__go_3_5_5 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var __t6 gopurs_runtime.Value
{
if (v_5 == nil) {
__t6 = b_4
goto end_branch_6
} else {

}
}
{
if (v_5 != nil) {
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3800170591_138441832(Rebox_Data_List_Types_138441832_3800170591(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V0.IntVal) + (int64(1))), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_5).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V1)}))}}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}
v_5_loop = (v_5).V1
continue go__go_3_5_5
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
go__go_3_5_5 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_3_5_5(b_4_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val))
})
})
// TAST (Let): v_3_4 shape=LetRec(App(Other)) bindingType=(ADT ["Data","Tuple","Tuple"] [Int, (ADT ["Data","List","Types","List"] [(TypeVar a$scope184)])])
v_3_4 := Rebox_Data_List_Types_138441832_3800170591(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](Call_local_Data_List_Types_go__go_3_5_5(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3800170591_138441832(Rebox_Data_List_Types_138441832_3800170591(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(int64(0)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_2))))
_ = v_3_4
var Call_local_Data_List_Types_go__go_4_7_6 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_4_7_6
var go__go_4_7_6 gopurs_runtime.Value
_ = go__go_4_7_6
Call_local_Data_List_Types_go__go_4_7_6 = func(b_5_loop gopurs_runtime.Value, v_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_4_7_6:
for {
if false { continue go__go_4_7_6 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_6_loop
_ = v_6
var __t8 gopurs_runtime.Value
{
if (v_6 == nil) {
__t8 = b_5
goto end_branch_8
} else {

}
}
{
if (v_6 != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1728839155_138441832(Rebox_Data_List_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V0.IntVal) - (int64(1))), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V0.IntVal) - (int64(1))), (v_6).V0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V1)}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}
v_6_loop = (v_6).V1
continue go__go_4_7_6
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
go__go_4_7_6 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_4_7_6(b_5_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_6_loop_val))
})
})
return Call_Data_Tuple_snd(Rebox_Data_List_Types_1728839155_138441832(Rebox_Data_List_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](Call_local_Data_List_Types_go__go_4_7_6(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1728839155_138441832(Rebox_Data_List_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int((v_3_4).V0), b_1}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}, (v_3_4).V1)))))
})})))}
	})
	return cache_Data_List_Types_foldableWithIndexList
}

var cache_Data_List_Types_foldableWithIndexNonEmpty gopurs_runtime.Value
var once_Data_List_Types_foldableWithIndexNonEmpty sync.Once
func Get_Data_List_Types_foldableWithIndexNonEmpty() gopurs_runtime.Value {
	once_Data_List_Types_foldableWithIndexNonEmpty.Do(func() {
		cache_Data_List_Types_foldableWithIndexNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_4016503379_3725484264(Rebox_Data_List_Types_3725484264_4016503379(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Call_Data_NonEmpty_foldableWithIndexNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_538456415_3725484264(Rebox_Data_List_Types_3725484264_538456415(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Types_foldableWithIndexList()))))})))))}
	})
	return cache_Data_List_Types_foldableWithIndexNonEmpty
}

var cache_Data_List_Types_foldableWithIndexNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_foldableWithIndexNonEmptyList sync.Once
func Get_Data_List_Types_foldableWithIndexNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_foldableWithIndexNonEmptyList.Do(func() {
		cache_Data_List_Types_foldableWithIndexNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1951493170_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_4064382095_1680800814(Rebox_Data_List_Types_1680800814_4064382095(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Data_NonEmpty_foldableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1022383170_1680800814(Rebox_Data_List_Types_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList()))))})))))}
}), gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldableWithIndexNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_538456415_3725484264(Rebox_Data_List_Types_3725484264_538456415(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Types_foldableWithIndexList()))))}), "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))}, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_1, gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
var __t_tag_0 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_List_Types_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_3))
_ = __t_tag_0
if (__t_tag_0 == nil) {
__t2 = gopurs_runtime.Int(int64(0))
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_List_Types_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_3))
_ = __t_tag_1
if (__t_tag_1 != nil) {
__t2 = gopurs_runtime.Int((int64(1)) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_3.UnsafePtr).V0.IntVal))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
})), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952(Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_2))))})
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldableWithIndexNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_538456415_3725484264(Rebox_Data_List_Types_3725484264_538456415(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Types_foldableWithIndexList()))))}), "foldlWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_0, gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
var __t_tag_3 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_List_Types_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_3))
_ = __t_tag_3
if (__t_tag_3 == nil) {
__t5 = gopurs_runtime.Int(int64(0))
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_List_Types_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_3))
_ = __t_tag_4
if (__t_tag_4 != nil) {
__t5 = gopurs_runtime.Int((int64(1)) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_3.UnsafePtr).V0.IntVal))
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
})), b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952(Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_2))))})
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Data_NonEmpty_foldableWithIndexNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_538456415_3725484264(Rebox_Data_List_Types_3725484264_538456415(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Types_foldableWithIndexList()))))}), "foldrWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_0, gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
var __t_tag_6 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_List_Types_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_3))
_ = __t_tag_6
if (__t_tag_6 == nil) {
__t8 = gopurs_runtime.Int(int64(0))
goto end_branch_8
} else {

}
}
{
var __t_tag_7 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_List_Types_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_3))
_ = __t_tag_7
if (__t_tag_7 != nil) {
__t8 = gopurs_runtime.Int((int64(1)) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_3.UnsafePtr).V0.IntVal))
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
})), b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952(Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_2))))})
})})))}
	})
	return cache_Data_List_Types_foldableWithIndexNonEmptyList
}

var cache_Data_List_Types_functorWithIndexList gopurs_runtime.Value
var once_Data_List_Types_functorWithIndexList sync.Once
func Get_Data_List_Types_functorWithIndexList() gopurs_runtime.Value {
	once_Data_List_Types_functorWithIndexList.Do(func() {
		cache_Data_List_Types_functorWithIndexList = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_721753375_2412140840((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3262795586_2812149806(Rebox_Data_List_Types_2812149806_3262795586(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Types_functorList()))))}
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_2_1_7 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_2_1_7
var go__go_2_1_7 gopurs_runtime.Value
_ = go__go_2_1_7
Call_local_Data_List_Types_go__go_2_1_7 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_1_7:
for {
if false { continue go__go_2_1_7 }
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
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3800170591_138441832(Rebox_Data_List_Types_138441832_3800170591(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_3.UnsafePtr).V0.IntVal) + (int64(1))), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_4).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_3.UnsafePtr).V1)}))}}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}
v_4_loop = (v_4).V1
continue go__go_2_1_7
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
go__go_2_1_7 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_2_1_7(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
// TAST (Let): v_2_0 shape=LetRec(App(Other)) bindingType=(ADT ["Data","Tuple","Tuple"] [Int, (ADT ["Data","List","Types","List"] [(TypeVar a$scope184)])])
v_2_0 := Rebox_Data_List_Types_138441832_3800170591(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](Call_local_Data_List_Types_go__go_2_1_7(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3800170591_138441832(Rebox_Data_List_Types_138441832_3800170591(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(int64(0)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1))))
_ = v_2_0
var Call_local_Data_List_Types_go__go_3_3_8 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_3_3_8
var go__go_3_3_8 gopurs_runtime.Value
_ = go__go_3_3_8
Call_local_Data_List_Types_go__go_3_3_8 = func(b_4_loop gopurs_runtime.Value, v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_3_3_8:
for {
if false { continue go__go_3_3_8 }
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
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1728839155_138441832(Rebox_Data_List_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V0.IntVal) - (int64(1))), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V0.IntVal) - (int64(1))), (v_5).V0), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_4.UnsafePtr).V1)}))}}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}
v_5_loop = (v_5).V1
continue go__go_3_3_8
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
go__go_3_3_8 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_3_3_8(b_4_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val))
})
})
return Call_Data_Tuple_snd(Rebox_Data_List_Types_1728839155_138441832(Rebox_Data_List_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](Call_local_Data_List_Types_go__go_3_3_8(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1728839155_138441832(Rebox_Data_List_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int((v_2_0).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}, (v_2_0).V1)))))
})})))}
	})
	return cache_Data_List_Types_functorWithIndexList
}

var cache_Data_List_Types_functorWithIndex gopurs_runtime.Value
var once_Data_List_Types_functorWithIndex sync.Once
func Get_Data_List_Types_functorWithIndex() gopurs_runtime.Value {
	once_Data_List_Types_functorWithIndex.Do(func() {
		cache_Data_List_Types_functorWithIndex = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_111597075_2412140840(Rebox_Data_List_Types_2412140840_111597075(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Call_Data_NonEmpty_functorWithIndex(gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_721753375_2412140840(Rebox_Data_List_Types_2412140840_721753375(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Types_functorWithIndexList()))))})))))}
	})
	return cache_Data_List_Types_functorWithIndex
}

var cache_Data_List_Types_functorWithIndexNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_functorWithIndexNonEmptyList sync.Once
func Get_Data_List_Types_functorWithIndexNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_functorWithIndexNonEmptyList.Do(func() {
		cache_Data_List_Types_functorWithIndexNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3339399026_2412140840((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2801299215_2812149806(Rebox_Data_List_Types_2812149806_2801299215(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Data_NonEmpty_functorNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3262795586_2812149806(Rebox_Data_List_Types_2812149806_3262795586(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Types_functorList()))))})))))}
}), gopurs_runtime.Func2(func(fn_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952(Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Call_Data_NonEmpty_functorWithIndex(gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_721753375_2412140840(Rebox_Data_List_Types_2412140840_721753375(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Types_functorWithIndexList()))))}), "mapWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), fn_0, gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
var __t_tag_0 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_List_Types_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_2))
_ = __t_tag_0
if (__t_tag_0 == nil) {
__t2 = gopurs_runtime.Int(int64(0))
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_List_Types_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_2))
_ = __t_tag_1
if (__t_tag_1 != nil) {
__t2 = gopurs_runtime.Int((int64(1)) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_2.UnsafePtr).V0.IntVal))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
})), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952(Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_1))))})))))}
})})))}
	})
	return cache_Data_List_Types_functorWithIndexNonEmptyList
}

var cache_Data_List_Types_semigroupList gopurs_runtime.Value
var once_Data_List_Types_semigroupList sync.Once
func Get_Data_List_Types_semigroupList() gopurs_runtime.Value {
	once_Data_List_Types_semigroupList.Do(func() {
		cache_Data_List_Types_semigroupList = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_4022093634_4179793454((&Constructor_Data_Semigroup_Semigroup[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(xs_0 gopurs_runtime.Value, ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_2_0_9 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_2_0_9
var go__go_2_0_9 gopurs_runtime.Value
_ = go__go_2_0_9
Call_local_Data_List_Types_go__go_2_0_9 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_0_9:
for {
if false { continue go__go_2_0_9 }
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
continue go__go_2_0_9
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
go__go_2_0_9 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_2_0_9(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
var Call_local_Data_List_Types_go__go_2_2_10 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_2_2_10
var go__go_2_2_10 gopurs_runtime.Value
_ = go__go_2_2_10
Call_local_Data_List_Types_go__go_2_2_10 = func(v_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_2_2_10:
for {
if false { continue go__go_2_2_10 }
var v_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_3_loop
_ = v_3
var v1_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_4_loop
_ = v1_4
var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_4 == nil) {
__t3 = v_3
goto end_branch_3
} else {

}
}
{
if (v1_4 != nil) {
v_3_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_4).V0, v_3})
v1_4_loop = (v1_4).V1
continue go__go_2_2_10
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
go__go_2_2_10 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_2_2_10(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_3_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_4_loop_val)))}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(go__go_2_0_9, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_1))}), gopurs_runtime.Apply(go__go_2_2_10, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_0))})))}
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
var Call_local_Data_List_Types_go__go_2_0_11 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_2_0_11
var go__go_2_0_11 gopurs_runtime.Value
_ = go__go_2_0_11
Call_local_Data_List_Types_go__go_2_0_11 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_0_11:
for {
if false { continue go__go_2_0_11 }
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
continue go__go_2_0_11
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
go__go_2_0_11 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_2_0_11(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
var Call_local_Data_List_Types_go__go_2_2_12 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_2_2_12
var go__go_2_2_12 gopurs_runtime.Value
_ = go__go_2_2_12
Call_local_Data_List_Types_go__go_2_2_12 = func(v_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_2_2_12:
for {
if false { continue go__go_2_2_12 }
var v_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_3_loop
_ = v_3
var v1_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_4_loop
_ = v1_4
var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_4 == nil) {
__t3 = v_3
goto end_branch_3
} else {

}
}
{
if (v1_4 != nil) {
v_3_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_4).V0, v_3})
v1_4_loop = (v1_4).V1
continue go__go_2_2_12
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
go__go_2_2_12 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_2_2_12(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_3_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_4_loop_val)))}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952((&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(go__go_2_0_11, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(as_prime__1.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(as_prime__1.UnsafePtr).V1)}))}), gopurs_runtime.Apply(go__go_2_2_12, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1))})))}})))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1022383170_1680800814(Rebox_Data_List_Types_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList()))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3262795586_2812149806(Rebox_Data_List_Types_2812149806_3262795586(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Types_functorList()))))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Rebox_Data_List_Types_3043886126_3037784642(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Types_traversableList())).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope51)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): Apply0_2_1 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m$scope51)])
Apply0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_4_2_13 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_4_2_13
var go__go_4_2_13 gopurs_runtime.Value
_ = go__go_4_2_13
Call_local_Data_List_Types_go__go_4_2_13 = func(b_5_loop gopurs_runtime.Value, v_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_4_2_13:
for {
if false { continue go__go_4_2_13 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_6_loop
_ = v_6
var __t3 gopurs_runtime.Value
{
if (v_6 == nil) {
__t3 = b_5
goto end_branch_3
} else {

}
}
{
if (v_6 != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_6).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_5)}))}
v_6_loop = (v_6).V1
continue go__go_4_2_13
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
go__go_4_2_13 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_4_2_13(b_5_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_6_loop_val))
})
})
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
// TAST (Let): Functor0_7_5 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f$scope44)])
Functor0_7_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(Apply0_2_1.V0, gopurs_runtime.Value{}))
_ = Functor0_7_5
b_5_loop = gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_2_1.V1, gopurs_runtime.Apply2(Functor0_7_5.V0, gopurs_runtime.Func2(func(b_9 gopurs_runtime.Value, a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_10, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_9)}))}
}), b_5), b_8)
}), f_3, (v_6).V0)
v_6_loop = (v_6).V1
continue go__go_4_4_14
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
go__go_4_4_14 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_4_4_14(b_5_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_6_loop_val))
})
})
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Functor0_1_0.V0, gopurs_runtime.Apply(go__go_4_2_13, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})), gopurs_runtime.Apply(go__go_4_4_14, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}))})))
})
})})))}
	})
	return cache_Data_List_Types_traversableList
}

var cache_Data_List_Types_traversableNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_traversableNonEmptyList sync.Once
func Get_Data_List_Types_traversableNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_traversableNonEmptyList.Do(func() {
		cache_Data_List_Types_traversableNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3425342735_3043886126(Rebox_Data_List_Types_3043886126_3425342735(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Data_NonEmpty_traversableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3037784642_3043886126(Rebox_Data_List_Types_3043886126_3037784642(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Types_traversableList()))))})))))}
	})
	return cache_Data_List_Types_traversableNonEmptyList
}

var cache_Data_List_Types_traversableWithIndexList gopurs_runtime.Value
var once_Data_List_Types_traversableWithIndexList sync.Once
func Get_Data_List_Types_traversableWithIndexList() gopurs_runtime.Value {
	once_Data_List_Types_traversableWithIndexList.Do(func() {
		cache_Data_List_Types_traversableWithIndexList = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1544744991_1812164904((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_538456415_3725484264(Rebox_Data_List_Types_3725484264_538456415(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Types_foldableWithIndexList()))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_721753375_2412140840(Rebox_Data_List_Types_2412140840_721753375(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Types_functorWithIndexList()))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3037784642_3043886126(Rebox_Data_List_Types_3043886126_3037784642(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Types_traversableList()))))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope36)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): Apply0_2_1 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m$scope36)])
Apply0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_4_2_15 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_4_2_15
var go__go_4_2_15 gopurs_runtime.Value
_ = go__go_4_2_15
Call_local_Data_List_Types_go__go_4_2_15 = func(b_5_loop gopurs_runtime.Value, v_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_4_2_15:
for {
if false { continue go__go_4_2_15 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_6_loop
_ = v_6
var __t3 gopurs_runtime.Value
{
if (v_6 == nil) {
__t3 = b_5
goto end_branch_3
} else {

}
}
{
if (v_6 != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_6).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_5)}))}
v_6_loop = (v_6).V1
continue go__go_4_2_15
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
go__go_4_2_15 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_4_2_15(b_5_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_6_loop_val))
})
})
var Call_local_Data_List_Types_go__go_4_4_16 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_4_4_16
var go__go_4_4_16 gopurs_runtime.Value
_ = go__go_4_4_16
Call_local_Data_List_Types_go__go_4_4_16 = func(b_5_loop gopurs_runtime.Value, v_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_4_4_16:
for {
if false { continue go__go_4_4_16 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_6_loop
_ = v_6
var __t7 gopurs_runtime.Value
{
if (v_6 == nil) {
__t7 = b_5
goto end_branch_7
} else {

}
}
{
if (v_6 != nil) {
// TAST (Let): Functor0_7_5 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f$scope44)])
Functor0_7_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(Apply0_2_1.V0, gopurs_runtime.Value{}))
_ = Functor0_7_5
// TAST (Let): __local_var_8_6 shape=Other bindingType=(TypeApp (TypeVar m$scope36) [(ADT ["Data","List","Types","List"] [(TypeVar b$scope35)])])
__local_var_8_6 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V1
_ = __local_var_8_6
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1728839155_138441832(Rebox_Data_List_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V0.IntVal) + (int64(1))), gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(b_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_2_1.V1, gopurs_runtime.Apply2(Functor0_7_5.V0, gopurs_runtime.Func2(func(b_10 gopurs_runtime.Value, a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_11, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_10)}))}
}), __local_var_8_6), b_9)
}), gopurs_runtime.Apply(f_3, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V0.IntVal)), (v_6).V0)}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}
v_6_loop = (v_6).V1
continue go__go_4_4_16
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
go__go_4_4_16 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_4_4_16(b_5_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_6_loop_val))
})
})
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Functor0_1_0.V0, gopurs_runtime.Apply(go__go_4_2_15, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Tuple_snd(), gopurs_runtime.Apply(go__go_4_4_16, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1728839155_138441832(Rebox_Data_List_Types_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(int64(0)), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}))})}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))})))
})
})})))}
	})
	return cache_Data_List_Types_traversableWithIndexList
}

var cache_Data_List_Types_traversableWithIndexNonEmpty gopurs_runtime.Value
var once_Data_List_Types_traversableWithIndexNonEmpty sync.Once
func Get_Data_List_Types_traversableWithIndexNonEmpty() gopurs_runtime.Value {
	once_Data_List_Types_traversableWithIndexNonEmpty.Do(func() {
		cache_Data_List_Types_traversableWithIndexNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_4122471187_1812164904(Rebox_Data_List_Types_1812164904_4122471187(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Call_Data_NonEmpty_traversableWithIndexNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1544744991_1812164904(Rebox_Data_List_Types_1812164904_1544744991(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Types_traversableWithIndexList()))))})))))}
	})
	return cache_Data_List_Types_traversableWithIndexNonEmpty
}

var cache_Data_List_Types_traversableWithIndexNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_traversableWithIndexNonEmptyList sync.Once
func Get_Data_List_Types_traversableWithIndexNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_traversableWithIndexNonEmptyList.Do(func() {
		cache_Data_List_Types_traversableWithIndexNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_677368690_1812164904((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1951493170_3725484264(Rebox_Data_List_Types_3725484264_1951493170(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Types_foldableWithIndexNonEmptyList()))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3339399026_2412140840(Rebox_Data_List_Types_2412140840_3339399026(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Types_functorWithIndexNonEmptyList()))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3425342735_3043886126(Rebox_Data_List_Types_3043886126_3425342735(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Data_NonEmpty_traversableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3037784642_3043886126(Rebox_Data_List_Types_3043886126_3037784642(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Types_traversableList()))))})))))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope26)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, Get_Data_List_Types_NonEmptyList(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Call_Data_NonEmpty_traversableWithIndexNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1544744991_1812164904(Rebox_Data_List_Types_1812164904_1544744991(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_List_Types_traversableWithIndexList()))))}), "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_2, gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t_tag_1 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_List_Types_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_4))
_ = __t_tag_1
if (__t_tag_1 == nil) {
__t3 = gopurs_runtime.Int(int64(0))
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Maybe_Just[int64] = Rebox_Data_List_Types_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_4))
_ = __t_tag_2
if (__t_tag_2 != nil) {
__t3 = gopurs_runtime.Int((int64(1)) + ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_4.UnsafePtr).V0.IntVal))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
})), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952(Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_3))))}))
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
var Call_local_Data_List_Types_go__go_2_0_17 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_2_0_17
var go__go_2_0_17 gopurs_runtime.Value
_ = go__go_2_0_17
Call_local_Data_List_Types_go__go_2_0_17 = func(source_3_loop gopurs_runtime.Value, memo_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_2_0_17:
for {
if false { continue go__go_2_0_17 }
var source_3 gopurs_runtime.Value = source_3_loop
_ = source_3
var memo_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = memo_4_loop
_ = memo_4
// TAST (Let): v_5_1 shape=App(Other) bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope12), (ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope13)])])
v_5_1 := Rebox_Data_List_Types_138441832_3804580809(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, source_3)))
_ = v_5_1
var __t6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = (v_5_1).V1
_ = __t_tag_2
if (__t_tag_2 != nil) {
source_3_loop = ((v_5_1).V1).V0
memo_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_5_1).V0, memo_4})
continue go__go_2_0_17
__t6 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_6
} else {

}
}
{
var __t_tag_3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = (v_5_1).V1
_ = __t_tag_3
if (__t_tag_3 == nil) {
var Call_local_Data_List_Types_go__go_6_4_18 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_6_4_18
var go__go_6_4_18 gopurs_runtime.Value
_ = go__go_6_4_18
Call_local_Data_List_Types_go__go_6_4_18 = func(b_7_loop gopurs_runtime.Value, v_8_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_6_4_18:
for {
if false { continue go__go_6_4_18 }
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
continue go__go_6_4_18
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
go__go_6_4_18 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_6_4_18(b_7_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_8_loop_val))
})
})
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_local_Data_List_Types_go__go_6_4_18(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_5_1).V0, memo_4})))
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
go__go_2_0_17 = gopurs_runtime.Func(func(source_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_2_0_17(source_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](memo_4_loop_val)))}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_2_0_17(b_1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)))}
})})))}
	})
	return cache_Data_List_Types_unfoldable1List
}

var cache_Data_List_Types_unfoldableList gopurs_runtime.Value
var once_Data_List_Types_unfoldableList sync.Once
func Get_Data_List_Types_unfoldableList() gopurs_runtime.Value {
	once_Data_List_Types_unfoldableList.Do(func() {
		cache_Data_List_Types_unfoldableList = gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_132451362_2738507278((&Constructor_Data_Unfoldable_Unfoldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_4073635714_2187088110(Rebox_Data_List_Types_2187088110_4073635714(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Get_Data_List_Types_unfoldable1List()))))}
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_2_0_19 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_2_0_19
var go__go_2_0_19 gopurs_runtime.Value
_ = go__go_2_0_19
Call_local_Data_List_Types_go__go_2_0_19 = func(source_3_loop gopurs_runtime.Value, memo_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_2_0_19:
for {
if false { continue go__go_2_0_19 }
var source_3 gopurs_runtime.Value = source_3_loop
_ = source_3
var memo_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = memo_4_loop
_ = memo_4
// TAST (Let): v_5_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope3), (TypeVar b$scope4)])])
v_5_1 := Rebox_Data_List_Types_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, source_3)))
_ = v_5_1
var __t4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v_5_1 == nil) {
var Call_local_Data_List_Types_go__go_6_2_20 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_6_2_20
var go__go_6_2_20 gopurs_runtime.Value
_ = go__go_6_2_20
Call_local_Data_List_Types_go__go_6_2_20 = func(b_7_loop gopurs_runtime.Value, v_8_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_6_2_20:
for {
if false { continue go__go_6_2_20 }
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
continue go__go_6_2_20
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
go__go_6_2_20 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_6_2_20(b_7_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_8_loop_val))
})
})
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_local_Data_List_Types_go__go_6_2_20(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, memo_4))
goto end_branch_4
} else {

}
}
{
if (v_5_1 != nil) {
source_3_loop = ((v_5_1).V0).V1
memo_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, ((v_5_1).V0).V0, memo_4})
continue go__go_2_0_19
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
go__go_2_0_19 = gopurs_runtime.Func(func(source_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_2_0_19(source_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](memo_4_loop_val)))}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_2_0_19(b_1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)))}
})})))}
	})
	return cache_Data_List_Types_unfoldableList
}

var cache_Data_List_Types_unfoldable1NonEmptyList gopurs_runtime.Value
var once_Data_List_Types_unfoldable1NonEmptyList sync.Once
func Get_Data_List_Types_unfoldable1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_unfoldable1NonEmptyList.Do(func() {
		cache_Data_List_Types_unfoldable1NonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2435023311_2187088110(Rebox_Data_List_Types_2187088110_2435023311(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Call_Data_NonEmpty_unfoldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_132451362_2738507278(Rebox_Data_List_Types_2738507278_132451362(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_List_Types_unfoldableList()))))})))))}
	})
	return cache_Data_List_Types_unfoldable1NonEmptyList
}

var cache_Data_List_Types_foldable1NonEmptyList gopurs_runtime.Value
var once_Data_List_Types_foldable1NonEmptyList sync.Once
func Get_Data_List_Types_foldable1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_foldable1NonEmptyList.Do(func() {
		cache_Data_List_Types_foldable1NonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3368202604_4151366573(Rebox_Data_List_Types_4151366573_3368202604(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Call_Data_NonEmpty_foldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1022383170_1680800814(Rebox_Data_List_Types_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList()))))})))))}
	})
	return cache_Data_List_Types_foldable1NonEmptyList
}

var cache_Data_List_Types_extendNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_extendNonEmptyList sync.Once
func Get_Data_List_Types_extendNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_extendNonEmptyList.Do(func() {
		cache_Data_List_Types_extendNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2476964120_3290176857((&Constructor_Control_Extend_Extend[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2801299215_2812149806(Rebox_Data_List_Types_2812149806_2801299215(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Data_NonEmpty_functorNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3262795586_2812149806(Rebox_Data_List_Types_2812149806_3262795586(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Types_functorList()))))})))))}
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_2_0_21 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_2_0_21
var go__go_2_0_21 gopurs_runtime.Value
_ = go__go_2_0_21
Call_local_Data_List_Types_go__go_2_0_21 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_0_21:
for {
if false { continue go__go_2_0_21 }
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
				return gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.acc)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.val)})
				}()
v_4_loop = (v_4).V1
continue go__go_2_0_21
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
go__go_2_0_21 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_2_0_21(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
var Call_local_Data_List_Types_go__go_2_2_22 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_2_2_22
var go__go_2_2_22 gopurs_runtime.Value
_ = go__go_2_2_22
Call_local_Data_List_Types_go__go_2_2_22 = func(v_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_2_2_22:
for {
if false { continue go__go_2_2_22 }
var v_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_3_loop
_ = v_3
var v1_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_4_loop
_ = v1_4
var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_4 == nil) {
__t3 = v_3
goto end_branch_3
} else {

}
}
{
if (v1_4 != nil) {
v_3_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_4).V0, v_3})
v1_4_loop = (v1_4).V1
continue go__go_2_2_22
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
go__go_2_2_22 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_2_2_22(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_3_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_4_loop_val)))}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952((&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952(Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_1))))}), gopurs_runtime.RecordGet(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(go__go_2_0_21, func() gopurs_runtime.Value {
				orig := struct{
	acc *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	val *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}{(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)}
				_ = orig
				return gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.acc)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.val)})
				}()), gopurs_runtime.Apply(go__go_2_2_22, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1))}), "val")})))}
})})))}
	})
	return cache_Data_List_Types_extendNonEmptyList
}

var cache_Data_List_Types_extendList gopurs_runtime.Value
var once_Data_List_Types_extendList sync.Once
func Get_Data_List_Types_extendList() gopurs_runtime.Value {
	once_Data_List_Types_extendList.Do(func() {
		cache_Data_List_Types_extendList = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2171623093_3290176857((&Constructor_Control_Extend_Extend[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3262795586_2812149806(Rebox_Data_List_Types_2812149806_3262795586(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Types_functorList()))))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1)
_ = __t_tag_0
if (__t_tag_0 == nil) {
__t6 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_6
} else {

}
}
{
var __t_tag_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1)
_ = __t_tag_1
if (__t_tag_1 != nil) {
var Call_local_Data_List_Types_go__go_2_2_23 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_2_2_23
var go__go_2_2_23 gopurs_runtime.Value
_ = go__go_2_2_23
Call_local_Data_List_Types_go__go_2_2_23 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_2_23:
for {
if false { continue go__go_2_2_23 }
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
				return gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.acc)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.val)})
				}()
v_4_loop = (v_4).V1
continue go__go_2_2_23
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
go__go_2_2_23 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_2_2_23(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
var Call_local_Data_List_Types_go__go_2_4_24 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_2_4_24
var go__go_2_4_24 gopurs_runtime.Value
_ = go__go_2_4_24
Call_local_Data_List_Types_go__go_2_4_24 = func(v_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_2_4_24:
for {
if false { continue go__go_2_4_24 }
var v_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_3_loop
_ = v_3
var v1_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_4_loop
_ = v1_4
var __t5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_4 == nil) {
__t5 = v_3
goto end_branch_5
} else {

}
}
{
if (v1_4 != nil) {
v_3_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_4).V0, v_3})
v1_4_loop = (v1_4).V1
continue go__go_2_4_24
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
go__go_2_4_24 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_2_4_24(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_3_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_4_loop_val)))}
})
})
__t6 = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1))}), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(go__go_2_2_23, func() gopurs_runtime.Value {
				orig := struct{
	acc *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	val *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}{(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)}
				_ = orig
				return gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.acc)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.val)})
				}()), gopurs_runtime.Apply(go__go_2_4_24, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_1.UnsafePtr).V1)}), "val"))})
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
var go__go_3_0_25 gopurs_runtime.Value
_ = go__go_3_0_25
var go__go_3_0_25_cell *gopurs_runtime.Value
_ = go__go_3_0_25_cell
// FALLBACK TCO: isLoop=false len=1
go__go_3_0_25 = gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
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
_ = __t_tag_4
if (__t_tag_4 == nil) {
var __t_tag_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5)
_ = __t_tag_5
__t6 = ((__t_tag_5 == nil)) && ((v2_6.IntVal) != (0))
goto end_branch_6
} else {

}
}
{
var __t_tag_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4)
_ = __t_tag_1
var __t_and_3 bool = false
if (__t_tag_1 != nil) {

var __t_tag_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5)
_ = __t_tag_2
__t_and_3 = ((__t_tag_2 != nil)) && ((gopurs_runtime.Bool((gopurs_runtime.Apply3((*go__go_3_0_25_cell), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}, gopurs_runtime.Bool(((v2_6.IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0).IntVal) != (0)))).IntVal) != (0)).IntVal) != (0))
}
__t6 = __t_and_3
}
end_branch_6:
return gopurs_runtime.Bool(__t6)
})
go__go_3_0_25_cell = &go__go_3_0_25
return gopurs_runtime.Bool((gopurs_runtime.Apply3(go__go_3_0_25, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_2))}, gopurs_runtime.Bool(true)).IntVal) != (0))
})})))}
	})
	return cache_Data_List_Types_eq1List
}

var cache_Data_List_Types_eq1 gopurs_runtime.Value
var once_Data_List_Types_eq1 sync.Once
func Get_Data_List_Types_eq1() gopurs_runtime.Value {
	once_Data_List_Types_eq1.Do(func() {
		cache_Data_List_Types_eq1 = Call_Data_Eq_eq1(Rebox_Data_List_Types_3554500787_1766074591(Rebox_Data_List_Types_1766074591_3554500787(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_List_Types_eq1List()))))
	})
	return cache_Data_List_Types_eq1
}

var cache_Data_List_Types_eqNonEmpty gopurs_runtime.Value
var once_Data_List_Types_eqNonEmpty sync.Once
func Get_Data_List_Types_eqNonEmpty() gopurs_runtime.Value {
	once_Data_List_Types_eqNonEmpty.Do(func() {
		cache_Data_List_Types_eqNonEmpty = gopurs_runtime.Apply(Get_Data_NonEmpty_eqNonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3554500787_1766074591(Rebox_Data_List_Types_1766074591_3554500787(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_List_Types_eq1List()))))})
	})
	return cache_Data_List_Types_eqNonEmpty
}

var cache_Data_List_Types_eq1NonEmptyList gopurs_runtime.Value
var once_Data_List_Types_eq1NonEmptyList sync.Once
func Get_Data_List_Types_eq1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_eq1NonEmptyList.Do(func() {
		cache_Data_List_Types_eq1NonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2625657118_1766074591(Rebox_Data_List_Types_1766074591_2625657118(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Call_Data_NonEmpty_eq1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3554500787_1766074591(Rebox_Data_List_Types_1766074591_3554500787(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_List_Types_eq1List()))))})))))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3554500787_1766074591(Rebox_Data_List_Types_1766074591_3554500787(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_List_Types_eq1List()))))}
}), gopurs_runtime.Func3(func(dictOrd_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value, ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_3_0_26 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) uint32
_ = Call_local_Data_List_Types_go__go_3_0_26
var go__go_3_0_26 gopurs_runtime.Value
_ = go__go_3_0_26
Call_local_Data_List_Types_go__go_3_0_26 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) uint32 {
go__go_3_0_26:
for {
if false { continue go__go_3_0_26 }
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
continue go__go_3_0_26
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
go__go_3_0_26 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_List_Types_go__go_3_0_26(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val))), UnsafePtr: nil}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_List_Types_go__go_3_0_26(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_2))), UnsafePtr: nil}
})})))}
	})
	return cache_Data_List_Types_ord1List
}

var cache_Data_List_Types_compare1 gopurs_runtime.Value
var once_Data_List_Types_compare1 sync.Once
func Get_Data_List_Types_compare1() gopurs_runtime.Value {
	once_Data_List_Types_compare1.Do(func() {
		cache_Data_List_Types_compare1 = Call_Data_Ord_compare1(Rebox_Data_List_Types_3739089555_3985601471(Rebox_Data_List_Types_3985601471_3739089555(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](Get_Data_List_Types_ord1List()))))
	})
	return cache_Data_List_Types_compare1
}

var cache_Data_List_Types_ordNonEmpty gopurs_runtime.Value
var once_Data_List_Types_ordNonEmpty sync.Once
func Get_Data_List_Types_ordNonEmpty() gopurs_runtime.Value {
	once_Data_List_Types_ordNonEmpty.Do(func() {
		cache_Data_List_Types_ordNonEmpty = Call_Data_NonEmpty_ordNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3739089555_3985601471(Rebox_Data_List_Types_3985601471_3739089555(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](Get_Data_List_Types_ord1List()))))})
	})
	return cache_Data_List_Types_ordNonEmpty
}

var cache_Data_List_Types_ord1NonEmptyList gopurs_runtime.Value
var once_Data_List_Types_ord1NonEmptyList sync.Once
func Get_Data_List_Types_ord1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_ord1NonEmptyList.Do(func() {
		cache_Data_List_Types_ord1NonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3290489278_3985601471(Rebox_Data_List_Types_3985601471_3290489278(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](Call_Data_NonEmpty_ord1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3739089555_3985601471(Rebox_Data_List_Types_3985601471_3739089555(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](Get_Data_List_Types_ord1List()))))})))))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2476964120_3290176857(Rebox_Data_List_Types_3290176857_2476964120(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend[gopurs_runtime.Value]](Get_Data_List_Types_extendNonEmptyList()))))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3262795586_2812149806(Rebox_Data_List_Types_2812149806_3262795586(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Types_functorList()))))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0)
_ = __t_tag_0
if (__t_tag_0 == nil) {
__t6 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_6
} else {

}
}
{
var __t_tag_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0)
_ = __t_tag_1
if (__t_tag_1 != nil) {
var Call_local_Data_List_Types_go__go_2_2_27 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_2_2_27
var go__go_2_2_27 gopurs_runtime.Value
_ = go__go_2_2_27
Call_local_Data_List_Types_go__go_2_2_27 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_2_27:
for {
if false { continue go__go_2_2_27 }
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
continue go__go_2_2_27
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
go__go_2_2_27 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_2_2_27(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
var Call_local_Data_List_Types_go__go_2_4_28 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_2_4_28
var go__go_2_4_28 gopurs_runtime.Value
_ = go__go_2_4_28
Call_local_Data_List_Types_go__go_2_4_28 = func(v_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_2_4_28:
for {
if false { continue go__go_2_4_28 }
var v_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_3_loop
_ = v_3
var v1_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_4_loop
_ = v1_4
var __t5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_4 == nil) {
__t5 = v_3
goto end_branch_5
} else {

}
}
{
if (v1_4 != nil) {
v_3_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_4).V0, v_3})
v1_4_loop = (v1_4).V1
continue go__go_2_4_28
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
go__go_2_4_28 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_2_4_28(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_3_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_4_loop_val)))}
})
})
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(go__go_2_2_27, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Data_List_Types_3741347833_1305434581(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_List_Types_applyList())).V1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1))})))}), gopurs_runtime.Apply(go__go_2_4_28, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_List_Types_listMap((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1))})))}))
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
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2801299215_2812149806(Rebox_Data_List_Types_2812149806_2801299215(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Data_NonEmpty_functorNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3262795586_2812149806(Rebox_Data_List_Types_2812149806_3262795586(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Types_functorList()))))})))))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_2_0_29 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_2_0_29
var go__go_2_0_29 gopurs_runtime.Value
_ = go__go_2_0_29
Call_local_Data_List_Types_go__go_2_0_29 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_0_29:
for {
if false { continue go__go_2_0_29 }
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
continue go__go_2_0_29
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
go__go_2_0_29 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_2_0_29(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
var Call_local_Data_List_Types_go__go_3_2_30 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_3_2_30
var go__go_3_2_30 gopurs_runtime.Value
_ = go__go_3_2_30
Call_local_Data_List_Types_go__go_3_2_30 = func(b_4_loop gopurs_runtime.Value, v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_3_2_30:
for {
if false { continue go__go_3_2_30 }
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
continue go__go_3_2_30
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
go__go_3_2_30 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_3_2_30(b_4_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val))
})
})
var Call_local_Data_List_Types_go__go_3_4_31 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_3_4_31
var go__go_3_4_31 gopurs_runtime.Value
_ = go__go_3_4_31
Call_local_Data_List_Types_go__go_3_4_31 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_3_4_31:
for {
if false { continue go__go_3_4_31 }
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
continue go__go_3_4_31
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
go__go_3_4_31 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_3_4_31(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val)))}
})
})
var Call_local_Data_List_Types_go__go_2_6_32 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_2_6_32
var go__go_2_6_32 gopurs_runtime.Value
_ = go__go_2_6_32
Call_local_Data_List_Types_go__go_2_6_32 = func(v_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_2_6_32:
for {
if false { continue go__go_2_6_32 }
var v_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_3_loop
_ = v_3
var v1_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_4_loop
_ = v1_4
var __t7 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_4 == nil) {
__t7 = v_3
goto end_branch_7
} else {

}
}
{
if (v1_4 != nil) {
v_3_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_4).V0, v_3})
v1_4_loop = (v1_4).V1
continue go__go_2_6_32
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
go__go_2_6_32 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_2_6_32(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_3_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_4_loop_val)))}
})
})
var __t14 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_8 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1)
_ = __t_tag_8
if (__t_tag_8 == nil) {
__t14 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_14
} else {

}
}
{
var __t_tag_9 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1)
_ = __t_tag_9
if (__t_tag_9 != nil) {
var Call_local_Data_List_Types_go__go_2_10_33 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_2_10_33
var go__go_2_10_33 gopurs_runtime.Value
_ = go__go_2_10_33
Call_local_Data_List_Types_go__go_2_10_33 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_10_33:
for {
if false { continue go__go_2_10_33 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var __t11 gopurs_runtime.Value
{
if (v_4 == nil) {
__t11 = b_3
goto end_branch_11
} else {

}
}
{
if (v_4 != nil) {
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_4).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_3)}))}
v_4_loop = (v_4).V1
continue go__go_2_10_33
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
go__go_2_10_33 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_2_10_33(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
var Call_local_Data_List_Types_go__go_2_12_34 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_2_12_34
var go__go_2_12_34 gopurs_runtime.Value
_ = go__go_2_12_34
Call_local_Data_List_Types_go__go_2_12_34 = func(v_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_2_12_34:
for {
if false { continue go__go_2_12_34 }
var v_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_3_loop
_ = v_3
var v1_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_4_loop
_ = v1_4
var __t13 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_4 == nil) {
__t13 = v_3
goto end_branch_13
} else {

}
}
{
if (v1_4 != nil) {
v_3_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_4).V0, v_3})
v1_4_loop = (v1_4).V1
continue go__go_2_12_34
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
go__go_2_12_34 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_2_12_34(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_3_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_4_loop_val)))}
})
})
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(go__go_2_10_33, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Data_List_Types_3741347833_1305434581(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_List_Types_applyList())).V1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)}))})))}), gopurs_runtime.Apply(go__go_2_12_34, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_List_Types_listMap((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)}))})))}))
goto end_branch_14
} else {

}
}
{
__t14 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_14:
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952((&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, gopurs_runtime.Apply((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(go__go_2_0_29, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(go__go_3_2_30, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Data_List_Types_3741347833_1305434581(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_List_Types_applyList())).V1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1))})))}), gopurs_runtime.Apply(go__go_3_4_31, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_List_Types_listMap((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1))})))})))}), gopurs_runtime.Apply(go__go_2_6_32, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t14)})))}})))}
})})))}
	})
	return cache_Data_List_Types_applyNonEmptyList
}

var cache_Data_List_Types_bindList gopurs_runtime.Value
var once_Data_List_Types_bindList sync.Once
func Get_Data_List_Types_bindList() gopurs_runtime.Value {
	once_Data_List_Types_bindList.Do(func() {
		cache_Data_List_Types_bindList = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2183599445_2748095225((&Constructor_Control_Bind_Bind[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1305434581_3741347833(Rebox_Data_List_Types_3741347833_1305434581(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_List_Types_applyList()))))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0)
_ = __t_tag_0
if (__t_tag_0 == nil) {
__t6 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_6
} else {

}
}
{
var __t_tag_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0)
_ = __t_tag_1
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
b_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_4).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_3)}))}
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
var Call_local_Data_List_Types_go__go_2_4_36 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_2_4_36
var go__go_2_4_36 gopurs_runtime.Value
_ = go__go_2_4_36
Call_local_Data_List_Types_go__go_2_4_36 = func(v_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_2_4_36:
for {
if false { continue go__go_2_4_36 }
var v_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_3_loop
_ = v_3
var v1_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_4_loop
_ = v1_4
var __t5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_4 == nil) {
__t5 = v_3
goto end_branch_5
} else {

}
}
{
if (v1_4 != nil) {
v_3_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_4).V0, v_3})
v1_4_loop = (v1_4).V1
continue go__go_2_4_36
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
go__go_2_4_36 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_2_4_36(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_3_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_4_loop_val)))}
})
})
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(go__go_2_2_35, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Data_List_Types_2748095225_2183599445(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Types_bindList())).V1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V1)}, v1_1)))}), gopurs_runtime.Apply(go__go_2_4_36, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V0)))}))
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
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3970790648_3741347833(Rebox_Data_List_Types_3741347833_3970790648(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_List_Types_applyNonEmptyList()))))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_2_0 shape=App(Other) bindingType=(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (TypeVar b$scope264)])
v1_2_0 := Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0)))
_ = v1_2_0
var Call_local_Data_List_Types_go__go_3_1_37 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_3_1_37
var go__go_3_1_37 gopurs_runtime.Value
_ = go__go_3_1_37
Call_local_Data_List_Types_go__go_3_1_37 = func(b_4_loop gopurs_runtime.Value, v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_3_1_37:
for {
if false { continue go__go_3_1_37 }
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
continue go__go_3_1_37
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
go__go_3_1_37 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_3_1_37(b_4_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val))
})
})
// TAST (Let): __local_var_4_3 shape=App(Var) bindingType=(Func [(TypeVar a$scope263)] (ADT ["Data","List","Types","List"] [(TypeVar b$scope264)]))
__local_var_4_3 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Types_toList(), f_1)
_ = __local_var_4_3
var __t10 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1)
_ = __t_tag_4
if (__t_tag_4 == nil) {
__t10 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_10
} else {

}
}
{
var __t_tag_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1)
_ = __t_tag_5
if (__t_tag_5 != nil) {
var Call_local_Data_List_Types_go__go_5_6_38 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_5_6_38
var go__go_5_6_38 gopurs_runtime.Value
_ = go__go_5_6_38
Call_local_Data_List_Types_go__go_5_6_38 = func(b_6_loop gopurs_runtime.Value, v_7_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_5_6_38:
for {
if false { continue go__go_5_6_38 }
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
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_7).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_6)}))}
v_7_loop = (v_7).V1
continue go__go_5_6_38
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
go__go_5_6_38 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_5_6_38(b_6_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_7_loop_val))
})
})
var Call_local_Data_List_Types_go__go_5_8_39 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_5_8_39
var go__go_5_8_39 gopurs_runtime.Value
_ = go__go_5_8_39
Call_local_Data_List_Types_go__go_5_8_39 = func(v_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_7_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_5_8_39:
for {
if false { continue go__go_5_8_39 }
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
continue go__go_5_8_39
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
go__go_5_8_39 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_5_8_39(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_6_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_7_loop_val)))}
})
})
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(go__go_5_6_38, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Data_List_Types_2748095225_2183599445(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Types_bindList())).V1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1.UnsafePtr).V1)}, __local_var_4_3)))}), gopurs_runtime.Apply(go__go_5_8_39, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_4_3, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1.UnsafePtr).V0)))}))
goto end_branch_10
} else {

}
}
{
__t10 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_10:
var Call_local_Data_List_Types_go__go_3_11_40 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Types_go__go_3_11_40
var go__go_3_11_40 gopurs_runtime.Value
_ = go__go_3_11_40
Call_local_Data_List_Types_go__go_3_11_40 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_3_11_40:
for {
if false { continue go__go_3_11_40 }
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_5_loop
_ = v1_5
var __t12 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_5 == nil) {
__t12 = v_4
goto end_branch_12
} else {

}
}
{
if (v1_5 != nil) {
v_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_5).V0, v_4})
v1_5_loop = (v1_5).V1
continue go__go_3_11_40
__t12 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_12
} else {

}
}
{
__t12 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
}
}
go__go_3_11_40 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_go__go_3_11_40(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val)))}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952((&Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{1, (v1_2_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(go__go_3_1_37, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t10)}), gopurs_runtime.Apply(go__go_3_11_40, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((v1_2_0).V1))})))}})))}
})})))}
	})
	return cache_Data_List_Types_bindNonEmptyList
}

var cache_Data_List_Types_applicativeList gopurs_runtime.Value
var once_Data_List_Types_applicativeList sync.Once
func Get_Data_List_Types_applicativeList() gopurs_runtime.Value {
	once_Data_List_Types_applicativeList.Do(func() {
		cache_Data_List_Types_applicativeList = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2870828117_1439734649((&Constructor_Control_Applicative_Applicative[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1305434581_3741347833(Rebox_Data_List_Types_3741347833_1305434581(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_List_Types_applyList()))))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2870828117_1439734649(Rebox_Data_List_Types_1439734649_2870828117(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_List_Types_applicativeList()))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2183599445_2748095225(Rebox_Data_List_Types_2748095225_2183599445(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Types_bindList()))))}
})})))}
	})
	return cache_Data_List_Types_monadList
}

var cache_Data_List_Types_altNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_altNonEmptyList sync.Once
func Get_Data_List_Types_altNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_altNonEmptyList.Do(func() {
		cache_Data_List_Types_altNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_203237368_3421983481((&Constructor_Control_Alt_Alt[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2801299215_2812149806(Rebox_Data_List_Types_2812149806_2801299215(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Data_NonEmpty_functorNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3262795586_2812149806(Rebox_Data_List_Types_2812149806_3262795586(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Types_functorList()))))})))))}
}), Call_Data_Semigroup_go__append(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_List_Types_semigroupNonEmptyList()))})))}
	})
	return cache_Data_List_Types_altNonEmptyList
}

var cache_Data_List_Types_altList gopurs_runtime.Value
var once_Data_List_Types_altList sync.Once
func Get_Data_List_Types_altList() gopurs_runtime.Value {
	once_Data_List_Types_altList.Do(func() {
		cache_Data_List_Types_altList = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_257347157_3421983481((&Constructor_Control_Alt_Alt[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3262795586_2812149806(Rebox_Data_List_Types_2812149806_3262795586(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Types_functorList()))))}
}), Call_Data_Semigroup_go__append(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_List_Types_semigroupList()))})))}
	})
	return cache_Data_List_Types_altList
}

var cache_Data_List_Types_plusList gopurs_runtime.Value
var once_Data_List_Types_plusList sync.Once
func Get_Data_List_Types_plusList() gopurs_runtime.Value {
	once_Data_List_Types_plusList.Do(func() {
		cache_Data_List_Types_plusList = gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2173899317_3706288089((&Constructor_Control_Plus_Plus[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_257347157_3421983481(Rebox_Data_List_Types_3421983481_257347157(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](Get_Data_List_Types_altList()))))}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}})))}
	})
	return cache_Data_List_Types_plusList
}

var cache_Data_List_Types_alternativeList gopurs_runtime.Value
var once_Data_List_Types_alternativeList sync.Once
func Get_Data_List_Types_alternativeList() gopurs_runtime.Value {
	once_Data_List_Types_alternativeList.Do(func() {
		cache_Data_List_Types_alternativeList = gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_739440277_2307501113((&Constructor_Control_Alternative_Alternative[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2870828117_1439734649(Rebox_Data_List_Types_1439734649_2870828117(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_List_Types_applicativeList()))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2173899317_3706288089(Rebox_Data_List_Types_3706288089_2173899317(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](Get_Data_List_Types_plusList()))))}
})})))}
	})
	return cache_Data_List_Types_alternativeList
}

var cache_Data_List_Types_monadPlusList gopurs_runtime.Value
var once_Data_List_Types_monadPlusList sync.Once
func Get_Data_List_Types_monadPlusList() gopurs_runtime.Value {
	once_Data_List_Types_monadPlusList.Do(func() {
		cache_Data_List_Types_monadPlusList = gopurs_runtime.Value{Type: 9, IntVal: 3236234573, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2656832021_2394491833((&Constructor_Control_MonadPlus_MonadPlus[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_739440277_2307501113(Rebox_Data_List_Types_2307501113_739440277(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](Get_Data_List_Types_alternativeList()))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3994390613_2568689657(Rebox_Data_List_Types_2568689657_3994390613(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Data_List_Types_monadList()))))}
})})))}
	})
	return cache_Data_List_Types_monadPlusList
}

var cache_Data_List_Types_applicativeNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_applicativeNonEmptyList sync.Once
func Get_Data_List_Types_applicativeNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_applicativeNonEmptyList.Do(func() {
		cache_Data_List_Types_applicativeNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2334388216_1439734649((&Constructor_Control_Applicative_Applicative[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3970790648_3741347833(Rebox_Data_List_Types_3741347833_3970790648(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_List_Types_applyNonEmptyList()))))}
}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Types_NonEmptyList(), Call_Data_NonEmpty_singleton(Rebox_Data_List_Types_2173899317_3706288089(Rebox_Data_List_Types_3706288089_2173899317(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](Get_Data_List_Types_plusList())))))})))}
	})
	return cache_Data_List_Types_applicativeNonEmptyList
}

var cache_Data_List_Types_pure gopurs_runtime.Value
var once_Data_List_Types_pure sync.Once
func Get_Data_List_Types_pure() gopurs_runtime.Value {
	once_Data_List_Types_pure.Do(func() {
		cache_Data_List_Types_pure = Call_Control_Applicative_pure(Rebox_Data_List_Types_2334388216_1439734649(Rebox_Data_List_Types_1439734649_2334388216(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_List_Types_applicativeNonEmptyList()))))
	})
	return cache_Data_List_Types_pure
}

var cache_Data_List_Types_monadNonEmptyList gopurs_runtime.Value
var once_Data_List_Types_monadNonEmptyList sync.Once
func Get_Data_List_Types_monadNonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_monadNonEmptyList.Do(func() {
		cache_Data_List_Types_monadNonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1757775096_2568689657((&Constructor_Control_Monad_Monad[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_2334388216_1439734649(Rebox_Data_List_Types_1439734649_2334388216(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_List_Types_applicativeNonEmptyList()))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3054666744_2748095225(Rebox_Data_List_Types_2748095225_3054666744(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Types_bindNonEmptyList()))))}
})})))}
	})
	return cache_Data_List_Types_monadNonEmptyList
}

var cache_Data_List_Types_traversable1NonEmptyList gopurs_runtime.Value
var once_Data_List_Types_traversable1NonEmptyList sync.Once
func Get_Data_List_Types_traversable1NonEmptyList() gopurs_runtime.Value {
	once_Data_List_Types_traversable1NonEmptyList.Do(func() {
		cache_Data_List_Types_traversable1NonEmptyList = gopurs_runtime.Value{Type: 9, IntVal: 1596088409, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3842311788_306175789((&Constructor_Data_Semigroup_Traversable_Traversable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3368202604_4151366573(Rebox_Data_List_Types_4151366573_3368202604(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Call_Data_NonEmpty_foldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1022383170_1680800814(Rebox_Data_List_Types_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList()))))})))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3425342735_3043886126(Rebox_Data_List_Types_3043886126_3425342735(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Call_Data_NonEmpty_traversableNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3037784642_3043886126(Rebox_Data_List_Types_3043886126_3037784642(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_List_Types_traversableList()))))})))))}
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Rebox_Data_List_Types_306175789_3842311788(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]](Get_Data_List_Types_traversable1NonEmptyList())).V3, gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0))}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f$scope111)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_4_3_42 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_4_3_42
var go__go_4_3_42 gopurs_runtime.Value
_ = go__go_4_3_42
Call_local_Data_List_Types_go__go_4_3_42 = func(b_5_loop gopurs_runtime.Value, v_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_4_3_42:
for {
if false { continue go__go_4_3_42 }
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
// TAST (Let): Functor0_7_4 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f$scope44)])
Functor0_7_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_4
b_5_loop = gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(Functor0_7_4.V0, gopurs_runtime.Func2(func(b_9 gopurs_runtime.Value, a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952(Call_Data_List_Types_nelCons(a_10, Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](b_9)))))}
}), b_5), b_8)
}), f_2, (v_6).V0)
v_6_loop = (v_6).V1
continue go__go_4_3_42
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
go__go_4_3_42 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_4_3_42(b_5_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_6_loop_val))
})
})
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Types_go__go_5_1_41 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Types_go__go_5_1_41
var go__go_5_1_41 gopurs_runtime.Value
_ = go__go_5_1_41
Call_local_Data_List_Types_go__go_5_1_41 = func(b_6_loop gopurs_runtime.Value, v_7_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_5_1_41:
for {
if false { continue go__go_5_1_41 }
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
b_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952(Call_Data_List_Types_nelCons((v_7).V0, Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](b_6)))))}
v_7_loop = (v_7).V1
continue go__go_5_1_41
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
go__go_5_1_41 = gopurs_runtime.Func(func(b_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Types_go__go_5_1_41(b_6_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_7_loop_val))
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952(Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](Call_local_Data_List_Types_go__go_5_1_41(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952(Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply3(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Types_NonEmptyList(), Call_Data_NonEmpty_singleton(Rebox_Data_List_Types_2173899317_3706288089(Rebox_Data_List_Types_3706288089_2173899317(gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](Get_Data_List_Types_plusList())))), (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0)))))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1))))))}
}), Call_local_Data_List_Types_go__go_4_3_42(gopurs_runtime.Apply2(Functor0_1_0.V0, Call_Control_Applicative_pure(Rebox_Data_List_Types_2334388216_1439734649(Rebox_Data_List_Types_1439734649_2334388216(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_List_Types_applicativeNonEmptyList())))), gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)))
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


func Call_Data_List_Types_Cons__1027980219(__eta_norm_1_0_loop *Constructor_Data_Tuple_Tuple[uint32, float64], __eta_norm_0_1_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
Cons__1027980219:
for {
if false { continue Cons__1027980219 }
var __eta_norm_1_0 *Constructor_Data_Tuple_Tuple[uint32, float64] = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return (&Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]{1, __eta_norm_1_0, __eta_norm_0_1})
}
}

func Call_Data_List_Types_Cons__351658551(__eta_norm_1_0_loop int64, __eta_norm_0_unused_1_loop *Constructor_Data_List_Types_Cons[int64]) *Constructor_Data_List_Types_Cons[int64] {
Cons__351658551:
for {
if false { continue Cons__351658551 }
var __eta_norm_1_0 int64 = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_unused_1 *Constructor_Data_List_Types_Cons[int64] = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return (&Constructor_Data_List_Types_Cons[int64]{1, __eta_norm_1_0, (*Constructor_Data_List_Types_Cons[int64])(nil)})
}
}

func Call_Data_List_Types_Cons__317927332(__eta_norm_1_0_loop int64, __eta_norm_0_1_loop *Constructor_Data_List_Types_Cons[int64]) *Constructor_Data_List_Types_Cons[int64] {
Cons__317927332:
for {
if false { continue Cons__317927332 }
var __eta_norm_1_0 int64 = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 *Constructor_Data_List_Types_Cons[int64] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return (&Constructor_Data_List_Types_Cons[int64]{1, __eta_norm_1_0, __eta_norm_0_1})
}
}

func Call_Data_List_Types_NonEmptyList(x_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var x_0 *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = x_0_loop
_ = x_0
return x_0
}

func Call_Data_List_Types_NonEmptyList__2988773242(x_0_loop *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
NonEmptyList__2988773242:
for {
if false { continue NonEmptyList__2988773242 }
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
_ = __t_tag_15
var __t_and_17 bool = false
if (__t_tag_15 != nil) {

var __t_tag_16 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = ((v1_3).V1).V1
_ = __t_tag_16
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
_ = __t_tag_2
var __t_and_6 bool = false
if (__t_tag_2 != nil) {

var __t_tag_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = ((v2_5).V0).V1
_ = __t_tag_3
var __t_and_5 bool = false
if (__t_tag_3 != nil) {

var __t_tag_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = (((v2_5).V0).V1).V1
_ = __t_tag_4
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
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_reverseUnrolledMap_4_1_1(Rebox_Data_List_Types_849153993_1116310629(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_5_loop_val)), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v3_6_loop_val)))}
})
})
var __t14 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_3 != nil) {
var __t13 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_9 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = (v1_3).V1
_ = __t_tag_9
if (__t_tag_9 != nil) {
var __t11 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_10 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = ((v1_3).V1).V1
_ = __t_tag_10
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
_ = __t_tag_12
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
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Types_chunkedRevMap_1_0_0(Rebox_Data_List_Types_849153993_1116310629(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val)), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val)))}
})
})
return gopurs_runtime.Apply(chunkedRevMap_1_0_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1116310629_849153993((*Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(nil)))})
}

func Call_Data_List_Types_showList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): show_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope68)] String)
show_1_0 := Call_Data_Show_show(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](dictShow_0))
_ = show_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3351995458_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 string
{
var __t_tag_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2)
_ = __t_tag_1
if (__t_tag_1 == nil) {
__t2 = "Nil"
goto end_branch_2
} else {

}
}
{
__t2 = (("(") + (Call_Data_Foldable_intercalate__4228433826(" : ", Rebox_Data_List_Types_849153993_128126966(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_List_Types_listMap(show_1_0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2))})))))) + (" : Nil)")
}
end_branch_2:
return gopurs_runtime.Str(__t2)
})})))}
}

func Call_Data_List_Types_showNonEmptyList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): showNonEmpty_1_0 shape=App(Var) bindingType=(ADT ["Data","Show","Show"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (TypeVar a$scope65)])])
showNonEmpty_1_0 := Rebox_Data_List_Types_1386611502_4061813263(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_NonEmpty_showNonEmpty(dictShow_0, Call_Data_List_Types_showList(dictShow_0))))
_ = showNonEmpty_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_4061813263_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(NonEmptyList ") + (gopurs_runtime.Apply(showNonEmpty_1_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3123684004_1293498952(Rebox_Data_List_Types_1293498952_3123684004(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_2))))}).StrVal())) + (")"))
})})))}
}

func Call_Data_List_Types_eqList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1636902754_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Apply(Call_Data_Eq_eq1(Rebox_Data_List_Types_3554500787_1766074591(Rebox_Data_List_Types_1766074591_3554500787(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_List_Types_eq1List())))), dictEq_0)})))}
}

func Call_Data_List_Types_eqNonEmptyList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3912963311_3790796878(Rebox_Data_List_Types_3790796878_3912963311(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_NonEmpty_eqNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3554500787_1766074591(Rebox_Data_List_Types_1766074591_3554500787(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_List_Types_eq1List()))))}, dictEq_0)))))}
}

func Call_Data_List_Types_ordList(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): eqList1_1_0 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","List","Types","List"] [(TypeVar a$scope87)])])
eqList1_1_0 := Rebox_Data_List_Types_3790796878_1636902754(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_List_Types_eqList(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))))
_ = eqList1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_4210054658_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1636902754_3790796878(eqList1_1_0))}
}), gopurs_runtime.Apply(Call_Data_Ord_compare1(Rebox_Data_List_Types_3739089555_3985601471(Rebox_Data_List_Types_3985601471_3739089555(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](Get_Data_List_Types_ord1List())))), dictOrd_0)})))}
}

func Call_Data_List_Types_ordNonEmptyList(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_1180711887_4177771502(Rebox_Data_List_Types_4177771502_1180711887(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_NonEmpty_ordNonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Types_3739089555_3985601471(Rebox_Data_List_Types_3985601471_3739089555(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](Get_Data_List_Types_ord1List()))))}), dictOrd_0)))))}
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

func Rebox_Data_List_Types_1386611502_4061813263(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Types_1439734649_2334388216(in *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) *Constructor_Control_Applicative_Applicative[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Applicative_Applicative[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_1439734649_2870828117(in *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) *Constructor_Control_Applicative_Applicative[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Applicative_Applicative[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
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

func Rebox_Data_List_Types_1680800814_1022383170(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_List_Types_1680800814_4064382095(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
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

func Rebox_Data_List_Types_1766074591_2625657118(in *Constructor_Data_Eq_Eq1[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Types_1766074591_3554500787(in *Constructor_Data_Eq_Eq1[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq1[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq1[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Types_1812164904_1544744991(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
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

func Rebox_Data_List_Types_2183599445_2748095225(in *Constructor_Control_Bind_Bind[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Control_Bind_Bind[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Bind_Bind[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_2187088110_2435023311(in *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]) *Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Types_2187088110_4073635714(in *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]) *Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Unfoldable1_Unfoldable1[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Types_2307501113_739440277(in *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]) *Constructor_Control_Alternative_Alternative[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Alternative_Alternative[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
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

func Rebox_Data_List_Types_2412140840_111597075(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_FunctorWithIndex_FunctorWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_2412140840_3339399026(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_2412140840_721753375(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
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

func Rebox_Data_List_Types_2568689657_3994390613(in *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) *Constructor_Control_Monad_Monad[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Monad[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
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

func Rebox_Data_List_Types_2738507278_132451362(in *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) *Constructor_Data_Unfoldable_Unfoldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Unfoldable_Unfoldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_2748095225_2183599445(in *Constructor_Control_Bind_Bind[gopurs_runtime.Value]) *Constructor_Control_Bind_Bind[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Bind_Bind[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_2748095225_3054666744(in *Constructor_Control_Bind_Bind[gopurs_runtime.Value]) *Constructor_Control_Bind_Bind[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Bind_Bind[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{}
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

func Rebox_Data_List_Types_2812149806_2801299215(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Types_2812149806_3262795586(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
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

func Rebox_Data_List_Types_3043886126_3037784642(in *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) *Constructor_Data_Traversable_Traversable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Traversable_Traversable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
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

func Rebox_Data_List_Types_306175789_3842311788(in *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Traversable_Traversable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Traversable_Traversable1[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_List_Types_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[int64]{}
		out.V0 = in.V0.IntVal
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

func Rebox_Data_List_Types_3290176857_2476964120(in *Constructor_Control_Extend_Extend[gopurs_runtime.Value]) *Constructor_Control_Extend_Extend[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Extend_Extend[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
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

func Rebox_Data_List_Types_3421983481_257347157(in *Constructor_Control_Alt_Alt[gopurs_runtime.Value]) *Constructor_Control_Alt_Alt[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Alt_Alt[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
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

func Rebox_Data_List_Types_3706288089_2173899317(in *Constructor_Control_Plus_Plus[gopurs_runtime.Value]) *Constructor_Control_Plus_Plus[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Plus_Plus[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_3725484264_1951493170(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
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

func Rebox_Data_List_Types_3725484264_538456415(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
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

func Rebox_Data_List_Types_3741347833_1305434581(in *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) *Constructor_Control_Apply_Apply[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Apply_Apply[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_Types_3741347833_3970790648(in *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) *Constructor_Control_Apply_Apply[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Apply_Apply[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{}
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

func Rebox_Data_List_Types_3790796878_1636902754(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Types_3790796878_3912963311(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{}
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

func Rebox_Data_List_Types_3985601471_3739089555(in *Constructor_Data_Ord_Ord1[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord1[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord1[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
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

func Rebox_Data_List_Types_849153993_1116310629(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](in.V0)
		out.V1 = Rebox_Data_List_Types_849153993_1116310629(in.V1)
	return out
}

func Rebox_Data_List_Types_849153993_128126966(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[string] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[string]{}
		out.V0 = in.V0.StrVal()
		out.V1 = Rebox_Data_List_Types_849153993_128126966(in.V1)
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



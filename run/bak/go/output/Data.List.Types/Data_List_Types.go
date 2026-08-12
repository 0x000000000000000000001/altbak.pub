package Data_List_Types

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_FoldableWithIndex "gopurs/output/Data.FoldableWithIndex"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_FunctorWithIndex "gopurs/output/Data.FunctorWithIndex"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_List_Lazy_Types "gopurs/output/Data.List.Lazy.Types"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Monoid "gopurs/output/Data.Monoid"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semigroup_Traversable "gopurs/output/Data.Semigroup.Traversable"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Data_TraversableWithIndex "gopurs/output/Data.TraversableWithIndex"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		cache_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity(x_0_box)
})
	})
	return cache_identity
}

var cache_identity1 gopurs_runtime.Value
var once_identity1 sync.Once
func Get_identity1() gopurs_runtime.Value {
	once_identity1.Do(func() {
		cache_identity1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity1(x_0_box)
})
	})
	return cache_identity1
}

var cache_Nil gopurs_runtime.Value
var once_Nil sync.Once
func Get_Nil() gopurs_runtime.Value {
	once_Nil.Do(func() {
		cache_Nil = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
	})
	return cache_Nil
}

var cache_Cons gopurs_runtime.Value
var once_Cons sync.Once
func Get_Cons() gopurs_runtime.Value {
	once_Cons.Do(func() {
		cache_Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, value0, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](value1)})}
})
})
	})
	return cache_Cons
}

var cache_NonEmptyList gopurs_runtime.Value
var once_NonEmptyList sync.Once
func Get_NonEmptyList() gopurs_runtime.Value {
	once_NonEmptyList.Do(func() {
		cache_NonEmptyList = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_NonEmptyList(x_0_box)
})
	})
	return cache_NonEmptyList
}

var cache_toList gopurs_runtime.Value
var once_toList sync.Once
func Get_toList() gopurs_runtime.Value {
	once_toList.Do(func() {
		cache_toList = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_toList(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_toList
}

var cache_newtypeNonEmptyList gopurs_runtime.Value
var once_newtypeNonEmptyList sync.Once
func Get_newtypeNonEmptyList() gopurs_runtime.Value {
	once_newtypeNonEmptyList.Do(func() {
		cache_newtypeNonEmptyList = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeNonEmptyList
}

var cache_nelCons gopurs_runtime.Value
var once_nelCons sync.Once
func Get_nelCons() gopurs_runtime.Value {
	once_nelCons.Do(func() {
		cache_nelCons = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_nelCons(a_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box)))}
})
	})
	return cache_nelCons
}

var cache_listMap gopurs_runtime.Value
var once_listMap sync.Once
func Get_listMap() gopurs_runtime.Value {
	once_listMap.Do(func() {
		cache_listMap = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_listMap(f_0_box)
})
	})
	return cache_listMap
}

var cache_functorList gopurs_runtime.Value
var once_functorList sync.Once
func Get_functorList() gopurs_runtime.Value {
	once_functorList.Do(func() {
		cache_functorList = gopurs_runtime.RecordDict1("map", Get_listMap())
	})
	return cache_functorList
}

var cache_functorNonEmptyList gopurs_runtime.Value
var once_functorNonEmptyList sync.Once
func Get_functorNonEmptyList() gopurs_runtime.Value {
	once_functorNonEmptyList.Do(func() {
		cache_functorNonEmptyList = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorList(), "map"), f_0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_functorNonEmptyList
}

var cache_foldableList gopurs_runtime.Value
var once_foldableList sync.Once
func Get_foldableList() gopurs_runtime.Value {
	once_foldableList.Do(func() {
		cache_foldableList = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(Semigroup0_1_0.V0, acc_4)
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
b_2_loop = gopurs_runtime.Apply2(f_0, b_2, (*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
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
__local_var_2_5 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
})
}), b_1)
_ = __local_var_2_5
var go__go_3_7_3 gopurs_runtime.Value
go__go_3_7_3 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_7_3:
for {
if false { continue go__go_3_7_3 }
var v_4 *Constructor_Cons[gopurs_runtime.Value] = v_4_loop
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
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
continue go__go_3_7_3
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t8))}
}
}()
})
})
__local_var_3_6 := gopurs_runtime.Apply(go__go_3_7_3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_3_6
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_5, gopurs_runtime.Apply(__local_var_3_6, x_4))
})
})
}))
	})
	return cache_foldableList
}

var cache_foldableNonEmptyList gopurs_runtime.Value
var once_foldableNonEmptyList sync.Once
func Get_foldableNonEmptyList() gopurs_runtime.Value {
	once_foldableNonEmptyList.Do(func() {
		cache_foldableNonEmptyList = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_foldableNonEmpty(), Get_foldableList())
	})
	return cache_foldableNonEmptyList
}

var cache_foldableWithIndexList gopurs_runtime.Value
var once_foldableWithIndexList sync.Once
func Get_foldableWithIndexList() gopurs_runtime.Value {
	once_foldableWithIndexList.Do(func() {
		cache_foldableWithIndexList = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableList()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableWithIndexList(), "foldlWithIndex"), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_2 := gopurs_runtime.Apply(Semigroup0_1_0.V0, acc_5)
_ = __local_var_6_2
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
__local_var_2_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_5 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_3_5
__local_var_4_6 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_4_6
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(__local_var_4_6.IntVal), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_4_6.IntVal), __local_var_3_5, a_5)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), acc_1})})
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(__local_var_2_4, x_3).UnsafePtr).V1
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_7 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, *Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_8 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_4_8
__local_var_5_9 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0
_ = __local_var_5_9
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(__local_var_5_9.IntVal), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_6, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__local_var_4_8)})}})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](xs_2))}))
_ = v_3_7
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_10 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_10
__local_var_6_11 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0
_ = __local_var_6_11
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(__local_var_6_11.IntVal), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(__local_var_6_11.IntVal), gopurs_runtime.Int(1)).IntVal), a_7, __local_var_5_10)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_7)}.UnsafePtr).V0.IntVal), b_1})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_7)}.UnsafePtr).V1))}).UnsafePtr).V1
})
})
}))
	})
	return cache_foldableWithIndexList
}

var cache_foldableWithIndexNonEmpty gopurs_runtime.Value
var once_foldableWithIndexNonEmpty sync.Once
func Get_foldableWithIndexNonEmpty() gopurs_runtime.Value {
	once_foldableWithIndexNonEmpty.Do(func() {
		cache_foldableWithIndexNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_foldableWithIndexNonEmpty(), Get_foldableWithIndexList())))}
	})
	return cache_foldableWithIndexNonEmpty
}

var cache_foldableWithIndexNonEmptyList gopurs_runtime.Value
var once_foldableWithIndexNonEmptyList sync.Once
func Get_foldableWithIndexNonEmptyList() gopurs_runtime.Value {
	once_foldableWithIndexNonEmptyList.Do(func() {
		cache_foldableWithIndexNonEmptyList = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableNonEmptyList()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply2(Get_maybe__919206801(), gopurs_runtime.Int(0), gopurs_runtime.Apply(Get_add__560788792(), gopurs_runtime.Int(1)))
_ = __local_var_3_0
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_foldableWithIndexNonEmpty()).V1, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_0))}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(__local_var_3_0, x_4))
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply2(Get_maybe__919206801(), gopurs_runtime.Int(0), gopurs_runtime.Apply(Get_add__560788792(), gopurs_runtime.Int(1)))
_ = __local_var_3_1
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_foldableWithIndexNonEmpty()).V2, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
}), b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_2 := gopurs_runtime.Apply2(Get_maybe__919206801(), gopurs_runtime.Int(0), gopurs_runtime.Apply(Get_add__560788792(), gopurs_runtime.Int(1)))
_ = __local_var_3_2
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_foldableWithIndexNonEmpty()).V3, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(__local_var_3_2, x_4))
}), b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2))})
})
})
}))
	})
	return cache_foldableWithIndexNonEmptyList
}

var cache_functorWithIndexList gopurs_runtime.Value
var once_functorWithIndexList sync.Once
func Get_functorWithIndexList() gopurs_runtime.Value {
	once_functorWithIndexList.Do(func() {
		cache_functorWithIndexList = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableWithIndexList(), "foldrWithIndex"), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, gopurs_runtime.Int(i_1.IntVal), x_2), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](acc_3)})}
})
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})
}))
	})
	return cache_functorWithIndexList
}

var cache_functorWithIndex gopurs_runtime.Value
var once_functorWithIndex sync.Once
func Get_functorWithIndex() gopurs_runtime.Value {
	once_functorWithIndex.Do(func() {
		cache_functorWithIndex = func() gopurs_runtime.Value {
__local_var_0_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_functorWithIndexList(), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_0_1
functorNonEmpty1_0_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_1, "map"), f_1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1)})}
})
}))
_ = functorNonEmpty1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return functorNonEmpty1_0_0
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorWithIndexList(), "mapWithIndex"), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, x_3})})
}), (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
})
}))))}
}()
	})
	return cache_functorWithIndex
}

var cache_functorWithIndexNonEmptyList gopurs_runtime.Value
var once_functorWithIndexNonEmptyList sync.Once
func Get_functorWithIndexNonEmptyList() gopurs_runtime.Value {
	once_functorWithIndexNonEmptyList.Do(func() {
		cache_functorWithIndexNonEmptyList = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
}), gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply2(Get_maybe__919206801(), gopurs_runtime.Int(0), gopurs_runtime.Apply(Get_add__560788792(), gopurs_runtime.Int(1)))
_ = __local_var_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](Call_mapWithIndex__598554346(gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(fn_0, gopurs_runtime.Apply(__local_var_2_0, x_3))
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1))})))}
})
}))
	})
	return cache_functorWithIndexNonEmptyList
}

var cache_semigroupList gopurs_runtime.Value
var once_semigroupList sync.Once
func Get_semigroupList() gopurs_runtime.Value {
	once_semigroupList.Do(func() {
		cache_semigroupList = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](ys_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](xs_0))})))}
})
}))
	})
	return cache_semigroupList
}

var cache_monoidList gopurs_runtime.Value
var once_monoidList sync.Once
func Get_monoidList() gopurs_runtime.Value {
	once_monoidList.Do(func() {
		cache_monoidList = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupList()
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})
	})
	return cache_monoidList
}

var cache_semigroupNonEmptyList gopurs_runtime.Value
var once_semigroupNonEmptyList sync.Once
func Get_semigroupNonEmptyList() gopurs_runtime.Value {
	once_semigroupNonEmptyList.Do(func() {
		cache_semigroupNonEmptyList = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(as_prime_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(as_prime_1.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(as_prime_1.UnsafePtr).V1)})})))}})}
})
}))
	})
	return cache_semigroupNonEmptyList
}

var cache_showList gopurs_runtime.Value
var once_showList sync.Once
func Get_showList() gopurs_runtime.Value {
	once_showList.Do(func() {
		cache_showList = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showList(dictShow_0_box)
})
	})
	return cache_showList
}

var cache_showNonEmptyList gopurs_runtime.Value
var once_showNonEmptyList sync.Once
func Get_showNonEmptyList() gopurs_runtime.Value {
	once_showNonEmptyList.Do(func() {
		cache_showNonEmptyList = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showNonEmptyList(dictShow_0_box)
})
	})
	return cache_showNonEmptyList
}

var cache_traversableList gopurs_runtime.Value
var once_traversableList sync.Once
func Get_traversableList() gopurs_runtime.Value {
	once_traversableList.Do(func() {
		cache_traversableList = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableList(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
Apply0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(Functor0_1_0.V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_5, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](b_4)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}))
_ = __local_var_4_2
__local_var_5_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_6_5 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(Apply0_2_1.V0, gopurs_runtime.Value{}))
_ = Functor0_6_5
__local_var_6_4 := gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_2_1.V1, gopurs_runtime.Apply2(Functor0_6_5.V0, gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_9, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](b_8)})}
})
}), acc_5), b_7)
})
_ = __local_var_6_4
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, gopurs_runtime.Apply(f_3, x_7))
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}))
_ = __local_var_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_5_3, x_6))
})
})
}))
	})
	return cache_traversableList
}

var cache_traversableNonEmptyList gopurs_runtime.Value
var once_traversableNonEmptyList sync.Once
func Get_traversableNonEmptyList() gopurs_runtime.Value {
	once_traversableNonEmptyList.Do(func() {
		cache_traversableNonEmptyList = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_traversableNonEmpty(), Get_traversableList())
	})
	return cache_traversableNonEmptyList
}

var cache_traversableWithIndexList gopurs_runtime.Value
var once_traversableWithIndexList sync.Once
func Get_traversableWithIndexList() gopurs_runtime.Value {
	once_traversableWithIndexList.Do(func() {
		cache_traversableWithIndexList = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableWithIndexList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorWithIndexList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_traversableList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
Apply0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(Functor0_1_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_5, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](b_4)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})))})
_ = __local_var_4_2
__local_var_5_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableWithIndexList(), "foldlWithIndex"), gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_6 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_7_5 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(Apply0_2_1.V0, gopurs_runtime.Value{}))
_ = Functor0_7_5
__local_var_7_4 := gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_2_1.V1, gopurs_runtime.Apply2(Functor0_7_5.V0, gopurs_runtime.Func(func(b_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_10, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](b_9)})}
})
}), acc_6), b_8)
})
_ = __local_var_7_4
__local_var_8_6 := gopurs_runtime.Apply(f_3, gopurs_runtime.Int(i_5.IntVal))
_ = __local_var_8_6
return gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_4, gopurs_runtime.Apply(__local_var_8_6, x_9))
})
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}))
_ = __local_var_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_5_3, x_6))
})
})
}))
	})
	return cache_traversableWithIndexList
}

var cache_traversableWithIndexNonEmpty gopurs_runtime.Value
var once_traversableWithIndexNonEmpty sync.Once
func Get_traversableWithIndexNonEmpty() gopurs_runtime.Value {
	once_traversableWithIndexNonEmpty.Do(func() {
		cache_traversableWithIndexNonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_TraversableWithIndex.Constructor_TraversableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_traversableWithIndexNonEmpty(), Get_traversableWithIndexList())))}
	})
	return cache_traversableWithIndexNonEmpty
}

var cache_traversableWithIndexNonEmptyList gopurs_runtime.Value
var once_traversableWithIndexNonEmptyList sync.Once
func Get_traversableWithIndexNonEmptyList() gopurs_runtime.Value {
	once_traversableWithIndexNonEmptyList.Do(func() {
		cache_traversableWithIndexNonEmptyList = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableWithIndexNonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorWithIndexNonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_traversableNonEmptyList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply2(Get_maybe__919206801(), gopurs_runtime.Int(0), gopurs_runtime.Apply(Get_add__560788792(), gopurs_runtime.Int(1)))
_ = __local_var_4_1
return gopurs_runtime.Apply2(Functor0_1_0.V0, Get_NonEmptyList(), gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_TraversableWithIndex.Constructor_TraversableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_traversableWithIndexNonEmpty()).V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(__local_var_4_1, x_5))
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_3))}))
})
})
}))
	})
	return cache_traversableWithIndexNonEmptyList
}

var cache_unfoldable1List gopurs_runtime.Value
var once_unfoldable1List sync.Once
func Get_unfoldable1List() gopurs_runtime.Value {
	once_unfoldable1List.Do(func() {
		cache_unfoldable1List = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_4 gopurs_runtime.Value
go__go_2_0_4 = gopurs_runtime.Func(func(source_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var source_3_loop gopurs_runtime.Value = source_3_loop_val
var memo_4_loop *Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](memo_4_loop_val)
go__go_2_0_4:
for {
if false { continue go__go_2_0_4 }
var source_3 gopurs_runtime.Value = source_3_loop
_ = source_3
var memo_4 *Constructor_Cons[gopurs_runtime.Value] = memo_4_loop
_ = memo_4
v_5_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]](gopurs_runtime.Apply(f_0, source_3))
_ = v_5_1
var __t4 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 930809136 && __t_tag_2.UnsafePtr != nil) {
source_3_loop = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V1.UnsafePtr).V0
memo_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V0, memo_4})})
continue go__go_2_0_4
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V1
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 930809136 && __t_tag_3.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_7, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](b_6)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V0, memo_4})})))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t4))}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_2_0_4, b_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})))}
})
}))
	})
	return cache_unfoldable1List
}

var cache_unfoldableList gopurs_runtime.Value
var once_unfoldableList sync.Once
func Get_unfoldableList() gopurs_runtime.Value {
	once_unfoldableList.Do(func() {
		cache_unfoldableList = gopurs_runtime.RecordDict2("Unfoldable10", "unfoldr", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_unfoldable1List()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_5 gopurs_runtime.Value
go__go_2_0_5 = gopurs_runtime.Func(func(source_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var source_3_loop gopurs_runtime.Value = source_3_loop_val
var memo_4_loop *Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](memo_4_loop_val)
go__go_2_0_5:
for {
if false { continue go__go_2_0_5 }
var source_3 gopurs_runtime.Value = source_3_loop
_ = source_3
var memo_4 *Constructor_Cons[gopurs_runtime.Value] = memo_4_loop
_ = memo_4
v_5_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(f_0, source_3))
_ = v_5_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_7, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](b_6)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(memo_4)})))}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr != nil) {
source_3_loop = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V0.UnsafePtr).V1
memo_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V0.UnsafePtr).V0, memo_4})})
continue go__go_2_0_5
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t2))}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_2_0_5, b_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})))}
})
}))
	})
	return cache_unfoldableList
}

var cache_unfoldable1NonEmptyList gopurs_runtime.Value
var once_unfoldable1NonEmptyList sync.Once
func Get_unfoldable1NonEmptyList() gopurs_runtime.Value {
	once_unfoldable1NonEmptyList.Do(func() {
		cache_unfoldable1NonEmptyList = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_unfoldable1NonEmpty(), Get_unfoldableList())
	})
	return cache_unfoldable1NonEmptyList
}

var cache_foldable1NonEmptyList gopurs_runtime.Value
var once_foldable1NonEmptyList sync.Once
func Get_foldable1NonEmptyList() gopurs_runtime.Value {
	once_foldable1NonEmptyList.Do(func() {
		cache_foldable1NonEmptyList = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_foldable1NonEmpty(), Get_foldableList())
	})
	return cache_foldable1NonEmptyList
}

var cache_extendNonEmptyList gopurs_runtime.Value
var once_extendNonEmptyList sync.Once
func Get_extendNonEmptyList() gopurs_runtime.Value {
	once_extendNonEmptyList.Do(func() {
		cache_extendNonEmptyList = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_2, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v1_3, "acc"))})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_2, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v1_3, "acc")))}})}), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v1_3, "val"))})})
})
}), gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1))}), "val")))}})}
})
}))
	})
	return cache_extendNonEmptyList
}

var cache_extendList gopurs_runtime.Value
var once_extendList sync.Once
func Get_extendList() gopurs_runtime.Value {
	once_extendList.Do(func() {
		cache_extendList = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 1358893437 && v1_1.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 1358893437 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v1_1))}), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), gopurs_runtime.Func(func(a_prime_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_prime_2, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v2_3, "acc"))})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_prime_2, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v2_3, "acc"))})}), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v2_3, "val"))})})
})
}), gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_1.UnsafePtr).V1)}), "val"))})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_extendList
}

var cache_eq1List gopurs_runtime.Value
var once_eq1List sync.Once
func Get_eq1List() gopurs_runtime.Value {
	once_eq1List.Do(func() {
		cache_eq1List = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
v2_6_loop = (Call_conj__3676519832(gopurs_runtime.Bool(v2_6), gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0).IntVal) != (0))).IntVal) != (0)
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
return gopurs_runtime.Bool((gopurs_runtime.Apply3(go__go_3_0_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](xs_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](ys_2))}, gopurs_runtime.Bool(true)).IntVal) != (0))
})
})
}))
	})
	return cache_eq1List
}

var cache_eq1NonEmptyList gopurs_runtime.Value
var once_eq1NonEmptyList sync.Once
func Get_eq1NonEmptyList() gopurs_runtime.Value {
	once_eq1NonEmptyList.Do(func() {
		cache_eq1NonEmptyList = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply2(pkg_Data_NonEmpty.Get_eqNonEmpty(), Get_eq1List(), dictEq_0), "eq")
}))
	})
	return cache_eq1NonEmptyList
}

var cache_eqList gopurs_runtime.Value
var once_eqList sync.Once
func Get_eqList() gopurs_runtime.Value {
	once_eqList.Do(func() {
		cache_eqList = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqList(dictEq_0_box)
})
	})
	return cache_eqList
}

var cache_eqNonEmptyList gopurs_runtime.Value
var once_eqNonEmptyList sync.Once
func Get_eqNonEmptyList() gopurs_runtime.Value {
	once_eqNonEmptyList.Do(func() {
		cache_eqNonEmptyList = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqNonEmptyList(dictEq_0_box)
})
	})
	return cache_eqNonEmptyList
}

var cache_ord1List gopurs_runtime.Value
var once_ord1List sync.Once
func Get_ord1List() gopurs_runtime.Value {
	once_ord1List.Do(func() {
		cache_ord1List = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1List()
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
v2_6_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0)
_ = v2_6_2
var __t3 uint32
{
if (uint32(v2_6_2.IntVal) == 902936544) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
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
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(go__go_3_0_7, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](xs_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](ys_2))}).IntVal)), UnsafePtr: nil}
})
})
}))
	})
	return cache_ord1List
}

var cache_ordNonEmpty gopurs_runtime.Value
var once_ordNonEmpty sync.Once
func Get_ordNonEmpty() gopurs_runtime.Value {
	once_ordNonEmpty.Do(func() {
		cache_ordNonEmpty = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_ordNonEmpty(), Get_ord1List())
	})
	return cache_ordNonEmpty
}

var cache_ord1NonEmptyList gopurs_runtime.Value
var once_ord1NonEmptyList sync.Once
func Get_ord1NonEmptyList() gopurs_runtime.Value {
	once_ord1NonEmptyList.Do(func() {
		cache_ord1NonEmptyList = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_ord1NonEmpty(), Get_ord1List())
	})
	return cache_ord1NonEmptyList
}

var cache_ordList gopurs_runtime.Value
var once_ordList sync.Once
func Get_ordList() gopurs_runtime.Value {
	once_ordList.Do(func() {
		cache_ordList = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordList(dictOrd_0_box)
})
	})
	return cache_ordList
}

var cache_ordNonEmptyList gopurs_runtime.Value
var once_ordNonEmptyList sync.Once
func Get_ordNonEmptyList() gopurs_runtime.Value {
	once_ordNonEmptyList.Do(func() {
		cache_ordNonEmptyList = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordNonEmptyList(dictOrd_0_box)
})
	})
	return cache_ordNonEmptyList
}

var cache_comonadNonEmptyList gopurs_runtime.Value
var once_comonadNonEmptyList sync.Once
func Get_comonadNonEmptyList() gopurs_runtime.Value {
	once_comonadNonEmptyList.Do(func() {
		cache_comonadNonEmptyList = gopurs_runtime.RecordDict2("Extend0", "extract", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_extendNonEmptyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0
}))
	})
	return cache_comonadNonEmptyList
}

var cache_applyList gopurs_runtime.Value
var once_applyList sync.Once
func Get_applyList() gopurs_runtime.Value {
	once_applyList.Do(func() {
		cache_applyList = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorList(), "map"), (*Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v1_1))})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v1_1))})))})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_applyList
}

var cache_applyNonEmptyList gopurs_runtime.Value
var once_applyNonEmptyList sync.Once
func Get_applyNonEmptyList() gopurs_runtime.Value {
	once_applyNonEmptyList.Do(func() {
		cache_applyNonEmptyList = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})})})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1))})))})))}})}
})
}))
	})
	return cache_applyNonEmptyList
}

var cache_bindList gopurs_runtime.Value
var once_bindList sync.Once
func Get_bindList() gopurs_runtime.Value {
	once_bindList.Do(func() {
		cache_bindList = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V0)))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V1)}, v1_1)))})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_bindList
}

var cache_bindNonEmptyList gopurs_runtime.Value
var once_bindNonEmptyList sync.Once
func Get_bindNonEmptyList() gopurs_runtime.Value {
	once_bindNonEmptyList.Do(func() {
		cache_bindNonEmptyList = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyNonEmptyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
v1_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0))
_ = v1_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v1_2_0)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v1_2_0)}.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1))}, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(f_1, x_3)
_ = __local_var_4_1
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1)})}
}))))})))}})}
})
}))
	})
	return cache_bindNonEmptyList
}

var cache_applicativeList gopurs_runtime.Value
var once_applicativeList sync.Once
func Get_applicativeList() gopurs_runtime.Value {
	once_applicativeList.Do(func() {
		cache_applicativeList = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyList()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_0, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})})}
}))
	})
	return cache_applicativeList
}

var cache_monadList gopurs_runtime.Value
var once_monadList sync.Once
func Get_monadList() gopurs_runtime.Value {
	once_monadList.Do(func() {
		cache_monadList = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindList()
}))
	})
	return cache_monadList
}

var cache_altNonEmptyList gopurs_runtime.Value
var once_altNonEmptyList sync.Once
func Get_altNonEmptyList() gopurs_runtime.Value {
	once_altNonEmptyList.Do(func() {
		cache_altNonEmptyList = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
}), gopurs_runtime.RecordGet(Get_semigroupNonEmptyList(), "append"))
	})
	return cache_altNonEmptyList
}

var cache_altList gopurs_runtime.Value
var once_altList sync.Once
func Get_altList() gopurs_runtime.Value {
	once_altList.Do(func() {
		cache_altList = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}), gopurs_runtime.RecordGet(Get_semigroupList(), "append"))
	})
	return cache_altList
}

var cache_plusList gopurs_runtime.Value
var once_plusList sync.Once
func Get_plusList() gopurs_runtime.Value {
	once_plusList.Do(func() {
		cache_plusList = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_altList()
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})
	})
	return cache_plusList
}

var cache_alternativeList gopurs_runtime.Value
var once_alternativeList sync.Once
func Get_alternativeList() gopurs_runtime.Value {
	once_alternativeList.Do(func() {
		cache_alternativeList = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_plusList()
}))
	})
	return cache_alternativeList
}

var cache_monadPlusList gopurs_runtime.Value
var once_monadPlusList sync.Once
func Get_monadPlusList() gopurs_runtime.Value {
	once_monadPlusList.Do(func() {
		cache_monadPlusList = gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_alternativeList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadList()
}))
	})
	return cache_monadPlusList
}

var cache_applicativeNonEmptyList gopurs_runtime.Value
var once_applicativeNonEmptyList sync.Once
func Get_applicativeNonEmptyList() gopurs_runtime.Value {
	once_applicativeNonEmptyList.Do(func() {
		cache_applicativeNonEmptyList = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyNonEmptyList()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_0, gopurs_runtime.RecordGet(Get_plusList(), "empty")})}
}))
	})
	return cache_applicativeNonEmptyList
}

var cache_monadNonEmptyList gopurs_runtime.Value
var once_monadNonEmptyList sync.Once
func Get_monadNonEmptyList() gopurs_runtime.Value {
	once_monadNonEmptyList.Do(func() {
		cache_monadNonEmptyList = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeNonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindNonEmptyList()
}))
	})
	return cache_monadNonEmptyList
}

var cache_traversable1NonEmptyList gopurs_runtime.Value
var once_traversable1NonEmptyList sync.Once
func Get_traversable1NonEmptyList() gopurs_runtime.Value {
	once_traversable1NonEmptyList.Do(func() {
		cache_traversable1NonEmptyList = gopurs_runtime.RecordDict4("Foldable10", "Traversable1", "sequence1", "traverse1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldable1NonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_traversableNonEmptyList()
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversable1NonEmptyList(), "traverse1"), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0))}, Get_identity1())
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V1)})}})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](Call_pure__575667894((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0)))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1))})))}
}), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_5_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_2
__local_var_5_1 := gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(Functor0_5_2.V0, gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_8, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_7.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_7.UnsafePtr).V1)})}})}
})
}), acc_4), b_6)
})
_ = __local_var_5_1
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_1, gopurs_runtime.Apply(f_2, x_6))
})
}), gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.RecordGet(Get_applicativeNonEmptyList(), "pure"), gopurs_runtime.Apply(f_2, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))}))
})
})
}))
	})
	return cache_traversable1NonEmptyList
}

var cache_pure__189931222 gopurs_runtime.Value
var once_pure__189931222 sync.Once
func Get_pure__189931222() gopurs_runtime.Value {
	once_pure__189931222.Do(func() {
		cache_pure__189931222 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__189931222(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__189931222
}

var cache_pure__2331489366 gopurs_runtime.Value
var once_pure__2331489366 sync.Once
func Get_pure__2331489366() gopurs_runtime.Value {
	once_pure__2331489366.Do(func() {
		cache_pure__2331489366 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__2331489366(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__2331489366
}

var cache_pure__993904534 gopurs_runtime.Value
var once_pure__993904534 sync.Once
func Get_pure__993904534() gopurs_runtime.Value {
	once_pure__993904534.Do(func() {
		cache_pure__993904534 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__993904534(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__993904534
}

var cache_pure__3215807376 gopurs_runtime.Value
var once_pure__3215807376 sync.Once
func Get_pure__3215807376() gopurs_runtime.Value {
	once_pure__3215807376.Do(func() {
		cache_pure__3215807376 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__3215807376(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__3215807376
}

var cache_pure__243192752 gopurs_runtime.Value
var once_pure__243192752 sync.Once
func Get_pure__243192752() gopurs_runtime.Value {
	once_pure__243192752.Do(func() {
		cache_pure__243192752 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__243192752(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__243192752
}

var cache_pure__575667894 gopurs_runtime.Value
var once_pure__575667894 sync.Once
func Get_pure__575667894() gopurs_runtime.Value {
	once_pure__575667894.Do(func() {
		cache_pure__575667894 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__575667894(__eta0_0_box)
})
	})
	return cache_pure__575667894
}

var cache_apply__1030762512 gopurs_runtime.Value
var once_apply__1030762512 sync.Once
func Get_apply__1030762512() gopurs_runtime.Value {
	once_apply__1030762512.Do(func() {
		cache_apply__1030762512 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__1030762512(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[*Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_apply__1030762512
}

var cache_apply__353515660 gopurs_runtime.Value
var once_apply__353515660 sync.Once
func Get_apply__353515660() gopurs_runtime.Value {
	once_apply__353515660.Do(func() {
		cache_apply__353515660 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__353515660(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__353515660
}

var cache_apply__2169384906 gopurs_runtime.Value
var once_apply__2169384906 sync.Once
func Get_apply__2169384906() gopurs_runtime.Value {
	once_apply__2169384906.Do(func() {
		cache_apply__2169384906 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_apply__2169384906(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v_0_box), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_apply__2169384906
}

var cache_lift2__1942544886 gopurs_runtime.Value
var once_lift2__1942544886 sync.Once
func Get_lift2__1942544886() gopurs_runtime.Value {
	once_lift2__1942544886.Do(func() {
		cache_lift2__1942544886 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__1942544886(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__1942544886
}

var cache_lift2__3139828374 gopurs_runtime.Value
var once_lift2__3139828374 sync.Once
func Get_lift2__3139828374() gopurs_runtime.Value {
	once_lift2__3139828374.Do(func() {
		cache_lift2__3139828374 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__3139828374(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__3139828374
}

var cache_lift2__3684551766 gopurs_runtime.Value
var once_lift2__3684551766 sync.Once
func Get_lift2__3684551766() gopurs_runtime.Value {
	once_lift2__3684551766.Do(func() {
		cache_lift2__3684551766 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__3684551766(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__3684551766
}

var cache_lift2__2286084406 gopurs_runtime.Value
var once_lift2__2286084406 sync.Once
func Get_lift2__2286084406() gopurs_runtime.Value {
	once_lift2__2286084406.Do(func() {
		cache_lift2__2286084406 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__2286084406(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__2286084406
}

var cache_lift2__3007543670 gopurs_runtime.Value
var once_lift2__3007543670 sync.Once
func Get_lift2__3007543670() gopurs_runtime.Value {
	once_lift2__3007543670.Do(func() {
		cache_lift2__3007543670 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__3007543670(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__3007543670
}

var cache_lift2__2762258480 gopurs_runtime.Value
var once_lift2__2762258480 sync.Once
func Get_lift2__2762258480() gopurs_runtime.Value {
	once_lift2__2762258480.Do(func() {
		cache_lift2__2762258480 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__2762258480(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__2762258480
}

var cache_lift2__2618178704 gopurs_runtime.Value
var once_lift2__2618178704 sync.Once
func Get_lift2__2618178704() gopurs_runtime.Value {
	once_lift2__2618178704.Do(func() {
		cache_lift2__2618178704 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__2618178704(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__2618178704
}

var cache_lift2__2114580400 gopurs_runtime.Value
var once_lift2__2114580400 sync.Once
func Get_lift2__2114580400() gopurs_runtime.Value {
	once_lift2__2114580400.Do(func() {
		cache_lift2__2114580400 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__2114580400(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__2114580400
}

var cache_lift2__3213187376 gopurs_runtime.Value
var once_lift2__3213187376 sync.Once
func Get_lift2__3213187376() gopurs_runtime.Value {
	once_lift2__3213187376.Do(func() {
		cache_lift2__3213187376 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__3213187376(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__3213187376
}

var cache_bind__3842572251 gopurs_runtime.Value
var once_bind__3842572251 sync.Once
func Get_bind__3842572251() gopurs_runtime.Value {
	once_bind__3842572251.Do(func() {
		cache_bind__3842572251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3842572251(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[*Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_bind__3842572251
}

var cache_bind__2601835655 gopurs_runtime.Value
var once_bind__2601835655 sync.Once
func Get_bind__2601835655() gopurs_runtime.Value {
	once_bind__2601835655.Do(func() {
		cache_bind__2601835655 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__2601835655(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__2601835655
}

var cache_bind__1872090113 gopurs_runtime.Value
var once_bind__1872090113 sync.Once
func Get_bind__1872090113() gopurs_runtime.Value {
	once_bind__1872090113.Do(func() {
		cache_bind__1872090113 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_bind__1872090113(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v_0_box), v1_1_box))}
})
	})
	return cache_bind__1872090113
}

var cache_compose__4141960292 gopurs_runtime.Value
var once_compose__4141960292 sync.Once
func Get_compose__4141960292() gopurs_runtime.Value {
	once_compose__4141960292.Do(func() {
		cache_compose__4141960292 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__4141960292(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[*Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_compose__4141960292
}

var cache_compose__1254722180 gopurs_runtime.Value
var once_compose__1254722180 sync.Once
func Get_compose__1254722180() gopurs_runtime.Value {
	once_compose__1254722180.Do(func() {
		cache_compose__1254722180 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__1254722180(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[*Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_compose__1254722180
}

var cache_compose__1555187646 gopurs_runtime.Value
var once_compose__1555187646 sync.Once
func Get_compose__1555187646() gopurs_runtime.Value {
	once_compose__1555187646.Do(func() {
		cache_compose__1555187646 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__1555187646(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__1555187646
}

var cache_compose__2527254334 gopurs_runtime.Value
var once_compose__2527254334 sync.Once
func Get_compose__2527254334() gopurs_runtime.Value {
	once_compose__2527254334.Do(func() {
		cache_compose__2527254334 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__2527254334(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__2527254334
}

var cache_compose__2532574046 gopurs_runtime.Value
var once_compose__2532574046 sync.Once
func Get_compose__2532574046() gopurs_runtime.Value {
	once_compose__2532574046.Do(func() {
		cache_compose__2532574046 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__2532574046(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__2532574046
}

var cache_compose__794534846 gopurs_runtime.Value
var once_compose__794534846 sync.Once
func Get_compose__794534846() gopurs_runtime.Value {
	once_compose__794534846.Do(func() {
		cache_compose__794534846 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__794534846(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__794534846
}

var cache_compose__4254807102 gopurs_runtime.Value
var once_compose__4254807102 sync.Once
func Get_compose__4254807102() gopurs_runtime.Value {
	once_compose__4254807102.Do(func() {
		cache_compose__4254807102 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__4254807102(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__4254807102
}

var cache_compose__2995688990 gopurs_runtime.Value
var once_compose__2995688990 sync.Once
func Get_compose__2995688990() gopurs_runtime.Value {
	once_compose__2995688990.Do(func() {
		cache_compose__2995688990 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__2995688990(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__2995688990
}

var cache_compose__3140790526 gopurs_runtime.Value
var once_compose__3140790526 sync.Once
func Get_compose__3140790526() gopurs_runtime.Value {
	once_compose__3140790526.Do(func() {
		cache_compose__3140790526 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__3140790526(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__3140790526
}

var cache_compose__3384557662 gopurs_runtime.Value
var once_compose__3384557662 sync.Once
func Get_compose__3384557662() gopurs_runtime.Value {
	once_compose__3384557662.Do(func() {
		cache_compose__3384557662 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__3384557662(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__3384557662
}

var cache_compose__2710321297 gopurs_runtime.Value
var once_compose__2710321297 sync.Once
func Get_compose__2710321297() gopurs_runtime.Value {
	once_compose__2710321297.Do(func() {
		cache_compose__2710321297 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__2710321297(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__2710321297
}

var cache_compose__1933206353 gopurs_runtime.Value
var once_compose__1933206353 sync.Once
func Get_compose__1933206353() gopurs_runtime.Value {
	once_compose__1933206353.Do(func() {
		cache_compose__1933206353 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__1933206353(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__1933206353
}

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_eq__2384498378 gopurs_runtime.Value
var once_eq__2384498378 sync.Once
func Get_eq__2384498378() gopurs_runtime.Value {
	once_eq__2384498378.Do(func() {
		cache_eq__2384498378 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__2384498378
}

var cache_foldl__3850309840 gopurs_runtime.Value
var once_foldl__3850309840 sync.Once
func Get_foldl__3850309840() gopurs_runtime.Value {
	once_foldl__3850309840.Do(func() {
		cache_foldl__3850309840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__3850309840(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_foldl__3850309840
}

var cache_foldl__2111289130 gopurs_runtime.Value
var once_foldl__2111289130 sync.Once
func Get_foldl__2111289130() gopurs_runtime.Value {
	once_foldl__2111289130.Do(func() {
		cache_foldl__2111289130 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__2111289130(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__2111289130
}

var cache_foldl__3041692656 gopurs_runtime.Value
var once_foldl__3041692656 sync.Once
func Get_foldl__3041692656() gopurs_runtime.Value {
	once_foldl__3041692656.Do(func() {
		cache_foldl__3041692656 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__3041692656(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_foldl__3041692656
}

var cache_foldl__66388714 gopurs_runtime.Value
var once_foldl__66388714 sync.Once
func Get_foldl__66388714() gopurs_runtime.Value {
	once_foldl__66388714.Do(func() {
		cache_foldl__66388714 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__66388714(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__66388714
}

var cache_foldl__1671904522 gopurs_runtime.Value
var once_foldl__1671904522 sync.Once
func Get_foldl__1671904522() gopurs_runtime.Value {
	once_foldl__1671904522.Do(func() {
		cache_foldl__1671904522 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__1671904522(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__1671904522
}

var cache_foldl__2602334544 gopurs_runtime.Value
var once_foldl__2602334544 sync.Once
func Get_foldl__2602334544() gopurs_runtime.Value {
	once_foldl__2602334544.Do(func() {
		cache_foldl__2602334544 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__2602334544(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_foldl__2602334544
}

var cache_foldl__371433392 gopurs_runtime.Value
var once_foldl__371433392 sync.Once
func Get_foldl__371433392() gopurs_runtime.Value {
	once_foldl__371433392.Do(func() {
		cache_foldl__371433392 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__371433392(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_foldl__371433392
}

var cache_foldl__3619558698 gopurs_runtime.Value
var once_foldl__3619558698 sync.Once
func Get_foldl__3619558698() gopurs_runtime.Value {
	once_foldl__3619558698.Do(func() {
		cache_foldl__3619558698 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__3619558698(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__3619558698
}

var cache_foldl__1601164432 gopurs_runtime.Value
var once_foldl__1601164432 sync.Once
func Get_foldl__1601164432() gopurs_runtime.Value {
	once_foldl__1601164432.Do(func() {
		cache_foldl__1601164432 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__1601164432(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_foldl__1601164432
}

var cache_foldl__2151204251 gopurs_runtime.Value
var once_foldl__2151204251 sync.Once
func Get_foldl__2151204251() gopurs_runtime.Value {
	once_foldl__2151204251.Do(func() {
		cache_foldl__2151204251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__2151204251(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__2151204251
}

var cache_foldl__1712912315 gopurs_runtime.Value
var once_foldl__1712912315 sync.Once
func Get_foldl__1712912315() gopurs_runtime.Value {
	once_foldl__1712912315.Do(func() {
		cache_foldl__1712912315 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__1712912315(f_0_box)
})
	})
	return cache_foldl__1712912315
}

var cache_foldl__2159564571 gopurs_runtime.Value
var once_foldl__2159564571 sync.Once
func Get_foldl__2159564571() gopurs_runtime.Value {
	once_foldl__2159564571.Do(func() {
		cache_foldl__2159564571 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__2159564571(f_0_box)
})
	})
	return cache_foldl__2159564571
}

var cache_foldl__3785384859 gopurs_runtime.Value
var once_foldl__3785384859 sync.Once
func Get_foldl__3785384859() gopurs_runtime.Value {
	once_foldl__3785384859.Do(func() {
		cache_foldl__3785384859 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__3785384859(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__3785384859
}

var cache_foldl__3943124669 gopurs_runtime.Value
var once_foldl__3943124669 sync.Once
func Get_foldl__3943124669() gopurs_runtime.Value {
	once_foldl__3943124669.Do(func() {
		cache_foldl__3943124669 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__3943124669(f_0_box)
})
	})
	return cache_foldl__3943124669
}

var cache_foldl__396932925 gopurs_runtime.Value
var once_foldl__396932925 sync.Once
func Get_foldl__396932925() gopurs_runtime.Value {
	once_foldl__396932925.Do(func() {
		cache_foldl__396932925 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__396932925(f_0_box)
})
	})
	return cache_foldl__396932925
}

var cache_foldl__2928402749 gopurs_runtime.Value
var once_foldl__2928402749 sync.Once
func Get_foldl__2928402749() gopurs_runtime.Value {
	once_foldl__2928402749.Do(func() {
		cache_foldl__2928402749 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__2928402749(f_0_box)
})
	})
	return cache_foldl__2928402749
}

var cache_foldl__255626813 gopurs_runtime.Value
var once_foldl__255626813 sync.Once
func Get_foldl__255626813() gopurs_runtime.Value {
	once_foldl__255626813.Do(func() {
		cache_foldl__255626813 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__255626813(f_0_box)
})
	})
	return cache_foldl__255626813
}

var cache_foldl__3915700701 gopurs_runtime.Value
var once_foldl__3915700701 sync.Once
func Get_foldl__3915700701() gopurs_runtime.Value {
	once_foldl__3915700701.Do(func() {
		cache_foldl__3915700701 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__3915700701(f_0_box)
})
	})
	return cache_foldl__3915700701
}

var cache_foldl__3459294429 gopurs_runtime.Value
var once_foldl__3459294429 sync.Once
func Get_foldl__3459294429() gopurs_runtime.Value {
	once_foldl__3459294429.Do(func() {
		cache_foldl__3459294429 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__3459294429(f_0_box)
})
	})
	return cache_foldl__3459294429
}

var cache_foldr__2512763050 gopurs_runtime.Value
var once_foldr__2512763050 sync.Once
func Get_foldr__2512763050() gopurs_runtime.Value {
	once_foldr__2512763050.Do(func() {
		cache_foldr__2512763050 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2512763050(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__2512763050
}

var cache_foldr__3673994608 gopurs_runtime.Value
var once_foldr__3673994608 sync.Once
func Get_foldr__3673994608() gopurs_runtime.Value {
	once_foldr__3673994608.Do(func() {
		cache_foldr__3673994608 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__3673994608(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_foldr__3673994608
}

var cache_foldr__2151204251 gopurs_runtime.Value
var once_foldr__2151204251 sync.Once
func Get_foldr__2151204251() gopurs_runtime.Value {
	once_foldr__2151204251.Do(func() {
		cache_foldr__2151204251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2151204251(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__2151204251
}

var cache_foldr__3943124669 gopurs_runtime.Value
var once_foldr__3943124669 sync.Once
func Get_foldr__3943124669() gopurs_runtime.Value {
	once_foldr__3943124669.Do(func() {
		cache_foldr__3943124669 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__3943124669(f_0_box, b_1_box)
})
	})
	return cache_foldr__3943124669
}

var cache_foldr__2979608669 gopurs_runtime.Value
var once_foldr__2979608669 sync.Once
func Get_foldr__2979608669() gopurs_runtime.Value {
	once_foldr__2979608669.Do(func() {
		cache_foldr__2979608669 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2979608669(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](b_1_box))
})
	})
	return cache_foldr__2979608669
}

var cache_foldr__4137485405 gopurs_runtime.Value
var once_foldr__4137485405 sync.Once
func Get_foldr__4137485405() gopurs_runtime.Value {
	once_foldr__4137485405.Do(func() {
		cache_foldr__4137485405 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__4137485405(f_0_box, b_1_box)
})
	})
	return cache_foldr__4137485405
}

var cache_intercalate__3939234276 gopurs_runtime.Value
var once_intercalate__3939234276 sync.Once
func Get_intercalate__3939234276() gopurs_runtime.Value {
	once_intercalate__3939234276.Do(func() {
		cache_intercalate__3939234276 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intercalate__3939234276(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_1_box))
})
	})
	return cache_intercalate__3939234276
}

var cache_intercalate__2937349250 gopurs_runtime.Value
var once_intercalate__2937349250 sync.Once
func Get_intercalate__2937349250() gopurs_runtime.Value {
	once_intercalate__2937349250.Do(func() {
		cache_intercalate__2937349250 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intercalate__2937349250(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_intercalate__2937349250
}

var cache_foldMapWithIndex__2292551140 gopurs_runtime.Value
var once_foldMapWithIndex__2292551140 sync.Once
func Get_foldMapWithIndex__2292551140() gopurs_runtime.Value {
	once_foldMapWithIndex__2292551140.Do(func() {
		cache_foldMapWithIndex__2292551140 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMapWithIndex__2292551140(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMapWithIndex__2292551140
}

var cache_foldMapWithIndex__1722031522 gopurs_runtime.Value
var once_foldMapWithIndex__1722031522 sync.Once
func Get_foldMapWithIndex__1722031522() gopurs_runtime.Value {
	once_foldMapWithIndex__1722031522.Do(func() {
		cache_foldMapWithIndex__1722031522 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMapWithIndex__1722031522(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_foldMapWithIndex__1722031522
}

var cache_foldMapWithIndex__2880267906 gopurs_runtime.Value
var once_foldMapWithIndex__2880267906 sync.Once
func Get_foldMapWithIndex__2880267906() gopurs_runtime.Value {
	once_foldMapWithIndex__2880267906.Do(func() {
		cache_foldMapWithIndex__2880267906 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMapWithIndex__2880267906(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_foldMapWithIndex__2880267906
}

var cache_foldlWithIndex__2808220203 gopurs_runtime.Value
var once_foldlWithIndex__2808220203 sync.Once
func Get_foldlWithIndex__2808220203() gopurs_runtime.Value {
	once_foldlWithIndex__2808220203.Do(func() {
		cache_foldlWithIndex__2808220203 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, acc_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlWithIndex__2808220203(f_0_box, acc_1_box)
})
	})
	return cache_foldlWithIndex__2808220203
}

var cache_foldlWithIndex__2764250251 gopurs_runtime.Value
var once_foldlWithIndex__2764250251 sync.Once
func Get_foldlWithIndex__2764250251() gopurs_runtime.Value {
	once_foldlWithIndex__2764250251.Do(func() {
		cache_foldlWithIndex__2764250251 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, acc_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlWithIndex__2764250251(f_0_box, acc_1_box)
})
	})
	return cache_foldlWithIndex__2764250251
}

var cache_foldlWithIndex__2986161357 gopurs_runtime.Value
var once_foldlWithIndex__2986161357 sync.Once
func Get_foldlWithIndex__2986161357() gopurs_runtime.Value {
	once_foldlWithIndex__2986161357.Do(func() {
		cache_foldlWithIndex__2986161357 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlWithIndex__2986161357(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldlWithIndex__2986161357
}

var cache_foldlWithIndex__3610348555 gopurs_runtime.Value
var once_foldlWithIndex__3610348555 sync.Once
func Get_foldlWithIndex__3610348555() gopurs_runtime.Value {
	once_foldlWithIndex__3610348555.Do(func() {
		cache_foldlWithIndex__3610348555 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlWithIndex__3610348555(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_foldlWithIndex__3610348555
}

var cache_foldlWithIndex__1651851147 gopurs_runtime.Value
var once_foldlWithIndex__1651851147 sync.Once
func Get_foldlWithIndex__1651851147() gopurs_runtime.Value {
	once_foldlWithIndex__1651851147.Do(func() {
		cache_foldlWithIndex__1651851147 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlWithIndex__1651851147(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_foldlWithIndex__1651851147
}

var cache_foldrWithIndex__2808220203 gopurs_runtime.Value
var once_foldrWithIndex__2808220203 sync.Once
func Get_foldrWithIndex__2808220203() gopurs_runtime.Value {
	once_foldrWithIndex__2808220203.Do(func() {
		cache_foldrWithIndex__2808220203 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrWithIndex__2808220203(f_0_box, b_1_box, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](xs_2_box))
})
	})
	return cache_foldrWithIndex__2808220203
}

var cache_foldrWithIndex__2439396107 gopurs_runtime.Value
var once_foldrWithIndex__2439396107 sync.Once
func Get_foldrWithIndex__2439396107() gopurs_runtime.Value {
	once_foldrWithIndex__2439396107.Do(func() {
		cache_foldrWithIndex__2439396107 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_foldrWithIndex__2439396107(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](b_1_box), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](xs_2_box)))}
})
	})
	return cache_foldrWithIndex__2439396107
}

var cache_foldrWithIndex__2986161357 gopurs_runtime.Value
var once_foldrWithIndex__2986161357 sync.Once
func Get_foldrWithIndex__2986161357() gopurs_runtime.Value {
	once_foldrWithIndex__2986161357.Do(func() {
		cache_foldrWithIndex__2986161357 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrWithIndex__2986161357(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldrWithIndex__2986161357
}

var cache_foldrWithIndex__3610348555 gopurs_runtime.Value
var once_foldrWithIndex__3610348555 sync.Once
func Get_foldrWithIndex__3610348555() gopurs_runtime.Value {
	once_foldrWithIndex__3610348555.Do(func() {
		cache_foldrWithIndex__3610348555 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrWithIndex__3610348555(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_foldrWithIndex__3610348555
}

var cache_foldrWithIndex__979136683 gopurs_runtime.Value
var once_foldrWithIndex__979136683 sync.Once
func Get_foldrWithIndex__979136683() gopurs_runtime.Value {
	once_foldrWithIndex__979136683.Do(func() {
		cache_foldrWithIndex__979136683 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrWithIndex__979136683(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_foldrWithIndex__979136683
}

var cache_flip__3525081280 gopurs_runtime.Value
var once_flip__3525081280 sync.Once
func Get_flip__3525081280() gopurs_runtime.Value {
	once_flip__3525081280.Do(func() {
		cache_flip__3525081280 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3525081280(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3525081280
}

var cache_flip__2974723072 gopurs_runtime.Value
var once_flip__2974723072 sync.Once
func Get_flip__2974723072() gopurs_runtime.Value {
	once_flip__2974723072.Do(func() {
		cache_flip__2974723072 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__2974723072(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__2974723072
}

var cache_flip__3709724320 gopurs_runtime.Value
var once_flip__3709724320 sync.Once
func Get_flip__3709724320() gopurs_runtime.Value {
	once_flip__3709724320.Do(func() {
		cache_flip__3709724320 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3709724320(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3709724320
}

var cache_flip__3563101792 gopurs_runtime.Value
var once_flip__3563101792 sync.Once
func Get_flip__3563101792() gopurs_runtime.Value {
	once_flip__3563101792.Do(func() {
		cache_flip__3563101792 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3563101792(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3563101792
}

var cache_flip__1833071808 gopurs_runtime.Value
var once_flip__1833071808 sync.Once
func Get_flip__1833071808() gopurs_runtime.Value {
	once_flip__1833071808.Do(func() {
		cache_flip__1833071808 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__1833071808(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__1833071808
}

var cache_flip__1673583840 gopurs_runtime.Value
var once_flip__1673583840 sync.Once
func Get_flip__1673583840() gopurs_runtime.Value {
	once_flip__1673583840.Do(func() {
		cache_flip__1673583840 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__1673583840(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__1673583840
}

var cache_flip__534748448 gopurs_runtime.Value
var once_flip__534748448 sync.Once
func Get_flip__534748448() gopurs_runtime.Value {
	once_flip__534748448.Do(func() {
		cache_flip__534748448 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__534748448(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__534748448
}

var cache_flip__1744188480 gopurs_runtime.Value
var once_flip__1744188480 sync.Once
func Get_flip__1744188480() gopurs_runtime.Value {
	once_flip__1744188480.Do(func() {
		cache_flip__1744188480 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__1744188480(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__1744188480
}

var cache_flip__3468792800 gopurs_runtime.Value
var once_flip__3468792800 sync.Once
func Get_flip__3468792800() gopurs_runtime.Value {
	once_flip__3468792800.Do(func() {
		cache_flip__3468792800 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3468792800(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3468792800
}

var cache_flip__872296768 gopurs_runtime.Value
var once_flip__872296768 sync.Once
func Get_flip__872296768() gopurs_runtime.Value {
	once_flip__872296768.Do(func() {
		cache_flip__872296768 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__872296768(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__872296768
}

var cache_flip__4091748192 gopurs_runtime.Value
var once_flip__4091748192 sync.Once
func Get_flip__4091748192() gopurs_runtime.Value {
	once_flip__4091748192.Do(func() {
		cache_flip__4091748192 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__4091748192(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__4091748192
}

var cache_map__3116241637 gopurs_runtime.Value
var once_map__3116241637 sync.Once
func Get_map__3116241637() gopurs_runtime.Value {
	once_map__3116241637.Do(func() {
		cache_map__3116241637 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3116241637(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[*Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_map__3116241637
}

var cache_map__2665381605 gopurs_runtime.Value
var once_map__2665381605 sync.Once
func Get_map__2665381605() gopurs_runtime.Value {
	once_map__2665381605.Do(func() {
		cache_map__2665381605 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2665381605(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2665381605
}

var cache_map__67411525 gopurs_runtime.Value
var once_map__67411525 sync.Once
func Get_map__67411525() gopurs_runtime.Value {
	once_map__67411525.Do(func() {
		cache_map__67411525 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__67411525(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__67411525
}

var cache_map__2876470885 gopurs_runtime.Value
var once_map__2876470885 sync.Once
func Get_map__2876470885() gopurs_runtime.Value {
	once_map__2876470885.Do(func() {
		cache_map__2876470885 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2876470885(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2876470885
}

var cache_map__2199395572 gopurs_runtime.Value
var once_map__2199395572 sync.Once
func Get_map__2199395572() gopurs_runtime.Value {
	once_map__2199395572.Do(func() {
		cache_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2199395572(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2199395572
}

var cache_map__1256368628 gopurs_runtime.Value
var once_map__1256368628 sync.Once
func Get_map__1256368628() gopurs_runtime.Value {
	once_map__1256368628.Do(func() {
		cache_map__1256368628 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1256368628(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1256368628
}

var cache_map__1483545076 gopurs_runtime.Value
var once_map__1483545076 sync.Once
func Get_map__1483545076() gopurs_runtime.Value {
	once_map__1483545076.Do(func() {
		cache_map__1483545076 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1483545076(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1483545076
}

var cache_map__3058795348 gopurs_runtime.Value
var once_map__3058795348 sync.Once
func Get_map__3058795348() gopurs_runtime.Value {
	once_map__3058795348.Do(func() {
		cache_map__3058795348 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3058795348(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__3058795348
}

var cache_map__2190988916 gopurs_runtime.Value
var once_map__2190988916 sync.Once
func Get_map__2190988916() gopurs_runtime.Value {
	once_map__2190988916.Do(func() {
		cache_map__2190988916 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2190988916(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2190988916
}

var cache_map__843173928 gopurs_runtime.Value
var once_map__843173928 sync.Once
func Get_map__843173928() gopurs_runtime.Value {
	once_map__843173928.Do(func() {
		cache_map__843173928 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__843173928(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_map__843173928
}

var cache_map__438443400 gopurs_runtime.Value
var once_map__438443400 sync.Once
func Get_map__438443400() gopurs_runtime.Value {
	once_map__438443400.Do(func() {
		cache_map__438443400 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__438443400(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_map__438443400
}

var cache_mapFlipped__2466386789 gopurs_runtime.Value
var once_mapFlipped__2466386789 sync.Once
func Get_mapFlipped__2466386789() gopurs_runtime.Value {
	once_mapFlipped__2466386789.Do(func() {
		cache_mapFlipped__2466386789 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapFlipped__2466386789(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_mapFlipped__2466386789
}

var cache_mapFlipped__4215217780 gopurs_runtime.Value
var once_mapFlipped__4215217780 sync.Once
func Get_mapFlipped__4215217780() gopurs_runtime.Value {
	once_mapFlipped__4215217780.Do(func() {
		cache_mapFlipped__4215217780 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapFlipped__4215217780(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_mapFlipped__4215217780
}

var cache_mapFlipped__2919806324 gopurs_runtime.Value
var once_mapFlipped__2919806324 sync.Once
func Get_mapFlipped__2919806324() gopurs_runtime.Value {
	once_mapFlipped__2919806324.Do(func() {
		cache_mapFlipped__2919806324 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapFlipped__2919806324(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), fa_1_box, f_2_box)
})
	})
	return cache_mapFlipped__2919806324
}

var cache_mapWithIndex__55256674 gopurs_runtime.Value
var once_mapWithIndex__55256674 sync.Once
func Get_mapWithIndex__55256674() gopurs_runtime.Value {
	once_mapWithIndex__55256674.Do(func() {
		cache_mapWithIndex__55256674 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapWithIndex__55256674(gopurs_runtime.CoerceToStruct[pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mapWithIndex__55256674
}

var cache_mapWithIndex__574674314 gopurs_runtime.Value
var once_mapWithIndex__574674314 sync.Once
func Get_mapWithIndex__574674314() gopurs_runtime.Value {
	once_mapWithIndex__574674314.Do(func() {
		cache_mapWithIndex__574674314 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapWithIndex__574674314(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_mapWithIndex__574674314
}

var cache_mapWithIndex__598554346 gopurs_runtime.Value
var once_mapWithIndex__598554346 sync.Once
func Get_mapWithIndex__598554346() gopurs_runtime.Value {
	once_mapWithIndex__598554346.Do(func() {
		cache_mapWithIndex__598554346 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapWithIndex__598554346(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_mapWithIndex__598554346
}

var cache_conj__3676519832 gopurs_runtime.Value
var once_conj__3676519832 sync.Once
func Get_conj__3676519832() gopurs_runtime.Value {
	once_conj__3676519832.Do(func() {
		cache_conj__3676519832 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__3676519832(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_conj__3676519832
}

var cache_conj__3472268504 gopurs_runtime.Value
var once_conj__3472268504 sync.Once
func Get_conj__3472268504() gopurs_runtime.Value {
	once_conj__3472268504.Do(func() {
		cache_conj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_conj__3472268504
}

var cache_disj__3676519832 gopurs_runtime.Value
var once_disj__3676519832 sync.Once
func Get_disj__3676519832() gopurs_runtime.Value {
	once_disj__3676519832.Do(func() {
		cache_disj__3676519832 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3676519832(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_disj__3676519832
}

var cache_disj__3472268504 gopurs_runtime.Value
var once_disj__3472268504 sync.Once
func Get_disj__3472268504() gopurs_runtime.Value {
	once_disj__3472268504.Do(func() {
		cache_disj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_disj__3472268504
}

var cache_not__3201284355 gopurs_runtime.Value
var once_not__3201284355 sync.Once
func Get_not__3201284355() gopurs_runtime.Value {
	once_not__3201284355.Do(func() {
		cache_not__3201284355 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__3201284355(__eta0_0_box)
})
	})
	return cache_not__3201284355
}

var cache_not__1505204753 gopurs_runtime.Value
var once_not__1505204753 sync.Once
func Get_not__1505204753() gopurs_runtime.Value {
	once_not__1505204753.Do(func() {
		cache_not__1505204753 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__1505204753(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_not__1505204753
}

var cache_altList__614667287 gopurs_runtime.Value
var once_altList__614667287 sync.Once
func Get_altList__614667287() gopurs_runtime.Value {
	once_altList__614667287.Do(func() {
		cache_altList__614667287 = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}), gopurs_runtime.RecordGet(Get_semigroupList(), "append"))
	})
	return cache_altList__614667287
}

var cache_alternativeList__4245871901 gopurs_runtime.Value
var once_alternativeList__4245871901 sync.Once
func Get_alternativeList__4245871901() gopurs_runtime.Value {
	once_alternativeList__4245871901.Do(func() {
		cache_alternativeList__4245871901 = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_plusList()
}))
	})
	return cache_alternativeList__4245871901
}

var cache_applicativeList__2027879016 gopurs_runtime.Value
var once_applicativeList__2027879016 sync.Once
func Get_applicativeList__2027879016() gopurs_runtime.Value {
	once_applicativeList__2027879016.Do(func() {
		cache_applicativeList__2027879016 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyList()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_0, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})})}
}))
	})
	return cache_applicativeList__2027879016
}

var cache_applicativeNonEmptyList__1156428081 gopurs_runtime.Value
var once_applicativeNonEmptyList__1156428081 sync.Once
func Get_applicativeNonEmptyList__1156428081() gopurs_runtime.Value {
	once_applicativeNonEmptyList__1156428081.Do(func() {
		cache_applicativeNonEmptyList__1156428081 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyNonEmptyList()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_0, gopurs_runtime.RecordGet(Get_plusList(), "empty")})}
}))
	})
	return cache_applicativeNonEmptyList__1156428081
}

var cache_applicativeNonEmptyList__3820246605 gopurs_runtime.Value
var once_applicativeNonEmptyList__3820246605 sync.Once
func Get_applicativeNonEmptyList__3820246605() gopurs_runtime.Value {
	once_applicativeNonEmptyList__3820246605.Do(func() {
		cache_applicativeNonEmptyList__3820246605 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyNonEmptyList()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_0, gopurs_runtime.RecordGet(Get_plusList(), "empty")})}
}))
	})
	return cache_applicativeNonEmptyList__3820246605
}

var cache_applicativeNonEmptyList__233027336 gopurs_runtime.Value
var once_applicativeNonEmptyList__233027336 sync.Once
func Get_applicativeNonEmptyList__233027336() gopurs_runtime.Value {
	once_applicativeNonEmptyList__233027336.Do(func() {
		cache_applicativeNonEmptyList__233027336 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyNonEmptyList()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_0, gopurs_runtime.RecordGet(Get_plusList(), "empty")})}
}))
	})
	return cache_applicativeNonEmptyList__233027336
}

var cache_applyList__3072763993 gopurs_runtime.Value
var once_applyList__3072763993 sync.Once
func Get_applyList__3072763993() gopurs_runtime.Value {
	once_applyList__3072763993.Do(func() {
		cache_applyList__3072763993 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorList(), "map"), (*Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v1_1))})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v1_1))})))})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_applyList__3072763993
}

var cache_applyList__1109325167 gopurs_runtime.Value
var once_applyList__1109325167 sync.Once
func Get_applyList__1109325167() gopurs_runtime.Value {
	once_applyList__1109325167.Do(func() {
		cache_applyList__1109325167 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorList(), "map"), (*Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v1_1))})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v1_1))})))})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_applyList__1109325167
}

var cache_applyNonEmptyList__602103086 gopurs_runtime.Value
var once_applyNonEmptyList__602103086 sync.Once
func Get_applyNonEmptyList__602103086() gopurs_runtime.Value {
	once_applyNonEmptyList__602103086.Do(func() {
		cache_applyNonEmptyList__602103086 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})})})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1))})))})))}})}
})
}))
	})
	return cache_applyNonEmptyList__602103086
}

var cache_bindList__241263065 gopurs_runtime.Value
var once_bindList__241263065 sync.Once
func Get_bindList__241263065() gopurs_runtime.Value {
	once_bindList__241263065.Do(func() {
		cache_bindList__241263065 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V0)))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V1)}, v1_1)))})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_bindList__241263065
}

var cache_bindList__3903132455 gopurs_runtime.Value
var once_bindList__3903132455 sync.Once
func Get_bindList__3903132455() gopurs_runtime.Value {
	once_bindList__3903132455.Do(func() {
		cache_bindList__3903132455 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V0)))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V1)}, v1_1)))})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_bindList__3903132455
}

var cache_bindNonEmptyList__1309464679 gopurs_runtime.Value
var once_bindNonEmptyList__1309464679 sync.Once
func Get_bindNonEmptyList__1309464679() gopurs_runtime.Value {
	once_bindNonEmptyList__1309464679.Do(func() {
		cache_bindNonEmptyList__1309464679 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyNonEmptyList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
v1_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0))
_ = v1_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v1_2_0)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v1_2_0)}.UnsafePtr).V1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1))}, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(f_1, x_3)
_ = __local_var_4_1
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1)})}
}))))})))}})}
})
}))
	})
	return cache_bindNonEmptyList__1309464679
}

var cache_eq1List__1109645470 gopurs_runtime.Value
var once_eq1List__1109645470 sync.Once
func Get_eq1List__1109645470() gopurs_runtime.Value {
	once_eq1List__1109645470.Do(func() {
		cache_eq1List__1109645470 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_0_19 gopurs_runtime.Value
go__go_3_0_19 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
var v2_6_loop bool = (v2_6_loop_val.IntVal) != (0)
go__go_3_0_19:
for {
if false { continue go__go_3_0_19 }
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
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
v2_6_loop = (v2_6) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0).IntVal) != (0))
continue go__go_3_0_19
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
return gopurs_runtime.Bool((gopurs_runtime.Apply3(go__go_3_0_19, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](xs_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](ys_2))}, gopurs_runtime.Bool(true)).IntVal) != (0))
})
})
}))
	})
	return cache_eq1List__1109645470
}

var cache_extendNonEmptyList__2163561628 gopurs_runtime.Value
var once_extendNonEmptyList__2163561628 sync.Once
func Get_extendNonEmptyList__2163561628() gopurs_runtime.Value {
	once_extendNonEmptyList__2163561628.Do(func() {
		cache_extendNonEmptyList__2163561628 = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_2, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v1_3, "acc"))})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_2, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v1_3, "acc")))}})}), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v1_3, "val"))})})
})
}), gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1))}), "val")))}})}
})
}))
	})
	return cache_extendNonEmptyList__2163561628
}

var cache_foldable1NonEmptyList__1746670655 gopurs_runtime.Value
var once_foldable1NonEmptyList__1746670655 sync.Once
func Get_foldable1NonEmptyList__1746670655() gopurs_runtime.Value {
	once_foldable1NonEmptyList__1746670655.Do(func() {
		cache_foldable1NonEmptyList__1746670655 = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_foldable1NonEmpty(), Get_foldableList())
	})
	return cache_foldable1NonEmptyList__1746670655
}

var cache_foldableList__1753400174 gopurs_runtime.Value
var once_foldableList__1753400174 sync.Once
func Get_foldableList__1753400174() gopurs_runtime.Value {
	once_foldableList__1753400174.Do(func() {
		cache_foldableList__1753400174 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(Semigroup0_1_0.V0, acc_4)
_ = __local_var_5_2
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, gopurs_runtime.Apply(f_3, x_6))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_3_20 gopurs_runtime.Value
go__go_1_3_20 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_3_20:
for {
if false { continue go__go_1_3_20 }
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
b_2_loop = gopurs_runtime.Apply2(f_0, b_2, (*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
continue go__go_1_3_20
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
return go__go_1_3_20
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_5 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
})
}), b_1)
_ = __local_var_2_5
var go__go_3_7_21 gopurs_runtime.Value
go__go_3_7_21 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_7_21:
for {
if false { continue go__go_3_7_21 }
var v_4 *Constructor_Cons[gopurs_runtime.Value] = v_4_loop
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
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
continue go__go_3_7_21
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t8))}
}
}()
})
})
__local_var_3_6 := gopurs_runtime.Apply(go__go_3_7_21, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_3_6
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_5, gopurs_runtime.Apply(__local_var_3_6, x_4))
})
})
}))
	})
	return cache_foldableList__1753400174
}

var cache_foldableList__46095397 gopurs_runtime.Value
var once_foldableList__46095397 sync.Once
func Get_foldableList__46095397() gopurs_runtime.Value {
	once_foldableList__46095397.Do(func() {
		cache_foldableList__46095397 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(Semigroup0_1_0.V0, acc_4)
_ = __local_var_5_2
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, gopurs_runtime.Apply(f_3, x_6))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_3_22 gopurs_runtime.Value
go__go_1_3_22 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_3_22:
for {
if false { continue go__go_1_3_22 }
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
b_2_loop = gopurs_runtime.Apply2(f_0, b_2, (*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
continue go__go_1_3_22
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
return go__go_1_3_22
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_5 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
})
}), b_1)
_ = __local_var_2_5
var go__go_3_7_23 gopurs_runtime.Value
go__go_3_7_23 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_7_23:
for {
if false { continue go__go_3_7_23 }
var v_4 *Constructor_Cons[gopurs_runtime.Value] = v_4_loop
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
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
continue go__go_3_7_23
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t8))}
}
}()
})
})
__local_var_3_6 := gopurs_runtime.Apply(go__go_3_7_23, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_3_6
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_5, gopurs_runtime.Apply(__local_var_3_6, x_4))
})
})
}))
	})
	return cache_foldableList__46095397
}

var cache_foldableNonEmptyList__3933232868 gopurs_runtime.Value
var once_foldableNonEmptyList__3933232868 sync.Once
func Get_foldableNonEmptyList__3933232868() gopurs_runtime.Value {
	once_foldableNonEmptyList__3933232868.Do(func() {
		cache_foldableNonEmptyList__3933232868 = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_foldableNonEmpty(), Get_foldableList())
	})
	return cache_foldableNonEmptyList__3933232868
}

var cache_foldableWithIndexList__3168570882 gopurs_runtime.Value
var once_foldableWithIndexList__3168570882 sync.Once
func Get_foldableWithIndexList__3168570882() gopurs_runtime.Value {
	once_foldableWithIndexList__3168570882.Do(func() {
		cache_foldableWithIndexList__3168570882 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableList()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableWithIndexList(), "foldlWithIndex"), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_2 := gopurs_runtime.Apply(Semigroup0_1_0.V0, acc_5)
_ = __local_var_6_2
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
__local_var_2_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_5 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_3_5
__local_var_4_6 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_4_6
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(__local_var_4_6.IntVal), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_4_6.IntVal), __local_var_3_5, a_5)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), acc_1})})
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(__local_var_2_4, x_3).UnsafePtr).V1
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_7 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, *Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_8 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_4_8
__local_var_5_9 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0
_ = __local_var_5_9
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(__local_var_5_9.IntVal), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_6, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__local_var_4_8)})}})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](xs_2))}))
_ = v_3_7
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_10 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_10
__local_var_6_11 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0
_ = __local_var_6_11
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(__local_var_6_11.IntVal), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(__local_var_6_11.IntVal), gopurs_runtime.Int(1)).IntVal), a_7, __local_var_5_10)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_7)}.UnsafePtr).V0.IntVal), b_1})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_7)}.UnsafePtr).V1))}).UnsafePtr).V1
})
})
}))
	})
	return cache_foldableWithIndexList__3168570882
}

var cache_foldableWithIndexList__184979230 gopurs_runtime.Value
var once_foldableWithIndexList__184979230 sync.Once
func Get_foldableWithIndexList__184979230() gopurs_runtime.Value {
	once_foldableWithIndexList__184979230.Do(func() {
		cache_foldableWithIndexList__184979230 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableList()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableWithIndexList(), "foldlWithIndex"), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_2 := gopurs_runtime.Apply(Semigroup0_1_0.V0, acc_5)
_ = __local_var_6_2
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
__local_var_2_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_5 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_3_5
__local_var_4_6 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_4_6
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(__local_var_4_6.IntVal), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_4_6.IntVal), __local_var_3_5, a_5)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), acc_1})})
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(__local_var_2_4, x_3).UnsafePtr).V1
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_7 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, *Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_8 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_4_8
__local_var_5_9 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0
_ = __local_var_5_9
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(__local_var_5_9.IntVal), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_6, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__local_var_4_8)})}})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](xs_2))}))
_ = v_3_7
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_10 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_10
__local_var_6_11 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0
_ = __local_var_6_11
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(__local_var_6_11.IntVal), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(__local_var_6_11.IntVal), gopurs_runtime.Int(1)).IntVal), a_7, __local_var_5_10)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_7)}.UnsafePtr).V0.IntVal), b_1})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_7)}.UnsafePtr).V1))}).UnsafePtr).V1
})
})
}))
	})
	return cache_foldableWithIndexList__184979230
}

var cache_foldableWithIndexNonEmptyList__3425888830 gopurs_runtime.Value
var once_foldableWithIndexNonEmptyList__3425888830 sync.Once
func Get_foldableWithIndexNonEmptyList__3425888830() gopurs_runtime.Value {
	once_foldableWithIndexNonEmptyList__3425888830.Do(func() {
		cache_foldableWithIndexNonEmptyList__3425888830 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableNonEmptyList()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply2(Get_maybe__919206801(), gopurs_runtime.Int(0), gopurs_runtime.Apply(Get_add__560788792(), gopurs_runtime.Int(1)))
_ = __local_var_3_0
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_foldableWithIndexNonEmpty()).V1, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_0))}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(__local_var_3_0, x_4))
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply2(Get_maybe__919206801(), gopurs_runtime.Int(0), gopurs_runtime.Apply(Get_add__560788792(), gopurs_runtime.Int(1)))
_ = __local_var_3_1
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_foldableWithIndexNonEmpty()).V2, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
}), b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_2 := gopurs_runtime.Apply2(Get_maybe__919206801(), gopurs_runtime.Int(0), gopurs_runtime.Apply(Get_add__560788792(), gopurs_runtime.Int(1)))
_ = __local_var_3_2
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_foldableWithIndexNonEmpty()).V3, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(__local_var_3_2, x_4))
}), b_1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2))})
})
})
}))
	})
	return cache_foldableWithIndexNonEmptyList__3425888830
}

var cache_functorList__4121998062 gopurs_runtime.Value
var once_functorList__4121998062 sync.Once
func Get_functorList__4121998062() gopurs_runtime.Value {
	once_functorList__4121998062.Do(func() {
		cache_functorList__4121998062 = gopurs_runtime.RecordDict1("map", Get_listMap())
	})
	return cache_functorList__4121998062
}

var cache_functorList__1783129585 gopurs_runtime.Value
var once_functorList__1783129585 sync.Once
func Get_functorList__1783129585() gopurs_runtime.Value {
	once_functorList__1783129585.Do(func() {
		cache_functorList__1783129585 = gopurs_runtime.RecordDict1("map", Get_listMap())
	})
	return cache_functorList__1783129585
}

var cache_functorNonEmptyList__257963697 gopurs_runtime.Value
var once_functorNonEmptyList__257963697 sync.Once
func Get_functorNonEmptyList__257963697() gopurs_runtime.Value {
	once_functorNonEmptyList__257963697.Do(func() {
		cache_functorNonEmptyList__257963697 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorList(), "map"), f_0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_functorNonEmptyList__257963697
}

var cache_functorWithIndexList__1995002722 gopurs_runtime.Value
var once_functorWithIndexList__1995002722 sync.Once
func Get_functorWithIndexList__1995002722() gopurs_runtime.Value {
	once_functorWithIndexList__1995002722.Do(func() {
		cache_functorWithIndexList__1995002722 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableWithIndexList(), "foldrWithIndex"), gopurs_runtime.Func(func(i_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, gopurs_runtime.Int(i_1.IntVal), x_2), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](acc_3)})}
})
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})
}))
	})
	return cache_functorWithIndexList__1995002722
}

var cache_functorWithIndexNonEmptyList__3683208290 gopurs_runtime.Value
var once_functorWithIndexNonEmptyList__3683208290 sync.Once
func Get_functorWithIndexNonEmptyList__3683208290() gopurs_runtime.Value {
	once_functorWithIndexNonEmptyList__3683208290.Do(func() {
		cache_functorWithIndexNonEmptyList__3683208290 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
}), gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply2(Get_maybe__919206801(), gopurs_runtime.Int(0), gopurs_runtime.Apply(Get_add__560788792(), gopurs_runtime.Int(1)))
_ = __local_var_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_functorWithIndex()).V1, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(fn_0, gopurs_runtime.Apply(__local_var_2_0, x_3))
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1))})))}
})
}))
	})
	return cache_functorWithIndexNonEmptyList__3683208290
}

var cache_listMap__858544730 gopurs_runtime.Value
var once_listMap__858544730 sync.Once
func Get_listMap__858544730() gopurs_runtime.Value {
	once_listMap__858544730.Do(func() {
		cache_listMap__858544730 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_listMap__858544730(f_0_box)
})
	})
	return cache_listMap__858544730
}

var cache_listMap__4135416762 gopurs_runtime.Value
var once_listMap__4135416762 sync.Once
func Get_listMap__4135416762() gopurs_runtime.Value {
	once_listMap__4135416762.Do(func() {
		cache_listMap__4135416762 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_listMap__4135416762(f_0_box)
})
	})
	return cache_listMap__4135416762
}

var cache_monadList__4139186259 gopurs_runtime.Value
var once_monadList__4139186259 sync.Once
func Get_monadList__4139186259() gopurs_runtime.Value {
	once_monadList__4139186259.Do(func() {
		cache_monadList__4139186259 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindList()
}))
	})
	return cache_monadList__4139186259
}

var cache_nelCons__195558898 gopurs_runtime.Value
var once_nelCons__195558898 sync.Once
func Get_nelCons__195558898() gopurs_runtime.Value {
	once_nelCons__195558898.Do(func() {
		cache_nelCons__195558898 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_nelCons__195558898(a_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_1_box)))}
})
	})
	return cache_nelCons__195558898
}

var cache_nelCons__2148523118 gopurs_runtime.Value
var once_nelCons__2148523118 sync.Once
func Get_nelCons__2148523118() gopurs_runtime.Value {
	once_nelCons__2148523118.Do(func() {
		cache_nelCons__2148523118 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_nelCons__2148523118(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](a_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], *Constructor_Cons[gopurs_runtime.Value]]](v_1_box)))}
})
	})
	return cache_nelCons__2148523118
}

var cache_plusList__2201439314 gopurs_runtime.Value
var once_plusList__2201439314 sync.Once
func Get_plusList__2201439314() gopurs_runtime.Value {
	once_plusList__2201439314.Do(func() {
		cache_plusList__2201439314 = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_altList()
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})
	})
	return cache_plusList__2201439314
}

var cache_semigroupList__2766094215 gopurs_runtime.Value
var once_semigroupList__2766094215 sync.Once
func Get_semigroupList__2766094215() gopurs_runtime.Value {
	once_semigroupList__2766094215.Do(func() {
		cache_semigroupList__2766094215 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](ys_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](xs_0))})))}
})
}))
	})
	return cache_semigroupList__2766094215
}

var cache_semigroupList__3410686552 gopurs_runtime.Value
var once_semigroupList__3410686552 sync.Once
func Get_semigroupList__3410686552() gopurs_runtime.Value {
	once_semigroupList__3410686552.Do(func() {
		cache_semigroupList__3410686552 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](ys_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](xs_0))})))}
})
}))
	})
	return cache_semigroupList__3410686552
}

var cache_toList__2859885498 gopurs_runtime.Value
var once_toList__2859885498 sync.Once
func Get_toList__2859885498() gopurs_runtime.Value {
	once_toList__2859885498.Do(func() {
		cache_toList__2859885498 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_toList__2859885498(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_toList__2859885498
}

var cache_toList__2402503393 gopurs_runtime.Value
var once_toList__2402503393 sync.Once
func Get_toList__2402503393() gopurs_runtime.Value {
	once_toList__2402503393.Do(func() {
		cache_toList__2402503393 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_toList__2402503393(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_toList__2402503393
}

var cache_traversable1NonEmptyList__1171985061 gopurs_runtime.Value
var once_traversable1NonEmptyList__1171985061 sync.Once
func Get_traversable1NonEmptyList__1171985061() gopurs_runtime.Value {
	once_traversable1NonEmptyList__1171985061.Do(func() {
		cache_traversable1NonEmptyList__1171985061 = gopurs_runtime.RecordDict4("Foldable10", "Traversable1", "sequence1", "traverse1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldable1NonEmptyList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_traversableNonEmptyList()
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversable1NonEmptyList(), "traverse1"), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0))}, Get_identity1())
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V1)})}})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0, gopurs_runtime.RecordGet(Get_plusList(), "empty")})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1))})))}
}), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_5_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_2
__local_var_5_1 := gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(Functor0_5_2.V0, gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_8, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_7.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_7.UnsafePtr).V1)})}})}
})
}), acc_4), b_6)
})
_ = __local_var_5_1
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_1, gopurs_runtime.Apply(f_2, x_6))
})
}), gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.RecordGet(Get_applicativeNonEmptyList(), "pure"), gopurs_runtime.Apply(f_2, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))}))
})
})
}))
	})
	return cache_traversable1NonEmptyList__1171985061
}

var cache_traversableList__3361437934 gopurs_runtime.Value
var once_traversableList__3361437934 sync.Once
func Get_traversableList__3361437934() gopurs_runtime.Value {
	once_traversableList__3361437934.Do(func() {
		cache_traversableList__3361437934 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableList(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
Apply0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(Functor0_1_0.V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_5, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](b_4)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}))
_ = __local_var_4_2
__local_var_5_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_6_5 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(Apply0_2_1.V0, gopurs_runtime.Value{}))
_ = Functor0_6_5
__local_var_6_4 := gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_2_1.V1, gopurs_runtime.Apply2(Functor0_6_5.V0, gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_9, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](b_8)})}
})
}), acc_5), b_7)
})
_ = __local_var_6_4
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, gopurs_runtime.Apply(f_3, x_7))
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}))
_ = __local_var_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_5_3, x_6))
})
})
}))
	})
	return cache_traversableList__3361437934
}

var cache_traversableList__365792206 gopurs_runtime.Value
var once_traversableList__365792206 sync.Once
func Get_traversableList__365792206() gopurs_runtime.Value {
	once_traversableList__365792206.Do(func() {
		cache_traversableList__365792206 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableList()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorList()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableList(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
Apply0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(Functor0_1_0.V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_5, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](b_4)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}))
_ = __local_var_4_2
__local_var_5_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_6_5 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(Apply0_2_1.V0, gopurs_runtime.Value{}))
_ = Functor0_6_5
__local_var_6_4 := gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_2_1.V1, gopurs_runtime.Apply2(Functor0_6_5.V0, gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_9, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](b_8)})}
})
}), acc_5), b_7)
})
_ = __local_var_6_4
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, gopurs_runtime.Apply(f_3, x_7))
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}))
_ = __local_var_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_5_3, x_6))
})
})
}))
	})
	return cache_traversableList__365792206
}

var cache_traversableNonEmptyList__1085933743 gopurs_runtime.Value
var once_traversableNonEmptyList__1085933743 sync.Once
func Get_traversableNonEmptyList__1085933743() gopurs_runtime.Value {
	once_traversableNonEmptyList__1085933743.Do(func() {
		cache_traversableNonEmptyList__1085933743 = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_traversableNonEmpty(), Get_traversableList())
	})
	return cache_traversableNonEmptyList__1085933743
}

var cache_unfoldable1List__3672302568 gopurs_runtime.Value
var once_unfoldable1List__3672302568 sync.Once
func Get_unfoldable1List__3672302568() gopurs_runtime.Value {
	once_unfoldable1List__3672302568.Do(func() {
		cache_unfoldable1List__3672302568 = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_28 gopurs_runtime.Value
go__go_2_0_28 = gopurs_runtime.Func(func(source_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var source_3_loop gopurs_runtime.Value = source_3_loop_val
var memo_4_loop *Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](memo_4_loop_val)
go__go_2_0_28:
for {
if false { continue go__go_2_0_28 }
var source_3 gopurs_runtime.Value = source_3_loop
_ = source_3
var memo_4 *Constructor_Cons[gopurs_runtime.Value] = memo_4_loop
_ = memo_4
v_5_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]](gopurs_runtime.Apply(f_0, source_3))
_ = v_5_1
var __t4 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 930809136 && __t_tag_2.UnsafePtr != nil) {
source_3_loop = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V1.UnsafePtr).V0
memo_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V0, memo_4})})
continue go__go_2_0_28
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V1
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 930809136 && __t_tag_3.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_7, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](b_6)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V0, memo_4})})))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t4))}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_2_0_28, b_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})))}
})
}))
	})
	return cache_unfoldable1List__3672302568
}

var cache_maybe__919206801 gopurs_runtime.Value
var once_maybe__919206801 sync.Once
func Get_maybe__919206801() gopurs_runtime.Value {
	once_maybe__919206801.Do(func() {
		cache_maybe__919206801 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__919206801(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__919206801
}

var cache_maybe__3658316244 gopurs_runtime.Value
var once_maybe__3658316244 sync.Once
func Get_maybe__3658316244() gopurs_runtime.Value {
	once_maybe__3658316244.Do(func() {
		cache_maybe__3658316244 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__3658316244(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__3658316244
}

var cache_compare__821463600 gopurs_runtime.Value
var once_compare__821463600 sync.Once
func Get_compare__821463600() gopurs_runtime.Value {
	once_compare__821463600.Do(func() {
		cache_compare__821463600 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__821463600(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__821463600
}

var cache_sub__1124926121 gopurs_runtime.Value
var once_sub__1124926121 sync.Once
func Get_sub__1124926121() gopurs_runtime.Value {
	once_sub__1124926121.Do(func() {
		cache_sub__1124926121 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__1124926121(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__1124926121
}

var cache_sub__1043827704 gopurs_runtime.Value
var once_sub__1043827704 sync.Once
func Get_sub__1043827704() gopurs_runtime.Value {
	once_sub__1043827704.Do(func() {
		cache_sub__1043827704 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__1043827704(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_sub__1043827704
}

var cache_sub__3675938712 gopurs_runtime.Value
var once_sub__3675938712 sync.Once
func Get_sub__3675938712() gopurs_runtime.Value {
	once_sub__3675938712.Do(func() {
		cache_sub__3675938712 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__3675938712(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__3675938712
}

var cache_traverse1__157785005 gopurs_runtime.Value
var once_traverse1__157785005 sync.Once
func Get_traverse1__157785005() gopurs_runtime.Value {
	once_traverse1__157785005.Do(func() {
		cache_traverse1__157785005 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse1__157785005(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Traversable.Constructor_Traversable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_traverse1__157785005
}

var cache_traverse1__42886725 gopurs_runtime.Value
var once_traverse1__42886725 sync.Once
func Get_traverse1__42886725() gopurs_runtime.Value {
	once_traverse1__42886725.Do(func() {
		cache_traverse1__42886725 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse1__42886725(dictApply_0_box)
})
	})
	return cache_traverse1__42886725
}

var cache_append__1124926121 gopurs_runtime.Value
var once_append__1124926121 sync.Once
func Get_append__1124926121() gopurs_runtime.Value {
	once_append__1124926121.Do(func() {
		cache_append__1124926121 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__1124926121(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__1124926121
}

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__493084344(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_append__493084344
}

var cache_append__1230318264 gopurs_runtime.Value
var once_append__1230318264 sync.Once
func Get_append__1230318264() gopurs_runtime.Value {
	once_append__1230318264.Do(func() {
		cache_append__1230318264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__1230318264(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__1230318264
}

var cache_append__2013893496 gopurs_runtime.Value
var once_append__2013893496 sync.Once
func Get_append__2013893496() gopurs_runtime.Value {
	once_append__2013893496.Do(func() {
		cache_append__2013893496 = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, ys_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_append__2013893496(gopurs_runtime.CoerceToStruct[Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]]](xs_0_box), gopurs_runtime.CoerceToStruct[Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]]](ys_1_box)))}
})
	})
	return cache_append__2013893496
}

var cache_add__1124926121 gopurs_runtime.Value
var once_add__1124926121 sync.Once
func Get_add__1124926121() gopurs_runtime.Value {
	once_add__1124926121.Do(func() {
		cache_add__1124926121 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__1124926121(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_add__1124926121
}

var cache_add__560788792 gopurs_runtime.Value
var once_add__560788792 sync.Once
func Get_add__560788792() gopurs_runtime.Value {
	once_add__560788792.Do(func() {
		cache_add__560788792 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__560788792(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_add__560788792
}

var cache_add__1614463960 gopurs_runtime.Value
var once_add__1614463960 sync.Once
func Get_add__1614463960() gopurs_runtime.Value {
	once_add__1614463960.Do(func() {
		cache_add__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_add__1614463960
}

var cache_show__2742601362 gopurs_runtime.Value
var once_show__2742601362 sync.Once
func Get_show__2742601362() gopurs_runtime.Value {
	once_show__2742601362.Do(func() {
		cache_show__2742601362 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__2742601362(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__2742601362
}

var cache_show__842947602 gopurs_runtime.Value
var once_show__842947602 sync.Once
func Get_show__842947602() gopurs_runtime.Value {
	once_show__842947602.Do(func() {
		cache_show__842947602 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__842947602(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[*pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_show__842947602
}

var cache_traverse__314957093 gopurs_runtime.Value
var once_traverse__314957093 sync.Once
func Get_traverse__314957093() gopurs_runtime.Value {
	once_traverse__314957093.Do(func() {
		cache_traverse__314957093 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse__314957093(gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_traverse__314957093
}

var cache_traverse__878259545 gopurs_runtime.Value
var once_traverse__878259545 sync.Once
func Get_traverse__878259545() gopurs_runtime.Value {
	once_traverse__878259545.Do(func() {
		cache_traverse__878259545 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse__878259545(dictApplicative_0_box)
})
	})
	return cache_traverse__878259545
}

var cache_traverse__2839486329 gopurs_runtime.Value
var once_traverse__2839486329 sync.Once
func Get_traverse__2839486329() gopurs_runtime.Value {
	once_traverse__2839486329.Do(func() {
		cache_traverse__2839486329 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse__2839486329(dictApplicative_0_box)
})
	})
	return cache_traverse__2839486329
}

var cache_traverseWithIndex__2726076659 gopurs_runtime.Value
var once_traverseWithIndex__2726076659 sync.Once
func Get_traverseWithIndex__2726076659() gopurs_runtime.Value {
	once_traverseWithIndex__2726076659.Do(func() {
		cache_traverseWithIndex__2726076659 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverseWithIndex__2726076659(gopurs_runtime.CoerceToStruct[pkg_Data_TraversableWithIndex.Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_traverseWithIndex__2726076659
}

var cache_traverseWithIndex__1901281819 gopurs_runtime.Value
var once_traverseWithIndex__1901281819 sync.Once
func Get_traverseWithIndex__1901281819() gopurs_runtime.Value {
	once_traverseWithIndex__1901281819.Do(func() {
		cache_traverseWithIndex__1901281819 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverseWithIndex__1901281819(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_traverseWithIndex__1901281819
}

var cache_snd__1234761462 gopurs_runtime.Value
var once_snd__1234761462 sync.Once
func Get_snd__1234761462() gopurs_runtime.Value {
	once_snd__1234761462.Do(func() {
		cache_snd__1234761462 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_snd__1234761462(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_snd__1234761462
}

var cache_snd__3058387254 gopurs_runtime.Value
var once_snd__3058387254 sync.Once
func Get_snd__3058387254() gopurs_runtime.Value {
	once_snd__3058387254.Do(func() {
		cache_snd__3058387254 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_snd__3058387254(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, *Constructor_Cons[gopurs_runtime.Value]]](v_0_box)))}
})
	})
	return cache_snd__3058387254
}

var cache_snd__21214742 gopurs_runtime.Value
var once_snd__21214742 sync.Once
func Get_snd__21214742() gopurs_runtime.Value {
	once_snd__21214742.Do(func() {
		cache_snd__21214742 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_snd__21214742(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_snd__21214742
}

type Constructor_Nil[T_a any] struct {
	Rc uint32
}


type Constructor_Cons[T_a any] struct {
	Rc uint32
	V0 T_a
	V1 *Constructor_Cons[gopurs_runtime.Value]
}


func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_identity1(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_NonEmptyList(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_toList(v_0_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Cons[gopurs_runtime.Value] {
var v_0 *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)})})
}

func Call_nelCons(a_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)})}})})
}

func Call_listMap(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var chunkedRevMap_1_0_0 gopurs_runtime.Value
chunkedRevMap_1_0_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]] = gopurs_runtime.CoerceToStruct[Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]]](v_2_loop_val)
var v1_3_loop *Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v1_3_loop_val)
chunkedRevMap_1_0_0:
for {
if false { continue chunkedRevMap_1_0_0 }
var v_2 *Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]] = v_2_loop
_ = v_2
var v1_3 *Constructor_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t19 *Constructor_Cons[gopurs_runtime.Value]
{
var __t_and_18 bool = false
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr != nil) {

var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
var __t_and_17 bool = false
if (__t_tag_15.Type == 9 && __t_tag_15.IntVal == 1358893437 && __t_tag_15.UnsafePtr != nil) {

var __t_tag_16 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}
__t_and_17 = (__t_tag_16.Type == 9 && __t_tag_16.IntVal == 1358893437 && __t_tag_16.UnsafePtr != nil)
}
__t_and_18 = __t_and_17
}
if __t_and_18 {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)})})})
v1_3_loop = (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr).V1
continue chunkedRevMap_1_0_0
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
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
var v3_6_loop *Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v3_6_loop_val)
reverseUnrolledMap_4_1_1:
for {
if false { continue reverseUnrolledMap_4_1_1 }
var v2_5 gopurs_runtime.Value = v2_5_loop
_ = v2_5
var v3_6 *Constructor_Cons[gopurs_runtime.Value] = v3_6_loop
_ = v3_6
var __t8 *Constructor_Cons[gopurs_runtime.Value]
{
var __t_and_7 bool = false
if (v2_5.Type == 9 && v2_5.IntVal == 1358893437 && v2_5.UnsafePtr != nil) {

var __t_tag_2 gopurs_runtime.Value = (*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0
var __t_and_6 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1358893437 && __t_tag_2.UnsafePtr != nil) {

var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])((*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}
var __t_and_5 bool = false
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 1358893437 && __t_tag_3.UnsafePtr != nil) {

var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])((*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V1)}
__t_and_5 = (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 1358893437 && __t_tag_4.UnsafePtr != nil)
}
__t_and_6 = __t_and_5
}
__t_and_7 = __t_and_6
}
if __t_and_7 {
v2_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V1)}
v3_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])((*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])((*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])((*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr).V0), v3_6})})})})})})
continue reverseUnrolledMap_4_1_1
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
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
var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 1358893437 && __t_tag_9.UnsafePtr != nil) {
var __t11 gopurs_runtime.Value
{
var __t_tag_10 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}
if (__t_tag_10.Type == 9 && __t_tag_10.IntVal == 1358893437 && __t_tag_10.UnsafePtr == nil) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})})})})}
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_11:
__t13 = __t11
goto end_branch_13
} else {

}
}
{
var __t_tag_12 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 1358893437 && __t_tag_12.UnsafePtr == nil) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})})}
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_13:
__t14 = __t13
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_14:
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(reverseUnrolledMap_4_1_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t14))}))
}
end_branch_19:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t19)}
}
}()
})
})
return gopurs_runtime.Apply(chunkedRevMap_1_0_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})
}

func Call_showList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
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
__t1 = Call_append__493084344(gopurs_runtime.Str("("), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(Call_intercalate__2937349250(gopurs_runtime.Str(" : "), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[string]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorList(), "map"), show_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v_2))})))}).StrVal()), gopurs_runtime.Str(" : Nil)")).StrVal())).StrVal()
}
end_branch_1:
return gopurs_runtime.Str(__t1)
}))
}

func Call_showNonEmptyList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
showNonEmpty_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[*pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](gopurs_runtime.Apply2(pkg_Data_NonEmpty.Get_showNonEmpty(), dictShow_0, Call_showList(dictShow_0)))
_ = showNonEmpty_1_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(NonEmptyList "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(showNonEmpty_1_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](v_2))}).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
}))
}

func Call_eqList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_eq1List(), "eq1"), dictEq_0))
}

func Call_eqNonEmptyList(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply2(pkg_Data_NonEmpty.Get_eqNonEmpty(), Get_eq1List(), dictEq_0)
}

func Call_ordList(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
eqList1_1_0 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_eq1List(), "eq1"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})))
_ = eqList1_1_0
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return eqList1_1_0
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_ord1List(), "compare1"), dictOrd_0))
}

func Call_ordNonEmptyList(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_ordNonEmpty(), dictOrd_0)
}

func Call_pure__189931222(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__2331489366(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__993904534(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__243192752(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__575667894(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __eta0_0, gopurs_runtime.RecordGet(Get_plusList(), "empty")})}
}

func Call_apply__1030762512(dict_0_loop *pkg_Control_Apply.Constructor_Apply[*Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[*Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__2169384906(v_0_loop *Constructor_Cons[gopurs_runtime.Value], v1_1_loop *Constructor_Cons[gopurs_runtime.Value]) *Constructor_Cons[gopurs_runtime.Value] {
var v_0 *Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var v1_1 *Constructor_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorList(), "map"), (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)})))})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t0)
}

func Call_lift2__1942544886(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
})
})
}

func Call_lift2__3139828374(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
})
})
}

func Call_lift2__3684551766(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
})
})
}

func Call_lift2__2286084406(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
})
})
}

func Call_lift2__3007543670(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
})
})
}

func Call_lift2__2762258480(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
})
})
}

func Call_lift2__2618178704(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
})
})
}

func Call_lift2__2114580400(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
})
})
}

func Call_lift2__3213187376(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
})
})
}

func Call_bind__3842572251(dict_0_loop *pkg_Control_Bind.Constructor_Bind[*Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[*Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__1872090113(v_0_loop *Constructor_Cons[gopurs_runtime.Value], v1_1_loop gopurs_runtime.Value) *Constructor_Cons[gopurs_runtime.Value] {
var v_0 *Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0)))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}, v1_1)))})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t0)
}

func Call_compose__4141960292(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[*Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[*Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__1254722180(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[*Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[*Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__1555187646(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__2527254334(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__2532574046(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__794534846(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__4254807102(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__2995688990(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__3140790526(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__3384557662(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__2710321297(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__1933206353(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_foldl__3850309840(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[*Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[*Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__2111289130(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__3041692656(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[*Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[*Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__66388714(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__1671904522(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__2602334544(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[*Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[*Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__371433392(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[*Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[*Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__3619558698(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__1601164432(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[*Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[*Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__1712912315(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_8 gopurs_runtime.Value
go__go_1_0_8 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop *Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](b_2_loop_val)
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_0_8:
for {
if false { continue go__go_1_0_8 }
var b_2 *Constructor_Cons[gopurs_runtime.Value] = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(b_2)}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(b_2)}, (*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0))
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
continue go__go_1_0_8
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t1))}
}
}()
})
})
return go__go_1_0_8
}

func Call_foldl__2159564571(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_9 gopurs_runtime.Value
go__go_1_0_9 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop *pkg_Data_Tuple.Constructor_Tuple[int64, *Constructor_Cons[gopurs_runtime.Value]] = gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, *Constructor_Cons[gopurs_runtime.Value]]](b_2_loop_val)
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_0_9:
for {
if false { continue go__go_1_0_9 }
var b_2 *pkg_Data_Tuple.Constructor_Tuple[int64, *Constructor_Cons[gopurs_runtime.Value]] = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(b_2)}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, *Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(b_2)}, (*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0))
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
continue go__go_1_0_9
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, *Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, *Constructor_Cons[gopurs_runtime.Value]]](__t1))}
}
}()
})
})
return go__go_1_0_9
}

func Call_foldl__3785384859(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__3943124669(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_10 gopurs_runtime.Value
go__go_1_0_10 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_0_10:
for {
if false { continue go__go_1_0_10 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t1 = b_2
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(f_0, b_2, (*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
continue go__go_1_0_10
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
return go__go_1_0_10
}

func Call_foldl__396932925(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_11 gopurs_runtime.Value
go__go_1_0_11 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop *Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](b_2_loop_val)
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_0_11:
for {
if false { continue go__go_1_0_11 }
var b_2 *Constructor_Cons[gopurs_runtime.Value] = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(b_2)}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(b_2)}, (*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0))
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
continue go__go_1_0_11
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t1))}
}
}()
})
})
return go__go_1_0_11
}

func Call_foldl__2928402749(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_12 gopurs_runtime.Value
go__go_1_0_12 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](b_2_loop_val)
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_0_12:
for {
if false { continue go__go_1_0_12 }
var b_2 *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(b_2)}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(b_2)}, (*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0))
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
continue go__go_1_0_12
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](__t1))}
}
}()
})
})
return go__go_1_0_12
}

func Call_foldl__255626813(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_13 gopurs_runtime.Value
go__go_1_0_13 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop *pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](b_2_loop_val)
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_0_13:
for {
if false { continue go__go_1_0_13 }
var b_2 *pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value] = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(b_2)}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(b_2)}, (*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0))
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
continue go__go_1_0_13
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]](__t1))}
}
}()
})
})
return go__go_1_0_13
}

func Call_foldl__3915700701(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_14 gopurs_runtime.Value
go__go_1_0_14 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_0_14:
for {
if false { continue go__go_1_0_14 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t1 = b_2
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(f_0, b_2, (*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
continue go__go_1_0_14
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
return go__go_1_0_14
}

func Call_foldl__3459294429(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_15 gopurs_runtime.Value
go__go_1_0_15 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_0_15:
for {
if false { continue go__go_1_0_15 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t1 = b_2
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(f_0, b_2, (*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
continue go__go_1_0_15
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
return go__go_1_0_15
}

func Call_foldr__2512763050(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__3673994608(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[*Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[*Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__3943124669(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
})
}), b_1)
_ = __local_var_2_0
var go__go_3_2_16 gopurs_runtime.Value
go__go_3_2_16 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_2_16:
for {
if false { continue go__go_3_2_16 }
var v_4 *Constructor_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
continue go__go_3_2_16
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t3))}
}
}()
})
})
__local_var_3_1 := gopurs_runtime.Apply(go__go_3_2_16, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
})
}

func Call_foldr__2979608669(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *Constructor_Cons[gopurs_runtime.Value] = b_1_loop
_ = b_1
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(b_1)})
_ = __local_var_2_0
var go__go_3_2_17 gopurs_runtime.Value
go__go_3_2_17 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_2_17:
for {
if false { continue go__go_3_2_17 }
var v_4 *Constructor_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
continue go__go_3_2_17
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t3))}
}
}()
})
})
__local_var_3_1 := gopurs_runtime.Apply(go__go_3_2_17, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
})
}

func Call_foldr__4137485405(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
})
}), b_1)
_ = __local_var_2_0
var go__go_3_2_18 gopurs_runtime.Value
go__go_3_2_18 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_2_18:
for {
if false { continue go__go_3_2_18 }
var v_4 *Constructor_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
continue go__go_3_2_18
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t3))}
}
}()
})
})
__local_var_3_1 := gopurs_runtime.Apply(go__go_3_2_18, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
})
}

func Call_intercalate__3939234276(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], dictMonoid_1_loop *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonoid_1 *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonoid_1.V0, gopurs_runtime.Value{}))
_ = Semigroup0_2_0
mempty_3_1 := dictMonoid_1.V1
_ = mempty_3_1
return gopurs_runtime.Func(func(sep_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(dictFoldable_0.V1, gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_6, "init").IntVal) != (0) {
__t2 = gopurs_runtime.RecordDict2("acc", "init", v1_7, gopurs_runtime.Bool(false))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Apply2(Semigroup0_2_0.V0, gopurs_runtime.RecordGet(v_6, "acc"), gopurs_runtime.Apply2(Semigroup0_2_0.V0, sep_4, v1_7)), gopurs_runtime.Bool(false))
}
end_branch_2:
return __t2
})
}), gopurs_runtime.RecordDict2("acc", "init", mempty_3_1, gopurs_runtime.Bool(true)), xs_5), "acc")
})
})
}

func Call_intercalate__2937349250(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_2, "init").IntVal) != (0) {
__t0 = gopurs_runtime.RecordDict2("acc", "init", v1_3, gopurs_runtime.Bool(false))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.RecordGet(v_2, "acc"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), __eta0_0, v1_3)), gopurs_runtime.Bool(false))
}
end_branch_0:
return __t0
})
}), gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.RecordGet(pkg_Data_Monoid.Get_monoidString(), "mempty"), gopurs_runtime.Bool(true)), __eta1_1), "acc")
}

func Call_foldMapWithIndex__2292551140(dict_0_loop *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldMapWithIndex__1722031522(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](pkg_Data_List_Lazy_Types.Get_foldableWithIndexNonEmpty()).V1, __eta0_0, __eta1_1)
}

func Call_foldMapWithIndex__2880267906(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_foldableWithIndexNonEmpty()).V1, __eta0_0, __eta1_1)
}

func Call_foldlWithIndex__2808220203(f_0_loop gopurs_runtime.Value, acc_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var acc_1 gopurs_runtime.Value = acc_1_loop
_ = acc_1
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_3_1
__local_var_4_2 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_4_2
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_4_2.IntVal) + (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_4_2.IntVal), __local_var_3_1, a_5)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), acc_1})})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(__local_var_2_0, x_3).UnsafePtr).V1
})
}

func Call_foldlWithIndex__2764250251(f_0_loop gopurs_runtime.Value, acc_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var acc_1 gopurs_runtime.Value = acc_1_loop
_ = acc_1
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_3_1
__local_var_4_2 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_4_2
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_4_2.IntVal) + (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_4_2.IntVal), __local_var_3_1, a_5)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), acc_1})})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(__local_var_2_0, x_3).UnsafePtr).V1
})
}

func Call_foldlWithIndex__2986161357(dict_0_loop *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldlWithIndex__3610348555(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](pkg_Data_List_Lazy_Types.Get_foldableWithIndexNonEmpty()).V2, __eta0_0, __eta1_1, __eta2_2)
}

func Call_foldlWithIndex__1651851147(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_foldableWithIndexNonEmpty()).V2, __eta0_0, __eta1_1, __eta2_2)
}

func Call_foldrWithIndex__2808220203(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, xs_2_loop *Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var xs_2 *Constructor_Cons[gopurs_runtime.Value] = xs_2_loop
_ = xs_2
v_3_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, *Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_4_1
__local_var_5_2 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0
_ = __local_var_5_2
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_5_2.IntVal) + (1)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_6, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__local_var_4_1)})}})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_2)}))
_ = v_3_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_3 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_3
__local_var_6_4 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0
_ = __local_var_6_4
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_6_4.IntVal) - (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((__local_var_6_4.IntVal) - (1)), a_7, __local_var_5_3)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_0)}.UnsafePtr).V0.IntVal), b_1})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_0)}.UnsafePtr).V1))}).UnsafePtr).V1
}

func Call_foldrWithIndex__2439396107(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Cons[gopurs_runtime.Value], xs_2_loop *Constructor_Cons[gopurs_runtime.Value]) *Constructor_Cons[gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *Constructor_Cons[gopurs_runtime.Value] = b_1_loop
_ = b_1
var xs_2 *Constructor_Cons[gopurs_runtime.Value] = xs_2_loop
_ = xs_2
v_3_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, *Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_4_1
__local_var_5_2 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0
_ = __local_var_5_2
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_5_2.IntVal) + (1)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_6, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__local_var_4_1)})}})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_2)}))
_ = v_3_0
return gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_3 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_3
__local_var_6_4 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0
_ = __local_var_6_4
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_6_4.IntVal) - (1)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((__local_var_6_4.IntVal) - (1)), a_7, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__local_var_5_3))})))}})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_0)}.UnsafePtr).V0.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(b_1)}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_0)}.UnsafePtr).V1))}).UnsafePtr).V1)
}

func Call_foldrWithIndex__2986161357(dict_0_loop *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_foldrWithIndex__3610348555(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](pkg_Data_List_Lazy_Types.Get_foldableWithIndexNonEmpty()).V3, __eta0_0, __eta1_1, __eta2_2)
}

func Call_foldrWithIndex__979136683(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply3(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_foldableWithIndexNonEmpty()).V3, __eta0_0, __eta1_1, __eta2_2)
}

func Call_flip__3525081280(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__2974723072(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__3709724320(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__3563101792(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__1833071808(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__1673583840(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__534748448(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__1744188480(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__3468792800(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__872296768(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__4091748192(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_map__3116241637(dict_0_loop *pkg_Data_Functor.Constructor_Functor[*Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[*Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2665381605(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__67411525(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2876470885(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1256368628(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1483545076(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__3058795348(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2190988916(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__843173928(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply(Call_listMap(__eta0_0), __eta1_1)
}

func Call_map__438443400(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply(Call_listMap(__eta0_0), __eta1_1)
}

func Call_mapFlipped__2466386789(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, f_2, fa_1)
}

func Call_mapFlipped__4215217780(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, f_2, fa_1)
}

func Call_mapFlipped__2919806324(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, f_2, fa_1)
}

func Call_mapWithIndex__55256674(dict_0_loop *pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_mapWithIndex__574674314(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](pkg_Data_List_Lazy_Types.Get_functorWithIndex()).V1, __eta0_0, __eta1_1)
}

func Call_mapWithIndex__598554346(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_functorWithIndex()).V1, __eta0_0, __eta1_1)
}

func Call_conj__3676519832(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) && ((__eta1_1.IntVal) != (0)))
}

func Call_conj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_disj__3676519832(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) || ((__eta1_1.IntVal) != (0)))
}

func Call_disj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_not__3201284355(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) != (true))
}

func Call_not__1505204753(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_listMap__858544730(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var chunkedRevMap_1_0_24 gopurs_runtime.Value
chunkedRevMap_1_0_24 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]] = gopurs_runtime.CoerceToStruct[Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]]](v_2_loop_val)
var v1_3_loop *Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v1_3_loop_val)
chunkedRevMap_1_0_24:
for {
if false { continue chunkedRevMap_1_0_24 }
var v_2 *Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]] = v_2_loop
_ = v_2
var v1_3 *Constructor_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t19 *Constructor_Cons[string]
{
var __t_and_18 bool = false
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr != nil) {

var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
var __t_and_17 bool = false
if (__t_tag_15.Type == 9 && __t_tag_15.IntVal == 1358893437 && __t_tag_15.UnsafePtr != nil) {

var __t_tag_16 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}
__t_and_17 = (__t_tag_16.Type == 9 && __t_tag_16.IntVal == 1358893437 && __t_tag_16.UnsafePtr != nil)
}
__t_and_18 = __t_and_17
}
if __t_and_18 {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)})})})
v1_3_loop = (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr).V1
continue chunkedRevMap_1_0_24
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Cons[string]](gopurs_runtime.Value{})
goto end_branch_19
} else {

}
}
{
var reverseUnrolledMap_4_1_25 gopurs_runtime.Value
reverseUnrolledMap_4_1_25 = gopurs_runtime.Func(func(v2_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v2_5_loop gopurs_runtime.Value = v2_5_loop_val
var v3_6_loop *Constructor_Cons[string] = gopurs_runtime.CoerceToStruct[Constructor_Cons[string]](v3_6_loop_val)
reverseUnrolledMap_4_1_25:
for {
if false { continue reverseUnrolledMap_4_1_25 }
var v2_5 gopurs_runtime.Value = v2_5_loop
_ = v2_5
var v3_6 *Constructor_Cons[string] = v3_6_loop
_ = v3_6
var __t8 *Constructor_Cons[string]
{
var __t_and_7 bool = false
if (v2_5.Type == 9 && v2_5.IntVal == 1358893437 && v2_5.UnsafePtr != nil) {

var __t_tag_2 gopurs_runtime.Value = (*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0
var __t_and_6 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1358893437 && __t_tag_2.UnsafePtr != nil) {

var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])((*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}
var __t_and_5 bool = false
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 1358893437 && __t_tag_3.UnsafePtr != nil) {

var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])((*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V1)}
__t_and_5 = (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 1358893437 && __t_tag_4.UnsafePtr != nil)
}
__t_and_6 = __t_and_5
}
__t_and_7 = __t_and_6
}
if __t_and_7 {
v2_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V1)}
v3_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Cons[string]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])((*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V0).StrVal()), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])((*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V0).StrVal()), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])((*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr).V0).StrVal()), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v3_6)})})})})})})})
continue reverseUnrolledMap_4_1_25
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Cons[string]](gopurs_runtime.Value{})
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
var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 1358893437 && __t_tag_9.UnsafePtr != nil) {
var __t11 gopurs_runtime.Value
{
var __t_tag_10 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}
if (__t_tag_10.Type == 9 && __t_tag_10.IntVal == 1358893437 && __t_tag_10.UnsafePtr == nil) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V0).StrVal()), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V0).StrVal()), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})})})})}
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_11:
__t13 = __t11
goto end_branch_13
} else {

}
}
{
var __t_tag_12 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 1358893437 && __t_tag_12.UnsafePtr == nil) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V0).StrVal()), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})})}
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_13:
__t14 = __t13
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_14:
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Cons[string]](gopurs_runtime.Apply2(reverseUnrolledMap_4_1_25, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[string]](__t14))}))
}
end_branch_19:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t19)}
}
}()
})
})
return gopurs_runtime.Apply(chunkedRevMap_1_0_24, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})
}

func Call_listMap__4135416762(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var chunkedRevMap_1_0_26 gopurs_runtime.Value
chunkedRevMap_1_0_26 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]] = gopurs_runtime.CoerceToStruct[Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]]](v_2_loop_val)
var v1_3_loop *Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v1_3_loop_val)
chunkedRevMap_1_0_26:
for {
if false { continue chunkedRevMap_1_0_26 }
var v_2 *Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]] = v_2_loop
_ = v_2
var v1_3 *Constructor_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t19 *Constructor_Cons[gopurs_runtime.Value]
{
var __t_and_18 bool = false
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr != nil) {

var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
var __t_and_17 bool = false
if (__t_tag_15.Type == 9 && __t_tag_15.IntVal == 1358893437 && __t_tag_15.UnsafePtr != nil) {

var __t_tag_16 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}
__t_and_17 = (__t_tag_16.Type == 9 && __t_tag_16.IntVal == 1358893437 && __t_tag_16.UnsafePtr != nil)
}
__t_and_18 = __t_and_17
}
if __t_and_18 {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)})})})
v1_3_loop = (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr).V1
continue chunkedRevMap_1_0_26
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
goto end_branch_19
} else {

}
}
{
var reverseUnrolledMap_4_1_27 gopurs_runtime.Value
reverseUnrolledMap_4_1_27 = gopurs_runtime.Func(func(v2_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v2_5_loop gopurs_runtime.Value = v2_5_loop_val
var v3_6_loop *Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](v3_6_loop_val)
reverseUnrolledMap_4_1_27:
for {
if false { continue reverseUnrolledMap_4_1_27 }
var v2_5 gopurs_runtime.Value = v2_5_loop
_ = v2_5
var v3_6 *Constructor_Cons[gopurs_runtime.Value] = v3_6_loop
_ = v3_6
var __t8 *Constructor_Cons[gopurs_runtime.Value]
{
var __t_and_7 bool = false
if (v2_5.Type == 9 && v2_5.IntVal == 1358893437 && v2_5.UnsafePtr != nil) {

var __t_tag_2 gopurs_runtime.Value = (*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0
var __t_and_6 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1358893437 && __t_tag_2.UnsafePtr != nil) {

var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])((*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}
var __t_and_5 bool = false
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 1358893437 && __t_tag_3.UnsafePtr != nil) {

var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])((*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V1)}
__t_and_5 = (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 1358893437 && __t_tag_4.UnsafePtr != nil)
}
__t_and_6 = __t_and_5
}
__t_and_7 = __t_and_6
}
if __t_and_7 {
v2_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V1)}
v3_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])((*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])((*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])((*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr).V0), v3_6})})})})})})
continue reverseUnrolledMap_4_1_27
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
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
var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 1358893437 && __t_tag_9.UnsafePtr != nil) {
var __t11 gopurs_runtime.Value
{
var __t_tag_10 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}
if (__t_tag_10.Type == 9 && __t_tag_10.IntVal == 1358893437 && __t_tag_10.UnsafePtr == nil) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})})})})}
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_11:
__t13 = __t11
goto end_branch_13
} else {

}
}
{
var __t_tag_12 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 1358893437 && __t_tag_12.UnsafePtr == nil) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})})}
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_13:
__t14 = __t13
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_14:
__t19 = gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(reverseUnrolledMap_4_1_27, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](__t14))}))
}
end_branch_19:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t19)}
}
}()
})
})
return gopurs_runtime.Apply(chunkedRevMap_1_0_26, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))})
}

func Call_nelCons__195558898(a_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)})}})})
}

func Call_nelCons__2148523118(a_0_loop *Constructor_Cons[gopurs_runtime.Value], v_1_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], *Constructor_Cons[gopurs_runtime.Value]]) *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], *Constructor_Cons[gopurs_runtime.Value]] {
var a_0 *Constructor_Cons[gopurs_runtime.Value] = a_0_loop
_ = a_0
var v_1 *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], *Constructor_Cons[gopurs_runtime.Value]] = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], *Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(a_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0))}, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1))})})}})})
}

func Call_toList__2859885498(v_0_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Cons[gopurs_runtime.Value] {
var v_0 *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)})})
}

func Call_toList__2402503393(v_0_loop *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Cons[gopurs_runtime.Value] {
var v_0 *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)})})
}

func Call_maybe__919206801(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_maybe__3658316244(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_sub__1124926121(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_sub__1043827704(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) - (__eta1_1.IntVal))
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_traverse1__157785005(dict_0_loop *pkg_Data_Semigroup_Traversable.Constructor_Traversable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup_Traversable.Constructor_Traversable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_traverse1__42886725(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], *Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V1)})}})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], *Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeNonEmptyList(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0))})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1))})))}
}), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_5_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_2
__local_var_5_1 := gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(Functor0_5_2.V0, gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_8, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_7.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_7.UnsafePtr).V1)})}})}
})
}), acc_4), b_6)
})
_ = __local_var_5_1
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_1, gopurs_runtime.Apply(f_2, x_6))
})
}), gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.RecordGet(Get_applicativeNonEmptyList(), "pure"), gopurs_runtime.Apply(f_2, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0))})), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]]]((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))}))
})
})
}

func Call_append__1124926121(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_append__493084344(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Str((__eta0_0.StrVal()) + (__eta1_1.StrVal()))
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_append__2013893496(xs_0_loop *Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]], ys_1_loop *Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]]) *Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]] {
var xs_0 *Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]] = xs_0_loop
_ = xs_0
var ys_1 *Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]] = ys_1_loop
_ = ys_1
return gopurs_runtime.CoerceToStruct[Constructor_Cons[*Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(ys_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_0)}))
}

func Call_add__1124926121(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_add__560788792(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) + (__eta1_1.IntVal))
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__842947602(dict_0_loop *pkg_Data_Show.Constructor_Show[*pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[*pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_traverse__314957093(dict_0_loop *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_traverse__878259545(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
Apply0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(Functor0_1_0.V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_5, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](b_4)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}))
_ = __local_var_4_2
__local_var_5_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_6_5 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(Apply0_2_1.V0, gopurs_runtime.Value{}))
_ = Functor0_6_5
__local_var_6_4 := gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_2_1.V1, gopurs_runtime.Apply2(Functor0_6_5.V0, gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_9, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](b_8)})}
})
}), acc_5), b_7)
})
_ = __local_var_6_4
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, gopurs_runtime.Apply(f_3, x_7))
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}))
_ = __local_var_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_5_3, x_6))
})
})
}

func Call_traverse__2839486329(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
Apply0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(Functor0_1_0.V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_5, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](b_4)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}))
_ = __local_var_4_2
__local_var_5_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_6_5 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(Apply0_2_1.V0, gopurs_runtime.Value{}))
_ = Functor0_6_5
__local_var_6_4 := gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_2_1.V1, gopurs_runtime.Apply2(Functor0_6_5.V0, gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{1, a_9, gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]](b_8)})}
})
}), acc_5), b_7)
})
_ = __local_var_6_4
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, gopurs_runtime.Apply(f_3, x_7))
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(nil))}))
_ = __local_var_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_5_3, x_6))
})
})
}

func Call_traverseWithIndex__2726076659(dict_0_loop *pkg_Data_TraversableWithIndex.Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_TraversableWithIndex.Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_traverseWithIndex__1901281819(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_TraversableWithIndex.Constructor_TraversableWithIndex[*pkg_Data_Maybe.Constructor_Just[int64], *pkg_Data_NonEmpty.Constructor_NonEmpty[*Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_traversableWithIndexNonEmpty()).V3, __eta0_0, __eta1_1)
}

func Call_snd__1234761462(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}

func Call_snd__3058387254(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[int64, *Constructor_Cons[gopurs_runtime.Value]]) *Constructor_Cons[gopurs_runtime.Value] {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[int64, *Constructor_Cons[gopurs_runtime.Value]] = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Cons[gopurs_runtime.Value]]((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)
}

func Call_snd__21214742(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[int64, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}



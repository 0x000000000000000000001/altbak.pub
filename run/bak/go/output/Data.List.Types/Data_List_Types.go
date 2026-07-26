package Data_List_Types

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	pkg_Data_Monoid "gopurs/output/Data.Monoid"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
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

var cache_Nil gopurs_runtime.Value
var once_Nil sync.Once
func Get_Nil() gopurs_runtime.Value {
	once_Nil.Do(func() {
		cache_Nil = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil}
	})
	return cache_Nil
}

var cache_Cons gopurs_runtime.Value
var once_Cons sync.Once
func Get_Cons() gopurs_runtime.Value {
	once_Cons.Do(func() {
		cache_Cons = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{value0, (*Constructor_Cons[gopurs_runtime.Value])(value1.UnsafePtr)})}
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
return Call_toList(v_0_box)
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
return Call_nelCons(a_0_box, v_1_box)
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
		cache_functorNonEmptyList = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply(f_0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorList(), "map"), f_0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V1)})}
}))
	})
	return cache_functorNonEmptyList
}

var cache_foldableList gopurs_runtime.Value
var once_foldableList sync.Once
func Get_foldableList() gopurs_runtime.Value {
	once_foldableList.Do(func() {
		cache_foldableList = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), acc_3)
_ = __local_var_4_1
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_1, gopurs_runtime.Apply(f_2, x_5))
})
}), mempty_1_0)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_2 gopurs_runtime.Value
go__1_2 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__1_2:
for {
if false { continue go__1_2 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t3 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t3 = b_2
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(f_0, b_2, (*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
continue go__1_2
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
return go__1_2
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func2(func(b_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
}), b_1)
_ = __local_var_2_4
var go__3_5 gopurs_runtime.Value
go__3_5 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__3_5:
for {
if false { continue go__3_5 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t6 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t6 = v_4
goto end_branch_6
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{(*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr)})}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
continue go__3_5
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
__local_var_4_7 := gopurs_runtime.Apply(go__3_5, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil})
_ = __local_var_4_7
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_4, gopurs_runtime.Apply(__local_var_4_7, x_5))
})
}))
	})
	return cache_foldableList
}

var cache_intercalate gopurs_runtime.Value
var once_intercalate sync.Once
func Get_intercalate() gopurs_runtime.Value {
	once_intercalate.Do(func() {
		cache_intercalate = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Monoid.Get_monoidString(), "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_0_0
return gopurs_runtime.Func2(func(sep_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_3, "init").IntVal) != (0) {
__t1 = gopurs_runtime.RecordDict2("acc", "init", v1_4, gopurs_runtime.Bool(false))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_0, "append"), gopurs_runtime.RecordGet(v_3, "acc"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_0, "append"), sep_1, v1_4)), gopurs_runtime.Bool(false))
}
end_branch_1:
return __t1
}), gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.RecordGet(pkg_Data_Monoid.Get_monoidString(), "mempty"), gopurs_runtime.Bool(true)), xs_2), "acc")
})
}()
	})
	return cache_intercalate
}

var cache_foldableNonEmptyList gopurs_runtime.Value
var once_foldableNonEmptyList sync.Once
func Get_foldableNonEmptyList() gopurs_runtime.Value {
	once_foldableNonEmptyList.Do(func() {
		cache_foldableNonEmptyList = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_foldableList(), "foldMap"), dictMonoid_0)
_ = foldMap1_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply(f_2, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), gopurs_runtime.Apply2(foldMap1_1_0, f_2, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))
})
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), f_0, gopurs_runtime.Apply2(f_0, b_1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), f_0, b_1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1))
}))
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
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableWithIndexList(), "foldlWithIndex"), gopurs_runtime.Func2(func(i_3 gopurs_runtime.Value, acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), acc_4)
_ = __local_var_5_1
__local_var_6_2 := gopurs_runtime.Apply(f_2, i_3)
_ = __local_var_6_2
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_1, gopurs_runtime.Apply(__local_var_6_2, x_7))
})
}), mempty_1_0)
})
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, acc_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_4 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_3_4
__local_var_4_5 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_4_5
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Int((__local_var_4_5.IntVal) + (1)), gopurs_runtime.Apply3(f_0, __local_var_4_5, __local_var_3_4, a_5)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Int(0), acc_1})})
_ = __local_var_2_3
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(__local_var_2_3, x_3).UnsafePtr).V1
})
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_6 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_7 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_4_7
__local_var_5_8 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0
_ = __local_var_5_8
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Int((__local_var_5_8.IntVal) + (1)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{a_6, (*Constructor_Cons[gopurs_runtime.Value])(__local_var_4_7.UnsafePtr)})}})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil}})}, xs_2)
_ = v_3_6
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_9 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_9
__local_var_6_10 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0
_ = __local_var_6_10
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Int((__local_var_6_10.IntVal) - (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((__local_var_6_10.IntVal) - (1)), a_7, __local_var_5_9)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3_6.UnsafePtr).V0, b_1})}, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3_6.UnsafePtr).V1).UnsafePtr).V1
}))
	})
	return cache_foldableWithIndexList
}

var cache_foldableWithIndexNonEmpty gopurs_runtime.Value
var once_foldableWithIndexNonEmpty sync.Once
func Get_foldableWithIndexNonEmpty() gopurs_runtime.Value {
	once_foldableWithIndexNonEmpty.Do(func() {
		cache_foldableWithIndexNonEmpty = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_foldableWithIndexNonEmpty(), Get_foldableWithIndexList())
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
foldMapWithIndex1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_foldableWithIndexNonEmpty(), "foldMapWithIndex"), dictMonoid_0)
_ = foldMapWithIndex1_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(foldMapWithIndex1_1_0, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = gopurs_runtime.Int((1) + ((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_4.UnsafePtr).V0.IntVal))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Apply(f_2, __t1)
}), v_3)
})
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableWithIndexNonEmpty(), "foldlWithIndex"), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t2 = gopurs_runtime.Int((1) + ((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0.IntVal))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Apply(f_0, __t2)
}), b_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableWithIndexNonEmpty(), "foldrWithIndex"), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Int(0)
goto end_branch_3
} else {

}
}
{
if (x_3.Type == 9 && x_3.IntVal == 930809136 && x_3.UnsafePtr != nil) {
__t3 = gopurs_runtime.Int((1) + ((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_3.UnsafePtr).V0.IntVal))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Apply(f_0, __t3)
}), b_1, v_2)
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableWithIndexList(), "foldrWithIndex"), gopurs_runtime.Func3(func(i_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value, acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{gopurs_runtime.Apply2(f_0, i_1, x_2), (*Constructor_Cons[gopurs_runtime.Value])(acc_3.UnsafePtr)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil})
}))
	})
	return cache_functorWithIndexList
}

var cache_mapWithIndex gopurs_runtime.Value
var once_mapWithIndex sync.Once
func Get_mapWithIndex() gopurs_runtime.Value {
	once_mapWithIndex.Do(func() {
		cache_mapWithIndex = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapWithIndex(f_0_box, v_1_box)
})
	})
	return cache_mapWithIndex
}

var cache_functorWithIndexNonEmptyList gopurs_runtime.Value
var once_functorWithIndexNonEmptyList sync.Once
func Get_functorWithIndexNonEmptyList() gopurs_runtime.Value {
	once_functorWithIndexNonEmptyList.Do(func() {
		cache_functorWithIndexNonEmptyList = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorNonEmptyList()
}), gopurs_runtime.Func2(func(fn_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapWithIndex(gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t0 = gopurs_runtime.Int((1) + ((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_2.UnsafePtr).V0.IntVal))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Apply(fn_0, __t0)
}), v_1)
}))
	})
	return cache_functorWithIndexNonEmptyList
}

var cache_semigroupList gopurs_runtime.Value
var once_semigroupList sync.Once
func Get_semigroupList() gopurs_runtime.Value {
	once_semigroupList.Do(func() {
		cache_semigroupList = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(xs_0 gopurs_runtime.Value, ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), Get_Cons(), ys_1, xs_0)
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
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil})
	})
	return cache_monoidList
}

var cache_semigroupNonEmptyList gopurs_runtime.Value
var once_semigroupNonEmptyList sync.Once
func Get_semigroupNonEmptyList() gopurs_runtime.Value {
	once_semigroupNonEmptyList.Do(func() {
		cache_semigroupNonEmptyList = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, as_prime_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{(*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(as_prime_1.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(as_prime_1.UnsafePtr).V1.UnsafePtr)})})})}
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableList(), "traverse"), dictApplicative_0, Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{a_4, (*Constructor_Cons[gopurs_runtime.Value])(b_3.UnsafePtr)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil}))
_ = __local_var_3_1
__local_var_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func2(func(acc_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func2(func(b_6 gopurs_runtime.Value, a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{a_7, (*Constructor_Cons[gopurs_runtime.Value])(b_6.UnsafePtr)})}
}), acc_4), gopurs_runtime.Apply(f_2, x_5))
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil}))
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(__local_var_4_2, x_5))
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
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{a_4, (*Constructor_Cons[gopurs_runtime.Value])(b_3.UnsafePtr)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil}))
_ = __local_var_3_1
__local_var_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableWithIndexList(), "foldlWithIndex"), gopurs_runtime.Func2(func(i_4 gopurs_runtime.Value, acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_3 := gopurs_runtime.Apply(f_2, i_4)
_ = __local_var_6_3
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func2(func(b_8 gopurs_runtime.Value, a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{a_9, (*Constructor_Cons[gopurs_runtime.Value])(b_8.UnsafePtr)})}
}), acc_5), gopurs_runtime.Apply(__local_var_6_3, x_7))
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil}))
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(__local_var_4_2, x_5))
})
})
}))
	})
	return cache_traversableWithIndexList
}

var cache_traverseWithIndex gopurs_runtime.Value
var once_traverseWithIndex sync.Once
func Get_traverseWithIndex() gopurs_runtime.Value {
	once_traverseWithIndex.Do(func() {
		cache_traverseWithIndex = gopurs_runtime.RecordGet(gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_traversableWithIndexNonEmpty(), Get_traversableWithIndexList()), "traverseWithIndex")
	})
	return cache_traverseWithIndex
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
traverseWithIndex1_1_0 := gopurs_runtime.Apply(Get_traverseWithIndex(), dictApplicative_0)
_ = traverseWithIndex1_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_NonEmptyList(), gopurs_runtime.Apply2(traverseWithIndex1_1_0, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = gopurs_runtime.Int((1) + ((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_4.UnsafePtr).V0.IntVal))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Apply(f_2, __t1)
}), v_3))
})
}))
	})
	return cache_traversableWithIndexNonEmptyList
}

var cache_unfoldable1List gopurs_runtime.Value
var once_unfoldable1List sync.Once
func Get_unfoldable1List() gopurs_runtime.Value {
	once_unfoldable1List.Do(func() {
		cache_unfoldable1List = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(source_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var source_3_loop gopurs_runtime.Value = source_3_loop_val
var memo_4_loop gopurs_runtime.Value = memo_4_loop_val
go__2_0:
for {
if false { continue go__2_0 }
var source_3 gopurs_runtime.Value = source_3_loop
_ = source_3
var memo_4 gopurs_runtime.Value = memo_4_loop
_ = memo_4
v_5_1 := gopurs_runtime.Apply(f_0, source_3)
_ = v_5_1
var __t2 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V1
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 930809136 && __t_tag_3.UnsafePtr != nil) {
source_3_loop = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V1.UnsafePtr).V0
memo_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])(memo_4.UnsafePtr)})}
continue go__2_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V1
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 930809136 && __t_tag_4.UnsafePtr == nil) {
__t2 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func2(func(b_6 gopurs_runtime.Value, a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{a_7, (*Constructor_Cons[gopurs_runtime.Value])(b_6.UnsafePtr)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5_1.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])(memo_4.UnsafePtr)})})
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
return gopurs_runtime.Apply2(go__2_0, b_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil})
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
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(source_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var source_3_loop gopurs_runtime.Value = source_3_loop_val
var memo_4_loop gopurs_runtime.Value = memo_4_loop_val
go__2_0:
for {
if false { continue go__2_0 }
var source_3 gopurs_runtime.Value = source_3_loop
_ = source_3
var memo_4 gopurs_runtime.Value = memo_4_loop
_ = memo_4
v_5_1 := gopurs_runtime.Apply(f_0, source_3)
_ = v_5_1
var __t2 gopurs_runtime.Value
{
if (v_5_1.Type == 9 && v_5_1.IntVal == 930809136 && v_5_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func2(func(b_6 gopurs_runtime.Value, a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{a_7, (*Constructor_Cons[gopurs_runtime.Value])(b_6.UnsafePtr)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil}, memo_4)
goto end_branch_2
} else {

}
}
{
if (v_5_1.Type == 9 && v_5_1.IntVal == 930809136 && v_5_1.UnsafePtr != nil) {
source_3_loop = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_5_1.UnsafePtr).V0.UnsafePtr).V1
memo_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_5_1.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])(memo_4.UnsafePtr)})}
continue go__2_0
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
return gopurs_runtime.Apply2(go__2_0, b_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil})
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
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply(f_0, v_1), gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{a_2, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.RecordGet(v1_3, "acc").UnsafePtr)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{a_2, gopurs_runtime.RecordGet(v1_3, "acc")})}), (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.RecordGet(v1_3, "val").UnsafePtr)})})
}), gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil}), (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1), "val")})}
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
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 1358893437 && v1_1.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 1358893437 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{gopurs_runtime.Apply(v_0, v1_1), (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldr"), gopurs_runtime.Func2(func(a_prime_2 gopurs_runtime.Value, v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{a_prime_2, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.RecordGet(v2_3, "acc").UnsafePtr)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{a_prime_2, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.RecordGet(v2_3, "acc").UnsafePtr)})}), (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.RecordGet(v2_3, "val").UnsafePtr)})})
}), gopurs_runtime.RecordDict2("acc", "val", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_1.UnsafePtr).V1)}), "val").UnsafePtr)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
	})
	return cache_extendList
}

var cache_eq1List gopurs_runtime.Value
var once_eq1List sync.Once
func Get_eq1List() gopurs_runtime.Value {
	once_eq1List.Do(func() {
		cache_eq1List = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func3(func(dictEq_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value, ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
_ = go__3_0
go__3_0 = gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if ((v2_6.IntVal) != (0)) != (true) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t1 = gopurs_runtime.Bool(((v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil)) && ((v2_6.IntVal) != (0)))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(((v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil)) && (((v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil)) && ((gopurs_runtime.Apply3(go__3_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), v2_6, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0))).IntVal) != (0))))
}
end_branch_1:
return __t1
})
return gopurs_runtime.Apply3(go__3_0, xs_1, ys_2, gopurs_runtime.Bool(true))
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
}), gopurs_runtime.Func3(func(dictOrd_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value, ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
go__3_0 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__3_0:
for {
if false { continue go__3_0 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
var __t2 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil)) {
v2_6_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0)
_ = v2_6_3
var __t4 gopurs_runtime.Value
{
if (v2_6_3.Type == 9 && v2_6_3.IntVal == 902936544) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
continue go__3_0
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
__t4 = v2_6_3
}
end_branch_4:
__t1 = __t4
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
return gopurs_runtime.Apply2(go__3_0, xs_1, ys_2)
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
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorList(), "map"), (*Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V0, v1_1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V1)}, v1_1))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
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
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyList(), "apply"), (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{(*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil}.UnsafePtr)})}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{(*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1.UnsafePtr)})}, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1))})}
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
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), gopurs_runtime.Apply(v1_1, (*Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V1)}, v1_1))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
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
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
v1_2_0 := gopurs_runtime.Apply(f_1, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0)
_ = v1_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2_0.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_semigroupList(), "append"), (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2_0.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_bindList(), "bind"), (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(f_1, x_3)
_ = __local_var_4_1
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{(*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1.UnsafePtr)})}
})))})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{a_0, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil}.UnsafePtr)})}
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
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil})
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
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{x_0, gopurs_runtime.RecordGet(Get_plusList(), "empty")})}
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversable1NonEmptyList(), "traverse1"), dictApply_0, Get_identity())
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})
_ = Functor0_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_1_0, "map"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func2(func(b_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{a_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{(*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_5.UnsafePtr).V1.UnsafePtr)})}})}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_applicativeNonEmptyList(), "pure"), (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0), (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1)
}), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableList(), "foldl"), gopurs_runtime.Func2(func(acc_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func2(func(b_6 gopurs_runtime.Value, a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{a_7, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{(*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_6.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(b_6.UnsafePtr).V1.UnsafePtr)})}})}
}), acc_4), gopurs_runtime.Apply(f_2, x_5))
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_1_0, "map"), gopurs_runtime.RecordGet(Get_applicativeNonEmptyList(), "pure"), gopurs_runtime.Apply(f_2, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)), (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))
})
}))
	})
	return cache_traversable1NonEmptyList
}

type Constructor_Nil[T_a any] struct {
	
}


type Constructor_Cons[T_a any] struct {
	V0 T_a
	V1 *Constructor_Cons[gopurs_runtime.Value]
}


func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_NonEmptyList(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_toList(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{(*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1.UnsafePtr)})}
}

func Call_nelCons(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{a_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{(*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1.UnsafePtr)})}})}
}

func Call_listMap(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var chunkedRevMap_1_0 gopurs_runtime.Value
chunkedRevMap_1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
chunkedRevMap_1_0:
for {
if false { continue chunkedRevMap_1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t15 gopurs_runtime.Value
{
var __t_and_19 bool = false
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {

var __t_tag_16 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
var __t_and_18 bool = false
if (__t_tag_16.Type == 9 && __t_tag_16.IntVal == 1358893437 && __t_tag_16.UnsafePtr != nil) {

var __t_tag_17 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}.UnsafePtr).V1)}
__t_and_18 = (__t_tag_17.Type == 9 && __t_tag_17.IntVal == 1358893437 && __t_tag_17.UnsafePtr != nil)
}
__t_and_19 = __t_and_18
}
if __t_and_19 {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{v1_3, (*Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr)})}
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr).V1)}
continue chunkedRevMap_1_0
__t15 = gopurs_runtime.Value{}
goto end_branch_15
} else {

}
}
{
var reverseUnrolledMap_4_1 gopurs_runtime.Value
reverseUnrolledMap_4_1 = gopurs_runtime.Func(func(v2_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v2_5_loop gopurs_runtime.Value = v2_5_loop_val
var v3_6_loop gopurs_runtime.Value = v3_6_loop_val
reverseUnrolledMap_4_1:
for {
if false { continue reverseUnrolledMap_4_1 }
var v2_5 gopurs_runtime.Value = v2_5_loop
_ = v2_5
var v3_6 gopurs_runtime.Value = v3_6_loop
_ = v3_6
var __t2 gopurs_runtime.Value
{
var __t_and_8 bool = false
if (v2_5.Type == 9 && v2_5.IntVal == 1358893437 && v2_5.UnsafePtr != nil) {

var __t_tag_3 gopurs_runtime.Value = (*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0
var __t_and_7 bool = false
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 1358893437 && __t_tag_3.UnsafePtr != nil) {

var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])((*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}
var __t_and_6 bool = false
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 1358893437 && __t_tag_4.UnsafePtr != nil) {

var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])((*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V1)}
__t_and_6 = (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 1358893437 && __t_tag_5.UnsafePtr != nil)
}
__t_and_7 = __t_and_6
}
__t_and_8 = __t_and_7
}
if __t_and_8 {
v2_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V1)}
v3_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])((*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V0), (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])((*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V0), (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])((*Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr).V0), (*Constructor_Cons[gopurs_runtime.Value])(v3_6.UnsafePtr)})}.UnsafePtr)})}.UnsafePtr)})}
continue reverseUnrolledMap_4_1
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
__t2 = v3_6
}
end_branch_2:
return __t2
}
}()
})
})
var __t9 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
var __t10 gopurs_runtime.Value
{
var __t_tag_11 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
if (__t_tag_11.Type == 9 && __t_tag_11.IntVal == 1358893437 && __t_tag_11.UnsafePtr != nil) {
var __t12 gopurs_runtime.Value
{
var __t_tag_13 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}.UnsafePtr).V1)}
if (__t_tag_13.Type == 9 && __t_tag_13.IntVal == 1358893437 && __t_tag_13.UnsafePtr == nil) {
__t12 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0), (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}.UnsafePtr).V0), (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil}.UnsafePtr)})}.UnsafePtr)})}
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil}
}
end_branch_12:
__t10 = __t12
goto end_branch_10
} else {

}
}
{
var __t_tag_14 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
if (__t_tag_14.Type == 9 && __t_tag_14.IntVal == 1358893437 && __t_tag_14.UnsafePtr == nil) {
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Cons[gopurs_runtime.Value]{gopurs_runtime.Apply(f_0, (*Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0), (*Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil}.UnsafePtr)})}
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil}
}
end_branch_10:
__t9 = __t10
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil}
}
end_branch_9:
__t15 = gopurs_runtime.Apply2(reverseUnrolledMap_4_1, v_2, __t9)
}
end_branch_15:
return __t15
}
}()
})
})
return gopurs_runtime.Apply(chunkedRevMap_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil})
}

func Call_mapWithIndex(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorWithIndexList(), "mapWithIndex"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{x_2})})
}), (*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1)})}
}

func Call_showList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
show_1_0 := gopurs_runtime.RecordGet(dictShow_0, "show")
_ = show_1_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Str("Nil")
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("("), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply2(Get_intercalate(), gopurs_runtime.Str(" : "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorList(), "map"), show_1_0, v_2)), gopurs_runtime.Str(" : Nil)")))
}
end_branch_1:
return __t1
}))
}

func Call_showNonEmptyList(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(NonEmptyList "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(pkg_Data_NonEmpty.Get_showNonEmpty(), dictShow_0, gopurs_runtime.Apply(Get_showList(), dictShow_0)), "show"), v_1), gopurs_runtime.Str(")")))
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



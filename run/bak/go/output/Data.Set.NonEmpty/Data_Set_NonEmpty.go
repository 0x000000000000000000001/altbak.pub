package Data_Set_NonEmpty

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Map_Internal "gopurs/output/Data.Map.Internal"
	pkg_Data_Function "gopurs/output/Data.Function"
	pkg_Partial "gopurs/output/Partial"
	pkg_Data_Unfoldable1 "gopurs/output/Data.Unfoldable1"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_Set "gopurs/output/Data.Set"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Eq "gopurs/output/Data.Eq"
)

var unionSet gopurs_runtime.Value
var once_unionSet sync.Once
func Get_unionSet() gopurs_runtime.Value {
	once_unionSet.Do(func() {
		unionSet = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), compare_1_0, pkg_Data_Function.Get_const_(), m1_2, m2_3)
})
})
	})
	return unionSet
}

var toUnfoldable1 gopurs_runtime.Value
var once_toUnfoldable1 sync.Once
func Get_toUnfoldable1() gopurs_runtime.Value {
	once_toUnfoldable1.Do(func() {
		toUnfoldable1 = gopurs_runtime.Func(func(dictUnfoldable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
stepNext_1_0 := gopurs_runtime.Apply3(pkg_Data_Map_Internal.Get_stepWith(), pkg_Data_Map_Internal.Get_iterMapL(), gopurs_runtime.Func3(func(k_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, next_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Tuple", k_1, next_3))
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Nothing")
}))
_ = stepNext_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable1_0, "unfoldr1"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.ConstructorGet(v_2, 0), gopurs_runtime.Apply(stepNext_1_0, gopurs_runtime.ConstructorGet(v_2, 1)))
}))
_ = __local_var_2_1
__local_var_3_2 := gopurs_runtime.Apply3(pkg_Data_Map_Internal.Get_stepWith(), pkg_Data_Map_Internal.Get_iterMapL(), gopurs_runtime.Func3(func(k_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value, next_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", k_3, next_5)
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str("toUnfoldable1: impossible"))
}))
_ = __local_var_3_2
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_1, gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Constructor2("IterNode", x_4, gopurs_runtime.Constructor0("IterLeaf"))))
})
})
	})
	return toUnfoldable1
}

var toUnfoldable11 gopurs_runtime.Value
var once_toUnfoldable11 sync.Once
func Get_toUnfoldable11() gopurs_runtime.Value {
	once_toUnfoldable11.Do(func() {
		toUnfoldable11 = gopurs_runtime.Apply(Get_toUnfoldable1(), pkg_Data_Unfoldable1.Get_unfoldable1Array())
	})
	return toUnfoldable11
}

var toUnfoldable12 gopurs_runtime.Value
var once_toUnfoldable12 sync.Once
func Get_toUnfoldable12() gopurs_runtime.Value {
	once_toUnfoldable12.Do(func() {
		toUnfoldable12 = gopurs_runtime.Apply(Get_toUnfoldable1(), pkg_Data_List_Types.Get_unfoldable1NonEmptyList())
	})
	return toUnfoldable12
}

var toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(xs_1.StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(xs_1.StrVal == "Cons")).IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Tuple", gopurs_runtime.ConstructorGet(xs_1, 0), gopurs_runtime.ConstructorGet(xs_1, 1)))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_2 gopurs_runtime.Value
_ = go__3_2
go__3_2 = gopurs_runtime.Func2(func(m_prime_4 gopurs_runtime.Value, z_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(m_prime_4.StrVal == "Leaf")).IntVal != 0 {
__t3 = z_prime_5
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(m_prime_4.StrVal == "Node")).IntVal != 0 {
__t3 = gopurs_runtime.UncurriedApp2(go__3_2, gopurs_runtime.ConstructorGet(m_prime_4, 4), gopurs_runtime.Constructor2("Cons", gopurs_runtime.ConstructorGet(m_prime_4, 2), gopurs_runtime.UncurriedApp2(go__3_2, gopurs_runtime.ConstructorGet(m_prime_4, 5), z_prime_5)))
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
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.UncurriedApp2(go__3_2, x_2, gopurs_runtime.Constructor0("Nil")))
})
})
	})
	return toUnfoldable
}

var toSet gopurs_runtime.Value
var once_toSet sync.Once
func Get_toSet() gopurs_runtime.Value {
	once_toSet.Do(func() {
		toSet = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
})
	})
	return toSet
}

var subset gopurs_runtime.Value
var once_subset sync.Once
func Get_subset() gopurs_runtime.Value {
	once_subset.Do(func() {
		subset = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func2(func(s1_2 gopurs_runtime.Value, s2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.UncurriedApp3(pkg_Data_Map_Internal.Get_unsafeDifference(), compare_1_0, s1_2, s2_3).StrVal == "Leaf")
})
})
	})
	return subset
}

var size gopurs_runtime.Value
var once_size sync.Once
func Get_size() gopurs_runtime.Value {
	once_size.Do(func() {
		size = pkg_Data_Map_Internal.Get_size()
	})
	return size
}

var singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		singleton = pkg_Data_Set.Get_singleton()
	})
	return singleton
}

var showNonEmptySet gopurs_runtime.Value
var once_showNonEmptySet sync.Once
func Get_showNonEmptySet() gopurs_runtime.Value {
	once_showNonEmptySet.Do(func() {
		showNonEmptySet = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str("(fromFoldable1 (NonEmptyArray ").StrVal + gopurs_runtime.Apply2(pkg_Data_Show.Get_showArrayImpl(), gopurs_runtime.RecordGet(dictShow_0, "show"), gopurs_runtime.Apply(Get_toUnfoldable11(), s_1)).StrVal).StrVal + gopurs_runtime.Str("))").StrVal)
}))
})
	})
	return showNonEmptySet
}

var semigroupNonEmptySet gopurs_runtime.Value
var once_semigroupNonEmptySet sync.Once
func Get_semigroupNonEmptySet() gopurs_runtime.Value {
	once_semigroupNonEmptySet.Do(func() {
		semigroupNonEmptySet = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), compare_1_0, pkg_Data_Function.Get_const_(), m1_2, m2_3)
}))
})
	})
	return semigroupNonEmptySet
}

var properSubset gopurs_runtime.Value
var once_properSubset sync.Once
func Get_properSubset() gopurs_runtime.Value {
	once_properSubset.Do(func() {
		properSubset = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Set.Get_properSubset(), dictOrd_0)
})
	})
	return properSubset
}

var ordNonEmptySet gopurs_runtime.Value
var once_ordNonEmptySet sync.Once
func Get_ordNonEmptySet() gopurs_runtime.Value {
	once_ordNonEmptySet.Do(func() {
		ordNonEmptySet = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Set.Get_ordSet(), dictOrd_0)
})
	})
	return ordNonEmptySet
}

var ord1NonEmptySet gopurs_runtime.Value
var once_ord1NonEmptySet sync.Once
func Get_ord1NonEmptySet() gopurs_runtime.Value {
	once_ord1NonEmptySet.Do(func() {
		ord1NonEmptySet = pkg_Data_Set.Get_ord1Set()
	})
	return ord1NonEmptySet
}

var min gopurs_runtime.Value
var once_min sync.Once
func Get_min() gopurs_runtime.Value {
	once_min.Do(func() {
		min = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_findMin(), v_0)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_1_0.StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.RecordGet(gopurs_runtime.ConstructorGet(__local_var_1_0, 0), "key")
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
	})
	return min
}

var member gopurs_runtime.Value
var once_member sync.Once
func Get_member() gopurs_runtime.Value {
	once_member.Do(func() {
		member = gopurs_runtime.Func2(func(dictOrd_0 gopurs_runtime.Value, k_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_0:
for {
if false { continue go__2_0 }
var v_3 = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_3.StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_3.StrVal == "Node")).IntVal != 0 {
v1_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_1, gopurs_runtime.ConstructorGet(v_3, 2))
_ = v1_4_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_4_2.StrVal == "LT")).IntVal != 0 {
v_3_loop = gopurs_runtime.ConstructorGet(v_3, 4)
continue go__2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v1_4_2.StrVal == "GT")).IntVal != 0 {
v_3_loop = gopurs_runtime.ConstructorGet(v_3, 5)
continue go__2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v1_4_2.StrVal == "EQ")).IntVal != 0 {
__t3 = gopurs_runtime.Bool(true)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t1 = __t3
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
return go__2_0
})
	})
	return member
}

var max gopurs_runtime.Value
var once_max sync.Once
func Get_max() gopurs_runtime.Value {
	once_max.Do(func() {
		max = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_findMax(), v_0)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_1_0.StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.RecordGet(gopurs_runtime.ConstructorGet(__local_var_1_0, 0), "key")
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
	})
	return max
}

var mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		mapMaybe = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Set.Get_mapMaybe(), dictOrd_0)
})
	})
	return mapMaybe
}

var map_ gopurs_runtime.Value
var once_map_ sync.Once
func Get_map_() gopurs_runtime.Value {
	once_map_.Do(func() {
		map_ = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Set.Get_map_(), dictOrd_0)
})
	})
	return map_
}

var insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		insert = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Set.Get_insert(), dictOrd_0)
})
	})
	return insert
}

var fromSet gopurs_runtime.Value
var once_fromSet sync.Once
func Get_fromSet() gopurs_runtime.Value {
	once_fromSet.Do(func() {
		fromSet = gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(s_0.StrVal == "Leaf")).IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor1("Just", s_0)
}
end_branch_0:
return __t0
})
	})
	return fromSet
}

var intersection gopurs_runtime.Value
var once_intersection sync.Once
func Get_intersection() gopurs_runtime.Value {
	once_intersection.Do(func() {
		intersection = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeIntersectionWith(), compare_1_0, pkg_Data_Function.Get_const_(), v_2, v1_3)
_ = __local_var_4_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_4_1.StrVal == "Leaf")).IntVal != 0 {
__t2 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Constructor1("Just", __local_var_4_1)
}
end_branch_2:
return __t2
})
})
	})
	return intersection
}

var fromFoldable1 gopurs_runtime.Value
var once_fromFoldable1 sync.Once
func Get_fromFoldable1() gopurs_runtime.Value {
	once_fromFoldable1.Do(func() {
		fromFoldable1 = gopurs_runtime.Func2(func(dictFoldable1_0 gopurs_runtime.Value, dictOrd_1 gopurs_runtime.Value) gopurs_runtime.Value {
compare_2_0 := gopurs_runtime.RecordGet(dictOrd_1, "compare")
_ = compare_2_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_0, "foldMap1"), gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(m1_3 gopurs_runtime.Value, m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), compare_2_0, pkg_Data_Function.Get_const_(), m1_3, m2_4)
})), pkg_Data_Set.Get_singleton())
})
	})
	return fromFoldable1
}

var fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		fromFoldable = gopurs_runtime.Func2(func(dictFoldable_0 gopurs_runtime.Value, dictOrd_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(m_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(pkg_Data_Map_Internal.Get_insert(), dictOrd_1, a_3, pkg_Data_Unit.Get_unit(), m_2)
}), gopurs_runtime.Constructor0("Leaf"))
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(__local_var_2_0, x_3)
_ = __local_var_4_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_4_1.StrVal == "Leaf")).IntVal != 0 {
__t2 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Constructor1("Just", __local_var_4_1)
}
end_branch_2:
return __t2
})
})
	})
	return fromFoldable
}

var foldableNonEmptySet gopurs_runtime.Value
var once_foldableNonEmptySet sync.Once
func Get_foldableNonEmptySet() gopurs_runtime.Value {
	once_foldableNonEmptySet.Do(func() {
		foldableNonEmptySet = pkg_Data_Set.Get_foldableSet()
	})
	return foldableNonEmptySet
}

var foldable1NonEmptySet gopurs_runtime.Value
var once_foldable1NonEmptySet sync.Once
func Get_foldable1NonEmptySet() gopurs_runtime.Value {
	once_foldable1NonEmptySet.Do(func() {
		foldable1NonEmptySet = gopurs_runtime.RecordDict4("foldMap1", "foldr1", "foldl1", "Foldable0", gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap11_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldable1NonEmptyList(), "foldMap1"), dictSemigroup_0)
_ = foldMap11_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(foldMap11_1_0, f_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(Get_toUnfoldable12(), x_4))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldable1NonEmptyList(), "foldr1"), f_0)
_ = __local_var_1_2
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_2, gopurs_runtime.Apply(Get_toUnfoldable12(), x_2))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldable1NonEmptyList(), "foldl1"), f_0)
_ = __local_var_1_3
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_3, gopurs_runtime.Apply(Get_toUnfoldable12(), x_2))
})
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Set.Get_foldableSet()
}))
	})
	return foldable1NonEmptySet
}

var filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		filter = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Set.Get_filter(), dictOrd_0)
})
	})
	return filter
}

var eqNonEmptySet gopurs_runtime.Value
var once_eqNonEmptySet sync.Once
func Get_eqNonEmptySet() gopurs_runtime.Value {
	once_eqNonEmptySet.Do(func() {
		eqNonEmptySet = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(pkg_Data_Map_Internal.Get_eqMap(), dictEq_0, pkg_Data_Eq.Get_eqUnit()), "eq"), v_1, v1_2)
}))
})
	})
	return eqNonEmptySet
}

var eq1NonEmptySet gopurs_runtime.Value
var once_eq1NonEmptySet sync.Once
func Get_eq1NonEmptySet() gopurs_runtime.Value {
	once_eq1NonEmptySet.Do(func() {
		eq1NonEmptySet = pkg_Data_Set.Get_eq1Set()
	})
	return eq1NonEmptySet
}

var difference gopurs_runtime.Value
var once_difference sync.Once
func Get_difference() gopurs_runtime.Value {
	once_difference.Do(func() {
		difference = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.UncurriedApp3(pkg_Data_Map_Internal.Get_unsafeDifference(), compare_1_0, v_2, v1_3)
_ = __local_var_4_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_4_1.StrVal == "Leaf")).IntVal != 0 {
__t2 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Constructor1("Just", __local_var_4_1)
}
end_branch_2:
return __t2
})
})
	})
	return difference
}

var delete_ gopurs_runtime.Value
var once_delete_ sync.Once
func Get_delete_() gopurs_runtime.Value {
	once_delete_.Do(func() {
		delete_ = gopurs_runtime.Func3(func(dictOrd_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply3(pkg_Data_Map_Internal.Get_delete_(), dictOrd_0, a_1, v_2)
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_3_0.StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor1("Just", __local_var_3_0)
}
end_branch_1:
return __t1
})
	})
	return delete_
}

var cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cons = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Set.Get_insert(), dictOrd_0)
})
	})
	return cons
}



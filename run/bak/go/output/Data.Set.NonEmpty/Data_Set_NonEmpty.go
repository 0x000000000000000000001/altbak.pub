package Data_Set_NonEmpty

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	pkg_Data_Set "gopurs/output/Data.Set"
	pkg_Data_Map_Internal "gopurs/output/Data.Map.Internal"
	pkg_Data_Function_Uncurried "gopurs/output/Data.Function.Uncurried"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
	pkg_Partial "gopurs/output/Partial"
	pkg_Data_Unfoldable1 "gopurs/output/Data.Unfoldable1"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Array_NonEmpty_Internal "gopurs/output/Data.Array.NonEmpty.Internal"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Eq "gopurs/output/Data.Eq"
)

var unionSet gopurs_runtime.Value
var once_unionSet sync.Once
func Get_unionSet() gopurs_runtime.Value {
	once_unionSet.Do(func() {
		unionSet = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(pkg_Data_Set.Get_union(), dictOrd_0))
})
	})
	return unionSet
}

var toUnfoldable1 gopurs_runtime.Value
var once_toUnfoldable1 sync.Once
func Get_toUnfoldable1() gopurs_runtime.Value {
	once_toUnfoldable1.Do(func() {
		toUnfoldable1 = gopurs_runtime.Func(func(dictUnfoldable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
stepNext_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_stepWith(), pkg_Data_Map_Internal.Get_iterMapL()), gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_mkFn3(), gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(next_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": k_1, "value1": next_3})})
})
})
}))), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}))
__local_var_2_1 := gopurs_runtime.Apply(dictUnfoldable1_0.PtrVal.(map[string]gopurs_runtime.Value)["unfoldr1"], gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(stepNext_1_0, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
}))
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_stepWith(), pkg_Data_Map_Internal.Get_iterMapL()), gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_mkFn3(), gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(next_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": k_3, "value1": next_5})
})
})
}))), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str("toUnfoldable1: impossible"))
}))
}))
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_1, gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterNode"), "value0": gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), x_4), "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("IterLeaf")})})))
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
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(pkg_Data_Set.Get_toUnfoldable(), dictUnfoldable_0))
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
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(pkg_Data_Set.Get_subset(), dictOrd_0))
})
	})
	return subset
}

var size gopurs_runtime.Value
var once_size sync.Once
func Get_size() gopurs_runtime.Value {
	once_size.Do(func() {
		size = gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), pkg_Data_Set.Get_size())
	})
	return size
}

var singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		singleton = gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), pkg_Data_Set.Get_singleton())
	})
	return singleton
}

var showNonEmptySet gopurs_runtime.Value
var once_showNonEmptySet sync.Once
func Get_showNonEmptySet() gopurs_runtime.Value {
	once_showNonEmptySet.Do(func() {
		showNonEmptySet = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(fromFoldable1 ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Array_NonEmpty_Internal.Get_showNonEmptyArray(), dictShow_0).PtrVal.(map[string]gopurs_runtime.Value)["show"], gopurs_runtime.Apply(Get_toUnfoldable11(), s_1))), gopurs_runtime.Str(")")))
})})
})
	})
	return showNonEmptySet
}

var semigroupNonEmptySet gopurs_runtime.Value
var once_semigroupNonEmptySet sync.Once
func Get_semigroupNonEmptySet() gopurs_runtime.Value {
	once_semigroupNonEmptySet.Do(func() {
		semigroupNonEmptySet = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Apply(pkg_Data_Set.Get_union(), dictOrd_0)})
})
	})
	return semigroupNonEmptySet
}

var properSubset gopurs_runtime.Value
var once_properSubset sync.Once
func Get_properSubset() gopurs_runtime.Value {
	once_properSubset.Do(func() {
		properSubset = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(pkg_Data_Set.Get_properSubset(), dictOrd_0))
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
return gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_findMin(), v_0)
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = __local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["key"]
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
})
	})
	return min
}

var member gopurs_runtime.Value
var once_member sync.Once
func Get_member() gopurs_runtime.Value {
	once_member.Do(func() {
		member = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(pkg_Data_Set.Get_member(), dictOrd_0))
})
	})
	return member
}

var max gopurs_runtime.Value
var once_max sync.Once
func Get_max() gopurs_runtime.Value {
	once_max.Do(func() {
		max = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_findMax(), v_0)
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = __local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["key"]
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
})
	})
	return max
}

var mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		mapMaybe = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(pkg_Data_Set.Get_mapMaybe(), dictOrd_0))
})
	})
	return mapMaybe
}

var map_ gopurs_runtime.Value
var once_map_ sync.Once
func Get_map_() gopurs_runtime.Value {
	once_map_.Do(func() {
		map_ = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(pkg_Data_Set.Get_map_(), dictOrd_0))
})
	})
	return map_
}

var insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		insert = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(pkg_Data_Set.Get_insert(), dictOrd_0))
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
if (gopurs_runtime.Apply(pkg_Data_Set.Get_isEmpty(), s_0)).IntVal != 0 {
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": s_0})
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
intersection1_1_0 := gopurs_runtime.Apply(pkg_Data_Set.Get_intersection(), dictOrd_0)
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(intersection1_1_0, v_2), v1_3)
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(pkg_Data_Set.Get_isEmpty(), __local_var_4_1)).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __local_var_4_1})
}
end_branch_2:
return __t2
})
})
})
	})
	return intersection
}

var fromFoldable1 gopurs_runtime.Value
var once_fromFoldable1 sync.Once
func Get_fromFoldable1() gopurs_runtime.Value {
	once_fromFoldable1.Do(func() {
		fromFoldable1 = gopurs_runtime.Func(func(dictFoldable1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictOrd_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable1_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap1"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Apply(pkg_Data_Set.Get_union(), dictOrd_1)})), Get_singleton())
})
})
	})
	return fromFoldable1
}

var fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		fromFoldable = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictOrd_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_insert(), dictOrd_1), a_3), pkg_Data_Unit.Get_unit()), m_2)
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")}))
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(__local_var_2_0, x_3)
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(pkg_Data_Set.Get_isEmpty(), __local_var_4_1)).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __local_var_4_1})
}
end_branch_2:
return __t2
})
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
		foldable1NonEmptySet = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldMap1": gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap11_1_0 := gopurs_runtime.Apply(pkg_Data_List_Types.Get_foldable1NonEmptyList().PtrVal.(map[string]gopurs_runtime.Value)["foldMap1"], dictSemigroup_0)
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(foldMap11_1_0, f_2)
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(Get_toUnfoldable12(), x_4))
})
})
}), "foldr1": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_2 := gopurs_runtime.Apply(pkg_Data_List_Types.Get_foldable1NonEmptyList().PtrVal.(map[string]gopurs_runtime.Value)["foldr1"], f_0)
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_2, gopurs_runtime.Apply(Get_toUnfoldable12(), x_2))
})
}), "foldl1": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_3 := gopurs_runtime.Apply(pkg_Data_List_Types.Get_foldable1NonEmptyList().PtrVal.(map[string]gopurs_runtime.Value)["foldl1"], f_0)
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_3, gopurs_runtime.Apply(Get_toUnfoldable12(), x_2))
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Set.Get_foldableSet()
})})
	})
	return foldable1NonEmptySet
}

var filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		filter = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_filterKeys(), dictOrd_0)))
})
	})
	return filter
}

var eqNonEmptySet gopurs_runtime.Value
var once_eqNonEmptySet sync.Once
func Get_eqNonEmptySet() gopurs_runtime.Value {
	once_eqNonEmptySet.Do(func() {
		eqNonEmptySet = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_eqMap(), dictEq_0), pkg_Data_Eq.Get_eqUnit()).PtrVal.(map[string]gopurs_runtime.Value)["eq"], v_1), v1_2)
})
})})
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
compare_1_1 := dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
difference1_1_0 := gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn3(), pkg_Data_Map_Internal.Get_unsafeDifference()), compare_1_1), m1_2), m2_3)
})
}))
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(difference1_1_0, v_2), v1_3)
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(pkg_Data_Set.Get_isEmpty(), __local_var_4_2)).IntVal != 0 {
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __local_var_4_2})
}
end_branch_3:
return __t3
})
})
})
	})
	return difference
}

var delete_ gopurs_runtime.Value
var once_delete_ sync.Once
func Get_delete_() gopurs_runtime.Value {
	once_delete_.Do(func() {
		delete_ = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
delete1_1_0 := gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_delete_(), dictOrd_0))
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(delete1_1_0, a_2), v_3)
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(pkg_Data_Set.Get_isEmpty(), __local_var_4_1)).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __local_var_4_1})
}
end_branch_2:
return __t2
})
})
})
	})
	return delete_
}

var cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cons = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(pkg_Data_Set.Get_insert(), dictOrd_0))
})
	})
	return cons
}



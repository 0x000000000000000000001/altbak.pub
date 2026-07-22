package Data_Set

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	pkg_Data_Function_Uncurried "gopurs/output/Data.Function.Uncurried"
	pkg_Data_Map_Internal "gopurs/output/Data.Map.Internal"
	pkg_Data_Function "gopurs/output/Data.Function"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Control_Category "gopurs/output/Control.Category"
)

var Set gopurs_runtime.Value
var once_Set sync.Once
func Get_Set() gopurs_runtime.Value {
	once_Set.Do(func() {
		Set = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Set
}

var union gopurs_runtime.Value
var once_union sync.Once
func Get_union() gopurs_runtime.Value {
	once_union.Do(func() {
		union = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), pkg_Data_Map_Internal.Get_unsafeUnionWith()), compare_1_0), pkg_Data_Function.Get_const_()), m1_2), m2_3)
})
}))
})
	})
	return union
}

var toggle gopurs_runtime.Value
var once_toggle sync.Once
func Get_toggle() gopurs_runtime.Value {
	once_toggle.Do(func() {
		toggle = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
alter_1_0 := gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_alter(), dictOrd_0)
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(alter_1_0, gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": pkg_Data_Unit.Get_unit()})
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v2_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
})), a_2), v_3)
})
})
})
	})
	return toggle
}

var toMap gopurs_runtime.Value
var once_toMap sync.Once
func Get_toMap() gopurs_runtime.Value {
	once_toMap.Do(func() {
		toMap = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
})
	})
	return toMap
}

var toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictUnfoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["unfoldr"], gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(xs_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(xs_1.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": xs_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": xs_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})})
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
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_keys(), x_2))
})
})
	})
	return toUnfoldable
}

var toUnfoldable1 gopurs_runtime.Value
var once_toUnfoldable1 sync.Once
func Get_toUnfoldable1() gopurs_runtime.Value {
	once_toUnfoldable1.Do(func() {
		toUnfoldable1 = gopurs_runtime.Apply(Get_toUnfoldable(), pkg_Data_Unfoldable.Get_unfoldableArray())
	})
	return toUnfoldable1
}

var size gopurs_runtime.Value
var once_size sync.Once
func Get_size() gopurs_runtime.Value {
	once_size.Do(func() {
		size = gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), pkg_Data_Map_Internal.Get_size())
	})
	return size
}

var singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		singleton = gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Node"), "value0": gopurs_runtime.Int(1), "value1": gopurs_runtime.Int(1), "value2": a_0, "value3": pkg_Data_Unit.Get_unit(), "value4": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")}), "value5": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")})})
})
	})
	return singleton
}

var showSet gopurs_runtime.Value
var once_showSet sync.Once
func Get_showSet() gopurs_runtime.Value {
	once_showSet.Do(func() {
		showSet = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"show": gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Str("(fromFoldable ")), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Semigroup.Get_concatString(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Show.Get_showArrayImpl(), dictShow_0.PtrVal.(map[string]gopurs_runtime.Value)["show"]), gopurs_runtime.Apply(Get_toUnfoldable1(), s_1))), gopurs_runtime.Str(")")))
})})
})
	})
	return showSet
}

var semigroupSet gopurs_runtime.Value
var once_semigroupSet sync.Once
func Get_semigroupSet() gopurs_runtime.Value {
	once_semigroupSet.Do(func() {
		semigroupSet = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Apply(Get_union(), dictOrd_0)})
})
	})
	return semigroupSet
}

var member gopurs_runtime.Value
var once_member sync.Once
func Get_member() gopurs_runtime.Value {
	once_member.Do(func() {
		member = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
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
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Node")).IntVal != 0 {
v1_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"], k_1), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value2"])
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "LT")).IntVal != 0 {
v_3_loop = v_3.PtrVal.(map[string]gopurs_runtime.Value)["value4"]
continue go__2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v1_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "GT")).IntVal != 0 {
v_3_loop = v_3.PtrVal.(map[string]gopurs_runtime.Value)["value5"]
continue go__2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v1_4_2.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "EQ")).IntVal != 0 {
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
}))
})
	})
	return member
}

var isEmpty gopurs_runtime.Value
var once_isEmpty sync.Once
func Get_isEmpty() gopurs_runtime.Value {
	once_isEmpty.Do(func() {
		isEmpty = gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), pkg_Data_Map_Internal.Get_isEmpty())
	})
	return isEmpty
}

var intersection gopurs_runtime.Value
var once_intersection sync.Once
func Get_intersection() gopurs_runtime.Value {
	once_intersection.Do(func() {
		intersection = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn4(), pkg_Data_Map_Internal.Get_unsafeIntersectionWith()), compare_1_0), pkg_Data_Function.Get_const_()), m1_2), m2_3)
})
}))
})
	})
	return intersection
}

var insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		insert = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_insert(), dictOrd_0), a_1), pkg_Data_Unit.Get_unit()), v_2)
})
})
})
	})
	return insert
}

var fromMap gopurs_runtime.Value
var once_fromMap sync.Once
func Get_fromMap() gopurs_runtime.Value {
	once_fromMap.Do(func() {
		fromMap = Get_Set()
	})
	return fromMap
}

var foldableSet gopurs_runtime.Value
var once_foldableSet sync.Once
func Get_foldableSet() gopurs_runtime.Value {
	once_foldableSet.Do(func() {
		foldableSet = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldMap": gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_1_0 := gopurs_runtime.Apply(pkg_Data_List_Types.Get_foldableList().PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_0)
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(foldMap1_1_0, f_2)
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_keys(), x_4))
})
})
}), "foldl": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_2 gopurs_runtime.Value
go__2_2 = gopurs_runtime.Func(func(b_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_2:
for {
if false { continue go__2_2 }
var b_3 = b_3_loop
_ = b_3
var v_4 = v_4_loop
_ = v_4
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nil")).IntVal != 0 {
__t3 = b_3
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(v_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Cons")).IntVal != 0 {
b_3_loop = gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, b_3), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
v_4_loop = v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
continue go__2_2
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
__local_var_3_4 := gopurs_runtime.Apply(go__2_2, x_1)
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_4, gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_keys(), x_4))
})
})
}), "foldr": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Types.Get_foldableList().PtrVal.(map[string]gopurs_runtime.Value)["foldr"], f_0), x_1)
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_5, gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_keys(), x_3))
})
})
})})
	})
	return foldableSet
}

var findMin gopurs_runtime.Value
var once_findMin sync.Once
func Get_findMin() gopurs_runtime.Value {
	once_findMin.Do(func() {
		findMin = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_findMin(), v_0)
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["key"]})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_1:
return __t1
})
	})
	return findMin
}

var findMax gopurs_runtime.Value
var once_findMax sync.Once
func Get_findMax() gopurs_runtime.Value {
	once_findMax.Do(func() {
		findMax = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_findMax(), v_0)
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Just"), "value0": __local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["key"]})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
}
end_branch_1:
return __t1
})
	})
	return findMax
}

var filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		filter = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_filterKeys(), dictOrd_0))
})
	})
	return filter
}

var eqSet gopurs_runtime.Value
var once_eqSet sync.Once
func Get_eqSet() gopurs_runtime.Value {
	once_eqSet.Do(func() {
		eqSet = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_eqMap(), dictEq_0), pkg_Data_Eq.Get_eqUnit()).PtrVal.(map[string]gopurs_runtime.Value)["eq"], v_1), v1_2)
})
})})
})
	})
	return eqSet
}

var ordSet gopurs_runtime.Value
var once_ordSet sync.Once
func Get_ordSet() gopurs_runtime.Value {
	once_ordSet.Do(func() {
		ordSet = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["Eq0"], gopurs_runtime.Value{})
eqSet1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_eqMap(), __local_var_1_0), pkg_Data_Eq.Get_eqUnit()).PtrVal.(map[string]gopurs_runtime.Value)["eq"], v_2), v1_3)
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare": gopurs_runtime.Func(func(s1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_List_Types.Get_ordList(), dictOrd_0).PtrVal.(map[string]gopurs_runtime.Value)["compare"], gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_keys(), s1_3)), gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_keys(), s2_4))
})
}), "Eq0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eqSet1_2_1
})})
})
	})
	return ordSet
}

var eq1Set gopurs_runtime.Value
var once_eq1Set sync.Once
func Get_eq1Set() gopurs_runtime.Value {
	once_eq1Set.Do(func() {
		eq1Set = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"eq1": gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_eqMap(), dictEq_0), pkg_Data_Eq.Get_eqUnit()).PtrVal.(map[string]gopurs_runtime.Value)["eq"], v_1), v1_2)
})
})
})})
	})
	return eq1Set
}

var ord1Set gopurs_runtime.Value
var once_ord1Set sync.Once
func Get_ord1Set() gopurs_runtime.Value {
	once_ord1Set.Do(func() {
		ord1Set = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"compare1": gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_ordSet(), dictOrd_0).PtrVal.(map[string]gopurs_runtime.Value)["compare"]
}), "Eq10": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Set()
})})
	})
	return ord1Set
}

var empty gopurs_runtime.Value
var once_empty sync.Once
func Get_empty() gopurs_runtime.Value {
	once_empty.Do(func() {
		empty = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")})
	})
	return empty
}

var fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		fromFoldable = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictOrd_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_insert(), dictOrd_1), a_3), pkg_Data_Unit.Get_unit()), m_2)
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")}))
})
})
	})
	return fromFoldable
}

var map_ gopurs_runtime.Value
var once_map_ sync.Once
func Get_map_() gopurs_runtime.Value {
	once_map_.Do(func() {
		map_ = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableSet().PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_insert(), dictOrd_0), gopurs_runtime.Apply(f_1, a_3)), pkg_Data_Unit.Get_unit()), m_2)
})
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")}))
})
})
	})
	return map_
}

var mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		mapMaybe = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_foldableSet().PtrVal.(map[string]gopurs_runtime.Value)["foldr"], gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(f_1, a_2)
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_4_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t1 = acc_3
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_4_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_insert(), dictOrd_0), __local_var_4_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), pkg_Data_Unit.Get_unit()), acc_3)
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
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")}))
})
})
	})
	return mapMaybe
}

var monoidSet gopurs_runtime.Value
var once_monoidSet sync.Once
func Get_monoidSet() gopurs_runtime.Value {
	once_monoidSet.Do(func() {
		monoidSet = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupSet1_1_0 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Apply(Get_union(), dictOrd_0)})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")}), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupSet1_1_0
})})
})
	})
	return monoidSet
}

var unions gopurs_runtime.Value
var once_unions sync.Once
func Get_unions() gopurs_runtime.Value {
	once_unions.Do(func() {
		unions = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictOrd_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], gopurs_runtime.Apply(Get_union(), dictOrd_1)), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Leaf")}))
})
})
	})
	return unions
}

var difference gopurs_runtime.Value
var once_difference sync.Once
func Get_difference() gopurs_runtime.Value {
	once_difference.Do(func() {
		difference = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn3(), pkg_Data_Map_Internal.Get_unsafeDifference()), compare_1_0), m1_2), m2_3)
})
}))
})
	})
	return difference
}

var subset gopurs_runtime.Value
var once_subset sync.Once
func Get_subset() gopurs_runtime.Value {
	once_subset.Do(func() {
		subset = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_1 := dictOrd_0.PtrVal.(map[string]gopurs_runtime.Value)["compare"]
difference1_1_0 := gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Func(func(m1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Function_Uncurried.Get_runFn3(), pkg_Data_Map_Internal.Get_unsafeDifference()), compare_1_1), m1_2), m2_3)
})
}))
return gopurs_runtime.Func(func(s1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_isEmpty(), gopurs_runtime.Apply(gopurs_runtime.Apply(difference1_1_0, s1_2), s2_3))
})
})
})
	})
	return subset
}

var properSubset gopurs_runtime.Value
var once_properSubset sync.Once
func Get_properSubset() gopurs_runtime.Value {
	once_properSubset.Do(func() {
		properSubset = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
subset1_1_0 := gopurs_runtime.Apply(Get_subset(), dictOrd_0)
return gopurs_runtime.Func(func(s1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_HeytingAlgebra.Get_boolConj(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqBooleanImpl(), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Eq.Get_eqIntImpl(), gopurs_runtime.Apply(Get_size(), s1_2)), gopurs_runtime.Apply(Get_size(), s2_3))), gopurs_runtime.Bool(false))), gopurs_runtime.Apply(gopurs_runtime.Apply(subset1_1_0, s1_2), s2_3))
})
})
})
	})
	return properSubset
}

var delete_ gopurs_runtime.Value
var once_delete_ sync.Once
func Get_delete_() gopurs_runtime.Value {
	once_delete_.Do(func() {
		delete_ = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_delete_(), dictOrd_0))
})
	})
	return delete_
}

var checkValid gopurs_runtime.Value
var once_checkValid sync.Once
func Get_checkValid() gopurs_runtime.Value {
	once_checkValid.Do(func() {
		checkValid = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_checkValid(), dictOrd_0))
})
	})
	return checkValid
}

var catMaybes gopurs_runtime.Value
var once_catMaybes sync.Once
func Get_catMaybes() gopurs_runtime.Value {
	once_catMaybes.Do(func() {
		catMaybes = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_mapMaybe(), dictOrd_0), pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
})
	})
	return catMaybes
}



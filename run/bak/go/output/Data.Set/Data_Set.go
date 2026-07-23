package Data_Set

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Map_Internal "gopurs/output/Data.Map.Internal"
	pkg_Data_Function "gopurs/output/Data.Function"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_Eq "gopurs/output/Data.Eq"
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
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), compare_1_0, pkg_Data_Function.Get_const_(), m1_2, m2_3)
})
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
_ = alter_1_0
return gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(alter_1_0, gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_4, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), pkg_Data_Unit.Get_unit())
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_4, "_tag").StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), a_2, v_3)
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
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(xs_1, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(xs_1, "_tag").StrVal == "Cons")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Tuple"), gopurs_runtime.RecordGet(xs_1, "value0"), gopurs_runtime.RecordGet(xs_1, "value1")))
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
go__3_2 = gopurs_runtime.Value{PtrVal: func(m_prime_4 gopurs_runtime.Value, z_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(m_prime_4, "_tag").StrVal == "Leaf")).IntVal != 0 {
__t3 = z_prime_5
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(m_prime_4, "_tag").StrVal == "Node")).IntVal != 0 {
__t3 = gopurs_runtime.UncurriedApp2(go__3_2, gopurs_runtime.RecordGet(m_prime_4, "value4"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(m_prime_4, "value2"), gopurs_runtime.UncurriedApp2(go__3_2, gopurs_runtime.RecordGet(m_prime_4, "value5"), z_prime_5)))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}}
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.UncurriedApp2(go__3_2, x_2, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))))
})
})
	})
	return toUnfoldable
}

var toUnfoldable1 gopurs_runtime.Value
var once_toUnfoldable1 sync.Once
func Get_toUnfoldable1() gopurs_runtime.Value {
	once_toUnfoldable1.Do(func() {
		toUnfoldable1 = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Unfoldable.Get_unfoldableArray(), "unfoldr"), gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(xs_0, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(xs_0, "_tag").StrVal == "Cons")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Tuple"), gopurs_runtime.RecordGet(xs_0, "value0"), gopurs_runtime.RecordGet(xs_0, "value1")))
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
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_2 gopurs_runtime.Value
_ = go__2_2
go__2_2 = gopurs_runtime.Value{PtrVal: func(m_prime_3 gopurs_runtime.Value, z_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(m_prime_3, "_tag").StrVal == "Leaf")).IntVal != 0 {
__t3 = z_prime_4
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(m_prime_3, "_tag").StrVal == "Node")).IntVal != 0 {
__t3 = gopurs_runtime.UncurriedApp2(go__2_2, gopurs_runtime.RecordGet(m_prime_3, "value4"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(m_prime_3, "value2"), gopurs_runtime.UncurriedApp2(go__2_2, gopurs_runtime.RecordGet(m_prime_3, "value5"), z_prime_4)))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}}
return gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.UncurriedApp2(go__2_2, x_1, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))))
})
}()
	})
	return toUnfoldable1
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
		singleton = gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict([]string{"_tag", "value0", "value1", "value2", "value3", "value4", "value5"}, []gopurs_runtime.Value{gopurs_runtime.Str("Node"), gopurs_runtime.Int(1), gopurs_runtime.Int(1), a_0, pkg_Data_Unit.Get_unit(), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Leaf")), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Leaf"))})
})
	})
	return singleton
}

var showSet gopurs_runtime.Value
var once_showSet sync.Once
func Get_showSet() gopurs_runtime.Value {
	once_showSet.Do(func() {
		showSet = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str("(fromFoldable ").StrVal + gopurs_runtime.Apply2(pkg_Data_Show.Get_showArrayImpl(), gopurs_runtime.RecordGet(dictShow_0, "show"), gopurs_runtime.Apply(Get_toUnfoldable1(), s_1)).StrVal).StrVal + gopurs_runtime.Str(")").StrVal)
}))
})
	})
	return showSet
}

var semigroupSet gopurs_runtime.Value
var once_semigroupSet sync.Once
func Get_semigroupSet() gopurs_runtime.Value {
	once_semigroupSet.Do(func() {
		semigroupSet = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), compare_1_0, pkg_Data_Function.Get_const_(), m1_2, m2_3)
}))
})
	})
	return semigroupSet
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
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_3, "_tag").StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_3, "_tag").StrVal == "Node")).IntVal != 0 {
v1_4_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), k_1, gopurs_runtime.RecordGet(v_3, "value2"))
_ = v1_4_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_4_2, "_tag").StrVal == "LT")).IntVal != 0 {
v_3_loop = gopurs_runtime.RecordGet(v_3, "value4")
continue go__2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_4_2, "_tag").StrVal == "GT")).IntVal != 0 {
v_3_loop = gopurs_runtime.RecordGet(v_3, "value5")
continue go__2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_4_2, "_tag").StrVal == "EQ")).IntVal != 0 {
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

var isEmpty gopurs_runtime.Value
var once_isEmpty sync.Once
func Get_isEmpty() gopurs_runtime.Value {
	once_isEmpty.Do(func() {
		isEmpty = pkg_Data_Map_Internal.Get_isEmpty()
	})
	return isEmpty
}

var intersection gopurs_runtime.Value
var once_intersection sync.Once
func Get_intersection() gopurs_runtime.Value {
	once_intersection.Do(func() {
		intersection = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeIntersectionWith(), compare_1_0, pkg_Data_Function.Get_const_(), m1_2, m2_3)
})
})
	})
	return intersection
}

var insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		insert = gopurs_runtime.Func3(func(dictOrd_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(pkg_Data_Map_Internal.Get_insert(), dictOrd_0, a_1, pkg_Data_Unit.Get_unit(), v_2)
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
		foldableSet = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldMap"), dictMonoid_0)
_ = foldMap1_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(foldMap1_1_0, f_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__5_2 gopurs_runtime.Value
_ = go__5_2
go__5_2 = gopurs_runtime.Value{PtrVal: func(m_prime_6 gopurs_runtime.Value, z_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(m_prime_6, "_tag").StrVal == "Leaf")).IntVal != 0 {
__t3 = z_prime_7
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(m_prime_6, "_tag").StrVal == "Node")).IntVal != 0 {
__t3 = gopurs_runtime.UncurriedApp2(go__5_2, gopurs_runtime.RecordGet(m_prime_6, "value4"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(m_prime_6, "value2"), gopurs_runtime.UncurriedApp2(go__5_2, gopurs_runtime.RecordGet(m_prime_6, "value5"), z_prime_7)))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}}
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.UncurriedApp2(go__5_2, x_4, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))))
})
})
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_4 gopurs_runtime.Value
go__2_4 = gopurs_runtime.Func(func(b_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_4:
for {
if false { continue go__2_4 }
var b_3 = b_3_loop
_ = b_3
var v_4 = v_4_loop
_ = v_4
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_4, "_tag").StrVal == "Nil")).IntVal != 0 {
__t5 = b_3
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_4, "_tag").StrVal == "Cons")).IntVal != 0 {
b_3_loop = gopurs_runtime.Apply2(f_0, b_3, gopurs_runtime.RecordGet(v_4, "value0"))
v_4_loop = gopurs_runtime.RecordGet(v_4, "value1")
continue go__2_4
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
__local_var_3_6 := gopurs_runtime.Apply(go__2_4, x_1)
_ = __local_var_3_6
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__5_7 gopurs_runtime.Value
_ = go__5_7
go__5_7 = gopurs_runtime.Value{PtrVal: func(m_prime_6 gopurs_runtime.Value, z_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(m_prime_6, "_tag").StrVal == "Leaf")).IntVal != 0 {
__t8 = z_prime_7
goto end_branch_8
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(m_prime_6, "_tag").StrVal == "Node")).IntVal != 0 {
__t8 = gopurs_runtime.UncurriedApp2(go__5_7, gopurs_runtime.RecordGet(m_prime_6, "value4"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(m_prime_6, "value2"), gopurs_runtime.UncurriedApp2(go__5_7, gopurs_runtime.RecordGet(m_prime_6, "value5"), z_prime_7)))
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
}}
return gopurs_runtime.Apply(__local_var_3_6, gopurs_runtime.UncurriedApp2(go__5_7, x_4, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))))
})
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_9 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), f_0, x_1)
_ = __local_var_2_9
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_10 gopurs_runtime.Value
_ = go__4_10
go__4_10 = gopurs_runtime.Value{PtrVal: func(m_prime_5 gopurs_runtime.Value, z_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(m_prime_5, "_tag").StrVal == "Leaf")).IntVal != 0 {
__t11 = z_prime_6
goto end_branch_11
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(m_prime_5, "_tag").StrVal == "Node")).IntVal != 0 {
__t11 = gopurs_runtime.UncurriedApp2(go__4_10, gopurs_runtime.RecordGet(m_prime_5, "value4"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(m_prime_5, "value2"), gopurs_runtime.UncurriedApp2(go__4_10, gopurs_runtime.RecordGet(m_prime_5, "value5"), z_prime_6)))
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
}}
return gopurs_runtime.Apply(__local_var_2_9, gopurs_runtime.UncurriedApp2(go__4_10, x_3, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))))
})
}))
	})
	return foldableSet
}

var findMin gopurs_runtime.Value
var once_findMin sync.Once
func Get_findMin() gopurs_runtime.Value {
	once_findMin.Do(func() {
		findMin = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_findMin(), v_0)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_1_0, "_tag").StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_1_0, "value0"), "key"))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
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
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_1_0, "_tag").StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_1_0, "value0"), "key"))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
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
return gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_filterKeys(), dictOrd_0)
})
	})
	return filter
}

var eqSet gopurs_runtime.Value
var once_eqSet sync.Once
func Get_eqSet() gopurs_runtime.Value {
	once_eqSet.Do(func() {
		eqSet = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(pkg_Data_Map_Internal.Get_eqMap(), dictEq_0, pkg_Data_Eq.Get_eqUnit()), "eq"), v_1, v1_2)
}))
})
	})
	return eqSet
}

var ordSet gopurs_runtime.Value
var once_ordSet sync.Once
func Get_ordSet() gopurs_runtime.Value {
	once_ordSet.Do(func() {
		ordSet = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
eqSet1_2_1 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(pkg_Data_Map_Internal.Get_eqMap(), __local_var_1_0, pkg_Data_Eq.Get_eqUnit()), "eq"), v_2, v1_3)
}))
_ = eqSet1_2_1
return gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(s1_3 gopurs_runtime.Value, s2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__5_2 gopurs_runtime.Value
_ = go__5_2
go__5_2 = gopurs_runtime.Value{PtrVal: func(m_prime_6 gopurs_runtime.Value, z_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(m_prime_6, "_tag").StrVal == "Leaf")).IntVal != 0 {
__t3 = z_prime_7
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(m_prime_6, "_tag").StrVal == "Node")).IntVal != 0 {
__t3 = gopurs_runtime.UncurriedApp2(go__5_2, gopurs_runtime.RecordGet(m_prime_6, "value4"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(m_prime_6, "value2"), gopurs_runtime.UncurriedApp2(go__5_2, gopurs_runtime.RecordGet(m_prime_6, "value5"), z_prime_7)))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}}
var go__5_4 gopurs_runtime.Value
_ = go__5_4
go__5_4 = gopurs_runtime.Value{PtrVal: func(m_prime_6 gopurs_runtime.Value, z_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(m_prime_6, "_tag").StrVal == "Leaf")).IntVal != 0 {
__t5 = z_prime_7
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(m_prime_6, "_tag").StrVal == "Node")).IntVal != 0 {
__t5 = gopurs_runtime.UncurriedApp2(go__5_4, gopurs_runtime.RecordGet(m_prime_6, "value4"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(m_prime_6, "value2"), gopurs_runtime.UncurriedApp2(go__5_4, gopurs_runtime.RecordGet(m_prime_6, "value5"), z_prime_7)))
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}}
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(pkg_Data_List_Types.Get_ordList(), dictOrd_0), "compare"), gopurs_runtime.UncurriedApp2(go__5_2, s1_3, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))), gopurs_runtime.UncurriedApp2(go__5_4, s2_4, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eqSet1_2_1
}))
})
	})
	return ordSet
}

var eq1Set gopurs_runtime.Value
var once_eq1Set sync.Once
func Get_eq1Set() gopurs_runtime.Value {
	once_eq1Set.Do(func() {
		eq1Set = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func3(func(dictEq_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(pkg_Data_Map_Internal.Get_eqMap(), dictEq_0, pkg_Data_Eq.Get_eqUnit()), "eq"), v_1, v1_2)
}))
	})
	return eq1Set
}

var ord1Set gopurs_runtime.Value
var once_ord1Set sync.Once
func Get_ord1Set() gopurs_runtime.Value {
	once_ord1Set.Do(func() {
		ord1Set = gopurs_runtime.RecordDict2("compare1", "Eq10", gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_ordSet(), dictOrd_0), "compare")
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Set()
}))
	})
	return ord1Set
}

var empty gopurs_runtime.Value
var once_empty sync.Once
func Get_empty() gopurs_runtime.Value {
	once_empty.Do(func() {
		empty = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Leaf"))
	})
	return empty
}

var fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		fromFoldable = gopurs_runtime.Func2(func(dictFoldable_0 gopurs_runtime.Value, dictOrd_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(m_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(pkg_Data_Map_Internal.Get_insert(), dictOrd_1, a_3, pkg_Data_Unit.Get_unit(), m_2)
}), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Leaf")))
})
	})
	return fromFoldable
}

var map_ gopurs_runtime.Value
var once_map_ sync.Once
func Get_map_() gopurs_runtime.Value {
	once_map_.Do(func() {
		map_ = gopurs_runtime.Func2(func(dictOrd_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(b_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_0:
for {
if false { continue go__2_0 }
var b_3 = b_3_loop
_ = b_3
var v_4 = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_4, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = b_3
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_4, "_tag").StrVal == "Cons")).IntVal != 0 {
b_3_loop = gopurs_runtime.Apply4(pkg_Data_Map_Internal.Get_insert(), dictOrd_0, gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet(v_4, "value0")), pkg_Data_Unit.Get_unit(), b_3)
v_4_loop = gopurs_runtime.RecordGet(v_4, "value1")
continue go__2_0
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
__local_var_3_2 := gopurs_runtime.Apply(go__2_0, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Leaf")))
_ = __local_var_3_2
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__5_3 gopurs_runtime.Value
_ = go__5_3
go__5_3 = gopurs_runtime.Value{PtrVal: func(m_prime_6 gopurs_runtime.Value, z_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(m_prime_6, "_tag").StrVal == "Leaf")).IntVal != 0 {
__t4 = z_prime_7
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(m_prime_6, "_tag").StrVal == "Node")).IntVal != 0 {
__t4 = gopurs_runtime.UncurriedApp2(go__5_3, gopurs_runtime.RecordGet(m_prime_6, "value4"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(m_prime_6, "value2"), gopurs_runtime.UncurriedApp2(go__5_3, gopurs_runtime.RecordGet(m_prime_6, "value5"), z_prime_7)))
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}}
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.UncurriedApp2(go__5_3, x_4, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))))
})
})
	})
	return map_
}

var mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		mapMaybe = gopurs_runtime.Func2(func(dictOrd_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableSet(), "foldr"), gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(f_1, a_2)
_ = __local_var_4_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_4_0, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t1 = acc_3
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_4_0, "_tag").StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Apply4(pkg_Data_Map_Internal.Get_insert(), dictOrd_0, gopurs_runtime.RecordGet(__local_var_4_0, "value0"), pkg_Data_Unit.Get_unit(), acc_3)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Leaf")))
})
	})
	return mapMaybe
}

var monoidSet gopurs_runtime.Value
var once_monoidSet sync.Once
func Get_monoidSet() gopurs_runtime.Value {
	once_monoidSet.Do(func() {
		monoidSet = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_1 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_1
semigroupSet1_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), compare_1_1, pkg_Data_Function.Get_const_(), m1_2, m2_3)
}))
_ = semigroupSet1_1_0
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Leaf")), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupSet1_1_0
}))
})
	})
	return monoidSet
}

var unions gopurs_runtime.Value
var once_unions sync.Once
func Get_unions() gopurs_runtime.Value {
	once_unions.Do(func() {
		unions = gopurs_runtime.Func2(func(dictFoldable_0 gopurs_runtime.Value, dictOrd_1 gopurs_runtime.Value) gopurs_runtime.Value {
compare_2_0 := gopurs_runtime.RecordGet(dictOrd_1, "compare")
_ = compare_2_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(m1_3 gopurs_runtime.Value, m2_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp4(pkg_Data_Map_Internal.Get_unsafeUnionWith(), compare_2_0, pkg_Data_Function.Get_const_(), m1_3, m2_4)
}), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Leaf")))
})
	})
	return unions
}

var difference gopurs_runtime.Value
var once_difference sync.Once
func Get_difference() gopurs_runtime.Value {
	once_difference.Do(func() {
		difference = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func2(func(m1_2 gopurs_runtime.Value, m2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(pkg_Data_Map_Internal.Get_unsafeDifference(), compare_1_0, m1_2, m2_3)
})
})
	})
	return difference
}

var subset gopurs_runtime.Value
var once_subset sync.Once
func Get_subset() gopurs_runtime.Value {
	once_subset.Do(func() {
		subset = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func2(func(s1_2 gopurs_runtime.Value, s2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.UncurriedApp3(pkg_Data_Map_Internal.Get_unsafeDifference(), compare_1_0, s1_2, s2_3), "_tag").StrVal == "Leaf")
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
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func2(func(s1_2 gopurs_runtime.Value, s2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(s1_2, "_tag").StrVal == "Leaf")).IntVal != 0 {
__t1 = gopurs_runtime.Int(0)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(s1_2, "_tag").StrVal == "Node")).IntVal != 0 {
__t1 = gopurs_runtime.RecordGet(s1_2, "value1")
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(s2_3, "_tag").StrVal == "Leaf")).IntVal != 0 {
__t2 = gopurs_runtime.Int(0)
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(s2_3, "_tag").StrVal == "Node")).IntVal != 0 {
__t2 = gopurs_runtime.RecordGet(s2_3, "value1")
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Bool(gopurs_runtime.Bool(__t1.IntVal != __t2.IntVal).IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.UncurriedApp3(pkg_Data_Map_Internal.Get_unsafeDifference(), compare_1_0, s1_2, s2_3), "_tag").StrVal == "Leaf").IntVal != 0)
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
return gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_delete_(), dictOrd_0)
})
	})
	return delete_
}

var checkValid gopurs_runtime.Value
var once_checkValid sync.Once
func Get_checkValid() gopurs_runtime.Value {
	once_checkValid.Do(func() {
		checkValid = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_checkValid(), dictOrd_0)
})
	})
	return checkValid
}

var catMaybes gopurs_runtime.Value
var once_catMaybes sync.Once
func Get_catMaybes() gopurs_runtime.Value {
	once_catMaybes.Do(func() {
		catMaybes = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_mapMaybe(), dictOrd_0, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
})
	})
	return catMaybes
}



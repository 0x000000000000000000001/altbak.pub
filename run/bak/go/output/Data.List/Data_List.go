package Data_List

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_List_Internal "gopurs/output/Data.List.Internal"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var any gopurs_runtime.Value
var once_any sync.Once
func Get_any() gopurs_runtime.Value {
	once_any.Do(func() {
		any = func() gopurs_runtime.Value {
semigroupDisj1_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(v_0.IntVal != 0 || v1_1.IntVal != 0)
}))
_ = semigroupDisj1_0_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldMap"), gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Bool(false), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_0_0
})))
}()
	})
	return any
}

var Pattern gopurs_runtime.Value
var once_Pattern sync.Once
func Get_Pattern() gopurs_runtime.Value {
	once_Pattern.Do(func() {
		Pattern = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Pattern
}

var updateAt gopurs_runtime.Value
var once_updateAt sync.Once
func Get_updateAt() gopurs_runtime.Value {
	once_updateAt.Do(func() {
		updateAt = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
updateAt:
for {
if false { continue updateAt }
var v_0 = v_0_loop
_ = v_0
var v1_1 = v1_1_loop
_ = v1_1
var v2_2 = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_2, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t3 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), v1_1, gopurs_runtime.RecordGet(v2_2, "value1")))
goto end_branch_3
} else {

}
}
{
__local_var_3_1 := gopurs_runtime.Apply3(Get_updateAt(), gopurs_runtime.Int(v_0.IntVal - gopurs_runtime.Int(1).IntVal), v1_1, gopurs_runtime.RecordGet(v2_2, "value1"))
_ = __local_var_3_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_3_1, "_tag").StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v2_2, "value0"), gopurs_runtime.RecordGet(__local_var_3_1, "value0")))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_2:
__t3 = __t2
}
end_branch_3:
__t0 = __t3
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_0:
return __t0
}
}()
})
})
})
	})
	return updateAt
}

var unzip gopurs_runtime.Value
var once_unzip sync.Once
func Get_unzip() gopurs_runtime.Value {
	once_unzip.Do(func() {
		unzip = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.RecordGet(v_0, "value0")
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.RecordGet(v_0, "value1")
_ = __local_var_2_1
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Tuple"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), __local_var_1_0, gopurs_runtime.RecordGet(v1_3, "value0")), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), __local_var_2_1, gopurs_runtime.RecordGet(v1_3, "value1")))
})
}), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Tuple"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))))
	})
	return unzip
}

var uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		uncons = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_0, "_tag").StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_0, "_tag").StrVal == "Cons")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.RecordGet(v_0, "value0"), gopurs_runtime.RecordGet(v_0, "value1")))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
	})
	return uncons
}

var toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(xs_1, "_tag").StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(xs_1, "_tag").StrVal == "Cons")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Tuple"), gopurs_runtime.RecordGet(xs_1, "value0"), gopurs_runtime.RecordGet(xs_1, "value1")))
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
	})
	return toUnfoldable
}

var tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		tail = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_0, "_tag").StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_0, "_tag").StrVal == "Cons")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordGet(v_0, "value1"))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
	})
	return tail
}

var stripPrefix gopurs_runtime.Value
var once_stripPrefix sync.Once
func Get_stripPrefix() gopurs_runtime.Value {
	once_stripPrefix.Do(func() {
		stripPrefix = gopurs_runtime.Func3(func(dictEq_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, s_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Value{PtrVal: func(prefix_3 gopurs_runtime.Value, input_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(input_4, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(prefix_3, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), gopurs_runtime.RecordGet(prefix_3, "value0"), gopurs_runtime.RecordGet(input_4, "value0"))).IntVal != 0 {
__t3 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Loop"), gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.RecordGet(prefix_3, "value1"), gopurs_runtime.RecordGet(input_4, "value1"))))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(prefix_3, "_tag").StrVal == "Nil")).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Done"), input_4))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(prefix_3, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Done"), input_4))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_1:
return __t1
}}
_ = __local_var_3_0
__local_var_4_4 := gopurs_runtime.Value{PtrVal: func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_4, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t5 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Done"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing")))
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_4, "_tag").StrVal == "Just")).IntVal != 0 {
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_4, "value0"), "_tag").StrVal == "Loop")).IntVal != 0 {
__t6 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Loop"), gopurs_runtime.UncurriedApp2(__local_var_3_0, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_4, "value0"), "value0"), "a"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_4, "value0"), "value0"), "b")))
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_4, "value0"), "_tag").StrVal == "Done")).IntVal != 0 {
__t6 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Done"), gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_4, "value0"), "value0")))
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
__t5 = __t6
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
_ = __local_var_4_4
var go__5_7 gopurs_runtime.Value
go__5_7 = gopurs_runtime.Func(func(v_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__5_7:
for {
if false { continue go__5_7 }
var v_6 = v_6_loop
_ = v_6
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_6, "_tag").StrVal == "Loop")).IntVal != 0 {
v_6_loop = gopurs_runtime.UncurriedApp(__local_var_4_4, gopurs_runtime.RecordGet(v_6, "value0"))
continue go__5_7
__t8 = gopurs_runtime.Value{}
goto end_branch_8
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_6, "_tag").StrVal == "Done")).IntVal != 0 {
__t8 = gopurs_runtime.RecordGet(v_6, "value0")
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
return gopurs_runtime.Apply(go__5_7, gopurs_runtime.UncurriedApp(__local_var_4_4, gopurs_runtime.UncurriedApp2(__local_var_3_0, v_1, s_2)))
})
	})
	return stripPrefix
}

var span gopurs_runtime.Value
var once_span sync.Once
func Get_span() gopurs_runtime.Value {
	once_span.Do(func() {
		span = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
span:
for {
if false { continue span }
var v_0 = v_0_loop
_ = v_0
var v1_1 = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_1, "_tag").StrVal == "Cons").IntVal != 0 && gopurs_runtime.Apply(v_0, gopurs_runtime.RecordGet(v1_1, "value0")).IntVal != 0)).IntVal != 0 {
v2_2_1 := gopurs_runtime.Apply2(Get_span(), v_0, gopurs_runtime.RecordGet(v1_1, "value1"))
_ = v2_2_1
__t0 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_1, "value0"), gopurs_runtime.RecordGet(v2_2_1, "init")), gopurs_runtime.RecordGet(v2_2_1, "rest"))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")), v1_1)
}
end_branch_0:
return __t0
}
}()
})
})
	})
	return span
}

var snoc gopurs_runtime.Value
var once_snoc sync.Once
func Get_snoc() gopurs_runtime.Value {
	once_snoc.Do(func() {
		snoc = gopurs_runtime.Func2(func(xs_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), x_1, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))), xs_0)
})
	})
	return snoc
}

var singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		singleton = gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), a_0, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")))
})
	})
	return singleton
}

var sortBy gopurs_runtime.Value
var once_sortBy sync.Once
func Get_sortBy() gopurs_runtime.Value {
	once_sortBy.Do(func() {
		sortBy = gopurs_runtime.Func(func(cmp_0 gopurs_runtime.Value) gopurs_runtime.Value {
var merge_1_0 gopurs_runtime.Value
_ = merge_1_0
merge_1_0 = gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(cmp_0, gopurs_runtime.RecordGet(v_2, "value0"), gopurs_runtime.RecordGet(v1_3, "value0")), "_tag").StrVal == "GT")).IntVal != 0 {
__t3 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_3, "value0"), gopurs_runtime.Apply2(merge_1_0, v_2, gopurs_runtime.RecordGet(v1_3, "value1")))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v_2, "value0"), gopurs_runtime.Apply2(merge_1_0, gopurs_runtime.RecordGet(v_2, "value1"), v1_3))
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Nil")).IntVal != 0 {
__t2 = v_2
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = v_2
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
var mergePairs_2_4 gopurs_runtime.Value
_ = mergePairs_2_4
mergePairs_2_4 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_3, "_tag").StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_3, "value1"), "_tag").StrVal == "Cons").IntVal != 0)).IntVal != 0 {
__t5 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.Apply2(merge_1_0, gopurs_runtime.RecordGet(v_3, "value0"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_3, "value1"), "value0")), gopurs_runtime.Apply(mergePairs_2_4, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_3, "value1"), "value1")))
goto end_branch_5
} else {

}
}
{
__t5 = v_3
}
end_branch_5:
return __t5
})
var mergeAll_3_6 gopurs_runtime.Value
mergeAll_3_6 = gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
mergeAll_3_6:
for {
if false { continue mergeAll_3_6 }
var v_4 = v_4_loop
_ = v_4
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_4, "_tag").StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_4, "value1"), "_tag").StrVal == "Nil").IntVal != 0)).IntVal != 0 {
__t7 = gopurs_runtime.RecordGet(v_4, "value0")
goto end_branch_7
} else {

}
}
{
v_4_loop = gopurs_runtime.Apply(mergePairs_2_4, v_4)
continue mergeAll_3_6
__t7 = gopurs_runtime.Value{}
}
end_branch_7:
return __t7
}
}()
})
var sequences_4_8 gopurs_runtime.Value
_ = sequences_4_8
var descending_4_9 gopurs_runtime.Value
_ = descending_4_9
var ascending_4_10 gopurs_runtime.Value
_ = ascending_4_10
sequences_4_8 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_5, "_tag").StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_5, "value1"), "_tag").StrVal == "Cons").IntVal != 0)).IntVal != 0 {
var __t13 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(cmp_0, gopurs_runtime.RecordGet(v_5, "value0"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_5, "value1"), "value0")), "_tag").StrVal == "GT")).IntVal != 0 {
__t13 = gopurs_runtime.Apply3(descending_4_9, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_5, "value1"), "value0"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v_5, "value0"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_5, "value1"), "value1"))
goto end_branch_13
} else {

}
}
{
__local_var_6_12 := gopurs_runtime.RecordGet(v_5, "value0")
_ = __local_var_6_12
__t13 = gopurs_runtime.Apply3(ascending_4_10, gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_5, "value1"), "value0"), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), __local_var_6_12, v1_7)
}), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_5, "value1"), "value1"))
}
end_branch_13:
__t11 = __t13
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), v_5, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")))
}
end_branch_11:
return __t11
})
descending_4_9 = gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_7, "_tag").StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(cmp_0, v_5, gopurs_runtime.RecordGet(v2_7, "value0")), "_tag").StrVal == "GT").IntVal != 0)).IntVal != 0 {
__t14 = gopurs_runtime.Apply3(descending_4_9, gopurs_runtime.RecordGet(v2_7, "value0"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), v_5, v1_6), gopurs_runtime.RecordGet(v2_7, "value1"))
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), v_5, v1_6), gopurs_runtime.Apply(sequences_4_8, v2_7))
}
end_branch_14:
return __t14
})
ascending_4_10 = gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
__local_var_8_16 := gopurs_runtime.Apply2(cmp_0, v_5, gopurs_runtime.RecordGet(v2_7, "value0"))
_ = __local_var_8_16
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_7, "_tag").StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_8_16, "_tag").StrVal == "LT").IntVal != 0 || gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_8_16, "_tag").StrVal == "GT").IntVal == 0).IntVal != 0).IntVal != 0)).IntVal != 0 {
__t15 = gopurs_runtime.Apply3(ascending_4_10, gopurs_runtime.RecordGet(v2_7, "value0"), gopurs_runtime.Func(func(ys_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_6, gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), v_5, ys_8))
}), gopurs_runtime.RecordGet(v2_7, "value1"))
goto end_branch_15
} else {

}
}
{
__t15 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.Apply(v1_6, gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), v_5, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")))), gopurs_runtime.Apply(sequences_4_8, v2_7))
}
end_branch_15:
return __t15
})
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(mergeAll_3_6, gopurs_runtime.Apply(sequences_4_8, x_5))
})
})
	})
	return sortBy
}

var sort gopurs_runtime.Value
var once_sort sync.Once
func Get_sort() gopurs_runtime.Value {
	once_sort.Do(func() {
		sort = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_sortBy(), compare_1_0, xs_2)
})
})
	})
	return sort
}

var showPattern gopurs_runtime.Value
var once_showPattern sync.Once
func Get_showPattern() gopurs_runtime.Value {
	once_showPattern.Do(func() {
		showPattern = gopurs_runtime.Func(func(dictShow_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str("(Pattern ").StrVal + gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(pkg_Data_List_Types.Get_showList(), dictShow_0), "show"), v_1).StrVal).StrVal + gopurs_runtime.Str(")").StrVal)
}))
})
	})
	return showPattern
}

var reverse gopurs_runtime.Value
var once_reverse sync.Once
func Get_reverse() gopurs_runtime.Value {
	once_reverse.Do(func() {
		reverse = func() gopurs_runtime.Value {
var go__0_0 gopurs_runtime.Value
go__0_0 = gopurs_runtime.Func(func(v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__0_0:
for {
if false { continue go__0_0 }
var v_1 = v_1_loop
_ = v_1
var v1_2 = v1_2_loop
_ = v1_2
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_2, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = v_1
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_2, "_tag").StrVal == "Cons")).IntVal != 0 {
v_1_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_2, "value0"), v_1)
v1_2_loop = gopurs_runtime.RecordGet(v1_2, "value1")
continue go__0_0
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
return gopurs_runtime.Apply(go__0_0, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")))
}()
	})
	return reverse
}

var take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		take = func() gopurs_runtime.Value {
var go__0_0 gopurs_runtime.Value
go__0_0 = gopurs_runtime.Func(func(v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__0_0:
for {
if false { continue go__0_0 }
var v_1 = v_1_loop
_ = v_1
var v1_2 = v1_2_loop
_ = v1_2
var v2_3 = v2_3_loop
_ = v2_3
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_2.IntVal < gopurs_runtime.Int(1).IntVal)).IntVal != 0 {
var go__4_2 gopurs_runtime.Value
go__4_2 = gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__4_2:
for {
if false { continue go__4_2 }
var v_5 = v_5_loop
_ = v_5
var v1_6 = v1_6_loop
_ = v1_6
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_6, "_tag").StrVal == "Nil")).IntVal != 0 {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_6, "_tag").StrVal == "Cons")).IntVal != 0 {
v_5_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_6, "value0"), v_5)
v1_6_loop = gopurs_runtime.RecordGet(v1_6, "value1")
continue go__4_2
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
__t1 = gopurs_runtime.Apply2(go__4_2, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")), v_1)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_3, "_tag").StrVal == "Nil")).IntVal != 0 {
var go__4_4 gopurs_runtime.Value
go__4_4 = gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__4_4:
for {
if false { continue go__4_4 }
var v_5 = v_5_loop
_ = v_5
var v1_6 = v1_6_loop
_ = v1_6
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_6, "_tag").StrVal == "Nil")).IntVal != 0 {
__t5 = v_5
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_6, "_tag").StrVal == "Cons")).IntVal != 0 {
v_5_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_6, "value0"), v_5)
v1_6_loop = gopurs_runtime.RecordGet(v1_6, "value1")
continue go__4_4
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
__t1 = gopurs_runtime.Apply2(go__4_4, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")), v_1)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_3, "_tag").StrVal == "Cons")).IntVal != 0 {
v_1_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v2_3, "value0"), v_1)
v1_2_loop = gopurs_runtime.Int(v1_2.IntVal - gopurs_runtime.Int(1).IntVal)
v2_3_loop = gopurs_runtime.RecordGet(v2_3, "value1")
continue go__0_0
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
})
return gopurs_runtime.Apply(go__0_0, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")))
}()
	})
	return take
}

var takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		takeWhile = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 = v_2_loop
_ = v_2
var v1_3 = v1_3_loop
_ = v1_3
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Cons").IntVal != 0 && gopurs_runtime.Apply(p_0, gopurs_runtime.RecordGet(v1_3, "value0")).IntVal != 0)).IntVal != 0 {
v_2_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_3, "value0"), v_2)
v1_3_loop = gopurs_runtime.RecordGet(v1_3, "value1")
continue go__1_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
var go__4_1 gopurs_runtime.Value
go__4_1 = gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__4_1:
for {
if false { continue go__4_1 }
var v_5 = v_5_loop
_ = v_5
var v1_6 = v1_6_loop
_ = v1_6
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_6, "_tag").StrVal == "Nil")).IntVal != 0 {
__t2 = v_5
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_6, "_tag").StrVal == "Cons")).IntVal != 0 {
v_5_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_6, "value0"), v_5)
v1_6_loop = gopurs_runtime.RecordGet(v1_6, "value1")
continue go__4_1
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
__t3 = gopurs_runtime.Apply2(go__4_1, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")), v_2)
}
end_branch_3:
return __t3
}
}()
})
})
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")))
})
	})
	return takeWhile
}

var unsnoc gopurs_runtime.Value
var once_unsnoc sync.Once
func Get_unsnoc() gopurs_runtime.Value {
	once_unsnoc.Do(func() {
		unsnoc = gopurs_runtime.Func(func(lst_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 = v_2_loop
_ = v_2
var v1_3 = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_2, "value1"), "_tag").StrVal == "Nil")).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict2("revInit", "last", v1_3, gopurs_runtime.RecordGet(v_2, "value0")))
goto end_branch_2
} else {

}
}
{
v_2_loop = gopurs_runtime.RecordGet(v_2, "value1")
v1_3_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v_2, "value0"), v1_3)
continue go__1_0
__t2 = gopurs_runtime.Value{}
}
end_branch_2:
__t1 = __t2
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
__local_var_2_3 := gopurs_runtime.Apply2(go__1_0, lst_0, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")))
_ = __local_var_2_3
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_2_3, "_tag").StrVal == "Just")).IntVal != 0 {
var go__3_5 gopurs_runtime.Value
go__3_5 = gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_5:
for {
if false { continue go__3_5 }
var v_4 = v_4_loop
_ = v_4
var v1_5 = v1_5_loop
_ = v1_5
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_5, "_tag").StrVal == "Nil")).IntVal != 0 {
__t6 = v_4
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_5, "_tag").StrVal == "Cons")).IntVal != 0 {
v_4_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_5, "value0"), v_4)
v1_5_loop = gopurs_runtime.RecordGet(v1_5, "value1")
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
__t4 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.Apply2(go__3_5, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_2_3, "value0"), "revInit")), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_2_3, "value0"), "last")))
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_4:
return __t4
})
	})
	return unsnoc
}

var zipWith gopurs_runtime.Value
var once_zipWith sync.Once
func Get_zipWith() gopurs_runtime.Value {
	once_zipWith.Do(func() {
		zipWith = gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value, ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
go__3_0 = gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_0:
for {
if false { continue go__3_0 }
var v_4 = v_4_loop
_ = v_4
var v1_5 = v1_5_loop
_ = v1_5
var v2_6 = v2_6_loop
_ = v2_6
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_4, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = v2_6
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_5, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = v2_6
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_4, "_tag").StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_5, "_tag").StrVal == "Cons").IntVal != 0)).IntVal != 0 {
v_4_loop = gopurs_runtime.RecordGet(v_4, "value1")
v1_5_loop = gopurs_runtime.RecordGet(v1_5, "value1")
v2_6_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.Apply2(f_0, gopurs_runtime.RecordGet(v_4, "value0"), gopurs_runtime.RecordGet(v1_5, "value0")), v2_6)
continue go__3_0
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
})
var go__4_2 gopurs_runtime.Value
go__4_2 = gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__4_2:
for {
if false { continue go__4_2 }
var v_5 = v_5_loop
_ = v_5
var v1_6 = v1_6_loop
_ = v1_6
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_6, "_tag").StrVal == "Nil")).IntVal != 0 {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_6, "_tag").StrVal == "Cons")).IntVal != 0 {
v_5_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_6, "value0"), v_5)
v1_6_loop = gopurs_runtime.RecordGet(v1_6, "value1")
continue go__4_2
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
return gopurs_runtime.Apply2(go__4_2, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")), gopurs_runtime.Apply3(go__3_0, xs_1, ys_2, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))))
})
	})
	return zipWith
}

var zip gopurs_runtime.Value
var once_zip sync.Once
func Get_zip() gopurs_runtime.Value {
	once_zip.Do(func() {
		zip = gopurs_runtime.Func2(func(xs_0 gopurs_runtime.Value, ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_0:
for {
if false { continue go__2_0 }
var v_3 = v_3_loop
_ = v_3
var v1_4 = v1_4_loop
_ = v1_4
var v2_5 = v2_5_loop
_ = v2_5
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_3, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = v2_5
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_4, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = v2_5
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_3, "_tag").StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_4, "_tag").StrVal == "Cons").IntVal != 0)).IntVal != 0 {
v_3_loop = gopurs_runtime.RecordGet(v_3, "value1")
v1_4_loop = gopurs_runtime.RecordGet(v1_4, "value1")
v2_5_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Tuple"), gopurs_runtime.RecordGet(v_3, "value0"), gopurs_runtime.RecordGet(v1_4, "value0")), v2_5)
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
})
var go__3_2 gopurs_runtime.Value
go__3_2 = gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_2:
for {
if false { continue go__3_2 }
var v_4 = v_4_loop
_ = v_4
var v1_5 = v1_5_loop
_ = v1_5
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_5, "_tag").StrVal == "Nil")).IntVal != 0 {
__t3 = v_4
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_5, "_tag").StrVal == "Cons")).IntVal != 0 {
v_4_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_5, "value0"), v_4)
v1_5_loop = gopurs_runtime.RecordGet(v1_5, "value1")
continue go__3_2
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
return gopurs_runtime.Apply2(go__3_2, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")), gopurs_runtime.Apply3(go__2_0, xs_0, ys_1, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))))
})
	})
	return zip
}

var zipWithA gopurs_runtime.Value
var once_zipWithA sync.Once
func Get_zipWithA() gopurs_runtime.Value {
	once_zipWithA.Do(func() {
		zipWithA = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
sequence1_1_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_traversableList(), "traverse"), dictApplicative_0, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
_ = sequence1_1_0
return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value, ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__5_1 gopurs_runtime.Value
go__5_1 = gopurs_runtime.Func(func(v_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_8_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__5_1:
for {
if false { continue go__5_1 }
var v_6 = v_6_loop
_ = v_6
var v1_7 = v1_7_loop
_ = v1_7
var v2_8 = v2_8_loop
_ = v2_8
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_6, "_tag").StrVal == "Nil")).IntVal != 0 {
__t2 = v2_8
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_7, "_tag").StrVal == "Nil")).IntVal != 0 {
__t2 = v2_8
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_6, "_tag").StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_7, "_tag").StrVal == "Cons").IntVal != 0)).IntVal != 0 {
v_6_loop = gopurs_runtime.RecordGet(v_6, "value1")
v1_7_loop = gopurs_runtime.RecordGet(v1_7, "value1")
v2_8_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.Apply2(f_2, gopurs_runtime.RecordGet(v_6, "value0"), gopurs_runtime.RecordGet(v1_7, "value0")), v2_8)
continue go__5_1
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
})
var go__6_3 gopurs_runtime.Value
go__6_3 = gopurs_runtime.Func(func(v_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__6_3:
for {
if false { continue go__6_3 }
var v_7 = v_7_loop
_ = v_7
var v1_8 = v1_8_loop
_ = v1_8
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_8, "_tag").StrVal == "Nil")).IntVal != 0 {
__t4 = v_7
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_8, "_tag").StrVal == "Cons")).IntVal != 0 {
v_7_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_8, "value0"), v_7)
v1_8_loop = gopurs_runtime.RecordGet(v1_8, "value1")
continue go__6_3
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
return gopurs_runtime.Apply(sequence1_1_0, gopurs_runtime.Apply2(go__6_3, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")), gopurs_runtime.Apply3(go__5_1, xs_3, ys_4, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")))))
})
})
	})
	return zipWithA
}

var range_ gopurs_runtime.Value
var once_range_ sync.Once
func Get_range_() gopurs_runtime.Value {
	once_range_.Do(func() {
		range_ = gopurs_runtime.Func2(func(start_0 gopurs_runtime.Value, end_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(start_0.IntVal == end_1.IntVal)).IntVal != 0 {
__t3 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), start_0, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")))
goto end_branch_3
} else {

}
}
{
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(s_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(step_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rest_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_0:
for {
if false { continue go__2_0 }
var s_3 = s_3_loop
_ = s_3
var e_4 = e_4_loop
_ = e_4
var step_5 = step_5_loop
_ = step_5
var rest_6 = rest_6_loop
_ = rest_6
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(s_3.IntVal == e_4.IntVal)).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), s_3, rest_6)
goto end_branch_1
} else {

}
}
{
s_3_loop = gopurs_runtime.Int(s_3.IntVal + step_5.IntVal)
e_4_loop = e_4
step_5_loop = step_5
rest_6_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), s_3, rest_6)
continue go__2_0
__t1 = gopurs_runtime.Value{}
}
end_branch_1:
return __t1
}
}()
})
})
})
})
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(start_0.IntVal > end_1.IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Int(1)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Int(-1)
}
end_branch_2:
__t3 = gopurs_runtime.Apply4(go__2_0, end_1, start_0, __t2, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")))
}
end_branch_3:
return __t3
})
	})
	return range_
}

var partition gopurs_runtime.Value
var once_partition sync.Once
func Get_partition() gopurs_runtime.Value {
	once_partition.Do(func() {
		partition = gopurs_runtime.Func2(func(p_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(p_0, x_2)).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.RecordGet(v_3, "no"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), x_2, gopurs_runtime.RecordGet(v_3, "yes")))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), x_2, gopurs_runtime.RecordGet(v_3, "no")), gopurs_runtime.RecordGet(v_3, "yes"))
}
end_branch_0:
return __t0
}), gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))), xs_1)
})
	})
	return partition
}

var null gopurs_runtime.Value
var once_null sync.Once
func Get_null() gopurs_runtime.Value {
	once_null.Do(func() {
		null = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_0, "_tag").StrVal == "Nil")
})
	})
	return null
}

var nubBy gopurs_runtime.Value
var once_nubBy sync.Once
func Get_nubBy() gopurs_runtime.Value {
	once_nubBy.Do(func() {
		nubBy = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 = v_2_loop
_ = v_2
var v1_3 = v1_3_loop
_ = v1_3
var v2_4 = v2_4_loop
_ = v2_4
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_4, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_4, "_tag").StrVal == "Cons")).IntVal != 0 {
v3_5_2 := gopurs_runtime.Apply3(pkg_Data_List_Internal.Get_insertAndLookupBy(), p_0, gopurs_runtime.RecordGet(v2_4, "value0"), v_2)
_ = v3_5_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v3_5_2, "found")).IntVal != 0 {
v_2_loop = gopurs_runtime.RecordGet(v3_5_2, "result")
v1_3_loop = v1_3
v2_4_loop = gopurs_runtime.RecordGet(v2_4, "value1")
continue go__1_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
v_2_loop = gopurs_runtime.RecordGet(v3_5_2, "result")
v1_3_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v2_4, "value0"), v1_3)
v2_4_loop = gopurs_runtime.RecordGet(v2_4, "value1")
continue go__1_0
__t3 = gopurs_runtime.Value{}
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
})
})
__local_var_2_4 := gopurs_runtime.Apply2(go__1_0, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Leaf")), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")))
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_5 gopurs_runtime.Value
go__4_5 = gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__4_5:
for {
if false { continue go__4_5 }
var v_5 = v_5_loop
_ = v_5
var v1_6 = v1_6_loop
_ = v1_6
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_6, "_tag").StrVal == "Nil")).IntVal != 0 {
__t6 = v_5
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_6, "_tag").StrVal == "Cons")).IntVal != 0 {
v_5_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_6, "value0"), v_5)
v1_6_loop = gopurs_runtime.RecordGet(v1_6, "value1")
continue go__4_5
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
return gopurs_runtime.Apply2(go__4_5, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")), gopurs_runtime.Apply(__local_var_2_4, x_3))
})
})
	})
	return nubBy
}

var nub gopurs_runtime.Value
var once_nub sync.Once
func Get_nub() gopurs_runtime.Value {
	once_nub.Do(func() {
		nub = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_nubBy(), gopurs_runtime.RecordGet(dictOrd_0, "compare"))
})
	})
	return nub
}

var newtypePattern gopurs_runtime.Value
var once_newtypePattern sync.Once
func Get_newtypePattern() gopurs_runtime.Value {
	once_newtypePattern.Do(func() {
		newtypePattern = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypePattern
}

var mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		mapMaybe = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 = v_2_loop
_ = v_2
var v1_3 = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Nil")).IntVal != 0 {
var go__4_2 gopurs_runtime.Value
go__4_2 = gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__4_2:
for {
if false { continue go__4_2 }
var v_5 = v_5_loop
_ = v_5
var v1_6 = v1_6_loop
_ = v1_6
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_6, "_tag").StrVal == "Nil")).IntVal != 0 {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_6, "_tag").StrVal == "Cons")).IntVal != 0 {
v_5_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_6, "value0"), v_5)
v1_6_loop = gopurs_runtime.RecordGet(v1_6, "value1")
continue go__4_2
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
__t1 = gopurs_runtime.Apply2(go__4_2, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")), v_2)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Cons")).IntVal != 0 {
v2_4_4 := gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v1_3, "value0"))
_ = v2_4_4
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_4_4, "_tag").StrVal == "Nothing")).IntVal != 0 {
v_2_loop = v_2
v1_3_loop = gopurs_runtime.RecordGet(v1_3, "value1")
continue go__1_0
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_4_4, "_tag").StrVal == "Just")).IntVal != 0 {
v_2_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v2_4_4, "value0"), v_2)
v1_3_loop = gopurs_runtime.RecordGet(v1_3, "value1")
continue go__1_0
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t1 = __t5
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
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")))
})
	})
	return mapMaybe
}

var manyRec gopurs_runtime.Value
var once_manyRec sync.Once
func Get_manyRec() gopurs_runtime.Value {
	once_manyRec.Do(func() {
		manyRec = gopurs_runtime.Func2(func(dictMonadRec_0 gopurs_runtime.Value, dictAlternative_1 gopurs_runtime.Value) gopurs_runtime.Value {
Alt0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_1, "Plus1"), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{})
_ = Alt0_2_0
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_1, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Func(func(p_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Alt0_2_0, "alt"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Alt0_2_0, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Control_Monad_Rec_Class.Get_Loop(), p_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "pure"), gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Done"), pkg_Data_Unit.Get_unit()))), gopurs_runtime.Func(func(aa_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(aa_6, "_tag").StrVal == "Loop")).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Loop"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(aa_6, "value0"), acc_5))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(aa_6, "_tag").StrVal == "Done")).IntVal != 0 {
var go__7_3 gopurs_runtime.Value
go__7_3 = gopurs_runtime.Func(func(v_8_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_9_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__7_3:
for {
if false { continue go__7_3 }
var v_8 = v_8_loop
_ = v_8
var v1_9 = v1_9_loop
_ = v1_9
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_9, "_tag").StrVal == "Nil")).IntVal != 0 {
__t4 = v_8
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_9, "_tag").StrVal == "Cons")).IntVal != 0 {
v_8_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_9, "value0"), v_8)
v1_9_loop = gopurs_runtime.RecordGet(v1_9, "value1")
continue go__7_3
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
__t2 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Done"), gopurs_runtime.Apply2(go__7_3, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")), acc_5))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "pure"), __t2)
}))
}), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")))
})
})
	})
	return manyRec
}

var someRec gopurs_runtime.Value
var once_someRec sync.Once
func Get_someRec() gopurs_runtime.Value {
	once_someRec.Do(func() {
		someRec = gopurs_runtime.Func2(func(dictMonadRec_0 gopurs_runtime.Value, dictAlternative_1 gopurs_runtime.Value) gopurs_runtime.Value {
manyRec2_2_0 := gopurs_runtime.Apply2(Get_manyRec(), dictMonadRec_0, dictAlternative_1)
_ = manyRec2_2_0
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_1, "Applicative0"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_1, "Plus1"), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_List_Types.Get_Cons(), v_3), gopurs_runtime.Apply(manyRec2_2_0, v_3))
})
})
	})
	return someRec
}

var some gopurs_runtime.Value
var once_some sync.Once
func Get_some() gopurs_runtime.Value {
	once_some.Do(func() {
		some = gopurs_runtime.Func3(func(dictAlternative_0 gopurs_runtime.Value, dictLazy_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_List_Types.Get_Cons(), v_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_1, "defer"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_many(), dictAlternative_0, dictLazy_1, v_2)
})))
})
	})
	return some
}

var many gopurs_runtime.Value
var once_many sync.Once
func Get_many() gopurs_runtime.Value {
	once_many.Do(func() {
		many = gopurs_runtime.Func3(func(dictAlternative_0 gopurs_runtime.Value, dictLazy_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "alt"), gopurs_runtime.Apply3(Get_some(), dictAlternative_0, dictLazy_1, v_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))))
})
	})
	return many
}

var length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		length = func() gopurs_runtime.Value {
var go__0_0 gopurs_runtime.Value
go__0_0 = gopurs_runtime.Func(func(b_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__0_0:
for {
if false { continue go__0_0 }
var b_1 = b_1_loop
_ = b_1
var v_2 = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = b_1
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Cons")).IntVal != 0 {
b_1_loop = gopurs_runtime.Int(b_1.IntVal + gopurs_runtime.Int(1).IntVal)
v_2_loop = gopurs_runtime.RecordGet(v_2, "value1")
continue go__0_0
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
return gopurs_runtime.Apply(go__0_0, gopurs_runtime.Int(0))
}()
	})
	return length
}

var last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		last = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
last:
for {
if false { continue last }
var v_0 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_0, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_0, "value1"), "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordGet(v_0, "value0"))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply(Get_last(), gopurs_runtime.RecordGet(v_0, "value1"))
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_0:
return __t0
}
}()
})
	})
	return last
}

var insertBy gopurs_runtime.Value
var once_insertBy sync.Once
func Get_insertBy() gopurs_runtime.Value {
	once_insertBy.Do(func() {
		insertBy = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
insertBy:
for {
if false { continue insertBy }
var v_0 = v_0_loop
_ = v_0
var v1_1 = v1_1_loop
_ = v1_1
var v2_2 = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_2, "_tag").StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), v1_1, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_2, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(v_0, v1_1, gopurs_runtime.RecordGet(v2_2, "value0")), "_tag").StrVal == "GT")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v2_2, "value0"), gopurs_runtime.Apply3(Get_insertBy(), v_0, v1_1, gopurs_runtime.RecordGet(v2_2, "value1")))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), v1_1, v2_2)
}
end_branch_1:
__t0 = __t1
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
}()
})
})
})
	})
	return insertBy
}

var insertAt gopurs_runtime.Value
var once_insertAt sync.Once
func Get_insertAt() gopurs_runtime.Value {
	once_insertAt.Do(func() {
		insertAt = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
insertAt:
for {
if false { continue insertAt }
var v_0 = v_0_loop
_ = v_0
var v1_1 = v1_1_loop
_ = v1_1
var v2_2 = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), v1_1, v2_2))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_2, "_tag").StrVal == "Cons")).IntVal != 0 {
__local_var_3_1 := gopurs_runtime.Apply3(Get_insertAt(), gopurs_runtime.Int(v_0.IntVal - gopurs_runtime.Int(1).IntVal), v1_1, gopurs_runtime.RecordGet(v2_2, "value1"))
_ = __local_var_3_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_3_1, "_tag").StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v2_2, "value0"), gopurs_runtime.RecordGet(__local_var_3_1, "value0")))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_2:
__t0 = __t2
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_0:
return __t0
}
}()
})
})
})
	})
	return insertAt
}

var insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		insert = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_insertBy(), gopurs_runtime.RecordGet(dictOrd_0, "compare"))
})
	})
	return insert
}

var init_ gopurs_runtime.Value
var once_init_ sync.Once
func Get_init_() gopurs_runtime.Value {
	once_init_.Do(func() {
		init_ = gopurs_runtime.Func(func(lst_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(Get_unsnoc(), lst_0)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_1_0, "_tag").StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(__local_var_1_0, "value0"), "init"))
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
	return init_
}

var index gopurs_runtime.Value
var once_index sync.Once
func Get_index() gopurs_runtime.Value {
	once_index.Do(func() {
		index = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
index:
for {
if false { continue index }
var v_0 = v_0_loop
_ = v_0
var v1_1 = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_0, "_tag").StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_0, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_1.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t1 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordGet(v_0, "value0"))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply2(Get_index(), gopurs_runtime.RecordGet(v_0, "value1"), gopurs_runtime.Int(v1_1.IntVal - gopurs_runtime.Int(1).IntVal))
}
end_branch_1:
__t0 = __t1
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
}()
})
})
	})
	return index
}

var head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		head = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_0, "_tag").StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_0, "_tag").StrVal == "Cons")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordGet(v_0, "value0"))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
	})
	return head
}

var transpose gopurs_runtime.Value
var once_transpose sync.Once
func Get_transpose() gopurs_runtime.Value {
	once_transpose.Do(func() {
		transpose = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
transpose:
for {
if false { continue transpose }
var v_0 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_0, "_tag").StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_0, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_0, "value0"), "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(Get_transpose(), gopurs_runtime.RecordGet(v_0, "value1"))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_0, "value0"), "_tag").StrVal == "Cons")).IntVal != 0 {
var go__1_2 gopurs_runtime.Value
go__1_2 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_2:
for {
if false { continue go__1_2 }
var v_2 = v_2_loop
_ = v_2
var v1_3 = v1_3_loop
_ = v1_3
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Nil")).IntVal != 0 {
var go__4_4 gopurs_runtime.Value
go__4_4 = gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__4_4:
for {
if false { continue go__4_4 }
var v_5 = v_5_loop
_ = v_5
var v1_6 = v1_6_loop
_ = v1_6
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_6, "_tag").StrVal == "Nil")).IntVal != 0 {
__t5 = v_5
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_6, "_tag").StrVal == "Cons")).IntVal != 0 {
v_5_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_6, "value0"), v_5)
v1_6_loop = gopurs_runtime.RecordGet(v1_6, "value1")
continue go__4_4
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
__t3 = gopurs_runtime.Apply2(go__4_4, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")), v_2)
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "value0"), "_tag").StrVal == "Nil")).IntVal != 0 {
v_2_loop = v_2
v1_3_loop = gopurs_runtime.RecordGet(v1_3, "value1")
continue go__1_2
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "value0"), "_tag").StrVal == "Cons")).IntVal != 0 {
v_2_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "value0"), "value0"), v_2)
v1_3_loop = gopurs_runtime.RecordGet(v1_3, "value1")
continue go__1_2
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
__t3 = __t6
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
var go__1_7 gopurs_runtime.Value
go__1_7 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_7:
for {
if false { continue go__1_7 }
var v_2 = v_2_loop
_ = v_2
var v1_3 = v1_3_loop
_ = v1_3
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Nil")).IntVal != 0 {
var go__4_9 gopurs_runtime.Value
go__4_9 = gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__4_9:
for {
if false { continue go__4_9 }
var v_5 = v_5_loop
_ = v_5
var v1_6 = v1_6_loop
_ = v1_6
var __t10 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_6, "_tag").StrVal == "Nil")).IntVal != 0 {
__t10 = v_5
goto end_branch_10
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_6, "_tag").StrVal == "Cons")).IntVal != 0 {
v_5_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_6, "value0"), v_5)
v1_6_loop = gopurs_runtime.RecordGet(v1_6, "value1")
continue go__4_9
__t10 = gopurs_runtime.Value{}
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
}
}()
})
})
__t8 = gopurs_runtime.Apply2(go__4_9, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")), v_2)
goto end_branch_8
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t11 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "value0"), "_tag").StrVal == "Nil")).IntVal != 0 {
v_2_loop = v_2
v1_3_loop = gopurs_runtime.RecordGet(v1_3, "value1")
continue go__1_7
__t11 = gopurs_runtime.Value{}
goto end_branch_11
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "value0"), "_tag").StrVal == "Cons")).IntVal != 0 {
v_2_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_3, "value0"), "value1"), v_2)
v1_3_loop = gopurs_runtime.RecordGet(v1_3, "value1")
continue go__1_7
__t11 = gopurs_runtime.Value{}
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
__t8 = __t11
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
__t1 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_0, "value0"), "value0"), gopurs_runtime.Apply2(go__1_2, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")), gopurs_runtime.RecordGet(v_0, "value1"))), gopurs_runtime.Apply(Get_transpose(), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v_0, "value0"), "value1"), gopurs_runtime.Apply2(go__1_7, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")), gopurs_runtime.RecordGet(v_0, "value1")))))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__t0 = __t1
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
}()
})
	})
	return transpose
}

var groupBy gopurs_runtime.Value
var once_groupBy sync.Once
func Get_groupBy() gopurs_runtime.Value {
	once_groupBy.Do(func() {
		groupBy = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
groupBy:
for {
if false { continue groupBy }
var v_0 = v_0_loop
_ = v_0
var v1_1 = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_1, "_tag").StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_1, "_tag").StrVal == "Cons")).IntVal != 0 {
v2_2_1 := gopurs_runtime.Apply2(Get_span(), gopurs_runtime.Apply(v_0, gopurs_runtime.RecordGet(v1_1, "value0")), gopurs_runtime.RecordGet(v1_1, "value1"))
_ = v2_2_1
__t0 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("NonEmpty"), gopurs_runtime.RecordGet(v1_1, "value0"), gopurs_runtime.RecordGet(v2_2_1, "init")), gopurs_runtime.Apply2(Get_groupBy(), v_0, gopurs_runtime.RecordGet(v2_2_1, "rest")))
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
}()
})
})
	})
	return groupBy
}

var groupAllBy gopurs_runtime.Value
var once_groupAllBy sync.Once
func Get_groupAllBy() gopurs_runtime.Value {
	once_groupAllBy.Do(func() {
		groupAllBy = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(Get_groupBy(), gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(p_0, x_1, y_2), "_tag").StrVal == "EQ")
}))
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(Get_sortBy(), p_0)
_ = __local_var_2_1
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply(__local_var_2_1, x_3))
})
})
	})
	return groupAllBy
}

var group gopurs_runtime.Value
var once_group sync.Once
func Get_group() gopurs_runtime.Value {
	once_group.Do(func() {
		group = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_groupBy(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
})
	})
	return group
}

var groupAll gopurs_runtime.Value
var once_groupAll sync.Once
func Get_groupAll() gopurs_runtime.Value {
	once_groupAll.Do(func() {
		groupAll = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(Get_group(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
_ = __local_var_1_0
compare_2_1 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_2_1
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply2(Get_sortBy(), compare_2_1, x_3))
})
})
	})
	return groupAll
}

var fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		fromFoldable = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")))
})
	})
	return fromFoldable
}

var foldM gopurs_runtime.Value
var once_foldM sync.Once
func Get_foldM() gopurs_runtime.Value {
	once_foldM.Do(func() {
		foldM = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
foldM:
for {
if false { continue foldM }
var dictMonad_0 = dictMonad_0_loop
_ = dictMonad_0
var v_1 = v_1_loop
_ = v_1
var v1_2 = v1_2_loop
_ = v1_2
var v2_3 = v2_3_loop
_ = v2_3
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_3, "_tag").StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), v1_2)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_3, "_tag").StrVal == "Cons")).IntVal != 0 {
__local_var_4_1 := gopurs_runtime.RecordGet(v2_3, "value1")
_ = __local_var_4_1
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(v_1, v1_2, gopurs_runtime.RecordGet(v2_3, "value0")), gopurs_runtime.Func(func(b_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(Get_foldM(), dictMonad_0, v_1, b_prime_5, __local_var_4_1)
}))
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
}()
})
})
})
})
	})
	return foldM
}

var findIndex gopurs_runtime.Value
var once_findIndex sync.Once
func Get_findIndex() gopurs_runtime.Value {
	once_findIndex.Do(func() {
		findIndex = gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 = v_2_loop
_ = v_2
var v1_3 = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(fn_0, gopurs_runtime.RecordGet(v1_3, "value0"))).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), v_2)
goto end_branch_2
} else {

}
}
{
v_2_loop = gopurs_runtime.Int(v_2.IntVal + gopurs_runtime.Int(1).IntVal)
v1_3_loop = gopurs_runtime.RecordGet(v1_3, "value1")
continue go__1_0
__t2 = gopurs_runtime.Value{}
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Nil")).IntVal != 0 {
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
}
}()
})
})
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.Int(0))
})
	})
	return findIndex
}

var findLastIndex gopurs_runtime.Value
var once_findLastIndex sync.Once
func Get_findLastIndex() gopurs_runtime.Value {
	once_findLastIndex.Do(func() {
		findLastIndex = gopurs_runtime.Func2(func(fn_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_0:
for {
if false { continue go__2_0 }
var v_3 = v_3_loop
_ = v_3
var v1_4 = v1_4_loop
_ = v1_4
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_4, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(fn_0, gopurs_runtime.RecordGet(v1_4, "value0"))).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), v_3)
goto end_branch_2
} else {

}
}
{
v_3_loop = gopurs_runtime.Int(v_3.IntVal + gopurs_runtime.Int(1).IntVal)
v1_4_loop = gopurs_runtime.RecordGet(v1_4, "value1")
continue go__2_0
__t2 = gopurs_runtime.Value{}
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_4, "_tag").StrVal == "Nil")).IntVal != 0 {
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
}
}()
})
})
var go__3_4 gopurs_runtime.Value
go__3_4 = gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_4:
for {
if false { continue go__3_4 }
var v_4 = v_4_loop
_ = v_4
var v1_5 = v1_5_loop
_ = v1_5
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_5, "_tag").StrVal == "Nil")).IntVal != 0 {
__t5 = v_4
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_5, "_tag").StrVal == "Cons")).IntVal != 0 {
v_4_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_5, "value0"), v_4)
v1_5_loop = gopurs_runtime.RecordGet(v1_5, "value1")
continue go__3_4
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
__local_var_3_3 := gopurs_runtime.Apply2(go__2_0, gopurs_runtime.Int(0), gopurs_runtime.Apply2(go__3_4, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")), xs_1))
_ = __local_var_3_3
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_3_3, "_tag").StrVal == "Just")).IntVal != 0 {
var go__4_7 gopurs_runtime.Value
go__4_7 = gopurs_runtime.Func(func(b_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__4_7:
for {
if false { continue go__4_7 }
var b_5 = b_5_loop
_ = b_5
var v_6 = v_6_loop
_ = v_6
var __t8 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_6, "_tag").StrVal == "Nil")).IntVal != 0 {
__t8 = b_5
goto end_branch_8
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_6, "_tag").StrVal == "Cons")).IntVal != 0 {
b_5_loop = gopurs_runtime.Int(b_5.IntVal + gopurs_runtime.Int(1).IntVal)
v_6_loop = gopurs_runtime.RecordGet(v_6, "value1")
continue go__4_7
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
__t6 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.Int(gopurs_runtime.Int(gopurs_runtime.Apply2(go__4_7, gopurs_runtime.Int(0), xs_1).IntVal - gopurs_runtime.Int(1).IntVal).IntVal - gopurs_runtime.RecordGet(__local_var_3_3, "value0").IntVal))
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_6:
return __t6
})
	})
	return findLastIndex
}

var filterM gopurs_runtime.Value
var once_filterM sync.Once
func Get_filterM() gopurs_runtime.Value {
	once_filterM.Do(func() {
		filterM = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
filterM:
for {
if false { continue filterM }
var dictMonad_0 = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_4, "_tag").StrVal == "Nil")).IntVal != 0 {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")))
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_4, "_tag").StrVal == "Cons")).IntVal != 0 {
__local_var_5_3 := gopurs_runtime.RecordGet(v1_4, "value0")
_ = __local_var_5_3
__local_var_6_4 := gopurs_runtime.RecordGet(v1_4, "value1")
_ = __local_var_6_4
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), gopurs_runtime.Apply(v_3, __local_var_5_3), gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), gopurs_runtime.Apply3(Get_filterM(), dictMonad_0, v_3, __local_var_6_4), gopurs_runtime.Func(func(xs_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (b_7).IntVal != 0 {
__t5 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), __local_var_5_3, xs_prime_8)
goto end_branch_5
} else {

}
}
{
__t5 = xs_prime_8
}
end_branch_5:
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), __t5)
}))
}))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
})
}
}()
})
	})
	return filterM
}

var filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		filter = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 = v_2_loop
_ = v_2
var v1_3 = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Nil")).IntVal != 0 {
var go__4_2 gopurs_runtime.Value
go__4_2 = gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__4_2:
for {
if false { continue go__4_2 }
var v_5 = v_5_loop
_ = v_5
var v1_6 = v1_6_loop
_ = v1_6
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_6, "_tag").StrVal == "Nil")).IntVal != 0 {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_6, "_tag").StrVal == "Cons")).IntVal != 0 {
v_5_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_6, "value0"), v_5)
v1_6_loop = gopurs_runtime.RecordGet(v1_6, "value1")
continue go__4_2
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
__t1 = gopurs_runtime.Apply2(go__4_2, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")), v_2)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_3, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(p_0, gopurs_runtime.RecordGet(v1_3, "value0"))).IntVal != 0 {
v_2_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_3, "value0"), v_2)
v1_3_loop = gopurs_runtime.RecordGet(v1_3, "value1")
continue go__1_0
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
v_2_loop = v_2
v1_3_loop = gopurs_runtime.RecordGet(v1_3, "value1")
continue go__1_0
__t4 = gopurs_runtime.Value{}
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
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")))
})
	})
	return filter
}

var intersectBy gopurs_runtime.Value
var once_intersectBy sync.Once
func Get_intersectBy() gopurs_runtime.Value {
	once_intersectBy.Do(func() {
		intersectBy = gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_1, "_tag").StrVal == "Nil")).IntVal != 0 {
__t5 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_2, "_tag").StrVal == "Nil")).IntVal != 0 {
__t5 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))
goto end_branch_5
} else {

}
}
{
var go__3_0 gopurs_runtime.Value
go__3_0 = gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_0:
for {
if false { continue go__3_0 }
var v_4 = v_4_loop
_ = v_4
var v1_5 = v1_5_loop
_ = v1_5
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_5, "_tag").StrVal == "Nil")).IntVal != 0 {
var go__6_2 gopurs_runtime.Value
go__6_2 = gopurs_runtime.Func(func(v_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__6_2:
for {
if false { continue go__6_2 }
var v_7 = v_7_loop
_ = v_7
var v1_8 = v1_8_loop
_ = v1_8
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_8, "_tag").StrVal == "Nil")).IntVal != 0 {
__t3 = v_7
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_8, "_tag").StrVal == "Cons")).IntVal != 0 {
v_7_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_8, "value0"), v_7)
v1_8_loop = gopurs_runtime.RecordGet(v1_8, "value1")
continue go__6_2
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
__t1 = gopurs_runtime.Apply2(go__6_2, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")), v_4)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_5, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_any(), gopurs_runtime.Apply(v_0, gopurs_runtime.RecordGet(v1_5, "value0")), v2_2)).IntVal != 0 {
v_4_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_5, "value0"), v_4)
v1_5_loop = gopurs_runtime.RecordGet(v1_5, "value1")
continue go__3_0
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
v_4_loop = v_4
v1_5_loop = gopurs_runtime.RecordGet(v1_5, "value1")
continue go__3_0
__t4 = gopurs_runtime.Value{}
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
__t5 = gopurs_runtime.Apply2(go__3_0, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")), v1_1)
}
end_branch_5:
return __t5
})
	})
	return intersectBy
}

var intersect gopurs_runtime.Value
var once_intersect sync.Once
func Get_intersect() gopurs_runtime.Value {
	once_intersect.Do(func() {
		intersect = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_intersectBy(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
})
	})
	return intersect
}

var nubByEq gopurs_runtime.Value
var once_nubByEq sync.Once
func Get_nubByEq() gopurs_runtime.Value {
	once_nubByEq.Do(func() {
		nubByEq = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
nubByEq:
for {
if false { continue nubByEq }
var v_0 = v_0_loop
_ = v_0
var v1_1 = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_1, "_tag").StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_1, "_tag").StrVal == "Cons")).IntVal != 0 {
__local_var_2_1 := gopurs_runtime.RecordGet(v1_1, "value0")
_ = __local_var_2_1
var go__3_2 gopurs_runtime.Value
go__3_2 = gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_2:
for {
if false { continue go__3_2 }
var v_4 = v_4_loop
_ = v_4
var v1_5 = v1_5_loop
_ = v1_5
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_5, "_tag").StrVal == "Nil")).IntVal != 0 {
var go__6_4 gopurs_runtime.Value
go__6_4 = gopurs_runtime.Func(func(v_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__6_4:
for {
if false { continue go__6_4 }
var v_7 = v_7_loop
_ = v_7
var v1_8 = v1_8_loop
_ = v1_8
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_8, "_tag").StrVal == "Nil")).IntVal != 0 {
__t5 = v_7
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_8, "_tag").StrVal == "Cons")).IntVal != 0 {
v_7_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_8, "value0"), v_7)
v1_8_loop = gopurs_runtime.RecordGet(v1_8, "value1")
continue go__6_4
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
__t3 = gopurs_runtime.Apply2(go__6_4, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")), v_4)
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_5, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Apply2(v_0, __local_var_2_1, gopurs_runtime.RecordGet(v1_5, "value0")).IntVal == 0)).IntVal != 0 {
v_4_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_5, "value0"), v_4)
v1_5_loop = gopurs_runtime.RecordGet(v1_5, "value1")
continue go__3_2
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
v_4_loop = v_4
v1_5_loop = gopurs_runtime.RecordGet(v1_5, "value1")
continue go__3_2
__t6 = gopurs_runtime.Value{}
}
end_branch_6:
__t3 = __t6
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
__t0 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), __local_var_2_1, gopurs_runtime.Apply2(Get_nubByEq(), v_0, gopurs_runtime.Apply2(go__3_2, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")), gopurs_runtime.RecordGet(v1_1, "value1"))))
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
}()
})
})
	})
	return nubByEq
}

var nubEq gopurs_runtime.Value
var once_nubEq sync.Once
func Get_nubEq() gopurs_runtime.Value {
	once_nubEq.Do(func() {
		nubEq = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_nubByEq(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
})
	})
	return nubEq
}

var eqPattern gopurs_runtime.Value
var once_eqPattern sync.Once
func Get_eqPattern() gopurs_runtime.Value {
	once_eqPattern.Do(func() {
		eqPattern = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
_ = go__3_0
go__3_0 = gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_6.IntVal == 0)).IntVal != 0 {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_4, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_5, "_tag").StrVal == "Nil").IntVal != 0 && v2_6.IntVal != 0)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_4, "_tag").StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_5, "_tag").StrVal == "Cons").IntVal != 0 && gopurs_runtime.Apply3(go__3_0, gopurs_runtime.RecordGet(v_4, "value1"), gopurs_runtime.RecordGet(v1_5, "value1"), gopurs_runtime.Bool(v2_6.IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), gopurs_runtime.RecordGet(v1_5, "value0"), gopurs_runtime.RecordGet(v_4, "value0")).IntVal != 0)).IntVal != 0).IntVal != 0)
}
end_branch_1:
return __t1
})
return gopurs_runtime.Apply3(go__3_0, x_1, y_2, gopurs_runtime.Bool(true))
}))
})
	})
	return eqPattern
}

var ordPattern gopurs_runtime.Value
var once_ordPattern sync.Once
func Get_ordPattern() gopurs_runtime.Value {
	once_ordPattern.Do(func() {
		ordPattern = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
eqPattern1_2_1 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_2 gopurs_runtime.Value
_ = go__4_2
go__4_2 = gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_7.IntVal == 0)).IntVal != 0 {
__t3 = gopurs_runtime.Bool(false)
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_5, "_tag").StrVal == "Nil")).IntVal != 0 {
__t3 = gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_6, "_tag").StrVal == "Nil").IntVal != 0 && v2_7.IntVal != 0)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_5, "_tag").StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_6, "_tag").StrVal == "Cons").IntVal != 0 && gopurs_runtime.Apply3(go__4_2, gopurs_runtime.RecordGet(v_5, "value1"), gopurs_runtime.RecordGet(v1_6, "value1"), gopurs_runtime.Bool(v2_7.IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), gopurs_runtime.RecordGet(v1_6, "value0"), gopurs_runtime.RecordGet(v_5, "value0")).IntVal != 0)).IntVal != 0).IntVal != 0)
}
end_branch_3:
return __t3
})
return gopurs_runtime.Apply3(go__4_2, x_2, y_3, gopurs_runtime.Bool(true))
}))
_ = eqPattern1_2_1
return gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(pkg_Data_List_Types.Get_ordList(), dictOrd_0), "compare"), x_3, y_4)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eqPattern1_2_1
}))
})
	})
	return ordPattern
}

var elemLastIndex gopurs_runtime.Value
var once_elemLastIndex sync.Once
func Get_elemLastIndex() gopurs_runtime.Value {
	once_elemLastIndex.Do(func() {
		elemLastIndex = gopurs_runtime.Func2(func(dictEq_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_findLastIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), v_2, x_1)
}))
})
	})
	return elemLastIndex
}

var elemIndex gopurs_runtime.Value
var once_elemIndex sync.Once
func Get_elemIndex() gopurs_runtime.Value {
	once_elemIndex.Do(func() {
		elemIndex = gopurs_runtime.Func2(func(dictEq_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_0:
for {
if false { continue go__2_0 }
var v_3 = v_3_loop
_ = v_3
var v1_4 = v1_4_loop
_ = v1_4
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_4, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), gopurs_runtime.RecordGet(v1_4, "value0"), x_1)).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), v_3)
goto end_branch_2
} else {

}
}
{
v_3_loop = gopurs_runtime.Int(v_3.IntVal + gopurs_runtime.Int(1).IntVal)
v1_4_loop = gopurs_runtime.RecordGet(v1_4, "value1")
continue go__2_0
__t2 = gopurs_runtime.Value{}
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_4, "_tag").StrVal == "Nil")).IntVal != 0 {
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
}
}()
})
})
return gopurs_runtime.Apply(go__2_0, gopurs_runtime.Int(0))
})
	})
	return elemIndex
}

var dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		dropWhile = gopurs_runtime.Func(func(p_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Cons").IntVal != 0 && gopurs_runtime.Apply(p_0, gopurs_runtime.RecordGet(v_2, "value0")).IntVal != 0)).IntVal != 0 {
v_2_loop = gopurs_runtime.RecordGet(v_2, "value1")
continue go__1_0
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
__t1 = v_2
}
end_branch_1:
return __t1
}
}()
})
return go__1_0
})
	})
	return dropWhile
}

var dropEnd gopurs_runtime.Value
var once_dropEnd sync.Once
func Get_dropEnd() gopurs_runtime.Value {
	once_dropEnd.Do(func() {
		dropEnd = gopurs_runtime.Func2(func(n_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
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
b_3_loop = gopurs_runtime.Int(b_3.IntVal + gopurs_runtime.Int(1).IntVal)
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
return gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Int(gopurs_runtime.Apply2(go__2_0, gopurs_runtime.Int(0), xs_1).IntVal - n_0.IntVal), xs_1)
})
	})
	return dropEnd
}

var drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		drop = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
drop:
for {
if false { continue drop }
var v_0 = v_0_loop
_ = v_0
var v1_1 = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.IntVal < gopurs_runtime.Int(1).IntVal)).IntVal != 0 {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_1, "_tag").StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_1, "_tag").StrVal == "Cons")).IntVal != 0 {
__t0 = gopurs_runtime.Apply2(Get_drop(), gopurs_runtime.Int(v_0.IntVal - gopurs_runtime.Int(1).IntVal), gopurs_runtime.RecordGet(v1_1, "value1"))
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
}()
})
})
	})
	return drop
}

var slice gopurs_runtime.Value
var once_slice sync.Once
func Get_slice() gopurs_runtime.Value {
	once_slice.Do(func() {
		slice = gopurs_runtime.Func3(func(start_0 gopurs_runtime.Value, end_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Int(end_1.IntVal - start_0.IntVal), gopurs_runtime.Apply2(Get_drop(), start_0, xs_2))
})
	})
	return slice
}

var takeEnd gopurs_runtime.Value
var once_takeEnd sync.Once
func Get_takeEnd() gopurs_runtime.Value {
	once_takeEnd.Do(func() {
		takeEnd = gopurs_runtime.Func2(func(n_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
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
b_3_loop = gopurs_runtime.Int(b_3.IntVal + gopurs_runtime.Int(1).IntVal)
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
return gopurs_runtime.Apply2(Get_drop(), gopurs_runtime.Int(gopurs_runtime.Apply2(go__2_0, gopurs_runtime.Int(0), xs_1).IntVal - n_0.IntVal), xs_1)
})
	})
	return takeEnd
}

var deleteBy gopurs_runtime.Value
var once_deleteBy sync.Once
func Get_deleteBy() gopurs_runtime.Value {
	once_deleteBy.Do(func() {
		deleteBy = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
deleteBy:
for {
if false { continue deleteBy }
var v_0 = v_0_loop
_ = v_0
var v1_1 = v1_1_loop
_ = v1_1
var v2_2 = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_2, "_tag").StrVal == "Nil")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_2, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(v_0, v1_1, gopurs_runtime.RecordGet(v2_2, "value0"))).IntVal != 0 {
__t1 = gopurs_runtime.RecordGet(v2_2, "value1")
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v2_2, "value0"), gopurs_runtime.Apply3(Get_deleteBy(), v_0, v1_1, gopurs_runtime.RecordGet(v2_2, "value1")))
}
end_branch_1:
__t0 = __t1
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
}()
})
})
})
	})
	return deleteBy
}

var unionBy gopurs_runtime.Value
var once_unionBy sync.Once
func Get_unionBy() gopurs_runtime.Value {
	once_unionBy.Do(func() {
		unionBy = gopurs_runtime.Func3(func(eq2_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value, ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
go__3_0 = gopurs_runtime.Func(func(b_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_0:
for {
if false { continue go__3_0 }
var b_4 = b_4_loop
_ = b_4
var v_5 = v_5_loop
_ = v_5
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_5, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = b_4
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_5, "_tag").StrVal == "Cons")).IntVal != 0 {
b_4_loop = gopurs_runtime.Apply3(Get_deleteBy(), eq2_0, gopurs_runtime.RecordGet(v_5, "value0"), b_4)
v_5_loop = gopurs_runtime.RecordGet(v_5, "value1")
continue go__3_0
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
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Apply2(go__3_0, gopurs_runtime.Apply2(Get_nubByEq(), eq2_0, ys_2), xs_1), xs_1)
})
	})
	return unionBy
}

var union gopurs_runtime.Value
var once_union sync.Once
func Get_union() gopurs_runtime.Value {
	once_union.Do(func() {
		union = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_unionBy(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
})
	})
	return union
}

var deleteAt gopurs_runtime.Value
var once_deleteAt sync.Once
func Get_deleteAt() gopurs_runtime.Value {
	once_deleteAt.Do(func() {
		deleteAt = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
deleteAt:
for {
if false { continue deleteAt }
var v_0 = v_0_loop
_ = v_0
var v1_1 = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_1, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t3 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordGet(v1_1, "value1"))
goto end_branch_3
} else {

}
}
{
__local_var_2_1 := gopurs_runtime.Apply2(Get_deleteAt(), gopurs_runtime.Int(v_0.IntVal - gopurs_runtime.Int(1).IntVal), gopurs_runtime.RecordGet(v1_1, "value1"))
_ = __local_var_2_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_2_1, "_tag").StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_1, "value0"), gopurs_runtime.RecordGet(__local_var_2_1, "value0")))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_2:
__t3 = __t2
}
end_branch_3:
__t0 = __t3
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_0:
return __t0
}
}()
})
})
	})
	return deleteAt
}

var delete_ gopurs_runtime.Value
var once_delete_ sync.Once
func Get_delete_() gopurs_runtime.Value {
	once_delete_.Do(func() {
		delete_ = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_deleteBy(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
})
	})
	return delete_
}

var difference gopurs_runtime.Value
var once_difference sync.Once
func Get_difference() gopurs_runtime.Value {
	once_difference.Do(func() {
		difference = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var b_2 = b_2_loop
_ = b_2
var v_3 = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_3, "_tag").StrVal == "Nil")).IntVal != 0 {
__t1 = b_2
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_3, "_tag").StrVal == "Cons")).IntVal != 0 {
b_2_loop = gopurs_runtime.Apply3(Get_deleteBy(), gopurs_runtime.RecordGet(dictEq_0, "eq"), gopurs_runtime.RecordGet(v_3, "value0"), b_2)
v_3_loop = gopurs_runtime.RecordGet(v_3, "value1")
continue go__1_0
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
return go__1_0
})
	})
	return difference
}

var concatMap gopurs_runtime.Value
var once_concatMap sync.Once
func Get_concatMap() gopurs_runtime.Value {
	once_concatMap.Do(func() {
		concatMap = gopurs_runtime.Func2(func(b_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_bindList(), "bind"), a_1, b_0)
})
	})
	return concatMap
}

var concat gopurs_runtime.Value
var once_concat sync.Once
func Get_concat() gopurs_runtime.Value {
	once_concat.Do(func() {
		concat = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_bindList(), "bind"), v_0, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
})
	})
	return concat
}

var catMaybes gopurs_runtime.Value
var once_catMaybes sync.Once
func Get_catMaybes() gopurs_runtime.Value {
	once_catMaybes.Do(func() {
		catMaybes = func() gopurs_runtime.Value {
var go__0_0 gopurs_runtime.Value
go__0_0 = gopurs_runtime.Func(func(v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__0_0:
for {
if false { continue go__0_0 }
var v_1 = v_1_loop
_ = v_1
var v1_2 = v1_2_loop
_ = v1_2
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_2, "_tag").StrVal == "Nil")).IntVal != 0 {
var go__3_2 gopurs_runtime.Value
go__3_2 = gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_2:
for {
if false { continue go__3_2 }
var v_4 = v_4_loop
_ = v_4
var v1_5 = v1_5_loop
_ = v1_5
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_5, "_tag").StrVal == "Nil")).IntVal != 0 {
__t3 = v_4
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_5, "_tag").StrVal == "Cons")).IntVal != 0 {
v_4_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v1_5, "value0"), v_4)
v1_5_loop = gopurs_runtime.RecordGet(v1_5, "value1")
continue go__3_2
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
__t1 = gopurs_runtime.Apply2(go__3_2, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")), v_1)
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_2, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_2, "value0"), "_tag").StrVal == "Nothing")).IntVal != 0 {
v_1_loop = v_1
v1_2_loop = gopurs_runtime.RecordGet(v1_2, "value1")
continue go__0_0
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_2, "value0"), "_tag").StrVal == "Just")).IntVal != 0 {
v_1_loop = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(gopurs_runtime.RecordGet(v1_2, "value0"), "value0"), v_1)
v1_2_loop = gopurs_runtime.RecordGet(v1_2, "value1")
continue go__0_0
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
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
return gopurs_runtime.Apply(go__0_0, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nil")))
}()
	})
	return catMaybes
}

var alterAt gopurs_runtime.Value
var once_alterAt sync.Once
func Get_alterAt() gopurs_runtime.Value {
	once_alterAt.Do(func() {
		alterAt = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
alterAt:
for {
if false { continue alterAt }
var v_0 = v_0_loop
_ = v_0
var v1_1 = v1_1_loop
_ = v1_1
var v2_2 = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_2, "_tag").StrVal == "Cons")).IntVal != 0 {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
v3_3_4 := gopurs_runtime.Apply(v1_1, gopurs_runtime.RecordGet(v2_2, "value0"))
_ = v3_3_4
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v3_3_4, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t5 = gopurs_runtime.RecordGet(v2_2, "value1")
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v3_3_4, "_tag").StrVal == "Just")).IntVal != 0 {
__t5 = gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v3_3_4, "value0"), gopurs_runtime.RecordGet(v2_2, "value1"))
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t3 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), __t5)
goto end_branch_3
} else {

}
}
{
__local_var_3_1 := gopurs_runtime.Apply3(Get_alterAt(), gopurs_runtime.Int(v_0.IntVal - gopurs_runtime.Int(1).IntVal), v1_1, gopurs_runtime.RecordGet(v2_2, "value1"))
_ = __local_var_3_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_3_1, "_tag").StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Cons"), gopurs_runtime.RecordGet(v2_2, "value0"), gopurs_runtime.RecordGet(__local_var_3_1, "value0")))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_2:
__t3 = __t2
}
end_branch_3:
__t0 = __t3
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
}
end_branch_0:
return __t0
}
}()
})
})
})
	})
	return alterAt
}

var modifyAt gopurs_runtime.Value
var once_modifyAt sync.Once
func Get_modifyAt() gopurs_runtime.Value {
	once_modifyAt.Do(func() {
		modifyAt = gopurs_runtime.Func2(func(n_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_alterAt(), n_0, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Just"), gopurs_runtime.Apply(f_1, x_2))
}))
})
	})
	return modifyAt
}



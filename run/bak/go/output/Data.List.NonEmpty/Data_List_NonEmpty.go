package Data_List_NonEmpty

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Partial "gopurs/output/Partial"
	pkg_Data_List "gopurs/output/Data.List"
)

var zipWith gopurs_runtime.Value
var once_zipWith sync.Once
func Get_zipWith() gopurs_runtime.Value {
	once_zipWith.Do(func() {
		zipWith = gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_0 gopurs_runtime.Value
go__3_0 = gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_0:
for {
if false { continue go__3_0 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var v2_6 gopurs_runtime.Value = v2_6_loop
_ = v2_6
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_4.StrVal == "Nil").IntVal != 0 {
__t1 = v2_6
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v1_5.StrVal == "Nil").IntVal != 0 {
__t1 = v2_6
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_4.StrVal == "Cons").IntVal != 0 && gopurs_runtime.Bool(v1_5.StrVal == "Cons").IntVal != 0 {
v_4_loop = (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1]
v1_5_loop = (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[1]
v2_6_loop = gopurs_runtime.Constructor2("Cons", gopurs_runtime.Apply2(f_0, (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[0]), v2_6)
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
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_6.StrVal == "Nil").IntVal != 0 {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v1_6.StrVal == "Cons").IntVal != 0 {
v_5_loop = gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[0], v_5)
v1_6_loop = (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[1]
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
return gopurs_runtime.Constructor2("NonEmpty", gopurs_runtime.Apply2(f_0, (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_2.UnsafePtr)[0]), gopurs_runtime.Apply2(go__4_2, gopurs_runtime.Constructor0("Nil"), gopurs_runtime.Apply3(go__3_0, (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_2.UnsafePtr)[1], gopurs_runtime.Constructor0("Nil"))))
})
	})
	return zipWith
}

var zipWithA gopurs_runtime.Value
var once_zipWithA sync.Once
func Get_zipWithA() gopurs_runtime.Value {
	once_zipWithA.Do(func() {
		zipWithA = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
sequence11_1_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_traversable1NonEmptyList(), "traverse1"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
_ = sequence11_1_0
return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value, ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence11_1_0, gopurs_runtime.Apply3(Get_zipWith(), f_2, xs_3, ys_4))
})
})
	})
	return zipWithA
}

var zip gopurs_runtime.Value
var once_zip sync.Once
func Get_zip() gopurs_runtime.Value {
	once_zip.Do(func() {
		zip = gopurs_runtime.Apply(Get_zipWith(), pkg_Data_Tuple.Get_Tuple())
	})
	return zip
}

var wrappedOperation2 gopurs_runtime.Value
var once_wrappedOperation2 sync.Once
func Get_wrappedOperation2() gopurs_runtime.Value {
	once_wrappedOperation2.Do(func() {
		wrappedOperation2 = gopurs_runtime.Func4(func(name_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
v2_4_0 := gopurs_runtime.Apply2(f_1, gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1]), gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[1]))
_ = v2_4_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_4_0.StrVal == "Cons").IntVal != 0 {
__t1 = gopurs_runtime.Constructor2("NonEmpty", (*[1024]gopurs_runtime.Value)(v2_4_0.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v2_4_0.UnsafePtr)[1])
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v2_4_0.StrVal == "Nil").IntVal != 0 {
__t1 = gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str("Impossible: empty list in NonEmptyList " + name_0.StrVal))
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
	return wrappedOperation2
}

var wrappedOperation gopurs_runtime.Value
var once_wrappedOperation sync.Once
func Get_wrappedOperation() gopurs_runtime.Value {
	once_wrappedOperation.Do(func() {
		wrappedOperation = gopurs_runtime.Func3(func(name_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
v1_3_0 := gopurs_runtime.Apply(f_1, gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1]))
_ = v1_3_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_3_0.StrVal == "Cons").IntVal != 0 {
__t1 = gopurs_runtime.Constructor2("NonEmpty", (*[1024]gopurs_runtime.Value)(v1_3_0.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_3_0.UnsafePtr)[1])
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v1_3_0.StrVal == "Nil").IntVal != 0 {
__t1 = gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str("Impossible: empty list in NonEmptyList " + name_0.StrVal))
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
	return wrappedOperation
}

var updateAt gopurs_runtime.Value
var once_updateAt sync.Once
func Get_updateAt() gopurs_runtime.Value {
	once_updateAt.Do(func() {
		updateAt = gopurs_runtime.Func3(func(i_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if i_0.IntVal == 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("NonEmpty", a_1, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1]))
goto end_branch_2
} else {

}
}
{
__local_var_3_0 := gopurs_runtime.Apply3(pkg_Data_List.Get_updateAt(), gopurs_runtime.Int(i_0.IntVal - 1), a_1, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1])
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_3_0.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("NonEmpty", (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0]))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
__t2 = __t1
}
end_branch_2:
return __t2
})
	})
	return updateAt
}

var unzip gopurs_runtime.Value
var once_unzip sync.Once
func Get_unzip() gopurs_runtime.Value {
	once_unzip.Do(func() {
		unzip = gopurs_runtime.Func(func(ts_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Constructor2("NonEmpty", (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(ts_0.UnsafePtr)[0].UnsafePtr)[0], gopurs_runtime.Apply2(pkg_Data_List_Types.Get_listMap(), pkg_Data_Tuple.Get_fst(), (*[1024]gopurs_runtime.Value)(ts_0.UnsafePtr)[1])), gopurs_runtime.Constructor2("NonEmpty", (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(ts_0.UnsafePtr)[0].UnsafePtr)[1], gopurs_runtime.Apply2(pkg_Data_List_Types.Get_listMap(), pkg_Data_Tuple.Get_snd(), (*[1024]gopurs_runtime.Value)(ts_0.UnsafePtr)[1])))
})
	})
	return unzip
}

var unsnoc gopurs_runtime.Value
var once_unsnoc sync.Once
func Get_unsnoc() gopurs_runtime.Value {
	once_unsnoc.Do(func() {
		unsnoc = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
v1_1_0 := gopurs_runtime.Apply(pkg_Data_List.Get_unsnoc(), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1])
_ = v1_1_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_1_0.StrVal == "Nothing").IntVal != 0 {
__t1 = gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.Constructor0("Nil"), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v1_1_0.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0], gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(v1_1_0.UnsafePtr)[0], "init")), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(v1_1_0.UnsafePtr)[0], "last"))
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
	return unsnoc
}

var unionBy gopurs_runtime.Value
var once_unionBy sync.Once
func Get_unionBy() gopurs_runtime.Value {
	once_unionBy.Do(func() {
		unionBy = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_wrappedOperation2(), gopurs_runtime.Str("unionBy"), gopurs_runtime.Apply(pkg_Data_List.Get_unionBy(), x_0))
})
	})
	return unionBy
}

var union gopurs_runtime.Value
var once_union sync.Once
func Get_union() gopurs_runtime.Value {
	once_union.Do(func() {
		union = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_wrappedOperation2(), gopurs_runtime.Str("union"), gopurs_runtime.Apply(pkg_Data_List.Get_union(), dictEq_0))
})
	})
	return union
}

var uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		uncons = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("head", "tail", (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1])
})
	})
	return uncons
}

var toList gopurs_runtime.Value
var once_toList sync.Once
func Get_toList() gopurs_runtime.Value {
	once_toList.Do(func() {
		toList = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1])
})
	})
	return toList
}

var toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(xs_1.StrVal == "Nil").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(xs_1.StrVal == "Cons").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Tuple", (*[1024]gopurs_runtime.Value)(xs_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(xs_1.UnsafePtr)[1]))
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
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(x_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(x_2.UnsafePtr)[1]))
})
})
	})
	return toUnfoldable
}

var tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		tail = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1]
})
	})
	return tail
}

var sortBy gopurs_runtime.Value
var once_sortBy sync.Once
func Get_sortBy() gopurs_runtime.Value {
	once_sortBy.Do(func() {
		sortBy = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_wrappedOperation(), gopurs_runtime.Str("sortBy"), gopurs_runtime.Apply(pkg_Data_List.Get_sortBy(), x_0))
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
return gopurs_runtime.Apply3(Get_wrappedOperation(), gopurs_runtime.Str("sortBy"), gopurs_runtime.Apply(pkg_Data_List.Get_sortBy(), compare_1_0), xs_2)
})
})
	})
	return sort
}

var snoc gopurs_runtime.Value
var once_snoc sync.Once
func Get_snoc() gopurs_runtime.Value {
	once_snoc.Do(func() {
		snoc = gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0], gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Constructor2("Cons", y_1, gopurs_runtime.Constructor0("Nil")), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1]))
})
	})
	return snoc
}

var singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		singleton = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", x_0, gopurs_runtime.Constructor0("Nil"))
})
	})
	return singleton
}

var snoc_prime gopurs_runtime.Value
var once_snoc_prime sync.Once
func Get_snoc_prime() gopurs_runtime.Value {
	once_snoc_prime.Do(func() {
		snoc_prime = gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_0.StrVal == "Cons").IntVal != 0 {
__t0 = gopurs_runtime.Constructor2("NonEmpty", (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0], gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Constructor2("Cons", v1_1, gopurs_runtime.Constructor0("Nil")), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1]))
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "Nil").IntVal != 0 {
__t0 = gopurs_runtime.Constructor2("NonEmpty", v1_1, gopurs_runtime.Constructor0("Nil"))
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
	return snoc_prime
}

var reverse gopurs_runtime.Value
var once_reverse sync.Once
func Get_reverse() gopurs_runtime.Value {
	once_reverse.Do(func() {
		reverse = gopurs_runtime.Apply2(Get_wrappedOperation(), gopurs_runtime.Str("reverse"), pkg_Data_List.Get_reverse())
	})
	return reverse
}

var nubEq gopurs_runtime.Value
var once_nubEq sync.Once
func Get_nubEq() gopurs_runtime.Value {
	once_nubEq.Do(func() {
		nubEq = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_wrappedOperation(), gopurs_runtime.Str("nubEq"), gopurs_runtime.Apply(pkg_Data_List.Get_nubEq(), dictEq_0))
})
	})
	return nubEq
}

var nubByEq gopurs_runtime.Value
var once_nubByEq sync.Once
func Get_nubByEq() gopurs_runtime.Value {
	once_nubByEq.Do(func() {
		nubByEq = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_wrappedOperation(), gopurs_runtime.Str("nubByEq"), gopurs_runtime.Apply(pkg_Data_List.Get_nubByEq(), x_0))
})
	})
	return nubByEq
}

var nubBy gopurs_runtime.Value
var once_nubBy sync.Once
func Get_nubBy() gopurs_runtime.Value {
	once_nubBy.Do(func() {
		nubBy = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_wrappedOperation(), gopurs_runtime.Str("nubBy"), gopurs_runtime.Apply(pkg_Data_List.Get_nubBy(), x_0))
})
	})
	return nubBy
}

var nub gopurs_runtime.Value
var once_nub sync.Once
func Get_nub() gopurs_runtime.Value {
	once_nub.Do(func() {
		nub = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_wrappedOperation(), gopurs_runtime.Str("nub"), gopurs_runtime.Apply(pkg_Data_List.Get_nubBy(), gopurs_runtime.RecordGet(dictOrd_0, "compare")))
})
	})
	return nub
}

var modifyAt gopurs_runtime.Value
var once_modifyAt sync.Once
func Get_modifyAt() gopurs_runtime.Value {
	once_modifyAt.Do(func() {
		modifyAt = gopurs_runtime.Func3(func(i_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if i_0.IntVal == 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("NonEmpty", gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1]))
goto end_branch_2
} else {

}
}
{
__local_var_3_0 := gopurs_runtime.Apply3(pkg_Data_List.Get_alterAt(), gopurs_runtime.Int(i_0.IntVal - 1), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(f_1, x_3))
}), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1])
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_3_0.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("NonEmpty", (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0]))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
__t2 = __t1
}
end_branch_2:
return __t2
})
	})
	return modifyAt
}

var mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		mapMaybe = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_3.StrVal == "Nil").IntVal != 0 {
var go__4_2 gopurs_runtime.Value
go__4_2 = gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__4_2:
for {
if false { continue go__4_2 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_6.StrVal == "Nil").IntVal != 0 {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v1_6.StrVal == "Cons").IntVal != 0 {
v_5_loop = gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[0], v_5)
v1_6_loop = (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[1]
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
__t1 = gopurs_runtime.Apply2(go__4_2, gopurs_runtime.Constructor0("Nil"), v_2)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v1_3.StrVal == "Cons").IntVal != 0 {
v2_4_4 := gopurs_runtime.Apply(x_0, (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[0])
_ = v2_4_4
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_4_4.StrVal == "Nothing").IntVal != 0 {
v_2_loop = v_2
v1_3_loop = (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[1]
continue go__1_0
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool(v2_4_4.StrVal == "Just").IntVal != 0 {
v_2_loop = gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v2_4_4.UnsafePtr)[0], v_2)
v1_3_loop = (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[1]
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
__local_var_2_6 := gopurs_runtime.Apply(go__1_0, gopurs_runtime.Constructor0("Nil"))
_ = __local_var_2_6
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_6, gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[1]))
})
})
	})
	return mapMaybe
}

var partition gopurs_runtime.Value
var once_partition sync.Once
func Get_partition() gopurs_runtime.Value {
	once_partition.Do(func() {
		partition = gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Data_List.Get_partition(), x_0, gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1]))
})
	})
	return partition
}

var span gopurs_runtime.Value
var once_span sync.Once
func Get_span() gopurs_runtime.Value {
	once_span.Do(func() {
		span = gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Data_List.Get_span(), x_0, gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1]))
})
	})
	return span
}

var take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		take = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_List.Get_take(), x_0)
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1]))
})
})
	})
	return take
}

var takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		takeWhile = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_3.StrVal == "Cons").IntVal != 0 && gopurs_runtime.Apply(x_0, (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[0]).IntVal != 0 {
v_2_loop = gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[0], v_2)
v1_3_loop = (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[1]
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
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_6.StrVal == "Nil").IntVal != 0 {
__t2 = v_5
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(v1_6.StrVal == "Cons").IntVal != 0 {
v_5_loop = gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[0], v_5)
v1_6_loop = (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[1]
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
__t3 = gopurs_runtime.Apply2(go__4_1, gopurs_runtime.Constructor0("Nil"), v_2)
}
end_branch_3:
return __t3
}
}()
})
})
__local_var_2_4 := gopurs_runtime.Apply(go__1_0, gopurs_runtime.Constructor0("Nil"))
_ = __local_var_2_4
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_4, gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[1]))
})
})
	})
	return takeWhile
}

var length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		length = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_3.StrVal == "Nil").IntVal != 0 {
__t1 = b_2
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v_3.StrVal == "Cons").IntVal != 0 {
b_2_loop = gopurs_runtime.Int(b_2.IntVal + 1)
v_3_loop = (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[1]
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
return 1 + gopurs_runtime.Apply2(go__1_0, gopurs_runtime.Int(0), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1]).IntVal
})
	})
	return length
}

var last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		last = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].StrVal == "Cons").IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].UnsafePtr)[1].StrVal == "Nil").IntVal != 0 {
__t1 = (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].UnsafePtr)[0]
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(gopurs_runtime.Apply(pkg_Data_List.Get_last(), (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].UnsafePtr)[1]).StrVal == "Nothing").IntVal != 0 {
__t1 = (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0]
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(gopurs_runtime.Apply(pkg_Data_List.Get_last(), (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].UnsafePtr)[1]).StrVal == "Just").IntVal != 0 {
__t1 = (*[1024]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Data_List.Get_last(), (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1].UnsafePtr)[1]).UnsafePtr)[0]
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
__t0 = (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0]
}
end_branch_0:
return __t0
})
	})
	return last
}

var intersectBy gopurs_runtime.Value
var once_intersectBy sync.Once
func Get_intersectBy() gopurs_runtime.Value {
	once_intersectBy.Do(func() {
		intersectBy = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_wrappedOperation2(), gopurs_runtime.Str("intersectBy"), gopurs_runtime.Apply(pkg_Data_List.Get_intersectBy(), x_0))
})
	})
	return intersectBy
}

var intersect gopurs_runtime.Value
var once_intersect sync.Once
func Get_intersect() gopurs_runtime.Value {
	once_intersect.Do(func() {
		intersect = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_wrappedOperation2(), gopurs_runtime.Str("intersect"), gopurs_runtime.Apply(pkg_Data_List.Get_intersect(), dictEq_0))
})
	})
	return intersect
}

var insertAt gopurs_runtime.Value
var once_insertAt sync.Once
func Get_insertAt() gopurs_runtime.Value {
	once_insertAt.Do(func() {
		insertAt = gopurs_runtime.Func3(func(i_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if i_0.IntVal == 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("NonEmpty", a_1, gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1])))
goto end_branch_2
} else {

}
}
{
__local_var_3_0 := gopurs_runtime.Apply3(pkg_Data_List.Get_insertAt(), gopurs_runtime.Int(i_0.IntVal - 1), a_1, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1])
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_3_0.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("NonEmpty", (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0]))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
__t2 = __t1
}
end_branch_2:
return __t2
})
	})
	return insertAt
}

var init_ gopurs_runtime.Value
var once_init_ sync.Once
func Get_init_() gopurs_runtime.Value {
	once_init_.Do(func() {
		init_ = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_List.Get_unsnoc(), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1])
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_1_0.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0], gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_1_0.UnsafePtr)[0], "init"))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nil")
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
		index = gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, i_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if i_1.IntVal == 0 {
__t0 = gopurs_runtime.Constructor1("Just", (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply2(pkg_Data_List.Get_index(), (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1], gopurs_runtime.Int(i_1.IntVal - 1))
}
end_branch_0:
return __t0
})
	})
	return index
}

var head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		head = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0]
})
	})
	return head
}

var groupBy gopurs_runtime.Value
var once_groupBy sync.Once
func Get_groupBy() gopurs_runtime.Value {
	once_groupBy.Do(func() {
		groupBy = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_wrappedOperation(), gopurs_runtime.Str("groupBy"), gopurs_runtime.Apply(pkg_Data_List.Get_groupBy(), x_0))
})
	})
	return groupBy
}

var groupAllBy gopurs_runtime.Value
var once_groupAllBy sync.Once
func Get_groupAllBy() gopurs_runtime.Value {
	once_groupAllBy.Do(func() {
		groupAllBy = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_wrappedOperation(), gopurs_runtime.Str("groupAllBy"), gopurs_runtime.Apply(pkg_Data_List.Get_groupAllBy(), x_0))
})
	})
	return groupAllBy
}

var groupAll gopurs_runtime.Value
var once_groupAll sync.Once
func Get_groupAll() gopurs_runtime.Value {
	once_groupAll.Do(func() {
		groupAll = gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_wrappedOperation(), gopurs_runtime.Str("groupAll"), gopurs_runtime.Apply(pkg_Data_List.Get_groupAll(), dictOrd_0))
})
	})
	return groupAll
}

var group gopurs_runtime.Value
var once_group sync.Once
func Get_group() gopurs_runtime.Value {
	once_group.Do(func() {
		group = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_wrappedOperation(), gopurs_runtime.Str("group"), gopurs_runtime.Apply(pkg_Data_List.Get_group(), dictEq_0))
})
	})
	return group
}

var fromList gopurs_runtime.Value
var once_fromList sync.Once
func Get_fromList() gopurs_runtime.Value {
	once_fromList.Do(func() {
		fromList = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_0.StrVal == "Nil").IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_0.StrVal == "Cons").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("NonEmpty", (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1]))
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
	return fromList
}

var fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		fromFoldable = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Constructor0("Nil"))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
_ = __local_var_3_1
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_3_1.StrVal == "Nil").IntVal != 0 {
__t2 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_3_1.StrVal == "Cons").IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("NonEmpty", (*[1024]gopurs_runtime.Value)(__local_var_3_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(__local_var_3_1.UnsafePtr)[1]))
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
})
	})
	return fromFoldable
}

var foldM gopurs_runtime.Value
var once_foldM sync.Once
func Get_foldM() gopurs_runtime.Value {
	once_foldM.Do(func() {
		foldM = gopurs_runtime.Func4(func(dictMonad_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[1]
_ = __local_var_4_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(f_1, b_2, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0]), gopurs_runtime.Func(func(b_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(pkg_Data_List.Get_foldM(), dictMonad_0, f_1, b_prime_5, __local_var_4_0)
}))
})
	})
	return foldM
}

var findLastIndex gopurs_runtime.Value
var once_findLastIndex sync.Once
func Get_findLastIndex() gopurs_runtime.Value {
	once_findLastIndex.Do(func() {
		findLastIndex = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
v1_2_0 := gopurs_runtime.Apply2(pkg_Data_List.Get_findLastIndex(), f_0, (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1])
_ = v1_2_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_2_0.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Int((*[1024]gopurs_runtime.Value)(v1_2_0.UnsafePtr)[0].IntVal + 1))
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v1_2_0.StrVal == "Nothing").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Apply(f_0, (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0]).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Int(0))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Constructor0("Nothing")
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
})
	})
	return findLastIndex
}

var findIndex gopurs_runtime.Value
var once_findIndex sync.Once
func Get_findIndex() gopurs_runtime.Value {
	once_findIndex.Do(func() {
		findIndex = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Apply(f_0, (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0]).IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Int(0))
goto end_branch_5
} else {

}
}
{
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_0:
for {
if false { continue go__2_0 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var v1_4 gopurs_runtime.Value = v1_4_loop
_ = v1_4
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_4.StrVal == "Cons").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Apply(f_0, (*[1024]gopurs_runtime.Value)(v1_4.UnsafePtr)[0]).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", v_3)
goto end_branch_2
} else {

}
}
{
v_3_loop = gopurs_runtime.Int(v_3.IntVal + 1)
v1_4_loop = (*[1024]gopurs_runtime.Value)(v1_4.UnsafePtr)[1]
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
if gopurs_runtime.Bool(v1_4.StrVal == "Nil").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Nothing")
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
__local_var_3_3 := gopurs_runtime.Apply2(go__2_0, gopurs_runtime.Int(0), (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1])
_ = __local_var_3_3
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_3_3.StrVal == "Just").IntVal != 0 {
__t4 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Int((*[1024]gopurs_runtime.Value)(__local_var_3_3.UnsafePtr)[0].IntVal + 1))
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_4:
__t5 = __t4
}
end_branch_5:
return __t5
})
	})
	return findIndex
}

var filterM gopurs_runtime.Value
var once_filterM sync.Once
func Get_filterM() gopurs_runtime.Value {
	once_filterM.Do(func() {
		filterM = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_List.Get_filterM(), dictMonad_0)
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1]))
})
})
})
	})
	return filterM
}

var filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		filter = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_3.StrVal == "Nil").IntVal != 0 {
var go__4_2 gopurs_runtime.Value
go__4_2 = gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__4_2:
for {
if false { continue go__4_2 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_6.StrVal == "Nil").IntVal != 0 {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v1_6.StrVal == "Cons").IntVal != 0 {
v_5_loop = gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[0], v_5)
v1_6_loop = (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[1]
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
__t1 = gopurs_runtime.Apply2(go__4_2, gopurs_runtime.Constructor0("Nil"), v_2)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v1_3.StrVal == "Cons").IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Apply(x_0, (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[0]).IntVal != 0 {
v_2_loop = gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[0], v_2)
v1_3_loop = (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[1]
continue go__1_0
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
v_2_loop = v_2
v1_3_loop = (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[1]
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
__local_var_2_5 := gopurs_runtime.Apply(go__1_0, gopurs_runtime.Constructor0("Nil"))
_ = __local_var_2_5
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_5, gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[1]))
})
})
	})
	return filter
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
		elemIndex = gopurs_runtime.Func3(func(dictEq_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0], x_1).IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Int(0))
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
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_5.StrVal == "Cons").IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[0], x_1).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", v_4)
goto end_branch_2
} else {

}
}
{
v_4_loop = gopurs_runtime.Int(v_4.IntVal + 1)
v1_5_loop = (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[1]
continue go__3_0
__t2 = gopurs_runtime.Value{}
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v1_5.StrVal == "Nil").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Nothing")
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
__local_var_4_3 := gopurs_runtime.Apply2(go__3_0, gopurs_runtime.Int(0), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1])
_ = __local_var_4_3
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_4_3.StrVal == "Just").IntVal != 0 {
__t4 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Int((*[1024]gopurs_runtime.Value)(__local_var_4_3.UnsafePtr)[0].IntVal + 1))
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_4:
__t5 = __t4
}
end_branch_5:
return __t5
})
	})
	return elemIndex
}

var dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		dropWhile = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_2.StrVal == "Cons").IntVal != 0 && gopurs_runtime.Apply(x_0, (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0]).IntVal != 0 {
v_2_loop = (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1]
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
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1]))
})
})
	})
	return dropWhile
}

var drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		drop = gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Data_List.Get_drop(), x_0, gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1]))
})
	})
	return drop
}

var cons_prime gopurs_runtime.Value
var once_cons_prime sync.Once
func Get_cons_prime() gopurs_runtime.Value {
	once_cons_prime.Do(func() {
		cons_prime = gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", x_0, xs_1)
})
	})
	return cons_prime
}

var cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cons = gopurs_runtime.Func2(func(y_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", y_0, gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1]))
})
	})
	return cons
}

var concatMap gopurs_runtime.Value
var once_concatMap sync.Once
func Get_concatMap() gopurs_runtime.Value {
	once_concatMap.Do(func() {
		concatMap = gopurs_runtime.Func2(func(b_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_bindNonEmptyList(), "bind"), a_1, b_0)
})
	})
	return concatMap
}

var concat gopurs_runtime.Value
var once_concat sync.Once
func Get_concat() gopurs_runtime.Value {
	once_concat.Do(func() {
		concat = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_bindNonEmptyList(), "bind"), v_0, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
})
	})
	return concat
}

var catMaybes gopurs_runtime.Value
var once_catMaybes sync.Once
func Get_catMaybes() gopurs_runtime.Value {
	once_catMaybes.Do(func() {
		catMaybes = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_3.StrVal == "Nil").IntVal != 0 {
var go__4_2 gopurs_runtime.Value
go__4_2 = gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__4_2:
for {
if false { continue go__4_2 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t3 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_6.StrVal == "Nil").IntVal != 0 {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if gopurs_runtime.Bool(v1_6.StrVal == "Cons").IntVal != 0 {
v_5_loop = gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[0], v_5)
v1_6_loop = (*[1024]gopurs_runtime.Value)(v1_6.UnsafePtr)[1]
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
__t1 = gopurs_runtime.Apply2(go__4_2, gopurs_runtime.Constructor0("Nil"), v_2)
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v1_3.StrVal == "Cons").IntVal != 0 {
var __t4 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[0].StrVal == "Nothing").IntVal != 0 {
v_2_loop = v_2
v1_3_loop = (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[1]
continue go__1_0
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[0].StrVal == "Just").IntVal != 0 {
v_2_loop = gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)((*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[0].UnsafePtr)[0], v_2)
v1_3_loop = (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[1]
continue go__1_0
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
return gopurs_runtime.Apply2(go__1_0, gopurs_runtime.Constructor0("Nil"), gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1]))
})
	})
	return catMaybes
}

var appendFoldable gopurs_runtime.Value
var once_appendFoldable sync.Once
func Get_appendFoldable() gopurs_runtime.Value {
	once_appendFoldable.Do(func() {
		appendFoldable = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
fromFoldable1_1_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Constructor0("Nil"))
_ = fromFoldable1_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("NonEmpty", (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0], gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Apply(fromFoldable1_1_0, ys_3), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1]))
})
})
	})
	return appendFoldable
}





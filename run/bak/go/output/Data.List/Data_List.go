package Data_List

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_List_Internal "gopurs/output/Data.List.Internal"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	unsafe "unsafe"
)

var cache_any gopurs_runtime.Value
var once_any sync.Once
func Get_any() gopurs_runtime.Value {
	once_any.Do(func() {
		cache_any = func() gopurs_runtime.Value {
semigroupDisj1_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((v_0.IntVal) != (0)) || ((v1_1.IntVal) != (0)))
}))
_ = semigroupDisj1_0_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldMap"), gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Bool(false), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_0_0
})))
}()
	})
	return cache_any
}

var cache_Pattern gopurs_runtime.Value
var once_Pattern sync.Once
func Get_Pattern() gopurs_runtime.Value {
	once_Pattern.Do(func() {
		cache_Pattern = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return cache_Pattern
}

var cache_updateAt gopurs_runtime.Value
var once_updateAt sync.Once
func Get_updateAt() gopurs_runtime.Value {
	once_updateAt.Do(func() {
		cache_updateAt = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_updateAt(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_updateAt
}

var cache_unzip gopurs_runtime.Value
var once_unzip sync.Once
func Get_unzip() gopurs_runtime.Value {
	once_unzip.Do(func() {
		cache_unzip = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_0.UnsafePtr).V0
_ = __local_var_1_0
__local_var_2_1 := (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_0.UnsafePtr).V1
_ = __local_var_2_1
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{__local_var_1_0, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v1_3.UnsafePtr).V0})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{__local_var_2_1, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v1_3.UnsafePtr).V1})}})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}})})
	})
	return cache_unzip
}

var cache_uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		cache_uncons = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 786377863) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.RecordDict2("head", "tail", (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}()
})
	})
	return cache_uncons
}

var cache_toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		cache_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (xs_1.Type == 9 && xs_1.IntVal == 786377863) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_0
} else {

}
}
{
if (xs_1.Type == 9 && xs_1.IntVal == 1358893437) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(xs_1.UnsafePtr).V0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(xs_1.UnsafePtr).V1})}})}
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
}()
})
	})
	return cache_toUnfoldable
}

var cache_tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		cache_tail = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 786377863) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V1})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}()
})
	})
	return cache_tail
}

var cache_stripPrefix gopurs_runtime.Value
var once_stripPrefix sync.Once
func Get_stripPrefix() gopurs_runtime.Value {
	once_stripPrefix.Do(func() {
		cache_stripPrefix = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stripPrefix(dictEq_0_box, v_1_box, s_2_box)
})
	})
	return cache_stripPrefix
}

var cache_span gopurs_runtime.Value
var once_span sync.Once
func Get_span() gopurs_runtime.Value {
	once_span.Do(func() {
		cache_span = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_span(v_0_box, v1_1_box)
})
	})
	return cache_span
}

var cache_snoc gopurs_runtime.Value
var once_snoc sync.Once
func Get_snoc() gopurs_runtime.Value {
	once_snoc.Do(func() {
		cache_snoc = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_snoc(xs_0_box, x_1_box)
})
	})
	return cache_snoc
}

var cache_singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		cache_singleton = gopurs_runtime.Func(func(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{a_0, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}})}
}()
})
	})
	return cache_singleton
}

var cache_sortBy gopurs_runtime.Value
var once_sortBy sync.Once
func Get_sortBy() gopurs_runtime.Value {
	once_sortBy.Do(func() {
		cache_sortBy = gopurs_runtime.Func(func(cmp_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
var merge_1_0 gopurs_runtime.Value
_ = merge_1_0
merge_1_0 = gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437) {
var __t2 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437) {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(cmp_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_2.UnsafePtr).V0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0).Type == 9 && gopurs_runtime.Apply2(cmp_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_2.UnsafePtr).V0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0).IntVal == 380165415) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0, gopurs_runtime.Apply2(merge_1_0, v_2, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(merge_1_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_2.UnsafePtr).V1, v1_3)})}
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 786377863) {
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
if (v_2.Type == 9 && v_2.IntVal == 786377863) {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 786377863) {
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
if ((v_3.Type == 9 && v_3.IntVal == 1358893437)) && (((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_3.UnsafePtr).V1.Type == 9 && (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_3.UnsafePtr).V1.IntVal == 1358893437)) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Apply2(merge_1_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_3.UnsafePtr).V0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_3.UnsafePtr).V1.UnsafePtr).V0), gopurs_runtime.Apply(mergePairs_2_4, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_3.UnsafePtr).V1.UnsafePtr).V1)})}
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
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t7 gopurs_runtime.Value
{
if ((v_4.Type == 9 && v_4.IntVal == 1358893437)) && (((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V1.Type == 9 && (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V1.IntVal == 786377863)) {
__t7 = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0
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
if ((v_5.Type == 9 && v_5.IntVal == 1358893437)) && (((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_5.UnsafePtr).V1.Type == 9 && (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_5.UnsafePtr).V1.IntVal == 1358893437)) {
var __t13 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(cmp_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_5.UnsafePtr).V0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_5.UnsafePtr).V1.UnsafePtr).V0).Type == 9 && gopurs_runtime.Apply2(cmp_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_5.UnsafePtr).V0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_5.UnsafePtr).V1.UnsafePtr).V0).IntVal == 380165415) {
__t13 = gopurs_runtime.Apply3(descending_4_9, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_5.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_5.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}})}, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_5.UnsafePtr).V1.UnsafePtr).V1)
goto end_branch_13
} else {

}
}
{
__local_var_6_12 := (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_5.UnsafePtr).V0
_ = __local_var_6_12
__t13 = gopurs_runtime.Apply3(ascending_4_10, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_5.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{__local_var_6_12, v1_7})}
}), (*pkg_Data_List_Types.Data_Data_List_Types_Cons)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_5.UnsafePtr).V1.UnsafePtr).V1)
}
end_branch_13:
__t11 = __t13
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{v_5, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}})}
}
end_branch_11:
return __t11
})
descending_4_9 = gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 gopurs_runtime.Value
{
if ((v2_7.Type == 9 && v2_7.IntVal == 1358893437)) && ((gopurs_runtime.Apply2(cmp_0, v_5, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_7.UnsafePtr).V0).Type == 9 && gopurs_runtime.Apply2(cmp_0, v_5, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_7.UnsafePtr).V0).IntVal == 380165415)) {
__t14 = gopurs_runtime.Apply3(descending_4_9, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_7.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{v_5, v1_6})}, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_7.UnsafePtr).V1)
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{v_5, v1_6})}, gopurs_runtime.Apply(sequences_4_8, v2_7)})}
}
end_branch_14:
return __t14
})
ascending_4_10 = gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
__local_var_8_16 := gopurs_runtime.Apply2(cmp_0, v_5, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_7.UnsafePtr).V0)
_ = __local_var_8_16
if ((v2_7.Type == 9 && v2_7.IntVal == 1358893437)) && (((__local_var_8_16.Type == 9 && __local_var_8_16.IntVal == 1527465420)) || (((__local_var_8_16.Type == 9 && __local_var_8_16.IntVal == 380165415)) != (true))) {
__t15 = gopurs_runtime.Apply3(ascending_4_10, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_7.UnsafePtr).V0, gopurs_runtime.Func(func(ys_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{v_5, ys_8})})
}), (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_7.UnsafePtr).V1)
goto end_branch_15
} else {

}
}
{
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Apply(v1_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{v_5, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}})}), gopurs_runtime.Apply(sequences_4_8, v2_7)})}
}
end_branch_15:
return __t15
})
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(mergeAll_3_6, gopurs_runtime.Apply(sequences_4_8, x_5))
})
}()
})
	})
	return cache_sortBy
}

var cache_sort gopurs_runtime.Value
var once_sort sync.Once
func Get_sort() gopurs_runtime.Value {
	once_sort.Do(func() {
		cache_sort = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_1_0
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_sortBy(), compare_1_0, xs_2)
})
}()
})
	})
	return cache_sort
}

var cache_showPattern gopurs_runtime.Value
var once_showPattern sync.Once
func Get_showPattern() gopurs_runtime.Value {
	once_showPattern.Do(func() {
		cache_showPattern = gopurs_runtime.Func(func(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Pattern ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(pkg_Data_List_Types.Get_showList(), dictShow_0), "show"), v_1).StrVal())) + (")"))
}))
}()
})
	})
	return cache_showPattern
}

var cache_reverse gopurs_runtime.Value
var once_reverse sync.Once
func Get_reverse() gopurs_runtime.Value {
	once_reverse.Do(func() {
		cache_reverse = func() gopurs_runtime.Value {
var go__0_0 gopurs_runtime.Value
go__0_0 = gopurs_runtime.Func(func(v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__0_0:
for {
if false { continue go__0_0 }
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t1 gopurs_runtime.Value
{
if (v1_2.Type == 9 && v1_2.IntVal == 786377863) {
__t1 = v_1
goto end_branch_1
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437) {
v_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_2.UnsafePtr).V0, v_1})}
v1_2_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_2.UnsafePtr).V1
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
return gopurs_runtime.Apply(go__0_0, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})})
}()
	})
	return cache_reverse
}

var cache_take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		cache_take = func() gopurs_runtime.Value {
var go__0_0 gopurs_runtime.Value
go__0_0 = gopurs_runtime.Func(func(v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__0_0:
for {
if false { continue go__0_0 }
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var v2_3 gopurs_runtime.Value = v2_3_loop
_ = v2_3
var __t1 gopurs_runtime.Value
{
if (v1_2.IntVal) < (1) {
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
if (v1_6.Type == 9 && v1_6.IntVal == 786377863) {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})}
v1_6_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V1
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
__t1 = gopurs_runtime.Apply2(go__4_2, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, v_1)
goto end_branch_1
} else {

}
}
{
if (v2_3.Type == 9 && v2_3.IntVal == 786377863) {
var go__4_4 gopurs_runtime.Value
go__4_4 = gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__4_4:
for {
if false { continue go__4_4 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t5 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 786377863) {
__t5 = v_5
goto end_branch_5
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})}
v1_6_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V1
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
__t1 = gopurs_runtime.Apply2(go__4_4, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, v_1)
goto end_branch_1
} else {

}
}
{
if (v2_3.Type == 9 && v2_3.IntVal == 1358893437) {
v_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_3.UnsafePtr).V0, v_1})}
v1_2_loop = gopurs_runtime.Int((v1_2.IntVal) - (1))
v2_3_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_3.UnsafePtr).V1
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
return gopurs_runtime.Apply(go__0_0, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})})
}()
	})
	return cache_take
}

var cache_takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		cache_takeWhile = gopurs_runtime.Func(func(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
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
if ((v1_3.Type == 9 && v1_3.IntVal == 1358893437)) && ((gopurs_runtime.Apply(p_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0).IntVal) != (0)) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0, v_2})}
v1_3_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V1
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
if (v1_6.Type == 9 && v1_6.IntVal == 786377863) {
__t2 = v_5
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})}
v1_6_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V1
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
__t3 = gopurs_runtime.Apply2(go__4_1, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, v_2)
}
end_branch_3:
return __t3
}
}()
})
})
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})})
}()
})
	})
	return cache_takeWhile
}

var cache_unsnoc gopurs_runtime.Value
var once_unsnoc sync.Once
func Get_unsnoc() gopurs_runtime.Value {
	once_unsnoc.Do(func() {
		cache_unsnoc = gopurs_runtime.Func(func(lst_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var lst_0 gopurs_runtime.Value = lst_0_loop
_ = lst_0
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
if (v_2.Type == 9 && v_2.IntVal == 786377863) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437) {
var __t2 gopurs_runtime.Value
{
if ((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_2.UnsafePtr).V1.Type == 9 && (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_2.UnsafePtr).V1.IntVal == 786377863) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.RecordDict2("revInit", "last", v1_3, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_2.UnsafePtr).V0)})}
goto end_branch_2
} else {

}
}
{
v_2_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_2.UnsafePtr).V1
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_2.UnsafePtr).V0, v1_3})}
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
__local_var_2_3 := gopurs_runtime.Apply2(go__1_0, lst_0, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})})
_ = __local_var_2_3
var __t4 gopurs_runtime.Value
{
if (__local_var_2_3.Type == 9 && __local_var_2_3.IntVal == 930809136) {
var go__3_5 gopurs_runtime.Value
go__3_5 = gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_5:
for {
if false { continue go__3_5 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t6 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 786377863) {
__t6 = v_4
goto end_branch_6
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})}
v1_5_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_5.UnsafePtr).V1
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
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.Apply2(go__3_5, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, gopurs_runtime.RecordGet((*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_3.UnsafePtr).V0, "revInit")), gopurs_runtime.RecordGet((*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_3.UnsafePtr).V0, "last"))})}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_4:
return __t4
}()
})
	})
	return cache_unsnoc
}

var cache_zipWith gopurs_runtime.Value
var once_zipWith sync.Once
func Get_zipWith() gopurs_runtime.Value {
	once_zipWith.Do(func() {
		cache_zipWith = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWith(f_0_box, xs_1_box, ys_2_box)
})
	})
	return cache_zipWith
}

var cache_zip gopurs_runtime.Value
var once_zip sync.Once
func Get_zip() gopurs_runtime.Value {
	once_zip.Do(func() {
		cache_zip = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, ys_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zip(xs_0_box, ys_1_box)
})
	})
	return cache_zip
}

var cache_zipWithA gopurs_runtime.Value
var once_zipWithA sync.Once
func Get_zipWithA() gopurs_runtime.Value {
	once_zipWithA.Do(func() {
		cache_zipWithA = gopurs_runtime.Func(func(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
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
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var v2_8 gopurs_runtime.Value = v2_8_loop
_ = v2_8
var __t2 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 786377863) {
__t2 = v2_8
goto end_branch_2
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 786377863) {
__t2 = v2_8
goto end_branch_2
} else {

}
}
{
if ((v_6.Type == 9 && v_6.IntVal == 1358893437)) && ((v1_7.Type == 9 && v1_7.IntVal == 1358893437)) {
v_6_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_6.UnsafePtr).V1
v1_7_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_7.UnsafePtr).V1
v2_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Apply2(f_2, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_6.UnsafePtr).V0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_7.UnsafePtr).V0), v2_8})}
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
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var v1_8 gopurs_runtime.Value = v1_8_loop
_ = v1_8
var __t4 gopurs_runtime.Value
{
if (v1_8.Type == 9 && v1_8.IntVal == 786377863) {
__t4 = v_7
goto end_branch_4
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 1358893437) {
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_8.UnsafePtr).V0, v_7})}
v1_8_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_8.UnsafePtr).V1
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
return gopurs_runtime.Apply(sequence1_1_0, gopurs_runtime.Apply2(go__6_3, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, gopurs_runtime.Apply3(go__5_1, xs_3, ys_4, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})})))
})
}()
})
	})
	return cache_zipWithA
}

var cache_range_ gopurs_runtime.Value
var once_range_ sync.Once
func Get_range_() gopurs_runtime.Value {
	once_range_.Do(func() {
		cache_range_ = gopurs_runtime.Func2(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_range_(start_0_box, end_1_box)
})
	})
	return cache_range_
}

var cache_partition gopurs_runtime.Value
var once_partition sync.Once
func Get_partition() gopurs_runtime.Value {
	once_partition.Do(func() {
		cache_partition = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_partition(p_0_box, xs_1_box)
})
	})
	return cache_partition
}

var cache_null gopurs_runtime.Value
var once_null sync.Once
func Get_null() gopurs_runtime.Value {
	once_null.Do(func() {
		cache_null = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Bool((v_0.Type == 9 && v_0.IntVal == 786377863))
}()
})
	})
	return cache_null
}

var cache_nubBy gopurs_runtime.Value
var once_nubBy sync.Once
func Get_nubBy() gopurs_runtime.Value {
	once_nubBy.Do(func() {
		cache_nubBy = gopurs_runtime.Func(func(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_0:
for {
if false { continue go__1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var v2_4 gopurs_runtime.Value = v2_4_loop
_ = v2_4
var __t1 gopurs_runtime.Value
{
if (v2_4.Type == 9 && v2_4.IntVal == 786377863) {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 1358893437) {
v3_5_2 := gopurs_runtime.Apply3(pkg_Data_List_Internal.Get_insertAndLookupBy(), p_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_4.UnsafePtr).V0, v_2)
_ = v3_5_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v3_5_2, "found").IntVal) != (0) {
v_2_loop = gopurs_runtime.RecordGet(v3_5_2, "result")
v1_3_loop = v1_3
v2_4_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_4.UnsafePtr).V1
continue go__1_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
v_2_loop = gopurs_runtime.RecordGet(v3_5_2, "result")
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_4.UnsafePtr).V0, v1_3})}
v2_4_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_4.UnsafePtr).V1
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
__local_var_2_4 := gopurs_runtime.Apply2(go__1_0, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Internal.Data_Data_List_Internal_Leaf{})}, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})})
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_5 gopurs_runtime.Value
go__4_5 = gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__4_5:
for {
if false { continue go__4_5 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t6 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 786377863) {
__t6 = v_5
goto end_branch_6
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})}
v1_6_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V1
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
return gopurs_runtime.Apply2(go__4_5, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, gopurs_runtime.Apply(__local_var_2_4, x_3))
})
}()
})
	})
	return cache_nubBy
}

var cache_nub gopurs_runtime.Value
var once_nub sync.Once
func Get_nub() gopurs_runtime.Value {
	once_nub.Do(func() {
		cache_nub = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_nubBy(), gopurs_runtime.RecordGet(dictOrd_0, "compare"))
}()
})
	})
	return cache_nub
}

var cache_newtypePattern gopurs_runtime.Value
var once_newtypePattern sync.Once
func Get_newtypePattern() gopurs_runtime.Value {
	once_newtypePattern.Do(func() {
		cache_newtypePattern = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypePattern
}

var cache_mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		cache_mapMaybe = gopurs_runtime.Func(func(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
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
if (v1_3.Type == 9 && v1_3.IntVal == 786377863) {
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
if (v1_6.Type == 9 && v1_6.IntVal == 786377863) {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})}
v1_6_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V1
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
__t1 = gopurs_runtime.Apply2(go__4_2, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, v_2)
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437) {
v2_4_4 := gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0)
_ = v2_4_4
var __t5 gopurs_runtime.Value
{
if (v2_4_4.Type == 9 && v2_4_4.IntVal == 3589588149) {
v_2_loop = v_2
v1_3_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V1
continue go__1_0
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
if (v2_4_4.Type == 9 && v2_4_4.IntVal == 930809136) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_Maybe.Data_Data_Maybe_Just)(v2_4_4.UnsafePtr).V0, v_2})}
v1_3_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V1
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
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})})
}()
})
	})
	return cache_mapMaybe
}

var cache_manyRec gopurs_runtime.Value
var once_manyRec sync.Once
func Get_manyRec() gopurs_runtime.Value {
	once_manyRec.Do(func() {
		cache_manyRec = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictAlternative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_manyRec(dictMonadRec_0_box, dictAlternative_1_box)
})
	})
	return cache_manyRec
}

var cache_someRec gopurs_runtime.Value
var once_someRec sync.Once
func Get_someRec() gopurs_runtime.Value {
	once_someRec.Do(func() {
		cache_someRec = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictAlternative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_someRec(dictMonadRec_0_box, dictAlternative_1_box)
})
	})
	return cache_someRec
}

var cache_some gopurs_runtime.Value
var once_some sync.Once
func Get_some() gopurs_runtime.Value {
	once_some.Do(func() {
		cache_some = gopurs_runtime.Func3(func(dictAlternative_0_box gopurs_runtime.Value, dictLazy_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_some(dictAlternative_0_box, dictLazy_1_box, v_2_box)
})
	})
	return cache_some
}

var cache_many gopurs_runtime.Value
var once_many sync.Once
func Get_many() gopurs_runtime.Value {
	once_many.Do(func() {
		cache_many = gopurs_runtime.Func3(func(dictAlternative_0_box gopurs_runtime.Value, dictLazy_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_many(dictAlternative_0_box, dictLazy_1_box, v_2_box)
})
	})
	return cache_many
}

var cache_length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		cache_length = func() gopurs_runtime.Value {
var go__0_0 gopurs_runtime.Value
go__0_0 = gopurs_runtime.Func(func(b_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__0_0:
for {
if false { continue go__0_0 }
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 786377863) {
__t1 = b_1
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437) {
b_1_loop = gopurs_runtime.Int((b_1.IntVal) + (1))
v_2_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_2.UnsafePtr).V1
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
	return cache_length
}

var cache_last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		cache_last = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
last:
for {
if false { continue last }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437) {
var __t1 gopurs_runtime.Value
{
if ((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V1.Type == 9 && (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V1.IntVal == 786377863) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0})}
goto end_branch_1
} else {

}
}
{
v_0_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V1
continue last
__t1 = gopurs_runtime.Value{}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_0:
return __t0
}
}()
})
	})
	return cache_last
}

var cache_insertBy gopurs_runtime.Value
var once_insertBy sync.Once
func Get_insertBy() gopurs_runtime.Value {
	once_insertBy.Do(func() {
		cache_insertBy = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insertBy(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_insertBy
}

var cache_insertAt gopurs_runtime.Value
var once_insertAt sync.Once
func Get_insertAt() gopurs_runtime.Value {
	once_insertAt.Do(func() {
		cache_insertAt = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insertAt(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_insertAt
}

var cache_insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		cache_insert = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_insertBy(), gopurs_runtime.RecordGet(dictOrd_0, "compare"))
}()
})
	})
	return cache_insert
}

var cache_init_ gopurs_runtime.Value
var once_init_ sync.Once
func Get_init_() gopurs_runtime.Value {
	once_init_.Do(func() {
		cache_init_ = gopurs_runtime.Func(func(lst_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var lst_0 gopurs_runtime.Value = lst_0_loop
_ = lst_0
__local_var_1_0 := gopurs_runtime.Apply(Get_unsnoc(), lst_0)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.RecordGet((*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_1_0.UnsafePtr).V0, "init")})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_1:
return __t1
}()
})
	})
	return cache_init_
}

var cache_index gopurs_runtime.Value
var once_index sync.Once
func Get_index() gopurs_runtime.Value {
	once_index.Do(func() {
		cache_index = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_index(v_0_box, v1_1_box)
})
	})
	return cache_index
}

var cache_head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		cache_head = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 786377863) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}()
})
	})
	return cache_head
}

var cache_transpose gopurs_runtime.Value
var once_transpose sync.Once
func Get_transpose() gopurs_runtime.Value {
	once_transpose.Do(func() {
		cache_transpose = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
transpose:
for {
if false { continue transpose }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 786377863) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437) {
var __t1 gopurs_runtime.Value
{
if ((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.Type == 9 && (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.IntVal == 786377863) {
v_0_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V1
continue transpose
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if ((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.Type == 9 && (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.IntVal == 1358893437) {
var go__1_2 gopurs_runtime.Value
go__1_2 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__1_2:
for {
if false { continue go__1_2 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t3 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 786377863) {
var go__4_4 gopurs_runtime.Value
go__4_4 = gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__4_4:
for {
if false { continue go__4_4 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t5 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 786377863) {
__t5 = v_5
goto end_branch_5
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})}
v1_6_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V1
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
__t3 = gopurs_runtime.Apply2(go__4_4, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, v_2)
goto end_branch_3
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437) {
var __t6 gopurs_runtime.Value
{
if ((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0.Type == 9 && (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0.IntVal == 786377863) {
v_2_loop = v_2
v1_3_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V1
continue go__1_2
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
if ((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0.Type == 9 && (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0.IntVal == 1358893437) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0.UnsafePtr).V0, v_2})}
v1_3_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V1
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
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t8 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 786377863) {
var go__4_9 gopurs_runtime.Value
go__4_9 = gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__4_9:
for {
if false { continue go__4_9 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t10 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 786377863) {
__t10 = v_5
goto end_branch_10
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})}
v1_6_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V1
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
__t8 = gopurs_runtime.Apply2(go__4_9, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, v_2)
goto end_branch_8
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437) {
var __t11 gopurs_runtime.Value
{
if ((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0.Type == 9 && (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0.IntVal == 786377863) {
v_2_loop = v_2
v1_3_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V1
continue go__1_7
__t11 = gopurs_runtime.Value{}
goto end_branch_11
} else {

}
}
{
if ((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0.Type == 9 && (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0.IntVal == 1358893437) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0.UnsafePtr).V1, v_2})}
v1_3_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V1
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.Apply2(go__1_2, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V1)})}, gopurs_runtime.Apply(Get_transpose(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V1, gopurs_runtime.Apply2(go__1_7, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V1)})})})}
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
	return cache_transpose
}

var cache_groupBy gopurs_runtime.Value
var once_groupBy sync.Once
func Get_groupBy() gopurs_runtime.Value {
	once_groupBy.Do(func() {
		cache_groupBy = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_groupBy(v_0_box, v1_1_box)
})
	})
	return cache_groupBy
}

var cache_groupAllBy gopurs_runtime.Value
var once_groupAllBy sync.Once
func Get_groupAllBy() gopurs_runtime.Value {
	once_groupAllBy.Do(func() {
		cache_groupAllBy = gopurs_runtime.Func(func(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
__local_var_1_0 := gopurs_runtime.Apply(Get_groupBy(), gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(p_0, x_1, y_2).Type == 9 && gopurs_runtime.Apply2(p_0, x_1, y_2).IntVal == 902936544))
}))
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(Get_sortBy(), p_0)
_ = __local_var_2_1
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply(__local_var_2_1, x_3))
})
}()
})
	})
	return cache_groupAllBy
}

var cache_group gopurs_runtime.Value
var once_group sync.Once
func Get_group() gopurs_runtime.Value {
	once_group.Do(func() {
		cache_group = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_groupBy(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
}()
})
	})
	return cache_group
}

var cache_groupAll gopurs_runtime.Value
var once_groupAll sync.Once
func Get_groupAll() gopurs_runtime.Value {
	once_groupAll.Do(func() {
		cache_groupAll = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
__local_var_1_0 := gopurs_runtime.Apply(Get_group(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
_ = __local_var_1_0
compare_2_1 := gopurs_runtime.RecordGet(dictOrd_0, "compare")
_ = compare_2_1
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply2(Get_sortBy(), compare_2_1, x_3))
})
}()
})
	})
	return cache_groupAll
}

var cache_fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		cache_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})})
}()
})
	})
	return cache_fromFoldable
}

var cache_foldM gopurs_runtime.Value
var once_foldM sync.Once
func Get_foldM() gopurs_runtime.Value {
	once_foldM.Do(func() {
		cache_foldM = gopurs_runtime.Func4(func(dictMonad_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value, v2_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldM(dictMonad_0_box, v_1_box, v1_2_box, v2_3_box)
})
	})
	return cache_foldM
}

var cache_findIndex gopurs_runtime.Value
var once_findIndex sync.Once
func Get_findIndex() gopurs_runtime.Value {
	once_findIndex.Do(func() {
		cache_findIndex = gopurs_runtime.Func(func(fn_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
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
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(fn_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0).IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{v_2})}
goto end_branch_2
} else {

}
}
{
v_2_loop = gopurs_runtime.Int((v_2.IntVal) + (1))
v1_3_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V1
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
if (v1_3.Type == 9 && v1_3.IntVal == 786377863) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
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
}()
})
	})
	return cache_findIndex
}

var cache_findLastIndex gopurs_runtime.Value
var once_findLastIndex sync.Once
func Get_findLastIndex() gopurs_runtime.Value {
	once_findLastIndex.Do(func() {
		cache_findLastIndex = gopurs_runtime.Func2(func(fn_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_findLastIndex(fn_0_box, xs_1_box)
})
	})
	return cache_findLastIndex
}

var cache_filterM gopurs_runtime.Value
var once_filterM sync.Once
func Get_filterM() gopurs_runtime.Value {
	once_filterM.Do(func() {
		cache_filterM = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
filterM:
for {
if false { continue filterM }
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 786377863) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})})
goto end_branch_2
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437) {
__local_var_5_3 := (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_4.UnsafePtr).V0
_ = __local_var_5_3
__local_var_6_4 := (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_4.UnsafePtr).V1
_ = __local_var_6_4
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), gopurs_runtime.Apply(v_3, __local_var_5_3), gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), gopurs_runtime.Apply3(Get_filterM(), dictMonad_0, v_3, __local_var_6_4), gopurs_runtime.Func(func(xs_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (b_7.IntVal) != (0) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{__local_var_5_3, xs_prime_8})}
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
	return cache_filterM
}

var cache_filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		cache_filter = gopurs_runtime.Func(func(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
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
if (v1_3.Type == 9 && v1_3.IntVal == 786377863) {
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
if (v1_6.Type == 9 && v1_6.IntVal == 786377863) {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})}
v1_6_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V1
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
__t1 = gopurs_runtime.Apply2(go__4_2, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, v_2)
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437) {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(p_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0).IntVal) != (0) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0, v_2})}
v1_3_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V1
continue go__1_0
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
v_2_loop = v_2
v1_3_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V1
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
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})})
}()
})
	})
	return cache_filter
}

var cache_intersectBy gopurs_runtime.Value
var once_intersectBy sync.Once
func Get_intersectBy() gopurs_runtime.Value {
	once_intersectBy.Do(func() {
		cache_intersectBy = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intersectBy(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_intersectBy
}

var cache_intersect gopurs_runtime.Value
var once_intersect sync.Once
func Get_intersect() gopurs_runtime.Value {
	once_intersect.Do(func() {
		cache_intersect = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_intersectBy(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
}()
})
	})
	return cache_intersect
}

var cache_nubByEq gopurs_runtime.Value
var once_nubByEq sync.Once
func Get_nubByEq() gopurs_runtime.Value {
	once_nubByEq.Do(func() {
		cache_nubByEq = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_nubByEq(v_0_box, v1_1_box)
})
	})
	return cache_nubByEq
}

var cache_nubEq gopurs_runtime.Value
var once_nubEq sync.Once
func Get_nubEq() gopurs_runtime.Value {
	once_nubEq.Do(func() {
		cache_nubEq = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_nubByEq(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
}()
})
	})
	return cache_nubEq
}

var cache_eqPattern gopurs_runtime.Value
var once_eqPattern sync.Once
func Get_eqPattern() gopurs_runtime.Value {
	once_eqPattern.Do(func() {
		cache_eqPattern = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
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
if (v_4.Type == 9 && v_4.IntVal == 786377863) {
__t1 = gopurs_runtime.Bool(((v1_5.Type == 9 && v1_5.IntVal == 786377863)) && ((v2_6.IntVal) != (0)))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(((v_4.Type == 9 && v_4.IntVal == 1358893437)) && (((v1_5.Type == 9 && v1_5.IntVal == 1358893437)) && ((gopurs_runtime.Apply3(go__3_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V1, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_5.UnsafePtr).V1, gopurs_runtime.Bool(((v2_6.IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0).IntVal) != (0)))).IntVal) != (0))))
}
end_branch_1:
return __t1
})
return gopurs_runtime.Apply3(go__3_0, x_1, y_2, gopurs_runtime.Bool(true))
}))
}()
})
	})
	return cache_eqPattern
}

var cache_ordPattern gopurs_runtime.Value
var once_ordPattern sync.Once
func Get_ordPattern() gopurs_runtime.Value {
	once_ordPattern.Do(func() {
		cache_ordPattern = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
eqPattern1_2_1 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_2 gopurs_runtime.Value
_ = go__4_2
go__4_2 = gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if ((v2_7.IntVal) != (0)) != (true) {
__t3 = gopurs_runtime.Bool(false)
goto end_branch_3
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 786377863) {
__t3 = gopurs_runtime.Bool(((v1_6.Type == 9 && v1_6.IntVal == 786377863)) && ((v2_7.IntVal) != (0)))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Bool(((v_5.Type == 9 && v_5.IntVal == 1358893437)) && (((v1_6.Type == 9 && v1_6.IntVal == 1358893437)) && ((gopurs_runtime.Apply3(go__4_2, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_5.UnsafePtr).V1, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V1, gopurs_runtime.Bool(((v2_7.IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_5.UnsafePtr).V0).IntVal) != (0)))).IntVal) != (0))))
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
}()
})
	})
	return cache_ordPattern
}

var cache_elemLastIndex gopurs_runtime.Value
var once_elemLastIndex sync.Once
func Get_elemLastIndex() gopurs_runtime.Value {
	once_elemLastIndex.Do(func() {
		cache_elemLastIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_elemLastIndex(dictEq_0_box, x_1_box)
})
	})
	return cache_elemLastIndex
}

var cache_elemIndex gopurs_runtime.Value
var once_elemIndex sync.Once
func Get_elemIndex() gopurs_runtime.Value {
	once_elemIndex.Do(func() {
		cache_elemIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_elemIndex(dictEq_0_box, x_1_box)
})
	})
	return cache_elemIndex
}

var cache_dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		cache_dropWhile = gopurs_runtime.Func(func(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
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
if ((v_2.Type == 9 && v_2.IntVal == 1358893437)) && ((gopurs_runtime.Apply(p_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_2.UnsafePtr).V0).IntVal) != (0)) {
v_2_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_2.UnsafePtr).V1
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
}()
})
	})
	return cache_dropWhile
}

var cache_dropEnd gopurs_runtime.Value
var once_dropEnd sync.Once
func Get_dropEnd() gopurs_runtime.Value {
	once_dropEnd.Do(func() {
		cache_dropEnd = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_dropEnd(n_0_box, xs_1_box)
})
	})
	return cache_dropEnd
}

var cache_drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		cache_drop = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_drop(v_0_box, v1_1_box)
})
	})
	return cache_drop
}

var cache_slice gopurs_runtime.Value
var once_slice sync.Once
func Get_slice() gopurs_runtime.Value {
	once_slice.Do(func() {
		cache_slice = gopurs_runtime.Func3(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_slice(start_0_box, end_1_box, xs_2_box)
})
	})
	return cache_slice
}

var cache_takeEnd gopurs_runtime.Value
var once_takeEnd sync.Once
func Get_takeEnd() gopurs_runtime.Value {
	once_takeEnd.Do(func() {
		cache_takeEnd = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_takeEnd(n_0_box, xs_1_box)
})
	})
	return cache_takeEnd
}

var cache_deleteBy gopurs_runtime.Value
var once_deleteBy sync.Once
func Get_deleteBy() gopurs_runtime.Value {
	once_deleteBy.Do(func() {
		cache_deleteBy = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_deleteBy(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_deleteBy
}

var cache_unionBy gopurs_runtime.Value
var once_unionBy sync.Once
func Get_unionBy() gopurs_runtime.Value {
	once_unionBy.Do(func() {
		cache_unionBy = gopurs_runtime.Func3(func(eq2_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unionBy(eq2_0_box, xs_1_box, ys_2_box)
})
	})
	return cache_unionBy
}

var cache_union gopurs_runtime.Value
var once_union sync.Once
func Get_union() gopurs_runtime.Value {
	once_union.Do(func() {
		cache_union = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_unionBy(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
}()
})
	})
	return cache_union
}

var cache_deleteAt gopurs_runtime.Value
var once_deleteAt sync.Once
func Get_deleteAt() gopurs_runtime.Value {
	once_deleteAt.Do(func() {
		cache_deleteAt = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_deleteAt(v_0_box, v1_1_box)
})
	})
	return cache_deleteAt
}

var cache_delete_ gopurs_runtime.Value
var once_delete_ sync.Once
func Get_delete_() gopurs_runtime.Value {
	once_delete_.Do(func() {
		cache_delete_ = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_deleteBy(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
}()
})
	})
	return cache_delete_
}

var cache_difference gopurs_runtime.Value
var once_difference sync.Once
func Get_difference() gopurs_runtime.Value {
	once_difference.Do(func() {
		cache_difference = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
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
if (v_3.Type == 9 && v_3.IntVal == 786377863) {
__t1 = b_2
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437) {
b_2_loop = Call_deleteBy(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_3.UnsafePtr).V0, b_2)
v_3_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_3.UnsafePtr).V1
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
}()
})
	})
	return cache_difference
}

var cache_concatMap gopurs_runtime.Value
var once_concatMap sync.Once
func Get_concatMap() gopurs_runtime.Value {
	once_concatMap.Do(func() {
		cache_concatMap = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_concatMap(b_0_box, a_1_box)
})
	})
	return cache_concatMap
}

var cache_concat gopurs_runtime.Value
var once_concat sync.Once
func Get_concat() gopurs_runtime.Value {
	once_concat.Do(func() {
		cache_concat = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_bindList(), "bind"), v_0, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
}()
})
	})
	return cache_concat
}

var cache_catMaybes gopurs_runtime.Value
var once_catMaybes sync.Once
func Get_catMaybes() gopurs_runtime.Value {
	once_catMaybes.Do(func() {
		cache_catMaybes = func() gopurs_runtime.Value {
var go__0_0 gopurs_runtime.Value
go__0_0 = gopurs_runtime.Func(func(v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__0_0:
for {
if false { continue go__0_0 }
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t1 gopurs_runtime.Value
{
if (v1_2.Type == 9 && v1_2.IntVal == 786377863) {
var go__3_2 gopurs_runtime.Value
go__3_2 = gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_2:
for {
if false { continue go__3_2 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 786377863) {
__t3 = v_4
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})}
v1_5_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_5.UnsafePtr).V1
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
__t1 = gopurs_runtime.Apply2(go__3_2, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, v_1)
goto end_branch_1
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437) {
var __t4 gopurs_runtime.Value
{
if ((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_2.UnsafePtr).V0.Type == 9 && (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_2.UnsafePtr).V0.IntVal == 3589588149) {
v_1_loop = v_1
v1_2_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_2.UnsafePtr).V1
continue go__0_0
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
if ((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_2.UnsafePtr).V0.Type == 9 && (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_2.UnsafePtr).V0.IntVal == 930809136) {
v_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_Maybe.Data_Data_Maybe_Just)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_2.UnsafePtr).V0.UnsafePtr).V0, v_1})}
v1_2_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_2.UnsafePtr).V1
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
return gopurs_runtime.Apply(go__0_0, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})})
}()
	})
	return cache_catMaybes
}

var cache_alterAt gopurs_runtime.Value
var once_alterAt sync.Once
func Get_alterAt() gopurs_runtime.Value {
	once_alterAt.Do(func() {
		cache_alterAt = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alterAt(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_alterAt
}

var cache_modifyAt gopurs_runtime.Value
var once_modifyAt sync.Once
func Get_modifyAt() gopurs_runtime.Value {
	once_modifyAt.Do(func() {
		cache_modifyAt = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modifyAt(n_0_box, f_1_box)
})
	})
	return cache_modifyAt
}

func Call_updateAt(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
updateAt:
for {
if false { continue updateAt }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 1358893437) {
var __t3 gopurs_runtime.Value
{
if (v_0.IntVal) == (0) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{v1_1, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V1})}})}
goto end_branch_3
} else {

}
}
{
__local_var_3_1 := Call_updateAt(gopurs_runtime.Int((v_0.IntVal) - (1)), v1_1, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V1)
_ = __local_var_3_1
var __t2 gopurs_runtime.Value
{
if (__local_var_3_1.Type == 9 && __local_var_3_1.IntVal == 930809136) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V0, (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_3_1.UnsafePtr).V0})}})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_0:
return __t0
}
}

func Call_stripPrefix(dictEq_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
__local_var_3_0 := gopurs_runtime.Func2(func(prefix_3 gopurs_runtime.Value, input_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (input_4.Type == 9 && input_4.IntVal == 1358893437) {
var __t2 gopurs_runtime.Value
{
if (prefix_3.Type == 9 && prefix_3.IntVal == 1358893437) {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(prefix_3.UnsafePtr).V0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(input_4.UnsafePtr).V0).IntVal) != (0) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Data_Control_Monad_Rec_Class_Loop{gopurs_runtime.RecordDict2("a", "b", (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(prefix_3.UnsafePtr).V1, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(input_4.UnsafePtr).V1)})}})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
if (prefix_3.Type == 9 && prefix_3.IntVal == 786377863) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Data_Control_Monad_Rec_Class_Done{input_4})}})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if (prefix_3.Type == 9 && prefix_3.IntVal == 786377863) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Data_Control_Monad_Rec_Class_Done{input_4})}})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_1:
return __t1
})
_ = __local_var_3_0
__local_var_4_4 := gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 3589588149) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Data_Control_Monad_Rec_Class_Done{gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}})}
goto end_branch_5
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 930809136) {
var __t6 gopurs_runtime.Value
{
if ((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_4.UnsafePtr).V0.Type == 9 && (*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_4.UnsafePtr).V0.IntVal == 525585346) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Data_Control_Monad_Rec_Class_Loop{gopurs_runtime.UncurriedApp2(__local_var_3_0, gopurs_runtime.RecordGet((*pkg_Control_Monad_Rec_Class.Data_Control_Monad_Rec_Class_Loop)((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_4.UnsafePtr).V0.UnsafePtr).V0, "a"), gopurs_runtime.RecordGet((*pkg_Control_Monad_Rec_Class.Data_Control_Monad_Rec_Class_Loop)((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_4.UnsafePtr).V0.UnsafePtr).V0, "b"))})}
goto end_branch_6
} else {

}
}
{
if ((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_4.UnsafePtr).V0.Type == 9 && (*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_4.UnsafePtr).V0.IntVal == 60402430) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Data_Control_Monad_Rec_Class_Done{gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*pkg_Control_Monad_Rec_Class.Data_Control_Monad_Rec_Class_Done)((*pkg_Data_Maybe.Data_Data_Maybe_Just)(v_4.UnsafePtr).V0.UnsafePtr).V0})}})}
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
})
_ = __local_var_4_4
var go__5_7 gopurs_runtime.Value
go__5_7 = gopurs_runtime.Func(func(v_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__5_7:
for {
if false { continue go__5_7 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t8 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 525585346) {
v_6_loop = gopurs_runtime.UncurriedApp(__local_var_4_4, (*pkg_Control_Monad_Rec_Class.Data_Control_Monad_Rec_Class_Loop)(v_6.UnsafePtr).V0)
continue go__5_7
__t8 = gopurs_runtime.Value{}
goto end_branch_8
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 60402430) {
__t8 = (*pkg_Control_Monad_Rec_Class.Data_Control_Monad_Rec_Class_Done)(v_6.UnsafePtr).V0
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
}

func Call_span(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
span:
for {
if false { continue span }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if ((v1_1.Type == 9 && v1_1.IntVal == 1358893437)) && ((gopurs_runtime.Apply(v_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_1.UnsafePtr).V0).IntVal) != (0)) {
v2_2_1 := Call_span(v_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_1.UnsafePtr).V1)
_ = v2_2_1
__t0 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_1.UnsafePtr).V0, gopurs_runtime.RecordGet(v2_2_1, "init")})}, gopurs_runtime.RecordGet(v2_2_1, "rest"))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, v1_1)
}
end_branch_0:
return __t0
}
}

func Call_snoc(xs_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{x_1, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}})}, xs_0)
}

func Call_zipWith(f_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
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
if (v_4.Type == 9 && v_4.IntVal == 786377863) {
__t1 = v2_6
goto end_branch_1
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 786377863) {
__t1 = v2_6
goto end_branch_1
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 1358893437)) && ((v1_5.Type == 9 && v1_5.IntVal == 1358893437)) {
v_4_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V1
v1_5_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_5.UnsafePtr).V1
v2_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Apply2(f_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_5.UnsafePtr).V0), v2_6})}
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
if (v1_6.Type == 9 && v1_6.IntVal == 786377863) {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})}
v1_6_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V1
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
return gopurs_runtime.Apply2(go__4_2, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, gopurs_runtime.Apply3(go__3_0, xs_1, ys_2, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}))
}

func Call_zip(xs_0_loop gopurs_runtime.Value, ys_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var ys_1 gopurs_runtime.Value = ys_1_loop
_ = ys_1
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_0:
for {
if false { continue go__2_0 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var v1_4 gopurs_runtime.Value = v1_4_loop
_ = v1_4
var v2_5 gopurs_runtime.Value = v2_5_loop
_ = v2_5
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 786377863) {
__t1 = v2_5
goto end_branch_1
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 786377863) {
__t1 = v2_5
goto end_branch_1
} else {

}
}
{
if ((v_3.Type == 9 && v_3.IntVal == 1358893437)) && ((v1_4.Type == 9 && v1_4.IntVal == 1358893437)) {
v_3_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_3.UnsafePtr).V1
v1_4_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_4.UnsafePtr).V1
v2_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_3.UnsafePtr).V0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_4.UnsafePtr).V0})}, v2_5})}
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
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 786377863) {
__t3 = v_4
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})}
v1_5_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_5.UnsafePtr).V1
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
return gopurs_runtime.Apply2(go__3_2, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, gopurs_runtime.Apply3(go__2_0, xs_0, ys_1, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}))
}

func Call_range_(start_0_loop gopurs_runtime.Value, end_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var start_0 gopurs_runtime.Value = start_0_loop
_ = start_0
var end_1 gopurs_runtime.Value = end_1_loop
_ = end_1
var __t3 gopurs_runtime.Value
{
if (start_0.IntVal) == (end_1.IntVal) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{start_0, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}})}
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
var s_3 gopurs_runtime.Value = s_3_loop
_ = s_3
var e_4 gopurs_runtime.Value = e_4_loop
_ = e_4
var step_5 gopurs_runtime.Value = step_5_loop
_ = step_5
var rest_6 gopurs_runtime.Value = rest_6_loop
_ = rest_6
var __t1 gopurs_runtime.Value
{
if (s_3.IntVal) == (e_4.IntVal) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{s_3, rest_6})}
goto end_branch_1
} else {

}
}
{
s_3_loop = gopurs_runtime.Int((s_3.IntVal) + (step_5.IntVal))
e_4_loop = e_4
step_5_loop = step_5
rest_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{s_3, rest_6})}
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
if (start_0.IntVal) > (end_1.IntVal) {
__t2 = gopurs_runtime.Int(1)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Int(-1)
}
end_branch_2:
__t3 = gopurs_runtime.Apply4(go__2_0, end_1, start_0, __t2, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})})
}
end_branch_3:
return __t3
}

func Call_partition(p_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(p_0, x_2).IntVal) != (0) {
__t0 = gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.RecordGet(v_3, "no"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{x_2, gopurs_runtime.RecordGet(v_3, "yes")})})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{x_2, gopurs_runtime.RecordGet(v_3, "no")})}, gopurs_runtime.RecordGet(v_3, "yes"))
}
end_branch_0:
return __t0
}), gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}), xs_1)
}

func Call_manyRec(dictMonadRec_0_loop gopurs_runtime.Value, dictAlternative_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictAlternative_1 gopurs_runtime.Value = dictAlternative_1_loop
_ = dictAlternative_1
Alt0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_1, "Plus1"), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{})
_ = Alt0_2_0
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_1, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Func(func(p_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Alt0_2_0, "alt"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Alt0_2_0, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Control_Monad_Rec_Class.Get_Loop(), p_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Data_Control_Monad_Rec_Class_Done{pkg_Data_Unit.Get_unit()})})), gopurs_runtime.Func(func(aa_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (aa_6.Type == 9 && aa_6.IntVal == 525585346) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Data_Control_Monad_Rec_Class_Loop{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Control_Monad_Rec_Class.Data_Control_Monad_Rec_Class_Loop)(aa_6.UnsafePtr).V0, acc_5})}})}
goto end_branch_2
} else {

}
}
{
if (aa_6.Type == 9 && aa_6.IntVal == 60402430) {
var go__7_3 gopurs_runtime.Value
go__7_3 = gopurs_runtime.Func(func(v_8_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_9_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__7_3:
for {
if false { continue go__7_3 }
var v_8 gopurs_runtime.Value = v_8_loop
_ = v_8
var v1_9 gopurs_runtime.Value = v1_9_loop
_ = v1_9
var __t4 gopurs_runtime.Value
{
if (v1_9.Type == 9 && v1_9.IntVal == 786377863) {
__t4 = v_8
goto end_branch_4
} else {

}
}
{
if (v1_9.Type == 9 && v1_9.IntVal == 1358893437) {
v_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_9.UnsafePtr).V0, v_8})}
v1_9_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_9.UnsafePtr).V1
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
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Data_Control_Monad_Rec_Class_Done{gopurs_runtime.Apply2(go__7_3, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, acc_5)})}
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
}), gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})})
})
}

func Call_someRec(dictMonadRec_0_loop gopurs_runtime.Value, dictAlternative_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictAlternative_1 gopurs_runtime.Value = dictAlternative_1_loop
_ = dictAlternative_1
manyRec2_2_0 := Call_manyRec(dictMonadRec_0, dictAlternative_1)
_ = manyRec2_2_0
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_1, "Applicative0"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_1, "Plus1"), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_List_Types.Get_Cons(), v_3), gopurs_runtime.Apply(manyRec2_2_0, v_3))
})
}

func Call_some(dictAlternative_0_loop gopurs_runtime.Value, dictLazy_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
var dictLazy_1 gopurs_runtime.Value = dictLazy_1_loop
_ = dictLazy_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_List_Types.Get_Cons(), v_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_1, "defer"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_many(dictAlternative_0, dictLazy_1, v_2)
})))
}

func Call_many(dictAlternative_0_loop gopurs_runtime.Value, dictLazy_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
var dictLazy_1 gopurs_runtime.Value = dictLazy_1_loop
_ = dictLazy_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "alt"), Call_some(dictAlternative_0, dictLazy_1, v_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}))
}

func Call_insertBy(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
insertBy:
for {
if false { continue insertBy }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 786377863) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{v1_1, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 1358893437) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(v_0, v1_1, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V0).Type == 9 && gopurs_runtime.Apply2(v_0, v1_1, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V0).IntVal == 380165415) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V0, Call_insertBy(v_0, v1_1, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V1)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{v1_1, v2_2})}
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
}

func Call_insertAt(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
insertAt:
for {
if false { continue insertAt }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v_0.IntVal) == (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{v1_1, v2_2})}})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 1358893437) {
__local_var_3_1 := Call_insertAt(gopurs_runtime.Int((v_0.IntVal) - (1)), v1_1, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V1)
_ = __local_var_3_1
var __t2 gopurs_runtime.Value
{
if (__local_var_3_1.Type == 9 && __local_var_3_1.IntVal == 930809136) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V0, (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_3_1.UnsafePtr).V0})}})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_2:
__t0 = __t2
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_0:
return __t0
}
}

func Call_index(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
index:
for {
if false { continue index }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 786377863) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437) {
var __t1 gopurs_runtime.Value
{
if (v1_1.IntVal) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0})}
goto end_branch_1
} else {

}
}
{
v_0_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V1
v1_1_loop = gopurs_runtime.Int((v1_1.IntVal) - (1))
continue index
__t1 = gopurs_runtime.Value{}
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
}

func Call_groupBy(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
groupBy:
for {
if false { continue groupBy }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 786377863) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 1358893437) {
v2_2_1 := Call_span(gopurs_runtime.Apply(v_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_1.UnsafePtr).V0), (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_1.UnsafePtr).V1)
_ = v2_2_1
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Data_Data_NonEmpty_NonEmpty{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_1.UnsafePtr).V0, gopurs_runtime.RecordGet(v2_2_1, "init")})}, Call_groupBy(v_0, gopurs_runtime.RecordGet(v2_2_1, "rest"))})}
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
}

func Call_foldM(dictMonad_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, v1_2_loop gopurs_runtime.Value, v2_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
foldM:
for {
if false { continue foldM }
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var v2_3 gopurs_runtime.Value = v2_3_loop
_ = v2_3
var __t0 gopurs_runtime.Value
{
if (v2_3.Type == 9 && v2_3.IntVal == 786377863) {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), v1_2)
goto end_branch_0
} else {

}
}
{
if (v2_3.Type == 9 && v2_3.IntVal == 1358893437) {
__local_var_4_1 := (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_3.UnsafePtr).V1
_ = __local_var_4_1
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(v_1, v1_2, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_3.UnsafePtr).V0), gopurs_runtime.Func(func(b_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldM(dictMonad_0, v_1, b_prime_5, __local_var_4_1)
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
}

func Call_findLastIndex(fn_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
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
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(fn_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_4.UnsafePtr).V0).IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{v_3})}
goto end_branch_2
} else {

}
}
{
v_3_loop = gopurs_runtime.Int((v_3.IntVal) + (1))
v1_4_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_4.UnsafePtr).V1
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
if (v1_4.Type == 9 && v1_4.IntVal == 786377863) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
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
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t5 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 786377863) {
__t5 = v_4
goto end_branch_5
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})}
v1_5_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_5.UnsafePtr).V1
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
__local_var_3_3 := gopurs_runtime.Apply2(go__2_0, gopurs_runtime.Int(0), gopurs_runtime.Apply2(go__3_4, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, xs_1))
_ = __local_var_3_3
var __t6 gopurs_runtime.Value
{
if (__local_var_3_3.Type == 9 && __local_var_3_3.IntVal == 930809136) {
var go__4_7 gopurs_runtime.Value
go__4_7 = gopurs_runtime.Func(func(b_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__4_7:
for {
if false { continue go__4_7 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t8 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 786377863) {
__t8 = b_5
goto end_branch_8
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437) {
b_5_loop = gopurs_runtime.Int((b_5.IntVal) + (1))
v_6_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_6.UnsafePtr).V1
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
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{((gopurs_runtime.Apply2(go__4_7, gopurs_runtime.Int(0), xs_1).IntVal) - (1)) - ((*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_3_3.UnsafePtr).V0.IntVal)})}
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_6:
return __t6
}

func Call_intersectBy(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t5 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 786377863) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}
goto end_branch_5
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 786377863) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}
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
if (v1_5.Type == 9 && v1_5.IntVal == 786377863) {
var go__6_2 gopurs_runtime.Value
go__6_2 = gopurs_runtime.Func(func(v_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__6_2:
for {
if false { continue go__6_2 }
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var v1_8 gopurs_runtime.Value = v1_8_loop
_ = v1_8
var __t3 gopurs_runtime.Value
{
if (v1_8.Type == 9 && v1_8.IntVal == 786377863) {
__t3 = v_7
goto end_branch_3
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 1358893437) {
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_8.UnsafePtr).V0, v_7})}
v1_8_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_8.UnsafePtr).V1
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
__t1 = gopurs_runtime.Apply2(go__6_2, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, v_4)
goto end_branch_1
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437) {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_any(), gopurs_runtime.Apply(v_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_5.UnsafePtr).V0), v2_2).IntVal) != (0) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})}
v1_5_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_5.UnsafePtr).V1
continue go__3_0
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
v_4_loop = v_4
v1_5_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_5.UnsafePtr).V1
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
__t5 = gopurs_runtime.Apply2(go__3_0, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, v1_1)
}
end_branch_5:
return __t5
}

func Call_nubByEq(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
nubByEq:
for {
if false { continue nubByEq }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 786377863) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 1358893437) {
__local_var_2_1 := (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_1.UnsafePtr).V0
_ = __local_var_2_1
var go__3_2 gopurs_runtime.Value
go__3_2 = gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_2:
for {
if false { continue go__3_2 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 786377863) {
var go__6_4 gopurs_runtime.Value
go__6_4 = gopurs_runtime.Func(func(v_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__6_4:
for {
if false { continue go__6_4 }
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var v1_8 gopurs_runtime.Value = v1_8_loop
_ = v1_8
var __t5 gopurs_runtime.Value
{
if (v1_8.Type == 9 && v1_8.IntVal == 786377863) {
__t5 = v_7
goto end_branch_5
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 1358893437) {
v_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_8.UnsafePtr).V0, v_7})}
v1_8_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_8.UnsafePtr).V1
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
__t3 = gopurs_runtime.Apply2(go__6_4, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, v_4)
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437) {
var __t6 gopurs_runtime.Value
{
if ((gopurs_runtime.Apply2(v_0, __local_var_2_1, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_5.UnsafePtr).V0).IntVal) != (0)) != (true) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})}
v1_5_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_5.UnsafePtr).V1
continue go__3_2
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
v_4_loop = v_4
v1_5_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_5.UnsafePtr).V1
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{__local_var_2_1, Call_nubByEq(v_0, gopurs_runtime.Apply2(go__3_2, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_1.UnsafePtr).V1))})}
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
}

func Call_elemLastIndex(dictEq_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_findLastIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), v_2, x_1)
}))
}

func Call_elemIndex(dictEq_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
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
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_4.UnsafePtr).V0, x_1).IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{v_3})}
goto end_branch_2
} else {

}
}
{
v_3_loop = gopurs_runtime.Int((v_3.IntVal) + (1))
v1_4_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_4.UnsafePtr).V1
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
if (v1_4.Type == 9 && v1_4.IntVal == 786377863) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
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
}

func Call_dropEnd(n_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 gopurs_runtime.Value = n_0_loop
_ = n_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(b_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_0:
for {
if false { continue go__2_0 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 786377863) {
__t1 = b_3
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437) {
b_3_loop = gopurs_runtime.Int((b_3.IntVal) + (1))
v_4_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V1
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
return gopurs_runtime.Apply2(Get_take(), (gopurs_runtime.Apply2(go__2_0, gopurs_runtime.Int(0), xs_1).IntVal) - (n_0.IntVal), xs_1)
}

func Call_drop(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
drop:
for {
if false { continue drop }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v_0.IntVal) < (1) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 786377863) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 1358893437) {
v_0_loop = gopurs_runtime.Int((v_0.IntVal) - (1))
v1_1_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_1.UnsafePtr).V1
continue drop
__t0 = gopurs_runtime.Value{}
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
}

func Call_slice(start_0_loop gopurs_runtime.Value, end_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var start_0 gopurs_runtime.Value = start_0_loop
_ = start_0
var end_1 gopurs_runtime.Value = end_1_loop
_ = end_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Int((end_1.IntVal) - (start_0.IntVal)), Call_drop(start_0, xs_2))
}

func Call_takeEnd(n_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 gopurs_runtime.Value = n_0_loop
_ = n_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(b_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__2_0:
for {
if false { continue go__2_0 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 786377863) {
__t1 = b_3
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437) {
b_3_loop = gopurs_runtime.Int((b_3.IntVal) + (1))
v_4_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V1
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
return Call_drop((gopurs_runtime.Apply2(go__2_0, gopurs_runtime.Int(0), xs_1).IntVal) - (n_0.IntVal), xs_1)
}

func Call_deleteBy(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
deleteBy:
for {
if false { continue deleteBy }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 786377863) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 1358893437) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(v_0, v1_1, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V0).IntVal) != (0) {
__t1 = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V1
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V0, Call_deleteBy(v_0, v1_1, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V1)})}
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
}

func Call_unionBy(eq2_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var eq2_0 gopurs_runtime.Value = eq2_0_loop
_ = eq2_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
var go__3_0 gopurs_runtime.Value
go__3_0 = gopurs_runtime.Func(func(b_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
go__3_0:
for {
if false { continue go__3_0 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t1 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 786377863) {
__t1 = b_4
goto end_branch_1
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437) {
b_4_loop = Call_deleteBy(eq2_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_5.UnsafePtr).V0, b_4)
v_5_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_5.UnsafePtr).V1
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
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Apply2(go__3_0, Call_nubByEq(eq2_0, ys_2), xs_1), xs_1)
}

func Call_deleteAt(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
deleteAt:
for {
if false { continue deleteAt }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 1358893437) {
var __t3 gopurs_runtime.Value
{
if (v_0.IntVal) == (0) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_1.UnsafePtr).V1})}
goto end_branch_3
} else {

}
}
{
__local_var_2_1 := Call_deleteAt(gopurs_runtime.Int((v_0.IntVal) - (1)), (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_1.UnsafePtr).V1)
_ = __local_var_2_1
var __t2 gopurs_runtime.Value
{
if (__local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 930809136) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_1.UnsafePtr).V0, (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_1.UnsafePtr).V0})}})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_0:
return __t0
}
}

func Call_concatMap(b_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_bindList(), "bind"), a_1, b_0)
}

func Call_alterAt(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
alterAt:
for {
if false { continue alterAt }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 1358893437) {
var __t3 gopurs_runtime.Value
{
if (v_0.IntVal) == (0) {
v3_3_4 := gopurs_runtime.Apply(v1_1, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V0)
_ = v3_3_4
var __t5 gopurs_runtime.Value
{
if (v3_3_4.Type == 9 && v3_3_4.IntVal == 3589588149) {
__t5 = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V1
goto end_branch_5
} else {

}
}
{
if (v3_3_4.Type == 9 && v3_3_4.IntVal == 930809136) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_Maybe.Data_Data_Maybe_Just)(v3_3_4.UnsafePtr).V0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V1})}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{__t5})}
goto end_branch_3
} else {

}
}
{
__local_var_3_1 := Call_alterAt(gopurs_runtime.Int((v_0.IntVal) - (1)), v1_1, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V1)
_ = __local_var_3_1
var __t2 gopurs_runtime.Value
{
if (__local_var_3_1.Type == 9 && __local_var_3_1.IntVal == 930809136) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V0, (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_3_1.UnsafePtr).V0})}})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_0:
return __t0
}
}

func Call_modifyAt(n_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 gopurs_runtime.Value = n_0_loop
_ = n_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(Get_alterAt(), n_0, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Apply(f_1, x_2)})}
}))
}



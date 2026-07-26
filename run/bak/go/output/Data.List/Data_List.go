package Data_List

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_List_Internal "gopurs/output/Data.List.Internal"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	unsafe "unsafe"
)

var cache_tailRecM2 gopurs_runtime.Value
var once_tailRecM2 sync.Once
func Get_tailRecM2() gopurs_runtime.Value {
	once_tailRecM2.Do(func() {
		cache_tailRecM2 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tailRecM2(f_0_box, a_1_box, b_2_box)
})
	})
	return cache_tailRecM2
}

var cache_lessThan gopurs_runtime.Value
var once_lessThan sync.Once
func Get_lessThan() gopurs_runtime.Value {
	once_lessThan.Do(func() {
		cache_lessThan = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1527465420))
})
}()
	})
	return cache_lessThan
}

var cache_greaterThan gopurs_runtime.Value
var once_greaterThan sync.Once
func Get_greaterThan() gopurs_runtime.Value {
	once_greaterThan.Do(func() {
		cache_greaterThan = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 380165415))
})
}()
	})
	return cache_greaterThan
}

var cache_any gopurs_runtime.Value
var once_any sync.Once
func Get_any() gopurs_runtime.Value {
	once_any.Do(func() {
		cache_any = func() gopurs_runtime.Value {
semigroupDisj1_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), v_0, v1_1)
}))
_ = semigroupDisj1_0_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldMap"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_0_0
}), gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "ff")))
}()
	})
	return cache_any
}

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

var cache_Pattern gopurs_runtime.Value
var once_Pattern sync.Once
func Get_Pattern() gopurs_runtime.Value {
	once_Pattern.Do(func() {
		cache_Pattern = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Pattern(x_0_box)
})
	})
	return cache_Pattern
}

var cache_updateAt gopurs_runtime.Value
var once_updateAt sync.Once
func Get_updateAt() gopurs_runtime.Value {
	once_updateAt.Do(func() {
		cache_updateAt = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_updateAt(v_0_box.IntVal, v1_1_box, v2_2_box)
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
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}})})
	})
	return cache_unzip
}

var cache_uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		cache_uncons = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_uncons(v_0_box)
})
	})
	return cache_uncons
}

var cache_toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		cache_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toUnfoldable((*Record_unfoldr_gopurs_runtime_Value)(dictUnfoldable_0_box.UnsafePtr))
})
	})
	return cache_toUnfoldable
}

var cache_tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		cache_tail = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tail(v_0_box)
})
	})
	return cache_tail
}

var cache_stripPrefix gopurs_runtime.Value
var once_stripPrefix sync.Once
func Get_stripPrefix() gopurs_runtime.Value {
	once_stripPrefix.Do(func() {
		cache_stripPrefix = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stripPrefix((*Record_eq_gopurs_runtime_Value)(dictEq_0_box.UnsafePtr), v_1_box, s_2_box)
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
		cache_singleton = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_singleton(a_0_box)
})
	})
	return cache_singleton
}

var cache_sortBy gopurs_runtime.Value
var once_sortBy sync.Once
func Get_sortBy() gopurs_runtime.Value {
	once_sortBy.Do(func() {
		cache_sortBy = gopurs_runtime.Func(func(cmp_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sortBy(cmp_0_box)
})
	})
	return cache_sortBy
}

var cache_sort gopurs_runtime.Value
var once_sort sync.Once
func Get_sort() gopurs_runtime.Value {
	once_sort.Do(func() {
		cache_sort = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sort((*Record_compare_gopurs_runtime_Value)(dictOrd_0_box.UnsafePtr))
})
	})
	return cache_sort
}

var cache_showPattern gopurs_runtime.Value
var once_showPattern sync.Once
func Get_showPattern() gopurs_runtime.Value {
	once_showPattern.Do(func() {
		cache_showPattern = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showPattern((*Record_show_gopurs_runtime_Value)(dictShow_0_box.UnsafePtr))
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
go__0_0 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop gopurs_runtime.Value = v_1_loop_val
var v1_2_loop gopurs_runtime.Value = v1_2_loop_val
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
return gopurs_runtime.Apply(go__0_0, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil})
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
go__0_0 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop gopurs_runtime.Value = v_1_loop_val
var v1_2_loop gopurs_runtime.Value = v1_2_loop_val
var v2_3_loop gopurs_runtime.Value = v2_3_loop_val
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
if (gopurs_runtime.Apply2(Get_lessThan(), v1_2, gopurs_runtime.Int(1)).IntVal) != (0) {
var go__4_2 gopurs_runtime.Value
go__4_2 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
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
__t1 = gopurs_runtime.Apply2(go__4_2, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, v_1)
goto end_branch_1
} else {

}
}
{
if (v2_3.Type == 9 && v2_3.IntVal == 786377863) {
var go__4_4 gopurs_runtime.Value
go__4_4 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
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
__t1 = gopurs_runtime.Apply2(go__4_4, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, v_1)
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
return gopurs_runtime.Apply(go__0_0, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil})
}()
	})
	return cache_take
}

var cache_takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		cache_takeWhile = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_takeWhile(p_0_box)
})
	})
	return cache_takeWhile
}

var cache_unsnoc gopurs_runtime.Value
var once_unsnoc sync.Once
func Get_unsnoc() gopurs_runtime.Value {
	once_unsnoc.Do(func() {
		cache_unsnoc = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsnoc(lst_0_box)
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
		cache_zipWithA = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWithA((*Record_pure_gopurs_runtime_Value)(dictApplicative_0_box.UnsafePtr))
})
	})
	return cache_zipWithA
}

var cache_range_ gopurs_runtime.Value
var once_range_ sync.Once
func Get_range_() gopurs_runtime.Value {
	once_range_.Do(func() {
		cache_range_ = gopurs_runtime.Func2(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_range_(start_0_box.IntVal, end_1_box.IntVal)
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
		cache_null = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_null(v_0_box))
})
	})
	return cache_null
}

var cache_nubBy gopurs_runtime.Value
var once_nubBy sync.Once
func Get_nubBy() gopurs_runtime.Value {
	once_nubBy.Do(func() {
		cache_nubBy = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_nubBy(p_0_box)
})
	})
	return cache_nubBy
}

var cache_nub gopurs_runtime.Value
var once_nub sync.Once
func Get_nub() gopurs_runtime.Value {
	once_nub.Do(func() {
		cache_nub = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_nub((*Record_compare_gopurs_runtime_Value)(dictOrd_0_box.UnsafePtr))
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
		cache_mapMaybe = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybe(f_0_box)
})
	})
	return cache_mapMaybe
}

var cache_manyRec gopurs_runtime.Value
var once_manyRec sync.Once
func Get_manyRec() gopurs_runtime.Value {
	once_manyRec.Do(func() {
		cache_manyRec = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictAlternative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_manyRec((*Record_tailRecM_gopurs_runtime_Value)(dictMonadRec_0_box.UnsafePtr), (*Record_)(dictAlternative_1_box.UnsafePtr))
})
	})
	return cache_manyRec
}

var cache_someRec gopurs_runtime.Value
var once_someRec sync.Once
func Get_someRec() gopurs_runtime.Value {
	once_someRec.Do(func() {
		cache_someRec = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictAlternative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_someRec((*Record_tailRecM_gopurs_runtime_Value)(dictMonadRec_0_box.UnsafePtr), (*Record_)(dictAlternative_1_box.UnsafePtr))
})
	})
	return cache_someRec
}

var cache_some gopurs_runtime.Value
var once_some sync.Once
func Get_some() gopurs_runtime.Value {
	once_some.Do(func() {
		cache_some = gopurs_runtime.Func3(func(dictAlternative_0_box gopurs_runtime.Value, dictLazy_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_some((*Record_)(dictAlternative_0_box.UnsafePtr), (*Record_defer__gopurs_runtime_Value)(dictLazy_1_box.UnsafePtr), v_2_box)
})
	})
	return cache_some
}

var cache_many gopurs_runtime.Value
var once_many sync.Once
func Get_many() gopurs_runtime.Value {
	once_many.Do(func() {
		cache_many = gopurs_runtime.Func3(func(dictAlternative_0_box gopurs_runtime.Value, dictLazy_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_many((*Record_)(dictAlternative_0_box.UnsafePtr), (*Record_defer__gopurs_runtime_Value)(dictLazy_1_box.UnsafePtr), v_2_box)
})
	})
	return cache_many
}

var cache_length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		cache_length = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func2(func(acc_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((acc_0.IntVal) + (1))
}), gopurs_runtime.Int(0))
	})
	return cache_length
}

var cache_last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		cache_last = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_last(v_0_box)
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
return Call_insertAt(v_0_box.IntVal, v1_1_box, v2_2_box)
})
	})
	return cache_insertAt
}

var cache_insert gopurs_runtime.Value
var once_insert sync.Once
func Get_insert() gopurs_runtime.Value {
	once_insert.Do(func() {
		cache_insert = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insert((*Record_compare_gopurs_runtime_Value)(dictOrd_0_box.UnsafePtr))
})
	})
	return cache_insert
}

var cache_init_ gopurs_runtime.Value
var once_init_ sync.Once
func Get_init_() gopurs_runtime.Value {
	once_init_.Do(func() {
		cache_init_ = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_init_(lst_0_box)
})
	})
	return cache_init_
}

var cache_index gopurs_runtime.Value
var once_index sync.Once
func Get_index() gopurs_runtime.Value {
	once_index.Do(func() {
		cache_index = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_index(v_0_box, v1_1_box.IntVal)
})
	})
	return cache_index
}

var cache_head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		cache_head = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_head(v_0_box)
})
	})
	return cache_head
}

var cache_transpose gopurs_runtime.Value
var once_transpose sync.Once
func Get_transpose() gopurs_runtime.Value {
	once_transpose.Do(func() {
		cache_transpose = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_transpose(v_0_box)
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
		cache_groupAllBy = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_groupAllBy(p_0_box)
})
	})
	return cache_groupAllBy
}

var cache_group gopurs_runtime.Value
var once_group sync.Once
func Get_group() gopurs_runtime.Value {
	once_group.Do(func() {
		cache_group = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_group((*Record_eq_gopurs_runtime_Value)(dictEq_0_box.UnsafePtr))
})
	})
	return cache_group
}

var cache_groupAll gopurs_runtime.Value
var once_groupAll sync.Once
func Get_groupAll() gopurs_runtime.Value {
	once_groupAll.Do(func() {
		cache_groupAll = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_groupAll((*Record_compare_gopurs_runtime_Value)(dictOrd_0_box.UnsafePtr))
})
	})
	return cache_groupAll
}

var cache_fromFoldable gopurs_runtime.Value
var once_fromFoldable sync.Once
func Get_fromFoldable() gopurs_runtime.Value {
	once_fromFoldable.Do(func() {
		cache_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldable((*Record_foldMap_gopurs_runtime_Value_foldl_gopurs_runtime_Value_foldr_gopurs_runtime_Value)(dictFoldable_0_box.UnsafePtr))
})
	})
	return cache_fromFoldable
}

var cache_foldM gopurs_runtime.Value
var once_foldM sync.Once
func Get_foldM() gopurs_runtime.Value {
	once_foldM.Do(func() {
		cache_foldM = gopurs_runtime.Func4(func(dictMonad_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value, v2_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldM((*Record_)(dictMonad_0_box.UnsafePtr), v_1_box, v1_2_box, v2_3_box)
})
	})
	return cache_foldM
}

var cache_findIndex gopurs_runtime.Value
var once_findIndex sync.Once
func Get_findIndex() gopurs_runtime.Value {
	once_findIndex.Do(func() {
		cache_findIndex = gopurs_runtime.Func(func(fn_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_findIndex(fn_0_box)
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
		cache_filterM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filterM((*Record_)(dictMonad_0_box.UnsafePtr))
})
	})
	return cache_filterM
}

var cache_filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		cache_filter = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_filter(p_0_box)
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
		cache_intersect = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intersect((*Record_eq_gopurs_runtime_Value)(dictEq_0_box.UnsafePtr))
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
		cache_nubEq = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_nubEq((*Record_eq_gopurs_runtime_Value)(dictEq_0_box.UnsafePtr))
})
	})
	return cache_nubEq
}

var cache_eqPattern gopurs_runtime.Value
var once_eqPattern sync.Once
func Get_eqPattern() gopurs_runtime.Value {
	once_eqPattern.Do(func() {
		cache_eqPattern = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqPattern((*Record_eq_gopurs_runtime_Value)(dictEq_0_box.UnsafePtr))
})
	})
	return cache_eqPattern
}

var cache_ordPattern gopurs_runtime.Value
var once_ordPattern sync.Once
func Get_ordPattern() gopurs_runtime.Value {
	once_ordPattern.Do(func() {
		cache_ordPattern = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordPattern((*Record_compare_gopurs_runtime_Value)(dictOrd_0_box.UnsafePtr))
})
	})
	return cache_ordPattern
}

var cache_elemLastIndex gopurs_runtime.Value
var once_elemLastIndex sync.Once
func Get_elemLastIndex() gopurs_runtime.Value {
	once_elemLastIndex.Do(func() {
		cache_elemLastIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_elemLastIndex((*Record_eq_gopurs_runtime_Value)(dictEq_0_box.UnsafePtr), x_1_box)
})
	})
	return cache_elemLastIndex
}

var cache_elemIndex gopurs_runtime.Value
var once_elemIndex sync.Once
func Get_elemIndex() gopurs_runtime.Value {
	once_elemIndex.Do(func() {
		cache_elemIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_elemIndex((*Record_eq_gopurs_runtime_Value)(dictEq_0_box.UnsafePtr), x_1_box)
})
	})
	return cache_elemIndex
}

var cache_dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		cache_dropWhile = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_dropWhile(p_0_box)
})
	})
	return cache_dropWhile
}

var cache_dropEnd gopurs_runtime.Value
var once_dropEnd sync.Once
func Get_dropEnd() gopurs_runtime.Value {
	once_dropEnd.Do(func() {
		cache_dropEnd = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_dropEnd(n_0_box.IntVal, xs_1_box)
})
	})
	return cache_dropEnd
}

var cache_drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		cache_drop = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_drop(v_0_box.IntVal, v1_1_box)
})
	})
	return cache_drop
}

var cache_slice gopurs_runtime.Value
var once_slice sync.Once
func Get_slice() gopurs_runtime.Value {
	once_slice.Do(func() {
		cache_slice = gopurs_runtime.Func3(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_slice(start_0_box.IntVal, end_1_box.IntVal, xs_2_box)
})
	})
	return cache_slice
}

var cache_takeEnd gopurs_runtime.Value
var once_takeEnd sync.Once
func Get_takeEnd() gopurs_runtime.Value {
	once_takeEnd.Do(func() {
		cache_takeEnd = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_takeEnd(n_0_box.IntVal, xs_1_box)
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
		cache_union = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_union((*Record_eq_gopurs_runtime_Value)(dictEq_0_box.UnsafePtr))
})
	})
	return cache_union
}

var cache_deleteAt gopurs_runtime.Value
var once_deleteAt sync.Once
func Get_deleteAt() gopurs_runtime.Value {
	once_deleteAt.Do(func() {
		cache_deleteAt = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_deleteAt(v_0_box.IntVal, v1_1_box)
})
	})
	return cache_deleteAt
}

var cache_delete_ gopurs_runtime.Value
var once_delete_ sync.Once
func Get_delete_() gopurs_runtime.Value {
	once_delete_.Do(func() {
		cache_delete_ = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_delete_((*Record_eq_gopurs_runtime_Value)(dictEq_0_box.UnsafePtr))
})
	})
	return cache_delete_
}

var cache_difference gopurs_runtime.Value
var once_difference sync.Once
func Get_difference() gopurs_runtime.Value {
	once_difference.Do(func() {
		cache_difference = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_difference((*Record_eq_gopurs_runtime_Value)(dictEq_0_box.UnsafePtr))
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
		cache_concat = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_concat(v_0_box)
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
go__0_0 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop gopurs_runtime.Value = v_1_loop_val
var v1_2_loop gopurs_runtime.Value = v1_2_loop_val
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
go__3_2 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
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
__t1 = gopurs_runtime.Apply2(go__3_2, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, v_1)
goto end_branch_1
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437) {
var __t4 gopurs_runtime.Value
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_2.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 3589588149) {
v_1_loop = v_1
v1_2_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_2.UnsafePtr).V1
continue go__0_0
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_2.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 930809136) {
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
return gopurs_runtime.Apply(go__0_0, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil})
}()
	})
	return cache_catMaybes
}

var cache_alterAt gopurs_runtime.Value
var once_alterAt sync.Once
func Get_alterAt() gopurs_runtime.Value {
	once_alterAt.Do(func() {
		cache_alterAt = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alterAt(v_0_box.IntVal, v1_1_box, v2_2_box)
})
	})
	return cache_alterAt
}

var cache_modifyAt gopurs_runtime.Value
var once_modifyAt sync.Once
func Get_modifyAt() gopurs_runtime.Value {
	once_modifyAt.Do(func() {
		cache_modifyAt = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modifyAt(n_0_box.IntVal, f_1_box)
})
	})
	return cache_modifyAt
}

type Record_alt_gopurs_runtime_Value struct {
	alt gopurs_runtime.Value
}

type Record_ struct {
	
}

type Record_pure_gopurs_runtime_Value struct {
	pure gopurs_runtime.Value
}

type Record_apply_gopurs_runtime_Value struct {
	apply gopurs_runtime.Value
}

type Record_bipure_gopurs_runtime_Value struct {
	bipure gopurs_runtime.Value
}

type Record_biapply_gopurs_runtime_Value struct {
	biapply gopurs_runtime.Value
}

type Record_bind_gopurs_runtime_Value struct {
	bind gopurs_runtime.Value
}

type Record_discard_gopurs_runtime_Value struct {
	discard gopurs_runtime.Value
}

type Record_identity_gopurs_runtime_Value struct {
	identity gopurs_runtime.Value
}

type Record_ask_gopurs_runtime_Value struct {
	ask gopurs_runtime.Value
}

type Record_local_gopurs_runtime_Value struct {
	local gopurs_runtime.Value
}

type Record_peek_gopurs_runtime_Value_pos_gopurs_runtime_Value struct {
	peek gopurs_runtime.Value
	pos gopurs_runtime.Value
}

type Record_track_gopurs_runtime_Value struct {
	track gopurs_runtime.Value
}

type Record_extract_gopurs_runtime_Value struct {
	extract gopurs_runtime.Value
}

type Record_extend_gopurs_runtime_Value struct {
	extend gopurs_runtime.Value
}

type Record_defer__gopurs_runtime_Value struct {
	defer_ gopurs_runtime.Value
}

type Record_callCC_gopurs_runtime_Value struct {
	callCC gopurs_runtime.Value
}

type Record_catchError_gopurs_runtime_Value struct {
	catchError gopurs_runtime.Value
}

type Record_throwError_gopurs_runtime_Value struct {
	throwError gopurs_runtime.Value
}

type Record_chooseBool_gopurs_runtime_Value_chooseFloat_gopurs_runtime_Value_chooseInt_gopurs_runtime_Value_resize_gopurs_runtime_Value_sized_gopurs_runtime_Value struct {
	chooseBool gopurs_runtime.Value
	chooseFloat gopurs_runtime.Value
	chooseInt gopurs_runtime.Value
	resize gopurs_runtime.Value
	sized gopurs_runtime.Value
}

type Record_foldMap1_gopurs_runtime_Value_foldl1_gopurs_runtime_Value_foldr1_gopurs_runtime_Value struct {
	foldMap1 gopurs_runtime.Value
	foldl1 gopurs_runtime.Value
	foldr1 gopurs_runtime.Value
}

type Record_append__gopurs_runtime_Value struct {
	append_ gopurs_runtime.Value
}

type Record_tailRecM_gopurs_runtime_Value struct {
	tailRecM gopurs_runtime.Value
}

type Record_unfoldr_gopurs_runtime_Value struct {
	unfoldr gopurs_runtime.Value
}

type Record_map__gopurs_runtime_Value struct {
	map_ gopurs_runtime.Value
}

type Record_state_gopurs_runtime_Value struct {
	state gopurs_runtime.Value
}

type Record_lift_gopurs_runtime_Value struct {
	lift gopurs_runtime.Value
}

type Record_listen_gopurs_runtime_Value_pass_gopurs_runtime_Value struct {
	listen gopurs_runtime.Value
	pass gopurs_runtime.Value
}

type Record_parallel_gopurs_runtime_Value_sequential_gopurs_runtime_Value struct {
	parallel gopurs_runtime.Value
	sequential gopurs_runtime.Value
}

type Record_foldMap_gopurs_runtime_Value_foldl_gopurs_runtime_Value_foldr_gopurs_runtime_Value struct {
	foldMap gopurs_runtime.Value
	foldl gopurs_runtime.Value
	foldr gopurs_runtime.Value
}

type Record_mempty_gopurs_runtime_Value struct {
	mempty gopurs_runtime.Value
}

type Record_sequence_gopurs_runtime_Value_traverse_gopurs_runtime_Value struct {
	sequence gopurs_runtime.Value
	traverse gopurs_runtime.Value
}

type Record_empty_gopurs_runtime_Value struct {
	empty gopurs_runtime.Value
}

type Record_compose_gopurs_runtime_Value struct {
	compose gopurs_runtime.Value
}

type Record_eq_gopurs_runtime_Value struct {
	eq gopurs_runtime.Value
}

type Record_compare_gopurs_runtime_Value struct {
	compare gopurs_runtime.Value
}

type Record_bifoldMap_gopurs_runtime_Value_bifoldl_gopurs_runtime_Value_bifoldr_gopurs_runtime_Value struct {
	bifoldMap gopurs_runtime.Value
	bifoldl gopurs_runtime.Value
	bifoldr gopurs_runtime.Value
}

type Record_bimap_gopurs_runtime_Value struct {
	bimap gopurs_runtime.Value
}

type Record_bisequence_gopurs_runtime_Value_bitraverse_gopurs_runtime_Value struct {
	bisequence gopurs_runtime.Value
	bitraverse gopurs_runtime.Value
}

type Record_genericBottom_prime_gopurs_runtime_Value struct {
	genericBottom_prime gopurs_runtime.Value
}

type Record_genericTop_prime_gopurs_runtime_Value struct {
	genericTop_prime gopurs_runtime.Value
}

type Record_bottom_gopurs_runtime_Value_top_gopurs_runtime_Value struct {
	bottom gopurs_runtime.Value
	top gopurs_runtime.Value
}

type Record_lose_gopurs_runtime_Value struct {
	lose gopurs_runtime.Value
}

type Record_choose_gopurs_runtime_Value struct {
	choose gopurs_runtime.Value
}

type Record_collect_gopurs_runtime_Value_distribute_gopurs_runtime_Value struct {
	collect gopurs_runtime.Value
	distribute gopurs_runtime.Value
}

type Record_divide_gopurs_runtime_Value struct {
	divide gopurs_runtime.Value
}

type Record_recip_gopurs_runtime_Value struct {
	recip gopurs_runtime.Value
}

type Record_genericCardinality_prime_gopurs_runtime_Value_genericFromEnum_prime_gopurs_runtime_Value_genericToEnum_prime_gopurs_runtime_Value struct {
	genericCardinality_prime gopurs_runtime.Value
	genericFromEnum_prime gopurs_runtime.Value
	genericToEnum_prime gopurs_runtime.Value
}

type Record_genericPred_prime_gopurs_runtime_Value_genericSucc_prime_gopurs_runtime_Value struct {
	genericPred_prime gopurs_runtime.Value
	genericSucc_prime gopurs_runtime.Value
}

type Record_pred_gopurs_runtime_Value_succ_gopurs_runtime_Value struct {
	pred gopurs_runtime.Value
	succ gopurs_runtime.Value
}

type Record_unfoldr1_gopurs_runtime_Value struct {
	unfoldr1 gopurs_runtime.Value
}

type Record_cardinality_gopurs_runtime_Value_fromEnum_gopurs_runtime_Value_toEnum_gopurs_runtime_Value struct {
	cardinality gopurs_runtime.Value
	fromEnum gopurs_runtime.Value
	toEnum gopurs_runtime.Value
}

type Record_genericEq_prime_gopurs_runtime_Value struct {
	genericEq_prime gopurs_runtime.Value
}

type Record_eq1_gopurs_runtime_Value struct {
	eq1 gopurs_runtime.Value
}

type Record_degree_gopurs_runtime_Value_div_gopurs_runtime_Value_mod_gopurs_runtime_Value struct {
	degree gopurs_runtime.Value
	div gopurs_runtime.Value
	mod gopurs_runtime.Value
}

type Record_conj_gopurs_runtime_Value_disj_gopurs_runtime_Value_ff_gopurs_runtime_Value_implies_gopurs_runtime_Value_not_gopurs_runtime_Value_tt_gopurs_runtime_Value struct {
	conj gopurs_runtime.Value
	disj gopurs_runtime.Value
	ff gopurs_runtime.Value
	implies gopurs_runtime.Value
	not gopurs_runtime.Value
	tt gopurs_runtime.Value
}

type Record_add_gopurs_runtime_Value_mul_gopurs_runtime_Value_one_gopurs_runtime_Value_zero_gopurs_runtime_Value struct {
	add gopurs_runtime.Value
	mul gopurs_runtime.Value
	one gopurs_runtime.Value
	zero gopurs_runtime.Value
}

type Record_foldMapWithIndex_gopurs_runtime_Value_foldlWithIndex_gopurs_runtime_Value_foldrWithIndex_gopurs_runtime_Value struct {
	foldMapWithIndex gopurs_runtime.Value
	foldlWithIndex gopurs_runtime.Value
	foldrWithIndex gopurs_runtime.Value
}

type Record_cmap_gopurs_runtime_Value struct {
	cmap gopurs_runtime.Value
}

type Record_imap_gopurs_runtime_Value struct {
	imap gopurs_runtime.Value
}

type Record_mapWithIndex_gopurs_runtime_Value struct {
	mapWithIndex gopurs_runtime.Value
}

type Record_from_gopurs_runtime_Value_to_gopurs_runtime_Value struct {
	from gopurs_runtime.Value
	to gopurs_runtime.Value
}

type Record_genericConj_prime_gopurs_runtime_Value_genericDisj_prime_gopurs_runtime_Value_genericFF_prime_gopurs_runtime_Value_genericImplies_prime_gopurs_runtime_Value_genericNot_prime_gopurs_runtime_Value_genericTT_prime_gopurs_runtime_Value struct {
	genericConj_prime gopurs_runtime.Value
	genericDisj_prime gopurs_runtime.Value
	genericFF_prime gopurs_runtime.Value
	genericImplies_prime gopurs_runtime.Value
	genericNot_prime gopurs_runtime.Value
	genericTT_prime gopurs_runtime.Value
}

type Record_conj_gopurs_runtime_Value_disj_gopurs_runtime_Value_ff_bool_implies_gopurs_runtime_Value_not_gopurs_runtime_Value_tt_bool struct {
	conj gopurs_runtime.Value
	disj gopurs_runtime.Value
	ff bool
	implies gopurs_runtime.Value
	not gopurs_runtime.Value
	tt bool
}

type Record_genericMempty_prime_gopurs_runtime_Value struct {
	genericMempty_prime gopurs_runtime.Value
}

type Record_genericCompare_prime_gopurs_runtime_Value struct {
	genericCompare_prime gopurs_runtime.Value
}

type Record_sub_gopurs_runtime_Value struct {
	sub gopurs_runtime.Value
}

type Record_compare1_gopurs_runtime_Value struct {
	compare1 gopurs_runtime.Value
}

type Record_left_gopurs_runtime_Value_right_gopurs_runtime_Value struct {
	left gopurs_runtime.Value
	right gopurs_runtime.Value
}

type Record_first_gopurs_runtime_Value_second_gopurs_runtime_Value struct {
	first gopurs_runtime.Value
	second gopurs_runtime.Value
}

type Record_dimap_gopurs_runtime_Value struct {
	dimap gopurs_runtime.Value
}

type Record_genericSub_prime_gopurs_runtime_Value struct {
	genericSub_prime gopurs_runtime.Value
}

type Record_genericAppend_prime_gopurs_runtime_Value struct {
	genericAppend_prime gopurs_runtime.Value
}

type Record_sequence1_gopurs_runtime_Value_traverse1_gopurs_runtime_Value struct {
	sequence1 gopurs_runtime.Value
	traverse1 gopurs_runtime.Value
}

type Record_genericAdd_prime_gopurs_runtime_Value_genericMul_prime_gopurs_runtime_Value_genericOne_prime_gopurs_runtime_Value_genericZero_prime_gopurs_runtime_Value struct {
	genericAdd_prime gopurs_runtime.Value
	genericMul_prime gopurs_runtime.Value
	genericOne_prime gopurs_runtime.Value
	genericZero_prime gopurs_runtime.Value
}

type Record_genericShow_prime_gopurs_runtime_Value struct {
	genericShow_prime gopurs_runtime.Value
}

type Record_genericShowArgs_gopurs_runtime_Value struct {
	genericShowArgs gopurs_runtime.Value
}

type Record_show_gopurs_runtime_Value struct {
	show gopurs_runtime.Value
}

type Record_fromDuration_gopurs_runtime_Value_toDuration_gopurs_runtime_Value struct {
	fromDuration gopurs_runtime.Value
	toDuration gopurs_runtime.Value
}

type Record_traverseWithIndex_gopurs_runtime_Value struct {
	traverseWithIndex gopurs_runtime.Value
}

type Record_liftEffect_gopurs_runtime_Value struct {
	liftEffect gopurs_runtime.Value
}

type Record_mappend__gopurs_runtime_Value_mempty__gopurs_runtime_Value struct {
	mappend_ gopurs_runtime.Value
	mempty_ gopurs_runtime.Value
}

type Record_proof_gopurs_runtime_Value struct {
	proof gopurs_runtime.Value
}

type Record_lower_gopurs_runtime_Value struct {
	lower gopurs_runtime.Value
}

type Record_liftST_gopurs_runtime_Value struct {
	liftST gopurs_runtime.Value
}

type Record_tell_gopurs_runtime_Value struct {
	tell gopurs_runtime.Value
}

type Record_reflectSymbol_gopurs_runtime_Value struct {
	reflectSymbol gopurs_runtime.Value
}

type Record_bottomRecord_gopurs_runtime_Value_topRecord_gopurs_runtime_Value struct {
	bottomRecord gopurs_runtime.Value
	topRecord gopurs_runtime.Value
}

type Record_conquer_gopurs_runtime_Value struct {
	conquer gopurs_runtime.Value
}

type Record_inj_gopurs_runtime_Value_prj_gopurs_runtime_Value struct {
	inj gopurs_runtime.Value
	prj gopurs_runtime.Value
}

type Record_eqRecord_gopurs_runtime_Value struct {
	eqRecord gopurs_runtime.Value
}

type Record_conjRecord_gopurs_runtime_Value_disjRecord_gopurs_runtime_Value_ffRecord_gopurs_runtime_Value_impliesRecord_gopurs_runtime_Value_notRecord_gopurs_runtime_Value_ttRecord_gopurs_runtime_Value struct {
	conjRecord gopurs_runtime.Value
	disjRecord gopurs_runtime.Value
	ffRecord gopurs_runtime.Value
	impliesRecord gopurs_runtime.Value
	notRecord gopurs_runtime.Value
	ttRecord gopurs_runtime.Value
}

type Record_memptyRecord_gopurs_runtime_Value struct {
	memptyRecord gopurs_runtime.Value
}

type Record_compareRecord_gopurs_runtime_Value struct {
	compareRecord gopurs_runtime.Value
}

type Record_closed_gopurs_runtime_Value struct {
	closed gopurs_runtime.Value
}

type Record_unleft_gopurs_runtime_Value_unright_gopurs_runtime_Value struct {
	unleft gopurs_runtime.Value
	unright gopurs_runtime.Value
}

type Record_unfirst_gopurs_runtime_Value_unsecond_gopurs_runtime_Value struct {
	unfirst gopurs_runtime.Value
	unsecond gopurs_runtime.Value
}

type Record_reflectType_gopurs_runtime_Value struct {
	reflectType gopurs_runtime.Value
}

type Record_subRecord_gopurs_runtime_Value struct {
	subRecord gopurs_runtime.Value
}

type Record_appendRecord_gopurs_runtime_Value struct {
	appendRecord gopurs_runtime.Value
}

type Record_addRecord_gopurs_runtime_Value_mulRecord_gopurs_runtime_Value_oneRecord_gopurs_runtime_Value_zeroRecord_gopurs_runtime_Value struct {
	addRecord gopurs_runtime.Value
	mulRecord gopurs_runtime.Value
	oneRecord gopurs_runtime.Value
	zeroRecord gopurs_runtime.Value
}

type Record_showRecordFields_gopurs_runtime_Value struct {
	showRecordFields gopurs_runtime.Value
}

type Record_nes_gopurs_runtime_Value struct {
	nes gopurs_runtime.Value
}

type Record_liftAff_gopurs_runtime_Value struct {
	liftAff gopurs_runtime.Value
}

func Call_tailRecM2(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_Rec_Class.Get_monadRecMaybe(), "tailRecM"), gopurs_runtime.Func(func(o_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, gopurs_runtime.RecordGet(o_3, "a"), gopurs_runtime.RecordGet(o_3, "b"))
}), gopurs_runtime.RecordDict2("a", "b", a_1, b_2))
}

func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Pattern(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_updateAt(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
updateAt:
for {
if false { continue updateAt }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 1358893437) {
var __t2 gopurs_runtime.Value
{
if (v_0) == (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{v1_1, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V1})}})}
goto end_branch_2
} else {

}
}
{
__local_var_3_1 := (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V0
_ = __local_var_3_1
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v3_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{__local_var_3_1, v3_4})}
}), Call_updateAt((v_0) - (1), v1_1, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V1))
}
end_branch_2:
__t0 = __t2
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
}
end_branch_0:
return __t0
}
}

func Call_uncons(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 786377863) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
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
}

func Call_toUnfoldable(dictUnfoldable_0_loop *Record_unfoldr_gopurs_runtime_Value) gopurs_runtime.Value {
var dictUnfoldable_0 *Record_unfoldr_gopurs_runtime_Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
return gopurs_runtime.Apply(dictUnfoldable_0.unfoldr, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (xs_1.Type == 9 && xs_1.IntVal == 786377863) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (xs_1.Type == 9 && xs_1.IntVal == 1358893437) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.RecordDict2("head", "tail", (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(xs_1.UnsafePtr).V0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(xs_1.UnsafePtr).V1)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(rec_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{gopurs_runtime.RecordGet(rec_2, "head"), gopurs_runtime.RecordGet(rec_2, "tail")})}
}), __t0)
}))
}

func Call_tail(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 786377863) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
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
}

func Call_stripPrefix(dictEq_0_loop *Record_eq_gopurs_runtime_Value, v_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 *Record_eq_gopurs_runtime_Value = dictEq_0_loop
_ = dictEq_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
return Call_tailRecM2(gopurs_runtime.Func2(func(prefix_3 gopurs_runtime.Value, input_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (input_4.Type == 9 && input_4.IntVal == 1358893437) {
var __t1 gopurs_runtime.Value
{
if (prefix_3.Type == 9 && prefix_3.IntVal == 1358893437) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(dictEq_0.eq, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(prefix_3.UnsafePtr).V0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(input_4.UnsafePtr).V0).IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Data_Control_Monad_Rec_Class_Loop{gopurs_runtime.RecordDict2("a", "b", (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(prefix_3.UnsafePtr).V1, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(input_4.UnsafePtr).V1)})}})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if (prefix_3.Type == 9 && prefix_3.IntVal == 786377863) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Data_Control_Monad_Rec_Class_Done{input_4})}})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
}
end_branch_0:
return __t0
}), v_1, s_2)
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
__t0 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, v1_1)
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
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{x_1, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}})}, xs_0)
}

func Call_singleton(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{a_0, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}})}
}

func Call_sortBy(cmp_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Apply2(cmp_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_2.UnsafePtr).V0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}).IntVal) != (0) {
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
var __t_and_7 bool = false
if (v_3.Type == 9 && v_3.IntVal == 1358893437) {

var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_3.UnsafePtr).V1
__t_and_7 = (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 1358893437)
}
if __t_and_7 {
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
var mergeAll_3_8 gopurs_runtime.Value
mergeAll_3_8 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
mergeAll_3_8:
for {
if false { continue mergeAll_3_8 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t9 gopurs_runtime.Value
{
var __t_and_11 bool = false
if (v_4.Type == 9 && v_4.IntVal == 1358893437) {

var __t_tag_10 gopurs_runtime.Value = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V1
__t_and_11 = (__t_tag_10.Type == 9 && __t_tag_10.IntVal == 786377863)
}
if __t_and_11 {
__t9 = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0
goto end_branch_9
} else {

}
}
{
v_4_loop = gopurs_runtime.Apply(mergePairs_2_4, v_4)
continue mergeAll_3_8
__t9 = gopurs_runtime.Value{}
}
end_branch_9:
return __t9
}
}()
})
var sequences_4_12 gopurs_runtime.Value
_ = sequences_4_12
var descending_4_13 gopurs_runtime.Value
_ = descending_4_13
var ascending_4_14 gopurs_runtime.Value
_ = ascending_4_14
sequences_4_12 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
var __t_and_17 bool = false
if (v_5.Type == 9 && v_5.IntVal == 1358893437) {

var __t_tag_16 gopurs_runtime.Value = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_5.UnsafePtr).V1
__t_and_17 = (__t_tag_16.Type == 9 && __t_tag_16.IntVal == 1358893437)
}
if __t_and_17 {
var __t19 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Apply2(cmp_0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_5.UnsafePtr).V0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_5.UnsafePtr).V1.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}).IntVal) != (0) {
__t19 = gopurs_runtime.Apply3(descending_4_13, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_5.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_5.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}})}, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_5.UnsafePtr).V1.UnsafePtr).V1)
goto end_branch_19
} else {

}
}
{
__local_var_6_18 := (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_5.UnsafePtr).V0
_ = __local_var_6_18
__t19 = gopurs_runtime.Apply3(ascending_4_14, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_5.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{__local_var_6_18, v1_7})}
}), (*pkg_Data_List_Types.Data_Data_List_Types_Cons)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_5.UnsafePtr).V1.UnsafePtr).V1)
}
end_branch_19:
__t15 = __t19
goto end_branch_15
} else {

}
}
{
__t15 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{v_5, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}})}
}
end_branch_15:
return __t15
})
descending_4_13 = gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t20 gopurs_runtime.Value
{
if ((v2_7.Type == 9 && v2_7.IntVal == 1358893437)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Apply2(cmp_0, v_5, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_7.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}).IntVal) != (0)) {
__t20 = gopurs_runtime.Apply3(descending_4_13, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_7.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{v_5, v1_6})}, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_7.UnsafePtr).V1)
goto end_branch_20
} else {

}
}
{
__t20 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{v_5, v1_6})}, gopurs_runtime.Apply(sequences_4_12, v2_7)})}
}
end_branch_20:
return __t20
})
ascending_4_14 = gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t21 gopurs_runtime.Value
{
if ((v2_7.Type == 9 && v2_7.IntVal == 1358893437)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Apply2(cmp_0, v_5, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_7.UnsafePtr).V0), gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}), gopurs_runtime.Bool(false)).IntVal) != (0)) {
__t21 = gopurs_runtime.Apply3(ascending_4_14, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_7.UnsafePtr).V0, gopurs_runtime.Func(func(ys_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{v_5, ys_8})})
}), (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_7.UnsafePtr).V1)
goto end_branch_21
} else {

}
}
{
__t21 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Apply(v1_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{v_5, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}})}), gopurs_runtime.Apply(sequences_4_12, v2_7)})}
}
end_branch_21:
return __t21
})
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(mergeAll_3_8, gopurs_runtime.Apply(sequences_4_12, x_5))
})
}

func Call_sort(dictOrd_0_loop *Record_compare_gopurs_runtime_Value) gopurs_runtime.Value {
var dictOrd_0 *Record_compare_gopurs_runtime_Value = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.compare
_ = compare_1_0
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_sortBy(), compare_1_0, xs_2)
})
}

func Call_showPattern(dictShow_0_loop *Record_show_gopurs_runtime_Value) gopurs_runtime.Value {
var dictShow_0 *Record_show_gopurs_runtime_Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Pattern "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(pkg_Data_List_Types.Get_showList(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictShow_0)}), "show"), v_1), gopurs_runtime.Str(")")))
}))
}

func Call_takeWhile(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
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
go__4_1 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
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
__t3 = gopurs_runtime.Apply2(go__4_1, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, v_2)
}
end_branch_3:
return __t3
}
}()
})
})
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil})
}

func Call_unsnoc(lst_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var lst_0 gopurs_runtime.Value = lst_0_loop
_ = lst_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437) {
var __t2 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_2.UnsafePtr).V1
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 786377863) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.RecordDict2("last", "revInit", (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_2.UnsafePtr).V0, v1_3)})}
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(h_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__3_4 gopurs_runtime.Value
go__3_4 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
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
return gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.Apply2(go__3_4, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, gopurs_runtime.RecordGet(h_2, "revInit")), gopurs_runtime.RecordGet(h_2, "last"))
}), gopurs_runtime.Apply2(go__1_0, lst_0, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}))
}

func Call_zipWith(f_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
var go__3_0 gopurs_runtime.Value
go__3_0 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
var v2_6_loop gopurs_runtime.Value = v2_6_loop_val
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
go__4_2 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
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
return gopurs_runtime.Apply2(go__4_2, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, gopurs_runtime.Apply3(go__3_0, xs_1, ys_2, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}))
}

func Call_zip(xs_0_loop gopurs_runtime.Value, ys_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var ys_1 gopurs_runtime.Value = ys_1_loop
_ = ys_1
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
var v1_4_loop gopurs_runtime.Value = v1_4_loop_val
var v2_5_loop gopurs_runtime.Value = v2_5_loop_val
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
go__3_2 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
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
return gopurs_runtime.Apply2(go__3_2, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, gopurs_runtime.Apply3(go__2_0, xs_0, ys_1, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}))
}

func Call_zipWithA(dictApplicative_0_loop *Record_pure_gopurs_runtime_Value) gopurs_runtime.Value {
var dictApplicative_0 *Record_pure_gopurs_runtime_Value = dictApplicative_0_loop
_ = dictApplicative_0
sequence1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_traversableList(), "sequence"), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictApplicative_0)})
_ = sequence1_1_0
return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value, ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__5_1 gopurs_runtime.Value
go__5_1 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
var v2_8_loop gopurs_runtime.Value = v2_8_loop_val
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
go__6_3 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_7_loop gopurs_runtime.Value = v_7_loop_val
var v1_8_loop gopurs_runtime.Value = v1_8_loop_val
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
return gopurs_runtime.Apply(sequence1_1_0, gopurs_runtime.Apply2(go__6_3, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, gopurs_runtime.Apply3(go__5_1, xs_3, ys_4, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil})))
})
}

func Call_range_(start_0_loop int64, end_1_loop int64) gopurs_runtime.Value {
var start_0 int64 = start_0_loop
_ = start_0
var end_1 int64 = end_1_loop
_ = end_1
var __t3 gopurs_runtime.Value
{
if (start_0) == (end_1) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Int(start_0), gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}})}
goto end_branch_3
} else {

}
}
{
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(s_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(step_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rest_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var s_3_loop gopurs_runtime.Value = s_3_loop_val
var e_4_loop gopurs_runtime.Value = e_4_loop_val
var step_5_loop gopurs_runtime.Value = step_5_loop_val
var rest_6_loop gopurs_runtime.Value = rest_6_loop_val
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
if (gopurs_runtime.Apply2(Get_greaterThan(), gopurs_runtime.Int(start_0), gopurs_runtime.Int(end_1)).IntVal) != (0) {
__t2 = gopurs_runtime.Int(1)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Int(-1)
}
end_branch_2:
__t3 = gopurs_runtime.Apply4(go__2_0, gopurs_runtime.Int(end_1), gopurs_runtime.Int(start_0), __t2, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil})
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
}), gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}), xs_1)
}

func Call_null(v_0_loop gopurs_runtime.Value) bool {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (v_0.Type == 9 && v_0.IntVal == 786377863)
}

func Call_nubBy(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
var v2_4_loop gopurs_runtime.Value = v2_4_loop_val
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
__local_var_2_4 := gopurs_runtime.Apply2(go__1_0, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil})
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__4_5 gopurs_runtime.Value
go__4_5 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
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
return gopurs_runtime.Apply2(go__4_5, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, gopurs_runtime.Apply(__local_var_2_4, x_3))
})
}

func Call_nub(dictOrd_0_loop *Record_compare_gopurs_runtime_Value) gopurs_runtime.Value {
var dictOrd_0 *Record_compare_gopurs_runtime_Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_nubBy(), dictOrd_0.compare)
}

func Call_mapMaybe(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
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
go__4_2 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
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
__t1 = gopurs_runtime.Apply2(go__4_2, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, v_2)
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
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil})
}

func Call_manyRec(dictMonadRec_0_loop *Record_tailRecM_gopurs_runtime_Value, dictAlternative_1_loop *Record_) gopurs_runtime.Value {
var dictMonadRec_0 *Record_tailRecM_gopurs_runtime_Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictAlternative_1 *Record_ = dictAlternative_1_loop
_ = dictAlternative_1
Alt0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictAlternative_1)}, "Plus1_NOT_FOUND"), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{})
_ = Alt0_2_0
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictAlternative_1)}, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_3_1
return gopurs_runtime.Func(func(p_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictMonadRec_0.tailRecM, gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonadRec_0)}, "Monad0_NOT_FOUND"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Alt0_2_0, "alt"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Alt0_2_0, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Control_Monad_Rec_Class.Get_Loop(), p_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Data_Control_Monad_Rec_Class_Done{pkg_Data_Unit.Get_unit()})})), gopurs_runtime.Func(func(aa_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "pure"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Control_Monad_Rec_Class.Get_bifunctorStep(), "bimap"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{v_7, acc_5})}
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var go__8_2 gopurs_runtime.Value
go__8_2 = gopurs_runtime.Func(func(v_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_9_loop gopurs_runtime.Value = v_9_loop_val
var v1_10_loop gopurs_runtime.Value = v1_10_loop_val
go__8_2:
for {
if false { continue go__8_2 }
var v_9 gopurs_runtime.Value = v_9_loop
_ = v_9
var v1_10 gopurs_runtime.Value = v1_10_loop
_ = v1_10
var __t3 gopurs_runtime.Value
{
if (v1_10.Type == 9 && v1_10.IntVal == 786377863) {
__t3 = v_9
goto end_branch_3
} else {

}
}
{
if (v1_10.Type == 9 && v1_10.IntVal == 1358893437) {
v_9_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_10.UnsafePtr).V0, v_9})}
v1_10_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_10.UnsafePtr).V1
continue go__8_2
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
return gopurs_runtime.Apply2(go__8_2, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, acc_5)
}), aa_6))
}))
}), gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil})
})
}

func Call_someRec(dictMonadRec_0_loop *Record_tailRecM_gopurs_runtime_Value, dictAlternative_1_loop *Record_) gopurs_runtime.Value {
var dictMonadRec_0 *Record_tailRecM_gopurs_runtime_Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictAlternative_1 *Record_ = dictAlternative_1_loop
_ = dictAlternative_1
manyRec2_2_0 := Call_manyRec(dictMonadRec_0, dictAlternative_1)
_ = manyRec2_2_0
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictAlternative_1)}, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictAlternative_1)}, "Plus1_NOT_FOUND"), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_List_Types.Get_Cons(), v_3), gopurs_runtime.Apply(manyRec2_2_0, v_3))
})
}

func Call_some(dictAlternative_0_loop *Record_, dictLazy_1_loop *Record_defer__gopurs_runtime_Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 *Record_ = dictAlternative_0_loop
_ = dictAlternative_0
var dictLazy_1 *Record_defer__gopurs_runtime_Value = dictLazy_1_loop
_ = dictLazy_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictAlternative_0)}, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictAlternative_0)}, "Plus1_NOT_FOUND"), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_List_Types.Get_Cons(), v_2), gopurs_runtime.Apply(dictLazy_1.defer_, gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_many(dictAlternative_0, dictLazy_1, v_2)
})))
}

func Call_many(dictAlternative_0_loop *Record_, dictLazy_1_loop *Record_defer__gopurs_runtime_Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 *Record_ = dictAlternative_0_loop
_ = dictAlternative_0
var dictLazy_1 *Record_defer__gopurs_runtime_Value = dictLazy_1_loop
_ = dictLazy_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictAlternative_0)}, "Plus1_NOT_FOUND"), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "alt"), Call_some(dictAlternative_0, dictLazy_1, v_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictAlternative_0)}, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}))
}

func Call_last(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 786377863) {
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
}
end_branch_0:
return __t0
}
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{v1_1, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 1358893437) {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(v_0, v1_1, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V0)
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 380165415) {
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

func Call_insertAt(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
insertAt:
for {
if false { continue insertAt }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v_0) == (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{v1_1, v2_2})}})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 1358893437) {
__local_var_3_1 := (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V0
_ = __local_var_3_1
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v3_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{__local_var_3_1, v3_4})}
}), Call_insertAt((v_0) - (1), v1_1, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V1))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
}
end_branch_0:
return __t0
}
}

func Call_insert(dictOrd_0_loop *Record_compare_gopurs_runtime_Value) gopurs_runtime.Value {
var dictOrd_0 *Record_compare_gopurs_runtime_Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_insertBy(), dictOrd_0.compare)
}

func Call_init_(lst_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var lst_0 gopurs_runtime.Value = lst_0_loop
_ = lst_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v_1, "init")
}), gopurs_runtime.Apply(Get_unsnoc(), lst_0))
}

func Call_index(v_0_loop gopurs_runtime.Value, v1_1_loop int64) gopurs_runtime.Value {
index:
for {
if false { continue index }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 int64 = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 786377863) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437) {
var __t1 gopurs_runtime.Value
{
if (v1_1) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0})}
goto end_branch_1
} else {

}
}
{
v_0_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V1
v1_1_loop = (v1_1) - (1)
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

func Call_head(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 786377863) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
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
}

func Call_transpose(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
transpose:
for {
if false { continue transpose }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 786377863) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437) {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 786377863) {
v_0_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V1
continue transpose
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 1358893437) {
var go__1_4 gopurs_runtime.Value
go__1_4 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__1_4:
for {
if false { continue go__1_4 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t5 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 786377863) {
var go__4_6 gopurs_runtime.Value
go__4_6 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__4_6:
for {
if false { continue go__4_6 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t7 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 786377863) {
__t7 = v_5
goto end_branch_7
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})}
v1_6_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V1
continue go__4_6
__t7 = gopurs_runtime.Value{}
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
}()
})
})
__t5 = gopurs_runtime.Apply2(go__4_6, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, v_2)
goto end_branch_5
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437) {
var __t8 gopurs_runtime.Value
{
var __t_tag_9 gopurs_runtime.Value = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 786377863) {
v_2_loop = v_2
v1_3_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V1
continue go__1_4
__t8 = gopurs_runtime.Value{}
goto end_branch_8
} else {

}
}
{
var __t_tag_10 gopurs_runtime.Value = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0
if (__t_tag_10.Type == 9 && __t_tag_10.IntVal == 1358893437) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0.UnsafePtr).V0, v_2})}
v1_3_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V1
continue go__1_4
__t8 = gopurs_runtime.Value{}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
__t5 = __t8
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
var go__1_11 gopurs_runtime.Value
go__1_11 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__1_11:
for {
if false { continue go__1_11 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t12 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 786377863) {
var go__4_13 gopurs_runtime.Value
go__4_13 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__4_13:
for {
if false { continue go__4_13 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t14 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 786377863) {
__t14 = v_5
goto end_branch_14
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})}
v1_6_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_6.UnsafePtr).V1
continue go__4_13
__t14 = gopurs_runtime.Value{}
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
return __t14
}
}()
})
})
__t12 = gopurs_runtime.Apply2(go__4_13, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, v_2)
goto end_branch_12
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437) {
var __t15 gopurs_runtime.Value
{
var __t_tag_16 gopurs_runtime.Value = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0
if (__t_tag_16.Type == 9 && __t_tag_16.IntVal == 786377863) {
v_2_loop = v_2
v1_3_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V1
continue go__1_11
__t15 = gopurs_runtime.Value{}
goto end_branch_15
} else {

}
}
{
var __t_tag_17 gopurs_runtime.Value = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0
if (__t_tag_17.Type == 9 && __t_tag_17.IntVal == 1358893437) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V0.UnsafePtr).V1, v_2})}
v1_3_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_3.UnsafePtr).V1
continue go__1_11
__t15 = gopurs_runtime.Value{}
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
__t12 = __t15
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
}
}()
})
})
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.Apply2(go__1_4, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V1)})}, gopurs_runtime.Apply(Get_transpose(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V1, gopurs_runtime.Apply2(go__1_11, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V1)})})})}
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}
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

func Call_groupAllBy(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
__local_var_1_0 := gopurs_runtime.Apply(Get_groupBy(), gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Apply2(p_0, x_1, y_2), gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil})
}))
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(Get_sortBy(), p_0)
_ = __local_var_2_1
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply(__local_var_2_1, x_3))
})
}

func Call_group(dictEq_0_loop *Record_eq_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEq_0 *Record_eq_gopurs_runtime_Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_groupBy(), dictEq_0.eq)
}

func Call_groupAll(dictOrd_0_loop *Record_compare_gopurs_runtime_Value) gopurs_runtime.Value {
var dictOrd_0 *Record_compare_gopurs_runtime_Value = dictOrd_0_loop
_ = dictOrd_0
__local_var_1_0 := gopurs_runtime.Apply(Get_group(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictOrd_0)}, "Eq0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = __local_var_1_0
compare_2_1 := dictOrd_0.compare
_ = compare_2_1
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply2(Get_sortBy(), compare_2_1, x_3))
})
}

func Call_fromFoldable(dictFoldable_0_loop *Record_foldMap_gopurs_runtime_Value_foldl_gopurs_runtime_Value_foldr_gopurs_runtime_Value) gopurs_runtime.Value {
var dictFoldable_0 *Record_foldMap_gopurs_runtime_Value_foldl_gopurs_runtime_Value_foldr_gopurs_runtime_Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply2(dictFoldable_0.foldr, pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil})
}

func Call_foldM(dictMonad_0_loop *Record_, v_1_loop gopurs_runtime.Value, v1_2_loop gopurs_runtime.Value, v2_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
foldM:
for {
if false { continue foldM }
var dictMonad_0 *Record_ = dictMonad_0_loop
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
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{}), "pure"), v1_2)
goto end_branch_0
} else {

}
}
{
if (v2_3.Type == 9 && v2_3.IntVal == 1358893437) {
__local_var_4_1 := (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_3.UnsafePtr).V1
_ = __local_var_4_1
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Bind1_NOT_FOUND"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(v_1, v1_2, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_3.UnsafePtr).V0), gopurs_runtime.Func(func(b_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_findIndex(fn_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
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
}

func Call_findLastIndex(fn_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
var v1_4_loop gopurs_runtime.Value = v1_4_loop_val
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
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
var go__3_3 gopurs_runtime.Value
go__3_3 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__3_3:
for {
if false { continue go__3_3 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t4 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 786377863) {
__t4 = v_4
goto end_branch_4
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})}
v1_5_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_5.UnsafePtr).V1
continue go__3_3
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(((gopurs_runtime.Apply(Get_length(), xs_1).IntVal) - (1)) - (v_2.IntVal))
}), gopurs_runtime.Apply2(go__2_0, gopurs_runtime.Int(0), gopurs_runtime.Apply2(go__3_3, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, xs_1)))
}

func Call_filterM(dictMonad_0_loop *Record_) gopurs_runtime.Value {
filterM:
for {
if false { continue filterM }
var dictMonad_0 *Record_ = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Bind1_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 786377863) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil})
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), gopurs_runtime.Apply3(Get_filterM(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictMonad_0)}, v_3, __local_var_6_4), gopurs_runtime.Func(func(xs_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
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
}

func Call_filter(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
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
go__4_2 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
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
__t1 = gopurs_runtime.Apply2(go__4_2, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, v_2)
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
return gopurs_runtime.Apply(go__1_0, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil})
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
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}
goto end_branch_5
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 786377863) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}
goto end_branch_5
} else {

}
}
{
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
if (v1_5.Type == 9 && v1_5.IntVal == 786377863) {
var go__6_2 gopurs_runtime.Value
go__6_2 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_7_loop gopurs_runtime.Value = v_7_loop_val
var v1_8_loop gopurs_runtime.Value = v1_8_loop_val
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
__t1 = gopurs_runtime.Apply2(go__6_2, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, v_4)
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
__t5 = gopurs_runtime.Apply2(go__3_0, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, v1_1)
}
end_branch_5:
return __t5
}

func Call_intersect(dictEq_0_loop *Record_eq_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEq_0 *Record_eq_gopurs_runtime_Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_intersectBy(), dictEq_0.eq)
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 1358893437) {
__local_var_2_1 := (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_1.UnsafePtr).V0
_ = __local_var_2_1
var go__3_2 gopurs_runtime.Value
go__3_2 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
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
go__6_4 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_7_loop gopurs_runtime.Value = v_7_loop_val
var v1_8_loop gopurs_runtime.Value = v1_8_loop_val
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
__t3 = gopurs_runtime.Apply2(go__6_4, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, v_4)
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437) {
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Apply2(v_0, __local_var_2_1, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_5.UnsafePtr).V0)).IntVal) != (0) {
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{__local_var_2_1, Call_nubByEq(v_0, gopurs_runtime.Apply2(go__3_2, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_1.UnsafePtr).V1))})}
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

func Call_nubEq(dictEq_0_loop *Record_eq_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEq_0 *Record_eq_gopurs_runtime_Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_nubByEq(), dictEq_0.eq)
}

func Call_eqPattern(dictEq_0_loop *Record_eq_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEq_0 *Record_eq_gopurs_runtime_Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_eq1List(), "eq1"), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictEq_0)}, x_1, y_2)
}))
}

func Call_ordPattern(dictOrd_0_loop *Record_compare_gopurs_runtime_Value) gopurs_runtime.Value {
var dictOrd_0 *Record_compare_gopurs_runtime_Value = dictOrd_0_loop
_ = dictOrd_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictOrd_0)}, "Eq0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
eqPattern1_2_1 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_eq1List(), "eq1"), __local_var_1_0, x_2, y_3)
}))
_ = eqPattern1_2_1
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eqPattern1_2_1
}), gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(pkg_Data_List_Types.Get_ordList(), gopurs_runtime.Value{Type: 8, IntVal: 0, UnsafePtr: unsafe.Pointer(dictOrd_0)}), "compare"), x_3, y_4)
}))
}

func Call_elemLastIndex(dictEq_0_loop *Record_eq_gopurs_runtime_Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 *Record_eq_gopurs_runtime_Value = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_findLastIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictEq_0.eq, v_2, x_1)
}))
}

func Call_elemIndex(dictEq_0_loop *Record_eq_gopurs_runtime_Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 *Record_eq_gopurs_runtime_Value = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var go__2_0 gopurs_runtime.Value
go__2_0 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
var v1_4_loop gopurs_runtime.Value = v1_4_loop_val
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
if (gopurs_runtime.Apply2(dictEq_0.eq, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_4.UnsafePtr).V0, x_1).IntVal) != (0) {
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
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

func Call_dropWhile(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
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
}

func Call_dropEnd(n_0_loop int64, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Int((gopurs_runtime.Apply(Get_length(), xs_1).IntVal) - (n_0)), xs_1)
}

func Call_drop(v_0_loop int64, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
drop:
for {
if false { continue drop }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThan(), gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal) != (0) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 786377863) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (v1_1.Type == 9 && v1_1.IntVal == 1358893437) {
v_0_loop = (v_0) - (1)
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

func Call_slice(start_0_loop int64, end_1_loop int64, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var start_0 int64 = start_0_loop
_ = start_0
var end_1 int64 = end_1_loop
_ = end_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Int((end_1) - (start_0)), Call_drop(start_0, xs_2))
}

func Call_takeEnd(n_0_loop int64, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return Call_drop((gopurs_runtime.Apply(Get_length(), xs_1).IntVal) - (n_0), xs_1)
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: nil}
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_semigroupList(), "append"), xs_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_deleteBy(eq2_0, a_4, b_3)
}), Call_nubByEq(eq2_0, ys_2), xs_1))
}

func Call_union(dictEq_0_loop *Record_eq_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEq_0 *Record_eq_gopurs_runtime_Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_unionBy(), dictEq_0.eq)
}

func Call_deleteAt(v_0_loop int64, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
deleteAt:
for {
if false { continue deleteAt }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 1358893437) {
var __t2 gopurs_runtime.Value
{
if (v_0) == (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_1.UnsafePtr).V1})}
goto end_branch_2
} else {

}
}
{
__local_var_2_1 := (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_1.UnsafePtr).V0
_ = __local_var_2_1
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{__local_var_2_1, v2_3})}
}), Call_deleteAt((v_0) - (1), (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v1_1.UnsafePtr).V1))
}
end_branch_2:
__t0 = __t2
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
}
end_branch_0:
return __t0
}
}

func Call_delete_(dictEq_0_loop *Record_eq_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEq_0 *Record_eq_gopurs_runtime_Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_deleteBy(), dictEq_0.eq)
}

func Call_difference(dictEq_0_loop *Record_eq_gopurs_runtime_Value) gopurs_runtime.Value {
var dictEq_0 *Record_eq_gopurs_runtime_Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func2(func(b_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_deleteBy(dictEq_0.eq, a_2, b_1)
}))
}

func Call_concatMap(b_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_bindList(), "bind"), a_1, b_0)
}

func Call_concat(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_bindList(), "bind"), v_0, Get_identity())
}

func Call_alterAt(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
alterAt:
for {
if false { continue alterAt }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 1358893437) {
var __t2 gopurs_runtime.Value
{
if (v_0) == (0) {
v3_3_3 := gopurs_runtime.Apply(v1_1, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V0)
_ = v3_3_3
var __t4 gopurs_runtime.Value
{
if (v3_3_3.Type == 9 && v3_3_3.IntVal == 3589588149) {
__t4 = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V1
goto end_branch_4
} else {

}
}
{
if (v3_3_3.Type == 9 && v3_3_3.IntVal == 930809136) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{(*pkg_Data_Maybe.Data_Data_Maybe_Just)(v3_3_3.UnsafePtr).V0, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V1})}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{__t4})}
goto end_branch_2
} else {

}
}
{
__local_var_3_1 := (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V0
_ = __local_var_3_1
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v3_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{__local_var_3_1, v3_4})}
}), Call_alterAt((v_0) - (1), v1_1, (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v2_2.UnsafePtr).V1))
}
end_branch_2:
__t0 = __t2
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
}
end_branch_0:
return __t0
}
}

func Call_modifyAt(n_0_loop int64, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(Get_alterAt(), gopurs_runtime.Int(n_0), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Apply(f_1, x_2)})}
}))
}



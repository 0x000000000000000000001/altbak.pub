package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_List_identity gopurs_runtime.Value
var once_Data_List_identity sync.Once
func Get_Data_List_identity() gopurs_runtime.Value {
	once_Data_List_identity.Do(func() {
		cache_Data_List_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_identity(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](x_0_box)))}
})
	})
	return cache_Data_List_identity
}

var cache_Data_List_Pattern gopurs_runtime.Value
var once_Data_List_Pattern sync.Once
func Get_Data_List_Pattern() gopurs_runtime.Value {
	once_Data_List_Pattern.Do(func() {
		cache_Data_List_Pattern = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Pattern(x_0_box)
})
	})
	return cache_Data_List_Pattern
}

var cache_Data_List_updateAt gopurs_runtime.Value
var once_Data_List_updateAt sync.Once
func Get_Data_List_updateAt() gopurs_runtime.Value {
	once_Data_List_updateAt.Do(func() {
		cache_Data_List_updateAt = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_updateAt(v_0_box.IntVal, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v2_2_box)))}
})
	})
	return cache_Data_List_updateAt
}

var cache_Data_List_unzip gopurs_runtime.Value
var once_Data_List_unzip sync.Once
func Get_Data_List_unzip() gopurs_runtime.Value {
	once_Data_List_unzip.Do(func() {
		cache_Data_List_unzip = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldr"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := (*Constructor_Data_Tuple_Tuple)(v_0.UnsafePtr).V0
_ = __local_var_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := (*Constructor_Data_Tuple_Tuple)(v_0.UnsafePtr).V1
_ = __local_var_2_1
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, __local_var_1_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_Tuple_Tuple)(v1_3.UnsafePtr).V0)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, __local_var_2_1, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_Tuple_Tuple)(v1_3.UnsafePtr).V1)})}})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}})})
	})
	return cache_Data_List_unzip
}

var cache_Data_List_uncons gopurs_runtime.Value
var once_Data_List_uncons sync.Once
func Get_Data_List_uncons() gopurs_runtime.Value {
	once_Data_List_uncons.Do(func() {
		cache_Data_List_uncons = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_uncons(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box)))}
})
	})
	return cache_Data_List_uncons
}

var cache_Data_List_toUnfoldable gopurs_runtime.Value
var once_Data_List_toUnfoldable sync.Once
func Get_Data_List_toUnfoldable() gopurs_runtime.Value {
	once_Data_List_toUnfoldable.Do(func() {
		cache_Data_List_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dictUnfoldable_0_box))
})
	})
	return cache_Data_List_toUnfoldable
}

var cache_Data_List_tail gopurs_runtime.Value
var once_Data_List_tail sync.Once
func Get_Data_List_tail() gopurs_runtime.Value {
	once_Data_List_tail.Do(func() {
		cache_Data_List_tail = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_tail(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box)))}
})
	})
	return cache_Data_List_tail
}

var cache_Data_List_stripPrefix gopurs_runtime.Value
var once_Data_List_stripPrefix sync.Once
func Get_Data_List_stripPrefix() gopurs_runtime.Value {
	once_Data_List_stripPrefix.Do(func() {
		cache_Data_List_stripPrefix = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_stripPrefix(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](s_2_box)))}
})
	})
	return cache_Data_List_stripPrefix
}

var cache_Data_List_span gopurs_runtime.Value
var once_Data_List_span sync.Once
func Get_Data_List_span() gopurs_runtime.Value {
	once_Data_List_span.Do(func() {
		cache_Data_List_span = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_span(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_1_box))
})
	})
	return cache_Data_List_span
}

var cache_Data_List_snoc gopurs_runtime.Value
var once_Data_List_snoc sync.Once
func Get_Data_List_snoc() gopurs_runtime.Value {
	once_Data_List_snoc.Do(func() {
		cache_Data_List_snoc = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_snoc(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_0_box), x_1_box))}
})
	})
	return cache_Data_List_snoc
}

var cache_Data_List_singleton gopurs_runtime.Value
var once_Data_List_singleton sync.Once
func Get_Data_List_singleton() gopurs_runtime.Value {
	once_Data_List_singleton.Do(func() {
		cache_Data_List_singleton = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_singleton(a_0_box))}
})
	})
	return cache_Data_List_singleton
}

var cache_Data_List_sortBy gopurs_runtime.Value
var once_Data_List_sortBy sync.Once
func Get_Data_List_sortBy() gopurs_runtime.Value {
	once_Data_List_sortBy.Do(func() {
		cache_Data_List_sortBy = gopurs_runtime.Func(func(cmp_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_sortBy(cmp_0_box)
})
	})
	return cache_Data_List_sortBy
}

var cache_Data_List_sort gopurs_runtime.Value
var once_Data_List_sort sync.Once
func Get_Data_List_sort() gopurs_runtime.Value {
	once_Data_List_sort.Do(func() {
		cache_Data_List_sort = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_sort(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_List_sort
}

var cache_Data_List_tails gopurs_runtime.Value
var once_Data_List_tails sync.Once
func Get_Data_List_tails() gopurs_runtime.Value {
	once_Data_List_tails.Do(func() {
		cache_Data_List_tails = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_tails(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box)))}
})
	})
	return cache_Data_List_tails
}

var cache_Data_List_showPattern gopurs_runtime.Value
var once_Data_List_showPattern sync.Once
func Get_Data_List_showPattern() gopurs_runtime.Value {
	once_Data_List_showPattern.Do(func() {
		cache_Data_List_showPattern = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_showPattern(dictShow_0_box)
})
	})
	return cache_Data_List_showPattern
}

var cache_Data_List_reverse gopurs_runtime.Value
var once_Data_List_reverse sync.Once
func Get_Data_List_reverse() gopurs_runtime.Value {
	once_Data_List_reverse.Do(func() {
		cache_Data_List_reverse = func() gopurs_runtime.Value {
var go__go_0_0_6 gopurs_runtime.Value
go__go_0_0_6 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_1_loop_val)
var v1_2_loop gopurs_runtime.Value = v1_2_loop_val
go__go_0_0_6:
for {
if false { continue go__go_0_0_6 }
var v_1 *Constructor_Data_List_Types_Cons = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t1 *Constructor_Data_List_Types_Cons
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr == nil) {
__t1 = v_1
goto end_branch_1
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr != nil) {
v_1_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_2.UnsafePtr).V0, v_1})})
v1_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_2.UnsafePtr).V1)}
continue go__go_0_0_6
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
})
return gopurs_runtime.Apply(go__go_0_0_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}()
	})
	return cache_Data_List_reverse
}

var cache_Data_List_take gopurs_runtime.Value
var once_Data_List_take sync.Once
func Get_Data_List_take() gopurs_runtime.Value {
	once_Data_List_take.Do(func() {
		cache_Data_List_take = func() gopurs_runtime.Value {
var go__go_0_0_7 gopurs_runtime.Value
go__go_0_0_7 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_1_loop_val)
var v1_2_loop int64 = v1_2_loop_val.IntVal
var v2_3_loop gopurs_runtime.Value = v2_3_loop_val
go__go_0_0_7:
for {
if false { continue go__go_0_0_7 }
var v_1 *Constructor_Data_List_Types_Cons = v_1_loop
_ = v_1
var v1_2 int64 = v1_2_loop
_ = v1_2
var v2_3 gopurs_runtime.Value = v2_3_loop
_ = v2_3
var __t6 *Constructor_Data_List_Types_Cons
{
var __t1 bool
{
if (v1_2) < (1) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
if __t1 {
var go__go_4_2_8 gopurs_runtime.Value
go__go_4_2_8 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_2_8:
for {
if false { continue go__go_4_2_8 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t3 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_2_8
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_2_8, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}))
goto end_branch_6
} else {

}
}
{
if (v2_3.Type == 9 && v2_3.IntVal == 1358893437 && v2_3.UnsafePtr == nil) {
var go__go_4_4_9 gopurs_runtime.Value
go__go_4_4_9 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_4_9:
for {
if false { continue go__go_4_4_9 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t5 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t5 = v_5
goto end_branch_5
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_4_9
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t5)}
}
}()
})
})
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_4_9, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}))
goto end_branch_6
} else {

}
}
{
if (v2_3.Type == 9 && v2_3.IntVal == 1358893437 && v2_3.UnsafePtr != nil) {
v_1_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v2_3.UnsafePtr).V0, v_1})})
v1_2_loop = (v1_2) - (1)
v2_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v2_3.UnsafePtr).V1)}
continue go__go_0_0_7
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t6)}
}
}()
})
})
})
return gopurs_runtime.Apply(go__go_0_0_7, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}()
	})
	return cache_Data_List_take
}

var cache_Data_List_takeWhile gopurs_runtime.Value
var once_Data_List_takeWhile sync.Once
func Get_Data_List_takeWhile() gopurs_runtime.Value {
	once_Data_List_takeWhile.Do(func() {
		cache_Data_List_takeWhile = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_takeWhile(p_0_box)
})
	})
	return cache_Data_List_takeWhile
}

var cache_Data_List_unsnoc gopurs_runtime.Value
var once_Data_List_unsnoc sync.Once
func Get_Data_List_unsnoc() gopurs_runtime.Value {
	once_Data_List_unsnoc.Do(func() {
		cache_Data_List_unsnoc = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_unsnoc(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](lst_0_box)))}
})
	})
	return cache_Data_List_unsnoc
}

var cache_Data_List_zipWith gopurs_runtime.Value
var once_Data_List_zipWith sync.Once
func Get_Data_List_zipWith() gopurs_runtime.Value {
	once_Data_List_zipWith.Do(func() {
		cache_Data_List_zipWith = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_zipWith(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_2_box)))}
})
	})
	return cache_Data_List_zipWith
}

var cache_Data_List_zip gopurs_runtime.Value
var once_Data_List_zip sync.Once
func Get_Data_List_zip() gopurs_runtime.Value {
	once_Data_List_zip.Do(func() {
		cache_Data_List_zip = gopurs_runtime.Apply(Get_Data_List_zipWith(), Get_Data_Tuple_Tuple())
	})
	return cache_Data_List_zip
}

var cache_Data_List_zipWithA gopurs_runtime.Value
var once_Data_List_zipWithA sync.Once
func Get_Data_List_zipWithA() gopurs_runtime.Value {
	once_Data_List_zipWithA.Do(func() {
		cache_Data_List_zipWithA = gopurs_runtime.Func4(func(dictApplicative_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value, ys_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_zipWithA(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box), f_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_2_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_3_box))
})
	})
	return cache_Data_List_zipWithA
}

var cache_Data_List_go__range gopurs_runtime.Value
var once_Data_List_go__range sync.Once
func Get_Data_List_go__range() gopurs_runtime.Value {
	once_Data_List_go__range.Do(func() {
		cache_Data_List_go__range = gopurs_runtime.Func2(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_go__range(start_0_box.IntVal, end_1_box.IntVal))}
})
	})
	return cache_Data_List_go__range
}

var cache_Data_List_partition gopurs_runtime.Value
var once_Data_List_partition sync.Once
func Get_Data_List_partition() gopurs_runtime.Value {
	once_Data_List_partition.Do(func() {
		cache_Data_List_partition = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_partition(p_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_1_box))
})
	})
	return cache_Data_List_partition
}

var cache_Data_List_null gopurs_runtime.Value
var once_Data_List_null sync.Once
func Get_Data_List_null() gopurs_runtime.Value {
	once_Data_List_null.Do(func() {
		cache_Data_List_null = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_List_null(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box)))
})
	})
	return cache_Data_List_null
}

var cache_Data_List_nubBy gopurs_runtime.Value
var once_Data_List_nubBy sync.Once
func Get_Data_List_nubBy() gopurs_runtime.Value {
	once_Data_List_nubBy.Do(func() {
		cache_Data_List_nubBy = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_nubBy(p_0_box)
})
	})
	return cache_Data_List_nubBy
}

var cache_Data_List_nub gopurs_runtime.Value
var once_Data_List_nub sync.Once
func Get_Data_List_nub() gopurs_runtime.Value {
	once_Data_List_nub.Do(func() {
		cache_Data_List_nub = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_nub(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_List_nub
}

var cache_Data_List_newtypePattern gopurs_runtime.Value
var once_Data_List_newtypePattern sync.Once
func Get_Data_List_newtypePattern() gopurs_runtime.Value {
	once_Data_List_newtypePattern.Do(func() {
		cache_Data_List_newtypePattern = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_List_newtypePattern
}

var cache_Data_List_mapMaybe gopurs_runtime.Value
var once_Data_List_mapMaybe sync.Once
func Get_Data_List_mapMaybe() gopurs_runtime.Value {
	once_Data_List_mapMaybe.Do(func() {
		cache_Data_List_mapMaybe = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_mapMaybe(f_0_box)
})
	})
	return cache_Data_List_mapMaybe
}

var cache_Data_List_manyRec gopurs_runtime.Value
var once_Data_List_manyRec sync.Once
func Get_Data_List_manyRec() gopurs_runtime.Value {
	once_Data_List_manyRec.Do(func() {
		cache_Data_List_manyRec = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_manyRec(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box))
})
	})
	return cache_Data_List_manyRec
}

var cache_Data_List_someRec gopurs_runtime.Value
var once_Data_List_someRec sync.Once
func Get_Data_List_someRec() gopurs_runtime.Value {
	once_Data_List_someRec.Do(func() {
		cache_Data_List_someRec = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictAlternative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_someRec(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative](dictAlternative_1_box))
})
	})
	return cache_Data_List_someRec
}

var cache_Data_List_some gopurs_runtime.Value
var once_Data_List_some sync.Once
func Get_Data_List_some() gopurs_runtime.Value {
	once_Data_List_some.Do(func() {
		cache_Data_List_some = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_some(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative](dictAlternative_0_box))
})
	})
	return cache_Data_List_some
}

var cache_Data_List_many gopurs_runtime.Value
var once_Data_List_many sync.Once
func Get_Data_List_many() gopurs_runtime.Value {
	once_Data_List_many.Do(func() {
		cache_Data_List_many = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_many(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative](dictAlternative_0_box))
})
	})
	return cache_Data_List_many
}

var cache_Data_List_length gopurs_runtime.Value
var once_Data_List_length sync.Once
func Get_Data_List_length() gopurs_runtime.Value {
	once_Data_List_length.Do(func() {
		cache_Data_List_length = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((acc_0.IntVal) + (1))
})
}), gopurs_runtime.Int(0))
	})
	return cache_Data_List_length
}

var cache_Data_List_last gopurs_runtime.Value
var once_Data_List_last sync.Once
func Get_Data_List_last() gopurs_runtime.Value {
	once_Data_List_last.Do(func() {
		cache_Data_List_last = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_last(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box)))}
})
	})
	return cache_Data_List_last
}

var cache_Data_List_insertBy gopurs_runtime.Value
var once_Data_List_insertBy sync.Once
func Get_Data_List_insertBy() gopurs_runtime.Value {
	once_Data_List_insertBy.Do(func() {
		cache_Data_List_insertBy = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_insertBy(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v2_2_box)))}
})
	})
	return cache_Data_List_insertBy
}

var cache_Data_List_insertAt gopurs_runtime.Value
var once_Data_List_insertAt sync.Once
func Get_Data_List_insertAt() gopurs_runtime.Value {
	once_Data_List_insertAt.Do(func() {
		cache_Data_List_insertAt = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_insertAt(v_0_box.IntVal, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v2_2_box)))}
})
	})
	return cache_Data_List_insertAt
}

var cache_Data_List_insert gopurs_runtime.Value
var once_Data_List_insert sync.Once
func Get_Data_List_insert() gopurs_runtime.Value {
	once_Data_List_insert.Do(func() {
		cache_Data_List_insert = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_insert(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_List_insert
}

var cache_Data_List_init gopurs_runtime.Value
var once_Data_List_init sync.Once
func Get_Data_List_init() gopurs_runtime.Value {
	once_Data_List_init.Do(func() {
		cache_Data_List_init = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_init(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](lst_0_box)))}
})
	})
	return cache_Data_List_init
}

var cache_Data_List_index gopurs_runtime.Value
var once_Data_List_index sync.Once
func Get_Data_List_index() gopurs_runtime.Value {
	once_Data_List_index.Do(func() {
		cache_Data_List_index = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_index(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box), v1_1_box.IntVal))}
})
	})
	return cache_Data_List_index
}

var cache_Data_List_head gopurs_runtime.Value
var once_Data_List_head sync.Once
func Get_Data_List_head() gopurs_runtime.Value {
	once_Data_List_head.Do(func() {
		cache_Data_List_head = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_head(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box)))}
})
	})
	return cache_Data_List_head
}

var cache_Data_List_transpose gopurs_runtime.Value
var once_Data_List_transpose sync.Once
func Get_Data_List_transpose() gopurs_runtime.Value {
	once_Data_List_transpose.Do(func() {
		cache_Data_List_transpose = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_transpose(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box)))}
})
	})
	return cache_Data_List_transpose
}

var cache_Data_List_groupBy gopurs_runtime.Value
var once_Data_List_groupBy sync.Once
func Get_Data_List_groupBy() gopurs_runtime.Value {
	once_Data_List_groupBy.Do(func() {
		cache_Data_List_groupBy = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_groupBy(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_1_box)))}
})
	})
	return cache_Data_List_groupBy
}

var cache_Data_List_groupAllBy gopurs_runtime.Value
var once_Data_List_groupAllBy sync.Once
func Get_Data_List_groupAllBy() gopurs_runtime.Value {
	once_Data_List_groupAllBy.Do(func() {
		cache_Data_List_groupAllBy = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_groupAllBy(p_0_box)
})
	})
	return cache_Data_List_groupAllBy
}

var cache_Data_List_group gopurs_runtime.Value
var once_Data_List_group sync.Once
func Get_Data_List_group() gopurs_runtime.Value {
	once_Data_List_group.Do(func() {
		cache_Data_List_group = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_group(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_List_group
}

var cache_Data_List_groupAll gopurs_runtime.Value
var once_Data_List_groupAll sync.Once
func Get_Data_List_groupAll() gopurs_runtime.Value {
	once_Data_List_groupAll.Do(func() {
		cache_Data_List_groupAll = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_groupAll(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_List_groupAll
}

var cache_Data_List_fromFoldable gopurs_runtime.Value
var once_Data_List_fromFoldable sync.Once
func Get_Data_List_fromFoldable() gopurs_runtime.Value {
	once_Data_List_fromFoldable.Do(func() {
		cache_Data_List_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box))
})
	})
	return cache_Data_List_fromFoldable
}

var cache_Data_List_foldM gopurs_runtime.Value
var once_Data_List_foldM sync.Once
func Get_Data_List_foldM() gopurs_runtime.Value {
	once_Data_List_foldM.Do(func() {
		cache_Data_List_foldM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_foldM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Data_List_foldM
}

var cache_Data_List_findIndex gopurs_runtime.Value
var once_Data_List_findIndex sync.Once
func Get_Data_List_findIndex() gopurs_runtime.Value {
	once_Data_List_findIndex.Do(func() {
		cache_Data_List_findIndex = gopurs_runtime.Func(func(fn_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_findIndex(fn_0_box)
})
	})
	return cache_Data_List_findIndex
}

var cache_Data_List_findLastIndex gopurs_runtime.Value
var once_Data_List_findLastIndex sync.Once
func Get_Data_List_findLastIndex() gopurs_runtime.Value {
	once_Data_List_findLastIndex.Do(func() {
		cache_Data_List_findLastIndex = gopurs_runtime.Func2(func(fn_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_findLastIndex(fn_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_1_box)))}
})
	})
	return cache_Data_List_findLastIndex
}

var cache_Data_List_filterM gopurs_runtime.Value
var once_Data_List_filterM sync.Once
func Get_Data_List_filterM() gopurs_runtime.Value {
	once_Data_List_filterM.Do(func() {
		cache_Data_List_filterM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_filterM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Data_List_filterM
}

var cache_Data_List_filter gopurs_runtime.Value
var once_Data_List_filter sync.Once
func Get_Data_List_filter() gopurs_runtime.Value {
	once_Data_List_filter.Do(func() {
		cache_Data_List_filter = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_filter(p_0_box)
})
	})
	return cache_Data_List_filter
}

var cache_Data_List_intersectBy gopurs_runtime.Value
var once_Data_List_intersectBy sync.Once
func Get_Data_List_intersectBy() gopurs_runtime.Value {
	once_Data_List_intersectBy.Do(func() {
		cache_Data_List_intersectBy = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_intersectBy(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v2_2_box)))}
})
	})
	return cache_Data_List_intersectBy
}

var cache_Data_List_intersect gopurs_runtime.Value
var once_Data_List_intersect sync.Once
func Get_Data_List_intersect() gopurs_runtime.Value {
	once_Data_List_intersect.Do(func() {
		cache_Data_List_intersect = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_intersect(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_List_intersect
}

var cache_Data_List_nubByEq gopurs_runtime.Value
var once_Data_List_nubByEq sync.Once
func Get_Data_List_nubByEq() gopurs_runtime.Value {
	once_Data_List_nubByEq.Do(func() {
		cache_Data_List_nubByEq = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_nubByEq(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_1_box)))}
})
	})
	return cache_Data_List_nubByEq
}

var cache_Data_List_nubEq gopurs_runtime.Value
var once_Data_List_nubEq sync.Once
func Get_Data_List_nubEq() gopurs_runtime.Value {
	once_Data_List_nubEq.Do(func() {
		cache_Data_List_nubEq = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_nubEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_List_nubEq
}

var cache_Data_List_eqPattern gopurs_runtime.Value
var once_Data_List_eqPattern sync.Once
func Get_Data_List_eqPattern() gopurs_runtime.Value {
	once_Data_List_eqPattern.Do(func() {
		cache_Data_List_eqPattern = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_eqPattern(dictEq_0_box)
})
	})
	return cache_Data_List_eqPattern
}

var cache_Data_List_ordPattern gopurs_runtime.Value
var once_Data_List_ordPattern sync.Once
func Get_Data_List_ordPattern() gopurs_runtime.Value {
	once_Data_List_ordPattern.Do(func() {
		cache_Data_List_ordPattern = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_ordPattern(dictOrd_0_box)
})
	})
	return cache_Data_List_ordPattern
}

var cache_Data_List_elemLastIndex gopurs_runtime.Value
var once_Data_List_elemLastIndex sync.Once
func Get_Data_List_elemLastIndex() gopurs_runtime.Value {
	once_Data_List_elemLastIndex.Do(func() {
		cache_Data_List_elemLastIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_elemLastIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), x_1_box)
})
	})
	return cache_Data_List_elemLastIndex
}

var cache_Data_List_elemIndex gopurs_runtime.Value
var once_Data_List_elemIndex sync.Once
func Get_Data_List_elemIndex() gopurs_runtime.Value {
	once_Data_List_elemIndex.Do(func() {
		cache_Data_List_elemIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_elemIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), x_1_box)
})
	})
	return cache_Data_List_elemIndex
}

var cache_Data_List_dropWhile gopurs_runtime.Value
var once_Data_List_dropWhile sync.Once
func Get_Data_List_dropWhile() gopurs_runtime.Value {
	once_Data_List_dropWhile.Do(func() {
		cache_Data_List_dropWhile = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_dropWhile(p_0_box)
})
	})
	return cache_Data_List_dropWhile
}

var cache_Data_List_dropEnd gopurs_runtime.Value
var once_Data_List_dropEnd sync.Once
func Get_Data_List_dropEnd() gopurs_runtime.Value {
	once_Data_List_dropEnd.Do(func() {
		cache_Data_List_dropEnd = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_dropEnd(n_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_1_box)))}
})
	})
	return cache_Data_List_dropEnd
}

var cache_Data_List_drop gopurs_runtime.Value
var once_Data_List_drop sync.Once
func Get_Data_List_drop() gopurs_runtime.Value {
	once_Data_List_drop.Do(func() {
		cache_Data_List_drop = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_drop(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_1_box)))}
})
	})
	return cache_Data_List_drop
}

var cache_Data_List_slice gopurs_runtime.Value
var once_Data_List_slice sync.Once
func Get_Data_List_slice() gopurs_runtime.Value {
	once_Data_List_slice.Do(func() {
		cache_Data_List_slice = gopurs_runtime.Func3(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_slice(start_0_box.IntVal, end_1_box.IntVal, xs_2_box)
})
	})
	return cache_Data_List_slice
}

var cache_Data_List_takeEnd gopurs_runtime.Value
var once_Data_List_takeEnd sync.Once
func Get_Data_List_takeEnd() gopurs_runtime.Value {
	once_Data_List_takeEnd.Do(func() {
		cache_Data_List_takeEnd = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_takeEnd(n_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_1_box)))}
})
	})
	return cache_Data_List_takeEnd
}

var cache_Data_List_deleteBy gopurs_runtime.Value
var once_Data_List_deleteBy sync.Once
func Get_Data_List_deleteBy() gopurs_runtime.Value {
	once_Data_List_deleteBy.Do(func() {
		cache_Data_List_deleteBy = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_deleteBy(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v2_2_box)))}
})
	})
	return cache_Data_List_deleteBy
}

var cache_Data_List_unionBy gopurs_runtime.Value
var once_Data_List_unionBy sync.Once
func Get_Data_List_unionBy() gopurs_runtime.Value {
	once_Data_List_unionBy.Do(func() {
		cache_Data_List_unionBy = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_unionBy(eq_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_2_box)))}
})
	})
	return cache_Data_List_unionBy
}

var cache_Data_List_union gopurs_runtime.Value
var once_Data_List_union sync.Once
func Get_Data_List_union() gopurs_runtime.Value {
	once_Data_List_union.Do(func() {
		cache_Data_List_union = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_union(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_List_union
}

var cache_Data_List_deleteAt gopurs_runtime.Value
var once_Data_List_deleteAt sync.Once
func Get_Data_List_deleteAt() gopurs_runtime.Value {
	once_Data_List_deleteAt.Do(func() {
		cache_Data_List_deleteAt = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_deleteAt(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_1_box)))}
})
	})
	return cache_Data_List_deleteAt
}

var cache_Data_List_delete gopurs_runtime.Value
var once_Data_List_delete sync.Once
func Get_Data_List_delete() gopurs_runtime.Value {
	once_Data_List_delete.Do(func() {
		cache_Data_List_delete = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_delete(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_List_delete
}

var cache_Data_List_difference gopurs_runtime.Value
var once_Data_List_difference sync.Once
func Get_Data_List_difference() gopurs_runtime.Value {
	once_Data_List_difference.Do(func() {
		cache_Data_List_difference = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_difference(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_List_difference
}

var cache_Data_List_concatMap gopurs_runtime.Value
var once_Data_List_concatMap sync.Once
func Get_Data_List_concatMap() gopurs_runtime.Value {
	once_Data_List_concatMap.Do(func() {
		cache_Data_List_concatMap = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_concatMap(b_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](a_1_box)))}
})
	})
	return cache_Data_List_concatMap
}

var cache_Data_List_concat gopurs_runtime.Value
var once_Data_List_concat sync.Once
func Get_Data_List_concat() gopurs_runtime.Value {
	once_Data_List_concat.Do(func() {
		cache_Data_List_concat = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_concat(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box)))}
})
	})
	return cache_Data_List_concat
}

var cache_Data_List_catMaybes gopurs_runtime.Value
var once_Data_List_catMaybes sync.Once
func Get_Data_List_catMaybes() gopurs_runtime.Value {
	once_Data_List_catMaybes.Do(func() {
		cache_Data_List_catMaybes = Call_Data_List_mapMaybe(gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_Data_List_catMaybes
}

var cache_Data_List_alterAt gopurs_runtime.Value
var once_Data_List_alterAt sync.Once
func Get_Data_List_alterAt() gopurs_runtime.Value {
	once_Data_List_alterAt.Do(func() {
		cache_Data_List_alterAt = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_alterAt(v_0_box.IntVal, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v2_2_box)))}
})
	})
	return cache_Data_List_alterAt
}

var cache_Data_List_modifyAt gopurs_runtime.Value
var once_Data_List_modifyAt sync.Once
func Get_Data_List_modifyAt() gopurs_runtime.Value {
	once_Data_List_modifyAt.Do(func() {
		cache_Data_List_modifyAt = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_modifyAt(n_0_box.IntVal, f_1_box)
})
	})
	return cache_Data_List_modifyAt
}

var cache_Data_List_alterAt__3453373293 gopurs_runtime.Value
var once_Data_List_alterAt__3453373293 sync.Once
func Get_Data_List_alterAt__3453373293() gopurs_runtime.Value {
	once_Data_List_alterAt__3453373293.Do(func() {
		cache_Data_List_alterAt__3453373293 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_alterAt__3453373293(v_0_box.IntVal, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v2_2_box)))}
})
	})
	return cache_Data_List_alterAt__3453373293
}

var cache_Data_List_catMaybes__3687414234 gopurs_runtime.Value
var once_Data_List_catMaybes__3687414234 sync.Once
func Get_Data_List_catMaybes__3687414234() gopurs_runtime.Value {
	once_Data_List_catMaybes__3687414234.Do(func() {
		cache_Data_List_catMaybes__3687414234 = Call_Data_List_mapMaybe(gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_Data_List_catMaybes__3687414234
}

var cache_Data_List_deleteAt__2845095501 gopurs_runtime.Value
var once_Data_List_deleteAt__2845095501 sync.Once
func Get_Data_List_deleteAt__2845095501() gopurs_runtime.Value {
	once_Data_List_deleteAt__2845095501.Do(func() {
		cache_Data_List_deleteAt__2845095501 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_deleteAt__2845095501(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_1_box)))}
})
	})
	return cache_Data_List_deleteAt__2845095501
}

var cache_Data_List_deleteBy__697302515 gopurs_runtime.Value
var once_Data_List_deleteBy__697302515 sync.Once
func Get_Data_List_deleteBy__697302515() gopurs_runtime.Value {
	once_Data_List_deleteBy__697302515.Do(func() {
		cache_Data_List_deleteBy__697302515 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_deleteBy__697302515(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v2_2_box)))}
})
	})
	return cache_Data_List_deleteBy__697302515
}

var cache_Data_List_drop__551729751 gopurs_runtime.Value
var once_Data_List_drop__551729751 sync.Once
func Get_Data_List_drop__551729751() gopurs_runtime.Value {
	once_Data_List_drop__551729751.Do(func() {
		cache_Data_List_drop__551729751 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_drop__551729751(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_1_box)))}
})
	})
	return cache_Data_List_drop__551729751
}

var cache_Data_List_drop__1836090668 gopurs_runtime.Value
var once_Data_List_drop__1836090668 sync.Once
func Get_Data_List_drop__1836090668() gopurs_runtime.Value {
	once_Data_List_drop__1836090668.Do(func() {
		cache_Data_List_drop__1836090668 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_drop__1836090668(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_1_box)))}
})
	})
	return cache_Data_List_drop__1836090668
}

var cache_Data_List_dropWhile__2352021032 gopurs_runtime.Value
var once_Data_List_dropWhile__2352021032 sync.Once
func Get_Data_List_dropWhile__2352021032() gopurs_runtime.Value {
	once_Data_List_dropWhile__2352021032.Do(func() {
		cache_Data_List_dropWhile__2352021032 = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_dropWhile__2352021032(p_0_box)
})
	})
	return cache_Data_List_dropWhile__2352021032
}

var cache_Data_List_filter__2352021032 gopurs_runtime.Value
var once_Data_List_filter__2352021032 sync.Once
func Get_Data_List_filter__2352021032() gopurs_runtime.Value {
	once_Data_List_filter__2352021032.Do(func() {
		cache_Data_List_filter__2352021032 = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_filter__2352021032(p_0_box)
})
	})
	return cache_Data_List_filter__2352021032
}

var cache_Data_List_filter__1617261107 gopurs_runtime.Value
var once_Data_List_filter__1617261107 sync.Once
func Get_Data_List_filter__1617261107() gopurs_runtime.Value {
	once_Data_List_filter__1617261107.Do(func() {
		cache_Data_List_filter__1617261107 = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_filter__1617261107(p_0_box)
})
	})
	return cache_Data_List_filter__1617261107
}

var cache_Data_List_filterM__14771079 gopurs_runtime.Value
var once_Data_List_filterM__14771079 sync.Once
func Get_Data_List_filterM__14771079() gopurs_runtime.Value {
	once_Data_List_filterM__14771079.Do(func() {
		cache_Data_List_filterM__14771079 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_filterM__14771079(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Data_List_filterM__14771079
}

var cache_Data_List_findIndex__2366045378 gopurs_runtime.Value
var once_Data_List_findIndex__2366045378 sync.Once
func Get_Data_List_findIndex__2366045378() gopurs_runtime.Value {
	once_Data_List_findIndex__2366045378.Do(func() {
		cache_Data_List_findIndex__2366045378 = gopurs_runtime.Func(func(fn_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_findIndex__2366045378(fn_0_box)
})
	})
	return cache_Data_List_findIndex__2366045378
}

var cache_Data_List_findLastIndex__2366045378 gopurs_runtime.Value
var once_Data_List_findLastIndex__2366045378 sync.Once
func Get_Data_List_findLastIndex__2366045378() gopurs_runtime.Value {
	once_Data_List_findLastIndex__2366045378.Do(func() {
		cache_Data_List_findLastIndex__2366045378 = gopurs_runtime.Func2(func(fn_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_findLastIndex__2366045378(fn_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_1_box)))}
})
	})
	return cache_Data_List_findLastIndex__2366045378
}

var cache_Data_List_foldM__3577257629 gopurs_runtime.Value
var once_Data_List_foldM__3577257629 sync.Once
func Get_Data_List_foldM__3577257629() gopurs_runtime.Value {
	once_Data_List_foldM__3577257629.Do(func() {
		cache_Data_List_foldM__3577257629 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_foldM__3577257629(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Data_List_foldM__3577257629
}

var cache_Data_List_fromFoldable__614070391 gopurs_runtime.Value
var once_Data_List_fromFoldable__614070391 sync.Once
func Get_Data_List_fromFoldable__614070391() gopurs_runtime.Value {
	once_Data_List_fromFoldable__614070391.Do(func() {
		cache_Data_List_fromFoldable__614070391 = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_fromFoldable__614070391(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box))
})
	})
	return cache_Data_List_fromFoldable__614070391
}

var cache_Data_List_groupAllBy__3934374991 gopurs_runtime.Value
var once_Data_List_groupAllBy__3934374991 sync.Once
func Get_Data_List_groupAllBy__3934374991() gopurs_runtime.Value {
	once_Data_List_groupAllBy__3934374991.Do(func() {
		cache_Data_List_groupAllBy__3934374991 = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_groupAllBy__3934374991(p_0_box)
})
	})
	return cache_Data_List_groupAllBy__3934374991
}

var cache_Data_List_groupBy__2162447253 gopurs_runtime.Value
var once_Data_List_groupBy__2162447253 sync.Once
func Get_Data_List_groupBy__2162447253() gopurs_runtime.Value {
	once_Data_List_groupBy__2162447253.Do(func() {
		cache_Data_List_groupBy__2162447253 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_groupBy__2162447253(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_1_box)))}
})
	})
	return cache_Data_List_groupBy__2162447253
}

var cache_Data_List_groupBy__1039549870 gopurs_runtime.Value
var once_Data_List_groupBy__1039549870 sync.Once
func Get_Data_List_groupBy__1039549870() gopurs_runtime.Value {
	once_Data_List_groupBy__1039549870.Do(func() {
		cache_Data_List_groupBy__1039549870 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_groupBy__1039549870(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_1_box)))}
})
	})
	return cache_Data_List_groupBy__1039549870
}

var cache_Data_List_head__3729839663 gopurs_runtime.Value
var once_Data_List_head__3729839663 sync.Once
func Get_Data_List_head__3729839663() gopurs_runtime.Value {
	once_Data_List_head__3729839663.Do(func() {
		cache_Data_List_head__3729839663 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_head__3729839663(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box)))}
})
	})
	return cache_Data_List_head__3729839663
}

var cache_Data_List_index__304299960 gopurs_runtime.Value
var once_Data_List_index__304299960 sync.Once
func Get_Data_List_index__304299960() gopurs_runtime.Value {
	once_Data_List_index__304299960.Do(func() {
		cache_Data_List_index__304299960 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_index__304299960(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box), v1_1_box.IntVal))}
})
	})
	return cache_Data_List_index__304299960
}

var cache_Data_List_init__2496605985 gopurs_runtime.Value
var once_Data_List_init__2496605985 sync.Once
func Get_Data_List_init__2496605985() gopurs_runtime.Value {
	once_Data_List_init__2496605985.Do(func() {
		cache_Data_List_init__2496605985 = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_init__2496605985(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](lst_0_box)))}
})
	})
	return cache_Data_List_init__2496605985
}

var cache_Data_List_insertAt__2634211748 gopurs_runtime.Value
var once_Data_List_insertAt__2634211748 sync.Once
func Get_Data_List_insertAt__2634211748() gopurs_runtime.Value {
	once_Data_List_insertAt__2634211748.Do(func() {
		cache_Data_List_insertAt__2634211748 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_insertAt__2634211748(v_0_box.IntVal, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v2_2_box)))}
})
	})
	return cache_Data_List_insertAt__2634211748
}

var cache_Data_List_insertBy__1738998985 gopurs_runtime.Value
var once_Data_List_insertBy__1738998985 sync.Once
func Get_Data_List_insertBy__1738998985() gopurs_runtime.Value {
	once_Data_List_insertBy__1738998985.Do(func() {
		cache_Data_List_insertBy__1738998985 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_insertBy__1738998985(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v2_2_box)))}
})
	})
	return cache_Data_List_insertBy__1738998985
}

var cache_Data_List_intersectBy__588351261 gopurs_runtime.Value
var once_Data_List_intersectBy__588351261 sync.Once
func Get_Data_List_intersectBy__588351261() gopurs_runtime.Value {
	once_Data_List_intersectBy__588351261.Do(func() {
		cache_Data_List_intersectBy__588351261 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_intersectBy__588351261(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v2_2_box)))}
})
	})
	return cache_Data_List_intersectBy__588351261
}

var cache_Data_List_intersectBy__1190504998 gopurs_runtime.Value
var once_Data_List_intersectBy__1190504998 sync.Once
func Get_Data_List_intersectBy__1190504998() gopurs_runtime.Value {
	once_Data_List_intersectBy__1190504998.Do(func() {
		cache_Data_List_intersectBy__1190504998 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_intersectBy__1190504998(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v2_2_box)))}
})
	})
	return cache_Data_List_intersectBy__1190504998
}

var cache_Data_List_last__4043133652 gopurs_runtime.Value
var once_Data_List_last__4043133652 sync.Once
func Get_Data_List_last__4043133652() gopurs_runtime.Value {
	once_Data_List_last__4043133652.Do(func() {
		cache_Data_List_last__4043133652 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_last__4043133652(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box)))}
})
	})
	return cache_Data_List_last__4043133652
}

var cache_Data_List_length__3003998832 gopurs_runtime.Value
var once_Data_List_length__3003998832 sync.Once
func Get_Data_List_length__3003998832() gopurs_runtime.Value {
	once_Data_List_length__3003998832.Do(func() {
		cache_Data_List_length__3003998832 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((acc_0.IntVal) + (1))
})
}), gopurs_runtime.Int(0))
	})
	return cache_Data_List_length__3003998832
}

var cache_Data_List_many__542682753 gopurs_runtime.Value
var once_Data_List_many__542682753 sync.Once
func Get_Data_List_many__542682753() gopurs_runtime.Value {
	once_Data_List_many__542682753.Do(func() {
		cache_Data_List_many__542682753 = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_many__542682753(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative](dictAlternative_0_box))
})
	})
	return cache_Data_List_many__542682753
}

var cache_Data_List_manyRec__4046352885 gopurs_runtime.Value
var once_Data_List_manyRec__4046352885 sync.Once
func Get_Data_List_manyRec__4046352885() gopurs_runtime.Value {
	once_Data_List_manyRec__4046352885.Do(func() {
		cache_Data_List_manyRec__4046352885 = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_manyRec__4046352885(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box))
})
	})
	return cache_Data_List_manyRec__4046352885
}

var cache_Data_List_mapMaybe__3262563995 gopurs_runtime.Value
var once_Data_List_mapMaybe__3262563995 sync.Once
func Get_Data_List_mapMaybe__3262563995() gopurs_runtime.Value {
	once_Data_List_mapMaybe__3262563995.Do(func() {
		cache_Data_List_mapMaybe__3262563995 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_mapMaybe__3262563995(f_0_box)
})
	})
	return cache_Data_List_mapMaybe__3262563995
}

var cache_Data_List_mapMaybe__1486753757 gopurs_runtime.Value
var once_Data_List_mapMaybe__1486753757 sync.Once
func Get_Data_List_mapMaybe__1486753757() gopurs_runtime.Value {
	once_Data_List_mapMaybe__1486753757.Do(func() {
		cache_Data_List_mapMaybe__1486753757 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_mapMaybe__1486753757(f_0_box)
})
	})
	return cache_Data_List_mapMaybe__1486753757
}

var cache_Data_List_mapMaybe__1640531773 gopurs_runtime.Value
var once_Data_List_mapMaybe__1640531773 sync.Once
func Get_Data_List_mapMaybe__1640531773() gopurs_runtime.Value {
	once_Data_List_mapMaybe__1640531773.Do(func() {
		cache_Data_List_mapMaybe__1640531773 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_mapMaybe__1640531773(f_0_box)
})
	})
	return cache_Data_List_mapMaybe__1640531773
}

var cache_Data_List_mapMaybe__748617661 gopurs_runtime.Value
var once_Data_List_mapMaybe__748617661 sync.Once
func Get_Data_List_mapMaybe__748617661() gopurs_runtime.Value {
	once_Data_List_mapMaybe__748617661.Do(func() {
		cache_Data_List_mapMaybe__748617661 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_mapMaybe__748617661(f_0_box)
})
	})
	return cache_Data_List_mapMaybe__748617661
}

var cache_Data_List_mapMaybe__4251473821 gopurs_runtime.Value
var once_Data_List_mapMaybe__4251473821 sync.Once
func Get_Data_List_mapMaybe__4251473821() gopurs_runtime.Value {
	once_Data_List_mapMaybe__4251473821.Do(func() {
		cache_Data_List_mapMaybe__4251473821 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_mapMaybe__4251473821(f_0_box)
})
	})
	return cache_Data_List_mapMaybe__4251473821
}

var cache_Data_List_mapMaybe__2491277821 gopurs_runtime.Value
var once_Data_List_mapMaybe__2491277821 sync.Once
func Get_Data_List_mapMaybe__2491277821() gopurs_runtime.Value {
	once_Data_List_mapMaybe__2491277821.Do(func() {
		cache_Data_List_mapMaybe__2491277821 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_mapMaybe__2491277821(f_0_box)
})
	})
	return cache_Data_List_mapMaybe__2491277821
}

var cache_Data_List_modifyAt__1886983628 gopurs_runtime.Value
var once_Data_List_modifyAt__1886983628 sync.Once
func Get_Data_List_modifyAt__1886983628() gopurs_runtime.Value {
	once_Data_List_modifyAt__1886983628.Do(func() {
		cache_Data_List_modifyAt__1886983628 = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_modifyAt__1886983628(n_0_box.IntVal, f_1_box)
})
	})
	return cache_Data_List_modifyAt__1886983628
}

var cache_Data_List_nubBy__2103943131 gopurs_runtime.Value
var once_Data_List_nubBy__2103943131 sync.Once
func Get_Data_List_nubBy__2103943131() gopurs_runtime.Value {
	once_Data_List_nubBy__2103943131.Do(func() {
		cache_Data_List_nubBy__2103943131 = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_nubBy__2103943131(p_0_box)
})
	})
	return cache_Data_List_nubBy__2103943131
}

var cache_Data_List_nubBy__1502591776 gopurs_runtime.Value
var once_Data_List_nubBy__1502591776 sync.Once
func Get_Data_List_nubBy__1502591776() gopurs_runtime.Value {
	once_Data_List_nubBy__1502591776.Do(func() {
		cache_Data_List_nubBy__1502591776 = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_nubBy__1502591776(p_0_box)
})
	})
	return cache_Data_List_nubBy__1502591776
}

var cache_Data_List_nubByEq__3956095361 gopurs_runtime.Value
var once_Data_List_nubByEq__3956095361 sync.Once
func Get_Data_List_nubByEq__3956095361() gopurs_runtime.Value {
	once_Data_List_nubByEq__3956095361.Do(func() {
		cache_Data_List_nubByEq__3956095361 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_nubByEq__3956095361(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_1_box)))}
})
	})
	return cache_Data_List_nubByEq__3956095361
}

var cache_Data_List_nubByEq__3655321914 gopurs_runtime.Value
var once_Data_List_nubByEq__3655321914 sync.Once
func Get_Data_List_nubByEq__3655321914() gopurs_runtime.Value {
	once_Data_List_nubByEq__3655321914.Do(func() {
		cache_Data_List_nubByEq__3655321914 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_nubByEq__3655321914(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_1_box)))}
})
	})
	return cache_Data_List_nubByEq__3655321914
}

var cache_Data_List_null__74357383 gopurs_runtime.Value
var once_Data_List_null__74357383 sync.Once
func Get_Data_List_null__74357383() gopurs_runtime.Value {
	once_Data_List_null__74357383.Do(func() {
		cache_Data_List_null__74357383 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_List_null__74357383(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box)))
})
	})
	return cache_Data_List_null__74357383
}

var cache_Data_List_null__2437342685 gopurs_runtime.Value
var once_Data_List_null__2437342685 sync.Once
func Get_Data_List_null__2437342685() gopurs_runtime.Value {
	once_Data_List_null__2437342685.Do(func() {
		cache_Data_List_null__2437342685 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_List_null__2437342685(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box)))
})
	})
	return cache_Data_List_null__2437342685
}

var cache_Data_List_partition__1623965204 gopurs_runtime.Value
var once_Data_List_partition__1623965204 sync.Once
func Get_Data_List_partition__1623965204() gopurs_runtime.Value {
	once_Data_List_partition__1623965204.Do(func() {
		cache_Data_List_partition__1623965204 = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_partition__1623965204(p_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_1_box))
})
	})
	return cache_Data_List_partition__1623965204
}

var cache_Data_List_reverse__1174136571 gopurs_runtime.Value
var once_Data_List_reverse__1174136571 sync.Once
func Get_Data_List_reverse__1174136571() gopurs_runtime.Value {
	once_Data_List_reverse__1174136571.Do(func() {
		cache_Data_List_reverse__1174136571 = func() gopurs_runtime.Value {
var go__go_0_0_54 gopurs_runtime.Value
go__go_0_0_54 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_1_loop_val)
var v1_2_loop gopurs_runtime.Value = v1_2_loop_val
go__go_0_0_54:
for {
if false { continue go__go_0_0_54 }
var v_1 *Constructor_Data_List_Types_Cons = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t1 *Constructor_Data_List_Types_Cons
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr == nil) {
__t1 = v_1
goto end_branch_1
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr != nil) {
v_1_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_2.UnsafePtr).V0, v_1})})
v1_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_2.UnsafePtr).V1)}
continue go__go_0_0_54
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
})
return gopurs_runtime.Apply(go__go_0_0_54, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}()
	})
	return cache_Data_List_reverse__1174136571
}

var cache_Data_List_reverse__2479281851 gopurs_runtime.Value
var once_Data_List_reverse__2479281851 sync.Once
func Get_Data_List_reverse__2479281851() gopurs_runtime.Value {
	once_Data_List_reverse__2479281851.Do(func() {
		cache_Data_List_reverse__2479281851 = func() gopurs_runtime.Value {
var go__go_0_0_55 gopurs_runtime.Value
go__go_0_0_55 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_1_loop_val)
var v1_2_loop gopurs_runtime.Value = v1_2_loop_val
go__go_0_0_55:
for {
if false { continue go__go_0_0_55 }
var v_1 *Constructor_Data_List_Types_Cons = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t1 *Constructor_Data_List_Types_Cons
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr == nil) {
__t1 = v_1
goto end_branch_1
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr != nil) {
v_1_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple]((*Constructor_Data_List_Types_Cons)(v1_2.UnsafePtr).V0))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)})})})
v1_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_2.UnsafePtr).V1)}
continue go__go_0_0_55
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
})
return gopurs_runtime.Apply(go__go_0_0_55, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}()
	})
	return cache_Data_List_reverse__2479281851
}

var cache_Data_List_reverse__103602267 gopurs_runtime.Value
var once_Data_List_reverse__103602267 sync.Once
func Get_Data_List_reverse__103602267() gopurs_runtime.Value {
	once_Data_List_reverse__103602267.Do(func() {
		cache_Data_List_reverse__103602267 = func() gopurs_runtime.Value {
var go__go_0_0_56 gopurs_runtime.Value
go__go_0_0_56 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_1_loop_val)
var v1_2_loop gopurs_runtime.Value = v1_2_loop_val
go__go_0_0_56:
for {
if false { continue go__go_0_0_56 }
var v_1 *Constructor_Data_List_Types_Cons = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t1 *Constructor_Data_List_Types_Cons
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr == nil) {
__t1 = v_1
goto end_branch_1
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr != nil) {
v_1_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_2.UnsafePtr).V0, v_1})})
v1_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_2.UnsafePtr).V1)}
continue go__go_0_0_56
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
})
return gopurs_runtime.Apply(go__go_0_0_56, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}()
	})
	return cache_Data_List_reverse__103602267
}

var cache_Data_List_reverse__4230102656 gopurs_runtime.Value
var once_Data_List_reverse__4230102656 sync.Once
func Get_Data_List_reverse__4230102656() gopurs_runtime.Value {
	once_Data_List_reverse__4230102656.Do(func() {
		cache_Data_List_reverse__4230102656 = func() gopurs_runtime.Value {
var go__go_0_0_57 gopurs_runtime.Value
go__go_0_0_57 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_1_loop_val)
var v1_2_loop gopurs_runtime.Value = v1_2_loop_val
go__go_0_0_57:
for {
if false { continue go__go_0_0_57 }
var v_1 *Constructor_Data_List_Types_Cons = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t1 *Constructor_Data_List_Types_Cons
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr == nil) {
__t1 = v_1
goto end_branch_1
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr != nil) {
v_1_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_2.UnsafePtr).V0, v_1})})
v1_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_2.UnsafePtr).V1)}
continue go__go_0_0_57
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
})
return gopurs_runtime.Apply(go__go_0_0_57, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}()
	})
	return cache_Data_List_reverse__4230102656
}

var cache_Data_List_reverse__682228544 gopurs_runtime.Value
var once_Data_List_reverse__682228544 sync.Once
func Get_Data_List_reverse__682228544() gopurs_runtime.Value {
	once_Data_List_reverse__682228544.Do(func() {
		cache_Data_List_reverse__682228544 = func() gopurs_runtime.Value {
var go__go_0_0_58 gopurs_runtime.Value
go__go_0_0_58 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_1_loop_val)
var v1_2_loop gopurs_runtime.Value = v1_2_loop_val
go__go_0_0_58:
for {
if false { continue go__go_0_0_58 }
var v_1 *Constructor_Data_List_Types_Cons = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t1 *Constructor_Data_List_Types_Cons
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr == nil) {
__t1 = v_1
goto end_branch_1
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr != nil) {
v_1_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_List_Types_Cons)(v1_2.UnsafePtr).V0))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)})})})
v1_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_2.UnsafePtr).V1)}
continue go__go_0_0_58
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
})
return gopurs_runtime.Apply(go__go_0_0_58, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}()
	})
	return cache_Data_List_reverse__682228544
}

var cache_Data_List_reverse__684006912 gopurs_runtime.Value
var once_Data_List_reverse__684006912 sync.Once
func Get_Data_List_reverse__684006912() gopurs_runtime.Value {
	once_Data_List_reverse__684006912.Do(func() {
		cache_Data_List_reverse__684006912 = func() gopurs_runtime.Value {
var go__go_0_0_59 gopurs_runtime.Value
go__go_0_0_59 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_1_loop_val)
var v1_2_loop gopurs_runtime.Value = v1_2_loop_val
go__go_0_0_59:
for {
if false { continue go__go_0_0_59 }
var v_1 *Constructor_Data_List_Types_Cons = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t1 *Constructor_Data_List_Types_Cons
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr == nil) {
__t1 = v_1
goto end_branch_1
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr != nil) {
v_1_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((*Constructor_Data_List_Types_Cons)(v1_2.UnsafePtr).V0))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)})})})
v1_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_2.UnsafePtr).V1)}
continue go__go_0_0_59
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
})
return gopurs_runtime.Apply(go__go_0_0_59, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}()
	})
	return cache_Data_List_reverse__684006912
}

var cache_Data_List_reverse__1758336384 gopurs_runtime.Value
var once_Data_List_reverse__1758336384 sync.Once
func Get_Data_List_reverse__1758336384() gopurs_runtime.Value {
	once_Data_List_reverse__1758336384.Do(func() {
		cache_Data_List_reverse__1758336384 = func() gopurs_runtime.Value {
var go__go_0_0_60 gopurs_runtime.Value
go__go_0_0_60 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_1_loop_val)
var v1_2_loop gopurs_runtime.Value = v1_2_loop_val
go__go_0_0_60:
for {
if false { continue go__go_0_0_60 }
var v_1 *Constructor_Data_List_Types_Cons = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t1 *Constructor_Data_List_Types_Cons
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr == nil) {
__t1 = v_1
goto end_branch_1
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr != nil) {
v_1_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple]((*Constructor_Data_List_Types_Cons)(v1_2.UnsafePtr).V0))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)})})})
v1_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_2.UnsafePtr).V1)}
continue go__go_0_0_60
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
})
return gopurs_runtime.Apply(go__go_0_0_60, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}()
	})
	return cache_Data_List_reverse__1758336384
}

var cache_Data_List_singleton__2450819477 gopurs_runtime.Value
var once_Data_List_singleton__2450819477 sync.Once
func Get_Data_List_singleton__2450819477() gopurs_runtime.Value {
	once_Data_List_singleton__2450819477.Do(func() {
		cache_Data_List_singleton__2450819477 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_singleton__2450819477(a_0_box.IntVal))}
})
	})
	return cache_Data_List_singleton__2450819477
}

var cache_Data_List_singleton__707062261 gopurs_runtime.Value
var once_Data_List_singleton__707062261 sync.Once
func Get_Data_List_singleton__707062261() gopurs_runtime.Value {
	once_Data_List_singleton__707062261.Do(func() {
		cache_Data_List_singleton__707062261 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_singleton__707062261(a_0_box))}
})
	})
	return cache_Data_List_singleton__707062261
}

var cache_Data_List_singleton__3932757557 gopurs_runtime.Value
var once_Data_List_singleton__3932757557 sync.Once
func Get_Data_List_singleton__3932757557() gopurs_runtime.Value {
	once_Data_List_singleton__3932757557.Do(func() {
		cache_Data_List_singleton__3932757557 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_singleton__3932757557(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](a_0_box)))}
})
	})
	return cache_Data_List_singleton__3932757557
}

var cache_Data_List_snoc__4290067657 gopurs_runtime.Value
var once_Data_List_snoc__4290067657 sync.Once
func Get_Data_List_snoc__4290067657() gopurs_runtime.Value {
	once_Data_List_snoc__4290067657.Do(func() {
		cache_Data_List_snoc__4290067657 = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_snoc__4290067657(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_0_box), x_1_box))}
})
	})
	return cache_Data_List_snoc__4290067657
}

var cache_Data_List_some__542682753 gopurs_runtime.Value
var once_Data_List_some__542682753 sync.Once
func Get_Data_List_some__542682753() gopurs_runtime.Value {
	once_Data_List_some__542682753.Do(func() {
		cache_Data_List_some__542682753 = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_some__542682753(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative](dictAlternative_0_box))
})
	})
	return cache_Data_List_some__542682753
}

var cache_Data_List_sortBy__2103943131 gopurs_runtime.Value
var once_Data_List_sortBy__2103943131 sync.Once
func Get_Data_List_sortBy__2103943131() gopurs_runtime.Value {
	once_Data_List_sortBy__2103943131.Do(func() {
		cache_Data_List_sortBy__2103943131 = gopurs_runtime.Func(func(cmp_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_sortBy__2103943131(cmp_0_box)
})
	})
	return cache_Data_List_sortBy__2103943131
}

var cache_Data_List_sortBy__1502591776 gopurs_runtime.Value
var once_Data_List_sortBy__1502591776 sync.Once
func Get_Data_List_sortBy__1502591776() gopurs_runtime.Value {
	once_Data_List_sortBy__1502591776.Do(func() {
		cache_Data_List_sortBy__1502591776 = gopurs_runtime.Func(func(cmp_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_sortBy__1502591776(cmp_0_box)
})
	})
	return cache_Data_List_sortBy__1502591776
}

var cache_Data_List_span__1918198736 gopurs_runtime.Value
var once_Data_List_span__1918198736 sync.Once
func Get_Data_List_span__1918198736() gopurs_runtime.Value {
	once_Data_List_span__1918198736.Do(func() {
		cache_Data_List_span__1918198736 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_span__1918198736(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_1_box))
})
	})
	return cache_Data_List_span__1918198736
}

var cache_Data_List_span__799093643 gopurs_runtime.Value
var once_Data_List_span__799093643 sync.Once
func Get_Data_List_span__799093643() gopurs_runtime.Value {
	once_Data_List_span__799093643.Do(func() {
		cache_Data_List_span__799093643 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_span__799093643(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_1_box))
})
	})
	return cache_Data_List_span__799093643
}

var cache_Data_List_span__2133741451 gopurs_runtime.Value
var once_Data_List_span__2133741451 sync.Once
func Get_Data_List_span__2133741451() gopurs_runtime.Value {
	once_Data_List_span__2133741451.Do(func() {
		cache_Data_List_span__2133741451 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_span__2133741451(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_1_box))
})
	})
	return cache_Data_List_span__2133741451
}

var cache_Data_List_tail__1771843450 gopurs_runtime.Value
var once_Data_List_tail__1771843450 sync.Once
func Get_Data_List_tail__1771843450() gopurs_runtime.Value {
	once_Data_List_tail__1771843450.Do(func() {
		cache_Data_List_tail__1771843450 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_tail__1771843450(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box)))}
})
	})
	return cache_Data_List_tail__1771843450
}

var cache_Data_List_tails__3932757557 gopurs_runtime.Value
var once_Data_List_tails__3932757557 sync.Once
func Get_Data_List_tails__3932757557() gopurs_runtime.Value {
	once_Data_List_tails__3932757557.Do(func() {
		cache_Data_List_tails__3932757557 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_tails__3932757557(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box)))}
})
	})
	return cache_Data_List_tails__3932757557
}

var cache_Data_List_take__551729751 gopurs_runtime.Value
var once_Data_List_take__551729751 sync.Once
func Get_Data_List_take__551729751() gopurs_runtime.Value {
	once_Data_List_take__551729751.Do(func() {
		cache_Data_List_take__551729751 = func() gopurs_runtime.Value {
var go__go_0_0_73 gopurs_runtime.Value
go__go_0_0_73 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_1_loop_val)
var v1_2_loop int64 = v1_2_loop_val.IntVal
var v2_3_loop gopurs_runtime.Value = v2_3_loop_val
go__go_0_0_73:
for {
if false { continue go__go_0_0_73 }
var v_1 *Constructor_Data_List_Types_Cons = v_1_loop
_ = v_1
var v1_2 int64 = v1_2_loop
_ = v1_2
var v2_3 gopurs_runtime.Value = v2_3_loop
_ = v2_3
var __t6 *Constructor_Data_List_Types_Cons
{
var __t1 bool
{
if (v1_2) < (1) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
if __t1 {
var go__go_4_2_74 gopurs_runtime.Value
go__go_4_2_74 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_2_74:
for {
if false { continue go__go_4_2_74 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t3 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_2_74
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_2_74, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}))
goto end_branch_6
} else {

}
}
{
if (v2_3.Type == 9 && v2_3.IntVal == 1358893437 && v2_3.UnsafePtr == nil) {
var go__go_4_4_75 gopurs_runtime.Value
go__go_4_4_75 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_4_75:
for {
if false { continue go__go_4_4_75 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t5 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t5 = v_5
goto end_branch_5
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_4_75
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t5)}
}
}()
})
})
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_4_75, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}))
goto end_branch_6
} else {

}
}
{
if (v2_3.Type == 9 && v2_3.IntVal == 1358893437 && v2_3.UnsafePtr != nil) {
v_1_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v2_3.UnsafePtr).V0, v_1})})
v1_2_loop = (v1_2) - (1)
v2_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v2_3.UnsafePtr).V1)}
continue go__go_0_0_73
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t6)}
}
}()
})
})
})
return gopurs_runtime.Apply(go__go_0_0_73, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}()
	})
	return cache_Data_List_take__551729751
}

var cache_Data_List_take__1836090668 gopurs_runtime.Value
var once_Data_List_take__1836090668 sync.Once
func Get_Data_List_take__1836090668() gopurs_runtime.Value {
	once_Data_List_take__1836090668.Do(func() {
		cache_Data_List_take__1836090668 = func() gopurs_runtime.Value {
var go__go_0_0_76 gopurs_runtime.Value
go__go_0_0_76 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_1_loop_val)
var v1_2_loop int64 = v1_2_loop_val.IntVal
var v2_3_loop gopurs_runtime.Value = v2_3_loop_val
go__go_0_0_76:
for {
if false { continue go__go_0_0_76 }
var v_1 *Constructor_Data_List_Types_Cons = v_1_loop
_ = v_1
var v1_2 int64 = v1_2_loop
_ = v1_2
var v2_3 gopurs_runtime.Value = v2_3_loop
_ = v2_3
var __t6 *Constructor_Data_List_Types_Cons
{
var __t1 bool
{
if (v1_2) < (1) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
if __t1 {
var go__go_4_2_77 gopurs_runtime.Value
go__go_4_2_77 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_2_77:
for {
if false { continue go__go_4_2_77 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t3 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_2_77
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_2_77, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}))
goto end_branch_6
} else {

}
}
{
if (v2_3.Type == 9 && v2_3.IntVal == 1358893437 && v2_3.UnsafePtr == nil) {
var go__go_4_4_78 gopurs_runtime.Value
go__go_4_4_78 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_4_78:
for {
if false { continue go__go_4_4_78 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t5 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t5 = v_5
goto end_branch_5
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_4_78
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t5)}
}
}()
})
})
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_4_78, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}))
goto end_branch_6
} else {

}
}
{
if (v2_3.Type == 9 && v2_3.IntVal == 1358893437 && v2_3.UnsafePtr != nil) {
v_1_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v2_3.UnsafePtr).V0, v_1})})
v1_2_loop = (v1_2) - (1)
v2_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v2_3.UnsafePtr).V1)}
continue go__go_0_0_76
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t6)}
}
}()
})
})
})
return gopurs_runtime.Apply(go__go_0_0_76, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}()
	})
	return cache_Data_List_take__1836090668
}

var cache_Data_List_takeWhile__2352021032 gopurs_runtime.Value
var once_Data_List_takeWhile__2352021032 sync.Once
func Get_Data_List_takeWhile__2352021032() gopurs_runtime.Value {
	once_Data_List_takeWhile__2352021032.Do(func() {
		cache_Data_List_takeWhile__2352021032 = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_takeWhile__2352021032(p_0_box)
})
	})
	return cache_Data_List_takeWhile__2352021032
}

var cache_Data_List_transpose__682228544 gopurs_runtime.Value
var once_Data_List_transpose__682228544 sync.Once
func Get_Data_List_transpose__682228544() gopurs_runtime.Value {
	once_Data_List_transpose__682228544.Do(func() {
		cache_Data_List_transpose__682228544 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_transpose__682228544(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box)))}
})
	})
	return cache_Data_List_transpose__682228544
}

var cache_Data_List_uncons__3009258782 gopurs_runtime.Value
var once_Data_List_uncons__3009258782 sync.Once
func Get_Data_List_uncons__3009258782() gopurs_runtime.Value {
	once_Data_List_uncons__3009258782.Do(func() {
		cache_Data_List_uncons__3009258782 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_uncons__3009258782(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_0_box)))}
})
	})
	return cache_Data_List_uncons__3009258782
}

var cache_Data_List_unionBy__588351261 gopurs_runtime.Value
var once_Data_List_unionBy__588351261 sync.Once
func Get_Data_List_unionBy__588351261() gopurs_runtime.Value {
	once_Data_List_unionBy__588351261.Do(func() {
		cache_Data_List_unionBy__588351261 = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_unionBy__588351261(eq_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_2_box)))}
})
	})
	return cache_Data_List_unionBy__588351261
}

var cache_Data_List_unionBy__1190504998 gopurs_runtime.Value
var once_Data_List_unionBy__1190504998 sync.Once
func Get_Data_List_unionBy__1190504998() gopurs_runtime.Value {
	once_Data_List_unionBy__1190504998.Do(func() {
		cache_Data_List_unionBy__1190504998 = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_unionBy__1190504998(eq_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_2_box)))}
})
	})
	return cache_Data_List_unionBy__1190504998
}

var cache_Data_List_unsnoc__2942606998 gopurs_runtime.Value
var once_Data_List_unsnoc__2942606998 sync.Once
func Get_Data_List_unsnoc__2942606998() gopurs_runtime.Value {
	once_Data_List_unsnoc__2942606998.Do(func() {
		cache_Data_List_unsnoc__2942606998 = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_unsnoc__2942606998(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](lst_0_box)))}
})
	})
	return cache_Data_List_unsnoc__2942606998
}

var cache_Data_List_updateAt__2634211748 gopurs_runtime.Value
var once_Data_List_updateAt__2634211748 sync.Once
func Get_Data_List_updateAt__2634211748() gopurs_runtime.Value {
	once_Data_List_updateAt__2634211748.Do(func() {
		cache_Data_List_updateAt__2634211748 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_updateAt__2634211748(v_0_box.IntVal, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v2_2_box)))}
})
	})
	return cache_Data_List_updateAt__2634211748
}

var cache_Data_List_zipWith__884793877 gopurs_runtime.Value
var once_Data_List_zipWith__884793877 sync.Once
func Get_Data_List_zipWith__884793877() gopurs_runtime.Value {
	once_Data_List_zipWith__884793877.Do(func() {
		cache_Data_List_zipWith__884793877 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_zipWith__884793877(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_2_box)))}
})
	})
	return cache_Data_List_zipWith__884793877
}

var cache_Data_List_zipWith__4203240021 gopurs_runtime.Value
var once_Data_List_zipWith__4203240021 sync.Once
func Get_Data_List_zipWith__4203240021() gopurs_runtime.Value {
	once_Data_List_zipWith__4203240021.Do(func() {
		cache_Data_List_zipWith__4203240021 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_zipWith__4203240021(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_2_box)))}
})
	})
	return cache_Data_List_zipWith__4203240021
}

var cache_Data_List_zipWith__3856182069 gopurs_runtime.Value
var once_Data_List_zipWith__3856182069 sync.Once
func Get_Data_List_zipWith__3856182069() gopurs_runtime.Value {
	once_Data_List_zipWith__3856182069.Do(func() {
		cache_Data_List_zipWith__3856182069 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_zipWith__3856182069(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_2_box)))}
})
	})
	return cache_Data_List_zipWith__3856182069
}

func Call_Data_List_identity(x_0_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var x_0 *Constructor_Data_List_Types_Cons = x_0_loop
_ = x_0
return x_0
}

func Call_Data_List_Pattern(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_List_updateAt(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
updateAt:
for {
if false { continue updateAt }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_List_Types_Cons = v2_2_loop
_ = v2_2
var __t2 *Constructor_Data_Maybe_Just
{
if (v2_2 != nil) {
var __t1 *Constructor_Data_Maybe_Just
{
if (v_0) == (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, v1_1, (v2_2).V1})}})})
goto end_branch_1
} else {

}
}
{
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := (v2_2).V0
_ = __local_var_3_0
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(v3_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, __local_var_3_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v3_4)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_updateAt((v_0) - (1), v1_1, (v2_2).V1))}))
}
end_branch_1:
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)})
goto end_branch_2
} else {

}
}
{
__t2 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)})
}
}

func Call_Data_List_uncons(v_0_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 == nil) {
__t0 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_0
} else {

}
}
{
if (v_0 != nil) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("head", "tail", (v_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)})
}

func Call_Data_List_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable = dictUnfoldable_0_loop
_ = dictUnfoldable_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable_0.V1), gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Maybe_Just
{
var __t_tag_0 *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_1)
if (__t_tag_0 == nil) {
__t2 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_1)
if (__t_tag_1 != nil) {
__t2 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("head", "tail", (*Constructor_Data_List_Types_Cons)(xs_1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(xs_1.UnsafePtr).V1)})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(rec_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.RecordGet(rec_2, "head"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(rec_2, "tail")))}})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)})))}
}))
}

func Call_Data_List_tail(v_0_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 == nil) {
__t0 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_0
} else {

}
}
{
if (v_0 != nil) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)}}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)})
}

func Call_Data_List_stripPrefix(dictEq_0_loop *Constructor_Data_Eq_Eq, v_1_loop *Constructor_Data_List_Types_Cons, s_2_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var v_1 *Constructor_Data_List_Types_Cons = v_1_loop
_ = v_1
var s_2 *Constructor_Data_List_Types_Cons = s_2_loop
_ = s_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_Rec_Class_monadRecMaybe(), "tailRecM"), gopurs_runtime.Func(func(o_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 *Constructor_Data_Maybe_Just
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.RecordGet(o_3, "b")
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 1358893437 && __t_tag_0.UnsafePtr != nil) {
var __t4 *Constructor_Data_Maybe_Just
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.RecordGet(o_3, "a")
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1358893437 && __t_tag_1.UnsafePtr != nil) {
var __t2 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), (*Constructor_Data_List_Types_Cons)(gopurs_runtime.RecordGet(o_3, "a").UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(gopurs_runtime.RecordGet(o_3, "b").UnsafePtr).V0).IntVal) != (0) {
__t2 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(gopurs_runtime.RecordGet(o_3, "a").UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(gopurs_runtime.RecordGet(o_3, "b").UnsafePtr).V1)})})}}
goto end_branch_2
} else {

}
}
{
__t2 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_2:
__t4 = __t2
goto end_branch_4
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.RecordGet(o_3, "a")
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 1358893437 && __t_tag_3.UnsafePtr == nil) {
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(o_3, "b")))}})}}
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_4:
__t6 = __t4
goto end_branch_6
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.RecordGet(o_3, "a")
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 1358893437 && __t_tag_5.UnsafePtr == nil) {
__t6 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(o_3, "b")))}})}}
goto end_branch_6
} else {

}
}
{
__t6 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t6)}
}), gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(s_2)})))
}

func Call_Data_List_span(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_List_Types_Cons) gopurs_runtime.Value {
span:
for {
if false { continue span }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_List_Types_Cons = v1_1_loop
_ = v1_1
var __t1 gopurs_runtime.Value
{
if ((v1_1 != nil)) && ((gopurs_runtime.Apply(v_0, (v1_1).V0).IntVal) != (0)) {
// TAST (Let): v2_2_0 -> gopurs_runtime.Value
var v2_2_0 gopurs_runtime.Value = Call_Data_List_span(v_0, (v1_1).V1)
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v1_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v2_2_0, "init"))})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v2_2_0, "rest")))})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)})
}
end_branch_1:
return __t1
}
}

func Call_Data_List_snoc(xs_0_loop *Constructor_Data_List_Types_Cons, x_1_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons {
var xs_0 *Constructor_Data_List_Types_Cons = xs_0_loop
_ = xs_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldr"), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, x_1, (*Constructor_Data_List_Types_Cons)(nil)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_0)}))
}

func Call_Data_List_singleton(a_0_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return &Constructor_Data_List_Types_Cons{1, a_0, (*Constructor_Data_List_Types_Cons)(nil)}
}

func Call_Data_List_sortBy(cmp_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
var merge_1_0_0 gopurs_runtime.Value
_ = merge_1_0_0
merge_1_0_0 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_List_Types_Cons
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr != nil) {
var __t4 *Constructor_Data_List_Types_Cons
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
var __t3 *Constructor_Data_List_Types_Cons
{
// TAST (Let): __local_var_4_1 -> uint32
__local_var_4_1 := uint32(gopurs_runtime.Apply2(cmp_0, (*Constructor_Data_List_Types_Cons)(v_2.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0).IntVal)
_ = __local_var_4_1
var __t2 bool
{
if (__local_var_4_1 == 1527465420) {
__t2 = false
goto end_branch_2
} else {

}
}
{
if (__local_var_4_1 == 380165415) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
if __t2 {
__t3 = &Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(merge_1_0_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}))}
goto end_branch_3
} else {

}
}
{
__t3 = &Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(merge_1_0_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_2.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_3))}))}
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_3)
goto end_branch_5
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2)
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t5)}
})
})
var mergePairs_2_6_1 gopurs_runtime.Value
_ = mergePairs_2_6_1
mergePairs_2_6_1 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 *Constructor_Data_List_Types_Cons
{
var __t_and_8 bool = false
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {

var __t_tag_7 *Constructor_Data_List_Types_Cons = (*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V1
__t_and_8 = (__t_tag_7 != nil)
}
if __t_and_8 {
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(merge_1_0_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V0))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V1).V0))})))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(mergePairs_2_6_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V1).V1)})))})})})
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_3)
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t9)}
})
var mergeAll_3_10_2 gopurs_runtime.Value
mergeAll_3_10_2 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
mergeAll_3_10_2:
for {
if false { continue mergeAll_3_10_2 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var __t13 *Constructor_Data_List_Types_Cons
{
var __t_and_12 bool = false
if (v_4 != nil) {

var __t_tag_11 *Constructor_Data_List_Types_Cons = (v_4).V1
__t_and_12 = (__t_tag_11 == nil)
}
if __t_and_12 {
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_4).V0)
goto end_branch_13
} else {

}
}
{
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(mergePairs_2_6_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}))
continue mergeAll_3_10_2
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
}
end_branch_13:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t13)}
}
}()
})
var sequences_4_14_3 gopurs_runtime.Value
_ = sequences_4_14_3
var descending_4_15_4 gopurs_runtime.Value
_ = descending_4_15_4
var ascending_4_16_5 gopurs_runtime.Value
_ = ascending_4_16_5
sequences_4_14_3 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t23 *Constructor_Data_List_Types_Cons
{
var __t_and_18 bool = false
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {

var __t_tag_17 *Constructor_Data_List_Types_Cons = (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1
__t_and_18 = (__t_tag_17 != nil)
}
if __t_and_18 {
var __t22 *Constructor_Data_List_Types_Cons
{
// TAST (Let): __local_var_6_20 -> uint32
__local_var_6_20 := uint32(gopurs_runtime.Apply2(cmp_0, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0, ((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1).V0).IntVal)
_ = __local_var_6_20
var __t21 bool
{
if (__local_var_6_20 == 1527465420) {
__t21 = false
goto end_branch_21
} else {

}
}
{
if (__local_var_6_20 == 380165415) {
__t21 = true
goto end_branch_21
} else {

}
}
{
__t21 = false
}
end_branch_21:
if __t21 {
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(descending_4_15_4, ((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(nil)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1).V1)}))
goto end_branch_22
} else {

}
}
{
// TAST (Let): __local_var_6_19 -> gopurs_runtime.Value
__local_var_6_19 := (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0
_ = __local_var_6_19
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(ascending_4_16_5, ((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1).V0, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, __local_var_6_19, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_7)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1).V1)}))
}
end_branch_22:
__t23 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t22)})
goto end_branch_23
} else {

}
}
{
__t23 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5))}, (*Constructor_Data_List_Types_Cons)(nil)}
}
end_branch_23:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t23)}
})
descending_4_15_4 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t27 *Constructor_Data_List_Types_Cons
{
var __t_and_26 bool = false
if (v2_7.Type == 9 && v2_7.IntVal == 1358893437 && v2_7.UnsafePtr != nil) {

// TAST (Let): __local_var_8_24 -> uint32
__local_var_8_24 := uint32(gopurs_runtime.Apply2(cmp_0, v_5, (*Constructor_Data_List_Types_Cons)(v2_7.UnsafePtr).V0).IntVal)
_ = __local_var_8_24
var __t25 bool
{
if (__local_var_8_24 == 1527465420) {
__t25 = false
goto end_branch_25
} else {

}
}
{
if (__local_var_8_24 == 380165415) {
__t25 = true
goto end_branch_25
} else {

}
}
{
__t25 = false
}
end_branch_25:
__t_and_26 = __t25
}
if __t_and_26 {
__t27 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(descending_4_15_4, (*Constructor_Data_List_Types_Cons)(v2_7.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, v_5, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_6)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v2_7.UnsafePtr).V1)})))})
goto end_branch_27
} else {

}
}
{
__t27 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, v_5, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_6)})}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(sequences_4_14_3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v2_7))})))})}
}
end_branch_27:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t27)}
})
})
})
ascending_4_16_5 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t28 *Constructor_Data_List_Types_Cons
{
if ((v2_7.Type == 9 && v2_7.IntVal == 1358893437 && v2_7.UnsafePtr != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Ordering_eqOrdering(), "eq"), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(cmp_0, v_5, (*Constructor_Data_List_Types_Cons)(v2_7.UnsafePtr).V0).IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}).IntVal) != (0)) != (true)) {
__t28 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(ascending_4_16_5, (*Constructor_Data_List_Types_Cons)(v2_7.UnsafePtr).V0, gopurs_runtime.Func(func(ys_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(v1_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, v_5, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_8)})})))}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v2_7.UnsafePtr).V1)})))})
goto end_branch_28
} else {

}
}
{
__t28 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(v1_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, v_5, (*Constructor_Data_List_Types_Cons)(nil)})})))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(sequences_4_14_3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v2_7))})))})}
}
end_branch_28:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t28)}
})
})
})
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(mergeAll_3_10_2, gopurs_runtime.Apply(sequences_4_14_3, x_5))
})
}

func Call_Data_List_sort(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(Call_Data_List_sortBy(compare_1_0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_2))})))}
})
}

func Call_Data_List_tails(v_0_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
tails:
for {
if false { continue tails }
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
var __t0 *Constructor_Data_List_Types_Cons
{
if (v_0 == nil) {
__t0 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, (*Constructor_Data_List_Types_Cons)(nil)}
goto end_branch_0
} else {

}
}
{
if (v_0 != nil) {
__t0 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_tails((v_0).V1))})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t0)})
}
}

func Call_Data_List_showPattern(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): showList_1_0 -> *Constructor_Data_Show_Show
showList_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](gopurs_runtime.Apply(Get_Data_List_Types_showList(), dictShow_0))
_ = showList_1_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Pattern ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showList_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2))}).StrVal())) + (")"))
}))
}

func Call_Data_List_takeWhile(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_10 gopurs_runtime.Value
go__go_1_0_10 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_10:
for {
if false { continue go__go_1_0_10 }
var v_2 *Constructor_Data_List_Types_Cons = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t3 *Constructor_Data_List_Types_Cons
{
if ((v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil)) && ((gopurs_runtime.Apply(p_0, (*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0).IntVal) != (0)) {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0, v_2})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}
continue go__go_1_0_10
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
var go__go_4_1_11 gopurs_runtime.Value
go__go_4_1_11 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_11:
for {
if false { continue go__go_4_1_11 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = v_5
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_1_11
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t2)}
}
}()
})
})
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_1_11, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}))
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_10, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}

func Call_Data_List_unsnoc(lst_0_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
var lst_0 *Constructor_Data_List_Types_Cons = lst_0_loop
_ = lst_0
var go__go_1_0_12 gopurs_runtime.Value
go__go_1_0_12 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_3_loop_val)
go__go_1_0_12:
for {
if false { continue go__go_1_0_12 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Types_Cons = v1_3_loop
_ = v1_3
var __t3 *Constructor_Data_Maybe_Just
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr == nil) {
__t3 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr != nil) {
var __t2 *Constructor_Data_Maybe_Just
{
var __t_tag_1 *Constructor_Data_List_Types_Cons = (*Constructor_Data_List_Types_Cons)(v_2.UnsafePtr).V1
if (__t_tag_1 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("last", "revInit", (*Constructor_Data_List_Types_Cons)(v_2.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)})})})
goto end_branch_2
} else {

}
}
{
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_2.UnsafePtr).V1)}
v1_3_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_2.UnsafePtr).V0, v1_3})})
continue go__go_1_0_12
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
}
end_branch_2:
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(h_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_4_13 gopurs_runtime.Value
go__go_3_4_13 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_4_13:
for {
if false { continue go__go_3_4_13 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t5 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t5 = v_4
goto end_branch_5
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_4_13
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t5)}
}
}()
})
})
return gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_3_4_13, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(h_2, "revInit")))})))}, gopurs_runtime.RecordGet(h_2, "last"))
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(go__go_1_0_12, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(lst_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}))
}

func Call_Data_List_zipWith(f_0_loop gopurs_runtime.Value, xs_1_loop *Constructor_Data_List_Types_Cons, ys_2_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 *Constructor_Data_List_Types_Cons = xs_1_loop
_ = xs_1
var ys_2 *Constructor_Data_List_Types_Cons = ys_2_loop
_ = ys_2
var go__go_3_0_14 gopurs_runtime.Value
go__go_3_0_14 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
var v2_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v2_6_loop_val)
go__go_3_0_14:
for {
if false { continue go__go_3_0_14 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var v2_6 *Constructor_Data_List_Types_Cons = v2_6_loop
_ = v2_6
var __t1 *Constructor_Data_List_Types_Cons
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t1 = v2_6
goto end_branch_1
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t1 = v2_6
goto end_branch_1
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil)) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
v2_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Apply2(f_0, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0), v2_6})})
continue go__go_3_0_14
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
})
})
var go__go_4_2_15 gopurs_runtime.Value
go__go_4_2_15 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_2_15:
for {
if false { continue go__go_4_2_15 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t3 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_2_15
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_2_15, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(go__go_3_0_14, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(ys_2)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}))
}

func Call_Data_List_zipWithA(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative, f_1_loop gopurs_runtime.Value, xs_2_loop *Constructor_Data_List_Types_Cons, ys_3_loop *Constructor_Data_List_Types_Cons) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 *Constructor_Data_List_Types_Cons = xs_2_loop
_ = xs_2
var ys_3 *Constructor_Data_List_Types_Cons = ys_3_loop
_ = ys_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_traversableList(), "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_zipWith(f_1, xs_2, ys_3))})
}

func Call_Data_List_go__range(start_0_loop int64, end_1_loop int64) *Constructor_Data_List_Types_Cons {
var start_0 int64 = start_0_loop
_ = start_0
var end_1 int64 = end_1_loop
_ = end_1
var __t4 *Constructor_Data_List_Types_Cons
{
if (start_0) == (end_1) {
__t4 = Call_Data_List_singleton__2450819477(start_0)
goto end_branch_4
} else {

}
}
{
var go__go_2_0_16 gopurs_runtime.Value
go__go_2_0_16 = gopurs_runtime.Func(func(s_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(step_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rest_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var s_3_loop int64 = s_3_loop_val.IntVal
var e_4_loop int64 = e_4_loop_val.IntVal
var step_5_loop int64 = step_5_loop_val.IntVal
var rest_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](rest_6_loop_val)
go__go_2_0_16:
for {
if false { continue go__go_2_0_16 }
var s_3 int64 = s_3_loop
_ = s_3
var e_4 int64 = e_4_loop
_ = e_4
var step_5 int64 = step_5_loop
_ = step_5
var rest_6 *Constructor_Data_List_Types_Cons = rest_6_loop
_ = rest_6
var __t1 *Constructor_Data_List_Types_Cons
{
if (s_3) == (e_4) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Int(s_3), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(rest_6)})})})
goto end_branch_1
} else {

}
}
{
s_3_loop = (s_3) + (step_5)
e_4_loop = e_4
step_5_loop = step_5
rest_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Int(s_3), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(rest_6)})})})
continue go__go_2_0_16
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
})
})
})
var __t3 int64
{
var __t2 bool
{
if (start_0) > (end_1) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
if __t2 {
__t3 = 1
goto end_branch_3
} else {

}
}
{
__t3 = -1
}
end_branch_3:
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply4(go__go_2_0_16, gopurs_runtime.Int(end_1), gopurs_runtime.Int(start_0), gopurs_runtime.Int(__t3), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}))
}
end_branch_4:
return __t4
}

func Call_Data_List_partition(p_0_loop gopurs_runtime.Value, xs_1_loop *Constructor_Data_List_Types_Cons) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var xs_1 *Constructor_Data_List_Types_Cons = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldr"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(p_0, x_2).IntVal) != (0) {
__t0 = gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v_3, "no")))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, x_2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v_3, "yes"))})})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, x_2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v_3, "no"))})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v_3, "yes")))})
}
end_branch_0:
return __t0
})
}), gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)})
}

func Call_Data_List_null(v_0_loop *Constructor_Data_List_Types_Cons) bool {
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
var __t0 bool
{
if (v_0 == nil) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return __t0
}

func Call_Data_List_nubBy(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_17 gopurs_runtime.Value
go__go_1_0_17 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_3_loop_val)
var v2_4_loop gopurs_runtime.Value = v2_4_loop_val
go__go_1_0_17:
for {
if false { continue go__go_1_0_17 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Types_Cons = v1_3_loop
_ = v1_3
var v2_4 gopurs_runtime.Value = v2_4_loop
_ = v2_4
var __t3 *Constructor_Data_List_Types_Cons
{
if (v2_4.Type == 9 && v2_4.IntVal == 1358893437 && v2_4.UnsafePtr == nil) {
__t3 = v1_3
goto end_branch_3
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 1358893437 && v2_4.UnsafePtr != nil) {
// TAST (Let): v3_5_1 -> gopurs_runtime.Value
v3_5_1 := gopurs_runtime.Apply3(Get_Data_List_Internal_insertAndLookupBy(), p_0, (*Constructor_Data_List_Types_Cons)(v2_4.UnsafePtr).V0, v_2)
_ = v3_5_1
var __t2 *Constructor_Data_List_Types_Cons
{
if (gopurs_runtime.RecordGet(v3_5_1, "found").IntVal) != (0) {
v_2_loop = gopurs_runtime.RecordGet(v3_5_1, "result")
v1_3_loop = v1_3
v2_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v2_4.UnsafePtr).V1)}
continue go__go_1_0_17
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
v_2_loop = gopurs_runtime.RecordGet(v3_5_1, "result")
v1_3_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v2_4.UnsafePtr).V0, v1_3})})
v2_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v2_4.UnsafePtr).V1)}
continue go__go_1_0_17
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
})
// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
__local_var_2_4 := gopurs_runtime.Apply2(go__go_1_0_17, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_5_18 gopurs_runtime.Value
go__go_4_5_18 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_5_18:
for {
if false { continue go__go_4_5_18 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t6 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t6 = v_5
goto end_branch_6
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_5_18
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t6)}
}
}()
})
})
return gopurs_runtime.Apply2(go__go_4_5_18, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Apply(__local_var_2_4, x_3))
})
}

func Call_Data_List_nub(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
return Call_Data_List_nubBy(gopurs_runtime.Box(dictOrd_0.V1))
}

func Call_Data_List_mapMaybe(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_19 gopurs_runtime.Value
go__go_1_0_19 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_19:
for {
if false { continue go__go_1_0_19 }
var v_2 *Constructor_Data_List_Types_Cons = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t5 *Constructor_Data_List_Types_Cons
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
var go__go_4_1_20 gopurs_runtime.Value
go__go_4_1_20 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_20:
for {
if false { continue go__go_4_1_20 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = v_5
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_1_20
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t2)}
}
}()
})
})
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_1_20, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}))
goto end_branch_5
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
// TAST (Let): v2_4_3 -> *Constructor_Data_Maybe_Just
v2_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0))
_ = v2_4_3
var __t4 *Constructor_Data_List_Types_Cons
{
if (v2_4_3 == nil) {
v_2_loop = v_2
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}
continue go__go_1_0_19
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_4
} else {

}
}
{
if (v2_4_3 != nil) {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v2_4_3).V0, v_2})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}
continue go__go_1_0_19
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t5)}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_19, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}

func Call_Data_List_manyRec(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V0), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.Func(func(dictAlternative_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Plus1_3_1 -> gopurs_runtime.Value
Plus1_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Plus1"), gopurs_runtime.Value{})
_ = Plus1_3_1
// TAST (Let): Alt0_4_2 -> *Constructor_Control_Alt_Alt
Alt0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Plus1_3_1, "Alt0"), gopurs_runtime.Value{}))
_ = Alt0_4_2
// TAST (Let): Functor0_5_3 -> *Constructor_Data_Functor_Functor
Functor0_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Plus1_3_1, "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_3
// TAST (Let): Applicative0_6_4 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_4
// TAST (Let): pure_7_5 -> gopurs_runtime.Value
pure_7_5 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_5
return gopurs_runtime.Func(func(p_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(acc_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Alt0_4_2.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_3.V0), Get_Control_Monad_Rec_Class_Loop(), p_8), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_4.V1), gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, Get_Data_Unit_unit()})})), gopurs_runtime.Func(func(aa_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_7_5, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Control_Monad_Rec_Class_bifunctorStep(), "bimap"), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, v_11, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](acc_9)})}
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_12_6_21 gopurs_runtime.Value
go__go_12_6_21 = gopurs_runtime.Func(func(v_13_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_14_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_13_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_13_loop_val)
var v1_14_loop gopurs_runtime.Value = v1_14_loop_val
go__go_12_6_21:
for {
if false { continue go__go_12_6_21 }
var v_13 *Constructor_Data_List_Types_Cons = v_13_loop
_ = v_13
var v1_14 gopurs_runtime.Value = v1_14_loop
_ = v1_14
var __t7 *Constructor_Data_List_Types_Cons
{
if (v1_14.Type == 9 && v1_14.IntVal == 1358893437 && v1_14.UnsafePtr == nil) {
__t7 = v_13
goto end_branch_7
} else {

}
}
{
if (v1_14.Type == 9 && v1_14.IntVal == 1358893437 && v1_14.UnsafePtr != nil) {
v_13_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_14.UnsafePtr).V0, v_13})})
v1_14_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_14.UnsafePtr).V1)}
continue go__go_12_6_21
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t7)}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_12_6_21, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](acc_9))})))}
}), aa_10))
}))
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
})
})
}

func Call_Data_List_someRec(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec, dictAlternative_1_loop *Constructor_Control_Alternative_Alternative) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictAlternative_1 *Constructor_Control_Alternative_Alternative = dictAlternative_1_loop
_ = dictAlternative_1
// TAST (Let): Apply0_2_0 -> *Constructor_Control_Apply_Apply
Apply0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_1.V0), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_0
// TAST (Let): Functor0_3_1 -> *Constructor_Data_Functor_Functor
Functor0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_1.V1), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_1
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_2_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_1.V0), Get_Data_List_Types_Cons(), v_4), gopurs_runtime.Apply2(Call_Data_List_manyRec(dictMonadRec_0), gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(dictAlternative_1)}, v_4))
})
}

func Call_Data_List_some(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative) gopurs_runtime.Value {
var dictAlternative_0 *Constructor_Control_Alternative_Alternative = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V0), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V1), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(dictLazy_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_List_Types_Cons(), v_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_3, "defer"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Data_List_many(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy](dictLazy_3))}, v_4)
})))
})
})
}

func Call_Data_List_many(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative) gopurs_runtime.Value {
var dictAlternative_0 *Constructor_Control_Alternative_Alternative = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): Alt0_1_0 -> *Constructor_Control_Alt_Alt
Alt0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V1), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}))
_ = Alt0_1_0
// TAST (Let): Applicative0_2_1 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func(func(dictLazy_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Alt0_1_0.V1), gopurs_runtime.Apply2(Call_Data_List_some(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy](dictLazy_3))}, v_4), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}))
})
})
}

func Call_Data_List_last(v_0_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
last:
for {
if false { continue last }
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
var __t2 *Constructor_Data_Maybe_Just
{
if (v_0 != nil) {
var __t1 *Constructor_Data_Maybe_Just
{
var __t_tag_0 *Constructor_Data_List_Types_Cons = (v_0).V1
if (__t_tag_0 == nil) {
__t1 = &Constructor_Data_Maybe_Just{1, (v_0).V0}
goto end_branch_1
} else {

}
}
{
v_0_loop = (v_0).V1
continue last
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_2:
return __t2
}
}

func Call_Data_List_insertBy(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
insertBy:
for {
if false { continue insertBy }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_List_Types_Cons = v2_2_loop
_ = v2_2
var __t2 *Constructor_Data_List_Types_Cons
{
if (v2_2 == nil) {
__t2 = &Constructor_Data_List_Types_Cons{1, v1_1, (*Constructor_Data_List_Types_Cons)(nil)}
goto end_branch_2
} else {

}
}
{
if (v2_2 != nil) {
var __t1 *Constructor_Data_List_Types_Cons
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(v_0, v1_1, (v2_2).V0)
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = &Constructor_Data_List_Types_Cons{1, (v2_2).V0, Call_Data_List_insertBy(v_0, v1_1, (v2_2).V1)}
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_List_Types_Cons{1, v1_1, v2_2}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return __t2
}
}

func Call_Data_List_insertAt(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
insertAt:
for {
if false { continue insertAt }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_List_Types_Cons = v2_2_loop
_ = v2_2
var __t1 *Constructor_Data_Maybe_Just
{
if (v_0) == (0) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, v1_1, v2_2})}}
goto end_branch_1
} else {

}
}
{
if (v2_2 != nil) {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := (v2_2).V0
_ = __local_var_3_0
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(v3_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, __local_var_3_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v3_4)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_insertAt((v_0) - (1), v1_1, (v2_2).V1))})))})
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)})
}
}

func Call_Data_List_insert(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_List_insertBy(), gopurs_runtime.Box(dictOrd_0.V1))
}

func Call_Data_List_init(lst_0_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
var lst_0 *Constructor_Data_List_Types_Cons = lst_0_loop
_ = lst_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v_1, "init")))}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_unsnoc(lst_0))}))
}

func Call_Data_List_index(v_0_loop *Constructor_Data_List_Types_Cons, v1_1_loop int64) *Constructor_Data_Maybe_Just {
index:
for {
if false { continue index }
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
var v1_1 int64 = v1_1_loop
_ = v1_1
var __t1 *Constructor_Data_Maybe_Just
{
if (v_0 == nil) {
__t1 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_1
} else {

}
}
{
if (v_0 != nil) {
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1) == (0) {
__t0 = &Constructor_Data_Maybe_Just{1, (v_0).V0}
goto end_branch_0
} else {

}
}
{
v_0_loop = (v_0).V1
v1_1_loop = (v1_1) - (1)
continue index
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return __t1
}
}

func Call_Data_List_head(v_0_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 == nil) {
__t0 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_0
} else {

}
}
{
if (v_0 != nil) {
__t0 = &Constructor_Data_Maybe_Just{1, (v_0).V0}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return __t0
}

func Call_Data_List_transpose(v_0_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
transpose:
for {
if false { continue transpose }
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
var __t3 *Constructor_Data_List_Types_Cons
{
if (v_0 == nil) {
__t3 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_3
} else {

}
}
{
if (v_0 != nil) {
var __t2 *Constructor_Data_List_Types_Cons
{
var __t_tag_0 gopurs_runtime.Value = (v_0).V0
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 1358893437 && __t_tag_0.UnsafePtr == nil) {
v_0_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)})
continue transpose
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
var __t_tag_1 gopurs_runtime.Value = (v_0).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1358893437 && __t_tag_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)((v_0).V0.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(Call_Data_List_mapMaybe(Get_Data_List_head()), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)}))})}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_transpose(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)((v_0).V0.UnsafePtr).V1)}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(Call_Data_List_mapMaybe(Get_Data_List_tail()), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)})))})})})))})})})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t2)})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)})
}
}

func Call_Data_List_groupBy(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
groupBy:
for {
if false { continue groupBy }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_List_Types_Cons = v1_1_loop
_ = v1_1
var __t1 *Constructor_Data_List_Types_Cons
{
if (v1_1 == nil) {
__t1 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_1
} else {

}
}
{
if (v1_1 != nil) {
// TAST (Let): v2_2_0 -> gopurs_runtime.Value
var v2_2_0 gopurs_runtime.Value = Call_Data_List_span(gopurs_runtime.Apply(v_0, (v1_1).V0), (v1_1).V1)
__t1 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (v1_1).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v2_2_0, "init")))}})}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_groupBy(v_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v2_2_0, "rest"))))})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)})
}
}

func Call_Data_List_groupAllBy(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_List_groupBy(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> uint32
__local_var_3_1 := uint32(gopurs_runtime.Apply2(p_0, x_1, y_2).IntVal)
_ = __local_var_3_1
var __t2 bool
{
if (__local_var_3_1 == 1527465420) {
__t2 = false
goto end_branch_2
} else {

}
}
{
if (__local_var_3_1 == 380165415) {
__t2 = false
goto end_branch_2
} else {

}
}
{
if (__local_var_3_1 == 902936544) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
return gopurs_runtime.Bool(__t2)
})
}))
_ = __local_var_1_0
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := Call_Data_List_sortBy(p_0)
_ = __local_var_2_3
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply(__local_var_2_3, x_3))
})
}

func Call_Data_List_group(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_List_groupBy(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_List_groupAll(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := Call_Data_List_group(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.Apply(gopurs_runtime.Box(dictOrd_0.V0), gopurs_runtime.Value{})))
_ = __local_var_1_0
// TAST (Let): compare_2_2 -> gopurs_runtime.Value
compare_2_2 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_2_2
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(Call_Data_List_sortBy(compare_2_2), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_3))})))}
})
_ = __local_var_2_1
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply(__local_var_2_1, x_3))
})
}

func Call_Data_List_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V2), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}

func Call_Data_List_foldM(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
foldM:
for {
if false { continue foldM }
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 1358893437 && v2_5.UnsafePtr == nil) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), v1_4)
goto end_branch_3
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 1358893437 && v2_5.UnsafePtr != nil) {
// TAST (Let): __local_var_6_2 -> gopurs_runtime.Value
var __local_var_6_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v2_5.UnsafePtr).V1)}
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(v_3, v1_4, (*Constructor_Data_List_Types_Cons)(v2_5.UnsafePtr).V0), gopurs_runtime.Func(func(b_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Call_Data_List_foldM(dictMonad_0), v_3, b_prime_7, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](__local_var_6_2))})
}))
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
})
})
}
}

func Call_Data_List_findIndex(fn_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var go__go_1_0_22 gopurs_runtime.Value
go__go_1_0_22 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop int64 = v_2_loop_val.IntVal
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_22:
for {
if false { continue go__go_1_0_22 }
var v_2 int64 = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t2 *Constructor_Data_Maybe_Just
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
var __t1 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.Apply(fn_0, (*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(v_2)})})
goto end_branch_1
} else {

}
}
{
v_2_loop = (v_2) + (1)
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}
continue go__go_1_0_22
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_22, gopurs_runtime.Int(0))
}

func Call_Data_List_findLastIndex(fn_0_loop gopurs_runtime.Value, xs_1_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var xs_1 *Constructor_Data_List_Types_Cons = xs_1_loop
_ = xs_1
var go__go_2_1_23 gopurs_runtime.Value
go__go_2_1_23 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop int64 = v_3_loop_val.IntVal
var v1_4_loop gopurs_runtime.Value = v1_4_loop_val
go__go_2_1_23:
for {
if false { continue go__go_2_1_23 }
var v_3 int64 = v_3_loop
_ = v_3
var v1_4 gopurs_runtime.Value = v1_4_loop
_ = v1_4
var __t3 *Constructor_Data_Maybe_Just
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr != nil) {
var __t2 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.Apply(fn_0, (*Constructor_Data_List_Types_Cons)(v1_4.UnsafePtr).V0).IntVal) != (0) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(v_3)})})
goto end_branch_2
} else {

}
}
{
v_3_loop = (v_3) + (1)
v1_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_4.UnsafePtr).V1)}
continue go__go_2_1_23
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
var go__go_3_4_24 gopurs_runtime.Value
go__go_3_4_24 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_4_24:
for {
if false { continue go__go_3_4_24 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t5 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t5 = v_4
goto end_branch_5
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_4_24
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t5)}
}
}()
})
})
// TAST (Let): __local_var_2_0 -> *Constructor_Data_Maybe_Just
__local_var_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(go__go_2_1_23, gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_3_4_24, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)})))}))
_ = __local_var_2_0
var __t6 *Constructor_Data_Maybe_Just
{
if (__local_var_2_0 != nil) {
__t6 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(((gopurs_runtime.Apply(Get_Data_List_length(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}).IntVal) - (1)) - ((__local_var_2_0).V0.IntVal))}
goto end_branch_6
} else {

}
}
{
__t6 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_6:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t6)})
}

func Call_Data_List_filterM(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
filterM:
for {
if false { continue filterM }
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
goto end_branch_5
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr != nil) {
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := (*Constructor_Data_List_Types_Cons)(v1_4.UnsafePtr).V0
_ = __local_var_5_2
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
var __local_var_6_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_4.UnsafePtr).V1)}
__t5 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(v_3, __local_var_5_2), gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(Call_Data_List_filterM(dictMonad_0), v_3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](__local_var_6_3))}), gopurs_runtime.Func(func(xs_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_List_Types_Cons
{
if (b_7.IntVal) != (0) {
__t4 = &Constructor_Data_List_Types_Cons{1, __local_var_5_2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_prime_8)}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_prime_8)
}
end_branch_4:
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t4)})
}))
}))
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
})
}
}

func Call_Data_List_filter(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_25 gopurs_runtime.Value
go__go_1_0_25 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_25:
for {
if false { continue go__go_1_0_25 }
var v_2 *Constructor_Data_List_Types_Cons = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t4 *Constructor_Data_List_Types_Cons
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
var go__go_4_1_26 gopurs_runtime.Value
go__go_4_1_26 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_26:
for {
if false { continue go__go_4_1_26 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = v_5
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_1_26
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t2)}
}
}()
})
})
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_1_26, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}))
goto end_branch_4
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
var __t3 *Constructor_Data_List_Types_Cons
{
if (gopurs_runtime.Apply(p_0, (*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0).IntVal) != (0) {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0, v_2})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}
continue go__go_1_0_25
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
v_2_loop = v_2
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}
continue go__go_1_0_25
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t4)}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_25, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}

func Call_Data_List_intersectBy(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_List_Types_Cons, v2_2_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_List_Types_Cons = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_List_Types_Cons = v2_2_loop
_ = v2_2
var __t1 *Constructor_Data_List_Types_Cons
{
if (v1_1 == nil) {
__t1 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_1
} else {

}
}
{
if (v2_2 == nil) {
__t1 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(Call_Data_List_filter(gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupDisj1_4_0 -> gopurs_runtime.Value
semigroupDisj1_4_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_HeytingAlgebra_heytingAlgebraBoolean(), "disj"), v_4, v1_5)
})
}))
_ = semigroupDisj1_4_0
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldMap"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_4_0
}), gopurs_runtime.RecordGet(Get_Data_HeytingAlgebra_heytingAlgebraBoolean(), "ff")), gopurs_runtime.Apply(v_0, x_3), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}).IntVal) != (0))
})), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}))
}
end_branch_1:
return __t1
}

func Call_Data_List_intersect(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_List_intersectBy(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_List_nubByEq(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
nubByEq:
for {
if false { continue nubByEq }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_List_Types_Cons = v1_1_loop
_ = v1_1
var __t1 *Constructor_Data_List_Types_Cons
{
if (v1_1 == nil) {
__t1 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_1
} else {

}
}
{
if (v1_1 != nil) {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := (v1_1).V0
_ = __local_var_2_0
__t1 = &Constructor_Data_List_Types_Cons{1, __local_var_2_0, Call_Data_List_nubByEq(v_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(Call_Data_List_filter(gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(v_0, __local_var_2_0, y_3).IntVal) != (0)) != (true))
})), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v1_1).V1)})))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return __t1
}
}

func Call_Data_List_nubEq(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_List_nubByEq(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_List_eqPattern(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
// TAST (Let): eqList_1_0 -> *Constructor_Data_Eq_Eq
eqList_1_0 := &Constructor_Data_Eq_Eq{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Types_eq1List(), "eq1"), dictEq_0)}
_ = eqList_1_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(eqList_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](x_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](y_3))}).IntVal) != (0))
})
}))
}

func Call_Data_List_ordPattern(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): ordList_1_0 -> *Constructor_Data_Ord_Ord
ordList_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](gopurs_runtime.Apply(Get_Data_List_Types_ordList(), dictOrd_0))
_ = ordList_1_0
// TAST (Let): eqPattern1_2_1 -> gopurs_runtime.Value
eqPattern1_2_1 := Call_Data_List_eqPattern(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
_ = eqPattern1_2_1
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eqPattern1_2_1
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(ordList_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](y_4))}).IntVal)), UnsafePtr: nil}
})
}))
}

func Call_Data_List_elemLastIndex(dictEq_0_loop *Constructor_Data_Eq_Eq, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_Data_List_findLastIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), v_2, x_1).IntVal) != (0))
}))
}

func Call_Data_List_elemIndex(dictEq_0_loop *Constructor_Data_Eq_Eq, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var go__go_2_0_27 gopurs_runtime.Value
go__go_2_0_27 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop int64 = v_3_loop_val.IntVal
var v1_4_loop gopurs_runtime.Value = v1_4_loop_val
go__go_2_0_27:
for {
if false { continue go__go_2_0_27 }
var v_3 int64 = v_3_loop
_ = v_3
var v1_4 gopurs_runtime.Value = v1_4_loop
_ = v1_4
var __t2 *Constructor_Data_Maybe_Just
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr != nil) {
var __t1 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), (*Constructor_Data_List_Types_Cons)(v1_4.UnsafePtr).V0, x_1).IntVal) != (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(v_3)})})
goto end_branch_1
} else {

}
}
{
v_3_loop = (v_3) + (1)
v1_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_4.UnsafePtr).V1)}
continue go__go_2_0_27
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
}
}()
})
})
return gopurs_runtime.Apply(go__go_2_0_27, gopurs_runtime.Int(0))
}

func Call_Data_List_dropWhile(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_28 gopurs_runtime.Value
go__go_1_0_28 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2_loop_val)
go__go_1_0_28:
for {
if false { continue go__go_1_0_28 }
var v_2 *Constructor_Data_List_Types_Cons = v_2_loop
_ = v_2
var __t1 *Constructor_Data_List_Types_Cons
{
if ((v_2 != nil)) && ((gopurs_runtime.Apply(p_0, (v_2).V0).IntVal) != (0)) {
v_2_loop = (v_2).V1
continue go__go_1_0_28
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = v_2
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
return go__go_1_0_28
}

func Call_Data_List_dropEnd(n_0_loop int64, xs_1_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 *Constructor_Data_List_Types_Cons = xs_1_loop
_ = xs_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(Get_Data_List_take(), gopurs_runtime.Int((gopurs_runtime.Apply(Get_Data_List_length(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}).IntVal) - (n_0)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}))
}

func Call_Data_List_drop(v_0_loop int64, v1_1_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
drop:
for {
if false { continue drop }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_List_Types_Cons = v1_1_loop
_ = v1_1
var __t1 *Constructor_Data_List_Types_Cons
{
var __t0 bool
{
if (v_0) < (1) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
if __t0 {
__t1 = v1_1
goto end_branch_1
} else {

}
}
{
if (v1_1 == nil) {
__t1 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_1
} else {

}
}
{
if (v1_1 != nil) {
v_0_loop = (v_0) - (1)
v1_1_loop = (v1_1).V1
continue drop
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return __t1
}
}

func Call_Data_List_slice(start_0_loop int64, end_1_loop int64, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var start_0 int64 = start_0_loop
_ = start_0
var end_1 int64 = end_1_loop
_ = end_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(Get_Data_List_take(), gopurs_runtime.Int((end_1) - (start_0)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_drop(start_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_2)))})))}
}

func Call_Data_List_takeEnd(n_0_loop int64, xs_1_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 *Constructor_Data_List_Types_Cons = xs_1_loop
_ = xs_1
return Call_Data_List_drop((gopurs_runtime.Apply(Get_Data_List_length(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}).IntVal) - (n_0), xs_1)
}

func Call_Data_List_deleteBy(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
deleteBy:
for {
if false { continue deleteBy }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_List_Types_Cons = v2_2_loop
_ = v2_2
var __t1 *Constructor_Data_List_Types_Cons
{
if (v2_2 == nil) {
__t1 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_1
} else {

}
}
{
if (v2_2 != nil) {
var __t0 *Constructor_Data_List_Types_Cons
{
if (gopurs_runtime.Apply2(v_0, v1_1, (v2_2).V0).IntVal) != (0) {
__t0 = (v2_2).V1
goto end_branch_0
} else {

}
}
{
__t0 = &Constructor_Data_List_Types_Cons{1, (v2_2).V0, Call_Data_List_deleteBy(v_0, v1_1, (v2_2).V1)}
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return __t1
}
}

func Call_Data_List_unionBy(eq_0_loop gopurs_runtime.Value, xs_1_loop *Constructor_Data_List_Types_Cons, ys_2_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 *Constructor_Data_List_Types_Cons = xs_1_loop
_ = xs_1
var ys_2 *Constructor_Data_List_Types_Cons = ys_2_loop
_ = ys_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_deleteBy(eq_0, a_4, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_3)))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_nubByEq(eq_0, ys_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)})))}))
}

func Call_Data_List_union(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_List_unionBy(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_List_deleteAt(v_0_loop int64, v1_1_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
deleteAt:
for {
if false { continue deleteAt }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_List_Types_Cons = v1_1_loop
_ = v1_1
var __t2 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
var __t1 *Constructor_Data_Maybe_Just
{
if (v_0) == (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v1_1).V1)}})})
goto end_branch_1
} else {

}
}
{
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := (v1_1).V0
_ = __local_var_2_0
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, __local_var_2_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v2_3)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_deleteAt((v_0) - (1), (v1_1).V1))}))
}
end_branch_1:
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)})
goto end_branch_2
} else {

}
}
{
__t2 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)})
}
}

func Call_Data_List_delete(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_List_deleteBy(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_List_difference(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_deleteBy(gopurs_runtime.Box(dictEq_0.V0), a_2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_1)))}
})
}))
}

func Call_Data_List_concatMap(b_0_loop gopurs_runtime.Value, a_1_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 *Constructor_Data_List_Types_Cons = a_1_loop
_ = a_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_bindList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(a_1)}, b_0))
}

func Call_Data_List_concat(v_0_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_bindList(), "bind"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}, Get_Data_List_identity()))
}

func Call_Data_List_alterAt(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
alterAt:
for {
if false { continue alterAt }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_List_Types_Cons = v2_2_loop
_ = v2_2
var __t4 *Constructor_Data_Maybe_Just
{
if (v2_2 != nil) {
var __t3 *Constructor_Data_Maybe_Just
{
if (v_0) == (0) {
// TAST (Let): v3_3_1 -> *Constructor_Data_Maybe_Just
v3_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(v1_1, (v2_2).V0))
_ = v3_3_1
var __t2 *Constructor_Data_List_Types_Cons
{
if (v3_3_1 == nil) {
__t2 = (v2_2).V1
goto end_branch_2
} else {

}
}
{
if (v3_3_1 != nil) {
__t2 = &Constructor_Data_List_Types_Cons{1, (v3_3_1).V0, (v2_2).V1}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t2)}})})
goto end_branch_3
} else {

}
}
{
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := (v2_2).V0
_ = __local_var_3_0
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(v3_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, __local_var_3_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v3_4)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_alterAt((v_0) - (1), v1_1, (v2_2).V1))}))
}
end_branch_3:
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)})
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_4:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)})
}
}

func Call_Data_List_modifyAt(n_0_loop int64, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(Get_Data_List_alterAt(), gopurs_runtime.Int(n_0), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_1, x_2)})}
}))
}

func Call_Data_List_alterAt__3453373293(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_List_Types_Cons = v2_2_loop
_ = v2_2
var __t4 *Constructor_Data_Maybe_Just
{
if (v2_2 != nil) {
var __t3 *Constructor_Data_Maybe_Just
{
if (v_0) == (0) {
// TAST (Let): v3_3_1 -> *Constructor_Data_Maybe_Just
v3_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(v1_1, (v2_2).V0))
_ = v3_3_1
var __t2 *Constructor_Data_List_Types_Cons
{
if (v3_3_1 == nil) {
__t2 = (v2_2).V1
goto end_branch_2
} else {

}
}
{
if (v3_3_1 != nil) {
__t2 = &Constructor_Data_List_Types_Cons{1, (v3_3_1).V0, (v2_2).V1}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t2)}})})
goto end_branch_3
} else {

}
}
{
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := (v2_2).V0
_ = __local_var_3_0
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(v3_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, __local_var_3_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v3_4)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_alterAt((v_0) - (1), v1_1, (v2_2).V1))}))
}
end_branch_3:
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)})
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_4:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)})
}

func Call_Data_List_deleteAt__2845095501(v_0_loop int64, v1_1_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_List_Types_Cons = v1_1_loop
_ = v1_1
var __t2 *Constructor_Data_Maybe_Just
{
if (v1_1 != nil) {
var __t1 *Constructor_Data_Maybe_Just
{
if (v_0) == (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v1_1).V1)}})})
goto end_branch_1
} else {

}
}
{
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := (v1_1).V0
_ = __local_var_2_0
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, __local_var_2_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v2_3)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_deleteAt((v_0) - (1), (v1_1).V1))}))
}
end_branch_1:
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)})
goto end_branch_2
} else {

}
}
{
__t2 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)})
}

func Call_Data_List_deleteBy__697302515(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_List_Types_Cons = v2_2_loop
_ = v2_2
var __t1 *Constructor_Data_List_Types_Cons
{
if (v2_2 == nil) {
__t1 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_1
} else {

}
}
{
if (v2_2 != nil) {
var __t0 *Constructor_Data_List_Types_Cons
{
if (gopurs_runtime.Apply2(v_0, v1_1, (v2_2).V0).IntVal) != (0) {
__t0 = (v2_2).V1
goto end_branch_0
} else {

}
}
{
__t0 = &Constructor_Data_List_Types_Cons{1, (v2_2).V0, Call_Data_List_deleteBy(v_0, v1_1, (v2_2).V1)}
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return __t1
}

func Call_Data_List_drop__551729751(v_0_loop int64, v1_1_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_List_Types_Cons = v1_1_loop
_ = v1_1
var __t1 *Constructor_Data_List_Types_Cons
{
var __t0 bool
{
if (v_0) < (1) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
if __t0 {
__t1 = v1_1
goto end_branch_1
} else {

}
}
{
if (v1_1 == nil) {
__t1 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_1
} else {

}
}
{
if (v1_1 != nil) {
__t1 = Call_Data_List_drop((v_0) - (1), (v1_1).V1)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return __t1
}

func Call_Data_List_drop__1836090668(v_0_loop int64, v1_1_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_List_Types_Cons = v1_1_loop
_ = v1_1
var __t1 *Constructor_Data_List_Types_Cons
{
var __t0 bool
{
if (v_0) < (1) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
if __t0 {
__t1 = v1_1
goto end_branch_1
} else {

}
}
{
if (v1_1 == nil) {
__t1 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_1
} else {

}
}
{
if (v1_1 != nil) {
__t1 = Call_Data_List_drop((v_0) - (1), (v1_1).V1)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return __t1
}

func Call_Data_List_dropWhile__2352021032(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_29 gopurs_runtime.Value
go__go_1_0_29 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2_loop_val)
go__go_1_0_29:
for {
if false { continue go__go_1_0_29 }
var v_2 *Constructor_Data_List_Types_Cons = v_2_loop
_ = v_2
var __t1 *Constructor_Data_List_Types_Cons
{
if ((v_2 != nil)) && ((gopurs_runtime.Apply(p_0, (v_2).V0).IntVal) != (0)) {
v_2_loop = (v_2).V1
continue go__go_1_0_29
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = v_2
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
return go__go_1_0_29
}

func Call_Data_List_filter__2352021032(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_30 gopurs_runtime.Value
go__go_1_0_30 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_30:
for {
if false { continue go__go_1_0_30 }
var v_2 *Constructor_Data_List_Types_Cons = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t4 *Constructor_Data_List_Types_Cons
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
var go__go_4_1_31 gopurs_runtime.Value
go__go_4_1_31 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_31:
for {
if false { continue go__go_4_1_31 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = v_5
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_1_31
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t2)}
}
}()
})
})
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_1_31, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}))
goto end_branch_4
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
var __t3 *Constructor_Data_List_Types_Cons
{
if (gopurs_runtime.Apply(p_0, (*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0).IntVal) != (0) {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0, v_2})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}
continue go__go_1_0_30
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
v_2_loop = v_2
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}
continue go__go_1_0_30
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t4)}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_30, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}

func Call_Data_List_filter__1617261107(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_32 gopurs_runtime.Value
go__go_1_0_32 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_32:
for {
if false { continue go__go_1_0_32 }
var v_2 *Constructor_Data_List_Types_Cons = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t4 *Constructor_Data_List_Types_Cons
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
var go__go_4_1_33 gopurs_runtime.Value
go__go_4_1_33 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_33:
for {
if false { continue go__go_4_1_33 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = v_5
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_1_33
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t2)}
}
}()
})
})
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_1_33, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}))
goto end_branch_4
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
var __t3 *Constructor_Data_List_Types_Cons
{
if (gopurs_runtime.Apply(p_0, (*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0).IntVal) != (0) {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0, v_2})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}
continue go__go_1_0_32
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
v_2_loop = v_2
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}
continue go__go_1_0_32
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t4)}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_32, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}

func Call_Data_List_filterM__14771079(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
goto end_branch_5
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr != nil) {
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := (*Constructor_Data_List_Types_Cons)(v1_4.UnsafePtr).V0
_ = __local_var_5_2
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
var __local_var_6_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_4.UnsafePtr).V1)}
__t5 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(v_3, __local_var_5_2), gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(Call_Data_List_filterM(dictMonad_0), v_3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](__local_var_6_3))}), gopurs_runtime.Func(func(xs_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_List_Types_Cons
{
if (b_7.IntVal) != (0) {
__t4 = &Constructor_Data_List_Types_Cons{1, __local_var_5_2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_prime_8)}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_prime_8)
}
end_branch_4:
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t4)})
}))
}))
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
})
}

func Call_Data_List_findIndex__2366045378(fn_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var go__go_1_0_34 gopurs_runtime.Value
go__go_1_0_34 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop int64 = v_2_loop_val.IntVal
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_34:
for {
if false { continue go__go_1_0_34 }
var v_2 int64 = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t2 *Constructor_Data_Maybe_Just
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
var __t1 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.Apply(fn_0, (*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(v_2)})})
goto end_branch_1
} else {

}
}
{
v_2_loop = (v_2) + (1)
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}
continue go__go_1_0_34
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_34, gopurs_runtime.Int(0))
}

func Call_Data_List_findLastIndex__2366045378(fn_0_loop gopurs_runtime.Value, xs_1_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var xs_1 *Constructor_Data_List_Types_Cons = xs_1_loop
_ = xs_1
var go__go_2_1_35 gopurs_runtime.Value
go__go_2_1_35 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop int64 = v_3_loop_val.IntVal
var v1_4_loop gopurs_runtime.Value = v1_4_loop_val
go__go_2_1_35:
for {
if false { continue go__go_2_1_35 }
var v_3 int64 = v_3_loop
_ = v_3
var v1_4 gopurs_runtime.Value = v1_4_loop
_ = v1_4
var __t3 *Constructor_Data_Maybe_Just
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr != nil) {
var __t2 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.Apply(fn_0, (*Constructor_Data_List_Types_Cons)(v1_4.UnsafePtr).V0).IntVal) != (0) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(v_3)})})
goto end_branch_2
} else {

}
}
{
v_3_loop = (v_3) + (1)
v1_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_4.UnsafePtr).V1)}
continue go__go_2_1_35
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 1358893437 && v1_4.UnsafePtr == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
var go__go_3_4_36 gopurs_runtime.Value
go__go_3_4_36 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_4_36:
for {
if false { continue go__go_3_4_36 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t5 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t5 = v_4
goto end_branch_5
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_4_36
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t5)}
}
}()
})
})
// TAST (Let): __local_var_2_0 -> *Constructor_Data_Maybe_Just
__local_var_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(go__go_2_1_35, gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_3_4_36, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)})))}))
_ = __local_var_2_0
var __t6 *Constructor_Data_Maybe_Just
{
if (__local_var_2_0 != nil) {
__t6 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(((gopurs_runtime.Apply(Get_Data_List_length(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}).IntVal) - (1)) - ((__local_var_2_0).V0.IntVal))}
goto end_branch_6
} else {

}
}
{
__t6 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_6:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t6)})
}

func Call_Data_List_foldM__3577257629(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 1358893437 && v2_5.UnsafePtr == nil) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), v1_4)
goto end_branch_3
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 1358893437 && v2_5.UnsafePtr != nil) {
// TAST (Let): __local_var_6_2 -> gopurs_runtime.Value
var __local_var_6_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v2_5.UnsafePtr).V1)}
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(v_3, v1_4, (*Constructor_Data_List_Types_Cons)(v2_5.UnsafePtr).V0), gopurs_runtime.Func(func(b_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Call_Data_List_foldM(dictMonad_0), v_3, b_prime_7, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](__local_var_6_2))})
}))
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
})
})
}

func Call_Data_List_fromFoldable__614070391(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V2), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}

func Call_Data_List_groupAllBy__3934374991(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_List_groupBy(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> uint32
__local_var_3_1 := uint32(gopurs_runtime.Apply2(p_0, x_1, y_2).IntVal)
_ = __local_var_3_1
var __t2 bool
{
if (__local_var_3_1 == 1527465420) {
__t2 = false
goto end_branch_2
} else {

}
}
{
if (__local_var_3_1 == 380165415) {
__t2 = false
goto end_branch_2
} else {

}
}
{
if (__local_var_3_1 == 902936544) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
return gopurs_runtime.Bool(__t2)
})
}))
_ = __local_var_1_0
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := Call_Data_List_sortBy(p_0)
_ = __local_var_2_3
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply(__local_var_2_3, x_3))
})
}

func Call_Data_List_groupBy__2162447253(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_List_Types_Cons = v1_1_loop
_ = v1_1
var __t1 *Constructor_Data_List_Types_Cons
{
if (v1_1 == nil) {
__t1 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_1
} else {

}
}
{
if (v1_1 != nil) {
// TAST (Let): v2_2_0 -> gopurs_runtime.Value
var v2_2_0 gopurs_runtime.Value = Call_Data_List_span(gopurs_runtime.Apply(v_0, (v1_1).V0), (v1_1).V1)
__t1 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (v1_1).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v2_2_0, "init")))}})}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_groupBy(v_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v2_2_0, "rest"))))})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)})
}

func Call_Data_List_groupBy__1039549870(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_List_Types_Cons = v1_1_loop
_ = v1_1
var __t1 *Constructor_Data_List_Types_Cons
{
if (v1_1 == nil) {
__t1 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_1
} else {

}
}
{
if (v1_1 != nil) {
// TAST (Let): v2_2_0 -> gopurs_runtime.Value
var v2_2_0 gopurs_runtime.Value = Call_Data_List_span(gopurs_runtime.Apply(v_0, (v1_1).V0), (v1_1).V1)
__t1 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (v1_1).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v2_2_0, "init")))}})}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_groupBy(v_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v2_2_0, "rest"))))})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)})
}

func Call_Data_List_head__3729839663(v_0_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 == nil) {
__t0 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_0
} else {

}
}
{
if (v_0 != nil) {
__t0 = &Constructor_Data_Maybe_Just{1, (v_0).V0}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return __t0
}

func Call_Data_List_index__304299960(v_0_loop *Constructor_Data_List_Types_Cons, v1_1_loop int64) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
var v1_1 int64 = v1_1_loop
_ = v1_1
var __t1 *Constructor_Data_Maybe_Just
{
if (v_0 == nil) {
__t1 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_1
} else {

}
}
{
if (v_0 != nil) {
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_1) == (0) {
__t0 = &Constructor_Data_Maybe_Just{1, (v_0).V0}
goto end_branch_0
} else {

}
}
{
__t0 = Call_Data_List_index((v_0).V1, (v1_1) - (1))
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return __t1
}

func Call_Data_List_init__2496605985(lst_0_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
var lst_0 *Constructor_Data_List_Types_Cons = lst_0_loop
_ = lst_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v_1, "init")))}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_unsnoc(lst_0))}))
}

func Call_Data_List_insertAt__2634211748(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_List_Types_Cons = v2_2_loop
_ = v2_2
var __t1 *Constructor_Data_Maybe_Just
{
if (v_0) == (0) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, v1_1, v2_2})}}
goto end_branch_1
} else {

}
}
{
if (v2_2 != nil) {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := (v2_2).V0
_ = __local_var_3_0
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(v3_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, __local_var_3_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v3_4)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_insertAt((v_0) - (1), v1_1, (v2_2).V1))})))})
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)})
}

func Call_Data_List_insertBy__1738998985(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_List_Types_Cons = v2_2_loop
_ = v2_2
var __t2 *Constructor_Data_List_Types_Cons
{
if (v2_2 == nil) {
__t2 = &Constructor_Data_List_Types_Cons{1, v1_1, (*Constructor_Data_List_Types_Cons)(nil)}
goto end_branch_2
} else {

}
}
{
if (v2_2 != nil) {
var __t1 *Constructor_Data_List_Types_Cons
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(v_0, v1_1, (v2_2).V0)
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = &Constructor_Data_List_Types_Cons{1, (v2_2).V0, Call_Data_List_insertBy(v_0, v1_1, (v2_2).V1)}
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_List_Types_Cons{1, v1_1, v2_2}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return __t2
}

func Call_Data_List_intersectBy__588351261(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_List_Types_Cons, v2_2_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_List_Types_Cons = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_List_Types_Cons = v2_2_loop
_ = v2_2
var __t1 *Constructor_Data_List_Types_Cons
{
if (v1_1 == nil) {
__t1 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_1
} else {

}
}
{
if (v2_2 == nil) {
__t1 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(Call_Data_List_filter(gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupDisj1_4_0 -> gopurs_runtime.Value
semigroupDisj1_4_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_HeytingAlgebra_heytingAlgebraBoolean(), "disj"), v_4, v1_5)
})
}))
_ = semigroupDisj1_4_0
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldMap"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_4_0
}), gopurs_runtime.RecordGet(Get_Data_HeytingAlgebra_heytingAlgebraBoolean(), "ff")), gopurs_runtime.Apply(v_0, x_3), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}).IntVal) != (0))
})), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}))
}
end_branch_1:
return __t1
}

func Call_Data_List_intersectBy__1190504998(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_List_Types_Cons, v2_2_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_List_Types_Cons = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_List_Types_Cons = v2_2_loop
_ = v2_2
var __t1 *Constructor_Data_List_Types_Cons
{
if (v1_1 == nil) {
__t1 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_1
} else {

}
}
{
if (v2_2 == nil) {
__t1 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(Call_Data_List_filter(gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupDisj1_4_0 -> gopurs_runtime.Value
semigroupDisj1_4_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_HeytingAlgebra_heytingAlgebraBoolean(), "disj"), v_4, v1_5)
})
}))
_ = semigroupDisj1_4_0
return gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldMap"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_4_0
}), gopurs_runtime.RecordGet(Get_Data_HeytingAlgebra_heytingAlgebraBoolean(), "ff")), gopurs_runtime.Apply(v_0, x_3), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}).IntVal) != (0))
})), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}))
}
end_branch_1:
return __t1
}

func Call_Data_List_last__4043133652(v_0_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
var __t6 *Constructor_Data_Maybe_Just
{
if (v_0 != nil) {
var __t5 *Constructor_Data_Maybe_Just
{
var __t_tag_4 *Constructor_Data_List_Types_Cons = (v_0).V1
if (__t_tag_4 == nil) {
__t5 = &Constructor_Data_Maybe_Just{1, (v_0).V0}
goto end_branch_5
} else {

}
}
{
var __t3 *Constructor_Data_Maybe_Just
{
var __t_tag_0 *Constructor_Data_List_Types_Cons = (v_0).V1
if (__t_tag_0 != nil) {
var __t2 *Constructor_Data_Maybe_Just
{
var __t_tag_1 *Constructor_Data_List_Types_Cons = ((v_0).V1).V1
if (__t_tag_1 == nil) {
__t2 = &Constructor_Data_Maybe_Just{1, ((v_0).V1).V0}
goto end_branch_2
} else {

}
}
{
__t2 = Call_Data_List_last(((v_0).V1).V1)
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_3:
__t5 = __t3
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
__t6 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_6:
return __t6
}

func Call_Data_List_many__542682753(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative) gopurs_runtime.Value {
var dictAlternative_0 *Constructor_Control_Alternative_Alternative = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): Alt0_1_0 -> *Constructor_Control_Alt_Alt
Alt0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V1), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}))
_ = Alt0_1_0
// TAST (Let): Applicative0_2_1 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func(func(dictLazy_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Alt0_1_0.V1), gopurs_runtime.Apply2(Call_Data_List_some(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy](dictLazy_3))}, v_4), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}))
})
})
}

func Call_Data_List_manyRec__4046352885(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V0), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.Func(func(dictAlternative_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Plus1_3_1 -> gopurs_runtime.Value
Plus1_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Plus1"), gopurs_runtime.Value{})
_ = Plus1_3_1
// TAST (Let): Alt0_4_2 -> *Constructor_Control_Alt_Alt
Alt0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Plus1_3_1, "Alt0"), gopurs_runtime.Value{}))
_ = Alt0_4_2
// TAST (Let): Functor0_5_3 -> *Constructor_Data_Functor_Functor
Functor0_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Plus1_3_1, "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_3
// TAST (Let): Applicative0_6_4 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_4
// TAST (Let): pure_7_5 -> gopurs_runtime.Value
pure_7_5 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_5
return gopurs_runtime.Func(func(p_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(acc_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Alt0_4_2.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_3.V0), Get_Control_Monad_Rec_Class_Loop(), p_8), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_4.V1), gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, Get_Data_Unit_unit()})})), gopurs_runtime.Func(func(aa_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_7_5, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Control_Monad_Rec_Class_bifunctorStep(), "bimap"), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, v_11, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](acc_9)})}
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_12_6_37 gopurs_runtime.Value
go__go_12_6_37 = gopurs_runtime.Func(func(v_13_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_14_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_13_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_13_loop_val)
var v1_14_loop gopurs_runtime.Value = v1_14_loop_val
go__go_12_6_37:
for {
if false { continue go__go_12_6_37 }
var v_13 *Constructor_Data_List_Types_Cons = v_13_loop
_ = v_13
var v1_14 gopurs_runtime.Value = v1_14_loop
_ = v1_14
var __t7 *Constructor_Data_List_Types_Cons
{
if (v1_14.Type == 9 && v1_14.IntVal == 1358893437 && v1_14.UnsafePtr == nil) {
__t7 = v_13
goto end_branch_7
} else {

}
}
{
if (v1_14.Type == 9 && v1_14.IntVal == 1358893437 && v1_14.UnsafePtr != nil) {
v_13_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_14.UnsafePtr).V0, v_13})})
v1_14_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_14.UnsafePtr).V1)}
continue go__go_12_6_37
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t7)}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_12_6_37, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](acc_9))})))}
}), aa_10))
}))
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
})
})
}

func Call_Data_List_mapMaybe__3262563995(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_38 gopurs_runtime.Value
go__go_1_0_38 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_38:
for {
if false { continue go__go_1_0_38 }
var v_2 *Constructor_Data_List_Types_Cons = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t5 *Constructor_Data_List_Types_Cons
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
var go__go_4_1_39 gopurs_runtime.Value
go__go_4_1_39 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_39:
for {
if false { continue go__go_4_1_39 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = v_5
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_1_39
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t2)}
}
}()
})
})
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_1_39, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}))
goto end_branch_5
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
// TAST (Let): v2_4_3 -> *Constructor_Data_Maybe_Just
v2_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0))
_ = v2_4_3
var __t4 *Constructor_Data_List_Types_Cons
{
if (v2_4_3 == nil) {
v_2_loop = v_2
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}
continue go__go_1_0_38
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_4
} else {

}
}
{
if (v2_4_3 != nil) {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v2_4_3).V0, v_2})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}
continue go__go_1_0_38
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t5)}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_38, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}

func Call_Data_List_mapMaybe__1486753757(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_40 gopurs_runtime.Value
go__go_1_0_40 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_40:
for {
if false { continue go__go_1_0_40 }
var v_2 *Constructor_Data_List_Types_Cons = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t5 *Constructor_Data_List_Types_Cons
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
var go__go_4_1_41 gopurs_runtime.Value
go__go_4_1_41 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_41:
for {
if false { continue go__go_4_1_41 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = v_5
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_1_41
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t2)}
}
}()
})
})
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_1_41, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}))
goto end_branch_5
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
// TAST (Let): v2_4_3 -> *Constructor_Data_Maybe_Just
v2_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0))
_ = v2_4_3
var __t4 *Constructor_Data_List_Types_Cons
{
if (v2_4_3 == nil) {
v_2_loop = v_2
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}
continue go__go_1_0_40
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_4
} else {

}
}
{
if (v2_4_3 != nil) {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v2_4_3).V0, v_2})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}
continue go__go_1_0_40
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t5)}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_40, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}

func Call_Data_List_mapMaybe__1640531773(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_42 gopurs_runtime.Value
go__go_1_0_42 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_42:
for {
if false { continue go__go_1_0_42 }
var v_2 *Constructor_Data_List_Types_Cons = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t5 *Constructor_Data_List_Types_Cons
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
var go__go_4_1_43 gopurs_runtime.Value
go__go_4_1_43 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_43:
for {
if false { continue go__go_4_1_43 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = v_5
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_1_43
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t2)}
}
}()
})
})
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_1_43, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}))
goto end_branch_5
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
// TAST (Let): v2_4_3 -> *Constructor_Data_Maybe_Just
v2_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0))
_ = v2_4_3
var __t4 *Constructor_Data_List_Types_Cons
{
if (v2_4_3 == nil) {
v_2_loop = v_2
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}
continue go__go_1_0_42
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_4
} else {

}
}
{
if (v2_4_3 != nil) {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v2_4_3).V0))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)})})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}
continue go__go_1_0_42
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t5)}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_42, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}

func Call_Data_List_mapMaybe__748617661(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_44 gopurs_runtime.Value
go__go_1_0_44 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_44:
for {
if false { continue go__go_1_0_44 }
var v_2 *Constructor_Data_List_Types_Cons = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t5 *Constructor_Data_List_Types_Cons
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
var go__go_4_1_45 gopurs_runtime.Value
go__go_4_1_45 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_45:
for {
if false { continue go__go_4_1_45 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = v_5
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_1_45
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t2)}
}
}()
})
})
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_1_45, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}))
goto end_branch_5
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
// TAST (Let): v2_4_3 -> *Constructor_Data_Maybe_Just
v2_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0))}))
_ = v2_4_3
var __t4 *Constructor_Data_List_Types_Cons
{
if (v2_4_3 == nil) {
v_2_loop = v_2
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}
continue go__go_1_0_44
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_4
} else {

}
}
{
if (v2_4_3 != nil) {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v2_4_3).V0, v_2})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}
continue go__go_1_0_44
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t5)}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_44, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}

func Call_Data_List_mapMaybe__4251473821(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_46 gopurs_runtime.Value
go__go_1_0_46 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_46:
for {
if false { continue go__go_1_0_46 }
var v_2 *Constructor_Data_List_Types_Cons = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t5 *Constructor_Data_List_Types_Cons
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
var go__go_4_1_47 gopurs_runtime.Value
go__go_4_1_47 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_47:
for {
if false { continue go__go_4_1_47 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = v_5
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_1_47
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t2)}
}
}()
})
})
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_1_47, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}))
goto end_branch_5
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
// TAST (Let): v2_4_3 -> *Constructor_Data_Maybe_Just
v2_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0))}))
_ = v2_4_3
var __t4 *Constructor_Data_List_Types_Cons
{
if (v2_4_3 == nil) {
v_2_loop = v_2
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}
continue go__go_1_0_46
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_4
} else {

}
}
{
if (v2_4_3 != nil) {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v2_4_3).V0))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)})})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}
continue go__go_1_0_46
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t5)}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_46, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}

func Call_Data_List_mapMaybe__2491277821(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_48 gopurs_runtime.Value
go__go_1_0_48 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_48:
for {
if false { continue go__go_1_0_48 }
var v_2 *Constructor_Data_List_Types_Cons = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t5 *Constructor_Data_List_Types_Cons
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
var go__go_4_1_49 gopurs_runtime.Value
go__go_4_1_49 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_49:
for {
if false { continue go__go_4_1_49 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = v_5
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_1_49
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t2)}
}
}()
})
})
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_1_49, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}))
goto end_branch_5
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
// TAST (Let): v2_4_3 -> *Constructor_Data_Maybe_Just
v2_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0))}))
_ = v2_4_3
var __t4 *Constructor_Data_List_Types_Cons
{
if (v2_4_3 == nil) {
v_2_loop = v_2
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}
continue go__go_1_0_48
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_4
} else {

}
}
{
if (v2_4_3 != nil) {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v2_4_3).V0, v_2})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}
continue go__go_1_0_48
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t5)}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_48, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}

func Call_Data_List_modifyAt__1886983628(n_0_loop int64, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(Get_Data_List_alterAt(), gopurs_runtime.Int(n_0), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_1, x_2)})}
}))
}

func Call_Data_List_nubBy__2103943131(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_50 gopurs_runtime.Value
go__go_1_0_50 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_3_loop_val)
var v2_4_loop gopurs_runtime.Value = v2_4_loop_val
go__go_1_0_50:
for {
if false { continue go__go_1_0_50 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Types_Cons = v1_3_loop
_ = v1_3
var v2_4 gopurs_runtime.Value = v2_4_loop
_ = v2_4
var __t3 *Constructor_Data_List_Types_Cons
{
if (v2_4.Type == 9 && v2_4.IntVal == 1358893437 && v2_4.UnsafePtr == nil) {
__t3 = v1_3
goto end_branch_3
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 1358893437 && v2_4.UnsafePtr != nil) {
// TAST (Let): v3_5_1 -> gopurs_runtime.Value
v3_5_1 := gopurs_runtime.Apply3(Get_Data_List_Internal_insertAndLookupBy(), p_0, (*Constructor_Data_List_Types_Cons)(v2_4.UnsafePtr).V0, v_2)
_ = v3_5_1
var __t2 *Constructor_Data_List_Types_Cons
{
if (gopurs_runtime.RecordGet(v3_5_1, "found").IntVal) != (0) {
v_2_loop = gopurs_runtime.RecordGet(v3_5_1, "result")
v1_3_loop = v1_3
v2_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v2_4.UnsafePtr).V1)}
continue go__go_1_0_50
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
v_2_loop = gopurs_runtime.RecordGet(v3_5_1, "result")
v1_3_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v2_4.UnsafePtr).V0, v1_3})})
v2_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v2_4.UnsafePtr).V1)}
continue go__go_1_0_50
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
})
// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
__local_var_2_4 := gopurs_runtime.Apply2(go__go_1_0_50, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_5_51 gopurs_runtime.Value
go__go_4_5_51 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_5_51:
for {
if false { continue go__go_4_5_51 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t6 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t6 = v_5
goto end_branch_6
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_5_51
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t6)}
}
}()
})
})
return gopurs_runtime.Apply2(go__go_4_5_51, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Apply(__local_var_2_4, x_3))
})
}

func Call_Data_List_nubBy__1502591776(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_52 gopurs_runtime.Value
go__go_1_0_52 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_3_loop_val)
var v2_4_loop gopurs_runtime.Value = v2_4_loop_val
go__go_1_0_52:
for {
if false { continue go__go_1_0_52 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Types_Cons = v1_3_loop
_ = v1_3
var v2_4 gopurs_runtime.Value = v2_4_loop
_ = v2_4
var __t3 *Constructor_Data_List_Types_Cons
{
if (v2_4.Type == 9 && v2_4.IntVal == 1358893437 && v2_4.UnsafePtr == nil) {
__t3 = v1_3
goto end_branch_3
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 1358893437 && v2_4.UnsafePtr != nil) {
// TAST (Let): v3_5_1 -> gopurs_runtime.Value
v3_5_1 := gopurs_runtime.Apply3(Get_Data_List_Internal_insertAndLookupBy(), p_0, (*Constructor_Data_List_Types_Cons)(v2_4.UnsafePtr).V0, v_2)
_ = v3_5_1
var __t2 *Constructor_Data_List_Types_Cons
{
if (gopurs_runtime.RecordGet(v3_5_1, "found").IntVal) != (0) {
v_2_loop = gopurs_runtime.RecordGet(v3_5_1, "result")
v1_3_loop = v1_3
v2_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v2_4.UnsafePtr).V1)}
continue go__go_1_0_52
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
v_2_loop = gopurs_runtime.RecordGet(v3_5_1, "result")
v1_3_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v2_4.UnsafePtr).V0, v1_3})})
v2_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v2_4.UnsafePtr).V1)}
continue go__go_1_0_52
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
})
// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
__local_var_2_4 := gopurs_runtime.Apply2(go__go_1_0_52, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_5_53 gopurs_runtime.Value
go__go_4_5_53 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_5_53:
for {
if false { continue go__go_4_5_53 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t6 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t6 = v_5
goto end_branch_6
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_5_53
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t6)}
}
}()
})
})
return gopurs_runtime.Apply2(go__go_4_5_53, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Apply(__local_var_2_4, x_3))
})
}

func Call_Data_List_nubByEq__3956095361(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_List_Types_Cons = v1_1_loop
_ = v1_1
var __t1 *Constructor_Data_List_Types_Cons
{
if (v1_1 == nil) {
__t1 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_1
} else {

}
}
{
if (v1_1 != nil) {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := (v1_1).V0
_ = __local_var_2_0
__t1 = &Constructor_Data_List_Types_Cons{1, __local_var_2_0, Call_Data_List_nubByEq(v_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(Call_Data_List_filter(gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(v_0, __local_var_2_0, y_3).IntVal) != (0)) != (true))
})), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v1_1).V1)})))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return __t1
}

func Call_Data_List_nubByEq__3655321914(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_List_Types_Cons = v1_1_loop
_ = v1_1
var __t1 *Constructor_Data_List_Types_Cons
{
if (v1_1 == nil) {
__t1 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_1
} else {

}
}
{
if (v1_1 != nil) {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := (v1_1).V0
_ = __local_var_2_0
__t1 = &Constructor_Data_List_Types_Cons{1, __local_var_2_0, Call_Data_List_nubByEq(v_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(Call_Data_List_filter(gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(v_0, __local_var_2_0, y_3).IntVal) != (0)) != (true))
})), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v1_1).V1)})))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return __t1
}

func Call_Data_List_null__74357383(v_0_loop *Constructor_Data_List_Types_Cons) bool {
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
var __t0 bool
{
if (v_0 == nil) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return __t0
}

func Call_Data_List_null__2437342685(v_0_loop *Constructor_Data_List_Types_Cons) bool {
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
var __t0 bool
{
if (v_0 == nil) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return __t0
}

func Call_Data_List_partition__1623965204(p_0_loop gopurs_runtime.Value, xs_1_loop *Constructor_Data_List_Types_Cons) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var xs_1 *Constructor_Data_List_Types_Cons = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldr"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(p_0, x_2).IntVal) != (0) {
__t0 = gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v_3, "no")))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, x_2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v_3, "yes"))})})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, x_2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v_3, "no"))})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v_3, "yes")))})
}
end_branch_0:
return __t0
})
}), gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)})
}

func Call_Data_List_singleton__2450819477(a_0_loop int64) *Constructor_Data_List_Types_Cons {
var a_0 int64 = a_0_loop
_ = a_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Int(a_0), (*Constructor_Data_List_Types_Cons)(nil)})})
}

func Call_Data_List_singleton__707062261(a_0_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return &Constructor_Data_List_Types_Cons{1, a_0, (*Constructor_Data_List_Types_Cons)(nil)}
}

func Call_Data_List_singleton__3932757557(a_0_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var a_0 *Constructor_Data_List_Types_Cons = a_0_loop
_ = a_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(a_0)}, (*Constructor_Data_List_Types_Cons)(nil)})})
}

func Call_Data_List_snoc__4290067657(xs_0_loop *Constructor_Data_List_Types_Cons, x_1_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons {
var xs_0 *Constructor_Data_List_Types_Cons = xs_0_loop
_ = xs_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldr"), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, x_1, (*Constructor_Data_List_Types_Cons)(nil)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_0)}))
}

func Call_Data_List_some__542682753(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative) gopurs_runtime.Value {
var dictAlternative_0 *Constructor_Control_Alternative_Alternative = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V0), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V1), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(dictLazy_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_List_Types_Cons(), v_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_3, "defer"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Data_List_many(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy](dictLazy_3))}, v_4)
})))
})
})
}

func Call_Data_List_sortBy__2103943131(cmp_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
var merge_1_0_61 gopurs_runtime.Value
_ = merge_1_0_61
merge_1_0_61 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_List_Types_Cons
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr != nil) {
var __t4 *Constructor_Data_List_Types_Cons
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
var __t3 *Constructor_Data_List_Types_Cons
{
// TAST (Let): __local_var_4_1 -> uint32
__local_var_4_1 := uint32(gopurs_runtime.Apply2(cmp_0, (*Constructor_Data_List_Types_Cons)(v_2.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0).IntVal)
_ = __local_var_4_1
var __t2 bool
{
if (__local_var_4_1 == 1527465420) {
__t2 = false
goto end_branch_2
} else {

}
}
{
if (__local_var_4_1 == 380165415) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
if __t2 {
__t3 = &Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(merge_1_0_61, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}))}
goto end_branch_3
} else {

}
}
{
__t3 = &Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(merge_1_0_61, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_2.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_3))}))}
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_3)
goto end_branch_5
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2)
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t5)}
})
})
var mergePairs_2_6_62 gopurs_runtime.Value
_ = mergePairs_2_6_62
mergePairs_2_6_62 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 *Constructor_Data_List_Types_Cons
{
var __t_and_8 bool = false
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {

var __t_tag_7 *Constructor_Data_List_Types_Cons = (*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V1
__t_and_8 = (__t_tag_7 != nil)
}
if __t_and_8 {
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(merge_1_0_61, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V0))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V1).V0))})))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(mergePairs_2_6_62, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V1).V1)})))})})})
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_3)
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t9)}
})
var mergeAll_3_10_63 gopurs_runtime.Value
mergeAll_3_10_63 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
mergeAll_3_10_63:
for {
if false { continue mergeAll_3_10_63 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var __t13 *Constructor_Data_List_Types_Cons
{
var __t_and_12 bool = false
if (v_4 != nil) {

var __t_tag_11 *Constructor_Data_List_Types_Cons = (v_4).V1
__t_and_12 = (__t_tag_11 == nil)
}
if __t_and_12 {
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_4).V0)
goto end_branch_13
} else {

}
}
{
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(mergePairs_2_6_62, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}))
continue mergeAll_3_10_63
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
}
end_branch_13:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t13)}
}
}()
})
var sequences_4_14_64 gopurs_runtime.Value
_ = sequences_4_14_64
var descending_4_15_65 gopurs_runtime.Value
_ = descending_4_15_65
var ascending_4_16_66 gopurs_runtime.Value
_ = ascending_4_16_66
sequences_4_14_64 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t23 *Constructor_Data_List_Types_Cons
{
var __t_and_18 bool = false
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {

var __t_tag_17 *Constructor_Data_List_Types_Cons = (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1
__t_and_18 = (__t_tag_17 != nil)
}
if __t_and_18 {
var __t22 *Constructor_Data_List_Types_Cons
{
// TAST (Let): __local_var_6_20 -> uint32
__local_var_6_20 := uint32(gopurs_runtime.Apply2(cmp_0, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0, ((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1).V0).IntVal)
_ = __local_var_6_20
var __t21 bool
{
if (__local_var_6_20 == 1527465420) {
__t21 = false
goto end_branch_21
} else {

}
}
{
if (__local_var_6_20 == 380165415) {
__t21 = true
goto end_branch_21
} else {

}
}
{
__t21 = false
}
end_branch_21:
if __t21 {
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(descending_4_15_65, ((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(nil)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1).V1)}))
goto end_branch_22
} else {

}
}
{
// TAST (Let): __local_var_6_19 -> gopurs_runtime.Value
__local_var_6_19 := (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0
_ = __local_var_6_19
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(ascending_4_16_66, ((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1).V0, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, __local_var_6_19, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_7)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1).V1)}))
}
end_branch_22:
__t23 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t22)})
goto end_branch_23
} else {

}
}
{
__t23 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5))}, (*Constructor_Data_List_Types_Cons)(nil)}
}
end_branch_23:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t23)}
})
descending_4_15_65 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t27 *Constructor_Data_List_Types_Cons
{
var __t_and_26 bool = false
if (v2_7.Type == 9 && v2_7.IntVal == 1358893437 && v2_7.UnsafePtr != nil) {

// TAST (Let): __local_var_8_24 -> uint32
__local_var_8_24 := uint32(gopurs_runtime.Apply2(cmp_0, v_5, (*Constructor_Data_List_Types_Cons)(v2_7.UnsafePtr).V0).IntVal)
_ = __local_var_8_24
var __t25 bool
{
if (__local_var_8_24 == 1527465420) {
__t25 = false
goto end_branch_25
} else {

}
}
{
if (__local_var_8_24 == 380165415) {
__t25 = true
goto end_branch_25
} else {

}
}
{
__t25 = false
}
end_branch_25:
__t_and_26 = __t25
}
if __t_and_26 {
__t27 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(descending_4_15_65, (*Constructor_Data_List_Types_Cons)(v2_7.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, v_5, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_6)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v2_7.UnsafePtr).V1)})))})
goto end_branch_27
} else {

}
}
{
__t27 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, v_5, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_6)})}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(sequences_4_14_64, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v2_7))})))})}
}
end_branch_27:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t27)}
})
})
})
ascending_4_16_66 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t28 *Constructor_Data_List_Types_Cons
{
if ((v2_7.Type == 9 && v2_7.IntVal == 1358893437 && v2_7.UnsafePtr != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Ordering_eqOrdering(), "eq"), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(cmp_0, v_5, (*Constructor_Data_List_Types_Cons)(v2_7.UnsafePtr).V0).IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}).IntVal) != (0)) != (true)) {
__t28 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(ascending_4_16_66, (*Constructor_Data_List_Types_Cons)(v2_7.UnsafePtr).V0, gopurs_runtime.Func(func(ys_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(v1_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, v_5, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_8)})})))}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v2_7.UnsafePtr).V1)})))})
goto end_branch_28
} else {

}
}
{
__t28 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(v1_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, v_5, (*Constructor_Data_List_Types_Cons)(nil)})})))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(sequences_4_14_64, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v2_7))})))})}
}
end_branch_28:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t28)}
})
})
})
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(mergeAll_3_10_63, gopurs_runtime.Apply(sequences_4_14_64, x_5))
})
}

func Call_Data_List_sortBy__1502591776(cmp_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
var merge_1_0_67 gopurs_runtime.Value
_ = merge_1_0_67
merge_1_0_67 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_List_Types_Cons
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr != nil) {
var __t4 *Constructor_Data_List_Types_Cons
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
var __t3 *Constructor_Data_List_Types_Cons
{
// TAST (Let): __local_var_4_1 -> uint32
__local_var_4_1 := uint32(gopurs_runtime.Apply2(cmp_0, (*Constructor_Data_List_Types_Cons)(v_2.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0).IntVal)
_ = __local_var_4_1
var __t2 bool
{
if (__local_var_4_1 == 1527465420) {
__t2 = false
goto end_branch_2
} else {

}
}
{
if (__local_var_4_1 == 380165415) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
if __t2 {
__t3 = &Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(merge_1_0_67, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}))}
goto end_branch_3
} else {

}
}
{
__t3 = &Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_2.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(merge_1_0_67, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_2.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_3))}))}
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_3)
goto end_branch_5
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2)
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t5)}
})
})
var mergePairs_2_6_68 gopurs_runtime.Value
_ = mergePairs_2_6_68
mergePairs_2_6_68 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 *Constructor_Data_List_Types_Cons
{
var __t_and_8 bool = false
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {

var __t_tag_7 *Constructor_Data_List_Types_Cons = (*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V1
__t_and_8 = (__t_tag_7 != nil)
}
if __t_and_8 {
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(merge_1_0_67, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V0))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V1).V0))})))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(mergePairs_2_6_68, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V1).V1)})))})})})
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_3)
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t9)}
})
var mergeAll_3_10_69 gopurs_runtime.Value
mergeAll_3_10_69 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
mergeAll_3_10_69:
for {
if false { continue mergeAll_3_10_69 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var __t13 *Constructor_Data_List_Types_Cons
{
var __t_and_12 bool = false
if (v_4 != nil) {

var __t_tag_11 *Constructor_Data_List_Types_Cons = (v_4).V1
__t_and_12 = (__t_tag_11 == nil)
}
if __t_and_12 {
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons]((v_4).V0)
goto end_branch_13
} else {

}
}
{
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(mergePairs_2_6_68, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}))
continue mergeAll_3_10_69
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
}
end_branch_13:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t13)}
}
}()
})
var sequences_4_14_70 gopurs_runtime.Value
_ = sequences_4_14_70
var descending_4_15_71 gopurs_runtime.Value
_ = descending_4_15_71
var ascending_4_16_72 gopurs_runtime.Value
_ = ascending_4_16_72
sequences_4_14_70 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t23 *Constructor_Data_List_Types_Cons
{
var __t_and_18 bool = false
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {

var __t_tag_17 *Constructor_Data_List_Types_Cons = (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1
__t_and_18 = (__t_tag_17 != nil)
}
if __t_and_18 {
var __t22 *Constructor_Data_List_Types_Cons
{
// TAST (Let): __local_var_6_20 -> uint32
__local_var_6_20 := uint32(gopurs_runtime.Apply2(cmp_0, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0, ((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1).V0).IntVal)
_ = __local_var_6_20
var __t21 bool
{
if (__local_var_6_20 == 1527465420) {
__t21 = false
goto end_branch_21
} else {

}
}
{
if (__local_var_6_20 == 380165415) {
__t21 = true
goto end_branch_21
} else {

}
}
{
__t21 = false
}
end_branch_21:
if __t21 {
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(descending_4_15_71, ((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(nil)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1).V1)}))
goto end_branch_22
} else {

}
}
{
// TAST (Let): __local_var_6_19 -> gopurs_runtime.Value
__local_var_6_19 := (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0
_ = __local_var_6_19
__t22 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(ascending_4_16_72, ((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1).V0, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, __local_var_6_19, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_7)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1).V1)}))
}
end_branch_22:
__t23 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t22)})
goto end_branch_23
} else {

}
}
{
__t23 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5))}, (*Constructor_Data_List_Types_Cons)(nil)}
}
end_branch_23:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t23)}
})
descending_4_15_71 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t27 *Constructor_Data_List_Types_Cons
{
var __t_and_26 bool = false
if (v2_7.Type == 9 && v2_7.IntVal == 1358893437 && v2_7.UnsafePtr != nil) {

// TAST (Let): __local_var_8_24 -> uint32
__local_var_8_24 := uint32(gopurs_runtime.Apply2(cmp_0, v_5, (*Constructor_Data_List_Types_Cons)(v2_7.UnsafePtr).V0).IntVal)
_ = __local_var_8_24
var __t25 bool
{
if (__local_var_8_24 == 1527465420) {
__t25 = false
goto end_branch_25
} else {

}
}
{
if (__local_var_8_24 == 380165415) {
__t25 = true
goto end_branch_25
} else {

}
}
{
__t25 = false
}
end_branch_25:
__t_and_26 = __t25
}
if __t_and_26 {
__t27 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(descending_4_15_71, (*Constructor_Data_List_Types_Cons)(v2_7.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, v_5, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_6)})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v2_7.UnsafePtr).V1)})))})
goto end_branch_27
} else {

}
}
{
__t27 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, v_5, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_6)})}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(sequences_4_14_70, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v2_7))})))})}
}
end_branch_27:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t27)}
})
})
})
ascending_4_16_72 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t28 *Constructor_Data_List_Types_Cons
{
if ((v2_7.Type == 9 && v2_7.IntVal == 1358893437 && v2_7.UnsafePtr != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Ordering_eqOrdering(), "eq"), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(cmp_0, v_5, (*Constructor_Data_List_Types_Cons)(v2_7.UnsafePtr).V0).IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}).IntVal) != (0)) != (true)) {
__t28 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(ascending_4_16_72, (*Constructor_Data_List_Types_Cons)(v2_7.UnsafePtr).V0, gopurs_runtime.Func(func(ys_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(v1_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, v_5, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_8)})})))}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v2_7.UnsafePtr).V1)})))})
goto end_branch_28
} else {

}
}
{
__t28 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(v1_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, v_5, (*Constructor_Data_List_Types_Cons)(nil)})})))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(sequences_4_14_70, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v2_7))})))})}
}
end_branch_28:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t28)}
})
})
})
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(mergeAll_3_10_69, gopurs_runtime.Apply(sequences_4_14_70, x_5))
})
}

func Call_Data_List_span__1918198736(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_List_Types_Cons) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_List_Types_Cons = v1_1_loop
_ = v1_1
var __t1 gopurs_runtime.Value
{
if ((v1_1 != nil)) && ((gopurs_runtime.Apply(v_0, (v1_1).V0).IntVal) != (0)) {
// TAST (Let): v2_2_0 -> gopurs_runtime.Value
var v2_2_0 gopurs_runtime.Value = Call_Data_List_span(v_0, (v1_1).V1)
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v1_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v2_2_0, "init"))})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v2_2_0, "rest")))})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)})
}
end_branch_1:
return __t1
}

func Call_Data_List_span__799093643(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_List_Types_Cons) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_List_Types_Cons = v1_1_loop
_ = v1_1
var __t1 gopurs_runtime.Value
{
if ((v1_1 != nil)) && ((gopurs_runtime.Apply(v_0, (v1_1).V0).IntVal) != (0)) {
// TAST (Let): v2_2_0 -> gopurs_runtime.Value
var v2_2_0 gopurs_runtime.Value = Call_Data_List_span(v_0, (v1_1).V1)
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (v1_1).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v2_2_0, "init"))})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v2_2_0, "rest")))})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)})
}
end_branch_1:
return __t1
}

func Call_Data_List_span__2133741451(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_List_Types_Cons) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_List_Types_Cons = v1_1_loop
_ = v1_1
var __t1 gopurs_runtime.Value
{
if ((v1_1 != nil)) && ((gopurs_runtime.Apply(v_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple]((v1_1).V0))}).IntVal) != (0)) {
// TAST (Let): v2_2_0 -> gopurs_runtime.Value
var v2_2_0 gopurs_runtime.Value = Call_Data_List_span(v_0, (v1_1).V1)
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple]((v1_1).V0))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v2_2_0, "init")))})})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v2_2_0, "rest")))})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)})
}
end_branch_1:
return __t1
}

func Call_Data_List_tail__1771843450(v_0_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 == nil) {
__t0 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_0
} else {

}
}
{
if (v_0 != nil) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)}}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)})
}

func Call_Data_List_tails__3932757557(v_0_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
var __t0 *Constructor_Data_List_Types_Cons
{
if (v_0 == nil) {
__t0 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, (*Constructor_Data_List_Types_Cons)(nil)}
goto end_branch_0
} else {

}
}
{
if (v_0 != nil) {
__t0 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_tails((v_0).V1))})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t0)})
}

func Call_Data_List_takeWhile__2352021032(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_79 gopurs_runtime.Value
go__go_1_0_79 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_79:
for {
if false { continue go__go_1_0_79 }
var v_2 *Constructor_Data_List_Types_Cons = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t3 *Constructor_Data_List_Types_Cons
{
if ((v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil)) && ((gopurs_runtime.Apply(p_0, (*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0).IntVal) != (0)) {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0, v_2})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}
continue go__go_1_0_79
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
var go__go_4_1_80 gopurs_runtime.Value
go__go_4_1_80 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_80:
for {
if false { continue go__go_4_1_80 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t2 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t2 = v_5
goto end_branch_2
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_1_80
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t2)}
}
}()
})
})
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_1_80, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}))
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_79, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}

func Call_Data_List_transpose__682228544(v_0_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
var __t3 *Constructor_Data_List_Types_Cons
{
if (v_0 == nil) {
__t3 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_3
} else {

}
}
{
if (v_0 != nil) {
var __t2 *Constructor_Data_List_Types_Cons
{
var __t_tag_0 gopurs_runtime.Value = (v_0).V0
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 1358893437 && __t_tag_0.UnsafePtr == nil) {
__t2 = Call_Data_List_transpose(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)}))
goto end_branch_2
} else {

}
}
{
var __t_tag_1 gopurs_runtime.Value = (v_0).V0
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1358893437 && __t_tag_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)((v_0).V0.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(Call_Data_List_mapMaybe(Get_Data_List_head()), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)}))})}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_transpose(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)((v_0).V0.UnsafePtr).V1)}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(Call_Data_List_mapMaybe(Get_Data_List_tail()), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)})))})})})))})})})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t2)})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)})
}

func Call_Data_List_uncons__3009258782(v_0_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
var v_0 *Constructor_Data_List_Types_Cons = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Maybe_Just
{
if (v_0 == nil) {
__t0 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_0
} else {

}
}
{
if (v_0 != nil) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("head", "tail", (v_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)})
}

func Call_Data_List_unionBy__588351261(eq_0_loop gopurs_runtime.Value, xs_1_loop *Constructor_Data_List_Types_Cons, ys_2_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 *Constructor_Data_List_Types_Cons = xs_1_loop
_ = xs_1
var ys_2 *Constructor_Data_List_Types_Cons = ys_2_loop
_ = ys_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_deleteBy(eq_0, a_4, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_3)))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_nubByEq(eq_0, ys_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)})))}))
}

func Call_Data_List_unionBy__1190504998(eq_0_loop gopurs_runtime.Value, xs_1_loop *Constructor_Data_List_Types_Cons, ys_2_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 *Constructor_Data_List_Types_Cons = xs_1_loop
_ = xs_1
var ys_2 *Constructor_Data_List_Types_Cons = ys_2_loop
_ = ys_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_deleteBy(eq_0, a_4, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_3)))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_nubByEq(eq_0, ys_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)})))}))
}

func Call_Data_List_unsnoc__2942606998(lst_0_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
var lst_0 *Constructor_Data_List_Types_Cons = lst_0_loop
_ = lst_0
var go__go_1_0_81 gopurs_runtime.Value
go__go_1_0_81 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v1_3_loop_val)
go__go_1_0_81:
for {
if false { continue go__go_1_0_81 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Types_Cons = v1_3_loop
_ = v1_3
var __t3 *Constructor_Data_Maybe_Just
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr == nil) {
__t3 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 1358893437 && v_2.UnsafePtr != nil) {
var __t2 *Constructor_Data_Maybe_Just
{
var __t_tag_1 *Constructor_Data_List_Types_Cons = (*Constructor_Data_List_Types_Cons)(v_2.UnsafePtr).V1
if (__t_tag_1 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("last", "revInit", (*Constructor_Data_List_Types_Cons)(v_2.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)})})})
goto end_branch_2
} else {

}
}
{
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_2.UnsafePtr).V1)}
v1_3_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_2.UnsafePtr).V0, v1_3})})
continue go__go_1_0_81
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
}
end_branch_2:
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(h_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_4_82 gopurs_runtime.Value
go__go_3_4_82 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_4_82:
for {
if false { continue go__go_3_4_82 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t5 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t5 = v_4
goto end_branch_5
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_4_82
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t5)}
}
}()
})
})
return gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_3_4_82, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(h_2, "revInit")))})))}, gopurs_runtime.RecordGet(h_2, "last"))
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(go__go_1_0_81, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(lst_0)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}))
}

func Call_Data_List_updateAt__2634211748(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_Maybe_Just {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_List_Types_Cons = v2_2_loop
_ = v2_2
var __t2 *Constructor_Data_Maybe_Just
{
if (v2_2 != nil) {
var __t1 *Constructor_Data_Maybe_Just
{
if (v_0) == (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, v1_1, (v2_2).V1})}})})
goto end_branch_1
} else {

}
}
{
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := (v2_2).V0
_ = __local_var_3_0
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(v3_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, __local_var_3_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v3_4)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_updateAt((v_0) - (1), v1_1, (v2_2).V1))}))
}
end_branch_1:
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)})
goto end_branch_2
} else {

}
}
{
__t2 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)})
}

func Call_Data_List_zipWith__884793877(f_0_loop gopurs_runtime.Value, xs_1_loop *Constructor_Data_List_Types_Cons, ys_2_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 *Constructor_Data_List_Types_Cons = xs_1_loop
_ = xs_1
var ys_2 *Constructor_Data_List_Types_Cons = ys_2_loop
_ = ys_2
var go__go_3_0_83 gopurs_runtime.Value
go__go_3_0_83 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
var v2_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v2_6_loop_val)
go__go_3_0_83:
for {
if false { continue go__go_3_0_83 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var v2_6 *Constructor_Data_List_Types_Cons = v2_6_loop
_ = v2_6
var __t1 *Constructor_Data_List_Types_Cons
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t1 = v2_6
goto end_branch_1
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t1 = v2_6
goto end_branch_1
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil)) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
v2_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Apply2(f_0, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0), v2_6})})
continue go__go_3_0_83
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
})
})
var go__go_4_2_84 gopurs_runtime.Value
go__go_4_2_84 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_2_84:
for {
if false { continue go__go_4_2_84 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t3 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_2_84
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_2_84, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(go__go_3_0_83, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(ys_2)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}))
}

func Call_Data_List_zipWith__4203240021(f_0_loop gopurs_runtime.Value, xs_1_loop *Constructor_Data_List_Types_Cons, ys_2_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 *Constructor_Data_List_Types_Cons = xs_1_loop
_ = xs_1
var ys_2 *Constructor_Data_List_Types_Cons = ys_2_loop
_ = ys_2
var go__go_3_0_85 gopurs_runtime.Value
go__go_3_0_85 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
var v2_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v2_6_loop_val)
go__go_3_0_85:
for {
if false { continue go__go_3_0_85 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var v2_6 *Constructor_Data_List_Types_Cons = v2_6_loop
_ = v2_6
var __t1 *Constructor_Data_List_Types_Cons
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t1 = v2_6
goto end_branch_1
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t1 = v2_6
goto end_branch_1
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil)) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
v2_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(f_0, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0)))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_6)})})})
continue go__go_3_0_85
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
})
})
var go__go_4_2_86 gopurs_runtime.Value
go__go_4_2_86 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_2_86:
for {
if false { continue go__go_4_2_86 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t3 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_2_86
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_2_86, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(go__go_3_0_85, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(ys_2)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}))
}

func Call_Data_List_zipWith__3856182069(f_0_loop gopurs_runtime.Value, xs_1_loop *Constructor_Data_List_Types_Cons, ys_2_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 *Constructor_Data_List_Types_Cons = xs_1_loop
_ = xs_1
var ys_2 *Constructor_Data_List_Types_Cons = ys_2_loop
_ = ys_2
var go__go_3_0_87 gopurs_runtime.Value
go__go_3_0_87 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
var v2_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v2_6_loop_val)
go__go_3_0_87:
for {
if false { continue go__go_3_0_87 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var v2_6 *Constructor_Data_List_Types_Cons = v2_6_loop
_ = v2_6
var __t1 *Constructor_Data_List_Types_Cons
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t1 = v2_6
goto end_branch_1
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t1 = v2_6
goto end_branch_1
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil)) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
v2_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Apply2(f_0, (*Constructor_Data_List_Types_Cons)(v_4.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0), v2_6})})
continue go__go_3_0_87
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
})
})
var go__go_4_2_88 gopurs_runtime.Value
go__go_4_2_88 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_2_88:
for {
if false { continue go__go_4_2_88 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t3 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_2_88
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_4_2_88, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(go__go_3_0_87, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(ys_2)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}))
}



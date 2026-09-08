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
		cache_Data_List_identity = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
	})
	return cache_Data_List_identity
}

var cache_Data_List_Pattern gopurs_runtime.Value
var once_Data_List_Pattern sync.Once
func Get_Data_List_Pattern() gopurs_runtime.Value {
	once_Data_List_Pattern.Do(func() {
		cache_Data_List_Pattern = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_Pattern(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](x_0_box)))}
})
	})
	return cache_Data_List_Pattern
}

var cache_Data_List_updateAt gopurs_runtime.Value
var once_Data_List_updateAt sync.Once
func Get_Data_List_updateAt() gopurs_runtime.Value {
	once_Data_List_updateAt.Do(func() {
		cache_Data_List_updateAt = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_updateAt(v_0_box.IntVal, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_2_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_updateAt
}

var cache_Data_List_unzip gopurs_runtime.Value
var once_Data_List_unzip sync.Once
func Get_Data_List_unzip() gopurs_runtime.Value {
	once_Data_List_unzip.Do(func() {
		cache_Data_List_unzip = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_22134120_138441832(Rebox_Data_List_138441832_22134120(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_foldableList()).V2), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_0 shape=Other bindingType=Any
__local_var_1_0 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0
_ = __local_var_1_0
// TAST (Let): __local_var_2_1 shape=Other bindingType=Any
__local_var_2_1 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1
_ = __local_var_2_1
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_22134120_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, __local_var_1_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0)}))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, __local_var_2_1, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_22134120_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))})))))}
	})
	return cache_Data_List_unzip
}

var cache_Data_List_uncons gopurs_runtime.Value
var once_Data_List_uncons sync.Once
func Get_Data_List_uncons() gopurs_runtime.Value {
	once_Data_List_uncons.Do(func() {
		cache_Data_List_uncons = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_uncons(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_uncons
}

var cache_Data_List_toUnfoldable gopurs_runtime.Value
var once_Data_List_toUnfoldable sync.Once
func Get_Data_List_toUnfoldable() gopurs_runtime.Value {
	once_Data_List_toUnfoldable.Do(func() {
		cache_Data_List_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
})
	})
	return cache_Data_List_toUnfoldable
}

var cache_Data_List_tail gopurs_runtime.Value
var once_Data_List_tail sync.Once
func Get_Data_List_tail() gopurs_runtime.Value {
	once_Data_List_tail.Do(func() {
		cache_Data_List_tail = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_tail(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_tail
}

var cache_Data_List_stripPrefix gopurs_runtime.Value
var once_Data_List_stripPrefix sync.Once
func Get_Data_List_stripPrefix() gopurs_runtime.Value {
	once_Data_List_stripPrefix.Do(func() {
		cache_Data_List_stripPrefix = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_stripPrefix(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](s_2_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_stripPrefix
}

var cache_Data_List_span gopurs_runtime.Value
var once_Data_List_span sync.Once
func Get_Data_List_span() gopurs_runtime.Value {
	once_Data_List_span.Do(func() {
		cache_Data_List_span = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := Call_Data_List_span(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1_box))
				_ = orig
				return gopurs_runtime.RecordDict([]string{"init", "rest"}, []gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.go__init)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.rest)}})
				}()
})
	})
	return cache_Data_List_span
}

var cache_Data_List_snoc gopurs_runtime.Value
var once_Data_List_snoc sync.Once
func Get_Data_List_snoc() gopurs_runtime.Value {
	once_Data_List_snoc.Do(func() {
		cache_Data_List_snoc = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_snoc(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_0_box), x_1_box))}
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

var cache_Data_List_singleton__371370577 gopurs_runtime.Value
var once_Data_List_singleton__371370577 sync.Once
func Get_Data_List_singleton__371370577() gopurs_runtime.Value {
	once_Data_List_singleton__371370577.Do(func() {
		cache_Data_List_singleton__371370577 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_3704040722_849153993(Call_Data_List_singleton__371370577(a_0_box.IntVal)))}
})
	})
	return cache_Data_List_singleton__371370577
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
return Call_Data_List_sort(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_List_sort
}

var cache_Data_List_tails gopurs_runtime.Value
var once_Data_List_tails sync.Once
func Get_Data_List_tails() gopurs_runtime.Value {
	once_Data_List_tails.Do(func() {
		cache_Data_List_tails = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1116310629_849153993(Call_Data_List_tails(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0_box))))}
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
var Call_local_Data_List_go__go_0_0_7 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_0_0_7
var go__go_0_0_7 gopurs_runtime.Value
_ = go__go_0_0_7
Call_local_Data_List_go__go_0_0_7 = func(v_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_0_0_7:
for {
if false { continue go__go_0_0_7 }
var v_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_1_loop
_ = v_1
var v1_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_2_loop
_ = v1_2
var __t1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_2 == nil) {
__t1 = v_1
goto end_branch_1
} else {

}
}
{
if (v1_2 != nil) {
v_1_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_2).V0, v_1})
v1_2_loop = (v1_2).V1
continue go__go_0_0_7
__t1 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_1
} else {

}
}
{
__t1 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}
go__go_0_0_7 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_0_0_7(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_1_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_2_loop_val)))}
})
})
return gopurs_runtime.Apply(go__go_0_0_7, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
}()
	})
	return cache_Data_List_reverse
}

var cache_Data_List_reverse__2762254084 gopurs_runtime.Value
var once_Data_List_reverse__2762254084 sync.Once
func Get_Data_List_reverse__2762254084() gopurs_runtime.Value {
	once_Data_List_reverse__2762254084.Do(func() {
		cache_Data_List_reverse__2762254084 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_2442833393_849153993(Call_Data_List_reverse__2762254084(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]](__eta_norm_0_0_box))))}
})
	})
	return cache_Data_List_reverse__2762254084
}

var cache_Data_List_take gopurs_runtime.Value
var once_Data_List_take sync.Once
func Get_Data_List_take() gopurs_runtime.Value {
	once_Data_List_take.Do(func() {
		cache_Data_List_take = func() gopurs_runtime.Value {
var Call_local_Data_List_go__go_0_0_10 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_0_0_10
var go__go_0_0_10 gopurs_runtime.Value
_ = go__go_0_0_10
Call_local_Data_List_go__go_0_0_10 = func(v_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_2_loop int64, v2_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_0_0_10:
for {
if false { continue go__go_0_0_10 }
var v_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_1_loop
_ = v_1
var v1_2 int64 = v1_2_loop
_ = v1_2
var v2_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v2_3_loop
_ = v2_3
var __t5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_2) < (int64(1)) {
var Call_local_Data_List_go__go_4_1_11 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_4_1_11
var go__go_4_1_11 gopurs_runtime.Value
_ = go__go_4_1_11
Call_local_Data_List_go__go_4_1_11 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_4_1_11:
for {
if false { continue go__go_4_1_11 }
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var __t2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_6 == nil) {
__t2 = v_5
goto end_branch_2
} else {

}
}
{
if (v1_6 != nil) {
v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
v1_6_loop = (v1_6).V1
continue go__go_4_1_11
__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}
go__go_4_1_11 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_4_1_11(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
})
})
__t5 = Call_local_Data_List_go__go_4_1_11((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v_1)
goto end_branch_5
} else {

}
}
{
if (v2_3 == nil) {
var Call_local_Data_List_go__go_4_3_12 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_4_3_12
var go__go_4_3_12 gopurs_runtime.Value
_ = go__go_4_3_12
Call_local_Data_List_go__go_4_3_12 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_4_3_12:
for {
if false { continue go__go_4_3_12 }
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var __t4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_6 == nil) {
__t4 = v_5
goto end_branch_4
} else {

}
}
{
if (v1_6 != nil) {
v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
v1_6_loop = (v1_6).V1
continue go__go_4_3_12
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}
go__go_4_3_12 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_4_3_12(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
})
})
__t5 = Call_local_Data_List_go__go_4_3_12((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v_1)
goto end_branch_5
} else {

}
}
{
if (v2_3 != nil) {
v_1_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v2_3).V0, v_1})
v1_2_loop = (v1_2) - (int64(1))
v2_3_loop = (v2_3).V1
continue go__go_0_0_10
__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}
go__go_0_0_10 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_0_0_10(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_1_loop_val), v1_2_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_3_loop_val)))}
})
})
})
return gopurs_runtime.Apply(go__go_0_0_10, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
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
return func() gopurs_runtime.Value {
				_v := Call_Data_List_unsnoc(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](lst_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_unsnoc
}

var cache_Data_List_zipWith gopurs_runtime.Value
var once_Data_List_zipWith sync.Once
func Get_Data_List_zipWith() gopurs_runtime.Value {
	once_Data_List_zipWith.Do(func() {
		cache_Data_List_zipWith = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_zipWith(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_2_box)))}
})
	})
	return cache_Data_List_zipWith
}

var cache_Data_List_zip gopurs_runtime.Value
var once_Data_List_zip sync.Once
func Get_Data_List_zip() gopurs_runtime.Value {
	once_Data_List_zip.Do(func() {
		cache_Data_List_zip = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, ys_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_3266529732_849153993(Call_Data_List_zip(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_1_box))))}
})
	})
	return cache_Data_List_zip
}

var cache_Data_List_zipWithA gopurs_runtime.Value
var once_Data_List_zipWithA sync.Once
func Get_Data_List_zipWithA() gopurs_runtime.Value {
	once_Data_List_zipWithA.Do(func() {
		cache_Data_List_zipWithA = gopurs_runtime.Func4(func(dictApplicative_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value, ys_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_zipWithA(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), f_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_2_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_3_box))
})
	})
	return cache_Data_List_zipWithA
}

var cache_Data_List_go__range gopurs_runtime.Value
var once_Data_List_go__range sync.Once
func Get_Data_List_go__range() gopurs_runtime.Value {
	once_Data_List_go__range.Do(func() {
		cache_Data_List_go__range = gopurs_runtime.Func2(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_3704040722_849153993(Call_Data_List_go__range(start_0_box.IntVal, end_1_box.IntVal)))}
})
	})
	return cache_Data_List_go__range
}

var cache_Data_List_partition gopurs_runtime.Value
var once_Data_List_partition sync.Once
func Get_Data_List_partition() gopurs_runtime.Value {
	once_Data_List_partition.Do(func() {
		cache_Data_List_partition = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := Call_Data_List_partition(p_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1_box))
				_ = orig
				return gopurs_runtime.RecordDict([]string{"no", "yes"}, []gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.no)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.yes)}})
				}()
})
	})
	return cache_Data_List_partition
}

var cache_Data_List_null gopurs_runtime.Value
var once_Data_List_null sync.Once
func Get_Data_List_null() gopurs_runtime.Value {
	once_Data_List_null.Do(func() {
		cache_Data_List_null = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_List_null(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0_box)))
})
	})
	return cache_Data_List_null
}

var cache_Data_List_null__196316409 gopurs_runtime.Value
var once_Data_List_null__196316409 sync.Once
func Get_Data_List_null__196316409() gopurs_runtime.Value {
	once_Data_List_null__196316409.Do(func() {
		cache_Data_List_null__196316409 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_List_null__196316409(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]](v_0_box)))
})
	})
	return cache_Data_List_null__196316409
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
return Call_Data_List_nub(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_List_nub
}

var cache_Data_List_newtypePattern gopurs_runtime.Value
var once_Data_List_newtypePattern sync.Once
func Get_Data_List_newtypePattern() gopurs_runtime.Value {
	once_Data_List_newtypePattern.Do(func() {
		cache_Data_List_newtypePattern = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1501865320_385277032((&Constructor_Data_Newtype_Newtype[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})))}
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
return Call_Data_List_manyRec(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box))
})
	})
	return cache_Data_List_manyRec
}

var cache_Data_List_someRec gopurs_runtime.Value
var once_Data_List_someRec sync.Once
func Get_Data_List_someRec() gopurs_runtime.Value {
	once_Data_List_someRec.Do(func() {
		cache_Data_List_someRec = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictAlternative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_someRec(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](dictAlternative_1_box))
})
	})
	return cache_Data_List_someRec
}

var cache_Data_List_some gopurs_runtime.Value
var once_Data_List_some sync.Once
func Get_Data_List_some() gopurs_runtime.Value {
	once_Data_List_some.Do(func() {
		cache_Data_List_some = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_some(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](dictAlternative_0_box))
})
	})
	return cache_Data_List_some
}

var cache_Data_List_many gopurs_runtime.Value
var once_Data_List_many sync.Once
func Get_Data_List_many() gopurs_runtime.Value {
	once_Data_List_many.Do(func() {
		cache_Data_List_many = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_many(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](dictAlternative_0_box))
})
	})
	return cache_Data_List_many
}

var cache_Data_List_length gopurs_runtime.Value
var once_Data_List_length sync.Once
func Get_Data_List_length() gopurs_runtime.Value {
	once_Data_List_length.Do(func() {
		cache_Data_List_length = func() gopurs_runtime.Value {
var Call_local_Data_List_go__go_0_0_30 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_go__go_0_0_30
var go__go_0_0_30 gopurs_runtime.Value
_ = go__go_0_0_30
Call_local_Data_List_go__go_0_0_30 = func(b_1_loop gopurs_runtime.Value, v_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_0_0_30:
for {
if false { continue go__go_0_0_30 }
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var v_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (v_2 == nil) {
__t1 = b_1
goto end_branch_1
} else {

}
}
{
if (v_2 != nil) {
b_1_loop = gopurs_runtime.Int((b_1.IntVal) + (int64(1)))
v_2_loop = (v_2).V1
continue go__go_0_0_30
__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_0_0_30 = gopurs_runtime.Func(func(b_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_go__go_0_0_30(b_1_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val))
})
})
return gopurs_runtime.Int(gopurs_runtime.Apply(go__go_0_0_30, gopurs_runtime.Int(int64(0))).IntVal)
}()
	})
	return cache_Data_List_length
}

var cache_Data_List_last gopurs_runtime.Value
var once_Data_List_last sync.Once
func Get_Data_List_last() gopurs_runtime.Value {
	once_Data_List_last.Do(func() {
		cache_Data_List_last = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_last(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_last
}

var cache_Data_List_insertBy gopurs_runtime.Value
var once_Data_List_insertBy sync.Once
func Get_Data_List_insertBy() gopurs_runtime.Value {
	once_Data_List_insertBy.Do(func() {
		cache_Data_List_insertBy = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_insertBy(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_2_box)))}
})
	})
	return cache_Data_List_insertBy
}

var cache_Data_List_insertAt gopurs_runtime.Value
var once_Data_List_insertAt sync.Once
func Get_Data_List_insertAt() gopurs_runtime.Value {
	once_Data_List_insertAt.Do(func() {
		cache_Data_List_insertAt = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_insertAt(v_0_box.IntVal, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_2_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_insertAt
}

var cache_Data_List_insert gopurs_runtime.Value
var once_Data_List_insert sync.Once
func Get_Data_List_insert() gopurs_runtime.Value {
	once_Data_List_insert.Do(func() {
		cache_Data_List_insert = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_insert(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_List_insert
}

var cache_Data_List_go__init gopurs_runtime.Value
var once_Data_List_go__init sync.Once
func Get_Data_List_go__init() gopurs_runtime.Value {
	once_Data_List_go__init.Do(func() {
		cache_Data_List_go__init = gopurs_runtime.Func(func(lst_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_go__init(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](lst_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_go__init
}

var cache_Data_List_index gopurs_runtime.Value
var once_Data_List_index sync.Once
func Get_Data_List_index() gopurs_runtime.Value {
	once_Data_List_index.Do(func() {
		cache_Data_List_index = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_index(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0_box), v1_1_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_index
}

var cache_Data_List_head gopurs_runtime.Value
var once_Data_List_head sync.Once
func Get_Data_List_head() gopurs_runtime.Value {
	once_Data_List_head.Do(func() {
		cache_Data_List_head = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_head(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_head
}

var cache_Data_List_transpose gopurs_runtime.Value
var once_Data_List_transpose sync.Once
func Get_Data_List_transpose() gopurs_runtime.Value {
	once_Data_List_transpose.Do(func() {
		cache_Data_List_transpose = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1116310629_849153993(Call_Data_List_transpose(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](v_0_box))))}
})
	})
	return cache_Data_List_transpose
}

var cache_Data_List_groupBy gopurs_runtime.Value
var once_Data_List_groupBy sync.Once
func Get_Data_List_groupBy() gopurs_runtime.Value {
	once_Data_List_groupBy.Do(func() {
		cache_Data_List_groupBy = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_587109416_849153993(Call_Data_List_groupBy(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1_box))))}
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
return Call_Data_List_group(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_Data_List_group
}

var cache_Data_List_groupAll gopurs_runtime.Value
var once_Data_List_groupAll sync.Once
func Get_Data_List_groupAll() gopurs_runtime.Value {
	once_Data_List_groupAll.Do(func() {
		cache_Data_List_groupAll = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_groupAll(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_List_groupAll
}

var cache_Data_List_fromFoldable gopurs_runtime.Value
var once_Data_List_fromFoldable sync.Once
func Get_Data_List_fromFoldable() gopurs_runtime.Value {
	once_Data_List_fromFoldable.Do(func() {
		cache_Data_List_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
})
	})
	return cache_Data_List_fromFoldable
}

var cache_Data_List_foldM gopurs_runtime.Value
var once_Data_List_foldM sync.Once
func Get_Data_List_foldM() gopurs_runtime.Value {
	once_Data_List_foldM.Do(func() {
		cache_Data_List_foldM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_foldM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_Data_List_foldM
}

var cache_Data_List_findIndex gopurs_runtime.Value
var once_Data_List_findIndex sync.Once
func Get_Data_List_findIndex() gopurs_runtime.Value {
	once_Data_List_findIndex.Do(func() {
		cache_Data_List_findIndex = gopurs_runtime.Func(func(fn_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_findIndex(fn_0_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_findIndex
}

var cache_Data_List_findLastIndex gopurs_runtime.Value
var once_Data_List_findLastIndex sync.Once
func Get_Data_List_findLastIndex() gopurs_runtime.Value {
	once_Data_List_findLastIndex.Do(func() {
		cache_Data_List_findLastIndex = gopurs_runtime.Func2(func(fn_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_findLastIndex(fn_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_findLastIndex
}

var cache_Data_List_filterM gopurs_runtime.Value
var once_Data_List_filterM sync.Once
func Get_Data_List_filterM() gopurs_runtime.Value {
	once_Data_List_filterM.Do(func() {
		cache_Data_List_filterM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_filterM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
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
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_intersectBy(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_2_box)))}
})
	})
	return cache_Data_List_intersectBy
}

var cache_Data_List_intersect gopurs_runtime.Value
var once_Data_List_intersect sync.Once
func Get_Data_List_intersect() gopurs_runtime.Value {
	once_Data_List_intersect.Do(func() {
		cache_Data_List_intersect = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_intersect(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_Data_List_intersect
}

var cache_Data_List_nubByEq gopurs_runtime.Value
var once_Data_List_nubByEq sync.Once
func Get_Data_List_nubByEq() gopurs_runtime.Value {
	once_Data_List_nubByEq.Do(func() {
		cache_Data_List_nubByEq = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_nubByEq(v_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_Data_List_nubByEq
}

var cache_Data_List_nubEq gopurs_runtime.Value
var once_Data_List_nubEq sync.Once
func Get_Data_List_nubEq() gopurs_runtime.Value {
	once_Data_List_nubEq.Do(func() {
		cache_Data_List_nubEq = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_nubEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
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
return func() gopurs_runtime.Value {
				_v := Call_Data_List_elemLastIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_elemLastIndex
}

var cache_Data_List_elemIndex gopurs_runtime.Value
var once_Data_List_elemIndex sync.Once
func Get_Data_List_elemIndex() gopurs_runtime.Value {
	once_Data_List_elemIndex.Do(func() {
		cache_Data_List_elemIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_elemIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
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
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_dropEnd(n_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1_box)))}
})
	})
	return cache_Data_List_dropEnd
}

var cache_Data_List_drop gopurs_runtime.Value
var once_Data_List_drop sync.Once
func Get_Data_List_drop() gopurs_runtime.Value {
	once_Data_List_drop.Do(func() {
		cache_Data_List_drop = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_drop(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1_box)))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_takeEnd(n_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1_box)))}
})
	})
	return cache_Data_List_takeEnd
}

var cache_Data_List_deleteBy gopurs_runtime.Value
var once_Data_List_deleteBy sync.Once
func Get_Data_List_deleteBy() gopurs_runtime.Value {
	once_Data_List_deleteBy.Do(func() {
		cache_Data_List_deleteBy = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_deleteBy(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_2_box)))}
})
	})
	return cache_Data_List_deleteBy
}

var cache_Data_List_unionBy gopurs_runtime.Value
var once_Data_List_unionBy sync.Once
func Get_Data_List_unionBy() gopurs_runtime.Value {
	once_Data_List_unionBy.Do(func() {
		cache_Data_List_unionBy = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_unionBy(eq_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_2_box)))}
})
	})
	return cache_Data_List_unionBy
}

var cache_Data_List_union gopurs_runtime.Value
var once_Data_List_union sync.Once
func Get_Data_List_union() gopurs_runtime.Value {
	once_Data_List_union.Do(func() {
		cache_Data_List_union = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_union(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_Data_List_union
}

var cache_Data_List_deleteAt gopurs_runtime.Value
var once_Data_List_deleteAt sync.Once
func Get_Data_List_deleteAt() gopurs_runtime.Value {
	once_Data_List_deleteAt.Do(func() {
		cache_Data_List_deleteAt = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_deleteAt(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_1_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_deleteAt
}

var cache_Data_List_go__delete gopurs_runtime.Value
var once_Data_List_go__delete sync.Once
func Get_Data_List_go__delete() gopurs_runtime.Value {
	once_Data_List_go__delete.Do(func() {
		cache_Data_List_go__delete = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_go__delete(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_Data_List_go__delete
}

var cache_Data_List_difference gopurs_runtime.Value
var once_Data_List_difference sync.Once
func Get_Data_List_difference() gopurs_runtime.Value {
	once_Data_List_difference.Do(func() {
		cache_Data_List_difference = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_difference(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_Data_List_difference
}

var cache_Data_List_concatMap gopurs_runtime.Value
var once_Data_List_concatMap sync.Once
func Get_Data_List_concatMap() gopurs_runtime.Value {
	once_Data_List_concatMap.Do(func() {
		cache_Data_List_concatMap = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_concatMap(b_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](a_1_box)))}
})
	})
	return cache_Data_List_concatMap
}

var cache_Data_List_concat gopurs_runtime.Value
var once_Data_List_concat sync.Once
func Get_Data_List_concat() gopurs_runtime.Value {
	once_Data_List_concat.Do(func() {
		cache_Data_List_concat = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_concat(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](v_0_box)))}
})
	})
	return cache_Data_List_concat
}

var cache_Data_List_catMaybes gopurs_runtime.Value
var once_Data_List_catMaybes sync.Once
func Get_Data_List_catMaybes() gopurs_runtime.Value {
	once_Data_List_catMaybes.Do(func() {
		cache_Data_List_catMaybes = func() gopurs_runtime.Value {
var Call_local_Data_List_go__go_0_0_58 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_0_0_58
var go__go_0_0_58 gopurs_runtime.Value
_ = go__go_0_0_58
Call_local_Data_List_go__go_0_0_58 = func(v_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_0_0_58:
for {
if false { continue go__go_0_0_58 }
var v_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_1_loop
_ = v_1
var v1_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_2_loop
_ = v1_2
var __t6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_2 == nil) {
var Call_local_Data_List_go__go_3_1_59 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_3_1_59
var go__go_3_1_59 gopurs_runtime.Value
_ = go__go_3_1_59
Call_local_Data_List_go__go_3_1_59 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_3_1_59:
for {
if false { continue go__go_3_1_59 }
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_5_loop
_ = v1_5
var __t2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_5 == nil) {
__t2 = v_4
goto end_branch_2
} else {

}
}
{
if (v1_5 != nil) {
v_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_5).V0, v_4})
v1_5_loop = (v1_5).V1
continue go__go_3_1_59
__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}
go__go_3_1_59 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_3_1_59(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val)))}
})
})
__t6 = Call_local_Data_List_go__go_3_1_59((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v_1)
goto end_branch_6
} else {

}
}
{
if (v1_2 != nil) {
var __t5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]]((v1_2).V0)
if (__t_tag_3 == nil) {
v_1_loop = v_1
v1_2_loop = (v1_2).V1
continue go__go_0_0_58
__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]]((v1_2).V0)
if (__t_tag_4 != nil) {
v_1_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])((v1_2).V0.UnsafePtr).V0, v_1})
v1_2_loop = (v1_2).V1
continue go__go_0_0_58
__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}
}
go__go_0_0_58 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_0_0_58(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_1_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_2_loop_val)))}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_0_0_58, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})))}
}()
	})
	return cache_Data_List_catMaybes
}

var cache_Data_List_alterAt gopurs_runtime.Value
var once_Data_List_alterAt sync.Once
func Get_Data_List_alterAt() gopurs_runtime.Value {
	once_Data_List_alterAt.Do(func() {
		cache_Data_List_alterAt = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_alterAt(v_0_box.IntVal, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_2_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_alterAt
}

var cache_Data_List_modifyAt gopurs_runtime.Value
var once_Data_List_modifyAt sync.Once
func Get_Data_List_modifyAt() gopurs_runtime.Value {
	once_Data_List_modifyAt.Do(func() {
		cache_Data_List_modifyAt = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_modifyAt(n_0_box.IntVal, f_1_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_modifyAt
}

func Call_Data_List_Pattern(x_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
var x_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = x_0_loop
_ = x_0
return x_0
}

func Call_Data_List_updateAt(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
updateAt:
for {
if false { continue updateAt }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t3 gopurs_runtime.Value
{
if (v2_2 != nil) {
var __t2 gopurs_runtime.Value
{
if (v_0) == (int64(0)) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1978018568_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, v1_1, (v2_2).V1}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
goto end_branch_2
} else {

}
}
{
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","List","Types","List"] [(TypeVar a)])])
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := Call_Data_List_updateAt((v_0) - (int64(1)), v1_1, (v2_2).V1)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = __local_var_3_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_3_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v2_2).V0, (__local_var_3_0).V0}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1978018568_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_3:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1978018568_3094389156(Rebox_Data_List_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t3))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_List_uncons(v_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var v_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}]
{
if (v_0 == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_0
} else {

}
}
{
if (v_0 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.RecordDict2("head", "tail", (v_0).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)}), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_0
} else {

}
}
{
__t0 = func() *Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}] { panic("Failed pattern match") }()
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_2681346401_3094389156(__t0))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_List_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable_0.V1), gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1)
if (__t_tag_0 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1)
if (__t_tag_1 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_2351501316_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(xs_1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(xs_1.UnsafePtr).V1)}}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
}))
}

func Call_Data_List_tail(v_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var v_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]
{
if (v_0 == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_0
} else {

}
}
{
if (v_0 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v_0).V1)}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_0
} else {

}
}
{
__t0 = func() *Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] { panic("Failed pattern match") }()
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1978018568_3094389156(__t0))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_List_stripPrefix(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], v_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], s_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var v_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_1_loop
_ = v_1
var s_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = s_2_loop
_ = s_2
var Call_local_Data_List___local_var_3_0 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List___local_var_3_0
var __local_var_3_0 gopurs_runtime.Value
_ = __local_var_3_0
Call_local_Data_List___local_var_3_0 = func(prefix_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], input_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var prefix_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = prefix_3_loop
_ = prefix_3
var input_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = input_4_loop
_ = input_4
var __t3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (input_4 != nil) {
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (prefix_3 != nil) {
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), (prefix_3).V0, (input_4).V0).IntVal) != (0) {
__t1 = (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(Rebox_Data_List_3718343566_4008603408((&Constructor_Control_Monad_Rec_Class_Loop[struct{
	a *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	b *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, func() struct{
	a *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	b *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
} {
					orig := gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((prefix_3).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((input_4).V1)})
					_ = orig
					clone := struct{
	a *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	b *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}{}
					clone.a = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "a"))
					clone.b = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "b"))
					return clone
				}()})))}})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if (prefix_3 == nil) {
__t2 = (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(Rebox_Data_List_4239369586_3603546092((&Constructor_Control_Monad_Rec_Class_Done[struct{
	a *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	b *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, input_4})))}})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
if (prefix_3 == nil) {
__t3 = (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(Rebox_Data_List_4239369586_3603546092((&Constructor_Control_Monad_Rec_Class_Done[struct{
	a *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	b *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, input_4})))}})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
}
__local_var_3_0 = gopurs_runtime.Func(func(prefix_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(input_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List___local_var_3_0(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](prefix_3_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](input_4_loop_val))
})
})
// TAST (Let): __local_var_4_4 shape=LitRecord bindingType=(Record (Row [a: (TypeVar a), b: (TypeVar b)] Any))
__local_var_4_4 := struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(s_2)}}
_ = __local_var_4_4
var Call_local_Data_List___local_var_5_5 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List___local_var_5_5
var __local_var_5_5 gopurs_runtime.Value
_ = __local_var_5_5
Call_local_Data_List___local_var_5_5 = func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t11 gopurs_runtime.Value
{
var __t_tag_6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_5)
if (__t_tag_6 == nil) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(Rebox_Data_List_802708012_3603546092((&Constructor_Control_Monad_Rec_Class_Done[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})})))}
goto end_branch_11
} else {

}
}
{
var __t_tag_7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_5)
if (__t_tag_7 != nil) {
var __t10 gopurs_runtime.Value
{
var __t_tag_8 gopurs_runtime.Value = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_5.UnsafePtr).V0
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 525585346) {
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(Rebox_Data_List_3406595152_4008603408((&Constructor_Control_Monad_Rec_Class_Loop[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](Call_local_Data_List___local_var_3_0(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet((*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_5.UnsafePtr).V0.UnsafePtr).V0, "a")), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet((*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_5.UnsafePtr).V0.UnsafePtr).V0, "b"))))})))}
goto end_branch_10
} else {

}
}
{
var __t_tag_9 gopurs_runtime.Value = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_5.UnsafePtr).V0
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 60402430) {
__t10 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(Rebox_Data_List_802708012_3603546092((&Constructor_Control_Monad_Rec_Class_Done[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_5.UnsafePtr).V0.UnsafePtr).V0})})))}
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
__t11 = __t10
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
}
__local_var_5_5 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List___local_var_5_5(v_5_loop_val)
})
var Call_local_Data_List_go__go_6_12_0 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_go__go_6_12_0
var go__go_6_12_0 gopurs_runtime.Value
_ = go__go_6_12_0
Call_local_Data_List_go__go_6_12_0 = func(v_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_6_12_0:
for {
if false { continue go__go_6_12_0 }
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t13 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 525585346) {
v_7_loop = Call_local_Data_List___local_var_5_5((*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0)
continue go__go_6_12_0
__t13 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_13
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 60402430) {
__t13 = (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
}
}
go__go_6_12_0 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_go__go_6_12_0(v_7_loop_val)
})
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](Call_local_Data_List_go__go_6_12_0(Call_local_Data_List___local_var_5_5(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](Call_local_Data_List___local_var_3_0(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__local_var_4_4.a), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__local_var_4_4.b))))}))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_List_span(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct{
	go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	rest *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
} {
span:
for {
if false { continue span }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t1 struct{
	go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	rest *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}
{
if ((v1_1 != nil)) && ((gopurs_runtime.Apply(v_0, (v1_1).V0).IntVal) != (0)) {
// TAST (Let): v2_2_0 shape=App(Var) bindingType=(Record (Row [init: (ADT ["Data","List","Types","List"] [(TypeVar a)]), rest: (ADT ["Data","List","Types","List"] [(TypeVar a)])] Any))
v2_2_0 := Call_Data_List_span(v_0, (v1_1).V1)
_ = v2_2_0
__t1 = struct{
	go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	rest *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}{(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_1).V0, v2_2_0.go__init}), v2_2_0.rest}
goto end_branch_1
} else {

}
}
{
__t1 = struct{
	go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	rest *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}{(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v1_1}
}
end_branch_1:
return __t1
}
}

func Call_Data_List_snoc(xs_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
var xs_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = xs_0_loop
_ = xs_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_foldableList()).V2), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, x_1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)}))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_0)}))
}

func Call_Data_List_singleton(a_0_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, a_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)})
}

func Call_Data_List_singleton__371370577(a_0_loop int64) *Constructor_Data_List_Types_Cons[int64] {
singleton__371370577:
for {
if false { continue singleton__371370577 }
var a_0 int64 = a_0_loop
_ = a_0
return (&Constructor_Data_List_Types_Cons[int64]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_3704040722_849153993((*Constructor_Data_List_Types_Cons[int64])(nil)))}.IntVal, (*Constructor_Data_List_Types_Cons[int64])(nil)})
}
}

func Call_Data_List_sortBy(cmp_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
var merge_1_0_1 gopurs_runtime.Value
_ = merge_1_0_1
// FALLBACK TCO: isLoop=false len=1
merge_1_0_1 = gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2)
if (__t_tag_1 != nil) {
var __t4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3)
if (__t_tag_2 != nil) {
__t4 = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(merge_1_0_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}))})
goto end_branch_4
} else {

}
}
{
var __t_tag_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3)
if (__t_tag_3 == nil) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2)
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
__t7 = __t4
goto end_branch_7
} else {

}
}
{
var __t_tag_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2)
if (__t_tag_5 == nil) {
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3)
goto end_branch_7
} else {

}
}
{
var __t_tag_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3)
if (__t_tag_6 == nil) {
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2)
goto end_branch_7
} else {

}
}
{
__t7 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t7)}
})
var mergePairs_2_8_2 gopurs_runtime.Value
_ = mergePairs_2_8_2
// FALLBACK TCO: isLoop=false len=1
mergePairs_2_8_2 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]
{
var __t_tag_9 *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] = Rebox_Data_List_849153993_1116310629(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_3))
var __t_and_11 bool = false
if (__t_tag_9 != nil) {

var __t_tag_10 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1
__t_and_11 = (__t_tag_10 != nil)
}
if __t_and_11 {
__t12 = (&Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(merge_1_0_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1).V0))})), Rebox_Data_List_849153993_1116310629(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(mergePairs_2_8_2, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1).V1)})))})
goto end_branch_12
} else {

}
}
{
__t12 = Rebox_Data_List_849153993_1116310629(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_3))
}
end_branch_12:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1116310629_849153993(__t12))}
})
var Call_local_Data_List_mergeAll_3_13_3 func(*Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_mergeAll_3_13_3
var mergeAll_3_13_3 gopurs_runtime.Value
_ = mergeAll_3_13_3
Call_local_Data_List_mergeAll_3_13_3 = func(v_4_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
mergeAll_3_13_3:
for {
if false { continue mergeAll_3_13_3 }
var v_4 *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] = v_4_loop
_ = v_4
var __t16 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_and_15 bool = false
if (v_4 != nil) {

var __t_tag_14 *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] = (v_4).V1
__t_and_15 = (__t_tag_14 == nil)
}
if __t_and_15 {
__t16 = (v_4).V0
goto end_branch_16
} else {

}
}
{
v_4_loop = Rebox_Data_List_849153993_1116310629(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(mergePairs_2_8_2, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1116310629_849153993(v_4))})))
continue mergeAll_3_13_3
__t16 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
}
end_branch_16:
return __t16
}
}
mergeAll_3_13_3 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_mergeAll_3_13_3(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](v_4_loop_val)))}
})
var sequences_4_17_4 gopurs_runtime.Value
_ = sequences_4_17_4
// FALLBACK TCO: isLoop=false len=3
var descending_4_18_5 gopurs_runtime.Value
_ = descending_4_18_5
// FALLBACK TCO: isLoop=false len=3
var ascending_4_19_6 gopurs_runtime.Value
_ = ascending_4_19_6
// FALLBACK TCO: isLoop=false len=3
sequences_4_17_4 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t23 gopurs_runtime.Value
{
var __t_tag_20 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5)
var __t_and_22 bool = false
if (__t_tag_20 != nil) {

var __t_tag_21 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1
__t_and_22 = (__t_tag_21 != nil)
}
if __t_and_22 {
__t23 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1116310629_849153993(Rebox_Data_List_849153993_1116310629(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(descending_4_18_5, ((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)}))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1).V1)})))))}
goto end_branch_23
} else {

}
}
{
__t23 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5))}, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)}))}
}
end_branch_23:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1116310629_849153993(Rebox_Data_List_849153993_1116310629(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__t23))))}
})
descending_4_18_5 = gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t25 *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]
{
var __t_tag_24 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_7)
if (__t_tag_24 != nil) {
__t25 = Rebox_Data_List_849153993_1116310629(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(descending_4_18_5, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v2_7.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, v_5, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6)}))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v2_7.UnsafePtr).V1)})))
goto end_branch_25
} else {

}
}
{
__t25 = (&Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, v_5, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6)}), Rebox_Data_List_849153993_1116310629(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(sequences_4_17_4, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_7))})))})
}
end_branch_25:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1116310629_849153993(__t25))}
})
ascending_4_19_6 = gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1116310629_849153993((&Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, v_5, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)}))})), Rebox_Data_List_849153993_1116310629(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(sequences_4_17_4, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_7))})))})))}
})
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_mergeAll_3_13_3(Rebox_Data_List_849153993_1116310629(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(sequences_4_17_4, x_5)))))}
})
}

func Call_Data_List_sort(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 shape=Other bindingType=(Func [(TypeVar a), (TypeVar a)] (ADT ["Data","Ordering","Ordering"] []))
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_List_sortBy(compare_1_0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_2))})))}
})
}

func Call_Data_List_tails(v_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
tails:
for {
if false { continue tails }
var v_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0 == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)}))}
goto end_branch_0
} else {

}
}
{
if (v_0 != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1116310629_849153993((&Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, v_0, Call_Data_List_tails((v_0).V1)})))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1116310629_849153993(func() *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] { panic("Failed pattern match") }()))}
}
end_branch_0:
return Rebox_Data_List_849153993_1116310629(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__t0))
}
}

func Call_Data_List_showPattern(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): showList_1_0 shape=App(Var) bindingType=(ADT ["Data","Show","Show"] [(ADT ["Data","List","Types","List"] [(TypeVar a)])])
showList_1_0 := Rebox_Data_List_1386611502_3351995458(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Types_showList(), dictShow_0)))
_ = showList_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_List_3351995458_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Pattern ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showList_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2))}).StrVal())) + (")"))
})})))}
}

func Call_Data_List_reverse__2762254084(__eta_norm_0_0_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
reverse__2762254084:
for {
if false { continue reverse__2762254084 }
var __eta_norm_0_0 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] = __eta_norm_0_0_loop
_ = __eta_norm_0_0
var Call_local_Data_List_go__952457181_1_0_8 func(*Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]], *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
_ = Call_local_Data_List_go__952457181_1_0_8
var go__952457181_1_0_8 gopurs_runtime.Value
_ = go__952457181_1_0_8
Call_local_Data_List_go__952457181_1_0_8 = func(v_2_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]], v1_3_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
go__952457181_1_0_8:
for {
if false { continue go__952457181_1_0_8 }
var v_2 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] = v1_3_loop
_ = v1_3
var __t1 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
{
if (v1_3 == nil) {
__t1 = v_2
goto end_branch_1
} else {

}
}
{
if (v1_3 != nil) {
v_2_loop = (&Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]{1, (v1_3).V0, v_2})
v1_3_loop = (v1_3).V1
continue go__952457181_1_0_8
__t1 = func() *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] { panic("unreachable") }()
goto end_branch_1
} else {

}
}
{
__t1 = func() *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}
go__952457181_1_0_8 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_2442833393_849153993(Call_local_Data_List_go__952457181_1_0_8(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]](v_2_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]](v1_3_loop_val))))}
})
})
var go__go_2_2_9 gopurs_runtime.Value
_ = go__go_2_2_9
// FALLBACK TCO: isLoop=false len=1
go__go_2_2_9 = gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
{
var __t_tag_3 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] = Rebox_Data_List_849153993_2442833393(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_4))
if (__t_tag_3 == nil) {
__t5 = Rebox_Data_List_849153993_2442833393(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_3))
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] = Rebox_Data_List_849153993_2442833393(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_4))
if (__t_tag_4 != nil) {
__t5 = Call_local_Data_List_go__952457181_1_0_8((&Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]{1, Rebox_Data_List_138441832_3132786365(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V0)), Rebox_Data_List_849153993_2442833393(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_3))}), Rebox_Data_List_849153993_2442833393((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V1))
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_2442833393_849153993(__t5))}
})
return Call_local_Data_List_go__952457181_1_0_8((*Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]])(nil), __eta_norm_0_0)
}
}

func Call_Data_List_takeWhile(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var Call_local_Data_List_go__go_1_0_13 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_1_0_13
var go__go_1_0_13 gopurs_runtime.Value
_ = go__go_1_0_13
Call_local_Data_List_go__go_1_0_13 = func(v_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_1_0_13:
for {
if false { continue go__go_1_0_13 }
var v_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if ((v1_3 != nil)) && ((gopurs_runtime.Apply(p_0, (v1_3).V0).IntVal) != (0)) {
v_2_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_3).V0, v_2})
v1_3_loop = (v1_3).V1
continue go__go_1_0_13
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_3
} else {

}
}
{
var Call_local_Data_List_go__go_4_1_14 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_4_1_14
var go__go_4_1_14 gopurs_runtime.Value
_ = go__go_4_1_14
Call_local_Data_List_go__go_4_1_14 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_4_1_14:
for {
if false { continue go__go_4_1_14 }
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var __t2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_6 == nil) {
__t2 = v_5
goto end_branch_2
} else {

}
}
{
if (v1_6 != nil) {
v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
v1_6_loop = (v1_6).V1
continue go__go_4_1_14
__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}
go__go_4_1_14 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_4_1_14(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
})
})
__t3 = Call_local_Data_List_go__go_4_1_14((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v_2)
}
end_branch_3:
return __t3
}
}
go__go_1_0_13 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_1_0_13(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val)))}
})
})
return gopurs_runtime.Apply(go__go_1_0_13, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
}

func Call_Data_List_unsnoc(lst_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var lst_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = lst_0_loop
_ = lst_0
var Call_local_Data_List_go__go_1_0_15 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct{
	last gopurs_runtime.Value
	revInit *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}]
_ = Call_local_Data_List_go__go_1_0_15
var go__go_1_0_15 gopurs_runtime.Value
_ = go__go_1_0_15
Call_local_Data_List_go__go_1_0_15 = func(v_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct{
	last gopurs_runtime.Value
	revInit *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}] {
go__go_1_0_15:
for {
if false { continue go__go_1_0_15 }
var v_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t3 *Constructor_Data_Maybe_Just[struct{
	last gopurs_runtime.Value
	revInit *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}]
{
if (v_2 == nil) {
__t3 = Rebox_Data_List_3094389156_3270864776(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
goto end_branch_3
} else {

}
}
{
if (v_2 != nil) {
var __t2 *Constructor_Data_Maybe_Just[struct{
	last gopurs_runtime.Value
	revInit *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}]
{
var __t_tag_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = (v_2).V1
if (__t_tag_1 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	last gopurs_runtime.Value
	revInit *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.RecordDict2("last", "revInit", (v_2).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_2
} else {

}
}
{
v_2_loop = (v_2).V1
v1_3_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_2).V0, v1_3})
continue go__go_1_0_15
__t2 = func() *Constructor_Data_Maybe_Just[struct{
	last gopurs_runtime.Value
	revInit *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}] { panic("unreachable") }()
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_Maybe_Just[struct{
	last gopurs_runtime.Value
	revInit *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}] { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}
go__go_1_0_15 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_3270864776_3094389156(Call_local_Data_List_go__go_1_0_15(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val))))}
})
})
// TAST (Let): __local_var_2_4 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [revInit: (ADT ["Data","List","Types","List"] [(TypeVar a)]), last: (TypeVar a)] Any))])
__local_var_2_4 := Call_local_Data_List_go__go_1_0_15(lst_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))
_ = __local_var_2_4
var __t7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_2_4 != nil) {
var Call_local_Data_List_go__go_3_5_16 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_3_5_16
var go__go_3_5_16 gopurs_runtime.Value
_ = go__go_3_5_16
Call_local_Data_List_go__go_3_5_16 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_3_5_16:
for {
if false { continue go__go_3_5_16 }
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_5_loop
_ = v1_5
var __t6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_5 == nil) {
__t6 = v_4
goto end_branch_6
} else {

}
}
{
if (v1_5 != nil) {
v_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_5).V0, v_4})
v1_5_loop = (v1_5).V1
continue go__go_3_5_16
__t6 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}
}
go__go_3_5_16 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_3_5_16(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val)))}
})
})
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{func() gopurs_runtime.Value {
				orig := struct{
	go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	last gopurs_runtime.Value
}{Call_local_Data_List_go__go_3_5_16((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), (__local_var_2_4).V0.revInit), (__local_var_2_4).V0.last}
				_ = orig
				return gopurs_runtime.RecordDict([]string{"init", "last"}, []gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.go__init)}, orig.last})
				}(), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_7:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t7)}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_List_zipWith(f_0_loop gopurs_runtime.Value, xs_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], ys_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = xs_1_loop
_ = xs_1
var ys_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = ys_2_loop
_ = ys_2
var Call_local_Data_List_go__go_3_0_17 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_3_0_17
var go__go_3_0_17 gopurs_runtime.Value
_ = go__go_3_0_17
Call_local_Data_List_go__go_3_0_17 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v2_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_3_0_17:
for {
if false { continue go__go_3_0_17 }
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_5_loop
_ = v1_5
var v2_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v2_6_loop
_ = v2_6
var __t1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v_4 == nil) {
__t1 = v2_6
goto end_branch_1
} else {

}
}
{
if (v1_5 == nil) {
__t1 = v2_6
goto end_branch_1
} else {

}
}
{
if ((v_4 != nil)) && ((v1_5 != nil)) {
v_4_loop = (v_4).V1
v1_5_loop = (v1_5).V1
v2_6_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (v_4).V0, (v1_5).V0), v2_6})
continue go__go_3_0_17
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](func() gopurs_runtime.Value { panic("unreachable") }())
goto end_branch_1
} else {

}
}
{
__t1 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}
go__go_3_0_17 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_3_0_17(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_6_loop_val)))}
})
})
})
var Call_local_Data_List_go__go_4_2_18 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_4_2_18
var go__go_4_2_18 gopurs_runtime.Value
_ = go__go_4_2_18
Call_local_Data_List_go__go_4_2_18 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_4_2_18:
for {
if false { continue go__go_4_2_18 }
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_6 == nil) {
__t3 = v_5
goto end_branch_3
} else {

}
}
{
if (v1_6 != nil) {
v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
v1_6_loop = (v1_6).V1
continue go__go_4_2_18
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}
go__go_4_2_18 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_4_2_18(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
})
})
return Call_local_Data_List_go__go_4_2_18((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), Call_local_Data_List_go__go_3_0_17(xs_1, ys_2, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)))
}

func Call_Data_List_zip(xs_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], ys_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
var xs_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = xs_0_loop
_ = xs_0
var ys_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = ys_1_loop
_ = ys_1
var Call_local_Data_List_go__go_2_0_19 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_2_0_19
var go__go_2_0_19 gopurs_runtime.Value
_ = go__go_2_0_19
Call_local_Data_List_go__go_2_0_19 = func(v_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v2_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_2_0_19:
for {
if false { continue go__go_2_0_19 }
var v_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_3_loop
_ = v_3
var v1_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_4_loop
_ = v1_4
var v2_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v2_5_loop
_ = v2_5
var __t1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v_3 == nil) {
__t1 = v2_5
goto end_branch_1
} else {

}
}
{
if (v1_4 == nil) {
__t1 = v2_5
goto end_branch_1
} else {

}
}
{
if ((v_3 != nil)) && ((v1_4 != nil)) {
v_3_loop = (v_3).V1
v1_4_loop = (v1_4).V1
v2_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v_3).V0, (v1_4).V0}))}, v2_5})
continue go__go_2_0_19
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](func() gopurs_runtime.Value { panic("unreachable") }())
goto end_branch_1
} else {

}
}
{
__t1 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}
go__go_2_0_19 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_2_0_19(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_3_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_5_loop_val)))}
})
})
})
var Call_local_Data_List_go__go_3_2_20 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_3_2_20
var go__go_3_2_20 gopurs_runtime.Value
_ = go__go_3_2_20
Call_local_Data_List_go__go_3_2_20 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_3_2_20:
for {
if false { continue go__go_3_2_20 }
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_5_loop
_ = v1_5
var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_5 == nil) {
__t3 = v_4
goto end_branch_3
} else {

}
}
{
if (v1_5 != nil) {
v_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_5).V0, v_4})
v1_5_loop = (v1_5).V1
continue go__go_3_2_20
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}
go__go_3_2_20 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_3_2_20(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val)))}
})
})
return Rebox_Data_List_849153993_3266529732(Call_local_Data_List_go__go_3_2_20((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), Call_local_Data_List_go__go_2_0_19(xs_0, ys_1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))))
}

func Call_Data_List_zipWithA(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, xs_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], ys_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = xs_2_loop
_ = xs_2
var ys_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = ys_3_loop
_ = ys_3
var Call_local_Data_List_go__go_4_0_21 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_4_0_21
var go__go_4_0_21 gopurs_runtime.Value
_ = go__go_4_0_21
Call_local_Data_List_go__go_4_0_21 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v2_7_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_4_0_21:
for {
if false { continue go__go_4_0_21 }
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var v2_7 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v2_7_loop
_ = v2_7
var __t1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v_5 == nil) {
__t1 = v2_7
goto end_branch_1
} else {

}
}
{
if (v1_6 == nil) {
__t1 = v2_7
goto end_branch_1
} else {

}
}
{
if ((v_5 != nil)) && ((v1_6 != nil)) {
v_5_loop = (v_5).V1
v1_6_loop = (v1_6).V1
v2_7_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_1, (v_5).V0, (v1_6).V0), v2_7})
continue go__go_4_0_21
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](func() gopurs_runtime.Value { panic("unreachable") }())
goto end_branch_1
} else {

}
}
{
__t1 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}
go__go_4_0_21 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_4_0_21(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_7_loop_val)))}
})
})
})
var Call_local_Data_List_go__go_5_2_22 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_5_2_22
var go__go_5_2_22 gopurs_runtime.Value
_ = go__go_5_2_22
Call_local_Data_List_go__go_5_2_22 = func(v_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_7_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_5_2_22:
for {
if false { continue go__go_5_2_22 }
var v_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_6_loop
_ = v_6
var v1_7 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_7_loop
_ = v1_7
var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_7 == nil) {
__t3 = v_6
goto end_branch_3
} else {

}
}
{
if (v1_7 != nil) {
v_6_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_7).V0, v_6})
v1_7_loop = (v1_7).V1
continue go__go_5_2_22
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}
go__go_5_2_22 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_5_2_22(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_6_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_7_loop_val)))}
})
})
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_traversableList()).V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_5_2_22((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), Call_local_Data_List_go__go_4_0_21(xs_2, ys_3, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))))})
}

func Call_Data_List_go__range(start_0_loop int64, end_1_loop int64) *Constructor_Data_List_Types_Cons[int64] {
var start_0 int64 = start_0_loop
_ = start_0
var end_1 int64 = end_1_loop
_ = end_1
var __t5 *Constructor_Data_List_Types_Cons[int64]
{
if (start_0) == (end_1) {
__t5 = (&Constructor_Data_List_Types_Cons[int64]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_3704040722_849153993((*Constructor_Data_List_Types_Cons[int64])(nil)))}.IntVal, (*Constructor_Data_List_Types_Cons[int64])(nil)})
goto end_branch_5
} else {

}
}
{
var Call_local_Data_List_go__535684919_2_0_23 func(int64, int64, int64, *Constructor_Data_List_Types_Cons[int64]) *Constructor_Data_List_Types_Cons[int64]
_ = Call_local_Data_List_go__535684919_2_0_23
var go__535684919_2_0_23 gopurs_runtime.Value
_ = go__535684919_2_0_23
Call_local_Data_List_go__535684919_2_0_23 = func(s_3_loop int64, e_4_loop int64, step_5_loop int64, rest_6_loop *Constructor_Data_List_Types_Cons[int64]) *Constructor_Data_List_Types_Cons[int64] {
go__535684919_2_0_23:
for {
if false { continue go__535684919_2_0_23 }
var s_3 int64 = s_3_loop
_ = s_3
var e_4 int64 = e_4_loop
_ = e_4
var step_5 int64 = step_5_loop
_ = step_5
var rest_6 *Constructor_Data_List_Types_Cons[int64] = rest_6_loop
_ = rest_6
var __t1 *Constructor_Data_List_Types_Cons[int64]
{
if (s_3) == (e_4) {
__t1 = (&Constructor_Data_List_Types_Cons[int64]{1, s_3, rest_6})
goto end_branch_1
} else {

}
}
{
s_3_loop = (s_3) + (step_5)
e_4_loop = e_4
step_5_loop = step_5
rest_6_loop = (&Constructor_Data_List_Types_Cons[int64]{1, s_3, rest_6})
continue go__535684919_2_0_23
__t1 = func() *Constructor_Data_List_Types_Cons[int64] { panic("unreachable") }()
}
end_branch_1:
return __t1
}
}
go__535684919_2_0_23 = gopurs_runtime.Func(func(s_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(e_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(step_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rest_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_3704040722_849153993(Call_local_Data_List_go__535684919_2_0_23(s_3_loop_val.IntVal, e_4_loop_val.IntVal, step_5_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[int64]](rest_6_loop_val))))}
})
})
})
})
var go__go_3_2_24 gopurs_runtime.Value
_ = go__go_3_2_24
// FALLBACK TCO: isLoop=false len=1
go__go_3_2_24 = gopurs_runtime.Func4(func(s_4 gopurs_runtime.Value, e_5 gopurs_runtime.Value, step_6 gopurs_runtime.Value, rest_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_List_Types_Cons[int64]
{
if (s_4.IntVal) == (e_5.IntVal) {
__t3 = (&Constructor_Data_List_Types_Cons[int64]{1, gopurs_runtime.Int(s_4.IntVal).IntVal, Rebox_Data_List_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](rest_7))})
goto end_branch_3
} else {

}
}
{
__t3 = Call_local_Data_List_go__535684919_2_0_23((s_4.IntVal) + (step_6.IntVal), e_5.IntVal, step_6.IntVal, (&Constructor_Data_List_Types_Cons[int64]{1, gopurs_runtime.Int(s_4.IntVal).IntVal, Rebox_Data_List_849153993_3704040722(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](rest_7))}))
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_3704040722_849153993(__t3))}
})
var __t4 int64
{
if (start_0) > (end_1) {
__t4 = int64(1)
goto end_branch_4
} else {

}
}
{
__t4 = int64(-1)
}
end_branch_4:
__t5 = Call_local_Data_List_go__535684919_2_0_23(end_1, start_0, __t4, (*Constructor_Data_List_Types_Cons[int64])(nil))
}
end_branch_5:
return __t5
}

func Call_Data_List_partition(p_0_loop gopurs_runtime.Value, xs_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct{
	no *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	yes *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
} {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var xs_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = xs_1_loop
_ = xs_1
return func() struct{
	no *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	yes *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
} {
					orig := gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_foldableList()).V2), gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 struct{
	no *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	yes *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}
{
if (gopurs_runtime.Apply(p_0, x_2).IntVal) != (0) {
__t0 = struct{
	no *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	yes *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}{gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v_3, "no")), (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, x_2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v_3, "yes"))})}
goto end_branch_0
} else {

}
}
{
__t0 = struct{
	no *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	yes *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}{(&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, x_2, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v_3, "no"))}), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v_3, "yes"))}
}
end_branch_0:
return func() gopurs_runtime.Value {
				orig := __t0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"no", "yes"}, []gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.no)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.yes)}})
				}()
}), func() gopurs_runtime.Value {
				orig := struct{
	no *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	yes *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}{(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)}
				_ = orig
				return gopurs_runtime.RecordDict([]string{"no", "yes"}, []gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.no)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.yes)}})
				}(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)})
					_ = orig
					clone := struct{
	no *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	yes *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}{}
					clone.no = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "no"))
					clone.yes = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "yes"))
					return clone
				}()
}

func Call_Data_List_null(v_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) bool {
var v_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
return (v_0 == nil)
}

func Call_Data_List_null__196316409(v_0_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) bool {
null__196316409:
for {
if false { continue null__196316409 }
var v_0 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] = v_0_loop
_ = v_0
return (v_0 == nil)
}
}

func Call_Data_List_nubBy(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var Call_local_Data_List_go__go_1_0_25 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_1_0_25
var go__go_1_0_25 gopurs_runtime.Value
_ = go__go_1_0_25
Call_local_Data_List_go__go_1_0_25 = func(v_2_loop gopurs_runtime.Value, v1_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v2_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_1_0_25:
for {
if false { continue go__go_1_0_25 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var v2_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v2_4_loop
_ = v2_4
var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v2_4 == nil) {
__t3 = v1_3
goto end_branch_3
} else {

}
}
{
if (v2_4 != nil) {
// TAST (Let): v3_5_1 shape=App(Var) bindingType=(Record (Row [found: Boolean, result: (ADT ["Data","List","Internal","Set"] [(TypeVar a)])] Any))
v3_5_1 := func() struct{
	found bool
	result gopurs_runtime.Value
} {
					orig := gopurs_runtime.Apply3(Get_Data_List_Internal_insertAndLookupBy(), p_0, (v2_4).V0, v_2)
					_ = orig
					clone := struct{
	found bool
	result gopurs_runtime.Value
}{}
					clone.found = (gopurs_runtime.RecordGet(orig, "found").IntVal) != (0)
					clone.result = gopurs_runtime.RecordGet(orig, "result")
					return clone
				}()
_ = v3_5_1
var __t2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if v3_5_1.found {
v_2_loop = v3_5_1.result
v1_3_loop = v1_3
v2_4_loop = (v2_4).V1
continue go__go_1_0_25
__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
v_2_loop = v3_5_1.result
v1_3_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v2_4).V0, v1_3})
v2_4_loop = (v2_4).V1
continue go__go_1_0_25
__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}
go__go_1_0_25 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_1_0_25(v_2_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_4_loop_val)))}
})
})
})
// TAST (Let): __local_var_2_4 shape=App(Other) bindingType=(Func [(ADT ["Data","List","Types","List"] [(TypeVar a)])] (ADT ["Data","List","Types","List"] [(TypeVar a)]))
__local_var_2_4 := gopurs_runtime.Apply2(go__go_1_0_25, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_2_4
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_go__go_4_5_26 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_4_5_26
var go__go_4_5_26 gopurs_runtime.Value
_ = go__go_4_5_26
Call_local_Data_List_go__go_4_5_26 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_4_5_26:
for {
if false { continue go__go_4_5_26 }
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var __t6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_6 == nil) {
__t6 = v_5
goto end_branch_6
} else {

}
}
{
if (v1_6 != nil) {
v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
v1_6_loop = (v1_6).V1
continue go__go_4_5_26
__t6 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}
}
go__go_4_5_26 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_4_5_26(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_4_5_26((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_2_4, x_3))))}
})
}

func Call_Data_List_nub(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return Call_Data_List_nubBy(gopurs_runtime.Box(dictOrd_0.V1))
}

func Call_Data_List_mapMaybe(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var Call_local_Data_List_go__go_1_0_27 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_1_0_27
var go__go_1_0_27 gopurs_runtime.Value
_ = go__go_1_0_27
Call_local_Data_List_go__go_1_0_27 = func(v_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_1_0_27:
for {
if false { continue go__go_1_0_27 }
var v_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_3 == nil) {
var Call_local_Data_List_go__go_4_1_28 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_4_1_28
var go__go_4_1_28 gopurs_runtime.Value
_ = go__go_4_1_28
Call_local_Data_List_go__go_4_1_28 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_4_1_28:
for {
if false { continue go__go_4_1_28 }
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var __t2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_6 == nil) {
__t2 = v_5
goto end_branch_2
} else {

}
}
{
if (v1_6 != nil) {
v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
v1_6_loop = (v1_6).V1
continue go__go_4_1_28
__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}
go__go_4_1_28 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_4_1_28(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
})
})
__t5 = Call_local_Data_List_go__go_4_1_28((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v_2)
goto end_branch_5
} else {

}
}
{
if (v1_3 != nil) {
// TAST (Let): v2_4_3 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
v2_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, (v1_3).V0))
_ = v2_4_3
var __t4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v2_4_3 == nil) {
v_2_loop = v_2
v1_3_loop = (v1_3).V1
continue go__go_1_0_27
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_4
} else {

}
}
{
if (v2_4_3 != nil) {
v_2_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v2_4_3).V0, v_2})
v1_3_loop = (v1_3).V1
continue go__go_1_0_27
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}
go__go_1_0_27 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_1_0_27(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val)))}
})
})
return gopurs_runtime.Apply(go__go_1_0_27, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
}

func Call_Data_List_manyRec(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar f)])
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V0), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.Func(func(dictAlternative_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Plus1_3_1 shape=App(Other) bindingType=Any
Plus1_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Plus1"), gopurs_runtime.Value{})
_ = Plus1_3_1
// TAST (Let): Alt0_4_2 shape=App(Other) bindingType=(ADT ["Control","Alt","Alt"] [(TypeVar f)])
Alt0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Plus1_3_1, "Alt0"), gopurs_runtime.Value{}))
_ = Alt0_4_2
// TAST (Let): Functor0_5_3 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Plus1_3_1, "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_3
// TAST (Let): Applicative0_6_4 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar f)])
Applicative0_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_4
// TAST (Let): pure_7_5 shape=Other bindingType=(Func [(ADT ["Control","Monad","Rec","Class","Step"] [(ADT ["Data","List","Types","List"] [(TypeVar a)]), (ADT ["Data","List","Types","List"] [(TypeVar a)])])] (TypeApp (TypeVar f) [(ADT ["Control","Monad","Rec","Class","Step"] [(ADT ["Data","List","Types","List"] [(TypeVar a)]), (ADT ["Data","List","Types","List"] [(TypeVar a)])])]))
pure_7_5 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_5
return gopurs_runtime.Func(func(p_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(acc_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Alt0_4_2.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_3.V0), Get_Control_Monad_Rec_Class_Loop(), p_8), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_4.V1), gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, Get_Data_Unit_unit()}))})), gopurs_runtime.Func(func(aa_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (aa_10.Type == 9 && aa_10.IntVal == 525585346) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(aa_10.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](acc_9)}))}}))}
goto end_branch_8
} else {

}
}
{
if (aa_10.Type == 9 && aa_10.IntVal == 60402430) {
var Call_local_Data_List_go__go_11_6_29 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_11_6_29
var go__go_11_6_29 gopurs_runtime.Value
_ = go__go_11_6_29
Call_local_Data_List_go__go_11_6_29 = func(v_12_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_13_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_11_6_29:
for {
if false { continue go__go_11_6_29 }
var v_12 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_12_loop
_ = v_12
var v1_13 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_13_loop
_ = v1_13
var __t7 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_13 == nil) {
__t7 = v_12
goto end_branch_7
} else {

}
}
{
if (v1_13 != nil) {
v_12_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_13).V0, v_12})
v1_13_loop = (v1_13).V1
continue go__go_11_6_29
__t7 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_7
} else {

}
}
{
__t7 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}
}
go__go_11_6_29 = gopurs_runtime.Func(func(v_12_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_13_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_11_6_29(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_12_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_13_loop_val)))}
})
})
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_11_6_29((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](acc_9)))}}))}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return gopurs_runtime.Apply(pure_7_5, __t8)
}))
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
})
})
}

func Call_Data_List_someRec(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value], dictAlternative_1_loop *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictAlternative_1 *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] = dictAlternative_1_loop
_ = dictAlternative_1
// TAST (Let): Apply0_2_0 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar f)])
Apply0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_1.V0), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_0
// TAST (Let): Functor0_3_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_1.V1), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_1
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_2_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_1.V0), Get_Data_List_Types_Cons(), v_4), gopurs_runtime.Apply2(Call_Data_List_manyRec(dictMonadRec_0), gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(dictAlternative_1)}, v_4))
})
}

func Call_Data_List_some(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): Apply0_1_0 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar f)])
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V0), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V1), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func2(func(dictLazy_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_List_Types_Cons(), v_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_3, "defer"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Data_List_many(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dictLazy_3))}, v_4)
})))
})
}

func Call_Data_List_many(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): Alt0_1_0 shape=App(Other) bindingType=(ADT ["Control","Alt","Alt"] [(TypeVar f)])
Alt0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V1), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}))
_ = Alt0_1_0
// TAST (Let): Applicative0_2_1 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar f)])
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func2(func(dictLazy_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Alt0_1_0.V1), gopurs_runtime.Apply2(Call_Data_List_some(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dictLazy_3))}, v_4), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}))}))
})
}

func Call_Data_List_last(v_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
last:
for {
if false { continue last }
var v_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v_0 != nil) {
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = (v_0).V1
if (__t_tag_0 == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(v_0).V0, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
v_0_loop = (v_0).V1
continue last
__t1 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("unreachable") }()
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_2:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_List_insertBy(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
insertBy:
for {
if false { continue insertBy }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v2_2 == nil) {
__t2 = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, v1_1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)})
goto end_branch_2
} else {

}
}
{
if (v2_2 != nil) {
var __t1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_0 uint32 = uint32(gopurs_runtime.Apply2(v_0, v1_1, (v2_2).V0).IntVal)
if (uint32(__t_tag_0) == 380165415) {
__t1 = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v2_2).V0, Call_Data_List_insertBy(v_0, v1_1, (v2_2).V1)})
goto end_branch_1
} else {

}
}
{
__t1 = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, v1_1, v2_2})
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}

func Call_Data_List_insertAt(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
insertAt:
for {
if false { continue insertAt }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t2 gopurs_runtime.Value
{
if (v_0) == (int64(0)) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1978018568_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, v1_1, v2_2}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
goto end_branch_2
} else {

}
}
{
if (v2_2 != nil) {
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","List","Types","List"] [(TypeVar a)])])
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := Call_Data_List_insertAt((v_0) - (int64(1)), v1_1, (v2_2).V1)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = __local_var_3_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_3_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v2_2).V0, (__local_var_3_0).V0}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1978018568_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_2:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1978018568_3094389156(Rebox_Data_List_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t2))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_List_insert(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_List_insertBy(), gopurs_runtime.Box(dictOrd_0.V1))
}

func Call_Data_List_go__init(lst_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var lst_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = lst_0_loop
_ = lst_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [init: (ADT ["Data","List","Types","List"] [(TypeVar a)]), last: (TypeVar a)] Any))])
__local_var_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	go__init *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	last gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := Call_Data_List_unsnoc(lst_0)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = __local_var_1_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_1_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((__local_var_1_0).V0.go__init)}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_1:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_List_index(v_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_1_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
index:
for {
if false { continue index }
var v_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var v1_1 int64 = v1_1_loop
_ = v1_1
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v_0 == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
if (v_0 != nil) {
var __t0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v1_1) == (int64(0)) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(v_0).V0, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_0
} else {

}
}
{
v_0_loop = (v_0).V1
v1_1_loop = (v1_1) - (int64(1))
continue index
__t0 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("unreachable") }()
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
__t1 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_1:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_List_head(v_0_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var v_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v_0 == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_0
} else {

}
}
{
if (v_0 != nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(v_0).V0, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_0
} else {

}
}
{
__t0 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_List_transpose(v_0_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
transpose:
for {
if false { continue transpose }
var v_0 *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] = v_0_loop
_ = v_0
var __t17 *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]
{
if (v_0 == nil) {
__t17 = (*Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]])(nil)
goto end_branch_17
} else {

}
}
{
if (v_0 != nil) {
var __t16 *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]
{
var __t_tag_0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = (v_0).V0
if (__t_tag_0 == nil) {
v_0_loop = (v_0).V1
continue transpose
__t16 = func() *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] { panic("unreachable") }()
goto end_branch_16
} else {

}
}
{
var __t_tag_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = (v_0).V0
if (__t_tag_1 != nil) {
var Call_local_Data_List_go__go_1_2_31 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_1_2_31
var go__go_1_2_31 gopurs_runtime.Value
_ = go__go_1_2_31
Call_local_Data_List_go__go_1_2_31 = func(v_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_1_2_31:
for {
if false { continue go__go_1_2_31 }
var v_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t8 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_3 == nil) {
var Call_local_Data_List_go__go_4_3_32 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_4_3_32
var go__go_4_3_32 gopurs_runtime.Value
_ = go__go_4_3_32
Call_local_Data_List_go__go_4_3_32 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_4_3_32:
for {
if false { continue go__go_4_3_32 }
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var __t4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_6 == nil) {
__t4 = v_5
goto end_branch_4
} else {

}
}
{
if (v1_6 != nil) {
v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
v1_6_loop = (v1_6).V1
continue go__go_4_3_32
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}
go__go_4_3_32 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_4_3_32(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
})
})
__t8 = Call_local_Data_List_go__go_4_3_32((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v_2)
goto end_branch_8
} else {

}
}
{
if (v1_3 != nil) {
var __t7 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_5 gopurs_runtime.Value = (v1_3).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 1358893437 && __t_tag_5.UnsafePtr == nil) {
v_2_loop = v_2
v1_3_loop = (v1_3).V1
continue go__go_1_2_31
__t7 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_7
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (v1_3).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 1358893437 && __t_tag_6.UnsafePtr != nil) {
v_2_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((v1_3).V0.UnsafePtr).V0, v_2})
v1_3_loop = (v1_3).V1
continue go__go_1_2_31
__t7 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_7
} else {

}
}
{
__t7 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_7:
__t8 = __t7
goto end_branch_8
} else {

}
}
{
__t8 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
}
}
go__go_1_2_31 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_1_2_31(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val)))}
})
})
var Call_local_Data_List_go__go_1_9_33 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_1_9_33
var go__go_1_9_33 gopurs_runtime.Value
_ = go__go_1_9_33
Call_local_Data_List_go__go_1_9_33 = func(v_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_1_9_33:
for {
if false { continue go__go_1_9_33 }
var v_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t15 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_3 == nil) {
var Call_local_Data_List_go__go_4_10_34 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_4_10_34
var go__go_4_10_34 gopurs_runtime.Value
_ = go__go_4_10_34
Call_local_Data_List_go__go_4_10_34 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_4_10_34:
for {
if false { continue go__go_4_10_34 }
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var __t11 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_6 == nil) {
__t11 = v_5
goto end_branch_11
} else {

}
}
{
if (v1_6 != nil) {
v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
v1_6_loop = (v1_6).V1
continue go__go_4_10_34
__t11 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_11
} else {

}
}
{
__t11 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
}
}
go__go_4_10_34 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_4_10_34(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
})
})
__t15 = Call_local_Data_List_go__go_4_10_34((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v_2)
goto end_branch_15
} else {

}
}
{
if (v1_3 != nil) {
var __t14 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_12 gopurs_runtime.Value = (v1_3).V0
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 1358893437 && __t_tag_12.UnsafePtr == nil) {
v_2_loop = v_2
v1_3_loop = (v1_3).V1
continue go__go_1_9_33
__t14 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_14
} else {

}
}
{
var __t_tag_13 gopurs_runtime.Value = (v1_3).V0
if (__t_tag_13.Type == 9 && __t_tag_13.IntVal == 1358893437 && __t_tag_13.UnsafePtr != nil) {
v_2_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])((v1_3).V0.UnsafePtr).V1)}, v_2})
v1_3_loop = (v1_3).V1
continue go__go_1_9_33
__t14 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_14
} else {

}
}
{
__t14 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_14:
__t15 = __t14
goto end_branch_15
} else {

}
}
{
__t15 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
}
}
go__go_1_9_33 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_1_9_33(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val)))}
})
})
__t16 = (&Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, ((v_0).V0).V0, Call_local_Data_List_go__go_1_2_31((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), Rebox_Data_List_1116310629_849153993((v_0).V1))}), Call_Data_List_transpose((&Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, ((v_0).V0).V1, Rebox_Data_List_849153993_1116310629(Call_local_Data_List_go__go_1_9_33((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), Rebox_Data_List_1116310629_849153993((v_0).V1)))}))})
goto end_branch_16
} else {

}
}
{
__t16 = func() *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] { panic("Failed pattern match") }()
}
end_branch_16:
__t17 = __t16
goto end_branch_17
} else {

}
}
{
__t17 = func() *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] { panic("Failed pattern match") }()
}
end_branch_17:
return __t17
}
}

func Call_Data_List_groupBy(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
groupBy:
for {
if false { continue groupBy }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t1 *Constructor_Data_List_Types_Cons[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]
{
if (v1_1 == nil) {
__t1 = (*Constructor_Data_List_Types_Cons[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]])(nil)
goto end_branch_1
} else {

}
}
{
if (v1_1 != nil) {
// TAST (Let): v2_2_0 shape=App(Var) bindingType=(Record (Row [init: (ADT ["Data","List","Types","List"] [(TypeVar a)]), rest: (ADT ["Data","List","Types","List"] [(TypeVar a)])] Any))
v2_2_0 := Call_Data_List_span(gopurs_runtime.Apply(v_0, (v1_1).V0), (v1_1).V1)
_ = v2_2_0
__t1 = (&Constructor_Data_List_Types_Cons[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]{1, Rebox_Data_List_1293498952_3123684004((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (v1_1).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2_0.go__init)}})), Call_Data_List_groupBy(v_0, v2_2_0.rest)})
goto end_branch_1
} else {

}
}
{
__t1 = func() *Constructor_Data_List_Types_Cons[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}

func Call_Data_List_groupAllBy(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [(ADT ["Data","List","Types","List"] [(TypeVar a)])] (ADT ["Data","List","Types","List"] [(TypeVar a)]))
__local_var_1_0 := Call_Data_List_sortBy(p_0)
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_587109416_849153993(Call_Data_List_groupBy(gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_1_0, x_2)))))}
})
}

func Call_Data_List_group(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_List_groupBy(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_List_groupAll(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
__local_var_1_0 := Call_Data_List_group(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictOrd_0.V0), gopurs_runtime.Value{})))
_ = __local_var_1_0
// TAST (Let): compare_2_2 shape=Other bindingType=(Func [(TypeVar a), (TypeVar a)] (ADT ["Data","Ordering","Ordering"] []))
compare_2_2 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_2_2
// TAST (Let): __local_var_2_1 shape=Let(Abs(App(Var))) bindingType=(Func [(ADT ["Data","List","Types","List"] [(TypeVar a)])] (ADT ["Data","List","Types","List"] [(TypeVar a)]))
__local_var_2_1 := gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Call_Data_List_sortBy(compare_2_2), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_3))})))}
})
_ = __local_var_2_1
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply(__local_var_2_1, x_3))
})
}

func Call_Data_List_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V2), Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
}

func Call_Data_List_foldM(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
foldM:
for {
if false { continue foldM }
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
var __t_tag_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_5)
if (__t_tag_2 == nil) {
__t5 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), v1_4)
goto end_branch_5
} else {

}
}
{
var __t_tag_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v2_5)
if (__t_tag_3 != nil) {
// TAST (Let): __local_var_6_4 shape=Other bindingType=Any
__local_var_6_4 := (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V1
_ = __local_var_6_4
__t5 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(v_3, v1_4, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0), gopurs_runtime.Func(func(b_prime__7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Call_Data_List_foldM(dictMonad_0), v_3, b_prime__7, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__local_var_6_4)})
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
}
}

func Call_Data_List_findIndex(fn_0_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var Call_local_Data_List_go__816245670_1_0_35 func(int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64]
_ = Call_local_Data_List_go__816245670_1_0_35
var go__816245670_1_0_35 gopurs_runtime.Value
_ = go__816245670_1_0_35
Call_local_Data_List_go__816245670_1_0_35 = func(v_2_loop int64, v1_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
go__816245670_1_0_35:
for {
if false { continue go__816245670_1_0_35 }
var v_2 int64 = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t2 *Constructor_Data_Maybe_Just[int64]
{
if (v1_3 != nil) {
var __t1 *Constructor_Data_Maybe_Just[int64]
{
if (gopurs_runtime.Apply(fn_0, (v1_3).V0).IntVal) != (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int(v_2), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
v_2_loop = (v_2) + (int64(1))
v1_3_loop = (v1_3).V1
continue go__816245670_1_0_35
__t1 = func() *Constructor_Data_Maybe_Just[int64] { panic("unreachable") }()
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if (v1_3 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_Maybe_Just[int64] { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}
go__816245670_1_0_35 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1170268447_3094389156(Call_local_Data_List_go__816245670_1_0_35(v_2_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val))))}
})
})
var go__go_2_3_36 gopurs_runtime.Value
_ = go__go_2_3_36
// FALLBACK TCO: isLoop=false len=1
go__go_2_3_36 = gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
var __t_tag_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_4)
if (__t_tag_4 != nil) {
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(fn_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V0).IntVal) != (0) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(v_3.IntVal)}))}
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1170268447_3094389156(Call_local_Data_List_go__816245670_1_0_35((v_3.IntVal) + (int64(1)), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V1)))}
}
end_branch_5:
__t7 = __t5
goto end_branch_7
} else {

}
}
{
var __t_tag_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_4)
if (__t_tag_6 == nil) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1170268447_3094389156(func() *Constructor_Data_Maybe_Just[int64] { panic("Failed pattern match") }()))}
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1170268447_3094389156(Rebox_Data_List_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t7))))}
})
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply(go__816245670_1_0_35, gopurs_runtime.Int(int64(0)))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_List_findLastIndex(fn_0_loop gopurs_runtime.Value, xs_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var xs_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = xs_1_loop
_ = xs_1
var Call_local_Data_List_go__816245670_2_1_37 func(int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64]
_ = Call_local_Data_List_go__816245670_2_1_37
var go__816245670_2_1_37 gopurs_runtime.Value
_ = go__816245670_2_1_37
Call_local_Data_List_go__816245670_2_1_37 = func(v_3_loop int64, v1_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
go__816245670_2_1_37:
for {
if false { continue go__816245670_2_1_37 }
var v_3 int64 = v_3_loop
_ = v_3
var v1_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_4_loop
_ = v1_4
var __t3 *Constructor_Data_Maybe_Just[int64]
{
if (v1_4 != nil) {
var __t2 *Constructor_Data_Maybe_Just[int64]
{
if (gopurs_runtime.Apply(fn_0, (v1_4).V0).IntVal) != (0) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int(v_3), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_2
} else {

}
}
{
v_3_loop = (v_3) + (int64(1))
v1_4_loop = (v1_4).V1
continue go__816245670_2_1_37
__t2 = func() *Constructor_Data_Maybe_Just[int64] { panic("unreachable") }()
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
if (v1_4 == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_Maybe_Just[int64] { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}
go__816245670_2_1_37 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1170268447_3094389156(Call_local_Data_List_go__816245670_2_1_37(v_3_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_4_loop_val))))}
})
})
var go__go_3_4_38 gopurs_runtime.Value
_ = go__go_3_4_38
// FALLBACK TCO: isLoop=false len=1
go__go_3_4_38 = gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
var __t_tag_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5)
if (__t_tag_5 != nil) {
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(fn_0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0).IntVal) != (0) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(v_4.IntVal)}))}
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1170268447_3094389156(Call_local_Data_List_go__816245670_2_1_37((v_4.IntVal) + (int64(1)), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)))}
}
end_branch_6:
__t8 = __t6
goto end_branch_8
} else {

}
}
{
var __t_tag_7 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5)
if (__t_tag_7 == nil) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1170268447_3094389156(func() *Constructor_Data_Maybe_Just[int64] { panic("Failed pattern match") }()))}
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1170268447_3094389156(Rebox_Data_List_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t8))))}
})
var Call_local_Data_List_go__go_4_9_39 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_4_9_39
var go__go_4_9_39 gopurs_runtime.Value
_ = go__go_4_9_39
Call_local_Data_List_go__go_4_9_39 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_4_9_39:
for {
if false { continue go__go_4_9_39 }
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var __t10 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_6 == nil) {
__t10 = v_5
goto end_branch_10
} else {

}
}
{
if (v1_6 != nil) {
v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
v1_6_loop = (v1_6).V1
continue go__go_4_9_39
__t10 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_10
} else {

}
}
{
__t10 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
}
}
go__go_4_9_39 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_4_9_39(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
})
})
// TAST (Let): __local_var_2_0 shape=LetRec(LetRec(App(Other))) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
__local_var_2_0 := Call_local_Data_List_go__816245670_2_1_37(int64(0), Call_local_Data_List_go__go_4_9_39((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), xs_1))
_ = __local_var_2_0
var __t13 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_2_0 != nil) {
var Call_local_Data_List_go__go_3_11_40 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_go__go_3_11_40
var go__go_3_11_40 gopurs_runtime.Value
_ = go__go_3_11_40
Call_local_Data_List_go__go_3_11_40 = func(b_4_loop gopurs_runtime.Value, v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_3_11_40:
for {
if false { continue go__go_3_11_40 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var __t12 gopurs_runtime.Value
{
if (v_5 == nil) {
__t12 = b_4
goto end_branch_12
} else {

}
}
{
if (v_5 != nil) {
b_4_loop = gopurs_runtime.Int((b_4.IntVal) + (int64(1)))
v_5_loop = (v_5).V1
continue go__go_3_11_40
__t12 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_3_11_40 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_go__go_3_11_40(b_4_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val))
})
})
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int(((Call_local_Data_List_go__go_3_11_40(gopurs_runtime.Int(int64(0)), xs_1).IntVal) - (int64(1))) - (gopurs_runtime.Int((__local_var_2_0).V0).IntVal)), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_13:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t13)}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_List_filterM(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
filterM:
for {
if false { continue filterM }
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
var __t_tag_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_4)
if (__t_tag_2 == nil) {
__t7 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}))})
goto end_branch_7
} else {

}
}
{
var __t_tag_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_4)
if (__t_tag_3 != nil) {
// TAST (Let): __local_var_5_4 shape=Other bindingType=Any
__local_var_5_4 := (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V0
_ = __local_var_5_4
// TAST (Let): __local_var_6_5 shape=Other bindingType=Any
__local_var_6_5 := (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_6_5
__t7 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(v_3, __local_var_5_4), gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(Call_Data_List_filterM(dictMonad_0), v_3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__local_var_6_5)}), gopurs_runtime.Func(func(xs_prime__8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (b_7.IntVal) != (0) {
__t6 = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, __local_var_5_4, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_prime__8)})
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_prime__8)
}
end_branch_6:
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t6)})
}))
}))
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
})
}
}

func Call_Data_List_filter(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var Call_local_Data_List_go__go_1_0_41 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_1_0_41
var go__go_1_0_41 gopurs_runtime.Value
_ = go__go_1_0_41
Call_local_Data_List_go__go_1_0_41 = func(v_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_1_0_41:
for {
if false { continue go__go_1_0_41 }
var v_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_3 == nil) {
var Call_local_Data_List_go__go_4_1_42 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_4_1_42
var go__go_4_1_42 gopurs_runtime.Value
_ = go__go_4_1_42
Call_local_Data_List_go__go_4_1_42 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_4_1_42:
for {
if false { continue go__go_4_1_42 }
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var __t2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_6 == nil) {
__t2 = v_5
goto end_branch_2
} else {

}
}
{
if (v1_6 != nil) {
v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
v1_6_loop = (v1_6).V1
continue go__go_4_1_42
__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}
go__go_4_1_42 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_4_1_42(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
})
})
__t4 = Call_local_Data_List_go__go_4_1_42((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v_2)
goto end_branch_4
} else {

}
}
{
if (v1_3 != nil) {
var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply(p_0, (v1_3).V0).IntVal) != (0) {
v_2_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_3).V0, v_2})
v1_3_loop = (v1_3).V1
continue go__go_1_0_41
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_3
} else {

}
}
{
v_2_loop = v_2
v1_3_loop = (v1_3).V1
continue go__go_1_0_41
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}
go__go_1_0_41 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_1_0_41(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val)))}
})
})
return gopurs_runtime.Apply(go__go_1_0_41, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))})
}

func Call_Data_List_intersectBy(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v2_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_1 == nil) {
__t6 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_6
} else {

}
}
{
if (v2_2 == nil) {
__t6 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_6
} else {

}
}
{
var Call_local_Data_List_go__go_3_0_43 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_3_0_43
var go__go_3_0_43 gopurs_runtime.Value
_ = go__go_3_0_43
Call_local_Data_List_go__go_3_0_43 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_3_0_43:
for {
if false { continue go__go_3_0_43 }
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_5_loop
_ = v1_5
var __t5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_5 == nil) {
var Call_local_Data_List_go__go_6_1_44 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_6_1_44
var go__go_6_1_44 gopurs_runtime.Value
_ = go__go_6_1_44
Call_local_Data_List_go__go_6_1_44 = func(v_7_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_8_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_6_1_44:
for {
if false { continue go__go_6_1_44 }
var v_7 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_7_loop
_ = v_7
var v1_8 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_8_loop
_ = v1_8
var __t2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_8 == nil) {
__t2 = v_7
goto end_branch_2
} else {

}
}
{
if (v1_8 != nil) {
v_7_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_8).V0, v_7})
v1_8_loop = (v1_8).V1
continue go__go_6_1_44
__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}
go__go_6_1_44 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_6_1_44(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_7_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_8_loop_val)))}
})
})
__t5 = Call_local_Data_List_go__go_6_1_44((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v_4)
goto end_branch_5
} else {

}
}
{
if (v1_5 != nil) {
var __t4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
// TAST (Let): semigroupDisj1_6_3 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar a)])
semigroupDisj1_6_3 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((v_6.IntVal) != (0)) || ((v1_7.IntVal) != (0)))
})})
_ = semigroupDisj1_6_3
if (gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_foldableList()).V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupDisj1_6_3)}
}), gopurs_runtime.Bool(false)}))}, gopurs_runtime.Apply(v_0, (v1_5).V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v2_2)}).IntVal) != (0) {
v_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_5).V0, v_4})
v1_5_loop = (v1_5).V1
continue go__go_3_0_43
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_4
} else {

}
}
{
v_4_loop = v_4
v1_5_loop = (v1_5).V1
continue go__go_3_0_43
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}
go__go_3_0_43 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_3_0_43(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val)))}
})
})
__t6 = Call_local_Data_List_go__go_3_0_43((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v1_1)
}
end_branch_6:
return __t6
}

func Call_Data_List_intersect(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_List_intersectBy(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_List_nubByEq(v_0_loop gopurs_runtime.Value, v1_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
nubByEq:
for {
if false { continue nubByEq }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_1 == nil) {
__t6 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_6
} else {

}
}
{
if (v1_1 != nil) {
// TAST (Let): __local_var_2_0 shape=Other bindingType=Any
__local_var_2_0 := (v1_1).V0
_ = __local_var_2_0
var Call_local_Data_List_go__go_3_1_45 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_3_1_45
var go__go_3_1_45 gopurs_runtime.Value
_ = go__go_3_1_45
Call_local_Data_List_go__go_3_1_45 = func(v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_3_1_45:
for {
if false { continue go__go_3_1_45 }
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_5_loop
_ = v1_5
var __t5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_5 == nil) {
var Call_local_Data_List_go__go_6_2_46 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_6_2_46
var go__go_6_2_46 gopurs_runtime.Value
_ = go__go_6_2_46
Call_local_Data_List_go__go_6_2_46 = func(v_7_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_8_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_6_2_46:
for {
if false { continue go__go_6_2_46 }
var v_7 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_7_loop
_ = v_7
var v1_8 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_8_loop
_ = v1_8
var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_8 == nil) {
__t3 = v_7
goto end_branch_3
} else {

}
}
{
if (v1_8 != nil) {
v_7_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_8).V0, v_7})
v1_8_loop = (v1_8).V1
continue go__go_6_2_46
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}
go__go_6_2_46 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_6_2_46(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_7_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_8_loop_val)))}
})
})
__t5 = Call_local_Data_List_go__go_6_2_46((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), v_4)
goto end_branch_5
} else {

}
}
{
if (v1_5 != nil) {
var __t4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if ((gopurs_runtime.Apply2(v_0, __local_var_2_0, (v1_5).V0).IntVal) != (0)) != (true) {
v_4_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_5).V0, v_4})
v1_5_loop = (v1_5).V1
continue go__go_3_1_45
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_4
} else {

}
}
{
v_4_loop = v_4
v1_5_loop = (v1_5).V1
continue go__go_3_1_45
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}
go__go_3_1_45 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_3_1_45(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5_loop_val)))}
})
})
__t6 = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, __local_var_2_0, Call_Data_List_nubByEq(v_0, Call_local_Data_List_go__go_3_1_45((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), (v1_1).V1))})
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}
}

func Call_Data_List_nubEq(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_List_nubByEq(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_List_eqPattern(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
// TAST (Let): eqList_1_0 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","List","Types","List"] [(TypeVar a)])])
eqList_1_0 := (&Constructor_Data_Eq_Eq[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(xs_1 gopurs_runtime.Value, ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_47 gopurs_runtime.Value
_ = go__go_3_1_47
// FALLBACK TCO: isLoop=false len=1
go__go_3_1_47 = gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 bool
{
if ((v2_6.IntVal) != (0)) != (true) {
__t7 = false
goto end_branch_7
} else {

}
}
{
var __t_tag_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4)
if (__t_tag_5 == nil) {
var __t_tag_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5)
__t7 = ((__t_tag_6 == nil)) && ((v2_6.IntVal) != (0))
goto end_branch_7
} else {

}
}
{
var __t_tag_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4)
var __t_and_4 bool = false
if (__t_tag_2 != nil) {

var __t_tag_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5)
__t_and_4 = ((__t_tag_3 != nil)) && ((gopurs_runtime.Apply3(go__go_3_1_47, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}, gopurs_runtime.Bool(((v2_6.IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0).IntVal) != (0)))).IntVal) != (0))
}
__t7 = __t_and_4
}
end_branch_7:
return gopurs_runtime.Bool(__t7)
})
return gopurs_runtime.Bool((gopurs_runtime.Apply3(go__go_3_1_47, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_2))}, gopurs_runtime.Bool(true)).IntVal) != (0))
})})
_ = eqList_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1636902754_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(eqList_1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](x_2))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](y_3))}).IntVal) != (0))
})})))}
}

func Call_Data_List_ordPattern(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): ordList_1_0 shape=App(Var) bindingType=(ADT ["Data","Ord","Ord"] [(ADT ["Data","List","Types","List"] [(TypeVar a)])])
ordList_1_0 := Rebox_Data_List_4177771502_4210054658(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Types_ordList(), dictOrd_0)))
_ = ordList_1_0
// TAST (Let): __local_var_2_2 shape=App(Other) bindingType=Any
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): eqList_3_3 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","List","Types","List"] [(TypeVar a)])])
eqList_3_3 := (&Constructor_Data_Eq_Eq[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(xs_3 gopurs_runtime.Value, ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_4_48 gopurs_runtime.Value
_ = go__go_5_4_48
// FALLBACK TCO: isLoop=false len=1
go__go_5_4_48 = gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value, v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 bool
{
if ((v2_8.IntVal) != (0)) != (true) {
__t10 = false
goto end_branch_10
} else {

}
}
{
var __t_tag_8 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_6)
if (__t_tag_8 == nil) {
var __t_tag_9 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_7)
__t10 = ((__t_tag_9 == nil)) && ((v2_8.IntVal) != (0))
goto end_branch_10
} else {

}
}
{
var __t_tag_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_6)
var __t_and_7 bool = false
if (__t_tag_5 != nil) {

var __t_tag_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_7)
__t_and_7 = ((__t_tag_6 != nil)) && ((gopurs_runtime.Apply3(go__go_5_4_48, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_6.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_7.UnsafePtr).V1)}, gopurs_runtime.Bool(((v2_8.IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_2, "eq"), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_7.UnsafePtr).V0, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v_6.UnsafePtr).V0).IntVal) != (0)))).IntVal) != (0))
}
__t10 = __t_and_7
}
end_branch_10:
return gopurs_runtime.Bool(__t10)
})
return gopurs_runtime.Bool((gopurs_runtime.Apply3(go__go_5_4_48, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](ys_4))}, gopurs_runtime.Bool(true)).IntVal) != (0))
})})
_ = eqList_3_3
// TAST (Let): eqPattern1_2_1 shape=Let(Let(LitRecord)) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","List","Types","List"] [(TypeVar a)])])
eqPattern1_2_1 := (&Constructor_Data_Eq_Eq[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(eqList_3_3.V0), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](x_4))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](y_5))}).IntVal) != (0))
})})
_ = eqPattern1_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_List_4210054658_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1636902754_3790796878(eqPattern1_2_1))}
}), gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(ordList_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](x_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](y_4))}).IntVal)), UnsafePtr: nil}
})})))}
}

func Call_Data_List_elemLastIndex(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply(Get_Data_List_findLastIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), v_2, x_1).IntVal) != (0))
}))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_List_elemIndex(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var Call_local_Data_List_go__816245670_2_0_49 func(int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64]
_ = Call_local_Data_List_go__816245670_2_0_49
var go__816245670_2_0_49 gopurs_runtime.Value
_ = go__816245670_2_0_49
Call_local_Data_List_go__816245670_2_0_49 = func(v_3_loop int64, v1_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
go__816245670_2_0_49:
for {
if false { continue go__816245670_2_0_49 }
var v_3 int64 = v_3_loop
_ = v_3
var v1_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_4_loop
_ = v1_4
var __t2 *Constructor_Data_Maybe_Just[int64]
{
if (v1_4 != nil) {
var __t1 *Constructor_Data_Maybe_Just[int64]
{
if (gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), (v1_4).V0, x_1).IntVal) != (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int(v_3), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
v_3_loop = (v_3) + (int64(1))
v1_4_loop = (v1_4).V1
continue go__816245670_2_0_49
__t1 = func() *Constructor_Data_Maybe_Just[int64] { panic("unreachable") }()
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if (v1_4 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_Maybe_Just[int64] { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}
go__816245670_2_0_49 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1170268447_3094389156(Call_local_Data_List_go__816245670_2_0_49(v_3_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_4_loop_val))))}
})
})
var go__go_3_3_50 gopurs_runtime.Value
_ = go__go_3_3_50
// FALLBACK TCO: isLoop=false len=1
go__go_3_3_50 = gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
var __t_tag_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5)
if (__t_tag_4 != nil) {
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, x_1).IntVal) != (0) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(v_4.IntVal)}))}
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1170268447_3094389156(Call_local_Data_List_go__816245670_2_0_49((v_4.IntVal) + (int64(1)), (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)))}
}
end_branch_5:
__t7 = __t5
goto end_branch_7
} else {

}
}
{
var __t_tag_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5)
if (__t_tag_6 == nil) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1170268447_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1170268447_3094389156(func() *Constructor_Data_Maybe_Just[int64] { panic("Failed pattern match") }()))}
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1170268447_3094389156(Rebox_Data_List_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t7))))}
})
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1170268447_3094389156(Rebox_Data_List_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__816245670_2_0_49, gopurs_runtime.Int(int64(0)))))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_List_dropWhile(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var Call_local_Data_List_go__go_1_0_51 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_1_0_51
var go__go_1_0_51 gopurs_runtime.Value
_ = go__go_1_0_51
Call_local_Data_List_go__go_1_0_51 = func(v_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_1_0_51:
for {
if false { continue go__go_1_0_51 }
var v_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var __t1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if ((v_2 != nil)) && ((gopurs_runtime.Apply(p_0, (v_2).V0).IntVal) != (0)) {
v_2_loop = (v_2).V1
continue go__go_1_0_51
__t1 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
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
}
go__go_1_0_51 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_1_0_51(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_2_loop_val)))}
})
return go__go_1_0_51
}

func Call_Data_List_dropEnd(n_0_loop int64, xs_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = xs_1_loop
_ = xs_1
var Call_local_Data_List_go__go_2_0_52 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_go__go_2_0_52
var go__go_2_0_52 gopurs_runtime.Value
_ = go__go_2_0_52
Call_local_Data_List_go__go_2_0_52 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_0_52:
for {
if false { continue go__go_2_0_52 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4 == nil) {
__t1 = b_3
goto end_branch_1
} else {

}
}
{
if (v_4 != nil) {
b_3_loop = gopurs_runtime.Int((b_3.IntVal) + (int64(1)))
v_4_loop = (v_4).V1
continue go__go_2_0_52
__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_2_0_52 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_go__go_2_0_52(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_List_take(), gopurs_runtime.Int((Call_local_Data_List_go__go_2_0_52(gopurs_runtime.Int(int64(0)), xs_1).IntVal) - (n_0)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_1)}))
}

func Call_Data_List_drop(v_0_loop int64, v1_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
drop:
for {
if false { continue drop }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v_0) < (int64(1)) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (v1_1 == nil) {
__t0 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_0
} else {

}
}
{
if (v1_1 != nil) {
v_0_loop = (v_0) - (int64(1))
v1_1_loop = (v1_1).V1
continue drop
__t0 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_0
} else {

}
}
{
__t0 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}
}

func Call_Data_List_slice(start_0_loop int64, end_1_loop int64, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var start_0 int64 = start_0_loop
_ = start_0
var end_1 int64 = end_1_loop
_ = end_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_List_take(), gopurs_runtime.Int((end_1) - (start_0)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_drop(start_0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_2)))})))}
}

func Call_Data_List_takeEnd(n_0_loop int64, xs_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = xs_1_loop
_ = xs_1
var Call_local_Data_List_go__go_2_0_53 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_go__go_2_0_53
var go__go_2_0_53 gopurs_runtime.Value
_ = go__go_2_0_53
Call_local_Data_List_go__go_2_0_53 = func(b_3_loop gopurs_runtime.Value, v_4_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_2_0_53:
for {
if false { continue go__go_2_0_53 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var v_4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4 == nil) {
__t1 = b_3
goto end_branch_1
} else {

}
}
{
if (v_4 != nil) {
b_3_loop = gopurs_runtime.Int((b_3.IntVal) + (int64(1)))
v_4_loop = (v_4).V1
continue go__go_2_0_53
__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_2_0_53 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_go__go_2_0_53(b_3_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4_loop_val))
})
})
return Call_Data_List_drop((Call_local_Data_List_go__go_2_0_53(gopurs_runtime.Int(int64(0)), xs_1).IntVal) - (n_0), xs_1)
}

func Call_Data_List_deleteBy(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
deleteBy:
for {
if false { continue deleteBy }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v2_2 == nil) {
__t1 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_1
} else {

}
}
{
if (v2_2 != nil) {
var __t0 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply2(v_0, v1_1, (v2_2).V0).IntVal) != (0) {
__t0 = (v2_2).V1
goto end_branch_0
} else {

}
}
{
__t0 = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v2_2).V0, Call_Data_List_deleteBy(v_0, v1_1, (v2_2).V1)})
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
__t1 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}

func Call_Data_List_unionBy(eq_0_loop gopurs_runtime.Value, xs_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], ys_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = xs_1_loop
_ = xs_1
var ys_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = ys_2_loop
_ = ys_2
var Call_local_Data_List_go__go_3_0_54 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_go__go_3_0_54
var go__go_3_0_54 gopurs_runtime.Value
_ = go__go_3_0_54
Call_local_Data_List_go__go_3_0_54 = func(b_4_loop gopurs_runtime.Value, v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_3_0_54:
for {
if false { continue go__go_3_0_54 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var __t1 gopurs_runtime.Value
{
if (v_5 == nil) {
__t1 = b_4
goto end_branch_1
} else {

}
}
{
if (v_5 != nil) {
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_5).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_4)}))}
v_5_loop = (v_5).V1
continue go__go_3_0_54
__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_3_0_54 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_go__go_3_0_54(b_4_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val))
})
})
var Call_local_Data_List_go__go_4_2_55 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_go__go_4_2_55
var go__go_4_2_55 gopurs_runtime.Value
_ = go__go_4_2_55
Call_local_Data_List_go__go_4_2_55 = func(b_5_loop gopurs_runtime.Value, v_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_4_2_55:
for {
if false { continue go__go_4_2_55 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_6_loop
_ = v_6
var __t3 gopurs_runtime.Value
{
if (v_6 == nil) {
__t3 = b_5
goto end_branch_3
} else {

}
}
{
if (v_6 != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_deleteBy(eq_0, (v_6).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_5)))}
v_6_loop = (v_6).V1
continue go__go_4_2_55
__t3 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_4_2_55 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_go__go_4_2_55(b_5_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_6_loop_val))
})
})
var Call_local_Data_List_go__go_4_4_56 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_go__go_4_4_56
var go__go_4_4_56 gopurs_runtime.Value
_ = go__go_4_4_56
Call_local_Data_List_go__go_4_4_56 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_4_4_56:
for {
if false { continue go__go_4_4_56 }
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var __t5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_6 == nil) {
__t5 = v_5
goto end_branch_5
} else {

}
}
{
if (v1_6 != nil) {
v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
v1_6_loop = (v1_6).V1
continue go__go_4_4_56
__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}
go__go_4_4_56 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_List_go__go_4_4_56(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
})
})
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_local_Data_List_go__go_3_0_54(Call_local_Data_List_go__go_4_2_55(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_nubByEq(eq_0, ys_2))}, xs_1), Call_local_Data_List_go__go_4_4_56((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), xs_1)))
}

func Call_Data_List_union(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_List_unionBy(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_List_deleteAt(v_0_loop int64, v1_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
deleteAt:
for {
if false { continue deleteAt }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t3 gopurs_runtime.Value
{
if (v1_1 != nil) {
var __t2 gopurs_runtime.Value
{
if (v_0) == (int64(0)) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1978018568_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v1_1).V1)}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
goto end_branch_2
} else {

}
}
{
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","List","Types","List"] [(TypeVar a)])])
__local_var_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := Call_Data_List_deleteAt((v_0) - (int64(1)), (v1_1).V1)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = __local_var_2_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_2_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_1).V0, (__local_var_2_0).V0}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1978018568_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_3:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1978018568_3094389156(Rebox_Data_List_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t3))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_List_go__delete(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_List_deleteBy(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_List_difference(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var Call_local_Data_List_go__go_1_0_57 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_go__go_1_0_57
var go__go_1_0_57 gopurs_runtime.Value
_ = go__go_1_0_57
Call_local_Data_List_go__go_1_0_57 = func(b_2_loop gopurs_runtime.Value, v_3_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_1_0_57:
for {
if false { continue go__go_1_0_57 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var v_3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (v_3 == nil) {
__t1 = b_2
goto end_branch_1
} else {

}
}
{
if (v_3 != nil) {
b_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_deleteBy(gopurs_runtime.Box(dictEq_0.V0), (v_3).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_2)))}
v_3_loop = (v_3).V1
continue go__go_1_0_57
__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_1_0_57 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_go__go_1_0_57(b_2_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_3_loop_val))
})
})
return go__go_1_0_57
}

func Call_Data_List_concatMap(b_0_loop gopurs_runtime.Value, a_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = a_1_loop
_ = a_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_bindList()).V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(a_1)}, b_0))
}

func Call_Data_List_concat(v_0_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
var v_0 *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_bindList()).V1), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1116310629_849153993(v_0))}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)))
}

func Call_Data_List_alterAt(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
alterAt:
for {
if false { continue alterAt }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t5 gopurs_runtime.Value
{
if (v2_2 != nil) {
var __t4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v_0) == (int64(0)) {
// TAST (Let): v3_3_2 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
v3_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (v2_2).V0))
_ = v3_3_2
var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v3_3_2 == nil) {
__t3 = (v2_2).V1
goto end_branch_3
} else {

}
}
{
if (v3_3_2 != nil) {
__t3 = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v3_3_2).V0, (v2_2).V1})
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}})
goto end_branch_4
} else {

}
}
{
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","List","Types","List"] [(TypeVar a)])])
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := Call_Data_List_alterAt((v_0) - (int64(1)), v1_1, (v2_2).V1)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = __local_var_3_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_3_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v2_2).V0, (__local_var_3_0).V0}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_1:
__t4 = __t1
}
end_branch_4:
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1978018568_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}
end_branch_5:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_1978018568_3094389156(Rebox_Data_List_3094389156_1978018568(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t5))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_List_modifyAt(n_0_loop int64, f_1_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var n_0 int64 = n_0_loop
_ = n_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply2(Get_Data_List_alterAt(), gopurs_runtime.Int(n_0), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, x_2)}))}
}))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Rebox_Data_List_1116310629_849153993(in *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(in.V0)}
		out.V1 = Rebox_Data_List_1116310629_849153993(in.V1)
	return out
}

func Rebox_Data_List_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_List_1293498952_3123684004(in *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_138441832_22134120(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](in.V0)
		out.V1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](in.V1)
	return out
}

func Rebox_Data_List_138441832_3132786365(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[uint32, float64] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[uint32, float64]{}
		out.V0 = uint32(in.V0.IntVal)
		out.V1 = in.V1.FloatVal()
	return out
}

func Rebox_Data_List_1386611502_3351995458(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_1501865320_385277032(in *Constructor_Data_Newtype_Newtype[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_1636902754_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_1978018568_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_List_22134120_138441832(in *Constructor_Data_Tuple_Tuple[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(in.V0)}
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}

func Rebox_Data_List_2351501316_138441832(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}

func Rebox_Data_List_2442833393_849153993(in *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_3132786365_138441832(in.V0))}
		out.V1 = Rebox_Data_List_2442833393_849153993(in.V1)
	return out
}

func Rebox_Data_List_2681346401_3094389156(in *Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = func() gopurs_runtime.Value {
				orig := in.V0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"head", "tail"}, []gopurs_runtime.Value{orig.head, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.tail)}})
				}()
	return out
}

func Rebox_Data_List_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[int64]{}
		out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_List_3094389156_1978018568(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](in.V0)
	return out
}

func Rebox_Data_List_3094389156_3270864776(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct{
	last gopurs_runtime.Value
	revInit *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[struct{
	last gopurs_runtime.Value
	revInit *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}]{}
		out.V0 = func() struct{
	last gopurs_runtime.Value
	revInit *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
} {
					orig := in.V0
					_ = orig
					clone := struct{
	last gopurs_runtime.Value
	revInit *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}{}
					clone.last = gopurs_runtime.RecordGet(orig, "last")
					clone.revInit = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "revInit"))
					return clone
				}()
	return out
}

func Rebox_Data_List_3123684004_1293498952(in *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_3132786365_138441832(in *Constructor_Data_Tuple_Tuple[uint32, float64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
		out.V1 = gopurs_runtime.Float(in.V1)
	return out
}

func Rebox_Data_List_3266529732_849153993(in *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V0)}
		out.V1 = Rebox_Data_List_3266529732_849153993(in.V1)
	return out
}

func Rebox_Data_List_3270864776_3094389156(in *Constructor_Data_Maybe_Just[struct{
	last gopurs_runtime.Value
	revInit *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = func() gopurs_runtime.Value {
				orig := in.V0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"last", "revInit"}, []gopurs_runtime.Value{orig.last, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.revInit)}})
				}()
	return out
}

func Rebox_Data_List_3351995458_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_3406595152_4008603408(in *Constructor_Control_Monad_Rec_Class_Loop[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_List_3704040722_849153993(in *Constructor_Data_List_Types_Cons[int64]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
		out.V1 = Rebox_Data_List_3704040722_849153993(in.V1)
	return out
}

func Rebox_Data_List_3718343566_4008603408(in *Constructor_Control_Monad_Rec_Class_Loop[struct{
	a *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	b *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = func() gopurs_runtime.Value {
				orig := in.V0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"a", "b"}, []gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.a)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(orig.b)}})
				}()
	return out
}

func Rebox_Data_List_4177771502_4210054658(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_4210054658_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_List_4239369586_3603546092(in *Constructor_Control_Monad_Rec_Class_Done[struct{
	a *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
	b *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
}, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]) *Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_List_587109416_849153993(in *Constructor_Data_List_Types_Cons[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_List_3123684004_1293498952(in.V0))}
		out.V1 = Rebox_Data_List_587109416_849153993(in.V1)
	return out
}

func Rebox_Data_List_802708012_3603546092(in *Constructor_Control_Monad_Rec_Class_Done[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_List_849153993_1116310629(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](in.V0)
		out.V1 = Rebox_Data_List_849153993_1116310629(in.V1)
	return out
}

func Rebox_Data_List_849153993_2442833393(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]{}
		out.V0 = Rebox_Data_List_138441832_3132786365(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
		out.V1 = Rebox_Data_List_849153993_2442833393(in.V1)
	return out
}

func Rebox_Data_List_849153993_3266529732(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0)
		out.V1 = Rebox_Data_List_849153993_3266529732(in.V1)
	return out
}

func Rebox_Data_List_849153993_3704040722(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[int64]{}
		out.V0 = in.V0.IntVal
		out.V1 = Rebox_Data_List_849153993_3704040722(in.V1)
	return out
}



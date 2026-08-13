package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_List_Lazy_unwrap gopurs_runtime.Value
var once_Data_List_Lazy_unwrap sync.Once
func Get_Data_List_Lazy_unwrap() gopurs_runtime.Value {
	once_Data_List_Lazy_unwrap.Do(func() {
		cache_Data_List_Lazy_unwrap = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_List_Lazy_unwrap
}

var cache_Data_List_Lazy_one gopurs_runtime.Value
var once_Data_List_Lazy_one sync.Once
func Get_Data_List_Lazy_one() gopurs_runtime.Value {
	once_Data_List_Lazy_one.Do(func() {
		cache_Data_List_Lazy_one = gopurs_runtime.Int(1)
	})
	return cache_Data_List_Lazy_one
}

var cache_Data_List_Lazy_identity gopurs_runtime.Value
var once_Data_List_Lazy_identity sync.Once
func Get_Data_List_Lazy_identity() gopurs_runtime.Value {
	once_Data_List_Lazy_identity.Do(func() {
		cache_Data_List_Lazy_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_identity(x_0_box)
})
	})
	return cache_Data_List_Lazy_identity
}

var cache_Data_List_Lazy_Pattern gopurs_runtime.Value
var once_Data_List_Lazy_Pattern sync.Once
func Get_Data_List_Lazy_Pattern() gopurs_runtime.Value {
	once_Data_List_Lazy_Pattern.Do(func() {
		cache_Data_List_Lazy_Pattern = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_Pattern(x_0_box)
})
	})
	return cache_Data_List_Lazy_Pattern
}

var cache_Data_List_Lazy_zipWith gopurs_runtime.Value
var once_Data_List_Lazy_zipWith sync.Once
func Get_Data_List_Lazy_zipWith() gopurs_runtime.Value {
	once_Data_List_Lazy_zipWith.Do(func() {
		cache_Data_List_Lazy_zipWith = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_zipWith(f_0_box, xs_1_box, ys_2_box)
})
	})
	return cache_Data_List_Lazy_zipWith
}

var cache_Data_List_Lazy_zipWithA gopurs_runtime.Value
var once_Data_List_Lazy_zipWithA sync.Once
func Get_Data_List_Lazy_zipWithA() gopurs_runtime.Value {
	once_Data_List_Lazy_zipWithA.Do(func() {
		cache_Data_List_Lazy_zipWithA = gopurs_runtime.Func4(func(dictApplicative_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value, ys_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_zipWithA(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box), f_1_box, xs_2_box, ys_3_box)
})
	})
	return cache_Data_List_Lazy_zipWithA
}

var cache_Data_List_Lazy_zip gopurs_runtime.Value
var once_Data_List_Lazy_zip sync.Once
func Get_Data_List_Lazy_zip() gopurs_runtime.Value {
	once_Data_List_Lazy_zip.Do(func() {
		cache_Data_List_Lazy_zip = gopurs_runtime.Apply(Get_Data_List_Lazy_zipWith(), Get_Data_Tuple_Tuple())
	})
	return cache_Data_List_Lazy_zip
}

var cache_Data_List_Lazy_updateAt gopurs_runtime.Value
var once_Data_List_Lazy_updateAt sync.Once
func Get_Data_List_Lazy_updateAt() gopurs_runtime.Value {
	once_Data_List_Lazy_updateAt.Do(func() {
		cache_Data_List_Lazy_updateAt = gopurs_runtime.Func3(func(n_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_updateAt(n_0_box.IntVal, x_1_box, xs_2_box)
})
	})
	return cache_Data_List_Lazy_updateAt
}

var cache_Data_List_Lazy_unzip gopurs_runtime.Value
var once_Data_List_Lazy_unzip sync.Once
func Get_Data_List_Lazy_unzip() gopurs_runtime.Value {
	once_Data_List_Lazy_unzip.Do(func() {
		cache_Data_List_Lazy_unzip = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_unzip(xs_0_box))}
})
	})
	return cache_Data_List_Lazy_unzip
}

var cache_Data_List_Lazy_uncons gopurs_runtime.Value
var once_Data_List_Lazy_uncons sync.Once
func Get_Data_List_Lazy_uncons() gopurs_runtime.Value {
	once_Data_List_Lazy_uncons.Do(func() {
		cache_Data_List_Lazy_uncons = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_uncons(xs_0_box))}
})
	})
	return cache_Data_List_Lazy_uncons
}

var cache_Data_List_Lazy_toUnfoldable gopurs_runtime.Value
var once_Data_List_Lazy_toUnfoldable sync.Once
func Get_Data_List_Lazy_toUnfoldable() gopurs_runtime.Value {
	once_Data_List_Lazy_toUnfoldable.Do(func() {
		cache_Data_List_Lazy_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dictUnfoldable_0_box))
})
	})
	return cache_Data_List_Lazy_toUnfoldable
}

var cache_Data_List_Lazy_takeWhile gopurs_runtime.Value
var once_Data_List_Lazy_takeWhile sync.Once
func Get_Data_List_Lazy_takeWhile() gopurs_runtime.Value {
	once_Data_List_Lazy_takeWhile.Do(func() {
		cache_Data_List_Lazy_takeWhile = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_takeWhile(p_0_box, x_1_box)
})
	})
	return cache_Data_List_Lazy_takeWhile
}

var cache_Data_List_Lazy_take gopurs_runtime.Value
var once_Data_List_Lazy_take sync.Once
func Get_Data_List_Lazy_take() gopurs_runtime.Value {
	once_Data_List_Lazy_take.Do(func() {
		cache_Data_List_Lazy_take = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_take(n_0_box.IntVal)
})
	})
	return cache_Data_List_Lazy_take
}

var cache_Data_List_Lazy_tail gopurs_runtime.Value
var once_Data_List_Lazy_tail sync.Once
func Get_Data_List_Lazy_tail() gopurs_runtime.Value {
	once_Data_List_Lazy_tail.Do(func() {
		cache_Data_List_Lazy_tail = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_tail(xs_0_box))}
})
	})
	return cache_Data_List_Lazy_tail
}

var cache_Data_List_Lazy_stripPrefix gopurs_runtime.Value
var once_Data_List_Lazy_stripPrefix sync.Once
func Get_Data_List_Lazy_stripPrefix() gopurs_runtime.Value {
	once_Data_List_Lazy_stripPrefix.Do(func() {
		cache_Data_List_Lazy_stripPrefix = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_stripPrefix(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), v_1_box, s_2_box))}
})
	})
	return cache_Data_List_Lazy_stripPrefix
}

var cache_Data_List_Lazy_span gopurs_runtime.Value
var once_Data_List_Lazy_span sync.Once
func Get_Data_List_Lazy_span() gopurs_runtime.Value {
	once_Data_List_Lazy_span.Do(func() {
		cache_Data_List_Lazy_span = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_span(p_0_box, xs_1_box)
})
	})
	return cache_Data_List_Lazy_span
}

var cache_Data_List_Lazy_snoc gopurs_runtime.Value
var once_Data_List_Lazy_snoc sync.Once
func Get_Data_List_Lazy_snoc() gopurs_runtime.Value {
	once_Data_List_Lazy_snoc.Do(func() {
		cache_Data_List_Lazy_snoc = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_snoc(xs_0_box, x_1_box)
})
	})
	return cache_Data_List_Lazy_snoc
}

var cache_Data_List_Lazy_singleton gopurs_runtime.Value
var once_Data_List_Lazy_singleton sync.Once
func Get_Data_List_Lazy_singleton() gopurs_runtime.Value {
	once_Data_List_Lazy_singleton.Do(func() {
		cache_Data_List_Lazy_singleton = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_singleton(a_0_box)
})
	})
	return cache_Data_List_Lazy_singleton
}

var cache_Data_List_Lazy_showPattern gopurs_runtime.Value
var once_Data_List_Lazy_showPattern sync.Once
func Get_Data_List_Lazy_showPattern() gopurs_runtime.Value {
	once_Data_List_Lazy_showPattern.Do(func() {
		cache_Data_List_Lazy_showPattern = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_showPattern(dictShow_0_box)
})
	})
	return cache_Data_List_Lazy_showPattern
}

var cache_Data_List_Lazy_scanlLazy gopurs_runtime.Value
var once_Data_List_Lazy_scanlLazy sync.Once
func Get_Data_List_Lazy_scanlLazy() gopurs_runtime.Value {
	once_Data_List_Lazy_scanlLazy.Do(func() {
		cache_Data_List_Lazy_scanlLazy = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, acc_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_scanlLazy(f_0_box, acc_1_box, xs_2_box)
})
	})
	return cache_Data_List_Lazy_scanlLazy
}

var cache_Data_List_Lazy_reverse gopurs_runtime.Value
var once_Data_List_Lazy_reverse sync.Once
func Get_Data_List_Lazy_reverse() gopurs_runtime.Value {
	once_Data_List_Lazy_reverse.Do(func() {
		cache_Data_List_Lazy_reverse = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_reverse(xs_0_box)
})
	})
	return cache_Data_List_Lazy_reverse
}

var cache_Data_List_Lazy_replicateM gopurs_runtime.Value
var once_Data_List_Lazy_replicateM sync.Once
func Get_Data_List_Lazy_replicateM() gopurs_runtime.Value {
	once_Data_List_Lazy_replicateM.Do(func() {
		cache_Data_List_Lazy_replicateM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_replicateM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Data_List_Lazy_replicateM
}

var cache_Data_List_Lazy_repeat gopurs_runtime.Value
var once_Data_List_Lazy_repeat sync.Once
func Get_Data_List_Lazy_repeat() gopurs_runtime.Value {
	once_Data_List_Lazy_repeat.Do(func() {
		cache_Data_List_Lazy_repeat = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_repeat(x_0_box)
})
	})
	return cache_Data_List_Lazy_repeat
}

var cache_Data_List_Lazy_replicate gopurs_runtime.Value
var once_Data_List_Lazy_replicate sync.Once
func Get_Data_List_Lazy_replicate() gopurs_runtime.Value {
	once_Data_List_Lazy_replicate.Do(func() {
		cache_Data_List_Lazy_replicate = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_replicate(i_0_box.IntVal, xs_1_box)
})
	})
	return cache_Data_List_Lazy_replicate
}

var cache_Data_List_Lazy_go__range gopurs_runtime.Value
var once_Data_List_Lazy_go__range sync.Once
func Get_Data_List_Lazy_go__range() gopurs_runtime.Value {
	once_Data_List_Lazy_go__range.Do(func() {
		cache_Data_List_Lazy_go__range = gopurs_runtime.Func2(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_go__range(start_0_box.IntVal, end_1_box.IntVal)
})
	})
	return cache_Data_List_Lazy_go__range
}

var cache_Data_List_Lazy_partition gopurs_runtime.Value
var once_Data_List_Lazy_partition sync.Once
func Get_Data_List_Lazy_partition() gopurs_runtime.Value {
	once_Data_List_Lazy_partition.Do(func() {
		cache_Data_List_Lazy_partition = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_partition(f_0_box, xs_1_box)
})
	})
	return cache_Data_List_Lazy_partition
}

var cache_Data_List_Lazy_null gopurs_runtime.Value
var once_Data_List_Lazy_null sync.Once
func Get_Data_List_Lazy_null() gopurs_runtime.Value {
	once_Data_List_Lazy_null.Do(func() {
		cache_Data_List_Lazy_null = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_List_Lazy_null(x_0_box))
})
	})
	return cache_Data_List_Lazy_null
}

var cache_Data_List_Lazy_nubBy gopurs_runtime.Value
var once_Data_List_Lazy_nubBy sync.Once
func Get_Data_List_Lazy_nubBy() gopurs_runtime.Value {
	once_Data_List_Lazy_nubBy.Do(func() {
		cache_Data_List_Lazy_nubBy = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_nubBy(p_0_box)
})
	})
	return cache_Data_List_Lazy_nubBy
}

var cache_Data_List_Lazy_nub gopurs_runtime.Value
var once_Data_List_Lazy_nub sync.Once
func Get_Data_List_Lazy_nub() gopurs_runtime.Value {
	once_Data_List_Lazy_nub.Do(func() {
		cache_Data_List_Lazy_nub = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_nub(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_List_Lazy_nub
}

var cache_Data_List_Lazy_newtypePattern gopurs_runtime.Value
var once_Data_List_Lazy_newtypePattern sync.Once
func Get_Data_List_Lazy_newtypePattern() gopurs_runtime.Value {
	once_Data_List_Lazy_newtypePattern.Do(func() {
		cache_Data_List_Lazy_newtypePattern = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(&Constructor_Data_Newtype_Newtype{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})}
	})
	return cache_Data_List_Lazy_newtypePattern
}

var cache_Data_List_Lazy_mapMaybe gopurs_runtime.Value
var once_Data_List_Lazy_mapMaybe sync.Once
func Get_Data_List_Lazy_mapMaybe() gopurs_runtime.Value {
	once_Data_List_Lazy_mapMaybe.Do(func() {
		cache_Data_List_Lazy_mapMaybe = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_mapMaybe(f_0_box)
})
	})
	return cache_Data_List_Lazy_mapMaybe
}

var cache_Data_List_Lazy_some gopurs_runtime.Value
var once_Data_List_Lazy_some sync.Once
func Get_Data_List_Lazy_some() gopurs_runtime.Value {
	once_Data_List_Lazy_some.Do(func() {
		cache_Data_List_Lazy_some = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_some(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative](dictAlternative_0_box))
})
	})
	return cache_Data_List_Lazy_some
}

var cache_Data_List_Lazy_many gopurs_runtime.Value
var once_Data_List_Lazy_many sync.Once
func Get_Data_List_Lazy_many() gopurs_runtime.Value {
	once_Data_List_Lazy_many.Do(func() {
		cache_Data_List_Lazy_many = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_many(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative](dictAlternative_0_box))
})
	})
	return cache_Data_List_Lazy_many
}

var cache_Data_List_Lazy_length gopurs_runtime.Value
var once_Data_List_Lazy_length sync.Once
func Get_Data_List_Lazy_length() gopurs_runtime.Value {
	once_Data_List_Lazy_length.Do(func() {
		cache_Data_List_Lazy_length = func() gopurs_runtime.Value {
var go__go_0_0_15 gopurs_runtime.Value
go__go_0_0_15 = gopurs_runtime.Func(func(b_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_1_loop gopurs_runtime.Value = b_1_loop_val
var xs_2_loop gopurs_runtime.Value = xs_2_loop_val
go__go_0_0_15:
for {
if false { continue go__go_0_0_15 }
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
// TAST (Let): v_3_1 -> *Constructor_Data_List_Lazy_Types_Cons
v_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2))
_ = v_3_1
var __t2 gopurs_runtime.Value
{
if (v_3_1 == nil) {
__t2 = b_1
goto end_branch_2
} else {

}
}
{
if (v_3_1 != nil) {
b_1_loop = gopurs_runtime.Int((b_1.IntVal) + (1))
xs_2_loop = (v_3_1).V1
continue go__go_0_0_15
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
return gopurs_runtime.Apply(go__go_0_0_15, gopurs_runtime.Int(0))
}()
	})
	return cache_Data_List_Lazy_length
}

var cache_Data_List_Lazy_last gopurs_runtime.Value
var once_Data_List_Lazy_last sync.Once
func Get_Data_List_Lazy_last() gopurs_runtime.Value {
	once_Data_List_Lazy_last.Do(func() {
		cache_Data_List_Lazy_last = func() gopurs_runtime.Value {
var go__go_0_0_16 gopurs_runtime.Value
go__go_0_0_16 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop gopurs_runtime.Value = v_1_loop_val
go__go_0_0_16:
for {
if false { continue go__go_0_0_16 }
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var __t4 *Constructor_Data_Maybe_Just
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil) {
var __t3 *Constructor_Data_Maybe_Just
{
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
var __local_var_2_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_uncons((*Constructor_Data_List_Lazy_Types_Cons)(v_1.UnsafePtr).V1))}
var __t2 gopurs_runtime.Value
{
if (__local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 930809136 && __local_var_2_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Bool(true)
goto end_branch_2
} else {

}
}
{
if (__local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 930809136 && __local_var_2_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.Bool(false)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
if (__t2.IntVal) != (0) {
__t3 = &Constructor_Data_Maybe_Just{1, (*Constructor_Data_List_Lazy_Types_Cons)(v_1.UnsafePtr).V0}
goto end_branch_3
} else {

}
}
{
v_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_1.UnsafePtr).V1)))}
continue go__go_0_0_16
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
}
}()
})
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_0_0_16, gopurs_runtime.Apply(Get_Data_Lazy_force(), x_1))
})
}()
	})
	return cache_Data_List_Lazy_last
}

var cache_Data_List_Lazy_iterate gopurs_runtime.Value
var once_Data_List_Lazy_iterate sync.Once
func Get_Data_List_Lazy_iterate() gopurs_runtime.Value {
	once_Data_List_Lazy_iterate.Do(func() {
		cache_Data_List_Lazy_iterate = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_iterate(f_0_box, x_1_box)
})
	})
	return cache_Data_List_Lazy_iterate
}

var cache_Data_List_Lazy_insertAt gopurs_runtime.Value
var once_Data_List_Lazy_insertAt sync.Once
func Get_Data_List_Lazy_insertAt() gopurs_runtime.Value {
	once_Data_List_Lazy_insertAt.Do(func() {
		cache_Data_List_Lazy_insertAt = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_insertAt(v_0_box.IntVal, v1_1_box, v2_2_box)
})
	})
	return cache_Data_List_Lazy_insertAt
}

var cache_Data_List_Lazy_init gopurs_runtime.Value
var once_Data_List_Lazy_init sync.Once
func Get_Data_List_Lazy_init() gopurs_runtime.Value {
	once_Data_List_Lazy_init.Do(func() {
		cache_Data_List_Lazy_init = func() gopurs_runtime.Value {
var go__go_0_0_18 gopurs_runtime.Value
_ = go__go_0_0_18
go__go_0_0_18 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 *Constructor_Data_Maybe_Just
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil) {
var __t7 *Constructor_Data_Maybe_Just
{
// TAST (Let): __local_var_2_5 -> gopurs_runtime.Value
var __local_var_2_5 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_uncons((*Constructor_Data_List_Lazy_Types_Cons)(v_1.UnsafePtr).V1))}
var __t6 gopurs_runtime.Value
{
if (__local_var_2_5.Type == 9 && __local_var_2_5.IntVal == 930809136 && __local_var_2_5.UnsafePtr == nil) {
__t6 = gopurs_runtime.Bool(true)
goto end_branch_6
} else {

}
}
{
if (__local_var_2_5.Type == 9 && __local_var_2_5.IntVal == 930809136 && __local_var_2_5.UnsafePtr != nil) {
__t6 = gopurs_runtime.Bool(false)
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
if (__t6.IntVal) != (0) {
__t7 = &Constructor_Data_Maybe_Just{1, Get_Data_List_Lazy_Types_nil()}
goto end_branch_7
} else {

}
}
{
// TAST (Let): __local_var_2_1 -> *Constructor_Data_Maybe_Just
__local_var_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(go__go_0_0_18, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_1.UnsafePtr).V1)))}))
_ = __local_var_2_1
var __t4 *Constructor_Data_Maybe_Just
{
if (__local_var_2_1 != nil) {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := (*Constructor_Data_List_Lazy_Types_Cons)(v_1.UnsafePtr).V0
_ = __local_var_3_2
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := (__local_var_2_1).V0
_ = __local_var_4_3
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_3_2, __local_var_4_3})}
}))}
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_4:
__t7 = __t4
}
end_branch_7:
__t8 = __t7
goto end_branch_8
} else {

}
}
{
__t8 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t8)}
})
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_0_0_18, gopurs_runtime.Apply(Get_Data_Lazy_force(), x_1))
})
}()
	})
	return cache_Data_List_Lazy_init
}

var cache_Data_List_Lazy_index gopurs_runtime.Value
var once_Data_List_Lazy_index sync.Once
func Get_Data_List_Lazy_index() gopurs_runtime.Value {
	once_Data_List_Lazy_index.Do(func() {
		cache_Data_List_Lazy_index = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_index(xs_0_box)
})
	})
	return cache_Data_List_Lazy_index
}

var cache_Data_List_Lazy_head gopurs_runtime.Value
var once_Data_List_Lazy_head sync.Once
func Get_Data_List_Lazy_head() gopurs_runtime.Value {
	once_Data_List_Lazy_head.Do(func() {
		cache_Data_List_Lazy_head = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_head(xs_0_box))}
})
	})
	return cache_Data_List_Lazy_head
}

var cache_Data_List_Lazy_transpose gopurs_runtime.Value
var once_Data_List_Lazy_transpose sync.Once
func Get_Data_List_Lazy_transpose() gopurs_runtime.Value {
	once_Data_List_Lazy_transpose.Do(func() {
		cache_Data_List_Lazy_transpose = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_transpose(xs_0_box)
})
	})
	return cache_Data_List_Lazy_transpose
}

var cache_Data_List_Lazy_groupBy gopurs_runtime.Value
var once_Data_List_Lazy_groupBy sync.Once
func Get_Data_List_Lazy_groupBy() gopurs_runtime.Value {
	once_Data_List_Lazy_groupBy.Do(func() {
		cache_Data_List_Lazy_groupBy = gopurs_runtime.Func2(func(eq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_groupBy(eq_0_box, x_1_box)
})
	})
	return cache_Data_List_Lazy_groupBy
}

var cache_Data_List_Lazy_group gopurs_runtime.Value
var once_Data_List_Lazy_group sync.Once
func Get_Data_List_Lazy_group() gopurs_runtime.Value {
	once_Data_List_Lazy_group.Do(func() {
		cache_Data_List_Lazy_group = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_group(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_List_Lazy_group
}

var cache_Data_List_Lazy_fromStep gopurs_runtime.Value
var once_Data_List_Lazy_fromStep sync.Once
func Get_Data_List_Lazy_fromStep() gopurs_runtime.Value {
	once_Data_List_Lazy_fromStep.Do(func() {
		cache_Data_List_Lazy_fromStep = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_fromStep(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](x_0_box))
})
	})
	return cache_Data_List_Lazy_fromStep
}

var cache_Data_List_Lazy_insertBy gopurs_runtime.Value
var once_Data_List_Lazy_insertBy sync.Once
func Get_Data_List_Lazy_insertBy() gopurs_runtime.Value {
	once_Data_List_Lazy_insertBy.Do(func() {
		cache_Data_List_Lazy_insertBy = gopurs_runtime.Func3(func(cmp_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_insertBy(cmp_0_box, x_1_box, xs_2_box)
})
	})
	return cache_Data_List_Lazy_insertBy
}

var cache_Data_List_Lazy_insert gopurs_runtime.Value
var once_Data_List_Lazy_insert sync.Once
func Get_Data_List_Lazy_insert() gopurs_runtime.Value {
	once_Data_List_Lazy_insert.Do(func() {
		cache_Data_List_Lazy_insert = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_insert(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_List_Lazy_insert
}

var cache_Data_List_Lazy_fromFoldable gopurs_runtime.Value
var once_Data_List_Lazy_fromFoldable sync.Once
func Get_Data_List_Lazy_fromFoldable() gopurs_runtime.Value {
	once_Data_List_Lazy_fromFoldable.Do(func() {
		cache_Data_List_Lazy_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box))
})
	})
	return cache_Data_List_Lazy_fromFoldable
}

var cache_Data_List_Lazy_foldrLazy gopurs_runtime.Value
var once_Data_List_Lazy_foldrLazy sync.Once
func Get_Data_List_Lazy_foldrLazy() gopurs_runtime.Value {
	once_Data_List_Lazy_foldrLazy.Do(func() {
		cache_Data_List_Lazy_foldrLazy = gopurs_runtime.Func3(func(dictLazy_0_box gopurs_runtime.Value, op_1_box gopurs_runtime.Value, z_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_foldrLazy(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy](dictLazy_0_box), op_1_box, z_2_box)
})
	})
	return cache_Data_List_Lazy_foldrLazy
}

var cache_Data_List_Lazy_foldM gopurs_runtime.Value
var once_Data_List_Lazy_foldM sync.Once
func Get_Data_List_Lazy_foldM() gopurs_runtime.Value {
	once_Data_List_Lazy_foldM.Do(func() {
		cache_Data_List_Lazy_foldM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_foldM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Data_List_Lazy_foldM
}

var cache_Data_List_Lazy_findIndex gopurs_runtime.Value
var once_Data_List_Lazy_findIndex sync.Once
func Get_Data_List_Lazy_findIndex() gopurs_runtime.Value {
	once_Data_List_Lazy_findIndex.Do(func() {
		cache_Data_List_Lazy_findIndex = gopurs_runtime.Func(func(fn_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_findIndex(fn_0_box)
})
	})
	return cache_Data_List_Lazy_findIndex
}

var cache_Data_List_Lazy_findLastIndex gopurs_runtime.Value
var once_Data_List_Lazy_findLastIndex sync.Once
func Get_Data_List_Lazy_findLastIndex() gopurs_runtime.Value {
	once_Data_List_Lazy_findLastIndex.Do(func() {
		cache_Data_List_Lazy_findLastIndex = gopurs_runtime.Func2(func(fn_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_findLastIndex(fn_0_box, xs_1_box))}
})
	})
	return cache_Data_List_Lazy_findLastIndex
}

var cache_Data_List_Lazy_filterM gopurs_runtime.Value
var once_Data_List_Lazy_filterM sync.Once
func Get_Data_List_Lazy_filterM() gopurs_runtime.Value {
	once_Data_List_Lazy_filterM.Do(func() {
		cache_Data_List_Lazy_filterM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_filterM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Data_List_Lazy_filterM
}

var cache_Data_List_Lazy_filter gopurs_runtime.Value
var once_Data_List_Lazy_filter sync.Once
func Get_Data_List_Lazy_filter() gopurs_runtime.Value {
	once_Data_List_Lazy_filter.Do(func() {
		cache_Data_List_Lazy_filter = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_filter(p_0_box)
})
	})
	return cache_Data_List_Lazy_filter
}

var cache_Data_List_Lazy_intersectBy gopurs_runtime.Value
var once_Data_List_Lazy_intersectBy sync.Once
func Get_Data_List_Lazy_intersectBy() gopurs_runtime.Value {
	once_Data_List_Lazy_intersectBy.Do(func() {
		cache_Data_List_Lazy_intersectBy = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_intersectBy(eq_0_box, xs_1_box, ys_2_box)
})
	})
	return cache_Data_List_Lazy_intersectBy
}

var cache_Data_List_Lazy_intersect gopurs_runtime.Value
var once_Data_List_Lazy_intersect sync.Once
func Get_Data_List_Lazy_intersect() gopurs_runtime.Value {
	once_Data_List_Lazy_intersect.Do(func() {
		cache_Data_List_Lazy_intersect = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_intersect(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_List_Lazy_intersect
}

var cache_Data_List_Lazy_nubByEq gopurs_runtime.Value
var once_Data_List_Lazy_nubByEq sync.Once
func Get_Data_List_Lazy_nubByEq() gopurs_runtime.Value {
	once_Data_List_Lazy_nubByEq.Do(func() {
		cache_Data_List_Lazy_nubByEq = gopurs_runtime.Func2(func(eq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_nubByEq(eq_0_box, x_1_box)
})
	})
	return cache_Data_List_Lazy_nubByEq
}

var cache_Data_List_Lazy_nubEq gopurs_runtime.Value
var once_Data_List_Lazy_nubEq sync.Once
func Get_Data_List_Lazy_nubEq() gopurs_runtime.Value {
	once_Data_List_Lazy_nubEq.Do(func() {
		cache_Data_List_Lazy_nubEq = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_nubEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_List_Lazy_nubEq
}

var cache_Data_List_Lazy_eqPattern gopurs_runtime.Value
var once_Data_List_Lazy_eqPattern sync.Once
func Get_Data_List_Lazy_eqPattern() gopurs_runtime.Value {
	once_Data_List_Lazy_eqPattern.Do(func() {
		cache_Data_List_Lazy_eqPattern = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_eqPattern(dictEq_0_box)
})
	})
	return cache_Data_List_Lazy_eqPattern
}

var cache_Data_List_Lazy_ordPattern gopurs_runtime.Value
var once_Data_List_Lazy_ordPattern sync.Once
func Get_Data_List_Lazy_ordPattern() gopurs_runtime.Value {
	once_Data_List_Lazy_ordPattern.Do(func() {
		cache_Data_List_Lazy_ordPattern = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_ordPattern(dictOrd_0_box)
})
	})
	return cache_Data_List_Lazy_ordPattern
}

var cache_Data_List_Lazy_elemLastIndex gopurs_runtime.Value
var once_Data_List_Lazy_elemLastIndex sync.Once
func Get_Data_List_Lazy_elemLastIndex() gopurs_runtime.Value {
	once_Data_List_Lazy_elemLastIndex.Do(func() {
		cache_Data_List_Lazy_elemLastIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_elemLastIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), x_1_box)
})
	})
	return cache_Data_List_Lazy_elemLastIndex
}

var cache_Data_List_Lazy_elemIndex gopurs_runtime.Value
var once_Data_List_Lazy_elemIndex sync.Once
func Get_Data_List_Lazy_elemIndex() gopurs_runtime.Value {
	once_Data_List_Lazy_elemIndex.Do(func() {
		cache_Data_List_Lazy_elemIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_elemIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), x_1_box)
})
	})
	return cache_Data_List_Lazy_elemIndex
}

var cache_Data_List_Lazy_dropWhile gopurs_runtime.Value
var once_Data_List_Lazy_dropWhile sync.Once
func Get_Data_List_Lazy_dropWhile() gopurs_runtime.Value {
	once_Data_List_Lazy_dropWhile.Do(func() {
		cache_Data_List_Lazy_dropWhile = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_dropWhile(p_0_box)
})
	})
	return cache_Data_List_Lazy_dropWhile
}

var cache_Data_List_Lazy_drop gopurs_runtime.Value
var once_Data_List_Lazy_drop sync.Once
func Get_Data_List_Lazy_drop() gopurs_runtime.Value {
	once_Data_List_Lazy_drop.Do(func() {
		cache_Data_List_Lazy_drop = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_drop(n_0_box.IntVal)
})
	})
	return cache_Data_List_Lazy_drop
}

var cache_Data_List_Lazy_slice gopurs_runtime.Value
var once_Data_List_Lazy_slice sync.Once
func Get_Data_List_Lazy_slice() gopurs_runtime.Value {
	once_Data_List_Lazy_slice.Do(func() {
		cache_Data_List_Lazy_slice = gopurs_runtime.Func3(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_slice(start_0_box.IntVal, end_1_box.IntVal, xs_2_box)
})
	})
	return cache_Data_List_Lazy_slice
}

var cache_Data_List_Lazy_deleteBy gopurs_runtime.Value
var once_Data_List_Lazy_deleteBy sync.Once
func Get_Data_List_Lazy_deleteBy() gopurs_runtime.Value {
	once_Data_List_Lazy_deleteBy.Do(func() {
		cache_Data_List_Lazy_deleteBy = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_deleteBy(eq_0_box, x_1_box, xs_2_box)
})
	})
	return cache_Data_List_Lazy_deleteBy
}

var cache_Data_List_Lazy_unionBy gopurs_runtime.Value
var once_Data_List_Lazy_unionBy sync.Once
func Get_Data_List_Lazy_unionBy() gopurs_runtime.Value {
	once_Data_List_Lazy_unionBy.Do(func() {
		cache_Data_List_Lazy_unionBy = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_unionBy(eq_0_box, xs_1_box, ys_2_box)
})
	})
	return cache_Data_List_Lazy_unionBy
}

var cache_Data_List_Lazy_union gopurs_runtime.Value
var once_Data_List_Lazy_union sync.Once
func Get_Data_List_Lazy_union() gopurs_runtime.Value {
	once_Data_List_Lazy_union.Do(func() {
		cache_Data_List_Lazy_union = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_union(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_List_Lazy_union
}

var cache_Data_List_Lazy_deleteAt gopurs_runtime.Value
var once_Data_List_Lazy_deleteAt sync.Once
func Get_Data_List_Lazy_deleteAt() gopurs_runtime.Value {
	once_Data_List_Lazy_deleteAt.Do(func() {
		cache_Data_List_Lazy_deleteAt = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_deleteAt(n_0_box.IntVal, xs_1_box)
})
	})
	return cache_Data_List_Lazy_deleteAt
}

var cache_Data_List_Lazy_delete gopurs_runtime.Value
var once_Data_List_Lazy_delete sync.Once
func Get_Data_List_Lazy_delete() gopurs_runtime.Value {
	once_Data_List_Lazy_delete.Do(func() {
		cache_Data_List_Lazy_delete = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_delete(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_List_Lazy_delete
}

var cache_Data_List_Lazy_difference gopurs_runtime.Value
var once_Data_List_Lazy_difference sync.Once
func Get_Data_List_Lazy_difference() gopurs_runtime.Value {
	once_Data_List_Lazy_difference.Do(func() {
		cache_Data_List_Lazy_difference = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_difference(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_List_Lazy_difference
}

var cache_Data_List_Lazy_cycle gopurs_runtime.Value
var once_Data_List_Lazy_cycle sync.Once
func Get_Data_List_Lazy_cycle() gopurs_runtime.Value {
	once_Data_List_Lazy_cycle.Do(func() {
		cache_Data_List_Lazy_cycle = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_cycle(xs_0_box)
})
	})
	return cache_Data_List_Lazy_cycle
}

var cache_Data_List_Lazy_concatMap gopurs_runtime.Value
var once_Data_List_Lazy_concatMap sync.Once
func Get_Data_List_Lazy_concatMap() gopurs_runtime.Value {
	once_Data_List_Lazy_concatMap.Do(func() {
		cache_Data_List_Lazy_concatMap = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_concatMap(b_0_box, a_1_box)
})
	})
	return cache_Data_List_Lazy_concatMap
}

var cache_Data_List_Lazy_concat gopurs_runtime.Value
var once_Data_List_Lazy_concat sync.Once
func Get_Data_List_Lazy_concat() gopurs_runtime.Value {
	once_Data_List_Lazy_concat.Do(func() {
		cache_Data_List_Lazy_concat = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_concat(v_0_box)
})
	})
	return cache_Data_List_Lazy_concat
}

var cache_Data_List_Lazy_catMaybes gopurs_runtime.Value
var once_Data_List_Lazy_catMaybes sync.Once
func Get_Data_List_Lazy_catMaybes() gopurs_runtime.Value {
	once_Data_List_Lazy_catMaybes.Do(func() {
		cache_Data_List_Lazy_catMaybes = Call_Data_List_Lazy_mapMaybe(gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_Data_List_Lazy_catMaybes
}

var cache_Data_List_Lazy_alterAt gopurs_runtime.Value
var once_Data_List_Lazy_alterAt sync.Once
func Get_Data_List_Lazy_alterAt() gopurs_runtime.Value {
	once_Data_List_Lazy_alterAt.Do(func() {
		cache_Data_List_Lazy_alterAt = gopurs_runtime.Func3(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_alterAt(n_0_box.IntVal, f_1_box, xs_2_box)
})
	})
	return cache_Data_List_Lazy_alterAt
}

var cache_Data_List_Lazy_modifyAt gopurs_runtime.Value
var once_Data_List_Lazy_modifyAt sync.Once
func Get_Data_List_Lazy_modifyAt() gopurs_runtime.Value {
	once_Data_List_Lazy_modifyAt.Do(func() {
		cache_Data_List_Lazy_modifyAt = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_modifyAt(n_0_box.IntVal, f_1_box)
})
	})
	return cache_Data_List_Lazy_modifyAt
}

var cache_Data_List_Lazy_alterAt__950766476 gopurs_runtime.Value
var once_Data_List_Lazy_alterAt__950766476 sync.Once
func Get_Data_List_Lazy_alterAt__950766476() gopurs_runtime.Value {
	once_Data_List_Lazy_alterAt__950766476.Do(func() {
		cache_Data_List_Lazy_alterAt__950766476 = gopurs_runtime.Func3(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_alterAt__950766476(n_0_box.IntVal, f_1_box, xs_2_box)
})
	})
	return cache_Data_List_Lazy_alterAt__950766476
}

var cache_Data_List_Lazy_deleteAt__4024047148 gopurs_runtime.Value
var once_Data_List_Lazy_deleteAt__4024047148 sync.Once
func Get_Data_List_Lazy_deleteAt__4024047148() gopurs_runtime.Value {
	once_Data_List_Lazy_deleteAt__4024047148.Do(func() {
		cache_Data_List_Lazy_deleteAt__4024047148 = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_deleteAt__4024047148(n_0_box.IntVal, xs_1_box)
})
	})
	return cache_Data_List_Lazy_deleteAt__4024047148
}

var cache_Data_List_Lazy_deleteBy__501100275 gopurs_runtime.Value
var once_Data_List_Lazy_deleteBy__501100275 sync.Once
func Get_Data_List_Lazy_deleteBy__501100275() gopurs_runtime.Value {
	once_Data_List_Lazy_deleteBy__501100275.Do(func() {
		cache_Data_List_Lazy_deleteBy__501100275 = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_deleteBy__501100275(eq_0_box, x_1_box, xs_2_box)
})
	})
	return cache_Data_List_Lazy_deleteBy__501100275
}

var cache_Data_List_Lazy_drop__4024047148 gopurs_runtime.Value
var once_Data_List_Lazy_drop__4024047148 sync.Once
func Get_Data_List_Lazy_drop__4024047148() gopurs_runtime.Value {
	once_Data_List_Lazy_drop__4024047148.Do(func() {
		cache_Data_List_Lazy_drop__4024047148 = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_drop__4024047148(n_0_box.IntVal)
})
	})
	return cache_Data_List_Lazy_drop__4024047148
}

var cache_Data_List_Lazy_filter__638755635 gopurs_runtime.Value
var once_Data_List_Lazy_filter__638755635 sync.Once
func Get_Data_List_Lazy_filter__638755635() gopurs_runtime.Value {
	once_Data_List_Lazy_filter__638755635.Do(func() {
		cache_Data_List_Lazy_filter__638755635 = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_filter__638755635(p_0_box)
})
	})
	return cache_Data_List_Lazy_filter__638755635
}

var cache_Data_List_Lazy_filterM__647926151 gopurs_runtime.Value
var once_Data_List_Lazy_filterM__647926151 sync.Once
func Get_Data_List_Lazy_filterM__647926151() gopurs_runtime.Value {
	once_Data_List_Lazy_filterM__647926151.Do(func() {
		cache_Data_List_Lazy_filterM__647926151 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_filterM__647926151(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Data_List_Lazy_filterM__647926151
}

var cache_Data_List_Lazy_findIndex__1594900290 gopurs_runtime.Value
var once_Data_List_Lazy_findIndex__1594900290 sync.Once
func Get_Data_List_Lazy_findIndex__1594900290() gopurs_runtime.Value {
	once_Data_List_Lazy_findIndex__1594900290.Do(func() {
		cache_Data_List_Lazy_findIndex__1594900290 = gopurs_runtime.Func(func(fn_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_findIndex__1594900290(fn_0_box)
})
	})
	return cache_Data_List_Lazy_findIndex__1594900290
}

var cache_Data_List_Lazy_findLastIndex__1594900290 gopurs_runtime.Value
var once_Data_List_Lazy_findLastIndex__1594900290 sync.Once
func Get_Data_List_Lazy_findLastIndex__1594900290() gopurs_runtime.Value {
	once_Data_List_Lazy_findLastIndex__1594900290.Do(func() {
		cache_Data_List_Lazy_findLastIndex__1594900290 = gopurs_runtime.Func2(func(fn_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_findLastIndex__1594900290(fn_0_box, xs_1_box))}
})
	})
	return cache_Data_List_Lazy_findLastIndex__1594900290
}

var cache_Data_List_Lazy_foldM__3505933597 gopurs_runtime.Value
var once_Data_List_Lazy_foldM__3505933597 sync.Once
func Get_Data_List_Lazy_foldM__3505933597() gopurs_runtime.Value {
	once_Data_List_Lazy_foldM__3505933597.Do(func() {
		cache_Data_List_Lazy_foldM__3505933597 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_foldM__3505933597(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Data_List_Lazy_foldM__3505933597
}

var cache_Data_List_Lazy_fromFoldable__4212258679 gopurs_runtime.Value
var once_Data_List_Lazy_fromFoldable__4212258679 sync.Once
func Get_Data_List_Lazy_fromFoldable__4212258679() gopurs_runtime.Value {
	once_Data_List_Lazy_fromFoldable__4212258679.Do(func() {
		cache_Data_List_Lazy_fromFoldable__4212258679 = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_fromFoldable__4212258679(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box))
})
	})
	return cache_Data_List_Lazy_fromFoldable__4212258679
}

var cache_Data_List_Lazy_fromStep__1398792641 gopurs_runtime.Value
var once_Data_List_Lazy_fromStep__1398792641 sync.Once
func Get_Data_List_Lazy_fromStep__1398792641() gopurs_runtime.Value {
	once_Data_List_Lazy_fromStep__1398792641.Do(func() {
		cache_Data_List_Lazy_fromStep__1398792641 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_fromStep__1398792641(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](x_0_box))
})
	})
	return cache_Data_List_Lazy_fromStep__1398792641
}

var cache_Data_List_Lazy_groupBy__1659362014 gopurs_runtime.Value
var once_Data_List_Lazy_groupBy__1659362014 sync.Once
func Get_Data_List_Lazy_groupBy__1659362014() gopurs_runtime.Value {
	once_Data_List_Lazy_groupBy__1659362014.Do(func() {
		cache_Data_List_Lazy_groupBy__1659362014 = gopurs_runtime.Func2(func(eq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_groupBy__1659362014(eq_0_box, x_1_box)
})
	})
	return cache_Data_List_Lazy_groupBy__1659362014
}

var cache_Data_List_Lazy_head__2155426095 gopurs_runtime.Value
var once_Data_List_Lazy_head__2155426095 sync.Once
func Get_Data_List_Lazy_head__2155426095() gopurs_runtime.Value {
	once_Data_List_Lazy_head__2155426095.Do(func() {
		cache_Data_List_Lazy_head__2155426095 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_head__2155426095(xs_0_box))}
})
	})
	return cache_Data_List_Lazy_head__2155426095
}

var cache_Data_List_Lazy_init__109194273 gopurs_runtime.Value
var once_Data_List_Lazy_init__109194273 sync.Once
func Get_Data_List_Lazy_init__109194273() gopurs_runtime.Value {
	once_Data_List_Lazy_init__109194273.Do(func() {
		cache_Data_List_Lazy_init__109194273 = func() gopurs_runtime.Value {
var go__go_0_0_36 gopurs_runtime.Value
_ = go__go_0_0_36
go__go_0_0_36 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 *Constructor_Data_Maybe_Just
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil) {
var __t7 *Constructor_Data_Maybe_Just
{
// TAST (Let): __local_var_2_5 -> gopurs_runtime.Value
var __local_var_2_5 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_uncons((*Constructor_Data_List_Lazy_Types_Cons)(v_1.UnsafePtr).V1))}
var __t6 gopurs_runtime.Value
{
if (__local_var_2_5.Type == 9 && __local_var_2_5.IntVal == 930809136 && __local_var_2_5.UnsafePtr == nil) {
__t6 = gopurs_runtime.Bool(true)
goto end_branch_6
} else {

}
}
{
if (__local_var_2_5.Type == 9 && __local_var_2_5.IntVal == 930809136 && __local_var_2_5.UnsafePtr != nil) {
__t6 = gopurs_runtime.Bool(false)
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
if (__t6.IntVal) != (0) {
__t7 = &Constructor_Data_Maybe_Just{1, Get_Data_List_Lazy_Types_nil()}
goto end_branch_7
} else {

}
}
{
// TAST (Let): __local_var_2_1 -> *Constructor_Data_Maybe_Just
__local_var_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(go__go_0_0_36, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_1.UnsafePtr).V1)))}))
_ = __local_var_2_1
var __t4 *Constructor_Data_Maybe_Just
{
if (__local_var_2_1 != nil) {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := (*Constructor_Data_List_Lazy_Types_Cons)(v_1.UnsafePtr).V0
_ = __local_var_3_2
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := (__local_var_2_1).V0
_ = __local_var_4_3
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_3_2, __local_var_4_3})}
}))}
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_4:
__t7 = __t4
}
end_branch_7:
__t8 = __t7
goto end_branch_8
} else {

}
}
{
__t8 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t8)}
})
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_0_0_36, gopurs_runtime.Apply(Get_Data_Lazy_force(), x_1))
})
}()
	})
	return cache_Data_List_Lazy_init__109194273
}

var cache_Data_List_Lazy_insertAt__725610501 gopurs_runtime.Value
var once_Data_List_Lazy_insertAt__725610501 sync.Once
func Get_Data_List_Lazy_insertAt__725610501() gopurs_runtime.Value {
	once_Data_List_Lazy_insertAt__725610501.Do(func() {
		cache_Data_List_Lazy_insertAt__725610501 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_insertAt__725610501(v_0_box.IntVal, v1_1_box, v2_2_box)
})
	})
	return cache_Data_List_Lazy_insertAt__725610501
}

var cache_Data_List_Lazy_insertBy__2098566601 gopurs_runtime.Value
var once_Data_List_Lazy_insertBy__2098566601 sync.Once
func Get_Data_List_Lazy_insertBy__2098566601() gopurs_runtime.Value {
	once_Data_List_Lazy_insertBy__2098566601.Do(func() {
		cache_Data_List_Lazy_insertBy__2098566601 = gopurs_runtime.Func3(func(cmp_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_insertBy__2098566601(cmp_0_box, x_1_box, xs_2_box)
})
	})
	return cache_Data_List_Lazy_insertBy__2098566601
}

var cache_Data_List_Lazy_intersectBy__3844889126 gopurs_runtime.Value
var once_Data_List_Lazy_intersectBy__3844889126 sync.Once
func Get_Data_List_Lazy_intersectBy__3844889126() gopurs_runtime.Value {
	once_Data_List_Lazy_intersectBy__3844889126.Do(func() {
		cache_Data_List_Lazy_intersectBy__3844889126 = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_intersectBy__3844889126(eq_0_box, xs_1_box, ys_2_box)
})
	})
	return cache_Data_List_Lazy_intersectBy__3844889126
}

var cache_Data_List_Lazy_iterate__455058292 gopurs_runtime.Value
var once_Data_List_Lazy_iterate__455058292 sync.Once
func Get_Data_List_Lazy_iterate__455058292() gopurs_runtime.Value {
	once_Data_List_Lazy_iterate__455058292.Do(func() {
		cache_Data_List_Lazy_iterate__455058292 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_iterate__455058292(f_0_box, x_1_box)
})
	})
	return cache_Data_List_Lazy_iterate__455058292
}

var cache_Data_List_Lazy_last__1102843348 gopurs_runtime.Value
var once_Data_List_Lazy_last__1102843348 sync.Once
func Get_Data_List_Lazy_last__1102843348() gopurs_runtime.Value {
	once_Data_List_Lazy_last__1102843348.Do(func() {
		cache_Data_List_Lazy_last__1102843348 = func() gopurs_runtime.Value {
var go__go_0_0_39 gopurs_runtime.Value
go__go_0_0_39 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop gopurs_runtime.Value = v_1_loop_val
go__go_0_0_39:
for {
if false { continue go__go_0_0_39 }
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var __t4 *Constructor_Data_Maybe_Just
{
if (v_1.Type == 9 && v_1.IntVal == 218341868 && v_1.UnsafePtr != nil) {
var __t3 *Constructor_Data_Maybe_Just
{
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
var __local_var_2_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_uncons((*Constructor_Data_List_Lazy_Types_Cons)(v_1.UnsafePtr).V1))}
var __t2 gopurs_runtime.Value
{
if (__local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 930809136 && __local_var_2_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Bool(true)
goto end_branch_2
} else {

}
}
{
if (__local_var_2_1.Type == 9 && __local_var_2_1.IntVal == 930809136 && __local_var_2_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.Bool(false)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
if (__t2.IntVal) != (0) {
__t3 = &Constructor_Data_Maybe_Just{1, (*Constructor_Data_List_Lazy_Types_Cons)(v_1.UnsafePtr).V0}
goto end_branch_3
} else {

}
}
{
v_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_1.UnsafePtr).V1)))}
continue go__go_0_0_39
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
}
}()
})
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_0_0_39, gopurs_runtime.Apply(Get_Data_Lazy_force(), x_1))
})
}()
	})
	return cache_Data_List_Lazy_last__1102843348
}

var cache_Data_List_Lazy_length__162861552 gopurs_runtime.Value
var once_Data_List_Lazy_length__162861552 sync.Once
func Get_Data_List_Lazy_length__162861552() gopurs_runtime.Value {
	once_Data_List_Lazy_length__162861552.Do(func() {
		cache_Data_List_Lazy_length__162861552 = func() gopurs_runtime.Value {
var go__go_0_0_40 gopurs_runtime.Value
go__go_0_0_40 = gopurs_runtime.Func(func(b_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_1_loop gopurs_runtime.Value = b_1_loop_val
var xs_2_loop gopurs_runtime.Value = xs_2_loop_val
go__go_0_0_40:
for {
if false { continue go__go_0_0_40 }
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
// TAST (Let): v_3_1 -> *Constructor_Data_List_Lazy_Types_Cons
v_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2))
_ = v_3_1
var __t2 gopurs_runtime.Value
{
if (v_3_1 == nil) {
__t2 = b_1
goto end_branch_2
} else {

}
}
{
if (v_3_1 != nil) {
b_1_loop = gopurs_runtime.Int((b_1.IntVal) + (1))
xs_2_loop = (v_3_1).V1
continue go__go_0_0_40
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
return gopurs_runtime.Apply(go__go_0_0_40, gopurs_runtime.Int(0))
}()
	})
	return cache_Data_List_Lazy_length__162861552
}

var cache_Data_List_Lazy_many__956417025 gopurs_runtime.Value
var once_Data_List_Lazy_many__956417025 sync.Once
func Get_Data_List_Lazy_many__956417025() gopurs_runtime.Value {
	once_Data_List_Lazy_many__956417025.Do(func() {
		cache_Data_List_Lazy_many__956417025 = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_many__956417025(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative](dictAlternative_0_box))
})
	})
	return cache_Data_List_Lazy_many__956417025
}

var cache_Data_List_Lazy_mapMaybe__3574309085 gopurs_runtime.Value
var once_Data_List_Lazy_mapMaybe__3574309085 sync.Once
func Get_Data_List_Lazy_mapMaybe__3574309085() gopurs_runtime.Value {
	once_Data_List_Lazy_mapMaybe__3574309085.Do(func() {
		cache_Data_List_Lazy_mapMaybe__3574309085 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_mapMaybe__3574309085(f_0_box)
})
	})
	return cache_Data_List_Lazy_mapMaybe__3574309085
}

var cache_Data_List_Lazy_mapMaybe__2519317725 gopurs_runtime.Value
var once_Data_List_Lazy_mapMaybe__2519317725 sync.Once
func Get_Data_List_Lazy_mapMaybe__2519317725() gopurs_runtime.Value {
	once_Data_List_Lazy_mapMaybe__2519317725.Do(func() {
		cache_Data_List_Lazy_mapMaybe__2519317725 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_mapMaybe__2519317725(f_0_box)
})
	})
	return cache_Data_List_Lazy_mapMaybe__2519317725
}

var cache_Data_List_Lazy_mapMaybe__1687744733 gopurs_runtime.Value
var once_Data_List_Lazy_mapMaybe__1687744733 sync.Once
func Get_Data_List_Lazy_mapMaybe__1687744733() gopurs_runtime.Value {
	once_Data_List_Lazy_mapMaybe__1687744733.Do(func() {
		cache_Data_List_Lazy_mapMaybe__1687744733 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_mapMaybe__1687744733(f_0_box)
})
	})
	return cache_Data_List_Lazy_mapMaybe__1687744733
}

var cache_Data_List_Lazy_mapMaybe__899591645 gopurs_runtime.Value
var once_Data_List_Lazy_mapMaybe__899591645 sync.Once
func Get_Data_List_Lazy_mapMaybe__899591645() gopurs_runtime.Value {
	once_Data_List_Lazy_mapMaybe__899591645.Do(func() {
		cache_Data_List_Lazy_mapMaybe__899591645 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_mapMaybe__899591645(f_0_box)
})
	})
	return cache_Data_List_Lazy_mapMaybe__899591645
}

var cache_Data_List_Lazy_mapMaybe__600226685 gopurs_runtime.Value
var once_Data_List_Lazy_mapMaybe__600226685 sync.Once
func Get_Data_List_Lazy_mapMaybe__600226685() gopurs_runtime.Value {
	once_Data_List_Lazy_mapMaybe__600226685.Do(func() {
		cache_Data_List_Lazy_mapMaybe__600226685 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_mapMaybe__600226685(f_0_box)
})
	})
	return cache_Data_List_Lazy_mapMaybe__600226685
}

var cache_Data_List_Lazy_nubBy__2220739616 gopurs_runtime.Value
var once_Data_List_Lazy_nubBy__2220739616 sync.Once
func Get_Data_List_Lazy_nubBy__2220739616() gopurs_runtime.Value {
	once_Data_List_Lazy_nubBy__2220739616.Do(func() {
		cache_Data_List_Lazy_nubBy__2220739616 = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_nubBy__2220739616(p_0_box)
})
	})
	return cache_Data_List_Lazy_nubBy__2220739616
}

var cache_Data_List_Lazy_nubByEq__616397370 gopurs_runtime.Value
var once_Data_List_Lazy_nubByEq__616397370 sync.Once
func Get_Data_List_Lazy_nubByEq__616397370() gopurs_runtime.Value {
	once_Data_List_Lazy_nubByEq__616397370.Do(func() {
		cache_Data_List_Lazy_nubByEq__616397370 = gopurs_runtime.Func2(func(eq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_nubByEq__616397370(eq_0_box, x_1_box)
})
	})
	return cache_Data_List_Lazy_nubByEq__616397370
}

var cache_Data_List_Lazy_null__1674339719 gopurs_runtime.Value
var once_Data_List_Lazy_null__1674339719 sync.Once
func Get_Data_List_Lazy_null__1674339719() gopurs_runtime.Value {
	once_Data_List_Lazy_null__1674339719.Do(func() {
		cache_Data_List_Lazy_null__1674339719 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_List_Lazy_null__1674339719(x_0_box))
})
	})
	return cache_Data_List_Lazy_null__1674339719
}

var cache_Data_List_Lazy_repeat__2462085934 gopurs_runtime.Value
var once_Data_List_Lazy_repeat__2462085934 sync.Once
func Get_Data_List_Lazy_repeat__2462085934() gopurs_runtime.Value {
	once_Data_List_Lazy_repeat__2462085934.Do(func() {
		cache_Data_List_Lazy_repeat__2462085934 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_repeat__2462085934(x_0_box)
})
	})
	return cache_Data_List_Lazy_repeat__2462085934
}

var cache_Data_List_Lazy_repeat__2149902581 gopurs_runtime.Value
var once_Data_List_Lazy_repeat__2149902581 sync.Once
func Get_Data_List_Lazy_repeat__2149902581() gopurs_runtime.Value {
	once_Data_List_Lazy_repeat__2149902581.Do(func() {
		cache_Data_List_Lazy_repeat__2149902581 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_repeat__2149902581(x_0_box)
})
	})
	return cache_Data_List_Lazy_repeat__2149902581
}

var cache_Data_List_Lazy_replicateM__3816548429 gopurs_runtime.Value
var once_Data_List_Lazy_replicateM__3816548429 sync.Once
func Get_Data_List_Lazy_replicateM__3816548429() gopurs_runtime.Value {
	once_Data_List_Lazy_replicateM__3816548429.Do(func() {
		cache_Data_List_Lazy_replicateM__3816548429 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_replicateM__3816548429(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Data_List_Lazy_replicateM__3816548429
}

var cache_Data_List_Lazy_reverse__1315655552 gopurs_runtime.Value
var once_Data_List_Lazy_reverse__1315655552 sync.Once
func Get_Data_List_Lazy_reverse__1315655552() gopurs_runtime.Value {
	once_Data_List_Lazy_reverse__1315655552.Do(func() {
		cache_Data_List_Lazy_reverse__1315655552 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_reverse__1315655552(xs_0_box)
})
	})
	return cache_Data_List_Lazy_reverse__1315655552
}

var cache_Data_List_Lazy_scanlLazy__3747838620 gopurs_runtime.Value
var once_Data_List_Lazy_scanlLazy__3747838620 sync.Once
func Get_Data_List_Lazy_scanlLazy__3747838620() gopurs_runtime.Value {
	once_Data_List_Lazy_scanlLazy__3747838620.Do(func() {
		cache_Data_List_Lazy_scanlLazy__3747838620 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, acc_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_scanlLazy__3747838620(f_0_box, acc_1_box, xs_2_box)
})
	})
	return cache_Data_List_Lazy_scanlLazy__3747838620
}

var cache_Data_List_Lazy_some__956417025 gopurs_runtime.Value
var once_Data_List_Lazy_some__956417025 sync.Once
func Get_Data_List_Lazy_some__956417025() gopurs_runtime.Value {
	once_Data_List_Lazy_some__956417025.Do(func() {
		cache_Data_List_Lazy_some__956417025 = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_some__956417025(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative](dictAlternative_0_box))
})
	})
	return cache_Data_List_Lazy_some__956417025
}

var cache_Data_List_Lazy_span__776304907 gopurs_runtime.Value
var once_Data_List_Lazy_span__776304907 sync.Once
func Get_Data_List_Lazy_span__776304907() gopurs_runtime.Value {
	once_Data_List_Lazy_span__776304907.Do(func() {
		cache_Data_List_Lazy_span__776304907 = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_span__776304907(p_0_box, xs_1_box)
})
	})
	return cache_Data_List_Lazy_span__776304907
}

var cache_Data_List_Lazy_tail__1935051898 gopurs_runtime.Value
var once_Data_List_Lazy_tail__1935051898 sync.Once
func Get_Data_List_Lazy_tail__1935051898() gopurs_runtime.Value {
	once_Data_List_Lazy_tail__1935051898.Do(func() {
		cache_Data_List_Lazy_tail__1935051898 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_tail__1935051898(xs_0_box))}
})
	})
	return cache_Data_List_Lazy_tail__1935051898
}

var cache_Data_List_Lazy_take__4024047148 gopurs_runtime.Value
var once_Data_List_Lazy_take__4024047148 sync.Once
func Get_Data_List_Lazy_take__4024047148() gopurs_runtime.Value {
	once_Data_List_Lazy_take__4024047148.Do(func() {
		cache_Data_List_Lazy_take__4024047148 = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_take__4024047148(n_0_box.IntVal)
})
	})
	return cache_Data_List_Lazy_take__4024047148
}

var cache_Data_List_Lazy_takeWhile__638755635 gopurs_runtime.Value
var once_Data_List_Lazy_takeWhile__638755635 sync.Once
func Get_Data_List_Lazy_takeWhile__638755635() gopurs_runtime.Value {
	once_Data_List_Lazy_takeWhile__638755635.Do(func() {
		cache_Data_List_Lazy_takeWhile__638755635 = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_takeWhile__638755635(p_0_box, x_1_box)
})
	})
	return cache_Data_List_Lazy_takeWhile__638755635
}

var cache_Data_List_Lazy_transpose__1534541312 gopurs_runtime.Value
var once_Data_List_Lazy_transpose__1534541312 sync.Once
func Get_Data_List_Lazy_transpose__1534541312() gopurs_runtime.Value {
	once_Data_List_Lazy_transpose__1534541312.Do(func() {
		cache_Data_List_Lazy_transpose__1534541312 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_transpose__1534541312(xs_0_box)
})
	})
	return cache_Data_List_Lazy_transpose__1534541312
}

var cache_Data_List_Lazy_uncons__3647012005 gopurs_runtime.Value
var once_Data_List_Lazy_uncons__3647012005 sync.Once
func Get_Data_List_Lazy_uncons__3647012005() gopurs_runtime.Value {
	once_Data_List_Lazy_uncons__3647012005.Do(func() {
		cache_Data_List_Lazy_uncons__3647012005 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_uncons__3647012005(xs_0_box))}
})
	})
	return cache_Data_List_Lazy_uncons__3647012005
}

var cache_Data_List_Lazy_uncons__1321764894 gopurs_runtime.Value
var once_Data_List_Lazy_uncons__1321764894 sync.Once
func Get_Data_List_Lazy_uncons__1321764894() gopurs_runtime.Value {
	once_Data_List_Lazy_uncons__1321764894.Do(func() {
		cache_Data_List_Lazy_uncons__1321764894 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_uncons__1321764894(xs_0_box))}
})
	})
	return cache_Data_List_Lazy_uncons__1321764894
}

var cache_Data_List_Lazy_uncons__974566859 gopurs_runtime.Value
var once_Data_List_Lazy_uncons__974566859 sync.Once
func Get_Data_List_Lazy_uncons__974566859() gopurs_runtime.Value {
	once_Data_List_Lazy_uncons__974566859.Do(func() {
		cache_Data_List_Lazy_uncons__974566859 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_uncons__974566859(xs_0_box))}
})
	})
	return cache_Data_List_Lazy_uncons__974566859
}

var cache_Data_List_Lazy_uncons__1420258522 gopurs_runtime.Value
var once_Data_List_Lazy_uncons__1420258522 sync.Once
func Get_Data_List_Lazy_uncons__1420258522() gopurs_runtime.Value {
	once_Data_List_Lazy_uncons__1420258522.Do(func() {
		cache_Data_List_Lazy_uncons__1420258522 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_uncons__1420258522(xs_0_box))}
})
	})
	return cache_Data_List_Lazy_uncons__1420258522
}

var cache_Data_List_Lazy_unionBy__3844889126 gopurs_runtime.Value
var once_Data_List_Lazy_unionBy__3844889126 sync.Once
func Get_Data_List_Lazy_unionBy__3844889126() gopurs_runtime.Value {
	once_Data_List_Lazy_unionBy__3844889126.Do(func() {
		cache_Data_List_Lazy_unionBy__3844889126 = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_unionBy__3844889126(eq_0_box, xs_1_box, ys_2_box)
})
	})
	return cache_Data_List_Lazy_unionBy__3844889126
}

var cache_Data_List_Lazy_updateAt__725610501 gopurs_runtime.Value
var once_Data_List_Lazy_updateAt__725610501 sync.Once
func Get_Data_List_Lazy_updateAt__725610501() gopurs_runtime.Value {
	once_Data_List_Lazy_updateAt__725610501.Do(func() {
		cache_Data_List_Lazy_updateAt__725610501 = gopurs_runtime.Func3(func(n_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_updateAt__725610501(n_0_box.IntVal, x_1_box, xs_2_box)
})
	})
	return cache_Data_List_Lazy_updateAt__725610501
}

var cache_Data_List_Lazy_zipWith__3539178005 gopurs_runtime.Value
var once_Data_List_Lazy_zipWith__3539178005 sync.Once
func Get_Data_List_Lazy_zipWith__3539178005() gopurs_runtime.Value {
	once_Data_List_Lazy_zipWith__3539178005.Do(func() {
		cache_Data_List_Lazy_zipWith__3539178005 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_zipWith__3539178005(f_0_box, xs_1_box, ys_2_box)
})
	})
	return cache_Data_List_Lazy_zipWith__3539178005
}

var cache_Data_List_Lazy_zipWith__3210333397 gopurs_runtime.Value
var once_Data_List_Lazy_zipWith__3210333397 sync.Once
func Get_Data_List_Lazy_zipWith__3210333397() gopurs_runtime.Value {
	once_Data_List_Lazy_zipWith__3210333397.Do(func() {
		cache_Data_List_Lazy_zipWith__3210333397 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_zipWith__3210333397(f_0_box, xs_1_box, ys_2_box)
})
	})
	return cache_Data_List_Lazy_zipWith__3210333397
}

var cache_Data_List_Lazy_zipWith__3984071349 gopurs_runtime.Value
var once_Data_List_Lazy_zipWith__3984071349 sync.Once
func Get_Data_List_Lazy_zipWith__3984071349() gopurs_runtime.Value {
	once_Data_List_Lazy_zipWith__3984071349.Do(func() {
		cache_Data_List_Lazy_zipWith__3984071349 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_zipWith__3984071349(f_0_box, xs_1_box, ys_2_box)
})
	})
	return cache_Data_List_Lazy_zipWith__3984071349
}

var cache_Data_List_Lazy_zipWith__2064722709 gopurs_runtime.Value
var once_Data_List_Lazy_zipWith__2064722709 sync.Once
func Get_Data_List_Lazy_zipWith__2064722709() gopurs_runtime.Value {
	once_Data_List_Lazy_zipWith__2064722709.Do(func() {
		cache_Data_List_Lazy_zipWith__2064722709 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_zipWith__2064722709(f_0_box, xs_1_box, ys_2_box)
})
	})
	return cache_Data_List_Lazy_zipWith__2064722709
}

func Call_Data_List_Lazy_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_List_Lazy_Pattern(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_List_Lazy_zipWith(f_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
zipWith:
for {
if false { continue zipWith }
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)
_ = __local_var_4_1
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr == nil) {
__t2 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_2
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr == nil) {
__t2 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_2
} else {

}
}
{
if ((__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr != nil)) {
__t2 = &Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Apply2(f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_1.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons)(v1_5.UnsafePtr).V0), Call_Data_List_Lazy_zipWith(f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_1.UnsafePtr).V1, (*Constructor_Data_List_Lazy_Types_Cons)(v1_5.UnsafePtr).V1)}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
})
}))
_ = __local_var_3_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Lazy_force(), __local_var_3_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_2))
}))
}
}

func Call_Data_List_Lazy_zipWithA(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative, f_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value, ys_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
var ys_3 gopurs_runtime.Value = ys_3_loop
_ = ys_3
// TAST (Let): Apply0_4_0 -> *Constructor_Control_Apply_Apply
Apply0_4_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}))
_ = Apply0_4_0
// TAST (Let): Functor0_5_1 -> *Constructor_Data_Functor_Functor
Functor0_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_1
var go__go_6_2_0 gopurs_runtime.Value
go__go_6_2_0 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_7_loop gopurs_runtime.Value = b_7_loop_val
var xs_8_loop gopurs_runtime.Value = xs_8_loop_val
go__go_6_2_0:
for {
if false { continue go__go_6_2_0 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var xs_8 gopurs_runtime.Value = xs_8_loop
_ = xs_8
// TAST (Let): v_9_3 -> *Constructor_Data_List_Lazy_Types_Cons
v_9_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_8))
_ = v_9_3
var __t4 gopurs_runtime.Value
{
if (v_9_3 == nil) {
__t4 = b_7
goto end_branch_4
} else {

}
}
{
if (v_9_3 != nil) {
b_7_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_4_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_1.V0), Get_Data_List_Lazy_Types_cons(), (v_9_3).V0), b_7)
xs_8_loop = (v_9_3).V1
continue go__go_6_2_0
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
var go__go_7_5_1 gopurs_runtime.Value
go__go_7_5_1 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_8_loop gopurs_runtime.Value = b_8_loop_val
var xs_9_loop gopurs_runtime.Value = xs_9_loop_val
go__go_7_5_1:
for {
if false { continue go__go_7_5_1 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var xs_9 gopurs_runtime.Value = xs_9_loop
_ = xs_9
// TAST (Let): v_10_6 -> *Constructor_Data_List_Lazy_Types_Cons
v_10_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_9))
_ = v_10_6
var __t8 gopurs_runtime.Value
{
if (v_10_6 == nil) {
__t8 = b_8
goto end_branch_8
} else {

}
}
{
if (v_10_6 != nil) {
// TAST (Let): __local_var_11_7 -> gopurs_runtime.Value
__local_var_11_7 := (v_10_6).V0
_ = __local_var_11_7
b_8_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_11_7, b_8})}
}))
xs_9_loop = (v_10_6).V1
continue go__go_7_5_1
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
return gopurs_runtime.Apply2(go__go_6_2_0, gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), Get_Data_List_Lazy_Types_nil()), gopurs_runtime.Apply2(go__go_7_5_1, Get_Data_List_Lazy_Types_nil(), Call_Data_List_Lazy_zipWith(f_1, xs_2, ys_3)))
}

func Call_Data_List_Lazy_updateAt(n_0_loop int64, x_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
updateAt:
for {
if false { continue updateAt }
var n_0 int64 = n_0_loop
_ = n_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)
_ = __local_var_4_0
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr == nil) {
__t2 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_2
} else {

}
}
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr != nil) {
var __t1 *Constructor_Data_List_Lazy_Types_Cons
{
if (n_0) == (0) {
__t1 = &Constructor_Data_List_Lazy_Types_Cons{1, x_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V1}
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V0, Call_Data_List_Lazy_updateAt((n_0) - (1), x_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V1)}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
}))
}
}

func Call_Data_List_Lazy_unzip(xs_0_loop gopurs_runtime.Value) *Constructor_Data_Tuple_Tuple {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var go__go_1_0_2 gopurs_runtime.Value
go__go_1_0_2 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_0_2:
for {
if false { continue go__go_1_0_2 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
// TAST (Let): v_4_1 -> *Constructor_Data_List_Lazy_Types_Cons
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_3))
_ = v_4_1
var __t4 gopurs_runtime.Value
{
if (v_4_1 == nil) {
__t4 = b_2
goto end_branch_4
} else {

}
}
{
if (v_4_1 != nil) {
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := (*Constructor_Data_Tuple_Tuple)((v_4_1).V0.UnsafePtr).V0
_ = __local_var_5_2
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
__local_var_6_3 := (*Constructor_Data_Tuple_Tuple)((v_4_1).V0.UnsafePtr).V1
_ = __local_var_6_3
b_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_5_2, (*Constructor_Data_Tuple_Tuple)(b_2.UnsafePtr).V0})}
})), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_6_3, (*Constructor_Data_Tuple_Tuple)(b_2.UnsafePtr).V1})}
}))})}
xs_3_loop = (v_4_1).V1
continue go__go_1_0_2
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
var go__go_2_5_3 gopurs_runtime.Value
go__go_2_5_3 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var xs_4_loop gopurs_runtime.Value = xs_4_loop_val
go__go_2_5_3:
for {
if false { continue go__go_2_5_3 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var xs_4 gopurs_runtime.Value = xs_4_loop
_ = xs_4
// TAST (Let): v_5_6 -> *Constructor_Data_List_Lazy_Types_Cons
v_5_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_4))
_ = v_5_6
var __t8 gopurs_runtime.Value
{
if (v_5_6 == nil) {
__t8 = b_3
goto end_branch_8
} else {

}
}
{
if (v_5_6 != nil) {
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := (v_5_6).V0
_ = __local_var_6_7
b_3_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_6_7, b_3})}
}))
xs_4_loop = (v_5_6).V1
continue go__go_2_5_3
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
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(go__go_1_0_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, Get_Data_List_Lazy_Types_nil(), Get_Data_List_Lazy_Types_nil()})}, gopurs_runtime.Apply2(go__go_2_5_3, Get_Data_List_Lazy_Types_nil(), xs_0)))
}

func Call_Data_List_Lazy_uncons(xs_0_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
// TAST (Let): v_1_0 -> *Constructor_Data_List_Lazy_Types_Cons
v_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_0))
_ = v_1_0
var __t1 *Constructor_Data_Maybe_Just
{
if (v_1_0 == nil) {
__t1 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_1
} else {

}
}
{
if (v_1_0 != nil) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("head", "tail", (v_1_0).V0, (v_1_0).V1)}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)})
}

func Call_Data_List_Lazy_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable = dictUnfoldable_0_loop
_ = dictUnfoldable_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable_0.V1), gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> *Constructor_Data_Maybe_Just
__local_var_2_0 := Call_Data_List_Lazy_uncons(xs_1)
_ = __local_var_2_0
var __t1 *Constructor_Data_Maybe_Just
{
if (__local_var_2_0 != nil) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.RecordGet((__local_var_2_0).V0, "head"), gopurs_runtime.RecordGet((__local_var_2_0).V0, "tail")})}}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}))
}

func Call_Data_List_Lazy_takeWhile(p_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
takeWhile:
for {
if false { continue takeWhile }
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), x_1)
_ = __local_var_3_0
var __t1 *Constructor_Data_List_Lazy_Types_Cons
{
if ((__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil)) && ((gopurs_runtime.Apply(p_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V0).IntVal) != (0)) {
__t1 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V0, Call_Data_List_Lazy_takeWhile(p_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V1)}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t1)}
}))
}
}

func Call_Data_List_Lazy_take(n_0_loop int64) gopurs_runtime.Value {
take:
for {
if false { continue take }
var n_0 int64 = n_0_loop
_ = n_0
var __t3 gopurs_runtime.Value
{
var __t2 bool
{
if (n_0) > (0) {
__t2 = false
goto end_branch_2
} else {

}
}
{
__t2 = true
}
end_branch_2:
if __t2 {
__t3 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_nil()
})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), x_1)
_ = __local_var_3_0
var __t1 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t1 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_1
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
__t1 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V0, gopurs_runtime.Apply(Call_Data_List_Lazy_take((n_0) - (1)), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V1)}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t1)}
}))
})
}
end_branch_3:
return __t3
}
}

func Call_Data_List_Lazy_tail(xs_0_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
// TAST (Let): __local_var_1_0 -> *Constructor_Data_Maybe_Just
__local_var_1_0 := Call_Data_List_Lazy_uncons(xs_0)
_ = __local_var_1_0
var __t1 *Constructor_Data_Maybe_Just
{
if (__local_var_1_0 != nil) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordGet((__local_var_1_0).V0, "tail")}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
return __t1
}

func Call_Data_List_Lazy_stripPrefix(dictEq_0_loop *Constructor_Data_Eq_Eq, v_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Func(func(prefix_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(input_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_5_1 -> *Constructor_Data_List_Lazy_Types_Cons
v1_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), prefix_3))
_ = v1_5_1
var __t4 *Constructor_Data_Maybe_Just
{
if (v1_5_1 == nil) {
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, input_4})}}
goto end_branch_4
} else {

}
}
{
if (v1_5_1 != nil) {
// TAST (Let): v2_6_2 -> *Constructor_Data_List_Lazy_Types_Cons
v2_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), input_4))
_ = v2_6_2
var __t3 *Constructor_Data_Maybe_Just
{
if ((v2_6_2 != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), (v1_5_1).V0, (v2_6_2).V0).IntVal) != (0)) {
__t3 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.RecordDict2("a", "b", (v1_5_1).V1, (v2_6_2).V1)})}}
goto end_branch_3
} else {

}
}
{
__t3 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
})
})
_ = __local_var_3_0
// TAST (Let): __local_var_4_5 -> gopurs_runtime.Value
__local_var_4_5 := gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 *Constructor_Control_Monad_Rec_Class_Done
{
if (v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr == nil) {
__t9 = &Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}}
goto end_branch_9
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 930809136 && v_4.UnsafePtr != nil) {
var __t8 gopurs_runtime.Value
{
var __t_tag_6 gopurs_runtime.Value = (*Constructor_Data_Maybe_Just)(v_4.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 525585346) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(__local_var_3_0, gopurs_runtime.RecordGet((*Constructor_Control_Monad_Rec_Class_Loop)((*Constructor_Data_Maybe_Just)(v_4.UnsafePtr).V0.UnsafePtr).V0, "a"), gopurs_runtime.RecordGet((*Constructor_Control_Monad_Rec_Class_Loop)((*Constructor_Data_Maybe_Just)(v_4.UnsafePtr).V0.UnsafePtr).V0, "b"))))}})}
goto end_branch_8
} else {

}
}
{
var __t_tag_7 gopurs_runtime.Value = (*Constructor_Data_Maybe_Just)(v_4.UnsafePtr).V0
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 60402430) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, (*Constructor_Control_Monad_Rec_Class_Done)((*Constructor_Data_Maybe_Just)(v_4.UnsafePtr).V0.UnsafePtr).V0})}})}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_Done](__t8)
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_Done](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(__t9)}
})
_ = __local_var_4_5
var go__go_5_10_4 gopurs_runtime.Value
go__go_5_10_4 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_5_10_4:
for {
if false { continue go__go_5_10_4 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t11 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 525585346) {
v_6_loop = gopurs_runtime.Apply(__local_var_4_5, (*Constructor_Control_Monad_Rec_Class_Loop)(v_6.UnsafePtr).V0)
continue go__go_5_10_4
__t11 = gopurs_runtime.Value{}
goto end_branch_11
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 60402430) {
__t11 = (*Constructor_Control_Monad_Rec_Class_Done)(v_6.UnsafePtr).V0
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
}()
})
// TAST (Let): __local_var_6_12 -> gopurs_runtime.Value
__local_var_6_12 := gopurs_runtime.RecordDict2("a", "b", v_1, s_2)
_ = __local_var_6_12
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(go__go_5_10_4, gopurs_runtime.Apply(__local_var_4_5, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(__local_var_3_0, gopurs_runtime.RecordGet(__local_var_6_12, "a"), gopurs_runtime.RecordGet(__local_var_6_12, "b"))))})))
}

func Call_Data_List_Lazy_span(p_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
span:
for {
if false { continue span }
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
// TAST (Let): v_2_0 -> *Constructor_Data_Maybe_Just
v_2_0 := Call_Data_List_Lazy_uncons(xs_1)
_ = v_2_0
var __t2 gopurs_runtime.Value
{
if ((v_2_0 != nil)) && ((gopurs_runtime.Apply(p_0, gopurs_runtime.RecordGet((v_2_0).V0, "head")).IntVal) != (0)) {
// TAST (Let): v1_3_1 -> gopurs_runtime.Value
var v1_3_1 gopurs_runtime.Value = Call_Data_List_Lazy_span(p_0, gopurs_runtime.RecordGet((v_2_0).V0, "tail"))
__t2 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.RecordGet((v_2_0).V0, "head"), gopurs_runtime.RecordGet(v1_3_1, "init")})}
})), gopurs_runtime.RecordGet(v1_3_1, "rest"))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict2("init", "rest", Get_Data_List_Lazy_Types_nil(), xs_1)
}
end_branch_2:
return __t2
}
}

func Call_Data_List_Lazy_snoc(xs_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var go__go_2_0_5 gopurs_runtime.Value
go__go_2_0_5 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var xs_4_loop gopurs_runtime.Value = xs_4_loop_val
go__go_2_0_5:
for {
if false { continue go__go_2_0_5 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var xs_4 gopurs_runtime.Value = xs_4_loop
_ = xs_4
// TAST (Let): v_5_1 -> *Constructor_Data_List_Lazy_Types_Cons
v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_4))
_ = v_5_1
var __t2 gopurs_runtime.Value
{
if (v_5_1 == nil) {
__t2 = b_3
goto end_branch_2
} else {

}
}
{
if (v_5_1 != nil) {
b_3_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, (v_5_1).V0, b_3})}
}))
xs_4_loop = (v_5_1).V1
continue go__go_2_0_5
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
var go__go_3_3_6 gopurs_runtime.Value
go__go_3_3_6 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var xs_5_loop gopurs_runtime.Value = xs_5_loop_val
go__go_3_3_6:
for {
if false { continue go__go_3_3_6 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_4 -> *Constructor_Data_List_Lazy_Types_Cons
v_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
_ = v_6_4
var __t6 gopurs_runtime.Value
{
if (v_6_4 == nil) {
__t6 = b_4
goto end_branch_6
} else {

}
}
{
if (v_6_4 != nil) {
// TAST (Let): __local_var_7_5 -> gopurs_runtime.Value
__local_var_7_5 := (v_6_4).V0
_ = __local_var_7_5
b_4_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_7_5, b_4})}
}))
xs_5_loop = (v_6_4).V1
continue go__go_3_3_6
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
return gopurs_runtime.Apply2(go__go_2_0_5, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, x_1, Get_Data_List_Lazy_Types_nil()})}
})), gopurs_runtime.Apply2(go__go_3_3_6, Get_Data_List_Lazy_Types_nil(), xs_0))
}

func Call_Data_List_Lazy_singleton(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, a_0, Get_Data_List_Lazy_Types_nil()})}
}))
}

func Call_Data_List_Lazy_showPattern(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): showList_1_0 -> *Constructor_Data_Show_Show
showList_1_0 := &Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_2_1 -> *Constructor_Data_List_Lazy_Types_Cons
v_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1))
_ = v_2_1
var __t5 string
{
if (v_2_1 == nil) {
__t5 = "(fromFoldable [])"
goto end_branch_5
} else {

}
}
{
if (v_2_1 != nil) {
var go__go_3_2_7 gopurs_runtime.Value
go__go_3_2_7 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var xs_5_loop gopurs_runtime.Value = xs_5_loop_val
go__go_3_2_7:
for {
if false { continue go__go_3_2_7 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_3 -> *Constructor_Data_List_Lazy_Types_Cons
v_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
_ = v_6_3
var __t4 gopurs_runtime.Value
{
if (v_6_3 == nil) {
__t4 = b_4
goto end_branch_4
} else {

}
}
{
if (v_6_3 != nil) {
b_4_loop = gopurs_runtime.Str(((b_4.StrVal()) + (",")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (v_6_3).V0).StrVal()))
xs_5_loop = (v_6_3).V1
continue go__go_3_2_7
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
__t5 = ((("(fromFoldable [") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (v_2_1).V0).StrVal())) + (gopurs_runtime.Apply2(go__go_3_2_7, gopurs_runtime.Str(""), (v_2_1).V1).StrVal())) + ("])")
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_5:
return gopurs_runtime.Str(__t5)
})}
_ = showList_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Pattern ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showList_1_0.V0), v_2).StrVal())) + (")"))
})})}
}

func Call_Data_List_Lazy_scanlLazy(f_0_loop gopurs_runtime.Value, acc_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
scanlLazy:
for {
if false { continue scanlLazy }
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var acc_1 gopurs_runtime.Value = acc_1_loop
_ = acc_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)
_ = __local_var_4_0
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr == nil) {
__t2 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_2
} else {

}
}
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr != nil) {
// TAST (Let): acc_prime_5_1 -> gopurs_runtime.Value
acc_prime_5_1 := gopurs_runtime.Apply2(f_0, acc_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V0)
_ = acc_prime_5_1
__t2 = &Constructor_Data_List_Lazy_Types_Cons{1, acc_prime_5_1, Call_Data_List_Lazy_scanlLazy(f_0, acc_prime_5_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V1)}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
}))
}
}

func Call_Data_List_Lazy_reverse(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_8 gopurs_runtime.Value
go__go_2_0_8 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var xs_4_loop gopurs_runtime.Value = xs_4_loop_val
go__go_2_0_8:
for {
if false { continue go__go_2_0_8 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var xs_4 gopurs_runtime.Value = xs_4_loop
_ = xs_4
// TAST (Let): v_5_1 -> *Constructor_Data_List_Lazy_Types_Cons
v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_4))
_ = v_5_1
var __t3 gopurs_runtime.Value
{
if (v_5_1 == nil) {
__t3 = b_3
goto end_branch_3
} else {

}
}
{
if (v_5_1 != nil) {
// TAST (Let): __local_var_6_2 -> gopurs_runtime.Value
__local_var_6_2 := (v_5_1).V0
_ = __local_var_6_2
b_3_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_6_2, b_3})}
}))
xs_4_loop = (v_5_1).V1
continue go__go_2_0_8
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
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply2(go__go_2_0_8, Get_Data_List_Lazy_Types_nil(), xs_0))
}))
}

func Call_Data_List_Lazy_replicateM(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
replicateM:
for {
if false { continue replicateM }
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(n_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t2 bool
{
if (n_3.IntVal) < (1) {
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
__t3 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), Get_Data_List_Lazy_Types_nil())
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(Call_Data_List_Lazy_replicateM(dictMonad_0), gopurs_runtime.Int((n_3.IntVal) - (1)), m_4), gopurs_runtime.Func(func(as_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, a_5, as_6})}
})))
}))
}))
}
end_branch_3:
return __t3
})
})
}
}

func Call_Data_List_Lazy_repeat(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var go__go_1_0_9 gopurs_runtime.Value
_ = go__go_1_0_9
go__go_1_0_9 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, x_0, go__go_1_0_9})}
})))
}))
return go__go_1_0_9
}

func Call_Data_List_Lazy_replicate(i_0_loop int64, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var i_0 int64 = i_0_loop
_ = i_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(Call_Data_List_Lazy_take(i_0), Call_Data_List_Lazy_repeat(xs_1))
}

func Call_Data_List_Lazy_go__range(start_0_loop int64, end_1_loop int64) gopurs_runtime.Value {
var start_0 int64 = start_0_loop
_ = start_0
var end_1 int64 = end_1_loop
_ = end_1
var __t5 gopurs_runtime.Value
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
__t5 = gopurs_runtime.Apply2(Get_Data_Unfoldable_unfoldr__193332035(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_Maybe_Just
{
var __t3 bool
{
if (x_2.IntVal) < (end_1) {
__t3 = false
goto end_branch_3
} else {

}
}
{
__t3 = true
}
end_branch_3:
if __t3 {
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(x_2.IntVal), gopurs_runtime.Int((x_2.IntVal) - (1))})}}
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
}), gopurs_runtime.Int(start_0))
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Apply2(Get_Data_Unfoldable_unfoldr__193332035(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just
{
var __t0 bool
{
if (x_2.IntVal) > (end_1) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
if __t0 {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(x_2.IntVal), gopurs_runtime.Int((x_2.IntVal) + (1))})}}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}), gopurs_runtime.Int(start_0))
}
end_branch_5:
return __t5
}

func Call_Data_List_Lazy_partition(f_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var go__go_2_0_10 gopurs_runtime.Value
go__go_2_0_10 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var xs_4_loop gopurs_runtime.Value = xs_4_loop_val
go__go_2_0_10:
for {
if false { continue go__go_2_0_10 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var xs_4 gopurs_runtime.Value = xs_4_loop
_ = xs_4
// TAST (Let): v_5_1 -> *Constructor_Data_List_Lazy_Types_Cons
v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_4))
_ = v_5_1
var __t3 gopurs_runtime.Value
{
if (v_5_1 == nil) {
__t3 = b_3
goto end_branch_3
} else {

}
}
{
if (v_5_1 != nil) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_0, (v_5_1).V0).IntVal) != (0) {
__t2 = gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.RecordGet(b_3, "no"), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, (v_5_1).V0, gopurs_runtime.RecordGet(b_3, "yes")})}
})))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, (v_5_1).V0, gopurs_runtime.RecordGet(b_3, "no")})}
})), gopurs_runtime.RecordGet(b_3, "yes"))
}
end_branch_2:
b_3_loop = __t2
xs_4_loop = (v_5_1).V1
continue go__go_2_0_10
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
var go__go_3_4_11 gopurs_runtime.Value
go__go_3_4_11 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var xs_5_loop gopurs_runtime.Value = xs_5_loop_val
go__go_3_4_11:
for {
if false { continue go__go_3_4_11 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_5 -> *Constructor_Data_List_Lazy_Types_Cons
v_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
_ = v_6_5
var __t7 gopurs_runtime.Value
{
if (v_6_5 == nil) {
__t7 = b_4
goto end_branch_7
} else {

}
}
{
if (v_6_5 != nil) {
// TAST (Let): __local_var_7_6 -> gopurs_runtime.Value
__local_var_7_6 := (v_6_5).V0
_ = __local_var_7_6
b_4_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_7_6, b_4})}
}))
xs_5_loop = (v_6_5).V1
continue go__go_3_4_11
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
return gopurs_runtime.Apply2(go__go_2_0_10, gopurs_runtime.RecordDict2("no", "yes", Get_Data_List_Lazy_Types_nil(), Get_Data_List_Lazy_Types_nil()), gopurs_runtime.Apply2(go__go_3_4_11, Get_Data_List_Lazy_Types_nil(), xs_1))
}

func Call_Data_List_Lazy_null(x_0_loop gopurs_runtime.Value) bool {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
var __local_var_1_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_uncons(x_0))}
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136 && __local_var_1_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136 && __local_var_1_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return (__t1.IntVal) != (0)
}

func Call_Data_List_Lazy_nubBy(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var goStep_1_0_12 gopurs_runtime.Value
_ = goStep_1_0_12
var go__go_1_1_13 gopurs_runtime.Value
_ = go__go_1_1_13
goStep_1_0_12 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_List_Lazy_Types_Cons
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr == nil) {
__t4 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_4
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr != nil) {
// TAST (Let): v2_4_2 -> gopurs_runtime.Value
v2_4_2 := gopurs_runtime.Apply3(Get_Data_List_Internal_insertAndLookupBy(), p_0, (*Constructor_Data_List_Lazy_Types_Cons)(v1_3.UnsafePtr).V0, v_2)
_ = v2_4_2
var __t3 *Constructor_Data_List_Lazy_Types_Cons
{
if (gopurs_runtime.RecordGet(v2_4_2, "found").IntVal) != (0) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply2(go__go_1_1_13, gopurs_runtime.RecordGet(v2_4_2, "result"), (*Constructor_Data_List_Lazy_Types_Cons)(v1_3.UnsafePtr).V1)))
goto end_branch_3
} else {

}
}
{
__t3 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(v1_3.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_1_1_13, gopurs_runtime.RecordGet(v2_4_2, "result"), (*Constructor_Data_List_Lazy_Types_Cons)(v1_3.UnsafePtr).V1)}
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t4)}
})
})
go__go_1_1_13 = gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_5 -> gopurs_runtime.Value
__local_var_4_5 := gopurs_runtime.Apply(goStep_1_0_12, s_2)
_ = __local_var_4_5
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_5, gopurs_runtime.Apply(Get_Data_Lazy_force(), v_3))
}))
})
})
return gopurs_runtime.Apply(go__go_1_1_13, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)})
}

func Call_Data_List_Lazy_nub(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
return Call_Data_List_Lazy_nubBy(gopurs_runtime.Box(dictOrd_0.V1))
}

func Call_Data_List_Lazy_mapMaybe(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
mapMaybe:
for {
if false { continue mapMaybe }
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_14 gopurs_runtime.Value
go__go_1_0_14 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_14:
for {
if false { continue go__go_1_0_14 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t3 *Constructor_Data_List_Lazy_Types_Cons
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t3 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
// TAST (Let): v1_3_1 -> *Constructor_Data_Maybe_Just
v1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V0))
_ = v1_3_1
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (v1_3_1 == nil) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V1)))}
continue go__go_1_0_14
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
if (v1_3_1 != nil) {
__t2 = &Constructor_Data_List_Lazy_Types_Cons{1, (v1_3_1).V0, gopurs_runtime.Apply(Call_Data_List_Lazy_mapMaybe(f_0), (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V1)}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_1_0_14, gopurs_runtime.Apply(Get_Data_Lazy_force(), x_2))
}))
})
}
}

func Call_Data_List_Lazy_some(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative) gopurs_runtime.Value {
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
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_List_Lazy_Types_cons(), v_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_3, "defer"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Data_List_Lazy_many(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy](dictLazy_3))}, v_4)
})))
})
})
}

func Call_Data_List_Lazy_many(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative) gopurs_runtime.Value {
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
return gopurs_runtime.Apply2(gopurs_runtime.Box(Alt0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V0), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V1), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_List_Lazy_Types_cons(), v_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_3, "defer"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Data_List_Lazy_many(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy](dictLazy_3))}, v_4)
}))), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), Get_Data_List_Lazy_Types_nil()))
})
})
}

func Call_Data_List_Lazy_iterate(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var go__go_2_0_17 gopurs_runtime.Value
_ = go__go_2_0_17
go__go_2_0_17 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.Apply(Get_Data_Lazy_force(), go__go_2_0_17)
_ = __local_var_5_2
var __t3 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_5_2.Type == 9 && __local_var_5_2.IntVal == 218341868 && __local_var_5_2.UnsafePtr == nil) {
__t3 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_3
} else {

}
}
{
if (__local_var_5_2.Type == 9 && __local_var_5_2.IntVal == 218341868 && __local_var_5_2.UnsafePtr != nil) {
__t3 = &Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_5_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()).V0), f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_5_2.UnsafePtr).V1)}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t3)}
}))
_ = __local_var_4_1
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, x_1, __local_var_4_1})}
})))
}))
return go__go_2_0_17
}

func Call_Data_List_Lazy_insertAt(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
insertAt:
for {
if false { continue insertAt }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t2 gopurs_runtime.Value
{
if (v_0) == (0) {
__t2 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, v1_1, v2_2})}
}))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v2_2)
_ = __local_var_4_0
var __t1 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr == nil) {
__t1 = &Constructor_Data_List_Lazy_Types_Cons{1, v1_1, Get_Data_List_Lazy_Types_nil()}
goto end_branch_1
} else {

}
}
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr != nil) {
__t1 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V0, Call_Data_List_Lazy_insertAt((v_0) - (1), v1_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V1)}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t1)}
}))
}
end_branch_2:
return __t2
}
}

func Call_Data_List_Lazy_index(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var go__go_1_0_19 gopurs_runtime.Value
go__go_1_0_19 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop int64 = v1_3_loop_val.IntVal
go__go_1_0_19:
for {
if false { continue go__go_1_0_19 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 int64 = v1_3_loop
_ = v1_3
var __t2 *Constructor_Data_Maybe_Just
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t2 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_2
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
var __t1 *Constructor_Data_Maybe_Just
{
if (v1_3) == (0) {
__t1 = &Constructor_Data_Maybe_Just{1, (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V0}
goto end_branch_1
} else {

}
}
{
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V1)))}
v1_3_loop = (v1_3) - (1)
continue go__go_1_0_19
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
}
end_branch_1:
__t2 = __t1
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
return gopurs_runtime.Apply(go__go_1_0_19, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_0)))})
}

func Call_Data_List_Lazy_head(xs_0_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
// TAST (Let): __local_var_1_0 -> *Constructor_Data_Maybe_Just
__local_var_1_0 := Call_Data_List_Lazy_uncons(xs_0)
_ = __local_var_1_0
var __t1 *Constructor_Data_Maybe_Just
{
if (__local_var_1_0 != nil) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordGet((__local_var_1_0).V0, "head")}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
return __t1
}

func Call_Data_List_Lazy_transpose(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
transpose:
for {
if false { continue transpose }
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
// TAST (Let): v_1_0 -> *Constructor_Data_Maybe_Just
v_1_0 := Call_Data_List_Lazy_uncons(xs_0)
_ = v_1_0
var __t7 gopurs_runtime.Value
{
if (v_1_0 == nil) {
__t7 = xs_0
goto end_branch_7
} else {

}
}
{
if (v_1_0 != nil) {
// TAST (Let): v1_2_1 -> *Constructor_Data_Maybe_Just
v1_2_1 := Call_Data_List_Lazy_uncons(gopurs_runtime.RecordGet((v_1_0).V0, "head"))
_ = v1_2_1
var __t6 gopurs_runtime.Value
{
if (v1_2_1 == nil) {
xs_0_loop = gopurs_runtime.RecordGet((v_1_0).V0, "tail")
continue transpose
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
if (v1_2_1 != nil) {
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(Call_Data_List_Lazy_mapMaybe(Get_Data_List_Lazy_head()), gopurs_runtime.RecordGet((v_1_0).V0, "tail"))
_ = __local_var_3_3
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.RecordGet((v1_2_1).V0, "head"), __local_var_3_3})}
}))
_ = __local_var_3_2
// TAST (Let): __local_var_4_5 -> gopurs_runtime.Value
__local_var_4_5 := gopurs_runtime.Apply(Call_Data_List_Lazy_mapMaybe(Get_Data_List_Lazy_tail()), gopurs_runtime.RecordGet((v_1_0).V0, "tail"))
_ = __local_var_4_5
// TAST (Let): __local_var_4_4 -> gopurs_runtime.Value
__local_var_4_4 := Call_Data_List_Lazy_transpose(gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.RecordGet((v1_2_1).V0, "tail"), __local_var_4_5})}
})))
_ = __local_var_4_4
__t6 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_3_2, __local_var_4_4})}
}))
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
__t7 = __t6
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
}

func Call_Data_List_Lazy_groupBy(eq_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
groupBy:
for {
if false { continue groupBy }
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), x_1)
_ = __local_var_3_0
var __t4 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t4 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_4
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V0
_ = __local_var_4_1
// TAST (Let): v1_5_2 -> gopurs_runtime.Value
var v1_5_2 gopurs_runtime.Value = Call_Data_List_Lazy_span(gopurs_runtime.Apply(eq_0, __local_var_4_1), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V1)
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
__local_var_6_3 := gopurs_runtime.RecordGet(v1_5_2, "init")
_ = __local_var_6_3
__t4 = &Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, __local_var_4_1, __local_var_6_3})}
})), Call_Data_List_Lazy_groupBy(eq_0, gopurs_runtime.RecordGet(v1_5_2, "rest"))}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t4)}
}))
}
}

func Call_Data_List_Lazy_group(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_List_Lazy_groupBy(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_List_Lazy_fromStep(x_0_loop *Constructor_Data_List_Lazy_Types_Cons) gopurs_runtime.Value {
var x_0 *Constructor_Data_List_Lazy_Types_Cons = x_0_loop
_ = x_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(x_0)}
}))
}

func Call_Data_List_Lazy_insertBy(cmp_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
insertBy:
for {
if false { continue insertBy }
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)
_ = __local_var_4_0
var __t3 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr == nil) {
__t3 = &Constructor_Data_List_Lazy_Types_Cons{1, x_1, Get_Data_List_Lazy_Types_nil()}
goto end_branch_3
} else {

}
}
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr != nil) {
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(cmp_0, x_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V0)
if (uint32(__t_tag_1.IntVal) == 380165415) {
__t2 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V0, Call_Data_List_Lazy_insertBy(cmp_0, x_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V1)}
goto end_branch_2
} else {

}
}
{
__t2 = &Constructor_Data_List_Lazy_Types_Cons{1, x_1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](__local_var_4_0))}
}))}
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t3)}
}))
}
}

func Call_Data_List_Lazy_insert(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_List_Lazy_insertBy(), gopurs_runtime.Box(dictOrd_0.V1))
}

func Call_Data_List_Lazy_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V2), Get_Data_List_Lazy_Types_cons(), Get_Data_List_Lazy_Types_nil())
}

func Call_Data_List_Lazy_foldrLazy(dictLazy_0_loop *Constructor_Control_Lazy_Lazy, op_1_loop gopurs_runtime.Value, z_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictLazy_0 *Constructor_Control_Lazy_Lazy = dictLazy_0_loop
_ = dictLazy_0
var op_1 gopurs_runtime.Value = op_1_loop
_ = op_1
var z_2 gopurs_runtime.Value = z_2_loop
_ = z_2
var go__go_3_0_20 gopurs_runtime.Value
_ = go__go_3_0_20
go__go_3_0_20 = gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_5_1 -> *Constructor_Data_List_Lazy_Types_Cons
v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_4))
_ = v_5_1
var __t4 gopurs_runtime.Value
{
if (v_5_1 != nil) {
// TAST (Let): __local_var_6_2 -> gopurs_runtime.Value
__local_var_6_2 := (v_5_1).V0
_ = __local_var_6_2
// TAST (Let): __local_var_7_3 -> gopurs_runtime.Value
__local_var_7_3 := (v_5_1).V1
_ = __local_var_7_3
__t4 = gopurs_runtime.Apply(gopurs_runtime.Box(dictLazy_0.V0), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_1, __local_var_6_2, gopurs_runtime.Apply(go__go_3_0_20, __local_var_7_3))
}))
goto end_branch_4
} else {

}
}
{
if (v_5_1 == nil) {
__t4 = z_2
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
})
return go__go_3_0_20
}

func Call_Data_List_Lazy_foldM(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
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
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_6_2 -> *Constructor_Data_Maybe_Just
v_6_2 := Call_Data_List_Lazy_uncons(xs_5)
_ = v_6_2
var __t4 gopurs_runtime.Value
{
if (v_6_2 == nil) {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), b_4)
goto end_branch_4
} else {

}
}
{
if (v_6_2 != nil) {
// TAST (Let): __local_var_7_3 -> gopurs_runtime.Value
__local_var_7_3 := gopurs_runtime.RecordGet((v_6_2).V0, "tail")
_ = __local_var_7_3
__t4 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(f_3, b_4, gopurs_runtime.RecordGet((v_6_2).V0, "head")), gopurs_runtime.Func(func(b_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Call_Data_List_Lazy_foldM(dictMonad_0), f_3, b_prime_8, __local_var_7_3)
}))
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
})
})
})
}
}

func Call_Data_List_Lazy_findIndex(fn_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var go__go_1_0_21 gopurs_runtime.Value
go__go_1_0_21 = gopurs_runtime.Func(func(n_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(list_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var n_2_loop int64 = n_2_loop_val.IntVal
var list_3_loop gopurs_runtime.Value = list_3_loop_val
go__go_1_0_21:
for {
if false { continue go__go_1_0_21 }
var n_2 int64 = n_2_loop
_ = n_2
var list_3 gopurs_runtime.Value = list_3_loop
_ = list_3
// TAST (Let): __local_var_4_1 -> *Constructor_Data_Maybe_Just
__local_var_4_1 := Call_Data_List_Lazy_uncons(list_3)
_ = __local_var_4_1
var __t3 *Constructor_Data_Maybe_Just
{
if (__local_var_4_1 != nil) {
var __t2 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.Apply(fn_0, gopurs_runtime.RecordGet((__local_var_4_1).V0, "head")).IntVal) != (0) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(n_2)})})
goto end_branch_2
} else {

}
}
{
n_2_loop = (n_2) + (1)
list_3_loop = gopurs_runtime.RecordGet((__local_var_4_1).V0, "tail")
continue go__go_1_0_21
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
if (__local_var_4_1 == nil) {
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
return gopurs_runtime.Apply(go__go_1_0_21, gopurs_runtime.Int(0))
}

func Call_Data_List_Lazy_findLastIndex(fn_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
// TAST (Let): __local_var_2_0 -> *Constructor_Data_Maybe_Just
__local_var_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Call_Data_List_Lazy_findIndex(fn_0), Call_Data_List_Lazy_reverse(xs_1)))
_ = __local_var_2_0
var __t1 *Constructor_Data_Maybe_Just
{
if (__local_var_2_0 != nil) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(((gopurs_runtime.Apply(Get_Data_List_Lazy_length(), xs_1).IntVal) - (1)) - ((__local_var_2_0).V0.IntVal))}
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

func Call_Data_List_Lazy_filterM(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
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
return gopurs_runtime.Func(func(p_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(list_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_5_2 -> *Constructor_Data_Maybe_Just
v_5_2 := Call_Data_List_Lazy_uncons(list_4)
_ = v_5_2
var __t6 gopurs_runtime.Value
{
if (v_5_2 == nil) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), Get_Data_List_Lazy_Types_nil())
goto end_branch_6
} else {

}
}
{
if (v_5_2 != nil) {
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
__local_var_6_3 := gopurs_runtime.RecordGet((v_5_2).V0, "head")
_ = __local_var_6_3
// TAST (Let): __local_var_7_4 -> gopurs_runtime.Value
__local_var_7_4 := gopurs_runtime.RecordGet((v_5_2).V0, "tail")
_ = __local_var_7_4
__t6 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(p_3, __local_var_6_3), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(Call_Data_List_Lazy_filterM(dictMonad_0), p_3, __local_var_7_4), gopurs_runtime.Func(func(xs_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (b_8.IntVal) != (0) {
__t5 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_6_3, xs_prime_9})}
}))
goto end_branch_5
} else {

}
}
{
__t5 = xs_prime_9
}
end_branch_5:
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), __t5)
}))
}))
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
})
})
}
}

func Call_Data_List_Lazy_filter(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
filter:
for {
if false { continue filter }
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_22 gopurs_runtime.Value
go__go_1_0_22 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_22:
for {
if false { continue go__go_1_0_22 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t2 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_2
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
var __t1 *Constructor_Data_List_Lazy_Types_Cons
{
if (gopurs_runtime.Apply(p_0, (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V0).IntVal) != (0) {
__t1 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V0, gopurs_runtime.Apply(Call_Data_List_Lazy_filter(p_0), (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V1)}
goto end_branch_1
} else {

}
}
{
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V1)))}
continue go__go_1_0_22
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Value{})
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
}
}()
})
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_1_0_22, gopurs_runtime.Apply(Get_Data_Lazy_force(), x_2))
}))
})
}
}

func Call_Data_List_Lazy_intersectBy(eq_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
return gopurs_runtime.Apply(Call_Data_List_Lazy_filter(gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupDisj1_4_0 -> *Constructor_Data_Semigroup_Semigroup
semigroupDisj1_4_0 := &Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((v_4.IntVal) != (0)) || ((v1_5.IntVal) != (0)))
})
})}
_ = semigroupDisj1_4_0
// TAST (Let): __local_var_5_1 -> *Constructor_Data_Monoid_Monoid
__local_var_5_1 := &Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupDisj1_4_0)}
}), gopurs_runtime.Bool(false)}
_ = __local_var_5_1
// TAST (Let): Semigroup0_6_2 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_5_1.V0), gopurs_runtime.Value{}))
_ = Semigroup0_6_2
// TAST (Let): __local_var_7_3 -> gopurs_runtime.Value
__local_var_7_3 := gopurs_runtime.Apply(eq_0, x_3)
_ = __local_var_7_3
var go__go_8_4_23 gopurs_runtime.Value
go__go_8_4_23 = gopurs_runtime.Func(func(b_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_9_loop gopurs_runtime.Value = b_9_loop_val
var xs_10_loop gopurs_runtime.Value = xs_10_loop_val
go__go_8_4_23:
for {
if false { continue go__go_8_4_23 }
var b_9 gopurs_runtime.Value = b_9_loop
_ = b_9
var xs_10 gopurs_runtime.Value = xs_10_loop
_ = xs_10
// TAST (Let): v_11_5 -> *Constructor_Data_List_Lazy_Types_Cons
v_11_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_10))
_ = v_11_5
var __t6 gopurs_runtime.Value
{
if (v_11_5 == nil) {
__t6 = b_9
goto end_branch_6
} else {

}
}
{
if (v_11_5 != nil) {
b_9_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_2.V0), b_9, gopurs_runtime.Apply(__local_var_7_3, (v_11_5).V0))
xs_10_loop = (v_11_5).V1
continue go__go_8_4_23
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
return gopurs_runtime.Bool((gopurs_runtime.Apply2(go__go_8_4_23, gopurs_runtime.Box(__local_var_5_1.V1), ys_2).IntVal) != (0))
})), xs_1)
}

func Call_Data_List_Lazy_intersect(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_List_Lazy_intersectBy(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_List_Lazy_nubByEq(eq_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
nubByEq:
for {
if false { continue nubByEq }
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), x_1)
_ = __local_var_3_0
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t2 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_2
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V0
_ = __local_var_4_1
__t2 = &Constructor_Data_List_Lazy_Types_Cons{1, __local_var_4_1, Call_Data_List_Lazy_nubByEq(eq_0, gopurs_runtime.Apply(Call_Data_List_Lazy_filter(gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(eq_0, __local_var_4_1, y_5).IntVal) != (0)) != (true))
})), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V1))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
}))
}
}

func Call_Data_List_Lazy_nubEq(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_List_Lazy_nubByEq(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_List_Lazy_eqPattern(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
// TAST (Let): eqList_1_0 -> *Constructor_Data_Eq_Eq
eqList_1_0 := &Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_24 gopurs_runtime.Value
go__go_3_1_24 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_1_24:
for {
if false { continue go__go_3_1_24 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 bool
{
if (v_4.Type == 9 && v_4.IntVal == 218341868 && v_4.UnsafePtr == nil) {
var __t2 bool
{
if (v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr == nil) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
if ((v_4.Type == 9 && v_4.IntVal == 218341868 && v_4.UnsafePtr != nil)) && (((v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Lazy_Types_Cons)(v_4.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons)(v1_5.UnsafePtr).V0).IntVal) != (0))) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_4.UnsafePtr).V1)))}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v1_5.UnsafePtr).V1)))}
continue go__go_3_1_24
__t3 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
return gopurs_runtime.Bool(__t3)
}
}()
})
})
return gopurs_runtime.Bool((gopurs_runtime.Apply2(go__go_3_1_24, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_2)))}).IntVal) != (0))
})
})}
_ = eqList_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(eqList_1_0.V0), x_2, y_3).IntVal) != (0))
})
})})}
}

func Call_Data_List_Lazy_ordPattern(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_2
// TAST (Let): eqList1_1_1 -> *Constructor_Data_Eq_Eq
eqList1_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_3_25 gopurs_runtime.Value
go__go_4_3_25 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_3_25:
for {
if false { continue go__go_4_3_25 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t5 bool
{
if (v_5.Type == 9 && v_5.IntVal == 218341868 && v_5.UnsafePtr == nil) {
var __t4 bool
{
if (v1_6.Type == 9 && v1_6.IntVal == 218341868 && v1_6.UnsafePtr == nil) {
__t4 = true
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
if ((v_5.Type == 9 && v_5.IntVal == 218341868 && v_5.UnsafePtr != nil)) && (((v1_6.Type == 9 && v1_6.IntVal == 218341868 && v1_6.UnsafePtr != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "eq"), (*Constructor_Data_List_Lazy_Types_Cons)(v_5.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons)(v1_6.UnsafePtr).V0).IntVal) != (0))) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_5.UnsafePtr).V1)))}
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v1_6.UnsafePtr).V1)))}
continue go__go_4_3_25
__t5 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_5
} else {

}
}
{
__t5 = false
}
end_branch_5:
return gopurs_runtime.Bool(__t5)
}
}()
})
})
return gopurs_runtime.Bool((gopurs_runtime.Apply2(go__go_4_3_25, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_3)))}).IntVal) != (0))
})
})))
_ = eqList1_1_1
// TAST (Let): ordList_1_0 -> *Constructor_Data_Ord_Ord
ordList_1_0 := &Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqList1_1_1)}
}), gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_6_26 gopurs_runtime.Value
go__go_4_6_26 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_6_26:
for {
if false { continue go__go_4_6_26 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t10 uint32
{
if (v_5.Type == 9 && v_5.IntVal == 218341868 && v_5.UnsafePtr == nil) {
var __t7 uint32
{
if (v1_6.Type == 9 && v1_6.IntVal == 218341868 && v1_6.UnsafePtr == nil) {
__t7 = 902936544
goto end_branch_7
} else {

}
}
{
__t7 = 1527465420
}
end_branch_7:
__t10 = __t7
goto end_branch_10
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 218341868 && v1_6.UnsafePtr == nil) {
__t10 = 380165415
goto end_branch_10
} else {

}
}
{
if ((v_5.Type == 9 && v_5.IntVal == 218341868 && v_5.UnsafePtr != nil)) && ((v1_6.Type == 9 && v1_6.IntVal == 218341868 && v1_6.UnsafePtr != nil)) {
// TAST (Let): v2_7_8 -> gopurs_runtime.Value
v2_7_8 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (*Constructor_Data_List_Lazy_Types_Cons)(v_5.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons)(v1_6.UnsafePtr).V0)
_ = v2_7_8
var __t9 uint32
{
if (uint32(v2_7_8.IntVal) == 902936544) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_5.UnsafePtr).V1)))}
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v1_6.UnsafePtr).V1)))}
continue go__go_4_6_26
__t9 = uint32(gopurs_runtime.Value{}.IntVal)
goto end_branch_9
} else {

}
}
{
__t9 = uint32(v2_7_8.IntVal)
}
end_branch_9:
__t10 = __t9
goto end_branch_10
} else {

}
}
{
__t10 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t10), UnsafePtr: nil}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(go__go_4_6_26, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_3)))}).IntVal)), UnsafePtr: nil}
})
})}
_ = ordList_1_0
// TAST (Let): __local_var_2_12 -> gopurs_runtime.Value
__local_var_2_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_12
// TAST (Let): eqList_3_13 -> *Constructor_Data_Eq_Eq
eqList_3_13 := &Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_14_27 gopurs_runtime.Value
go__go_5_14_27 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop gopurs_runtime.Value = v_6_loop_val
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_14_27:
for {
if false { continue go__go_5_14_27 }
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t16 bool
{
if (v_6.Type == 9 && v_6.IntVal == 218341868 && v_6.UnsafePtr == nil) {
var __t15 bool
{
if (v1_7.Type == 9 && v1_7.IntVal == 218341868 && v1_7.UnsafePtr == nil) {
__t15 = true
goto end_branch_15
} else {

}
}
{
__t15 = false
}
end_branch_15:
__t16 = __t15
goto end_branch_16
} else {

}
}
{
if ((v_6.Type == 9 && v_6.IntVal == 218341868 && v_6.UnsafePtr != nil)) && (((v1_7.Type == 9 && v1_7.IntVal == 218341868 && v1_7.UnsafePtr != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_12, "eq"), (*Constructor_Data_List_Lazy_Types_Cons)(v_6.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons)(v1_7.UnsafePtr).V0).IntVal) != (0))) {
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_6.UnsafePtr).V1)))}
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v1_7.UnsafePtr).V1)))}
continue go__go_5_14_27
__t16 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_16
} else {

}
}
{
__t16 = false
}
end_branch_16:
return gopurs_runtime.Bool(__t16)
}
}()
})
})
return gopurs_runtime.Bool((gopurs_runtime.Apply2(go__go_5_14_27, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_3)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_4)))}).IntVal) != (0))
})
})}
_ = eqList_3_13
// TAST (Let): eqPattern1_2_11 -> *Constructor_Data_Eq_Eq
eqPattern1_2_11 := &Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(eqList_3_13.V0), x_4, y_5).IntVal) != (0))
})
})}
_ = eqPattern1_2_11
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqPattern1_2_11)}
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(ordList_1_0.V1), x_3, y_4).IntVal)), UnsafePtr: nil}
})
})})}
}

func Call_Data_List_Lazy_elemLastIndex(dictEq_0_loop *Constructor_Data_Eq_Eq, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_Data_List_Lazy_findLastIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), v_2, x_1).IntVal) != (0))
}))
}

func Call_Data_List_Lazy_elemIndex(dictEq_0_loop *Constructor_Data_Eq_Eq, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return Call_Data_List_Lazy_findIndex(gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), v_2, x_1).IntVal) != (0))
}))
}

func Call_Data_List_Lazy_dropWhile(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_28 gopurs_runtime.Value
go__go_1_0_28 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Data_List_Lazy_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](v_2_loop_val)
go__go_1_0_28:
for {
if false { continue go__go_1_0_28 }
var v_2 *Constructor_Data_List_Lazy_Types_Cons = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if ((v_2 != nil)) && ((gopurs_runtime.Apply(p_0, (v_2).V0).IntVal) != (0)) {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (v_2).V1))
continue go__go_1_0_28
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2)}
}))
}
end_branch_1:
return __t1
}
}()
})
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_1_0_28, gopurs_runtime.Apply(Get_Data_Lazy_force(), x_2))
})
}

func Call_Data_List_Lazy_drop(n_0_loop int64) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var go__go_1_0_29 gopurs_runtime.Value
go__go_1_0_29 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop int64 = v_2_loop_val.IntVal
var v1_3_loop *Constructor_Data_List_Lazy_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](v1_3_loop_val)
go__go_1_0_29:
for {
if false { continue go__go_1_0_29 }
var v_2 int64 = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Lazy_Types_Cons = v1_3_loop
_ = v1_3
var __t1 *Constructor_Data_List_Lazy_Types_Cons
{
if (v_2) == (0) {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
if (v1_3 == nil) {
__t1 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_1
} else {

}
}
{
if (v1_3 != nil) {
v_2_loop = (v_2) - (1)
v1_3_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (v1_3).V1))
continue go__go_1_0_29
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
})
// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
__local_var_2_4 := gopurs_runtime.Apply(go__go_1_0_29, gopurs_runtime.Int(n_0))
_ = __local_var_2_4
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Func(func(l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_4, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_3))
}))
})
_ = __local_var_2_3
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_3, x_3)
})
_ = __local_var_2_2
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_2, x_3)
})
}

func Call_Data_List_Lazy_slice(start_0_loop int64, end_1_loop int64, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var start_0 int64 = start_0_loop
_ = start_0
var end_1 int64 = end_1_loop
_ = end_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply(Call_Data_List_Lazy_take((end_1) - (start_0)), gopurs_runtime.Apply(Call_Data_List_Lazy_drop(start_0), xs_2))
}

func Call_Data_List_Lazy_deleteBy(eq_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
deleteBy:
for {
if false { continue deleteBy }
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)
_ = __local_var_4_0
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr == nil) {
__t2 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_2
} else {

}
}
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr != nil) {
var __t1 *Constructor_Data_List_Lazy_Types_Cons
{
if (gopurs_runtime.Apply2(eq_0, x_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V1))
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V0, Call_Data_List_Lazy_deleteBy(eq_0, x_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V1)}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
}))
}
}

func Call_Data_List_Lazy_unionBy(eq_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
var go__go_3_1_30 gopurs_runtime.Value
go__go_3_1_30 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var xs_5_loop gopurs_runtime.Value = xs_5_loop_val
go__go_3_1_30:
for {
if false { continue go__go_3_1_30 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_2 -> *Constructor_Data_List_Lazy_Types_Cons
v_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
_ = v_6_2
var __t3 gopurs_runtime.Value
{
if (v_6_2 == nil) {
__t3 = b_4
goto end_branch_3
} else {

}
}
{
if (v_6_2 != nil) {
b_4_loop = Call_Data_List_Lazy_deleteBy(eq_0, (v_6_2).V0, b_4)
xs_5_loop = (v_6_2).V1
continue go__go_3_1_30
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
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply2(go__go_3_1_30, Call_Data_List_Lazy_nubByEq(eq_0, ys_2), xs_1)
_ = __local_var_3_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)
_ = __local_var_5_4
var __t5 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_5_4.Type == 9 && __local_var_5_4.IntVal == 218341868 && __local_var_5_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_3_0))
goto end_branch_5
} else {

}
}
{
if (__local_var_5_4.Type == 9 && __local_var_5_4.IntVal == 218341868 && __local_var_5_4.UnsafePtr != nil) {
__t5 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_5_4.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_5_4.UnsafePtr).V1, __local_var_3_0)}
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t5)}
}))
}

func Call_Data_List_Lazy_union(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_List_Lazy_unionBy(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_List_Lazy_deleteAt(n_0_loop int64, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
deleteAt:
for {
if false { continue deleteAt }
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)
_ = __local_var_3_0
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t2 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_2
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
var __t1 *Constructor_Data_List_Lazy_Types_Cons
{
if (n_0) == (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V1))
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V0, Call_Data_List_Lazy_deleteAt((n_0) - (1), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V1)}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
}))
}
}

func Call_Data_List_Lazy_delete(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_List_Lazy_deleteBy(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_List_Lazy_difference(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var go__go_1_0_31 gopurs_runtime.Value
go__go_1_0_31 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_0_31:
for {
if false { continue go__go_1_0_31 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
// TAST (Let): v_4_1 -> *Constructor_Data_List_Lazy_Types_Cons
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_3))
_ = v_4_1
var __t2 gopurs_runtime.Value
{
if (v_4_1 == nil) {
__t2 = b_2
goto end_branch_2
} else {

}
}
{
if (v_4_1 != nil) {
b_2_loop = Call_Data_List_Lazy_deleteBy(gopurs_runtime.Box(dictEq_0.V0), (v_4_1).V0, b_2)
xs_3_loop = (v_4_1).V1
continue go__go_1_0_31
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
return go__go_1_0_31
}

func Call_Data_List_Lazy_cycle(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var go__go_1_0_32 gopurs_runtime.Value
_ = go__go_1_0_32
go__go_1_0_32 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_0)
_ = __local_var_4_1
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), go__go_1_0_32))
goto end_branch_2
} else {

}
}
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr != nil) {
__t2 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_1.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_1.UnsafePtr).V1, go__go_1_0_32)}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
})))
}))
return go__go_1_0_32
}

func Call_Data_List_Lazy_concatMap(b_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), a_1)
_ = __local_var_3_0
var __t8 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t8 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_8
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(b_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V0)
_ = __local_var_4_1
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V1, b_0)
_ = __local_var_5_2
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_3 -> gopurs_runtime.Value
__local_var_7_3 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_1)
_ = __local_var_7_3
var __t7 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_7_3.Type == 9 && __local_var_7_3.IntVal == 218341868 && __local_var_7_3.UnsafePtr == nil) {
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_2))
goto end_branch_7
} else {

}
}
{
if (__local_var_7_3.Type == 9 && __local_var_7_3.IntVal == 218341868 && __local_var_7_3.UnsafePtr != nil) {
// TAST (Let): __local_var_8_4 -> gopurs_runtime.Value
__local_var_8_4 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_7_3.UnsafePtr).V1
_ = __local_var_8_4
__t7 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_7_3.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_5 -> gopurs_runtime.Value
__local_var_10_5 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_8_4)
_ = __local_var_10_5
var __t6 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_10_5.Type == 9 && __local_var_10_5.IntVal == 218341868 && __local_var_10_5.UnsafePtr == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_2))
goto end_branch_6
} else {

}
}
{
if (__local_var_10_5.Type == 9 && __local_var_10_5.IntVal == 218341868 && __local_var_10_5.UnsafePtr != nil) {
__t6 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_10_5.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_10_5.UnsafePtr).V1, __local_var_5_2)}
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t6)}
}))}
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t7)}
}))))
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t8)}
}))
}

func Call_Data_List_Lazy_concat(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)
_ = __local_var_2_0
var __t8 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 218341868 && __local_var_2_0.UnsafePtr == nil) {
__t8 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_8
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 218341868 && __local_var_2_0.UnsafePtr != nil) {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_2_0.UnsafePtr).V0
_ = __local_var_3_1
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_2_0.UnsafePtr).V1, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return x_4
}))
_ = __local_var_4_2
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
__local_var_6_3 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_3_1)
_ = __local_var_6_3
var __t7 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_6_3.Type == 9 && __local_var_6_3.IntVal == 218341868 && __local_var_6_3.UnsafePtr == nil) {
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_2))
goto end_branch_7
} else {

}
}
{
if (__local_var_6_3.Type == 9 && __local_var_6_3.IntVal == 218341868 && __local_var_6_3.UnsafePtr != nil) {
// TAST (Let): __local_var_7_4 -> gopurs_runtime.Value
__local_var_7_4 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_6_3.UnsafePtr).V1
_ = __local_var_7_4
__t7 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_6_3.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_5 -> gopurs_runtime.Value
__local_var_9_5 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_7_4)
_ = __local_var_9_5
var __t6 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_9_5.Type == 9 && __local_var_9_5.IntVal == 218341868 && __local_var_9_5.UnsafePtr == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_2))
goto end_branch_6
} else {

}
}
{
if (__local_var_9_5.Type == 9 && __local_var_9_5.IntVal == 218341868 && __local_var_9_5.UnsafePtr != nil) {
__t6 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_9_5.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_9_5.UnsafePtr).V1, __local_var_4_2)}
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t6)}
}))}
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t7)}
}))))
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t8)}
}))
}

func Call_Data_List_Lazy_alterAt(n_0_loop int64, f_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
alterAt:
for {
if false { continue alterAt }
var n_0 int64 = n_0_loop
_ = n_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)
_ = __local_var_4_0
var __t4 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr == nil) {
__t4 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_4
} else {

}
}
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr != nil) {
var __t3 *Constructor_Data_List_Lazy_Types_Cons
{
if (n_0) == (0) {
// TAST (Let): v2_5_1 -> *Constructor_Data_Maybe_Just
v2_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V0))
_ = v2_5_1
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (v2_5_1 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5_1 != nil) {
__t2 = &Constructor_Data_List_Lazy_Types_Cons{1, (v2_5_1).V0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V1}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V0, Call_Data_List_Lazy_alterAt((n_0) - (1), f_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V1)}
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t4)}
}))
}
}

func Call_Data_List_Lazy_modifyAt(n_0_loop int64, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(Get_Data_List_Lazy_alterAt(), gopurs_runtime.Int(n_0), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_1, x_2)})}
}))
}

func Call_Data_List_Lazy_alterAt__950766476(n_0_loop int64, f_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)
_ = __local_var_4_0
var __t4 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr == nil) {
__t4 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_4
} else {

}
}
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr != nil) {
var __t3 *Constructor_Data_List_Lazy_Types_Cons
{
if (n_0) == (0) {
// TAST (Let): v2_5_1 -> *Constructor_Data_Maybe_Just
v2_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V0))
_ = v2_5_1
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (v2_5_1 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5_1 != nil) {
__t2 = &Constructor_Data_List_Lazy_Types_Cons{1, (v2_5_1).V0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V1}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V0, Call_Data_List_Lazy_alterAt((n_0) - (1), f_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V1)}
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t4)}
}))
}

func Call_Data_List_Lazy_deleteAt__4024047148(n_0_loop int64, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)
_ = __local_var_3_0
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t2 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_2
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
var __t1 *Constructor_Data_List_Lazy_Types_Cons
{
if (n_0) == (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V1))
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V0, Call_Data_List_Lazy_deleteAt((n_0) - (1), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V1)}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
}))
}

func Call_Data_List_Lazy_deleteBy__501100275(eq_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)
_ = __local_var_4_0
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr == nil) {
__t2 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_2
} else {

}
}
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr != nil) {
var __t1 *Constructor_Data_List_Lazy_Types_Cons
{
if (gopurs_runtime.Apply2(eq_0, x_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V1))
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V0, Call_Data_List_Lazy_deleteBy(eq_0, x_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V1)}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
}))
}

func Call_Data_List_Lazy_drop__4024047148(n_0_loop int64) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var go__go_1_0_33 gopurs_runtime.Value
go__go_1_0_33 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop int64 = v_2_loop_val.IntVal
var v1_3_loop *Constructor_Data_List_Lazy_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](v1_3_loop_val)
go__go_1_0_33:
for {
if false { continue go__go_1_0_33 }
var v_2 int64 = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Lazy_Types_Cons = v1_3_loop
_ = v1_3
var __t1 *Constructor_Data_List_Lazy_Types_Cons
{
if (v_2) == (0) {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
if (v1_3 == nil) {
__t1 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_1
} else {

}
}
{
if (v1_3 != nil) {
v_2_loop = (v_2) - (1)
v1_3_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (v1_3).V1))
continue go__go_1_0_33
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
})
// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
__local_var_2_4 := gopurs_runtime.Apply(go__go_1_0_33, gopurs_runtime.Int(n_0))
_ = __local_var_2_4
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Func(func(l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_4, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_3))
}))
})
_ = __local_var_2_3
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_3, x_3)
})
_ = __local_var_2_2
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_2, x_3)
})
}

func Call_Data_List_Lazy_filter__638755635(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var go__go_1_0_34 gopurs_runtime.Value
go__go_1_0_34 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_34:
for {
if false { continue go__go_1_0_34 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t2 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_2
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
var __t1 *Constructor_Data_List_Lazy_Types_Cons
{
if (gopurs_runtime.Apply(p_0, (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V0).IntVal) != (0) {
__t1 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V0, gopurs_runtime.Apply(Call_Data_List_Lazy_filter(p_0), (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V1)}
goto end_branch_1
} else {

}
}
{
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V1)))}
continue go__go_1_0_34
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Value{})
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
}
}()
})
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_1_0_34, gopurs_runtime.Apply(Get_Data_Lazy_force(), x_2))
}))
})
}

func Call_Data_List_Lazy_filterM__647926151(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(p_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(list_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_5_2 -> *Constructor_Data_Maybe_Just
v_5_2 := Call_Data_List_Lazy_uncons(list_4)
_ = v_5_2
var __t13 gopurs_runtime.Value
{
if (v_5_2 == nil) {
__t13 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), Get_Data_List_Lazy_Types_nil())
goto end_branch_13
} else {

}
}
{
if (v_5_2 != nil) {
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
__local_var_6_3 := gopurs_runtime.RecordGet((v_5_2).V0, "head")
_ = __local_var_6_3
// TAST (Let): __local_var_7_4 -> gopurs_runtime.Value
__local_var_7_4 := gopurs_runtime.RecordGet((v_5_2).V0, "tail")
_ = __local_var_7_4
__t13 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(p_3, __local_var_6_3), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Applicative0_9_5 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_9_5
// TAST (Let): Bind1_10_6 -> *Constructor_Control_Bind_Bind
Bind1_10_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_10_6
// TAST (Let): v_11_7 -> *Constructor_Data_Maybe_Just
v_11_7 := Call_Data_List_Lazy_uncons(__local_var_7_4)
_ = v_11_7
var __t11 gopurs_runtime.Value
{
if (v_11_7 == nil) {
__t11 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_5.V1), Get_Data_List_Lazy_Types_nil())
goto end_branch_11
} else {

}
}
{
if (v_11_7 != nil) {
// TAST (Let): __local_var_12_8 -> gopurs_runtime.Value
__local_var_12_8 := gopurs_runtime.RecordGet((v_11_7).V0, "head")
_ = __local_var_12_8
// TAST (Let): __local_var_13_9 -> gopurs_runtime.Value
__local_var_13_9 := gopurs_runtime.RecordGet((v_11_7).V0, "tail")
_ = __local_var_13_9
__t11 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_6.V1), gopurs_runtime.Apply(p_3, __local_var_12_8), gopurs_runtime.Func(func(b_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_6.V1), gopurs_runtime.Apply2(Call_Data_List_Lazy_filterM(dictMonad_0), p_3, __local_var_13_9), gopurs_runtime.Func(func(xs_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
if (b_14.IntVal) != (0) {
__t10 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_12_8, xs_prime_15})}
}))
goto end_branch_10
} else {

}
}
{
__t10 = xs_prime_15
}
end_branch_10:
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_5.V1), __t10)
}))
}))
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), __t11, gopurs_runtime.Func(func(xs_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
if (b_8.IntVal) != (0) {
__t12 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_6_3, xs_prime_9})}
}))
goto end_branch_12
} else {

}
}
{
__t12 = xs_prime_9
}
end_branch_12:
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), __t12)
}))
}))
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
})
})
}

func Call_Data_List_Lazy_findIndex__1594900290(fn_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var go__go_1_0_35 gopurs_runtime.Value
go__go_1_0_35 = gopurs_runtime.Func(func(n_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(list_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var n_2_loop int64 = n_2_loop_val.IntVal
var list_3_loop gopurs_runtime.Value = list_3_loop_val
go__go_1_0_35:
for {
if false { continue go__go_1_0_35 }
var n_2 int64 = n_2_loop
_ = n_2
var list_3 gopurs_runtime.Value = list_3_loop
_ = list_3
// TAST (Let): __local_var_4_1 -> *Constructor_Data_Maybe_Just
__local_var_4_1 := Call_Data_List_Lazy_uncons(list_3)
_ = __local_var_4_1
var __t3 *Constructor_Data_Maybe_Just
{
if (__local_var_4_1 != nil) {
var __t2 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.Apply(fn_0, gopurs_runtime.RecordGet((__local_var_4_1).V0, "head")).IntVal) != (0) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(n_2)})})
goto end_branch_2
} else {

}
}
{
n_2_loop = (n_2) + (1)
list_3_loop = gopurs_runtime.RecordGet((__local_var_4_1).V0, "tail")
continue go__go_1_0_35
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{})
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
if (__local_var_4_1 == nil) {
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
return gopurs_runtime.Apply(go__go_1_0_35, gopurs_runtime.Int(0))
}

func Call_Data_List_Lazy_findLastIndex__1594900290(fn_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
// TAST (Let): __local_var_2_0 -> *Constructor_Data_Maybe_Just
__local_var_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Call_Data_List_Lazy_findIndex(fn_0), Call_Data_List_Lazy_reverse(xs_1)))
_ = __local_var_2_0
var __t1 *Constructor_Data_Maybe_Just
{
if (__local_var_2_0 != nil) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(((gopurs_runtime.Apply(Get_Data_List_Lazy_length(), xs_1).IntVal) - (1)) - ((__local_var_2_0).V0.IntVal))}
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

func Call_Data_List_Lazy_foldM__3505933597(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_6_2 -> *Constructor_Data_Maybe_Just
v_6_2 := Call_Data_List_Lazy_uncons(xs_5)
_ = v_6_2
var __t9 gopurs_runtime.Value
{
if (v_6_2 == nil) {
__t9 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), b_4)
goto end_branch_9
} else {

}
}
{
if (v_6_2 != nil) {
// TAST (Let): __local_var_7_3 -> gopurs_runtime.Value
__local_var_7_3 := gopurs_runtime.RecordGet((v_6_2).V0, "tail")
_ = __local_var_7_3
__t9 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(f_3, b_4, gopurs_runtime.RecordGet((v_6_2).V0, "head")), gopurs_runtime.Func(func(b_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Applicative0_9_4 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_9_4
// TAST (Let): Bind1_10_5 -> *Constructor_Control_Bind_Bind
Bind1_10_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_10_5
// TAST (Let): v_11_6 -> *Constructor_Data_Maybe_Just
v_11_6 := Call_Data_List_Lazy_uncons(__local_var_7_3)
_ = v_11_6
var __t8 gopurs_runtime.Value
{
if (v_11_6 == nil) {
__t8 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_4.V1), b_prime_8)
goto end_branch_8
} else {

}
}
{
if (v_11_6 != nil) {
// TAST (Let): __local_var_12_7 -> gopurs_runtime.Value
__local_var_12_7 := gopurs_runtime.RecordGet((v_11_6).V0, "tail")
_ = __local_var_12_7
__t8 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_5.V1), gopurs_runtime.Apply2(f_3, b_prime_8, gopurs_runtime.RecordGet((v_11_6).V0, "head")), gopurs_runtime.Func(func(b_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Call_Data_List_Lazy_foldM(dictMonad_0), f_3, b_prime_13, __local_var_12_7)
}))
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
}))
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
})
})
})
}

func Call_Data_List_Lazy_fromFoldable__4212258679(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V2), Get_Data_List_Lazy_Types_cons(), Get_Data_List_Lazy_Types_nil())
}

func Call_Data_List_Lazy_fromStep__1398792641(x_0_loop *Constructor_Data_List_Lazy_Types_Cons) gopurs_runtime.Value {
var x_0 *Constructor_Data_List_Lazy_Types_Cons = x_0_loop
_ = x_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(x_0)}
}))
}

func Call_Data_List_Lazy_groupBy__1659362014(eq_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), x_1)
_ = __local_var_3_0
var __t4 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t4 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_4
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V0
_ = __local_var_4_1
// TAST (Let): v1_5_2 -> gopurs_runtime.Value
var v1_5_2 gopurs_runtime.Value = Call_Data_List_Lazy_span(gopurs_runtime.Apply(eq_0, __local_var_4_1), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V1)
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
__local_var_6_3 := gopurs_runtime.RecordGet(v1_5_2, "init")
_ = __local_var_6_3
__t4 = &Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, __local_var_4_1, __local_var_6_3})}
})), Call_Data_List_Lazy_groupBy(eq_0, gopurs_runtime.RecordGet(v1_5_2, "rest"))}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t4)}
}))
}

func Call_Data_List_Lazy_head__2155426095(xs_0_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
// TAST (Let): __local_var_1_0 -> *Constructor_Data_Maybe_Just
__local_var_1_0 := Call_Data_List_Lazy_uncons(xs_0)
_ = __local_var_1_0
var __t1 *Constructor_Data_Maybe_Just
{
if (__local_var_1_0 != nil) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordGet((__local_var_1_0).V0, "head")}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
return __t1
}

func Call_Data_List_Lazy_insertAt__725610501(v_0_loop int64, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t2 gopurs_runtime.Value
{
if (v_0) == (0) {
__t2 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, v1_1, v2_2})}
}))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v2_2)
_ = __local_var_4_0
var __t1 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr == nil) {
__t1 = &Constructor_Data_List_Lazy_Types_Cons{1, v1_1, Get_Data_List_Lazy_Types_nil()}
goto end_branch_1
} else {

}
}
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr != nil) {
__t1 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V0, Call_Data_List_Lazy_insertAt((v_0) - (1), v1_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V1)}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t1)}
}))
}
end_branch_2:
return __t2
}

func Call_Data_List_Lazy_insertBy__2098566601(cmp_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)
_ = __local_var_4_0
var __t3 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr == nil) {
__t3 = &Constructor_Data_List_Lazy_Types_Cons{1, x_1, Get_Data_List_Lazy_Types_nil()}
goto end_branch_3
} else {

}
}
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr != nil) {
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(cmp_0, x_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V0)
if (uint32(__t_tag_1.IntVal) == 380165415) {
__t2 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V0, Call_Data_List_Lazy_insertBy(cmp_0, x_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V1)}
goto end_branch_2
} else {

}
}
{
__t2 = &Constructor_Data_List_Lazy_Types_Cons{1, x_1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](__local_var_4_0))}
}))}
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t3)}
}))
}

func Call_Data_List_Lazy_intersectBy__3844889126(eq_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
return gopurs_runtime.Apply(Call_Data_List_Lazy_filter(gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupDisj1_4_0 -> *Constructor_Data_Semigroup_Semigroup
semigroupDisj1_4_0 := &Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((v_4.IntVal) != (0)) || ((v1_5.IntVal) != (0)))
})
})}
_ = semigroupDisj1_4_0
// TAST (Let): __local_var_5_1 -> *Constructor_Data_Monoid_Monoid
__local_var_5_1 := &Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupDisj1_4_0)}
}), gopurs_runtime.Bool(false)}
_ = __local_var_5_1
// TAST (Let): Semigroup0_6_2 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_5_1.V0), gopurs_runtime.Value{}))
_ = Semigroup0_6_2
// TAST (Let): __local_var_7_3 -> gopurs_runtime.Value
__local_var_7_3 := gopurs_runtime.Apply(eq_0, x_3)
_ = __local_var_7_3
var go__go_8_4_37 gopurs_runtime.Value
go__go_8_4_37 = gopurs_runtime.Func(func(b_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_9_loop gopurs_runtime.Value = b_9_loop_val
var xs_10_loop gopurs_runtime.Value = xs_10_loop_val
go__go_8_4_37:
for {
if false { continue go__go_8_4_37 }
var b_9 gopurs_runtime.Value = b_9_loop
_ = b_9
var xs_10 gopurs_runtime.Value = xs_10_loop
_ = xs_10
// TAST (Let): v_11_5 -> *Constructor_Data_List_Lazy_Types_Cons
v_11_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_10))
_ = v_11_5
var __t6 gopurs_runtime.Value
{
if (v_11_5 == nil) {
__t6 = b_9
goto end_branch_6
} else {

}
}
{
if (v_11_5 != nil) {
b_9_loop = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_2.V0), b_9, gopurs_runtime.Apply(__local_var_7_3, (v_11_5).V0))
xs_10_loop = (v_11_5).V1
continue go__go_8_4_37
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
return gopurs_runtime.Bool((gopurs_runtime.Apply2(go__go_8_4_37, gopurs_runtime.Box(__local_var_5_1.V1), ys_2).IntVal) != (0))
})), xs_1)
}

func Call_Data_List_Lazy_iterate__455058292(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var go__go_2_0_38 gopurs_runtime.Value
_ = go__go_2_0_38
go__go_2_0_38 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.Apply(Get_Data_Lazy_force(), go__go_2_0_38)
_ = __local_var_5_2
var __t3 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_5_2.Type == 9 && __local_var_5_2.IntVal == 218341868 && __local_var_5_2.UnsafePtr == nil) {
__t3 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_3
} else {

}
}
{
if (__local_var_5_2.Type == 9 && __local_var_5_2.IntVal == 218341868 && __local_var_5_2.UnsafePtr != nil) {
__t3 = &Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_5_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_List_Lazy_Types_functorList()).V0), f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_5_2.UnsafePtr).V1)}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t3)}
}))
_ = __local_var_4_1
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, x_1, __local_var_4_1})}
})))
}))
return go__go_2_0_38
}

func Call_Data_List_Lazy_many__956417025(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative) gopurs_runtime.Value {
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
return gopurs_runtime.Apply2(gopurs_runtime.Box(Alt0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V0), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V1), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_List_Lazy_Types_cons(), v_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_3, "defer"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V1), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "alt"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V0), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V1), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_List_Lazy_Types_cons(), v_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_3, "defer"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Data_List_Lazy_many(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy](dictLazy_3))}, v_4)
}))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V0), gopurs_runtime.Value{}), "pure"), Get_Data_List_Lazy_Types_nil()))
}))), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), Get_Data_List_Lazy_Types_nil()))
})
})
}

func Call_Data_List_Lazy_mapMaybe__3574309085(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_41 gopurs_runtime.Value
go__go_1_0_41 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_41:
for {
if false { continue go__go_1_0_41 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t3 *Constructor_Data_List_Lazy_Types_Cons
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t3 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
// TAST (Let): v1_3_1 -> *Constructor_Data_Maybe_Just
v1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V0))
_ = v1_3_1
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (v1_3_1 == nil) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V1)))}
continue go__go_1_0_41
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
if (v1_3_1 != nil) {
__t2 = &Constructor_Data_List_Lazy_Types_Cons{1, (v1_3_1).V0, gopurs_runtime.Apply(Call_Data_List_Lazy_mapMaybe(f_0), (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V1)}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_1_0_41, gopurs_runtime.Apply(Get_Data_Lazy_force(), x_2))
}))
})
}

func Call_Data_List_Lazy_mapMaybe__2519317725(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_42 gopurs_runtime.Value
go__go_1_0_42 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_42:
for {
if false { continue go__go_1_0_42 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t3 *Constructor_Data_List_Lazy_Types_Cons
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t3 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
// TAST (Let): v1_3_1 -> *Constructor_Data_Maybe_Just
v1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V0))
_ = v1_3_1
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (v1_3_1 == nil) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V1)))}
continue go__go_1_0_42
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
if (v1_3_1 != nil) {
__t2 = &Constructor_Data_List_Lazy_Types_Cons{1, (v1_3_1).V0, gopurs_runtime.Apply(Call_Data_List_Lazy_mapMaybe(f_0), (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V1)}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_1_0_42, gopurs_runtime.Apply(Get_Data_Lazy_force(), x_2))
}))
})
}

func Call_Data_List_Lazy_mapMaybe__1687744733(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_43 gopurs_runtime.Value
go__go_1_0_43 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_43:
for {
if false { continue go__go_1_0_43 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t3 *Constructor_Data_List_Lazy_Types_Cons
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t3 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
// TAST (Let): v1_3_1 -> *Constructor_Data_Maybe_Just
v1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V0))
_ = v1_3_1
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (v1_3_1 == nil) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V1)))}
continue go__go_1_0_43
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
if (v1_3_1 != nil) {
__t2 = &Constructor_Data_List_Lazy_Types_Cons{1, (v1_3_1).V0, gopurs_runtime.Apply(Call_Data_List_Lazy_mapMaybe(f_0), (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V1)}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_1_0_43, gopurs_runtime.Apply(Get_Data_Lazy_force(), x_2))
}))
})
}

func Call_Data_List_Lazy_mapMaybe__899591645(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_44 gopurs_runtime.Value
go__go_1_0_44 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_44:
for {
if false { continue go__go_1_0_44 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t3 *Constructor_Data_List_Lazy_Types_Cons
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t3 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
// TAST (Let): v1_3_1 -> *Constructor_Data_Maybe_Just
v1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V0))
_ = v1_3_1
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (v1_3_1 == nil) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V1)))}
continue go__go_1_0_44
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
if (v1_3_1 != nil) {
__t2 = &Constructor_Data_List_Lazy_Types_Cons{1, (v1_3_1).V0, gopurs_runtime.Apply(Call_Data_List_Lazy_mapMaybe(f_0), (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V1)}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_1_0_44, gopurs_runtime.Apply(Get_Data_Lazy_force(), x_2))
}))
})
}

func Call_Data_List_Lazy_mapMaybe__600226685(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_45 gopurs_runtime.Value
go__go_1_0_45 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_45:
for {
if false { continue go__go_1_0_45 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t3 *Constructor_Data_List_Lazy_Types_Cons
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t3 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
// TAST (Let): v1_3_1 -> *Constructor_Data_Maybe_Just
v1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V0))}))
_ = v1_3_1
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (v1_3_1 == nil) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V1)))}
continue go__go_1_0_45
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
if (v1_3_1 != nil) {
__t2 = &Constructor_Data_List_Lazy_Types_Cons{1, (v1_3_1).V0, gopurs_runtime.Apply(Call_Data_List_Lazy_mapMaybe(f_0), (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V1)}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_1_0_45, gopurs_runtime.Apply(Get_Data_Lazy_force(), x_2))
}))
})
}

func Call_Data_List_Lazy_nubBy__2220739616(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var goStep_1_0_46 gopurs_runtime.Value
_ = goStep_1_0_46
var go__go_1_1_47 gopurs_runtime.Value
_ = go__go_1_1_47
goStep_1_0_46 = gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_List_Lazy_Types_Cons
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr == nil) {
__t4 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_4
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 218341868 && v1_3.UnsafePtr != nil) {
// TAST (Let): v2_4_2 -> gopurs_runtime.Value
v2_4_2 := gopurs_runtime.Apply3(Get_Data_List_Internal_insertAndLookupBy(), p_0, (*Constructor_Data_List_Lazy_Types_Cons)(v1_3.UnsafePtr).V0, v_2)
_ = v2_4_2
var __t3 *Constructor_Data_List_Lazy_Types_Cons
{
if (gopurs_runtime.RecordGet(v2_4_2, "found").IntVal) != (0) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply2(go__go_1_1_47, gopurs_runtime.RecordGet(v2_4_2, "result"), (*Constructor_Data_List_Lazy_Types_Cons)(v1_3.UnsafePtr).V1)))
goto end_branch_3
} else {

}
}
{
__t3 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(v1_3.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_1_1_47, gopurs_runtime.RecordGet(v2_4_2, "result"), (*Constructor_Data_List_Lazy_Types_Cons)(v1_3.UnsafePtr).V1)}
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t4)}
})
})
go__go_1_1_47 = gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_5 -> gopurs_runtime.Value
__local_var_4_5 := gopurs_runtime.Apply(goStep_1_0_46, s_2)
_ = __local_var_4_5
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_5, gopurs_runtime.Apply(Get_Data_Lazy_force(), v_3))
}))
})
})
return gopurs_runtime.Apply(go__go_1_1_47, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)})
}

func Call_Data_List_Lazy_nubByEq__616397370(eq_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), x_1)
_ = __local_var_3_0
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t2 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_2
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V0
_ = __local_var_4_1
__t2 = &Constructor_Data_List_Lazy_Types_Cons{1, __local_var_4_1, Call_Data_List_Lazy_nubByEq(eq_0, gopurs_runtime.Apply(Call_Data_List_Lazy_filter(gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(eq_0, __local_var_4_1, y_5).IntVal) != (0)) != (true))
})), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V1))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
}))
}

func Call_Data_List_Lazy_null__1674339719(x_0_loop gopurs_runtime.Value) bool {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
var __local_var_1_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_List_Lazy_uncons(x_0))}
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136 && __local_var_1_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136 && __local_var_1_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return (__t1.IntVal) != (0)
}

func Call_Data_List_Lazy_repeat__2462085934(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var go__go_1_0_48 gopurs_runtime.Value
_ = go__go_1_0_48
go__go_1_0_48 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, x_0, go__go_1_0_48})}
})))
}))
return go__go_1_0_48
}

func Call_Data_List_Lazy_repeat__2149902581(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var go__go_1_0_49 gopurs_runtime.Value
_ = go__go_1_0_49
go__go_1_0_49 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, x_0, go__go_1_0_49})}
})))
}))
return go__go_1_0_49
}

func Call_Data_List_Lazy_replicateM__3816548429(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(n_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
var __t7 bool
{
if (n_3.IntVal) < (1) {
__t7 = true
goto end_branch_7
} else {

}
}
{
__t7 = false
}
end_branch_7:
if __t7 {
__t8 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), Get_Data_List_Lazy_Types_nil())
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Applicative0_6_2 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_6_2
// TAST (Let): Bind1_7_3 -> *Constructor_Control_Bind_Bind
Bind1_7_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_7_3
// TAST (Let): __local_var_8_4 -> int64
__local_var_8_4 := (n_3.IntVal) - (1)
_ = __local_var_8_4
var __t6 gopurs_runtime.Value
{
var __t5 bool
{
if (__local_var_8_4) < (1) {
__t5 = true
goto end_branch_5
} else {

}
}
{
__t5 = false
}
end_branch_5:
if __t5 {
__t6 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_2.V1), Get_Data_List_Lazy_Types_nil())
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_3.V1), m_4, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_3.V1), gopurs_runtime.Apply2(Call_Data_List_Lazy_replicateM(dictMonad_0), gopurs_runtime.Int((__local_var_8_4) - (1)), m_4), gopurs_runtime.Func(func(as_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_2.V1), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, a_9, as_10})}
})))
}))
}))
}
end_branch_6:
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), __t6, gopurs_runtime.Func(func(as_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, a_5, as_6})}
})))
}))
}))
}
end_branch_8:
return __t8
})
})
}

func Call_Data_List_Lazy_reverse__1315655552(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_50 gopurs_runtime.Value
go__go_2_0_50 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_3_loop gopurs_runtime.Value = b_3_loop_val
var xs_4_loop gopurs_runtime.Value = xs_4_loop_val
go__go_2_0_50:
for {
if false { continue go__go_2_0_50 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var xs_4 gopurs_runtime.Value = xs_4_loop
_ = xs_4
// TAST (Let): v_5_1 -> *Constructor_Data_List_Lazy_Types_Cons
v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_4))
_ = v_5_1
var __t3 gopurs_runtime.Value
{
if (v_5_1 == nil) {
__t3 = b_3
goto end_branch_3
} else {

}
}
{
if (v_5_1 != nil) {
// TAST (Let): __local_var_6_2 -> gopurs_runtime.Value
__local_var_6_2 := (v_5_1).V0
_ = __local_var_6_2
b_3_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_6_2, b_3})}
}))
xs_4_loop = (v_5_1).V1
continue go__go_2_0_50
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
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply2(go__go_2_0_50, Get_Data_List_Lazy_Types_nil(), xs_0))
}))
}

func Call_Data_List_Lazy_scanlLazy__3747838620(f_0_loop gopurs_runtime.Value, acc_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var acc_1 gopurs_runtime.Value = acc_1_loop
_ = acc_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)
_ = __local_var_4_0
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr == nil) {
__t2 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_2
} else {

}
}
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr != nil) {
// TAST (Let): acc_prime_5_1 -> gopurs_runtime.Value
acc_prime_5_1 := gopurs_runtime.Apply2(f_0, acc_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V0)
_ = acc_prime_5_1
__t2 = &Constructor_Data_List_Lazy_Types_Cons{1, acc_prime_5_1, Call_Data_List_Lazy_scanlLazy(f_0, acc_prime_5_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V1)}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
}))
}

func Call_Data_List_Lazy_some__956417025(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative) gopurs_runtime.Value {
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
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_List_Lazy_Types_cons(), v_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_3, "defer"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V1), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "alt"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V0), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V1), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_List_Lazy_Types_cons(), v_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_3, "defer"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Data_List_Lazy_many(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy](dictLazy_3))}, v_4)
}))), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V0), gopurs_runtime.Value{}), "pure"), Get_Data_List_Lazy_Types_nil()))
})))
})
})
}

func Call_Data_List_Lazy_span__776304907(p_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
// TAST (Let): v_2_0 -> *Constructor_Data_Maybe_Just
v_2_0 := Call_Data_List_Lazy_uncons(xs_1)
_ = v_2_0
var __t2 gopurs_runtime.Value
{
if ((v_2_0 != nil)) && ((gopurs_runtime.Apply(p_0, gopurs_runtime.RecordGet((v_2_0).V0, "head")).IntVal) != (0)) {
// TAST (Let): v1_3_1 -> gopurs_runtime.Value
var v1_3_1 gopurs_runtime.Value = Call_Data_List_Lazy_span(p_0, gopurs_runtime.RecordGet((v_2_0).V0, "tail"))
__t2 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.RecordGet((v_2_0).V0, "head"), gopurs_runtime.RecordGet(v1_3_1, "init")})}
})), gopurs_runtime.RecordGet(v1_3_1, "rest"))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict2("init", "rest", Get_Data_List_Lazy_Types_nil(), xs_1)
}
end_branch_2:
return __t2
}

func Call_Data_List_Lazy_tail__1935051898(xs_0_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
// TAST (Let): __local_var_1_0 -> *Constructor_Data_Maybe_Just
__local_var_1_0 := Call_Data_List_Lazy_uncons(xs_0)
_ = __local_var_1_0
var __t1 *Constructor_Data_Maybe_Just
{
if (__local_var_1_0 != nil) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordGet((__local_var_1_0).V0, "tail")}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_1:
return __t1
}

func Call_Data_List_Lazy_take__4024047148(n_0_loop int64) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var __t3 gopurs_runtime.Value
{
var __t2 bool
{
if (n_0) > (0) {
__t2 = false
goto end_branch_2
} else {

}
}
{
__t2 = true
}
end_branch_2:
if __t2 {
__t3 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_nil()
})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), x_1)
_ = __local_var_3_0
var __t1 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t1 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_1
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
__t1 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V0, gopurs_runtime.Apply(Call_Data_List_Lazy_take((n_0) - (1)), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V1)}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t1)}
}))
})
}
end_branch_3:
return __t3
}

func Call_Data_List_Lazy_takeWhile__638755635(p_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), x_1)
_ = __local_var_3_0
var __t1 *Constructor_Data_List_Lazy_Types_Cons
{
if ((__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil)) && ((gopurs_runtime.Apply(p_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V0).IntVal) != (0)) {
__t1 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V0, Call_Data_List_Lazy_takeWhile(p_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_3_0.UnsafePtr).V1)}
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t1)}
}))
}

func Call_Data_List_Lazy_transpose__1534541312(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
// TAST (Let): v_1_0 -> *Constructor_Data_Maybe_Just
v_1_0 := Call_Data_List_Lazy_uncons(xs_0)
_ = v_1_0
var __t7 gopurs_runtime.Value
{
if (v_1_0 == nil) {
__t7 = xs_0
goto end_branch_7
} else {

}
}
{
if (v_1_0 != nil) {
// TAST (Let): v1_2_1 -> *Constructor_Data_Maybe_Just
v1_2_1 := Call_Data_List_Lazy_uncons(gopurs_runtime.RecordGet((v_1_0).V0, "head"))
_ = v1_2_1
var __t6 gopurs_runtime.Value
{
if (v1_2_1 == nil) {
__t6 = Call_Data_List_Lazy_transpose(gopurs_runtime.RecordGet((v_1_0).V0, "tail"))
goto end_branch_6
} else {

}
}
{
if (v1_2_1 != nil) {
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(Call_Data_List_Lazy_mapMaybe(Get_Data_List_Lazy_head()), gopurs_runtime.RecordGet((v_1_0).V0, "tail"))
_ = __local_var_3_3
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.RecordGet((v1_2_1).V0, "head"), __local_var_3_3})}
}))
_ = __local_var_3_2
// TAST (Let): __local_var_4_5 -> gopurs_runtime.Value
__local_var_4_5 := gopurs_runtime.Apply(Call_Data_List_Lazy_mapMaybe(Get_Data_List_Lazy_tail()), gopurs_runtime.RecordGet((v_1_0).V0, "tail"))
_ = __local_var_4_5
// TAST (Let): __local_var_4_4 -> gopurs_runtime.Value
__local_var_4_4 := Call_Data_List_Lazy_transpose(gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.RecordGet((v1_2_1).V0, "tail"), __local_var_4_5})}
})))
_ = __local_var_4_4
__t6 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Lazy_Types_Cons{1, __local_var_3_2, __local_var_4_4})}
}))
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
__t7 = __t6
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

func Call_Data_List_Lazy_uncons__3647012005(xs_0_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
// TAST (Let): v_1_0 -> *Constructor_Data_List_Lazy_Types_Cons
v_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_0))
_ = v_1_0
var __t1 *Constructor_Data_Maybe_Just
{
if (v_1_0 == nil) {
__t1 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_1
} else {

}
}
{
if (v_1_0 != nil) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("head", "tail", (v_1_0).V0, (v_1_0).V1)}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)})
}

func Call_Data_List_Lazy_uncons__1321764894(xs_0_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
// TAST (Let): v_1_0 -> *Constructor_Data_List_Lazy_Types_Cons
v_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_0))
_ = v_1_0
var __t1 *Constructor_Data_Maybe_Just
{
if (v_1_0 == nil) {
__t1 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_1
} else {

}
}
{
if (v_1_0 != nil) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("head", "tail", (v_1_0).V0, (v_1_0).V1)}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)})
}

func Call_Data_List_Lazy_uncons__974566859(xs_0_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
// TAST (Let): v_1_0 -> *Constructor_Data_List_Lazy_Types_Cons
v_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_0))
_ = v_1_0
var __t1 *Constructor_Data_Maybe_Just
{
if (v_1_0 == nil) {
__t1 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_1
} else {

}
}
{
if (v_1_0 != nil) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("head", "tail", (v_1_0).V0, (v_1_0).V1)}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)})
}

func Call_Data_List_Lazy_uncons__1420258522(xs_0_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
// TAST (Let): v_1_0 -> *Constructor_Data_List_Lazy_Types_Cons
v_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_0))
_ = v_1_0
var __t1 *Constructor_Data_Maybe_Just
{
if (v_1_0 == nil) {
__t1 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_1
} else {

}
}
{
if (v_1_0 != nil) {
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("head", "tail", (v_1_0).V0, (v_1_0).V1)}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)})
}

func Call_Data_List_Lazy_unionBy__3844889126(eq_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
var go__go_3_1_51 gopurs_runtime.Value
go__go_3_1_51 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var xs_5_loop gopurs_runtime.Value = xs_5_loop_val
go__go_3_1_51:
for {
if false { continue go__go_3_1_51 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_2 -> *Constructor_Data_List_Lazy_Types_Cons
v_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
_ = v_6_2
var __t3 gopurs_runtime.Value
{
if (v_6_2 == nil) {
__t3 = b_4
goto end_branch_3
} else {

}
}
{
if (v_6_2 != nil) {
b_4_loop = Call_Data_List_Lazy_deleteBy(eq_0, (v_6_2).V0, b_4)
xs_5_loop = (v_6_2).V1
continue go__go_3_1_51
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
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply2(go__go_3_1_51, Call_Data_List_Lazy_nubByEq(eq_0, ys_2), xs_1)
_ = __local_var_3_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)
_ = __local_var_5_4
var __t5 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_5_4.Type == 9 && __local_var_5_4.IntVal == 218341868 && __local_var_5_4.UnsafePtr == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_3_0))
goto end_branch_5
} else {

}
}
{
if (__local_var_5_4.Type == 9 && __local_var_5_4.IntVal == 218341868 && __local_var_5_4.UnsafePtr != nil) {
__t5 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_5_4.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_List_Lazy_Types_semigroupList()).V0), (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_5_4.UnsafePtr).V1, __local_var_3_0)}
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t5)}
}))
}

func Call_Data_List_Lazy_updateAt__725610501(n_0_loop int64, x_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)
_ = __local_var_4_0
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr == nil) {
__t2 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_2
} else {

}
}
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr != nil) {
var __t1 *Constructor_Data_List_Lazy_Types_Cons
{
if (n_0) == (0) {
__t1 = &Constructor_Data_List_Lazy_Types_Cons{1, x_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V1}
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_List_Lazy_Types_Cons{1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V0, Call_Data_List_Lazy_updateAt((n_0) - (1), x_1, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_0.UnsafePtr).V1)}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
}))
}

func Call_Data_List_Lazy_zipWith__3539178005(f_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)
_ = __local_var_4_1
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr == nil) {
__t2 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_2
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr == nil) {
__t2 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_2
} else {

}
}
{
if ((__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr != nil)) {
__t2 = &Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Apply2(f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_1.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons)(v1_5.UnsafePtr).V0), Call_Data_List_Lazy_zipWith(f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_1.UnsafePtr).V1, (*Constructor_Data_List_Lazy_Types_Cons)(v1_5.UnsafePtr).V1)}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
})
}))
_ = __local_var_3_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Lazy_force(), __local_var_3_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_2))
}))
}

func Call_Data_List_Lazy_zipWith__3210333397(f_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)
_ = __local_var_4_1
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr == nil) {
__t2 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_2
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr == nil) {
__t2 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_2
} else {

}
}
{
if ((__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr != nil)) {
__t2 = &Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_1.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons)(v1_5.UnsafePtr).V0)))}, Call_Data_List_Lazy_zipWith(f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_1.UnsafePtr).V1, (*Constructor_Data_List_Lazy_Types_Cons)(v1_5.UnsafePtr).V1)}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
})
}))
_ = __local_var_3_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Lazy_force(), __local_var_3_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_2))
}))
}

func Call_Data_List_Lazy_zipWith__3984071349(f_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)
_ = __local_var_4_1
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr == nil) {
__t2 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_2
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr == nil) {
__t2 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_2
} else {

}
}
{
if ((__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr != nil)) {
__t2 = &Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Apply2(f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_1.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons)(v1_5.UnsafePtr).V0), Call_Data_List_Lazy_zipWith(f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_1.UnsafePtr).V1, (*Constructor_Data_List_Lazy_Types_Cons)(v1_5.UnsafePtr).V1)}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
})
}))
_ = __local_var_3_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Lazy_force(), __local_var_3_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_2))
}))
}

func Call_Data_List_Lazy_zipWith__2064722709(f_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)
_ = __local_var_4_1
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_List_Lazy_Types_Cons
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr == nil) {
__t2 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_2
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr == nil) {
__t2 = (*Constructor_Data_List_Lazy_Types_Cons)(nil)
goto end_branch_2
} else {

}
}
{
if ((__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr != nil)) && ((v1_5.Type == 9 && v1_5.IntVal == 218341868 && v1_5.UnsafePtr != nil)) {
__t2 = &Constructor_Data_List_Lazy_Types_Cons{1, gopurs_runtime.Apply2(f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_1.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons)(v1_5.UnsafePtr).V0), Call_Data_List_Lazy_zipWith(f_0, (*Constructor_Data_List_Lazy_Types_Cons)(__local_var_4_1.UnsafePtr).V1, (*Constructor_Data_List_Lazy_Types_Cons)(v1_5.UnsafePtr).V1)}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
})
}))
_ = __local_var_3_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Lazy_force(), __local_var_3_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_2))
}))
}



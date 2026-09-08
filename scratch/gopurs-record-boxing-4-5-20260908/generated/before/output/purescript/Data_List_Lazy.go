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
		cache_Data_List_Lazy_one = gopurs_runtime.Int(int64(1))
	})
	return cache_Data_List_Lazy_one
}

var cache_Data_List_Lazy_identity gopurs_runtime.Value
var once_Data_List_Lazy_identity sync.Once
func Get_Data_List_Lazy_identity() gopurs_runtime.Value {
	once_Data_List_Lazy_identity.Do(func() {
		cache_Data_List_Lazy_identity = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
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
return Call_Data_List_Lazy_zipWithA(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), f_1_box, xs_2_box, ys_3_box)
})
	})
	return cache_Data_List_Lazy_zipWithA
}

var cache_Data_List_Lazy_zip gopurs_runtime.Value
var once_Data_List_Lazy_zip sync.Once
func Get_Data_List_Lazy_zip() gopurs_runtime.Value {
	once_Data_List_Lazy_zip.Do(func() {
		cache_Data_List_Lazy_zip = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, ys_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_zip(xs_0_box, ys_1_box)
})
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
return func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_unzip(xs_0_box)
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()
})
	})
	return cache_Data_List_Lazy_unzip
}

var cache_Data_List_Lazy_uncons gopurs_runtime.Value
var once_Data_List_Lazy_uncons sync.Once
func Get_Data_List_Lazy_uncons() gopurs_runtime.Value {
	once_Data_List_Lazy_uncons.Do(func() {
		cache_Data_List_Lazy_uncons = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_uncons(xs_0_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_Lazy_uncons
}

var cache_Data_List_Lazy_toUnfoldable gopurs_runtime.Value
var once_Data_List_Lazy_toUnfoldable sync.Once
func Get_Data_List_Lazy_toUnfoldable() gopurs_runtime.Value {
	once_Data_List_Lazy_toUnfoldable.Do(func() {
		cache_Data_List_Lazy_toUnfoldable = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
})
	})
	return cache_Data_List_Lazy_toUnfoldable
}

var cache_Data_List_Lazy_takeWhile gopurs_runtime.Value
var once_Data_List_Lazy_takeWhile sync.Once
func Get_Data_List_Lazy_takeWhile() gopurs_runtime.Value {
	once_Data_List_Lazy_takeWhile.Do(func() {
		cache_Data_List_Lazy_takeWhile = gopurs_runtime.Func(func(p_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_takeWhile(p_0_box)
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
return func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_tail(xs_0_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_Lazy_tail
}

var cache_Data_List_Lazy_stripPrefix gopurs_runtime.Value
var once_Data_List_Lazy_stripPrefix sync.Once
func Get_Data_List_Lazy_stripPrefix() gopurs_runtime.Value {
	once_Data_List_Lazy_stripPrefix.Do(func() {
		cache_Data_List_Lazy_stripPrefix = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_stripPrefix(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), v_1_box, s_2_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_Lazy_stripPrefix
}

var cache_Data_List_Lazy_span gopurs_runtime.Value
var once_Data_List_Lazy_span sync.Once
func Get_Data_List_Lazy_span() gopurs_runtime.Value {
	once_Data_List_Lazy_span.Do(func() {
		cache_Data_List_Lazy_span = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := Call_Data_List_Lazy_span(p_0_box, xs_1_box)
				_ = orig
				return gopurs_runtime.RecordDict([]string{"init", "rest"}, []gopurs_runtime.Value{orig.go__init, orig.rest})
				}()
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
return Call_Data_List_Lazy_replicateM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
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
return func() gopurs_runtime.Value {
				orig := Call_Data_List_Lazy_partition(f_0_box, xs_1_box)
				_ = orig
				return gopurs_runtime.RecordDict([]string{"no", "yes"}, []gopurs_runtime.Value{orig.no, orig.yes})
				}()
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
return Call_Data_List_Lazy_nub(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_List_Lazy_nub
}

var cache_Data_List_Lazy_newtypePattern gopurs_runtime.Value
var once_Data_List_Lazy_newtypePattern sync.Once
func Get_Data_List_Lazy_newtypePattern() gopurs_runtime.Value {
	once_Data_List_Lazy_newtypePattern.Do(func() {
		cache_Data_List_Lazy_newtypePattern = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})}))}
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
return Call_Data_List_Lazy_some(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](dictAlternative_0_box))
})
	})
	return cache_Data_List_Lazy_some
}

var cache_Data_List_Lazy_many gopurs_runtime.Value
var once_Data_List_Lazy_many sync.Once
func Get_Data_List_Lazy_many() gopurs_runtime.Value {
	once_Data_List_Lazy_many.Do(func() {
		cache_Data_List_Lazy_many = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_many(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](dictAlternative_0_box))
})
	})
	return cache_Data_List_Lazy_many
}

var cache_Data_List_Lazy_length gopurs_runtime.Value
var once_Data_List_Lazy_length sync.Once
func Get_Data_List_Lazy_length() gopurs_runtime.Value {
	once_Data_List_Lazy_length.Do(func() {
		cache_Data_List_Lazy_length = func() gopurs_runtime.Value {
var Call_local_Data_List_Lazy_go__go_0_0_17 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_0_0_17
var go__go_0_0_17 gopurs_runtime.Value
_ = go__go_0_0_17
Call_local_Data_List_Lazy_go__go_0_0_17 = func(b_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_0_0_17:
for {
if false { continue go__go_0_0_17 }
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
// TAST (Let): v_3_1 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2))
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
b_1_loop = gopurs_runtime.Int((b_1.IntVal) + (int64(1)))
xs_2_loop = (v_3_1).V1
continue go__go_0_0_17
__t2 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_0_0_17 = gopurs_runtime.Func(func(b_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_0_0_17(b_1_loop_val, xs_2_loop_val)
})
})
return gopurs_runtime.Int(gopurs_runtime.Apply(go__go_0_0_17, gopurs_runtime.Int(int64(0))).IntVal)
}()
	})
	return cache_Data_List_Lazy_length
}

var cache_Data_List_Lazy_last gopurs_runtime.Value
var once_Data_List_Lazy_last sync.Once
func Get_Data_List_Lazy_last() gopurs_runtime.Value {
	once_Data_List_Lazy_last.Do(func() {
		cache_Data_List_Lazy_last = func() gopurs_runtime.Value {
var Call_local_Data_List_Lazy_go__go_0_0_18 func(*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
_ = Call_local_Data_List_Lazy_go__go_0_0_18
var go__go_0_0_18 gopurs_runtime.Value
_ = go__go_0_0_18
Call_local_Data_List_Lazy_go__go_0_0_18 = func(v_1_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
go__go_0_0_18:
for {
if false { continue go__go_0_0_18 }
var v_1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v_1_loop
_ = v_1
var __t4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v_1 != nil) {
var __t3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
// TAST (Let): __local_var_2_1 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: (TypeVar a), tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])] Any))])
__local_var_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_uncons((v_1).V1)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = __local_var_2_1
var __t2 gopurs_runtime.Value
{
if (__local_var_2_1 == nil) {
__t2 = gopurs_runtime.Bool(true)
goto end_branch_2
} else {

}
}
{
if (__local_var_2_1 != nil) {
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
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(v_1).V0, true}
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
v_1_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (v_1).V1))
continue go__go_0_0_18
__t3 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("unreachable") }()
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
}
end_branch_4:
return __t4
}
}
go__go_0_0_18 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Lazy_go__go_0_0_18(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_1_loop_val)))}
})
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Lazy_go__go_0_0_18(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), x_1))))}
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

var cache_Data_List_Lazy_go__init gopurs_runtime.Value
var once_Data_List_Lazy_go__init sync.Once
func Get_Data_List_Lazy_go__init() gopurs_runtime.Value {
	once_Data_List_Lazy_go__init.Do(func() {
		cache_Data_List_Lazy_go__init = func() gopurs_runtime.Value {
var go__go_0_0_20 gopurs_runtime.Value
_ = go__go_0_0_20
// FALLBACK TCO: isLoop=false len=1
go__go_0_0_20 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_1)
if (__t_tag_1 != nil) {
var __t8 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
// TAST (Let): __local_var_2_6 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: (TypeVar a), tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])] Any))])
__local_var_2_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_uncons((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = __local_var_2_6
var __t7 gopurs_runtime.Value
{
if (__local_var_2_6 == nil) {
__t7 = gopurs_runtime.Bool(true)
goto end_branch_7
} else {

}
}
{
if (__local_var_2_6 != nil) {
__t7 = gopurs_runtime.Bool(false)
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
if (__t7.IntVal) != (0) {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{Get_Data_List_Lazy_Types_nil(), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_8
} else {

}
}
{
// TAST (Let): __local_var_2_2 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])])
__local_var_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(go__go_0_0_20, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)))}))
_ = __local_var_2_2
var __t5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_2_2 != nil) {
// TAST (Let): __local_var_3_3 shape=Other bindingType=Any
__local_var_3_3 := (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0
_ = __local_var_3_3
// TAST (Let): __local_var_4_4 shape=Other bindingType=Any
__local_var_4_4 := (__local_var_2_2).V0
_ = __local_var_4_4
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_3_3, __local_var_4_4}))}
})), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_5:
__t8 = __t5
}
end_branch_8:
__t9 = __t8
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t9)}
})
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(go__go_0_0_20, gopurs_runtime.Apply(Get_Data_Lazy_force(), x_1))
})
}()
	})
	return cache_Data_List_Lazy_go__init
}

var cache_Data_List_Lazy_index gopurs_runtime.Value
var once_Data_List_Lazy_index sync.Once
func Get_Data_List_Lazy_index() gopurs_runtime.Value {
	once_Data_List_Lazy_index.Do(func() {
		cache_Data_List_Lazy_index = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_index(xs_0_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_Lazy_index
}

var cache_Data_List_Lazy_head gopurs_runtime.Value
var once_Data_List_Lazy_head sync.Once
func Get_Data_List_Lazy_head() gopurs_runtime.Value {
	once_Data_List_Lazy_head.Do(func() {
		cache_Data_List_Lazy_head = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_head(xs_0_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
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
		cache_Data_List_Lazy_groupBy = gopurs_runtime.Func(func(eq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_groupBy(eq_0_box)
})
	})
	return cache_Data_List_Lazy_groupBy
}

var cache_Data_List_Lazy_group gopurs_runtime.Value
var once_Data_List_Lazy_group sync.Once
func Get_Data_List_Lazy_group() gopurs_runtime.Value {
	once_Data_List_Lazy_group.Do(func() {
		cache_Data_List_Lazy_group = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_group(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_Data_List_Lazy_group
}

var cache_Data_List_Lazy_fromStep gopurs_runtime.Value
var once_Data_List_Lazy_fromStep sync.Once
func Get_Data_List_Lazy_fromStep() gopurs_runtime.Value {
	once_Data_List_Lazy_fromStep.Do(func() {
		cache_Data_List_Lazy_fromStep = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_fromStep(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](x_0_box))
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
return Call_Data_List_Lazy_insert(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_List_Lazy_insert
}

var cache_Data_List_Lazy_fromFoldable gopurs_runtime.Value
var once_Data_List_Lazy_fromFoldable sync.Once
func Get_Data_List_Lazy_fromFoldable() gopurs_runtime.Value {
	once_Data_List_Lazy_fromFoldable.Do(func() {
		cache_Data_List_Lazy_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
})
	})
	return cache_Data_List_Lazy_fromFoldable
}

var cache_Data_List_Lazy_foldrLazy gopurs_runtime.Value
var once_Data_List_Lazy_foldrLazy sync.Once
func Get_Data_List_Lazy_foldrLazy() gopurs_runtime.Value {
	once_Data_List_Lazy_foldrLazy.Do(func() {
		cache_Data_List_Lazy_foldrLazy = gopurs_runtime.Func3(func(dictLazy_0_box gopurs_runtime.Value, op_1_box gopurs_runtime.Value, z_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_foldrLazy(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dictLazy_0_box), op_1_box, z_2_box)
})
	})
	return cache_Data_List_Lazy_foldrLazy
}

var cache_Data_List_Lazy_foldM gopurs_runtime.Value
var once_Data_List_Lazy_foldM sync.Once
func Get_Data_List_Lazy_foldM() gopurs_runtime.Value {
	once_Data_List_Lazy_foldM.Do(func() {
		cache_Data_List_Lazy_foldM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_foldM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_Data_List_Lazy_foldM
}

var cache_Data_List_Lazy_findIndex gopurs_runtime.Value
var once_Data_List_Lazy_findIndex sync.Once
func Get_Data_List_Lazy_findIndex() gopurs_runtime.Value {
	once_Data_List_Lazy_findIndex.Do(func() {
		cache_Data_List_Lazy_findIndex = gopurs_runtime.Func(func(fn_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_findIndex(fn_0_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_Lazy_findIndex
}

var cache_Data_List_Lazy_findLastIndex gopurs_runtime.Value
var once_Data_List_Lazy_findLastIndex sync.Once
func Get_Data_List_Lazy_findLastIndex() gopurs_runtime.Value {
	once_Data_List_Lazy_findLastIndex.Do(func() {
		cache_Data_List_Lazy_findLastIndex = gopurs_runtime.Func2(func(fn_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_findLastIndex(fn_0_box, xs_1_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_Lazy_findLastIndex
}

var cache_Data_List_Lazy_filterM gopurs_runtime.Value
var once_Data_List_Lazy_filterM sync.Once
func Get_Data_List_Lazy_filterM() gopurs_runtime.Value {
	once_Data_List_Lazy_filterM.Do(func() {
		cache_Data_List_Lazy_filterM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_filterM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
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
return Call_Data_List_Lazy_intersect(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_Data_List_Lazy_intersect
}

var cache_Data_List_Lazy_nubByEq gopurs_runtime.Value
var once_Data_List_Lazy_nubByEq sync.Once
func Get_Data_List_Lazy_nubByEq() gopurs_runtime.Value {
	once_Data_List_Lazy_nubByEq.Do(func() {
		cache_Data_List_Lazy_nubByEq = gopurs_runtime.Func(func(eq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_nubByEq(eq_0_box)
})
	})
	return cache_Data_List_Lazy_nubByEq
}

var cache_Data_List_Lazy_nubEq gopurs_runtime.Value
var once_Data_List_Lazy_nubEq sync.Once
func Get_Data_List_Lazy_nubEq() gopurs_runtime.Value {
	once_Data_List_Lazy_nubEq.Do(func() {
		cache_Data_List_Lazy_nubEq = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_nubEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
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
return func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_elemLastIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_List_Lazy_elemLastIndex
}

var cache_Data_List_Lazy_elemIndex gopurs_runtime.Value
var once_Data_List_Lazy_elemIndex sync.Once
func Get_Data_List_Lazy_elemIndex() gopurs_runtime.Value {
	once_Data_List_Lazy_elemIndex.Do(func() {
		cache_Data_List_Lazy_elemIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_elemIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
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
return Call_Data_List_Lazy_union(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
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

var cache_Data_List_Lazy_go__delete gopurs_runtime.Value
var once_Data_List_Lazy_go__delete sync.Once
func Get_Data_List_Lazy_go__delete() gopurs_runtime.Value {
	once_Data_List_Lazy_go__delete.Do(func() {
		cache_Data_List_Lazy_go__delete = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_go__delete(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_Data_List_Lazy_go__delete
}

var cache_Data_List_Lazy_difference gopurs_runtime.Value
var once_Data_List_Lazy_difference sync.Once
func Get_Data_List_Lazy_difference() gopurs_runtime.Value {
	once_Data_List_Lazy_difference.Do(func() {
		cache_Data_List_Lazy_difference = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_difference(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
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
		cache_Data_List_Lazy_catMaybes = Call_Data_List_Lazy_mapMaybe(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
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
		cache_Data_List_Lazy_modifyAt = gopurs_runtime.Func3(func(n_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_modifyAt(n_0_box.IntVal, f_1_box, xs_2_box)
})
	})
	return cache_Data_List_Lazy_modifyAt
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
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeVar b)])
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=(TypeVar a)
__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)
_ = __local_var_4_1
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr == nil) {
__t5 = (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_5
} else {

}
}
{
var __t_tag_2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_5)
if (__t_tag_2 == nil) {
__t5 = (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_5
} else {

}
}
{
var __t_and_4 bool = false
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr != nil) {

var __t_tag_3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_5)
__t_and_4 = (__t_tag_3 != nil)
}
if __t_and_4 {
__t5 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0), Call_Data_List_Lazy_zipWith(f_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)})
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t5)}
})
}))
_ = __local_var_3_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Lazy_force(), __local_var_3_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_2))
}))
}
}

func Call_Data_List_Lazy_zipWithA(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value, ys_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
var ys_3 gopurs_runtime.Value = ys_3_loop
_ = ys_3
// TAST (Let): Apply0_4_0 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m)])
Apply0_4_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}))
_ = Apply0_4_0
// TAST (Let): Functor0_5_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_1
var Call_local_Data_List_Lazy_go__go_6_2_0 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_6_2_0
var go__go_6_2_0 gopurs_runtime.Value
_ = go__go_6_2_0
Call_local_Data_List_Lazy_go__go_6_2_0 = func(b_7_loop gopurs_runtime.Value, xs_8_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_6_2_0:
for {
if false { continue go__go_6_2_0 }
var b_7 gopurs_runtime.Value = b_7_loop
_ = b_7
var xs_8 gopurs_runtime.Value = xs_8_loop
_ = xs_8
// TAST (Let): v_9_3 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_9_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_8))
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
__t4 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_6_2_0 = gopurs_runtime.Func(func(b_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_6_2_0(b_7_loop_val, xs_8_loop_val)
})
})
var Call_local_Data_List_Lazy_go__go_7_5_1 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_7_5_1
var go__go_7_5_1 gopurs_runtime.Value
_ = go__go_7_5_1
Call_local_Data_List_Lazy_go__go_7_5_1 = func(b_8_loop gopurs_runtime.Value, xs_9_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_7_5_1:
for {
if false { continue go__go_7_5_1 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var xs_9 gopurs_runtime.Value = xs_9_loop
_ = xs_9
// TAST (Let): v_10_6 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_10_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_9))
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
// TAST (Let): __local_var_11_7 shape=Other bindingType=(TypeVar a)
__local_var_11_7 := (v_10_6).V0
_ = __local_var_11_7
b_8_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_11_7, b_8}))}
}))
xs_9_loop = (v_10_6).V1
continue go__go_7_5_1
__t8 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_7_5_1 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_7_5_1(b_8_loop_val, xs_9_loop_val)
})
})
// TAST (Let): __local_var_8_9 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeVar b)])
__local_var_8_9 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_10 shape=App(Var) bindingType=(TypeVar a)
__local_var_9_10 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)
_ = __local_var_9_10
return gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_9_10.Type == 9 && __local_var_9_10.IntVal == 218341868 && __local_var_9_10.UnsafePtr == nil) {
__t14 = (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_14
} else {

}
}
{
var __t_tag_11 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_10)
if (__t_tag_11 == nil) {
__t14 = (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_14
} else {

}
}
{
var __t_and_13 bool = false
if (__local_var_9_10.Type == 9 && __local_var_9_10.IntVal == 218341868 && __local_var_9_10.UnsafePtr != nil) {

var __t_tag_12 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_10)
__t_and_13 = (__t_tag_12 != nil)
}
if __t_and_13 {
__t14 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_9_10.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_10.UnsafePtr).V0), Call_Data_List_Lazy_zipWith(f_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_9_10.UnsafePtr).V1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_10.UnsafePtr).V1)})
goto end_branch_14
} else {

}
}
{
__t14 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_14:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t14)}
})
}))
_ = __local_var_8_9
return Call_local_Data_List_Lazy_go__go_6_2_0(gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), Get_Data_List_Lazy_Types_nil()), Call_local_Data_List_Lazy_go__go_7_5_1(Get_Data_List_Lazy_Types_nil(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Lazy_force(), __local_var_8_9, gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_3))
}))))
}

func Call_Data_List_Lazy_zip(xs_0_loop gopurs_runtime.Value, ys_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var ys_1 gopurs_runtime.Value = ys_1_loop
_ = ys_1
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeVar b)])
__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 shape=App(Var) bindingType=(TypeVar a)
__local_var_3_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_0)
_ = __local_var_3_1
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_3_1.Type == 9 && __local_var_3_1.IntVal == 218341868 && __local_var_3_1.UnsafePtr == nil) {
__t5 = (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_5
} else {

}
}
{
var __t_tag_2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_4)
if (__t_tag_2 == nil) {
__t5 = (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_5
} else {

}
}
{
var __t_and_4 bool = false
if (__local_var_3_1.Type == 9 && __local_var_3_1.IntVal == 218341868 && __local_var_3_1.UnsafePtr != nil) {

var __t_tag_3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_4)
__t_and_4 = (__t_tag_3 != nil)
}
if __t_and_4 {
__t5 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_3_1.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V0}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}, Call_Data_List_Lazy_zipWith(Get_Data_Tuple_Tuple(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_3_1.UnsafePtr).V1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V1)})
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t5)}
})
}))
_ = __local_var_2_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Lazy_force(), __local_var_2_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_1))
}))
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
// TAST (Let): __local_var_4_0 shape=App(Var) bindingType=(TypeVar a)
__local_var_4_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)
_ = __local_var_4_0
var __t2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_2
} else {

}
}
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr != nil) {
var __t1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (n_0) == (int64(0)) {
__t1 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V1})
goto end_branch_1
} else {

}
}
{
__t1 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V0, Call_Data_List_Lazy_updateAt((n_0) - (int64(1)), x_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V1)})
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
}))
}
}

func Call_Data_List_Lazy_unzip(xs_0_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var Call_local_Data_List_Lazy_go__go_1_0_2 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_1_0_2
var go__go_1_0_2 gopurs_runtime.Value
_ = go__go_1_0_2
Call_local_Data_List_Lazy_go__go_1_0_2 = func(b_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_1_0_2:
for {
if false { continue go__go_1_0_2 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
// TAST (Let): v_4_1 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_3))
_ = v_4_1
var __t6 gopurs_runtime.Value
{
if (v_4_1 == nil) {
__t6 = b_2
goto end_branch_6
} else {

}
}
{
if (v_4_1 != nil) {
// TAST (Let): __local_var_5_2 shape=Other bindingType=Any
__local_var_5_2 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((v_4_1).V0.UnsafePtr).V0
_ = __local_var_5_2
// TAST (Let): __local_var_6_3 shape=Other bindingType=Any
__local_var_6_3 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((v_4_1).V0.UnsafePtr).V1
_ = __local_var_6_3
// TAST (Let): __local_var_7_4 shape=Other bindingType=Any
__local_var_7_4 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_2.UnsafePtr).V0
_ = __local_var_7_4
// TAST (Let): __local_var_8_5 shape=Other bindingType=Any
__local_var_8_5 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(b_2.UnsafePtr).V1
_ = __local_var_8_5
b_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_5_2, __local_var_7_4}))}
})), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_6_3, __local_var_8_5}))}
}))}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
xs_3_loop = (v_4_1).V1
continue go__go_1_0_2
__t6 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_1_0_2 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_1_0_2(b_2_loop_val, xs_3_loop_val)
})
})
var Call_local_Data_List_Lazy_go__go_2_7_3 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_2_7_3
var go__go_2_7_3 gopurs_runtime.Value
_ = go__go_2_7_3
Call_local_Data_List_Lazy_go__go_2_7_3 = func(b_3_loop gopurs_runtime.Value, xs_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_2_7_3:
for {
if false { continue go__go_2_7_3 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var xs_4 gopurs_runtime.Value = xs_4_loop
_ = xs_4
// TAST (Let): v_5_8 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_5_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_4))
_ = v_5_8
var __t10 gopurs_runtime.Value
{
if (v_5_8 == nil) {
__t10 = b_3
goto end_branch_10
} else {

}
}
{
if (v_5_8 != nil) {
// TAST (Let): __local_var_6_9 shape=Other bindingType=(TypeVar a)
__local_var_6_9 := (v_5_8).V0
_ = __local_var_6_9
b_3_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_6_9, b_3}))}
}))
xs_4_loop = (v_5_8).V1
continue go__go_2_7_3
__t10 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_2_7_3 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_2_7_3(b_3_loop_val, xs_4_loop_val)
})
})
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
				_v := Call_local_Data_List_Lazy_go__go_1_0_2(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{Get_Data_List_Lazy_Types_nil(), Get_Data_List_Lazy_Types_nil()}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}, Call_local_Data_List_Lazy_go__go_2_7_3(Get_Data_List_Lazy_Types_nil(), xs_0))
				_p := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{V0: _p.V0, V1: _p.V1}
			}()
}

func Call_Data_List_Lazy_uncons(xs_0_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
// TAST (Let): v_1_0 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_0))
_ = v_1_0
var __t1 *Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail gopurs_runtime.Value
}]
{
if (v_1_0 == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
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
if (v_1_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.RecordDict2("head", "tail", (v_1_0).V0, (v_1_0).V1), true}
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
__t1 = func() *Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail gopurs_runtime.Value
}] { panic("Failed pattern match") }()
}
end_branch_1:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_3833657837_3094389156(__t1))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_List_Lazy_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dictUnfoldable_0.V1), gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: (TypeVar a), tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])] Any))])
__local_var_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_uncons(xs_1)
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
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(__local_var_2_0).V0.head, (__local_var_2_0).V0.tail}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}, true}
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
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}))
}

func Call_Data_List_Lazy_takeWhile(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
takeWhile:
for {
if false { continue takeWhile }
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_1)
if ((__t_tag_2 != nil)) && ((gopurs_runtime.Apply(p_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0).IntVal) != (0)) {
__t3 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0, gopurs_runtime.Apply(Call_Data_List_Lazy_takeWhile(p_0), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t3)}
}))
_ = __local_var_1_1
// TAST (Let): __local_var_1_0 shape=Let(Abs(App(Other))) bindingType=(Func [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])] (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]))
__local_var_1_0 := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_1, x_2)
})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
}
}

func Call_Data_List_Lazy_take(n_0_loop int64) gopurs_runtime.Value {
take:
for {
if false { continue take }
var n_0 int64 = n_0_loop
_ = n_0
var __t5 gopurs_runtime.Value
{
if (n_0) <= (int64(0)) {
__t5 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_nil()
})
goto end_branch_5
} else {

}
}
{
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_1)
if (__t_tag_2 == nil) {
__t4 = (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_4
} else {

}
}
{
var __t_tag_3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_1)
if (__t_tag_3 != nil) {
__t4 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_1.UnsafePtr).V0, gopurs_runtime.Apply(Call_Data_List_Lazy_take((n_0) - (int64(1))), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_1.UnsafePtr).V1)})
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t4)}
}))
_ = __local_var_1_1
// TAST (Let): __local_var_1_0 shape=Let(Abs(App(Other))) bindingType=(Func [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])] (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]))
__local_var_1_0 := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_1, x_2)
})
_ = __local_var_1_0
__t5 = gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
}
end_branch_5:
return __t5
}
}

func Call_Data_List_Lazy_tail(xs_0_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), head: (TypeVar a)] Any))])
__local_var_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_uncons(xs_0)
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
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(__local_var_1_0).V0.tail, true}
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

func Call_Data_List_Lazy_stripPrefix(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], v_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
var Call_local_Data_List_Lazy___local_var_3_0 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy___local_var_3_0
var __local_var_3_0 gopurs_runtime.Value
_ = __local_var_3_0
Call_local_Data_List_Lazy___local_var_3_0 = func(prefix_3_loop gopurs_runtime.Value, input_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var prefix_3 gopurs_runtime.Value = prefix_3_loop
_ = prefix_3
var input_4 gopurs_runtime.Value = input_4_loop
_ = input_4
// TAST (Let): v1_5_1 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v1_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), prefix_3))
_ = v1_5_1
var __t4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v1_5_1 == nil) {
__t4 = (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_1239976062_3603546092((&Constructor_Control_Monad_Rec_Class_Done[struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}, gopurs_runtime.Value]{1, input_4})))}})
goto end_branch_4
} else {

}
}
{
if (v1_5_1 != nil) {
// TAST (Let): v2_6_2 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v2_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), input_4))
_ = v2_6_2
var __t3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((v2_6_2 != nil)) && ((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), (v1_5_1).V0, (v2_6_2).V0).IntVal) != (0)) {
__t3 = (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_784458114_4008603408((&Constructor_Control_Monad_Rec_Class_Loop[struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}, gopurs_runtime.Value]{1, func() struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
} {
					orig := gopurs_runtime.RecordDict2("a", "b", (v1_5_1).V1, (v2_6_2).V1)
					_ = orig
					clone := struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}{}
					clone.a = gopurs_runtime.RecordGet(orig, "a")
					clone.b = gopurs_runtime.RecordGet(orig, "b")
					return clone
				}()})))}})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
}
__local_var_3_0 = gopurs_runtime.Func(func(prefix_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(input_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy___local_var_3_0(prefix_3_loop_val, input_4_loop_val)
})
})
// TAST (Let): __local_var_4_5 shape=LitRecord bindingType=(Record (Row [a: (TypeVar a), b: (TypeVar b)] Any))
__local_var_4_5 := struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}{v_1, s_2}
_ = __local_var_4_5
var Call_local_Data_List_Lazy___local_var_5_6 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy___local_var_5_6
var __local_var_5_6 gopurs_runtime.Value
_ = __local_var_5_6
Call_local_Data_List_Lazy___local_var_5_6 = func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t12 gopurs_runtime.Value
{
var __t_tag_7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_5)
if (__t_tag_7 == nil) {
__t12 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_802708012_3603546092((&Constructor_Control_Monad_Rec_Class_Done[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})})))}
goto end_branch_12
} else {

}
}
{
var __t_tag_8 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_5)
if (__t_tag_8 != nil) {
var __t11 gopurs_runtime.Value
{
var __t_tag_9 gopurs_runtime.Value = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_5.UnsafePtr).V0
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 525585346) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_3406595152_4008603408((&Constructor_Control_Monad_Rec_Class_Loop[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](Call_local_Data_List_Lazy___local_var_3_0(gopurs_runtime.RecordGet((*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_5.UnsafePtr).V0.UnsafePtr).V0, "a"), gopurs_runtime.RecordGet((*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_5.UnsafePtr).V0.UnsafePtr).V0, "b")))})))}
goto end_branch_11
} else {

}
}
{
var __t_tag_10 gopurs_runtime.Value = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_5.UnsafePtr).V0
if (__t_tag_10.Type == 9 && __t_tag_10.IntVal == 60402430) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_802708012_3603546092((&Constructor_Control_Monad_Rec_Class_Done[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_5.UnsafePtr).V0.UnsafePtr).V0})})))}
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
__t12 = __t11
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
__local_var_5_6 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy___local_var_5_6(v_5_loop_val)
})
var Call_local_Data_List_Lazy_go__go_6_13_4 func(gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_6_13_4
var go__go_6_13_4 gopurs_runtime.Value
_ = go__go_6_13_4
Call_local_Data_List_Lazy_go__go_6_13_4 = func(v_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_6_13_4:
for {
if false { continue go__go_6_13_4 }
var v_7 gopurs_runtime.Value = v_7_loop
_ = v_7
var __t14 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 525585346) {
v_7_loop = Call_local_Data_List_Lazy___local_var_5_6((*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0)
continue go__go_6_13_4
__t14 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_14
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 60402430) {
__t14 = (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0
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
}
go__go_6_13_4 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_6_13_4(v_7_loop_val)
})
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](Call_local_Data_List_Lazy_go__go_6_13_4(Call_local_Data_List_Lazy___local_var_5_6(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](Call_local_Data_List_Lazy___local_var_3_0(__local_var_4_5.a, __local_var_4_5.b)))}))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_List_Lazy_span(p_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) struct{
	go__init gopurs_runtime.Value
	rest gopurs_runtime.Value
} {
span:
for {
if false { continue span }
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
// TAST (Let): v_2_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: (TypeVar a), tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])] Any))])
v_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_uncons(xs_1)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = v_2_0
var __t4 struct{
	go__init gopurs_runtime.Value
	rest gopurs_runtime.Value
}
{
if ((v_2_0 != nil)) && ((gopurs_runtime.Apply(p_0, (v_2_0).V0.head).IntVal) != (0)) {
// TAST (Let): __local_var_3_1 shape=Other bindingType=Any
__local_var_3_1 := (v_2_0).V0.head
_ = __local_var_3_1
// TAST (Let): v1_4_2 shape=App(Var) bindingType=(Record (Row [init: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), rest: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])] Any))
v1_4_2 := Call_Data_List_Lazy_span(p_0, (v_2_0).V0.tail)
_ = v1_4_2
// TAST (Let): __local_var_5_3 shape=Other bindingType=Any
__local_var_5_3 := v1_4_2.go__init
_ = __local_var_5_3
__t4 = struct{
	go__init gopurs_runtime.Value
	rest gopurs_runtime.Value
}{gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_3_1, __local_var_5_3}))}
})), v1_4_2.rest}
goto end_branch_4
} else {

}
}
{
__t4 = struct{
	go__init gopurs_runtime.Value
	rest gopurs_runtime.Value
}{Get_Data_List_Lazy_Types_nil(), xs_1}
}
end_branch_4:
return __t4
}
}

func Call_Data_List_Lazy_snoc(xs_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var Call_local_Data_List_Lazy_go__go_2_0_5 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_2_0_5
var go__go_2_0_5 gopurs_runtime.Value
_ = go__go_2_0_5
Call_local_Data_List_Lazy_go__go_2_0_5 = func(b_3_loop gopurs_runtime.Value, xs_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_2_0_5:
for {
if false { continue go__go_2_0_5 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var xs_4 gopurs_runtime.Value = xs_4_loop
_ = xs_4
// TAST (Let): v_5_1 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_4))
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
// TAST (Let): __local_var_6_2 shape=Other bindingType=(TypeVar a)
__local_var_6_2 := (v_5_1).V0
_ = __local_var_6_2
b_3_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_6_2, b_3}))}
}))
xs_4_loop = (v_5_1).V1
continue go__go_2_0_5
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
go__go_2_0_5 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_2_0_5(b_3_loop_val, xs_4_loop_val)
})
})
var Call_local_Data_List_Lazy_go__go_3_4_6 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_3_4_6
var go__go_3_4_6 gopurs_runtime.Value
_ = go__go_3_4_6
Call_local_Data_List_Lazy_go__go_3_4_6 = func(b_4_loop gopurs_runtime.Value, xs_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_3_4_6:
for {
if false { continue go__go_3_4_6 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_5 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
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
// TAST (Let): __local_var_7_6 shape=Other bindingType=(TypeVar a)
__local_var_7_6 := (v_6_5).V0
_ = __local_var_7_6
b_4_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_7_6, b_4}))}
}))
xs_5_loop = (v_6_5).V1
continue go__go_3_4_6
__t7 = func() gopurs_runtime.Value { panic("unreachable") }()
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
go__go_3_4_6 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_3_4_6(b_4_loop_val, xs_5_loop_val)
})
})
return Call_local_Data_List_Lazy_go__go_2_0_5(gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_1, Get_Data_List_Lazy_Types_nil()}))}
})), Call_local_Data_List_Lazy_go__go_3_4_6(Get_Data_List_Lazy_Types_nil(), xs_0))
}

func Call_Data_List_Lazy_singleton(a_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, a_0, Get_Data_List_Lazy_Types_nil()}))}
}))
}

func Call_Data_List_Lazy_showPattern(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): showList_1_0 shape=LitRecord bindingType=(ADT ["Data","Show","Show"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])])
showList_1_0 := (&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_2_1 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1))
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
var Call_local_Data_List_Lazy_go__go_3_2_7 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_3_2_7
var go__go_3_2_7 gopurs_runtime.Value
_ = go__go_3_2_7
Call_local_Data_List_Lazy_go__go_3_2_7 = func(b_4_loop gopurs_runtime.Value, xs_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_3_2_7:
for {
if false { continue go__go_3_2_7 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_3 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
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
__t4 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_3_2_7 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_3_2_7(b_4_loop_val, xs_5_loop_val)
})
})
__t5 = ((("(fromFoldable [") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (v_2_1).V0).StrVal())) + (Call_local_Data_List_Lazy_go__go_3_2_7(gopurs_runtime.Str(""), (v_2_1).V1).StrVal())) + ("])")
goto end_branch_5
} else {

}
}
{
__t5 = func() string { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Str(__t5)
})})
_ = showList_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Pattern ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showList_1_0.V0), v_2).StrVal())) + (")"))
})}))}
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
// TAST (Let): __local_var_4_0 shape=App(Var) bindingType=(TypeVar a)
__local_var_4_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)
_ = __local_var_4_0
var __t2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr == nil) {
__t2 = (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_2
} else {

}
}
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr != nil) {
// TAST (Let): acc_prime__5_1 shape=App(Other) bindingType=(TypeVar b)
acc_prime__5_1 := gopurs_runtime.Apply2(f_0, acc_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V0)
_ = acc_prime__5_1
__t2 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, acc_prime__5_1, Call_Data_List_Lazy_scanlLazy(f_0, acc_prime__5_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V1)})
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
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
var Call_local_Data_List_Lazy_go__go_2_0_8 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_2_0_8
var go__go_2_0_8 gopurs_runtime.Value
_ = go__go_2_0_8
Call_local_Data_List_Lazy_go__go_2_0_8 = func(b_3_loop gopurs_runtime.Value, xs_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_2_0_8:
for {
if false { continue go__go_2_0_8 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var xs_4 gopurs_runtime.Value = xs_4_loop
_ = xs_4
// TAST (Let): v_5_1 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_4))
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
// TAST (Let): __local_var_6_2 shape=Other bindingType=(TypeVar a)
__local_var_6_2 := (v_5_1).V0
_ = __local_var_6_2
b_3_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_6_2, b_3}))}
}))
xs_4_loop = (v_5_1).V1
continue go__go_2_0_8
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
go__go_2_0_8 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_2_0_8(b_3_loop_val, xs_4_loop_val)
})
})
return gopurs_runtime.Apply(Get_Data_Lazy_force(), Call_local_Data_List_Lazy_go__go_2_0_8(Get_Data_List_Lazy_Types_nil(), xs_0))
}))
}

func Call_Data_List_Lazy_replicateM(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
replicateM:
for {
if false { continue replicateM }
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func2(func(n_3 gopurs_runtime.Value, m_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (n_3.IntVal) < (int64(1)) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), Get_Data_List_Lazy_Types_nil())
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(Call_Data_List_Lazy_replicateM(dictMonad_0), gopurs_runtime.Int((n_3.IntVal) - (int64(1))), m_4), gopurs_runtime.Func(func(as_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, a_5, as_6}))}
})))
}))
}))
}
end_branch_2:
return __t2
})
}
}

func Call_Data_List_Lazy_repeat(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var go__go_1_0_9 gopurs_runtime.Value
_ = go__go_1_0_9
// FALLBACK TCO: isLoop=false len=1
go__go_1_0_9 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_0, go__go_1_0_9}))}
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
var __t12 gopurs_runtime.Value
{
if (start_0) > (end_1) {
var go__go_2_6_11 gopurs_runtime.Value
_ = go__go_2_6_11
// FALLBACK TCO: isLoop=false len=1
go__go_2_6_11 = gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_6_7 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar b)])])
v1_6_7 := Rebox_Data_List_Lazy_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_3, b_4)))
_ = v1_6_7
var __t10 gopurs_runtime.Value
{
if (v1_6_7 == nil) {
__t10 = Get_Data_List_Lazy_Types_nil()
goto end_branch_10
} else {

}
}
{
if (v1_6_7 != nil) {
// TAST (Let): __local_var_7_8 shape=Other bindingType=Any
__local_var_7_8 := ((v1_6_7).V0).V0
_ = __local_var_7_8
// TAST (Let): __local_var_8_9 shape=App(Other) bindingType=Any
__local_var_8_9 := gopurs_runtime.Apply2(go__go_2_6_11, f_3, ((v1_6_7).V0).V1)
_ = __local_var_8_9
__t10 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_7_8, __local_var_8_9}))}
}))
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return gopurs_runtime.Apply(Get_Data_Lazy_force(), __t10)
}))
})
__t12 = gopurs_runtime.Apply2(go__go_2_6_11, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
{
if (x_3.IntVal) >= (end_1) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_3363075976_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(x_3.IntVal), gopurs_runtime.Int((x_3.IntVal) - (int64(1)))}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}}))}
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_1415037225_3094389156(Rebox_Data_List_Lazy_3094389156_1415037225(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))))}
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_1415037225_3094389156(Rebox_Data_List_Lazy_3094389156_1415037225(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t11))))}
}), gopurs_runtime.Int(start_0))
goto end_branch_12
} else {

}
}
{
var go__go_2_0_10 gopurs_runtime.Value
_ = go__go_2_0_10
// FALLBACK TCO: isLoop=false len=1
go__go_2_0_10 = gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_6_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar b)])])
v1_6_1 := Rebox_Data_List_Lazy_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_3, b_4)))
_ = v1_6_1
var __t4 gopurs_runtime.Value
{
if (v1_6_1 == nil) {
__t4 = Get_Data_List_Lazy_Types_nil()
goto end_branch_4
} else {

}
}
{
if (v1_6_1 != nil) {
// TAST (Let): __local_var_7_2 shape=Other bindingType=Any
__local_var_7_2 := ((v1_6_1).V0).V0
_ = __local_var_7_2
// TAST (Let): __local_var_8_3 shape=App(Other) bindingType=Any
__local_var_8_3 := gopurs_runtime.Apply2(go__go_2_0_10, f_3, ((v1_6_1).V0).V1)
_ = __local_var_8_3
__t4 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_7_2, __local_var_8_3}))}
}))
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Apply(Get_Data_Lazy_force(), __t4)
}))
})
__t12 = gopurs_runtime.Apply2(go__go_2_0_10, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (x_3.IntVal) <= (end_1) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_3363075976_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, int64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(x_3.IntVal), gopurs_runtime.Int((x_3.IntVal) + (int64(1)))}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}}))}
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_1415037225_3094389156(Rebox_Data_List_Lazy_3094389156_1415037225(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))))}
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_1415037225_3094389156(Rebox_Data_List_Lazy_3094389156_1415037225(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t5))))}
}), gopurs_runtime.Int(start_0))
}
end_branch_12:
return __t12
}

func Call_Data_List_Lazy_partition(f_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) struct{
	no gopurs_runtime.Value
	yes gopurs_runtime.Value
} {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var Call_local_Data_List_Lazy_go__go_2_0_12 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_2_0_12
var go__go_2_0_12 gopurs_runtime.Value
_ = go__go_2_0_12
Call_local_Data_List_Lazy_go__go_2_0_12 = func(b_3_loop gopurs_runtime.Value, xs_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_2_0_12:
for {
if false { continue go__go_2_0_12 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var xs_4 gopurs_runtime.Value = xs_4_loop
_ = xs_4
// TAST (Let): v_5_1 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_4))
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
var __t2 struct{
	no gopurs_runtime.Value
	yes gopurs_runtime.Value
}
{
if (gopurs_runtime.Apply(f_0, (v_5_1).V0).IntVal) != (0) {
__t2 = struct{
	no gopurs_runtime.Value
	yes gopurs_runtime.Value
}{gopurs_runtime.RecordGet(b_3, "no"), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (v_5_1).V0, gopurs_runtime.RecordGet(b_3, "yes")}))}
}))}
goto end_branch_2
} else {

}
}
{
__t2 = struct{
	no gopurs_runtime.Value
	yes gopurs_runtime.Value
}{gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (v_5_1).V0, gopurs_runtime.RecordGet(b_3, "no")}))}
})), gopurs_runtime.RecordGet(b_3, "yes")}
}
end_branch_2:
b_3_loop = func() gopurs_runtime.Value {
				orig := __t2
				_ = orig
				return gopurs_runtime.RecordDict([]string{"no", "yes"}, []gopurs_runtime.Value{orig.no, orig.yes})
				}()
xs_4_loop = (v_5_1).V1
continue go__go_2_0_12
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
go__go_2_0_12 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_2_0_12(b_3_loop_val, xs_4_loop_val)
})
})
var Call_local_Data_List_Lazy_go__go_3_4_13 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_3_4_13
var go__go_3_4_13 gopurs_runtime.Value
_ = go__go_3_4_13
Call_local_Data_List_Lazy_go__go_3_4_13 = func(b_4_loop gopurs_runtime.Value, xs_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_3_4_13:
for {
if false { continue go__go_3_4_13 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_5 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
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
// TAST (Let): __local_var_7_6 shape=Other bindingType=(TypeVar a)
__local_var_7_6 := (v_6_5).V0
_ = __local_var_7_6
b_4_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_7_6, b_4}))}
}))
xs_5_loop = (v_6_5).V1
continue go__go_3_4_13
__t7 = func() gopurs_runtime.Value { panic("unreachable") }()
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
go__go_3_4_13 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_3_4_13(b_4_loop_val, xs_5_loop_val)
})
})
return func() struct{
	no gopurs_runtime.Value
	yes gopurs_runtime.Value
} {
					orig := Call_local_Data_List_Lazy_go__go_2_0_12(func() gopurs_runtime.Value {
				orig := struct{
	no gopurs_runtime.Value
	yes gopurs_runtime.Value
}{Get_Data_List_Lazy_Types_nil(), Get_Data_List_Lazy_Types_nil()}
				_ = orig
				return gopurs_runtime.RecordDict([]string{"no", "yes"}, []gopurs_runtime.Value{orig.no, orig.yes})
				}(), Call_local_Data_List_Lazy_go__go_3_4_13(Get_Data_List_Lazy_Types_nil(), xs_1))
					_ = orig
					clone := struct{
	no gopurs_runtime.Value
	yes gopurs_runtime.Value
}{}
					clone.no = gopurs_runtime.RecordGet(orig, "no")
					clone.yes = gopurs_runtime.RecordGet(orig, "yes")
					return clone
				}()
}

func Call_Data_List_Lazy_null(x_0_loop gopurs_runtime.Value) bool {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: (TypeVar a), tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])] Any))])
__local_var_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_uncons(x_0)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0 == nil) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
if (__local_var_1_0 != nil) {
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
var goStep_1_0_14 gopurs_runtime.Value
_ = goStep_1_0_14
// FALLBACK TCO: isLoop=false len=2
var go__go_1_1_15 gopurs_runtime.Value
_ = go__go_1_1_15
// FALLBACK TCO: isLoop=false len=2
goStep_1_0_14 = gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_3)
if (__t_tag_2 == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_6
} else {

}
}
{
var __t_tag_3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_3)
if (__t_tag_3 != nil) {
// TAST (Let): v2_4_4 shape=App(Var) bindingType=(Record (Row [found: Boolean, result: (ADT ["Data","List","Internal","Set"] [(TypeVar a)])] Any))
v2_4_4 := func() struct{
	found bool
	result gopurs_runtime.Value
} {
					orig := gopurs_runtime.Apply3(Get_Data_List_Internal_insertAndLookupBy(), p_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, v_2)
					_ = orig
					clone := struct{
	found bool
	result gopurs_runtime.Value
}{}
					clone.found = (gopurs_runtime.RecordGet(orig, "found").IntVal) != (0)
					clone.result = gopurs_runtime.RecordGet(orig, "result")
					return clone
				}()
_ = v2_4_4
var __t5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if v2_4_4.found {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply2(go__go_1_1_15, v2_4_4.result, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)))
goto end_branch_5
} else {

}
}
{
__t5 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, gopurs_runtime.Apply2(go__go_1_1_15, v2_4_4.result, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t6)}
})
go__go_1_1_15 = gopurs_runtime.Func2(func(s_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_7 shape=App(Other) bindingType=Any
__local_var_4_7 := gopurs_runtime.Apply(goStep_1_0_14, s_2)
_ = __local_var_4_7
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_7, gopurs_runtime.Apply(Get_Data_Lazy_force(), v_3))
}))
})
return gopurs_runtime.Apply(go__go_1_1_15, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)})
}

func Call_Data_List_Lazy_nub(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return Call_Data_List_Lazy_nubBy(gopurs_runtime.Box(dictOrd_0.V1))
}

func Call_Data_List_Lazy_mapMaybe(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
mapMaybe:
for {
if false { continue mapMaybe }
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var Call_local_Data_List_Lazy_go__go_1_0_16 func(*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Lazy_go__go_1_0_16
var go__go_1_0_16 gopurs_runtime.Value
_ = go__go_1_0_16
Call_local_Data_List_Lazy_go__go_1_0_16 = func(v_2_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] {
go__go_1_0_16:
for {
if false { continue go__go_1_0_16 }
var v_2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var __t3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (v_2 == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_3
} else {

}
}
{
if (v_2 != nil) {
// TAST (Let): v1_3_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])
v1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, (v_2).V0))
_ = v1_3_1
var __t2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (v1_3_1 == nil) {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (v_2).V1))
continue go__go_1_0_16
__t2 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_2
} else {

}
}
{
if (v1_3_1 != nil) {
__t2 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (v1_3_1).V0, gopurs_runtime.Apply(Call_Data_List_Lazy_mapMaybe(f_0), (v_2).V1)})
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}
go__go_1_0_16 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Lazy_go__go_1_0_16(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_2_loop_val)))}
})
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Lazy_go__go_1_0_16(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), x_2))))}
}))
})
}
}

func Call_Data_List_Lazy_some(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): Apply0_1_0 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar f)])
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V0), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V1), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func2(func(dictLazy_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_List_Lazy_Types_cons(), v_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_3, "defer"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Data_List_Lazy_many(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dictLazy_3))}, v_4)
})))
})
}

func Call_Data_List_Lazy_many(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): Alt0_1_0 shape=App(Other) bindingType=(ADT ["Control","Alt","Alt"] [(TypeVar f)])
Alt0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V1), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}))
_ = Alt0_1_0
// TAST (Let): Applicative0_2_1 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar f)])
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func2(func(dictLazy_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Alt0_1_0.V1), gopurs_runtime.Apply2(Call_Data_List_Lazy_some(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dictLazy_3))}, v_4), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), Get_Data_List_Lazy_Types_nil()))
})
}

func Call_Data_List_Lazy_iterate(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var go__go_2_0_19 gopurs_runtime.Value
_ = go__go_2_0_19
// FALLBACK TCO: isLoop=false len=1
go__go_2_0_19 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(TypeVar b)])
__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_2 shape=App(Var) bindingType=(TypeVar a)
__local_var_5_2 := gopurs_runtime.Apply(Get_Data_Lazy_force(), go__go_2_0_19)
_ = __local_var_5_2
var __t3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_5_2.Type == 9 && __local_var_5_2.IntVal == 218341868 && __local_var_5_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_3
} else {

}
}
{
if (__local_var_5_2.Type == 9 && __local_var_5_2.IntVal == 218341868 && __local_var_5_2.UnsafePtr != nil) {
__t3 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()).V0), f_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_2.UnsafePtr).V1)})
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t3)}
}))
_ = __local_var_4_1
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_1, __local_var_4_1}))}
})))
}))
return go__go_2_0_19
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
if (v_0) == (int64(0)) {
__t2 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, v1_1, v2_2}))}
}))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 shape=App(Var) bindingType=(TypeVar a)
__local_var_4_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v2_2)
_ = __local_var_4_0
var __t1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr == nil) {
__t1 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, v1_1, Get_Data_List_Lazy_Types_nil()})
goto end_branch_1
} else {

}
}
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr != nil) {
__t1 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V0, Call_Data_List_Lazy_insertAt((v_0) - (int64(1)), v1_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V1)})
goto end_branch_1
} else {

}
}
{
__t1 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t1)}
}))
}
end_branch_2:
return __t2
}
}

func Call_Data_List_Lazy_index(xs_0_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var Call_local_Data_List_Lazy_go__go_1_0_21 func(*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value], int64) *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
_ = Call_local_Data_List_Lazy_go__go_1_0_21
var go__go_1_0_21 gopurs_runtime.Value
_ = go__go_1_0_21
Call_local_Data_List_Lazy_go__go_1_0_21 = func(v_2_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value], v1_3_loop int64) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
go__go_1_0_21:
for {
if false { continue go__go_1_0_21 }
var v_2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 int64 = v1_3_loop
_ = v1_3
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v_2 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
goto end_branch_2
} else {

}
}
{
if (v_2 != nil) {
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v1_3) == (int64(0)) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(v_2).V0, true}
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
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (v_2).V1))
v1_3_loop = (v1_3) - (int64(1))
continue go__go_1_0_21
__t1 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("unreachable") }()
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}
go__go_1_0_21 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Lazy_go__go_1_0_21(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_2_loop_val), v1_3_loop_val.IntVal))}
})
})
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply(go__go_1_0_21, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_0)))})
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_List_Lazy_head(xs_0_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: (TypeVar a), tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])] Any))])
__local_var_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_uncons(xs_0)
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
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(__local_var_1_0).V0.head, true}
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

func Call_Data_List_Lazy_transpose(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
transpose:
for {
if false { continue transpose }
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
// TAST (Let): v_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(ADT ["Data","List","Lazy","Types","List"] [(TypeVar a)])])])] Any))])
v_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_uncons(xs_0)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = v_1_0
var __t9 gopurs_runtime.Value
{
if (v_1_0 == nil) {
__t9 = xs_0
goto end_branch_9
} else {

}
}
{
if (v_1_0 != nil) {
// TAST (Let): v1_2_1 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: (TypeVar a), tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])] Any))])
v1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_uncons((v_1_0).V0.head)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
_ = v1_2_1
var __t8 gopurs_runtime.Value
{
if (v1_2_1 == nil) {
xs_0_loop = (v_1_0).V0.tail
continue transpose
__t8 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_8
} else {

}
}
{
if (v1_2_1 != nil) {
// TAST (Let): __local_var_3_2 shape=Other bindingType=Any
__local_var_3_2 := (v1_2_1).V0.head
_ = __local_var_3_2
// TAST (Let): __local_var_4_3 shape=Other bindingType=Any
__local_var_4_3 := (v1_2_1).V0.tail
_ = __local_var_4_3
// TAST (Let): __local_var_5_4 shape=App(Var) bindingType=Any
__local_var_5_4 := gopurs_runtime.Apply(Call_Data_List_Lazy_mapMaybe(Get_Data_List_Lazy_head()), (v_1_0).V0.tail)
_ = __local_var_5_4
// TAST (Let): __local_var_6_5 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])
__local_var_6_5 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_3_2, __local_var_5_4}))}
}))
_ = __local_var_6_5
// TAST (Let): __local_var_7_7 shape=App(Var) bindingType=Any
__local_var_7_7 := gopurs_runtime.Apply(Call_Data_List_Lazy_mapMaybe(Get_Data_List_Lazy_tail()), (v_1_0).V0.tail)
_ = __local_var_7_7
// TAST (Let): __local_var_7_6 shape=App(Var) bindingType=Any
__local_var_7_6 := Call_Data_List_Lazy_transpose(gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_4_3, __local_var_7_7}))}
})))
_ = __local_var_7_6
__t8 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_6_5, __local_var_7_6}))}
}))
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
__t9 = __t8
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
}
}

func Call_Data_List_Lazy_groupBy(eq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
groupBy:
for {
if false { continue groupBy }
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_1)
if (__t_tag_2 == nil) {
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_7
} else {

}
}
{
var __t_tag_3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_1)
if (__t_tag_3 != nil) {
// TAST (Let): __local_var_2_4 shape=Other bindingType=Any
__local_var_2_4 := (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0
_ = __local_var_2_4
// TAST (Let): v1_3_5 shape=App(Var) bindingType=(Record (Row [init: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), rest: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])] Any))
v1_3_5 := Call_Data_List_Lazy_span(gopurs_runtime.Apply(eq_0, __local_var_2_4), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)
_ = v1_3_5
// TAST (Let): __local_var_4_6 shape=Other bindingType=Any
__local_var_4_6 := v1_3_5.go__init
_ = __local_var_4_6
__t7 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_2_4, __local_var_4_6}))}
})), gopurs_runtime.Apply(Call_Data_List_Lazy_groupBy(eq_0), v1_3_5.rest)})
goto end_branch_7
} else {

}
}
{
__t7 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t7)}
}))
_ = __local_var_1_1
// TAST (Let): __local_var_1_0 shape=Let(Abs(App(Other))) bindingType=(Func [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])] (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a)])])])]))
__local_var_1_0 := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_1, x_2)
})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
}
}

func Call_Data_List_Lazy_group(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return Call_Data_List_Lazy_groupBy(gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_List_Lazy_fromStep(x_0_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var x_0 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = x_0_loop
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
// TAST (Let): __local_var_4_0 shape=App(Var) bindingType=(TypeVar a)
__local_var_4_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)
_ = __local_var_4_0
var __t3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr == nil) {
__t3 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_1, Get_Data_List_Lazy_Types_nil()})
goto end_branch_3
} else {

}
}
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr != nil) {
var __t2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_1 uint32 = uint32(gopurs_runtime.Apply2(cmp_0, x_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V0).IntVal)
if (uint32(__t_tag_1) == 380165415) {
__t2 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V0, Call_Data_List_Lazy_insertBy(cmp_0, x_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V1)})
goto end_branch_2
} else {

}
}
{
__t2 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](__local_var_4_0))}
}))})
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t3)}
}))
}
}

func Call_Data_List_Lazy_insert(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_0 shape=Other bindingType=(Func [(TypeVar a), (TypeVar a)] (ADT ["Data","Ordering","Ordering"] []))
__local_var_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = __local_var_1_0
return gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_1 shape=App(Var) bindingType=(TypeVar a)
__local_var_5_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_3)
_ = __local_var_5_1
var __t4 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_5_1.Type == 9 && __local_var_5_1.IntVal == 218341868 && __local_var_5_1.UnsafePtr == nil) {
__t4 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_2, Get_Data_List_Lazy_Types_nil()})
goto end_branch_4
} else {

}
}
{
if (__local_var_5_1.Type == 9 && __local_var_5_1.IntVal == 218341868 && __local_var_5_1.UnsafePtr != nil) {
var __t3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_1_0, x_2, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V0)
if (uint32(__t_tag_2.IntVal) == 380165415) {
__t3 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V0, Call_Data_List_Lazy_insertBy(__local_var_1_0, x_2, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V1)})
goto end_branch_3
} else {

}
}
{
__t3 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_2, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](__local_var_5_1))}
}))})
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t4)}
}))
})
}

func Call_Data_List_Lazy_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V2), Get_Data_List_Lazy_Types_cons(), Get_Data_List_Lazy_Types_nil())
}

func Call_Data_List_Lazy_foldrLazy(dictLazy_0_loop *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value], op_1_loop gopurs_runtime.Value, z_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictLazy_0 *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value] = dictLazy_0_loop
_ = dictLazy_0
var op_1 gopurs_runtime.Value = op_1_loop
_ = op_1
var z_2 gopurs_runtime.Value = z_2_loop
_ = z_2
var go__go_3_0_22 gopurs_runtime.Value
_ = go__go_3_0_22
// FALLBACK TCO: isLoop=false len=1
go__go_3_0_22 = gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_5_1 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_4))
_ = v_5_1
var __t4 gopurs_runtime.Value
{
if (v_5_1 != nil) {
// TAST (Let): __local_var_6_2 shape=Other bindingType=Any
__local_var_6_2 := (v_5_1).V0
_ = __local_var_6_2
// TAST (Let): __local_var_7_3 shape=Other bindingType=Any
__local_var_7_3 := (v_5_1).V1
_ = __local_var_7_3
__t4 = gopurs_runtime.Apply(gopurs_runtime.Box(dictLazy_0.V0), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_1, __local_var_6_2, gopurs_runtime.Apply(go__go_3_0_22, __local_var_7_3))
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
return go__go_3_0_22
}

func Call_Data_List_Lazy_foldM(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
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
return gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value, xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_6_2 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: (TypeVar a), tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])] Any))])
v_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_uncons(xs_5)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
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
// TAST (Let): __local_var_7_3 shape=Other bindingType=Any
__local_var_7_3 := (v_6_2).V0.tail
_ = __local_var_7_3
__t4 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(f_3, b_4, (v_6_2).V0.head), gopurs_runtime.Func(func(b_prime__8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Call_Data_List_Lazy_foldM(dictMonad_0), f_3, b_prime__8, __local_var_7_3)
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
}
}

func Call_Data_List_Lazy_findIndex(fn_0_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var Call_local_Data_List_Lazy_go__2988133926_1_0_23 func(int64, gopurs_runtime.Value) *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
_ = Call_local_Data_List_Lazy_go__2988133926_1_0_23
var go__2988133926_1_0_23 gopurs_runtime.Value
_ = go__2988133926_1_0_23
Call_local_Data_List_Lazy_go__2988133926_1_0_23 = func(n_2_loop int64, list_3_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
go__2988133926_1_0_23:
for {
if false { continue go__2988133926_1_0_23 }
var n_2 int64 = n_2_loop
_ = n_2
var list_3 gopurs_runtime.Value = list_3_loop
_ = list_3
// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=Any
__local_var_4_1 := Call_Data_List_Lazy_uncons(list_3)
_ = __local_var_4_1
var __t3 gopurs_runtime.Value
{
if (func() gopurs_runtime.Value {
				_v := __local_var_4_1
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}().Type == 9 && func() gopurs_runtime.Value {
				_v := __local_var_4_1
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}().IntVal == 930809136 && func() gopurs_runtime.Value {
				_v := __local_var_4_1
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}().UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(fn_0, gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(func() gopurs_runtime.Value {
				_v := __local_var_4_1
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}().UnsafePtr).V0, "head")).IntVal) != (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int(n_2), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}
goto end_branch_2
} else {

}
}
{
n_2_loop = (n_2) + (int64(1))
list_3_loop = gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(func() gopurs_runtime.Value {
				_v := __local_var_4_1
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}().UnsafePtr).V0, "tail")
continue go__2988133926_1_0_23
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_1170268447_3094389156(func() *Constructor_Data_Maybe_Just[int64] { panic("unreachable") }()))}
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_1170268447_3094389156(Rebox_Data_List_Lazy_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t2))))}
goto end_branch_3
} else {

}
}
{
if (func() gopurs_runtime.Value {
				_v := __local_var_4_1
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}().Type == 9 && func() gopurs_runtime.Value {
				_v := __local_var_4_1
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}().IntVal == 930809136 && func() gopurs_runtime.Value {
				_v := __local_var_4_1
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}().UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }())}
}
end_branch_3:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t3)
}
}
go__2988133926_1_0_23 = gopurs_runtime.Func(func(n_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(list_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Lazy_go__2988133926_1_0_23(n_2_loop_val.IntVal, list_3_loop_val))}
})
})
var go__go_2_4_24 gopurs_runtime.Value
_ = go__go_2_4_24
// FALLBACK TCO: isLoop=false len=1
go__go_2_4_24 = gopurs_runtime.Func2(func(n_3 gopurs_runtime.Value, list_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_5 shape=App(Var) bindingType=Any
__local_var_5_5 := Call_Data_List_Lazy_uncons(list_4)
_ = __local_var_5_5
var __t7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (func() gopurs_runtime.Value {
				_v := __local_var_5_5
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}().Type == 9 && func() gopurs_runtime.Value {
				_v := __local_var_5_5
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}().IntVal == 930809136 && func() gopurs_runtime.Value {
				_v := __local_var_5_5
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}().UnsafePtr != nil) {
var __t6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply(fn_0, gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(func() gopurs_runtime.Value {
				_v := __local_var_5_5
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}().UnsafePtr).V0, "head")).IntVal) != (0) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{n_3, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_6
} else {

}
}
{
__t6 = Call_local_Data_List_Lazy_go__2988133926_1_0_23((n_3.IntVal) + (int64(1)), gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(func() gopurs_runtime.Value {
				_v := __local_var_5_5
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}().UnsafePtr).V0, "tail"))
}
end_branch_6:
__t7 = __t6
goto end_branch_7
} else {

}
}
{
if (func() gopurs_runtime.Value {
				_v := __local_var_5_5
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}().Type == 9 && func() gopurs_runtime.Value {
				_v := __local_var_5_5
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}().IntVal == 930809136 && func() gopurs_runtime.Value {
				_v := __local_var_5_5
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}().UnsafePtr == nil) {
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
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
__t7 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t7)}
})
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply(go__2988133926_1_0_23, gopurs_runtime.Int(int64(0)))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_List_Lazy_findLastIndex(fn_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
__local_var_2_0 := Rebox_Data_List_Lazy_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_findIndex(fn_0)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}(), Call_Data_List_Lazy_reverse(xs_1))))
_ = __local_var_2_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_2_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int(((gopurs_runtime.Apply(Get_Data_List_Lazy_length(), xs_1).IntVal) - (int64(1))) - (gopurs_runtime.Int((__local_var_2_0).V0).IntVal)), true}
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

func Call_Data_List_Lazy_filterM(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
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
return gopurs_runtime.Func2(func(p_3 gopurs_runtime.Value, list_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_5_2 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: (TypeVar a), tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])] Any))])
v_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_uncons(list_4)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
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
// TAST (Let): __local_var_6_3 shape=Other bindingType=Any
__local_var_6_3 := (v_5_2).V0.head
_ = __local_var_6_3
// TAST (Let): __local_var_7_4 shape=Other bindingType=Any
__local_var_7_4 := (v_5_2).V0.tail
_ = __local_var_7_4
__t6 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(p_3, __local_var_6_3), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(Call_Data_List_Lazy_filterM(dictMonad_0), p_3, __local_var_7_4), gopurs_runtime.Func(func(xs_prime__9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (b_8.IntVal) != (0) {
__t5 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_6_3, xs_prime__9}))}
}))
goto end_branch_5
} else {

}
}
{
__t5 = xs_prime__9
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
}
}

func Call_Data_List_Lazy_filter(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
filter:
for {
if false { continue filter }
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var Call_local_Data_List_Lazy_go__go_1_0_25 func(*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Lazy_go__go_1_0_25
var go__go_1_0_25 gopurs_runtime.Value
_ = go__go_1_0_25
Call_local_Data_List_Lazy_go__go_1_0_25 = func(v_2_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] {
go__go_1_0_25:
for {
if false { continue go__go_1_0_25 }
var v_2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var __t2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (v_2 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_2
} else {

}
}
{
if (v_2 != nil) {
var __t1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply(p_0, (v_2).V0).IntVal) != (0) {
__t1 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (v_2).V0, gopurs_runtime.Apply(Call_Data_List_Lazy_filter(p_0), (v_2).V1)})
goto end_branch_1
} else {

}
}
{
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (v_2).V1))
continue go__go_1_0_25
__t1 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}
go__go_1_0_25 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Lazy_go__go_1_0_25(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_2_loop_val)))}
})
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Lazy_go__go_1_0_25(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), x_2))))}
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
// TAST (Let): semigroupDisj1_4_0 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar a)])
semigroupDisj1_4_0 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((v_4.IntVal) != (0)) || ((v1_5.IntVal) != (0)))
})})
_ = semigroupDisj1_4_0
// TAST (Let): __local_var_5_1 shape=LitRecord bindingType=(ADT ["Data","Monoid","Monoid"] [(TypeVar a)])
__local_var_5_1 := (&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupDisj1_4_0)}
}), gopurs_runtime.Bool(false)})
_ = __local_var_5_1
// TAST (Let): Semigroup0_6_2 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
Semigroup0_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_5_1.V0), gopurs_runtime.Value{}))
_ = Semigroup0_6_2
// TAST (Let): __local_var_7_3 shape=App(Other) bindingType=(Func [(TypeVar a)] Boolean)
__local_var_7_3 := gopurs_runtime.Apply(eq_0, x_3)
_ = __local_var_7_3
var Call_local_Data_List_Lazy_go__go_8_4_26 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_8_4_26
var go__go_8_4_26 gopurs_runtime.Value
_ = go__go_8_4_26
Call_local_Data_List_Lazy_go__go_8_4_26 = func(b_9_loop gopurs_runtime.Value, xs_10_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_8_4_26:
for {
if false { continue go__go_8_4_26 }
var b_9 gopurs_runtime.Value = b_9_loop
_ = b_9
var xs_10 gopurs_runtime.Value = xs_10_loop
_ = xs_10
// TAST (Let): v_11_5 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_11_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_10))
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
continue go__go_8_4_26
__t6 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_8_4_26 = gopurs_runtime.Func(func(b_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_8_4_26(b_9_loop_val, xs_10_loop_val)
})
})
return gopurs_runtime.Bool((Call_local_Data_List_Lazy_go__go_8_4_26(gopurs_runtime.Box(__local_var_5_1.V1), ys_2).IntVal) != (0))
})), xs_1)
}

func Call_Data_List_Lazy_intersect(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_List_Lazy_intersectBy(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_List_Lazy_nubByEq(eq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
nubByEq:
for {
if false { continue nubByEq }
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_1)
if (__t_tag_2 == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_5
} else {

}
}
{
var __t_tag_3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_1)
if (__t_tag_3 != nil) {
// TAST (Let): __local_var_2_4 shape=Other bindingType=Any
__local_var_2_4 := (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0
_ = __local_var_2_4
__t5 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_2_4, gopurs_runtime.Apply(Call_Data_List_Lazy_nubByEq(eq_0), gopurs_runtime.Apply(Call_Data_List_Lazy_filter(gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(eq_0, __local_var_2_4, y_3).IntVal) != (0)) != (true))
})), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1))})
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t5)}
}))
_ = __local_var_1_1
// TAST (Let): __local_var_1_0 shape=Let(Abs(App(Other))) bindingType=(Func [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])] (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]))
__local_var_1_0 := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_1, x_2)
})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
}
}

func Call_Data_List_Lazy_nubEq(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return Call_Data_List_Lazy_nubByEq(gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_List_Lazy_eqPattern(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
// TAST (Let): eqList_1_0 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])])
eqList_1_0 := (&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(xs_1 gopurs_runtime.Value, ys_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_27 gopurs_runtime.Value
_ = go__go_3_1_27
// FALLBACK TCO: isLoop=false len=1
go__go_3_1_27 = gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 bool
{
var __t_tag_5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_4)
if (__t_tag_5 == nil) {
var __t_tag_6 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_5)
__t7 = (__t_tag_6 == nil)
goto end_branch_7
} else {

}
}
{
var __t_tag_2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_4)
var __t_and_4 bool = false
if (__t_tag_2 != nil) {

var __t_tag_3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_5)
__t_and_4 = ((__t_tag_3 != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(go__go_3_1_27, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)))}).IntVal) != (0)))
}
__t7 = __t_and_4
}
end_branch_7:
return gopurs_runtime.Bool(__t7)
})
return gopurs_runtime.Bool((gopurs_runtime.Apply2(go__go_3_1_27, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_2)))}).IntVal) != (0))
})})
_ = eqList_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(eqList_1_0.V0), x_2, y_3).IntVal) != (0))
})}))}
}

func Call_Data_List_Lazy_ordPattern(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_2 shape=App(Other) bindingType=Any
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_2
// TAST (Let): eqList1_1_1 shape=Let(LitRecord) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])])
eqList1_1_1 := (&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(xs_2 gopurs_runtime.Value, ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_3_28 gopurs_runtime.Value
_ = go__go_4_3_28
// FALLBACK TCO: isLoop=false len=1
go__go_4_3_28 = gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 bool
{
var __t_tag_7 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_5)
if (__t_tag_7 == nil) {
var __t_tag_8 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_6)
__t9 = (__t_tag_8 == nil)
goto end_branch_9
} else {

}
}
{
var __t_tag_4 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_5)
var __t_and_6 bool = false
if (__t_tag_4 != nil) {

var __t_tag_5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_6)
__t_and_6 = ((__t_tag_5 != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "eq"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(go__go_4_3_28, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_5.UnsafePtr).V1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)))}).IntVal) != (0)))
}
__t9 = __t_and_6
}
end_branch_9:
return gopurs_runtime.Bool(__t9)
})
return gopurs_runtime.Bool((gopurs_runtime.Apply2(go__go_4_3_28, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_3)))}).IntVal) != (0))
})})
_ = eqList1_1_1
// TAST (Let): ordList_1_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Ord","Ord"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])])
ordList_1_0 := (&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqList1_1_1)}
}), gopurs_runtime.Func2(func(xs_2 gopurs_runtime.Value, ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Lazy_go__go_4_10_29 func(*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) uint32
_ = Call_local_Data_List_Lazy_go__go_4_10_29
var go__go_4_10_29 gopurs_runtime.Value
_ = go__go_4_10_29
Call_local_Data_List_Lazy_go__go_4_10_29 = func(v_5_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) uint32 {
go__go_4_10_29:
for {
if false { continue go__go_4_10_29 }
var v_5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var __t14 uint32
{
if (v_5 == nil) {
var __t11 uint32
{
if (v1_6 == nil) {
__t11 = 902936544
goto end_branch_11
} else {

}
}
{
__t11 = 1527465420
}
end_branch_11:
__t14 = __t11
goto end_branch_14
} else {

}
}
{
if (v1_6 == nil) {
__t14 = 380165415
goto end_branch_14
} else {

}
}
{
if ((v_5 != nil)) && ((v1_6 != nil)) {
// TAST (Let): v2_7_12 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v2_7_12 := uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), (v_5).V0, (v1_6).V0).IntVal)
_ = v2_7_12
var __t13 uint32
{
if (v2_7_12 == 902936544) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (v_5).V1))
v1_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (v1_6).V1))
continue go__go_4_10_29
__t13 = func() uint32 { panic("unreachable") }()
goto end_branch_13
} else {

}
}
{
__t13 = v2_7_12
}
end_branch_13:
__t14 = __t13
goto end_branch_14
} else {

}
}
{
__t14 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_14:
return __t14
}
}
go__go_4_10_29 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_List_Lazy_go__go_4_10_29(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val))), UnsafePtr: nil}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_local_Data_List_Lazy_go__go_4_10_29(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_3)))), UnsafePtr: nil}
})})
_ = ordList_1_0
// TAST (Let): __local_var_2_16 shape=App(Other) bindingType=Any
__local_var_2_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_16
// TAST (Let): eqList_3_17 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])])
eqList_3_17 := (&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(xs_3 gopurs_runtime.Value, ys_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_5_18_30 gopurs_runtime.Value
_ = go__go_5_18_30
// FALLBACK TCO: isLoop=false len=1
go__go_5_18_30 = gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t24 bool
{
var __t_tag_22 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_6)
if (__t_tag_22 == nil) {
var __t_tag_23 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_7)
__t24 = (__t_tag_23 == nil)
goto end_branch_24
} else {

}
}
{
var __t_tag_19 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_6)
var __t_and_21 bool = false
if (__t_tag_19 != nil) {

var __t_tag_20 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_7)
__t_and_21 = ((__t_tag_20 != nil)) && (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_16, "eq"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_6.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_7.UnsafePtr).V0).IntVal) != (0)) && ((gopurs_runtime.Apply2(go__go_5_18_30, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_6.UnsafePtr).V1)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_7.UnsafePtr).V1)))}).IntVal) != (0)))
}
__t24 = __t_and_21
}
end_branch_24:
return gopurs_runtime.Bool(__t24)
})
return gopurs_runtime.Bool((gopurs_runtime.Apply2(go__go_5_18_30, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_3)))}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), ys_4)))}).IntVal) != (0))
})})
_ = eqList_3_17
// TAST (Let): eqPattern1_2_15 shape=Let(Let(LitRecord)) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])])
eqPattern1_2_15 := (&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(eqList_3_17.V0), x_4, y_5).IntVal) != (0))
})})
_ = eqPattern1_2_15
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqPattern1_2_15)}
}), gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(ordList_1_0.V1), x_3, y_4).IntVal)), UnsafePtr: nil}
})}))}
}

func Call_Data_List_Lazy_elemLastIndex(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply(Get_Data_List_Lazy_findLastIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), v_2, x_1).IntVal) != (0))
}))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_List_Lazy_elemIndex(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
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
var Call_local_Data_List_Lazy_go__go_1_0_31 func(*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_1_0_31
var go__go_1_0_31 gopurs_runtime.Value
_ = go__go_1_0_31
Call_local_Data_List_Lazy_go__go_1_0_31 = func(v_2_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_1_0_31:
for {
if false { continue go__go_1_0_31 }
var v_2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if ((v_2 != nil)) && ((gopurs_runtime.Apply(p_0, (v_2).V0).IntVal) != (0)) {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (v_2).V1))
continue go__go_1_0_31
__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_1_0_31 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_1_0_31(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_2_loop_val))
})
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_1_0_31(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), x_2)))
})
}

func Call_Data_List_Lazy_drop(n_0_loop int64) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var Call_local_Data_List_Lazy_go__2018343127_1_0_32 func(int64, *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Lazy_go__2018343127_1_0_32
var go__2018343127_1_0_32 gopurs_runtime.Value
_ = go__2018343127_1_0_32
Call_local_Data_List_Lazy_go__2018343127_1_0_32 = func(v_2_loop int64, v1_3_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] {
go__2018343127_1_0_32:
for {
if false { continue go__2018343127_1_0_32 }
var v_2 int64 = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (v_2) == (int64(0)) {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
if (v1_3 == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_1
} else {

}
}
{
if (v1_3 != nil) {
v_2_loop = (v_2) - (int64(1))
v1_3_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (v1_3).V1))
continue go__2018343127_1_0_32
__t1 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_1
} else {

}
}
{
__t1 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}
go__2018343127_1_0_32 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Lazy_go__2018343127_1_0_32(v_2_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val)))}
})
})
var go__go_2_2_33 gopurs_runtime.Value
_ = go__go_2_2_33
// FALLBACK TCO: isLoop=false len=1
go__go_2_2_33 = gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (v_3.IntVal) == (int64(0)) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_4)
goto end_branch_5
} else {

}
}
{
var __t_tag_3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_4)
if (__t_tag_3 == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_4)
if (__t_tag_4 != nil) {
__t5 = Call_local_Data_List_Lazy_go__2018343127_1_0_32((v_3.IntVal) - (int64(1)), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_4.UnsafePtr).V1)))
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t5)}
})
// TAST (Let): __local_var_3_7 shape=App(Other) bindingType=Any
__local_var_3_7 := gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0), gopurs_runtime.Apply(go__2018343127_1_0_32, gopurs_runtime.Int(n_0)))
_ = __local_var_3_7
// TAST (Let): __local_var_3_6 shape=Let(Abs(App(Other))) bindingType=(Func [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])] (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]))
__local_var_3_6 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_7, x_4)
})
_ = __local_var_3_6
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_6, x_4)
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
// TAST (Let): __local_var_4_0 shape=App(Var) bindingType=(TypeVar a)
__local_var_4_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)
_ = __local_var_4_0
var __t2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_2
} else {

}
}
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr != nil) {
var __t1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply2(eq_0, x_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V1))
goto end_branch_1
} else {

}
}
{
__t1 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V0, Call_Data_List_Lazy_deleteBy(eq_0, x_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V1)})
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
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
var Call_local_Data_List_Lazy_go__go_3_1_34 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_3_1_34
var go__go_3_1_34 gopurs_runtime.Value
_ = go__go_3_1_34
Call_local_Data_List_Lazy_go__go_3_1_34 = func(b_4_loop gopurs_runtime.Value, xs_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_3_1_34:
for {
if false { continue go__go_3_1_34 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_2 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_5))
_ = v_6_2
var __t7 gopurs_runtime.Value
{
if (v_6_2 == nil) {
__t7 = b_4
goto end_branch_7
} else {

}
}
{
if (v_6_2 != nil) {
// TAST (Let): __local_var_7_3 shape=Other bindingType=(TypeVar a)
__local_var_7_3 := (v_6_2).V0
_ = __local_var_7_3
b_4_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_4 shape=App(Var) bindingType=(TypeVar a)
__local_var_9_4 := gopurs_runtime.Apply(Get_Data_Lazy_force(), b_4)
_ = __local_var_9_4
var __t6 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_9_4.Type == 9 && __local_var_9_4.IntVal == 218341868 && __local_var_9_4.UnsafePtr == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_6
} else {

}
}
{
if (__local_var_9_4.Type == 9 && __local_var_9_4.IntVal == 218341868 && __local_var_9_4.UnsafePtr != nil) {
var __t5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply2(eq_0, __local_var_7_3, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_9_4.UnsafePtr).V0).IntVal) != (0) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_9_4.UnsafePtr).V1))
goto end_branch_5
} else {

}
}
{
__t5 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_9_4.UnsafePtr).V0, Call_Data_List_Lazy_deleteBy(eq_0, __local_var_7_3, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_9_4.UnsafePtr).V1)})
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t6)}
}))
xs_5_loop = (v_6_2).V1
continue go__go_3_1_34
__t7 = func() gopurs_runtime.Value { panic("unreachable") }()
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
go__go_3_1_34 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_3_1_34(b_4_loop_val, xs_5_loop_val)
})
})
// TAST (Let): __local_var_3_0 shape=LetRec(App(Other)) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])
__local_var_3_0 := Call_local_Data_List_Lazy_go__go_3_1_34(gopurs_runtime.Apply(Call_Data_List_Lazy_nubByEq(eq_0), ys_2), xs_1)
_ = __local_var_3_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_8 shape=App(Var) bindingType=(TypeVar a)
__local_var_5_8 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)
_ = __local_var_5_8
var __t9 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_5_8.Type == 9 && __local_var_5_8.IntVal == 218341868 && __local_var_5_8.UnsafePtr == nil) {
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_3_0))
goto end_branch_9
} else {

}
}
{
if (__local_var_5_8.Type == 9 && __local_var_5_8.IntVal == 218341868 && __local_var_5_8.UnsafePtr != nil) {
__t9 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_8.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_8.UnsafePtr).V1, __local_var_3_0)})
goto end_branch_9
} else {

}
}
{
__t9 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t9)}
}))
}

func Call_Data_List_Lazy_union(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
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
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(TypeVar a)
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_1)
_ = __local_var_3_0
var __t2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_2
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
var __t1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (n_0) == (int64(0)) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V1))
goto end_branch_1
} else {

}
}
{
__t1 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, Call_Data_List_Lazy_deleteAt((n_0) - (int64(1)), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V1)})
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
}))
}
}

func Call_Data_List_Lazy_go__delete(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
// TAST (Let): __local_var_1_0 shape=Other bindingType=(Func [(TypeVar a), (TypeVar a)] Boolean)
__local_var_1_0 := gopurs_runtime.Box(dictEq_0.V0)
_ = __local_var_1_0
return gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_1 shape=App(Var) bindingType=(TypeVar a)
__local_var_5_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_3)
_ = __local_var_5_1
var __t3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_5_1.Type == 9 && __local_var_5_1.IntVal == 218341868 && __local_var_5_1.UnsafePtr == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_3
} else {

}
}
{
if (__local_var_5_1.Type == 9 && __local_var_5_1.IntVal == 218341868 && __local_var_5_1.UnsafePtr != nil) {
var __t2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply2(__local_var_1_0, x_2, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V0).IntVal) != (0) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
__t2 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V0, Call_Data_List_Lazy_deleteBy(__local_var_1_0, x_2, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V1)})
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t3)}
}))
})
}

func Call_Data_List_Lazy_difference(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
// TAST (Let): __local_var_1_1 shape=App(Var) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(Get_Data_List_Lazy_go__delete(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_0)})
_ = __local_var_1_1
// TAST (Let): __local_var_1_0 shape=Let(Abs(Abs(App(Other)))) bindingType=(Func [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]), (TypeVar a)] (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])]))
__local_var_1_0 := gopurs_runtime.Func2(func(b_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_1_1, a_3, b_2)
})
_ = __local_var_1_0
var Call_local_Data_List_Lazy_go__go_2_2_35 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_2_2_35
var go__go_2_2_35 gopurs_runtime.Value
_ = go__go_2_2_35
Call_local_Data_List_Lazy_go__go_2_2_35 = func(b_3_loop gopurs_runtime.Value, xs_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_2_2_35:
for {
if false { continue go__go_2_2_35 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var xs_4 gopurs_runtime.Value = xs_4_loop
_ = xs_4
// TAST (Let): v_5_3 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])
v_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_4))
_ = v_5_3
var __t4 gopurs_runtime.Value
{
if (v_5_3 == nil) {
__t4 = b_3
goto end_branch_4
} else {

}
}
{
if (v_5_3 != nil) {
b_3_loop = gopurs_runtime.Apply2(__local_var_1_0, b_3, (v_5_3).V0)
xs_4_loop = (v_5_3).V1
continue go__go_2_2_35
__t4 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_2_2_35 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_2_2_35(b_3_loop_val, xs_4_loop_val)
})
})
return go__go_2_2_35
}

func Call_Data_List_Lazy_cycle(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var go__go_1_0_36 gopurs_runtime.Value
_ = go__go_1_0_36
// FALLBACK TCO: isLoop=false len=1
go__go_1_0_36 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=(TypeVar a)
__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_0)
_ = __local_var_4_1
var __t2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), go__go_1_0_36))
goto end_branch_2
} else {

}
}
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr != nil) {
__t2 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1, go__go_1_0_36)})
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
})))
}))
return go__go_1_0_36
}

func Call_Data_List_Lazy_concatMap(b_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(TypeVar a)
__local_var_3_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), a_1)
_ = __local_var_3_0
var __t8 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr == nil) {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_8
} else {

}
}
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 218341868 && __local_var_3_0.UnsafePtr != nil) {
// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b)])])
__local_var_4_1 := gopurs_runtime.Apply(b_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0)
_ = __local_var_4_1
// TAST (Let): __local_var_5_2 shape=App(Other) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b)])])
__local_var_5_2 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V1, b_0)
_ = __local_var_5_2
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_3 shape=App(Var) bindingType=(TypeVar a)
__local_var_7_3 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_1)
_ = __local_var_7_3
var __t7 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_7_3.Type == 9 && __local_var_7_3.IntVal == 218341868 && __local_var_7_3.UnsafePtr == nil) {
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_2))
goto end_branch_7
} else {

}
}
{
if (__local_var_7_3.Type == 9 && __local_var_7_3.IntVal == 218341868 && __local_var_7_3.UnsafePtr != nil) {
// TAST (Let): __local_var_8_4 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])
__local_var_8_4 := (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_7_3.UnsafePtr).V1
_ = __local_var_8_4
__t7 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_7_3.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_5 shape=App(Var) bindingType=(TypeVar a)
__local_var_10_5 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_8_4)
_ = __local_var_10_5
var __t6 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_10_5.Type == 9 && __local_var_10_5.IntVal == 218341868 && __local_var_10_5.UnsafePtr == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_5_2))
goto end_branch_6
} else {

}
}
{
if (__local_var_10_5.Type == 9 && __local_var_10_5.IntVal == 218341868 && __local_var_10_5.UnsafePtr != nil) {
__t6 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_10_5.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_10_5.UnsafePtr).V1, __local_var_5_2)})
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t6)}
}))})
goto end_branch_7
} else {

}
}
{
__t7 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t7)}
}))))
goto end_branch_8
} else {

}
}
{
__t8 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t8)}
}))
}

func Call_Data_List_Lazy_concat(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(TypeVar a)
__local_var_2_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), v_0)
_ = __local_var_2_0
var __t8 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 218341868 && __local_var_2_0.UnsafePtr == nil) {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_8
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 218341868 && __local_var_2_0.UnsafePtr != nil) {
// TAST (Let): __local_var_3_1 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b)])])
__local_var_3_1 := (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_2_0.UnsafePtr).V0
_ = __local_var_3_1
// TAST (Let): __local_var_4_2 shape=App(Other) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b)])])
__local_var_4_2 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_bindList()).V1), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_2_0.UnsafePtr).V1, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
_ = __local_var_4_2
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_3 shape=App(Var) bindingType=(TypeVar a)
__local_var_6_3 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_3_1)
_ = __local_var_6_3
var __t7 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_6_3.Type == 9 && __local_var_6_3.IntVal == 218341868 && __local_var_6_3.UnsafePtr == nil) {
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_2))
goto end_branch_7
} else {

}
}
{
if (__local_var_6_3.Type == 9 && __local_var_6_3.IntVal == 218341868 && __local_var_6_3.UnsafePtr != nil) {
// TAST (Let): __local_var_7_4 shape=Other bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a)])])
__local_var_7_4 := (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_3.UnsafePtr).V1
_ = __local_var_7_4
__t7 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_3.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_5 shape=App(Var) bindingType=(TypeVar a)
__local_var_9_5 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_7_4)
_ = __local_var_9_5
var __t6 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_9_5.Type == 9 && __local_var_9_5.IntVal == 218341868 && __local_var_9_5.UnsafePtr == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_2))
goto end_branch_6
} else {

}
}
{
if (__local_var_9_5.Type == 9 && __local_var_9_5.IntVal == 218341868 && __local_var_9_5.UnsafePtr != nil) {
__t6 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_9_5.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_9_5.UnsafePtr).V1, __local_var_4_2)})
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t6)}
}))})
goto end_branch_7
} else {

}
}
{
__t7 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t7)}
}))))
goto end_branch_8
} else {

}
}
{
__t8 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
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
// TAST (Let): __local_var_4_0 shape=App(Var) bindingType=(TypeVar a)
__local_var_4_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)
_ = __local_var_4_0
var __t4 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr == nil) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_4
} else {

}
}
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr != nil) {
var __t3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (n_0) == (int64(0)) {
// TAST (Let): v2_5_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
v2_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V0))
_ = v2_5_1
var __t2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (v2_5_1 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Lazy_force(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5_1 != nil) {
__t2 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (v2_5_1).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V1})
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V0, Call_Data_List_Lazy_alterAt((n_0) - (int64(1)), f_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V1)})
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t4)}
}))
}
}

func Call_Data_List_Lazy_modifyAt(n_0_loop int64, f_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 shape=App(Var) bindingType=(TypeVar a)
__local_var_4_0 := gopurs_runtime.Apply(Get_Data_Lazy_force(), xs_2)
_ = __local_var_4_0
var __t2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_2
} else {

}
}
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 218341868 && __local_var_4_0.UnsafePtr != nil) {
var __t1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (n_0) == (int64(0)) {
__t1 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V0), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V1})
goto end_branch_1
} else {

}
}
{
__t1 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V0, Call_Data_List_Lazy_alterAt((n_0) - (int64(1)), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, x_5)}))}
}), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_0.UnsafePtr).V1)})
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
}))
}

func Rebox_Data_List_Lazy_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_List_Lazy_1239976062_3603546092(in *Constructor_Control_Monad_Rec_Class_Done[struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}, gopurs_runtime.Value]) *Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_List_Lazy_138441832_3363075976(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[int64, int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[int64, int64]{}
		out.V0 = in.V0.IntVal
		out.V1 = in.V1.IntVal
	return out
}

func Rebox_Data_List_Lazy_1415037225_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, int64]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_3363075976_138441832(in.V0))}
	return out
}

func Rebox_Data_List_Lazy_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[int64]{}
		out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_List_Lazy_3094389156_1415037225(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, int64]] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, int64]]{}
		out.V0 = Rebox_Data_List_Lazy_138441832_3363075976(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
	return out
}

func Rebox_Data_List_Lazy_3094389156_4010058633(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0)
	return out
}

func Rebox_Data_List_Lazy_3363075976_138441832(in *Constructor_Data_Tuple_Tuple[int64, int64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
		out.V1 = gopurs_runtime.Int(in.V1)
	return out
}

func Rebox_Data_List_Lazy_3406595152_4008603408(in *Constructor_Control_Monad_Rec_Class_Loop[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}

func Rebox_Data_List_Lazy_3833657837_3094389156(in *Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail gopurs_runtime.Value
}]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = func() gopurs_runtime.Value {
				orig := in.V0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"head", "tail"}, []gopurs_runtime.Value{orig.head, orig.tail})
				}()
	return out
}

func Rebox_Data_List_Lazy_784458114_4008603408(in *Constructor_Control_Monad_Rec_Class_Loop[struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}, gopurs_runtime.Value]) *Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = func() gopurs_runtime.Value {
				orig := in.V0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"a", "b"}, []gopurs_runtime.Value{orig.a, orig.b})
				}()
	return out
}

func Rebox_Data_List_Lazy_802708012_3603546092(in *Constructor_Control_Monad_Rec_Class_Done[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
}



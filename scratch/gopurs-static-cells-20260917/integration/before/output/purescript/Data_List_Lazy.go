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
		cache_Data_List_Lazy_unwrap = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_List_Lazy_unwrap
}

var cache_Data_List_Lazy_unwrap1 gopurs_runtime.Value
var once_Data_List_Lazy_unwrap1 sync.Once
func Get_Data_List_Lazy_unwrap1() gopurs_runtime.Value {
	once_Data_List_Lazy_unwrap1.Do(func() {
		cache_Data_List_Lazy_unwrap1 = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_List_Lazy_unwrap1
}

var cache_Data_List_Lazy_one gopurs_runtime.Value
var once_Data_List_Lazy_one sync.Once
func Get_Data_List_Lazy_one() gopurs_runtime.Value {
	once_Data_List_Lazy_one.Do(func() {
		cache_Data_List_Lazy_one = gopurs_runtime.Int(Call_Data_Semiring_one(gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_348932501_2826095630(Rebox_Data_List_Lazy_2826095630_348932501(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringInt()))))}).IntVal)
	})
	return cache_Data_List_Lazy_one
}

var cache_Data_List_Lazy_unwrap2 gopurs_runtime.Value
var once_Data_List_Lazy_unwrap2 sync.Once
func Get_Data_List_Lazy_unwrap2() gopurs_runtime.Value {
	once_Data_List_Lazy_unwrap2.Do(func() {
		cache_Data_List_Lazy_unwrap2 = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_List_Lazy_unwrap2
}

var cache_Data_List_Lazy_unwrap3 gopurs_runtime.Value
var once_Data_List_Lazy_unwrap3 sync.Once
func Get_Data_List_Lazy_unwrap3() gopurs_runtime.Value {
	once_Data_List_Lazy_unwrap3.Do(func() {
		cache_Data_List_Lazy_unwrap3 = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_List_Lazy_unwrap3
}

var cache_Data_List_Lazy_unwrap4 gopurs_runtime.Value
var once_Data_List_Lazy_unwrap4 sync.Once
func Get_Data_List_Lazy_unwrap4() gopurs_runtime.Value {
	once_Data_List_Lazy_unwrap4.Do(func() {
		cache_Data_List_Lazy_unwrap4 = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_List_Lazy_unwrap4
}

var cache_Data_List_Lazy_unwrap5 gopurs_runtime.Value
var once_Data_List_Lazy_unwrap5 sync.Once
func Get_Data_List_Lazy_unwrap5() gopurs_runtime.Value {
	once_Data_List_Lazy_unwrap5.Do(func() {
		cache_Data_List_Lazy_unwrap5 = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_List_Lazy_unwrap5
}

var cache_Data_List_Lazy_unwrap6 gopurs_runtime.Value
var once_Data_List_Lazy_unwrap6 sync.Once
func Get_Data_List_Lazy_unwrap6() gopurs_runtime.Value {
	once_Data_List_Lazy_unwrap6.Do(func() {
		cache_Data_List_Lazy_unwrap6 = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_List_Lazy_unwrap6
}

var cache_Data_List_Lazy_identity gopurs_runtime.Value
var once_Data_List_Lazy_identity sync.Once
func Get_Data_List_Lazy_identity() gopurs_runtime.Value {
	once_Data_List_Lazy_identity.Do(func() {
		cache_Data_List_Lazy_identity = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
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
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
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
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
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
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
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
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
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
				return gopurs_runtime.RecordDict2("init", "rest", orig.go__init, orig.rest)
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
				return gopurs_runtime.RecordDict2("no", "yes", orig.no, orig.yes)
				}()
})
	})
	return cache_Data_List_Lazy_partition
}

var cache_Data_List_Lazy_null gopurs_runtime.Value
var once_Data_List_Lazy_null sync.Once
func Get_Data_List_Lazy_null() gopurs_runtime.Value {
	once_Data_List_Lazy_null.Do(func() {
		cache_Data_List_Lazy_null = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_isNothing(), Get_Data_List_Lazy_uncons())
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
var Call_local_Data_List_Lazy_go__go_0_0_15 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_0_0_15
var go__go_0_0_15 gopurs_runtime.Value
_ = go__go_0_0_15
Call_local_Data_List_Lazy_go__go_0_0_15 = func(b_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_0_0_15:
for {
if false { continue go__go_0_0_15 }
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
// TAST (Let): v_3_1 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_2))
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
continue go__go_0_0_15
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
go__go_0_0_15 = gopurs_runtime.Func(func(b_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_0_0_15(b_1_loop_val, xs_2_loop_val)
})
})
return gopurs_runtime.Apply(go__go_0_0_15, gopurs_runtime.Int(int64(0)))
}()
	})
	return cache_Data_List_Lazy_length
}

var cache_Data_List_Lazy_last gopurs_runtime.Value
var once_Data_List_Lazy_last sync.Once
func Get_Data_List_Lazy_last() gopurs_runtime.Value {
	once_Data_List_Lazy_last.Do(func() {
		cache_Data_List_Lazy_last = func() gopurs_runtime.Value {
var Call_local_Data_List_Lazy_go__go_0_0_16 func(*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
_ = Call_local_Data_List_Lazy_go__go_0_0_16
var go__go_0_0_16 gopurs_runtime.Value
_ = go__go_0_0_16
Call_local_Data_List_Lazy_go__go_0_0_16 = func(v_1_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
go__go_0_0_16:
for {
if false { continue go__go_0_0_16 }
var v_1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v_1_loop
_ = v_1
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v_1 != nil) {
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply(Get_Data_List_Lazy_null(), (v_1).V1).IntVal) != (0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(v_1).V0, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_1
} else {

}
}
{
v_1_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (v_1).V1))
continue go__go_0_0_16
__t1 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("unreachable") }()
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
}
end_branch_2:
return __t2
}
}
go__go_0_0_16 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Lazy_go__go_0_0_16(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_1_loop_val)))}
})
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), go__go_0_0_16, Get_Data_List_Lazy_Types_step())
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
var go__go_0_0_18 gopurs_runtime.Value
_ = go__go_0_0_18
var go__go_0_0_18_cell *gopurs_runtime.Value
_ = go__go_0_0_18_cell
// FALLBACK TCO: isLoop=false len=1
go__go_0_0_18 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_1)
_ = __t_tag_1
if (__t_tag_1 != nil) {
var __t6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply(Get_Data_List_Lazy_null(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1).IntVal) != (0) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{Get_Data_List_Lazy_Types_nil(), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_6
} else {

}
}
{
// TAST (Let): __local_var_2_2 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope177)])])])
__local_var_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply((*go__go_0_0_18_cell), gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)))}))
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
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_5:
__t6 = __t5
}
end_branch_6:
__t7 = __t6
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t7)}
})
go__go_0_0_18_cell = &go__go_0_0_18
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), go__go_0_0_18, Get_Data_List_Lazy_Types_step())
}()
	})
	return cache_Data_List_Lazy_go__init
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
return func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_head(xs_0_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
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
		cache_Data_List_Lazy_fromStep = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_List(), Call_Control_Applicative_pure(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Lazy_applicativeLazy())))
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
return func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_findLastIndex(fn_0_box, xs_1_box)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
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
return Call_Data_List_Lazy_elemLastIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
})
	})
	return cache_Data_List_Lazy_elemLastIndex
}

var cache_Data_List_Lazy_elemIndex gopurs_runtime.Value
var once_Data_List_Lazy_elemIndex sync.Once
func Get_Data_List_Lazy_elemIndex() gopurs_runtime.Value {
	once_Data_List_Lazy_elemIndex.Do(func() {
		cache_Data_List_Lazy_elemIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Lazy_elemIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
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
		cache_Data_List_Lazy_concatMap = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
__local_var_0_0 := Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_bindList()))
_ = __local_var_0_0
return gopurs_runtime.Func2(func(b_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_0_0, a_2, b_1)
})
}()
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
		cache_Data_List_Lazy_catMaybes = Call_Data_List_Lazy_mapMaybe(Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))
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
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope3)])])
__local_var_3_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_1)
_ = __local_var_3_0
// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b$scope4)])])
__local_var_4_1 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), ys_2)
_ = __local_var_4_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_2 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_7_2 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_3_0)
_ = __local_var_7_2
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_7_2.Type == 9 && __local_var_7_2.IntVal == 218341868 && __local_var_7_2.UnsafePtr == nil) {
__t6 = (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_6
} else {

}
}
{
var __t_tag_3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_8)
_ = __t_tag_3
if (__t_tag_3 == nil) {
__t6 = (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_6
} else {

}
}
{
var __t_and_5 bool = false
if (__local_var_7_2.Type == 9 && __local_var_7_2.IntVal == 218341868 && __local_var_7_2.UnsafePtr != nil) {

var __t_tag_4 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_8)
_ = __t_tag_4
__t_and_5 = (__t_tag_4 != nil)
}
if __t_and_5 {
__t6 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_7_2.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_8.UnsafePtr).V0), Call_Data_List_Lazy_zipWith(f_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_7_2.UnsafePtr).V1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_8.UnsafePtr).V1)})
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
})), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_1))
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
// TAST (Let): Apply0_4_0 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m$scope256)])
Apply0_4_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{}))
_ = Apply0_4_0
// TAST (Let): Functor0_5_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope256)])
Functor0_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_1
// TAST (Let): __local_var_6_2 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope258) [(TypeVar a$scope257)])] (TypeApp (TypeVar m$scope258) [(TypeVar a$scope257)]))
__local_var_6_2 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_6_2
var Call_local_Data_List_Lazy_go__go_7_3_0 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_7_3_0
var go__go_7_3_0 gopurs_runtime.Value
_ = go__go_7_3_0
Call_local_Data_List_Lazy_go__go_7_3_0 = func(b_8_loop gopurs_runtime.Value, xs_9_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_7_3_0:
for {
if false { continue go__go_7_3_0 }
var b_8 gopurs_runtime.Value = b_8_loop
_ = b_8
var xs_9 gopurs_runtime.Value = xs_9_loop
_ = xs_9
// TAST (Let): v_10_4 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_10_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_9))
_ = v_10_4
var __t5 gopurs_runtime.Value
{
if (v_10_4 == nil) {
__t5 = b_8
goto end_branch_5
} else {

}
}
{
if (v_10_4 != nil) {
b_8_loop = gopurs_runtime.Apply2(Apply0_4_0.V1, gopurs_runtime.Apply2(Functor0_5_1.V0, Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Apply(__local_var_6_2, (v_10_4).V0)), b_8)
xs_9_loop = (v_10_4).V1
continue go__go_7_3_0
__t5 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_7_3_0 = gopurs_runtime.Func(func(b_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_7_3_0(b_8_loop_val, xs_9_loop_val)
})
})
var Call_local_Data_List_Lazy_go__go_8_6_1 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_8_6_1
var go__go_8_6_1 gopurs_runtime.Value
_ = go__go_8_6_1
Call_local_Data_List_Lazy_go__go_8_6_1 = func(b_9_loop gopurs_runtime.Value, xs_10_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_8_6_1:
for {
if false { continue go__go_8_6_1 }
var b_9 gopurs_runtime.Value = b_9_loop
_ = b_9
var xs_10 gopurs_runtime.Value = xs_10_loop
_ = xs_10
// TAST (Let): v_11_7 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_11_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_10))
_ = v_11_7
var __t8 gopurs_runtime.Value
{
if (v_11_7 == nil) {
__t8 = b_9
goto end_branch_8
} else {

}
}
{
if (v_11_7 != nil) {
b_9_loop = Call_Data_List_Lazy_Types_cons((v_11_7).V0, b_9)
xs_10_loop = (v_11_7).V1
continue go__go_8_6_1
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
go__go_8_6_1 = gopurs_runtime.Func(func(b_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_8_6_1(b_9_loop_val, xs_10_loop_val)
})
})
// TAST (Let): __local_var_9_9 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope18)])])
__local_var_9_9 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_2)
_ = __local_var_9_9
// TAST (Let): __local_var_10_10 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b$scope4)])])
__local_var_10_10 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), ys_3)
_ = __local_var_10_10
return Call_local_Data_List_Lazy_go__go_7_3_0(gopurs_runtime.Apply(dictApplicative_0.V1, Get_Data_List_Lazy_Types_nil()), Call_local_Data_List_Lazy_go__go_8_6_1(Get_Data_List_Lazy_Types_nil(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_11 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_13_11 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_9_9)
_ = __local_var_13_11
return gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_13_11.Type == 9 && __local_var_13_11.IntVal == 218341868 && __local_var_13_11.UnsafePtr == nil) {
__t15 = (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_15
} else {

}
}
{
var __t_tag_12 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_14)
_ = __t_tag_12
if (__t_tag_12 == nil) {
__t15 = (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_15
} else {

}
}
{
var __t_and_14 bool = false
if (__local_var_13_11.Type == 9 && __local_var_13_11.IntVal == 218341868 && __local_var_13_11.UnsafePtr != nil) {

var __t_tag_13 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_14)
_ = __t_tag_13
__t_and_14 = (__t_tag_13 != nil)
}
if __t_and_14 {
__t15 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_13_11.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_14.UnsafePtr).V0), Call_Data_List_Lazy_zipWith(f_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_13_11.UnsafePtr).V1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_14.UnsafePtr).V1)})
goto end_branch_15
} else {

}
}
{
__t15 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_15:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t15)}
})
})), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_10_10))
}))))
}

func Call_Data_List_Lazy_zip(xs_0_loop gopurs_runtime.Value, ys_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var ys_1 gopurs_runtime.Value = ys_1_loop
_ = ys_1
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope23)])])
__local_var_2_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_0)
_ = __local_var_2_0
// TAST (Let): __local_var_3_1 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b$scope4)])])
__local_var_3_1 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), ys_1)
_ = __local_var_3_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Lazy_force(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_6_2 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_2_0)
_ = __local_var_6_2
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr == nil) {
__t6 = (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_6
} else {

}
}
{
var __t_tag_3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_7)
_ = __t_tag_3
if (__t_tag_3 == nil) {
__t6 = (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_6
} else {

}
}
{
var __t_and_5 bool = false
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr != nil) {

var __t_tag_4 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_7)
_ = __t_tag_4
__t_and_5 = (__t_tag_4 != nil)
}
if __t_and_5 {
__t6 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_7.UnsafePtr).V0}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))}, Call_Data_List_Lazy_zipWith(Get_Data_Tuple_Tuple(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_7.UnsafePtr).V1)})
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
})), gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_3_1))
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
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope28)])])
__local_var_3_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_2)
_ = __local_var_3_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_1 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_5_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_3_0)
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
if (n_0) == (int64(0)) {
__t2 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V1})
goto end_branch_2
} else {

}
}
{
__t2 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V0, Call_Data_List_Lazy_updateAt((n_0) - (int64(1)), x_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V1)})
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
// TAST (Let): v_4_1 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_3))
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
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
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
// TAST (Let): v_5_8 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_5_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_4))
_ = v_5_8
var __t9 gopurs_runtime.Value
{
if (v_5_8 == nil) {
__t9 = b_3
goto end_branch_9
} else {

}
}
{
if (v_5_8 != nil) {
b_3_loop = Call_Data_List_Lazy_Types_cons((v_5_8).V0, b_3)
xs_4_loop = (v_5_8).V1
continue go__go_2_7_3
__t9 = func() gopurs_runtime.Value { panic("unreachable") }()
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
go__go_2_7_3 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_2_7_3(b_3_loop_val, xs_4_loop_val)
})
})
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
				_v := Call_local_Data_List_Lazy_go__go_1_0_2(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{Get_Data_List_Lazy_Types_nil(), Get_Data_List_Lazy_Types_nil()}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))}, Call_local_Data_List_Lazy_go__go_2_7_3(Get_Data_List_Lazy_Types_nil(), xs_0))
				_p := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{V0: _p.V0, V1: _p.V1}
			}()
}

func Call_Data_List_Lazy_uncons(xs_0_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
// TAST (Let): v_1_0 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope43)])
v_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_0))
_ = v_1_0
var __t1 *Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail gopurs_runtime.Value
}]
{
if (v_1_0 == nil) {
__t1 = Rebox_Data_List_Lazy_3094389156_3833657837(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_1
} else {

}
}
{
if (v_1_0 != nil) {
__t1 = Rebox_Data_List_Lazy_3094389156_3833657837(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.RecordDict2("head", "tail", (v_1_0).V0, (v_1_0).V1), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
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
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_List_Lazy_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
return gopurs_runtime.Apply(dictUnfoldable_0.V1, gopurs_runtime.Func(func(xs_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: (TypeVar a$scope49), tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope49)])])] Empty))])
__local_var_2_0 := Rebox_Data_List_Lazy_3094389156_3833657837(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_uncons(xs_1)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
_ = __local_var_2_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_2_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(__local_var_2_0).V0.head, (__local_var_2_0).V0.tail}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
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
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_List(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_1)
_ = __t_tag_0
if ((__t_tag_0 != nil)) && ((gopurs_runtime.Apply(p_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0).IntVal) != (0)) {
__t1 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0, gopurs_runtime.Apply(Call_Data_List_Lazy_takeWhile(p_0), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t1)}
})), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))))
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
if (n_0) <= (int64(0)) {
__t3 = gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_List_Lazy_Types_nil()
})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_List(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0, gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_1)
_ = __t_tag_0
if (__t_tag_0 == nil) {
__t2 = (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_1)
_ = __t_tag_1
if (__t_tag_1 != nil) {
__t2 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_1.UnsafePtr).V0, gopurs_runtime.Apply(Call_Data_List_Lazy_take((n_0) - (int64(1))), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_1.UnsafePtr).V1)})
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t2)}
})), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))))
}
end_branch_3:
return __t3
}
}

func Call_Data_List_Lazy_tail(xs_0_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope69)])]), head: (TypeVar a$scope69)] Empty))])
__local_var_1_0 := Rebox_Data_List_Lazy_3094389156_3833657837(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_uncons(xs_0)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
_ = __local_var_1_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_1_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(__local_var_1_0).V0.tail, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_1:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Rebox_Data_List_Lazy_4130553207_1542299734(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](Get_Control_Monad_Rec_Class_monadRecMaybe())).V1, gopurs_runtime.Func(func(o_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_4_0 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope71)])
v1_4_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), gopurs_runtime.RecordGet(o_3, "a")))
_ = v1_4_0
var __t3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v1_4_0 == nil) {
__t3 = (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_1239976062_3603546092((&Constructor_Control_Monad_Rec_Class_Done[struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(o_3, "b")})))}})
goto end_branch_3
} else {

}
}
{
if (v1_4_0 != nil) {
// TAST (Let): v2_5_1 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope71)])
v2_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), gopurs_runtime.RecordGet(o_3, "b")))
_ = v2_5_1
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((v2_5_1 != nil)) && ((gopurs_runtime.Apply2(dictEq_0.V0, (v1_4_0).V0, (v2_5_1).V0).IntVal) != (0)) {
__t2 = (&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_784458114_4008603408((&Constructor_Control_Monad_Rec_Class_Loop[struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}, gopurs_runtime.Value]{1, func() struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
} {
					orig := gopurs_runtime.RecordDict2("a", "b", (v1_4_0).V1, (v2_5_1).V1)
					_ = orig
					clone := struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}{}
					clone.a = gopurs_runtime.RecordGet(orig, "a")
					clone.b = gopurs_runtime.RecordGet(orig, "b")
					return clone
				}()})))}})
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
__t3 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
}), func() gopurs_runtime.Value {
				orig := struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}{v_1, s_2}
				_ = orig
				return gopurs_runtime.RecordDict2("a", "b", orig.a, orig.b)
				}())))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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
// TAST (Let): v_2_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: (TypeVar a$scope80), tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope80)])])] Empty))])
v_2_0 := Rebox_Data_List_Lazy_3094389156_3833657837(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_uncons(xs_1)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
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
// TAST (Let): v1_4_2 shape=App(Var) bindingType=(Record (Row [init: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope80)])]), rest: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope80)])])] Empty))
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
var Call_local_Data_List_Lazy_go__go_2_0_4 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_2_0_4
var go__go_2_0_4 gopurs_runtime.Value
_ = go__go_2_0_4
Call_local_Data_List_Lazy_go__go_2_0_4 = func(b_3_loop gopurs_runtime.Value, xs_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_2_0_4:
for {
if false { continue go__go_2_0_4 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var xs_4 gopurs_runtime.Value = xs_4_loop
_ = xs_4
// TAST (Let): v_5_1 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_4))
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
b_3_loop = Call_Data_List_Lazy_Types_cons((v_5_1).V0, b_3)
xs_4_loop = (v_5_1).V1
continue go__go_2_0_4
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
go__go_2_0_4 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_2_0_4(b_3_loop_val, xs_4_loop_val)
})
})
var Call_local_Data_List_Lazy_go__go_3_3_5 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_3_3_5
var go__go_3_3_5 gopurs_runtime.Value
_ = go__go_3_3_5
Call_local_Data_List_Lazy_go__go_3_3_5 = func(b_4_loop gopurs_runtime.Value, xs_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_3_3_5:
for {
if false { continue go__go_3_3_5 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_4 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_5))
_ = v_6_4
var __t5 gopurs_runtime.Value
{
if (v_6_4 == nil) {
__t5 = b_4
goto end_branch_5
} else {

}
}
{
if (v_6_4 != nil) {
b_4_loop = Call_Data_List_Lazy_Types_cons((v_6_4).V0, b_4)
xs_5_loop = (v_6_4).V1
continue go__go_3_3_5
__t5 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_3_3_5 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_3_3_5(b_4_loop_val, xs_5_loop_val)
})
})
return Call_local_Data_List_Lazy_go__go_2_0_4(gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_1, Get_Data_List_Lazy_Types_nil()}))}
})), Call_local_Data_List_Lazy_go__go_3_3_5(Get_Data_List_Lazy_Types_nil(), xs_0))
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
// TAST (Let): showList_1_0 shape=App(Var) bindingType=(ADT ["Data","Show","Show"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope87)])])])
showList_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Call_Data_List_Lazy_Types_showList(dictShow_0))
_ = showList_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Pattern ") + (gopurs_runtime.Apply(showList_1_0.V0, v_2).StrVal())) + (")"))
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
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope91)])])
__local_var_3_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_2)
_ = __local_var_3_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_1 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_5_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_3_0)
_ = __local_var_5_1
var __t3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_5_1.Type == 9 && __local_var_5_1.IntVal == 218341868 && __local_var_5_1.UnsafePtr == nil) {
__t3 = (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_3
} else {

}
}
{
if (__local_var_5_1.Type == 9 && __local_var_5_1.IntVal == 218341868 && __local_var_5_1.UnsafePtr != nil) {
// TAST (Let): acc_prime__6_2 shape=App(Other) bindingType=(TypeVar b$scope92)
acc_prime__6_2 := gopurs_runtime.Apply2(f_0, acc_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V0)
_ = acc_prime__6_2
__t3 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, acc_prime__6_2, Call_Data_List_Lazy_scanlLazy(f_0, acc_prime__6_2, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V1)})
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

func Call_Data_List_Lazy_reverse(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_step(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_List_Lazy_go__go_2_0_6 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_2_0_6
var go__go_2_0_6 gopurs_runtime.Value
_ = go__go_2_0_6
Call_local_Data_List_Lazy_go__go_2_0_6 = func(b_3_loop gopurs_runtime.Value, xs_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_2_0_6:
for {
if false { continue go__go_2_0_6 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var xs_4 gopurs_runtime.Value = xs_4_loop
_ = xs_4
// TAST (Let): v_5_1 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_4))
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
b_3_loop = Call_Data_List_Lazy_Types_cons((v_5_1).V0, b_3)
xs_4_loop = (v_5_1).V1
continue go__go_2_0_6
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
go__go_2_0_6 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_2_0_6(b_3_loop_val, xs_4_loop_val)
})
})
return Call_local_Data_List_Lazy_go__go_2_0_6(Get_Data_List_Lazy_Types_nil(), xs_0)
})))
}

func Call_Data_List_Lazy_replicateM(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
replicateM:
for {
if false { continue replicateM }
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m$scope102)])
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m$scope102)])
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func2(func(n_3 gopurs_runtime.Value, m_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (n_3.IntVal) < (Call_Data_Semiring_one(gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_348932501_2826095630(Rebox_Data_List_Lazy_2826095630_348932501(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringInt()))))}).IntVal) {
__t2 = gopurs_runtime.Apply(Applicative0_1_0.V1, Get_Data_List_Lazy_Types_nil())
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(Bind1_2_1.V1, m_4, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(Call_Data_List_Lazy_replicateM(dictMonad_0), gopurs_runtime.Int((n_3.IntVal) - (Call_Data_Semiring_one(gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_348932501_2826095630(Rebox_Data_List_Lazy_2826095630_348932501(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringInt()))))}).IntVal)), m_4), gopurs_runtime.Func(func(as_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_0.V1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
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
var go__go_1_0_7 gopurs_runtime.Value
_ = go__go_1_0_7
var go__go_1_0_7_cell *gopurs_runtime.Value
_ = go__go_1_0_7_cell
// FALLBACK TCO: isLoop=false len=1
go__go_1_0_7 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_step(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_0, (*go__go_1_0_7_cell)}))}
}))
})))
go__go_1_0_7_cell = &go__go_1_0_7
return go__go_1_0_7
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
var go__go_2_6_9 gopurs_runtime.Value
_ = go__go_2_6_9
var go__go_2_6_9_cell *gopurs_runtime.Value
_ = go__go_2_6_9_cell
// FALLBACK TCO: isLoop=false len=1
go__go_2_6_9 = gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_step(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_6_7 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope276), (TypeVar b$scope277)])])
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
// TAST (Let): __local_var_8_9 shape=App(Other) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope276)])])
__local_var_8_9 := gopurs_runtime.Apply2((*go__go_2_6_9_cell), f_3, ((v1_6_7).V0).V1)
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
return __t10
})))
})
go__go_2_6_9_cell = &go__go_2_6_9
__t12 = gopurs_runtime.Apply2(go__go_2_6_9, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, int64]]
{
if (x_3.IntVal) >= (end_1) {
__t11 = Rebox_Data_List_Lazy_3094389156_1415037225(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_3363075976_138441832(Rebox_Data_List_Lazy_138441832_3363075976(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(x_3.IntVal), gopurs_runtime.Int((x_3.IntVal) - (int64(1)))}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_11
} else {

}
}
{
__t11 = Rebox_Data_List_Lazy_3094389156_1415037225(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_1415037225_3094389156(__t11))}
}), gopurs_runtime.Int(start_0))
goto end_branch_12
} else {

}
}
{
var go__go_2_0_8 gopurs_runtime.Value
_ = go__go_2_0_8
var go__go_2_0_8_cell *gopurs_runtime.Value
_ = go__go_2_0_8_cell
// FALLBACK TCO: isLoop=false len=1
go__go_2_0_8 = gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_step(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v1_6_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a$scope276), (TypeVar b$scope277)])])
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
// TAST (Let): __local_var_8_3 shape=App(Other) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope276)])])
__local_var_8_3 := gopurs_runtime.Apply2((*go__go_2_0_8_cell), f_3, ((v1_6_1).V0).V1)
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
return __t4
})))
})
go__go_2_0_8_cell = &go__go_2_0_8
__t12 = gopurs_runtime.Apply2(go__go_2_0_8, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, int64]]
{
if (x_3.IntVal) <= (end_1) {
__t5 = Rebox_Data_List_Lazy_3094389156_1415037225(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_3363075976_138441832(Rebox_Data_List_Lazy_138441832_3363075976(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Int(x_3.IntVal), gopurs_runtime.Int((x_3.IntVal) + (int64(1)))}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))))}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_5
} else {

}
}
{
__t5 = Rebox_Data_List_Lazy_3094389156_1415037225(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_1415037225_3094389156(__t5))}
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
var Call_local_Data_List_Lazy_go__go_2_0_10 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_2_0_10
var go__go_2_0_10 gopurs_runtime.Value
_ = go__go_2_0_10
Call_local_Data_List_Lazy_go__go_2_0_10 = func(b_3_loop gopurs_runtime.Value, xs_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_2_0_10:
for {
if false { continue go__go_2_0_10 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var xs_4 gopurs_runtime.Value = xs_4_loop
_ = xs_4
// TAST (Let): v_5_1 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_4))
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
				return gopurs_runtime.RecordDict2("no", "yes", orig.no, orig.yes)
				}()
xs_4_loop = (v_5_1).V1
continue go__go_2_0_10
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
go__go_2_0_10 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_2_0_10(b_3_loop_val, xs_4_loop_val)
})
})
var Call_local_Data_List_Lazy_go__go_3_4_11 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_3_4_11
var go__go_3_4_11 gopurs_runtime.Value
_ = go__go_3_4_11
Call_local_Data_List_Lazy_go__go_3_4_11 = func(b_4_loop gopurs_runtime.Value, xs_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_3_4_11:
for {
if false { continue go__go_3_4_11 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_5 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_5))
_ = v_6_5
var __t6 gopurs_runtime.Value
{
if (v_6_5 == nil) {
__t6 = b_4
goto end_branch_6
} else {

}
}
{
if (v_6_5 != nil) {
b_4_loop = Call_Data_List_Lazy_Types_cons((v_6_5).V0, b_4)
xs_5_loop = (v_6_5).V1
continue go__go_3_4_11
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
go__go_3_4_11 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_3_4_11(b_4_loop_val, xs_5_loop_val)
})
})
return func() struct{
	no gopurs_runtime.Value
	yes gopurs_runtime.Value
} {
					orig := Call_local_Data_List_Lazy_go__go_2_0_10(func() gopurs_runtime.Value {
				orig := struct{
	no gopurs_runtime.Value
	yes gopurs_runtime.Value
}{Get_Data_List_Lazy_Types_nil(), Get_Data_List_Lazy_Types_nil()}
				_ = orig
				return gopurs_runtime.RecordDict2("no", "yes", orig.no, orig.yes)
				}(), Call_local_Data_List_Lazy_go__go_3_4_11(Get_Data_List_Lazy_Types_nil(), xs_1))
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

func Call_Data_List_Lazy_nubBy(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var goStep_1_0_12 gopurs_runtime.Value
_ = goStep_1_0_12
var goStep_1_0_12_cell *gopurs_runtime.Value
_ = goStep_1_0_12_cell
// FALLBACK TCO: isLoop=false len=2
var go__go_1_1_13 gopurs_runtime.Value
_ = go__go_1_1_13
var go__go_1_1_13_cell *gopurs_runtime.Value
_ = go__go_1_1_13_cell
// FALLBACK TCO: isLoop=false len=2
goStep_1_0_12 = gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_3)
_ = __t_tag_2
if (__t_tag_2 == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_6
} else {

}
}
{
var __t_tag_3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_3)
_ = __t_tag_3
if (__t_tag_3 != nil) {
// TAST (Let): v2_4_4 shape=App(Var) bindingType=(Record (Row [found: Boolean, result: (ADT ["Data","List","Internal","Set"] [(TypeVar a$scope127)])] Empty))
v2_4_4 := Call_Data_List_Internal_insertAndLookupBy(p_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, v_2)
_ = v2_4_4
var __t5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if v2_4_4.found {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), gopurs_runtime.Apply2((*go__go_1_1_13_cell), v2_4_4.result, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)))
goto end_branch_5
} else {

}
}
{
__t5 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, gopurs_runtime.Apply2((*go__go_1_1_13_cell), v2_4_4.result, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})
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
goStep_1_0_12_cell = &goStep_1_0_12
go__go_1_1_13 = gopurs_runtime.Func2(func(s_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_7 shape=App(Other) bindingType=(Func [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope127)])] (ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope127)]))
__local_var_4_7 := gopurs_runtime.Apply((*goStep_1_0_12_cell), s_2)
_ = __local_var_4_7
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_7, gopurs_runtime.Apply(Get_Data_Lazy_force(), v_3))
}))
})
go__go_1_1_13_cell = &go__go_1_1_13
return gopurs_runtime.Apply(go__go_1_1_13, Get_Data_List_Internal_emptySet())
}

func Call_Data_List_Lazy_nub(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return Call_Data_List_Lazy_nubBy(Call_Data_Ord_compare(dictOrd_0))
}

func Call_Data_List_Lazy_mapMaybe(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
mapMaybe:
for {
if false { continue mapMaybe }
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var Call_local_Data_List_Lazy_go__go_1_0_14 func(*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Lazy_go__go_1_0_14
var go__go_1_0_14 gopurs_runtime.Value
_ = go__go_1_0_14
Call_local_Data_List_Lazy_go__go_1_0_14 = func(v_2_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] {
go__go_1_0_14:
for {
if false { continue go__go_1_0_14 }
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
// TAST (Let): v1_3_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar b$scope141)])
v1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, (v_2).V0))
_ = v1_3_1
var __t2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (v1_3_1 == nil) {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (v_2).V1))
continue go__go_1_0_14
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
go__go_1_0_14 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Lazy_go__go_1_0_14(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_2_loop_val)))}
})
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_List(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0, go__go_1_0_14), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))))
}
}

func Call_Data_List_Lazy_some(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): Apply0_1_0 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar f$scope151)])
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V0, gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f$scope151)])
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V1, gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func2(func(dictLazy_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, Get_Data_List_Lazy_Types_cons(), v_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_3, "defer"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Data_List_Lazy_many(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dictLazy_3))}, v_4)
})))
})
}

func Call_Data_List_Lazy_many(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): Alt0_1_0 shape=App(Other) bindingType=(ADT ["Control","Alt","Alt"] [(TypeVar f$scope157)])
Alt0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V1, gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}))
_ = Alt0_1_0
// TAST (Let): Applicative0_2_1 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar f$scope157)])
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictAlternative_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func2(func(dictLazy_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Alt0_1_0.V1, gopurs_runtime.Apply2(Call_Data_List_Lazy_some(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dictLazy_3))}, v_4), gopurs_runtime.Apply(Applicative0_2_1.V1, Get_Data_List_Lazy_Types_nil()))
})
}

func Call_Data_List_Lazy_iterate(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var go__go_2_0_17 gopurs_runtime.Value
_ = go__go_2_0_17
var go__go_2_0_17_cell *gopurs_runtime.Value
_ = go__go_2_0_17_cell
// FALLBACK TCO: isLoop=false len=1
go__go_2_0_17 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_step(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope114)])])
__local_var_4_2 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), (*go__go_2_0_17_cell))
_ = __local_var_4_2
// TAST (Let): __local_var_4_1 shape=Let(App(Var)) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope169)])])
__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_3 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_6_3 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_2)
_ = __local_var_6_3
var __t4 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_6_3.Type == 9 && __local_var_6_3.IntVal == 218341868 && __local_var_6_3.UnsafePtr == nil) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_4
} else {

}
}
{
if (__local_var_6_3.Type == 9 && __local_var_6_3.IntVal == 218341868 && __local_var_6_3.UnsafePtr != nil) {
__t4 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_functorList()).V0, f_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_3.UnsafePtr).V1)})
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
_ = __local_var_4_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_1, __local_var_4_1}))}
}))
})))
go__go_2_0_17_cell = &go__go_2_0_17
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
var __t3 gopurs_runtime.Value
{
if (v_0) == (int64(0)) {
__t3 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, v1_1, v2_2}))}
}))
goto end_branch_3
} else {

}
}
{
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope171)])])
__local_var_3_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), v2_2)
_ = __local_var_3_0
__t3 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_1 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_5_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_3_0)
_ = __local_var_5_1
var __t2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_5_1.Type == 9 && __local_var_5_1.IntVal == 218341868 && __local_var_5_1.UnsafePtr == nil) {
__t2 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, v1_1, Get_Data_List_Lazy_Types_nil()})
goto end_branch_2
} else {

}
}
{
if (__local_var_5_1.Type == 9 && __local_var_5_1.IntVal == 218341868 && __local_var_5_1.UnsafePtr != nil) {
__t2 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V0, Call_Data_List_Lazy_insertAt((v_0) - (int64(1)), v1_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V1)})
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
end_branch_3:
return __t3
}
}

func Call_Data_List_Lazy_index(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var Call_local_Data_List_Lazy_go__go_1_0_19 func(*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value], int64) *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
_ = Call_local_Data_List_Lazy_go__go_1_0_19
var go__go_1_0_19 gopurs_runtime.Value
_ = go__go_1_0_19
Call_local_Data_List_Lazy_go__go_1_0_19 = func(v_2_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value], v1_3_loop int64) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
go__go_1_0_19:
for {
if false { continue go__go_1_0_19 }
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
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_1
} else {

}
}
{
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (v_2).V1))
v1_3_loop = (v1_3) - (int64(1))
continue go__go_1_0_19
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
go__go_1_0_19 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Lazy_go__go_1_0_19(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_2_loop_val), v1_3_loop_val.IntVal))}
})
})
return gopurs_runtime.Apply(go__go_1_0_19, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_0)))})
}

func Call_Data_List_Lazy_head(xs_0_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: (TypeVar a$scope186), tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope186)])])] Empty))])
__local_var_1_0 := Rebox_Data_List_Lazy_3094389156_3833657837(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_uncons(xs_0)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
_ = __local_var_1_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_1_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(__local_var_1_0).V0.head, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_1:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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
// TAST (Let): v_1_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope188)])]), tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(ADT ["Data","List","Lazy","Types","List"] [(TypeVar a$scope188)])])])] Empty))])
v_1_0 := Rebox_Data_List_Lazy_3094389156_3833657837(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_uncons(xs_0)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
_ = v_1_0
var __t8 gopurs_runtime.Value
{
if (v_1_0 == nil) {
__t8 = xs_0
goto end_branch_8
} else {

}
}
{
if (v_1_0 != nil) {
// TAST (Let): v1_2_1 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: (TypeVar a$scope188), tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope188)])])] Empty))])
v1_2_1 := Rebox_Data_List_Lazy_3094389156_3833657837(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_uncons((v_1_0).V0.head)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
_ = v1_2_1
var __t7 gopurs_runtime.Value
{
if (v1_2_1 == nil) {
xs_0_loop = (v_1_0).V0.tail
continue transpose
__t7 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_7
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
// TAST (Let): __local_var_6_6 shape=App(Var) bindingType=Any
__local_var_6_6 := gopurs_runtime.Apply(Call_Data_List_Lazy_mapMaybe(Get_Data_List_Lazy_tail()), (v_1_0).V0.tail)
_ = __local_var_6_6
// TAST (Let): __local_var_6_5 shape=App(Var) bindingType=Any
__local_var_6_5 := Call_Data_List_Lazy_transpose(gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_4_3, __local_var_6_6}))}
})))
_ = __local_var_6_5
__t7 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_3_2, __local_var_5_4}))}
})), __local_var_6_5}))}
}))
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
__t8 = __t7
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

func Call_Data_List_Lazy_groupBy(eq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
groupBy:
for {
if false { continue groupBy }
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_List(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_1)
_ = __t_tag_0
if (__t_tag_0 == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_5
} else {

}
}
{
var __t_tag_1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_1)
_ = __t_tag_1
if (__t_tag_1 != nil) {
// TAST (Let): __local_var_2_2 shape=Other bindingType=Any
__local_var_2_2 := (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0
_ = __local_var_2_2
// TAST (Let): v1_3_3 shape=App(Var) bindingType=(Record (Row [init: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope192)])]), rest: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope192)])])] Empty))
v1_3_3 := Call_Data_List_Lazy_span(gopurs_runtime.Apply(eq_0, __local_var_2_2), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1)
_ = v1_3_3
// TAST (Let): __local_var_4_4 shape=Other bindingType=Any
__local_var_4_4 := v1_3_3.go__init
_ = __local_var_4_4
__t5 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_2_2, __local_var_4_4}))}
})), gopurs_runtime.Apply(Call_Data_List_Lazy_groupBy(eq_0), v1_3_3.rest)})
goto end_branch_5
} else {

}
}
{
__t5 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t5)}
})), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))))
}
}

func Call_Data_List_Lazy_group(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return Call_Data_List_Lazy_groupBy(Call_Data_Eq_eq(dictEq_0))
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
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope210)])])
__local_var_3_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_2)
_ = __local_var_3_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_1 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_5_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_3_0)
_ = __local_var_5_1
var __t4 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_5_1.Type == 9 && __local_var_5_1.IntVal == 218341868 && __local_var_5_1.UnsafePtr == nil) {
__t4 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_1, Get_Data_List_Lazy_Types_nil()})
goto end_branch_4
} else {

}
}
{
if (__local_var_5_1.Type == 9 && __local_var_5_1.IntVal == 218341868 && __local_var_5_1.UnsafePtr != nil) {
var __t3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_2 uint32 = uint32(gopurs_runtime.Apply2(cmp_0, x_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V0).IntVal)
_ = __t_tag_2
if (uint32(__t_tag_2) == 380165415) {
__t3 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V0, Call_Data_List_Lazy_insertBy(cmp_0, x_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V1)})
goto end_branch_3
} else {

}
}
{
__t3 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_1, gopurs_runtime.Apply(Get_Data_List_Lazy_fromStep(), gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](__local_var_5_1))})})
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

func Call_Data_List_Lazy_insert(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope217), (TypeVar a$scope217)] (ADT ["Data","Ordering","Ordering"] []))
__local_var_1_0 := Call_Data_Ord_compare(dictOrd_0)
_ = __local_var_1_0
return gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope217)])])
__local_var_4_1 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_3)
_ = __local_var_4_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_6_2 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_1)
_ = __local_var_6_2
var __t5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr == nil) {
__t5 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_2, Get_Data_List_Lazy_Types_nil()})
goto end_branch_5
} else {

}
}
{
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr != nil) {
var __t4 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_1_0, x_2, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V0)
_ = __t_tag_3
if (uint32(__t_tag_3.IntVal) == 380165415) {
__t4 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V0, Call_Data_List_Lazy_insertBy(__local_var_1_0, x_2, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V1)})
goto end_branch_4
} else {

}
}
{
__t4 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, x_2, gopurs_runtime.Apply(Get_Data_List_Lazy_fromStep(), gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](__local_var_6_2))})})
}
end_branch_4:
__t5 = __t4
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
})
}

func Call_Data_List_Lazy_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply2(dictFoldable_0.V2, Get_Data_List_Lazy_Types_cons(), Get_Data_List_Lazy_Types_nil())
}

func Call_Data_List_Lazy_foldrLazy(dictLazy_0_loop *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value], op_1_loop gopurs_runtime.Value, z_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictLazy_0 *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value] = dictLazy_0_loop
_ = dictLazy_0
var op_1 gopurs_runtime.Value = op_1_loop
_ = op_1
var z_2 gopurs_runtime.Value = z_2_loop
_ = z_2
var go__go_3_0_20 gopurs_runtime.Value
_ = go__go_3_0_20
var go__go_3_0_20_cell *gopurs_runtime.Value
_ = go__go_3_0_20_cell
// FALLBACK TCO: isLoop=false len=1
go__go_3_0_20 = gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_5_1 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope224)])
v_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_4))
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
__t4 = gopurs_runtime.Apply(dictLazy_0.V0, gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_1, __local_var_6_2, gopurs_runtime.Apply((*go__go_3_0_20_cell), __local_var_7_3))
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
go__go_3_0_20_cell = &go__go_3_0_20
return go__go_3_0_20
}

func Call_Data_List_Lazy_foldM(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
foldM:
for {
if false { continue foldM }
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m$scope229)])
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m$scope229)])
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value, xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_6_2 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: (TypeVar a$scope230), tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope230)])])] Empty))])
v_6_2 := Rebox_Data_List_Lazy_3094389156_3833657837(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_uncons(xs_5)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
_ = v_6_2
var __t4 gopurs_runtime.Value
{
if (v_6_2 == nil) {
__t4 = gopurs_runtime.Apply(Applicative0_1_0.V1, b_4)
goto end_branch_4
} else {

}
}
{
if (v_6_2 != nil) {
// TAST (Let): __local_var_7_3 shape=Other bindingType=Any
__local_var_7_3 := (v_6_2).V0.tail
_ = __local_var_7_3
__t4 = gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(f_3, b_4, (v_6_2).V0.head), gopurs_runtime.Func(func(b_prime__8 gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_Data_List_Lazy_findIndex(fn_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var Call_local_Data_List_Lazy_go__2988133926_1_0_21 func(int64, gopurs_runtime.Value) *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
_ = Call_local_Data_List_Lazy_go__2988133926_1_0_21
var go__2988133926_1_0_21 gopurs_runtime.Value
_ = go__2988133926_1_0_21
var Call_local_Data_List_Lazy_go__go_1_1_22 func(int64, gopurs_runtime.Value) *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
_ = Call_local_Data_List_Lazy_go__go_1_1_22
var go__go_1_1_22 gopurs_runtime.Value
_ = go__go_1_1_22
Call_local_Data_List_Lazy_go__2988133926_1_0_21 = func(n_2_loop int64, list_3_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
go__2988133926_1_0_21:
for {
if false { continue go__2988133926_1_0_21 }
var n_2 int64 = n_2_loop
_ = n_2
var list_3 gopurs_runtime.Value = list_3_loop
_ = list_3
// TAST (Let): __local_var_4_2 shape=App(Var) bindingType=Any
__local_var_4_2 := Call_Data_List_Lazy_uncons(list_3)
_ = __local_var_4_2
var __t4 gopurs_runtime.Value
{
if __local_var_4_2.V1 {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(fn_0, gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(func() gopurs_runtime.Value {
				_v := __local_var_4_2
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}().UnsafePtr).V0, "head")).IntVal) != (0) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int(n_2), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))}
goto end_branch_3
} else {

}
}
{
n_2_loop = (n_2) + (int64(1))
list_3_loop = gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(func() gopurs_runtime.Value {
				_v := __local_var_4_2
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}().UnsafePtr).V0, "tail")
continue go__2988133926_1_0_21
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_1170268447_3094389156(func() *Constructor_Data_Maybe_Just[int64] { panic("unreachable") }()))}
}
end_branch_3:
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_1170268447_3094389156(Rebox_Data_List_Lazy_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t3))))}
goto end_branch_4
} else {

}
}
{
if (!__local_var_4_2.V1) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }())}
}
end_branch_4:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t4)
}
}
go__2988133926_1_0_21 = gopurs_runtime.Func(func(n_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(list_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Lazy_go__2988133926_1_0_21(n_2_loop_val.IntVal, list_3_loop_val))}
})
})
Call_local_Data_List_Lazy_go__go_1_1_22 = func(n_2_loop int64, list_3_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
go__go_1_1_22:
for {
if false { continue go__go_1_1_22 }
var n_2 int64 = n_2_loop
_ = n_2
var list_3 gopurs_runtime.Value = list_3_loop
_ = list_3
// TAST (Let): __local_var_4_5 shape=App(Var) bindingType=Any
__local_var_4_5 := Call_Data_List_Lazy_uncons(list_3)
_ = __local_var_4_5
var __t7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if __local_var_4_5.V1 {
var __t6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply(fn_0, gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(func() gopurs_runtime.Value {
				_v := __local_var_4_5
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}().UnsafePtr).V0, "head")).IntVal) != (0) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int(n_2), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_6
} else {

}
}
{
__t6 = Call_local_Data_List_Lazy_go__2988133926_1_0_21((n_2) + (int64(1)), gopurs_runtime.RecordGet((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(func() gopurs_runtime.Value {
				_v := __local_var_4_5
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}().UnsafePtr).V0, "tail"))
}
end_branch_6:
__t7 = __t6
goto end_branch_7
} else {

}
}
{
if (!__local_var_4_5.V1) {
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_7
} else {

}
}
{
__t7 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}
}
go__go_1_1_22 = gopurs_runtime.Func(func(n_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(list_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Lazy_go__go_1_1_22(n_2_loop_val.IntVal, list_3_loop_val))}
})
})
return gopurs_runtime.Apply(go__2988133926_1_0_21, gopurs_runtime.Int(int64(0)))
}

func Call_Data_List_Lazy_findLastIndex(fn_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply(Call_Data_List_Lazy_findIndex(fn_0), Call_Data_List_Lazy_reverse(xs_1))
_ = __local_var_2_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136 && __local_var_2_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Int(((gopurs_runtime.Apply(Get_Data_List_Lazy_length(), xs_1).IntVal) - (int64(1))) - ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_2_0.UnsafePtr).V0.IntVal)), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_1:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
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
// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m$scope242)])
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m$scope242)])
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func2(func(p_3 gopurs_runtime.Value, list_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_5_2 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Record (Row [head: (TypeVar a$scope241), tail: (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope241)])])] Empty))])
v_5_2 := Rebox_Data_List_Lazy_3094389156_3833657837(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Data_List_Lazy_uncons(list_4)
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
_ = v_5_2
var __t6 gopurs_runtime.Value
{
if (v_5_2 == nil) {
__t6 = gopurs_runtime.Apply(Applicative0_1_0.V1, Get_Data_List_Lazy_Types_nil())
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
__t6 = gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply(p_3, __local_var_6_3), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(Call_Data_List_Lazy_filterM(dictMonad_0), p_3, __local_var_7_4), gopurs_runtime.Func(func(xs_prime__9 gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.Apply(Applicative0_1_0.V1, __t5)
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
var Call_local_Data_List_Lazy_go__go_1_0_23 func(*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Lazy_go__go_1_0_23
var go__go_1_0_23 gopurs_runtime.Value
_ = go__go_1_0_23
Call_local_Data_List_Lazy_go__go_1_0_23 = func(v_2_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] {
go__go_1_0_23:
for {
if false { continue go__go_1_0_23 }
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
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (v_2).V1))
continue go__go_1_0_23
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
go__go_1_0_23 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Lazy_go__go_1_0_23(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_2_loop_val)))}
})
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_List(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0, go__go_1_0_23), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))))
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
return gopurs_runtime.Bool((gopurs_runtime.Apply3(Call_Data_Foldable_any(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_foldableList())), gopurs_runtime.Value{Type: 9, IntVal: 926771738, UnsafePtr: unsafe.Pointer(Rebox_Data_List_Lazy_3591112874_2663347022(Rebox_Data_List_Lazy_2663347022_3591112874(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]](Get_Data_HeytingAlgebra_heytingAlgebraBoolean()))))}, gopurs_runtime.Apply(eq_0, x_3), ys_2).IntVal) != (0))
})), xs_1)
}

func Call_Data_List_Lazy_intersect(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_List_Lazy_intersectBy(), Call_Data_Eq_eq(dictEq_0))
}

func Call_Data_List_Lazy_nubByEq(eq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
nubByEq:
for {
if false { continue nubByEq }
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_List(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_1)
_ = __t_tag_0
if (__t_tag_0 == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_3
} else {

}
}
{
var __t_tag_1 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_1)
_ = __t_tag_1
if (__t_tag_1 != nil) {
// TAST (Let): __local_var_2_2 shape=Other bindingType=Any
__local_var_2_2 := (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V0
_ = __local_var_2_2
__t3 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, __local_var_2_2, gopurs_runtime.Apply(Call_Data_List_Lazy_nubByEq(eq_0), gopurs_runtime.Apply(Call_Data_List_Lazy_filter(gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(eq_0, __local_var_2_2, y_3).IntVal) != (0)) != (true))
})), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(v_1.UnsafePtr).V1))})
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t3)}
})), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))))
}
}

func Call_Data_List_Lazy_nubEq(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return Call_Data_List_Lazy_nubByEq(Call_Data_Eq_eq(dictEq_0))
}

func Call_Data_List_Lazy_eqPattern(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
// TAST (Let): eqList_1_0 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope268)])])])
eqList_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_List_Lazy_Types_eqList(dictEq_0))
_ = eqList_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(eqList_1_0.V0, x_2, y_3).IntVal) != (0))
})}))}
}

func Call_Data_List_Lazy_ordPattern(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): ordList_1_0 shape=App(Var) bindingType=(ADT ["Data","Ord","Ord"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope121)])])])
ordList_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Call_Data_List_Lazy_Types_ordList(dictOrd_0))
_ = ordList_1_0
// TAST (Let): eqPattern1_2_1 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope121)])])])
eqPattern1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_List_Lazy_eqPattern(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})))
_ = eqPattern1_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqPattern1_2_1)}
}), gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(ordList_1_0.V1, x_3, y_4).IntVal)), UnsafePtr: nil}
})}))}
}

func Call_Data_List_Lazy_elemLastIndex(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_Data_List_Lazy_findLastIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(dictEq_0.V0, v_2, x_1).IntVal) != (0))
}))
}

func Call_Data_List_Lazy_elemIndex(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return Call_Data_List_Lazy_findIndex(gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(dictEq_0.V0, v_2, x_1).IntVal) != (0))
}))
}

func Call_Data_List_Lazy_dropWhile(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var Call_local_Data_List_Lazy_go__go_1_0_24 func(*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_1_0_24
var go__go_1_0_24 gopurs_runtime.Value
_ = go__go_1_0_24
Call_local_Data_List_Lazy_go__go_1_0_24 = func(v_2_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_1_0_24:
for {
if false { continue go__go_1_0_24 }
var v_2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if ((v_2 != nil)) && ((gopurs_runtime.Apply(p_0, (v_2).V0).IntVal) != (0)) {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (v_2).V1))
continue go__go_1_0_24
__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply(Get_Data_List_Lazy_fromStep(), gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(v_2)})
}
end_branch_1:
return __t1
}
}
go__go_1_0_24 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_1_0_24(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v_2_loop_val))
})
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), go__go_1_0_24, Get_Data_List_Lazy_Types_step())
}

func Call_Data_List_Lazy_drop(n_0_loop int64) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var Call_local_Data_List_Lazy_go__2018343127_1_0_25 func(int64, *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Lazy_go__2018343127_1_0_25
var go__2018343127_1_0_25 gopurs_runtime.Value
_ = go__2018343127_1_0_25
var Call_local_Data_List_Lazy_go__go_1_1_26 func(int64, *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_List_Lazy_go__go_1_1_26
var go__go_1_1_26 gopurs_runtime.Value
_ = go__go_1_1_26
Call_local_Data_List_Lazy_go__2018343127_1_0_25 = func(v_2_loop int64, v1_3_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] {
go__2018343127_1_0_25:
for {
if false { continue go__2018343127_1_0_25 }
var v_2 int64 = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (v_2) == (int64(0)) {
__t2 = v1_3
goto end_branch_2
} else {

}
}
{
if (v1_3 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_2
} else {

}
}
{
if (v1_3 != nil) {
v_2_loop = (v_2) - (int64(1))
v1_3_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (v1_3).V1))
continue go__2018343127_1_0_25
__t2 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
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
go__2018343127_1_0_25 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Lazy_go__2018343127_1_0_25(v_2_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val)))}
})
})
Call_local_Data_List_Lazy_go__go_1_1_26 = func(v_2_loop int64, v1_3_loop *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] {
go__go_1_1_26:
for {
if false { continue go__go_1_1_26 }
var v_2 int64 = v_2_loop
_ = v_2
var v1_3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (v_2) == (int64(0)) {
__t3 = v1_3
goto end_branch_3
} else {

}
}
{
if (v1_3 == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_3
} else {

}
}
{
if (v1_3 != nil) {
__t3 = Call_local_Data_List_Lazy_go__2018343127_1_0_25((v_2) - (int64(1)), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (v1_3).V1)))
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
go__go_1_1_26 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(Call_local_Data_List_Lazy_go__go_1_1_26(v_2_loop_val.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](v1_3_loop_val)))}
})
})
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_List(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()).V0, gopurs_runtime.Apply(go__2018343127_1_0_25, gopurs_runtime.Int(n_0))), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))))
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
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope286)])])
__local_var_3_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_2)
_ = __local_var_3_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_1 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_5_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_3_0)
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
if (gopurs_runtime.Apply2(eq_0, x_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V0).IntVal) != (0) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
__t2 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V0, Call_Data_List_Lazy_deleteBy(eq_0, x_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V1)})
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

func Call_Data_List_Lazy_unionBy(eq_0_loop gopurs_runtime.Value, xs_1_loop gopurs_runtime.Value, ys_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 gopurs_runtime.Value = ys_2_loop
_ = ys_2
var Call_local_Data_List_Lazy_go__go_3_1_27 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_3_1_27
var go__go_3_1_27 gopurs_runtime.Value
_ = go__go_3_1_27
Call_local_Data_List_Lazy_go__go_3_1_27 = func(b_4_loop gopurs_runtime.Value, xs_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_3_1_27:
for {
if false { continue go__go_3_1_27 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var xs_5 gopurs_runtime.Value = xs_5_loop
_ = xs_5
// TAST (Let): v_6_2 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_5))
_ = v_6_2
var __t8 gopurs_runtime.Value
{
if (v_6_2 == nil) {
__t8 = b_4
goto end_branch_8
} else {

}
}
{
if (v_6_2 != nil) {
// TAST (Let): __local_var_7_3 shape=Other bindingType=(TypeVar a$scope199)
__local_var_7_3 := (v_6_2).V0
_ = __local_var_7_3
// TAST (Let): __local_var_8_4 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope292)])])
__local_var_8_4 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), b_4)
_ = __local_var_8_4
b_4_loop = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_5 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_10_5 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_8_4)
_ = __local_var_10_5
var __t7 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_10_5.Type == 9 && __local_var_10_5.IntVal == 218341868 && __local_var_10_5.UnsafePtr == nil) {
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_7
} else {

}
}
{
if (__local_var_10_5.Type == 9 && __local_var_10_5.IntVal == 218341868 && __local_var_10_5.UnsafePtr != nil) {
var __t6 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply2(eq_0, __local_var_7_3, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_10_5.UnsafePtr).V0).IntVal) != (0) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_10_5.UnsafePtr).V1))
goto end_branch_6
} else {

}
}
{
__t6 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_10_5.UnsafePtr).V0, Call_Data_List_Lazy_deleteBy(eq_0, __local_var_7_3, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_10_5.UnsafePtr).V1)})
}
end_branch_6:
__t7 = __t6
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
xs_5_loop = (v_6_2).V1
continue go__go_3_1_27
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
go__go_3_1_27 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_3_1_27(b_4_loop_val, xs_5_loop_val)
})
})
// TAST (Let): __local_var_3_0 shape=LetRec(App(Other)) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope292)])])
__local_var_3_0 := Call_local_Data_List_Lazy_go__go_3_1_27(gopurs_runtime.Apply(Call_Data_List_Lazy_nubByEq(eq_0), ys_2), xs_1)
_ = __local_var_3_0
// TAST (Let): __local_var_4_9 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope50)])])
__local_var_4_9 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_1)
_ = __local_var_4_9
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_10 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_6_10 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_9)
_ = __local_var_6_10
var __t11 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_6_10.Type == 9 && __local_var_6_10.IntVal == 218341868 && __local_var_6_10.UnsafePtr == nil) {
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), __local_var_3_0))
goto end_branch_11
} else {

}
}
{
if (__local_var_6_10.Type == 9 && __local_var_6_10.IntVal == 218341868 && __local_var_6_10.UnsafePtr != nil) {
__t11 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_10.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_10.UnsafePtr).V1, __local_var_3_0)})
goto end_branch_11
} else {

}
}
{
__t11 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t11)}
}))
}

func Call_Data_List_Lazy_union(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_List_Lazy_unionBy(), Call_Data_Eq_eq(dictEq_0))
}

func Call_Data_List_Lazy_deleteAt(n_0_loop int64, xs_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
deleteAt:
for {
if false { continue deleteAt }
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 gopurs_runtime.Value = xs_1_loop
_ = xs_1
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope296)])])
__local_var_2_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_1)
_ = __local_var_2_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_4_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_2_0)
_ = __local_var_4_1
var __t3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_3
} else {

}
}
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 218341868 && __local_var_4_1.UnsafePtr != nil) {
var __t2 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (n_0) == (int64(0)) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
__t2 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0, Call_Data_List_Lazy_deleteAt((n_0) - (int64(1)), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1)})
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

func Call_Data_List_Lazy_go__delete(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope302), (TypeVar a$scope302)] Boolean)
__local_var_1_0 := Call_Data_Eq_eq(dictEq_0)
_ = __local_var_1_0
return gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope302)])])
__local_var_4_1 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_3)
_ = __local_var_4_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_6_2 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_1)
_ = __local_var_6_2
var __t4 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr == nil) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_4
} else {

}
}
{
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr != nil) {
var __t3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply2(__local_var_1_0, x_2, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V0).IntVal) != (0) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V1))
goto end_branch_3
} else {

}
}
{
__t3 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V0, Call_Data_List_Lazy_deleteBy(__local_var_1_0, x_2, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V1)})
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

func Call_Data_List_Lazy_difference(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
// TAST (Let): __local_var_1_1 shape=App(Var) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(Get_Data_List_Lazy_go__delete(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_0)})
_ = __local_var_1_1
// TAST (Let): __local_var_1_0 shape=Let(Abs(Abs(App(Other)))) bindingType=(Func [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope304)])]), (TypeVar a$scope304)] (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope304)])]))
__local_var_1_0 := gopurs_runtime.Func2(func(b_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_1_1, a_3, b_2)
})
_ = __local_var_1_0
var Call_local_Data_List_Lazy_go__go_2_2_28 func(gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_List_Lazy_go__go_2_2_28
var go__go_2_2_28 gopurs_runtime.Value
_ = go__go_2_2_28
Call_local_Data_List_Lazy_go__go_2_2_28 = func(b_3_loop gopurs_runtime.Value, xs_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_2_2_28:
for {
if false { continue go__go_2_2_28 }
var b_3 gopurs_runtime.Value = b_3_loop
_ = b_3
var xs_4 gopurs_runtime.Value = xs_4_loop
_ = xs_4
// TAST (Let): v_5_3 shape=App(Var) bindingType=(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope199)])
v_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_4))
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
continue go__go_2_2_28
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
go__go_2_2_28 = gopurs_runtime.Func(func(b_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_List_Lazy_go__go_2_2_28(b_3_loop_val, xs_4_loop_val)
})
})
return go__go_2_2_28
}

func Call_Data_List_Lazy_cycle(xs_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var go__go_1_0_29 gopurs_runtime.Value
_ = go__go_1_0_29
var go__go_1_0_29_cell *gopurs_runtime.Value
_ = go__go_1_0_29_cell
// FALLBACK TCO: isLoop=false len=1
go__go_1_0_29 = gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_List_Lazy_Types_step(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope50)])])
__local_var_3_1 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_0)
_ = __local_var_3_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_2 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_5_2 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_3_1)
_ = __local_var_5_2
var __t3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_5_2.Type == 9 && __local_var_5_2.IntVal == 218341868 && __local_var_5_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (*go__go_1_0_29_cell)))
goto end_branch_3
} else {

}
}
{
if (__local_var_5_2.Type == 9 && __local_var_5_2.IntVal == 218341868 && __local_var_5_2.UnsafePtr != nil) {
__t3 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_2.UnsafePtr).V1, (*go__go_1_0_29_cell))})
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
})))
go__go_1_0_29_cell = &go__go_1_0_29
return go__go_1_0_29
}

func Call_Data_List_Lazy_concat(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope312)])])] (ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope312)])]))
__local_var_1_0 := Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
_ = __local_var_1_0
// TAST (Let): __local_var_2_1 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope297)])])
__local_var_2_1 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), v_0)
_ = __local_var_2_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_4_2 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_2_1)
_ = __local_var_4_2
var __t10 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_4_2.Type == 9 && __local_var_4_2.IntVal == 218341868 && __local_var_4_2.UnsafePtr == nil) {
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_10
} else {

}
}
{
if (__local_var_4_2.Type == 9 && __local_var_4_2.IntVal == 218341868 && __local_var_4_2.UnsafePtr != nil) {
// TAST (Let): __local_var_5_3 shape=App(Other) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar b$scope298)])])
__local_var_5_3 := gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_List_Lazy_Types_bindList()).V1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_2.UnsafePtr).V1, __local_var_1_0)
_ = __local_var_5_3
// TAST (Let): __local_var_6_4 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope50)])])
__local_var_6_4 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Apply(__local_var_1_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_4_2.UnsafePtr).V0))
_ = __local_var_6_4
__t10 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_5 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_8_5 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_6_4)
_ = __local_var_8_5
var __t9 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_8_5.Type == 9 && __local_var_8_5.IntVal == 218341868 && __local_var_8_5.UnsafePtr == nil) {
__t9 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), __local_var_5_3))
goto end_branch_9
} else {

}
}
{
if (__local_var_8_5.Type == 9 && __local_var_8_5.IntVal == 218341868 && __local_var_8_5.UnsafePtr != nil) {
// TAST (Let): __local_var_9_6 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope50)])])
__local_var_9_6 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_8_5.UnsafePtr).V1)
_ = __local_var_9_6
__t9 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_8_5.UnsafePtr).V0, gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_7 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_11_7 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_9_6)
_ = __local_var_11_7
var __t8 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_11_7.Type == 9 && __local_var_11_7.IntVal == 218341868 && __local_var_11_7.UnsafePtr == nil) {
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), __local_var_5_3))
goto end_branch_8
} else {

}
}
{
if (__local_var_11_7.Type == 9 && __local_var_11_7.IntVal == 218341868 && __local_var_11_7.UnsafePtr != nil) {
__t8 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_11_7.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_11_7.UnsafePtr).V1, __local_var_5_3)})
goto end_branch_8
} else {

}
}
{
__t8 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t8)}
}))})
goto end_branch_9
} else {

}
}
{
__t9 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t9)}
}))))
goto end_branch_10
} else {

}
}
{
__t10 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(__t10)}
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
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope318)])])
__local_var_3_0 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_2)
_ = __local_var_3_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_1 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_5_1 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_3_0)
_ = __local_var_5_1
var __t5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_5_1.Type == 9 && __local_var_5_1.IntVal == 218341868 && __local_var_5_1.UnsafePtr == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_5
} else {

}
}
{
if (__local_var_5_1.Type == 9 && __local_var_5_1.IntVal == 218341868 && __local_var_5_1.UnsafePtr != nil) {
var __t4 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (n_0) == (int64(0)) {
// TAST (Let): v2_6_2 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope318)])
v2_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V0))
_ = v2_6_2
var __t3 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (v2_6_2 == nil) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V1))
goto end_branch_3
} else {

}
}
{
if (v2_6_2 != nil) {
__t3 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (v2_6_2).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V1})
goto end_branch_3
} else {

}
}
{
__t3 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
__t4 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V0, Call_Data_List_Lazy_alterAt((n_0) - (int64(1)), f_1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_5_1.UnsafePtr).V1)})
}
end_branch_4:
__t5 = __t4
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
}
}

func Call_Data_List_Lazy_modifyAt(n_0_loop int64, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope325)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope325)]))
__local_var_2_0 := gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Maybe_Just(), f_1)
_ = __local_var_2_0
return gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(ADT ["Data","List","Lazy","Types","Step"] [(TypeVar a$scope325)])])
__local_var_4_1 := gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), xs_3)
_ = __local_var_4_1
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 shape=App(Var) bindingType=(TypeVar a$scope106)
__local_var_6_2 := gopurs_runtime.Apply(Get_Data_Lazy_force(), __local_var_4_1)
_ = __local_var_6_2
var __t6 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(nil))})
goto end_branch_6
} else {

}
}
{
if (__local_var_6_2.Type == 9 && __local_var_6_2.IntVal == 218341868 && __local_var_6_2.UnsafePtr != nil) {
var __t5 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (n_0) == (int64(0)) {
// TAST (Let): v2_7_3 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope325)])
v2_7_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(__local_var_2_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V0))
_ = v2_7_3
var __t4 *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]
{
if (v2_7_3 == nil) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V1))
goto end_branch_4
} else {

}
}
{
if (v2_7_3 != nil) {
__t4 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (v2_7_3).V0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V1})
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = __t4
goto end_branch_5
} else {

}
}
{
__t5 = (&Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]{1, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V0, Call_Data_List_Lazy_alterAt((n_0) - (int64(1)), __local_var_2_0, (*Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value])(__local_var_6_2.UnsafePtr).V1)})
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
})
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

func Rebox_Data_List_Lazy_2663347022_3591112874(in *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]) *Constructor_Data_HeytingAlgebra_HeytingAlgebra[bool] {
	if in == nil { return nil }
	out := &Constructor_Data_HeytingAlgebra_HeytingAlgebra[bool]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = (in.V2.IntVal) != (0)
		out.V3 = in.V3
		out.V4 = in.V4
		out.V5 = (in.V5.IntVal) != (0)
	return out
}

func Rebox_Data_List_Lazy_2826095630_348932501(in *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]) *Constructor_Data_Semiring_Semiring[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Semiring_Semiring[int64]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2.IntVal
		out.V3 = in.V3.IntVal
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

func Rebox_Data_List_Lazy_3094389156_3833657837(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail gopurs_runtime.Value
}] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail gopurs_runtime.Value
}]{}
		out.V0 = func() struct{
	head gopurs_runtime.Value
	tail gopurs_runtime.Value
} {
					orig := in.V0
					_ = orig
					clone := struct{
	head gopurs_runtime.Value
	tail gopurs_runtime.Value
}{}
					clone.head = gopurs_runtime.RecordGet(orig, "head")
					clone.tail = gopurs_runtime.RecordGet(orig, "tail")
					return clone
				}()
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

func Rebox_Data_List_Lazy_348932501_2826095630(in *Constructor_Data_Semiring_Semiring[int64]) *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = gopurs_runtime.Int(in.V2)
		out.V3 = gopurs_runtime.Int(in.V3)
	return out
}

func Rebox_Data_List_Lazy_3591112874_2663347022(in *Constructor_Data_HeytingAlgebra_HeytingAlgebra[bool]) *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = gopurs_runtime.Bool(in.V2)
		out.V3 = in.V3
		out.V4 = in.V4
		out.V5 = gopurs_runtime.Bool(in.V5)
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
				return gopurs_runtime.RecordDict2("head", "tail", orig.head, orig.tail)
				}()
	return out
}

func Rebox_Data_List_Lazy_4130553207_1542299734(in *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]) *Constructor_Control_Monad_Rec_Class_MonadRec[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Rec_Class_MonadRec[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
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
				return gopurs_runtime.RecordDict2("a", "b", orig.a, orig.b)
				}()
	return out
}



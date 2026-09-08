package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Array_NonEmpty_intercalate1 gopurs_runtime.Value
var once_Data_Array_NonEmpty_intercalate1 sync.Once
func Get_Data_Array_NonEmpty_intercalate1() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_intercalate1.Do(func() {
		cache_Data_Array_NonEmpty_intercalate1 = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_intercalate1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_0_box))
})
	})
	return cache_Data_Array_NonEmpty_intercalate1
}

var cache_Data_Array_NonEmpty_foldMap11 gopurs_runtime.Value
var once_Data_Array_NonEmpty_foldMap11 sync.Once
func Get_Data_Array_NonEmpty_foldMap11() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_foldMap11.Do(func() {
		cache_Data_Array_NonEmpty_foldMap11 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[[]gopurs_runtime.Value]](Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray()).V1)
	})
	return cache_Data_Array_NonEmpty_foldMap11
}

var cache_Data_Array_NonEmpty_fold11 gopurs_runtime.Value
var once_Data_Array_NonEmpty_fold11 sync.Once
func Get_Data_Array_NonEmpty_fold11() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_fold11.Do(func() {
		cache_Data_Array_NonEmpty_fold11 = gopurs_runtime.Apply(Get_Data_Semigroup_Foldable_fold1(), gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_3016437835_4151366573(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[[]gopurs_runtime.Value]](Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray())))})
	})
	return cache_Data_Array_NonEmpty_fold11
}

var cache_Data_Array_NonEmpty_fromJust gopurs_runtime.Value
var once_Data_Array_NonEmpty_fromJust sync.Once
func Get_Data_Array_NonEmpty_fromJust() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_fromJust.Do(func() {
		cache_Data_Array_NonEmpty_fromJust = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_fromJust(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_Data_Array_NonEmpty_fromJust
}

var cache_Data_Array_NonEmpty_unsafeIndex1 gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeIndex1 sync.Once
func Get_Data_Array_NonEmpty_unsafeIndex1() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeIndex1.Do(func() {
		cache_Data_Array_NonEmpty_unsafeIndex1 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_unsafeIndex1(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), __local_var_1_box.IntVal)
})
	})
	return cache_Data_Array_NonEmpty_unsafeIndex1
}

var cache_Data_Array_NonEmpty_unsafeFromArrayF gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeFromArrayF sync.Once
func Get_Data_Array_NonEmpty_unsafeFromArrayF() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeFromArrayF.Do(func() {
		cache_Data_Array_NonEmpty_unsafeFromArrayF = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Array_NonEmpty_unsafeFromArrayF
}

var cache_Data_Array_NonEmpty_unsafeFromArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeFromArray sync.Once
func Get_Data_Array_NonEmpty_unsafeFromArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeFromArray.Do(func() {
		cache_Data_Array_NonEmpty_unsafeFromArray = Get_Data_Array_NonEmpty_Internal_NonEmptyArray()
	})
	return cache_Data_Array_NonEmpty_unsafeFromArray
}

var cache_Data_Array_NonEmpty_unsafeFromArray__1127564228 gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeFromArray__1127564228 sync.Once
func Get_Data_Array_NonEmpty_unsafeFromArray__1127564228() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeFromArray__1127564228.Do(func() {
		cache_Data_Array_NonEmpty_unsafeFromArray__1127564228 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Array_NonEmpty_unsafeFromArray__1127564228(func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(__eta_norm_0_0_box.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_Array_NonEmpty_unsafeFromArray__1127564228
}

var cache_Data_Array_NonEmpty_transpose gopurs_runtime.Value
var once_Data_Array_NonEmpty_transpose sync.Once
func Get_Data_Array_NonEmpty_transpose() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_transpose.Do(func() {
		cache_Data_Array_NonEmpty_transpose = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_transpose(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_transpose
}

var cache_Data_Array_NonEmpty_toArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_toArray sync.Once
func Get_Data_Array_NonEmpty_toArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_toArray.Do(func() {
		cache_Data_Array_NonEmpty_toArray = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_toArray(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_toArray
}

var cache_Data_Array_NonEmpty_unionBy_prime_ gopurs_runtime.Value
var once_Data_Array_NonEmpty_unionBy_prime_ sync.Once
func Get_Data_Array_NonEmpty_unionBy_prime_() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unionBy_prime_.Do(func() {
		cache_Data_Array_NonEmpty_unionBy_prime_ = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_unionBy_prime_(eq_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_unionBy_prime_
}

var cache_Data_Array_NonEmpty_union_prime_ gopurs_runtime.Value
var once_Data_Array_NonEmpty_union_prime_ sync.Once
func Get_Data_Array_NonEmpty_union_prime_() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_union_prime_.Do(func() {
		cache_Data_Array_NonEmpty_union_prime_ = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_union_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_Data_Array_NonEmpty_union_prime_
}

var cache_Data_Array_NonEmpty_unionBy gopurs_runtime.Value
var once_Data_Array_NonEmpty_unionBy sync.Once
func Get_Data_Array_NonEmpty_unionBy() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unionBy.Do(func() {
		cache_Data_Array_NonEmpty_unionBy = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_unionBy(eq_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_unionBy
}

var cache_Data_Array_NonEmpty_union gopurs_runtime.Value
var once_Data_Array_NonEmpty_union sync.Once
func Get_Data_Array_NonEmpty_union() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_union.Do(func() {
		cache_Data_Array_NonEmpty_union = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_union(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_Data_Array_NonEmpty_union
}

var cache_Data_Array_NonEmpty_unzip gopurs_runtime.Value
var once_Data_Array_NonEmpty_unzip sync.Once
func Get_Data_Array_NonEmpty_unzip() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unzip.Do(func() {
		cache_Data_Array_NonEmpty_unzip = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_NonEmpty_unzip(func() []*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v) }
					return unboxed
				}())
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()
})
	})
	return cache_Data_Array_NonEmpty_unzip
}

var cache_Data_Array_NonEmpty_updateAt gopurs_runtime.Value
var once_Data_Array_NonEmpty_updateAt sync.Once
func Get_Data_Array_NonEmpty_updateAt() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_updateAt.Do(func() {
		cache_Data_Array_NonEmpty_updateAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_NonEmpty_updateAt(i_0_box.IntVal, x_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_NonEmpty_updateAt
}

var cache_Data_Array_NonEmpty_zip gopurs_runtime.Value
var once_Data_Array_NonEmpty_zip sync.Once
func Get_Data_Array_NonEmpty_zip() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_zip.Do(func() {
		cache_Data_Array_NonEmpty_zip = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, ys_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Array_NonEmpty_zip(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(ys_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_Array_NonEmpty_zip
}

var cache_Data_Array_NonEmpty_zipWith gopurs_runtime.Value
var once_Data_Array_NonEmpty_zipWith sync.Once
func Get_Data_Array_NonEmpty_zipWith() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_zipWith.Do(func() {
		cache_Data_Array_NonEmpty_zipWith = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_zipWith(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(ys_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_zipWith
}

var cache_Data_Array_NonEmpty_zipWithA gopurs_runtime.Value
var once_Data_Array_NonEmpty_zipWithA sync.Once
func Get_Data_Array_NonEmpty_zipWithA() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_zipWithA.Do(func() {
		cache_Data_Array_NonEmpty_zipWithA = gopurs_runtime.Func4(func(dictApplicative_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value, ys_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_zipWithA(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), f_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(ys_3_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_NonEmpty_zipWithA
}

var cache_Data_Array_NonEmpty_splitAt gopurs_runtime.Value
var once_Data_Array_NonEmpty_splitAt sync.Once
func Get_Data_Array_NonEmpty_splitAt() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_splitAt.Do(func() {
		cache_Data_Array_NonEmpty_splitAt = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := Call_Data_Array_NonEmpty_splitAt(i_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				_ = orig
				return gopurs_runtime.RecordDict([]string{"after", "before"}, []gopurs_runtime.Value{gopurs_runtime.Array(orig.after), gopurs_runtime.Array(orig.before)})
				}()
})
	})
	return cache_Data_Array_NonEmpty_splitAt
}

var cache_Data_Array_NonEmpty_some gopurs_runtime.Value
var once_Data_Array_NonEmpty_some sync.Once
func Get_Data_Array_NonEmpty_some() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_some.Do(func() {
		cache_Data_Array_NonEmpty_some = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_some(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](dictAlternative_0_box))
})
	})
	return cache_Data_Array_NonEmpty_some
}

var cache_Data_Array_NonEmpty_snoc_prime_ gopurs_runtime.Value
var once_Data_Array_NonEmpty_snoc_prime_ sync.Once
func Get_Data_Array_NonEmpty_snoc_prime_() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_snoc_prime_.Do(func() {
		cache_Data_Array_NonEmpty_snoc_prime_ = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_snoc_prime_(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), x_1_box))
})
	})
	return cache_Data_Array_NonEmpty_snoc_prime_
}

var cache_Data_Array_NonEmpty_snoc gopurs_runtime.Value
var once_Data_Array_NonEmpty_snoc sync.Once
func Get_Data_Array_NonEmpty_snoc() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_snoc.Do(func() {
		cache_Data_Array_NonEmpty_snoc = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_snoc(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), x_1_box))
})
	})
	return cache_Data_Array_NonEmpty_snoc
}

var cache_Data_Array_NonEmpty_singleton gopurs_runtime.Value
var once_Data_Array_NonEmpty_singleton sync.Once
func Get_Data_Array_NonEmpty_singleton() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_singleton.Do(func() {
		cache_Data_Array_NonEmpty_singleton = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_singleton(x_0_box))
})
	})
	return cache_Data_Array_NonEmpty_singleton
}

var cache_Data_Array_NonEmpty_replicate gopurs_runtime.Value
var once_Data_Array_NonEmpty_replicate sync.Once
func Get_Data_Array_NonEmpty_replicate() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_replicate.Do(func() {
		cache_Data_Array_NonEmpty_replicate = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_replicate(i_0_box.IntVal, x_1_box))
})
	})
	return cache_Data_Array_NonEmpty_replicate
}

var cache_Data_Array_NonEmpty_go__range gopurs_runtime.Value
var once_Data_Array_NonEmpty_go__range sync.Once
func Get_Data_Array_NonEmpty_go__range() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_go__range.Do(func() {
		cache_Data_Array_NonEmpty_go__range = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Array_NonEmpty_go__range(x_0_box.IntVal, y_1_box.IntVal)
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_Array_NonEmpty_go__range
}

var cache_Data_Array_NonEmpty_prependArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_prependArray sync.Once
func Get_Data_Array_NonEmpty_prependArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_prependArray.Do(func() {
		cache_Data_Array_NonEmpty_prependArray = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, ys_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_prependArray(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(ys_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_prependArray
}

var cache_Data_Array_NonEmpty_modifyAt gopurs_runtime.Value
var once_Data_Array_NonEmpty_modifyAt sync.Once
func Get_Data_Array_NonEmpty_modifyAt() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_modifyAt.Do(func() {
		cache_Data_Array_NonEmpty_modifyAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_NonEmpty_modifyAt(i_0_box.IntVal, f_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_NonEmpty_modifyAt
}

var cache_Data_Array_NonEmpty_intersectBy_prime_ gopurs_runtime.Value
var once_Data_Array_NonEmpty_intersectBy_prime_ sync.Once
func Get_Data_Array_NonEmpty_intersectBy_prime_() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_intersectBy_prime_.Do(func() {
		cache_Data_Array_NonEmpty_intersectBy_prime_ = gopurs_runtime.Func2(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_intersectBy_prime_(eq_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_NonEmpty_intersectBy_prime_
}

var cache_Data_Array_NonEmpty_intersectBy gopurs_runtime.Value
var once_Data_Array_NonEmpty_intersectBy sync.Once
func Get_Data_Array_NonEmpty_intersectBy() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_intersectBy.Do(func() {
		cache_Data_Array_NonEmpty_intersectBy = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_intersectBy(eq_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_intersectBy
}

var cache_Data_Array_NonEmpty_intersect_prime_ gopurs_runtime.Value
var once_Data_Array_NonEmpty_intersect_prime_ sync.Once
func Get_Data_Array_NonEmpty_intersect_prime_() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_intersect_prime_.Do(func() {
		cache_Data_Array_NonEmpty_intersect_prime_ = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_intersect_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_Data_Array_NonEmpty_intersect_prime_
}

var cache_Data_Array_NonEmpty_intersect gopurs_runtime.Value
var once_Data_Array_NonEmpty_intersect sync.Once
func Get_Data_Array_NonEmpty_intersect() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_intersect.Do(func() {
		cache_Data_Array_NonEmpty_intersect = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_intersect(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_Data_Array_NonEmpty_intersect
}

var cache_Data_Array_NonEmpty_intercalate gopurs_runtime.Value
var once_Data_Array_NonEmpty_intercalate sync.Once
func Get_Data_Array_NonEmpty_intercalate() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_intercalate.Do(func() {
		cache_Data_Array_NonEmpty_intercalate = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_intercalate(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_0_box))
})
	})
	return cache_Data_Array_NonEmpty_intercalate
}

var cache_Data_Array_NonEmpty_insertAt gopurs_runtime.Value
var once_Data_Array_NonEmpty_insertAt sync.Once
func Get_Data_Array_NonEmpty_insertAt() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_insertAt.Do(func() {
		cache_Data_Array_NonEmpty_insertAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_NonEmpty_insertAt(i_0_box.IntVal, x_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_NonEmpty_insertAt
}

var cache_Data_Array_NonEmpty_fromFoldable1 gopurs_runtime.Value
var once_Data_Array_NonEmpty_fromFoldable1 sync.Once
func Get_Data_Array_NonEmpty_fromFoldable1() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_fromFoldable1.Do(func() {
		cache_Data_Array_NonEmpty_fromFoldable1 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_fromFoldable1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box))
})
	})
	return cache_Data_Array_NonEmpty_fromFoldable1
}

var cache_Data_Array_NonEmpty_fromArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_fromArray sync.Once
func Get_Data_Array_NonEmpty_fromArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_fromArray.Do(func() {
		cache_Data_Array_NonEmpty_fromArray = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_NonEmpty_fromArray(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_NonEmpty_fromArray
}

var cache_Data_Array_NonEmpty_fromFoldable gopurs_runtime.Value
var once_Data_Array_NonEmpty_fromFoldable sync.Once
func Get_Data_Array_NonEmpty_fromFoldable() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_fromFoldable.Do(func() {
		cache_Data_Array_NonEmpty_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_NonEmpty_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_NonEmpty_fromFoldable
}

var cache_Data_Array_NonEmpty_transpose_prime_ gopurs_runtime.Value
var once_Data_Array_NonEmpty_transpose_prime_ sync.Once
func Get_Data_Array_NonEmpty_transpose_prime_() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_transpose_prime_.Do(func() {
		cache_Data_Array_NonEmpty_transpose_prime_ = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_NonEmpty_transpose_prime_(func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}() }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_NonEmpty_transpose_prime_
}

var cache_Data_Array_NonEmpty_foldr1 gopurs_runtime.Value
var once_Data_Array_NonEmpty_foldr1 sync.Once
func Get_Data_Array_NonEmpty_foldr1() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_foldr1.Do(func() {
		cache_Data_Array_NonEmpty_foldr1 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[[]gopurs_runtime.Value]](Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray()).V3)
	})
	return cache_Data_Array_NonEmpty_foldr1
}

var cache_Data_Array_NonEmpty_foldl1 gopurs_runtime.Value
var once_Data_Array_NonEmpty_foldl1 sync.Once
func Get_Data_Array_NonEmpty_foldl1() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_foldl1.Do(func() {
		cache_Data_Array_NonEmpty_foldl1 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[[]gopurs_runtime.Value]](Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray()).V2)
	})
	return cache_Data_Array_NonEmpty_foldl1
}

var cache_Data_Array_NonEmpty_foldMap1 gopurs_runtime.Value
var once_Data_Array_NonEmpty_foldMap1 sync.Once
func Get_Data_Array_NonEmpty_foldMap1() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_foldMap1.Do(func() {
		cache_Data_Array_NonEmpty_foldMap1 = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_foldMap1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_0_box))
})
	})
	return cache_Data_Array_NonEmpty_foldMap1
}

var cache_Data_Array_NonEmpty_fold1 gopurs_runtime.Value
var once_Data_Array_NonEmpty_fold1 sync.Once
func Get_Data_Array_NonEmpty_fold1() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_fold1.Do(func() {
		cache_Data_Array_NonEmpty_fold1 = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_fold1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_0_box))
})
	})
	return cache_Data_Array_NonEmpty_fold1
}

var cache_Data_Array_NonEmpty_difference_prime_ gopurs_runtime.Value
var once_Data_Array_NonEmpty_difference_prime_ sync.Once
func Get_Data_Array_NonEmpty_difference_prime_() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_difference_prime_.Do(func() {
		cache_Data_Array_NonEmpty_difference_prime_ = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_difference_prime_(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_Data_Array_NonEmpty_difference_prime_
}

var cache_Data_Array_NonEmpty_cons_prime_ gopurs_runtime.Value
var once_Data_Array_NonEmpty_cons_prime_ sync.Once
func Get_Data_Array_NonEmpty_cons_prime_() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_cons_prime_.Do(func() {
		cache_Data_Array_NonEmpty_cons_prime_ = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_cons_prime_(x_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_cons_prime_
}

var cache_Data_Array_NonEmpty_fromNonEmpty gopurs_runtime.Value
var once_Data_Array_NonEmpty_fromNonEmpty sync.Once
func Get_Data_Array_NonEmpty_fromNonEmpty() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_fromNonEmpty.Do(func() {
		cache_Data_Array_NonEmpty_fromNonEmpty = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_fromNonEmpty(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))
})
	})
	return cache_Data_Array_NonEmpty_fromNonEmpty
}

var cache_Data_Array_NonEmpty_concatMap gopurs_runtime.Value
var once_Data_Array_NonEmpty_concatMap sync.Once
func Get_Data_Array_NonEmpty_concatMap() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_concatMap.Do(func() {
		cache_Data_Array_NonEmpty_concatMap = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_concatMap(b_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(a_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_concatMap
}

var cache_Data_Array_NonEmpty_concat gopurs_runtime.Value
var once_Data_Array_NonEmpty_concat sync.Once
func Get_Data_Array_NonEmpty_concat() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_concat.Do(func() {
		cache_Data_Array_NonEmpty_concat = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_2 shape=App(Var) bindingType=(Func [(Array (ADT ["Data","Array","NonEmpty","Internal","NonEmptyArray"] [(TypeVar a)]))] (Array (Array (TypeVar a))))
__local_var_0_2 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), Get_Data_Array_NonEmpty_toArray())
_ = __local_var_0_2
// TAST (Let): __local_var_0_1 shape=Let(Abs(App(Other))) bindingType=(Func [(Array (ADT ["Data","Array","NonEmpty","Internal","NonEmptyArray"] [(TypeVar a)]))] (Array (Array (TypeVar a))))
__local_var_0_1 := gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(__local_var_0_2, x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
_ = __local_var_0_1
// TAST (Let): __local_var_0_0 shape=Let(Abs(App(Var))) bindingType=(Func [(Array (ADT ["Data","Array","NonEmpty","Internal","NonEmptyArray"] [(TypeVar a)]))] (Array (TypeVar a)))
__local_var_0_0 := gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Array_concat(), gopurs_runtime.Apply(__local_var_0_1, x_1))
})
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_0_0, x_1)
})
}()
	})
	return cache_Data_Array_NonEmpty_concat
}

var cache_Data_Array_NonEmpty_appendArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_appendArray sync.Once
func Get_Data_Array_NonEmpty_appendArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_appendArray.Do(func() {
		cache_Data_Array_NonEmpty_appendArray = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, ys_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_appendArray(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(ys_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_appendArray
}

var cache_Data_Array_NonEmpty_alterAt gopurs_runtime.Value
var once_Data_Array_NonEmpty_alterAt sync.Once
func Get_Data_Array_NonEmpty_alterAt() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_alterAt.Do(func() {
		cache_Data_Array_NonEmpty_alterAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_NonEmpty_alterAt(i_0_box.IntVal, f_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_NonEmpty_alterAt
}

var cache_Data_Array_NonEmpty_adaptMaybe gopurs_runtime.Value
var once_Data_Array_NonEmpty_adaptMaybe sync.Once
func Get_Data_Array_NonEmpty_adaptMaybe() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_adaptMaybe.Do(func() {
		cache_Data_Array_NonEmpty_adaptMaybe = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_adaptMaybe(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_NonEmpty_adaptMaybe
}

var cache_Data_Array_NonEmpty_head gopurs_runtime.Value
var once_Data_Array_NonEmpty_head sync.Once
func Get_Data_Array_NonEmpty_head() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_head.Do(func() {
		cache_Data_Array_NonEmpty_head = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_head(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_NonEmpty_head
}

var cache_Data_Array_NonEmpty_go__init gopurs_runtime.Value
var once_Data_Array_NonEmpty_go__init sync.Once
func Get_Data_Array_NonEmpty_go__init() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_go__init.Do(func() {
		cache_Data_Array_NonEmpty_go__init = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_go__init(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_go__init
}

var cache_Data_Array_NonEmpty_last gopurs_runtime.Value
var once_Data_Array_NonEmpty_last sync.Once
func Get_Data_Array_NonEmpty_last() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_last.Do(func() {
		cache_Data_Array_NonEmpty_last = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_last(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_NonEmpty_last
}

var cache_Data_Array_NonEmpty_tail gopurs_runtime.Value
var once_Data_Array_NonEmpty_tail sync.Once
func Get_Data_Array_NonEmpty_tail() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_tail.Do(func() {
		cache_Data_Array_NonEmpty_tail = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_tail(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_tail
}

var cache_Data_Array_NonEmpty_uncons gopurs_runtime.Value
var once_Data_Array_NonEmpty_uncons sync.Once
func Get_Data_Array_NonEmpty_uncons() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_uncons.Do(func() {
		cache_Data_Array_NonEmpty_uncons = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := Call_Data_Array_NonEmpty_uncons(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				_ = orig
				return gopurs_runtime.RecordDict([]string{"head", "tail"}, []gopurs_runtime.Value{orig.head, gopurs_runtime.Array(orig.tail)})
				}()
})
	})
	return cache_Data_Array_NonEmpty_uncons
}

var cache_Data_Array_NonEmpty_toNonEmpty gopurs_runtime.Value
var once_Data_Array_NonEmpty_toNonEmpty sync.Once
func Get_Data_Array_NonEmpty_toNonEmpty() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_toNonEmpty.Do(func() {
		cache_Data_Array_NonEmpty_toNonEmpty = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Call_Data_Array_NonEmpty_toNonEmpty(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_NonEmpty_toNonEmpty
}

var cache_Data_Array_NonEmpty_unsnoc gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsnoc sync.Once
func Get_Data_Array_NonEmpty_unsnoc() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsnoc.Do(func() {
		cache_Data_Array_NonEmpty_unsnoc = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := Call_Data_Array_NonEmpty_unsnoc(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				_ = orig
				return gopurs_runtime.RecordDict([]string{"init", "last"}, []gopurs_runtime.Value{gopurs_runtime.Array(orig.go__init), orig.last})
				}()
})
	})
	return cache_Data_Array_NonEmpty_unsnoc
}

var cache_Data_Array_NonEmpty_adaptAny gopurs_runtime.Value
var once_Data_Array_NonEmpty_adaptAny sync.Once
func Get_Data_Array_NonEmpty_adaptAny() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_adaptAny.Do(func() {
		cache_Data_Array_NonEmpty_adaptAny = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_adaptAny(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_NonEmpty_adaptAny
}

var cache_Data_Array_NonEmpty_all gopurs_runtime.Value
var once_Data_Array_NonEmpty_all sync.Once
func Get_Data_Array_NonEmpty_all() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_all.Do(func() {
		cache_Data_Array_NonEmpty_all = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Array_NonEmpty_all(p_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_all
}

var cache_Data_Array_NonEmpty_any gopurs_runtime.Value
var once_Data_Array_NonEmpty_any sync.Once
func Get_Data_Array_NonEmpty_any() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_any.Do(func() {
		cache_Data_Array_NonEmpty_any = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Array_NonEmpty_any(p_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_any
}

var cache_Data_Array_NonEmpty_catMaybes gopurs_runtime.Value
var once_Data_Array_NonEmpty_catMaybes sync.Once
func Get_Data_Array_NonEmpty_catMaybes() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_catMaybes.Do(func() {
		cache_Data_Array_NonEmpty_catMaybes = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_catMaybes(func() []*Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]*Constructor_Data_Maybe_Just[gopurs_runtime.Value], len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v) }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_catMaybes
}

var cache_Data_Array_NonEmpty_go__delete gopurs_runtime.Value
var once_Data_Array_NonEmpty_go__delete sync.Once
func Get_Data_Array_NonEmpty_go__delete() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_go__delete.Do(func() {
		cache_Data_Array_NonEmpty_go__delete = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_go__delete(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_go__delete
}

var cache_Data_Array_NonEmpty_deleteAt gopurs_runtime.Value
var once_Data_Array_NonEmpty_deleteAt sync.Once
func Get_Data_Array_NonEmpty_deleteAt() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_deleteAt.Do(func() {
		cache_Data_Array_NonEmpty_deleteAt = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_NonEmpty_deleteAt(i_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_NonEmpty_deleteAt
}

var cache_Data_Array_NonEmpty_deleteBy gopurs_runtime.Value
var once_Data_Array_NonEmpty_deleteBy sync.Once
func Get_Data_Array_NonEmpty_deleteBy() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_deleteBy.Do(func() {
		cache_Data_Array_NonEmpty_deleteBy = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_deleteBy(f_0_box, x_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_deleteBy
}

var cache_Data_Array_NonEmpty_difference gopurs_runtime.Value
var once_Data_Array_NonEmpty_difference sync.Once
func Get_Data_Array_NonEmpty_difference() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_difference.Do(func() {
		cache_Data_Array_NonEmpty_difference = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_difference(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_NonEmpty_difference
}

var cache_Data_Array_NonEmpty_drop gopurs_runtime.Value
var once_Data_Array_NonEmpty_drop sync.Once
func Get_Data_Array_NonEmpty_drop() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_drop.Do(func() {
		cache_Data_Array_NonEmpty_drop = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_drop(i_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_drop
}

var cache_Data_Array_NonEmpty_dropEnd gopurs_runtime.Value
var once_Data_Array_NonEmpty_dropEnd sync.Once
func Get_Data_Array_NonEmpty_dropEnd() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_dropEnd.Do(func() {
		cache_Data_Array_NonEmpty_dropEnd = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_dropEnd(i_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_dropEnd
}

var cache_Data_Array_NonEmpty_dropWhile gopurs_runtime.Value
var once_Data_Array_NonEmpty_dropWhile sync.Once
func Get_Data_Array_NonEmpty_dropWhile() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_dropWhile.Do(func() {
		cache_Data_Array_NonEmpty_dropWhile = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_dropWhile(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_dropWhile
}

var cache_Data_Array_NonEmpty_elem gopurs_runtime.Value
var once_Data_Array_NonEmpty_elem sync.Once
func Get_Data_Array_NonEmpty_elem() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_elem.Do(func() {
		cache_Data_Array_NonEmpty_elem = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Array_NonEmpty_elem(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_elem
}

var cache_Data_Array_NonEmpty_elemIndex gopurs_runtime.Value
var once_Data_Array_NonEmpty_elemIndex sync.Once
func Get_Data_Array_NonEmpty_elemIndex() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_elemIndex.Do(func() {
		cache_Data_Array_NonEmpty_elemIndex = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_NonEmpty_elemIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_NonEmpty_elemIndex
}

var cache_Data_Array_NonEmpty_elemLastIndex gopurs_runtime.Value
var once_Data_Array_NonEmpty_elemLastIndex sync.Once
func Get_Data_Array_NonEmpty_elemLastIndex() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_elemLastIndex.Do(func() {
		cache_Data_Array_NonEmpty_elemLastIndex = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_NonEmpty_elemLastIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_NonEmpty_elemLastIndex
}

var cache_Data_Array_NonEmpty_filter gopurs_runtime.Value
var once_Data_Array_NonEmpty_filter sync.Once
func Get_Data_Array_NonEmpty_filter() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_filter.Do(func() {
		cache_Data_Array_NonEmpty_filter = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_filter(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_filter
}

var cache_Data_Array_NonEmpty_filterA gopurs_runtime.Value
var once_Data_Array_NonEmpty_filterA sync.Once
func Get_Data_Array_NonEmpty_filterA() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_filterA.Do(func() {
		cache_Data_Array_NonEmpty_filterA = gopurs_runtime.Func2(func(dictApplicative_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_filterA(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), f_1_box)
})
	})
	return cache_Data_Array_NonEmpty_filterA
}

var cache_Data_Array_NonEmpty_find gopurs_runtime.Value
var once_Data_Array_NonEmpty_find sync.Once
func Get_Data_Array_NonEmpty_find() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_find.Do(func() {
		cache_Data_Array_NonEmpty_find = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_NonEmpty_find(p_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_NonEmpty_find
}

var cache_Data_Array_NonEmpty_findIndex gopurs_runtime.Value
var once_Data_Array_NonEmpty_findIndex sync.Once
func Get_Data_Array_NonEmpty_findIndex() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_findIndex.Do(func() {
		cache_Data_Array_NonEmpty_findIndex = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_NonEmpty_findIndex(p_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_NonEmpty_findIndex
}

var cache_Data_Array_NonEmpty_findLastIndex gopurs_runtime.Value
var once_Data_Array_NonEmpty_findLastIndex sync.Once
func Get_Data_Array_NonEmpty_findLastIndex() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_findLastIndex.Do(func() {
		cache_Data_Array_NonEmpty_findLastIndex = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_NonEmpty_findLastIndex(x_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_NonEmpty_findLastIndex
}

var cache_Data_Array_NonEmpty_findMap gopurs_runtime.Value
var once_Data_Array_NonEmpty_findMap sync.Once
func Get_Data_Array_NonEmpty_findMap() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_findMap.Do(func() {
		cache_Data_Array_NonEmpty_findMap = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_NonEmpty_findMap(p_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_NonEmpty_findMap
}

var cache_Data_Array_NonEmpty_foldM gopurs_runtime.Value
var once_Data_Array_NonEmpty_foldM sync.Once
func Get_Data_Array_NonEmpty_foldM() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_foldM.Do(func() {
		cache_Data_Array_NonEmpty_foldM = gopurs_runtime.Func3(func(dictMonad_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, acc_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_foldM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box), f_1_box, acc_2_box)
})
	})
	return cache_Data_Array_NonEmpty_foldM
}

var cache_Data_Array_NonEmpty_foldRecM gopurs_runtime.Value
var once_Data_Array_NonEmpty_foldRecM sync.Once
func Get_Data_Array_NonEmpty_foldRecM() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_foldRecM.Do(func() {
		cache_Data_Array_NonEmpty_foldRecM = gopurs_runtime.Func3(func(dictMonadRec_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, acc_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_foldRecM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), f_1_box, acc_2_box)
})
	})
	return cache_Data_Array_NonEmpty_foldRecM
}

var cache_Data_Array_NonEmpty_index gopurs_runtime.Value
var once_Data_Array_NonEmpty_index sync.Once
func Get_Data_Array_NonEmpty_index() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_index.Do(func() {
		cache_Data_Array_NonEmpty_index = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_NonEmpty_index(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_NonEmpty_index
}

var cache_Data_Array_NonEmpty_length gopurs_runtime.Value
var once_Data_Array_NonEmpty_length sync.Once
func Get_Data_Array_NonEmpty_length() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_length.Do(func() {
		cache_Data_Array_NonEmpty_length = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Array_NonEmpty_length(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_length
}

var cache_Data_Array_NonEmpty_mapMaybe gopurs_runtime.Value
var once_Data_Array_NonEmpty_mapMaybe sync.Once
func Get_Data_Array_NonEmpty_mapMaybe() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_mapMaybe.Do(func() {
		cache_Data_Array_NonEmpty_mapMaybe = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_mapMaybe(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_mapMaybe
}

var cache_Data_Array_NonEmpty_notElem gopurs_runtime.Value
var once_Data_Array_NonEmpty_notElem sync.Once
func Get_Data_Array_NonEmpty_notElem() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_notElem.Do(func() {
		cache_Data_Array_NonEmpty_notElem = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Array_NonEmpty_notElem(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_notElem
}

var cache_Data_Array_NonEmpty_partition gopurs_runtime.Value
var once_Data_Array_NonEmpty_partition sync.Once
func Get_Data_Array_NonEmpty_partition() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_partition.Do(func() {
		cache_Data_Array_NonEmpty_partition = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := Call_Data_Array_NonEmpty_partition(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				_ = orig
				return gopurs_runtime.RecordDict([]string{"no", "yes"}, []gopurs_runtime.Value{gopurs_runtime.Array(orig.no), gopurs_runtime.Array(orig.yes)})
				}()
})
	})
	return cache_Data_Array_NonEmpty_partition
}

var cache_Data_Array_NonEmpty_slice gopurs_runtime.Value
var once_Data_Array_NonEmpty_slice sync.Once
func Get_Data_Array_NonEmpty_slice() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_slice.Do(func() {
		cache_Data_Array_NonEmpty_slice = gopurs_runtime.Func3(func(start_0_box gopurs_runtime.Value, end_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_slice(start_0_box.IntVal, end_1_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_slice
}

var cache_Data_Array_NonEmpty_span gopurs_runtime.Value
var once_Data_Array_NonEmpty_span sync.Once
func Get_Data_Array_NonEmpty_span() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_span.Do(func() {
		cache_Data_Array_NonEmpty_span = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := Call_Data_Array_NonEmpty_span(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				_ = orig
				return gopurs_runtime.RecordDict([]string{"init", "rest"}, []gopurs_runtime.Value{gopurs_runtime.Array(orig.go__init), gopurs_runtime.Array(orig.rest)})
				}()
})
	})
	return cache_Data_Array_NonEmpty_span
}

var cache_Data_Array_NonEmpty_take gopurs_runtime.Value
var once_Data_Array_NonEmpty_take sync.Once
func Get_Data_Array_NonEmpty_take() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_take.Do(func() {
		cache_Data_Array_NonEmpty_take = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_take(i_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_take
}

var cache_Data_Array_NonEmpty_takeEnd gopurs_runtime.Value
var once_Data_Array_NonEmpty_takeEnd sync.Once
func Get_Data_Array_NonEmpty_takeEnd() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_takeEnd.Do(func() {
		cache_Data_Array_NonEmpty_takeEnd = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_takeEnd(i_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_takeEnd
}

var cache_Data_Array_NonEmpty_takeWhile gopurs_runtime.Value
var once_Data_Array_NonEmpty_takeWhile sync.Once
func Get_Data_Array_NonEmpty_takeWhile() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_takeWhile.Do(func() {
		cache_Data_Array_NonEmpty_takeWhile = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_takeWhile(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_takeWhile
}

var cache_Data_Array_NonEmpty_toUnfoldable gopurs_runtime.Value
var once_Data_Array_NonEmpty_toUnfoldable sync.Once
func Get_Data_Array_NonEmpty_toUnfoldable() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_toUnfoldable.Do(func() {
		cache_Data_Array_NonEmpty_toUnfoldable = gopurs_runtime.Func2(func(dictUnfoldable_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_NonEmpty_toUnfoldable
}

var cache_Data_Array_NonEmpty_unsafeAdapt gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeAdapt sync.Once
func Get_Data_Array_NonEmpty_unsafeAdapt() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeAdapt.Do(func() {
		cache_Data_Array_NonEmpty_unsafeAdapt = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_unsafeAdapt(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_unsafeAdapt
}

var cache_Data_Array_NonEmpty_cons gopurs_runtime.Value
var once_Data_Array_NonEmpty_cons sync.Once
func Get_Data_Array_NonEmpty_cons() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_cons.Do(func() {
		cache_Data_Array_NonEmpty_cons = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_cons(x_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_cons
}

var cache_Data_Array_NonEmpty_group gopurs_runtime.Value
var once_Data_Array_NonEmpty_group sync.Once
func Get_Data_Array_NonEmpty_group() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_group.Do(func() {
		cache_Data_Array_NonEmpty_group = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_group(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_Data_Array_NonEmpty_group
}

var cache_Data_Array_NonEmpty_groupAllBy gopurs_runtime.Value
var once_Data_Array_NonEmpty_groupAllBy sync.Once
func Get_Data_Array_NonEmpty_groupAllBy() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_groupAllBy.Do(func() {
		cache_Data_Array_NonEmpty_groupAllBy = gopurs_runtime.Func2(func(op_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_groupAllBy(op_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_groupAllBy
}

var cache_Data_Array_NonEmpty_groupAll gopurs_runtime.Value
var once_Data_Array_NonEmpty_groupAll sync.Once
func Get_Data_Array_NonEmpty_groupAll() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_groupAll.Do(func() {
		cache_Data_Array_NonEmpty_groupAll = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_groupAll(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Array_NonEmpty_groupAll
}

var cache_Data_Array_NonEmpty_groupBy gopurs_runtime.Value
var once_Data_Array_NonEmpty_groupBy sync.Once
func Get_Data_Array_NonEmpty_groupBy() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_groupBy.Do(func() {
		cache_Data_Array_NonEmpty_groupBy = gopurs_runtime.Func2(func(op_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_groupBy(op_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_groupBy
}

var cache_Data_Array_NonEmpty_insert gopurs_runtime.Value
var once_Data_Array_NonEmpty_insert sync.Once
func Get_Data_Array_NonEmpty_insert() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_insert.Do(func() {
		cache_Data_Array_NonEmpty_insert = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_insert(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), x_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_insert
}

var cache_Data_Array_NonEmpty_insertBy gopurs_runtime.Value
var once_Data_Array_NonEmpty_insertBy sync.Once
func Get_Data_Array_NonEmpty_insertBy() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_insertBy.Do(func() {
		cache_Data_Array_NonEmpty_insertBy = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_insertBy(f_0_box, x_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_insertBy
}

var cache_Data_Array_NonEmpty_intersperse gopurs_runtime.Value
var once_Data_Array_NonEmpty_intersperse sync.Once
func Get_Data_Array_NonEmpty_intersperse() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_intersperse.Do(func() {
		cache_Data_Array_NonEmpty_intersperse = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_intersperse(x_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_intersperse
}

var cache_Data_Array_NonEmpty_mapWithIndex gopurs_runtime.Value
var once_Data_Array_NonEmpty_mapWithIndex sync.Once
func Get_Data_Array_NonEmpty_mapWithIndex() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_mapWithIndex.Do(func() {
		cache_Data_Array_NonEmpty_mapWithIndex = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_mapWithIndex(f_0_box)
})
	})
	return cache_Data_Array_NonEmpty_mapWithIndex
}

var cache_Data_Array_NonEmpty_modifyAtIndices gopurs_runtime.Value
var once_Data_Array_NonEmpty_modifyAtIndices sync.Once
func Get_Data_Array_NonEmpty_modifyAtIndices() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_modifyAtIndices.Do(func() {
		cache_Data_Array_NonEmpty_modifyAtIndices = gopurs_runtime.Func4(func(dictFoldable_0_box gopurs_runtime.Value, is_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_modifyAtIndices(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), is_1_box, f_2_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_3_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_modifyAtIndices
}

var cache_Data_Array_NonEmpty_nub gopurs_runtime.Value
var once_Data_Array_NonEmpty_nub sync.Once
func Get_Data_Array_NonEmpty_nub() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_nub.Do(func() {
		cache_Data_Array_NonEmpty_nub = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_nub(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_nub
}

var cache_Data_Array_NonEmpty_nubBy gopurs_runtime.Value
var once_Data_Array_NonEmpty_nubBy sync.Once
func Get_Data_Array_NonEmpty_nubBy() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_nubBy.Do(func() {
		cache_Data_Array_NonEmpty_nubBy = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_nubBy(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_nubBy
}

var cache_Data_Array_NonEmpty_nubByEq gopurs_runtime.Value
var once_Data_Array_NonEmpty_nubByEq sync.Once
func Get_Data_Array_NonEmpty_nubByEq() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_nubByEq.Do(func() {
		cache_Data_Array_NonEmpty_nubByEq = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_nubByEq(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_nubByEq
}

var cache_Data_Array_NonEmpty_nubEq gopurs_runtime.Value
var once_Data_Array_NonEmpty_nubEq sync.Once
func Get_Data_Array_NonEmpty_nubEq() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_nubEq.Do(func() {
		cache_Data_Array_NonEmpty_nubEq = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_nubEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_nubEq
}

var cache_Data_Array_NonEmpty_reverse gopurs_runtime.Value
var once_Data_Array_NonEmpty_reverse sync.Once
func Get_Data_Array_NonEmpty_reverse() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_reverse.Do(func() {
		cache_Data_Array_NonEmpty_reverse = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_reverse(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_reverse
}

var cache_Data_Array_NonEmpty_scanl gopurs_runtime.Value
var once_Data_Array_NonEmpty_scanl sync.Once
func Get_Data_Array_NonEmpty_scanl() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_scanl.Do(func() {
		cache_Data_Array_NonEmpty_scanl = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_scanl(f_0_box, x_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_scanl
}

var cache_Data_Array_NonEmpty_scanr gopurs_runtime.Value
var once_Data_Array_NonEmpty_scanr sync.Once
func Get_Data_Array_NonEmpty_scanr() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_scanr.Do(func() {
		cache_Data_Array_NonEmpty_scanr = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_scanr(f_0_box, x_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_scanr
}

var cache_Data_Array_NonEmpty_sort gopurs_runtime.Value
var once_Data_Array_NonEmpty_sort sync.Once
func Get_Data_Array_NonEmpty_sort() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_sort.Do(func() {
		cache_Data_Array_NonEmpty_sort = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_sort(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Array_NonEmpty_sort
}

var cache_Data_Array_NonEmpty_sortBy gopurs_runtime.Value
var once_Data_Array_NonEmpty_sortBy sync.Once
func Get_Data_Array_NonEmpty_sortBy() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_sortBy.Do(func() {
		cache_Data_Array_NonEmpty_sortBy = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_sortBy(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_sortBy
}

var cache_Data_Array_NonEmpty_sortWith gopurs_runtime.Value
var once_Data_Array_NonEmpty_sortWith sync.Once
func Get_Data_Array_NonEmpty_sortWith() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_sortWith.Do(func() {
		cache_Data_Array_NonEmpty_sortWith = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_sortWith(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), f_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_sortWith
}

var cache_Data_Array_NonEmpty_updateAtIndices gopurs_runtime.Value
var once_Data_Array_NonEmpty_updateAtIndices sync.Once
func Get_Data_Array_NonEmpty_updateAtIndices() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_updateAtIndices.Do(func() {
		cache_Data_Array_NonEmpty_updateAtIndices = gopurs_runtime.Func3(func(dictFoldable_0_box gopurs_runtime.Value, pairs_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_updateAtIndices(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), pairs_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_updateAtIndices
}

var cache_Data_Array_NonEmpty_unsafeIndex gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeIndex sync.Once
func Get_Data_Array_NonEmpty_unsafeIndex() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeIndex.Do(func() {
		cache_Data_Array_NonEmpty_unsafeIndex = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_unsafeIndex(_dollar___unused_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), __local_var_2_box.IntVal)
})
	})
	return cache_Data_Array_NonEmpty_unsafeIndex
}

var cache_Data_Array_NonEmpty_toUnfoldable1 gopurs_runtime.Value
var once_Data_Array_NonEmpty_toUnfoldable1 sync.Once
func Get_Data_Array_NonEmpty_toUnfoldable1() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_toUnfoldable1.Do(func() {
		cache_Data_Array_NonEmpty_toUnfoldable1 = gopurs_runtime.Func2(func(dictUnfoldable1_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_toUnfoldable1(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_0_box), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_NonEmpty_toUnfoldable1
}

func Call_Data_Array_NonEmpty_intercalate1(dictSemigroup_0_loop *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictSemigroup_0 *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] = dictSemigroup_0_loop
_ = dictSemigroup_0
// TAST (Let): semigroupJoinWith1_1_0 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(Func [(TypeVar m)] (TypeVar m))])
semigroupJoinWith1_1_0 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value, j_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_0.V0), gopurs_runtime.Apply(v_1, j_3), gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_0.V0), j_3, gopurs_runtime.Apply(v1_2, j_3)))
})})
_ = semigroupJoinWith1_1_0
return gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, foldable_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[[]gopurs_runtime.Value]](Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray()).V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupJoinWith1_1_0)}, gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return x_4
}), foldable_3, a_2)
})
}

func Call_Data_Array_NonEmpty_fromJust(v_0_loop *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0 != nil) {
__t0 = (v_0).V0
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

func Call_Data_Array_NonEmpty_unsafeIndex1(__local_var_0_loop []gopurs_runtime.Value, __local_var_1_loop int64) gopurs_runtime.Value {
var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 int64 = __local_var_1_loop
_ = __local_var_1
return __local_var_0[__local_var_1]
}

func Call_Data_Array_NonEmpty_unsafeFromArray__1127564228(__eta_norm_0_0_loop []int64) []int64 {
unsafeFromArray__1127564228:
for {
if false { continue unsafeFromArray__1127564228 }
var __eta_norm_0_0 []int64 = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := __eta_norm_0_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
}
}

func Call_Data_Array_NonEmpty_transpose(x_0_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var x_0 []gopurs_runtime.Value = x_0_loop
_ = x_0
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_Array_transpose(), gopurs_runtime.Array(x_0)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_toArray(v_0_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var v_0 []gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Array_NonEmpty_unionBy_prime_(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(Get_Data_Array_unionBy(), eq_0, gopurs_runtime.Array(xs_1), gopurs_runtime.Array(x_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_union_prime_(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_Array_NonEmpty_unionBy_prime_(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_Array_NonEmpty_unionBy(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(Get_Data_Array_unionBy(), eq_0, gopurs_runtime.Array(xs_1), gopurs_runtime.Array(x_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_union(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_Array_NonEmpty_unionBy(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_Array_NonEmpty_unzip(x_0_loop []*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
var x_0 []*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(TypeVar c)
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_Array_unzip(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := x_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
_ = __local_var_1_0
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V1.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
				_p := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{V0: _p.V0, V1: _p.V1}
			}()
}

func Call_Data_Array_NonEmpty_updateAt(i_0_loop int64, x_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var i_0 int64 = i_0_loop
_ = i_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.UncurriedApp5(Get_Data_Array__updateAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Int(i_0), x_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_NonEmpty_zip(xs_0_loop []gopurs_runtime.Value, ys_1_loop []gopurs_runtime.Value) []*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
var ys_1 []gopurs_runtime.Value = ys_1_loop
_ = ys_1
return func() []*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_zipWithImpl(), Get_Data_Tuple_Tuple(), gopurs_runtime.Array(xs_0), gopurs_runtime.Array(ys_1)).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v) }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_zipWith(f_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value, ys_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 []gopurs_runtime.Value = ys_2_loop
_ = ys_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_zipWithImpl(), f_0, gopurs_runtime.Array(xs_1), gopurs_runtime.Array(ys_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_zipWithA(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, xs_2_loop []gopurs_runtime.Value, ys_3_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 []gopurs_runtime.Value = xs_2_loop
_ = xs_2
var ys_3 []gopurs_runtime.Value = ys_3_loop
_ = ys_3
return gopurs_runtime.Apply4(Get_Data_Array_zipWithA(), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, f_1, gopurs_runtime.Array(xs_2), gopurs_runtime.Array(ys_3))
}

func Call_Data_Array_NonEmpty_splitAt(i_0_loop int64, xs_1_loop []gopurs_runtime.Value) struct{
	after []gopurs_runtime.Value
	before []gopurs_runtime.Value
} {
var i_0 int64 = i_0_loop
_ = i_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() struct{
	after []gopurs_runtime.Value
	before []gopurs_runtime.Value
} {
					orig := gopurs_runtime.Apply2(Get_Data_Array_splitAt(), gopurs_runtime.Int(i_0), gopurs_runtime.Array(xs_1))
					_ = orig
					clone := struct{
	after []gopurs_runtime.Value
	before []gopurs_runtime.Value
}{}
					clone.after = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "after").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
					clone.before = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "before").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
					return clone
				}()
}

func Call_Data_Array_NonEmpty_some(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): some1__193435443_1_0 shape=App(Var) bindingType=Any
some1__193435443_1_0 := gopurs_runtime.Apply(Get_Data_Array_some(), gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(dictAlternative_0)})
_ = some1__193435443_1_0
return gopurs_runtime.Func(func(dictLazy_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(Array (TypeVar a))]))
__local_var_3_1 := gopurs_runtime.Apply(some1__193435443_1_0, dictLazy_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, x_4)
})
})
}

func Call_Data_Array_NonEmpty_snoc_prime_(xs_0_loop []gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) []gopurs_runtime.Value {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(Get_Data_Array_ST_withArray(), gopurs_runtime.Apply(Get_Data_Array_ST_push(), x_1), gopurs_runtime.Array(xs_0))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_snoc(xs_0_loop []gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) []gopurs_runtime.Value {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(Get_Data_Array_ST_withArray(), gopurs_runtime.Apply(Get_Data_Array_ST_push(), x_1), gopurs_runtime.Array(xs_0))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_singleton(x_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{x_0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_replicate(i_0_loop int64, x_1_loop gopurs_runtime.Value) []gopurs_runtime.Value {
var i_0 int64 = i_0_loop
_ = i_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_Data_Array_replicateImpl(), gopurs_runtime.Apply2(Get_Data_Ord_max__3052583336(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(i_0)), x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_go__range(x_0_loop int64, y_1_loop int64) []int64 {
var x_0 int64 = x_0_loop
_ = x_0
var y_1 int64 = y_1_loop
_ = y_1
return func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_Data_Array_rangeImpl(), gopurs_runtime.Int(x_0), gopurs_runtime.Int(y_1)).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_prependArray(xs_0_loop []gopurs_runtime.Value, ys_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
var ys_1 []gopurs_runtime.Value = ys_1_loop
_ = ys_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array(xs_0), gopurs_runtime.Array(ys_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_modifyAt(i_0_loop int64, f_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var i_0 int64 = i_0_loop
_ = i_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply3(Get_Data_Array_modifyAt(), gopurs_runtime.Int(i_0), f_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_NonEmpty_intersectBy_prime_(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply2(Get_Data_Array_intersectBy(), eq_0, gopurs_runtime.Array(xs_1))
}

func Call_Data_Array_NonEmpty_intersectBy(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(Get_Data_Array_intersectBy(), eq_0, gopurs_runtime.Array(xs_1), gopurs_runtime.Array(x_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_intersect_prime_(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_Array_NonEmpty_intersectBy_prime_(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_Array_NonEmpty_intersect(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_Array_NonEmpty_intersectBy(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_Array_NonEmpty_intercalate(dictSemigroup_0_loop *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictSemigroup_0 *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] = dictSemigroup_0_loop
_ = dictSemigroup_0
// TAST (Let): semigroupJoinWith1_1_0 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(Func [(TypeVar m)] (TypeVar m))])
semigroupJoinWith1_1_0 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value, j_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_0.V0), gopurs_runtime.Apply(v_1, j_3), gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_0.V0), j_3, gopurs_runtime.Apply(v1_2, j_3)))
})})
_ = semigroupJoinWith1_1_0
return gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, foldable_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[[]gopurs_runtime.Value]](Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray()).V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupJoinWith1_1_0)}, gopurs_runtime.Func2(func(x_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return x_4
}), foldable_3, a_2)
})
}

func Call_Data_Array_NonEmpty_insertAt(i_0_loop int64, x_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var i_0 int64 = i_0_loop
_ = i_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.UncurriedApp5(Get_Data_Array__insertAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Int(i_0), x_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_NonEmpty_fromFoldable1(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
// TAST (Let): __local_var_1_1 shape=Other bindingType=(Func [(Func [(TypeVar a), (TypeVar b)] (TypeVar b)), (TypeVar b), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeVar b))
__local_var_1_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldable1_0.V0), gopurs_runtime.Value{}), "foldr")
_ = __local_var_1_1
// TAST (Let): __local_var_1_0 shape=Let(Abs(UncurriedApp(Var))) bindingType=(Func [(TypeApp (TypeVar f) [(TypeVar a)])] (Array (TypeVar a)))
__local_var_1_0 := gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_fromFoldableImpl(), __local_var_1_1, __local_var_2)
})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
}

func Call_Data_Array_NonEmpty_fromArray(xs_0_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
var __t0 *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]
{
if (gopurs_runtime.Int(int64(len(xs_0))).IntVal) > (int64(0)) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Array(xs_0), true}
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
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_2807397954_3094389156(__t0))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_NonEmpty_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
// TAST (Let): __local_var_1_1 shape=Other bindingType=(Func [(Func [(TypeVar a), (TypeVar b)] (TypeVar b)), (TypeVar b), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeVar b))
__local_var_1_1 := gopurs_runtime.Box(dictFoldable_0.V2)
_ = __local_var_1_1
// TAST (Let): __local_var_1_0 shape=Let(Abs(UncurriedApp(Var))) bindingType=(Func [(TypeApp (TypeVar f) [(TypeVar a)])] (Array (TypeVar a)))
__local_var_1_0 := gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_fromFoldableImpl(), __local_var_1_1, __local_var_2)
})
_ = __local_var_1_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 shape=App(Other) bindingType=Any
__local_var_3_2 := gopurs_runtime.Apply(__local_var_1_0, x_2)
_ = __local_var_3_2
var __t3 *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]
{
if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(__local_var_3_2))).IntVal) > (int64(0)) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_3_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), true}
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
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_2807397954_3094389156(__t3))}
})
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_NonEmpty_transpose_prime_(x_0_loop [][]gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var x_0 [][]gopurs_runtime.Value = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(TypeVar d)
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_Array_transpose(), func() gopurs_runtime.Value {
					arr := x_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}())
_ = __local_var_1_0
var __t1 *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]
{
if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(__local_var_1_0))).IntVal) > (int64(0)) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_0.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), true}
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
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_1:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_2807397954_3094389156(__t1))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_NonEmpty_foldMap1(dictSemigroup_0_loop *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictSemigroup_0 *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[[]gopurs_runtime.Value]](Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray()).V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(dictSemigroup_0)})
}

func Call_Data_Array_NonEmpty_fold1(dictSemigroup_0_loop *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictSemigroup_0 *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[[]gopurs_runtime.Value]](Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray()).V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(dictSemigroup_0)}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}

func Call_Data_Array_NonEmpty_difference_prime_(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
// TAST (Let): difference1_1_0 shape=App(Var) bindingType=(Func [(Array (TypeVar a)), (Array (TypeVar a))] (Array (TypeVar a)))
difference1_1_0 := gopurs_runtime.Apply(Get_Data_Foldable_foldrArray(), gopurs_runtime.Apply(Get_Data_Array_go__delete(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_0)}))
_ = difference1_1_0
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(difference1_1_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
}

func Call_Data_Array_NonEmpty_cons_prime_(x_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array([]gopurs_runtime.Value{x_0}), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_fromNonEmpty(v_0_loop *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]) []gopurs_runtime.Value {
var v_0 *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array([]gopurs_runtime.Value{(v_0).V0}), (v_0).V1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_concatMap(b_0_loop gopurs_runtime.Value, a_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 []gopurs_runtime.Value = a_1_loop
_ = a_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Control_Bind_arrayBind(), gopurs_runtime.Array(a_1), b_0).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_appendArray(xs_0_loop []gopurs_runtime.Value, ys_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
var ys_1 []gopurs_runtime.Value = ys_1_loop
_ = ys_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array(xs_0), gopurs_runtime.Array(ys_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_alterAt(i_0_loop int64, f_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var i_0 int64 = i_0_loop
_ = i_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply3(Get_Data_Array_alterAt(), gopurs_runtime.Int(i_0), f_1, gopurs_runtime.Array(x_2))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_NonEmpty_adaptMaybe(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=(TypeVar c)
__local_var_2_0 := gopurs_runtime.Apply(f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136 && __local_var_2_0.UnsafePtr != nil) {
__t1 = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_2_0.UnsafePtr).V0
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

func Call_Data_Array_NonEmpty_head(x_0_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 []gopurs_runtime.Value = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 shape=UncurriedApp(Var) bindingType=(TypeVar c)
__local_var_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Array(x_0), gopurs_runtime.Int(int64(0))))
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0 != nil) {
__t1 = (__local_var_1_0).V0
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

func Call_Data_Array_NonEmpty_go__init(x_0_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var x_0 []gopurs_runtime.Value = x_0_loop
_ = x_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Int(int64(len(x_0))).IntVal) == (int64(0)) {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int((gopurs_runtime.Int(int64(len(x_0))).IntVal) - (int64(1))), gopurs_runtime.Array(x_0))
}
end_branch_0:
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t0.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_last(x_0_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 []gopurs_runtime.Value = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 shape=UncurriedApp(Var) bindingType=(TypeVar c)
__local_var_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Array(x_0), gopurs_runtime.Int((gopurs_runtime.Int(int64(len(x_0))).IntVal) - (int64(1)))))
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0 != nil) {
__t1 = (__local_var_1_0).V0
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

func Call_Data_Array_NonEmpty_tail(x_0_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var x_0 []gopurs_runtime.Value = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 shape=UncurriedApp(Var) bindingType=(TypeVar c)
__local_var_1_0 := gopurs_runtime.UncurriedApp3(Get_Data_Array_unconsImpl(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_2807397954_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{xs_2, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}), gopurs_runtime.Array(x_0))
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136 && __local_var_1_0.UnsafePtr != nil) {
__t1 = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t1.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_uncons(x_0_loop []gopurs_runtime.Value) struct{
	head gopurs_runtime.Value
	tail []gopurs_runtime.Value
} {
var x_0 []gopurs_runtime.Value = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 shape=UncurriedApp(Var) bindingType=(TypeVar c)
__local_var_1_0 := gopurs_runtime.UncurriedApp3(Get_Data_Array_unconsImpl(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
}), gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_2762242827_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail []gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.RecordDict2("head", "tail", x_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}), gopurs_runtime.Array(x_0))
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136 && __local_var_1_0.UnsafePtr != nil) {
__t1 = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return func() struct{
	head gopurs_runtime.Value
	tail []gopurs_runtime.Value
} {
					orig := __t1
					_ = orig
					clone := struct{
	head gopurs_runtime.Value
	tail []gopurs_runtime.Value
}{}
					clone.head = gopurs_runtime.RecordGet(orig, "head")
					clone.tail = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "tail").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
					return clone
				}()
}

func Call_Data_Array_NonEmpty_toNonEmpty(x_0_loop []gopurs_runtime.Value) *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 []gopurs_runtime.Value = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(TypeVar c)
__local_var_1_0 := func() gopurs_runtime.Value {
				orig := Call_Data_Array_NonEmpty_uncons(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_0).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				_ = orig
				return gopurs_runtime.RecordDict([]string{"head", "tail"}, []gopurs_runtime.Value{orig.head, gopurs_runtime.Array(orig.tail)})
				}()
_ = __local_var_1_0
return (&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(__local_var_1_0, "head"), gopurs_runtime.RecordGet(__local_var_1_0, "tail")})
}

func Call_Data_Array_NonEmpty_unsnoc(x_0_loop []gopurs_runtime.Value) struct{
	go__init []gopurs_runtime.Value
	last gopurs_runtime.Value
} {
var x_0 []gopurs_runtime.Value = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(TypeVar c)
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_Array_unsnoc(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_0).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136 && __local_var_1_0.UnsafePtr != nil) {
__t1 = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return func() struct{
	go__init []gopurs_runtime.Value
	last gopurs_runtime.Value
} {
					orig := __t1
					_ = orig
					clone := struct{
	go__init []gopurs_runtime.Value
	last gopurs_runtime.Value
}{}
					clone.go__init = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "init").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
					clone.last = gopurs_runtime.RecordGet(orig, "last")
					return clone
				}()
}

func Call_Data_Array_NonEmpty_adaptAny(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Array(x_1))
}

func Call_Data_Array_NonEmpty_all(p_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) bool {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return (gopurs_runtime.UncurriedApp2(Get_Data_Array_allImpl(), p_0, gopurs_runtime.Array(x_1)).IntVal) != (0)
}

func Call_Data_Array_NonEmpty_any(p_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) bool {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return (gopurs_runtime.UncurriedApp2(Get_Data_Array_anyImpl(), p_0, gopurs_runtime.Array(x_1)).IntVal) != (0)
}

func Call_Data_Array_NonEmpty_catMaybes(x_0_loop []*Constructor_Data_Maybe_Just[gopurs_runtime.Value]) []gopurs_runtime.Value {
var x_0 []*Constructor_Data_Maybe_Just[gopurs_runtime.Value] = x_0_loop
_ = x_0
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Array_mapMaybe(), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := x_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_go__delete(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(Get_Data_Array_deleteBy(), gopurs_runtime.Box(dictEq_0.V0), x_1, gopurs_runtime.Array(x_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_deleteAt(i_0_loop int64, x_1_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var i_0 int64 = i_0_loop
_ = i_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.UncurriedApp4(Get_Data_Array__deleteAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Int(i_0), gopurs_runtime.Array(x_1))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_NonEmpty_deleteBy(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(Get_Data_Array_deleteBy(), f_0, x_1, gopurs_runtime.Array(x_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_difference(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(Func [(Array (TypeVar a))] (Array (TypeVar a)))
__local_var_2_0 := gopurs_runtime.Apply2(Get_Data_Foldable_foldrArray(), gopurs_runtime.Apply(Get_Data_Array_go__delete(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_0)}), gopurs_runtime.Array(xs_1))
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_3.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
}

func Call_Data_Array_NonEmpty_drop(i_0_loop int64, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var i_0 int64 = i_0_loop
_ = i_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
var __t0 []gopurs_runtime.Value
{
if (i_0) < (int64(1)) {
__t0 = x_1
goto end_branch_0
} else {

}
}
{
__t0 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(i_0), gopurs_runtime.Int(int64(len(x_1))), gopurs_runtime.Array(x_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_0:
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(__t0).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_dropEnd(i_0_loop int64, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var i_0 int64 = i_0_loop
_ = i_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Array_dropEnd(), gopurs_runtime.Int(i_0), gopurs_runtime.Array(x_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_dropWhile(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_Data_Array_span(), f_0, gopurs_runtime.Array(x_1)), "rest").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_elem(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) bool {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return (gopurs_runtime.Bool((gopurs_runtime.Apply3(Get_Data_Array_elem(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_0)}, x_1, gopurs_runtime.Array(x_2)).IntVal) != (0)).IntVal) != (0)
}

func Call_Data_Array_NonEmpty_elemIndex(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_1170268447_3094389156(Rebox_Data_Array_NonEmpty_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), v_3, x_1).IntVal) != (0))
}), gopurs_runtime.Array(x_2))))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_NonEmpty_elemLastIndex(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_1170268447_3094389156(Rebox_Data_Array_NonEmpty_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_findLastIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), v_3, x_1).IntVal) != (0))
}), gopurs_runtime.Array(x_2))))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_NonEmpty_filter(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
arr_val_filterImpl0 := x_1
_ = arr_val_filterImpl0
_ = arr_val_filterImpl0
arr_go_filterImpl0 := arr_val_filterImpl0
_ = arr_go_filterImpl0
res_go_filterImpl0 := make([]gopurs_runtime.Value, 0)
_ = res_go_filterImpl0
for _, v_filterImpl0 := range arr_go_filterImpl0 {
if gopurs_runtime.Apply(f_0, v_filterImpl0).BoolVal() {
res_go_filterImpl0 = append(res_go_filterImpl0, v_filterImpl0)
} else {

}
}
return res_go_filterImpl0
}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_filterA(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=(Func [(Array (TypeVar a))] (TypeApp (TypeVar f) [(Array (TypeVar a))]))
__local_var_2_0 := gopurs_runtime.Apply2(Get_Data_Array_filterA(), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, f_1)
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_3.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
}

func Call_Data_Array_NonEmpty_find(p_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply2(Get_Data_Array_find(), p_0, gopurs_runtime.Array(x_1))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_NonEmpty_findIndex(p_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, p_0, gopurs_runtime.Array(x_1))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_NonEmpty_findLastIndex(x_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.UncurriedApp4(Get_Data_Array_findLastIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, x_0, gopurs_runtime.Array(x_1))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_NonEmpty_findMap(p_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.UncurriedApp4(Get_Data_Array_findMapImpl(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, Get_Data_Maybe_isJust(), p_0, gopurs_runtime.Array(x_1))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_NonEmpty_foldM(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, acc_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var acc_2 gopurs_runtime.Value = acc_2_loop
_ = acc_2
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(Func [(Array (TypeVar a))] (TypeApp (TypeVar m) [(TypeVar b)]))
__local_var_3_0 := gopurs_runtime.Apply3(Get_Data_Array_foldM(), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(dictMonad_0)}, f_1, acc_2)
_ = __local_var_3_0
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_4.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
}

func Call_Data_Array_NonEmpty_foldRecM(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, acc_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var acc_2 gopurs_runtime.Value = acc_2_loop
_ = acc_2
// TAST (Let): __local_var_3_0 shape=App(Var) bindingType=(Func [(Array (TypeVar a))] (TypeApp (TypeVar m) [(TypeVar b)]))
__local_var_3_0 := gopurs_runtime.Apply3(Get_Data_Array_foldRecM(), gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(dictMonadRec_0)}, f_1, acc_2)
_ = __local_var_3_0
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_4.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
}

func Call_Data_Array_NonEmpty_index(x_0_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var x_0 []gopurs_runtime.Value = x_0_loop
_ = x_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply(Get_Data_Array_index(), gopurs_runtime.Array(x_0))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_NonEmpty_length(x_0_loop []gopurs_runtime.Value) int64 {
var x_0 []gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Int(int64(len(x_0))).IntVal
}

func Call_Data_Array_NonEmpty_mapMaybe(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Array_mapMaybe(), f_0, gopurs_runtime.Array(x_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_notElem(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) bool {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return (gopurs_runtime.Bool((gopurs_runtime.Apply3(Get_Data_Array_notElem(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_0)}, x_1, gopurs_runtime.Array(x_2)).IntVal) != (0)).IntVal) != (0)
}

func Call_Data_Array_NonEmpty_partition(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) struct{
	no []gopurs_runtime.Value
	yes []gopurs_runtime.Value
} {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() struct{
	no []gopurs_runtime.Value
	yes []gopurs_runtime.Value
} {
					orig := gopurs_runtime.UncurriedApp2(Get_Data_Array_partitionImpl(), f_0, gopurs_runtime.Array(x_1))
					_ = orig
					clone := struct{
	no []gopurs_runtime.Value
	yes []gopurs_runtime.Value
}{}
					clone.no = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "no").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
					clone.yes = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "yes").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
					return clone
				}()
}

func Call_Data_Array_NonEmpty_slice(start_0_loop int64, end_1_loop int64, x_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var start_0 int64 = start_0_loop
_ = start_0
var end_1 int64 = end_1_loop
_ = end_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(start_0), gopurs_runtime.Int(end_1), gopurs_runtime.Array(x_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_span(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) struct{
	go__init []gopurs_runtime.Value
	rest []gopurs_runtime.Value
} {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() struct{
	go__init []gopurs_runtime.Value
	rest []gopurs_runtime.Value
} {
					orig := gopurs_runtime.Apply2(Get_Data_Array_span(), f_0, gopurs_runtime.Array(x_1))
					_ = orig
					clone := struct{
	go__init []gopurs_runtime.Value
	rest []gopurs_runtime.Value
}{}
					clone.go__init = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "init").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
					clone.rest = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "rest").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
					return clone
				}()
}

func Call_Data_Array_NonEmpty_take(i_0_loop int64, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var i_0 int64 = i_0_loop
_ = i_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
var __t0 []gopurs_runtime.Value
{
if (i_0) < (int64(1)) {
__t0 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
goto end_branch_0
} else {

}
}
{
__t0 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(i_0), gopurs_runtime.Array(x_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_0:
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(__t0).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_takeEnd(i_0_loop int64, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var i_0 int64 = i_0_loop
_ = i_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Array_takeEnd(), gopurs_runtime.Int(i_0), gopurs_runtime.Array(x_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_takeWhile(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_Data_Array_span(), f_0, gopurs_runtime.Array(x_1)), "init").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value], x_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
// TAST (Let): len_2_0 shape=Other bindingType=Int
var len_2_0 int64 = gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Array(x_1)))).IntVal
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictUnfoldable_0.V1), gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]]
{
if (i_3.IntVal) < (len_2_0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()[i_3.IntVal], gopurs_runtime.Int((i_3.IntVal) + (int64(1)))}))}, true}
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
__t1 = Rebox_Data_Array_NonEmpty_3094389156_1413047506(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_1413047506_3094389156(__t1))}
}), gopurs_runtime.Int(int64(0)))
}

func Call_Data_Array_NonEmpty_unsafeAdapt(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_cons(x_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array([]gopurs_runtime.Value{x_0}), gopurs_runtime.Array(x_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_group(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
// TAST (Let): eq_1_1 shape=Other bindingType=(Func [(TypeVar a), (TypeVar a)] Boolean)
eq_1_1 := gopurs_runtime.Box(dictEq_0.V0)
_ = eq_1_1
// TAST (Let): __local_var_1_0 shape=Let(Abs(App(Var))) bindingType=(Func [(Array (TypeVar a))] (Array (Array (TypeVar a))))
__local_var_1_0 := gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Array_groupBy(), eq_1_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}() }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
}

func Call_Data_Array_NonEmpty_groupAllBy(op_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Array_groupBy(), gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}), gopurs_runtime.Apply2(Get_Data_Array_sortBy(), op_0, gopurs_runtime.Array(x_1))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_groupAll(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Array_NonEmpty_groupAllBy(), gopurs_runtime.Box(dictOrd_0.V1))
}

func Call_Data_Array_NonEmpty_groupBy(op_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Array_groupBy(), op_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_insert(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(Get_Data_Array_insertBy(), gopurs_runtime.Box(dictOrd_0.V1), x_1, gopurs_runtime.Array(x_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_insertBy(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(Get_Data_Array_insertBy(), f_0, x_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_intersperse(x_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Array_intersperse(), x_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_mapWithIndex(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [(Array (TypeVar a))] (Array (TypeVar b)))
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_FunctorWithIndex_mapWithIndexArray(), f_0)
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
}

func Call_Data_Array_NonEmpty_modifyAtIndices(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], is_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value, x_3_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var is_1 gopurs_runtime.Value = is_1_loop
_ = is_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
var x_3 []gopurs_runtime.Value = x_3_loop
_ = x_3
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply4(Get_Data_Array_modifyAtIndices(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_0)}, is_1, f_2, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_3).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_nub(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Array_nubBy(), gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Array(x_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_nubBy(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Array_nubBy(), f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_nubByEq(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Array_nubByEq(), f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_nubEq(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Array_nubByEq(), gopurs_runtime.Box(dictEq_0.V0), gopurs_runtime.Array(x_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_reverse(x_0_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var x_0 []gopurs_runtime.Value = x_0_loop
_ = x_0
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_Array_reverse(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_0).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_scanl(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_scanlImpl(), f_0, x_1, gopurs_runtime.Array(x_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_scanr(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_scanrImpl(), f_0, x_1, gopurs_runtime.Array(x_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_sort(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 shape=Other bindingType=(Func [(TypeVar a), (TypeVar a)] (ADT ["Data","Ordering","Ordering"] []))
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Array_sortBy(), compare_1_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
}

func Call_Data_Array_NonEmpty_sortBy(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Array_sortBy(), f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_sortWith(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(Get_Data_Array_sortWith(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, f_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_updateAtIndices(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], pairs_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var pairs_1 gopurs_runtime.Value = pairs_1_loop
_ = pairs_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(Get_Data_Array_updateAtIndices(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_0)}, pairs_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_unsafeIndex(_dollar___unused_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value, __local_var_2_loop int64) gopurs_runtime.Value {
var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
_ = _dollar___unused_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
var __local_var_2 int64 = __local_var_2_loop
_ = __local_var_2
return gopurs_runtime.ArrayAccess(gopurs_runtime.Array(x_1), int(__local_var_2))
}

func Call_Data_Array_NonEmpty_toUnfoldable1(dictUnfoldable1_0_loop *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value], xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable1_0 *Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
// TAST (Let): len_2_0 shape=Other bindingType=(TypeVar d)
len_2_0 := gopurs_runtime.Int(int64(len(xs_1)))
_ = len_2_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictUnfoldable1_0.V0), gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (i_3.IntVal) < ((len_2_0.IntVal) - (int64(1))) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int((i_3.IntVal) + (int64(1)))}))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_1170268447_3094389156(Rebox_Data_Array_NonEmpty_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))))}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.ArrayAccess(gopurs_runtime.Array(xs_1), int(i_3.IntVal)), __t1}))}
}), gopurs_runtime.Int(int64(0)))
}

func Rebox_Data_Array_NonEmpty_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_Array_NonEmpty_138441832_3415943795(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]{}
		out.V0 = in.V0
		out.V1 = in.V1.IntVal
	return out
}

func Rebox_Data_Array_NonEmpty_1413047506_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_3415943795_138441832(in.V0))}
	return out
}

func Rebox_Data_Array_NonEmpty_2762242827_3094389156(in *Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail []gopurs_runtime.Value
}]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = func() gopurs_runtime.Value {
				orig := in.V0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"head", "tail"}, []gopurs_runtime.Value{orig.head, gopurs_runtime.Array(orig.tail)})
				}()
	return out
}

func Rebox_Data_Array_NonEmpty_2807397954_3094389156(in *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Array(in.V0)
	return out
}

func Rebox_Data_Array_NonEmpty_3016437835_4151366573(in *Constructor_Data_Semigroup_Foldable_Foldable1[[]gopurs_runtime.Value]) *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Array_NonEmpty_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[int64]{}
		out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_Array_NonEmpty_3094389156_1413047506(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]]{}
		out.V0 = Rebox_Data_Array_NonEmpty_138441832_3415943795(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
	return out
}

func Rebox_Data_Array_NonEmpty_3415943795_138441832(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Int(in.V1)
	return out
}



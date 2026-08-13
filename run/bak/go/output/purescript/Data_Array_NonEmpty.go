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
return Call_Data_Array_NonEmpty_intercalate1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dictSemigroup_0_box))
})
	})
	return cache_Data_Array_NonEmpty_intercalate1
}

var cache_Data_Array_NonEmpty_foldMap11 gopurs_runtime.Value
var once_Data_Array_NonEmpty_foldMap11 sync.Once
func Get_Data_Array_NonEmpty_foldMap11() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_foldMap11.Do(func() {
		cache_Data_Array_NonEmpty_foldMap11 = gopurs_runtime.RecordGet(Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray(), "foldMap1")
	})
	return cache_Data_Array_NonEmpty_foldMap11
}

var cache_Data_Array_NonEmpty_fold11 gopurs_runtime.Value
var once_Data_Array_NonEmpty_fold11 sync.Once
func Get_Data_Array_NonEmpty_fold11() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_fold11.Do(func() {
		cache_Data_Array_NonEmpty_fold11 = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_fold11(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dictSemigroup_0_box))
})
	})
	return cache_Data_Array_NonEmpty_fold11
}

var cache_Data_Array_NonEmpty_fromJust gopurs_runtime.Value
var once_Data_Array_NonEmpty_fromJust sync.Once
func Get_Data_Array_NonEmpty_fromJust() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_fromJust.Do(func() {
		cache_Data_Array_NonEmpty_fromJust = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_fromJust(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box))
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

var cache_Data_Array_NonEmpty_unionBy_prime gopurs_runtime.Value
var once_Data_Array_NonEmpty_unionBy_prime sync.Once
func Get_Data_Array_NonEmpty_unionBy_prime() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unionBy_prime.Do(func() {
		cache_Data_Array_NonEmpty_unionBy_prime = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_unionBy_prime(eq_0_box, func() []gopurs_runtime.Value {
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
	return cache_Data_Array_NonEmpty_unionBy_prime
}

var cache_Data_Array_NonEmpty_union_prime gopurs_runtime.Value
var once_Data_Array_NonEmpty_union_prime sync.Once
func Get_Data_Array_NonEmpty_union_prime() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_union_prime.Do(func() {
		cache_Data_Array_NonEmpty_union_prime = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_union_prime(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_Array_NonEmpty_union_prime
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
return Call_Data_Array_NonEmpty_union(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_Array_NonEmpty_union
}

var cache_Data_Array_NonEmpty_unzip gopurs_runtime.Value
var once_Data_Array_NonEmpty_unzip sync.Once
func Get_Data_Array_NonEmpty_unzip() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unzip.Do(func() {
		cache_Data_Array_NonEmpty_unzip = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Array_NonEmpty_unzip(func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_NonEmpty_unzip
}

var cache_Data_Array_NonEmpty_updateAt gopurs_runtime.Value
var once_Data_Array_NonEmpty_updateAt sync.Once
func Get_Data_Array_NonEmpty_updateAt() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_updateAt.Do(func() {
		cache_Data_Array_NonEmpty_updateAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_NonEmpty_updateAt(i_0_box.IntVal, x_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
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
return Call_Data_Array_NonEmpty_zipWithA(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box), f_1_box, func() []gopurs_runtime.Value {
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
return Call_Data_Array_NonEmpty_splitAt(i_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_NonEmpty_splitAt
}

var cache_Data_Array_NonEmpty_some gopurs_runtime.Value
var once_Data_Array_NonEmpty_some sync.Once
func Get_Data_Array_NonEmpty_some() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_some.Do(func() {
		cache_Data_Array_NonEmpty_some = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_some(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative](dictAlternative_0_box))
})
	})
	return cache_Data_Array_NonEmpty_some
}

var cache_Data_Array_NonEmpty_snoc_prime gopurs_runtime.Value
var once_Data_Array_NonEmpty_snoc_prime sync.Once
func Get_Data_Array_NonEmpty_snoc_prime() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_snoc_prime.Do(func() {
		cache_Data_Array_NonEmpty_snoc_prime = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_snoc_prime(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), x_1_box))
})
	})
	return cache_Data_Array_NonEmpty_snoc_prime
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
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_NonEmpty_modifyAt(i_0_box.IntVal, f_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_NonEmpty_modifyAt
}

var cache_Data_Array_NonEmpty_intersectBy_prime gopurs_runtime.Value
var once_Data_Array_NonEmpty_intersectBy_prime sync.Once
func Get_Data_Array_NonEmpty_intersectBy_prime() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_intersectBy_prime.Do(func() {
		cache_Data_Array_NonEmpty_intersectBy_prime = gopurs_runtime.Func2(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_intersectBy_prime(eq_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_NonEmpty_intersectBy_prime
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

var cache_Data_Array_NonEmpty_intersect_prime gopurs_runtime.Value
var once_Data_Array_NonEmpty_intersect_prime sync.Once
func Get_Data_Array_NonEmpty_intersect_prime() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_intersect_prime.Do(func() {
		cache_Data_Array_NonEmpty_intersect_prime = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_intersect_prime(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_Array_NonEmpty_intersect_prime
}

var cache_Data_Array_NonEmpty_intersect gopurs_runtime.Value
var once_Data_Array_NonEmpty_intersect sync.Once
func Get_Data_Array_NonEmpty_intersect() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_intersect.Do(func() {
		cache_Data_Array_NonEmpty_intersect = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_intersect(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_Array_NonEmpty_intersect
}

var cache_Data_Array_NonEmpty_intercalate gopurs_runtime.Value
var once_Data_Array_NonEmpty_intercalate sync.Once
func Get_Data_Array_NonEmpty_intercalate() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_intercalate.Do(func() {
		cache_Data_Array_NonEmpty_intercalate = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_intercalate(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dictSemigroup_0_box))
})
	})
	return cache_Data_Array_NonEmpty_intercalate
}

var cache_Data_Array_NonEmpty_insertAt gopurs_runtime.Value
var once_Data_Array_NonEmpty_insertAt sync.Once
func Get_Data_Array_NonEmpty_insertAt() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_insertAt.Do(func() {
		cache_Data_Array_NonEmpty_insertAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_NonEmpty_insertAt(i_0_box.IntVal, x_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_NonEmpty_insertAt
}

var cache_Data_Array_NonEmpty_fromFoldable1 gopurs_runtime.Value
var once_Data_Array_NonEmpty_fromFoldable1 sync.Once
func Get_Data_Array_NonEmpty_fromFoldable1() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_fromFoldable1.Do(func() {
		cache_Data_Array_NonEmpty_fromFoldable1 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_fromFoldable1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_0_box))
})
	})
	return cache_Data_Array_NonEmpty_fromFoldable1
}

var cache_Data_Array_NonEmpty_fromArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_fromArray sync.Once
func Get_Data_Array_NonEmpty_fromArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_fromArray.Do(func() {
		cache_Data_Array_NonEmpty_fromArray = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_NonEmpty_fromArray(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_NonEmpty_fromArray
}

var cache_Data_Array_NonEmpty_fromFoldable gopurs_runtime.Value
var once_Data_Array_NonEmpty_fromFoldable sync.Once
func Get_Data_Array_NonEmpty_fromFoldable() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_fromFoldable.Do(func() {
		cache_Data_Array_NonEmpty_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box))
})
	})
	return cache_Data_Array_NonEmpty_fromFoldable
}

var cache_Data_Array_NonEmpty_transpose_prime gopurs_runtime.Value
var once_Data_Array_NonEmpty_transpose_prime sync.Once
func Get_Data_Array_NonEmpty_transpose_prime() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_transpose_prime.Do(func() {
		cache_Data_Array_NonEmpty_transpose_prime = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_NonEmpty_transpose_prime(func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}() }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_NonEmpty_transpose_prime
}

var cache_Data_Array_NonEmpty_foldr1 gopurs_runtime.Value
var once_Data_Array_NonEmpty_foldr1 sync.Once
func Get_Data_Array_NonEmpty_foldr1() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_foldr1.Do(func() {
		cache_Data_Array_NonEmpty_foldr1 = gopurs_runtime.RecordGet(Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray(), "foldr1")
	})
	return cache_Data_Array_NonEmpty_foldr1
}

var cache_Data_Array_NonEmpty_foldl1 gopurs_runtime.Value
var once_Data_Array_NonEmpty_foldl1 sync.Once
func Get_Data_Array_NonEmpty_foldl1() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_foldl1.Do(func() {
		cache_Data_Array_NonEmpty_foldl1 = gopurs_runtime.RecordGet(Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray(), "foldl1")
	})
	return cache_Data_Array_NonEmpty_foldl1
}

var cache_Data_Array_NonEmpty_foldMap1 gopurs_runtime.Value
var once_Data_Array_NonEmpty_foldMap1 sync.Once
func Get_Data_Array_NonEmpty_foldMap1() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_foldMap1.Do(func() {
		cache_Data_Array_NonEmpty_foldMap1 = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_foldMap1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dictSemigroup_0_box))
})
	})
	return cache_Data_Array_NonEmpty_foldMap1
}

var cache_Data_Array_NonEmpty_fold1 gopurs_runtime.Value
var once_Data_Array_NonEmpty_fold1 sync.Once
func Get_Data_Array_NonEmpty_fold1() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_fold1.Do(func() {
		cache_Data_Array_NonEmpty_fold1 = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_fold1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dictSemigroup_0_box))
})
	})
	return cache_Data_Array_NonEmpty_fold1
}

var cache_Data_Array_NonEmpty_difference_prime gopurs_runtime.Value
var once_Data_Array_NonEmpty_difference_prime sync.Once
func Get_Data_Array_NonEmpty_difference_prime() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_difference_prime.Do(func() {
		cache_Data_Array_NonEmpty_difference_prime = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_difference_prime(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_Array_NonEmpty_difference_prime
}

var cache_Data_Array_NonEmpty_cons_prime gopurs_runtime.Value
var once_Data_Array_NonEmpty_cons_prime sync.Once
func Get_Data_Array_NonEmpty_cons_prime() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_cons_prime.Do(func() {
		cache_Data_Array_NonEmpty_cons_prime = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_cons_prime(x_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_cons_prime
}

var cache_Data_Array_NonEmpty_fromNonEmpty gopurs_runtime.Value
var once_Data_Array_NonEmpty_fromNonEmpty sync.Once
func Get_Data_Array_NonEmpty_fromNonEmpty() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_fromNonEmpty.Do(func() {
		cache_Data_Array_NonEmpty_fromNonEmpty = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_fromNonEmpty(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](v_0_box)))
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
// TAST (Let): __local_var_0_2 -> gopurs_runtime.Value
__local_var_0_2 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), Get_Data_Array_NonEmpty_toArray())
_ = __local_var_0_2
// TAST (Let): __local_var_0_1 -> gopurs_runtime.Value
__local_var_0_1 := gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(__local_var_0_2, x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
_ = __local_var_0_1
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
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
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_NonEmpty_alterAt(i_0_box.IntVal, f_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_NonEmpty_alterAt
}

var cache_Data_Array_NonEmpty_adaptMaybe gopurs_runtime.Value
var once_Data_Array_NonEmpty_adaptMaybe sync.Once
func Get_Data_Array_NonEmpty_adaptMaybe() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_adaptMaybe.Do(func() {
		cache_Data_Array_NonEmpty_adaptMaybe = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_adaptMaybe(f_0_box)
})
	})
	return cache_Data_Array_NonEmpty_adaptMaybe
}

var cache_Data_Array_NonEmpty_head gopurs_runtime.Value
var once_Data_Array_NonEmpty_head sync.Once
func Get_Data_Array_NonEmpty_head() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_head.Do(func() {
		cache_Data_Array_NonEmpty_head = Call_Data_Array_NonEmpty_adaptMaybe(Get_Data_Array_head())
	})
	return cache_Data_Array_NonEmpty_head
}

var cache_Data_Array_NonEmpty_init gopurs_runtime.Value
var once_Data_Array_NonEmpty_init sync.Once
func Get_Data_Array_NonEmpty_init() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_init.Do(func() {
		cache_Data_Array_NonEmpty_init = Call_Data_Array_NonEmpty_adaptMaybe(Get_Data_Array_init())
	})
	return cache_Data_Array_NonEmpty_init
}

var cache_Data_Array_NonEmpty_last gopurs_runtime.Value
var once_Data_Array_NonEmpty_last sync.Once
func Get_Data_Array_NonEmpty_last() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_last.Do(func() {
		cache_Data_Array_NonEmpty_last = Call_Data_Array_NonEmpty_adaptMaybe(Get_Data_Array_last())
	})
	return cache_Data_Array_NonEmpty_last
}

var cache_Data_Array_NonEmpty_tail gopurs_runtime.Value
var once_Data_Array_NonEmpty_tail sync.Once
func Get_Data_Array_NonEmpty_tail() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_tail.Do(func() {
		cache_Data_Array_NonEmpty_tail = Call_Data_Array_NonEmpty_adaptMaybe(Get_Data_Array_tail())
	})
	return cache_Data_Array_NonEmpty_tail
}

var cache_Data_Array_NonEmpty_uncons gopurs_runtime.Value
var once_Data_Array_NonEmpty_uncons sync.Once
func Get_Data_Array_NonEmpty_uncons() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_uncons.Do(func() {
		cache_Data_Array_NonEmpty_uncons = Call_Data_Array_NonEmpty_adaptMaybe(Get_Data_Array_uncons())
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
		cache_Data_Array_NonEmpty_unsnoc = Call_Data_Array_NonEmpty_adaptMaybe(Get_Data_Array_unsnoc())
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
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_catMaybes(func() []*Constructor_Data_Maybe_Just {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]*Constructor_Data_Maybe_Just, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v) }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_catMaybes
}

var cache_Data_Array_NonEmpty_delete gopurs_runtime.Value
var once_Data_Array_NonEmpty_delete sync.Once
func Get_Data_Array_NonEmpty_delete() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_delete.Do(func() {
		cache_Data_Array_NonEmpty_delete = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_delete(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), x_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_delete
}

var cache_Data_Array_NonEmpty_deleteAt gopurs_runtime.Value
var once_Data_Array_NonEmpty_deleteAt sync.Once
func Get_Data_Array_NonEmpty_deleteAt() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_deleteAt.Do(func() {
		cache_Data_Array_NonEmpty_deleteAt = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_NonEmpty_deleteAt(i_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
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
return Call_Data_Array_NonEmpty_difference(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), func() []gopurs_runtime.Value {
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
return gopurs_runtime.Bool(Call_Data_Array_NonEmpty_elem(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), x_1_box, func() []gopurs_runtime.Value {
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
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_NonEmpty_elemIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), x_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_NonEmpty_elemIndex
}

var cache_Data_Array_NonEmpty_elemLastIndex gopurs_runtime.Value
var once_Data_Array_NonEmpty_elemLastIndex sync.Once
func Get_Data_Array_NonEmpty_elemLastIndex() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_elemLastIndex.Do(func() {
		cache_Data_Array_NonEmpty_elemLastIndex = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_NonEmpty_elemLastIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), x_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
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
return Call_Data_Array_NonEmpty_filterA(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box), f_1_box)
})
	})
	return cache_Data_Array_NonEmpty_filterA
}

var cache_Data_Array_NonEmpty_find gopurs_runtime.Value
var once_Data_Array_NonEmpty_find sync.Once
func Get_Data_Array_NonEmpty_find() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_find.Do(func() {
		cache_Data_Array_NonEmpty_find = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_NonEmpty_find(p_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_NonEmpty_find
}

var cache_Data_Array_NonEmpty_findIndex gopurs_runtime.Value
var once_Data_Array_NonEmpty_findIndex sync.Once
func Get_Data_Array_NonEmpty_findIndex() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_findIndex.Do(func() {
		cache_Data_Array_NonEmpty_findIndex = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_NonEmpty_findIndex(p_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_NonEmpty_findIndex
}

var cache_Data_Array_NonEmpty_findLastIndex gopurs_runtime.Value
var once_Data_Array_NonEmpty_findLastIndex sync.Once
func Get_Data_Array_NonEmpty_findLastIndex() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_findLastIndex.Do(func() {
		cache_Data_Array_NonEmpty_findLastIndex = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_NonEmpty_findLastIndex(x_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_NonEmpty_findLastIndex
}

var cache_Data_Array_NonEmpty_findMap gopurs_runtime.Value
var once_Data_Array_NonEmpty_findMap sync.Once
func Get_Data_Array_NonEmpty_findMap() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_findMap.Do(func() {
		cache_Data_Array_NonEmpty_findMap = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_NonEmpty_findMap(p_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_NonEmpty_findMap
}

var cache_Data_Array_NonEmpty_foldM gopurs_runtime.Value
var once_Data_Array_NonEmpty_foldM sync.Once
func Get_Data_Array_NonEmpty_foldM() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_foldM.Do(func() {
		cache_Data_Array_NonEmpty_foldM = gopurs_runtime.Func3(func(dictMonad_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, acc_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_foldM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box), f_1_box, acc_2_box)
})
	})
	return cache_Data_Array_NonEmpty_foldM
}

var cache_Data_Array_NonEmpty_foldRecM gopurs_runtime.Value
var once_Data_Array_NonEmpty_foldRecM sync.Once
func Get_Data_Array_NonEmpty_foldRecM() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_foldRecM.Do(func() {
		cache_Data_Array_NonEmpty_foldRecM = gopurs_runtime.Func3(func(dictMonadRec_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, acc_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_foldRecM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box), f_1_box, acc_2_box)
})
	})
	return cache_Data_Array_NonEmpty_foldRecM
}

var cache_Data_Array_NonEmpty_index gopurs_runtime.Value
var once_Data_Array_NonEmpty_index sync.Once
func Get_Data_Array_NonEmpty_index() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_index.Do(func() {
		cache_Data_Array_NonEmpty_index = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_index(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
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
return gopurs_runtime.Bool(Call_Data_Array_NonEmpty_notElem(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), x_1_box, func() []gopurs_runtime.Value {
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
return Call_Data_Array_NonEmpty_partition(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
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
return Call_Data_Array_NonEmpty_span(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
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
return Call_Data_Array_NonEmpty_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dictUnfoldable_0_box), func() []gopurs_runtime.Value {
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
return Call_Data_Array_NonEmpty_group(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_Array_NonEmpty_group
}

var cache_Data_Array_NonEmpty_groupAllBy gopurs_runtime.Value
var once_Data_Array_NonEmpty_groupAllBy sync.Once
func Get_Data_Array_NonEmpty_groupAllBy() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_groupAllBy.Do(func() {
		cache_Data_Array_NonEmpty_groupAllBy = gopurs_runtime.Func(func(op_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_groupAllBy(op_0_box)
})
	})
	return cache_Data_Array_NonEmpty_groupAllBy
}

var cache_Data_Array_NonEmpty_groupAll gopurs_runtime.Value
var once_Data_Array_NonEmpty_groupAll sync.Once
func Get_Data_Array_NonEmpty_groupAll() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_groupAll.Do(func() {
		cache_Data_Array_NonEmpty_groupAll = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_groupAll(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
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
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_insert(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), x_1_box, func() []gopurs_runtime.Value {
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
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_modifyAtIndices(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), is_1_box, f_2_box, func() []gopurs_runtime.Value {
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
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_nub(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), func() []gopurs_runtime.Value {
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
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_nubEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), func() []gopurs_runtime.Value {
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
return Call_Data_Array_NonEmpty_sort(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
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
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_sortWith(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), f_1_box, func() []gopurs_runtime.Value {
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
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_updateAtIndices(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), pairs_1_box, func() []gopurs_runtime.Value {
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
		cache_Data_Array_NonEmpty_unsafeIndex = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_unsafeIndex(_dollar__unused_0_box, func() []gopurs_runtime.Value {
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
return Call_Data_Array_NonEmpty_toUnfoldable1(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](dictUnfoldable1_0_box), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_NonEmpty_toUnfoldable1
}

var cache_Data_Array_NonEmpty_adaptAny__4166803194 gopurs_runtime.Value
var once_Data_Array_NonEmpty_adaptAny__4166803194 sync.Once
func Get_Data_Array_NonEmpty_adaptAny__4166803194() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_adaptAny__4166803194.Do(func() {
		cache_Data_Array_NonEmpty_adaptAny__4166803194 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Array_NonEmpty_adaptAny__4166803194(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_adaptAny__4166803194
}

var cache_Data_Array_NonEmpty_adaptAny__2550017530 gopurs_runtime.Value
var once_Data_Array_NonEmpty_adaptAny__2550017530 sync.Once
func Get_Data_Array_NonEmpty_adaptAny__2550017530() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_adaptAny__2550017530.Do(func() {
		cache_Data_Array_NonEmpty_adaptAny__2550017530 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_adaptAny__2550017530(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_adaptAny__2550017530
}

var cache_Data_Array_NonEmpty_adaptAny__3782176954 gopurs_runtime.Value
var once_Data_Array_NonEmpty_adaptAny__3782176954 sync.Once
func Get_Data_Array_NonEmpty_adaptAny__3782176954() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_adaptAny__3782176954.Do(func() {
		cache_Data_Array_NonEmpty_adaptAny__3782176954 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_NonEmpty_adaptAny__3782176954(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_NonEmpty_adaptAny__3782176954
}

var cache_Data_Array_NonEmpty_adaptAny__2822562266 gopurs_runtime.Value
var once_Data_Array_NonEmpty_adaptAny__2822562266 sync.Once
func Get_Data_Array_NonEmpty_adaptAny__2822562266() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_adaptAny__2822562266.Do(func() {
		cache_Data_Array_NonEmpty_adaptAny__2822562266 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_NonEmpty_adaptAny__2822562266(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_NonEmpty_adaptAny__2822562266
}

var cache_Data_Array_NonEmpty_adaptAny__4223607898 gopurs_runtime.Value
var once_Data_Array_NonEmpty_adaptAny__4223607898 sync.Once
func Get_Data_Array_NonEmpty_adaptAny__4223607898() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_adaptAny__4223607898.Do(func() {
		cache_Data_Array_NonEmpty_adaptAny__4223607898 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_NonEmpty_adaptAny__4223607898(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_NonEmpty_adaptAny__4223607898
}

var cache_Data_Array_NonEmpty_adaptAny__3353605050 gopurs_runtime.Value
var once_Data_Array_NonEmpty_adaptAny__3353605050 sync.Once
func Get_Data_Array_NonEmpty_adaptAny__3353605050() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_adaptAny__3353605050.Do(func() {
		cache_Data_Array_NonEmpty_adaptAny__3353605050 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_adaptAny__3353605050(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_NonEmpty_adaptAny__3353605050
}

var cache_Data_Array_NonEmpty_adaptAny__2575348090 gopurs_runtime.Value
var once_Data_Array_NonEmpty_adaptAny__2575348090 sync.Once
func Get_Data_Array_NonEmpty_adaptAny__2575348090() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_adaptAny__2575348090.Do(func() {
		cache_Data_Array_NonEmpty_adaptAny__2575348090 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_adaptAny__2575348090(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_NonEmpty_adaptAny__2575348090
}

var cache_Data_Array_NonEmpty_adaptAny__31046170 gopurs_runtime.Value
var once_Data_Array_NonEmpty_adaptAny__31046170 sync.Once
func Get_Data_Array_NonEmpty_adaptAny__31046170() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_adaptAny__31046170.Do(func() {
		cache_Data_Array_NonEmpty_adaptAny__31046170 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_adaptAny__31046170(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_NonEmpty_adaptAny__31046170
}

var cache_Data_Array_NonEmpty_adaptAny__2003479482 gopurs_runtime.Value
var once_Data_Array_NonEmpty_adaptAny__2003479482 sync.Once
func Get_Data_Array_NonEmpty_adaptAny__2003479482() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_adaptAny__2003479482.Do(func() {
		cache_Data_Array_NonEmpty_adaptAny__2003479482 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_adaptAny__2003479482(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_NonEmpty_adaptAny__2003479482
}

var cache_Data_Array_NonEmpty_adaptAny__2668197852 gopurs_runtime.Value
var once_Data_Array_NonEmpty_adaptAny__2668197852 sync.Once
func Get_Data_Array_NonEmpty_adaptAny__2668197852() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_adaptAny__2668197852.Do(func() {
		cache_Data_Array_NonEmpty_adaptAny__2668197852 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Array_NonEmpty_adaptAny__2668197852(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_adaptAny__2668197852
}

var cache_Data_Array_NonEmpty_adaptAny__162062748 gopurs_runtime.Value
var once_Data_Array_NonEmpty_adaptAny__162062748 sync.Once
func Get_Data_Array_NonEmpty_adaptAny__162062748() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_adaptAny__162062748.Do(func() {
		cache_Data_Array_NonEmpty_adaptAny__162062748 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_adaptAny__162062748(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_NonEmpty_adaptAny__162062748
}

var cache_Data_Array_NonEmpty_adaptAny__4201103260 gopurs_runtime.Value
var once_Data_Array_NonEmpty_adaptAny__4201103260 sync.Once
func Get_Data_Array_NonEmpty_adaptAny__4201103260() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_adaptAny__4201103260.Do(func() {
		cache_Data_Array_NonEmpty_adaptAny__4201103260 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_adaptAny__4201103260(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_adaptAny__4201103260
}

var cache_Data_Array_NonEmpty_adaptAny__2840548764 gopurs_runtime.Value
var once_Data_Array_NonEmpty_adaptAny__2840548764 sync.Once
func Get_Data_Array_NonEmpty_adaptAny__2840548764() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_adaptAny__2840548764.Do(func() {
		cache_Data_Array_NonEmpty_adaptAny__2840548764 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_adaptAny__2840548764(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_NonEmpty_adaptAny__2840548764
}

var cache_Data_Array_NonEmpty_adaptAny__1291335164 gopurs_runtime.Value
var once_Data_Array_NonEmpty_adaptAny__1291335164 sync.Once
func Get_Data_Array_NonEmpty_adaptAny__1291335164() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_adaptAny__1291335164.Do(func() {
		cache_Data_Array_NonEmpty_adaptAny__1291335164 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_NonEmpty_adaptAny__1291335164(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_NonEmpty_adaptAny__1291335164
}

var cache_Data_Array_NonEmpty_adaptAny__120724188 gopurs_runtime.Value
var once_Data_Array_NonEmpty_adaptAny__120724188 sync.Once
func Get_Data_Array_NonEmpty_adaptAny__120724188() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_adaptAny__120724188.Do(func() {
		cache_Data_Array_NonEmpty_adaptAny__120724188 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_adaptAny__120724188(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_NonEmpty_adaptAny__120724188
}

var cache_Data_Array_NonEmpty_adaptAny__339533500 gopurs_runtime.Value
var once_Data_Array_NonEmpty_adaptAny__339533500 sync.Once
func Get_Data_Array_NonEmpty_adaptAny__339533500() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_adaptAny__339533500.Do(func() {
		cache_Data_Array_NonEmpty_adaptAny__339533500 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_adaptAny__339533500(f_0_box, func() []*Constructor_Data_Maybe_Just {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]*Constructor_Data_Maybe_Just, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v) }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_adaptAny__339533500
}

var cache_Data_Array_NonEmpty_adaptMaybe__3884694685 gopurs_runtime.Value
var once_Data_Array_NonEmpty_adaptMaybe__3884694685 sync.Once
func Get_Data_Array_NonEmpty_adaptMaybe__3884694685() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_adaptMaybe__3884694685.Do(func() {
		cache_Data_Array_NonEmpty_adaptMaybe__3884694685 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_adaptMaybe__3884694685(f_0_box)
})
	})
	return cache_Data_Array_NonEmpty_adaptMaybe__3884694685
}

var cache_Data_Array_NonEmpty_adaptMaybe__2747357853 gopurs_runtime.Value
var once_Data_Array_NonEmpty_adaptMaybe__2747357853 sync.Once
func Get_Data_Array_NonEmpty_adaptMaybe__2747357853() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_adaptMaybe__2747357853.Do(func() {
		cache_Data_Array_NonEmpty_adaptMaybe__2747357853 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_adaptMaybe__2747357853(f_0_box)
})
	})
	return cache_Data_Array_NonEmpty_adaptMaybe__2747357853
}

var cache_Data_Array_NonEmpty_adaptMaybe__3944647165 gopurs_runtime.Value
var once_Data_Array_NonEmpty_adaptMaybe__3944647165 sync.Once
func Get_Data_Array_NonEmpty_adaptMaybe__3944647165() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_adaptMaybe__3944647165.Do(func() {
		cache_Data_Array_NonEmpty_adaptMaybe__3944647165 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_adaptMaybe__3944647165(f_0_box)
})
	})
	return cache_Data_Array_NonEmpty_adaptMaybe__3944647165
}

var cache_Data_Array_NonEmpty_adaptMaybe__1892156733 gopurs_runtime.Value
var once_Data_Array_NonEmpty_adaptMaybe__1892156733 sync.Once
func Get_Data_Array_NonEmpty_adaptMaybe__1892156733() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_adaptMaybe__1892156733.Do(func() {
		cache_Data_Array_NonEmpty_adaptMaybe__1892156733 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_adaptMaybe__1892156733(f_0_box)
})
	})
	return cache_Data_Array_NonEmpty_adaptMaybe__1892156733
}

var cache_Data_Array_NonEmpty_cons_prime__4002752745 gopurs_runtime.Value
var once_Data_Array_NonEmpty_cons_prime__4002752745 sync.Once
func Get_Data_Array_NonEmpty_cons_prime__4002752745() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_cons_prime__4002752745.Do(func() {
		cache_Data_Array_NonEmpty_cons_prime__4002752745 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_cons_prime__4002752745(x_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_cons_prime__4002752745
}

var cache_Data_Array_NonEmpty_difference_prime__564981534 gopurs_runtime.Value
var once_Data_Array_NonEmpty_difference_prime__564981534 sync.Once
func Get_Data_Array_NonEmpty_difference_prime__564981534() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_difference_prime__564981534.Do(func() {
		cache_Data_Array_NonEmpty_difference_prime__564981534 = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_difference_prime__564981534(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_Array_NonEmpty_difference_prime__564981534
}

var cache_Data_Array_NonEmpty_fromArray__2195001498 gopurs_runtime.Value
var once_Data_Array_NonEmpty_fromArray__2195001498 sync.Once
func Get_Data_Array_NonEmpty_fromArray__2195001498() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_fromArray__2195001498.Do(func() {
		cache_Data_Array_NonEmpty_fromArray__2195001498 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_NonEmpty_fromArray__2195001498(func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_NonEmpty_fromArray__2195001498
}

var cache_Data_Array_NonEmpty_fromArray__260997498 gopurs_runtime.Value
var once_Data_Array_NonEmpty_fromArray__260997498 sync.Once
func Get_Data_Array_NonEmpty_fromArray__260997498() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_fromArray__260997498.Do(func() {
		cache_Data_Array_NonEmpty_fromArray__260997498 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_NonEmpty_fromArray__260997498(func() []string {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_NonEmpty_fromArray__260997498
}

var cache_Data_Array_NonEmpty_fromArray__7061562 gopurs_runtime.Value
var once_Data_Array_NonEmpty_fromArray__7061562 sync.Once
func Get_Data_Array_NonEmpty_fromArray__7061562() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_fromArray__7061562.Do(func() {
		cache_Data_Array_NonEmpty_fromArray__7061562 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_NonEmpty_fromArray__7061562(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_NonEmpty_fromArray__7061562
}

var cache_Data_Array_NonEmpty_fromArray__1294666874 gopurs_runtime.Value
var once_Data_Array_NonEmpty_fromArray__1294666874 sync.Once
func Get_Data_Array_NonEmpty_fromArray__1294666874() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_fromArray__1294666874.Do(func() {
		cache_Data_Array_NonEmpty_fromArray__1294666874 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_NonEmpty_fromArray__1294666874(func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}() }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_NonEmpty_fromArray__1294666874
}

var cache_Data_Array_NonEmpty_groupAllBy__3210946726 gopurs_runtime.Value
var once_Data_Array_NonEmpty_groupAllBy__3210946726 sync.Once
func Get_Data_Array_NonEmpty_groupAllBy__3210946726() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_groupAllBy__3210946726.Do(func() {
		cache_Data_Array_NonEmpty_groupAllBy__3210946726 = gopurs_runtime.Func(func(op_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_groupAllBy__3210946726(op_0_box)
})
	})
	return cache_Data_Array_NonEmpty_groupAllBy__3210946726
}

var cache_Data_Array_NonEmpty_intersectBy__145374773 gopurs_runtime.Value
var once_Data_Array_NonEmpty_intersectBy__145374773 sync.Once
func Get_Data_Array_NonEmpty_intersectBy__145374773() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_intersectBy__145374773.Do(func() {
		cache_Data_Array_NonEmpty_intersectBy__145374773 = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_intersectBy__145374773(eq_0_box, func() []gopurs_runtime.Value {
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
	return cache_Data_Array_NonEmpty_intersectBy__145374773
}

var cache_Data_Array_NonEmpty_intersectBy_prime__145374773 gopurs_runtime.Value
var once_Data_Array_NonEmpty_intersectBy_prime__145374773 sync.Once
func Get_Data_Array_NonEmpty_intersectBy_prime__145374773() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_intersectBy_prime__145374773.Do(func() {
		cache_Data_Array_NonEmpty_intersectBy_prime__145374773 = gopurs_runtime.Func2(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_intersectBy_prime__145374773(eq_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_NonEmpty_intersectBy_prime__145374773
}

var cache_Data_Array_NonEmpty_length__4151727363 gopurs_runtime.Value
var once_Data_Array_NonEmpty_length__4151727363 sync.Once
func Get_Data_Array_NonEmpty_length__4151727363() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_length__4151727363.Do(func() {
		cache_Data_Array_NonEmpty_length__4151727363 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Array_NonEmpty_length__4151727363(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_length__4151727363
}

var cache_Data_Array_NonEmpty_toArray__2781090619 gopurs_runtime.Value
var once_Data_Array_NonEmpty_toArray__2781090619 sync.Once
func Get_Data_Array_NonEmpty_toArray__2781090619() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_toArray__2781090619.Do(func() {
		cache_Data_Array_NonEmpty_toArray__2781090619 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Array_NonEmpty_toArray__2781090619(func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
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
	return cache_Data_Array_NonEmpty_toArray__2781090619
}

var cache_Data_Array_NonEmpty_toArray__1949224283 gopurs_runtime.Value
var once_Data_Array_NonEmpty_toArray__1949224283 sync.Once
func Get_Data_Array_NonEmpty_toArray__1949224283() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_toArray__1949224283.Do(func() {
		cache_Data_Array_NonEmpty_toArray__1949224283 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Array_NonEmpty_toArray__1949224283(func() []string {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_Array_NonEmpty_toArray__1949224283
}

var cache_Data_Array_NonEmpty_toArray__1136335355 gopurs_runtime.Value
var once_Data_Array_NonEmpty_toArray__1136335355 sync.Once
func Get_Data_Array_NonEmpty_toArray__1136335355() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_toArray__1136335355.Do(func() {
		cache_Data_Array_NonEmpty_toArray__1136335355 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_toArray__1136335355(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_toArray__1136335355
}

var cache_Data_Array_NonEmpty_toArray__4194748859 gopurs_runtime.Value
var once_Data_Array_NonEmpty_toArray__4194748859 sync.Once
func Get_Data_Array_NonEmpty_toArray__4194748859() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_toArray__4194748859.Do(func() {
		cache_Data_Array_NonEmpty_toArray__4194748859 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Array_NonEmpty_toArray__4194748859(func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}() }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_Array_NonEmpty_toArray__4194748859
}

var cache_Data_Array_NonEmpty_toArray__3370789435 gopurs_runtime.Value
var once_Data_Array_NonEmpty_toArray__3370789435 sync.Once
func Get_Data_Array_NonEmpty_toArray__3370789435() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_toArray__3370789435.Do(func() {
		cache_Data_Array_NonEmpty_toArray__3370789435 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Array_NonEmpty_toArray__3370789435(func() []*Constructor_Data_Maybe_Just {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
					unboxed := make([]*Constructor_Data_Maybe_Just, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v) }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_Array_NonEmpty_toArray__3370789435
}

var cache_Data_Array_NonEmpty_toArray__1068313307 gopurs_runtime.Value
var once_Data_Array_NonEmpty_toArray__1068313307 sync.Once
func Get_Data_Array_NonEmpty_toArray__1068313307() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_toArray__1068313307.Do(func() {
		cache_Data_Array_NonEmpty_toArray__1068313307 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Array_NonEmpty_toArray__1068313307(func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_Array_NonEmpty_toArray__1068313307
}

var cache_Data_Array_NonEmpty_toArray__4222275968 gopurs_runtime.Value
var once_Data_Array_NonEmpty_toArray__4222275968 sync.Once
func Get_Data_Array_NonEmpty_toArray__4222275968() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_toArray__4222275968.Do(func() {
		cache_Data_Array_NonEmpty_toArray__4222275968 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_toArray__4222275968(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_toArray__4222275968
}

var cache_Data_Array_NonEmpty_uncons__3248730276 gopurs_runtime.Value
var once_Data_Array_NonEmpty_uncons__3248730276 sync.Once
func Get_Data_Array_NonEmpty_uncons__3248730276() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_uncons__3248730276.Do(func() {
		cache_Data_Array_NonEmpty_uncons__3248730276 = Call_Data_Array_NonEmpty_adaptMaybe(Get_Data_Array_uncons())
	})
	return cache_Data_Array_NonEmpty_uncons__3248730276
}

var cache_Data_Array_NonEmpty_unionBy__145374773 gopurs_runtime.Value
var once_Data_Array_NonEmpty_unionBy__145374773 sync.Once
func Get_Data_Array_NonEmpty_unionBy__145374773() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unionBy__145374773.Do(func() {
		cache_Data_Array_NonEmpty_unionBy__145374773 = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_unionBy__145374773(eq_0_box, func() []gopurs_runtime.Value {
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
	return cache_Data_Array_NonEmpty_unionBy__145374773
}

var cache_Data_Array_NonEmpty_unionBy_prime__145374773 gopurs_runtime.Value
var once_Data_Array_NonEmpty_unionBy_prime__145374773 sync.Once
func Get_Data_Array_NonEmpty_unionBy_prime__145374773() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unionBy_prime__145374773.Do(func() {
		cache_Data_Array_NonEmpty_unionBy_prime__145374773 = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_unionBy_prime__145374773(eq_0_box, func() []gopurs_runtime.Value {
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
	return cache_Data_Array_NonEmpty_unionBy_prime__145374773
}

var cache_Data_Array_NonEmpty_unsafeAdapt__2550017530 gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeAdapt__2550017530 sync.Once
func Get_Data_Array_NonEmpty_unsafeAdapt__2550017530() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeAdapt__2550017530.Do(func() {
		cache_Data_Array_NonEmpty_unsafeAdapt__2550017530 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_unsafeAdapt__2550017530(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_unsafeAdapt__2550017530
}

var cache_Data_Array_NonEmpty_unsafeAdapt__2996396282 gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeAdapt__2996396282 sync.Once
func Get_Data_Array_NonEmpty_unsafeAdapt__2996396282() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeAdapt__2996396282.Do(func() {
		cache_Data_Array_NonEmpty_unsafeAdapt__2996396282 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_unsafeAdapt__2996396282(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_unsafeAdapt__2996396282
}

var cache_Data_Array_NonEmpty_unsafeAdapt__4201103260 gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeAdapt__4201103260 sync.Once
func Get_Data_Array_NonEmpty_unsafeAdapt__4201103260() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeAdapt__4201103260.Do(func() {
		cache_Data_Array_NonEmpty_unsafeAdapt__4201103260 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_unsafeAdapt__4201103260(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_unsafeAdapt__4201103260
}

var cache_Data_Array_NonEmpty_unsafeFromArray__2781090619 gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeFromArray__2781090619 sync.Once
func Get_Data_Array_NonEmpty_unsafeFromArray__2781090619() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeFromArray__2781090619.Do(func() {
		cache_Data_Array_NonEmpty_unsafeFromArray__2781090619 = Get_Data_Array_NonEmpty_Internal_NonEmptyArray()
	})
	return cache_Data_Array_NonEmpty_unsafeFromArray__2781090619
}

var cache_Data_Array_NonEmpty_unsafeFromArray__1136335355 gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeFromArray__1136335355 sync.Once
func Get_Data_Array_NonEmpty_unsafeFromArray__1136335355() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeFromArray__1136335355.Do(func() {
		cache_Data_Array_NonEmpty_unsafeFromArray__1136335355 = Get_Data_Array_NonEmpty_Internal_NonEmptyArray()
	})
	return cache_Data_Array_NonEmpty_unsafeFromArray__1136335355
}

var cache_Data_Array_NonEmpty_unsafeFromArray__1596955163 gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeFromArray__1596955163 sync.Once
func Get_Data_Array_NonEmpty_unsafeFromArray__1596955163() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeFromArray__1596955163.Do(func() {
		cache_Data_Array_NonEmpty_unsafeFromArray__1596955163 = Get_Data_Array_NonEmpty_Internal_NonEmptyArray()
	})
	return cache_Data_Array_NonEmpty_unsafeFromArray__1596955163
}

var cache_Data_Array_NonEmpty_unsafeFromArray__1068313307 gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeFromArray__1068313307 sync.Once
func Get_Data_Array_NonEmpty_unsafeFromArray__1068313307() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeFromArray__1068313307.Do(func() {
		cache_Data_Array_NonEmpty_unsafeFromArray__1068313307 = Get_Data_Array_NonEmpty_Internal_NonEmptyArray()
	})
	return cache_Data_Array_NonEmpty_unsafeFromArray__1068313307
}

var cache_Data_Array_NonEmpty_unsafeFromArray__3704033600 gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeFromArray__3704033600 sync.Once
func Get_Data_Array_NonEmpty_unsafeFromArray__3704033600() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeFromArray__3704033600.Do(func() {
		cache_Data_Array_NonEmpty_unsafeFromArray__3704033600 = Get_Data_Array_NonEmpty_Internal_NonEmptyArray()
	})
	return cache_Data_Array_NonEmpty_unsafeFromArray__3704033600
}

var cache_Data_Array_NonEmpty_unsafeFromArray__3238140064 gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeFromArray__3238140064 sync.Once
func Get_Data_Array_NonEmpty_unsafeFromArray__3238140064() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeFromArray__3238140064.Do(func() {
		cache_Data_Array_NonEmpty_unsafeFromArray__3238140064 = Get_Data_Array_NonEmpty_Internal_NonEmptyArray()
	})
	return cache_Data_Array_NonEmpty_unsafeFromArray__3238140064
}

var cache_Data_Array_NonEmpty_unsafeFromArray__4222275968 gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeFromArray__4222275968 sync.Once
func Get_Data_Array_NonEmpty_unsafeFromArray__4222275968() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeFromArray__4222275968.Do(func() {
		cache_Data_Array_NonEmpty_unsafeFromArray__4222275968 = Get_Data_Array_NonEmpty_Internal_NonEmptyArray()
	})
	return cache_Data_Array_NonEmpty_unsafeFromArray__4222275968
}

var cache_Data_Array_NonEmpty_unsafeFromArray__2747209664 gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeFromArray__2747209664 sync.Once
func Get_Data_Array_NonEmpty_unsafeFromArray__2747209664() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeFromArray__2747209664.Do(func() {
		cache_Data_Array_NonEmpty_unsafeFromArray__2747209664 = Get_Data_Array_NonEmpty_Internal_NonEmptyArray()
	})
	return cache_Data_Array_NonEmpty_unsafeFromArray__2747209664
}

var cache_Data_Array_NonEmpty_unsafeFromArrayF__3145969851 gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeFromArrayF__3145969851 sync.Once
func Get_Data_Array_NonEmpty_unsafeFromArrayF__3145969851() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeFromArrayF__3145969851.Do(func() {
		cache_Data_Array_NonEmpty_unsafeFromArrayF__3145969851 = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Array_NonEmpty_unsafeFromArrayF__3145969851
}

var cache_Data_Array_NonEmpty_unsafeFromArrayF__2527304315 gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeFromArrayF__2527304315 sync.Once
func Get_Data_Array_NonEmpty_unsafeFromArrayF__2527304315() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeFromArrayF__2527304315.Do(func() {
		cache_Data_Array_NonEmpty_unsafeFromArrayF__2527304315 = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Array_NonEmpty_unsafeFromArrayF__2527304315
}

var cache_Data_Array_NonEmpty_unsafeIndex__2808089623 gopurs_runtime.Value
var once_Data_Array_NonEmpty_unsafeIndex__2808089623 sync.Once
func Get_Data_Array_NonEmpty_unsafeIndex__2808089623() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_unsafeIndex__2808089623.Do(func() {
		cache_Data_Array_NonEmpty_unsafeIndex__2808089623 = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_unsafeIndex__2808089623(_dollar__unused_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), __local_var_2_box.IntVal)
})
	})
	return cache_Data_Array_NonEmpty_unsafeIndex__2808089623
}

func Call_Data_Array_NonEmpty_intercalate1(dictSemigroup_0_loop *Constructor_Data_Semigroup_Semigroup) gopurs_runtime.Value {
var dictSemigroup_0 *Constructor_Data_Semigroup_Semigroup = dictSemigroup_0_loop
_ = dictSemigroup_0
// TAST (Let): semigroupJoinWith1_1_1 -> *Constructor_Data_Semigroup_Semigroup
semigroupJoinWith1_1_1 := &Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(j_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_0.V0), gopurs_runtime.Apply(v_1, j_3), gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_0.V0), j_3, gopurs_runtime.Apply(v1_2, j_3)))
})
})
})}
_ = semigroupJoinWith1_1_1
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Func(func(j_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(foldable_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray(), "foldMap1"), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupJoinWith1_1_1)}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 -> gopurs_runtime.Value
__local_var_6_2 := gopurs_runtime.Apply(f_3, x_5)
_ = __local_var_6_2
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_6_2
})
}), foldable_4, j_2)
})
})
})
_ = __local_var_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_1_0, a_2, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_3
}))
})
}

func Call_Data_Array_NonEmpty_fold11(dictSemigroup_0_loop *Constructor_Data_Semigroup_Semigroup) gopurs_runtime.Value {
var dictSemigroup_0 *Constructor_Data_Semigroup_Semigroup = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray(), "foldMap1"), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(dictSemigroup_0)}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_Data_Array_NonEmpty_fromJust(v_0_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
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

func Call_Data_Array_NonEmpty_unionBy_prime(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(Get_Data_Array_unionBy(), eq_0, gopurs_runtime.Array(xs_1), gopurs_runtime.Array(x_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_union_prime(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_Array_NonEmpty_unionBy_prime(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_Array_NonEmpty_unionBy(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(Get_Data_Array_unionBy(), eq_0, gopurs_runtime.Array(xs_1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_union(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_Array_NonEmpty_unionBy(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_Array_NonEmpty_unzip(x_0_loop []*Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var x_0 []*Constructor_Data_Tuple_Tuple = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
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
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(__local_var_1_0.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(__local_var_1_0.UnsafePtr).V1})})
}

func Call_Data_Array_NonEmpty_updateAt(i_0_loop int64, x_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var i_0 int64 = i_0_loop
_ = i_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp5(Get_Data_Array__updateAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Int(i_0), x_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))
}

func Call_Data_Array_NonEmpty_zip(xs_0_loop []gopurs_runtime.Value, ys_1_loop []gopurs_runtime.Value) []*Constructor_Data_Tuple_Tuple {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
var ys_1 []gopurs_runtime.Value = ys_1_loop
_ = ys_1
return func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_zipWithImpl(), Get_Data_Tuple_Tuple(), gopurs_runtime.Array(xs_0), gopurs_runtime.Array(ys_1)).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
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
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_zipWithImpl(), f_0, gopurs_runtime.Array(xs_1), gopurs_runtime.Array(ys_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_zipWithA(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative, f_1_loop gopurs_runtime.Value, xs_2_loop []gopurs_runtime.Value, ys_3_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 []gopurs_runtime.Value = xs_2_loop
_ = xs_2
var ys_3 []gopurs_runtime.Value = ys_3_loop
_ = ys_3
return gopurs_runtime.Apply4(Get_Data_Array_zipWithA(), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, f_1, gopurs_runtime.Array(xs_2), gopurs_runtime.Array(ys_3))
}

func Call_Data_Array_NonEmpty_splitAt(i_0_loop int64, xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var i_0 int64 = i_0_loop
_ = i_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply2(Get_Data_Array_splitAt(), gopurs_runtime.Int(i_0), gopurs_runtime.Array(xs_1))
}

func Call_Data_Array_NonEmpty_some(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative) gopurs_runtime.Value {
var dictAlternative_0 *Constructor_Control_Alternative_Alternative = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): some1_1_0 -> gopurs_runtime.Value
some1_1_0 := gopurs_runtime.Apply(Get_Data_Array_some(), gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(dictAlternative_0)})
_ = some1_1_0
return gopurs_runtime.Func(func(dictLazy_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(some1_1_0, dictLazy_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, x_4)
})
})
}

func Call_Data_Array_NonEmpty_snoc_prime(xs_0_loop []gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) []gopurs_runtime.Value {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Array_snoc(), gopurs_runtime.Array(xs_0), x_1).UnsafePtr)
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
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Array_snoc(), gopurs_runtime.Array(xs_0), x_1).UnsafePtr)
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
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_Data_Array_replicateImpl(), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_Data_Ord_max__2538992856(), gopurs_runtime.Int(1), gopurs_runtime.Int(i_0)).IntVal), x_1).UnsafePtr)
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
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_Data_Array_rangeImpl(), gopurs_runtime.Int(x_0), gopurs_runtime.Int(y_1)).UnsafePtr)
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
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array(xs_0), gopurs_runtime.Array(ys_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_modifyAt(i_0_loop int64, f_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var i_0 int64 = i_0_loop
_ = i_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_Array_modifyAt(), gopurs_runtime.Int(i_0), f_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))
}

func Call_Data_Array_NonEmpty_intersectBy_prime(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
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
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(Get_Data_Array_intersectBy(), eq_0, gopurs_runtime.Array(xs_1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_intersect_prime(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_Array_NonEmpty_intersectBy_prime(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_Array_NonEmpty_intersect(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_Array_NonEmpty_intersectBy(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_Array_NonEmpty_intercalate(dictSemigroup_0_loop *Constructor_Data_Semigroup_Semigroup) gopurs_runtime.Value {
var dictSemigroup_0 *Constructor_Data_Semigroup_Semigroup = dictSemigroup_0_loop
_ = dictSemigroup_0
// TAST (Let): semigroupJoinWith1_1_1 -> *Constructor_Data_Semigroup_Semigroup
semigroupJoinWith1_1_1 := &Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(j_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_0.V0), gopurs_runtime.Apply(v_1, j_3), gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_0.V0), j_3, gopurs_runtime.Apply(v1_2, j_3)))
})
})
})}
_ = semigroupJoinWith1_1_1
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Func(func(j_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(foldable_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray(), "foldMap1"), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupJoinWith1_1_1)}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 -> gopurs_runtime.Value
__local_var_6_2 := gopurs_runtime.Apply(f_3, x_5)
_ = __local_var_6_2
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_6_2
})
}), foldable_4, j_2)
})
})
})
_ = __local_var_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_1_0, a_2, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_3
}))
})
}

func Call_Data_Array_NonEmpty_insertAt(i_0_loop int64, x_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var i_0 int64 = i_0_loop
_ = i_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp5(Get_Data_Array__insertAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Int(i_0), x_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))
}

func Call_Data_Array_NonEmpty_fromFoldable1(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dictFoldable1_0_loop
_ = dictFoldable1_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldable1_0.V0), gopurs_runtime.Value{}), "foldr")
_ = __local_var_1_1
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_fromFoldableImpl(), __local_var_1_1, __local_var_2)
})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
}

func Call_Data_Array_NonEmpty_fromArray(xs_0_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
var __t1 *Constructor_Data_Maybe_Just
{
var __t0 bool
{
if (gopurs_runtime.Int(int64(len(xs_0))).IntVal) > (0) {
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
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Array(xs_0)}
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

func Call_Data_Array_NonEmpty_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Box(dictFoldable_0.V2)
_ = __local_var_1_1
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_fromFoldableImpl(), __local_var_1_1, __local_var_2)
})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(__local_var_1_0, x_2)
_ = __local_var_3_2
var __t4 *Constructor_Data_Maybe_Just
{
var __t3 bool
{
if (gopurs_runtime.Int(int64(len(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_3_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))).IntVal) > (0) {
__t3 = true
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
if __t3 {
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_3_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())}
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
})
}

func Call_Data_Array_NonEmpty_transpose_prime(x_0_loop [][]gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var x_0 [][]gopurs_runtime.Value = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_Array_transpose(), func() gopurs_runtime.Value {
					arr := x_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}())
_ = __local_var_1_0
var __t2 *Constructor_Data_Maybe_Just
{
var __t1 bool
{
if (gopurs_runtime.Int(int64(len(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_0.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))).IntVal) > (0) {
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
__t2 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_0.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())}
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

func Call_Data_Array_NonEmpty_foldMap1(dictSemigroup_0_loop *Constructor_Data_Semigroup_Semigroup) gopurs_runtime.Value {
var dictSemigroup_0 *Constructor_Data_Semigroup_Semigroup = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray(), "foldMap1"), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(dictSemigroup_0)})
}

func Call_Data_Array_NonEmpty_fold1(dictSemigroup_0_loop *Constructor_Data_Semigroup_Semigroup) gopurs_runtime.Value {
var dictSemigroup_0 *Constructor_Data_Semigroup_Semigroup = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray(), "foldMap1"), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(dictSemigroup_0)}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_Data_Array_NonEmpty_difference_prime(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
// TAST (Let): difference1_1_0 -> gopurs_runtime.Value
difference1_1_0 := gopurs_runtime.Apply(Get_Data_Foldable_foldrArray(), gopurs_runtime.Apply(Get_Data_Array_deleteBy(), gopurs_runtime.Box(dictEq_0.V0)))
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

func Call_Data_Array_NonEmpty_cons_prime(x_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array([]gopurs_runtime.Value{x_0}), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_fromNonEmpty(v_0_loop *Constructor_Data_NonEmpty_NonEmpty) []gopurs_runtime.Value {
var v_0 *Constructor_Data_NonEmpty_NonEmpty = v_0_loop
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
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array(xs_0), gopurs_runtime.Array(ys_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_alterAt(i_0_loop int64, f_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var i_0 int64 = i_0_loop
_ = i_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_Array_alterAt(), gopurs_runtime.Int(i_0), f_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))
}

func Call_Data_Array_NonEmpty_adaptMaybe(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 930809136 && __local_var_3_0.UnsafePtr != nil) {
__t1 = (*Constructor_Data_Maybe_Just)(__local_var_3_0.UnsafePtr).V0
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
}))
}

func Call_Data_Array_NonEmpty_toNonEmpty(x_0_loop []gopurs_runtime.Value) *Constructor_Data_NonEmpty_NonEmpty {
var x_0 []gopurs_runtime.Value = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_Array_NonEmpty_uncons(), gopurs_runtime.Array(x_0))
_ = __local_var_1_0
return &Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.RecordGet(__local_var_1_0, "head"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(__local_var_1_0, "tail").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())}
}

func Call_Data_Array_NonEmpty_adaptAny(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
}

func Call_Data_Array_NonEmpty_all(p_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) bool {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return (gopurs_runtime.UncurriedApp2(Get_Data_Array_allImpl(), p_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).IntVal) != (0)
}

func Call_Data_Array_NonEmpty_any(p_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) bool {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return (gopurs_runtime.UncurriedApp2(Get_Data_Array_anyImpl(), p_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).IntVal) != (0)
}

func Call_Data_Array_NonEmpty_catMaybes(x_0_loop []*Constructor_Data_Maybe_Just) []gopurs_runtime.Value {
var x_0 []*Constructor_Data_Maybe_Just = x_0_loop
_ = x_0
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Array_mapMaybe(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
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

func Call_Data_Array_NonEmpty_delete(dictEq_0_loop *Constructor_Data_Eq_Eq, x_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(Get_Data_Array_deleteBy(), gopurs_runtime.Box(dictEq_0.V0), x_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_deleteAt(i_0_loop int64, x_1_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var i_0 int64 = i_0_loop
_ = i_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array__deleteAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Int(i_0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))
}

func Call_Data_Array_NonEmpty_deleteBy(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(Get_Data_Array_deleteBy(), f_0, x_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_difference(dictEq_0_loop *Constructor_Data_Eq_Eq, xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(Get_Data_Foldable_foldrArray(), gopurs_runtime.Apply(Get_Data_Array_deleteBy(), gopurs_runtime.Box(dictEq_0.V0)), gopurs_runtime.Array(xs_1))
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
var __t1 []gopurs_runtime.Value
{
var __t0 bool
{
if (i_0) < (1) {
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
__t1 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
goto end_branch_1
} else {

}
}
{
__t1 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(i_0), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))).IntVal), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
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
				}()
}
end_branch_1:
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(__t1).UnsafePtr)
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
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Array_dropEnd(), gopurs_runtime.Int(i_0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
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

func Call_Data_Array_NonEmpty_dropWhile(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_Data_Array_span(), f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())), "rest").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_elem(dictEq_0_loop *Constructor_Data_Eq_Eq, x_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) bool {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return (gopurs_runtime.Apply3(Get_Data_Array_elem(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_0)}, x_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).IntVal) != (0)
}

func Call_Data_Array_NonEmpty_elemIndex(dictEq_0_loop *Constructor_Data_Eq_Eq, x_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), v_3, x_1).IntVal) != (0))
}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))
}

func Call_Data_Array_NonEmpty_elemLastIndex(dictEq_0_loop *Constructor_Data_Eq_Eq, x_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_findLastIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), v_3, x_1).IntVal) != (0))
}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))
}

func Call_Data_Array_NonEmpty_filter(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
arr_val_filterImpl0 := gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
_ = arr_val_filterImpl0
_ = arr_val_filterImpl0
arr_go_filterImpl0 := (*[]gopurs_runtime.Value)(arr_val_filterImpl0.UnsafePtr)
_ = arr_go_filterImpl0
res_go_filterImpl0 := make([]gopurs_runtime.Value, 0)
_ = res_go_filterImpl0
for _, v_filterImpl0 := range *arr_go_filterImpl0 {
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

func Call_Data_Array_NonEmpty_filterA(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
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

func Call_Data_Array_NonEmpty_find(p_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(Get_Data_Array_find(), p_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))
}

func Call_Data_Array_NonEmpty_findIndex(p_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, p_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))
}

func Call_Data_Array_NonEmpty_findLastIndex(x_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_findLastIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, x_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))
}

func Call_Data_Array_NonEmpty_findMap(p_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_findMapImpl(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, Get_Data_Maybe_isJust(), p_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))
}

func Call_Data_Array_NonEmpty_foldM(dictMonad_0_loop *Constructor_Control_Monad_Monad, f_1_loop gopurs_runtime.Value, acc_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var acc_2 gopurs_runtime.Value = acc_2_loop
_ = acc_2
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
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

func Call_Data_Array_NonEmpty_foldRecM(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec, f_1_loop gopurs_runtime.Value, acc_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var acc_2 gopurs_runtime.Value = acc_2_loop
_ = acc_2
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
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

func Call_Data_Array_NonEmpty_index(x_0_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 []gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(Get_Data_Array_index(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_0).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
}

func Call_Data_Array_NonEmpty_length(x_0_loop []gopurs_runtime.Value) int64 {
var x_0 []gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_0).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))).IntVal
}

func Call_Data_Array_NonEmpty_mapMaybe(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Array_mapMaybe(), f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
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

func Call_Data_Array_NonEmpty_notElem(dictEq_0_loop *Constructor_Data_Eq_Eq, x_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) bool {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return (gopurs_runtime.Apply3(Get_Data_Array_notElem(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_0)}, x_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).IntVal) != (0)
}

func Call_Data_Array_NonEmpty_partition(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.UncurriedApp2(Get_Data_Array_partitionImpl(), f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
}

func Call_Data_Array_NonEmpty_slice(start_0_loop int64, end_1_loop int64, x_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var start_0 int64 = start_0_loop
_ = start_0
var end_1 int64 = end_1_loop
_ = end_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(start_0), gopurs_runtime.Int(end_1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_span(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply2(Get_Data_Array_span(), f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
}

func Call_Data_Array_NonEmpty_take(i_0_loop int64, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var i_0 int64 = i_0_loop
_ = i_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
var __t1 []gopurs_runtime.Value
{
var __t0 bool
{
if (i_0) < (1) {
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
__t1 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
goto end_branch_1
} else {

}
}
{
__t1 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int(i_0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
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
				}()
}
end_branch_1:
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(__t1).UnsafePtr)
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
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Array_takeEnd(), gopurs_runtime.Int(i_0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
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

func Call_Data_Array_NonEmpty_takeWhile(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_Data_Array_span(), f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())), "init").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable, x_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable = dictUnfoldable_0_loop
_ = dictUnfoldable_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply2(Get_Data_Array_toUnfoldable(), gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(dictUnfoldable_0)}, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
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

func Call_Data_Array_NonEmpty_group(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
// TAST (Let): eq_1_1 -> gopurs_runtime.Value
eq_1_1 := gopurs_runtime.Box(dictEq_0.V0)
_ = eq_1_1
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
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

func Call_Data_Array_NonEmpty_groupAllBy(op_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_Array_groupAllBy(), op_0)
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

func Call_Data_Array_NonEmpty_groupAll(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_Array_groupAllBy(), gopurs_runtime.Box(dictOrd_0.V1))
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

func Call_Data_Array_NonEmpty_insert(dictOrd_0_loop *Constructor_Data_Ord_Ord, x_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(Get_Data_Array_insertBy(), gopurs_runtime.Box(dictOrd_0.V1), x_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
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
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(Get_Data_Array_insertBy(), f_0, x_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
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
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
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

func Call_Data_Array_NonEmpty_modifyAtIndices(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, is_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value, x_3_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
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

func Call_Data_Array_NonEmpty_nub(dictOrd_0_loop *Constructor_Data_Ord_Ord, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Array_nubBy(), gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
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

func Call_Data_Array_NonEmpty_nubEq(dictEq_0_loop *Constructor_Data_Eq_Eq, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Array_nubByEq(), gopurs_runtime.Box(dictEq_0.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
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
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_scanlImpl(), f_0, x_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
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
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_scanrImpl(), f_0, x_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_sort(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_1 -> gopurs_runtime.Value
compare_1_1 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_1
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Array_sortBy(), compare_1_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
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

func Call_Data_Array_NonEmpty_sortWith(dictOrd_0_loop *Constructor_Data_Ord_Ord, f_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(Get_Data_Array_sortWith(), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(dictOrd_0)}, f_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_updateAtIndices(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, pairs_1_loop gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var pairs_1 gopurs_runtime.Value = pairs_1_loop
_ = pairs_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(Get_Data_Array_updateAtIndices(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_0)}, pairs_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_unsafeIndex(_dollar__unused_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value, __local_var_2_loop int64) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
var __local_var_2 int64 = __local_var_2_loop
_ = __local_var_2
return gopurs_runtime.ArrayAccess(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), int(__local_var_2))
}

func Call_Data_Array_NonEmpty_toUnfoldable1(dictUnfoldable1_0_loop *Constructor_Data_Unfoldable1_Unfoldable1, xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable1_0 *Constructor_Data_Unfoldable1_Unfoldable1 = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
// TAST (Let): len_2_0 -> int64
len_2_0 := gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(xs_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))).IntVal
_ = len_2_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictUnfoldable1_0.V0), gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Maybe_Just
{
var __t1 bool
{
if (i_3.IntVal) < ((len_2_0) - (1)) {
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
__t2 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((i_3.IntVal) + (1))}
goto end_branch_2
} else {

}
}
{
__t2 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply3(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.ArrayAccess(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_5.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), int(__local_var_6.IntVal))
})
})
}), gopurs_runtime.Array(xs_1), gopurs_runtime.Int(i_3.IntVal)), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}})}
}), gopurs_runtime.Int(0))
}

func Call_Data_Array_NonEmpty_adaptAny__4166803194(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) bool {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return (gopurs_runtime.Apply(f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).IntVal) != (0)
}

func Call_Data_Array_NonEmpty_adaptAny__2550017530(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
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

func Call_Data_Array_NonEmpty_adaptAny__3782176954(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))
}

func Call_Data_Array_NonEmpty_adaptAny__2822562266(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))
}

func Call_Data_Array_NonEmpty_adaptAny__4223607898(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))
}

func Call_Data_Array_NonEmpty_adaptAny__3353605050(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
}

func Call_Data_Array_NonEmpty_adaptAny__2575348090(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
}

func Call_Data_Array_NonEmpty_adaptAny__31046170(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
}

func Call_Data_Array_NonEmpty_adaptAny__2003479482(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
}

func Call_Data_Array_NonEmpty_adaptAny__2668197852(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) int64 {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).IntVal
}

func Call_Data_Array_NonEmpty_adaptAny__162062748(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
}

func Call_Data_Array_NonEmpty_adaptAny__4201103260(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
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

func Call_Data_Array_NonEmpty_adaptAny__2840548764(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
}

func Call_Data_Array_NonEmpty_adaptAny__1291335164(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))
}

func Call_Data_Array_NonEmpty_adaptAny__120724188(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
}

func Call_Data_Array_NonEmpty_adaptAny__339533500(f_0_loop gopurs_runtime.Value, x_1_loop []*Constructor_Data_Maybe_Just) []gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 []*Constructor_Data_Maybe_Just = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := x_1
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

func Call_Data_Array_NonEmpty_adaptMaybe__3884694685(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 930809136 && __local_var_3_0.UnsafePtr != nil) {
__t1 = (*Constructor_Data_Maybe_Just)(__local_var_3_0.UnsafePtr).V0
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
}))
}

func Call_Data_Array_NonEmpty_adaptMaybe__2747357853(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 930809136 && __local_var_3_0.UnsafePtr != nil) {
__t1 = (*Constructor_Data_Maybe_Just)(__local_var_3_0.UnsafePtr).V0
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
}))
}

func Call_Data_Array_NonEmpty_adaptMaybe__3944647165(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 930809136 && __local_var_3_0.UnsafePtr != nil) {
__t1 = (*Constructor_Data_Maybe_Just)(__local_var_3_0.UnsafePtr).V0
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
}))
}

func Call_Data_Array_NonEmpty_adaptMaybe__1892156733(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(f_0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (__local_var_3_0.Type == 9 && __local_var_3_0.IntVal == 930809136 && __local_var_3_0.UnsafePtr != nil) {
__t1 = (*Constructor_Data_Maybe_Just)(__local_var_3_0.UnsafePtr).V0
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
}))
}

func Call_Data_Array_NonEmpty_cons_prime__4002752745(x_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array([]gopurs_runtime.Value{x_0}), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_difference_prime__564981534(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
// TAST (Let): difference1_1_0 -> gopurs_runtime.Value
difference1_1_0 := gopurs_runtime.Apply(Get_Data_Foldable_foldrArray(), gopurs_runtime.Apply(Get_Data_Array_deleteBy(), gopurs_runtime.Box(dictEq_0.V0)))
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

func Call_Data_Array_NonEmpty_fromArray__2195001498(xs_0_loop []int64) *Constructor_Data_Maybe_Just {
var xs_0 []int64 = xs_0_loop
_ = xs_0
var __t1 *Constructor_Data_Maybe_Just
{
var __t0 bool
{
if (gopurs_runtime.Int(int64(len(xs_0))).IntVal) > (0) {
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
__t1 = &Constructor_Data_Maybe_Just{1, func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()}
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

func Call_Data_Array_NonEmpty_fromArray__260997498(xs_0_loop []string) *Constructor_Data_Maybe_Just {
var xs_0 []string = xs_0_loop
_ = xs_0
var __t1 *Constructor_Data_Maybe_Just
{
var __t0 bool
{
if (gopurs_runtime.Int(int64(len(xs_0))).IntVal) > (0) {
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
__t1 = &Constructor_Data_Maybe_Just{1, func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()}
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

func Call_Data_Array_NonEmpty_fromArray__7061562(xs_0_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
var __t1 *Constructor_Data_Maybe_Just
{
var __t0 bool
{
if (gopurs_runtime.Int(int64(len(xs_0))).IntVal) > (0) {
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
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Array(xs_0)}
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

func Call_Data_Array_NonEmpty_fromArray__1294666874(xs_0_loop [][]gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var xs_0 [][]gopurs_runtime.Value = xs_0_loop
_ = xs_0
var __t1 *Constructor_Data_Maybe_Just
{
var __t0 bool
{
if (gopurs_runtime.Int(int64(len(xs_0))).IntVal) > (0) {
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
__t1 = &Constructor_Data_Maybe_Just{1, func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}()}
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

func Call_Data_Array_NonEmpty_groupAllBy__3210946726(op_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_Array_groupAllBy(), op_0)
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

func Call_Data_Array_NonEmpty_intersectBy__145374773(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(Get_Data_Array_intersectBy(), eq_0, gopurs_runtime.Array(xs_1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_intersectBy_prime__145374773(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.Apply2(Get_Data_Array_intersectBy(), eq_0, gopurs_runtime.Array(xs_1))
}

func Call_Data_Array_NonEmpty_length__4151727363(x_0_loop []gopurs_runtime.Value) int64 {
var x_0 []gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_0).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))).IntVal
}

func Call_Data_Array_NonEmpty_toArray__2781090619(v_0_loop []int64) []int64 {
var v_0 []int64 = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Array_NonEmpty_toArray__1949224283(v_0_loop []string) []string {
var v_0 []string = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Array_NonEmpty_toArray__1136335355(v_0_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var v_0 []gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Array_NonEmpty_toArray__4194748859(v_0_loop [][]gopurs_runtime.Value) [][]gopurs_runtime.Value {
var v_0 [][]gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Array_NonEmpty_toArray__3370789435(v_0_loop []*Constructor_Data_Maybe_Just) []*Constructor_Data_Maybe_Just {
var v_0 []*Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Array_NonEmpty_toArray__1068313307(v_0_loop []*Constructor_Data_Tuple_Tuple) []*Constructor_Data_Tuple_Tuple {
var v_0 []*Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Array_NonEmpty_toArray__4222275968(v_0_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var v_0 []gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Array_NonEmpty_unionBy__145374773(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(Get_Data_Array_unionBy(), eq_0, gopurs_runtime.Array(xs_1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_2).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_unionBy_prime__145374773(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value, x_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var x_2 []gopurs_runtime.Value = x_2_loop
_ = x_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(Get_Data_Array_unionBy(), eq_0, gopurs_runtime.Array(xs_1), gopurs_runtime.Array(x_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_NonEmpty_unsafeAdapt__2550017530(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
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

func Call_Data_Array_NonEmpty_unsafeAdapt__2996396282(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
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

func Call_Data_Array_NonEmpty_unsafeAdapt__4201103260(f_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
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

func Call_Data_Array_NonEmpty_unsafeIndex__2808089623(_dollar__unused_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value, __local_var_2_loop int64) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
var __local_var_2 int64 = __local_var_2_loop
_ = __local_var_2
return gopurs_runtime.ArrayAccess(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), int(__local_var_2))
}


